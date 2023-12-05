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

// checks if the V0TestDeviceResponseItemModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0TestDeviceResponseItemModel{}

// V0TestDeviceResponseItemModel struct for V0TestDeviceResponseItemModel
type V0TestDeviceResponseItemModel struct {
	DeviceId *string `json:"device_id,omitempty"`
	DeviceType *string `json:"device_type,omitempty"`
	Owner *string `json:"owner,omitempty"`
}

// NewV0TestDeviceResponseItemModel instantiates a new V0TestDeviceResponseItemModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0TestDeviceResponseItemModel() *V0TestDeviceResponseItemModel {
	this := V0TestDeviceResponseItemModel{}
	return &this
}

// NewV0TestDeviceResponseItemModelWithDefaults instantiates a new V0TestDeviceResponseItemModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0TestDeviceResponseItemModelWithDefaults() *V0TestDeviceResponseItemModel {
	this := V0TestDeviceResponseItemModel{}
	return &this
}

// GetDeviceId returns the DeviceId field value if set, zero value otherwise.
func (o *V0TestDeviceResponseItemModel) GetDeviceId() string {
	if o == nil || IsNil(o.DeviceId) {
		var ret string
		return ret
	}
	return *o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0TestDeviceResponseItemModel) GetDeviceIdOk() (*string, bool) {
	if o == nil || IsNil(o.DeviceId) {
		return nil, false
	}
	return o.DeviceId, true
}

// HasDeviceId returns a boolean if a field has been set.
func (o *V0TestDeviceResponseItemModel) HasDeviceId() bool {
	if o != nil && !IsNil(o.DeviceId) {
		return true
	}

	return false
}

// SetDeviceId gets a reference to the given string and assigns it to the DeviceId field.
func (o *V0TestDeviceResponseItemModel) SetDeviceId(v string) {
	o.DeviceId = &v
}

// GetDeviceType returns the DeviceType field value if set, zero value otherwise.
func (o *V0TestDeviceResponseItemModel) GetDeviceType() string {
	if o == nil || IsNil(o.DeviceType) {
		var ret string
		return ret
	}
	return *o.DeviceType
}

// GetDeviceTypeOk returns a tuple with the DeviceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0TestDeviceResponseItemModel) GetDeviceTypeOk() (*string, bool) {
	if o == nil || IsNil(o.DeviceType) {
		return nil, false
	}
	return o.DeviceType, true
}

// HasDeviceType returns a boolean if a field has been set.
func (o *V0TestDeviceResponseItemModel) HasDeviceType() bool {
	if o != nil && !IsNil(o.DeviceType) {
		return true
	}

	return false
}

// SetDeviceType gets a reference to the given string and assigns it to the DeviceType field.
func (o *V0TestDeviceResponseItemModel) SetDeviceType(v string) {
	o.DeviceType = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *V0TestDeviceResponseItemModel) GetOwner() string {
	if o == nil || IsNil(o.Owner) {
		var ret string
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0TestDeviceResponseItemModel) GetOwnerOk() (*string, bool) {
	if o == nil || IsNil(o.Owner) {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *V0TestDeviceResponseItemModel) HasOwner() bool {
	if o != nil && !IsNil(o.Owner) {
		return true
	}

	return false
}

// SetOwner gets a reference to the given string and assigns it to the Owner field.
func (o *V0TestDeviceResponseItemModel) SetOwner(v string) {
	o.Owner = &v
}

func (o V0TestDeviceResponseItemModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0TestDeviceResponseItemModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DeviceId) {
		toSerialize["device_id"] = o.DeviceId
	}
	if !IsNil(o.DeviceType) {
		toSerialize["device_type"] = o.DeviceType
	}
	if !IsNil(o.Owner) {
		toSerialize["owner"] = o.Owner
	}
	return toSerialize, nil
}

type NullableV0TestDeviceResponseItemModel struct {
	value *V0TestDeviceResponseItemModel
	isSet bool
}

func (v NullableV0TestDeviceResponseItemModel) Get() *V0TestDeviceResponseItemModel {
	return v.value
}

func (v *NullableV0TestDeviceResponseItemModel) Set(val *V0TestDeviceResponseItemModel) {
	v.value = val
	v.isSet = true
}

func (v NullableV0TestDeviceResponseItemModel) IsSet() bool {
	return v.isSet
}

func (v *NullableV0TestDeviceResponseItemModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0TestDeviceResponseItemModel(val *V0TestDeviceResponseItemModel) *NullableV0TestDeviceResponseItemModel {
	return &NullableV0TestDeviceResponseItemModel{value: val, isSet: true}
}

func (v NullableV0TestDeviceResponseItemModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0TestDeviceResponseItemModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


