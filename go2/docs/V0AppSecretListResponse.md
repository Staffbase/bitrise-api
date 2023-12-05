# V0AppSecretListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Secrets** | Pointer to [**[]V0AppSecret**](V0AppSecret.md) |  | [optional] 

## Methods

### NewV0AppSecretListResponse

`func NewV0AppSecretListResponse() *V0AppSecretListResponse`

NewV0AppSecretListResponse instantiates a new V0AppSecretListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0AppSecretListResponseWithDefaults

`func NewV0AppSecretListResponseWithDefaults() *V0AppSecretListResponse`

NewV0AppSecretListResponseWithDefaults instantiates a new V0AppSecretListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecrets

`func (o *V0AppSecretListResponse) GetSecrets() []V0AppSecret`

GetSecrets returns the Secrets field if non-nil, zero value otherwise.

### GetSecretsOk

`func (o *V0AppSecretListResponse) GetSecretsOk() (*[]V0AppSecret, bool)`

GetSecretsOk returns a tuple with the Secrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecrets

`func (o *V0AppSecretListResponse) SetSecrets(v []V0AppSecret)`

SetSecrets sets Secrets field to given value.

### HasSecrets

`func (o *V0AppSecretListResponse) HasSecrets() bool`

HasSecrets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


