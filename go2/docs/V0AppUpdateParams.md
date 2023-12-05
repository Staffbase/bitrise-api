# V0AppUpdateParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppleCredentialUserId** | Pointer to **int32** | The new apple credential user ID (recommendation: use the UI to set this) | [optional] 
**AppleCredentialUserSlug** | Pointer to **string** | The new apple credential user slug (recommendation: use the UI to set this) | [optional] 
**DefaultBranch** | Pointer to **string** | The new default branch for the application. | [optional] 
**IsPublic** | Pointer to **bool** | The new value of whether an application should be publicly visible | [optional] 
**RepositoryUrl** | Pointer to **string** | The new repository URL for the application. | [optional] 
**ServicesCredentialUserId** | Pointer to **int32** | The new service credential user ID (recommendation: use the UI to set this) | [optional] 
**Title** | Pointer to **string** | The new title of the application. | [optional] 

## Methods

### NewV0AppUpdateParams

`func NewV0AppUpdateParams() *V0AppUpdateParams`

NewV0AppUpdateParams instantiates a new V0AppUpdateParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0AppUpdateParamsWithDefaults

`func NewV0AppUpdateParamsWithDefaults() *V0AppUpdateParams`

NewV0AppUpdateParamsWithDefaults instantiates a new V0AppUpdateParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppleCredentialUserId

`func (o *V0AppUpdateParams) GetAppleCredentialUserId() int32`

GetAppleCredentialUserId returns the AppleCredentialUserId field if non-nil, zero value otherwise.

### GetAppleCredentialUserIdOk

`func (o *V0AppUpdateParams) GetAppleCredentialUserIdOk() (*int32, bool)`

GetAppleCredentialUserIdOk returns a tuple with the AppleCredentialUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppleCredentialUserId

`func (o *V0AppUpdateParams) SetAppleCredentialUserId(v int32)`

SetAppleCredentialUserId sets AppleCredentialUserId field to given value.

### HasAppleCredentialUserId

`func (o *V0AppUpdateParams) HasAppleCredentialUserId() bool`

HasAppleCredentialUserId returns a boolean if a field has been set.

### GetAppleCredentialUserSlug

`func (o *V0AppUpdateParams) GetAppleCredentialUserSlug() string`

GetAppleCredentialUserSlug returns the AppleCredentialUserSlug field if non-nil, zero value otherwise.

### GetAppleCredentialUserSlugOk

`func (o *V0AppUpdateParams) GetAppleCredentialUserSlugOk() (*string, bool)`

GetAppleCredentialUserSlugOk returns a tuple with the AppleCredentialUserSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppleCredentialUserSlug

`func (o *V0AppUpdateParams) SetAppleCredentialUserSlug(v string)`

SetAppleCredentialUserSlug sets AppleCredentialUserSlug field to given value.

### HasAppleCredentialUserSlug

`func (o *V0AppUpdateParams) HasAppleCredentialUserSlug() bool`

HasAppleCredentialUserSlug returns a boolean if a field has been set.

### GetDefaultBranch

`func (o *V0AppUpdateParams) GetDefaultBranch() string`

GetDefaultBranch returns the DefaultBranch field if non-nil, zero value otherwise.

### GetDefaultBranchOk

`func (o *V0AppUpdateParams) GetDefaultBranchOk() (*string, bool)`

GetDefaultBranchOk returns a tuple with the DefaultBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultBranch

`func (o *V0AppUpdateParams) SetDefaultBranch(v string)`

SetDefaultBranch sets DefaultBranch field to given value.

### HasDefaultBranch

`func (o *V0AppUpdateParams) HasDefaultBranch() bool`

HasDefaultBranch returns a boolean if a field has been set.

### GetIsPublic

`func (o *V0AppUpdateParams) GetIsPublic() bool`

GetIsPublic returns the IsPublic field if non-nil, zero value otherwise.

### GetIsPublicOk

`func (o *V0AppUpdateParams) GetIsPublicOk() (*bool, bool)`

GetIsPublicOk returns a tuple with the IsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublic

`func (o *V0AppUpdateParams) SetIsPublic(v bool)`

SetIsPublic sets IsPublic field to given value.

### HasIsPublic

`func (o *V0AppUpdateParams) HasIsPublic() bool`

HasIsPublic returns a boolean if a field has been set.

### GetRepositoryUrl

`func (o *V0AppUpdateParams) GetRepositoryUrl() string`

GetRepositoryUrl returns the RepositoryUrl field if non-nil, zero value otherwise.

### GetRepositoryUrlOk

`func (o *V0AppUpdateParams) GetRepositoryUrlOk() (*string, bool)`

GetRepositoryUrlOk returns a tuple with the RepositoryUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryUrl

`func (o *V0AppUpdateParams) SetRepositoryUrl(v string)`

SetRepositoryUrl sets RepositoryUrl field to given value.

### HasRepositoryUrl

`func (o *V0AppUpdateParams) HasRepositoryUrl() bool`

HasRepositoryUrl returns a boolean if a field has been set.

### GetServicesCredentialUserId

`func (o *V0AppUpdateParams) GetServicesCredentialUserId() int32`

GetServicesCredentialUserId returns the ServicesCredentialUserId field if non-nil, zero value otherwise.

### GetServicesCredentialUserIdOk

`func (o *V0AppUpdateParams) GetServicesCredentialUserIdOk() (*int32, bool)`

GetServicesCredentialUserIdOk returns a tuple with the ServicesCredentialUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicesCredentialUserId

`func (o *V0AppUpdateParams) SetServicesCredentialUserId(v int32)`

SetServicesCredentialUserId sets ServicesCredentialUserId field to given value.

### HasServicesCredentialUserId

`func (o *V0AppUpdateParams) HasServicesCredentialUserId() bool`

HasServicesCredentialUserId returns a boolean if a field has been set.

### GetTitle

`func (o *V0AppUpdateParams) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *V0AppUpdateParams) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *V0AppUpdateParams) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *V0AppUpdateParams) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


