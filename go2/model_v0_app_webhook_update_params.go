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

// checks if the V0AppWebhookUpdateParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0AppWebhookUpdateParams{}

// V0AppWebhookUpdateParams struct for V0AppWebhookUpdateParams
type V0AppWebhookUpdateParams struct {
	Events []string `json:"events"`
	Headers []int32 `json:"headers,omitempty"`
	Url string `json:"url"`
}

type _V0AppWebhookUpdateParams V0AppWebhookUpdateParams

// NewV0AppWebhookUpdateParams instantiates a new V0AppWebhookUpdateParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0AppWebhookUpdateParams(events []string, url string) *V0AppWebhookUpdateParams {
	this := V0AppWebhookUpdateParams{}
	this.Events = events
	this.Url = url
	return &this
}

// NewV0AppWebhookUpdateParamsWithDefaults instantiates a new V0AppWebhookUpdateParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0AppWebhookUpdateParamsWithDefaults() *V0AppWebhookUpdateParams {
	this := V0AppWebhookUpdateParams{}
	return &this
}

// GetEvents returns the Events field value
func (o *V0AppWebhookUpdateParams) GetEvents() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Events
}

// GetEventsOk returns a tuple with the Events field value
// and a boolean to check if the value has been set.
func (o *V0AppWebhookUpdateParams) GetEventsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Events, true
}

// SetEvents sets field value
func (o *V0AppWebhookUpdateParams) SetEvents(v []string) {
	o.Events = v
}

// GetHeaders returns the Headers field value if set, zero value otherwise.
func (o *V0AppWebhookUpdateParams) GetHeaders() []int32 {
	if o == nil || IsNil(o.Headers) {
		var ret []int32
		return ret
	}
	return o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0AppWebhookUpdateParams) GetHeadersOk() ([]int32, bool) {
	if o == nil || IsNil(o.Headers) {
		return nil, false
	}
	return o.Headers, true
}

// HasHeaders returns a boolean if a field has been set.
func (o *V0AppWebhookUpdateParams) HasHeaders() bool {
	if o != nil && !IsNil(o.Headers) {
		return true
	}

	return false
}

// SetHeaders gets a reference to the given []int32 and assigns it to the Headers field.
func (o *V0AppWebhookUpdateParams) SetHeaders(v []int32) {
	o.Headers = v
}

// GetUrl returns the Url field value
func (o *V0AppWebhookUpdateParams) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *V0AppWebhookUpdateParams) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *V0AppWebhookUpdateParams) SetUrl(v string) {
	o.Url = v
}

func (o V0AppWebhookUpdateParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0AppWebhookUpdateParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["events"] = o.Events
	if !IsNil(o.Headers) {
		toSerialize["headers"] = o.Headers
	}
	toSerialize["url"] = o.Url
	return toSerialize, nil
}

func (o *V0AppWebhookUpdateParams) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"events",
		"url",
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

	varV0AppWebhookUpdateParams := _V0AppWebhookUpdateParams{}

	err = json.Unmarshal(bytes, &varV0AppWebhookUpdateParams)

	if err != nil {
		return err
	}

	*o = V0AppWebhookUpdateParams(varV0AppWebhookUpdateParams)

	return err
}

type NullableV0AppWebhookUpdateParams struct {
	value *V0AppWebhookUpdateParams
	isSet bool
}

func (v NullableV0AppWebhookUpdateParams) Get() *V0AppWebhookUpdateParams {
	return v.value
}

func (v *NullableV0AppWebhookUpdateParams) Set(val *V0AppWebhookUpdateParams) {
	v.value = val
	v.isSet = true
}

func (v NullableV0AppWebhookUpdateParams) IsSet() bool {
	return v.isSet
}

func (v *NullableV0AppWebhookUpdateParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0AppWebhookUpdateParams(val *V0AppWebhookUpdateParams) *NullableV0AppWebhookUpdateParams {
	return &NullableV0AppWebhookUpdateParams{value: val, isSet: true}
}

func (v NullableV0AppWebhookUpdateParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0AppWebhookUpdateParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


