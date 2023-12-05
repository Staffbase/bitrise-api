# \AddonsAPI

All URIs are relative to *https://api.bitrise.io/v0.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddonListByApp**](AddonsAPI.md#AddonListByApp) | **Get** /apps/{app-slug}/addons | Get list of the addons for apps
[**AddonListByOrganization**](AddonsAPI.md#AddonListByOrganization) | **Get** /organizations/{organization-slug}/addons | Get list of the addons for organization
[**AddonListByUser**](AddonsAPI.md#AddonListByUser) | **Get** /users/{user-slug}/addons | Get list of the addons for user
[**AddonsList**](AddonsAPI.md#AddonsList) | **Get** /addons | Get list of available Bitrise addons
[**AddonsShow**](AddonsAPI.md#AddonsShow) | **Get** /addons/{addon-id} | Get a specific Bitrise addon



## AddonListByApp

> V0AppAddOnsListResponseModel AddonListByApp(ctx, appSlug).Execute()

Get list of the addons for apps



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
    resp, r, err := apiClient.AddonsAPI.AddonListByApp(context.Background(), appSlug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AddonsAPI.AddonListByApp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddonListByApp`: V0AppAddOnsListResponseModel
    fmt.Fprintf(os.Stdout, "Response from `AddonsAPI.AddonListByApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appSlug** | **string** | App slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddonListByAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V0AppAddOnsListResponseModel**](V0AppAddOnsListResponseModel.md)

### Authorization

[AddonAuthToken](../README.md#AddonAuthToken), [PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddonListByOrganization

> V0OwnerAddOnsListResponseModel AddonListByOrganization(ctx, organizationSlug).Execute()

Get list of the addons for organization



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
    organizationSlug := "organizationSlug_example" // string | Organization slug

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AddonsAPI.AddonListByOrganization(context.Background(), organizationSlug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AddonsAPI.AddonListByOrganization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddonListByOrganization`: V0OwnerAddOnsListResponseModel
    fmt.Fprintf(os.Stdout, "Response from `AddonsAPI.AddonListByOrganization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationSlug** | **string** | Organization slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddonListByOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V0OwnerAddOnsListResponseModel**](V0OwnerAddOnsListResponseModel.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddonListByUser

> V0OwnerAddOnsListResponseModel AddonListByUser(ctx, userSlug).Execute()

Get list of the addons for user



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
    resp, r, err := apiClient.AddonsAPI.AddonListByUser(context.Background(), userSlug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AddonsAPI.AddonListByUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddonListByUser`: V0OwnerAddOnsListResponseModel
    fmt.Fprintf(os.Stdout, "Response from `AddonsAPI.AddonListByUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userSlug** | **string** | User slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddonListByUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V0OwnerAddOnsListResponseModel**](V0OwnerAddOnsListResponseModel.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddonsList

> V0AddonsListResponseModel AddonsList(ctx).Execute()

Get list of available Bitrise addons



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
    resp, r, err := apiClient.AddonsAPI.AddonsList(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AddonsAPI.AddonsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddonsList`: V0AddonsListResponseModel
    fmt.Fprintf(os.Stdout, "Response from `AddonsAPI.AddonsList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAddonsListRequest struct via the builder pattern


### Return type

[**V0AddonsListResponseModel**](V0AddonsListResponseModel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddonsShow

> V0AddonsShowResponseModel AddonsShow(ctx, addonId).Execute()

Get a specific Bitrise addon



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
    addonId := "addonId_example" // string | Addon ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AddonsAPI.AddonsShow(context.Background(), addonId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AddonsAPI.AddonsShow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddonsShow`: V0AddonsShowResponseModel
    fmt.Fprintf(os.Stdout, "Response from `AddonsAPI.AddonsShow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**addonId** | **string** | Addon ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddonsShowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V0AddonsShowResponseModel**](V0AddonsShowResponseModel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

