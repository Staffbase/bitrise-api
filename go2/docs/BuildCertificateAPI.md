# \BuildCertificateAPI

All URIs are relative to *https://api.bitrise.io/v0.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BuildCertificateConfirm**](BuildCertificateAPI.md#BuildCertificateConfirm) | **Post** /apps/{app-slug}/build-certificates/{build-certificate-slug}/uploaded | Confirm a build certificate upload
[**BuildCertificateCreate**](BuildCertificateAPI.md#BuildCertificateCreate) | **Post** /apps/{app-slug}/build-certificates | Create a build certificate
[**BuildCertificateDelete**](BuildCertificateAPI.md#BuildCertificateDelete) | **Delete** /apps/{app-slug}/build-certificates/{build-certificate-slug} | Delete a build certificate
[**BuildCertificateList**](BuildCertificateAPI.md#BuildCertificateList) | **Get** /apps/{app-slug}/build-certificates | Get a list of the build certificates
[**BuildCertificateShow**](BuildCertificateAPI.md#BuildCertificateShow) | **Get** /apps/{app-slug}/build-certificates/{build-certificate-slug} | Get a specific build certificate
[**BuildCertificateUpdate**](BuildCertificateAPI.md#BuildCertificateUpdate) | **Patch** /apps/{app-slug}/build-certificates/{build-certificate-slug} | Update a build certificate



## BuildCertificateConfirm

> V0BuildCertificateResponseModel BuildCertificateConfirm(ctx, appSlug, buildCertificateSlug).Execute()

Confirm a build certificate upload



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
    buildCertificateSlug := "buildCertificateSlug_example" // string | Build certificate slug

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuildCertificateAPI.BuildCertificateConfirm(context.Background(), appSlug, buildCertificateSlug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuildCertificateAPI.BuildCertificateConfirm``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BuildCertificateConfirm`: V0BuildCertificateResponseModel
    fmt.Fprintf(os.Stdout, "Response from `BuildCertificateAPI.BuildCertificateConfirm`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appSlug** | **string** | App slug | 
**buildCertificateSlug** | **string** | Build certificate slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiBuildCertificateConfirmRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**V0BuildCertificateResponseModel**](V0BuildCertificateResponseModel.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BuildCertificateCreate

> V0BuildCertificateResponseModel BuildCertificateCreate(ctx, appSlug).BuildCertificate(buildCertificate).Execute()

Create a build certificate



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
    buildCertificate := *openapiclient.NewV0BuildCertificateUploadParams("UploadFileName_example", int32(123)) // V0BuildCertificateUploadParams | Build certificate parameters such as file name and its file size

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuildCertificateAPI.BuildCertificateCreate(context.Background(), appSlug).BuildCertificate(buildCertificate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuildCertificateAPI.BuildCertificateCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BuildCertificateCreate`: V0BuildCertificateResponseModel
    fmt.Fprintf(os.Stdout, "Response from `BuildCertificateAPI.BuildCertificateCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appSlug** | **string** | App slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiBuildCertificateCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **buildCertificate** | [**V0BuildCertificateUploadParams**](V0BuildCertificateUploadParams.md) | Build certificate parameters such as file name and its file size | 

### Return type

[**V0BuildCertificateResponseModel**](V0BuildCertificateResponseModel.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BuildCertificateDelete

> V0BuildCertificateResponseModel BuildCertificateDelete(ctx, appSlug, buildCertificateSlug).Execute()

Delete a build certificate



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
    buildCertificateSlug := "buildCertificateSlug_example" // string | Build certificate slug

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuildCertificateAPI.BuildCertificateDelete(context.Background(), appSlug, buildCertificateSlug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuildCertificateAPI.BuildCertificateDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BuildCertificateDelete`: V0BuildCertificateResponseModel
    fmt.Fprintf(os.Stdout, "Response from `BuildCertificateAPI.BuildCertificateDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appSlug** | **string** | App slug | 
**buildCertificateSlug** | **string** | Build certificate slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiBuildCertificateDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**V0BuildCertificateResponseModel**](V0BuildCertificateResponseModel.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BuildCertificateList

> V0BuildCertificateListResponseModel BuildCertificateList(ctx, appSlug).Next(next).Limit(limit).Execute()

Get a list of the build certificates



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
    next := "next_example" // string | Slug of the first build certificate in the response (optional)
    limit := int32(56) // int32 | Max number of build certificates per page is 50. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuildCertificateAPI.BuildCertificateList(context.Background(), appSlug).Next(next).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuildCertificateAPI.BuildCertificateList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BuildCertificateList`: V0BuildCertificateListResponseModel
    fmt.Fprintf(os.Stdout, "Response from `BuildCertificateAPI.BuildCertificateList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appSlug** | **string** | App slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiBuildCertificateListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **next** | **string** | Slug of the first build certificate in the response | 
 **limit** | **int32** | Max number of build certificates per page is 50. | 

### Return type

[**V0BuildCertificateListResponseModel**](V0BuildCertificateListResponseModel.md)

### Authorization

[AddonAuthToken](../README.md#AddonAuthToken), [PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BuildCertificateShow

> V0BuildCertificateResponseModel BuildCertificateShow(ctx, appSlug, buildCertificateSlug).Execute()

Get a specific build certificate



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
    buildCertificateSlug := "buildCertificateSlug_example" // string | Build certificate slug

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuildCertificateAPI.BuildCertificateShow(context.Background(), appSlug, buildCertificateSlug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuildCertificateAPI.BuildCertificateShow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BuildCertificateShow`: V0BuildCertificateResponseModel
    fmt.Fprintf(os.Stdout, "Response from `BuildCertificateAPI.BuildCertificateShow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appSlug** | **string** | App slug | 
**buildCertificateSlug** | **string** | Build certificate slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiBuildCertificateShowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**V0BuildCertificateResponseModel**](V0BuildCertificateResponseModel.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BuildCertificateUpdate

> V0BuildCertificateResponseModel BuildCertificateUpdate(ctx, appSlug, buildCertificateSlug).BuildCertificate(buildCertificate).Execute()

Update a build certificate



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
    buildCertificateSlug := "buildCertificateSlug_example" // string | Build certificate slug
    buildCertificate := *openapiclient.NewV0BuildCertificateUpdateParams() // V0BuildCertificateUpdateParams | Build certificate parameters

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuildCertificateAPI.BuildCertificateUpdate(context.Background(), appSlug, buildCertificateSlug).BuildCertificate(buildCertificate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuildCertificateAPI.BuildCertificateUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BuildCertificateUpdate`: V0BuildCertificateResponseModel
    fmt.Fprintf(os.Stdout, "Response from `BuildCertificateAPI.BuildCertificateUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appSlug** | **string** | App slug | 
**buildCertificateSlug** | **string** | Build certificate slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiBuildCertificateUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **buildCertificate** | [**V0BuildCertificateUpdateParams**](V0BuildCertificateUpdateParams.md) | Build certificate parameters | 

### Return type

[**V0BuildCertificateResponseModel**](V0BuildCertificateResponseModel.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

