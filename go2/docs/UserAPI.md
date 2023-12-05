# \UserAPI

All URIs are relative to *https://api.bitrise.io/v0.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UserMachineTypeUpdate**](UserAPI.md#UserMachineTypeUpdate) | **Patch** /user/{user-slug}/apps/machine_types | Migrate machine types
[**UserPlan**](UserAPI.md#UserPlan) | **Get** /me/plan | The subscription plan of the user
[**UserProfile**](UserAPI.md#UserProfile) | **Get** /me | Get your profile info
[**UserShow**](UserAPI.md#UserShow) | **Get** /users/{user-slug} | Get a specific user



## UserMachineTypeUpdate

> V0OrganizationUpdateMachineTypeResponse UserMachineTypeUpdate(ctx, userSlug).Types(types).Execute()

Migrate machine types



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
    types := *openapiclient.NewV0OrganizationUpdateMachineTypeParams() // V0OrganizationUpdateMachineTypeParams | Machine type to migrate from and to

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserAPI.UserMachineTypeUpdate(context.Background(), userSlug).Types(types).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UserMachineTypeUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserMachineTypeUpdate`: V0OrganizationUpdateMachineTypeResponse
    fmt.Fprintf(os.Stdout, "Response from `UserAPI.UserMachineTypeUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userSlug** | **string** | User slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserMachineTypeUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **types** | [**V0OrganizationUpdateMachineTypeParams**](V0OrganizationUpdateMachineTypeParams.md) | Machine type to migrate from and to | 

### Return type

[**V0OrganizationUpdateMachineTypeResponse**](V0OrganizationUpdateMachineTypeResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserPlan

> V0UserPlanRespModel UserPlan(ctx).Execute()

The subscription plan of the user



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserAPI.UserPlan(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UserPlan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserPlan`: V0UserPlanRespModel
    fmt.Fprintf(os.Stdout, "Response from `UserAPI.UserPlan`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUserPlanRequest struct via the builder pattern


### Return type

[**V0UserPlanRespModel**](V0UserPlanRespModel.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserProfile

> V0UserProfileRespModel UserProfile(ctx).Execute()

Get your profile info



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserAPI.UserProfile(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UserProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserProfile`: V0UserProfileRespModel
    fmt.Fprintf(os.Stdout, "Response from `UserAPI.UserProfile`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUserProfileRequest struct via the builder pattern


### Return type

[**V0UserProfileRespModel**](V0UserProfileRespModel.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserShow

> V0UserProfileRespModel UserShow(ctx, userSlug).Execute()

Get a specific user



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
    resp, r, err := apiClient.UserAPI.UserShow(context.Background(), userSlug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UserShow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserShow`: V0UserProfileRespModel
    fmt.Fprintf(os.Stdout, "Response from `UserAPI.UserShow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userSlug** | **string** | User slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserShowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V0UserProfileRespModel**](V0UserProfileRespModel.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

