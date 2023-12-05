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

// checks if the V0AppDeleteRespModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0AppDeleteRespModel{}

// V0AppDeleteRespModel struct for V0AppDeleteRespModel
type V0AppDeleteRespModel struct {
	Msg *string `json:"msg,omitempty"`
}

// NewV0AppDeleteRespModel instantiates a new V0AppDeleteRespModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0AppDeleteRespModel() *V0AppDeleteRespModel {
	this := V0AppDeleteRespModel{}
	return &this
}

// NewV0AppDeleteRespModelWithDefaults instantiates a new V0AppDeleteRespModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0AppDeleteRespModelWithDefaults() *V0AppDeleteRespModel {
	this := V0AppDeleteRespModel{}
	return &this
}

// GetMsg returns the Msg field value if set, zero value otherwise.
func (o *V0AppDeleteRespModel) GetMsg() string {
	if o == nil || IsNil(o.Msg) {
		var ret string
		return ret
	}
	return *o.Msg
}

// GetMsgOk returns a tuple with the Msg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0AppDeleteRespModel) GetMsgOk() (*string, bool) {
	if o == nil || IsNil(o.Msg) {
		return nil, false
	}
	return o.Msg, true
}

// HasMsg returns a boolean if a field has been set.
func (o *V0AppDeleteRespModel) HasMsg() bool {
	if o != nil && !IsNil(o.Msg) {
		return true
	}

	return false
}

// SetMsg gets a reference to the given string and assigns it to the Msg field.
func (o *V0AppDeleteRespModel) SetMsg(v string) {
	o.Msg = &v
}

func (o V0AppDeleteRespModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0AppDeleteRespModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Msg) {
		toSerialize["msg"] = o.Msg
	}
	return toSerialize, nil
}

type NullableV0AppDeleteRespModel struct {
	value *V0AppDeleteRespModel
	isSet bool
}

func (v NullableV0AppDeleteRespModel) Get() *V0AppDeleteRespModel {
	return v.value
}

func (v *NullableV0AppDeleteRespModel) Set(val *V0AppDeleteRespModel) {
	v.value = val
	v.isSet = true
}

func (v NullableV0AppDeleteRespModel) IsSet() bool {
	return v.isSet
}

func (v *NullableV0AppDeleteRespModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0AppDeleteRespModel(val *V0AppDeleteRespModel) *NullableV0AppDeleteRespModel {
	return &NullableV0AppDeleteRespModel{value: val, isSet: true}
}

func (v NullableV0AppDeleteRespModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0AppDeleteRespModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

