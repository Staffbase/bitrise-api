# \KeyValueCacheAPI

All URIs are relative to *https://api.bitrise.io/v0.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CacheItemDelete**](KeyValueCacheAPI.md#CacheItemDelete) | **Delete** /apps/{app-slug}/cache-items/{cache-item-id} | Deletes a key-value cache item
[**CacheItemDeleteAll**](KeyValueCacheAPI.md#CacheItemDeleteAll) | **Delete** /apps/{app-slug}/cache-items | Deletes all key-value cache items belonging to an app
[**CacheItemDownload**](KeyValueCacheAPI.md#CacheItemDownload) | **Get** /apps/{app-slug}/cache-items/{cache-item-id}/download | Gets the download URL of a key-value cache item
[**CacheList**](KeyValueCacheAPI.md#CacheList) | **Get** /apps/{app-slug}/cache-items | List the key-value cache items belonging to an app



## CacheItemDelete

> CacheItemDelete(ctx, appSlug, cacheItemId).Execute()

Deletes a key-value cache item



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
    cacheItemId := "cacheItemId_example" // string | The ID of the cache item to be deleted

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.KeyValueCacheAPI.CacheItemDelete(context.Background(), appSlug, cacheItemId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyValueCacheAPI.CacheItemDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appSlug** | **string** | App slug | 
**cacheItemId** | **string** | The ID of the cache item to be deleted | 

### Other Parameters

Other parameters are passed through a pointer to a apiCacheItemDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CacheItemDeleteAll

> CacheItemDeleteAll(ctx, appSlug).Execute()

Deletes all key-value cache items belonging to an app



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
    r, err := apiClient.KeyValueCacheAPI.CacheItemDeleteAll(context.Background(), appSlug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyValueCacheAPI.CacheItemDeleteAll``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appSlug** | **string** | App slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiCacheItemDeleteAllRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CacheItemDownload

> V0CacheItemDownloadResponseModel CacheItemDownload(ctx, appSlug, cacheItemId).Execute()

Gets the download URL of a key-value cache item



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
    cacheItemId := "cacheItemId_example" // string | The ID of the cache item to be downloaded

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KeyValueCacheAPI.CacheItemDownload(context.Background(), appSlug, cacheItemId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyValueCacheAPI.CacheItemDownload``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CacheItemDownload`: V0CacheItemDownloadResponseModel
    fmt.Fprintf(os.Stdout, "Response from `KeyValueCacheAPI.CacheItemDownload`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appSlug** | **string** | App slug | 
**cacheItemId** | **string** | The ID of the cache item to be downloaded | 

### Other Parameters

Other parameters are passed through a pointer to a apiCacheItemDownloadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**V0CacheItemDownloadResponseModel**](V0CacheItemDownloadResponseModel.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CacheList

> V0CacheItemListResponseModel CacheList(ctx, appSlug).Next(next).Limit(limit).Execute()

List the key-value cache items belonging to an app



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
    next := "next_example" // string | Getting cache items created before the given parameter (RFC3339 time format, default: now) (optional)
    limit := int32(56) // int32 | Max number of elements per page (default: 100) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KeyValueCacheAPI.CacheList(context.Background(), appSlug).Next(next).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyValueCacheAPI.CacheList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CacheList`: V0CacheItemListResponseModel
    fmt.Fprintf(os.Stdout, "Response from `KeyValueCacheAPI.CacheList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appSlug** | **string** | App slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiCacheListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **next** | **string** | Getting cache items created before the given parameter (RFC3339 time format, default: now) | 
 **limit** | **int32** | Max number of elements per page (default: 100) | 

### Return type

[**V0CacheItemListResponseModel**](V0CacheItemListResponseModel.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

