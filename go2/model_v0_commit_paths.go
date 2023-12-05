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

// checks if the V0CommitPaths type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0CommitPaths{}

// V0CommitPaths struct for V0CommitPaths
type V0CommitPaths struct {
	Added []string `json:"added,omitempty"`
	Modified []string `json:"modified,omitempty"`
	Removed []string `json:"removed,omitempty"`
}

// NewV0CommitPaths instantiates a new V0CommitPaths object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0CommitPaths() *V0CommitPaths {
	this := V0CommitPaths{}
	return &this
}

// NewV0CommitPathsWithDefaults instantiates a new V0CommitPaths object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0CommitPathsWithDefaults() *V0CommitPaths {
	this := V0CommitPaths{}
	return &this
}

// GetAdded returns the Added field value if set, zero value otherwise.
func (o *V0CommitPaths) GetAdded() []string {
	if o == nil || IsNil(o.Added) {
		var ret []string
		return ret
	}
	return o.Added
}

// GetAddedOk returns a tuple with the Added field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0CommitPaths) GetAddedOk() ([]string, bool) {
	if o == nil || IsNil(o.Added) {
		return nil, false
	}
	return o.Added, true
}

// HasAdded returns a boolean if a field has been set.
func (o *V0CommitPaths) HasAdded() bool {
	if o != nil && !IsNil(o.Added) {
		return true
	}

	return false
}

// SetAdded gets a reference to the given []string and assigns it to the Added field.
func (o *V0CommitPaths) SetAdded(v []string) {
	o.Added = v
}

// GetModified returns the Modified field value if set, zero value otherwise.
func (o *V0CommitPaths) GetModified() []string {
	if o == nil || IsNil(o.Modified) {
		var ret []string
		return ret
	}
	return o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0CommitPaths) GetModifiedOk() ([]string, bool) {
	if o == nil || IsNil(o.Modified) {
		return nil, false
	}
	return o.Modified, true
}

// HasModified returns a boolean if a field has been set.
func (o *V0CommitPaths) HasModified() bool {
	if o != nil && !IsNil(o.Modified) {
		return true
	}

	return false
}

// SetModified gets a reference to the given []string and assigns it to the Modified field.
func (o *V0CommitPaths) SetModified(v []string) {
	o.Modified = v
}

// GetRemoved returns the Removed field value if set, zero value otherwise.
func (o *V0CommitPaths) GetRemoved() []string {
	if o == nil || IsNil(o.Removed) {
		var ret []string
		return ret
	}
	return o.Removed
}

// GetRemovedOk returns a tuple with the Removed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0CommitPaths) GetRemovedOk() ([]string, bool) {
	if o == nil || IsNil(o.Removed) {
		return nil, false
	}
	return o.Removed, true
}

// HasRemoved returns a boolean if a field has been set.
func (o *V0CommitPaths) HasRemoved() bool {
	if o != nil && !IsNil(o.Removed) {
		return true
	}

	return false
}

// SetRemoved gets a reference to the given []string and assigns it to the Removed field.
func (o *V0CommitPaths) SetRemoved(v []string) {
	o.Removed = v
}

func (o V0CommitPaths) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0CommitPaths) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Added) {
		toSerialize["added"] = o.Added
	}
	if !IsNil(o.Modified) {
		toSerialize["modified"] = o.Modified
	}
	if !IsNil(o.Removed) {
		toSerialize["removed"] = o.Removed
	}
	return toSerialize, nil
}

type NullableV0CommitPaths struct {
	value *V0CommitPaths
	isSet bool
}

func (v NullableV0CommitPaths) Get() *V0CommitPaths {
	return v.value
}

func (v *NullableV0CommitPaths) Set(val *V0CommitPaths) {
	v.value = val
	v.isSet = true
}

func (v NullableV0CommitPaths) IsSet() bool {
	return v.isSet
}

func (v *NullableV0CommitPaths) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0CommitPaths(val *V0CommitPaths) *NullableV0CommitPaths {
	return &NullableV0CommitPaths{value: val, isSet: true}
}

func (v NullableV0CommitPaths) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0CommitPaths) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


