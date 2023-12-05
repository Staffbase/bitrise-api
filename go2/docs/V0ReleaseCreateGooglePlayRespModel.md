# V0ReleaseCreateGooglePlayRespModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The ID of the newly created release | [optional] 
**Name** | Pointer to **string** | The name/version of the release (e.g. &#x60;1.2&#x60;) | [optional] 
**PackageName** | Pointer to **string** | The package name of the app | [optional] 
**Platform** | Pointer to **string** | The platform of the release (&#x60;android&#x60;) | [optional] 
**Status** | Pointer to **string** | The status of the newly created release (&#x60;scheduled&#x60; or &#x60;in-progress&#x60;, depending on whether or not the release candidate settings were specified) | [optional] 

## Methods

### NewV0ReleaseCreateGooglePlayRespModel

`func NewV0ReleaseCreateGooglePlayRespModel() *V0ReleaseCreateGooglePlayRespModel`

NewV0ReleaseCreateGooglePlayRespModel instantiates a new V0ReleaseCreateGooglePlayRespModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0ReleaseCreateGooglePlayRespModelWithDefaults

`func NewV0ReleaseCreateGooglePlayRespModelWithDefaults() *V0ReleaseCreateGooglePlayRespModel`

NewV0ReleaseCreateGooglePlayRespModelWithDefaults instantiates a new V0ReleaseCreateGooglePlayRespModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *V0ReleaseCreateGooglePlayRespModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V0ReleaseCreateGooglePlayRespModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V0ReleaseCreateGooglePlayRespModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *V0ReleaseCreateGooglePlayRespModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *V0ReleaseCreateGooglePlayRespModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V0ReleaseCreateGooglePlayRespModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V0ReleaseCreateGooglePlayRespModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V0ReleaseCreateGooglePlayRespModel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPackageName

`func (o *V0ReleaseCreateGooglePlayRespModel) GetPackageName() string`

GetPackageName returns the PackageName field if non-nil, zero value otherwise.

### GetPackageNameOk

`func (o *V0ReleaseCreateGooglePlayRespModel) GetPackageNameOk() (*string, bool)`

GetPackageNameOk returns a tuple with the PackageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageName

`func (o *V0ReleaseCreateGooglePlayRespModel) SetPackageName(v string)`

SetPackageName sets PackageName field to given value.

### HasPackageName

`func (o *V0ReleaseCreateGooglePlayRespModel) HasPackageName() bool`

HasPackageName returns a boolean if a field has been set.

### GetPlatform

`func (o *V0ReleaseCreateGooglePlayRespModel) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *V0ReleaseCreateGooglePlayRespModel) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *V0ReleaseCreateGooglePlayRespModel) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *V0ReleaseCreateGooglePlayRespModel) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetStatus

`func (o *V0ReleaseCreateGooglePlayRespModel) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *V0ReleaseCreateGooglePlayRespModel) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *V0ReleaseCreateGooglePlayRespModel) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *V0ReleaseCreateGooglePlayRespModel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


