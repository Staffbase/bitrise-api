# V0PipelineListResponseItemModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Artifacts** | Pointer to [**[]PipelineabledomainArtifactMeta**](PipelineabledomainArtifactMeta.md) |  | [optional] 
**Branch** | Pointer to [**NullsString**](NullsString.md) |  | [optional] 
**BuildNumber** | Pointer to **int32** |  | [optional] 
**BuildToolInvocations** | Pointer to [**[]PipelineabledomainBuildToolInvocation**](PipelineabledomainBuildToolInvocation.md) |  | [optional] 
**CommitHash** | Pointer to [**NullsString**](NullsString.md) |  | [optional] 
**CommitMessage** | Pointer to [**NullsString**](NullsString.md) |  | [optional] 
**CreditCost** | Pointer to [**NullsInt64**](NullsInt64.md) |  | [optional] 
**FinishedAt** | Pointer to **string** |  | [optional] 
**IsOnHold** | Pointer to **bool** |  | [optional] 
**IsPipeline** | Pointer to **bool** |  | [optional] 
**IsProcessed** | Pointer to **bool** |  | [optional] 
**LocalConfig** | Pointer to [**PipelineabledomainLocalConfig**](PipelineabledomainLocalConfig.md) |  | [optional] 
**PullRequestId** | Pointer to **int32** |  | [optional] 
**PullRequestTargetBranch** | Pointer to [**NullsString**](NullsString.md) |  | [optional] 
**Slug** | Pointer to **string** |  | [optional] 
**StartedAt** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **int32** |  | [optional] 
**Tag** | Pointer to [**NullsString**](NullsString.md) |  | [optional] 
**TriggerParams** | Pointer to [**PipelineabledomainTriggerParams**](PipelineabledomainTriggerParams.md) |  | [optional] 
**TriggeredAt** | Pointer to **string** |  | [optional] 
**TriggeredBy** | Pointer to [**NullsString**](NullsString.md) |  | [optional] 
**TriggeredWorkflow** | Pointer to **string** |  | [optional] 

## Methods

### NewV0PipelineListResponseItemModel

`func NewV0PipelineListResponseItemModel() *V0PipelineListResponseItemModel`

NewV0PipelineListResponseItemModel instantiates a new V0PipelineListResponseItemModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0PipelineListResponseItemModelWithDefaults

`func NewV0PipelineListResponseItemModelWithDefaults() *V0PipelineListResponseItemModel`

NewV0PipelineListResponseItemModelWithDefaults instantiates a new V0PipelineListResponseItemModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArtifacts

`func (o *V0PipelineListResponseItemModel) GetArtifacts() []PipelineabledomainArtifactMeta`

GetArtifacts returns the Artifacts field if non-nil, zero value otherwise.

### GetArtifactsOk

`func (o *V0PipelineListResponseItemModel) GetArtifactsOk() (*[]PipelineabledomainArtifactMeta, bool)`

GetArtifactsOk returns a tuple with the Artifacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifacts

`func (o *V0PipelineListResponseItemModel) SetArtifacts(v []PipelineabledomainArtifactMeta)`

SetArtifacts sets Artifacts field to given value.

### HasArtifacts

`func (o *V0PipelineListResponseItemModel) HasArtifacts() bool`

HasArtifacts returns a boolean if a field has been set.

### GetBranch

`func (o *V0PipelineListResponseItemModel) GetBranch() NullsString`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *V0PipelineListResponseItemModel) GetBranchOk() (*NullsString, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *V0PipelineListResponseItemModel) SetBranch(v NullsString)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *V0PipelineListResponseItemModel) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### GetBuildNumber

`func (o *V0PipelineListResponseItemModel) GetBuildNumber() int32`

GetBuildNumber returns the BuildNumber field if non-nil, zero value otherwise.

### GetBuildNumberOk

`func (o *V0PipelineListResponseItemModel) GetBuildNumberOk() (*int32, bool)`

GetBuildNumberOk returns a tuple with the BuildNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildNumber

`func (o *V0PipelineListResponseItemModel) SetBuildNumber(v int32)`

SetBuildNumber sets BuildNumber field to given value.

### HasBuildNumber

`func (o *V0PipelineListResponseItemModel) HasBuildNumber() bool`

HasBuildNumber returns a boolean if a field has been set.

### GetBuildToolInvocations

`func (o *V0PipelineListResponseItemModel) GetBuildToolInvocations() []PipelineabledomainBuildToolInvocation`

GetBuildToolInvocations returns the BuildToolInvocations field if non-nil, zero value otherwise.

### GetBuildToolInvocationsOk

`func (o *V0PipelineListResponseItemModel) GetBuildToolInvocationsOk() (*[]PipelineabledomainBuildToolInvocation, bool)`

GetBuildToolInvocationsOk returns a tuple with the BuildToolInvocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildToolInvocations

`func (o *V0PipelineListResponseItemModel) SetBuildToolInvocations(v []PipelineabledomainBuildToolInvocation)`

SetBuildToolInvocations sets BuildToolInvocations field to given value.

### HasBuildToolInvocations

`func (o *V0PipelineListResponseItemModel) HasBuildToolInvocations() bool`

HasBuildToolInvocations returns a boolean if a field has been set.

### GetCommitHash

`func (o *V0PipelineListResponseItemModel) GetCommitHash() NullsString`

GetCommitHash returns the CommitHash field if non-nil, zero value otherwise.

### GetCommitHashOk

`func (o *V0PipelineListResponseItemModel) GetCommitHashOk() (*NullsString, bool)`

GetCommitHashOk returns a tuple with the CommitHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitHash

`func (o *V0PipelineListResponseItemModel) SetCommitHash(v NullsString)`

SetCommitHash sets CommitHash field to given value.

### HasCommitHash

`func (o *V0PipelineListResponseItemModel) HasCommitHash() bool`

HasCommitHash returns a boolean if a field has been set.

### GetCommitMessage

`func (o *V0PipelineListResponseItemModel) GetCommitMessage() NullsString`

GetCommitMessage returns the CommitMessage field if non-nil, zero value otherwise.

### GetCommitMessageOk

`func (o *V0PipelineListResponseItemModel) GetCommitMessageOk() (*NullsString, bool)`

GetCommitMessageOk returns a tuple with the CommitMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitMessage

`func (o *V0PipelineListResponseItemModel) SetCommitMessage(v NullsString)`

SetCommitMessage sets CommitMessage field to given value.

### HasCommitMessage

`func (o *V0PipelineListResponseItemModel) HasCommitMessage() bool`

HasCommitMessage returns a boolean if a field has been set.

### GetCreditCost

`func (o *V0PipelineListResponseItemModel) GetCreditCost() NullsInt64`

GetCreditCost returns the CreditCost field if non-nil, zero value otherwise.

### GetCreditCostOk

`func (o *V0PipelineListResponseItemModel) GetCreditCostOk() (*NullsInt64, bool)`

GetCreditCostOk returns a tuple with the CreditCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditCost

`func (o *V0PipelineListResponseItemModel) SetCreditCost(v NullsInt64)`

SetCreditCost sets CreditCost field to given value.

### HasCreditCost

`func (o *V0PipelineListResponseItemModel) HasCreditCost() bool`

HasCreditCost returns a boolean if a field has been set.

### GetFinishedAt

`func (o *V0PipelineListResponseItemModel) GetFinishedAt() string`

GetFinishedAt returns the FinishedAt field if non-nil, zero value otherwise.

### GetFinishedAtOk

`func (o *V0PipelineListResponseItemModel) GetFinishedAtOk() (*string, bool)`

GetFinishedAtOk returns a tuple with the FinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedAt

`func (o *V0PipelineListResponseItemModel) SetFinishedAt(v string)`

SetFinishedAt sets FinishedAt field to given value.

### HasFinishedAt

`func (o *V0PipelineListResponseItemModel) HasFinishedAt() bool`

HasFinishedAt returns a boolean if a field has been set.

### GetIsOnHold

`func (o *V0PipelineListResponseItemModel) GetIsOnHold() bool`

GetIsOnHold returns the IsOnHold field if non-nil, zero value otherwise.

### GetIsOnHoldOk

`func (o *V0PipelineListResponseItemModel) GetIsOnHoldOk() (*bool, bool)`

GetIsOnHoldOk returns a tuple with the IsOnHold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOnHold

`func (o *V0PipelineListResponseItemModel) SetIsOnHold(v bool)`

SetIsOnHold sets IsOnHold field to given value.

### HasIsOnHold

`func (o *V0PipelineListResponseItemModel) HasIsOnHold() bool`

HasIsOnHold returns a boolean if a field has been set.

### GetIsPipeline

`func (o *V0PipelineListResponseItemModel) GetIsPipeline() bool`

GetIsPipeline returns the IsPipeline field if non-nil, zero value otherwise.

### GetIsPipelineOk

`func (o *V0PipelineListResponseItemModel) GetIsPipelineOk() (*bool, bool)`

GetIsPipelineOk returns a tuple with the IsPipeline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPipeline

`func (o *V0PipelineListResponseItemModel) SetIsPipeline(v bool)`

SetIsPipeline sets IsPipeline field to given value.

### HasIsPipeline

`func (o *V0PipelineListResponseItemModel) HasIsPipeline() bool`

HasIsPipeline returns a boolean if a field has been set.

### GetIsProcessed

`func (o *V0PipelineListResponseItemModel) GetIsProcessed() bool`

GetIsProcessed returns the IsProcessed field if non-nil, zero value otherwise.

### GetIsProcessedOk

`func (o *V0PipelineListResponseItemModel) GetIsProcessedOk() (*bool, bool)`

GetIsProcessedOk returns a tuple with the IsProcessed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsProcessed

`func (o *V0PipelineListResponseItemModel) SetIsProcessed(v bool)`

SetIsProcessed sets IsProcessed field to given value.

### HasIsProcessed

`func (o *V0PipelineListResponseItemModel) HasIsProcessed() bool`

HasIsProcessed returns a boolean if a field has been set.

### GetLocalConfig

`func (o *V0PipelineListResponseItemModel) GetLocalConfig() PipelineabledomainLocalConfig`

GetLocalConfig returns the LocalConfig field if non-nil, zero value otherwise.

### GetLocalConfigOk

`func (o *V0PipelineListResponseItemModel) GetLocalConfigOk() (*PipelineabledomainLocalConfig, bool)`

GetLocalConfigOk returns a tuple with the LocalConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalConfig

`func (o *V0PipelineListResponseItemModel) SetLocalConfig(v PipelineabledomainLocalConfig)`

SetLocalConfig sets LocalConfig field to given value.

### HasLocalConfig

`func (o *V0PipelineListResponseItemModel) HasLocalConfig() bool`

HasLocalConfig returns a boolean if a field has been set.

### GetPullRequestId

`func (o *V0PipelineListResponseItemModel) GetPullRequestId() int32`

GetPullRequestId returns the PullRequestId field if non-nil, zero value otherwise.

### GetPullRequestIdOk

`func (o *V0PipelineListResponseItemModel) GetPullRequestIdOk() (*int32, bool)`

GetPullRequestIdOk returns a tuple with the PullRequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullRequestId

`func (o *V0PipelineListResponseItemModel) SetPullRequestId(v int32)`

SetPullRequestId sets PullRequestId field to given value.

### HasPullRequestId

`func (o *V0PipelineListResponseItemModel) HasPullRequestId() bool`

HasPullRequestId returns a boolean if a field has been set.

### GetPullRequestTargetBranch

`func (o *V0PipelineListResponseItemModel) GetPullRequestTargetBranch() NullsString`

GetPullRequestTargetBranch returns the PullRequestTargetBranch field if non-nil, zero value otherwise.

### GetPullRequestTargetBranchOk

`func (o *V0PipelineListResponseItemModel) GetPullRequestTargetBranchOk() (*NullsString, bool)`

GetPullRequestTargetBranchOk returns a tuple with the PullRequestTargetBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullRequestTargetBranch

`func (o *V0PipelineListResponseItemModel) SetPullRequestTargetBranch(v NullsString)`

SetPullRequestTargetBranch sets PullRequestTargetBranch field to given value.

### HasPullRequestTargetBranch

`func (o *V0PipelineListResponseItemModel) HasPullRequestTargetBranch() bool`

HasPullRequestTargetBranch returns a boolean if a field has been set.

### GetSlug

`func (o *V0PipelineListResponseItemModel) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *V0PipelineListResponseItemModel) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *V0PipelineListResponseItemModel) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *V0PipelineListResponseItemModel) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetStartedAt

`func (o *V0PipelineListResponseItemModel) GetStartedAt() string`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *V0PipelineListResponseItemModel) GetStartedAtOk() (*string, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *V0PipelineListResponseItemModel) SetStartedAt(v string)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *V0PipelineListResponseItemModel) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetStatus

`func (o *V0PipelineListResponseItemModel) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *V0PipelineListResponseItemModel) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *V0PipelineListResponseItemModel) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *V0PipelineListResponseItemModel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTag

`func (o *V0PipelineListResponseItemModel) GetTag() NullsString`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *V0PipelineListResponseItemModel) GetTagOk() (*NullsString, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *V0PipelineListResponseItemModel) SetTag(v NullsString)`

SetTag sets Tag field to given value.

### HasTag

`func (o *V0PipelineListResponseItemModel) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetTriggerParams

`func (o *V0PipelineListResponseItemModel) GetTriggerParams() PipelineabledomainTriggerParams`

GetTriggerParams returns the TriggerParams field if non-nil, zero value otherwise.

### GetTriggerParamsOk

`func (o *V0PipelineListResponseItemModel) GetTriggerParamsOk() (*PipelineabledomainTriggerParams, bool)`

GetTriggerParamsOk returns a tuple with the TriggerParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerParams

`func (o *V0PipelineListResponseItemModel) SetTriggerParams(v PipelineabledomainTriggerParams)`

SetTriggerParams sets TriggerParams field to given value.

### HasTriggerParams

`func (o *V0PipelineListResponseItemModel) HasTriggerParams() bool`

HasTriggerParams returns a boolean if a field has been set.

### GetTriggeredAt

`func (o *V0PipelineListResponseItemModel) GetTriggeredAt() string`

GetTriggeredAt returns the TriggeredAt field if non-nil, zero value otherwise.

### GetTriggeredAtOk

`func (o *V0PipelineListResponseItemModel) GetTriggeredAtOk() (*string, bool)`

GetTriggeredAtOk returns a tuple with the TriggeredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggeredAt

`func (o *V0PipelineListResponseItemModel) SetTriggeredAt(v string)`

SetTriggeredAt sets TriggeredAt field to given value.

### HasTriggeredAt

`func (o *V0PipelineListResponseItemModel) HasTriggeredAt() bool`

HasTriggeredAt returns a boolean if a field has been set.

### GetTriggeredBy

`func (o *V0PipelineListResponseItemModel) GetTriggeredBy() NullsString`

GetTriggeredBy returns the TriggeredBy field if non-nil, zero value otherwise.

### GetTriggeredByOk

`func (o *V0PipelineListResponseItemModel) GetTriggeredByOk() (*NullsString, bool)`

GetTriggeredByOk returns a tuple with the TriggeredBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggeredBy

`func (o *V0PipelineListResponseItemModel) SetTriggeredBy(v NullsString)`

SetTriggeredBy sets TriggeredBy field to given value.

### HasTriggeredBy

`func (o *V0PipelineListResponseItemModel) HasTriggeredBy() bool`

HasTriggeredBy returns a boolean if a field has been set.

### GetTriggeredWorkflow

`func (o *V0PipelineListResponseItemModel) GetTriggeredWorkflow() string`

GetTriggeredWorkflow returns the TriggeredWorkflow field if non-nil, zero value otherwise.

### GetTriggeredWorkflowOk

`func (o *V0PipelineListResponseItemModel) GetTriggeredWorkflowOk() (*string, bool)`

GetTriggeredWorkflowOk returns a tuple with the TriggeredWorkflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggeredWorkflow

`func (o *V0PipelineListResponseItemModel) SetTriggeredWorkflow(v string)`

SetTriggeredWorkflow sets TriggeredWorkflow field to given value.

### HasTriggeredWorkflow

`func (o *V0PipelineListResponseItemModel) HasTriggeredWorkflow() bool`

HasTriggeredWorkflow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


