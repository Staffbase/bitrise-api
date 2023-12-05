# V0BuildResponseItemModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AbortReason** | Pointer to [**NullsString**](NullsString.md) |  | [optional] 
**Branch** | Pointer to [**NullsString**](NullsString.md) |  | [optional] 
**BuildNumber** | Pointer to **int32** |  | [optional] 
**CommitHash** | Pointer to [**NullsString**](NullsString.md) |  | [optional] 
**CommitMessage** | Pointer to [**NullsString**](NullsString.md) |  | [optional] 
**CommitViewUrl** | Pointer to [**NullsString**](NullsString.md) |  | [optional] 
**CreditCost** | Pointer to [**NullsInt64**](NullsInt64.md) |  | [optional] 
**EnvironmentPrepareFinishedAt** | Pointer to **string** |  | [optional] 
**FinishedAt** | Pointer to **string** |  | [optional] 
**IsOnHold** | Pointer to **bool** |  | [optional] 
**IsProcessed** | Pointer to **bool** |  | [optional] 
**IsStatusSent** | Pointer to **bool** |  | [optional] 
**LogFormat** | Pointer to **string** |  | [optional] 
**MachineTypeId** | Pointer to [**NullsString**](NullsString.md) |  | [optional] 
**OriginalBuildParams** | Pointer to **[]int32** |  | [optional] 
**PipelineWorkflowId** | Pointer to **string** |  | [optional] 
**PullRequestId** | Pointer to **int32** |  | [optional] 
**PullRequestTargetBranch** | Pointer to [**NullsString**](NullsString.md) |  | [optional] 
**PullRequestViewUrl** | Pointer to [**NullsString**](NullsString.md) |  | [optional] 
**Slug** | Pointer to **string** |  | [optional] 
**StackIdentifier** | Pointer to [**NullsString**](NullsString.md) |  | [optional] 
**StartedOnWorkerAt** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **int32** |  | [optional] 
**StatusText** | Pointer to **string** |  | [optional] 
**Tag** | Pointer to [**NullsString**](NullsString.md) |  | [optional] 
**TriggeredAt** | Pointer to **string** |  | [optional] 
**TriggeredBy** | Pointer to [**NullsString**](NullsString.md) |  | [optional] 
**TriggeredWorkflow** | Pointer to **string** |  | [optional] 

## Methods

### NewV0BuildResponseItemModel

`func NewV0BuildResponseItemModel() *V0BuildResponseItemModel`

NewV0BuildResponseItemModel instantiates a new V0BuildResponseItemModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0BuildResponseItemModelWithDefaults

`func NewV0BuildResponseItemModelWithDefaults() *V0BuildResponseItemModel`

NewV0BuildResponseItemModelWithDefaults instantiates a new V0BuildResponseItemModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAbortReason

`func (o *V0BuildResponseItemModel) GetAbortReason() NullsString`

GetAbortReason returns the AbortReason field if non-nil, zero value otherwise.

### GetAbortReasonOk

`func (o *V0BuildResponseItemModel) GetAbortReasonOk() (*NullsString, bool)`

GetAbortReasonOk returns a tuple with the AbortReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbortReason

`func (o *V0BuildResponseItemModel) SetAbortReason(v NullsString)`

SetAbortReason sets AbortReason field to given value.

### HasAbortReason

`func (o *V0BuildResponseItemModel) HasAbortReason() bool`

HasAbortReason returns a boolean if a field has been set.

### GetBranch

`func (o *V0BuildResponseItemModel) GetBranch() NullsString`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *V0BuildResponseItemModel) GetBranchOk() (*NullsString, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *V0BuildResponseItemModel) SetBranch(v NullsString)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *V0BuildResponseItemModel) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### GetBuildNumber

`func (o *V0BuildResponseItemModel) GetBuildNumber() int32`

GetBuildNumber returns the BuildNumber field if non-nil, zero value otherwise.

### GetBuildNumberOk

`func (o *V0BuildResponseItemModel) GetBuildNumberOk() (*int32, bool)`

GetBuildNumberOk returns a tuple with the BuildNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildNumber

`func (o *V0BuildResponseItemModel) SetBuildNumber(v int32)`

SetBuildNumber sets BuildNumber field to given value.

### HasBuildNumber

`func (o *V0BuildResponseItemModel) HasBuildNumber() bool`

HasBuildNumber returns a boolean if a field has been set.

### GetCommitHash

`func (o *V0BuildResponseItemModel) GetCommitHash() NullsString`

GetCommitHash returns the CommitHash field if non-nil, zero value otherwise.

### GetCommitHashOk

`func (o *V0BuildResponseItemModel) GetCommitHashOk() (*NullsString, bool)`

GetCommitHashOk returns a tuple with the CommitHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitHash

`func (o *V0BuildResponseItemModel) SetCommitHash(v NullsString)`

SetCommitHash sets CommitHash field to given value.

### HasCommitHash

`func (o *V0BuildResponseItemModel) HasCommitHash() bool`

HasCommitHash returns a boolean if a field has been set.

### GetCommitMessage

`func (o *V0BuildResponseItemModel) GetCommitMessage() NullsString`

GetCommitMessage returns the CommitMessage field if non-nil, zero value otherwise.

### GetCommitMessageOk

`func (o *V0BuildResponseItemModel) GetCommitMessageOk() (*NullsString, bool)`

GetCommitMessageOk returns a tuple with the CommitMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitMessage

`func (o *V0BuildResponseItemModel) SetCommitMessage(v NullsString)`

SetCommitMessage sets CommitMessage field to given value.

### HasCommitMessage

`func (o *V0BuildResponseItemModel) HasCommitMessage() bool`

HasCommitMessage returns a boolean if a field has been set.

### GetCommitViewUrl

`func (o *V0BuildResponseItemModel) GetCommitViewUrl() NullsString`

GetCommitViewUrl returns the CommitViewUrl field if non-nil, zero value otherwise.

### GetCommitViewUrlOk

`func (o *V0BuildResponseItemModel) GetCommitViewUrlOk() (*NullsString, bool)`

GetCommitViewUrlOk returns a tuple with the CommitViewUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitViewUrl

`func (o *V0BuildResponseItemModel) SetCommitViewUrl(v NullsString)`

SetCommitViewUrl sets CommitViewUrl field to given value.

### HasCommitViewUrl

`func (o *V0BuildResponseItemModel) HasCommitViewUrl() bool`

HasCommitViewUrl returns a boolean if a field has been set.

### GetCreditCost

`func (o *V0BuildResponseItemModel) GetCreditCost() NullsInt64`

GetCreditCost returns the CreditCost field if non-nil, zero value otherwise.

### GetCreditCostOk

`func (o *V0BuildResponseItemModel) GetCreditCostOk() (*NullsInt64, bool)`

GetCreditCostOk returns a tuple with the CreditCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditCost

`func (o *V0BuildResponseItemModel) SetCreditCost(v NullsInt64)`

SetCreditCost sets CreditCost field to given value.

### HasCreditCost

`func (o *V0BuildResponseItemModel) HasCreditCost() bool`

HasCreditCost returns a boolean if a field has been set.

### GetEnvironmentPrepareFinishedAt

`func (o *V0BuildResponseItemModel) GetEnvironmentPrepareFinishedAt() string`

GetEnvironmentPrepareFinishedAt returns the EnvironmentPrepareFinishedAt field if non-nil, zero value otherwise.

### GetEnvironmentPrepareFinishedAtOk

`func (o *V0BuildResponseItemModel) GetEnvironmentPrepareFinishedAtOk() (*string, bool)`

GetEnvironmentPrepareFinishedAtOk returns a tuple with the EnvironmentPrepareFinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentPrepareFinishedAt

`func (o *V0BuildResponseItemModel) SetEnvironmentPrepareFinishedAt(v string)`

SetEnvironmentPrepareFinishedAt sets EnvironmentPrepareFinishedAt field to given value.

### HasEnvironmentPrepareFinishedAt

`func (o *V0BuildResponseItemModel) HasEnvironmentPrepareFinishedAt() bool`

HasEnvironmentPrepareFinishedAt returns a boolean if a field has been set.

### GetFinishedAt

`func (o *V0BuildResponseItemModel) GetFinishedAt() string`

GetFinishedAt returns the FinishedAt field if non-nil, zero value otherwise.

### GetFinishedAtOk

`func (o *V0BuildResponseItemModel) GetFinishedAtOk() (*string, bool)`

GetFinishedAtOk returns a tuple with the FinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedAt

`func (o *V0BuildResponseItemModel) SetFinishedAt(v string)`

SetFinishedAt sets FinishedAt field to given value.

### HasFinishedAt

`func (o *V0BuildResponseItemModel) HasFinishedAt() bool`

HasFinishedAt returns a boolean if a field has been set.

### GetIsOnHold

`func (o *V0BuildResponseItemModel) GetIsOnHold() bool`

GetIsOnHold returns the IsOnHold field if non-nil, zero value otherwise.

### GetIsOnHoldOk

`func (o *V0BuildResponseItemModel) GetIsOnHoldOk() (*bool, bool)`

GetIsOnHoldOk returns a tuple with the IsOnHold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOnHold

`func (o *V0BuildResponseItemModel) SetIsOnHold(v bool)`

SetIsOnHold sets IsOnHold field to given value.

### HasIsOnHold

`func (o *V0BuildResponseItemModel) HasIsOnHold() bool`

HasIsOnHold returns a boolean if a field has been set.

### GetIsProcessed

`func (o *V0BuildResponseItemModel) GetIsProcessed() bool`

GetIsProcessed returns the IsProcessed field if non-nil, zero value otherwise.

### GetIsProcessedOk

`func (o *V0BuildResponseItemModel) GetIsProcessedOk() (*bool, bool)`

GetIsProcessedOk returns a tuple with the IsProcessed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsProcessed

`func (o *V0BuildResponseItemModel) SetIsProcessed(v bool)`

SetIsProcessed sets IsProcessed field to given value.

### HasIsProcessed

`func (o *V0BuildResponseItemModel) HasIsProcessed() bool`

HasIsProcessed returns a boolean if a field has been set.

### GetIsStatusSent

`func (o *V0BuildResponseItemModel) GetIsStatusSent() bool`

GetIsStatusSent returns the IsStatusSent field if non-nil, zero value otherwise.

### GetIsStatusSentOk

`func (o *V0BuildResponseItemModel) GetIsStatusSentOk() (*bool, bool)`

GetIsStatusSentOk returns a tuple with the IsStatusSent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsStatusSent

`func (o *V0BuildResponseItemModel) SetIsStatusSent(v bool)`

SetIsStatusSent sets IsStatusSent field to given value.

### HasIsStatusSent

`func (o *V0BuildResponseItemModel) HasIsStatusSent() bool`

HasIsStatusSent returns a boolean if a field has been set.

### GetLogFormat

`func (o *V0BuildResponseItemModel) GetLogFormat() string`

GetLogFormat returns the LogFormat field if non-nil, zero value otherwise.

### GetLogFormatOk

`func (o *V0BuildResponseItemModel) GetLogFormatOk() (*string, bool)`

GetLogFormatOk returns a tuple with the LogFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFormat

`func (o *V0BuildResponseItemModel) SetLogFormat(v string)`

SetLogFormat sets LogFormat field to given value.

### HasLogFormat

`func (o *V0BuildResponseItemModel) HasLogFormat() bool`

HasLogFormat returns a boolean if a field has been set.

### GetMachineTypeId

`func (o *V0BuildResponseItemModel) GetMachineTypeId() NullsString`

GetMachineTypeId returns the MachineTypeId field if non-nil, zero value otherwise.

### GetMachineTypeIdOk

`func (o *V0BuildResponseItemModel) GetMachineTypeIdOk() (*NullsString, bool)`

GetMachineTypeIdOk returns a tuple with the MachineTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineTypeId

`func (o *V0BuildResponseItemModel) SetMachineTypeId(v NullsString)`

SetMachineTypeId sets MachineTypeId field to given value.

### HasMachineTypeId

`func (o *V0BuildResponseItemModel) HasMachineTypeId() bool`

HasMachineTypeId returns a boolean if a field has been set.

### GetOriginalBuildParams

`func (o *V0BuildResponseItemModel) GetOriginalBuildParams() []int32`

GetOriginalBuildParams returns the OriginalBuildParams field if non-nil, zero value otherwise.

### GetOriginalBuildParamsOk

`func (o *V0BuildResponseItemModel) GetOriginalBuildParamsOk() (*[]int32, bool)`

GetOriginalBuildParamsOk returns a tuple with the OriginalBuildParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalBuildParams

`func (o *V0BuildResponseItemModel) SetOriginalBuildParams(v []int32)`

SetOriginalBuildParams sets OriginalBuildParams field to given value.

### HasOriginalBuildParams

`func (o *V0BuildResponseItemModel) HasOriginalBuildParams() bool`

HasOriginalBuildParams returns a boolean if a field has been set.

### GetPipelineWorkflowId

`func (o *V0BuildResponseItemModel) GetPipelineWorkflowId() string`

GetPipelineWorkflowId returns the PipelineWorkflowId field if non-nil, zero value otherwise.

### GetPipelineWorkflowIdOk

`func (o *V0BuildResponseItemModel) GetPipelineWorkflowIdOk() (*string, bool)`

GetPipelineWorkflowIdOk returns a tuple with the PipelineWorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipelineWorkflowId

`func (o *V0BuildResponseItemModel) SetPipelineWorkflowId(v string)`

SetPipelineWorkflowId sets PipelineWorkflowId field to given value.

### HasPipelineWorkflowId

`func (o *V0BuildResponseItemModel) HasPipelineWorkflowId() bool`

HasPipelineWorkflowId returns a boolean if a field has been set.

### GetPullRequestId

`func (o *V0BuildResponseItemModel) GetPullRequestId() int32`

GetPullRequestId returns the PullRequestId field if non-nil, zero value otherwise.

### GetPullRequestIdOk

`func (o *V0BuildResponseItemModel) GetPullRequestIdOk() (*int32, bool)`

GetPullRequestIdOk returns a tuple with the PullRequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullRequestId

`func (o *V0BuildResponseItemModel) SetPullRequestId(v int32)`

SetPullRequestId sets PullRequestId field to given value.

### HasPullRequestId

`func (o *V0BuildResponseItemModel) HasPullRequestId() bool`

HasPullRequestId returns a boolean if a field has been set.

### GetPullRequestTargetBranch

`func (o *V0BuildResponseItemModel) GetPullRequestTargetBranch() NullsString`

GetPullRequestTargetBranch returns the PullRequestTargetBranch field if non-nil, zero value otherwise.

### GetPullRequestTargetBranchOk

`func (o *V0BuildResponseItemModel) GetPullRequestTargetBranchOk() (*NullsString, bool)`

GetPullRequestTargetBranchOk returns a tuple with the PullRequestTargetBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullRequestTargetBranch

`func (o *V0BuildResponseItemModel) SetPullRequestTargetBranch(v NullsString)`

SetPullRequestTargetBranch sets PullRequestTargetBranch field to given value.

### HasPullRequestTargetBranch

`func (o *V0BuildResponseItemModel) HasPullRequestTargetBranch() bool`

HasPullRequestTargetBranch returns a boolean if a field has been set.

### GetPullRequestViewUrl

`func (o *V0BuildResponseItemModel) GetPullRequestViewUrl() NullsString`

GetPullRequestViewUrl returns the PullRequestViewUrl field if non-nil, zero value otherwise.

### GetPullRequestViewUrlOk

`func (o *V0BuildResponseItemModel) GetPullRequestViewUrlOk() (*NullsString, bool)`

GetPullRequestViewUrlOk returns a tuple with the PullRequestViewUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullRequestViewUrl

`func (o *V0BuildResponseItemModel) SetPullRequestViewUrl(v NullsString)`

SetPullRequestViewUrl sets PullRequestViewUrl field to given value.

### HasPullRequestViewUrl

`func (o *V0BuildResponseItemModel) HasPullRequestViewUrl() bool`

HasPullRequestViewUrl returns a boolean if a field has been set.

### GetSlug

`func (o *V0BuildResponseItemModel) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *V0BuildResponseItemModel) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *V0BuildResponseItemModel) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *V0BuildResponseItemModel) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetStackIdentifier

`func (o *V0BuildResponseItemModel) GetStackIdentifier() NullsString`

GetStackIdentifier returns the StackIdentifier field if non-nil, zero value otherwise.

### GetStackIdentifierOk

`func (o *V0BuildResponseItemModel) GetStackIdentifierOk() (*NullsString, bool)`

GetStackIdentifierOk returns a tuple with the StackIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackIdentifier

`func (o *V0BuildResponseItemModel) SetStackIdentifier(v NullsString)`

SetStackIdentifier sets StackIdentifier field to given value.

### HasStackIdentifier

`func (o *V0BuildResponseItemModel) HasStackIdentifier() bool`

HasStackIdentifier returns a boolean if a field has been set.

### GetStartedOnWorkerAt

`func (o *V0BuildResponseItemModel) GetStartedOnWorkerAt() string`

GetStartedOnWorkerAt returns the StartedOnWorkerAt field if non-nil, zero value otherwise.

### GetStartedOnWorkerAtOk

`func (o *V0BuildResponseItemModel) GetStartedOnWorkerAtOk() (*string, bool)`

GetStartedOnWorkerAtOk returns a tuple with the StartedOnWorkerAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedOnWorkerAt

`func (o *V0BuildResponseItemModel) SetStartedOnWorkerAt(v string)`

SetStartedOnWorkerAt sets StartedOnWorkerAt field to given value.

### HasStartedOnWorkerAt

`func (o *V0BuildResponseItemModel) HasStartedOnWorkerAt() bool`

HasStartedOnWorkerAt returns a boolean if a field has been set.

### GetStatus

`func (o *V0BuildResponseItemModel) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *V0BuildResponseItemModel) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *V0BuildResponseItemModel) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *V0BuildResponseItemModel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusText

`func (o *V0BuildResponseItemModel) GetStatusText() string`

GetStatusText returns the StatusText field if non-nil, zero value otherwise.

### GetStatusTextOk

`func (o *V0BuildResponseItemModel) GetStatusTextOk() (*string, bool)`

GetStatusTextOk returns a tuple with the StatusText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusText

`func (o *V0BuildResponseItemModel) SetStatusText(v string)`

SetStatusText sets StatusText field to given value.

### HasStatusText

`func (o *V0BuildResponseItemModel) HasStatusText() bool`

HasStatusText returns a boolean if a field has been set.

### GetTag

`func (o *V0BuildResponseItemModel) GetTag() NullsString`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *V0BuildResponseItemModel) GetTagOk() (*NullsString, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *V0BuildResponseItemModel) SetTag(v NullsString)`

SetTag sets Tag field to given value.

### HasTag

`func (o *V0BuildResponseItemModel) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetTriggeredAt

`func (o *V0BuildResponseItemModel) GetTriggeredAt() string`

GetTriggeredAt returns the TriggeredAt field if non-nil, zero value otherwise.

### GetTriggeredAtOk

`func (o *V0BuildResponseItemModel) GetTriggeredAtOk() (*string, bool)`

GetTriggeredAtOk returns a tuple with the TriggeredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggeredAt

`func (o *V0BuildResponseItemModel) SetTriggeredAt(v string)`

SetTriggeredAt sets TriggeredAt field to given value.

### HasTriggeredAt

`func (o *V0BuildResponseItemModel) HasTriggeredAt() bool`

HasTriggeredAt returns a boolean if a field has been set.

### GetTriggeredBy

`func (o *V0BuildResponseItemModel) GetTriggeredBy() NullsString`

GetTriggeredBy returns the TriggeredBy field if non-nil, zero value otherwise.

### GetTriggeredByOk

`func (o *V0BuildResponseItemModel) GetTriggeredByOk() (*NullsString, bool)`

GetTriggeredByOk returns a tuple with the TriggeredBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggeredBy

`func (o *V0BuildResponseItemModel) SetTriggeredBy(v NullsString)`

SetTriggeredBy sets TriggeredBy field to given value.

### HasTriggeredBy

`func (o *V0BuildResponseItemModel) HasTriggeredBy() bool`

HasTriggeredBy returns a boolean if a field has been set.

### GetTriggeredWorkflow

`func (o *V0BuildResponseItemModel) GetTriggeredWorkflow() string`

GetTriggeredWorkflow returns the TriggeredWorkflow field if non-nil, zero value otherwise.

### GetTriggeredWorkflowOk

`func (o *V0BuildResponseItemModel) GetTriggeredWorkflowOk() (*string, bool)`

GetTriggeredWorkflowOk returns a tuple with the TriggeredWorkflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggeredWorkflow

`func (o *V0BuildResponseItemModel) SetTriggeredWorkflow(v string)`

SetTriggeredWorkflow sets TriggeredWorkflow field to given value.

### HasTriggeredWorkflow

`func (o *V0BuildResponseItemModel) HasTriggeredWorkflow() bool`

HasTriggeredWorkflow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


