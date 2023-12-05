# V0ProvisionProfileListResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]V0ProvisionProfileResponseItemModel**](V0ProvisionProfileResponseItemModel.md) |  | [optional] 
**Paging** | Pointer to [**V0PagingResponseModel**](V0PagingResponseModel.md) |  | [optional] 

## Methods

### NewV0ProvisionProfileListResponseModel

`func NewV0ProvisionProfileListResponseModel() *V0ProvisionProfileListResponseModel`

NewV0ProvisionProfileListResponseModel instantiates a new V0ProvisionProfileListResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0ProvisionProfileListResponseModelWithDefaults

`func NewV0ProvisionProfileListResponseModelWithDefaults() *V0ProvisionProfileListResponseModel`

NewV0ProvisionProfileListResponseModelWithDefaults instantiates a new V0ProvisionProfileListResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *V0ProvisionProfileListResponseModel) GetData() []V0ProvisionProfileResponseItemModel`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *V0ProvisionProfileListResponseModel) GetDataOk() (*[]V0ProvisionProfileResponseItemModel, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *V0ProvisionProfileListResponseModel) SetData(v []V0ProvisionProfileResponseItemModel)`

SetData sets Data field to given value.

### HasData

`func (o *V0ProvisionProfileListResponseModel) HasData() bool`

HasData returns a boolean if a field has been set.

### GetPaging

`func (o *V0ProvisionProfileListResponseModel) GetPaging() V0PagingResponseModel`

GetPaging returns the Paging field if non-nil, zero value otherwise.

### GetPagingOk

`func (o *V0ProvisionProfileListResponseModel) GetPagingOk() (*V0PagingResponseModel, bool)`

GetPagingOk returns a tuple with the Paging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaging

`func (o *V0ProvisionProfileListResponseModel) SetPaging(v V0PagingResponseModel)`

SetPaging sets Paging field to given value.

### HasPaging

`func (o *V0ProvisionProfileListResponseModel) HasPaging() bool`

HasPaging returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


