# \GenericProjectFileAPI

All URIs are relative to *https://api.bitrise.io/v0.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GenericProjectFileConfirm**](GenericProjectFileAPI.md#GenericProjectFileConfirm) | **Post** /apps/{app-slug}/generic-project-files/{generic-project-file-slug}/uploaded | Confirm a generic project file upload
[**GenericProjectFileDelete**](GenericProjectFileAPI.md#GenericProjectFileDelete) | **Delete** /apps/{app-slug}/generic-project-files/{generic-project-file-slug} | Delete a generic project file
[**GenericProjectFileList**](GenericProjectFileAPI.md#GenericProjectFileList) | **Get** /apps/{app-slug}/generic-project-files | Get a list of the generic project files
[**GenericProjectFileShow**](GenericProjectFileAPI.md#GenericProjectFileShow) | **Get** /apps/{app-slug}/generic-project-files/{generic-project-file-slug} | Get a specific generic project file
[**GenericProjectFileUpdate**](GenericProjectFileAPI.md#GenericProjectFileUpdate) | **Patch** /apps/{app-slug}/generic-project-files/{generic-project-file-slug} | Update a generic project file
[**GenericProjectFilesCreate**](GenericProjectFileAPI.md#GenericProjectFilesCreate) | **Post** /apps/{app-slug}/generic-project-files | Create a generic project file



## GenericProjectFileConfirm

> V0ProjectFileStorageResponseModel GenericProjectFileConfirm(ctx, appSlug, genericProjectFileSlug).Execute()

Confirm a generic project file upload



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
    genericProjectFileSlug := "genericProjectFileSlug_example" // string | Generic project file slug

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GenericProjectFileAPI.GenericProjectFileConfirm(context.Background(), appSlug, genericProjectFileSlug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GenericProjectFileAPI.GenericProjectFileConfirm``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenericProjectFileConfirm`: V0ProjectFileStorageResponseModel
    fmt.Fprintf(os.Stdout, "Response from `GenericProjectFileAPI.GenericProjectFileConfirm`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appSlug** | **string** | App slug | 
**genericProjectFileSlug** | **string** | Generic project file slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiGenericProjectFileConfirmRequest struct via the builder pattern


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


## GenericProjectFileDelete

> V0ProjectFileStorageResponseModel GenericProjectFileDelete(ctx, appSlug, genericProjectFileSlug).Execute()

Delete a generic project file



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
    genericProjectFileSlug := "genericProjectFileSlug_example" // string | Generic project file slug

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GenericProjectFileAPI.GenericProjectFileDelete(context.Background(), appSlug, genericProjectFileSlug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GenericProjectFileAPI.GenericProjectFileDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenericProjectFileDelete`: V0ProjectFileStorageResponseModel
    fmt.Fprintf(os.Stdout, "Response from `GenericProjectFileAPI.GenericProjectFileDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appSlug** | **string** | App slug | 
**genericProjectFileSlug** | **string** | Generic project file slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiGenericProjectFileDeleteRequest struct via the builder pattern


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


## GenericProjectFileList

> V0ProjectFileStorageListResponseModel GenericProjectFileList(ctx, appSlug).Next(next).Limit(limit).Execute()

Get a list of the generic project files



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
    next := "next_example" // string | Slug of the first generic project file in the response (optional)
    limit := int32(56) // int32 | Max number of build certificates per page is 50. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GenericProjectFileAPI.GenericProjectFileList(context.Background(), appSlug).Next(next).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GenericProjectFileAPI.GenericProjectFileList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenericProjectFileList`: V0ProjectFileStorageListResponseModel
    fmt.Fprintf(os.Stdout, "Response from `GenericProjectFileAPI.GenericProjectFileList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appSlug** | **string** | App slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiGenericProjectFileListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **next** | **string** | Slug of the first generic project file in the response | 
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


## GenericProjectFileShow

> V0ProjectFileStorageResponseModel GenericProjectFileShow(ctx, appSlug, genericProjectFileSlug).Execute()

Get a specific generic project file



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
    genericProjectFileSlug := "genericProjectFileSlug_example" // string | Generic project file slug

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GenericProjectFileAPI.GenericProjectFileShow(context.Background(), appSlug, genericProjectFileSlug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GenericProjectFileAPI.GenericProjectFileShow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenericProjectFileShow`: V0ProjectFileStorageResponseModel
    fmt.Fprintf(os.Stdout, "Response from `GenericProjectFileAPI.GenericProjectFileShow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appSlug** | **string** | App slug | 
**genericProjectFileSlug** | **string** | Generic project file slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiGenericProjectFileShowRequest struct via the builder pattern


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


## GenericProjectFileUpdate

> V0ProjectFileStorageResponseModel GenericProjectFileUpdate(ctx, appSlug, genericProjectFileSlug).GenericProjectFile(genericProjectFile).Execute()

Update a generic project file



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
    genericProjectFileSlug := "genericProjectFileSlug_example" // string | Generic project file slug
    genericProjectFile := *openapiclient.NewV0ProjectFileStorageDocumentUpdateParams() // V0ProjectFileStorageDocumentUpdateParams | Generic project file parameters

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GenericProjectFileAPI.GenericProjectFileUpdate(context.Background(), appSlug, genericProjectFileSlug).GenericProjectFile(genericProjectFile).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GenericProjectFileAPI.GenericProjectFileUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenericProjectFileUpdate`: V0ProjectFileStorageResponseModel
    fmt.Fprintf(os.Stdout, "Response from `GenericProjectFileAPI.GenericProjectFileUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appSlug** | **string** | App slug | 
**genericProjectFileSlug** | **string** | Generic project file slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiGenericProjectFileUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **genericProjectFile** | [**V0ProjectFileStorageDocumentUpdateParams**](V0ProjectFileStorageDocumentUpdateParams.md) | Generic project file parameters | 

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


## GenericProjectFilesCreate

> V0ProjectFileStorageResponseModel GenericProjectFilesCreate(ctx, appSlug).GenericProjectFile(genericProjectFile).Execute()

Create a generic project file



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
    genericProjectFile := *openapiclient.NewV0ProjectFileStorageUploadParams("UploadFileName_example", int32(123), "UserEnvKey_example") // V0ProjectFileStorageUploadParams | Generic project file parameters

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GenericProjectFileAPI.GenericProjectFilesCreate(context.Background(), appSlug).GenericProjectFile(genericProjectFile).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GenericProjectFileAPI.GenericProjectFilesCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenericProjectFilesCreate`: V0ProjectFileStorageResponseModel
    fmt.Fprintf(os.Stdout, "Response from `GenericProjectFileAPI.GenericProjectFilesCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appSlug** | **string** | App slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiGenericProjectFilesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **genericProjectFile** | [**V0ProjectFileStorageUploadParams**](V0ProjectFileStorageUploadParams.md) | Generic project file parameters | 

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

