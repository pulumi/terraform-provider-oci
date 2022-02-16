// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package aivision

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"net/http"
	"strings"
)

// ListModelsRequest wrapper for the ListModels operation
type ListModelsRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// The ID of the project for which to list the objects.
	ProjectId *string `mandatory:"false" contributesTo:"query" name:"projectId"`

	// Filter to match models with the given lifecycleState.
	LifecycleState ModelLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Filter to find the model with the given identifier.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListModelsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ListModelsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListModelsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListModelsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListModelsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListModelsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListModelsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := mappingModelLifecycleStateEnum[string(request.LifecycleState)]; !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetModelLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := mappingListModelsSortOrderEnum[string(request.SortOrder)]; !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListModelsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := mappingListModelsSortByEnum[string(request.SortBy)]; !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListModelsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListModelsResponse wrapper for the ListModels operation
type ListModelsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ModelCollection instances
	ModelCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListModelsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListModelsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListModelsSortOrderEnum Enum with underlying type: string
type ListModelsSortOrderEnum string

// Set of constants representing the allowable values for ListModelsSortOrderEnum
const (
	ListModelsSortOrderAsc  ListModelsSortOrderEnum = "ASC"
	ListModelsSortOrderDesc ListModelsSortOrderEnum = "DESC"
)

var mappingListModelsSortOrderEnum = map[string]ListModelsSortOrderEnum{
	"ASC":  ListModelsSortOrderAsc,
	"DESC": ListModelsSortOrderDesc,
}

// GetListModelsSortOrderEnumValues Enumerates the set of values for ListModelsSortOrderEnum
func GetListModelsSortOrderEnumValues() []ListModelsSortOrderEnum {
	values := make([]ListModelsSortOrderEnum, 0)
	for _, v := range mappingListModelsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListModelsSortOrderEnumStringValues Enumerates the set of values in String for ListModelsSortOrderEnum
func GetListModelsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// ListModelsSortByEnum Enum with underlying type: string
type ListModelsSortByEnum string

// Set of constants representing the allowable values for ListModelsSortByEnum
const (
	ListModelsSortByTimecreated ListModelsSortByEnum = "timeCreated"
	ListModelsSortByDisplayname ListModelsSortByEnum = "displayName"
)

var mappingListModelsSortByEnum = map[string]ListModelsSortByEnum{
	"timeCreated": ListModelsSortByTimecreated,
	"displayName": ListModelsSortByDisplayname,
}

// GetListModelsSortByEnumValues Enumerates the set of values for ListModelsSortByEnum
func GetListModelsSortByEnumValues() []ListModelsSortByEnum {
	values := make([]ListModelsSortByEnum, 0)
	for _, v := range mappingListModelsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListModelsSortByEnumStringValues Enumerates the set of values in String for ListModelsSortByEnum
func GetListModelsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}
