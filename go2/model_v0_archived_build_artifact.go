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

// checks if the V0ArchivedBuildArtifact type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0ArchivedBuildArtifact{}

// V0ArchivedBuildArtifact struct for V0ArchivedBuildArtifact
type V0ArchivedBuildArtifact struct {
	ArtifactMeta map[string]interface{} `json:"artifact_meta,omitempty"`
	ArtifactType *string `json:"artifact_type,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	FileSizeBytes *int32 `json:"file_size_bytes,omitempty"`
	Id *string `json:"id,omitempty"`
	IsPublicPageEnabled *bool `json:"is_public_page_enabled,omitempty"`
	PublicPageToken *string `json:"public_page_token,omitempty"`
	Title *string `json:"title,omitempty"`
}

// NewV0ArchivedBuildArtifact instantiates a new V0ArchivedBuildArtifact object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0ArchivedBuildArtifact() *V0ArchivedBuildArtifact {
	this := V0ArchivedBuildArtifact{}
	return &this
}

// NewV0ArchivedBuildArtifactWithDefaults instantiates a new V0ArchivedBuildArtifact object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0ArchivedBuildArtifactWithDefaults() *V0ArchivedBuildArtifact {
	this := V0ArchivedBuildArtifact{}
	return &this
}

// GetArtifactMeta returns the ArtifactMeta field value if set, zero value otherwise.
func (o *V0ArchivedBuildArtifact) GetArtifactMeta() map[string]interface{} {
	if o == nil || IsNil(o.ArtifactMeta) {
		var ret map[string]interface{}
		return ret
	}
	return o.ArtifactMeta
}

// GetArtifactMetaOk returns a tuple with the ArtifactMeta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0ArchivedBuildArtifact) GetArtifactMetaOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.ArtifactMeta) {
		return map[string]interface{}{}, false
	}
	return o.ArtifactMeta, true
}

// HasArtifactMeta returns a boolean if a field has been set.
func (o *V0ArchivedBuildArtifact) HasArtifactMeta() bool {
	if o != nil && !IsNil(o.ArtifactMeta) {
		return true
	}

	return false
}

// SetArtifactMeta gets a reference to the given map[string]interface{} and assigns it to the ArtifactMeta field.
func (o *V0ArchivedBuildArtifact) SetArtifactMeta(v map[string]interface{}) {
	o.ArtifactMeta = v
}

// GetArtifactType returns the ArtifactType field value if set, zero value otherwise.
func (o *V0ArchivedBuildArtifact) GetArtifactType() string {
	if o == nil || IsNil(o.ArtifactType) {
		var ret string
		return ret
	}
	return *o.ArtifactType
}

// GetArtifactTypeOk returns a tuple with the ArtifactType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0ArchivedBuildArtifact) GetArtifactTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ArtifactType) {
		return nil, false
	}
	return o.ArtifactType, true
}

// HasArtifactType returns a boolean if a field has been set.
func (o *V0ArchivedBuildArtifact) HasArtifactType() bool {
	if o != nil && !IsNil(o.ArtifactType) {
		return true
	}

	return false
}

// SetArtifactType gets a reference to the given string and assigns it to the ArtifactType field.
func (o *V0ArchivedBuildArtifact) SetArtifactType(v string) {
	o.ArtifactType = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *V0ArchivedBuildArtifact) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0ArchivedBuildArtifact) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *V0ArchivedBuildArtifact) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *V0ArchivedBuildArtifact) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetFileSizeBytes returns the FileSizeBytes field value if set, zero value otherwise.
func (o *V0ArchivedBuildArtifact) GetFileSizeBytes() int32 {
	if o == nil || IsNil(o.FileSizeBytes) {
		var ret int32
		return ret
	}
	return *o.FileSizeBytes
}

// GetFileSizeBytesOk returns a tuple with the FileSizeBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0ArchivedBuildArtifact) GetFileSizeBytesOk() (*int32, bool) {
	if o == nil || IsNil(o.FileSizeBytes) {
		return nil, false
	}
	return o.FileSizeBytes, true
}

// HasFileSizeBytes returns a boolean if a field has been set.
func (o *V0ArchivedBuildArtifact) HasFileSizeBytes() bool {
	if o != nil && !IsNil(o.FileSizeBytes) {
		return true
	}

	return false
}

// SetFileSizeBytes gets a reference to the given int32 and assigns it to the FileSizeBytes field.
func (o *V0ArchivedBuildArtifact) SetFileSizeBytes(v int32) {
	o.FileSizeBytes = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *V0ArchivedBuildArtifact) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0ArchivedBuildArtifact) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *V0ArchivedBuildArtifact) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *V0ArchivedBuildArtifact) SetId(v string) {
	o.Id = &v
}

// GetIsPublicPageEnabled returns the IsPublicPageEnabled field value if set, zero value otherwise.
func (o *V0ArchivedBuildArtifact) GetIsPublicPageEnabled() bool {
	if o == nil || IsNil(o.IsPublicPageEnabled) {
		var ret bool
		return ret
	}
	return *o.IsPublicPageEnabled
}

// GetIsPublicPageEnabledOk returns a tuple with the IsPublicPageEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0ArchivedBuildArtifact) GetIsPublicPageEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPublicPageEnabled) {
		return nil, false
	}
	return o.IsPublicPageEnabled, true
}

// HasIsPublicPageEnabled returns a boolean if a field has been set.
func (o *V0ArchivedBuildArtifact) HasIsPublicPageEnabled() bool {
	if o != nil && !IsNil(o.IsPublicPageEnabled) {
		return true
	}

	return false
}

// SetIsPublicPageEnabled gets a reference to the given bool and assigns it to the IsPublicPageEnabled field.
func (o *V0ArchivedBuildArtifact) SetIsPublicPageEnabled(v bool) {
	o.IsPublicPageEnabled = &v
}

// GetPublicPageToken returns the PublicPageToken field value if set, zero value otherwise.
func (o *V0ArchivedBuildArtifact) GetPublicPageToken() string {
	if o == nil || IsNil(o.PublicPageToken) {
		var ret string
		return ret
	}
	return *o.PublicPageToken
}

// GetPublicPageTokenOk returns a tuple with the PublicPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0ArchivedBuildArtifact) GetPublicPageTokenOk() (*string, bool) {
	if o == nil || IsNil(o.PublicPageToken) {
		return nil, false
	}
	return o.PublicPageToken, true
}

// HasPublicPageToken returns a boolean if a field has been set.
func (o *V0ArchivedBuildArtifact) HasPublicPageToken() bool {
	if o != nil && !IsNil(o.PublicPageToken) {
		return true
	}

	return false
}

// SetPublicPageToken gets a reference to the given string and assigns it to the PublicPageToken field.
func (o *V0ArchivedBuildArtifact) SetPublicPageToken(v string) {
	o.PublicPageToken = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *V0ArchivedBuildArtifact) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0ArchivedBuildArtifact) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *V0ArchivedBuildArtifact) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *V0ArchivedBuildArtifact) SetTitle(v string) {
	o.Title = &v
}

func (o V0ArchivedBuildArtifact) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0ArchivedBuildArtifact) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ArtifactMeta) {
		toSerialize["artifact_meta"] = o.ArtifactMeta
	}
	if !IsNil(o.ArtifactType) {
		toSerialize["artifact_type"] = o.ArtifactType
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.FileSizeBytes) {
		toSerialize["file_size_bytes"] = o.FileSizeBytes
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.IsPublicPageEnabled) {
		toSerialize["is_public_page_enabled"] = o.IsPublicPageEnabled
	}
	if !IsNil(o.PublicPageToken) {
		toSerialize["public_page_token"] = o.PublicPageToken
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	return toSerialize, nil
}

type NullableV0ArchivedBuildArtifact struct {
	value *V0ArchivedBuildArtifact
	isSet bool
}

func (v NullableV0ArchivedBuildArtifact) Get() *V0ArchivedBuildArtifact {
	return v.value
}

func (v *NullableV0ArchivedBuildArtifact) Set(val *V0ArchivedBuildArtifact) {
	v.value = val
	v.isSet = true
}

func (v NullableV0ArchivedBuildArtifact) IsSet() bool {
	return v.isSet
}

func (v *NullableV0ArchivedBuildArtifact) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0ArchivedBuildArtifact(val *V0ArchivedBuildArtifact) *NullableV0ArchivedBuildArtifact {
	return &NullableV0ArchivedBuildArtifact{value: val, isSet: true}
}

func (v NullableV0ArchivedBuildArtifact) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0ArchivedBuildArtifact) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


