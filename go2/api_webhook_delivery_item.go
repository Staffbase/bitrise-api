/*
Bitrise API

Official REST API for Bitrise.io

API version: 0.1
Contact: letsconnect@bitrise.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// WebhookDeliveryItemAPIService WebhookDeliveryItemAPI service
type WebhookDeliveryItemAPIService service

type ApiWebhookDeliveryItemListRequest struct {
	ctx context.Context
	ApiService *WebhookDeliveryItemAPIService
	appSlug string
	appWebhookSlug string
	next *string
	limit *int32
}

// Slug of the first delivery item in the response
func (r ApiWebhookDeliveryItemListRequest) Next(next string) ApiWebhookDeliveryItemListRequest {
	r.next = &next
	return r
}

// Max number of elements per page (default: 50)
func (r ApiWebhookDeliveryItemListRequest) Limit(limit int32) ApiWebhookDeliveryItemListRequest {
	r.limit = &limit
	return r
}

func (r ApiWebhookDeliveryItemListRequest) Execute() (*V0WebhookDeliveryItemShowResponseModel, *http.Response, error) {
	return r.ApiService.WebhookDeliveryItemListExecute(r)
}

/*
WebhookDeliveryItemList List the webhook delivery items of an app

List all the delivery items of an outgoing webhook of a Bitrise application

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param appSlug App slug
 @param appWebhookSlug App webhook slug
 @return ApiWebhookDeliveryItemListRequest
*/
func (a *WebhookDeliveryItemAPIService) WebhookDeliveryItemList(ctx context.Context, appSlug string, appWebhookSlug string) ApiWebhookDeliveryItemListRequest {
	return ApiWebhookDeliveryItemListRequest{
		ApiService: a,
		ctx: ctx,
		appSlug: appSlug,
		appWebhookSlug: appWebhookSlug,
	}
}

// Execute executes the request
//  @return V0WebhookDeliveryItemShowResponseModel
func (a *WebhookDeliveryItemAPIService) WebhookDeliveryItemListExecute(r ApiWebhookDeliveryItemListRequest) (*V0WebhookDeliveryItemShowResponseModel, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *V0WebhookDeliveryItemShowResponseModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhookDeliveryItemAPIService.WebhookDeliveryItemList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/apps/{app-slug}/outgoing-webhooks/{app-webhook-slug}/delivery-items"
	localVarPath = strings.Replace(localVarPath, "{"+"app-slug"+"}", url.PathEscape(parameterValueToString(r.appSlug, "appSlug")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"app-webhook-slug"+"}", url.PathEscape(parameterValueToString(r.appWebhookSlug, "appWebhookSlug")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.next != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "next", r.next, "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
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
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["PersonalAccessToken"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ServiceStandardErrorRespModel
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ServiceStandardErrorRespModel
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ServiceStandardErrorRespModel
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiWebhookDeliveryItemRedeliverRequest struct {
	ctx context.Context
	ApiService *WebhookDeliveryItemAPIService
	appSlug string
	appWebhookSlug string
	webhookDeliveryItemSlug string
}

func (r ApiWebhookDeliveryItemRedeliverRequest) Execute() (*ServiceStandardErrorRespModel, *http.Response, error) {
	return r.ApiService.WebhookDeliveryItemRedeliverExecute(r)
}

/*
WebhookDeliveryItemRedeliver Re-deliver the webhook delivery items of an app

Re-deliver the delivery item of a specified webhook of a Bitrise application

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param appSlug App slug
 @param appWebhookSlug App webhook slug
 @param webhookDeliveryItemSlug Webhook delivery item slug
 @return ApiWebhookDeliveryItemRedeliverRequest
*/
func (a *WebhookDeliveryItemAPIService) WebhookDeliveryItemRedeliver(ctx context.Context, appSlug string, appWebhookSlug string, webhookDeliveryItemSlug string) ApiWebhookDeliveryItemRedeliverRequest {
	return ApiWebhookDeliveryItemRedeliverRequest{
		ApiService: a,
		ctx: ctx,
		appSlug: appSlug,
		appWebhookSlug: appWebhookSlug,
		webhookDeliveryItemSlug: webhookDeliveryItemSlug,
	}
}

// Execute executes the request
//  @return ServiceStandardErrorRespModel
func (a *WebhookDeliveryItemAPIService) WebhookDeliveryItemRedeliverExecute(r ApiWebhookDeliveryItemRedeliverRequest) (*ServiceStandardErrorRespModel, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ServiceStandardErrorRespModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhookDeliveryItemAPIService.WebhookDeliveryItemRedeliver")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/apps/{app-slug}/outgoing-webhooks/{app-webhook-slug}/delivery-items/{webhook-delivery-item-slug}/redeliver"
	localVarPath = strings.Replace(localVarPath, "{"+"app-slug"+"}", url.PathEscape(parameterValueToString(r.appSlug, "appSlug")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"app-webhook-slug"+"}", url.PathEscape(parameterValueToString(r.appWebhookSlug, "appWebhookSlug")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"webhook-delivery-item-slug"+"}", url.PathEscape(parameterValueToString(r.webhookDeliveryItemSlug, "webhookDeliveryItemSlug")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["PersonalAccessToken"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ServiceStandardErrorRespModel
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ServiceStandardErrorRespModel
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ServiceStandardErrorRespModel
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiWebhookDeliveryItemShowRequest struct {
	ctx context.Context
	ApiService *WebhookDeliveryItemAPIService
	appSlug string
	appWebhookSlug string
	webhookDeliveryItemSlug string
}

func (r ApiWebhookDeliveryItemShowRequest) Execute() (*V0WebhookDeliveryItemResponseModel, *http.Response, error) {
	return r.ApiService.WebhookDeliveryItemShowExecute(r)
}

/*
WebhookDeliveryItemShow Get a specific delivery item of a webhook

Get the specified delivery item of an outgoing webhook of a Bitrise application

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param appSlug App slug
 @param appWebhookSlug App webhook slug
 @param webhookDeliveryItemSlug Webhook delivery item slug
 @return ApiWebhookDeliveryItemShowRequest
*/
func (a *WebhookDeliveryItemAPIService) WebhookDeliveryItemShow(ctx context.Context, appSlug string, appWebhookSlug string, webhookDeliveryItemSlug string) ApiWebhookDeliveryItemShowRequest {
	return ApiWebhookDeliveryItemShowRequest{
		ApiService: a,
		ctx: ctx,
		appSlug: appSlug,
		appWebhookSlug: appWebhookSlug,
		webhookDeliveryItemSlug: webhookDeliveryItemSlug,
	}
}

// Execute executes the request
//  @return V0WebhookDeliveryItemResponseModel
func (a *WebhookDeliveryItemAPIService) WebhookDeliveryItemShowExecute(r ApiWebhookDeliveryItemShowRequest) (*V0WebhookDeliveryItemResponseModel, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *V0WebhookDeliveryItemResponseModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhookDeliveryItemAPIService.WebhookDeliveryItemShow")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/apps/{app-slug}/outgoing-webhooks/{app-webhook-slug}/delivery-items/{webhook-delivery-item-slug}"
	localVarPath = strings.Replace(localVarPath, "{"+"app-slug"+"}", url.PathEscape(parameterValueToString(r.appSlug, "appSlug")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"app-webhook-slug"+"}", url.PathEscape(parameterValueToString(r.appWebhookSlug, "appWebhookSlug")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"webhook-delivery-item-slug"+"}", url.PathEscape(parameterValueToString(r.webhookDeliveryItemSlug, "webhookDeliveryItemSlug")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["PersonalAccessToken"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ServiceStandardErrorRespModel
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ServiceStandardErrorRespModel
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ServiceStandardErrorRespModel
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
