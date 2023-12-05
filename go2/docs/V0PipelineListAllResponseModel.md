# V0PipelineListAllResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]V0PipelineListAllResponseItemModel**](V0PipelineListAllResponseItemModel.md) |  | [optional] 
**Paging** | Pointer to [**V0PagingResponseModel**](V0PagingResponseModel.md) |  | [optional] 

## Methods

### NewV0PipelineListAllResponseModel

`func NewV0PipelineListAllResponseModel() *V0PipelineListAllResponseModel`

NewV0PipelineListAllResponseModel instantiates a new V0PipelineListAllResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0PipelineListAllResponseModelWithDefaults

`func NewV0PipelineListAllResponseModelWithDefaults() *V0PipelineListAllResponseModel`

NewV0PipelineListAllResponseModelWithDefaults instantiates a new V0PipelineListAllResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *V0PipelineListAllResponseModel) GetData() []V0PipelineListAllResponseItemModel`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *V0PipelineListAllResponseModel) GetDataOk() (*[]V0PipelineListAllResponseItemModel, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *V0PipelineListAllResponseModel) SetData(v []V0PipelineListAllResponseItemModel)`

SetData sets Data field to given value.

### HasData

`func (o *V0PipelineListAllResponseModel) HasData() bool`

HasData returns a boolean if a field has been set.

### GetPaging

`func (o *V0PipelineListAllResponseModel) GetPaging() V0PagingResponseModel`

GetPaging returns the Paging field if non-nil, zero value otherwise.

### GetPagingOk

`func (o *V0PipelineListAllResponseModel) GetPagingOk() (*V0PagingResponseModel, bool)`

GetPagingOk returns a tuple with the Paging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaging

`func (o *V0PipelineListAllResponseModel) SetPaging(v V0PagingResponseModel)`

SetPaging sets Paging field to given value.

### HasPaging

`func (o *V0PipelineListAllResponseModel) HasPaging() bool`

HasPaging returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


