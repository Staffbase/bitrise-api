# V0BuildListAllResponseItemModel

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
**MachineTypeId** | Pointer to [**NullsString**](NullsString.md) |  | [optional] 
**OriginalBuildParams** | Pointer to **[]int32** |  | [optional] 
**PullRequestId** | Pointer to **int32** |  | [optional] 
**PullRequestTargetBranch** | Pointer to [**NullsString**](NullsString.md) |  | [optional] 
**PullRequestViewUrl** | Pointer to [**NullsString**](NullsString.md) |  | [optional] 
**Repository** | Pointer to [**V0AppResponseItemModel**](V0AppResponseItemModel.md) |  | [optional] 
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

### NewV0BuildListAllResponseItemModel

`func NewV0BuildListAllResponseItemModel() *V0BuildListAllResponseItemModel`

NewV0BuildListAllResponseItemModel instantiates a new V0BuildListAllResponseItemModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0BuildListAllResponseItemModelWithDefaults

`func NewV0BuildListAllResponseItemModelWithDefaults() *V0BuildListAllResponseItemModel`

NewV0BuildListAllResponseItemModelWithDefaults instantiates a new V0BuildListAllResponseItemModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAbortReason

`func (o *V0BuildListAllResponseItemModel) GetAbortReason() NullsString`

GetAbortReason returns the AbortReason field if non-nil, zero value otherwise.

### GetAbortReasonOk

`func (o *V0BuildListAllResponseItemModel) GetAbortReasonOk() (*NullsString, bool)`

GetAbortReasonOk returns a tuple with the AbortReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbortReason

`func (o *V0BuildListAllResponseItemModel) SetAbortReason(v NullsString)`

SetAbortReason sets AbortReason field to given value.

### HasAbortReason

`func (o *V0BuildListAllResponseItemModel) HasAbortReason() bool`

HasAbortReason returns a boolean if a field has been set.

### GetBranch

`func (o *V0BuildListAllResponseItemModel) GetBranch() NullsString`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *V0BuildListAllResponseItemModel) GetBranchOk() (*NullsString, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *V0BuildListAllResponseItemModel) SetBranch(v NullsString)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *V0BuildListAllResponseItemModel) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### GetBuildNumber

`func (o *V0BuildListAllResponseItemModel) GetBuildNumber() int32`

GetBuildNumber returns the BuildNumber field if non-nil, zero value otherwise.

### GetBuildNumberOk

`func (o *V0BuildListAllResponseItemModel) GetBuildNumberOk() (*int32, bool)`

GetBuildNumberOk returns a tuple with the BuildNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildNumber

`func (o *V0BuildListAllResponseItemModel) SetBuildNumber(v int32)`

SetBuildNumber sets BuildNumber field to given value.

### HasBuildNumber

`func (o *V0BuildListAllResponseItemModel) HasBuildNumber() bool`

HasBuildNumber returns a boolean if a field has been set.

### GetCommitHash

`func (o *V0BuildListAllResponseItemModel) GetCommitHash() NullsString`

GetCommitHash returns the CommitHash field if non-nil, zero value otherwise.

### GetCommitHashOk

`func (o *V0BuildListAllResponseItemModel) GetCommitHashOk() (*NullsString, bool)`

GetCommitHashOk returns a tuple with the CommitHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitHash

`func (o *V0BuildListAllResponseItemModel) SetCommitHash(v NullsString)`

SetCommitHash sets CommitHash field to given value.

### HasCommitHash

`func (o *V0BuildListAllResponseItemModel) HasCommitHash() bool`

HasCommitHash returns a boolean if a field has been set.

### GetCommitMessage

`func (o *V0BuildListAllResponseItemModel) GetCommitMessage() NullsString`

GetCommitMessage returns the CommitMessage field if non-nil, zero value otherwise.

### GetCommitMessageOk

`func (o *V0BuildListAllResponseItemModel) GetCommitMessageOk() (*NullsString, bool)`

GetCommitMessageOk returns a tuple with the CommitMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitMessage

`func (o *V0BuildListAllResponseItemModel) SetCommitMessage(v NullsString)`

SetCommitMessage sets CommitMessage field to given value.

### HasCommitMessage

`func (o *V0BuildListAllResponseItemModel) HasCommitMessage() bool`

HasCommitMessage returns a boolean if a field has been set.

### GetCommitViewUrl

`func (o *V0BuildListAllResponseItemModel) GetCommitViewUrl() NullsString`

GetCommitViewUrl returns the CommitViewUrl field if non-nil, zero value otherwise.

### GetCommitViewUrlOk

`func (o *V0BuildListAllResponseItemModel) GetCommitViewUrlOk() (*NullsString, bool)`

GetCommitViewUrlOk returns a tuple with the CommitViewUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitViewUrl

`func (o *V0BuildListAllResponseItemModel) SetCommitViewUrl(v NullsString)`

SetCommitViewUrl sets CommitViewUrl field to given value.

### HasCommitViewUrl

`func (o *V0BuildListAllResponseItemModel) HasCommitViewUrl() bool`

HasCommitViewUrl returns a boolean if a field has been set.

### GetCreditCost

`func (o *V0BuildListAllResponseItemModel) GetCreditCost() NullsInt64`

GetCreditCost returns the CreditCost field if non-nil, zero value otherwise.

### GetCreditCostOk

`func (o *V0BuildListAllResponseItemModel) GetCreditCostOk() (*NullsInt64, bool)`

GetCreditCostOk returns a tuple with the CreditCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditCost

`func (o *V0BuildListAllResponseItemModel) SetCreditCost(v NullsInt64)`

SetCreditCost sets CreditCost field to given value.

### HasCreditCost

`func (o *V0BuildListAllResponseItemModel) HasCreditCost() bool`

HasCreditCost returns a boolean if a field has been set.

### GetEnvironmentPrepareFinishedAt

`func (o *V0BuildListAllResponseItemModel) GetEnvironmentPrepareFinishedAt() string`

GetEnvironmentPrepareFinishedAt returns the EnvironmentPrepareFinishedAt field if non-nil, zero value otherwise.

### GetEnvironmentPrepareFinishedAtOk

`func (o *V0BuildListAllResponseItemModel) GetEnvironmentPrepareFinishedAtOk() (*string, bool)`

GetEnvironmentPrepareFinishedAtOk returns a tuple with the EnvironmentPrepareFinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentPrepareFinishedAt

`func (o *V0BuildListAllResponseItemModel) SetEnvironmentPrepareFinishedAt(v string)`

SetEnvironmentPrepareFinishedAt sets EnvironmentPrepareFinishedAt field to given value.

### HasEnvironmentPrepareFinishedAt

`func (o *V0BuildListAllResponseItemModel) HasEnvironmentPrepareFinishedAt() bool`

HasEnvironmentPrepareFinishedAt returns a boolean if a field has been set.

### GetFinishedAt

`func (o *V0BuildListAllResponseItemModel) GetFinishedAt() string`

GetFinishedAt returns the FinishedAt field if non-nil, zero value otherwise.

### GetFinishedAtOk

`func (o *V0BuildListAllResponseItemModel) GetFinishedAtOk() (*string, bool)`

GetFinishedAtOk returns a tuple with the FinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedAt

`func (o *V0BuildListAllResponseItemModel) SetFinishedAt(v string)`

SetFinishedAt sets FinishedAt field to given value.

### HasFinishedAt

`func (o *V0BuildListAllResponseItemModel) HasFinishedAt() bool`

HasFinishedAt returns a boolean if a field has been set.

### GetIsOnHold

`func (o *V0BuildListAllResponseItemModel) GetIsOnHold() bool`

GetIsOnHold returns the IsOnHold field if non-nil, zero value otherwise.

### GetIsOnHoldOk

`func (o *V0BuildListAllResponseItemModel) GetIsOnHoldOk() (*bool, bool)`

GetIsOnHoldOk returns a tuple with the IsOnHold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOnHold

`func (o *V0BuildListAllResponseItemModel) SetIsOnHold(v bool)`

SetIsOnHold sets IsOnHold field to given value.

### HasIsOnHold

`func (o *V0BuildListAllResponseItemModel) HasIsOnHold() bool`

HasIsOnHold returns a boolean if a field has been set.

### GetIsProcessed

`func (o *V0BuildListAllResponseItemModel) GetIsProcessed() bool`

GetIsProcessed returns the IsProcessed field if non-nil, zero value otherwise.

### GetIsProcessedOk

`func (o *V0BuildListAllResponseItemModel) GetIsProcessedOk() (*bool, bool)`

GetIsProcessedOk returns a tuple with the IsProcessed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsProcessed

`func (o *V0BuildListAllResponseItemModel) SetIsProcessed(v bool)`

SetIsProcessed sets IsProcessed field to given value.

### HasIsProcessed

`func (o *V0BuildListAllResponseItemModel) HasIsProcessed() bool`

HasIsProcessed returns a boolean if a field has been set.

### GetMachineTypeId

`func (o *V0BuildListAllResponseItemModel) GetMachineTypeId() NullsString`

GetMachineTypeId returns the MachineTypeId field if non-nil, zero value otherwise.

### GetMachineTypeIdOk

`func (o *V0BuildListAllResponseItemModel) GetMachineTypeIdOk() (*NullsString, bool)`

GetMachineTypeIdOk returns a tuple with the MachineTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineTypeId

`func (o *V0BuildListAllResponseItemModel) SetMachineTypeId(v NullsString)`

SetMachineTypeId sets MachineTypeId field to given value.

### HasMachineTypeId

`func (o *V0BuildListAllResponseItemModel) HasMachineTypeId() bool`

HasMachineTypeId returns a boolean if a field has been set.

### GetOriginalBuildParams

`func (o *V0BuildListAllResponseItemModel) GetOriginalBuildParams() []int32`

GetOriginalBuildParams returns the OriginalBuildParams field if non-nil, zero value otherwise.

### GetOriginalBuildParamsOk

`func (o *V0BuildListAllResponseItemModel) GetOriginalBuildParamsOk() (*[]int32, bool)`

GetOriginalBuildParamsOk returns a tuple with the OriginalBuildParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalBuildParams

`func (o *V0BuildListAllResponseItemModel) SetOriginalBuildParams(v []int32)`

SetOriginalBuildParams sets OriginalBuildParams field to given value.

### HasOriginalBuildParams

`func (o *V0BuildListAllResponseItemModel) HasOriginalBuildParams() bool`

HasOriginalBuildParams returns a boolean if a field has been set.

### GetPullRequestId

`func (o *V0BuildListAllResponseItemModel) GetPullRequestId() int32`

GetPullRequestId returns the PullRequestId field if non-nil, zero value otherwise.

### GetPullRequestIdOk

`func (o *V0BuildListAllResponseItemModel) GetPullRequestIdOk() (*int32, bool)`

GetPullRequestIdOk returns a tuple with the PullRequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullRequestId

`func (o *V0BuildListAllResponseItemModel) SetPullRequestId(v int32)`

SetPullRequestId sets PullRequestId field to given value.

### HasPullRequestId

`func (o *V0BuildListAllResponseItemModel) HasPullRequestId() bool`

HasPullRequestId returns a boolean if a field has been set.

### GetPullRequestTargetBranch

`func (o *V0BuildListAllResponseItemModel) GetPullRequestTargetBranch() NullsString`

GetPullRequestTargetBranch returns the PullRequestTargetBranch field if non-nil, zero value otherwise.

### GetPullRequestTargetBranchOk

`func (o *V0BuildListAllResponseItemModel) GetPullRequestTargetBranchOk() (*NullsString, bool)`

GetPullRequestTargetBranchOk returns a tuple with the PullRequestTargetBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullRequestTargetBranch

`func (o *V0BuildListAllResponseItemModel) SetPullRequestTargetBranch(v NullsString)`

SetPullRequestTargetBranch sets PullRequestTargetBranch field to given value.

### HasPullRequestTargetBranch

`func (o *V0BuildListAllResponseItemModel) HasPullRequestTargetBranch() bool`

HasPullRequestTargetBranch returns a boolean if a field has been set.

### GetPullRequestViewUrl

`func (o *V0BuildListAllResponseItemModel) GetPullRequestViewUrl() NullsString`

GetPullRequestViewUrl returns the PullRequestViewUrl field if non-nil, zero value otherwise.

### GetPullRequestViewUrlOk

`func (o *V0BuildListAllResponseItemModel) GetPullRequestViewUrlOk() (*NullsString, bool)`

GetPullRequestViewUrlOk returns a tuple with the PullRequestViewUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullRequestViewUrl

`func (o *V0BuildListAllResponseItemModel) SetPullRequestViewUrl(v NullsString)`

SetPullRequestViewUrl sets PullRequestViewUrl field to given value.

### HasPullRequestViewUrl

`func (o *V0BuildListAllResponseItemModel) HasPullRequestViewUrl() bool`

HasPullRequestViewUrl returns a boolean if a field has been set.

### GetRepository

`func (o *V0BuildListAllResponseItemModel) GetRepository() V0AppResponseItemModel`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *V0BuildListAllResponseItemModel) GetRepositoryOk() (*V0AppResponseItemModel, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *V0BuildListAllResponseItemModel) SetRepository(v V0AppResponseItemModel)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *V0BuildListAllResponseItemModel) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### GetSlug

`func (o *V0BuildListAllResponseItemModel) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *V0BuildListAllResponseItemModel) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *V0BuildListAllResponseItemModel) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *V0BuildListAllResponseItemModel) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetStackIdentifier

`func (o *V0BuildListAllResponseItemModel) GetStackIdentifier() NullsString`

GetStackIdentifier returns the StackIdentifier field if non-nil, zero value otherwise.

### GetStackIdentifierOk

`func (o *V0BuildListAllResponseItemModel) GetStackIdentifierOk() (*NullsString, bool)`

GetStackIdentifierOk returns a tuple with the StackIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackIdentifier

`func (o *V0BuildListAllResponseItemModel) SetStackIdentifier(v NullsString)`

SetStackIdentifier sets StackIdentifier field to given value.

### HasStackIdentifier

`func (o *V0BuildListAllResponseItemModel) HasStackIdentifier() bool`

HasStackIdentifier returns a boolean if a field has been set.

### GetStartedOnWorkerAt

`func (o *V0BuildListAllResponseItemModel) GetStartedOnWorkerAt() string`

GetStartedOnWorkerAt returns the StartedOnWorkerAt field if non-nil, zero value otherwise.

### GetStartedOnWorkerAtOk

`func (o *V0BuildListAllResponseItemModel) GetStartedOnWorkerAtOk() (*string, bool)`

GetStartedOnWorkerAtOk returns a tuple with the StartedOnWorkerAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedOnWorkerAt

`func (o *V0BuildListAllResponseItemModel) SetStartedOnWorkerAt(v string)`

SetStartedOnWorkerAt sets StartedOnWorkerAt field to given value.

### HasStartedOnWorkerAt

`func (o *V0BuildListAllResponseItemModel) HasStartedOnWorkerAt() bool`

HasStartedOnWorkerAt returns a boolean if a field has been set.

### GetStatus

`func (o *V0BuildListAllResponseItemModel) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *V0BuildListAllResponseItemModel) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *V0BuildListAllResponseItemModel) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *V0BuildListAllResponseItemModel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusText

`func (o *V0BuildListAllResponseItemModel) GetStatusText() string`

GetStatusText returns the StatusText field if non-nil, zero value otherwise.

### GetStatusTextOk

`func (o *V0BuildListAllResponseItemModel) GetStatusTextOk() (*string, bool)`

GetStatusTextOk returns a tuple with the StatusText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusText

`func (o *V0BuildListAllResponseItemModel) SetStatusText(v string)`

SetStatusText sets StatusText field to given value.

### HasStatusText

`func (o *V0BuildListAllResponseItemModel) HasStatusText() bool`

HasStatusText returns a boolean if a field has been set.

### GetTag

`func (o *V0BuildListAllResponseItemModel) GetTag() NullsString`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *V0BuildListAllResponseItemModel) GetTagOk() (*NullsString, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *V0BuildListAllResponseItemModel) SetTag(v NullsString)`

SetTag sets Tag field to given value.

### HasTag

`func (o *V0BuildListAllResponseItemModel) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetTriggeredAt

`func (o *V0BuildListAllResponseItemModel) GetTriggeredAt() string`

GetTriggeredAt returns the TriggeredAt field if non-nil, zero value otherwise.

### GetTriggeredAtOk

`func (o *V0BuildListAllResponseItemModel) GetTriggeredAtOk() (*string, bool)`

GetTriggeredAtOk returns a tuple with the TriggeredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggeredAt

`func (o *V0BuildListAllResponseItemModel) SetTriggeredAt(v string)`

SetTriggeredAt sets TriggeredAt field to given value.

### HasTriggeredAt

`func (o *V0BuildListAllResponseItemModel) HasTriggeredAt() bool`

HasTriggeredAt returns a boolean if a field has been set.

### GetTriggeredBy

`func (o *V0BuildListAllResponseItemModel) GetTriggeredBy() NullsString`

GetTriggeredBy returns the TriggeredBy field if non-nil, zero value otherwise.

### GetTriggeredByOk

`func (o *V0BuildListAllResponseItemModel) GetTriggeredByOk() (*NullsString, bool)`

GetTriggeredByOk returns a tuple with the TriggeredBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggeredBy

`func (o *V0BuildListAllResponseItemModel) SetTriggeredBy(v NullsString)`

SetTriggeredBy sets TriggeredBy field to given value.

### HasTriggeredBy

`func (o *V0BuildListAllResponseItemModel) HasTriggeredBy() bool`

HasTriggeredBy returns a boolean if a field has been set.

### GetTriggeredWorkflow

`func (o *V0BuildListAllResponseItemModel) GetTriggeredWorkflow() string`

GetTriggeredWorkflow returns the TriggeredWorkflow field if non-nil, zero value otherwise.

### GetTriggeredWorkflowOk

`func (o *V0BuildListAllResponseItemModel) GetTriggeredWorkflowOk() (*string, bool)`

GetTriggeredWorkflowOk returns a tuple with the TriggeredWorkflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggeredWorkflow

`func (o *V0BuildListAllResponseItemModel) SetTriggeredWorkflow(v string)`

SetTriggeredWorkflow sets TriggeredWorkflow field to given value.

### HasTriggeredWorkflow

`func (o *V0BuildListAllResponseItemModel) HasTriggeredWorkflow() bool`

HasTriggeredWorkflow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


