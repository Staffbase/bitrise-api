# \AppleApiCredentialsAPI

All URIs are relative to *https://api.bitrise.io/v0.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppleApiCredentialList**](AppleApiCredentialsAPI.md#AppleApiCredentialList) | **Get** /users/{user-slug}/apple-api-credentials | List Apple API credentials for a specific user



## AppleApiCredentialList

> V0AppleAPICredentialsListResponse AppleApiCredentialList(ctx, userSlug).Execute()

List Apple API credentials for a specific user



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
    userSlug := "userSlug_example" // string | User slug

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppleApiCredentialsAPI.AppleApiCredentialList(context.Background(), userSlug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppleApiCredentialsAPI.AppleApiCredentialList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppleApiCredentialList`: V0AppleAPICredentialsListResponse
    fmt.Fprintf(os.Stdout, "Response from `AppleApiCredentialsAPI.AppleApiCredentialList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userSlug** | **string** | User slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppleApiCredentialListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V0AppleAPICredentialsListResponse**](V0AppleAPICredentialsListResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

