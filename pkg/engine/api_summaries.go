/*
Nextlinux Engine API Server

This is the Nextlinux Engine API. Provides the primary external API for users of the service.

API version: 0.6.0
Contact: nurmi@nextlinux.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package engine

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
)

// Linger please
var (
	_ _context.Context
)

type SummariesApi interface {

	/*
	ListImagetags List all visible image digests and tags

	List all image tags visible to the user

	 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @return ApiListImagetagsRequest
	*/
	ListImagetags(ctx _context.Context) ApiListImagetagsRequest

	// ListImagetagsExecute executes the request
	//  @return []NextlinuxImageTagSummary
	ListImagetagsExecute(r ApiListImagetagsRequest) ([]NextlinuxImageTagSummary, *_nethttp.Response, error)
}

// SummariesApiService SummariesApi service
type SummariesApiService service

type ApiListImagetagsRequest struct {
	ctx _context.Context
	ApiService SummariesApi
	imageStatus *[]string
	xNextlinuxAccount *string
}

// Filter images in one or more states such as active, deleting. Defaults to active images only if unspecified
func (r ApiListImagetagsRequest) ImageStatus(imageStatus []string) ApiListImagetagsRequest {
	r.imageStatus = &imageStatus
	return r
}
// An account name to change the resource scope of the request to that account, if permissions allow (admin only)
func (r ApiListImagetagsRequest) XNextlinuxAccount(xNextlinuxAccount string) ApiListImagetagsRequest {
	r.xNextlinuxAccount = &xNextlinuxAccount
	return r
}

func (r ApiListImagetagsRequest) Execute() ([]NextlinuxImageTagSummary, *_nethttp.Response, error) {
	return r.ApiService.ListImagetagsExecute(r)
}

/*
ListImagetags List all visible image digests and tags

List all image tags visible to the user

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListImagetagsRequest
*/
func (a *SummariesApiService) ListImagetags(ctx _context.Context) ApiListImagetagsRequest {
	return ApiListImagetagsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []NextlinuxImageTagSummary
func (a *SummariesApiService) ListImagetagsExecute(r ApiListImagetagsRequest) ([]NextlinuxImageTagSummary, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []NextlinuxImageTagSummary
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SummariesApiService.ListImagetags")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/summaries/imagetags"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.imageStatus != nil {
		localVarQueryParams.Add("image_status", parameterToString(*r.imageStatus, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.xNextlinuxAccount != nil {
		localVarHeaderParams["x-nextlinux-account"] = parameterToString(*r.xNextlinuxAccount, "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ApiErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}