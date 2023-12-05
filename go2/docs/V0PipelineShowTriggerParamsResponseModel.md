# V0PipelineShowTriggerParamsResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Branch** | Pointer to **string** |  | [optional] 
**BranchDest** | Pointer to **string** |  | [optional] 
**BranchDestRepoOwner** | Pointer to **string** |  | [optional] 
**BranchRepoOwner** | Pointer to **string** |  | [optional] 
**CommitHash** | Pointer to **string** |  | [optional] 
**CommitMessage** | Pointer to **string** |  | [optional] 
**Environments** | Pointer to [**[]V0PipelineShowEnvironmentsResponseModel**](V0PipelineShowEnvironmentsResponseModel.md) |  | [optional] 
**PullRequestAuthor** | Pointer to **string** |  | [optional] 
**PullRequestHeadBranch** | Pointer to **string** |  | [optional] 
**PullRequestId** | Pointer to **string** |  | [optional] 
**PullRequestMergeBranch** | Pointer to **string** |  | [optional] 
**PullRequestRepositoryUrl** | Pointer to **string** |  | [optional] 
**PullRequestUnverifiedMergeBranch** | Pointer to **string** |  | [optional] 
**Tag** | Pointer to **string** |  | [optional] 

## Methods

### NewV0PipelineShowTriggerParamsResponseModel

`func NewV0PipelineShowTriggerParamsResponseModel() *V0PipelineShowTriggerParamsResponseModel`

NewV0PipelineShowTriggerParamsResponseModel instantiates a new V0PipelineShowTriggerParamsResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0PipelineShowTriggerParamsResponseModelWithDefaults

`func NewV0PipelineShowTriggerParamsResponseModelWithDefaults() *V0PipelineShowTriggerParamsResponseModel`

NewV0PipelineShowTriggerParamsResponseModelWithDefaults instantiates a new V0PipelineShowTriggerParamsResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBranch

`func (o *V0PipelineShowTriggerParamsResponseModel) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *V0PipelineShowTriggerParamsResponseModel) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *V0PipelineShowTriggerParamsResponseModel) SetBranch(v string)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *V0PipelineShowTriggerParamsResponseModel) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### GetBranchDest

`func (o *V0PipelineShowTriggerParamsResponseModel) GetBranchDest() string`

GetBranchDest returns the BranchDest field if non-nil, zero value otherwise.

### GetBranchDestOk

`func (o *V0PipelineShowTriggerParamsResponseModel) GetBranchDestOk() (*string, bool)`

GetBranchDestOk returns a tuple with the BranchDest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchDest

`func (o *V0PipelineShowTriggerParamsResponseModel) SetBranchDest(v string)`

SetBranchDest sets BranchDest field to given value.

### HasBranchDest

`func (o *V0PipelineShowTriggerParamsResponseModel) HasBranchDest() bool`

HasBranchDest returns a boolean if a field has been set.

### GetBranchDestRepoOwner

`func (o *V0PipelineShowTriggerParamsResponseModel) GetBranchDestRepoOwner() string`

GetBranchDestRepoOwner returns the BranchDestRepoOwner field if non-nil, zero value otherwise.

### GetBranchDestRepoOwnerOk

`func (o *V0PipelineShowTriggerParamsResponseModel) GetBranchDestRepoOwnerOk() (*string, bool)`

GetBranchDestRepoOwnerOk returns a tuple with the BranchDestRepoOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchDestRepoOwner

`func (o *V0PipelineShowTriggerParamsResponseModel) SetBranchDestRepoOwner(v string)`

SetBranchDestRepoOwner sets BranchDestRepoOwner field to given value.

### HasBranchDestRepoOwner

`func (o *V0PipelineShowTriggerParamsResponseModel) HasBranchDestRepoOwner() bool`

HasBranchDestRepoOwner returns a boolean if a field has been set.

### GetBranchRepoOwner

`func (o *V0PipelineShowTriggerParamsResponseModel) GetBranchRepoOwner() string`

GetBranchRepoOwner returns the BranchRepoOwner field if non-nil, zero value otherwise.

### GetBranchRepoOwnerOk

`func (o *V0PipelineShowTriggerParamsResponseModel) GetBranchRepoOwnerOk() (*string, bool)`

GetBranchRepoOwnerOk returns a tuple with the BranchRepoOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchRepoOwner

`func (o *V0PipelineShowTriggerParamsResponseModel) SetBranchRepoOwner(v string)`

SetBranchRepoOwner sets BranchRepoOwner field to given value.

### HasBranchRepoOwner

`func (o *V0PipelineShowTriggerParamsResponseModel) HasBranchRepoOwner() bool`

HasBranchRepoOwner returns a boolean if a field has been set.

### GetCommitHash

`func (o *V0PipelineShowTriggerParamsResponseModel) GetCommitHash() string`

GetCommitHash returns the CommitHash field if non-nil, zero value otherwise.

### GetCommitHashOk

`func (o *V0PipelineShowTriggerParamsResponseModel) GetCommitHashOk() (*string, bool)`

GetCommitHashOk returns a tuple with the CommitHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitHash

`func (o *V0PipelineShowTriggerParamsResponseModel) SetCommitHash(v string)`

SetCommitHash sets CommitHash field to given value.

### HasCommitHash

`func (o *V0PipelineShowTriggerParamsResponseModel) HasCommitHash() bool`

HasCommitHash returns a boolean if a field has been set.

### GetCommitMessage

`func (o *V0PipelineShowTriggerParamsResponseModel) GetCommitMessage() string`

GetCommitMessage returns the CommitMessage field if non-nil, zero value otherwise.

### GetCommitMessageOk

`func (o *V0PipelineShowTriggerParamsResponseModel) GetCommitMessageOk() (*string, bool)`

GetCommitMessageOk returns a tuple with the CommitMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitMessage

`func (o *V0PipelineShowTriggerParamsResponseModel) SetCommitMessage(v string)`

SetCommitMessage sets CommitMessage field to given value.

### HasCommitMessage

`func (o *V0PipelineShowTriggerParamsResponseModel) HasCommitMessage() bool`

HasCommitMessage returns a boolean if a field has been set.

### GetEnvironments

`func (o *V0PipelineShowTriggerParamsResponseModel) GetEnvironments() []V0PipelineShowEnvironmentsResponseModel`

GetEnvironments returns the Environments field if non-nil, zero value otherwise.

### GetEnvironmentsOk

`func (o *V0PipelineShowTriggerParamsResponseModel) GetEnvironmentsOk() (*[]V0PipelineShowEnvironmentsResponseModel, bool)`

GetEnvironmentsOk returns a tuple with the Environments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironments

`func (o *V0PipelineShowTriggerParamsResponseModel) SetEnvironments(v []V0PipelineShowEnvironmentsResponseModel)`

SetEnvironments sets Environments field to given value.

### HasEnvironments

`func (o *V0PipelineShowTriggerParamsResponseModel) HasEnvironments() bool`

HasEnvironments returns a boolean if a field has been set.

### GetPullRequestAuthor

`func (o *V0PipelineShowTriggerParamsResponseModel) GetPullRequestAuthor() string`

GetPullRequestAuthor returns the PullRequestAuthor field if non-nil, zero value otherwise.

### GetPullRequestAuthorOk

`func (o *V0PipelineShowTriggerParamsResponseModel) GetPullRequestAuthorOk() (*string, bool)`

GetPullRequestAuthorOk returns a tuple with the PullRequestAuthor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullRequestAuthor

`func (o *V0PipelineShowTriggerParamsResponseModel) SetPullRequestAuthor(v string)`

SetPullRequestAuthor sets PullRequestAuthor field to given value.

### HasPullRequestAuthor

`func (o *V0PipelineShowTriggerParamsResponseModel) HasPullRequestAuthor() bool`

HasPullRequestAuthor returns a boolean if a field has been set.

### GetPullRequestHeadBranch

`func (o *V0PipelineShowTriggerParamsResponseModel) GetPullRequestHeadBranch() string`

GetPullRequestHeadBranch returns the PullRequestHeadBranch field if non-nil, zero value otherwise.

### GetPullRequestHeadBranchOk

`func (o *V0PipelineShowTriggerParamsResponseModel) GetPullRequestHeadBranchOk() (*string, bool)`

GetPullRequestHeadBranchOk returns a tuple with the PullRequestHeadBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullRequestHeadBranch

`func (o *V0PipelineShowTriggerParamsResponseModel) SetPullRequestHeadBranch(v string)`

SetPullRequestHeadBranch sets PullRequestHeadBranch field to given value.

### HasPullRequestHeadBranch

`func (o *V0PipelineShowTriggerParamsResponseModel) HasPullRequestHeadBranch() bool`

HasPullRequestHeadBranch returns a boolean if a field has been set.

### GetPullRequestId

`func (o *V0PipelineShowTriggerParamsResponseModel) GetPullRequestId() string`

GetPullRequestId returns the PullRequestId field if non-nil, zero value otherwise.

### GetPullRequestIdOk

`func (o *V0PipelineShowTriggerParamsResponseModel) GetPullRequestIdOk() (*string, bool)`

GetPullRequestIdOk returns a tuple with the PullRequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullRequestId

`func (o *V0PipelineShowTriggerParamsResponseModel) SetPullRequestId(v string)`

SetPullRequestId sets PullRequestId field to given value.

### HasPullRequestId

`func (o *V0PipelineShowTriggerParamsResponseModel) HasPullRequestId() bool`

HasPullRequestId returns a boolean if a field has been set.

### GetPullRequestMergeBranch

`func (o *V0PipelineShowTriggerParamsResponseModel) GetPullRequestMergeBranch() string`

GetPullRequestMergeBranch returns the PullRequestMergeBranch field if non-nil, zero value otherwise.

### GetPullRequestMergeBranchOk

`func (o *V0PipelineShowTriggerParamsResponseModel) GetPullRequestMergeBranchOk() (*string, bool)`

GetPullRequestMergeBranchOk returns a tuple with the PullRequestMergeBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullRequestMergeBranch

`func (o *V0PipelineShowTriggerParamsResponseModel) SetPullRequestMergeBranch(v string)`

SetPullRequestMergeBranch sets PullRequestMergeBranch field to given value.

### HasPullRequestMergeBranch

`func (o *V0PipelineShowTriggerParamsResponseModel) HasPullRequestMergeBranch() bool`

HasPullRequestMergeBranch returns a boolean if a field has been set.

### GetPullRequestRepositoryUrl

`func (o *V0PipelineShowTriggerParamsResponseModel) GetPullRequestRepositoryUrl() string`

GetPullRequestRepositoryUrl returns the PullRequestRepositoryUrl field if non-nil, zero value otherwise.

### GetPullRequestRepositoryUrlOk

`func (o *V0PipelineShowTriggerParamsResponseModel) GetPullRequestRepositoryUrlOk() (*string, bool)`

GetPullRequestRepositoryUrlOk returns a tuple with the PullRequestRepositoryUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullRequestRepositoryUrl

`func (o *V0PipelineShowTriggerParamsResponseModel) SetPullRequestRepositoryUrl(v string)`

SetPullRequestRepositoryUrl sets PullRequestRepositoryUrl field to given value.

### HasPullRequestRepositoryUrl

`func (o *V0PipelineShowTriggerParamsResponseModel) HasPullRequestRepositoryUrl() bool`

HasPullRequestRepositoryUrl returns a boolean if a field has been set.

### GetPullRequestUnverifiedMergeBranch

`func (o *V0PipelineShowTriggerParamsResponseModel) GetPullRequestUnverifiedMergeBranch() string`

GetPullRequestUnverifiedMergeBranch returns the PullRequestUnverifiedMergeBranch field if non-nil, zero value otherwise.

### GetPullRequestUnverifiedMergeBranchOk

`func (o *V0PipelineShowTriggerParamsResponseModel) GetPullRequestUnverifiedMergeBranchOk() (*string, bool)`

GetPullRequestUnverifiedMergeBranchOk returns a tuple with the PullRequestUnverifiedMergeBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullRequestUnverifiedMergeBranch

`func (o *V0PipelineShowTriggerParamsResponseModel) SetPullRequestUnverifiedMergeBranch(v string)`

SetPullRequestUnverifiedMergeBranch sets PullRequestUnverifiedMergeBranch field to given value.

### HasPullRequestUnverifiedMergeBranch

`func (o *V0PipelineShowTriggerParamsResponseModel) HasPullRequestUnverifiedMergeBranch() bool`

HasPullRequestUnverifiedMergeBranch returns a boolean if a field has been set.

### GetTag

`func (o *V0PipelineShowTriggerParamsResponseModel) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *V0PipelineShowTriggerParamsResponseModel) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *V0PipelineShowTriggerParamsResponseModel) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *V0PipelineShowTriggerParamsResponseModel) HasTag() bool`

HasTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


