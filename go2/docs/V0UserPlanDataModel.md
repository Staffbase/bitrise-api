# V0UserPlanDataModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentPlan** | Pointer to [**V0PlanDataModel**](V0PlanDataModel.md) |  | [optional] 
**PendingPlan** | Pointer to [**V0PlanDataModel**](V0PlanDataModel.md) |  | [optional] 
**TrialExpiresAt** | Pointer to **string** |  | [optional] 

## Methods

### NewV0UserPlanDataModel

`func NewV0UserPlanDataModel() *V0UserPlanDataModel`

NewV0UserPlanDataModel instantiates a new V0UserPlanDataModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0UserPlanDataModelWithDefaults

`func NewV0UserPlanDataModelWithDefaults() *V0UserPlanDataModel`

NewV0UserPlanDataModelWithDefaults instantiates a new V0UserPlanDataModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentPlan

`func (o *V0UserPlanDataModel) GetCurrentPlan() V0PlanDataModel`

GetCurrentPlan returns the CurrentPlan field if non-nil, zero value otherwise.

### GetCurrentPlanOk

`func (o *V0UserPlanDataModel) GetCurrentPlanOk() (*V0PlanDataModel, bool)`

GetCurrentPlanOk returns a tuple with the CurrentPlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPlan

`func (o *V0UserPlanDataModel) SetCurrentPlan(v V0PlanDataModel)`

SetCurrentPlan sets CurrentPlan field to given value.

### HasCurrentPlan

`func (o *V0UserPlanDataModel) HasCurrentPlan() bool`

HasCurrentPlan returns a boolean if a field has been set.

### GetPendingPlan

`func (o *V0UserPlanDataModel) GetPendingPlan() V0PlanDataModel`

GetPendingPlan returns the PendingPlan field if non-nil, zero value otherwise.

### GetPendingPlanOk

`func (o *V0UserPlanDataModel) GetPendingPlanOk() (*V0PlanDataModel, bool)`

GetPendingPlanOk returns a tuple with the PendingPlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingPlan

`func (o *V0UserPlanDataModel) SetPendingPlan(v V0PlanDataModel)`

SetPendingPlan sets PendingPlan field to given value.

### HasPendingPlan

`func (o *V0UserPlanDataModel) HasPendingPlan() bool`

HasPendingPlan returns a boolean if a field has been set.

### GetTrialExpiresAt

`func (o *V0UserPlanDataModel) GetTrialExpiresAt() string`

GetTrialExpiresAt returns the TrialExpiresAt field if non-nil, zero value otherwise.

### GetTrialExpiresAtOk

`func (o *V0UserPlanDataModel) GetTrialExpiresAtOk() (*string, bool)`

GetTrialExpiresAtOk returns a tuple with the TrialExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialExpiresAt

`func (o *V0UserPlanDataModel) SetTrialExpiresAt(v string)`

SetTrialExpiresAt sets TrialExpiresAt field to given value.

### HasTrialExpiresAt

`func (o *V0UserPlanDataModel) HasTrialExpiresAt() bool`

HasTrialExpiresAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


