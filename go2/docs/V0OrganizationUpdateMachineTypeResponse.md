# V0OrganizationUpdateMachineTypeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | Pointer to [**[]V0OrganizationUpdateMachineTypeResponseErrorsInner**](V0OrganizationUpdateMachineTypeResponseErrorsInner.md) |  | [optional] 
**Message** | Pointer to **string** | The result of the migration | [optional] 
**MigratedApps** | Pointer to **[]string** | The migrated apps&#39; identifiers in the following format \&quot;#{app.title} (#{app.slug})\&quot; | [optional] 

## Methods

### NewV0OrganizationUpdateMachineTypeResponse

`func NewV0OrganizationUpdateMachineTypeResponse() *V0OrganizationUpdateMachineTypeResponse`

NewV0OrganizationUpdateMachineTypeResponse instantiates a new V0OrganizationUpdateMachineTypeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0OrganizationUpdateMachineTypeResponseWithDefaults

`func NewV0OrganizationUpdateMachineTypeResponseWithDefaults() *V0OrganizationUpdateMachineTypeResponse`

NewV0OrganizationUpdateMachineTypeResponseWithDefaults instantiates a new V0OrganizationUpdateMachineTypeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrors

`func (o *V0OrganizationUpdateMachineTypeResponse) GetErrors() []V0OrganizationUpdateMachineTypeResponseErrorsInner`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *V0OrganizationUpdateMachineTypeResponse) GetErrorsOk() (*[]V0OrganizationUpdateMachineTypeResponseErrorsInner, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *V0OrganizationUpdateMachineTypeResponse) SetErrors(v []V0OrganizationUpdateMachineTypeResponseErrorsInner)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *V0OrganizationUpdateMachineTypeResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetMessage

`func (o *V0OrganizationUpdateMachineTypeResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *V0OrganizationUpdateMachineTypeResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *V0OrganizationUpdateMachineTypeResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *V0OrganizationUpdateMachineTypeResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetMigratedApps

`func (o *V0OrganizationUpdateMachineTypeResponse) GetMigratedApps() []string`

GetMigratedApps returns the MigratedApps field if non-nil, zero value otherwise.

### GetMigratedAppsOk

`func (o *V0OrganizationUpdateMachineTypeResponse) GetMigratedAppsOk() (*[]string, bool)`

GetMigratedAppsOk returns a tuple with the MigratedApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigratedApps

`func (o *V0OrganizationUpdateMachineTypeResponse) SetMigratedApps(v []string)`

SetMigratedApps sets MigratedApps field to given value.

### HasMigratedApps

`func (o *V0OrganizationUpdateMachineTypeResponse) HasMigratedApps() bool`

HasMigratedApps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


