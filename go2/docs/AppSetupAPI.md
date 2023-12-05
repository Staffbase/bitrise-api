# \AppSetupAPI

All URIs are relative to *https://api.bitrise.io/v0.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppConfigCreate**](AppSetupAPI.md#AppConfigCreate) | **Post** /apps/{app-slug}/bitrise.yml | Upload a new bitrise.yml for your application.
[**AppCreate**](AppSetupAPI.md#AppCreate) | **Post** /apps/register | Add a new app
[**AppFinish**](AppSetupAPI.md#AppFinish) | **Post** /apps/{app-slug}/finish | Save the application at the end of the app registration process
[**AppSetupBitriseYmlConfigGet**](AppSetupAPI.md#AppSetupBitriseYmlConfigGet) | **Get** /apps/{app-slug}/bitrise.yml/config | Getting the location of the application&#39;s bitrise.yaml
[**AppSetupBitriseYmlConfigUpdate**](AppSetupAPI.md#AppSetupBitriseYmlConfigUpdate) | **Put** /apps/{app-slug}/bitrise.yml/config | Changing the location of the application&#39;s bitrise.yaml
[**AppWebhookCreate**](AppSetupAPI.md#AppWebhookCreate) | **Post** /apps/{app-slug}/register-webhook | Register an incoming webhook for a specific application
[**SshKeyCreate**](AppSetupAPI.md#SshKeyCreate) | **Post** /apps/{app-slug}/register-ssh-key | Add an SSH-key to a specific app



## AppConfigCreate

> map[string]string AppConfigCreate(ctx, appSlug).AppConfig(appConfig).Execute()

Upload a new bitrise.yml for your application.



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
    appConfig := *openapiclient.NewV0AppConfigRequestParam("AppConfigDatastoreYaml_example") // V0AppConfigRequestParam | App config parameters

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppSetupAPI.AppConfigCreate(context.Background(), appSlug).AppConfig(appConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppSetupAPI.AppConfigCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppConfigCreate`: map[string]string
    fmt.Fprintf(os.Stdout, "Response from `AppSetupAPI.AppConfigCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appSlug** | **string** | App slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppConfigCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appConfig** | [**V0AppConfigRequestParam**](V0AppConfigRequestParam.md) | App config parameters | 

### Return type

**map[string]string**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppCreate

> V0AppRespModel AppCreate(ctx).App(app).Execute()

Add a new app



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
    app := *openapiclient.NewV0AppUploadParams(false, "RepoUrl_example") // V0AppUploadParams | App parameters

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppSetupAPI.AppCreate(context.Background()).App(app).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppSetupAPI.AppCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppCreate`: V0AppRespModel
    fmt.Fprintf(os.Stdout, "Response from `AppSetupAPI.AppCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **app** | [**V0AppUploadParams**](V0AppUploadParams.md) | App parameters | 

### Return type

[**V0AppRespModel**](V0AppRespModel.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppFinish

> V0AppFinishRespModel AppFinish(ctx, appSlug).App(app).Execute()

Save the application at the end of the app registration process



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
    app := *openapiclient.NewV0AppFinishParams("ProjectType_example", "StackId_example") // V0AppFinishParams | App finish parameters

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppSetupAPI.AppFinish(context.Background(), appSlug).App(app).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppSetupAPI.AppFinish``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppFinish`: V0AppFinishRespModel
    fmt.Fprintf(os.Stdout, "Response from `AppSetupAPI.AppFinish`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appSlug** | **string** | App slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppFinishRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **app** | [**V0AppFinishParams**](V0AppFinishParams.md) | App finish parameters | 

### Return type

[**V0AppFinishRespModel**](V0AppFinishRespModel.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppSetupBitriseYmlConfigGet

> V0BitriseYMLConfigGetResponse AppSetupBitriseYmlConfigGet(ctx, appSlug).Execute()

Getting the location of the application's bitrise.yaml



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
    resp, r, err := apiClient.AppSetupAPI.AppSetupBitriseYmlConfigGet(context.Background(), appSlug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppSetupAPI.AppSetupBitriseYmlConfigGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppSetupBitriseYmlConfigGet`: V0BitriseYMLConfigGetResponse
    fmt.Fprintf(os.Stdout, "Response from `AppSetupAPI.AppSetupBitriseYmlConfigGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appSlug** | **string** | App slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppSetupBitriseYmlConfigGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V0BitriseYMLConfigGetResponse**](V0BitriseYMLConfigGetResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppSetupBitriseYmlConfigUpdate

> AppSetupBitriseYmlConfigUpdate(ctx, appSlug).App(app).Execute()

Changing the location of the application's bitrise.yaml



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
    app := *openapiclient.NewV0BitriseYMLConfigUpdateParams() // V0BitriseYMLConfigUpdateParams | Bitrise YML Config Update Params

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AppSetupAPI.AppSetupBitriseYmlConfigUpdate(context.Background(), appSlug).App(app).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppSetupAPI.AppSetupBitriseYmlConfigUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appSlug** | **string** | App slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppSetupBitriseYmlConfigUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **app** | [**V0BitriseYMLConfigUpdateParams**](V0BitriseYMLConfigUpdateParams.md) | Bitrise YML Config Update Params | 

### Return type

 (empty response body)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppWebhookCreate

> V0WebhookRespModel AppWebhookCreate(ctx, appSlug).Execute()

Register an incoming webhook for a specific application



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
    resp, r, err := apiClient.AppSetupAPI.AppWebhookCreate(context.Background(), appSlug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppSetupAPI.AppWebhookCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppWebhookCreate`: V0WebhookRespModel
    fmt.Fprintf(os.Stdout, "Response from `AppSetupAPI.AppWebhookCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appSlug** | **string** | App slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppWebhookCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V0WebhookRespModel**](V0WebhookRespModel.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SshKeyCreate

> V0SSHKeyRespModel SshKeyCreate(ctx, appSlug).SshKey(sshKey).Execute()

Add an SSH-key to a specific app



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
    sshKey := *openapiclient.NewV0SSHKeyUploadParams("AuthSshPrivateKey_example", "AuthSshPublicKey_example") // V0SSHKeyUploadParams | SSH key parameters

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppSetupAPI.SshKeyCreate(context.Background(), appSlug).SshKey(sshKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppSetupAPI.SshKeyCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SshKeyCreate`: V0SSHKeyRespModel
    fmt.Fprintf(os.Stdout, "Response from `AppSetupAPI.SshKeyCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appSlug** | **string** | App slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiSshKeyCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sshKey** | [**V0SSHKeyUploadParams**](V0SSHKeyUploadParams.md) | SSH key parameters | 

### Return type

[**V0SSHKeyRespModel**](V0SSHKeyRespModel.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

