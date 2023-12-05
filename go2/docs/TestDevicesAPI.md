# \TestDevicesAPI

All URIs are relative to *https://api.bitrise.io/v0.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TestDeviceList**](TestDevicesAPI.md#TestDeviceList) | **Get** /apps/{app-slug}/test-devices | List the test devices for an app



## TestDeviceList

> V0TestDeviceListResponseModel TestDeviceList(ctx, appSlug).Execute()

List the test devices for an app



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
    resp, r, err := apiClient.TestDevicesAPI.TestDeviceList(context.Background(), appSlug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TestDevicesAPI.TestDeviceList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TestDeviceList`: V0TestDeviceListResponseModel
    fmt.Fprintf(os.Stdout, "Response from `TestDevicesAPI.TestDeviceList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appSlug** | **string** | App slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestDeviceListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V0TestDeviceListResponseModel**](V0TestDeviceListResponseModel.md)

### Authorization

[AddonAuthToken](../README.md#AddonAuthToken), [PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

