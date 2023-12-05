# PipelineabledomainTriggerParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Branch** | Pointer to **string** |  | [optional] 
**BranchDest** | Pointer to **string** |  | [optional] 
**BranchDestRepoOwner** | Pointer to **string** |  | [optional] 
**BranchRepoOwner** | Pointer to **string** |  | [optional] 
**CommitHash** | Pointer to **string** |  | [optional] 
**CommitMessage** | Pointer to **string** |  | [optional] 
**Environments** | Pointer to [**[]PipelineabledomainEnvironments**](PipelineabledomainEnvironments.md) |  | [optional] 
**PullRequestAuthor** | Pointer to **string** |  | [optional] 
**PullRequestHeadBranch** | Pointer to **string** |  | [optional] 
**PullRequestId** | Pointer to **map[string]interface{}** |  | [optional] 
**PullRequestMergeBranch** | Pointer to **string** |  | [optional] 
**PullRequestRepositoryUrl** | Pointer to **string** |  | [optional] 
**PullRequestUnverifiedMergeBranch** | Pointer to **string** |  | [optional] 
**Tag** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewPipelineabledomainTriggerParams

`func NewPipelineabledomainTriggerParams() *PipelineabledomainTriggerParams`

NewPipelineabledomainTriggerParams instantiates a new PipelineabledomainTriggerParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPipelineabledomainTriggerParamsWithDefaults

`func NewPipelineabledomainTriggerParamsWithDefaults() *PipelineabledomainTriggerParams`

NewPipelineabledomainTriggerParamsWithDefaults instantiates a new PipelineabledomainTriggerParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBranch

`func (o *PipelineabledomainTriggerParams) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *PipelineabledomainTriggerParams) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *PipelineabledomainTriggerParams) SetBranch(v string)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *PipelineabledomainTriggerParams) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### GetBranchDest

`func (o *PipelineabledomainTriggerParams) GetBranchDest() string`

GetBranchDest returns the BranchDest field if non-nil, zero value otherwise.

### GetBranchDestOk

`func (o *PipelineabledomainTriggerParams) GetBranchDestOk() (*string, bool)`

GetBranchDestOk returns a tuple with the BranchDest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchDest

`func (o *PipelineabledomainTriggerParams) SetBranchDest(v string)`

SetBranchDest sets BranchDest field to given value.

### HasBranchDest

`func (o *PipelineabledomainTriggerParams) HasBranchDest() bool`

HasBranchDest returns a boolean if a field has been set.

### GetBranchDestRepoOwner

`func (o *PipelineabledomainTriggerParams) GetBranchDestRepoOwner() string`

GetBranchDestRepoOwner returns the BranchDestRepoOwner field if non-nil, zero value otherwise.

### GetBranchDestRepoOwnerOk

`func (o *PipelineabledomainTriggerParams) GetBranchDestRepoOwnerOk() (*string, bool)`

GetBranchDestRepoOwnerOk returns a tuple with the BranchDestRepoOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchDestRepoOwner

`func (o *PipelineabledomainTriggerParams) SetBranchDestRepoOwner(v string)`

SetBranchDestRepoOwner sets BranchDestRepoOwner field to given value.

### HasBranchDestRepoOwner

`func (o *PipelineabledomainTriggerParams) HasBranchDestRepoOwner() bool`

HasBranchDestRepoOwner returns a boolean if a field has been set.

### GetBranchRepoOwner

`func (o *PipelineabledomainTriggerParams) GetBranchRepoOwner() string`

GetBranchRepoOwner returns the BranchRepoOwner field if non-nil, zero value otherwise.

### GetBranchRepoOwnerOk

`func (o *PipelineabledomainTriggerParams) GetBranchRepoOwnerOk() (*string, bool)`

GetBranchRepoOwnerOk returns a tuple with the BranchRepoOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchRepoOwner

`func (o *PipelineabledomainTriggerParams) SetBranchRepoOwner(v string)`

SetBranchRepoOwner sets BranchRepoOwner field to given value.

### HasBranchRepoOwner

`func (o *PipelineabledomainTriggerParams) HasBranchRepoOwner() bool`

HasBranchRepoOwner returns a boolean if a field has been set.

### GetCommitHash

`func (o *PipelineabledomainTriggerParams) GetCommitHash() string`

GetCommitHash returns the CommitHash field if non-nil, zero value otherwise.

### GetCommitHashOk

`func (o *PipelineabledomainTriggerParams) GetCommitHashOk() (*string, bool)`

GetCommitHashOk returns a tuple with the CommitHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitHash

`func (o *PipelineabledomainTriggerParams) SetCommitHash(v string)`

SetCommitHash sets CommitHash field to given value.

### HasCommitHash

`func (o *PipelineabledomainTriggerParams) HasCommitHash() bool`

HasCommitHash returns a boolean if a field has been set.

### GetCommitMessage

`func (o *PipelineabledomainTriggerParams) GetCommitMessage() string`

GetCommitMessage returns the CommitMessage field if non-nil, zero value otherwise.

### GetCommitMessageOk

`func (o *PipelineabledomainTriggerParams) GetCommitMessageOk() (*string, bool)`

GetCommitMessageOk returns a tuple with the CommitMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitMessage

`func (o *PipelineabledomainTriggerParams) SetCommitMessage(v string)`

SetCommitMessage sets CommitMessage field to given value.

### HasCommitMessage

`func (o *PipelineabledomainTriggerParams) HasCommitMessage() bool`

HasCommitMessage returns a boolean if a field has been set.

### GetEnvironments

`func (o *PipelineabledomainTriggerParams) GetEnvironments() []PipelineabledomainEnvironments`

GetEnvironments returns the Environments field if non-nil, zero value otherwise.

### GetEnvironmentsOk

`func (o *PipelineabledomainTriggerParams) GetEnvironmentsOk() (*[]PipelineabledomainEnvironments, bool)`

GetEnvironmentsOk returns a tuple with the Environments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironments

`func (o *PipelineabledomainTriggerParams) SetEnvironments(v []PipelineabledomainEnvironments)`

SetEnvironments sets Environments field to given value.

### HasEnvironments

`func (o *PipelineabledomainTriggerParams) HasEnvironments() bool`

HasEnvironments returns a boolean if a field has been set.

### GetPullRequestAuthor

`func (o *PipelineabledomainTriggerParams) GetPullRequestAuthor() string`

GetPullRequestAuthor returns the PullRequestAuthor field if non-nil, zero value otherwise.

### GetPullRequestAuthorOk

`func (o *PipelineabledomainTriggerParams) GetPullRequestAuthorOk() (*string, bool)`

GetPullRequestAuthorOk returns a tuple with the PullRequestAuthor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullRequestAuthor

`func (o *PipelineabledomainTriggerParams) SetPullRequestAuthor(v string)`

SetPullRequestAuthor sets PullRequestAuthor field to given value.

### HasPullRequestAuthor

`func (o *PipelineabledomainTriggerParams) HasPullRequestAuthor() bool`

HasPullRequestAuthor returns a boolean if a field has been set.

### GetPullRequestHeadBranch

`func (o *PipelineabledomainTriggerParams) GetPullRequestHeadBranch() string`

GetPullRequestHeadBranch returns the PullRequestHeadBranch field if non-nil, zero value otherwise.

### GetPullRequestHeadBranchOk

`func (o *PipelineabledomainTriggerParams) GetPullRequestHeadBranchOk() (*string, bool)`

GetPullRequestHeadBranchOk returns a tuple with the PullRequestHeadBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullRequestHeadBranch

`func (o *PipelineabledomainTriggerParams) SetPullRequestHeadBranch(v string)`

SetPullRequestHeadBranch sets PullRequestHeadBranch field to given value.

### HasPullRequestHeadBranch

`func (o *PipelineabledomainTriggerParams) HasPullRequestHeadBranch() bool`

HasPullRequestHeadBranch returns a boolean if a field has been set.

### GetPullRequestId

`func (o *PipelineabledomainTriggerParams) GetPullRequestId() map[string]interface{}`

GetPullRequestId returns the PullRequestId field if non-nil, zero value otherwise.

### GetPullRequestIdOk

`func (o *PipelineabledomainTriggerParams) GetPullRequestIdOk() (*map[string]interface{}, bool)`

GetPullRequestIdOk returns a tuple with the PullRequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullRequestId

`func (o *PipelineabledomainTriggerParams) SetPullRequestId(v map[string]interface{})`

SetPullRequestId sets PullRequestId field to given value.

### HasPullRequestId

`func (o *PipelineabledomainTriggerParams) HasPullRequestId() bool`

HasPullRequestId returns a boolean if a field has been set.

### GetPullRequestMergeBranch

`func (o *PipelineabledomainTriggerParams) GetPullRequestMergeBranch() string`

GetPullRequestMergeBranch returns the PullRequestMergeBranch field if non-nil, zero value otherwise.

### GetPullRequestMergeBranchOk

`func (o *PipelineabledomainTriggerParams) GetPullRequestMergeBranchOk() (*string, bool)`

GetPullRequestMergeBranchOk returns a tuple with the PullRequestMergeBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullRequestMergeBranch

`func (o *PipelineabledomainTriggerParams) SetPullRequestMergeBranch(v string)`

SetPullRequestMergeBranch sets PullRequestMergeBranch field to given value.

### HasPullRequestMergeBranch

`func (o *PipelineabledomainTriggerParams) HasPullRequestMergeBranch() bool`

HasPullRequestMergeBranch returns a boolean if a field has been set.

### GetPullRequestRepositoryUrl

`func (o *PipelineabledomainTriggerParams) GetPullRequestRepositoryUrl() string`

GetPullRequestRepositoryUrl returns the PullRequestRepositoryUrl field if non-nil, zero value otherwise.

### GetPullRequestRepositoryUrlOk

`func (o *PipelineabledomainTriggerParams) GetPullRequestRepositoryUrlOk() (*string, bool)`

GetPullRequestRepositoryUrlOk returns a tuple with the PullRequestRepositoryUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullRequestRepositoryUrl

`func (o *PipelineabledomainTriggerParams) SetPullRequestRepositoryUrl(v string)`

SetPullRequestRepositoryUrl sets PullRequestRepositoryUrl field to given value.

### HasPullRequestRepositoryUrl

`func (o *PipelineabledomainTriggerParams) HasPullRequestRepositoryUrl() bool`

HasPullRequestRepositoryUrl returns a boolean if a field has been set.

### GetPullRequestUnverifiedMergeBranch

`func (o *PipelineabledomainTriggerParams) GetPullRequestUnverifiedMergeBranch() string`

GetPullRequestUnverifiedMergeBranch returns the PullRequestUnverifiedMergeBranch field if non-nil, zero value otherwise.

### GetPullRequestUnverifiedMergeBranchOk

`func (o *PipelineabledomainTriggerParams) GetPullRequestUnverifiedMergeBranchOk() (*string, bool)`

GetPullRequestUnverifiedMergeBranchOk returns a tuple with the PullRequestUnverifiedMergeBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullRequestUnverifiedMergeBranch

`func (o *PipelineabledomainTriggerParams) SetPullRequestUnverifiedMergeBranch(v string)`

SetPullRequestUnverifiedMergeBranch sets PullRequestUnverifiedMergeBranch field to given value.

### HasPullRequestUnverifiedMergeBranch

`func (o *PipelineabledomainTriggerParams) HasPullRequestUnverifiedMergeBranch() bool`

HasPullRequestUnverifiedMergeBranch returns a boolean if a field has been set.

### GetTag

`func (o *PipelineabledomainTriggerParams) GetTag() map[string]interface{}`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *PipelineabledomainTriggerParams) GetTagOk() (*map[string]interface{}, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *PipelineabledomainTriggerParams) SetTag(v map[string]interface{})`

SetTag sets Tag field to given value.

### HasTag

`func (o *PipelineabledomainTriggerParams) HasTag() bool`

HasTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


