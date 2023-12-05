# V0SSHKeyUploadParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthSshPrivateKey** | **string** | The private part of the SSH key you would like to use | 
**AuthSshPublicKey** | **string** | The public part of the SSH key you would like to use | 
**IsRegisterKeyIntoProviderService** | Pointer to **bool** | If it&#39;s set to true, the provided SSH key will be registered at the provider of the application | [optional] 

## Methods

### NewV0SSHKeyUploadParams

`func NewV0SSHKeyUploadParams(authSshPrivateKey string, authSshPublicKey string, ) *V0SSHKeyUploadParams`

NewV0SSHKeyUploadParams instantiates a new V0SSHKeyUploadParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0SSHKeyUploadParamsWithDefaults

`func NewV0SSHKeyUploadParamsWithDefaults() *V0SSHKeyUploadParams`

NewV0SSHKeyUploadParamsWithDefaults instantiates a new V0SSHKeyUploadParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthSshPrivateKey

`func (o *V0SSHKeyUploadParams) GetAuthSshPrivateKey() string`

GetAuthSshPrivateKey returns the AuthSshPrivateKey field if non-nil, zero value otherwise.

### GetAuthSshPrivateKeyOk

`func (o *V0SSHKeyUploadParams) GetAuthSshPrivateKeyOk() (*string, bool)`

GetAuthSshPrivateKeyOk returns a tuple with the AuthSshPrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthSshPrivateKey

`func (o *V0SSHKeyUploadParams) SetAuthSshPrivateKey(v string)`

SetAuthSshPrivateKey sets AuthSshPrivateKey field to given value.


### GetAuthSshPublicKey

`func (o *V0SSHKeyUploadParams) GetAuthSshPublicKey() string`

GetAuthSshPublicKey returns the AuthSshPublicKey field if non-nil, zero value otherwise.

### GetAuthSshPublicKeyOk

`func (o *V0SSHKeyUploadParams) GetAuthSshPublicKeyOk() (*string, bool)`

GetAuthSshPublicKeyOk returns a tuple with the AuthSshPublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthSshPublicKey

`func (o *V0SSHKeyUploadParams) SetAuthSshPublicKey(v string)`

SetAuthSshPublicKey sets AuthSshPublicKey field to given value.


### GetIsRegisterKeyIntoProviderService

`func (o *V0SSHKeyUploadParams) GetIsRegisterKeyIntoProviderService() bool`

GetIsRegisterKeyIntoProviderService returns the IsRegisterKeyIntoProviderService field if non-nil, zero value otherwise.

### GetIsRegisterKeyIntoProviderServiceOk

`func (o *V0SSHKeyUploadParams) GetIsRegisterKeyIntoProviderServiceOk() (*bool, bool)`

GetIsRegisterKeyIntoProviderServiceOk returns a tuple with the IsRegisterKeyIntoProviderService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRegisterKeyIntoProviderService

`func (o *V0SSHKeyUploadParams) SetIsRegisterKeyIntoProviderService(v bool)`

SetIsRegisterKeyIntoProviderService sets IsRegisterKeyIntoProviderService field to given value.

### HasIsRegisterKeyIntoProviderService

`func (o *V0SSHKeyUploadParams) HasIsRegisterKeyIntoProviderService() bool`

HasIsRegisterKeyIntoProviderService returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


