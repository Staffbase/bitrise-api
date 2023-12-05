# V0WebhookDeliveryItemResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **string** |  | [optional] 
**RequestHeaders** | Pointer to **string** |  | [optional] 
**RequestPayload** | Pointer to **string** |  | [optional] 
**RequestUrl** | Pointer to **string** |  | [optional] 
**ResponseBody** | Pointer to [**NullsString**](NullsString.md) |  | [optional] 
**ResponseHeaders** | Pointer to [**NullsString**](NullsString.md) |  | [optional] 
**ResponseHttpStatus** | Pointer to [**NullsInt64**](NullsInt64.md) |  | [optional] 
**ResponseSeconds** | Pointer to [**NullsInt64**](NullsInt64.md) |  | [optional] 
**Slug** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 

## Methods

### NewV0WebhookDeliveryItemResponseModel

`func NewV0WebhookDeliveryItemResponseModel() *V0WebhookDeliveryItemResponseModel`

NewV0WebhookDeliveryItemResponseModel instantiates a new V0WebhookDeliveryItemResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0WebhookDeliveryItemResponseModelWithDefaults

`func NewV0WebhookDeliveryItemResponseModelWithDefaults() *V0WebhookDeliveryItemResponseModel`

NewV0WebhookDeliveryItemResponseModelWithDefaults instantiates a new V0WebhookDeliveryItemResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *V0WebhookDeliveryItemResponseModel) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *V0WebhookDeliveryItemResponseModel) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *V0WebhookDeliveryItemResponseModel) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *V0WebhookDeliveryItemResponseModel) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetRequestHeaders

`func (o *V0WebhookDeliveryItemResponseModel) GetRequestHeaders() string`

GetRequestHeaders returns the RequestHeaders field if non-nil, zero value otherwise.

### GetRequestHeadersOk

`func (o *V0WebhookDeliveryItemResponseModel) GetRequestHeadersOk() (*string, bool)`

GetRequestHeadersOk returns a tuple with the RequestHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestHeaders

`func (o *V0WebhookDeliveryItemResponseModel) SetRequestHeaders(v string)`

SetRequestHeaders sets RequestHeaders field to given value.

### HasRequestHeaders

`func (o *V0WebhookDeliveryItemResponseModel) HasRequestHeaders() bool`

HasRequestHeaders returns a boolean if a field has been set.

### GetRequestPayload

`func (o *V0WebhookDeliveryItemResponseModel) GetRequestPayload() string`

GetRequestPayload returns the RequestPayload field if non-nil, zero value otherwise.

### GetRequestPayloadOk

`func (o *V0WebhookDeliveryItemResponseModel) GetRequestPayloadOk() (*string, bool)`

GetRequestPayloadOk returns a tuple with the RequestPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestPayload

`func (o *V0WebhookDeliveryItemResponseModel) SetRequestPayload(v string)`

SetRequestPayload sets RequestPayload field to given value.

### HasRequestPayload

`func (o *V0WebhookDeliveryItemResponseModel) HasRequestPayload() bool`

HasRequestPayload returns a boolean if a field has been set.

### GetRequestUrl

`func (o *V0WebhookDeliveryItemResponseModel) GetRequestUrl() string`

GetRequestUrl returns the RequestUrl field if non-nil, zero value otherwise.

### GetRequestUrlOk

`func (o *V0WebhookDeliveryItemResponseModel) GetRequestUrlOk() (*string, bool)`

GetRequestUrlOk returns a tuple with the RequestUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestUrl

`func (o *V0WebhookDeliveryItemResponseModel) SetRequestUrl(v string)`

SetRequestUrl sets RequestUrl field to given value.

### HasRequestUrl

`func (o *V0WebhookDeliveryItemResponseModel) HasRequestUrl() bool`

HasRequestUrl returns a boolean if a field has been set.

### GetResponseBody

`func (o *V0WebhookDeliveryItemResponseModel) GetResponseBody() NullsString`

GetResponseBody returns the ResponseBody field if non-nil, zero value otherwise.

### GetResponseBodyOk

`func (o *V0WebhookDeliveryItemResponseModel) GetResponseBodyOk() (*NullsString, bool)`

GetResponseBodyOk returns a tuple with the ResponseBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseBody

`func (o *V0WebhookDeliveryItemResponseModel) SetResponseBody(v NullsString)`

SetResponseBody sets ResponseBody field to given value.

### HasResponseBody

`func (o *V0WebhookDeliveryItemResponseModel) HasResponseBody() bool`

HasResponseBody returns a boolean if a field has been set.

### GetResponseHeaders

`func (o *V0WebhookDeliveryItemResponseModel) GetResponseHeaders() NullsString`

GetResponseHeaders returns the ResponseHeaders field if non-nil, zero value otherwise.

### GetResponseHeadersOk

`func (o *V0WebhookDeliveryItemResponseModel) GetResponseHeadersOk() (*NullsString, bool)`

GetResponseHeadersOk returns a tuple with the ResponseHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseHeaders

`func (o *V0WebhookDeliveryItemResponseModel) SetResponseHeaders(v NullsString)`

SetResponseHeaders sets ResponseHeaders field to given value.

### HasResponseHeaders

`func (o *V0WebhookDeliveryItemResponseModel) HasResponseHeaders() bool`

HasResponseHeaders returns a boolean if a field has been set.

### GetResponseHttpStatus

`func (o *V0WebhookDeliveryItemResponseModel) GetResponseHttpStatus() NullsInt64`

GetResponseHttpStatus returns the ResponseHttpStatus field if non-nil, zero value otherwise.

### GetResponseHttpStatusOk

`func (o *V0WebhookDeliveryItemResponseModel) GetResponseHttpStatusOk() (*NullsInt64, bool)`

GetResponseHttpStatusOk returns a tuple with the ResponseHttpStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseHttpStatus

`func (o *V0WebhookDeliveryItemResponseModel) SetResponseHttpStatus(v NullsInt64)`

SetResponseHttpStatus sets ResponseHttpStatus field to given value.

### HasResponseHttpStatus

`func (o *V0WebhookDeliveryItemResponseModel) HasResponseHttpStatus() bool`

HasResponseHttpStatus returns a boolean if a field has been set.

### GetResponseSeconds

`func (o *V0WebhookDeliveryItemResponseModel) GetResponseSeconds() NullsInt64`

GetResponseSeconds returns the ResponseSeconds field if non-nil, zero value otherwise.

### GetResponseSecondsOk

`func (o *V0WebhookDeliveryItemResponseModel) GetResponseSecondsOk() (*NullsInt64, bool)`

GetResponseSecondsOk returns a tuple with the ResponseSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseSeconds

`func (o *V0WebhookDeliveryItemResponseModel) SetResponseSeconds(v NullsInt64)`

SetResponseSeconds sets ResponseSeconds field to given value.

### HasResponseSeconds

`func (o *V0WebhookDeliveryItemResponseModel) HasResponseSeconds() bool`

HasResponseSeconds returns a boolean if a field has been set.

### GetSlug

`func (o *V0WebhookDeliveryItemResponseModel) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *V0WebhookDeliveryItemResponseModel) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *V0WebhookDeliveryItemResponseModel) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *V0WebhookDeliveryItemResponseModel) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *V0WebhookDeliveryItemResponseModel) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *V0WebhookDeliveryItemResponseModel) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *V0WebhookDeliveryItemResponseModel) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *V0WebhookDeliveryItemResponseModel) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


