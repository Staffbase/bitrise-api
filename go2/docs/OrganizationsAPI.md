# \OrganizationsAPI

All URIs are relative to *https://api.bitrise.io/v0.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrgList**](OrganizationsAPI.md#OrgList) | **Get** /organizations | List the organizations that the user is part of
[**OrgShow**](OrganizationsAPI.md#OrgShow) | **Get** /organizations/{org-slug} | Get a specified organization.
[**OrganizationMachineTypeUpdate**](OrganizationsAPI.md#OrganizationMachineTypeUpdate) | **Patch** /organizations/{org-slug}/apps/machine_types | Migrate machine types
[**OrganzationGroupsList**](OrganizationsAPI.md#OrganzationGroupsList) | **Get** /organizations/{org-slug}/groups | List organizations groups



## OrgList

> V0OrganizationListRespModel OrgList(ctx).Execute()

List the organizations that the user is part of



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
    resp, r, err := apiClient.OrganizationsAPI.OrgList(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.OrgList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrgList`: V0OrganizationListRespModel
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.OrgList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiOrgListRequest struct via the builder pattern


### Return type

[**V0OrganizationListRespModel**](V0OrganizationListRespModel.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrgShow

> V0OrganizationRespModel OrgShow(ctx, orgSlug).Execute()

Get a specified organization.



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
    orgSlug := "orgSlug_example" // string | The organization slug

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsAPI.OrgShow(context.Background(), orgSlug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.OrgShow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrgShow`: V0OrganizationRespModel
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.OrgShow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgSlug** | **string** | The organization slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgShowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V0OrganizationRespModel**](V0OrganizationRespModel.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationMachineTypeUpdate

> V0OrganizationUpdateMachineTypeResponse OrganizationMachineTypeUpdate(ctx, orgSlug).Types(types).Execute()

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
    orgSlug := "orgSlug_example" // string | Organization slug
    types := *openapiclient.NewV0OrganizationUpdateMachineTypeParams() // V0OrganizationUpdateMachineTypeParams | Machine type to migrate from and to

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsAPI.OrganizationMachineTypeUpdate(context.Background(), orgSlug).Types(types).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.OrganizationMachineTypeUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationMachineTypeUpdate`: V0OrganizationUpdateMachineTypeResponse
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.OrganizationMachineTypeUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgSlug** | **string** | Organization slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationMachineTypeUpdateRequest struct via the builder pattern


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


## OrganzationGroupsList

> []OrganzationGroupsList200ResponseInner OrganzationGroupsList(ctx, orgSlug).Execute()

List organizations groups



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
    orgSlug := "orgSlug_example" // string | slug of the organization

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsAPI.OrganzationGroupsList(context.Background(), orgSlug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.OrganzationGroupsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganzationGroupsList`: []OrganzationGroupsList200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.OrganzationGroupsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgSlug** | **string** | slug of the organization | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganzationGroupsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]OrganzationGroupsList200ResponseInner**](OrganzationGroupsList200ResponseInner.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

