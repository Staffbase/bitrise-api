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

// checks if the AddonsAddon type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddonsAddon{}

// AddonsAddon struct for AddonsAddon
type AddonsAddon struct {
	BannerImage *string `json:"banner_image,omitempty"`
	CardHeaderColors []string `json:"card_header_colors,omitempty"`
	Categories []string `json:"categories,omitempty"`
	Description *string `json:"description,omitempty"`
	DeveloperLinks []AddonsDeveloperLink `json:"developer_links,omitempty"`
	DocumentationUrl *string `json:"documentation_url,omitempty"`
	HasUi *bool `json:"has_ui,omitempty"`
	Icon *string `json:"icon,omitempty"`
	Id *string `json:"id,omitempty"`
	IsBeta *bool `json:"is_beta,omitempty"`
	Plans []AddonsPlan `json:"plans,omitempty"`
	Platforms []string `json:"platforms,omitempty"`
	PreviewImages []string `json:"preview_images,omitempty"`
	SetupGuide *AddonsSetupGuide `json:"setup_guide,omitempty"`
	Subtitle *string `json:"subtitle,omitempty"`
	Summary *string `json:"summary,omitempty"`
	Title *string `json:"title,omitempty"`
}

// NewAddonsAddon instantiates a new AddonsAddon object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddonsAddon() *AddonsAddon {
	this := AddonsAddon{}
	return &this
}

// NewAddonsAddonWithDefaults instantiates a new AddonsAddon object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddonsAddonWithDefaults() *AddonsAddon {
	this := AddonsAddon{}
	return &this
}

// GetBannerImage returns the BannerImage field value if set, zero value otherwise.
func (o *AddonsAddon) GetBannerImage() string {
	if o == nil || IsNil(o.BannerImage) {
		var ret string
		return ret
	}
	return *o.BannerImage
}

// GetBannerImageOk returns a tuple with the BannerImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddonsAddon) GetBannerImageOk() (*string, bool) {
	if o == nil || IsNil(o.BannerImage) {
		return nil, false
	}
	return o.BannerImage, true
}

// HasBannerImage returns a boolean if a field has been set.
func (o *AddonsAddon) HasBannerImage() bool {
	if o != nil && !IsNil(o.BannerImage) {
		return true
	}

	return false
}

// SetBannerImage gets a reference to the given string and assigns it to the BannerImage field.
func (o *AddonsAddon) SetBannerImage(v string) {
	o.BannerImage = &v
}

// GetCardHeaderColors returns the CardHeaderColors field value if set, zero value otherwise.
func (o *AddonsAddon) GetCardHeaderColors() []string {
	if o == nil || IsNil(o.CardHeaderColors) {
		var ret []string
		return ret
	}
	return o.CardHeaderColors
}

// GetCardHeaderColorsOk returns a tuple with the CardHeaderColors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddonsAddon) GetCardHeaderColorsOk() ([]string, bool) {
	if o == nil || IsNil(o.CardHeaderColors) {
		return nil, false
	}
	return o.CardHeaderColors, true
}

// HasCardHeaderColors returns a boolean if a field has been set.
func (o *AddonsAddon) HasCardHeaderColors() bool {
	if o != nil && !IsNil(o.CardHeaderColors) {
		return true
	}

	return false
}

// SetCardHeaderColors gets a reference to the given []string and assigns it to the CardHeaderColors field.
func (o *AddonsAddon) SetCardHeaderColors(v []string) {
	o.CardHeaderColors = v
}

// GetCategories returns the Categories field value if set, zero value otherwise.
func (o *AddonsAddon) GetCategories() []string {
	if o == nil || IsNil(o.Categories) {
		var ret []string
		return ret
	}
	return o.Categories
}

// GetCategoriesOk returns a tuple with the Categories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddonsAddon) GetCategoriesOk() ([]string, bool) {
	if o == nil || IsNil(o.Categories) {
		return nil, false
	}
	return o.Categories, true
}

// HasCategories returns a boolean if a field has been set.
func (o *AddonsAddon) HasCategories() bool {
	if o != nil && !IsNil(o.Categories) {
		return true
	}

	return false
}

// SetCategories gets a reference to the given []string and assigns it to the Categories field.
func (o *AddonsAddon) SetCategories(v []string) {
	o.Categories = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddonsAddon) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddonsAddon) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddonsAddon) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddonsAddon) SetDescription(v string) {
	o.Description = &v
}

// GetDeveloperLinks returns the DeveloperLinks field value if set, zero value otherwise.
func (o *AddonsAddon) GetDeveloperLinks() []AddonsDeveloperLink {
	if o == nil || IsNil(o.DeveloperLinks) {
		var ret []AddonsDeveloperLink
		return ret
	}
	return o.DeveloperLinks
}

// GetDeveloperLinksOk returns a tuple with the DeveloperLinks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddonsAddon) GetDeveloperLinksOk() ([]AddonsDeveloperLink, bool) {
	if o == nil || IsNil(o.DeveloperLinks) {
		return nil, false
	}
	return o.DeveloperLinks, true
}

// HasDeveloperLinks returns a boolean if a field has been set.
func (o *AddonsAddon) HasDeveloperLinks() bool {
	if o != nil && !IsNil(o.DeveloperLinks) {
		return true
	}

	return false
}

// SetDeveloperLinks gets a reference to the given []AddonsDeveloperLink and assigns it to the DeveloperLinks field.
func (o *AddonsAddon) SetDeveloperLinks(v []AddonsDeveloperLink) {
	o.DeveloperLinks = v
}

// GetDocumentationUrl returns the DocumentationUrl field value if set, zero value otherwise.
func (o *AddonsAddon) GetDocumentationUrl() string {
	if o == nil || IsNil(o.DocumentationUrl) {
		var ret string
		return ret
	}
	return *o.DocumentationUrl
}

// GetDocumentationUrlOk returns a tuple with the DocumentationUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddonsAddon) GetDocumentationUrlOk() (*string, bool) {
	if o == nil || IsNil(o.DocumentationUrl) {
		return nil, false
	}
	return o.DocumentationUrl, true
}

// HasDocumentationUrl returns a boolean if a field has been set.
func (o *AddonsAddon) HasDocumentationUrl() bool {
	if o != nil && !IsNil(o.DocumentationUrl) {
		return true
	}

	return false
}

// SetDocumentationUrl gets a reference to the given string and assigns it to the DocumentationUrl field.
func (o *AddonsAddon) SetDocumentationUrl(v string) {
	o.DocumentationUrl = &v
}

// GetHasUi returns the HasUi field value if set, zero value otherwise.
func (o *AddonsAddon) GetHasUi() bool {
	if o == nil || IsNil(o.HasUi) {
		var ret bool
		return ret
	}
	return *o.HasUi
}

// GetHasUiOk returns a tuple with the HasUi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddonsAddon) GetHasUiOk() (*bool, bool) {
	if o == nil || IsNil(o.HasUi) {
		return nil, false
	}
	return o.HasUi, true
}

// HasHasUi returns a boolean if a field has been set.
func (o *AddonsAddon) HasHasUi() bool {
	if o != nil && !IsNil(o.HasUi) {
		return true
	}

	return false
}

// SetHasUi gets a reference to the given bool and assigns it to the HasUi field.
func (o *AddonsAddon) SetHasUi(v bool) {
	o.HasUi = &v
}

// GetIcon returns the Icon field value if set, zero value otherwise.
func (o *AddonsAddon) GetIcon() string {
	if o == nil || IsNil(o.Icon) {
		var ret string
		return ret
	}
	return *o.Icon
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddonsAddon) GetIconOk() (*string, bool) {
	if o == nil || IsNil(o.Icon) {
		return nil, false
	}
	return o.Icon, true
}

// HasIcon returns a boolean if a field has been set.
func (o *AddonsAddon) HasIcon() bool {
	if o != nil && !IsNil(o.Icon) {
		return true
	}

	return false
}

// SetIcon gets a reference to the given string and assigns it to the Icon field.
func (o *AddonsAddon) SetIcon(v string) {
	o.Icon = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AddonsAddon) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddonsAddon) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AddonsAddon) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AddonsAddon) SetId(v string) {
	o.Id = &v
}

// GetIsBeta returns the IsBeta field value if set, zero value otherwise.
func (o *AddonsAddon) GetIsBeta() bool {
	if o == nil || IsNil(o.IsBeta) {
		var ret bool
		return ret
	}
	return *o.IsBeta
}

// GetIsBetaOk returns a tuple with the IsBeta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddonsAddon) GetIsBetaOk() (*bool, bool) {
	if o == nil || IsNil(o.IsBeta) {
		return nil, false
	}
	return o.IsBeta, true
}

// HasIsBeta returns a boolean if a field has been set.
func (o *AddonsAddon) HasIsBeta() bool {
	if o != nil && !IsNil(o.IsBeta) {
		return true
	}

	return false
}

// SetIsBeta gets a reference to the given bool and assigns it to the IsBeta field.
func (o *AddonsAddon) SetIsBeta(v bool) {
	o.IsBeta = &v
}

// GetPlans returns the Plans field value if set, zero value otherwise.
func (o *AddonsAddon) GetPlans() []AddonsPlan {
	if o == nil || IsNil(o.Plans) {
		var ret []AddonsPlan
		return ret
	}
	return o.Plans
}

// GetPlansOk returns a tuple with the Plans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddonsAddon) GetPlansOk() ([]AddonsPlan, bool) {
	if o == nil || IsNil(o.Plans) {
		return nil, false
	}
	return o.Plans, true
}

// HasPlans returns a boolean if a field has been set.
func (o *AddonsAddon) HasPlans() bool {
	if o != nil && !IsNil(o.Plans) {
		return true
	}

	return false
}

// SetPlans gets a reference to the given []AddonsPlan and assigns it to the Plans field.
func (o *AddonsAddon) SetPlans(v []AddonsPlan) {
	o.Plans = v
}

// GetPlatforms returns the Platforms field value if set, zero value otherwise.
func (o *AddonsAddon) GetPlatforms() []string {
	if o == nil || IsNil(o.Platforms) {
		var ret []string
		return ret
	}
	return o.Platforms
}

// GetPlatformsOk returns a tuple with the Platforms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddonsAddon) GetPlatformsOk() ([]string, bool) {
	if o == nil || IsNil(o.Platforms) {
		return nil, false
	}
	return o.Platforms, true
}

// HasPlatforms returns a boolean if a field has been set.
func (o *AddonsAddon) HasPlatforms() bool {
	if o != nil && !IsNil(o.Platforms) {
		return true
	}

	return false
}

// SetPlatforms gets a reference to the given []string and assigns it to the Platforms field.
func (o *AddonsAddon) SetPlatforms(v []string) {
	o.Platforms = v
}

// GetPreviewImages returns the PreviewImages field value if set, zero value otherwise.
func (o *AddonsAddon) GetPreviewImages() []string {
	if o == nil || IsNil(o.PreviewImages) {
		var ret []string
		return ret
	}
	return o.PreviewImages
}

// GetPreviewImagesOk returns a tuple with the PreviewImages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddonsAddon) GetPreviewImagesOk() ([]string, bool) {
	if o == nil || IsNil(o.PreviewImages) {
		return nil, false
	}
	return o.PreviewImages, true
}

// HasPreviewImages returns a boolean if a field has been set.
func (o *AddonsAddon) HasPreviewImages() bool {
	if o != nil && !IsNil(o.PreviewImages) {
		return true
	}

	return false
}

// SetPreviewImages gets a reference to the given []string and assigns it to the PreviewImages field.
func (o *AddonsAddon) SetPreviewImages(v []string) {
	o.PreviewImages = v
}

// GetSetupGuide returns the SetupGuide field value if set, zero value otherwise.
func (o *AddonsAddon) GetSetupGuide() AddonsSetupGuide {
	if o == nil || IsNil(o.SetupGuide) {
		var ret AddonsSetupGuide
		return ret
	}
	return *o.SetupGuide
}

// GetSetupGuideOk returns a tuple with the SetupGuide field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddonsAddon) GetSetupGuideOk() (*AddonsSetupGuide, bool) {
	if o == nil || IsNil(o.SetupGuide) {
		return nil, false
	}
	return o.SetupGuide, true
}

// HasSetupGuide returns a boolean if a field has been set.
func (o *AddonsAddon) HasSetupGuide() bool {
	if o != nil && !IsNil(o.SetupGuide) {
		return true
	}

	return false
}

// SetSetupGuide gets a reference to the given AddonsSetupGuide and assigns it to the SetupGuide field.
func (o *AddonsAddon) SetSetupGuide(v AddonsSetupGuide) {
	o.SetupGuide = &v
}

// GetSubtitle returns the Subtitle field value if set, zero value otherwise.
func (o *AddonsAddon) GetSubtitle() string {
	if o == nil || IsNil(o.Subtitle) {
		var ret string
		return ret
	}
	return *o.Subtitle
}

// GetSubtitleOk returns a tuple with the Subtitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddonsAddon) GetSubtitleOk() (*string, bool) {
	if o == nil || IsNil(o.Subtitle) {
		return nil, false
	}
	return o.Subtitle, true
}

// HasSubtitle returns a boolean if a field has been set.
func (o *AddonsAddon) HasSubtitle() bool {
	if o != nil && !IsNil(o.Subtitle) {
		return true
	}

	return false
}

// SetSubtitle gets a reference to the given string and assigns it to the Subtitle field.
func (o *AddonsAddon) SetSubtitle(v string) {
	o.Subtitle = &v
}

// GetSummary returns the Summary field value if set, zero value otherwise.
func (o *AddonsAddon) GetSummary() string {
	if o == nil || IsNil(o.Summary) {
		var ret string
		return ret
	}
	return *o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddonsAddon) GetSummaryOk() (*string, bool) {
	if o == nil || IsNil(o.Summary) {
		return nil, false
	}
	return o.Summary, true
}

// HasSummary returns a boolean if a field has been set.
func (o *AddonsAddon) HasSummary() bool {
	if o != nil && !IsNil(o.Summary) {
		return true
	}

	return false
}

// SetSummary gets a reference to the given string and assigns it to the Summary field.
func (o *AddonsAddon) SetSummary(v string) {
	o.Summary = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *AddonsAddon) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddonsAddon) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *AddonsAddon) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *AddonsAddon) SetTitle(v string) {
	o.Title = &v
}

func (o AddonsAddon) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddonsAddon) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BannerImage) {
		toSerialize["banner_image"] = o.BannerImage
	}
	if !IsNil(o.CardHeaderColors) {
		toSerialize["card_header_colors"] = o.CardHeaderColors
	}
	if !IsNil(o.Categories) {
		toSerialize["categories"] = o.Categories
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.DeveloperLinks) {
		toSerialize["developer_links"] = o.DeveloperLinks
	}
	if !IsNil(o.DocumentationUrl) {
		toSerialize["documentation_url"] = o.DocumentationUrl
	}
	if !IsNil(o.HasUi) {
		toSerialize["has_ui"] = o.HasUi
	}
	if !IsNil(o.Icon) {
		toSerialize["icon"] = o.Icon
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.IsBeta) {
		toSerialize["is_beta"] = o.IsBeta
	}
	if !IsNil(o.Plans) {
		toSerialize["plans"] = o.Plans
	}
	if !IsNil(o.Platforms) {
		toSerialize["platforms"] = o.Platforms
	}
	if !IsNil(o.PreviewImages) {
		toSerialize["preview_images"] = o.PreviewImages
	}
	if !IsNil(o.SetupGuide) {
		toSerialize["setup_guide"] = o.SetupGuide
	}
	if !IsNil(o.Subtitle) {
		toSerialize["subtitle"] = o.Subtitle
	}
	if !IsNil(o.Summary) {
		toSerialize["summary"] = o.Summary
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	return toSerialize, nil
}

type NullableAddonsAddon struct {
	value *AddonsAddon
	isSet bool
}

func (v NullableAddonsAddon) Get() *AddonsAddon {
	return v.value
}

func (v *NullableAddonsAddon) Set(val *AddonsAddon) {
	v.value = val
	v.isSet = true
}

func (v NullableAddonsAddon) IsSet() bool {
	return v.isSet
}

func (v *NullableAddonsAddon) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddonsAddon(val *AddonsAddon) *NullableAddonsAddon {
	return &NullableAddonsAddon{value: val, isSet: true}
}

func (v NullableAddonsAddon) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddonsAddon) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

