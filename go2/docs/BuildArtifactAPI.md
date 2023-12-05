# \BuildArtifactAPI

All URIs are relative to *https://api.bitrise.io/v0.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ArtifactDelete**](BuildArtifactAPI.md#ArtifactDelete) | **Delete** /apps/{app-slug}/builds/{build-slug}/artifacts/{artifact-slug} | Delete a build artifact
[**ArtifactList**](BuildArtifactAPI.md#ArtifactList) | **Get** /apps/{app-slug}/builds/{build-slug}/artifacts | Get a list of all build artifacts
[**ArtifactShow**](BuildArtifactAPI.md#ArtifactShow) | **Get** /apps/{app-slug}/builds/{build-slug}/artifacts/{artifact-slug} | Get a specific build artifact
[**ArtifactUpdate**](BuildArtifactAPI.md#ArtifactUpdate) | **Patch** /apps/{app-slug}/builds/{build-slug}/artifacts/{artifact-slug} | Update a build artifact



## ArtifactDelete

> V0ArtifactDeleteResponseModel ArtifactDelete(ctx, appSlug, buildSlug, artifactSlug).Execute()

Delete a build artifact



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
    artifactSlug := "artifactSlug_example" // string | Artifact slug

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuildArtifactAPI.ArtifactDelete(context.Background(), appSlug, buildSlug, artifactSlug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuildArtifactAPI.ArtifactDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ArtifactDelete`: V0ArtifactDeleteResponseModel
    fmt.Fprintf(os.Stdout, "Response from `BuildArtifactAPI.ArtifactDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appSlug** | **string** | App slug | 
**buildSlug** | **string** | Build slug | 
**artifactSlug** | **string** | Artifact slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiArtifactDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**V0ArtifactDeleteResponseModel**](V0ArtifactDeleteResponseModel.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ArtifactList

> V0ArtifactListResponseModel ArtifactList(ctx, appSlug, buildSlug).Next(next).Limit(limit).Execute()

Get a list of all build artifacts



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
    next := "next_example" // string | Slug of the first build artifact in the response (optional)
    limit := int32(56) // int32 | Max number of build artifacts per page is 50. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuildArtifactAPI.ArtifactList(context.Background(), appSlug, buildSlug).Next(next).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuildArtifactAPI.ArtifactList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ArtifactList`: V0ArtifactListResponseModel
    fmt.Fprintf(os.Stdout, "Response from `BuildArtifactAPI.ArtifactList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appSlug** | **string** | App slug | 
**buildSlug** | **string** | Build slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiArtifactListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **next** | **string** | Slug of the first build artifact in the response | 
 **limit** | **int32** | Max number of build artifacts per page is 50. | 

### Return type

[**V0ArtifactListResponseModel**](V0ArtifactListResponseModel.md)

### Authorization

[AddonAuthToken](../README.md#AddonAuthToken), [PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ArtifactShow

> V0ArtifactShowResponseModel ArtifactShow(ctx, appSlug, buildSlug, artifactSlug).Download(download).Execute()

Get a specific build artifact



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
    artifactSlug := "artifactSlug_example" // string | Artifact slug
    download := int32(56) // int32 | Setting this will result in a redirect to the artifact download location (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuildArtifactAPI.ArtifactShow(context.Background(), appSlug, buildSlug, artifactSlug).Download(download).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuildArtifactAPI.ArtifactShow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ArtifactShow`: V0ArtifactShowResponseModel
    fmt.Fprintf(os.Stdout, "Response from `BuildArtifactAPI.ArtifactShow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appSlug** | **string** | App slug | 
**buildSlug** | **string** | Build slug | 
**artifactSlug** | **string** | Artifact slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiArtifactShowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **download** | **int32** | Setting this will result in a redirect to the artifact download location | 

### Return type

[**V0ArtifactShowResponseModel**](V0ArtifactShowResponseModel.md)

### Authorization

[AddonAuthToken](../README.md#AddonAuthToken), [PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ArtifactUpdate

> V0ArtifactShowResponseModel ArtifactUpdate(ctx, appSlug, buildSlug, artifactSlug).ArtifactParams(artifactParams).Execute()

Update a build artifact



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
    artifactSlug := "artifactSlug_example" // string | Artifact slug
    artifactParams := *openapiclient.NewV0ArtifactUpdateParams(false) // V0ArtifactUpdateParams | Artifact parameters

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuildArtifactAPI.ArtifactUpdate(context.Background(), appSlug, buildSlug, artifactSlug).ArtifactParams(artifactParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuildArtifactAPI.ArtifactUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ArtifactUpdate`: V0ArtifactShowResponseModel
    fmt.Fprintf(os.Stdout, "Response from `BuildArtifactAPI.ArtifactUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appSlug** | **string** | App slug | 
**buildSlug** | **string** | Build slug | 
**artifactSlug** | **string** | Artifact slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiArtifactUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **artifactParams** | [**V0ArtifactUpdateParams**](V0ArtifactUpdateParams.md) | Artifact parameters | 

### Return type

[**V0ArtifactShowResponseModel**](V0ArtifactShowResponseModel.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

