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

// checks if the V0ReleaseCreateAppStoreRespModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0ReleaseCreateAppStoreRespModel{}

// V0ReleaseCreateAppStoreRespModel struct for V0ReleaseCreateAppStoreRespModel
type V0ReleaseCreateAppStoreRespModel struct {
	// The bundle ID of the app
	BundleId *string `json:"bundle_id,omitempty"`
	// The ID of the newly created release
	Id *string `json:"id,omitempty"`
	// The name/version of the release (e.g. `1.2`)
	Name *string `json:"name,omitempty"`
	// The platform of the release (`ios`)
	Platform *string `json:"platform,omitempty"`
	// The status of the newly created release (`scheduled` or `in-progress`, depending on whether or not the release candidate settings were specified)
	Status *string `json:"status,omitempty"`
}

// NewV0ReleaseCreateAppStoreRespModel instantiates a new V0ReleaseCreateAppStoreRespModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0ReleaseCreateAppStoreRespModel() *V0ReleaseCreateAppStoreRespModel {
	this := V0ReleaseCreateAppStoreRespModel{}
	return &this
}

// NewV0ReleaseCreateAppStoreRespModelWithDefaults instantiates a new V0ReleaseCreateAppStoreRespModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0ReleaseCreateAppStoreRespModelWithDefaults() *V0ReleaseCreateAppStoreRespModel {
	this := V0ReleaseCreateAppStoreRespModel{}
	return &this
}

// GetBundleId returns the BundleId field value if set, zero value otherwise.
func (o *V0ReleaseCreateAppStoreRespModel) GetBundleId() string {
	if o == nil || IsNil(o.BundleId) {
		var ret string
		return ret
	}
	return *o.BundleId
}

// GetBundleIdOk returns a tuple with the BundleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0ReleaseCreateAppStoreRespModel) GetBundleIdOk() (*string, bool) {
	if o == nil || IsNil(o.BundleId) {
		return nil, false
	}
	return o.BundleId, true
}

// HasBundleId returns a boolean if a field has been set.
func (o *V0ReleaseCreateAppStoreRespModel) HasBundleId() bool {
	if o != nil && !IsNil(o.BundleId) {
		return true
	}

	return false
}

// SetBundleId gets a reference to the given string and assigns it to the BundleId field.
func (o *V0ReleaseCreateAppStoreRespModel) SetBundleId(v string) {
	o.BundleId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *V0ReleaseCreateAppStoreRespModel) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0ReleaseCreateAppStoreRespModel) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *V0ReleaseCreateAppStoreRespModel) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *V0ReleaseCreateAppStoreRespModel) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *V0ReleaseCreateAppStoreRespModel) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0ReleaseCreateAppStoreRespModel) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *V0ReleaseCreateAppStoreRespModel) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *V0ReleaseCreateAppStoreRespModel) SetName(v string) {
	o.Name = &v
}

// GetPlatform returns the Platform field value if set, zero value otherwise.
func (o *V0ReleaseCreateAppStoreRespModel) GetPlatform() string {
	if o == nil || IsNil(o.Platform) {
		var ret string
		return ret
	}
	return *o.Platform
}

// GetPlatformOk returns a tuple with the Platform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0ReleaseCreateAppStoreRespModel) GetPlatformOk() (*string, bool) {
	if o == nil || IsNil(o.Platform) {
		return nil, false
	}
	return o.Platform, true
}

// HasPlatform returns a boolean if a field has been set.
func (o *V0ReleaseCreateAppStoreRespModel) HasPlatform() bool {
	if o != nil && !IsNil(o.Platform) {
		return true
	}

	return false
}

// SetPlatform gets a reference to the given string and assigns it to the Platform field.
func (o *V0ReleaseCreateAppStoreRespModel) SetPlatform(v string) {
	o.Platform = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *V0ReleaseCreateAppStoreRespModel) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0ReleaseCreateAppStoreRespModel) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *V0ReleaseCreateAppStoreRespModel) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *V0ReleaseCreateAppStoreRespModel) SetStatus(v string) {
	o.Status = &v
}

func (o V0ReleaseCreateAppStoreRespModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0ReleaseCreateAppStoreRespModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BundleId) {
		toSerialize["bundle_id"] = o.BundleId
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Platform) {
		toSerialize["platform"] = o.Platform
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableV0ReleaseCreateAppStoreRespModel struct {
	value *V0ReleaseCreateAppStoreRespModel
	isSet bool
}

func (v NullableV0ReleaseCreateAppStoreRespModel) Get() *V0ReleaseCreateAppStoreRespModel {
	return v.value
}

func (v *NullableV0ReleaseCreateAppStoreRespModel) Set(val *V0ReleaseCreateAppStoreRespModel) {
	v.value = val
	v.isSet = true
}

func (v NullableV0ReleaseCreateAppStoreRespModel) IsSet() bool {
	return v.isSet
}

func (v *NullableV0ReleaseCreateAppStoreRespModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0ReleaseCreateAppStoreRespModel(val *V0ReleaseCreateAppStoreRespModel) *NullableV0ReleaseCreateAppStoreRespModel {
	return &NullableV0ReleaseCreateAppStoreRespModel{value: val, isSet: true}
}

func (v NullableV0ReleaseCreateAppStoreRespModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0ReleaseCreateAppStoreRespModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

