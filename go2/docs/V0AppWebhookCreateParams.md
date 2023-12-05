# V0AppWebhookCreateParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Events** | **[]string** |  | 
**Headers** | Pointer to **[]int32** |  | [optional] 
**Secret** | **string** |  | 
**Url** | **string** |  | 

## Methods

### NewV0AppWebhookCreateParams

`func NewV0AppWebhookCreateParams(events []string, secret string, url string, ) *V0AppWebhookCreateParams`

NewV0AppWebhookCreateParams instantiates a new V0AppWebhookCreateParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0AppWebhookCreateParamsWithDefaults

`func NewV0AppWebhookCreateParamsWithDefaults() *V0AppWebhookCreateParams`

NewV0AppWebhookCreateParamsWithDefaults instantiates a new V0AppWebhookCreateParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvents

`func (o *V0AppWebhookCreateParams) GetEvents() []string`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *V0AppWebhookCreateParams) GetEventsOk() (*[]string, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *V0AppWebhookCreateParams) SetEvents(v []string)`

SetEvents sets Events field to given value.


### GetHeaders

`func (o *V0AppWebhookCreateParams) GetHeaders() []int32`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *V0AppWebhookCreateParams) GetHeadersOk() (*[]int32, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *V0AppWebhookCreateParams) SetHeaders(v []int32)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *V0AppWebhookCreateParams) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetSecret

`func (o *V0AppWebhookCreateParams) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *V0AppWebhookCreateParams) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *V0AppWebhookCreateParams) SetSecret(v string)`

SetSecret sets Secret field to given value.


### GetUrl

`func (o *V0AppWebhookCreateParams) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *V0AppWebhookCreateParams) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *V0AppWebhookCreateParams) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


