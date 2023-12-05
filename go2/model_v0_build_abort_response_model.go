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

// checks if the V0BuildAbortResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0BuildAbortResponseModel{}

// V0BuildAbortResponseModel struct for V0BuildAbortResponseModel
type V0BuildAbortResponseModel struct {
	Status *string `json:"status,omitempty"`
}

// NewV0BuildAbortResponseModel instantiates a new V0BuildAbortResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0BuildAbortResponseModel() *V0BuildAbortResponseModel {
	this := V0BuildAbortResponseModel{}
	return &this
}

// NewV0BuildAbortResponseModelWithDefaults instantiates a new V0BuildAbortResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0BuildAbortResponseModelWithDefaults() *V0BuildAbortResponseModel {
	this := V0BuildAbortResponseModel{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *V0BuildAbortResponseModel) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0BuildAbortResponseModel) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *V0BuildAbortResponseModel) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *V0BuildAbortResponseModel) SetStatus(v string) {
	o.Status = &v
}

func (o V0BuildAbortResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0BuildAbortResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableV0BuildAbortResponseModel struct {
	value *V0BuildAbortResponseModel
	isSet bool
}

func (v NullableV0BuildAbortResponseModel) Get() *V0BuildAbortResponseModel {
	return v.value
}

func (v *NullableV0BuildAbortResponseModel) Set(val *V0BuildAbortResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableV0BuildAbortResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableV0BuildAbortResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0BuildAbortResponseModel(val *V0BuildAbortResponseModel) *NullableV0BuildAbortResponseModel {
	return &NullableV0BuildAbortResponseModel{value: val, isSet: true}
}

func (v NullableV0BuildAbortResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0BuildAbortResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


