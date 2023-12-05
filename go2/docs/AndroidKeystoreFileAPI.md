# \AndroidKeystoreFileAPI

All URIs are relative to *https://api.bitrise.io/v0.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AndroidKeystoreFileConfirm**](AndroidKeystoreFileAPI.md#AndroidKeystoreFileConfirm) | **Post** /apps/{app-slug}/android-keystore-files/{android-keystore-file-slug} | Confirm an android keystore file upload
[**AndroidKeystoreFileCreate**](AndroidKeystoreFileAPI.md#AndroidKeystoreFileCreate) | **Post** /apps/{app-slug}/android-keystore-files | Create an Android keystore file
[**AndroidKeystoreFileDelete**](AndroidKeystoreFileAPI.md#AndroidKeystoreFileDelete) | **Delete** /apps/{app-slug}/android-keystore-files/{android-keystore-file-slug} | Delete an android keystore file
[**AndroidKeystoreFileList**](AndroidKeystoreFileAPI.md#AndroidKeystoreFileList) | **Get** /apps/{app-slug}/android-keystore-files | Get a list of the android keystore files



## AndroidKeystoreFileConfirm

> V0ProjectFileStorageResponseModel AndroidKeystoreFileConfirm(ctx, appSlug, androidKeystoreFileSlug).Execute()

Confirm an android keystore file upload



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
    androidKeystoreFileSlug := "androidKeystoreFileSlug_example" // string | Android keystore file slug

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AndroidKeystoreFileAPI.AndroidKeystoreFileConfirm(context.Background(), appSlug, androidKeystoreFileSlug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AndroidKeystoreFileAPI.AndroidKeystoreFileConfirm``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AndroidKeystoreFileConfirm`: V0ProjectFileStorageResponseModel
    fmt.Fprintf(os.Stdout, "Response from `AndroidKeystoreFileAPI.AndroidKeystoreFileConfirm`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appSlug** | **string** | App slug | 
**androidKeystoreFileSlug** | **string** | Android keystore file slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiAndroidKeystoreFileConfirmRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**V0ProjectFileStorageResponseModel**](V0ProjectFileStorageResponseModel.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AndroidKeystoreFileCreate

> V0ProjectFileStorageResponseModel AndroidKeystoreFileCreate(ctx, appSlug).AndroidKeystoreFile(androidKeystoreFile).Execute()

Create an Android keystore file



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
    androidKeystoreFile := *openapiclient.NewV0AndroidKeystoreFileUploadParams("Alias_example", "Password_example", "PrivateKeyPassword_example", "UploadFileName_example", int32(123)) // V0AndroidKeystoreFileUploadParams | Android keystore file parameters

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AndroidKeystoreFileAPI.AndroidKeystoreFileCreate(context.Background(), appSlug).AndroidKeystoreFile(androidKeystoreFile).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AndroidKeystoreFileAPI.AndroidKeystoreFileCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AndroidKeystoreFileCreate`: V0ProjectFileStorageResponseModel
    fmt.Fprintf(os.Stdout, "Response from `AndroidKeystoreFileAPI.AndroidKeystoreFileCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appSlug** | **string** | App slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiAndroidKeystoreFileCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **androidKeystoreFile** | [**V0AndroidKeystoreFileUploadParams**](V0AndroidKeystoreFileUploadParams.md) | Android keystore file parameters | 

### Return type

[**V0ProjectFileStorageResponseModel**](V0ProjectFileStorageResponseModel.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AndroidKeystoreFileDelete

> V0ProjectFileStorageResponseModel AndroidKeystoreFileDelete(ctx, appSlug, androidKeystoreFileSlug).Execute()

Delete an android keystore file



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
    androidKeystoreFileSlug := "androidKeystoreFileSlug_example" // string | Android keystore file slug

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AndroidKeystoreFileAPI.AndroidKeystoreFileDelete(context.Background(), appSlug, androidKeystoreFileSlug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AndroidKeystoreFileAPI.AndroidKeystoreFileDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AndroidKeystoreFileDelete`: V0ProjectFileStorageResponseModel
    fmt.Fprintf(os.Stdout, "Response from `AndroidKeystoreFileAPI.AndroidKeystoreFileDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appSlug** | **string** | App slug | 
**androidKeystoreFileSlug** | **string** | Android keystore file slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiAndroidKeystoreFileDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**V0ProjectFileStorageResponseModel**](V0ProjectFileStorageResponseModel.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AndroidKeystoreFileList

> V0ProjectFileStorageListResponseModel AndroidKeystoreFileList(ctx, appSlug).Next(next).Limit(limit).Execute()

Get a list of the android keystore files



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
    next := "next_example" // string | Slug of the first android keystore file in the response (optional)
    limit := int32(56) // int32 | Max number of build certificates per page is 50. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AndroidKeystoreFileAPI.AndroidKeystoreFileList(context.Background(), appSlug).Next(next).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AndroidKeystoreFileAPI.AndroidKeystoreFileList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AndroidKeystoreFileList`: V0ProjectFileStorageListResponseModel
    fmt.Fprintf(os.Stdout, "Response from `AndroidKeystoreFileAPI.AndroidKeystoreFileList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appSlug** | **string** | App slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiAndroidKeystoreFileListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **next** | **string** | Slug of the first android keystore file in the response | 
 **limit** | **int32** | Max number of build certificates per page is 50. | 

### Return type

[**V0ProjectFileStorageListResponseModel**](V0ProjectFileStorageListResponseModel.md)

### Authorization

[AddonAuthToken](../README.md#AddonAuthToken), [PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

