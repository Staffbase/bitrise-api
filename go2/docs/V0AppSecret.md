# V0AppSecret

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpandInStepInputs** | Pointer to **bool** | Replace variable in inputs? https://devcenter.bitrise.io/en/references/steps-reference/step-inputs-reference.html#step-input-properties | [optional] 
**IsExposedForPullRequests** | Pointer to **bool** | Expose for Pull Requests? | [optional] 
**IsProtected** | Pointer to **bool** | Secret marked as protected? | [optional] 
**Name** | Pointer to **string** | Name of the secret | [optional] 

## Methods

### NewV0AppSecret

`func NewV0AppSecret() *V0AppSecret`

NewV0AppSecret instantiates a new V0AppSecret object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0AppSecretWithDefaults

`func NewV0AppSecretWithDefaults() *V0AppSecret`

NewV0AppSecretWithDefaults instantiates a new V0AppSecret object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpandInStepInputs

`func (o *V0AppSecret) GetExpandInStepInputs() bool`

GetExpandInStepInputs returns the ExpandInStepInputs field if non-nil, zero value otherwise.

### GetExpandInStepInputsOk

`func (o *V0AppSecret) GetExpandInStepInputsOk() (*bool, bool)`

GetExpandInStepInputsOk returns a tuple with the ExpandInStepInputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpandInStepInputs

`func (o *V0AppSecret) SetExpandInStepInputs(v bool)`

SetExpandInStepInputs sets ExpandInStepInputs field to given value.

### HasExpandInStepInputs

`func (o *V0AppSecret) HasExpandInStepInputs() bool`

HasExpandInStepInputs returns a boolean if a field has been set.

### GetIsExposedForPullRequests

`func (o *V0AppSecret) GetIsExposedForPullRequests() bool`

GetIsExposedForPullRequests returns the IsExposedForPullRequests field if non-nil, zero value otherwise.

### GetIsExposedForPullRequestsOk

`func (o *V0AppSecret) GetIsExposedForPullRequestsOk() (*bool, bool)`

GetIsExposedForPullRequestsOk returns a tuple with the IsExposedForPullRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExposedForPullRequests

`func (o *V0AppSecret) SetIsExposedForPullRequests(v bool)`

SetIsExposedForPullRequests sets IsExposedForPullRequests field to given value.

### HasIsExposedForPullRequests

`func (o *V0AppSecret) HasIsExposedForPullRequests() bool`

HasIsExposedForPullRequests returns a boolean if a field has been set.

### GetIsProtected

`func (o *V0AppSecret) GetIsProtected() bool`

GetIsProtected returns the IsProtected field if non-nil, zero value otherwise.

### GetIsProtectedOk

`func (o *V0AppSecret) GetIsProtectedOk() (*bool, bool)`

GetIsProtectedOk returns a tuple with the IsProtected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsProtected

`func (o *V0AppSecret) SetIsProtected(v bool)`

SetIsProtected sets IsProtected field to given value.

### HasIsProtected

`func (o *V0AppSecret) HasIsProtected() bool`

HasIsProtected returns a boolean if a field has been set.

### GetName

`func (o *V0AppSecret) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V0AppSecret) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V0AppSecret) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V0AppSecret) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


