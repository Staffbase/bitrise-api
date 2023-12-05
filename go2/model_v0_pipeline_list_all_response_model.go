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

// checks if the V0PipelineListAllResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0PipelineListAllResponseModel{}

// V0PipelineListAllResponseModel struct for V0PipelineListAllResponseModel
type V0PipelineListAllResponseModel struct {
	Data []V0PipelineListAllResponseItemModel `json:"data,omitempty"`
	Paging *V0PagingResponseModel `json:"paging,omitempty"`
}

// NewV0PipelineListAllResponseModel instantiates a new V0PipelineListAllResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0PipelineListAllResponseModel() *V0PipelineListAllResponseModel {
	this := V0PipelineListAllResponseModel{}
	return &this
}

// NewV0PipelineListAllResponseModelWithDefaults instantiates a new V0PipelineListAllResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0PipelineListAllResponseModelWithDefaults() *V0PipelineListAllResponseModel {
	this := V0PipelineListAllResponseModel{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *V0PipelineListAllResponseModel) GetData() []V0PipelineListAllResponseItemModel {
	if o == nil || IsNil(o.Data) {
		var ret []V0PipelineListAllResponseItemModel
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0PipelineListAllResponseModel) GetDataOk() ([]V0PipelineListAllResponseItemModel, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *V0PipelineListAllResponseModel) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []V0PipelineListAllResponseItemModel and assigns it to the Data field.
func (o *V0PipelineListAllResponseModel) SetData(v []V0PipelineListAllResponseItemModel) {
	o.Data = v
}

// GetPaging returns the Paging field value if set, zero value otherwise.
func (o *V0PipelineListAllResponseModel) GetPaging() V0PagingResponseModel {
	if o == nil || IsNil(o.Paging) {
		var ret V0PagingResponseModel
		return ret
	}
	return *o.Paging
}

// GetPagingOk returns a tuple with the Paging field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0PipelineListAllResponseModel) GetPagingOk() (*V0PagingResponseModel, bool) {
	if o == nil || IsNil(o.Paging) {
		return nil, false
	}
	return o.Paging, true
}

// HasPaging returns a boolean if a field has been set.
func (o *V0PipelineListAllResponseModel) HasPaging() bool {
	if o != nil && !IsNil(o.Paging) {
		return true
	}

	return false
}

// SetPaging gets a reference to the given V0PagingResponseModel and assigns it to the Paging field.
func (o *V0PipelineListAllResponseModel) SetPaging(v V0PagingResponseModel) {
	o.Paging = &v
}

func (o V0PipelineListAllResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0PipelineListAllResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Paging) {
		toSerialize["paging"] = o.Paging
	}
	return toSerialize, nil
}

type NullableV0PipelineListAllResponseModel struct {
	value *V0PipelineListAllResponseModel
	isSet bool
}

func (v NullableV0PipelineListAllResponseModel) Get() *V0PipelineListAllResponseModel {
	return v.value
}

func (v *NullableV0PipelineListAllResponseModel) Set(val *V0PipelineListAllResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableV0PipelineListAllResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableV0PipelineListAllResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0PipelineListAllResponseModel(val *V0PipelineListAllResponseModel) *NullableV0PipelineListAllResponseModel {
	return &NullableV0PipelineListAllResponseModel{value: val, isSet: true}
}

func (v NullableV0PipelineListAllResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0PipelineListAllResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

