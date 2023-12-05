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

// checks if the V0PlanDataModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0PlanDataModel{}

// V0PlanDataModel struct for V0PlanDataModel
type V0PlanDataModel struct {
	ContainerCount *int32 `json:"container_count,omitempty"`
	ExpiresAt *string `json:"expires_at,omitempty"`
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Price *NullsInt64 `json:"price,omitempty"`
}

// NewV0PlanDataModel instantiates a new V0PlanDataModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0PlanDataModel() *V0PlanDataModel {
	this := V0PlanDataModel{}
	return &this
}

// NewV0PlanDataModelWithDefaults instantiates a new V0PlanDataModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0PlanDataModelWithDefaults() *V0PlanDataModel {
	this := V0PlanDataModel{}
	return &this
}

// GetContainerCount returns the ContainerCount field value if set, zero value otherwise.
func (o *V0PlanDataModel) GetContainerCount() int32 {
	if o == nil || IsNil(o.ContainerCount) {
		var ret int32
		return ret
	}
	return *o.ContainerCount
}

// GetContainerCountOk returns a tuple with the ContainerCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0PlanDataModel) GetContainerCountOk() (*int32, bool) {
	if o == nil || IsNil(o.ContainerCount) {
		return nil, false
	}
	return o.ContainerCount, true
}

// HasContainerCount returns a boolean if a field has been set.
func (o *V0PlanDataModel) HasContainerCount() bool {
	if o != nil && !IsNil(o.ContainerCount) {
		return true
	}

	return false
}

// SetContainerCount gets a reference to the given int32 and assigns it to the ContainerCount field.
func (o *V0PlanDataModel) SetContainerCount(v int32) {
	o.ContainerCount = &v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *V0PlanDataModel) GetExpiresAt() string {
	if o == nil || IsNil(o.ExpiresAt) {
		var ret string
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0PlanDataModel) GetExpiresAtOk() (*string, bool) {
	if o == nil || IsNil(o.ExpiresAt) {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *V0PlanDataModel) HasExpiresAt() bool {
	if o != nil && !IsNil(o.ExpiresAt) {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given string and assigns it to the ExpiresAt field.
func (o *V0PlanDataModel) SetExpiresAt(v string) {
	o.ExpiresAt = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *V0PlanDataModel) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0PlanDataModel) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *V0PlanDataModel) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *V0PlanDataModel) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *V0PlanDataModel) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0PlanDataModel) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *V0PlanDataModel) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *V0PlanDataModel) SetName(v string) {
	o.Name = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *V0PlanDataModel) GetPrice() NullsInt64 {
	if o == nil || IsNil(o.Price) {
		var ret NullsInt64
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0PlanDataModel) GetPriceOk() (*NullsInt64, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *V0PlanDataModel) HasPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given NullsInt64 and assigns it to the Price field.
func (o *V0PlanDataModel) SetPrice(v NullsInt64) {
	o.Price = &v
}

func (o V0PlanDataModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0PlanDataModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ContainerCount) {
		toSerialize["container_count"] = o.ContainerCount
	}
	if !IsNil(o.ExpiresAt) {
		toSerialize["expires_at"] = o.ExpiresAt
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	return toSerialize, nil
}

type NullableV0PlanDataModel struct {
	value *V0PlanDataModel
	isSet bool
}

func (v NullableV0PlanDataModel) Get() *V0PlanDataModel {
	return v.value
}

func (v *NullableV0PlanDataModel) Set(val *V0PlanDataModel) {
	v.value = val
	v.isSet = true
}

func (v NullableV0PlanDataModel) IsSet() bool {
	return v.isSet
}

func (v *NullableV0PlanDataModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0PlanDataModel(val *V0PlanDataModel) *NullableV0PlanDataModel {
	return &NullableV0PlanDataModel{value: val, isSet: true}
}

func (v NullableV0PlanDataModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0PlanDataModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


