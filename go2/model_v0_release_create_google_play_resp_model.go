/*
Bitrise API

Official REST API for Bitrise.io

API version: 0.1
Contact: letsconnect@bitrise.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the V0ReleaseCreateGooglePlayRespModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0ReleaseCreateGooglePlayRespModel{}

// V0ReleaseCreateGooglePlayRespModel struct for V0ReleaseCreateGooglePlayRespModel
type V0ReleaseCreateGooglePlayRespModel struct {
	// The ID of the newly created release
	Id *string `json:"id,omitempty"`
	// The name/version of the release (e.g. `1.2`)
	Name *string `json:"name,omitempty"`
	// The package name of the app
	PackageName *string `json:"package_name,omitempty"`
	// The platform of the release (`android`)
	Platform *string `json:"platform,omitempty"`
	// The status of the newly created release (`scheduled` or `in-progress`, depending on whether or not the release candidate settings were specified)
	Status *string `json:"status,omitempty"`
}

// NewV0ReleaseCreateGooglePlayRespModel instantiates a new V0ReleaseCreateGooglePlayRespModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0ReleaseCreateGooglePlayRespModel() *V0ReleaseCreateGooglePlayRespModel {
	this := V0ReleaseCreateGooglePlayRespModel{}
	return &this
}

// NewV0ReleaseCreateGooglePlayRespModelWithDefaults instantiates a new V0ReleaseCreateGooglePlayRespModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0ReleaseCreateGooglePlayRespModelWithDefaults() *V0ReleaseCreateGooglePlayRespModel {
	this := V0ReleaseCreateGooglePlayRespModel{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *V0ReleaseCreateGooglePlayRespModel) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0ReleaseCreateGooglePlayRespModel) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *V0ReleaseCreateGooglePlayRespModel) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *V0ReleaseCreateGooglePlayRespModel) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *V0ReleaseCreateGooglePlayRespModel) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0ReleaseCreateGooglePlayRespModel) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *V0ReleaseCreateGooglePlayRespModel) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *V0ReleaseCreateGooglePlayRespModel) SetName(v string) {
	o.Name = &v
}

// GetPackageName returns the PackageName field value if set, zero value otherwise.
func (o *V0ReleaseCreateGooglePlayRespModel) GetPackageName() string {
	if o == nil || IsNil(o.PackageName) {
		var ret string
		return ret
	}
	return *o.PackageName
}

// GetPackageNameOk returns a tuple with the PackageName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0ReleaseCreateGooglePlayRespModel) GetPackageNameOk() (*string, bool) {
	if o == nil || IsNil(o.PackageName) {
		return nil, false
	}
	return o.PackageName, true
}

// HasPackageName returns a boolean if a field has been set.
func (o *V0ReleaseCreateGooglePlayRespModel) HasPackageName() bool {
	if o != nil && !IsNil(o.PackageName) {
		return true
	}

	return false
}

// SetPackageName gets a reference to the given string and assigns it to the PackageName field.
func (o *V0ReleaseCreateGooglePlayRespModel) SetPackageName(v string) {
	o.PackageName = &v
}

// GetPlatform returns the Platform field value if set, zero value otherwise.
func (o *V0ReleaseCreateGooglePlayRespModel) GetPlatform() string {
	if o == nil || IsNil(o.Platform) {
		var ret string
		return ret
	}
	return *o.Platform
}

// GetPlatformOk returns a tuple with the Platform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0ReleaseCreateGooglePlayRespModel) GetPlatformOk() (*string, bool) {
	if o == nil || IsNil(o.Platform) {
		return nil, false
	}
	return o.Platform, true
}

// HasPlatform returns a boolean if a field has been set.
func (o *V0ReleaseCreateGooglePlayRespModel) HasPlatform() bool {
	if o != nil && !IsNil(o.Platform) {
		return true
	}

	return false
}

// SetPlatform gets a reference to the given string and assigns it to the Platform field.
func (o *V0ReleaseCreateGooglePlayRespModel) SetPlatform(v string) {
	o.Platform = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *V0ReleaseCreateGooglePlayRespModel) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0ReleaseCreateGooglePlayRespModel) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *V0ReleaseCreateGooglePlayRespModel) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *V0ReleaseCreateGooglePlayRespModel) SetStatus(v string) {
	o.Status = &v
}

func (o V0ReleaseCreateGooglePlayRespModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0ReleaseCreateGooglePlayRespModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.PackageName) {
		toSerialize["package_name"] = o.PackageName
	}
	if !IsNil(o.Platform) {
		toSerialize["platform"] = o.Platform
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableV0ReleaseCreateGooglePlayRespModel struct {
	value *V0ReleaseCreateGooglePlayRespModel
	isSet bool
}

func (v NullableV0ReleaseCreateGooglePlayRespModel) Get() *V0ReleaseCreateGooglePlayRespModel {
	return v.value
}

func (v *NullableV0ReleaseCreateGooglePlayRespModel) Set(val *V0ReleaseCreateGooglePlayRespModel) {
	v.value = val
	v.isSet = true
}

func (v NullableV0ReleaseCreateGooglePlayRespModel) IsSet() bool {
	return v.isSet
}

func (v *NullableV0ReleaseCreateGooglePlayRespModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0ReleaseCreateGooglePlayRespModel(val *V0ReleaseCreateGooglePlayRespModel) *NullableV0ReleaseCreateGooglePlayRespModel {
	return &NullableV0ReleaseCreateGooglePlayRespModel{value: val, isSet: true}
}

func (v NullableV0ReleaseCreateGooglePlayRespModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0ReleaseCreateGooglePlayRespModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


