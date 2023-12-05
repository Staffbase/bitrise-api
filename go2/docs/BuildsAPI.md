# \BuildsAPI

All URIs are relative to *https://api.bitrise.io/v0.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ArchivedBuildsList**](BuildsAPI.md#ArchivedBuildsList) | **Get** /apps/{app-slug}/archived-builds | List 1000 archived builds of an app
[**BuildAbort**](BuildsAPI.md#BuildAbort) | **Post** /apps/{app-slug}/builds/{build-slug}/abort | Abort a specific build
[**BuildBitriseYmlShow**](BuildsAPI.md#BuildBitriseYmlShow) | **Get** /apps/{app-slug}/builds/{build-slug}/bitrise.yml | Get the bitrise.yml of a build
[**BuildList**](BuildsAPI.md#BuildList) | **Get** /apps/{app-slug}/builds | List all builds of an app
[**BuildListAll**](BuildsAPI.md#BuildListAll) | **Get** /builds | List all builds
[**BuildLog**](BuildsAPI.md#BuildLog) | **Get** /apps/{app-slug}/builds/{build-slug}/log | Get the build log of a build
[**BuildShow**](BuildsAPI.md#BuildShow) | **Get** /apps/{app-slug}/builds/{build-slug} | Get a build of a given app
[**BuildTrigger**](BuildsAPI.md#BuildTrigger) | **Post** /apps/{app-slug}/builds | Trigger a new build/pipeline
[**BuildWorkflowList**](BuildsAPI.md#BuildWorkflowList) | **Get** /apps/{app-slug}/build-workflows | List the workflows of an app



## ArchivedBuildsList

> V0ArchivedBuildListResponseModel ArchivedBuildsList(ctx, appSlug).After(after).Before(before).Execute()

List 1000 archived builds of an app



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
    after := int32(56) // int32 | Listed builds should only be the ones which ran after the given date (Unix Timestamp)
    before := int32(56) // int32 | Listed builds should only be the ones which ran before the given date (Unix Timestamp)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuildsAPI.ArchivedBuildsList(context.Background(), appSlug).After(after).Before(before).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuildsAPI.ArchivedBuildsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ArchivedBuildsList`: V0ArchivedBuildListResponseModel
    fmt.Fprintf(os.Stdout, "Response from `BuildsAPI.ArchivedBuildsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appSlug** | **string** | App slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiArchivedBuildsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **after** | **int32** | Listed builds should only be the ones which ran after the given date (Unix Timestamp) | 
 **before** | **int32** | Listed builds should only be the ones which ran before the given date (Unix Timestamp) | 

### Return type

[**V0ArchivedBuildListResponseModel**](V0ArchivedBuildListResponseModel.md)

### Authorization

[AddonAuthToken](../README.md#AddonAuthToken), [PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BuildAbort

> V0BuildAbortResponseModel BuildAbort(ctx, appSlug, buildSlug).BuildAbortParams(buildAbortParams).Execute()

Abort a specific build



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
    buildSlug := "buildSlug_example" // string | Build slug
    buildAbortParams := *openapiclient.NewV0BuildAbortParams("AbortReason_example", false, false) // V0BuildAbortParams | Build abort parameters

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuildsAPI.BuildAbort(context.Background(), appSlug, buildSlug).BuildAbortParams(buildAbortParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuildsAPI.BuildAbort``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BuildAbort`: V0BuildAbortResponseModel
    fmt.Fprintf(os.Stdout, "Response from `BuildsAPI.BuildAbort`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appSlug** | **string** | App slug | 
**buildSlug** | **string** | Build slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiBuildAbortRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **buildAbortParams** | [**V0BuildAbortParams**](V0BuildAbortParams.md) | Build abort parameters | 

### Return type

[**V0BuildAbortResponseModel**](V0BuildAbortResponseModel.md)

### Authorization

[AddonAuthToken](../README.md#AddonAuthToken), [PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BuildBitriseYmlShow

> string BuildBitriseYmlShow(ctx, appSlug, buildSlug).Execute()

Get the bitrise.yml of a build



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
    buildSlug := "buildSlug_example" // string | Build slug

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuildsAPI.BuildBitriseYmlShow(context.Background(), appSlug, buildSlug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuildsAPI.BuildBitriseYmlShow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BuildBitriseYmlShow`: string
    fmt.Fprintf(os.Stdout, "Response from `BuildsAPI.BuildBitriseYmlShow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appSlug** | **string** | App slug | 
**buildSlug** | **string** | Build slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiBuildBitriseYmlShowRequest struct via the builder pattern


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


## BuildList

> V0BuildListResponseModel BuildList(ctx, appSlug).SortBy(sortBy).Branch(branch).Workflow(workflow).CommitMessage(commitMessage).TriggerEventType(triggerEventType).PullRequestId(pullRequestId).BuildNumber(buildNumber).After(after).Before(before).Status(status).IsPipelineBuild(isPipelineBuild).Next(next).Limit(limit).Execute()

List all builds of an app



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
    sortBy := "sortBy_example" // string | Order of builds: sort them based on when they were created or the time when they were triggered (optional)
    branch := "branch_example" // string | The branch which was built (optional)
    workflow := "workflow_example" // string | The name of the workflow used for the build (optional)
    commitMessage := "commitMessage_example" // string | The commit message of the build (optional)
    triggerEventType := "triggerEventType_example" // string | The event that triggered the build (push, pull-request, tag) (optional)
    pullRequestId := int32(56) // int32 | The id of the pull request that triggered the build (optional)
    buildNumber := int32(56) // int32 | The build number (optional)
    after := int32(56) // int32 | List builds run after a given date (Unix Timestamp) (optional)
    before := int32(56) // int32 | List builds run before a given date (Unix Timestamp) (optional)
    status := int32(56) // int32 | The status of the build: not finished (0), successful (1), failed (2), aborted with failure (3), aborted with success (4) (optional)
    isPipelineBuild := true // bool | Whether the builds are part of a pipeline or not (optional)
    next := "next_example" // string | Slug of the first build in the response (optional)
    limit := int32(56) // int32 | Max number of elements per page (default: 50) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuildsAPI.BuildList(context.Background(), appSlug).SortBy(sortBy).Branch(branch).Workflow(workflow).CommitMessage(commitMessage).TriggerEventType(triggerEventType).PullRequestId(pullRequestId).BuildNumber(buildNumber).After(after).Before(before).Status(status).IsPipelineBuild(isPipelineBuild).Next(next).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuildsAPI.BuildList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BuildList`: V0BuildListResponseModel
    fmt.Fprintf(os.Stdout, "Response from `BuildsAPI.BuildList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appSlug** | **string** | App slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiBuildListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sortBy** | **string** | Order of builds: sort them based on when they were created or the time when they were triggered | 
 **branch** | **string** | The branch which was built | 
 **workflow** | **string** | The name of the workflow used for the build | 
 **commitMessage** | **string** | The commit message of the build | 
 **triggerEventType** | **string** | The event that triggered the build (push, pull-request, tag) | 
 **pullRequestId** | **int32** | The id of the pull request that triggered the build | 
 **buildNumber** | **int32** | The build number | 
 **after** | **int32** | List builds run after a given date (Unix Timestamp) | 
 **before** | **int32** | List builds run before a given date (Unix Timestamp) | 
 **status** | **int32** | The status of the build: not finished (0), successful (1), failed (2), aborted with failure (3), aborted with success (4) | 
 **isPipelineBuild** | **bool** | Whether the builds are part of a pipeline or not | 
 **next** | **string** | Slug of the first build in the response | 
 **limit** | **int32** | Max number of elements per page (default: 50) | 

### Return type

[**V0BuildListResponseModel**](V0BuildListResponseModel.md)

### Authorization

[AddonAuthToken](../README.md#AddonAuthToken), [PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BuildListAll

> V0BuildListAllResponseModel BuildListAll(ctx).OwnerSlug(ownerSlug).IsOnHold(isOnHold).Status(status).Next(next).Limit(limit).Execute()

List all builds



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
    ownerSlug := "ownerSlug_example" // string | The slug of the owner of the app or apps (optional)
    isOnHold := true // bool | Indicates whether the build has started yet (true: the build hasn't started) (optional)
    status := int32(56) // int32 | The status of the build: not finished (0), successful (1), failed (2), aborted with failure (3), aborted with success (4) (optional)
    next := "next_example" // string | Slug of the first build in the response (optional)
    limit := int32(56) // int32 | Max number of elements per page (default: 50) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuildsAPI.BuildListAll(context.Background()).OwnerSlug(ownerSlug).IsOnHold(isOnHold).Status(status).Next(next).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuildsAPI.BuildListAll``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BuildListAll`: V0BuildListAllResponseModel
    fmt.Fprintf(os.Stdout, "Response from `BuildsAPI.BuildListAll`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBuildListAllRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ownerSlug** | **string** | The slug of the owner of the app or apps | 
 **isOnHold** | **bool** | Indicates whether the build has started yet (true: the build hasn&#39;t started) | 
 **status** | **int32** | The status of the build: not finished (0), successful (1), failed (2), aborted with failure (3), aborted with success (4) | 
 **next** | **string** | Slug of the first build in the response | 
 **limit** | **int32** | Max number of elements per page (default: 50) | 

### Return type

[**V0BuildListAllResponseModel**](V0BuildListAllResponseModel.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BuildLog

> BuildLog(ctx, appSlug, buildSlug).Execute()

Get the build log of a build



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
    buildSlug := "buildSlug_example" // string | Build slug

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.BuildsAPI.BuildLog(context.Background(), appSlug, buildSlug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuildsAPI.BuildLog``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appSlug** | **string** | App slug | 
**buildSlug** | **string** | Build slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiBuildLogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[AddonAuthToken](../README.md#AddonAuthToken), [PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BuildShow

> V0BuildShowResponseModel BuildShow(ctx, appSlug, buildSlug).Execute()

Get a build of a given app



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
    buildSlug := "buildSlug_example" // string | Build slug

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuildsAPI.BuildShow(context.Background(), appSlug, buildSlug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuildsAPI.BuildShow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BuildShow`: V0BuildShowResponseModel
    fmt.Fprintf(os.Stdout, "Response from `BuildsAPI.BuildShow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appSlug** | **string** | App slug | 
**buildSlug** | **string** | Build slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiBuildShowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**V0BuildShowResponseModel**](V0BuildShowResponseModel.md)

### Authorization

[AddonAuthToken](../README.md#AddonAuthToken), [PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BuildTrigger

> V0BuildTriggerRespModel BuildTrigger(ctx, appSlug).BuildParams(buildParams).Execute()

Trigger a new build/pipeline



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
    buildParams := *openapiclient.NewV0BuildTriggerParams() // V0BuildTriggerParams | Build trigger parameters

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuildsAPI.BuildTrigger(context.Background(), appSlug).BuildParams(buildParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuildsAPI.BuildTrigger``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BuildTrigger`: V0BuildTriggerRespModel
    fmt.Fprintf(os.Stdout, "Response from `BuildsAPI.BuildTrigger`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appSlug** | **string** | App slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiBuildTriggerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **buildParams** | [**V0BuildTriggerParams**](V0BuildTriggerParams.md) | Build trigger parameters | 

### Return type

[**V0BuildTriggerRespModel**](V0BuildTriggerRespModel.md)

### Authorization

[AddonAuthToken](../README.md#AddonAuthToken), [PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BuildWorkflowList

> V0BuildWorkflowListResponseModel BuildWorkflowList(ctx, appSlug).Execute()

List the workflows of an app



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
    resp, r, err := apiClient.BuildsAPI.BuildWorkflowList(context.Background(), appSlug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuildsAPI.BuildWorkflowList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BuildWorkflowList`: V0BuildWorkflowListResponseModel
    fmt.Fprintf(os.Stdout, "Response from `BuildsAPI.BuildWorkflowList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appSlug** | **string** | App slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiBuildWorkflowListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V0BuildWorkflowListResponseModel**](V0BuildWorkflowListResponseModel.md)

### Authorization

[AddonAuthToken](../README.md#AddonAuthToken), [PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

