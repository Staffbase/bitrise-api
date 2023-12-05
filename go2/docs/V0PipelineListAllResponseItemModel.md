# V0PipelineListAllResponseItemModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Branch** | Pointer to [**NullsString**](NullsString.md) |  | [optional] 
**BuildNumber** | Pointer to **int32** |  | [optional] 
**CommitHash** | Pointer to [**NullsString**](NullsString.md) |  | [optional] 
**CommitMessage** | Pointer to [**NullsString**](NullsString.md) |  | [optional] 
**CreditCost** | Pointer to [**NullsInt64**](NullsInt64.md) |  | [optional] 
**FinishedAt** | Pointer to **string** |  | [optional] 
**IsOnHold** | Pointer to **bool** |  | [optional] 
**IsPipeline** | Pointer to **bool** |  | [optional] 
**IsProcessed** | Pointer to **bool** |  | [optional] 
**PullRequestId** | Pointer to **int32** |  | [optional] 
**PullRequestTargetBranch** | Pointer to [**NullsString**](NullsString.md) |  | [optional] 
**Repository** | Pointer to [**V0AppResponseItemModel**](V0AppResponseItemModel.md) |  | [optional] 
**Slug** | Pointer to **string** |  | [optional] 
**StartedAt** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **int32** |  | [optional] 
**Tag** | Pointer to [**NullsString**](NullsString.md) |  | [optional] 
**TriggeredAt** | Pointer to **string** |  | [optional] 
**TriggeredBy** | Pointer to [**NullsString**](NullsString.md) |  | [optional] 
**TriggeredWorkflow** | Pointer to **string** |  | [optional] 

## Methods

### NewV0PipelineListAllResponseItemModel

`func NewV0PipelineListAllResponseItemModel() *V0PipelineListAllResponseItemModel`

NewV0PipelineListAllResponseItemModel instantiates a new V0PipelineListAllResponseItemModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0PipelineListAllResponseItemModelWithDefaults

`func NewV0PipelineListAllResponseItemModelWithDefaults() *V0PipelineListAllResponseItemModel`

NewV0PipelineListAllResponseItemModelWithDefaults instantiates a new V0PipelineListAllResponseItemModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBranch

`func (o *V0PipelineListAllResponseItemModel) GetBranch() NullsString`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *V0PipelineListAllResponseItemModel) GetBranchOk() (*NullsString, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *V0PipelineListAllResponseItemModel) SetBranch(v NullsString)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *V0PipelineListAllResponseItemModel) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### GetBuildNumber

`func (o *V0PipelineListAllResponseItemModel) GetBuildNumber() int32`

GetBuildNumber returns the BuildNumber field if non-nil, zero value otherwise.

### GetBuildNumberOk

`func (o *V0PipelineListAllResponseItemModel) GetBuildNumberOk() (*int32, bool)`

GetBuildNumberOk returns a tuple with the BuildNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildNumber

`func (o *V0PipelineListAllResponseItemModel) SetBuildNumber(v int32)`

SetBuildNumber sets BuildNumber field to given value.

### HasBuildNumber

`func (o *V0PipelineListAllResponseItemModel) HasBuildNumber() bool`

HasBuildNumber returns a boolean if a field has been set.

### GetCommitHash

`func (o *V0PipelineListAllResponseItemModel) GetCommitHash() NullsString`

GetCommitHash returns the CommitHash field if non-nil, zero value otherwise.

### GetCommitHashOk

`func (o *V0PipelineListAllResponseItemModel) GetCommitHashOk() (*NullsString, bool)`

GetCommitHashOk returns a tuple with the CommitHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitHash

`func (o *V0PipelineListAllResponseItemModel) SetCommitHash(v NullsString)`

SetCommitHash sets CommitHash field to given value.

### HasCommitHash

`func (o *V0PipelineListAllResponseItemModel) HasCommitHash() bool`

HasCommitHash returns a boolean if a field has been set.

### GetCommitMessage

`func (o *V0PipelineListAllResponseItemModel) GetCommitMessage() NullsString`

GetCommitMessage returns the CommitMessage field if non-nil, zero value otherwise.

### GetCommitMessageOk

`func (o *V0PipelineListAllResponseItemModel) GetCommitMessageOk() (*NullsString, bool)`

GetCommitMessageOk returns a tuple with the CommitMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitMessage

`func (o *V0PipelineListAllResponseItemModel) SetCommitMessage(v NullsString)`

SetCommitMessage sets CommitMessage field to given value.

### HasCommitMessage

`func (o *V0PipelineListAllResponseItemModel) HasCommitMessage() bool`

HasCommitMessage returns a boolean if a field has been set.

### GetCreditCost

`func (o *V0PipelineListAllResponseItemModel) GetCreditCost() NullsInt64`

GetCreditCost returns the CreditCost field if non-nil, zero value otherwise.

### GetCreditCostOk

`func (o *V0PipelineListAllResponseItemModel) GetCreditCostOk() (*NullsInt64, bool)`

GetCreditCostOk returns a tuple with the CreditCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditCost

`func (o *V0PipelineListAllResponseItemModel) SetCreditCost(v NullsInt64)`

SetCreditCost sets CreditCost field to given value.

### HasCreditCost

`func (o *V0PipelineListAllResponseItemModel) HasCreditCost() bool`

HasCreditCost returns a boolean if a field has been set.

### GetFinishedAt

`func (o *V0PipelineListAllResponseItemModel) GetFinishedAt() string`

GetFinishedAt returns the FinishedAt field if non-nil, zero value otherwise.

### GetFinishedAtOk

`func (o *V0PipelineListAllResponseItemModel) GetFinishedAtOk() (*string, bool)`

GetFinishedAtOk returns a tuple with the FinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedAt

`func (o *V0PipelineListAllResponseItemModel) SetFinishedAt(v string)`

SetFinishedAt sets FinishedAt field to given value.

### HasFinishedAt

`func (o *V0PipelineListAllResponseItemModel) HasFinishedAt() bool`

HasFinishedAt returns a boolean if a field has been set.

### GetIsOnHold

`func (o *V0PipelineListAllResponseItemModel) GetIsOnHold() bool`

GetIsOnHold returns the IsOnHold field if non-nil, zero value otherwise.

### GetIsOnHoldOk

`func (o *V0PipelineListAllResponseItemModel) GetIsOnHoldOk() (*bool, bool)`

GetIsOnHoldOk returns a tuple with the IsOnHold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOnHold

`func (o *V0PipelineListAllResponseItemModel) SetIsOnHold(v bool)`

SetIsOnHold sets IsOnHold field to given value.

### HasIsOnHold

`func (o *V0PipelineListAllResponseItemModel) HasIsOnHold() bool`

HasIsOnHold returns a boolean if a field has been set.

### GetIsPipeline

`func (o *V0PipelineListAllResponseItemModel) GetIsPipeline() bool`

GetIsPipeline returns the IsPipeline field if non-nil, zero value otherwise.

### GetIsPipelineOk

`func (o *V0PipelineListAllResponseItemModel) GetIsPipelineOk() (*bool, bool)`

GetIsPipelineOk returns a tuple with the IsPipeline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPipeline

`func (o *V0PipelineListAllResponseItemModel) SetIsPipeline(v bool)`

SetIsPipeline sets IsPipeline field to given value.

### HasIsPipeline

`func (o *V0PipelineListAllResponseItemModel) HasIsPipeline() bool`

HasIsPipeline returns a boolean if a field has been set.

### GetIsProcessed

`func (o *V0PipelineListAllResponseItemModel) GetIsProcessed() bool`

GetIsProcessed returns the IsProcessed field if non-nil, zero value otherwise.

### GetIsProcessedOk

`func (o *V0PipelineListAllResponseItemModel) GetIsProcessedOk() (*bool, bool)`

GetIsProcessedOk returns a tuple with the IsProcessed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsProcessed

`func (o *V0PipelineListAllResponseItemModel) SetIsProcessed(v bool)`

SetIsProcessed sets IsProcessed field to given value.

### HasIsProcessed

`func (o *V0PipelineListAllResponseItemModel) HasIsProcessed() bool`

HasIsProcessed returns a boolean if a field has been set.

### GetPullRequestId

`func (o *V0PipelineListAllResponseItemModel) GetPullRequestId() int32`

GetPullRequestId returns the PullRequestId field if non-nil, zero value otherwise.

### GetPullRequestIdOk

`func (o *V0PipelineListAllResponseItemModel) GetPullRequestIdOk() (*int32, bool)`

GetPullRequestIdOk returns a tuple with the PullRequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullRequestId

`func (o *V0PipelineListAllResponseItemModel) SetPullRequestId(v int32)`

SetPullRequestId sets PullRequestId field to given value.

### HasPullRequestId

`func (o *V0PipelineListAllResponseItemModel) HasPullRequestId() bool`

HasPullRequestId returns a boolean if a field has been set.

### GetPullRequestTargetBranch

`func (o *V0PipelineListAllResponseItemModel) GetPullRequestTargetBranch() NullsString`

GetPullRequestTargetBranch returns the PullRequestTargetBranch field if non-nil, zero value otherwise.

### GetPullRequestTargetBranchOk

`func (o *V0PipelineListAllResponseItemModel) GetPullRequestTargetBranchOk() (*NullsString, bool)`

GetPullRequestTargetBranchOk returns a tuple with the PullRequestTargetBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullRequestTargetBranch

`func (o *V0PipelineListAllResponseItemModel) SetPullRequestTargetBranch(v NullsString)`

SetPullRequestTargetBranch sets PullRequestTargetBranch field to given value.

### HasPullRequestTargetBranch

`func (o *V0PipelineListAllResponseItemModel) HasPullRequestTargetBranch() bool`

HasPullRequestTargetBranch returns a boolean if a field has been set.

### GetRepository

`func (o *V0PipelineListAllResponseItemModel) GetRepository() V0AppResponseItemModel`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *V0PipelineListAllResponseItemModel) GetRepositoryOk() (*V0AppResponseItemModel, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *V0PipelineListAllResponseItemModel) SetRepository(v V0AppResponseItemModel)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *V0PipelineListAllResponseItemModel) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### GetSlug

`func (o *V0PipelineListAllResponseItemModel) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *V0PipelineListAllResponseItemModel) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *V0PipelineListAllResponseItemModel) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *V0PipelineListAllResponseItemModel) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetStartedAt

`func (o *V0PipelineListAllResponseItemModel) GetStartedAt() string`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *V0PipelineListAllResponseItemModel) GetStartedAtOk() (*string, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *V0PipelineListAllResponseItemModel) SetStartedAt(v string)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *V0PipelineListAllResponseItemModel) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetStatus

`func (o *V0PipelineListAllResponseItemModel) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *V0PipelineListAllResponseItemModel) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *V0PipelineListAllResponseItemModel) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *V0PipelineListAllResponseItemModel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTag

`func (o *V0PipelineListAllResponseItemModel) GetTag() NullsString`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *V0PipelineListAllResponseItemModel) GetTagOk() (*NullsString, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *V0PipelineListAllResponseItemModel) SetTag(v NullsString)`

SetTag sets Tag field to given value.

### HasTag

`func (o *V0PipelineListAllResponseItemModel) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetTriggeredAt

`func (o *V0PipelineListAllResponseItemModel) GetTriggeredAt() string`

GetTriggeredAt returns the TriggeredAt field if non-nil, zero value otherwise.

### GetTriggeredAtOk

`func (o *V0PipelineListAllResponseItemModel) GetTriggeredAtOk() (*string, bool)`

GetTriggeredAtOk returns a tuple with the TriggeredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggeredAt

`func (o *V0PipelineListAllResponseItemModel) SetTriggeredAt(v string)`

SetTriggeredAt sets TriggeredAt field to given value.

### HasTriggeredAt

`func (o *V0PipelineListAllResponseItemModel) HasTriggeredAt() bool`

HasTriggeredAt returns a boolean if a field has been set.

### GetTriggeredBy

`func (o *V0PipelineListAllResponseItemModel) GetTriggeredBy() NullsString`

GetTriggeredBy returns the TriggeredBy field if non-nil, zero value otherwise.

### GetTriggeredByOk

`func (o *V0PipelineListAllResponseItemModel) GetTriggeredByOk() (*NullsString, bool)`

GetTriggeredByOk returns a tuple with the TriggeredBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggeredBy

`func (o *V0PipelineListAllResponseItemModel) SetTriggeredBy(v NullsString)`

SetTriggeredBy sets TriggeredBy field to given value.

### HasTriggeredBy

`func (o *V0PipelineListAllResponseItemModel) HasTriggeredBy() bool`

HasTriggeredBy returns a boolean if a field has been set.

### GetTriggeredWorkflow

`func (o *V0PipelineListAllResponseItemModel) GetTriggeredWorkflow() string`

GetTriggeredWorkflow returns the TriggeredWorkflow field if non-nil, zero value otherwise.

### GetTriggeredWorkflowOk

`func (o *V0PipelineListAllResponseItemModel) GetTriggeredWorkflowOk() (*string, bool)`

GetTriggeredWorkflowOk returns a tuple with the TriggeredWorkflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggeredWorkflow

`func (o *V0PipelineListAllResponseItemModel) SetTriggeredWorkflow(v string)`

SetTriggeredWorkflow sets TriggeredWorkflow field to given value.

### HasTriggeredWorkflow

`func (o *V0PipelineListAllResponseItemModel) HasTriggeredWorkflow() bool`

HasTriggeredWorkflow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


