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

// checks if the V0AppFinishParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0AppFinishParams{}

// V0AppFinishParams struct for V0AppFinishParams
type V0AppFinishParams struct {
	// Which config to use. Specify a config that matches your project type (e. g. `default-android-config` for `android`, etc.). If not speficied, default value is `other-config`. The available values are `default-android-config`, `default-cordova-config`, `default-fastlane-android-config`, `default-fastlane-ios-config`, `flutter-config-notest-app-android`, `flutter-config-notest-app-both`, `flutter-config-notest-app-ios`, `flutter-config-test-app-android`, `flutter-config-test-app-both`, `flutter-config-test-app-ios`, `default-ionic-config`, `default-ios-config`, `default-macos-config`, `default-react-native-config`, `default-react-native-expo-config`, `other-config`.
	Config *string `json:"config,omitempty"`
	// Environment variables for the application workflows, e.g. {\"env1\":\"val1\",\"env2\":\"val2\"}
	Envs *map[string]string `json:"envs,omitempty"`
	// config specification mode, has to be specified with `manual` value
	Mode *string `json:"mode,omitempty"`
	// The slug of the organization, who will be the owner of the application, if it's not specified, then the authenticated user will be the owner
	OrganizationSlug *string `json:"organization_slug,omitempty"`
	// The type of your project (`android`, `ios`, `cordova`, `other`, `xamarin`, `macos`, `ionic`, `react-native`, `fastlane`, null)
	ProjectType string `json:"project_type"`
	// The id of the stack the application will be built (these can be found in the [system report](https://github.com/bitrise-io/bitrise.io/tree/master/system_reports) file names)
	StackId string `json:"stack_id"`
}

type _V0AppFinishParams V0AppFinishParams

// NewV0AppFinishParams instantiates a new V0AppFinishParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0AppFinishParams(projectType string, stackId string) *V0AppFinishParams {
	this := V0AppFinishParams{}
	this.ProjectType = projectType
	this.StackId = stackId
	return &this
}

// NewV0AppFinishParamsWithDefaults instantiates a new V0AppFinishParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0AppFinishParamsWithDefaults() *V0AppFinishParams {
	this := V0AppFinishParams{}
	return &this
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *V0AppFinishParams) GetConfig() string {
	if o == nil || IsNil(o.Config) {
		var ret string
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0AppFinishParams) GetConfigOk() (*string, bool) {
	if o == nil || IsNil(o.Config) {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *V0AppFinishParams) HasConfig() bool {
	if o != nil && !IsNil(o.Config) {
		return true
	}

	return false
}

// SetConfig gets a reference to the given string and assigns it to the Config field.
func (o *V0AppFinishParams) SetConfig(v string) {
	o.Config = &v
}

// GetEnvs returns the Envs field value if set, zero value otherwise.
func (o *V0AppFinishParams) GetEnvs() map[string]string {
	if o == nil || IsNil(o.Envs) {
		var ret map[string]string
		return ret
	}
	return *o.Envs
}

// GetEnvsOk returns a tuple with the Envs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0AppFinishParams) GetEnvsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Envs) {
		return nil, false
	}
	return o.Envs, true
}

// HasEnvs returns a boolean if a field has been set.
func (o *V0AppFinishParams) HasEnvs() bool {
	if o != nil && !IsNil(o.Envs) {
		return true
	}

	return false
}

// SetEnvs gets a reference to the given map[string]string and assigns it to the Envs field.
func (o *V0AppFinishParams) SetEnvs(v map[string]string) {
	o.Envs = &v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *V0AppFinishParams) GetMode() string {
	if o == nil || IsNil(o.Mode) {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0AppFinishParams) GetModeOk() (*string, bool) {
	if o == nil || IsNil(o.Mode) {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *V0AppFinishParams) HasMode() bool {
	if o != nil && !IsNil(o.Mode) {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *V0AppFinishParams) SetMode(v string) {
	o.Mode = &v
}

// GetOrganizationSlug returns the OrganizationSlug field value if set, zero value otherwise.
func (o *V0AppFinishParams) GetOrganizationSlug() string {
	if o == nil || IsNil(o.OrganizationSlug) {
		var ret string
		return ret
	}
	return *o.OrganizationSlug
}

// GetOrganizationSlugOk returns a tuple with the OrganizationSlug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0AppFinishParams) GetOrganizationSlugOk() (*string, bool) {
	if o == nil || IsNil(o.OrganizationSlug) {
		return nil, false
	}
	return o.OrganizationSlug, true
}

// HasOrganizationSlug returns a boolean if a field has been set.
func (o *V0AppFinishParams) HasOrganizationSlug() bool {
	if o != nil && !IsNil(o.OrganizationSlug) {
		return true
	}

	return false
}

// SetOrganizationSlug gets a reference to the given string and assigns it to the OrganizationSlug field.
func (o *V0AppFinishParams) SetOrganizationSlug(v string) {
	o.OrganizationSlug = &v
}

// GetProjectType returns the ProjectType field value
func (o *V0AppFinishParams) GetProjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectType
}

// GetProjectTypeOk returns a tuple with the ProjectType field value
// and a boolean to check if the value has been set.
func (o *V0AppFinishParams) GetProjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectType, true
}

// SetProjectType sets field value
func (o *V0AppFinishParams) SetProjectType(v string) {
	o.ProjectType = v
}

// GetStackId returns the StackId field value
func (o *V0AppFinishParams) GetStackId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StackId
}

// GetStackIdOk returns a tuple with the StackId field value
// and a boolean to check if the value has been set.
func (o *V0AppFinishParams) GetStackIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StackId, true
}

// SetStackId sets field value
func (o *V0AppFinishParams) SetStackId(v string) {
	o.StackId = v
}

func (o V0AppFinishParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0AppFinishParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Config) {
		toSerialize["config"] = o.Config
	}
	if !IsNil(o.Envs) {
		toSerialize["envs"] = o.Envs
	}
	if !IsNil(o.Mode) {
		toSerialize["mode"] = o.Mode
	}
	if !IsNil(o.OrganizationSlug) {
		toSerialize["organization_slug"] = o.OrganizationSlug
	}
	toSerialize["project_type"] = o.ProjectType
	toSerialize["stack_id"] = o.StackId
	return toSerialize, nil
}

func (o *V0AppFinishParams) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"project_type",
		"stack_id",
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

	varV0AppFinishParams := _V0AppFinishParams{}

	err = json.Unmarshal(bytes, &varV0AppFinishParams)

	if err != nil {
		return err
	}

	*o = V0AppFinishParams(varV0AppFinishParams)

	return err
}

type NullableV0AppFinishParams struct {
	value *V0AppFinishParams
	isSet bool
}

func (v NullableV0AppFinishParams) Get() *V0AppFinishParams {
	return v.value
}

func (v *NullableV0AppFinishParams) Set(val *V0AppFinishParams) {
	v.value = val
	v.isSet = true
}

func (v NullableV0AppFinishParams) IsSet() bool {
	return v.isSet
}

func (v *NullableV0AppFinishParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0AppFinishParams(val *V0AppFinishParams) *NullableV0AppFinishParams {
	return &NullableV0AppFinishParams{value: val, isSet: true}
}

func (v NullableV0AppFinishParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0AppFinishParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


