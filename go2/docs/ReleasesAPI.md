# \ReleasesAPI

All URIs are relative to *https://api.bitrise.io/v0.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReleaseCreateAppStore**](ReleasesAPI.md#ReleaseCreateAppStore) | **Post** /apps/{app-slug}/releases/app-store | Create a new Apple App Store release for the app.
[**ReleaseCreateGooglePlay**](ReleasesAPI.md#ReleaseCreateGooglePlay) | **Post** /apps/{app-slug}/releases/google-play | Create a new Google Play Store release for the app.



## ReleaseCreateAppStore

> V0ReleaseCreateAppStoreRespModel ReleaseCreateAppStore(ctx, appSlug).App(app).Execute()

Create a new Apple App Store release for the app.



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
    app := *openapiclient.NewV0ReleaseCreateAppStoreParams("BundleId_example", "Name_example") // V0ReleaseCreateAppStoreParams | App Store release parameters

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReleasesAPI.ReleaseCreateAppStore(context.Background(), appSlug).App(app).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReleasesAPI.ReleaseCreateAppStore``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReleaseCreateAppStore`: V0ReleaseCreateAppStoreRespModel
    fmt.Fprintf(os.Stdout, "Response from `ReleasesAPI.ReleaseCreateAppStore`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appSlug** | **string** | App slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiReleaseCreateAppStoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **app** | [**V0ReleaseCreateAppStoreParams**](V0ReleaseCreateAppStoreParams.md) | App Store release parameters | 

### Return type

[**V0ReleaseCreateAppStoreRespModel**](V0ReleaseCreateAppStoreRespModel.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReleaseCreateGooglePlay

> V0ReleaseCreateGooglePlayRespModel ReleaseCreateGooglePlay(ctx, appSlug).App(app).Execute()

Create a new Google Play Store release for the app.



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
    app := *openapiclient.NewV0ReleaseCreateGooglePlayParams("Name_example", "PackageName_example") // V0ReleaseCreateGooglePlayParams | Google Play release parameters

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReleasesAPI.ReleaseCreateGooglePlay(context.Background(), appSlug).App(app).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReleasesAPI.ReleaseCreateGooglePlay``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReleaseCreateGooglePlay`: V0ReleaseCreateGooglePlayRespModel
    fmt.Fprintf(os.Stdout, "Response from `ReleasesAPI.ReleaseCreateGooglePlay`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appSlug** | **string** | App slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiReleaseCreateGooglePlayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **app** | [**V0ReleaseCreateGooglePlayParams**](V0ReleaseCreateGooglePlayParams.md) | Google Play release parameters | 

### Return type

[**V0ReleaseCreateGooglePlayRespModel**](V0ReleaseCreateGooglePlayRespModel.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

