# V0PipelineShowStageResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Workflows** | Pointer to [**[]V0PipelineShowWorkflowResponseModel**](V0PipelineShowWorkflowResponseModel.md) |  | [optional] 

## Methods

### NewV0PipelineShowStageResponseModel

`func NewV0PipelineShowStageResponseModel() *V0PipelineShowStageResponseModel`

NewV0PipelineShowStageResponseModel instantiates a new V0PipelineShowStageResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0PipelineShowStageResponseModelWithDefaults

`func NewV0PipelineShowStageResponseModelWithDefaults() *V0PipelineShowStageResponseModel`

NewV0PipelineShowStageResponseModelWithDefaults instantiates a new V0PipelineShowStageResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *V0PipelineShowStageResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V0PipelineShowStageResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V0PipelineShowStageResponseModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *V0PipelineShowStageResponseModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *V0PipelineShowStageResponseModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V0PipelineShowStageResponseModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V0PipelineShowStageResponseModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V0PipelineShowStageResponseModel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetWorkflows

`func (o *V0PipelineShowStageResponseModel) GetWorkflows() []V0PipelineShowWorkflowResponseModel`

GetWorkflows returns the Workflows field if non-nil, zero value otherwise.

### GetWorkflowsOk

`func (o *V0PipelineShowStageResponseModel) GetWorkflowsOk() (*[]V0PipelineShowWorkflowResponseModel, bool)`

GetWorkflowsOk returns a tuple with the Workflows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflows

`func (o *V0PipelineShowStageResponseModel) SetWorkflows(v []V0PipelineShowWorkflowResponseModel)`

SetWorkflows sets Workflows field to given value.

### HasWorkflows

`func (o *V0PipelineShowStageResponseModel) HasWorkflows() bool`

HasWorkflows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


