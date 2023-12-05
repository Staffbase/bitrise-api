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

// checks if the NullsString type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NullsString{}

// NullsString struct for NullsString
type NullsString struct {
	String *string `json:"string,omitempty"`
	// Valid is true if String is not NULL
	Valid *bool `json:"valid,omitempty"`
}

// NewNullsString instantiates a new NullsString object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNullsString() *NullsString {
	this := NullsString{}
	return &this
}

// NewNullsStringWithDefaults instantiates a new NullsString object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNullsStringWithDefaults() *NullsString {
	this := NullsString{}
	return &this
}

// GetString returns the String field value if set, zero value otherwise.
func (o *NullsString) GetString() string {
	if o == nil || IsNil(o.String) {
		var ret string
		return ret
	}
	return *o.String
}

// GetStringOk returns a tuple with the String field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NullsString) GetStringOk() (*string, bool) {
	if o == nil || IsNil(o.String) {
		return nil, false
	}
	return o.String, true
}

// HasString returns a boolean if a field has been set.
func (o *NullsString) HasString() bool {
	if o != nil && !IsNil(o.String) {
		return true
	}

	return false
}

// SetString gets a reference to the given string and assigns it to the String field.
func (o *NullsString) SetString(v string) {
	o.String = &v
}

// GetValid returns the Valid field value if set, zero value otherwise.
func (o *NullsString) GetValid() bool {
	if o == nil || IsNil(o.Valid) {
		var ret bool
		return ret
	}
	return *o.Valid
}

// GetValidOk returns a tuple with the Valid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NullsString) GetValidOk() (*bool, bool) {
	if o == nil || IsNil(o.Valid) {
		return nil, false
	}
	return o.Valid, true
}

// HasValid returns a boolean if a field has been set.
func (o *NullsString) HasValid() bool {
	if o != nil && !IsNil(o.Valid) {
		return true
	}

	return false
}

// SetValid gets a reference to the given bool and assigns it to the Valid field.
func (o *NullsString) SetValid(v bool) {
	o.Valid = &v
}

func (o NullsString) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NullsString) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.String) {
		toSerialize["string"] = o.String
	}
	if !IsNil(o.Valid) {
		toSerialize["valid"] = o.Valid
	}
	return toSerialize, nil
}

type NullableNullsString struct {
	value *NullsString
	isSet bool
}

func (v NullableNullsString) Get() *NullsString {
	return v.value
}

func (v *NullableNullsString) Set(val *NullsString) {
	v.value = val
	v.isSet = true
}

func (v NullableNullsString) IsSet() bool {
	return v.isSet
}

func (v *NullableNullsString) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNullsString(val *NullsString) *NullableNullsString {
	return &NullableNullsString{value: val, isSet: true}
}

func (v NullableNullsString) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNullsString) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


