# \ApplicationAPI

All URIs are relative to *https://api.bitrise.io/v0.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppConfigDatastoreShow**](ApplicationAPI.md#AppConfigDatastoreShow) | **Get** /apps/{app-slug}/bitrise.yml | Get bitrise.yml of a specific app
[**AppDelete**](ApplicationAPI.md#AppDelete) | **Delete** /apps/{app-slug} | Deletes an app
[**AppList**](ApplicationAPI.md#AppList) | **Get** /apps | Get list of the apps
[**AppListByOrganization**](ApplicationAPI.md#AppListByOrganization) | **Get** /organizations/{org-slug}/apps | Get list of the apps for an organization
[**AppListByUser**](ApplicationAPI.md#AppListByUser) | **Get** /users/{user-slug}/apps | Get list of the apps for a user
[**AppNotifications**](ApplicationAPI.md#AppNotifications) | **Patch** /apps/{app-slug}/update-email-notifications | Updates the app&#39;s notification settings
[**AppRolesQuery**](ApplicationAPI.md#AppRolesQuery) | **Get** /apps/{app-slug}/roles/{role-name} | Lists group roles for an app
[**AppRolesUpdate**](ApplicationAPI.md#AppRolesUpdate) | **Put** /apps/{app-slug}/roles/{role-name} | Replaces group roles for an app
[**AppShow**](ApplicationAPI.md#AppShow) | **Get** /apps/{app-slug} | Get a specific app
[**AppUpdate**](ApplicationAPI.md#AppUpdate) | **Patch** /apps/{app-slug} | Updates an app
[**BranchList**](ApplicationAPI.md#BranchList) | **Get** /apps/{app-slug}/branches | List the branches with existing builds of an app&#39;s repository



## AppConfigDatastoreShow

> string AppConfigDatastoreShow(ctx, appSlug).Execute()

Get bitrise.yml of a specific app



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
    resp, r, err := apiClient.ApplicationAPI.AppConfigDatastoreShow(context.Background(), appSlug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationAPI.AppConfigDatastoreShow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppConfigDatastoreShow`: string
    fmt.Fprintf(os.Stdout, "Response from `ApplicationAPI.AppConfigDatastoreShow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appSlug** | **string** | App slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppConfigDatastoreShowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

[AddonAuthToken](../README.md#AddonAuthToken), [PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppDelete

> V0AppDeleteRespModel AppDelete(ctx, appSlug).Execute()

Deletes an app



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
    resp, r, err := apiClient.ApplicationAPI.AppDelete(context.Background(), appSlug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationAPI.AppDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppDelete`: V0AppDeleteRespModel
    fmt.Fprintf(os.Stdout, "Response from `ApplicationAPI.AppDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appSlug** | **string** | App slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V0AppDeleteRespModel**](V0AppDeleteRespModel.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppList

> V0AppListResponseModel AppList(ctx).SortBy(sortBy).Next(next).Limit(limit).Title(title).ProjectType(projectType).Execute()

Get list of the apps



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
    sortBy := "sortBy_example" // string | Order of the applications: sort them based on when they were created or the time of their last build (optional)
    next := "next_example" // string | Slug of the first app in the response (optional)
    limit := int32(56) // int32 | Max number of elements per page (default: 50) (optional)
    title := "title_example" // string | The title of the app (optional)
    projectType := "projectType_example" // string | The project type of the app (eg. 'ios', 'android') (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationAPI.AppList(context.Background()).SortBy(sortBy).Next(next).Limit(limit).Title(title).ProjectType(projectType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationAPI.AppList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppList`: V0AppListResponseModel
    fmt.Fprintf(os.Stdout, "Response from `ApplicationAPI.AppList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sortBy** | **string** | Order of the applications: sort them based on when they were created or the time of their last build | 
 **next** | **string** | Slug of the first app in the response | 
 **limit** | **int32** | Max number of elements per page (default: 50) | 
 **title** | **string** | The title of the app | 
 **projectType** | **string** | The project type of the app (eg. &#39;ios&#39;, &#39;android&#39;) | 

### Return type

[**V0AppListResponseModel**](V0AppListResponseModel.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppListByOrganization

> V0AppListResponseModel AppListByOrganization(ctx, orgSlug).SortBy(sortBy).Next(next).Limit(limit).Title(title).ProjectType(projectType).Execute()

Get list of the apps for an organization



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
    sortBy := "sortBy_example" // string | Order of applications: sort them based on when they were created or the time of their last build (optional)
    next := "next_example" // string | Slug of the first app in the response (optional)
    limit := int32(56) // int32 | Max number of elements per page (default: 50) (optional)
    title := "title_example" // string | The title of the app (optional)
    projectType := "projectType_example" // string | The project type of the app (eg. 'ios', 'android') (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationAPI.AppListByOrganization(context.Background(), orgSlug).SortBy(sortBy).Next(next).Limit(limit).Title(title).ProjectType(projectType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationAPI.AppListByOrganization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppListByOrganization`: V0AppListResponseModel
    fmt.Fprintf(os.Stdout, "Response from `ApplicationAPI.AppListByOrganization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgSlug** | **string** | Organization slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppListByOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sortBy** | **string** | Order of applications: sort them based on when they were created or the time of their last build | 
 **next** | **string** | Slug of the first app in the response | 
 **limit** | **int32** | Max number of elements per page (default: 50) | 
 **title** | **string** | The title of the app | 
 **projectType** | **string** | The project type of the app (eg. &#39;ios&#39;, &#39;android&#39;) | 

### Return type

[**V0AppListResponseModel**](V0AppListResponseModel.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppListByUser

> V0AppListResponseModel AppListByUser(ctx, userSlug).SortBy(sortBy).Next(next).Limit(limit).Title(title).ProjectType(projectType).Execute()

Get list of the apps for a user



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
    sortBy := "sortBy_example" // string | Order of applications (optional)
    next := "next_example" // string | Slug of the first app in the response (optional)
    limit := int32(56) // int32 | Max number of elements per page (default: 50) (optional)
    title := "title_example" // string | The title of the app (optional)
    projectType := "projectType_example" // string | The project type of the app (eg. 'ios', 'android') (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationAPI.AppListByUser(context.Background(), userSlug).SortBy(sortBy).Next(next).Limit(limit).Title(title).ProjectType(projectType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationAPI.AppListByUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppListByUser`: V0AppListResponseModel
    fmt.Fprintf(os.Stdout, "Response from `ApplicationAPI.AppListByUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userSlug** | **string** | User slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppListByUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sortBy** | **string** | Order of applications | 
 **next** | **string** | Slug of the first app in the response | 
 **limit** | **int32** | Max number of elements per page (default: 50) | 
 **title** | **string** | The title of the app | 
 **projectType** | **string** | The project type of the app (eg. &#39;ios&#39;, &#39;android&#39;) | 

### Return type

[**V0AppListResponseModel**](V0AppListResponseModel.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppNotifications

> V0AppNotificationSettingsUpdateResponse AppNotifications(ctx, appSlug).NotificationSettingsParams(notificationSettingsParams).Execute()

Updates the app's notification settings



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
    notificationSettingsParams := *openapiclient.NewV0AppNotificationSettingsParams() // V0AppNotificationSettingsParams | App notification settings parameters

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationAPI.AppNotifications(context.Background(), appSlug).NotificationSettingsParams(notificationSettingsParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationAPI.AppNotifications``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppNotifications`: V0AppNotificationSettingsUpdateResponse
    fmt.Fprintf(os.Stdout, "Response from `ApplicationAPI.AppNotifications`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appSlug** | **string** | App slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppNotificationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **notificationSettingsParams** | [**V0AppNotificationSettingsParams**](V0AppNotificationSettingsParams.md) | App notification settings parameters | 

### Return type

[**V0AppNotificationSettingsUpdateResponse**](V0AppNotificationSettingsUpdateResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppRolesQuery

> AppRolesQuery200Response AppRolesQuery(ctx, appSlug, roleName).Execute()

Lists group roles for an app

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
    appSlug := "appSlug_example" // string | Slug of the app
    roleName := "roleName_example" // string | Name of the role being queried, supported values are: admin, manager (equals developer), and member (equals tester/qa)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationAPI.AppRolesQuery(context.Background(), appSlug, roleName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationAPI.AppRolesQuery``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppRolesQuery`: AppRolesQuery200Response
    fmt.Fprintf(os.Stdout, "Response from `ApplicationAPI.AppRolesQuery`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appSlug** | **string** | Slug of the app | 
**roleName** | **string** | Name of the role being queried, supported values are: admin, manager (equals developer), and member (equals tester/qa) | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppRolesQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AppRolesQuery200Response**](AppRolesQuery200Response.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppRolesUpdate

> AppRolesQuery200Response AppRolesUpdate(ctx, appSlug, roleName).Groups(groups).Execute()

Replaces group roles for an app



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
    appSlug := "appSlug_example" // string | Slug of the app
    roleName := "roleName_example" // string | Name of the role being modified, supported values are: admin, manager (equals developer), and member (equals tester/qa)
    groups := *openapiclient.NewAppRolesQuery200Response() // AppRolesQuery200Response | List of group slugs

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationAPI.AppRolesUpdate(context.Background(), appSlug, roleName).Groups(groups).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationAPI.AppRolesUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppRolesUpdate`: AppRolesQuery200Response
    fmt.Fprintf(os.Stdout, "Response from `ApplicationAPI.AppRolesUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appSlug** | **string** | Slug of the app | 
**roleName** | **string** | Name of the role being modified, supported values are: admin, manager (equals developer), and member (equals tester/qa) | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppRolesUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **groups** | [**AppRolesQuery200Response**](AppRolesQuery200Response.md) | List of group slugs | 

### Return type

[**AppRolesQuery200Response**](AppRolesQuery200Response.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppShow

> V0AppShowResponseModel AppShow(ctx, appSlug).Execute()

Get a specific app



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
    resp, r, err := apiClient.ApplicationAPI.AppShow(context.Background(), appSlug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationAPI.AppShow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppShow`: V0AppShowResponseModel
    fmt.Fprintf(os.Stdout, "Response from `ApplicationAPI.AppShow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appSlug** | **string** | App slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppShowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V0AppShowResponseModel**](V0AppShowResponseModel.md)

### Authorization

[AddonAuthToken](../README.md#AddonAuthToken), [PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppUpdate

> V0AppUpdateRespModel AppUpdate(ctx, appSlug).App(app).Execute()

Updates an app



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
    app := *openapiclient.NewV0AppUpdateParams() // V0AppUpdateParams | App update params. All fields are optional, omit the fields you don't wish to update.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationAPI.AppUpdate(context.Background(), appSlug).App(app).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationAPI.AppUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppUpdate`: V0AppUpdateRespModel
    fmt.Fprintf(os.Stdout, "Response from `ApplicationAPI.AppUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appSlug** | **string** | App slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **app** | [**V0AppUpdateParams**](V0AppUpdateParams.md) | App update params. All fields are optional, omit the fields you don&#39;t wish to update. | 

### Return type

[**V0AppUpdateRespModel**](V0AppUpdateRespModel.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BranchList

> V0BranchListResponseModel BranchList(ctx, appSlug).Execute()

List the branches with existing builds of an app's repository



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
    resp, r, err := apiClient.ApplicationAPI.BranchList(context.Background(), appSlug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationAPI.BranchList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BranchList`: V0BranchListResponseModel
    fmt.Fprintf(os.Stdout, "Response from `ApplicationAPI.BranchList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appSlug** | **string** | App slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiBranchListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V0BranchListResponseModel**](V0BranchListResponseModel.md)

### Authorization

[AddonAuthToken](../README.md#AddonAuthToken), [PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

