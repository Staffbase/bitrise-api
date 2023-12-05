# V0BuildTriggerParamsBuildParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaseRepositoryUrl** | Pointer to **string** |  | [optional] 
**Branch** | Pointer to **string** |  | [optional] 
**BranchDest** | Pointer to **string** |  | [optional] 
**BranchDestRepoOwner** | Pointer to **string** |  | [optional] 
**BranchRepoOwner** | Pointer to **string** |  | [optional] 
**BuildRequestSlug** | Pointer to **string** |  | [optional] 
**CommitHash** | Pointer to **string** |  | [optional] 
**CommitMessage** | Pointer to **string** |  | [optional] 
**CommitPaths** | Pointer to [**[]V0CommitPaths**](V0CommitPaths.md) |  | [optional] 
**DiffUrl** | Pointer to **string** |  | [optional] 
**Environments** | Pointer to [**[]V0BuildParamsEnvironment**](V0BuildParamsEnvironment.md) |  | [optional] 
**HeadRepositoryUrl** | Pointer to **string** |  | [optional] 
**PipelineId** | Pointer to **string** |  | [optional] 
**PullRequestAuthor** | Pointer to **string** |  | [optional] 
**PullRequestHeadBranch** | Pointer to **string** |  | [optional] 
**PullRequestId** | Pointer to **map[string]interface{}** |  | [optional] 
**PullRequestMergeBranch** | Pointer to **string** |  | [optional] 
**PullRequestRepositoryUrl** | Pointer to **string** |  | [optional] 
**PullRequestUnverifiedMergeBranch** | Pointer to **string** |  | [optional] 
**SkipGitStatusReport** | Pointer to **bool** |  | [optional] 
**Tag** | Pointer to **string** |  | [optional] 
**WorkflowId** | Pointer to **string** |  | [optional] 

## Methods

### NewV0BuildTriggerParamsBuildParams

`func NewV0BuildTriggerParamsBuildParams() *V0BuildTriggerParamsBuildParams`

NewV0BuildTriggerParamsBuildParams instantiates a new V0BuildTriggerParamsBuildParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0BuildTriggerParamsBuildParamsWithDefaults

`func NewV0BuildTriggerParamsBuildParamsWithDefaults() *V0BuildTriggerParamsBuildParams`

NewV0BuildTriggerParamsBuildParamsWithDefaults instantiates a new V0BuildTriggerParamsBuildParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaseRepositoryUrl

`func (o *V0BuildTriggerParamsBuildParams) GetBaseRepositoryUrl() string`

GetBaseRepositoryUrl returns the BaseRepositoryUrl field if non-nil, zero value otherwise.

### GetBaseRepositoryUrlOk

`func (o *V0BuildTriggerParamsBuildParams) GetBaseRepositoryUrlOk() (*string, bool)`

GetBaseRepositoryUrlOk returns a tuple with the BaseRepositoryUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseRepositoryUrl

`func (o *V0BuildTriggerParamsBuildParams) SetBaseRepositoryUrl(v string)`

SetBaseRepositoryUrl sets BaseRepositoryUrl field to given value.

### HasBaseRepositoryUrl

`func (o *V0BuildTriggerParamsBuildParams) HasBaseRepositoryUrl() bool`

HasBaseRepositoryUrl returns a boolean if a field has been set.

### GetBranch

`func (o *V0BuildTriggerParamsBuildParams) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *V0BuildTriggerParamsBuildParams) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *V0BuildTriggerParamsBuildParams) SetBranch(v string)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *V0BuildTriggerParamsBuildParams) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### GetBranchDest

`func (o *V0BuildTriggerParamsBuildParams) GetBranchDest() string`

GetBranchDest returns the BranchDest field if non-nil, zero value otherwise.

### GetBranchDestOk

`func (o *V0BuildTriggerParamsBuildParams) GetBranchDestOk() (*string, bool)`

GetBranchDestOk returns a tuple with the BranchDest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchDest

`func (o *V0BuildTriggerParamsBuildParams) SetBranchDest(v string)`

SetBranchDest sets BranchDest field to given value.

### HasBranchDest

`func (o *V0BuildTriggerParamsBuildParams) HasBranchDest() bool`

HasBranchDest returns a boolean if a field has been set.

### GetBranchDestRepoOwner

`func (o *V0BuildTriggerParamsBuildParams) GetBranchDestRepoOwner() string`

GetBranchDestRepoOwner returns the BranchDestRepoOwner field if non-nil, zero value otherwise.

### GetBranchDestRepoOwnerOk

`func (o *V0BuildTriggerParamsBuildParams) GetBranchDestRepoOwnerOk() (*string, bool)`

GetBranchDestRepoOwnerOk returns a tuple with the BranchDestRepoOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchDestRepoOwner

`func (o *V0BuildTriggerParamsBuildParams) SetBranchDestRepoOwner(v string)`

SetBranchDestRepoOwner sets BranchDestRepoOwner field to given value.

### HasBranchDestRepoOwner

`func (o *V0BuildTriggerParamsBuildParams) HasBranchDestRepoOwner() bool`

HasBranchDestRepoOwner returns a boolean if a field has been set.

### GetBranchRepoOwner

`func (o *V0BuildTriggerParamsBuildParams) GetBranchRepoOwner() string`

GetBranchRepoOwner returns the BranchRepoOwner field if non-nil, zero value otherwise.

### GetBranchRepoOwnerOk

`func (o *V0BuildTriggerParamsBuildParams) GetBranchRepoOwnerOk() (*string, bool)`

GetBranchRepoOwnerOk returns a tuple with the BranchRepoOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchRepoOwner

`func (o *V0BuildTriggerParamsBuildParams) SetBranchRepoOwner(v string)`

SetBranchRepoOwner sets BranchRepoOwner field to given value.

### HasBranchRepoOwner

`func (o *V0BuildTriggerParamsBuildParams) HasBranchRepoOwner() bool`

HasBranchRepoOwner returns a boolean if a field has been set.

### GetBuildRequestSlug

`func (o *V0BuildTriggerParamsBuildParams) GetBuildRequestSlug() string`

GetBuildRequestSlug returns the BuildRequestSlug field if non-nil, zero value otherwise.

### GetBuildRequestSlugOk

`func (o *V0BuildTriggerParamsBuildParams) GetBuildRequestSlugOk() (*string, bool)`

GetBuildRequestSlugOk returns a tuple with the BuildRequestSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildRequestSlug

`func (o *V0BuildTriggerParamsBuildParams) SetBuildRequestSlug(v string)`

SetBuildRequestSlug sets BuildRequestSlug field to given value.

### HasBuildRequestSlug

`func (o *V0BuildTriggerParamsBuildParams) HasBuildRequestSlug() bool`

HasBuildRequestSlug returns a boolean if a field has been set.

### GetCommitHash

`func (o *V0BuildTriggerParamsBuildParams) GetCommitHash() string`

GetCommitHash returns the CommitHash field if non-nil, zero value otherwise.

### GetCommitHashOk

`func (o *V0BuildTriggerParamsBuildParams) GetCommitHashOk() (*string, bool)`

GetCommitHashOk returns a tuple with the CommitHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitHash

`func (o *V0BuildTriggerParamsBuildParams) SetCommitHash(v string)`

SetCommitHash sets CommitHash field to given value.

### HasCommitHash

`func (o *V0BuildTriggerParamsBuildParams) HasCommitHash() bool`

HasCommitHash returns a boolean if a field has been set.

### GetCommitMessage

`func (o *V0BuildTriggerParamsBuildParams) GetCommitMessage() string`

GetCommitMessage returns the CommitMessage field if non-nil, zero value otherwise.

### GetCommitMessageOk

`func (o *V0BuildTriggerParamsBuildParams) GetCommitMessageOk() (*string, bool)`

GetCommitMessageOk returns a tuple with the CommitMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitMessage

`func (o *V0BuildTriggerParamsBuildParams) SetCommitMessage(v string)`

SetCommitMessage sets CommitMessage field to given value.

### HasCommitMessage

`func (o *V0BuildTriggerParamsBuildParams) HasCommitMessage() bool`

HasCommitMessage returns a boolean if a field has been set.

### GetCommitPaths

`func (o *V0BuildTriggerParamsBuildParams) GetCommitPaths() []V0CommitPaths`

GetCommitPaths returns the CommitPaths field if non-nil, zero value otherwise.

### GetCommitPathsOk

`func (o *V0BuildTriggerParamsBuildParams) GetCommitPathsOk() (*[]V0CommitPaths, bool)`

GetCommitPathsOk returns a tuple with the CommitPaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitPaths

`func (o *V0BuildTriggerParamsBuildParams) SetCommitPaths(v []V0CommitPaths)`

SetCommitPaths sets CommitPaths field to given value.

### HasCommitPaths

`func (o *V0BuildTriggerParamsBuildParams) HasCommitPaths() bool`

HasCommitPaths returns a boolean if a field has been set.

### GetDiffUrl

`func (o *V0BuildTriggerParamsBuildParams) GetDiffUrl() string`

GetDiffUrl returns the DiffUrl field if non-nil, zero value otherwise.

### GetDiffUrlOk

`func (o *V0BuildTriggerParamsBuildParams) GetDiffUrlOk() (*string, bool)`

GetDiffUrlOk returns a tuple with the DiffUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiffUrl

`func (o *V0BuildTriggerParamsBuildParams) SetDiffUrl(v string)`

SetDiffUrl sets DiffUrl field to given value.

### HasDiffUrl

`func (o *V0BuildTriggerParamsBuildParams) HasDiffUrl() bool`

HasDiffUrl returns a boolean if a field has been set.

### GetEnvironments

`func (o *V0BuildTriggerParamsBuildParams) GetEnvironments() []V0BuildParamsEnvironment`

GetEnvironments returns the Environments field if non-nil, zero value otherwise.

### GetEnvironmentsOk

`func (o *V0BuildTriggerParamsBuildParams) GetEnvironmentsOk() (*[]V0BuildParamsEnvironment, bool)`

GetEnvironmentsOk returns a tuple with the Environments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironments

`func (o *V0BuildTriggerParamsBuildParams) SetEnvironments(v []V0BuildParamsEnvironment)`

SetEnvironments sets Environments field to given value.

### HasEnvironments

`func (o *V0BuildTriggerParamsBuildParams) HasEnvironments() bool`

HasEnvironments returns a boolean if a field has been set.

### GetHeadRepositoryUrl

`func (o *V0BuildTriggerParamsBuildParams) GetHeadRepositoryUrl() string`

GetHeadRepositoryUrl returns the HeadRepositoryUrl field if non-nil, zero value otherwise.

### GetHeadRepositoryUrlOk

`func (o *V0BuildTriggerParamsBuildParams) GetHeadRepositoryUrlOk() (*string, bool)`

GetHeadRepositoryUrlOk returns a tuple with the HeadRepositoryUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeadRepositoryUrl

`func (o *V0BuildTriggerParamsBuildParams) SetHeadRepositoryUrl(v string)`

SetHeadRepositoryUrl sets HeadRepositoryUrl field to given value.

### HasHeadRepositoryUrl

`func (o *V0BuildTriggerParamsBuildParams) HasHeadRepositoryUrl() bool`

HasHeadRepositoryUrl returns a boolean if a field has been set.

### GetPipelineId

`func (o *V0BuildTriggerParamsBuildParams) GetPipelineId() string`

GetPipelineId returns the PipelineId field if non-nil, zero value otherwise.

### GetPipelineIdOk

`func (o *V0BuildTriggerParamsBuildParams) GetPipelineIdOk() (*string, bool)`

GetPipelineIdOk returns a tuple with the PipelineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipelineId

`func (o *V0BuildTriggerParamsBuildParams) SetPipelineId(v string)`

SetPipelineId sets PipelineId field to given value.

### HasPipelineId

`func (o *V0BuildTriggerParamsBuildParams) HasPipelineId() bool`

HasPipelineId returns a boolean if a field has been set.

### GetPullRequestAuthor

`func (o *V0BuildTriggerParamsBuildParams) GetPullRequestAuthor() string`

GetPullRequestAuthor returns the PullRequestAuthor field if non-nil, zero value otherwise.

### GetPullRequestAuthorOk

`func (o *V0BuildTriggerParamsBuildParams) GetPullRequestAuthorOk() (*string, bool)`

GetPullRequestAuthorOk returns a tuple with the PullRequestAuthor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullRequestAuthor

`func (o *V0BuildTriggerParamsBuildParams) SetPullRequestAuthor(v string)`

SetPullRequestAuthor sets PullRequestAuthor field to given value.

### HasPullRequestAuthor

`func (o *V0BuildTriggerParamsBuildParams) HasPullRequestAuthor() bool`

HasPullRequestAuthor returns a boolean if a field has been set.

### GetPullRequestHeadBranch

`func (o *V0BuildTriggerParamsBuildParams) GetPullRequestHeadBranch() string`

GetPullRequestHeadBranch returns the PullRequestHeadBranch field if non-nil, zero value otherwise.

### GetPullRequestHeadBranchOk

`func (o *V0BuildTriggerParamsBuildParams) GetPullRequestHeadBranchOk() (*string, bool)`

GetPullRequestHeadBranchOk returns a tuple with the PullRequestHeadBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullRequestHeadBranch

`func (o *V0BuildTriggerParamsBuildParams) SetPullRequestHeadBranch(v string)`

SetPullRequestHeadBranch sets PullRequestHeadBranch field to given value.

### HasPullRequestHeadBranch

`func (o *V0BuildTriggerParamsBuildParams) HasPullRequestHeadBranch() bool`

HasPullRequestHeadBranch returns a boolean if a field has been set.

### GetPullRequestId

`func (o *V0BuildTriggerParamsBuildParams) GetPullRequestId() map[string]interface{}`

GetPullRequestId returns the PullRequestId field if non-nil, zero value otherwise.

### GetPullRequestIdOk

`func (o *V0BuildTriggerParamsBuildParams) GetPullRequestIdOk() (*map[string]interface{}, bool)`

GetPullRequestIdOk returns a tuple with the PullRequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullRequestId

`func (o *V0BuildTriggerParamsBuildParams) SetPullRequestId(v map[string]interface{})`

SetPullRequestId sets PullRequestId field to given value.

### HasPullRequestId

`func (o *V0BuildTriggerParamsBuildParams) HasPullRequestId() bool`

HasPullRequestId returns a boolean if a field has been set.

### GetPullRequestMergeBranch

`func (o *V0BuildTriggerParamsBuildParams) GetPullRequestMergeBranch() string`

GetPullRequestMergeBranch returns the PullRequestMergeBranch field if non-nil, zero value otherwise.

### GetPullRequestMergeBranchOk

`func (o *V0BuildTriggerParamsBuildParams) GetPullRequestMergeBranchOk() (*string, bool)`

GetPullRequestMergeBranchOk returns a tuple with the PullRequestMergeBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullRequestMergeBranch

`func (o *V0BuildTriggerParamsBuildParams) SetPullRequestMergeBranch(v string)`

SetPullRequestMergeBranch sets PullRequestMergeBranch field to given value.

### HasPullRequestMergeBranch

`func (o *V0BuildTriggerParamsBuildParams) HasPullRequestMergeBranch() bool`

HasPullRequestMergeBranch returns a boolean if a field has been set.

### GetPullRequestRepositoryUrl

`func (o *V0BuildTriggerParamsBuildParams) GetPullRequestRepositoryUrl() string`

GetPullRequestRepositoryUrl returns the PullRequestRepositoryUrl field if non-nil, zero value otherwise.

### GetPullRequestRepositoryUrlOk

`func (o *V0BuildTriggerParamsBuildParams) GetPullRequestRepositoryUrlOk() (*string, bool)`

GetPullRequestRepositoryUrlOk returns a tuple with the PullRequestRepositoryUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullRequestRepositoryUrl

`func (o *V0BuildTriggerParamsBuildParams) SetPullRequestRepositoryUrl(v string)`

SetPullRequestRepositoryUrl sets PullRequestRepositoryUrl field to given value.

### HasPullRequestRepositoryUrl

`func (o *V0BuildTriggerParamsBuildParams) HasPullRequestRepositoryUrl() bool`

HasPullRequestRepositoryUrl returns a boolean if a field has been set.

### GetPullRequestUnverifiedMergeBranch

`func (o *V0BuildTriggerParamsBuildParams) GetPullRequestUnverifiedMergeBranch() string`

GetPullRequestUnverifiedMergeBranch returns the PullRequestUnverifiedMergeBranch field if non-nil, zero value otherwise.

### GetPullRequestUnverifiedMergeBranchOk

`func (o *V0BuildTriggerParamsBuildParams) GetPullRequestUnverifiedMergeBranchOk() (*string, bool)`

GetPullRequestUnverifiedMergeBranchOk returns a tuple with the PullRequestUnverifiedMergeBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullRequestUnverifiedMergeBranch

`func (o *V0BuildTriggerParamsBuildParams) SetPullRequestUnverifiedMergeBranch(v string)`

SetPullRequestUnverifiedMergeBranch sets PullRequestUnverifiedMergeBranch field to given value.

### HasPullRequestUnverifiedMergeBranch

`func (o *V0BuildTriggerParamsBuildParams) HasPullRequestUnverifiedMergeBranch() bool`

HasPullRequestUnverifiedMergeBranch returns a boolean if a field has been set.

### GetSkipGitStatusReport

`func (o *V0BuildTriggerParamsBuildParams) GetSkipGitStatusReport() bool`

GetSkipGitStatusReport returns the SkipGitStatusReport field if non-nil, zero value otherwise.

### GetSkipGitStatusReportOk

`func (o *V0BuildTriggerParamsBuildParams) GetSkipGitStatusReportOk() (*bool, bool)`

GetSkipGitStatusReportOk returns a tuple with the SkipGitStatusReport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipGitStatusReport

`func (o *V0BuildTriggerParamsBuildParams) SetSkipGitStatusReport(v bool)`

SetSkipGitStatusReport sets SkipGitStatusReport field to given value.

### HasSkipGitStatusReport

`func (o *V0BuildTriggerParamsBuildParams) HasSkipGitStatusReport() bool`

HasSkipGitStatusReport returns a boolean if a field has been set.

### GetTag

`func (o *V0BuildTriggerParamsBuildParams) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *V0BuildTriggerParamsBuildParams) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *V0BuildTriggerParamsBuildParams) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *V0BuildTriggerParamsBuildParams) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetWorkflowId

`func (o *V0BuildTriggerParamsBuildParams) GetWorkflowId() string`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *V0BuildTriggerParamsBuildParams) GetWorkflowIdOk() (*string, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *V0BuildTriggerParamsBuildParams) SetWorkflowId(v string)`

SetWorkflowId sets WorkflowId field to given value.

### HasWorkflowId

`func (o *V0BuildTriggerParamsBuildParams) HasWorkflowId() bool`

HasWorkflowId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


