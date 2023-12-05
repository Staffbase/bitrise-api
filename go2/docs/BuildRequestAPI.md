# \BuildRequestAPI

All URIs are relative to *https://api.bitrise.io/v0.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BuildRequestList**](BuildRequestAPI.md#BuildRequestList) | **Get** /apps/{app-slug}/build-requests | List the open build requests for an app
[**BuildRequestUpdate**](BuildRequestAPI.md#BuildRequestUpdate) | **Patch** /apps/{app-slug}/build-requests/{build-request-slug} | Update a build request



## BuildRequestList

> V0BuildRequestListResponseModel BuildRequestList(ctx, appSlug).Execute()

List the open build requests for an app



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuildRequestAPI.BuildRequestList(context.Background(), appSlug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuildRequestAPI.BuildRequestList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BuildRequestList`: V0BuildRequestListResponseModel
    fmt.Fprintf(os.Stdout, "Response from `BuildRequestAPI.BuildRequestList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appSlug** | **string** | App slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiBuildRequestListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V0BuildRequestListResponseModel**](V0BuildRequestListResponseModel.md)

### Authorization

[AddonAuthToken](../README.md#AddonAuthToken), [PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BuildRequestUpdate

> V0BuildRequestUpdateResponseModel BuildRequestUpdate(ctx, appSlug, buildRequestSlug).BuildRequest(buildRequest).Execute()

Update a build request



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
    buildRequestSlug := "buildRequestSlug_example" // string | Build request slug
    buildRequest := *openapiclient.NewV0BuildRequestUpdateParams(false) // V0BuildRequestUpdateParams | Build request parameters

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuildRequestAPI.BuildRequestUpdate(context.Background(), appSlug, buildRequestSlug).BuildRequest(buildRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuildRequestAPI.BuildRequestUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BuildRequestUpdate`: V0BuildRequestUpdateResponseModel
    fmt.Fprintf(os.Stdout, "Response from `BuildRequestAPI.BuildRequestUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appSlug** | **string** | App slug | 
**buildRequestSlug** | **string** | Build request slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiBuildRequestUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **buildRequest** | [**V0BuildRequestUpdateParams**](V0BuildRequestUpdateParams.md) | Build request parameters | 

### Return type

[**V0BuildRequestUpdateResponseModel**](V0BuildRequestUpdateResponseModel.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

