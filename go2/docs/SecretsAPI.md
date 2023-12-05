# \SecretsAPI

All URIs are relative to *https://api.bitrise.io/v0.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SecretDelete**](SecretsAPI.md#SecretDelete) | **Delete** /apps/{app-slug}/secrets/{secret-name} | Delete an application secret
[**SecretList**](SecretsAPI.md#SecretList) | **Get** /apps/{app-slug}/secrets | List the application secrets (with no values)
[**SecretUpsert**](SecretsAPI.md#SecretUpsert) | **Put** /apps/{app-slug}/secrets/{secret-name} | Upsert an application secret
[**SecretValueGet**](SecretsAPI.md#SecretValueGet) | **Get** /apps/{app-slug}/secrets/{secret-name}/value | Get the value of an (unprotected) application secrets



## SecretDelete

> SecretDelete(ctx, appSlug, secretName).Execute()

Delete an application secret



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
    secretName := "secretName_example" // string | Secret name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SecretsAPI.SecretDelete(context.Background(), appSlug, secretName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsAPI.SecretDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appSlug** | **string** | App slug | 
**secretName** | **string** | Secret name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSecretDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## SecretList

> V0AppSecretListResponse SecretList(ctx, appSlug).Execute()

List the application secrets (with no values)



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
    resp, r, err := apiClient.SecretsAPI.SecretList(context.Background(), appSlug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsAPI.SecretList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SecretList`: V0AppSecretListResponse
    fmt.Fprintf(os.Stdout, "Response from `SecretsAPI.SecretList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appSlug** | **string** | App slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiSecretListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V0AppSecretListResponse**](V0AppSecretListResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecretUpsert

> SecretUpsert(ctx, appSlug, secretName).App(app).Execute()

Upsert an application secret



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
    secretName := "secretName_example" // string | Secret name
    app := *openapiclient.NewV0AppSecretUpsertParams() // V0AppSecretUpsertParams | Secret parameters

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SecretsAPI.SecretUpsert(context.Background(), appSlug, secretName).App(app).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsAPI.SecretUpsert``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appSlug** | **string** | App slug | 
**secretName** | **string** | Secret name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSecretUpsertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **app** | [**V0AppSecretUpsertParams**](V0AppSecretUpsertParams.md) | Secret parameters | 

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


## SecretValueGet

> V0AppSecretGetValueResponse SecretValueGet(ctx, appSlug, secretName).Execute()

Get the value of an (unprotected) application secrets



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
    secretName := "secretName_example" // string | Secret name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsAPI.SecretValueGet(context.Background(), appSlug, secretName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsAPI.SecretValueGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SecretValueGet`: V0AppSecretGetValueResponse
    fmt.Fprintf(os.Stdout, "Response from `SecretsAPI.SecretValueGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appSlug** | **string** | App slug | 
**secretName** | **string** | Secret name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSecretValueGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**V0AppSecretGetValueResponse**](V0AppSecretGetValueResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

