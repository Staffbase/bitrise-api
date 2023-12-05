# V0ProjectFileStorageListResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]V0ProjectFileStorageResponseItemModel**](V0ProjectFileStorageResponseItemModel.md) |  | [optional] 
**Paging** | Pointer to [**V0PagingResponseModel**](V0PagingResponseModel.md) |  | [optional] 

## Methods

### NewV0ProjectFileStorageListResponseModel

`func NewV0ProjectFileStorageListResponseModel() *V0ProjectFileStorageListResponseModel`

NewV0ProjectFileStorageListResponseModel instantiates a new V0ProjectFileStorageListResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0ProjectFileStorageListResponseModelWithDefaults

`func NewV0ProjectFileStorageListResponseModelWithDefaults() *V0ProjectFileStorageListResponseModel`

NewV0ProjectFileStorageListResponseModelWithDefaults instantiates a new V0ProjectFileStorageListResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *V0ProjectFileStorageListResponseModel) GetData() []V0ProjectFileStorageResponseItemModel`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *V0ProjectFileStorageListResponseModel) GetDataOk() (*[]V0ProjectFileStorageResponseItemModel, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *V0ProjectFileStorageListResponseModel) SetData(v []V0ProjectFileStorageResponseItemModel)`

SetData sets Data field to given value.

### HasData

`func (o *V0ProjectFileStorageListResponseModel) HasData() bool`

HasData returns a boolean if a field has been set.

### GetPaging

`func (o *V0ProjectFileStorageListResponseModel) GetPaging() V0PagingResponseModel`

GetPaging returns the Paging field if non-nil, zero value otherwise.

### GetPagingOk

`func (o *V0ProjectFileStorageListResponseModel) GetPagingOk() (*V0PagingResponseModel, bool)`

GetPagingOk returns a tuple with the Paging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaging

`func (o *V0ProjectFileStorageListResponseModel) SetPaging(v V0PagingResponseModel)`

SetPaging sets Paging field to given value.

### HasPaging

`func (o *V0ProjectFileStorageListResponseModel) HasPaging() bool`

HasPaging returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


