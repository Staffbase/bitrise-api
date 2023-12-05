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

// checks if the V0ProvisionProfileResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0ProvisionProfileResponseModel{}

// V0ProvisionProfileResponseModel struct for V0ProvisionProfileResponseModel
type V0ProvisionProfileResponseModel struct {
	Data *V0ProvisionProfileResponseItemModel `json:"data,omitempty"`
}

// NewV0ProvisionProfileResponseModel instantiates a new V0ProvisionProfileResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0ProvisionProfileResponseModel() *V0ProvisionProfileResponseModel {
	this := V0ProvisionProfileResponseModel{}
	return &this
}

// NewV0ProvisionProfileResponseModelWithDefaults instantiates a new V0ProvisionProfileResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0ProvisionProfileResponseModelWithDefaults() *V0ProvisionProfileResponseModel {
	this := V0ProvisionProfileResponseModel{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *V0ProvisionProfileResponseModel) GetData() V0ProvisionProfileResponseItemModel {
	if o == nil || IsNil(o.Data) {
		var ret V0ProvisionProfileResponseItemModel
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0ProvisionProfileResponseModel) GetDataOk() (*V0ProvisionProfileResponseItemModel, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *V0ProvisionProfileResponseModel) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given V0ProvisionProfileResponseItemModel and assigns it to the Data field.
func (o *V0ProvisionProfileResponseModel) SetData(v V0ProvisionProfileResponseItemModel) {
	o.Data = &v
}

func (o V0ProvisionProfileResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0ProvisionProfileResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableV0ProvisionProfileResponseModel struct {
	value *V0ProvisionProfileResponseModel
	isSet bool
}

func (v NullableV0ProvisionProfileResponseModel) Get() *V0ProvisionProfileResponseModel {
	return v.value
}

func (v *NullableV0ProvisionProfileResponseModel) Set(val *V0ProvisionProfileResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableV0ProvisionProfileResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableV0ProvisionProfileResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0ProvisionProfileResponseModel(val *V0ProvisionProfileResponseModel) *NullableV0ProvisionProfileResponseModel {
	return &NullableV0ProvisionProfileResponseModel{value: val, isSet: true}
}

func (v NullableV0ProvisionProfileResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0ProvisionProfileResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


