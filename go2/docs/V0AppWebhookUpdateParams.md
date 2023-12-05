# V0AppWebhookUpdateParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Events** | **[]string** |  | 
**Headers** | Pointer to **[]int32** |  | [optional] 
**Url** | **string** |  | 

## Methods

### NewV0AppWebhookUpdateParams

`func NewV0AppWebhookUpdateParams(events []string, url string, ) *V0AppWebhookUpdateParams`

NewV0AppWebhookUpdateParams instantiates a new V0AppWebhookUpdateParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0AppWebhookUpdateParamsWithDefaults

`func NewV0AppWebhookUpdateParamsWithDefaults() *V0AppWebhookUpdateParams`

NewV0AppWebhookUpdateParamsWithDefaults instantiates a new V0AppWebhookUpdateParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvents

`func (o *V0AppWebhookUpdateParams) GetEvents() []string`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *V0AppWebhookUpdateParams) GetEventsOk() (*[]string, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *V0AppWebhookUpdateParams) SetEvents(v []string)`

SetEvents sets Events field to given value.


### GetHeaders

`func (o *V0AppWebhookUpdateParams) GetHeaders() []int32`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *V0AppWebhookUpdateParams) GetHeadersOk() (*[]int32, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *V0AppWebhookUpdateParams) SetHeaders(v []int32)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *V0AppWebhookUpdateParams) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetUrl

`func (o *V0AppWebhookUpdateParams) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *V0AppWebhookUpdateParams) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *V0AppWebhookUpdateParams) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


