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

// checks if the V0CacheItemDownloadURLResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0CacheItemDownloadURLResponseModel{}

// V0CacheItemDownloadURLResponseModel struct for V0CacheItemDownloadURLResponseModel
type V0CacheItemDownloadURLResponseModel struct {
	DownloadUrl *string `json:"download_url,omitempty"`
}

// NewV0CacheItemDownloadURLResponseModel instantiates a new V0CacheItemDownloadURLResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0CacheItemDownloadURLResponseModel() *V0CacheItemDownloadURLResponseModel {
	this := V0CacheItemDownloadURLResponseModel{}
	return &this
}

// NewV0CacheItemDownloadURLResponseModelWithDefaults instantiates a new V0CacheItemDownloadURLResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0CacheItemDownloadURLResponseModelWithDefaults() *V0CacheItemDownloadURLResponseModel {
	this := V0CacheItemDownloadURLResponseModel{}
	return &this
}

// GetDownloadUrl returns the DownloadUrl field value if set, zero value otherwise.
func (o *V0CacheItemDownloadURLResponseModel) GetDownloadUrl() string {
	if o == nil || IsNil(o.DownloadUrl) {
		var ret string
		return ret
	}
	return *o.DownloadUrl
}

// GetDownloadUrlOk returns a tuple with the DownloadUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0CacheItemDownloadURLResponseModel) GetDownloadUrlOk() (*string, bool) {
	if o == nil || IsNil(o.DownloadUrl) {
		return nil, false
	}
	return o.DownloadUrl, true
}

// HasDownloadUrl returns a boolean if a field has been set.
func (o *V0CacheItemDownloadURLResponseModel) HasDownloadUrl() bool {
	if o != nil && !IsNil(o.DownloadUrl) {
		return true
	}

	return false
}

// SetDownloadUrl gets a reference to the given string and assigns it to the DownloadUrl field.
func (o *V0CacheItemDownloadURLResponseModel) SetDownloadUrl(v string) {
	o.DownloadUrl = &v
}

func (o V0CacheItemDownloadURLResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0CacheItemDownloadURLResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DownloadUrl) {
		toSerialize["download_url"] = o.DownloadUrl
	}
	return toSerialize, nil
}

type NullableV0CacheItemDownloadURLResponseModel struct {
	value *V0CacheItemDownloadURLResponseModel
	isSet bool
}

func (v NullableV0CacheItemDownloadURLResponseModel) Get() *V0CacheItemDownloadURLResponseModel {
	return v.value
}

func (v *NullableV0CacheItemDownloadURLResponseModel) Set(val *V0CacheItemDownloadURLResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableV0CacheItemDownloadURLResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableV0CacheItemDownloadURLResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0CacheItemDownloadURLResponseModel(val *V0CacheItemDownloadURLResponseModel) *NullableV0CacheItemDownloadURLResponseModel {
	return &NullableV0CacheItemDownloadURLResponseModel{value: val, isSet: true}
}

func (v NullableV0CacheItemDownloadURLResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0CacheItemDownloadURLResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


