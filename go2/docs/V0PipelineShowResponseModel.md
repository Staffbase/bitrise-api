# V0PipelineShowResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AbortReason** | Pointer to **string** |  | [optional] 
**App** | Pointer to [**V0PipelineShowAppResponseModel**](V0PipelineShowAppResponseModel.md) |  | [optional] 
**Attempts** | Pointer to [**[]V0PipelineShowAttemptResponseModel**](V0PipelineShowAttemptResponseModel.md) |  | [optional] 
**CreditCost** | Pointer to **map[string]interface{}** |  | [optional] 
**CurrentAttemptId** | Pointer to **string** |  | [optional] 
**FinishedAt** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NumberInAppScope** | Pointer to **int32** |  | [optional] 
**PutOnHoldAt** | Pointer to **map[string]interface{}** |  | [optional] 
**Stages** | Pointer to [**[]V0PipelineShowStageResponseModel**](V0PipelineShowStageResponseModel.md) |  | [optional] 
**StartedAt** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**TriggerParams** | Pointer to [**V0PipelineShowTriggerParamsResponseModel**](V0PipelineShowTriggerParamsResponseModel.md) |  | [optional] 
**TriggeredAt** | Pointer to **string** |  | [optional] 
**TriggeredBy** | Pointer to **string** |  | [optional] 

## Methods

### NewV0PipelineShowResponseModel

`func NewV0PipelineShowResponseModel() *V0PipelineShowResponseModel`

NewV0PipelineShowResponseModel instantiates a new V0PipelineShowResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0PipelineShowResponseModelWithDefaults

`func NewV0PipelineShowResponseModelWithDefaults() *V0PipelineShowResponseModel`

NewV0PipelineShowResponseModelWithDefaults instantiates a new V0PipelineShowResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAbortReason

`func (o *V0PipelineShowResponseModel) GetAbortReason() string`

GetAbortReason returns the AbortReason field if non-nil, zero value otherwise.

### GetAbortReasonOk

`func (o *V0PipelineShowResponseModel) GetAbortReasonOk() (*string, bool)`

GetAbortReasonOk returns a tuple with the AbortReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbortReason

`func (o *V0PipelineShowResponseModel) SetAbortReason(v string)`

SetAbortReason sets AbortReason field to given value.

### HasAbortReason

`func (o *V0PipelineShowResponseModel) HasAbortReason() bool`

HasAbortReason returns a boolean if a field has been set.

### GetApp

`func (o *V0PipelineShowResponseModel) GetApp() V0PipelineShowAppResponseModel`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *V0PipelineShowResponseModel) GetAppOk() (*V0PipelineShowAppResponseModel, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *V0PipelineShowResponseModel) SetApp(v V0PipelineShowAppResponseModel)`

SetApp sets App field to given value.

### HasApp

`func (o *V0PipelineShowResponseModel) HasApp() bool`

HasApp returns a boolean if a field has been set.

### GetAttempts

`func (o *V0PipelineShowResponseModel) GetAttempts() []V0PipelineShowAttemptResponseModel`

GetAttempts returns the Attempts field if non-nil, zero value otherwise.

### GetAttemptsOk

`func (o *V0PipelineShowResponseModel) GetAttemptsOk() (*[]V0PipelineShowAttemptResponseModel, bool)`

GetAttemptsOk returns a tuple with the Attempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttempts

`func (o *V0PipelineShowResponseModel) SetAttempts(v []V0PipelineShowAttemptResponseModel)`

SetAttempts sets Attempts field to given value.

### HasAttempts

`func (o *V0PipelineShowResponseModel) HasAttempts() bool`

HasAttempts returns a boolean if a field has been set.

### GetCreditCost

`func (o *V0PipelineShowResponseModel) GetCreditCost() map[string]interface{}`

GetCreditCost returns the CreditCost field if non-nil, zero value otherwise.

### GetCreditCostOk

`func (o *V0PipelineShowResponseModel) GetCreditCostOk() (*map[string]interface{}, bool)`

GetCreditCostOk returns a tuple with the CreditCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditCost

`func (o *V0PipelineShowResponseModel) SetCreditCost(v map[string]interface{})`

SetCreditCost sets CreditCost field to given value.

### HasCreditCost

`func (o *V0PipelineShowResponseModel) HasCreditCost() bool`

HasCreditCost returns a boolean if a field has been set.

### GetCurrentAttemptId

`func (o *V0PipelineShowResponseModel) GetCurrentAttemptId() string`

GetCurrentAttemptId returns the CurrentAttemptId field if non-nil, zero value otherwise.

### GetCurrentAttemptIdOk

`func (o *V0PipelineShowResponseModel) GetCurrentAttemptIdOk() (*string, bool)`

GetCurrentAttemptIdOk returns a tuple with the CurrentAttemptId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentAttemptId

`func (o *V0PipelineShowResponseModel) SetCurrentAttemptId(v string)`

SetCurrentAttemptId sets CurrentAttemptId field to given value.

### HasCurrentAttemptId

`func (o *V0PipelineShowResponseModel) HasCurrentAttemptId() bool`

HasCurrentAttemptId returns a boolean if a field has been set.

### GetFinishedAt

`func (o *V0PipelineShowResponseModel) GetFinishedAt() string`

GetFinishedAt returns the FinishedAt field if non-nil, zero value otherwise.

### GetFinishedAtOk

`func (o *V0PipelineShowResponseModel) GetFinishedAtOk() (*string, bool)`

GetFinishedAtOk returns a tuple with the FinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedAt

`func (o *V0PipelineShowResponseModel) SetFinishedAt(v string)`

SetFinishedAt sets FinishedAt field to given value.

### HasFinishedAt

`func (o *V0PipelineShowResponseModel) HasFinishedAt() bool`

HasFinishedAt returns a boolean if a field has been set.

### GetId

`func (o *V0PipelineShowResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V0PipelineShowResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V0PipelineShowResponseModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *V0PipelineShowResponseModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *V0PipelineShowResponseModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V0PipelineShowResponseModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V0PipelineShowResponseModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V0PipelineShowResponseModel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNumberInAppScope

`func (o *V0PipelineShowResponseModel) GetNumberInAppScope() int32`

GetNumberInAppScope returns the NumberInAppScope field if non-nil, zero value otherwise.

### GetNumberInAppScopeOk

`func (o *V0PipelineShowResponseModel) GetNumberInAppScopeOk() (*int32, bool)`

GetNumberInAppScopeOk returns a tuple with the NumberInAppScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberInAppScope

`func (o *V0PipelineShowResponseModel) SetNumberInAppScope(v int32)`

SetNumberInAppScope sets NumberInAppScope field to given value.

### HasNumberInAppScope

`func (o *V0PipelineShowResponseModel) HasNumberInAppScope() bool`

HasNumberInAppScope returns a boolean if a field has been set.

### GetPutOnHoldAt

`func (o *V0PipelineShowResponseModel) GetPutOnHoldAt() map[string]interface{}`

GetPutOnHoldAt returns the PutOnHoldAt field if non-nil, zero value otherwise.

### GetPutOnHoldAtOk

`func (o *V0PipelineShowResponseModel) GetPutOnHoldAtOk() (*map[string]interface{}, bool)`

GetPutOnHoldAtOk returns a tuple with the PutOnHoldAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPutOnHoldAt

`func (o *V0PipelineShowResponseModel) SetPutOnHoldAt(v map[string]interface{})`

SetPutOnHoldAt sets PutOnHoldAt field to given value.

### HasPutOnHoldAt

`func (o *V0PipelineShowResponseModel) HasPutOnHoldAt() bool`

HasPutOnHoldAt returns a boolean if a field has been set.

### GetStages

`func (o *V0PipelineShowResponseModel) GetStages() []V0PipelineShowStageResponseModel`

GetStages returns the Stages field if non-nil, zero value otherwise.

### GetStagesOk

`func (o *V0PipelineShowResponseModel) GetStagesOk() (*[]V0PipelineShowStageResponseModel, bool)`

GetStagesOk returns a tuple with the Stages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStages

`func (o *V0PipelineShowResponseModel) SetStages(v []V0PipelineShowStageResponseModel)`

SetStages sets Stages field to given value.

### HasStages

`func (o *V0PipelineShowResponseModel) HasStages() bool`

HasStages returns a boolean if a field has been set.

### GetStartedAt

`func (o *V0PipelineShowResponseModel) GetStartedAt() string`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *V0PipelineShowResponseModel) GetStartedAtOk() (*string, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *V0PipelineShowResponseModel) SetStartedAt(v string)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *V0PipelineShowResponseModel) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetStatus

`func (o *V0PipelineShowResponseModel) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *V0PipelineShowResponseModel) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *V0PipelineShowResponseModel) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *V0PipelineShowResponseModel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTriggerParams

`func (o *V0PipelineShowResponseModel) GetTriggerParams() V0PipelineShowTriggerParamsResponseModel`

GetTriggerParams returns the TriggerParams field if non-nil, zero value otherwise.

### GetTriggerParamsOk

`func (o *V0PipelineShowResponseModel) GetTriggerParamsOk() (*V0PipelineShowTriggerParamsResponseModel, bool)`

GetTriggerParamsOk returns a tuple with the TriggerParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerParams

`func (o *V0PipelineShowResponseModel) SetTriggerParams(v V0PipelineShowTriggerParamsResponseModel)`

SetTriggerParams sets TriggerParams field to given value.

### HasTriggerParams

`func (o *V0PipelineShowResponseModel) HasTriggerParams() bool`

HasTriggerParams returns a boolean if a field has been set.

### GetTriggeredAt

`func (o *V0PipelineShowResponseModel) GetTriggeredAt() string`

GetTriggeredAt returns the TriggeredAt field if non-nil, zero value otherwise.

### GetTriggeredAtOk

`func (o *V0PipelineShowResponseModel) GetTriggeredAtOk() (*string, bool)`

GetTriggeredAtOk returns a tuple with the TriggeredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggeredAt

`func (o *V0PipelineShowResponseModel) SetTriggeredAt(v string)`

SetTriggeredAt sets TriggeredAt field to given value.

### HasTriggeredAt

`func (o *V0PipelineShowResponseModel) HasTriggeredAt() bool`

HasTriggeredAt returns a boolean if a field has been set.

### GetTriggeredBy

`func (o *V0PipelineShowResponseModel) GetTriggeredBy() string`

GetTriggeredBy returns the TriggeredBy field if non-nil, zero value otherwise.

### GetTriggeredByOk

`func (o *V0PipelineShowResponseModel) GetTriggeredByOk() (*string, bool)`

GetTriggeredByOk returns a tuple with the TriggeredBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggeredBy

`func (o *V0PipelineShowResponseModel) SetTriggeredBy(v string)`

SetTriggeredBy sets TriggeredBy field to given value.

### HasTriggeredBy

`func (o *V0PipelineShowResponseModel) HasTriggeredBy() bool`

HasTriggeredBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


