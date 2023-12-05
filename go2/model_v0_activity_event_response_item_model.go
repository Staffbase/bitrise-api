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

// checks if the V0ActivityEventResponseItemModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0ActivityEventResponseItemModel{}

// V0ActivityEventResponseItemModel struct for V0ActivityEventResponseItemModel
type V0ActivityEventResponseItemModel struct {
	CreatedAt *string `json:"created_at,omitempty"`
	Description *NullsString `json:"description,omitempty"`
	EventIcon *NullsString `json:"event_icon,omitempty"`
	EventStype *NullsString `json:"event_stype,omitempty"`
	RepositoryAvatarIconUrl *string `json:"repository_avatar_icon_url,omitempty"`
	RepositoryTitle *string `json:"repository_title,omitempty"`
	Slug *string `json:"slug,omitempty"`
	TargetPathString *NullsString `json:"target_path_string,omitempty"`
	Title *NullsString `json:"title,omitempty"`
}

// NewV0ActivityEventResponseItemModel instantiates a new V0ActivityEventResponseItemModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0ActivityEventResponseItemModel() *V0ActivityEventResponseItemModel {
	this := V0ActivityEventResponseItemModel{}
	return &this
}

// NewV0ActivityEventResponseItemModelWithDefaults instantiates a new V0ActivityEventResponseItemModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0ActivityEventResponseItemModelWithDefaults() *V0ActivityEventResponseItemModel {
	this := V0ActivityEventResponseItemModel{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *V0ActivityEventResponseItemModel) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0ActivityEventResponseItemModel) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *V0ActivityEventResponseItemModel) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *V0ActivityEventResponseItemModel) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *V0ActivityEventResponseItemModel) GetDescription() NullsString {
	if o == nil || IsNil(o.Description) {
		var ret NullsString
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0ActivityEventResponseItemModel) GetDescriptionOk() (*NullsString, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *V0ActivityEventResponseItemModel) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullsString and assigns it to the Description field.
func (o *V0ActivityEventResponseItemModel) SetDescription(v NullsString) {
	o.Description = &v
}

// GetEventIcon returns the EventIcon field value if set, zero value otherwise.
func (o *V0ActivityEventResponseItemModel) GetEventIcon() NullsString {
	if o == nil || IsNil(o.EventIcon) {
		var ret NullsString
		return ret
	}
	return *o.EventIcon
}

// GetEventIconOk returns a tuple with the EventIcon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0ActivityEventResponseItemModel) GetEventIconOk() (*NullsString, bool) {
	if o == nil || IsNil(o.EventIcon) {
		return nil, false
	}
	return o.EventIcon, true
}

// HasEventIcon returns a boolean if a field has been set.
func (o *V0ActivityEventResponseItemModel) HasEventIcon() bool {
	if o != nil && !IsNil(o.EventIcon) {
		return true
	}

	return false
}

// SetEventIcon gets a reference to the given NullsString and assigns it to the EventIcon field.
func (o *V0ActivityEventResponseItemModel) SetEventIcon(v NullsString) {
	o.EventIcon = &v
}

// GetEventStype returns the EventStype field value if set, zero value otherwise.
func (o *V0ActivityEventResponseItemModel) GetEventStype() NullsString {
	if o == nil || IsNil(o.EventStype) {
		var ret NullsString
		return ret
	}
	return *o.EventStype
}

// GetEventStypeOk returns a tuple with the EventStype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0ActivityEventResponseItemModel) GetEventStypeOk() (*NullsString, bool) {
	if o == nil || IsNil(o.EventStype) {
		return nil, false
	}
	return o.EventStype, true
}

// HasEventStype returns a boolean if a field has been set.
func (o *V0ActivityEventResponseItemModel) HasEventStype() bool {
	if o != nil && !IsNil(o.EventStype) {
		return true
	}

	return false
}

// SetEventStype gets a reference to the given NullsString and assigns it to the EventStype field.
func (o *V0ActivityEventResponseItemModel) SetEventStype(v NullsString) {
	o.EventStype = &v
}

// GetRepositoryAvatarIconUrl returns the RepositoryAvatarIconUrl field value if set, zero value otherwise.
func (o *V0ActivityEventResponseItemModel) GetRepositoryAvatarIconUrl() string {
	if o == nil || IsNil(o.RepositoryAvatarIconUrl) {
		var ret string
		return ret
	}
	return *o.RepositoryAvatarIconUrl
}

// GetRepositoryAvatarIconUrlOk returns a tuple with the RepositoryAvatarIconUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0ActivityEventResponseItemModel) GetRepositoryAvatarIconUrlOk() (*string, bool) {
	if o == nil || IsNil(o.RepositoryAvatarIconUrl) {
		return nil, false
	}
	return o.RepositoryAvatarIconUrl, true
}

// HasRepositoryAvatarIconUrl returns a boolean if a field has been set.
func (o *V0ActivityEventResponseItemModel) HasRepositoryAvatarIconUrl() bool {
	if o != nil && !IsNil(o.RepositoryAvatarIconUrl) {
		return true
	}

	return false
}

// SetRepositoryAvatarIconUrl gets a reference to the given string and assigns it to the RepositoryAvatarIconUrl field.
func (o *V0ActivityEventResponseItemModel) SetRepositoryAvatarIconUrl(v string) {
	o.RepositoryAvatarIconUrl = &v
}

// GetRepositoryTitle returns the RepositoryTitle field value if set, zero value otherwise.
func (o *V0ActivityEventResponseItemModel) GetRepositoryTitle() string {
	if o == nil || IsNil(o.RepositoryTitle) {
		var ret string
		return ret
	}
	return *o.RepositoryTitle
}

// GetRepositoryTitleOk returns a tuple with the RepositoryTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0ActivityEventResponseItemModel) GetRepositoryTitleOk() (*string, bool) {
	if o == nil || IsNil(o.RepositoryTitle) {
		return nil, false
	}
	return o.RepositoryTitle, true
}

// HasRepositoryTitle returns a boolean if a field has been set.
func (o *V0ActivityEventResponseItemModel) HasRepositoryTitle() bool {
	if o != nil && !IsNil(o.RepositoryTitle) {
		return true
	}

	return false
}

// SetRepositoryTitle gets a reference to the given string and assigns it to the RepositoryTitle field.
func (o *V0ActivityEventResponseItemModel) SetRepositoryTitle(v string) {
	o.RepositoryTitle = &v
}

// GetSlug returns the Slug field value if set, zero value otherwise.
func (o *V0ActivityEventResponseItemModel) GetSlug() string {
	if o == nil || IsNil(o.Slug) {
		var ret string
		return ret
	}
	return *o.Slug
}

// GetSlugOk returns a tuple with the Slug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0ActivityEventResponseItemModel) GetSlugOk() (*string, bool) {
	if o == nil || IsNil(o.Slug) {
		return nil, false
	}
	return o.Slug, true
}

// HasSlug returns a boolean if a field has been set.
func (o *V0ActivityEventResponseItemModel) HasSlug() bool {
	if o != nil && !IsNil(o.Slug) {
		return true
	}

	return false
}

// SetSlug gets a reference to the given string and assigns it to the Slug field.
func (o *V0ActivityEventResponseItemModel) SetSlug(v string) {
	o.Slug = &v
}

// GetTargetPathString returns the TargetPathString field value if set, zero value otherwise.
func (o *V0ActivityEventResponseItemModel) GetTargetPathString() NullsString {
	if o == nil || IsNil(o.TargetPathString) {
		var ret NullsString
		return ret
	}
	return *o.TargetPathString
}

// GetTargetPathStringOk returns a tuple with the TargetPathString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0ActivityEventResponseItemModel) GetTargetPathStringOk() (*NullsString, bool) {
	if o == nil || IsNil(o.TargetPathString) {
		return nil, false
	}
	return o.TargetPathString, true
}

// HasTargetPathString returns a boolean if a field has been set.
func (o *V0ActivityEventResponseItemModel) HasTargetPathString() bool {
	if o != nil && !IsNil(o.TargetPathString) {
		return true
	}

	return false
}

// SetTargetPathString gets a reference to the given NullsString and assigns it to the TargetPathString field.
func (o *V0ActivityEventResponseItemModel) SetTargetPathString(v NullsString) {
	o.TargetPathString = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *V0ActivityEventResponseItemModel) GetTitle() NullsString {
	if o == nil || IsNil(o.Title) {
		var ret NullsString
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0ActivityEventResponseItemModel) GetTitleOk() (*NullsString, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *V0ActivityEventResponseItemModel) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullsString and assigns it to the Title field.
func (o *V0ActivityEventResponseItemModel) SetTitle(v NullsString) {
	o.Title = &v
}

func (o V0ActivityEventResponseItemModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0ActivityEventResponseItemModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.EventIcon) {
		toSerialize["event_icon"] = o.EventIcon
	}
	if !IsNil(o.EventStype) {
		toSerialize["event_stype"] = o.EventStype
	}
	if !IsNil(o.RepositoryAvatarIconUrl) {
		toSerialize["repository_avatar_icon_url"] = o.RepositoryAvatarIconUrl
	}
	if !IsNil(o.RepositoryTitle) {
		toSerialize["repository_title"] = o.RepositoryTitle
	}
	if !IsNil(o.Slug) {
		toSerialize["slug"] = o.Slug
	}
	if !IsNil(o.TargetPathString) {
		toSerialize["target_path_string"] = o.TargetPathString
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	return toSerialize, nil
}

type NullableV0ActivityEventResponseItemModel struct {
	value *V0ActivityEventResponseItemModel
	isSet bool
}

func (v NullableV0ActivityEventResponseItemModel) Get() *V0ActivityEventResponseItemModel {
	return v.value
}

func (v *NullableV0ActivityEventResponseItemModel) Set(val *V0ActivityEventResponseItemModel) {
	v.value = val
	v.isSet = true
}

func (v NullableV0ActivityEventResponseItemModel) IsSet() bool {
	return v.isSet
}

func (v *NullableV0ActivityEventResponseItemModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0ActivityEventResponseItemModel(val *V0ActivityEventResponseItemModel) *NullableV0ActivityEventResponseItemModel {
	return &NullableV0ActivityEventResponseItemModel{value: val, isSet: true}
}

func (v NullableV0ActivityEventResponseItemModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0ActivityEventResponseItemModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

