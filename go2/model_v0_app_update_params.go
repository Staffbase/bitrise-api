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

// checks if the V0AppUpdateParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0AppUpdateParams{}

// V0AppUpdateParams struct for V0AppUpdateParams
type V0AppUpdateParams struct {
	// The new apple credential user ID (recommendation: use the UI to set this)
	AppleCredentialUserId *int32 `json:"apple_credential_user_id,omitempty"`
	// The new apple credential user slug (recommendation: use the UI to set this)
	AppleCredentialUserSlug *string `json:"apple_credential_user_slug,omitempty"`
	// The new default branch for the application.
	DefaultBranch *string `json:"default_branch,omitempty"`
	// The new value of whether an application should be publicly visible
	IsPublic *bool `json:"is_public,omitempty"`
	// The new repository URL for the application.
	RepositoryUrl *string `json:"repository_url,omitempty"`
	// The new service credential user ID (recommendation: use the UI to set this)
	ServicesCredentialUserId *int32 `json:"services_credential_user_id,omitempty"`
	// The new title of the application.
	Title *string `json:"title,omitempty"`
}

// NewV0AppUpdateParams instantiates a new V0AppUpdateParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0AppUpdateParams() *V0AppUpdateParams {
	this := V0AppUpdateParams{}
	return &this
}

// NewV0AppUpdateParamsWithDefaults instantiates a new V0AppUpdateParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0AppUpdateParamsWithDefaults() *V0AppUpdateParams {
	this := V0AppUpdateParams{}
	return &this
}

// GetAppleCredentialUserId returns the AppleCredentialUserId field value if set, zero value otherwise.
func (o *V0AppUpdateParams) GetAppleCredentialUserId() int32 {
	if o == nil || IsNil(o.AppleCredentialUserId) {
		var ret int32
		return ret
	}
	return *o.AppleCredentialUserId
}

// GetAppleCredentialUserIdOk returns a tuple with the AppleCredentialUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0AppUpdateParams) GetAppleCredentialUserIdOk() (*int32, bool) {
	if o == nil || IsNil(o.AppleCredentialUserId) {
		return nil, false
	}
	return o.AppleCredentialUserId, true
}

// HasAppleCredentialUserId returns a boolean if a field has been set.
func (o *V0AppUpdateParams) HasAppleCredentialUserId() bool {
	if o != nil && !IsNil(o.AppleCredentialUserId) {
		return true
	}

	return false
}

// SetAppleCredentialUserId gets a reference to the given int32 and assigns it to the AppleCredentialUserId field.
func (o *V0AppUpdateParams) SetAppleCredentialUserId(v int32) {
	o.AppleCredentialUserId = &v
}

// GetAppleCredentialUserSlug returns the AppleCredentialUserSlug field value if set, zero value otherwise.
func (o *V0AppUpdateParams) GetAppleCredentialUserSlug() string {
	if o == nil || IsNil(o.AppleCredentialUserSlug) {
		var ret string
		return ret
	}
	return *o.AppleCredentialUserSlug
}

// GetAppleCredentialUserSlugOk returns a tuple with the AppleCredentialUserSlug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0AppUpdateParams) GetAppleCredentialUserSlugOk() (*string, bool) {
	if o == nil || IsNil(o.AppleCredentialUserSlug) {
		return nil, false
	}
	return o.AppleCredentialUserSlug, true
}

// HasAppleCredentialUserSlug returns a boolean if a field has been set.
func (o *V0AppUpdateParams) HasAppleCredentialUserSlug() bool {
	if o != nil && !IsNil(o.AppleCredentialUserSlug) {
		return true
	}

	return false
}

// SetAppleCredentialUserSlug gets a reference to the given string and assigns it to the AppleCredentialUserSlug field.
func (o *V0AppUpdateParams) SetAppleCredentialUserSlug(v string) {
	o.AppleCredentialUserSlug = &v
}

// GetDefaultBranch returns the DefaultBranch field value if set, zero value otherwise.
func (o *V0AppUpdateParams) GetDefaultBranch() string {
	if o == nil || IsNil(o.DefaultBranch) {
		var ret string
		return ret
	}
	return *o.DefaultBranch
}

// GetDefaultBranchOk returns a tuple with the DefaultBranch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0AppUpdateParams) GetDefaultBranchOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultBranch) {
		return nil, false
	}
	return o.DefaultBranch, true
}

// HasDefaultBranch returns a boolean if a field has been set.
func (o *V0AppUpdateParams) HasDefaultBranch() bool {
	if o != nil && !IsNil(o.DefaultBranch) {
		return true
	}

	return false
}

// SetDefaultBranch gets a reference to the given string and assigns it to the DefaultBranch field.
func (o *V0AppUpdateParams) SetDefaultBranch(v string) {
	o.DefaultBranch = &v
}

// GetIsPublic returns the IsPublic field value if set, zero value otherwise.
func (o *V0AppUpdateParams) GetIsPublic() bool {
	if o == nil || IsNil(o.IsPublic) {
		var ret bool
		return ret
	}
	return *o.IsPublic
}

// GetIsPublicOk returns a tuple with the IsPublic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0AppUpdateParams) GetIsPublicOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPublic) {
		return nil, false
	}
	return o.IsPublic, true
}

// HasIsPublic returns a boolean if a field has been set.
func (o *V0AppUpdateParams) HasIsPublic() bool {
	if o != nil && !IsNil(o.IsPublic) {
		return true
	}

	return false
}

// SetIsPublic gets a reference to the given bool and assigns it to the IsPublic field.
func (o *V0AppUpdateParams) SetIsPublic(v bool) {
	o.IsPublic = &v
}

// GetRepositoryUrl returns the RepositoryUrl field value if set, zero value otherwise.
func (o *V0AppUpdateParams) GetRepositoryUrl() string {
	if o == nil || IsNil(o.RepositoryUrl) {
		var ret string
		return ret
	}
	return *o.RepositoryUrl
}

// GetRepositoryUrlOk returns a tuple with the RepositoryUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0AppUpdateParams) GetRepositoryUrlOk() (*string, bool) {
	if o == nil || IsNil(o.RepositoryUrl) {
		return nil, false
	}
	return o.RepositoryUrl, true
}

// HasRepositoryUrl returns a boolean if a field has been set.
func (o *V0AppUpdateParams) HasRepositoryUrl() bool {
	if o != nil && !IsNil(o.RepositoryUrl) {
		return true
	}

	return false
}

// SetRepositoryUrl gets a reference to the given string and assigns it to the RepositoryUrl field.
func (o *V0AppUpdateParams) SetRepositoryUrl(v string) {
	o.RepositoryUrl = &v
}

// GetServicesCredentialUserId returns the ServicesCredentialUserId field value if set, zero value otherwise.
func (o *V0AppUpdateParams) GetServicesCredentialUserId() int32 {
	if o == nil || IsNil(o.ServicesCredentialUserId) {
		var ret int32
		return ret
	}
	return *o.ServicesCredentialUserId
}

// GetServicesCredentialUserIdOk returns a tuple with the ServicesCredentialUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0AppUpdateParams) GetServicesCredentialUserIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ServicesCredentialUserId) {
		return nil, false
	}
	return o.ServicesCredentialUserId, true
}

// HasServicesCredentialUserId returns a boolean if a field has been set.
func (o *V0AppUpdateParams) HasServicesCredentialUserId() bool {
	if o != nil && !IsNil(o.ServicesCredentialUserId) {
		return true
	}

	return false
}

// SetServicesCredentialUserId gets a reference to the given int32 and assigns it to the ServicesCredentialUserId field.
func (o *V0AppUpdateParams) SetServicesCredentialUserId(v int32) {
	o.ServicesCredentialUserId = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *V0AppUpdateParams) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0AppUpdateParams) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *V0AppUpdateParams) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *V0AppUpdateParams) SetTitle(v string) {
	o.Title = &v
}

func (o V0AppUpdateParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0AppUpdateParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AppleCredentialUserId) {
		toSerialize["apple_credential_user_id"] = o.AppleCredentialUserId
	}
	if !IsNil(o.AppleCredentialUserSlug) {
		toSerialize["apple_credential_user_slug"] = o.AppleCredentialUserSlug
	}
	if !IsNil(o.DefaultBranch) {
		toSerialize["default_branch"] = o.DefaultBranch
	}
	if !IsNil(o.IsPublic) {
		toSerialize["is_public"] = o.IsPublic
	}
	if !IsNil(o.RepositoryUrl) {
		toSerialize["repository_url"] = o.RepositoryUrl
	}
	if !IsNil(o.ServicesCredentialUserId) {
		toSerialize["services_credential_user_id"] = o.ServicesCredentialUserId
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	return toSerialize, nil
}

type NullableV0AppUpdateParams struct {
	value *V0AppUpdateParams
	isSet bool
}

func (v NullableV0AppUpdateParams) Get() *V0AppUpdateParams {
	return v.value
}

func (v *NullableV0AppUpdateParams) Set(val *V0AppUpdateParams) {
	v.value = val
	v.isSet = true
}

func (v NullableV0AppUpdateParams) IsSet() bool {
	return v.isSet
}

func (v *NullableV0AppUpdateParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0AppUpdateParams(val *V0AppUpdateParams) *NullableV0AppUpdateParams {
	return &NullableV0AppUpdateParams{value: val, isSet: true}
}

func (v NullableV0AppUpdateParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0AppUpdateParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


