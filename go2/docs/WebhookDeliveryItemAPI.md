# \WebhookDeliveryItemAPI

All URIs are relative to *https://api.bitrise.io/v0.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WebhookDeliveryItemList**](WebhookDeliveryItemAPI.md#WebhookDeliveryItemList) | **Get** /apps/{app-slug}/outgoing-webhooks/{app-webhook-slug}/delivery-items | List the webhook delivery items of an app
[**WebhookDeliveryItemRedeliver**](WebhookDeliveryItemAPI.md#WebhookDeliveryItemRedeliver) | **Post** /apps/{app-slug}/outgoing-webhooks/{app-webhook-slug}/delivery-items/{webhook-delivery-item-slug}/redeliver | Re-deliver the webhook delivery items of an app
[**WebhookDeliveryItemShow**](WebhookDeliveryItemAPI.md#WebhookDeliveryItemShow) | **Get** /apps/{app-slug}/outgoing-webhooks/{app-webhook-slug}/delivery-items/{webhook-delivery-item-slug} | Get a specific delivery item of a webhook



## WebhookDeliveryItemList

> V0WebhookDeliveryItemShowResponseModel WebhookDeliveryItemList(ctx, appSlug, appWebhookSlug).Next(next).Limit(limit).Execute()

List the webhook delivery items of an app



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/staffbase/bitrise-api"
)

func main() {
    appSlug := "appSlug_example" // string | App slug
    appWebhookSlug := "appWebhookSlug_example" // string | App webhook slug
    next := "next_example" // string | Slug of the first delivery item in the response (optional)
    limit := int32(56) // int32 | Max number of elements per page (default: 50) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhookDeliveryItemAPI.WebhookDeliveryItemList(context.Background(), appSlug, appWebhookSlug).Next(next).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhookDeliveryItemAPI.WebhookDeliveryItemList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WebhookDeliveryItemList`: V0WebhookDeliveryItemShowResponseModel
    fmt.Fprintf(os.Stdout, "Response from `WebhookDeliveryItemAPI.WebhookDeliveryItemList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appSlug** | **string** | App slug | 
**appWebhookSlug** | **string** | App webhook slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiWebhookDeliveryItemListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **next** | **string** | Slug of the first delivery item in the response | 
 **limit** | **int32** | Max number of elements per page (default: 50) | 

### Return type

[**V0WebhookDeliveryItemShowResponseModel**](V0WebhookDeliveryItemShowResponseModel.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhookDeliveryItemRedeliver

> ServiceStandardErrorRespModel WebhookDeliveryItemRedeliver(ctx, appSlug, appWebhookSlug, webhookDeliveryItemSlug).Execute()

Re-deliver the webhook delivery items of an app



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/staffbase/bitrise-api"
)

func main() {
    appSlug := "appSlug_example" // string | App slug
    appWebhookSlug := "appWebhookSlug_example" // string | App webhook slug
    webhookDeliveryItemSlug := "webhookDeliveryItemSlug_example" // string | Webhook delivery item slug

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhookDeliveryItemAPI.WebhookDeliveryItemRedeliver(context.Background(), appSlug, appWebhookSlug, webhookDeliveryItemSlug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhookDeliveryItemAPI.WebhookDeliveryItemRedeliver``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WebhookDeliveryItemRedeliver`: ServiceStandardErrorRespModel
    fmt.Fprintf(os.Stdout, "Response from `WebhookDeliveryItemAPI.WebhookDeliveryItemRedeliver`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appSlug** | **string** | App slug | 
**appWebhookSlug** | **string** | App webhook slug | 
**webhookDeliveryItemSlug** | **string** | Webhook delivery item slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiWebhookDeliveryItemRedeliverRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ServiceStandardErrorRespModel**](ServiceStandardErrorRespModel.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhookDeliveryItemShow

> V0WebhookDeliveryItemResponseModel WebhookDeliveryItemShow(ctx, appSlug, appWebhookSlug, webhookDeliveryItemSlug).Execute()

Get a specific delivery item of a webhook



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/staffbase/bitrise-api"
)

func main() {
    appSlug := "appSlug_example" // string | App slug
    appWebhookSlug := "appWebhookSlug_example" // string | App webhook slug
    webhookDeliveryItemSlug := "webhookDeliveryItemSlug_example" // string | Webhook delivery item slug

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhookDeliveryItemAPI.WebhookDeliveryItemShow(context.Background(), appSlug, appWebhookSlug, webhookDeliveryItemSlug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhookDeliveryItemAPI.WebhookDeliveryItemShow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WebhookDeliveryItemShow`: V0WebhookDeliveryItemResponseModel
    fmt.Fprintf(os.Stdout, "Response from `WebhookDeliveryItemAPI.WebhookDeliveryItemShow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appSlug** | **string** | App slug | 
**appWebhookSlug** | **string** | App webhook slug | 
**webhookDeliveryItemSlug** | **string** | Webhook delivery item slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiWebhookDeliveryItemShowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**V0WebhookDeliveryItemResponseModel**](V0WebhookDeliveryItemResponseModel.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

