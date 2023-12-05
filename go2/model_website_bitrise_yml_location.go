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

// WebsiteBitriseYMLLocation the model 'WebsiteBitriseYMLLocation'
type WebsiteBitriseYMLLocation string

// List of website.BitriseYMLLocation
const (
	LocationBitriseIo WebsiteBitriseYMLLocation = "bitrise.io"
	LocationRepository WebsiteBitriseYMLLocation = "repository"
)

// All allowed values of WebsiteBitriseYMLLocation enum
var AllowedWebsiteBitriseYMLLocationEnumValues = []WebsiteBitriseYMLLocation{
	"bitrise.io",
	"repository",
}

func (v *WebsiteBitriseYMLLocation) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := WebsiteBitriseYMLLocation(value)
	for _, existing := range AllowedWebsiteBitriseYMLLocationEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid WebsiteBitriseYMLLocation", value)
}

// NewWebsiteBitriseYMLLocationFromValue returns a pointer to a valid WebsiteBitriseYMLLocation
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewWebsiteBitriseYMLLocationFromValue(v string) (*WebsiteBitriseYMLLocation, error) {
	ev := WebsiteBitriseYMLLocation(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for WebsiteBitriseYMLLocation: valid values are %v", v, AllowedWebsiteBitriseYMLLocationEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v WebsiteBitriseYMLLocation) IsValid() bool {
	for _, existing := range AllowedWebsiteBitriseYMLLocationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to website.BitriseYMLLocation value
func (v WebsiteBitriseYMLLocation) Ptr() *WebsiteBitriseYMLLocation {
	return &v
}

type NullableWebsiteBitriseYMLLocation struct {
	value *WebsiteBitriseYMLLocation
	isSet bool
}

func (v NullableWebsiteBitriseYMLLocation) Get() *WebsiteBitriseYMLLocation {
	return v.value
}

func (v *NullableWebsiteBitriseYMLLocation) Set(val *WebsiteBitriseYMLLocation) {
	v.value = val
	v.isSet = true
}

func (v NullableWebsiteBitriseYMLLocation) IsSet() bool {
	return v.isSet
}

func (v *NullableWebsiteBitriseYMLLocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebsiteBitriseYMLLocation(val *WebsiteBitriseYMLLocation) *NullableWebsiteBitriseYMLLocation {
	return &NullableWebsiteBitriseYMLLocation{value: val, isSet: true}
}

func (v NullableWebsiteBitriseYMLLocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebsiteBitriseYMLLocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
