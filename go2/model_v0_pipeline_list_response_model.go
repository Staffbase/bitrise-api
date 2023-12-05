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

// checks if the V0PipelineListResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0PipelineListResponseModel{}

// V0PipelineListResponseModel struct for V0PipelineListResponseModel
type V0PipelineListResponseModel struct {
	Data []V0PipelineListResponseItemModel `json:"data,omitempty"`
	Paging *V0PagingResponseModel `json:"paging,omitempty"`
}

// NewV0PipelineListResponseModel instantiates a new V0PipelineListResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0PipelineListResponseModel() *V0PipelineListResponseModel {
	this := V0PipelineListResponseModel{}
	return &this
}

// NewV0PipelineListResponseModelWithDefaults instantiates a new V0PipelineListResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0PipelineListResponseModelWithDefaults() *V0PipelineListResponseModel {
	this := V0PipelineListResponseModel{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *V0PipelineListResponseModel) GetData() []V0PipelineListResponseItemModel {
	if o == nil || IsNil(o.Data) {
		var ret []V0PipelineListResponseItemModel
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0PipelineListResponseModel) GetDataOk() ([]V0PipelineListResponseItemModel, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *V0PipelineListResponseModel) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []V0PipelineListResponseItemModel and assigns it to the Data field.
func (o *V0PipelineListResponseModel) SetData(v []V0PipelineListResponseItemModel) {
	o.Data = v
}

// GetPaging returns the Paging field value if set, zero value otherwise.
func (o *V0PipelineListResponseModel) GetPaging() V0PagingResponseModel {
	if o == nil || IsNil(o.Paging) {
		var ret V0PagingResponseModel
		return ret
	}
	return *o.Paging
}

// GetPagingOk returns a tuple with the Paging field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0PipelineListResponseModel) GetPagingOk() (*V0PagingResponseModel, bool) {
	if o == nil || IsNil(o.Paging) {
		return nil, false
	}
	return o.Paging, true
}

// HasPaging returns a boolean if a field has been set.
func (o *V0PipelineListResponseModel) HasPaging() bool {
	if o != nil && !IsNil(o.Paging) {
		return true
	}

	return false
}

// SetPaging gets a reference to the given V0PagingResponseModel and assigns it to the Paging field.
func (o *V0PipelineListResponseModel) SetPaging(v V0PagingResponseModel) {
	o.Paging = &v
}

func (o V0PipelineListResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0PipelineListResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Paging) {
		toSerialize["paging"] = o.Paging
	}
	return toSerialize, nil
}

type NullableV0PipelineListResponseModel struct {
	value *V0PipelineListResponseModel
	isSet bool
}

func (v NullableV0PipelineListResponseModel) Get() *V0PipelineListResponseModel {
	return v.value
}

func (v *NullableV0PipelineListResponseModel) Set(val *V0PipelineListResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableV0PipelineListResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableV0PipelineListResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0PipelineListResponseModel(val *V0PipelineListResponseModel) *NullableV0PipelineListResponseModel {
	return &NullableV0PipelineListResponseModel{value: val, isSet: true}
}

func (v NullableV0PipelineListResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0PipelineListResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


