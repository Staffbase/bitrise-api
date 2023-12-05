# V0AppNotificationSettingsParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OnFailure** | Pointer to **string** | When should Bitrise send notifications for the users on build failure? Possible values: always, never, change | [optional] 
**OnSuccess** | Pointer to **string** | When should Bitrise send notifications for the users on build success? Possible values: always, never, change | [optional] 

## Methods

### NewV0AppNotificationSettingsParams

`func NewV0AppNotificationSettingsParams() *V0AppNotificationSettingsParams`

NewV0AppNotificationSettingsParams instantiates a new V0AppNotificationSettingsParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0AppNotificationSettingsParamsWithDefaults

`func NewV0AppNotificationSettingsParamsWithDefaults() *V0AppNotificationSettingsParams`

NewV0AppNotificationSettingsParamsWithDefaults instantiates a new V0AppNotificationSettingsParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOnFailure

`func (o *V0AppNotificationSettingsParams) GetOnFailure() string`

GetOnFailure returns the OnFailure field if non-nil, zero value otherwise.

### GetOnFailureOk

`func (o *V0AppNotificationSettingsParams) GetOnFailureOk() (*string, bool)`

GetOnFailureOk returns a tuple with the OnFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnFailure

`func (o *V0AppNotificationSettingsParams) SetOnFailure(v string)`

SetOnFailure sets OnFailure field to given value.

### HasOnFailure

`func (o *V0AppNotificationSettingsParams) HasOnFailure() bool`

HasOnFailure returns a boolean if a field has been set.

### GetOnSuccess

`func (o *V0AppNotificationSettingsParams) GetOnSuccess() string`

GetOnSuccess returns the OnSuccess field if non-nil, zero value otherwise.

### GetOnSuccessOk

`func (o *V0AppNotificationSettingsParams) GetOnSuccessOk() (*string, bool)`

GetOnSuccessOk returns a tuple with the OnSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnSuccess

`func (o *V0AppNotificationSettingsParams) SetOnSuccess(v string)`

SetOnSuccess sets OnSuccess field to given value.

### HasOnSuccess

`func (o *V0AppNotificationSettingsParams) HasOnSuccess() bool`

HasOnSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


