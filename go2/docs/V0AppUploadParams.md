# V0AppUploadParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultBranchName** | Pointer to **string** | The default branch of the repository. If it&#39;s not specified, it will be &#x60;master&#x60;. | [optional] 
**GitOwner** | Pointer to **string** | [Deprecated] You no longer need to provide this field. | [optional] 
**GitRepoSlug** | Pointer to **string** | [Deprecated] You no longer need to provide this field. | [optional] 
**IsPublic** | **bool** | If &#x60;true&#x60; then the repository visibility setting will be public, in case of &#x60;false&#x60; it will be private | 
**ManualApprovalEnabled** | Pointer to **bool** | Toggles whether manual approval should be enabled for the app&#39;s builds. If it&#39;s not specified, it will be &#x60;true&#x60;. | [optional] 
**OrganizationSlug** | Pointer to **string** | The slug of the organization, who will be the owner of the application. If it&#39;s not specified, then the authenticated user will be the owner. | [optional] 
**Provider** | Pointer to **string** | The git provider you are using, it can be &#x60;github&#x60;, &#x60;bitbucket&#x60;, &#x60;gitlab&#x60;, &#x60;gitlab-self-hosted&#x60; or &#x60;custom&#x60; | [optional] 
**RepoUrl** | **string** | The URL of your repository | 
**Title** | Pointer to **string** | The title of the application. If it&#39;s not specified, it will be the git repository&#39;s name. | [optional] 
**Type** | Pointer to **string** | [Deprecated] You no longer need to provide this field. | [optional] 

## Methods

### NewV0AppUploadParams

`func NewV0AppUploadParams(isPublic bool, repoUrl string, ) *V0AppUploadParams`

NewV0AppUploadParams instantiates a new V0AppUploadParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0AppUploadParamsWithDefaults

`func NewV0AppUploadParamsWithDefaults() *V0AppUploadParams`

NewV0AppUploadParamsWithDefaults instantiates a new V0AppUploadParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultBranchName

`func (o *V0AppUploadParams) GetDefaultBranchName() string`

GetDefaultBranchName returns the DefaultBranchName field if non-nil, zero value otherwise.

### GetDefaultBranchNameOk

`func (o *V0AppUploadParams) GetDefaultBranchNameOk() (*string, bool)`

GetDefaultBranchNameOk returns a tuple with the DefaultBranchName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultBranchName

`func (o *V0AppUploadParams) SetDefaultBranchName(v string)`

SetDefaultBranchName sets DefaultBranchName field to given value.

### HasDefaultBranchName

`func (o *V0AppUploadParams) HasDefaultBranchName() bool`

HasDefaultBranchName returns a boolean if a field has been set.

### GetGitOwner

`func (o *V0AppUploadParams) GetGitOwner() string`

GetGitOwner returns the GitOwner field if non-nil, zero value otherwise.

### GetGitOwnerOk

`func (o *V0AppUploadParams) GetGitOwnerOk() (*string, bool)`

GetGitOwnerOk returns a tuple with the GitOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitOwner

`func (o *V0AppUploadParams) SetGitOwner(v string)`

SetGitOwner sets GitOwner field to given value.

### HasGitOwner

`func (o *V0AppUploadParams) HasGitOwner() bool`

HasGitOwner returns a boolean if a field has been set.

### GetGitRepoSlug

`func (o *V0AppUploadParams) GetGitRepoSlug() string`

GetGitRepoSlug returns the GitRepoSlug field if non-nil, zero value otherwise.

### GetGitRepoSlugOk

`func (o *V0AppUploadParams) GetGitRepoSlugOk() (*string, bool)`

GetGitRepoSlugOk returns a tuple with the GitRepoSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitRepoSlug

`func (o *V0AppUploadParams) SetGitRepoSlug(v string)`

SetGitRepoSlug sets GitRepoSlug field to given value.

### HasGitRepoSlug

`func (o *V0AppUploadParams) HasGitRepoSlug() bool`

HasGitRepoSlug returns a boolean if a field has been set.

### GetIsPublic

`func (o *V0AppUploadParams) GetIsPublic() bool`

GetIsPublic returns the IsPublic field if non-nil, zero value otherwise.

### GetIsPublicOk

`func (o *V0AppUploadParams) GetIsPublicOk() (*bool, bool)`

GetIsPublicOk returns a tuple with the IsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublic

`func (o *V0AppUploadParams) SetIsPublic(v bool)`

SetIsPublic sets IsPublic field to given value.


### GetManualApprovalEnabled

`func (o *V0AppUploadParams) GetManualApprovalEnabled() bool`

GetManualApprovalEnabled returns the ManualApprovalEnabled field if non-nil, zero value otherwise.

### GetManualApprovalEnabledOk

`func (o *V0AppUploadParams) GetManualApprovalEnabledOk() (*bool, bool)`

GetManualApprovalEnabledOk returns a tuple with the ManualApprovalEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManualApprovalEnabled

`func (o *V0AppUploadParams) SetManualApprovalEnabled(v bool)`

SetManualApprovalEnabled sets ManualApprovalEnabled field to given value.

### HasManualApprovalEnabled

`func (o *V0AppUploadParams) HasManualApprovalEnabled() bool`

HasManualApprovalEnabled returns a boolean if a field has been set.

### GetOrganizationSlug

`func (o *V0AppUploadParams) GetOrganizationSlug() string`

GetOrganizationSlug returns the OrganizationSlug field if non-nil, zero value otherwise.

### GetOrganizationSlugOk

`func (o *V0AppUploadParams) GetOrganizationSlugOk() (*string, bool)`

GetOrganizationSlugOk returns a tuple with the OrganizationSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationSlug

`func (o *V0AppUploadParams) SetOrganizationSlug(v string)`

SetOrganizationSlug sets OrganizationSlug field to given value.

### HasOrganizationSlug

`func (o *V0AppUploadParams) HasOrganizationSlug() bool`

HasOrganizationSlug returns a boolean if a field has been set.

### GetProvider

`func (o *V0AppUploadParams) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *V0AppUploadParams) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *V0AppUploadParams) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *V0AppUploadParams) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetRepoUrl

`func (o *V0AppUploadParams) GetRepoUrl() string`

GetRepoUrl returns the RepoUrl field if non-nil, zero value otherwise.

### GetRepoUrlOk

`func (o *V0AppUploadParams) GetRepoUrlOk() (*string, bool)`

GetRepoUrlOk returns a tuple with the RepoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepoUrl

`func (o *V0AppUploadParams) SetRepoUrl(v string)`

SetRepoUrl sets RepoUrl field to given value.


### GetTitle

`func (o *V0AppUploadParams) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *V0AppUploadParams) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *V0AppUploadParams) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *V0AppUploadParams) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetType

`func (o *V0AppUploadParams) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *V0AppUploadParams) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *V0AppUploadParams) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *V0AppUploadParams) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


