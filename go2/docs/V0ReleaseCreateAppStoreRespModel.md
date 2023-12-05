# V0ReleaseCreateAppStoreRespModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BundleId** | Pointer to **string** | The bundle ID of the app | [optional] 
**Id** | Pointer to **string** | The ID of the newly created release | [optional] 
**Name** | Pointer to **string** | The name/version of the release (e.g. &#x60;1.2&#x60;) | [optional] 
**Platform** | Pointer to **string** | The platform of the release (&#x60;ios&#x60;) | [optional] 
**Status** | Pointer to **string** | The status of the newly created release (&#x60;scheduled&#x60; or &#x60;in-progress&#x60;, depending on whether or not the release candidate settings were specified) | [optional] 

## Methods

### NewV0ReleaseCreateAppStoreRespModel

`func NewV0ReleaseCreateAppStoreRespModel() *V0ReleaseCreateAppStoreRespModel`

NewV0ReleaseCreateAppStoreRespModel instantiates a new V0ReleaseCreateAppStoreRespModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0ReleaseCreateAppStoreRespModelWithDefaults

`func NewV0ReleaseCreateAppStoreRespModelWithDefaults() *V0ReleaseCreateAppStoreRespModel`

NewV0ReleaseCreateAppStoreRespModelWithDefaults instantiates a new V0ReleaseCreateAppStoreRespModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBundleId

`func (o *V0ReleaseCreateAppStoreRespModel) GetBundleId() string`

GetBundleId returns the BundleId field if non-nil, zero value otherwise.

### GetBundleIdOk

`func (o *V0ReleaseCreateAppStoreRespModel) GetBundleIdOk() (*string, bool)`

GetBundleIdOk returns a tuple with the BundleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleId

`func (o *V0ReleaseCreateAppStoreRespModel) SetBundleId(v string)`

SetBundleId sets BundleId field to given value.

### HasBundleId

`func (o *V0ReleaseCreateAppStoreRespModel) HasBundleId() bool`

HasBundleId returns a boolean if a field has been set.

### GetId

`func (o *V0ReleaseCreateAppStoreRespModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V0ReleaseCreateAppStoreRespModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V0ReleaseCreateAppStoreRespModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *V0ReleaseCreateAppStoreRespModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *V0ReleaseCreateAppStoreRespModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V0ReleaseCreateAppStoreRespModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V0ReleaseCreateAppStoreRespModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V0ReleaseCreateAppStoreRespModel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPlatform

`func (o *V0ReleaseCreateAppStoreRespModel) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *V0ReleaseCreateAppStoreRespModel) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *V0ReleaseCreateAppStoreRespModel) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *V0ReleaseCreateAppStoreRespModel) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetStatus

`func (o *V0ReleaseCreateAppStoreRespModel) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *V0ReleaseCreateAppStoreRespModel) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *V0ReleaseCreateAppStoreRespModel) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *V0ReleaseCreateAppStoreRespModel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


