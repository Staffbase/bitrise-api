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

// checks if the AddonsDeveloperLink type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddonsDeveloperLink{}

// AddonsDeveloperLink struct for AddonsDeveloperLink
type AddonsDeveloperLink struct {
	Title *string `json:"title,omitempty"`
	Url *string `json:"url,omitempty"`
}

// NewAddonsDeveloperLink instantiates a new AddonsDeveloperLink object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddonsDeveloperLink() *AddonsDeveloperLink {
	this := AddonsDeveloperLink{}
	return &this
}

// NewAddonsDeveloperLinkWithDefaults instantiates a new AddonsDeveloperLink object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddonsDeveloperLinkWithDefaults() *AddonsDeveloperLink {
	this := AddonsDeveloperLink{}
	return &this
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *AddonsDeveloperLink) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddonsDeveloperLink) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *AddonsDeveloperLink) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *AddonsDeveloperLink) SetTitle(v string) {
	o.Title = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *AddonsDeveloperLink) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddonsDeveloperLink) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *AddonsDeveloperLink) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *AddonsDeveloperLink) SetUrl(v string) {
	o.Url = &v
}

func (o AddonsDeveloperLink) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddonsDeveloperLink) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return toSerialize, nil
}

type NullableAddonsDeveloperLink struct {
	value *AddonsDeveloperLink
	isSet bool
}

func (v NullableAddonsDeveloperLink) Get() *AddonsDeveloperLink {
	return v.value
}

func (v *NullableAddonsDeveloperLink) Set(val *AddonsDeveloperLink) {
	v.value = val
	v.isSet = true
}

func (v NullableAddonsDeveloperLink) IsSet() bool {
	return v.isSet
}

func (v *NullableAddonsDeveloperLink) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddonsDeveloperLink(val *AddonsDeveloperLink) *NullableAddonsDeveloperLink {
	return &NullableAddonsDeveloperLink{value: val, isSet: true}
}

func (v NullableAddonsDeveloperLink) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddonsDeveloperLink) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

