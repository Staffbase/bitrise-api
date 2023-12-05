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
	"fmt"
)

// checks if the V0AppConfigRequestParam type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0AppConfigRequestParam{}

// V0AppConfigRequestParam struct for V0AppConfigRequestParam
type V0AppConfigRequestParam struct {
	// The bitrise.yml of your application, defined in JSON format
	AppConfigDatastoreYaml string `json:"app_config_datastore_yaml"`
}

type _V0AppConfigRequestParam V0AppConfigRequestParam

// NewV0AppConfigRequestParam instantiates a new V0AppConfigRequestParam object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0AppConfigRequestParam(appConfigDatastoreYaml string) *V0AppConfigRequestParam {
	this := V0AppConfigRequestParam{}
	this.AppConfigDatastoreYaml = appConfigDatastoreYaml
	return &this
}

// NewV0AppConfigRequestParamWithDefaults instantiates a new V0AppConfigRequestParam object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0AppConfigRequestParamWithDefaults() *V0AppConfigRequestParam {
	this := V0AppConfigRequestParam{}
	return &this
}

// GetAppConfigDatastoreYaml returns the AppConfigDatastoreYaml field value
func (o *V0AppConfigRequestParam) GetAppConfigDatastoreYaml() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AppConfigDatastoreYaml
}

// GetAppConfigDatastoreYamlOk returns a tuple with the AppConfigDatastoreYaml field value
// and a boolean to check if the value has been set.
func (o *V0AppConfigRequestParam) GetAppConfigDatastoreYamlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AppConfigDatastoreYaml, true
}

// SetAppConfigDatastoreYaml sets field value
func (o *V0AppConfigRequestParam) SetAppConfigDatastoreYaml(v string) {
	o.AppConfigDatastoreYaml = v
}

func (o V0AppConfigRequestParam) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0AppConfigRequestParam) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["app_config_datastore_yaml"] = o.AppConfigDatastoreYaml
	return toSerialize, nil
}

func (o *V0AppConfigRequestParam) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"app_config_datastore_yaml",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varV0AppConfigRequestParam := _V0AppConfigRequestParam{}

	err = json.Unmarshal(bytes, &varV0AppConfigRequestParam)

	if err != nil {
		return err
	}

	*o = V0AppConfigRequestParam(varV0AppConfigRequestParam)

	return err
}

type NullableV0AppConfigRequestParam struct {
	value *V0AppConfigRequestParam
	isSet bool
}

func (v NullableV0AppConfigRequestParam) Get() *V0AppConfigRequestParam {
	return v.value
}

func (v *NullableV0AppConfigRequestParam) Set(val *V0AppConfigRequestParam) {
	v.value = val
	v.isSet = true
}

func (v NullableV0AppConfigRequestParam) IsSet() bool {
	return v.isSet
}

func (v *NullableV0AppConfigRequestParam) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0AppConfigRequestParam(val *V0AppConfigRequestParam) *NullableV0AppConfigRequestParam {
	return &NullableV0AppConfigRequestParam{value: val, isSet: true}
}

func (v NullableV0AppConfigRequestParam) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0AppConfigRequestParam) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


