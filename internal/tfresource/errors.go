// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package tfresource

import (
	"fmt"
	"log"
	"reflect"
	"regexp"
	"strings"
	"time"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"

	"terraform-provider-oci/internal/globalvar"
)

type errorTypeEnum string

var serviceErrorCheck = func(err error) (failure oci_common.ServiceErrorRichInfo, ok bool) {
	return oci_common.IsServiceErrorRichInfo(err)
}

const (
	ServiceError         errorTypeEnum = "ServiceError"
	TimeoutError         errorTypeEnum = "TimeoutError"
	UnexpectedStateError errorTypeEnum = "UnexpectedStateError"
	WorkRequestError     errorTypeEnum = "WorkRequestError"
)

type customError struct {
	TypeOfError   errorTypeEnum
	ErrorCode     int
	ErrorCodeName string
	Service       string
	Message       string
	OpcRequestID  string
	ResourceOCID  string
	OperationName string
	RequestTarget string
	Suggestion    string
	VersionError  string
	ResourceDocs  string
	SdkApiDocs    string
}

// Create new error format for Terraform output
func newCustomError(sync interface{}, err error) error {
	var tfError customError
	errorMessage := err.Error()

	// Service error
	if failure, isServiceError := serviceErrorCheck(err); isServiceError {
		tfError = customError{
			TypeOfError:   ServiceError,
			ErrorCode:     failure.GetHTTPStatusCode(),
			ErrorCodeName: failure.GetCode(),
			Message:       failure.GetMessage(),
			OpcRequestID:  failure.GetOpcRequestID(),
			OperationName: failure.GetOperationName(),
			RequestTarget: failure.GetRequestTarget(),
			Service:       getServiceName(sync),
			ResourceDocs:  getResourceDocsURL(sync),
			SdkApiDocs:    failure.GetOperationReferenceLink(),
		}
	} else if strings.Contains(errorMessage, "timeout while waiting for state") {
		// Timeout error
		tfError = customError{
			TypeOfError:   TimeoutError,
			ErrorCodeName: "Operation Timeout",
			Message:       errorMessage,
			Service:       getServiceName(sync),
		}
		// Unexpected state error
	} else if strings.Contains(errorMessage, "unexpected state") {
		tfError = customError{
			TypeOfError:   UnexpectedStateError,
			ErrorCodeName: "Unexpected LifeCycle state",
			Message:       errorMessage,
			Service:       getServiceName(sync),
			ResourceOCID:  getResourceOCID(sync),
		}
	} else if strings.Contains(errorMessage, "work request") {
		tfError = customError{
			TypeOfError:   WorkRequestError,
			ErrorCodeName: "Work Request error",
			Message:       errorMessage,
			Service:       getServiceName(sync),
			ResourceOCID:  getResourceOCID(sync),
		}
	} else {
		// Terraform error return as is
		return err
	}

	tfError.VersionError = GetVersionAndDateError()
	tfError.Suggestion = getSuggestionFromError(tfError)
	return tfError.Error()
}

func (tfE customError) Error() error {
	switch tfE.TypeOfError {
	case ServiceError:
		return fmt.Errorf("%d-%s, %s \n"+
			"Suggestion: %s\n"+
			"Documentation: %s \n"+
			"API Reference: %s \n"+
			"Request Target: %s \n"+
			"%s \n"+
			"Service: %s \n"+
			"Operation Name: %s \n"+
			"OPC request ID: %s \n",
			tfE.ErrorCode, tfE.ErrorCodeName, tfE.Message, tfE.Suggestion, tfE.ResourceDocs, tfE.SdkApiDocs, tfE.RequestTarget,
			tfE.VersionError, tfE.Service, tfE.OperationName, tfE.OpcRequestID)
	case TimeoutError:
		return fmt.Errorf("%s \n"+
			"%s \n"+
			"Service: %s \n"+
			"Error Message: %s \n"+
			"Suggestion: %s\n",
			tfE.ErrorCodeName, tfE.VersionError, tfE.Service, tfE.Message, tfE.Suggestion)
	case UnexpectedStateError:
		return fmt.Errorf("%s \n"+
			"%s \n"+
			"Service: %s \n"+
			"Error Message: %s \n"+
			"Resource OCID: %s \n"+
			"Suggestion: %s\n",
			tfE.ErrorCodeName, tfE.VersionError, tfE.Service, tfE.Message, tfE.ResourceOCID, tfE.Suggestion)
	case WorkRequestError:
		return fmt.Errorf("%s \n"+
			"%s \n"+
			"Service: %s \n"+
			"Error Message: %s \n"+
			"Resource OCID: %s \n"+
			"Suggestion: %s\n",
			tfE.ErrorCodeName, tfE.VersionError, tfE.Service, tfE.Message, tfE.ResourceOCID, tfE.Suggestion)
	default:
		return fmt.Errorf(tfE.Message)
	}
}

func handleMissingResourceError(sync ResourceVoider, err *error) {

	if err != nil {
		// patch till OCE service returns correct error response code for invalid auth token
		if strings.Contains((*err).Error(), "IDCS token validation has failed") {
			return
		}

		if strings.Contains((*err).Error(), "does not exist") ||
			strings.Contains((*err).Error(), " not present in ") ||
			strings.Contains((*err).Error(), "not found") ||
			(strings.Contains((*err).Error(), "Load balancer") && strings.Contains((*err).Error(), " has no ")) ||
			strings.Contains(strings.ToLower((*err).Error()), "status code: 404") { // status code: 404 is not enough because the load balancer error responses don't include it for some reason
			log.Println("[DEBUG] Object does not exist, voiding resource and nullifying error")
			if sync != nil {
				sync.VoidState()
			}
			*err = nil
		}
	}
}

func HandleError(sync interface{}, err error) error {
	if err != nil {
		tfError := newCustomError(sync, err)
		return tfError
	}
	return err
}

func getServiceName(sync interface{}) string {
	syncTypeName := reflect.TypeOf(sync).String()
	if strings.Contains(syncTypeName, "ResourceCrud") {
		serviceName := syncTypeName[strings.Index(syncTypeName, ".")+1 : strings.Index(syncTypeName, "ResourceCrud")]
		return removeDuplicate(serviceName)
	}
	if strings.Contains(syncTypeName, "DataSourcesCrud") {
		serviceName := syncTypeName[strings.Index(syncTypeName, ".")+1 : strings.Index(syncTypeName, "DataSourcesCrud")]
		return removeDuplicate(serviceName)
	}
	if strings.Contains(syncTypeName, "DataSourceCrud") {
		serviceName := syncTypeName[strings.Index(syncTypeName, ".")+1 : strings.Index(syncTypeName, "DataSourceCrud")]
		return removeDuplicate(serviceName)
	}
	log.Printf("[DEBUG] Can't get the service name for: %v", syncTypeName)
	return ""
}

// Return the Terraform document for the resource/datasource
func getResourceDocsURL(sync interface{}) string {
	baseURL := globalvar.TerraformDocumentLink
	var result = baseURL
	syncTypeName := reflect.TypeOf(sync).String()
	if strings.Contains(syncTypeName, "ResourceCrud") {
		result += "resources/"
		resourceName := syncTypeName[strings.Index(syncTypeName, ".")+1 : strings.Index(syncTypeName, "ResourceCrud")]
		result += toSnakeCase(resourceName)
		return result
	}
	if strings.Contains(syncTypeName, "DataSourcesCrud") {
		result += "data-sources/"
		datasourceName := syncTypeName[strings.Index(syncTypeName, ".")+1 : strings.Index(syncTypeName, "DataSourcesCrud")]
		result += toSnakeCase(datasourceName)
		return result
	}
	if strings.Contains(syncTypeName, "DataSourceCrud") {
		result += "data-sources/"
		datasourceName := syncTypeName[strings.Index(syncTypeName, ".")+1 : strings.Index(syncTypeName, "DataSourceCrud")]
		result += toSnakeCase(datasourceName)
		return result
	}
	log.Printf("[DEBUG] Can't get the resource name for: %v", syncTypeName)
	return ""
}

// CoreBootVolume -> core_boot_volume
func toSnakeCase(name string) string {
	var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")
	snake := matchFirstCap.ReplaceAllString(name, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

func removeDuplicate(name string) string {
	re := regexp.MustCompile(`[A-Z][^A-Z]*`)
	subMatchAll := re.FindAllString(name, -1)
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range subMatchAll {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return strings.Join(list, " ")
}

// Use to get OCID from refresh state only
func getResourceOCID(sync interface{}) string {
	defer func() {
		if r := recover(); r != nil {
			log.Println("[WARN] ID() function panic recovered!", r)
		}
	}()
	if syn, ok := sync.(StatefulResource); ok {
		return syn.ID()
	}
	return ""
}

func GetVersionAndDateError() string {
	return getVersionAndDateErrorImpl(globalvar.Version, globalvar.ReleaseDate)
}

func getVersionAndDateErrorImpl(version string, date string) string {
	result := fmt.Sprintf("Provider version: %s, released on %s. ", version, date)
	today := time.Now()
	releaseDate, _ := time.Parse("2006-01-02", date)
	days := today.Sub(releaseDate).Hours() / 24

	if days > 8 {
		versionOld := int(days / 7)
		result += fmt.Sprintf("This provider is %v Update(s) behind to current.", versionOld)
	}
	return result
}
