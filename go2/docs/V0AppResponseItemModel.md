# V0AppResponseItemModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvatarUrl** | Pointer to [**NullsString**](NullsString.md) |  | [optional] 
**IsDisabled** | Pointer to **bool** |  | [optional] 
**IsGithubChecksEnabled** | Pointer to **bool** |  | [optional] 
**IsPublic** | Pointer to **bool** |  | [optional] 
**Owner** | Pointer to [**V0OwnerAccountResponseModel**](V0OwnerAccountResponseModel.md) |  | [optional] 
**ProjectType** | Pointer to [**NullsString**](NullsString.md) |  | [optional] 
**Provider** | Pointer to [**NullsString**](NullsString.md) |  | [optional] 
**RepoOwner** | Pointer to [**NullsString**](NullsString.md) |  | [optional] 
**RepoSlug** | Pointer to [**NullsString**](NullsString.md) |  | [optional] 
**RepoUrl** | Pointer to [**NullsString**](NullsString.md) |  | [optional] 
**Slug** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **int32** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 

## Methods

### NewV0AppResponseItemModel

`func NewV0AppResponseItemModel() *V0AppResponseItemModel`

NewV0AppResponseItemModel instantiates a new V0AppResponseItemModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0AppResponseItemModelWithDefaults

`func NewV0AppResponseItemModelWithDefaults() *V0AppResponseItemModel`

NewV0AppResponseItemModelWithDefaults instantiates a new V0AppResponseItemModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvatarUrl

`func (o *V0AppResponseItemModel) GetAvatarUrl() NullsString`

GetAvatarUrl returns the AvatarUrl field if non-nil, zero value otherwise.

### GetAvatarUrlOk

`func (o *V0AppResponseItemModel) GetAvatarUrlOk() (*NullsString, bool)`

GetAvatarUrlOk returns a tuple with the AvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrl

`func (o *V0AppResponseItemModel) SetAvatarUrl(v NullsString)`

SetAvatarUrl sets AvatarUrl field to given value.

### HasAvatarUrl

`func (o *V0AppResponseItemModel) HasAvatarUrl() bool`

HasAvatarUrl returns a boolean if a field has been set.

### GetIsDisabled

`func (o *V0AppResponseItemModel) GetIsDisabled() bool`

GetIsDisabled returns the IsDisabled field if non-nil, zero value otherwise.

### GetIsDisabledOk

`func (o *V0AppResponseItemModel) GetIsDisabledOk() (*bool, bool)`

GetIsDisabledOk returns a tuple with the IsDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDisabled

`func (o *V0AppResponseItemModel) SetIsDisabled(v bool)`

SetIsDisabled sets IsDisabled field to given value.

### HasIsDisabled

`func (o *V0AppResponseItemModel) HasIsDisabled() bool`

HasIsDisabled returns a boolean if a field has been set.

### GetIsGithubChecksEnabled

`func (o *V0AppResponseItemModel) GetIsGithubChecksEnabled() bool`

GetIsGithubChecksEnabled returns the IsGithubChecksEnabled field if non-nil, zero value otherwise.

### GetIsGithubChecksEnabledOk

`func (o *V0AppResponseItemModel) GetIsGithubChecksEnabledOk() (*bool, bool)`

GetIsGithubChecksEnabledOk returns a tuple with the IsGithubChecksEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGithubChecksEnabled

`func (o *V0AppResponseItemModel) SetIsGithubChecksEnabled(v bool)`

SetIsGithubChecksEnabled sets IsGithubChecksEnabled field to given value.

### HasIsGithubChecksEnabled

`func (o *V0AppResponseItemModel) HasIsGithubChecksEnabled() bool`

HasIsGithubChecksEnabled returns a boolean if a field has been set.

### GetIsPublic

`func (o *V0AppResponseItemModel) GetIsPublic() bool`

GetIsPublic returns the IsPublic field if non-nil, zero value otherwise.

### GetIsPublicOk

`func (o *V0AppResponseItemModel) GetIsPublicOk() (*bool, bool)`

GetIsPublicOk returns a tuple with the IsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublic

`func (o *V0AppResponseItemModel) SetIsPublic(v bool)`

SetIsPublic sets IsPublic field to given value.

### HasIsPublic

`func (o *V0AppResponseItemModel) HasIsPublic() bool`

HasIsPublic returns a boolean if a field has been set.

### GetOwner

`func (o *V0AppResponseItemModel) GetOwner() V0OwnerAccountResponseModel`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *V0AppResponseItemModel) GetOwnerOk() (*V0OwnerAccountResponseModel, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *V0AppResponseItemModel) SetOwner(v V0OwnerAccountResponseModel)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *V0AppResponseItemModel) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetProjectType

`func (o *V0AppResponseItemModel) GetProjectType() NullsString`

GetProjectType returns the ProjectType field if non-nil, zero value otherwise.

### GetProjectTypeOk

`func (o *V0AppResponseItemModel) GetProjectTypeOk() (*NullsString, bool)`

GetProjectTypeOk returns a tuple with the ProjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectType

`func (o *V0AppResponseItemModel) SetProjectType(v NullsString)`

SetProjectType sets ProjectType field to given value.

### HasProjectType

`func (o *V0AppResponseItemModel) HasProjectType() bool`

HasProjectType returns a boolean if a field has been set.

### GetProvider

`func (o *V0AppResponseItemModel) GetProvider() NullsString`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *V0AppResponseItemModel) GetProviderOk() (*NullsString, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *V0AppResponseItemModel) SetProvider(v NullsString)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *V0AppResponseItemModel) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetRepoOwner

`func (o *V0AppResponseItemModel) GetRepoOwner() NullsString`

GetRepoOwner returns the RepoOwner field if non-nil, zero value otherwise.

### GetRepoOwnerOk

`func (o *V0AppResponseItemModel) GetRepoOwnerOk() (*NullsString, bool)`

GetRepoOwnerOk returns a tuple with the RepoOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepoOwner

`func (o *V0AppResponseItemModel) SetRepoOwner(v NullsString)`

SetRepoOwner sets RepoOwner field to given value.

### HasRepoOwner

`func (o *V0AppResponseItemModel) HasRepoOwner() bool`

HasRepoOwner returns a boolean if a field has been set.

### GetRepoSlug

`func (o *V0AppResponseItemModel) GetRepoSlug() NullsString`

GetRepoSlug returns the RepoSlug field if non-nil, zero value otherwise.

### GetRepoSlugOk

`func (o *V0AppResponseItemModel) GetRepoSlugOk() (*NullsString, bool)`

GetRepoSlugOk returns a tuple with the RepoSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepoSlug

`func (o *V0AppResponseItemModel) SetRepoSlug(v NullsString)`

SetRepoSlug sets RepoSlug field to given value.

### HasRepoSlug

`func (o *V0AppResponseItemModel) HasRepoSlug() bool`

HasRepoSlug returns a boolean if a field has been set.

### GetRepoUrl

`func (o *V0AppResponseItemModel) GetRepoUrl() NullsString`

GetRepoUrl returns the RepoUrl field if non-nil, zero value otherwise.

### GetRepoUrlOk

`func (o *V0AppResponseItemModel) GetRepoUrlOk() (*NullsString, bool)`

GetRepoUrlOk returns a tuple with the RepoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepoUrl

`func (o *V0AppResponseItemModel) SetRepoUrl(v NullsString)`

SetRepoUrl sets RepoUrl field to given value.

### HasRepoUrl

`func (o *V0AppResponseItemModel) HasRepoUrl() bool`

HasRepoUrl returns a boolean if a field has been set.

### GetSlug

`func (o *V0AppResponseItemModel) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *V0AppResponseItemModel) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *V0AppResponseItemModel) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *V0AppResponseItemModel) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetStatus

`func (o *V0AppResponseItemModel) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *V0AppResponseItemModel) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *V0AppResponseItemModel) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *V0AppResponseItemModel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTitle

`func (o *V0AppResponseItemModel) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *V0AppResponseItemModel) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *V0AppResponseItemModel) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *V0AppResponseItemModel) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


