# \ProvisioningProfileAPI

All URIs are relative to *https://api.bitrise.io/v0.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProvisioningProfileConfirm**](ProvisioningProfileAPI.md#ProvisioningProfileConfirm) | **Post** /apps/{app-slug}/provisioning-profiles/{provisioning-profile-slug}/uploaded | Confirm a provisioning profile upload
[**ProvisioningProfileCreate**](ProvisioningProfileAPI.md#ProvisioningProfileCreate) | **Post** /apps/{app-slug}/provisioning-profiles | Create a provisioning profile
[**ProvisioningProfileDelete**](ProvisioningProfileAPI.md#ProvisioningProfileDelete) | **Delete** /apps/{app-slug}/provisioning-profiles/{provisioning-profile-slug} | Delete a provisioning profile
[**ProvisioningProfileList**](ProvisioningProfileAPI.md#ProvisioningProfileList) | **Get** /apps/{app-slug}/provisioning-profiles | Get a list of the provisioning profiles
[**ProvisioningProfileShow**](ProvisioningProfileAPI.md#ProvisioningProfileShow) | **Get** /apps/{app-slug}/provisioning-profiles/{provisioning-profile-slug} | Get a specific provisioning profile
[**ProvisioningProfileUpdate**](ProvisioningProfileAPI.md#ProvisioningProfileUpdate) | **Patch** /apps/{app-slug}/provisioning-profiles/{provisioning-profile-slug} | Update a provisioning profile



## ProvisioningProfileConfirm

> V0ProvisionProfileResponseModel ProvisioningProfileConfirm(ctx, appSlug, provisioningProfileSlug).Execute()

Confirm a provisioning profile upload



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
    provisioningProfileSlug := "provisioningProfileSlug_example" // string | Provisioning profile slug

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvisioningProfileAPI.ProvisioningProfileConfirm(context.Background(), appSlug, provisioningProfileSlug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvisioningProfileAPI.ProvisioningProfileConfirm``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProvisioningProfileConfirm`: V0ProvisionProfileResponseModel
    fmt.Fprintf(os.Stdout, "Response from `ProvisioningProfileAPI.ProvisioningProfileConfirm`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appSlug** | **string** | App slug | 
**provisioningProfileSlug** | **string** | Provisioning profile slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiProvisioningProfileConfirmRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**V0ProvisionProfileResponseModel**](V0ProvisionProfileResponseModel.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProvisioningProfileCreate

> V0ProvisionProfileResponseModel ProvisioningProfileCreate(ctx, appSlug).ProvisioningProfile(provisioningProfile).Execute()

Create a provisioning profile



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
    provisioningProfile := *openapiclient.NewV0ProvisionProfileUploadParams("UploadFileName_example", int32(123)) // V0ProvisionProfileUploadParams | Provisioning profile parameters such as file name and file size

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvisioningProfileAPI.ProvisioningProfileCreate(context.Background(), appSlug).ProvisioningProfile(provisioningProfile).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvisioningProfileAPI.ProvisioningProfileCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProvisioningProfileCreate`: V0ProvisionProfileResponseModel
    fmt.Fprintf(os.Stdout, "Response from `ProvisioningProfileAPI.ProvisioningProfileCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appSlug** | **string** | App slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiProvisioningProfileCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **provisioningProfile** | [**V0ProvisionProfileUploadParams**](V0ProvisionProfileUploadParams.md) | Provisioning profile parameters such as file name and file size | 

### Return type

[**V0ProvisionProfileResponseModel**](V0ProvisionProfileResponseModel.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProvisioningProfileDelete

> V0ProvisionProfileResponseModel ProvisioningProfileDelete(ctx, appSlug, provisioningProfileSlug).Execute()

Delete a provisioning profile



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
    provisioningProfileSlug := "provisioningProfileSlug_example" // string | Provisioning profile slug

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvisioningProfileAPI.ProvisioningProfileDelete(context.Background(), appSlug, provisioningProfileSlug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvisioningProfileAPI.ProvisioningProfileDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProvisioningProfileDelete`: V0ProvisionProfileResponseModel
    fmt.Fprintf(os.Stdout, "Response from `ProvisioningProfileAPI.ProvisioningProfileDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appSlug** | **string** | App slug | 
**provisioningProfileSlug** | **string** | Provisioning profile slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiProvisioningProfileDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**V0ProvisionProfileResponseModel**](V0ProvisionProfileResponseModel.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProvisioningProfileList

> V0ProvisionProfileListResponseModel ProvisioningProfileList(ctx, appSlug).Next(next).Limit(limit).Execute()

Get a list of the provisioning profiles



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
    next := "next_example" // string | Slug of the first provisioning profile in the response (optional)
    limit := int32(56) // int32 | Max number of elements per page (default: 50) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvisioningProfileAPI.ProvisioningProfileList(context.Background(), appSlug).Next(next).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvisioningProfileAPI.ProvisioningProfileList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProvisioningProfileList`: V0ProvisionProfileListResponseModel
    fmt.Fprintf(os.Stdout, "Response from `ProvisioningProfileAPI.ProvisioningProfileList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appSlug** | **string** | App slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiProvisioningProfileListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **next** | **string** | Slug of the first provisioning profile in the response | 
 **limit** | **int32** | Max number of elements per page (default: 50) | 

### Return type

[**V0ProvisionProfileListResponseModel**](V0ProvisionProfileListResponseModel.md)

### Authorization

[AddonAuthToken](../README.md#AddonAuthToken), [PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProvisioningProfileShow

> V0ProvisionProfileResponseModel ProvisioningProfileShow(ctx, appSlug, provisioningProfileSlug).Execute()

Get a specific provisioning profile



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
    provisioningProfileSlug := "provisioningProfileSlug_example" // string | Provisioning profile slug

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvisioningProfileAPI.ProvisioningProfileShow(context.Background(), appSlug, provisioningProfileSlug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvisioningProfileAPI.ProvisioningProfileShow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProvisioningProfileShow`: V0ProvisionProfileResponseModel
    fmt.Fprintf(os.Stdout, "Response from `ProvisioningProfileAPI.ProvisioningProfileShow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appSlug** | **string** | App slug | 
**provisioningProfileSlug** | **string** | Provisioning profile slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiProvisioningProfileShowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**V0ProvisionProfileResponseModel**](V0ProvisionProfileResponseModel.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProvisioningProfileUpdate

> V0ProvisionProfileResponseModel ProvisioningProfileUpdate(ctx, appSlug, provisioningProfileSlug).ProvisioningProfile(provisioningProfile).Execute()

Update a provisioning profile



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
    provisioningProfileSlug := "provisioningProfileSlug_example" // string | Provisioning profile slug
    provisioningProfile := *openapiclient.NewV0ProvProfileDocumentUpdateParams() // V0ProvProfileDocumentUpdateParams | Provisioning profile parameters

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvisioningProfileAPI.ProvisioningProfileUpdate(context.Background(), appSlug, provisioningProfileSlug).ProvisioningProfile(provisioningProfile).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvisioningProfileAPI.ProvisioningProfileUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProvisioningProfileUpdate`: V0ProvisionProfileResponseModel
    fmt.Fprintf(os.Stdout, "Response from `ProvisioningProfileAPI.ProvisioningProfileUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appSlug** | **string** | App slug | 
**provisioningProfileSlug** | **string** | Provisioning profile slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiProvisioningProfileUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **provisioningProfile** | [**V0ProvProfileDocumentUpdateParams**](V0ProvProfileDocumentUpdateParams.md) | Provisioning profile parameters | 

### Return type

[**V0ProvisionProfileResponseModel**](V0ProvisionProfileResponseModel.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

