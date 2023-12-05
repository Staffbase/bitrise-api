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

// checks if the V0AddonsListResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0AddonsListResponseModel{}

// V0AddonsListResponseModel struct for V0AddonsListResponseModel
type V0AddonsListResponseModel struct {
	Data []AddonsAddon `json:"data,omitempty"`
}

// NewV0AddonsListResponseModel instantiates a new V0AddonsListResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0AddonsListResponseModel() *V0AddonsListResponseModel {
	this := V0AddonsListResponseModel{}
	return &this
}

// NewV0AddonsListResponseModelWithDefaults instantiates a new V0AddonsListResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0AddonsListResponseModelWithDefaults() *V0AddonsListResponseModel {
	this := V0AddonsListResponseModel{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *V0AddonsListResponseModel) GetData() []AddonsAddon {
	if o == nil || IsNil(o.Data) {
		var ret []AddonsAddon
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0AddonsListResponseModel) GetDataOk() ([]AddonsAddon, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *V0AddonsListResponseModel) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []AddonsAddon and assigns it to the Data field.
func (o *V0AddonsListResponseModel) SetData(v []AddonsAddon) {
	o.Data = v
}

func (o V0AddonsListResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0AddonsListResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableV0AddonsListResponseModel struct {
	value *V0AddonsListResponseModel
	isSet bool
}

func (v NullableV0AddonsListResponseModel) Get() *V0AddonsListResponseModel {
	return v.value
}

func (v *NullableV0AddonsListResponseModel) Set(val *V0AddonsListResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableV0AddonsListResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableV0AddonsListResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0AddonsListResponseModel(val *V0AddonsListResponseModel) *NullableV0AddonsListResponseModel {
	return &NullableV0AddonsListResponseModel{value: val, isSet: true}
}

func (v NullableV0AddonsListResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0AddonsListResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

