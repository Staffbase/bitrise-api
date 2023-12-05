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

// checks if the V0WebhookDeliveryItemShowResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0WebhookDeliveryItemShowResponseModel{}

// V0WebhookDeliveryItemShowResponseModel struct for V0WebhookDeliveryItemShowResponseModel
type V0WebhookDeliveryItemShowResponseModel struct {
	Data *V0WebhookDeliveryItemResponseModel `json:"data,omitempty"`
}

// NewV0WebhookDeliveryItemShowResponseModel instantiates a new V0WebhookDeliveryItemShowResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0WebhookDeliveryItemShowResponseModel() *V0WebhookDeliveryItemShowResponseModel {
	this := V0WebhookDeliveryItemShowResponseModel{}
	return &this
}

// NewV0WebhookDeliveryItemShowResponseModelWithDefaults instantiates a new V0WebhookDeliveryItemShowResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0WebhookDeliveryItemShowResponseModelWithDefaults() *V0WebhookDeliveryItemShowResponseModel {
	this := V0WebhookDeliveryItemShowResponseModel{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *V0WebhookDeliveryItemShowResponseModel) GetData() V0WebhookDeliveryItemResponseModel {
	if o == nil || IsNil(o.Data) {
		var ret V0WebhookDeliveryItemResponseModel
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0WebhookDeliveryItemShowResponseModel) GetDataOk() (*V0WebhookDeliveryItemResponseModel, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *V0WebhookDeliveryItemShowResponseModel) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given V0WebhookDeliveryItemResponseModel and assigns it to the Data field.
func (o *V0WebhookDeliveryItemShowResponseModel) SetData(v V0WebhookDeliveryItemResponseModel) {
	o.Data = &v
}

func (o V0WebhookDeliveryItemShowResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0WebhookDeliveryItemShowResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableV0WebhookDeliveryItemShowResponseModel struct {
	value *V0WebhookDeliveryItemShowResponseModel
	isSet bool
}

func (v NullableV0WebhookDeliveryItemShowResponseModel) Get() *V0WebhookDeliveryItemShowResponseModel {
	return v.value
}

func (v *NullableV0WebhookDeliveryItemShowResponseModel) Set(val *V0WebhookDeliveryItemShowResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableV0WebhookDeliveryItemShowResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableV0WebhookDeliveryItemShowResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0WebhookDeliveryItemShowResponseModel(val *V0WebhookDeliveryItemShowResponseModel) *NullableV0WebhookDeliveryItemShowResponseModel {
	return &NullableV0WebhookDeliveryItemShowResponseModel{value: val, isSet: true}
}

func (v NullableV0WebhookDeliveryItemShowResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0WebhookDeliveryItemShowResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

