# V0AppSecretUpsertParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpandInStepInputs** | Pointer to **bool** | Replace variable in inputs? https://devcenter.bitrise.io/en/references/steps-reference/step-inputs-reference.html#step-input-properties | [optional] 
**IsExposedForPullRequests** | Pointer to **bool** | Expose for Pull Requests? | [optional] 
**IsProtected** | Pointer to **bool** | Secret marked as protected? | [optional] 
**Value** | Pointer to **string** | Value of the secret | [optional] 

## Methods

### NewV0AppSecretUpsertParams

`func NewV0AppSecretUpsertParams() *V0AppSecretUpsertParams`

NewV0AppSecretUpsertParams instantiates a new V0AppSecretUpsertParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0AppSecretUpsertParamsWithDefaults

`func NewV0AppSecretUpsertParamsWithDefaults() *V0AppSecretUpsertParams`

NewV0AppSecretUpsertParamsWithDefaults instantiates a new V0AppSecretUpsertParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpandInStepInputs

`func (o *V0AppSecretUpsertParams) GetExpandInStepInputs() bool`

GetExpandInStepInputs returns the ExpandInStepInputs field if non-nil, zero value otherwise.

### GetExpandInStepInputsOk

`func (o *V0AppSecretUpsertParams) GetExpandInStepInputsOk() (*bool, bool)`

GetExpandInStepInputsOk returns a tuple with the ExpandInStepInputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpandInStepInputs

`func (o *V0AppSecretUpsertParams) SetExpandInStepInputs(v bool)`

SetExpandInStepInputs sets ExpandInStepInputs field to given value.

### HasExpandInStepInputs

`func (o *V0AppSecretUpsertParams) HasExpandInStepInputs() bool`

HasExpandInStepInputs returns a boolean if a field has been set.

### GetIsExposedForPullRequests

`func (o *V0AppSecretUpsertParams) GetIsExposedForPullRequests() bool`

GetIsExposedForPullRequests returns the IsExposedForPullRequests field if non-nil, zero value otherwise.

### GetIsExposedForPullRequestsOk

`func (o *V0AppSecretUpsertParams) GetIsExposedForPullRequestsOk() (*bool, bool)`

GetIsExposedForPullRequestsOk returns a tuple with the IsExposedForPullRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExposedForPullRequests

`func (o *V0AppSecretUpsertParams) SetIsExposedForPullRequests(v bool)`

SetIsExposedForPullRequests sets IsExposedForPullRequests field to given value.

### HasIsExposedForPullRequests

`func (o *V0AppSecretUpsertParams) HasIsExposedForPullRequests() bool`

HasIsExposedForPullRequests returns a boolean if a field has been set.

### GetIsProtected

`func (o *V0AppSecretUpsertParams) GetIsProtected() bool`

GetIsProtected returns the IsProtected field if non-nil, zero value otherwise.

### GetIsProtectedOk

`func (o *V0AppSecretUpsertParams) GetIsProtectedOk() (*bool, bool)`

GetIsProtectedOk returns a tuple with the IsProtected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsProtected

`func (o *V0AppSecretUpsertParams) SetIsProtected(v bool)`

SetIsProtected sets IsProtected field to given value.

### HasIsProtected

`func (o *V0AppSecretUpsertParams) HasIsProtected() bool`

HasIsProtected returns a boolean if a field has been set.

### GetValue

`func (o *V0AppSecretUpsertParams) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *V0AppSecretUpsertParams) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *V0AppSecretUpsertParams) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *V0AppSecretUpsertParams) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


