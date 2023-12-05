# AddonsSetupGuide

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Instructions** | Pointer to [**[]AddonsSetupInstruction**](AddonsSetupInstruction.md) |  | [optional] 
**Notification** | Pointer to **string** |  | [optional] 

## Methods

### NewAddonsSetupGuide

`func NewAddonsSetupGuide() *AddonsSetupGuide`

NewAddonsSetupGuide instantiates a new AddonsSetupGuide object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddonsSetupGuideWithDefaults

`func NewAddonsSetupGuideWithDefaults() *AddonsSetupGuide`

NewAddonsSetupGuideWithDefaults instantiates a new AddonsSetupGuide object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstructions

`func (o *AddonsSetupGuide) GetInstructions() []AddonsSetupInstruction`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *AddonsSetupGuide) GetInstructionsOk() (*[]AddonsSetupInstruction, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *AddonsSetupGuide) SetInstructions(v []AddonsSetupInstruction)`

SetInstructions sets Instructions field to given value.

### HasInstructions

`func (o *AddonsSetupGuide) HasInstructions() bool`

HasInstructions returns a boolean if a field has been set.

### GetNotification

`func (o *AddonsSetupGuide) GetNotification() string`

GetNotification returns the Notification field if non-nil, zero value otherwise.

### GetNotificationOk

`func (o *AddonsSetupGuide) GetNotificationOk() (*string, bool)`

GetNotificationOk returns a tuple with the Notification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotification

`func (o *AddonsSetupGuide) SetNotification(v string)`

SetNotification sets Notification field to given value.

### HasNotification

`func (o *AddonsSetupGuide) HasNotification() bool`

HasNotification returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


