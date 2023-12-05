# \ActivityAPI

All URIs are relative to *https://api.bitrise.io/v0.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivityList**](ActivityAPI.md#ActivityList) | **Get** /me/activities | Get list of Bitrise activity events



## ActivityList

> V0ActivityEventListResponseModel ActivityList(ctx).Next(next).Limit(limit).Execute()

Get list of Bitrise activity events



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
    next := "next_example" // string | Slug of the first activity event in the response (optional)
    limit := int32(56) // int32 | Max number of elements per page (default: 50) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActivityAPI.ActivityList(context.Background()).Next(next).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActivityAPI.ActivityList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActivityList`: V0ActivityEventListResponseModel
    fmt.Fprintf(os.Stdout, "Response from `ActivityAPI.ActivityList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiActivityListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **next** | **string** | Slug of the first activity event in the response | 
 **limit** | **int32** | Max number of elements per page (default: 50) | 

### Return type

[**V0ActivityEventListResponseModel**](V0ActivityEventListResponseModel.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

