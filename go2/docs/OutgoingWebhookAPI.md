# \OutgoingWebhookAPI

All URIs are relative to *https://api.bitrise.io/v0.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OutgoingWebhookCreate**](OutgoingWebhookAPI.md#OutgoingWebhookCreate) | **Post** /apps/{app-slug}/outgoing-webhooks | Create an outgoing webhook for an app
[**OutgoingWebhookDelete**](OutgoingWebhookAPI.md#OutgoingWebhookDelete) | **Delete** /apps/{app-slug}/outgoing-webhooks/{app-webhook-slug} | Delete an outgoing webhook of an app
[**OutgoingWebhookList**](OutgoingWebhookAPI.md#OutgoingWebhookList) | **Get** /apps/{app-slug}/outgoing-webhooks | List the outgoing webhooks of an app
[**OutgoingWebhookUpdate**](OutgoingWebhookAPI.md#OutgoingWebhookUpdate) | **Put** /apps/{app-slug}/outgoing-webhooks/{app-webhook-slug} | Update an outgoing webhook of an app



## OutgoingWebhookCreate

> V0AppWebhookCreatedResponseModel OutgoingWebhookCreate(ctx, appSlug).AppWebhookCreateParams(appWebhookCreateParams).Execute()

Create an outgoing webhook for an app



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
    appWebhookCreateParams := *openapiclient.NewV0AppWebhookCreateParams([]string{"Events_example"}, "Secret_example", "Url_example") // V0AppWebhookCreateParams | App webhook creation params

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OutgoingWebhookAPI.OutgoingWebhookCreate(context.Background(), appSlug).AppWebhookCreateParams(appWebhookCreateParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OutgoingWebhookAPI.OutgoingWebhookCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OutgoingWebhookCreate`: V0AppWebhookCreatedResponseModel
    fmt.Fprintf(os.Stdout, "Response from `OutgoingWebhookAPI.OutgoingWebhookCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appSlug** | **string** | App slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiOutgoingWebhookCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appWebhookCreateParams** | [**V0AppWebhookCreateParams**](V0AppWebhookCreateParams.md) | App webhook creation params | 

### Return type

[**V0AppWebhookCreatedResponseModel**](V0AppWebhookCreatedResponseModel.md)

### Authorization

[AddonAuthToken](../README.md#AddonAuthToken), [PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OutgoingWebhookDelete

> V0AppWebhookDeletedResponseModel OutgoingWebhookDelete(ctx, appSlug, appWebhookSlug).Execute()

Delete an outgoing webhook of an app



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OutgoingWebhookAPI.OutgoingWebhookDelete(context.Background(), appSlug, appWebhookSlug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OutgoingWebhookAPI.OutgoingWebhookDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OutgoingWebhookDelete`: V0AppWebhookDeletedResponseModel
    fmt.Fprintf(os.Stdout, "Response from `OutgoingWebhookAPI.OutgoingWebhookDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appSlug** | **string** | App slug | 
**appWebhookSlug** | **string** | App webhook slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiOutgoingWebhookDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**V0AppWebhookDeletedResponseModel**](V0AppWebhookDeletedResponseModel.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OutgoingWebhookList

> V0AppWebhookListResponseModel OutgoingWebhookList(ctx, appSlug).Next(next).Limit(limit).Execute()

List the outgoing webhooks of an app



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
    next := "next_example" // string | Slug of the first webhook in the response (optional)
    limit := int32(56) // int32 | Max number of elements per page (default: 50) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OutgoingWebhookAPI.OutgoingWebhookList(context.Background(), appSlug).Next(next).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OutgoingWebhookAPI.OutgoingWebhookList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OutgoingWebhookList`: V0AppWebhookListResponseModel
    fmt.Fprintf(os.Stdout, "Response from `OutgoingWebhookAPI.OutgoingWebhookList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appSlug** | **string** | App slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiOutgoingWebhookListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **next** | **string** | Slug of the first webhook in the response | 
 **limit** | **int32** | Max number of elements per page (default: 50) | 

### Return type

[**V0AppWebhookListResponseModel**](V0AppWebhookListResponseModel.md)

### Authorization

[AddonAuthToken](../README.md#AddonAuthToken), [PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OutgoingWebhookUpdate

> V0AppWebhookResponseModel OutgoingWebhookUpdate(ctx, appSlug, appWebhookSlug).AppWebhookUpdateParams(appWebhookUpdateParams).Execute()

Update an outgoing webhook of an app



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
    appWebhookUpdateParams := *openapiclient.NewV0AppWebhookUpdateParams([]string{"Events_example"}, "Url_example") // V0AppWebhookUpdateParams | App webhook update params

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OutgoingWebhookAPI.OutgoingWebhookUpdate(context.Background(), appSlug, appWebhookSlug).AppWebhookUpdateParams(appWebhookUpdateParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OutgoingWebhookAPI.OutgoingWebhookUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OutgoingWebhookUpdate`: V0AppWebhookResponseModel
    fmt.Fprintf(os.Stdout, "Response from `OutgoingWebhookAPI.OutgoingWebhookUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appSlug** | **string** | App slug | 
**appWebhookSlug** | **string** | App webhook slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiOutgoingWebhookUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **appWebhookUpdateParams** | [**V0AppWebhookUpdateParams**](V0AppWebhookUpdateParams.md) | App webhook update params | 

### Return type

[**V0AppWebhookResponseModel**](V0AppWebhookResponseModel.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

