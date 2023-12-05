# V0PagingResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Next** | Pointer to **string** | Next is the \&quot;anchor\&quot; for pagination. This value should be passed to the same endpoint to get the next page. Empty/not included if there&#39;s no \&quot;next\&quot; page. Stop paging when there&#39;s no \&quot;Next\&quot; item in the response! | [optional] 
**PageItemLimit** | Pointer to **int32** | PageItemLimit - per-page item count. A given page might include less items if there&#39;s not enough items. This value is the \&quot;max item count per page\&quot;. | [optional] 
**TotalItemCount** | Pointer to **int32** | TotalItemCount - total item count, through \&quot;all pages\&quot; | [optional] 

## Methods

### NewV0PagingResponseModel

`func NewV0PagingResponseModel() *V0PagingResponseModel`

NewV0PagingResponseModel instantiates a new V0PagingResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0PagingResponseModelWithDefaults

`func NewV0PagingResponseModelWithDefaults() *V0PagingResponseModel`

NewV0PagingResponseModelWithDefaults instantiates a new V0PagingResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNext

`func (o *V0PagingResponseModel) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *V0PagingResponseModel) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *V0PagingResponseModel) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *V0PagingResponseModel) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetPageItemLimit

`func (o *V0PagingResponseModel) GetPageItemLimit() int32`

GetPageItemLimit returns the PageItemLimit field if non-nil, zero value otherwise.

### GetPageItemLimitOk

`func (o *V0PagingResponseModel) GetPageItemLimitOk() (*int32, bool)`

GetPageItemLimitOk returns a tuple with the PageItemLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageItemLimit

`func (o *V0PagingResponseModel) SetPageItemLimit(v int32)`

SetPageItemLimit sets PageItemLimit field to given value.

### HasPageItemLimit

`func (o *V0PagingResponseModel) HasPageItemLimit() bool`

HasPageItemLimit returns a boolean if a field has been set.

### GetTotalItemCount

`func (o *V0PagingResponseModel) GetTotalItemCount() int32`

GetTotalItemCount returns the TotalItemCount field if non-nil, zero value otherwise.

### GetTotalItemCountOk

`func (o *V0PagingResponseModel) GetTotalItemCountOk() (*int32, bool)`

GetTotalItemCountOk returns a tuple with the TotalItemCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalItemCount

`func (o *V0PagingResponseModel) SetTotalItemCount(v int32)`

SetTotalItemCount sets TotalItemCount field to given value.

### HasTotalItemCount

`func (o *V0PagingResponseModel) HasTotalItemCount() bool`

HasTotalItemCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


