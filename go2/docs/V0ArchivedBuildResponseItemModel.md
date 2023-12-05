# V0ArchivedBuildResponseItemModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AbortReason** | Pointer to **string** |  | [optional] 
**Branch** | Pointer to **string** |  | [optional] 
**BuildArtifacts** | Pointer to [**[]V0ArchivedBuildArtifact**](V0ArchivedBuildArtifact.md) |  | [optional] 
**BuildNumber** | Pointer to **int32** |  | [optional] 
**CommitHash** | Pointer to **string** |  | [optional] 
**CommitMessage** | Pointer to **string** |  | [optional] 
**CreditCost** | Pointer to **int32** |  | [optional] 
**FinishedAt** | Pointer to **string** |  | [optional] 
**MachineTypeId** | Pointer to **string** |  | [optional] 
**OriginalBuildParams** | Pointer to **[]int32** |  | [optional] 
**PipelineWorkflowId** | Pointer to **string** |  | [optional] 
**PullRequestId** | Pointer to **int32** |  | [optional] 
**PullRequestTargetBranch** | Pointer to **string** |  | [optional] 
**PullRequestViewUrl** | Pointer to **string** |  | [optional] 
**Slug** | Pointer to **string** |  | [optional] 
**StackIdentifier** | Pointer to **string** |  | [optional] 
**StartedOnWorkerAt** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **int32** |  | [optional] 
**StatusText** | Pointer to **string** |  | [optional] 
**Tag** | Pointer to **string** |  | [optional] 
**TriggeredBy** | Pointer to **string** |  | [optional] 
**TriggeredWorkflow** | Pointer to **string** |  | [optional] 

## Methods

### NewV0ArchivedBuildResponseItemModel

`func NewV0ArchivedBuildResponseItemModel() *V0ArchivedBuildResponseItemModel`

NewV0ArchivedBuildResponseItemModel instantiates a new V0ArchivedBuildResponseItemModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0ArchivedBuildResponseItemModelWithDefaults

`func NewV0ArchivedBuildResponseItemModelWithDefaults() *V0ArchivedBuildResponseItemModel`

NewV0ArchivedBuildResponseItemModelWithDefaults instantiates a new V0ArchivedBuildResponseItemModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAbortReason

`func (o *V0ArchivedBuildResponseItemModel) GetAbortReason() string`

GetAbortReason returns the AbortReason field if non-nil, zero value otherwise.

### GetAbortReasonOk

`func (o *V0ArchivedBuildResponseItemModel) GetAbortReasonOk() (*string, bool)`

GetAbortReasonOk returns a tuple with the AbortReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbortReason

`func (o *V0ArchivedBuildResponseItemModel) SetAbortReason(v string)`

SetAbortReason sets AbortReason field to given value.

### HasAbortReason

`func (o *V0ArchivedBuildResponseItemModel) HasAbortReason() bool`

HasAbortReason returns a boolean if a field has been set.

### GetBranch

`func (o *V0ArchivedBuildResponseItemModel) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *V0ArchivedBuildResponseItemModel) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *V0ArchivedBuildResponseItemModel) SetBranch(v string)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *V0ArchivedBuildResponseItemModel) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### GetBuildArtifacts

`func (o *V0ArchivedBuildResponseItemModel) GetBuildArtifacts() []V0ArchivedBuildArtifact`

GetBuildArtifacts returns the BuildArtifacts field if non-nil, zero value otherwise.

### GetBuildArtifactsOk

`func (o *V0ArchivedBuildResponseItemModel) GetBuildArtifactsOk() (*[]V0ArchivedBuildArtifact, bool)`

GetBuildArtifactsOk returns a tuple with the BuildArtifacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildArtifacts

`func (o *V0ArchivedBuildResponseItemModel) SetBuildArtifacts(v []V0ArchivedBuildArtifact)`

SetBuildArtifacts sets BuildArtifacts field to given value.

### HasBuildArtifacts

`func (o *V0ArchivedBuildResponseItemModel) HasBuildArtifacts() bool`

HasBuildArtifacts returns a boolean if a field has been set.

### GetBuildNumber

`func (o *V0ArchivedBuildResponseItemModel) GetBuildNumber() int32`

GetBuildNumber returns the BuildNumber field if non-nil, zero value otherwise.

### GetBuildNumberOk

`func (o *V0ArchivedBuildResponseItemModel) GetBuildNumberOk() (*int32, bool)`

GetBuildNumberOk returns a tuple with the BuildNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildNumber

`func (o *V0ArchivedBuildResponseItemModel) SetBuildNumber(v int32)`

SetBuildNumber sets BuildNumber field to given value.

### HasBuildNumber

`func (o *V0ArchivedBuildResponseItemModel) HasBuildNumber() bool`

HasBuildNumber returns a boolean if a field has been set.

### GetCommitHash

`func (o *V0ArchivedBuildResponseItemModel) GetCommitHash() string`

GetCommitHash returns the CommitHash field if non-nil, zero value otherwise.

### GetCommitHashOk

`func (o *V0ArchivedBuildResponseItemModel) GetCommitHashOk() (*string, bool)`

GetCommitHashOk returns a tuple with the CommitHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitHash

`func (o *V0ArchivedBuildResponseItemModel) SetCommitHash(v string)`

SetCommitHash sets CommitHash field to given value.

### HasCommitHash

`func (o *V0ArchivedBuildResponseItemModel) HasCommitHash() bool`

HasCommitHash returns a boolean if a field has been set.

### GetCommitMessage

`func (o *V0ArchivedBuildResponseItemModel) GetCommitMessage() string`

GetCommitMessage returns the CommitMessage field if non-nil, zero value otherwise.

### GetCommitMessageOk

`func (o *V0ArchivedBuildResponseItemModel) GetCommitMessageOk() (*string, bool)`

GetCommitMessageOk returns a tuple with the CommitMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitMessage

`func (o *V0ArchivedBuildResponseItemModel) SetCommitMessage(v string)`

SetCommitMessage sets CommitMessage field to given value.

### HasCommitMessage

`func (o *V0ArchivedBuildResponseItemModel) HasCommitMessage() bool`

HasCommitMessage returns a boolean if a field has been set.

### GetCreditCost

`func (o *V0ArchivedBuildResponseItemModel) GetCreditCost() int32`

GetCreditCost returns the CreditCost field if non-nil, zero value otherwise.

### GetCreditCostOk

`func (o *V0ArchivedBuildResponseItemModel) GetCreditCostOk() (*int32, bool)`

GetCreditCostOk returns a tuple with the CreditCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditCost

`func (o *V0ArchivedBuildResponseItemModel) SetCreditCost(v int32)`

SetCreditCost sets CreditCost field to given value.

### HasCreditCost

`func (o *V0ArchivedBuildResponseItemModel) HasCreditCost() bool`

HasCreditCost returns a boolean if a field has been set.

### GetFinishedAt

`func (o *V0ArchivedBuildResponseItemModel) GetFinishedAt() string`

GetFinishedAt returns the FinishedAt field if non-nil, zero value otherwise.

### GetFinishedAtOk

`func (o *V0ArchivedBuildResponseItemModel) GetFinishedAtOk() (*string, bool)`

GetFinishedAtOk returns a tuple with the FinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedAt

`func (o *V0ArchivedBuildResponseItemModel) SetFinishedAt(v string)`

SetFinishedAt sets FinishedAt field to given value.

### HasFinishedAt

`func (o *V0ArchivedBuildResponseItemModel) HasFinishedAt() bool`

HasFinishedAt returns a boolean if a field has been set.

### GetMachineTypeId

`func (o *V0ArchivedBuildResponseItemModel) GetMachineTypeId() string`

GetMachineTypeId returns the MachineTypeId field if non-nil, zero value otherwise.

### GetMachineTypeIdOk

`func (o *V0ArchivedBuildResponseItemModel) GetMachineTypeIdOk() (*string, bool)`

GetMachineTypeIdOk returns a tuple with the MachineTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineTypeId

`func (o *V0ArchivedBuildResponseItemModel) SetMachineTypeId(v string)`

SetMachineTypeId sets MachineTypeId field to given value.

### HasMachineTypeId

`func (o *V0ArchivedBuildResponseItemModel) HasMachineTypeId() bool`

HasMachineTypeId returns a boolean if a field has been set.

### GetOriginalBuildParams

`func (o *V0ArchivedBuildResponseItemModel) GetOriginalBuildParams() []int32`

GetOriginalBuildParams returns the OriginalBuildParams field if non-nil, zero value otherwise.

### GetOriginalBuildParamsOk

`func (o *V0ArchivedBuildResponseItemModel) GetOriginalBuildParamsOk() (*[]int32, bool)`

GetOriginalBuildParamsOk returns a tuple with the OriginalBuildParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalBuildParams

`func (o *V0ArchivedBuildResponseItemModel) SetOriginalBuildParams(v []int32)`

SetOriginalBuildParams sets OriginalBuildParams field to given value.

### HasOriginalBuildParams

`func (o *V0ArchivedBuildResponseItemModel) HasOriginalBuildParams() bool`

HasOriginalBuildParams returns a boolean if a field has been set.

### GetPipelineWorkflowId

`func (o *V0ArchivedBuildResponseItemModel) GetPipelineWorkflowId() string`

GetPipelineWorkflowId returns the PipelineWorkflowId field if non-nil, zero value otherwise.

### GetPipelineWorkflowIdOk

`func (o *V0ArchivedBuildResponseItemModel) GetPipelineWorkflowIdOk() (*string, bool)`

GetPipelineWorkflowIdOk returns a tuple with the PipelineWorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipelineWorkflowId

`func (o *V0ArchivedBuildResponseItemModel) SetPipelineWorkflowId(v string)`

SetPipelineWorkflowId sets PipelineWorkflowId field to given value.

### HasPipelineWorkflowId

`func (o *V0ArchivedBuildResponseItemModel) HasPipelineWorkflowId() bool`

HasPipelineWorkflowId returns a boolean if a field has been set.

### GetPullRequestId

`func (o *V0ArchivedBuildResponseItemModel) GetPullRequestId() int32`

GetPullRequestId returns the PullRequestId field if non-nil, zero value otherwise.

### GetPullRequestIdOk

`func (o *V0ArchivedBuildResponseItemModel) GetPullRequestIdOk() (*int32, bool)`

GetPullRequestIdOk returns a tuple with the PullRequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullRequestId

`func (o *V0ArchivedBuildResponseItemModel) SetPullRequestId(v int32)`

SetPullRequestId sets PullRequestId field to given value.

### HasPullRequestId

`func (o *V0ArchivedBuildResponseItemModel) HasPullRequestId() bool`

HasPullRequestId returns a boolean if a field has been set.

### GetPullRequestTargetBranch

`func (o *V0ArchivedBuildResponseItemModel) GetPullRequestTargetBranch() string`

GetPullRequestTargetBranch returns the PullRequestTargetBranch field if non-nil, zero value otherwise.

### GetPullRequestTargetBranchOk

`func (o *V0ArchivedBuildResponseItemModel) GetPullRequestTargetBranchOk() (*string, bool)`

GetPullRequestTargetBranchOk returns a tuple with the PullRequestTargetBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullRequestTargetBranch

`func (o *V0ArchivedBuildResponseItemModel) SetPullRequestTargetBranch(v string)`

SetPullRequestTargetBranch sets PullRequestTargetBranch field to given value.

### HasPullRequestTargetBranch

`func (o *V0ArchivedBuildResponseItemModel) HasPullRequestTargetBranch() bool`

HasPullRequestTargetBranch returns a boolean if a field has been set.

### GetPullRequestViewUrl

`func (o *V0ArchivedBuildResponseItemModel) GetPullRequestViewUrl() string`

GetPullRequestViewUrl returns the PullRequestViewUrl field if non-nil, zero value otherwise.

### GetPullRequestViewUrlOk

`func (o *V0ArchivedBuildResponseItemModel) GetPullRequestViewUrlOk() (*string, bool)`

GetPullRequestViewUrlOk returns a tuple with the PullRequestViewUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullRequestViewUrl

`func (o *V0ArchivedBuildResponseItemModel) SetPullRequestViewUrl(v string)`

SetPullRequestViewUrl sets PullRequestViewUrl field to given value.

### HasPullRequestViewUrl

`func (o *V0ArchivedBuildResponseItemModel) HasPullRequestViewUrl() bool`

HasPullRequestViewUrl returns a boolean if a field has been set.

### GetSlug

`func (o *V0ArchivedBuildResponseItemModel) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *V0ArchivedBuildResponseItemModel) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *V0ArchivedBuildResponseItemModel) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *V0ArchivedBuildResponseItemModel) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetStackIdentifier

`func (o *V0ArchivedBuildResponseItemModel) GetStackIdentifier() string`

GetStackIdentifier returns the StackIdentifier field if non-nil, zero value otherwise.

### GetStackIdentifierOk

`func (o *V0ArchivedBuildResponseItemModel) GetStackIdentifierOk() (*string, bool)`

GetStackIdentifierOk returns a tuple with the StackIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackIdentifier

`func (o *V0ArchivedBuildResponseItemModel) SetStackIdentifier(v string)`

SetStackIdentifier sets StackIdentifier field to given value.

### HasStackIdentifier

`func (o *V0ArchivedBuildResponseItemModel) HasStackIdentifier() bool`

HasStackIdentifier returns a boolean if a field has been set.

### GetStartedOnWorkerAt

`func (o *V0ArchivedBuildResponseItemModel) GetStartedOnWorkerAt() string`

GetStartedOnWorkerAt returns the StartedOnWorkerAt field if non-nil, zero value otherwise.

### GetStartedOnWorkerAtOk

`func (o *V0ArchivedBuildResponseItemModel) GetStartedOnWorkerAtOk() (*string, bool)`

GetStartedOnWorkerAtOk returns a tuple with the StartedOnWorkerAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedOnWorkerAt

`func (o *V0ArchivedBuildResponseItemModel) SetStartedOnWorkerAt(v string)`

SetStartedOnWorkerAt sets StartedOnWorkerAt field to given value.

### HasStartedOnWorkerAt

`func (o *V0ArchivedBuildResponseItemModel) HasStartedOnWorkerAt() bool`

HasStartedOnWorkerAt returns a boolean if a field has been set.

### GetStatus

`func (o *V0ArchivedBuildResponseItemModel) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *V0ArchivedBuildResponseItemModel) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *V0ArchivedBuildResponseItemModel) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *V0ArchivedBuildResponseItemModel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusText

`func (o *V0ArchivedBuildResponseItemModel) GetStatusText() string`

GetStatusText returns the StatusText field if non-nil, zero value otherwise.

### GetStatusTextOk

`func (o *V0ArchivedBuildResponseItemModel) GetStatusTextOk() (*string, bool)`

GetStatusTextOk returns a tuple with the StatusText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusText

`func (o *V0ArchivedBuildResponseItemModel) SetStatusText(v string)`

SetStatusText sets StatusText field to given value.

### HasStatusText

`func (o *V0ArchivedBuildResponseItemModel) HasStatusText() bool`

HasStatusText returns a boolean if a field has been set.

### GetTag

`func (o *V0ArchivedBuildResponseItemModel) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *V0ArchivedBuildResponseItemModel) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *V0ArchivedBuildResponseItemModel) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *V0ArchivedBuildResponseItemModel) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetTriggeredBy

`func (o *V0ArchivedBuildResponseItemModel) GetTriggeredBy() string`

GetTriggeredBy returns the TriggeredBy field if non-nil, zero value otherwise.

### GetTriggeredByOk

`func (o *V0ArchivedBuildResponseItemModel) GetTriggeredByOk() (*string, bool)`

GetTriggeredByOk returns a tuple with the TriggeredBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggeredBy

`func (o *V0ArchivedBuildResponseItemModel) SetTriggeredBy(v string)`

SetTriggeredBy sets TriggeredBy field to given value.

### HasTriggeredBy

`func (o *V0ArchivedBuildResponseItemModel) HasTriggeredBy() bool`

HasTriggeredBy returns a boolean if a field has been set.

### GetTriggeredWorkflow

`func (o *V0ArchivedBuildResponseItemModel) GetTriggeredWorkflow() string`

GetTriggeredWorkflow returns the TriggeredWorkflow field if non-nil, zero value otherwise.

### GetTriggeredWorkflowOk

`func (o *V0ArchivedBuildResponseItemModel) GetTriggeredWorkflowOk() (*string, bool)`

GetTriggeredWorkflowOk returns a tuple with the TriggeredWorkflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggeredWorkflow

`func (o *V0ArchivedBuildResponseItemModel) SetTriggeredWorkflow(v string)`

SetTriggeredWorkflow sets TriggeredWorkflow field to given value.

### HasTriggeredWorkflow

`func (o *V0ArchivedBuildResponseItemModel) HasTriggeredWorkflow() bool`

HasTriggeredWorkflow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


