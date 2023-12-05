# \LocalBuildsAPI

All URIs are relative to *https://api.bitrise.io/v0.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LocalBuildList**](LocalBuildsAPI.md#LocalBuildList) | **Get** /apps/{app-slug}/local-builds | List local builds of an app



## LocalBuildList

> V0PipelineListResponseModel LocalBuildList(ctx, appSlug).After(after).Before(before).Limit(limit).Execute()

List local builds of an app



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
    after := "after_example" // string | List local builds run after a given date (RFC3339 time format) (optional)
    before := "before_example" // string | List local builds run before a given date (RFC3339 time format) - was called 'next' earlier (optional)
    limit := int32(56) // int32 | Max number of elements per page (default: 10) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LocalBuildsAPI.LocalBuildList(context.Background(), appSlug).After(after).Before(before).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LocalBuildsAPI.LocalBuildList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LocalBuildList`: V0PipelineListResponseModel
    fmt.Fprintf(os.Stdout, "Response from `LocalBuildsAPI.LocalBuildList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appSlug** | **string** | App slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiLocalBuildListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **after** | **string** | List local builds run after a given date (RFC3339 time format) | 
 **before** | **string** | List local builds run before a given date (RFC3339 time format) - was called &#39;next&#39; earlier | 
 **limit** | **int32** | Max number of elements per page (default: 10) | 

### Return type

[**V0PipelineListResponseModel**](V0PipelineListResponseModel.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

