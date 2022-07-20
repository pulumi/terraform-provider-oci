// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// EmdwControlPlane API
//
// A description of the EmdwControlPlane API
//

package emwarehouse

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

//EmDataLakeClient a client for EmDataLake
type EmDataLakeClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewEmDataLakeClientWithConfigurationProvider Creates a new default EmDataLake client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewEmDataLakeClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client EmDataLakeClient, err error) {
	provider, err := auth.GetGenericConfigurationProvider(configProvider)
	if err != nil {
		return client, err
	}
	baseClient, e := common.NewClientWithConfig(provider)
	if e != nil {
		return client, e
	}
	return newEmDataLakeClientFromBaseClient(baseClient, provider)
}

// NewEmDataLakeClientWithOboToken Creates a new default EmDataLake client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//  as well as reading the region
func NewEmDataLakeClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client EmDataLakeClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newEmDataLakeClientFromBaseClient(baseClient, configProvider)
}

func newEmDataLakeClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client EmDataLakeClient, err error) {
	// EmDataLake service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("EmDataLake"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = EmDataLakeClient{BaseClient: baseClient}
	client.BasePath = "20180828"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *EmDataLakeClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("emwarehouse", "https://operationsinsights.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *EmDataLakeClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
	if ok, err := common.IsConfigurationProviderValid(configProvider); !ok {
		return err
	}

	// Error has been checked already
	region, _ := configProvider.Region()
	client.SetRegion(region)
	if client.Host == "" {
		return fmt.Errorf("Invalid region or Host. Endpoint cannot be constructed without endpointServiceName or serviceEndpointTemplate for a dotted region")
	}
	client.config = &configProvider
	return nil
}

// ConfigurationProvider the ConfigurationProvider used in this client, or null if none set
func (client *EmDataLakeClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// CancelWorkRequest Cancels the work request with the given ID.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/emwarehouse/CancelWorkRequest.go.html to see an example of how to use CancelWorkRequest API.
func (client EmDataLakeClient) CancelWorkRequest(ctx context.Context, request CancelWorkRequestRequest) (response CancelWorkRequestResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.cancelWorkRequest, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CancelWorkRequestResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CancelWorkRequestResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CancelWorkRequestResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CancelWorkRequestResponse")
	}
	return
}

// cancelWorkRequest implements the OCIOperation interface (enables retrying operations)
func (client EmDataLakeClient) cancelWorkRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/workRequests/{workRequestId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CancelWorkRequestResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "EmDataLake", "CancelWorkRequest", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeEmWarehouseCompartment Moves a EmWarehouse resource from one compartment identifier to another. When provided, If-Match is checked against ETag values of the resource.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/emwarehouse/ChangeEmWarehouseCompartment.go.html to see an example of how to use ChangeEmWarehouseCompartment API.
func (client EmDataLakeClient) ChangeEmWarehouseCompartment(ctx context.Context, request ChangeEmWarehouseCompartmentRequest) (response ChangeEmWarehouseCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.changeEmWarehouseCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeEmWarehouseCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeEmWarehouseCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeEmWarehouseCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeEmWarehouseCompartmentResponse")
	}
	return
}

// changeEmWarehouseCompartment implements the OCIOperation interface (enables retrying operations)
func (client EmDataLakeClient) changeEmWarehouseCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/emWarehouses/{emWarehouseId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeEmWarehouseCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "EmDataLake", "ChangeEmWarehouseCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateEmWarehouse Creates a new EmWarehouse.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/emwarehouse/CreateEmWarehouse.go.html to see an example of how to use CreateEmWarehouse API.
func (client EmDataLakeClient) CreateEmWarehouse(ctx context.Context, request CreateEmWarehouseRequest) (response CreateEmWarehouseResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createEmWarehouse, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateEmWarehouseResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateEmWarehouseResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateEmWarehouseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateEmWarehouseResponse")
	}
	return
}

// createEmWarehouse implements the OCIOperation interface (enables retrying operations)
func (client EmDataLakeClient) createEmWarehouse(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/emWarehouses", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateEmWarehouseResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "EmDataLake", "CreateEmWarehouse", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteEmWarehouse Deletes a EmWarehouse resource by identifier
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/emwarehouse/DeleteEmWarehouse.go.html to see an example of how to use DeleteEmWarehouse API.
func (client EmDataLakeClient) DeleteEmWarehouse(ctx context.Context, request DeleteEmWarehouseRequest) (response DeleteEmWarehouseResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteEmWarehouse, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteEmWarehouseResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteEmWarehouseResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteEmWarehouseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteEmWarehouseResponse")
	}
	return
}

// deleteEmWarehouse implements the OCIOperation interface (enables retrying operations)
func (client EmDataLakeClient) deleteEmWarehouse(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/emWarehouses/{emWarehouseId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteEmWarehouseResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "EmDataLake", "DeleteEmWarehouse", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetEmWarehouse Gets a EmWarehouse by identifier
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/emwarehouse/GetEmWarehouse.go.html to see an example of how to use GetEmWarehouse API.
func (client EmDataLakeClient) GetEmWarehouse(ctx context.Context, request GetEmWarehouseRequest) (response GetEmWarehouseResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getEmWarehouse, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetEmWarehouseResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetEmWarehouseResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetEmWarehouseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetEmWarehouseResponse")
	}
	return
}

// getEmWarehouse implements the OCIOperation interface (enables retrying operations)
func (client EmDataLakeClient) getEmWarehouse(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/emWarehouses/{emWarehouseId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetEmWarehouseResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "EmDataLake", "GetEmWarehouse", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetEmWarehouseResourceUsage Gets a EmWarehouseResourceUsage by identifier
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/emwarehouse/GetEmWarehouseResourceUsage.go.html to see an example of how to use GetEmWarehouseResourceUsage API.
func (client EmDataLakeClient) GetEmWarehouseResourceUsage(ctx context.Context, request GetEmWarehouseResourceUsageRequest) (response GetEmWarehouseResourceUsageResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getEmWarehouseResourceUsage, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetEmWarehouseResourceUsageResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetEmWarehouseResourceUsageResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetEmWarehouseResourceUsageResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetEmWarehouseResourceUsageResponse")
	}
	return
}

// getEmWarehouseResourceUsage implements the OCIOperation interface (enables retrying operations)
func (client EmDataLakeClient) getEmWarehouseResourceUsage(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/emWarehouses/{emWarehouseId}/resourceUsage", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetEmWarehouseResourceUsageResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "EmDataLake", "GetEmWarehouseResourceUsage", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetWorkRequest Gets details of the work request with the given ID.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/emwarehouse/GetWorkRequest.go.html to see an example of how to use GetWorkRequest API.
func (client EmDataLakeClient) GetWorkRequest(ctx context.Context, request GetWorkRequestRequest) (response GetWorkRequestResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getWorkRequest, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetWorkRequestResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetWorkRequestResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetWorkRequestResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetWorkRequestResponse")
	}
	return
}

// getWorkRequest implements the OCIOperation interface (enables retrying operations)
func (client EmDataLakeClient) getWorkRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workRequests/{workRequestId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetWorkRequestResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "EmDataLake", "GetWorkRequest", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListEmWarehouses Returns a list of EmWarehouses.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/emwarehouse/ListEmWarehouses.go.html to see an example of how to use ListEmWarehouses API.
func (client EmDataLakeClient) ListEmWarehouses(ctx context.Context, request ListEmWarehousesRequest) (response ListEmWarehousesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listEmWarehouses, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListEmWarehousesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListEmWarehousesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListEmWarehousesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListEmWarehousesResponse")
	}
	return
}

// listEmWarehouses implements the OCIOperation interface (enables retrying operations)
func (client EmDataLakeClient) listEmWarehouses(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/emWarehouses", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListEmWarehousesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "EmDataLake", "ListEmWarehouses", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListEtlRuns Gets a list of runs of an EmWarehouseResource by identifier
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/emwarehouse/ListEtlRuns.go.html to see an example of how to use ListEtlRuns API.
func (client EmDataLakeClient) ListEtlRuns(ctx context.Context, request ListEtlRunsRequest) (response ListEtlRunsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listEtlRuns, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListEtlRunsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListEtlRunsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListEtlRunsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListEtlRunsResponse")
	}
	return
}

// listEtlRuns implements the OCIOperation interface (enables retrying operations)
func (client EmDataLakeClient) listEtlRuns(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/emWarehouses/{emWarehouseId}/etlRuns", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListEtlRunsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "EmDataLake", "ListEtlRuns", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequestErrors Returns a (paginated) list of errors for the work request with the given ID.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/emwarehouse/ListWorkRequestErrors.go.html to see an example of how to use ListWorkRequestErrors API.
func (client EmDataLakeClient) ListWorkRequestErrors(ctx context.Context, request ListWorkRequestErrorsRequest) (response ListWorkRequestErrorsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listWorkRequestErrors, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListWorkRequestErrorsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListWorkRequestErrorsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListWorkRequestErrorsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListWorkRequestErrorsResponse")
	}
	return
}

// listWorkRequestErrors implements the OCIOperation interface (enables retrying operations)
func (client EmDataLakeClient) listWorkRequestErrors(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workRequests/{workRequestId}/errors", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListWorkRequestErrorsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "EmDataLake", "ListWorkRequestErrors", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequestLogs Returns a (paginated) list of logs for the work request with the given ID.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/emwarehouse/ListWorkRequestLogs.go.html to see an example of how to use ListWorkRequestLogs API.
func (client EmDataLakeClient) ListWorkRequestLogs(ctx context.Context, request ListWorkRequestLogsRequest) (response ListWorkRequestLogsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listWorkRequestLogs, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListWorkRequestLogsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListWorkRequestLogsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListWorkRequestLogsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListWorkRequestLogsResponse")
	}
	return
}

// listWorkRequestLogs implements the OCIOperation interface (enables retrying operations)
func (client EmDataLakeClient) listWorkRequestLogs(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workRequests/{workRequestId}/logs", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListWorkRequestLogsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "EmDataLake", "ListWorkRequestLogs", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequests Lists the work requests in a compartment.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/emwarehouse/ListWorkRequests.go.html to see an example of how to use ListWorkRequests API.
func (client EmDataLakeClient) ListWorkRequests(ctx context.Context, request ListWorkRequestsRequest) (response ListWorkRequestsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listWorkRequests, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListWorkRequestsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListWorkRequestsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListWorkRequestsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListWorkRequestsResponse")
	}
	return
}

// listWorkRequests implements the OCIOperation interface (enables retrying operations)
func (client EmDataLakeClient) listWorkRequests(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workRequests", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListWorkRequestsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "EmDataLake", "ListWorkRequests", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateEmWarehouse Updates the EmWarehouse
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/emwarehouse/UpdateEmWarehouse.go.html to see an example of how to use UpdateEmWarehouse API.
func (client EmDataLakeClient) UpdateEmWarehouse(ctx context.Context, request UpdateEmWarehouseRequest) (response UpdateEmWarehouseResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateEmWarehouse, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateEmWarehouseResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateEmWarehouseResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateEmWarehouseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateEmWarehouseResponse")
	}
	return
}

// updateEmWarehouse implements the OCIOperation interface (enables retrying operations)
func (client EmDataLakeClient) updateEmWarehouse(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/emWarehouses/{emWarehouseId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateEmWarehouseResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "EmDataLake", "UpdateEmWarehouse", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
