# \PipelinesAPI

All URIs are relative to *https://api.bitrise.io/v0.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PipelineAbort**](PipelinesAPI.md#PipelineAbort) | **Post** /apps/{app-slug}/pipelines/{pipeline-id}/abort | Aborts a pipeline
[**PipelineList**](PipelinesAPI.md#PipelineList) | **Get** /apps/{app-slug}/pipelines | List all pipelines and standalone builds of an app
[**PipelineListAll**](PipelinesAPI.md#PipelineListAll) | **Get** /pipelines | List all pipelines/standalone builds
[**PipelineRebuild**](PipelinesAPI.md#PipelineRebuild) | **Post** /apps/{app-slug}/pipelines/{pipeline-id}/rebuild | Rebuilds a pipeline
[**PipelineShow**](PipelinesAPI.md#PipelineShow) | **Get** /apps/{app-slug}/pipelines/{pipeline-id} | Get a pipeline of a given app



## PipelineAbort

> PipelineAbort(ctx, appSlug, pipelineId).PipelineAbortParams(pipelineAbortParams).Execute()

Aborts a pipeline



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
    pipelineId := "pipelineId_example" // string | Pipeline id
    pipelineAbortParams := *openapiclient.NewV0PipelineAbortParams("AbortReason_example", false, false) // V0PipelineAbortParams | Pipeline abort parameters

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PipelinesAPI.PipelineAbort(context.Background(), appSlug, pipelineId).PipelineAbortParams(pipelineAbortParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PipelinesAPI.PipelineAbort``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appSlug** | **string** | App slug | 
**pipelineId** | **string** | Pipeline id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPipelineAbortRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pipelineAbortParams** | [**V0PipelineAbortParams**](V0PipelineAbortParams.md) | Pipeline abort parameters | 

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


## PipelineList

> V0PipelineListResponseModel PipelineList(ctx, appSlug).After(after).Before(before).Branch(branch).BuildNumber(buildNumber).BuildEnvironment(buildEnvironment).CommitMessage(commitMessage).Limit(limit).Next(next).Pipeline(pipeline).Status(status).TriggerEventType(triggerEventType).Workflow(workflow).Execute()

List all pipelines and standalone builds of an app



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
    after := "after_example" // string | List pipelines/standalone builds run after a given date (RFC3339 time format) (optional)
    before := "before_example" // string | List pipelines/standalone builds run before a given date (RFC3339 time format) - was called 'next' earlier (optional)
    branch := "branch_example" // string | The branch which was built (optional)
    buildNumber := int32(56) // int32 | The pipeline/standalone build number (optional)
    buildEnvironment := "buildEnvironment_example" // string | The build environment of the listed builds (ci, local, all) - default: ci (optional)
    commitMessage := "commitMessage_example" // string | The commit message of the pipeline/standalone build (optional)
    limit := int32(56) // int32 | Max number of elements per page - default: 10 (optional)
    next := "next_example" // string | List pipelines/standalone builds run before a given date (RFC3339 time format) - deprecated (optional)
    pipeline := "pipeline_example" // string | Name of the pipeline (optional)
    status := "status_example" // string | The status of the pipeline/standalone build: initializing, on_hold, running, succeeded, failed, aborted, succeeded_with_abort (optional)
    triggerEventType := "triggerEventType_example" // string | The event that triggered the pipeline/standalone build (push, pull-request, tag) (optional)
    workflow := "workflow_example" // string | The name of the workflow used for the pipeline/standalone build (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PipelinesAPI.PipelineList(context.Background(), appSlug).After(after).Before(before).Branch(branch).BuildNumber(buildNumber).BuildEnvironment(buildEnvironment).CommitMessage(commitMessage).Limit(limit).Next(next).Pipeline(pipeline).Status(status).TriggerEventType(triggerEventType).Workflow(workflow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PipelinesAPI.PipelineList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PipelineList`: V0PipelineListResponseModel
    fmt.Fprintf(os.Stdout, "Response from `PipelinesAPI.PipelineList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appSlug** | **string** | App slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiPipelineListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **after** | **string** | List pipelines/standalone builds run after a given date (RFC3339 time format) | 
 **before** | **string** | List pipelines/standalone builds run before a given date (RFC3339 time format) - was called &#39;next&#39; earlier | 
 **branch** | **string** | The branch which was built | 
 **buildNumber** | **int32** | The pipeline/standalone build number | 
 **buildEnvironment** | **string** | The build environment of the listed builds (ci, local, all) - default: ci | 
 **commitMessage** | **string** | The commit message of the pipeline/standalone build | 
 **limit** | **int32** | Max number of elements per page - default: 10 | 
 **next** | **string** | List pipelines/standalone builds run before a given date (RFC3339 time format) - deprecated | 
 **pipeline** | **string** | Name of the pipeline | 
 **status** | **string** | The status of the pipeline/standalone build: initializing, on_hold, running, succeeded, failed, aborted, succeeded_with_abort | 
 **triggerEventType** | **string** | The event that triggered the pipeline/standalone build (push, pull-request, tag) | 
 **workflow** | **string** | The name of the workflow used for the pipeline/standalone build | 

### Return type

[**V0PipelineListResponseModel**](V0PipelineListResponseModel.md)

### Authorization

[AddonAuthToken](../README.md#AddonAuthToken), [PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PipelineListAll

> V0PipelineListAllResponseModel PipelineListAll(ctx).OwnerSlug(ownerSlug).Status(status).Next(next).Limit(limit).Execute()

List all pipelines/standalone builds



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
    status := int32(56) // int32 | The status of the pipelines/standalone build: initializing, on_hold, running, succeeded, failed, aborted, succeeded_with_abort (optional)
    next := "next_example" // string | Getting pipelines/standalone builds before the given parameter (RFC3339 time format) (optional)
    limit := int32(56) // int32 | Max number of elements per page (default: 10) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PipelinesAPI.PipelineListAll(context.Background()).OwnerSlug(ownerSlug).Status(status).Next(next).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PipelinesAPI.PipelineListAll``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PipelineListAll`: V0PipelineListAllResponseModel
    fmt.Fprintf(os.Stdout, "Response from `PipelinesAPI.PipelineListAll`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPipelineListAllRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ownerSlug** | **string** | The slug of the owner of the app or apps | 
 **status** | **int32** | The status of the pipelines/standalone build: initializing, on_hold, running, succeeded, failed, aborted, succeeded_with_abort | 
 **next** | **string** | Getting pipelines/standalone builds before the given parameter (RFC3339 time format) | 
 **limit** | **int32** | Max number of elements per page (default: 10) | 

### Return type

[**V0PipelineListAllResponseModel**](V0PipelineListAllResponseModel.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PipelineRebuild

> PipelineRebuild(ctx, appSlug, pipelineId).PipelineRebuildParams(pipelineRebuildParams).Execute()

Rebuilds a pipeline



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
    pipelineId := "pipelineId_example" // string | Pipeline id
    pipelineRebuildParams := *openapiclient.NewV0PipelineRebuildParams() // V0PipelineRebuildParams | Pipeline rebuild parameters

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PipelinesAPI.PipelineRebuild(context.Background(), appSlug, pipelineId).PipelineRebuildParams(pipelineRebuildParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PipelinesAPI.PipelineRebuild``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appSlug** | **string** | App slug | 
**pipelineId** | **string** | Pipeline id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPipelineRebuildRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pipelineRebuildParams** | [**V0PipelineRebuildParams**](V0PipelineRebuildParams.md) | Pipeline rebuild parameters | 

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


## PipelineShow

> V0PipelineShowResponseModel PipelineShow(ctx, appSlug, pipelineId).Execute()

Get a pipeline of a given app



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
    pipelineId := "pipelineId_example" // string | Pipeline id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PipelinesAPI.PipelineShow(context.Background(), appSlug, pipelineId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PipelinesAPI.PipelineShow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PipelineShow`: V0PipelineShowResponseModel
    fmt.Fprintf(os.Stdout, "Response from `PipelinesAPI.PipelineShow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appSlug** | **string** | App slug | 
**pipelineId** | **string** | Pipeline id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPipelineShowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**V0PipelineShowResponseModel**](V0PipelineShowResponseModel.md)

### Authorization

[AddonAuthToken](../README.md#AddonAuthToken), [PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

