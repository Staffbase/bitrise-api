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

// checks if the V0ArtifactResponseItemModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0ArtifactResponseItemModel{}

// V0ArtifactResponseItemModel struct for V0ArtifactResponseItemModel
type V0ArtifactResponseItemModel struct {
	ArtifactMeta []int32 `json:"artifact_meta,omitempty"`
	ArtifactType *NullsString `json:"artifact_type,omitempty"`
	ExpiringDownloadUrl *string `json:"expiring_download_url,omitempty"`
	FileSizeBytes *int32 `json:"file_size_bytes,omitempty"`
	IntermediateFileInfo []int32 `json:"intermediate_file_info,omitempty"`
	IsPublicPageEnabled *bool `json:"is_public_page_enabled,omitempty"`
	PublicInstallPageUrl *string `json:"public_install_page_url,omitempty"`
	Slug *string `json:"slug,omitempty"`
	Title *NullsString `json:"title,omitempty"`
}

// NewV0ArtifactResponseItemModel instantiates a new V0ArtifactResponseItemModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0ArtifactResponseItemModel() *V0ArtifactResponseItemModel {
	this := V0ArtifactResponseItemModel{}
	return &this
}

// NewV0ArtifactResponseItemModelWithDefaults instantiates a new V0ArtifactResponseItemModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0ArtifactResponseItemModelWithDefaults() *V0ArtifactResponseItemModel {
	this := V0ArtifactResponseItemModel{}
	return &this
}

// GetArtifactMeta returns the ArtifactMeta field value if set, zero value otherwise.
func (o *V0ArtifactResponseItemModel) GetArtifactMeta() []int32 {
	if o == nil || IsNil(o.ArtifactMeta) {
		var ret []int32
		return ret
	}
	return o.ArtifactMeta
}

// GetArtifactMetaOk returns a tuple with the ArtifactMeta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0ArtifactResponseItemModel) GetArtifactMetaOk() ([]int32, bool) {
	if o == nil || IsNil(o.ArtifactMeta) {
		return nil, false
	}
	return o.ArtifactMeta, true
}

// HasArtifactMeta returns a boolean if a field has been set.
func (o *V0ArtifactResponseItemModel) HasArtifactMeta() bool {
	if o != nil && !IsNil(o.ArtifactMeta) {
		return true
	}

	return false
}

// SetArtifactMeta gets a reference to the given []int32 and assigns it to the ArtifactMeta field.
func (o *V0ArtifactResponseItemModel) SetArtifactMeta(v []int32) {
	o.ArtifactMeta = v
}

// GetArtifactType returns the ArtifactType field value if set, zero value otherwise.
func (o *V0ArtifactResponseItemModel) GetArtifactType() NullsString {
	if o == nil || IsNil(o.ArtifactType) {
		var ret NullsString
		return ret
	}
	return *o.ArtifactType
}

// GetArtifactTypeOk returns a tuple with the ArtifactType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0ArtifactResponseItemModel) GetArtifactTypeOk() (*NullsString, bool) {
	if o == nil || IsNil(o.ArtifactType) {
		return nil, false
	}
	return o.ArtifactType, true
}

// HasArtifactType returns a boolean if a field has been set.
func (o *V0ArtifactResponseItemModel) HasArtifactType() bool {
	if o != nil && !IsNil(o.ArtifactType) {
		return true
	}

	return false
}

// SetArtifactType gets a reference to the given NullsString and assigns it to the ArtifactType field.
func (o *V0ArtifactResponseItemModel) SetArtifactType(v NullsString) {
	o.ArtifactType = &v
}

// GetExpiringDownloadUrl returns the ExpiringDownloadUrl field value if set, zero value otherwise.
func (o *V0ArtifactResponseItemModel) GetExpiringDownloadUrl() string {
	if o == nil || IsNil(o.ExpiringDownloadUrl) {
		var ret string
		return ret
	}
	return *o.ExpiringDownloadUrl
}

// GetExpiringDownloadUrlOk returns a tuple with the ExpiringDownloadUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0ArtifactResponseItemModel) GetExpiringDownloadUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ExpiringDownloadUrl) {
		return nil, false
	}
	return o.ExpiringDownloadUrl, true
}

// HasExpiringDownloadUrl returns a boolean if a field has been set.
func (o *V0ArtifactResponseItemModel) HasExpiringDownloadUrl() bool {
	if o != nil && !IsNil(o.ExpiringDownloadUrl) {
		return true
	}

	return false
}

// SetExpiringDownloadUrl gets a reference to the given string and assigns it to the ExpiringDownloadUrl field.
func (o *V0ArtifactResponseItemModel) SetExpiringDownloadUrl(v string) {
	o.ExpiringDownloadUrl = &v
}

// GetFileSizeBytes returns the FileSizeBytes field value if set, zero value otherwise.
func (o *V0ArtifactResponseItemModel) GetFileSizeBytes() int32 {
	if o == nil || IsNil(o.FileSizeBytes) {
		var ret int32
		return ret
	}
	return *o.FileSizeBytes
}

// GetFileSizeBytesOk returns a tuple with the FileSizeBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0ArtifactResponseItemModel) GetFileSizeBytesOk() (*int32, bool) {
	if o == nil || IsNil(o.FileSizeBytes) {
		return nil, false
	}
	return o.FileSizeBytes, true
}

// HasFileSizeBytes returns a boolean if a field has been set.
func (o *V0ArtifactResponseItemModel) HasFileSizeBytes() bool {
	if o != nil && !IsNil(o.FileSizeBytes) {
		return true
	}

	return false
}

// SetFileSizeBytes gets a reference to the given int32 and assigns it to the FileSizeBytes field.
func (o *V0ArtifactResponseItemModel) SetFileSizeBytes(v int32) {
	o.FileSizeBytes = &v
}

// GetIntermediateFileInfo returns the IntermediateFileInfo field value if set, zero value otherwise.
func (o *V0ArtifactResponseItemModel) GetIntermediateFileInfo() []int32 {
	if o == nil || IsNil(o.IntermediateFileInfo) {
		var ret []int32
		return ret
	}
	return o.IntermediateFileInfo
}

// GetIntermediateFileInfoOk returns a tuple with the IntermediateFileInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0ArtifactResponseItemModel) GetIntermediateFileInfoOk() ([]int32, bool) {
	if o == nil || IsNil(o.IntermediateFileInfo) {
		return nil, false
	}
	return o.IntermediateFileInfo, true
}

// HasIntermediateFileInfo returns a boolean if a field has been set.
func (o *V0ArtifactResponseItemModel) HasIntermediateFileInfo() bool {
	if o != nil && !IsNil(o.IntermediateFileInfo) {
		return true
	}

	return false
}

// SetIntermediateFileInfo gets a reference to the given []int32 and assigns it to the IntermediateFileInfo field.
func (o *V0ArtifactResponseItemModel) SetIntermediateFileInfo(v []int32) {
	o.IntermediateFileInfo = v
}

// GetIsPublicPageEnabled returns the IsPublicPageEnabled field value if set, zero value otherwise.
func (o *V0ArtifactResponseItemModel) GetIsPublicPageEnabled() bool {
	if o == nil || IsNil(o.IsPublicPageEnabled) {
		var ret bool
		return ret
	}
	return *o.IsPublicPageEnabled
}

// GetIsPublicPageEnabledOk returns a tuple with the IsPublicPageEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0ArtifactResponseItemModel) GetIsPublicPageEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPublicPageEnabled) {
		return nil, false
	}
	return o.IsPublicPageEnabled, true
}

// HasIsPublicPageEnabled returns a boolean if a field has been set.
func (o *V0ArtifactResponseItemModel) HasIsPublicPageEnabled() bool {
	if o != nil && !IsNil(o.IsPublicPageEnabled) {
		return true
	}

	return false
}

// SetIsPublicPageEnabled gets a reference to the given bool and assigns it to the IsPublicPageEnabled field.
func (o *V0ArtifactResponseItemModel) SetIsPublicPageEnabled(v bool) {
	o.IsPublicPageEnabled = &v
}

// GetPublicInstallPageUrl returns the PublicInstallPageUrl field value if set, zero value otherwise.
func (o *V0ArtifactResponseItemModel) GetPublicInstallPageUrl() string {
	if o == nil || IsNil(o.PublicInstallPageUrl) {
		var ret string
		return ret
	}
	return *o.PublicInstallPageUrl
}

// GetPublicInstallPageUrlOk returns a tuple with the PublicInstallPageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0ArtifactResponseItemModel) GetPublicInstallPageUrlOk() (*string, bool) {
	if o == nil || IsNil(o.PublicInstallPageUrl) {
		return nil, false
	}
	return o.PublicInstallPageUrl, true
}

// HasPublicInstallPageUrl returns a boolean if a field has been set.
func (o *V0ArtifactResponseItemModel) HasPublicInstallPageUrl() bool {
	if o != nil && !IsNil(o.PublicInstallPageUrl) {
		return true
	}

	return false
}

// SetPublicInstallPageUrl gets a reference to the given string and assigns it to the PublicInstallPageUrl field.
func (o *V0ArtifactResponseItemModel) SetPublicInstallPageUrl(v string) {
	o.PublicInstallPageUrl = &v
}

// GetSlug returns the Slug field value if set, zero value otherwise.
func (o *V0ArtifactResponseItemModel) GetSlug() string {
	if o == nil || IsNil(o.Slug) {
		var ret string
		return ret
	}
	return *o.Slug
}

// GetSlugOk returns a tuple with the Slug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0ArtifactResponseItemModel) GetSlugOk() (*string, bool) {
	if o == nil || IsNil(o.Slug) {
		return nil, false
	}
	return o.Slug, true
}

// HasSlug returns a boolean if a field has been set.
func (o *V0ArtifactResponseItemModel) HasSlug() bool {
	if o != nil && !IsNil(o.Slug) {
		return true
	}

	return false
}

// SetSlug gets a reference to the given string and assigns it to the Slug field.
func (o *V0ArtifactResponseItemModel) SetSlug(v string) {
	o.Slug = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *V0ArtifactResponseItemModel) GetTitle() NullsString {
	if o == nil || IsNil(o.Title) {
		var ret NullsString
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0ArtifactResponseItemModel) GetTitleOk() (*NullsString, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *V0ArtifactResponseItemModel) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullsString and assigns it to the Title field.
func (o *V0ArtifactResponseItemModel) SetTitle(v NullsString) {
	o.Title = &v
}

func (o V0ArtifactResponseItemModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0ArtifactResponseItemModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ArtifactMeta) {
		toSerialize["artifact_meta"] = o.ArtifactMeta
	}
	if !IsNil(o.ArtifactType) {
		toSerialize["artifact_type"] = o.ArtifactType
	}
	if !IsNil(o.ExpiringDownloadUrl) {
		toSerialize["expiring_download_url"] = o.ExpiringDownloadUrl
	}
	if !IsNil(o.FileSizeBytes) {
		toSerialize["file_size_bytes"] = o.FileSizeBytes
	}
	if !IsNil(o.IntermediateFileInfo) {
		toSerialize["intermediate_file_info"] = o.IntermediateFileInfo
	}
	if !IsNil(o.IsPublicPageEnabled) {
		toSerialize["is_public_page_enabled"] = o.IsPublicPageEnabled
	}
	if !IsNil(o.PublicInstallPageUrl) {
		toSerialize["public_install_page_url"] = o.PublicInstallPageUrl
	}
	if !IsNil(o.Slug) {
		toSerialize["slug"] = o.Slug
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	return toSerialize, nil
}

type NullableV0ArtifactResponseItemModel struct {
	value *V0ArtifactResponseItemModel
	isSet bool
}

func (v NullableV0ArtifactResponseItemModel) Get() *V0ArtifactResponseItemModel {
	return v.value
}

func (v *NullableV0ArtifactResponseItemModel) Set(val *V0ArtifactResponseItemModel) {
	v.value = val
	v.isSet = true
}

func (v NullableV0ArtifactResponseItemModel) IsSet() bool {
	return v.isSet
}

func (v *NullableV0ArtifactResponseItemModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0ArtifactResponseItemModel(val *V0ArtifactResponseItemModel) *NullableV0ArtifactResponseItemModel {
	return &NullableV0ArtifactResponseItemModel{value: val, isSet: true}
}

func (v NullableV0ArtifactResponseItemModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0ArtifactResponseItemModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


