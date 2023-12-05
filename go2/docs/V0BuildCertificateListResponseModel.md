# V0BuildCertificateListResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]V0BuildCertificateResponseItemModel**](V0BuildCertificateResponseItemModel.md) |  | [optional] 
**Paging** | Pointer to [**V0PagingResponseModel**](V0PagingResponseModel.md) |  | [optional] 

## Methods

### NewV0BuildCertificateListResponseModel

`func NewV0BuildCertificateListResponseModel() *V0BuildCertificateListResponseModel`

NewV0BuildCertificateListResponseModel instantiates a new V0BuildCertificateListResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0BuildCertificateListResponseModelWithDefaults

`func NewV0BuildCertificateListResponseModelWithDefaults() *V0BuildCertificateListResponseModel`

NewV0BuildCertificateListResponseModelWithDefaults instantiates a new V0BuildCertificateListResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *V0BuildCertificateListResponseModel) GetData() []V0BuildCertificateResponseItemModel`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *V0BuildCertificateListResponseModel) GetDataOk() (*[]V0BuildCertificateResponseItemModel, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *V0BuildCertificateListResponseModel) SetData(v []V0BuildCertificateResponseItemModel)`

SetData sets Data field to given value.

### HasData

`func (o *V0BuildCertificateListResponseModel) HasData() bool`

HasData returns a boolean if a field has been set.

### GetPaging

`func (o *V0BuildCertificateListResponseModel) GetPaging() V0PagingResponseModel`

GetPaging returns the Paging field if non-nil, zero value otherwise.

### GetPagingOk

`func (o *V0BuildCertificateListResponseModel) GetPagingOk() (*V0PagingResponseModel, bool)`

GetPagingOk returns a tuple with the Paging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaging

`func (o *V0BuildCertificateListResponseModel) SetPaging(v V0PagingResponseModel)`

SetPaging sets Paging field to given value.

### HasPaging

`func (o *V0BuildCertificateListResponseModel) HasPaging() bool`

HasPaging returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


