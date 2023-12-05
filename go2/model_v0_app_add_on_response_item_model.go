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

// checks if the V0AppAddOnResponseItemModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0AppAddOnResponseItemModel{}

// V0AppAddOnResponseItemModel struct for V0AppAddOnResponseItemModel
type V0AppAddOnResponseItemModel struct {
	Description *string `json:"description,omitempty"`
	DocumentationUrl *string `json:"documentation_url,omitempty"`
	HasUi *bool `json:"has_ui,omitempty"`
	Icon *string `json:"icon,omitempty"`
	Id *string `json:"id,omitempty"`
	IsBeta *bool `json:"is_beta,omitempty"`
	IsEnabled *bool `json:"is_enabled,omitempty"`
	LoginUrl *string `json:"login_url,omitempty"`
	Plan *AddonsPlan `json:"plan,omitempty"`
	Scopes []string `json:"scopes,omitempty"`
	SetupGuide *AddonsSetupGuide `json:"setup_guide,omitempty"`
	Summary *string `json:"summary,omitempty"`
	TermsUrl *string `json:"terms_url,omitempty"`
	Title *string `json:"title,omitempty"`
}

// NewV0AppAddOnResponseItemModel instantiates a new V0AppAddOnResponseItemModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0AppAddOnResponseItemModel() *V0AppAddOnResponseItemModel {
	this := V0AppAddOnResponseItemModel{}
	return &this
}

// NewV0AppAddOnResponseItemModelWithDefaults instantiates a new V0AppAddOnResponseItemModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0AppAddOnResponseItemModelWithDefaults() *V0AppAddOnResponseItemModel {
	this := V0AppAddOnResponseItemModel{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *V0AppAddOnResponseItemModel) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0AppAddOnResponseItemModel) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *V0AppAddOnResponseItemModel) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *V0AppAddOnResponseItemModel) SetDescription(v string) {
	o.Description = &v
}

// GetDocumentationUrl returns the DocumentationUrl field value if set, zero value otherwise.
func (o *V0AppAddOnResponseItemModel) GetDocumentationUrl() string {
	if o == nil || IsNil(o.DocumentationUrl) {
		var ret string
		return ret
	}
	return *o.DocumentationUrl
}

// GetDocumentationUrlOk returns a tuple with the DocumentationUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0AppAddOnResponseItemModel) GetDocumentationUrlOk() (*string, bool) {
	if o == nil || IsNil(o.DocumentationUrl) {
		return nil, false
	}
	return o.DocumentationUrl, true
}

// HasDocumentationUrl returns a boolean if a field has been set.
func (o *V0AppAddOnResponseItemModel) HasDocumentationUrl() bool {
	if o != nil && !IsNil(o.DocumentationUrl) {
		return true
	}

	return false
}

// SetDocumentationUrl gets a reference to the given string and assigns it to the DocumentationUrl field.
func (o *V0AppAddOnResponseItemModel) SetDocumentationUrl(v string) {
	o.DocumentationUrl = &v
}

// GetHasUi returns the HasUi field value if set, zero value otherwise.
func (o *V0AppAddOnResponseItemModel) GetHasUi() bool {
	if o == nil || IsNil(o.HasUi) {
		var ret bool
		return ret
	}
	return *o.HasUi
}

// GetHasUiOk returns a tuple with the HasUi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0AppAddOnResponseItemModel) GetHasUiOk() (*bool, bool) {
	if o == nil || IsNil(o.HasUi) {
		return nil, false
	}
	return o.HasUi, true
}

// HasHasUi returns a boolean if a field has been set.
func (o *V0AppAddOnResponseItemModel) HasHasUi() bool {
	if o != nil && !IsNil(o.HasUi) {
		return true
	}

	return false
}

// SetHasUi gets a reference to the given bool and assigns it to the HasUi field.
func (o *V0AppAddOnResponseItemModel) SetHasUi(v bool) {
	o.HasUi = &v
}

// GetIcon returns the Icon field value if set, zero value otherwise.
func (o *V0AppAddOnResponseItemModel) GetIcon() string {
	if o == nil || IsNil(o.Icon) {
		var ret string
		return ret
	}
	return *o.Icon
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0AppAddOnResponseItemModel) GetIconOk() (*string, bool) {
	if o == nil || IsNil(o.Icon) {
		return nil, false
	}
	return o.Icon, true
}

// HasIcon returns a boolean if a field has been set.
func (o *V0AppAddOnResponseItemModel) HasIcon() bool {
	if o != nil && !IsNil(o.Icon) {
		return true
	}

	return false
}

// SetIcon gets a reference to the given string and assigns it to the Icon field.
func (o *V0AppAddOnResponseItemModel) SetIcon(v string) {
	o.Icon = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *V0AppAddOnResponseItemModel) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0AppAddOnResponseItemModel) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *V0AppAddOnResponseItemModel) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *V0AppAddOnResponseItemModel) SetId(v string) {
	o.Id = &v
}

// GetIsBeta returns the IsBeta field value if set, zero value otherwise.
func (o *V0AppAddOnResponseItemModel) GetIsBeta() bool {
	if o == nil || IsNil(o.IsBeta) {
		var ret bool
		return ret
	}
	return *o.IsBeta
}

// GetIsBetaOk returns a tuple with the IsBeta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0AppAddOnResponseItemModel) GetIsBetaOk() (*bool, bool) {
	if o == nil || IsNil(o.IsBeta) {
		return nil, false
	}
	return o.IsBeta, true
}

// HasIsBeta returns a boolean if a field has been set.
func (o *V0AppAddOnResponseItemModel) HasIsBeta() bool {
	if o != nil && !IsNil(o.IsBeta) {
		return true
	}

	return false
}

// SetIsBeta gets a reference to the given bool and assigns it to the IsBeta field.
func (o *V0AppAddOnResponseItemModel) SetIsBeta(v bool) {
	o.IsBeta = &v
}

// GetIsEnabled returns the IsEnabled field value if set, zero value otherwise.
func (o *V0AppAddOnResponseItemModel) GetIsEnabled() bool {
	if o == nil || IsNil(o.IsEnabled) {
		var ret bool
		return ret
	}
	return *o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0AppAddOnResponseItemModel) GetIsEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsEnabled) {
		return nil, false
	}
	return o.IsEnabled, true
}

// HasIsEnabled returns a boolean if a field has been set.
func (o *V0AppAddOnResponseItemModel) HasIsEnabled() bool {
	if o != nil && !IsNil(o.IsEnabled) {
		return true
	}

	return false
}

// SetIsEnabled gets a reference to the given bool and assigns it to the IsEnabled field.
func (o *V0AppAddOnResponseItemModel) SetIsEnabled(v bool) {
	o.IsEnabled = &v
}

// GetLoginUrl returns the LoginUrl field value if set, zero value otherwise.
func (o *V0AppAddOnResponseItemModel) GetLoginUrl() string {
	if o == nil || IsNil(o.LoginUrl) {
		var ret string
		return ret
	}
	return *o.LoginUrl
}

// GetLoginUrlOk returns a tuple with the LoginUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0AppAddOnResponseItemModel) GetLoginUrlOk() (*string, bool) {
	if o == nil || IsNil(o.LoginUrl) {
		return nil, false
	}
	return o.LoginUrl, true
}

// HasLoginUrl returns a boolean if a field has been set.
func (o *V0AppAddOnResponseItemModel) HasLoginUrl() bool {
	if o != nil && !IsNil(o.LoginUrl) {
		return true
	}

	return false
}

// SetLoginUrl gets a reference to the given string and assigns it to the LoginUrl field.
func (o *V0AppAddOnResponseItemModel) SetLoginUrl(v string) {
	o.LoginUrl = &v
}

// GetPlan returns the Plan field value if set, zero value otherwise.
func (o *V0AppAddOnResponseItemModel) GetPlan() AddonsPlan {
	if o == nil || IsNil(o.Plan) {
		var ret AddonsPlan
		return ret
	}
	return *o.Plan
}

// GetPlanOk returns a tuple with the Plan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0AppAddOnResponseItemModel) GetPlanOk() (*AddonsPlan, bool) {
	if o == nil || IsNil(o.Plan) {
		return nil, false
	}
	return o.Plan, true
}

// HasPlan returns a boolean if a field has been set.
func (o *V0AppAddOnResponseItemModel) HasPlan() bool {
	if o != nil && !IsNil(o.Plan) {
		return true
	}

	return false
}

// SetPlan gets a reference to the given AddonsPlan and assigns it to the Plan field.
func (o *V0AppAddOnResponseItemModel) SetPlan(v AddonsPlan) {
	o.Plan = &v
}

// GetScopes returns the Scopes field value if set, zero value otherwise.
func (o *V0AppAddOnResponseItemModel) GetScopes() []string {
	if o == nil || IsNil(o.Scopes) {
		var ret []string
		return ret
	}
	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0AppAddOnResponseItemModel) GetScopesOk() ([]string, bool) {
	if o == nil || IsNil(o.Scopes) {
		return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *V0AppAddOnResponseItemModel) HasScopes() bool {
	if o != nil && !IsNil(o.Scopes) {
		return true
	}

	return false
}

// SetScopes gets a reference to the given []string and assigns it to the Scopes field.
func (o *V0AppAddOnResponseItemModel) SetScopes(v []string) {
	o.Scopes = v
}

// GetSetupGuide returns the SetupGuide field value if set, zero value otherwise.
func (o *V0AppAddOnResponseItemModel) GetSetupGuide() AddonsSetupGuide {
	if o == nil || IsNil(o.SetupGuide) {
		var ret AddonsSetupGuide
		return ret
	}
	return *o.SetupGuide
}

// GetSetupGuideOk returns a tuple with the SetupGuide field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0AppAddOnResponseItemModel) GetSetupGuideOk() (*AddonsSetupGuide, bool) {
	if o == nil || IsNil(o.SetupGuide) {
		return nil, false
	}
	return o.SetupGuide, true
}

// HasSetupGuide returns a boolean if a field has been set.
func (o *V0AppAddOnResponseItemModel) HasSetupGuide() bool {
	if o != nil && !IsNil(o.SetupGuide) {
		return true
	}

	return false
}

// SetSetupGuide gets a reference to the given AddonsSetupGuide and assigns it to the SetupGuide field.
func (o *V0AppAddOnResponseItemModel) SetSetupGuide(v AddonsSetupGuide) {
	o.SetupGuide = &v
}

// GetSummary returns the Summary field value if set, zero value otherwise.
func (o *V0AppAddOnResponseItemModel) GetSummary() string {
	if o == nil || IsNil(o.Summary) {
		var ret string
		return ret
	}
	return *o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0AppAddOnResponseItemModel) GetSummaryOk() (*string, bool) {
	if o == nil || IsNil(o.Summary) {
		return nil, false
	}
	return o.Summary, true
}

// HasSummary returns a boolean if a field has been set.
func (o *V0AppAddOnResponseItemModel) HasSummary() bool {
	if o != nil && !IsNil(o.Summary) {
		return true
	}

	return false
}

// SetSummary gets a reference to the given string and assigns it to the Summary field.
func (o *V0AppAddOnResponseItemModel) SetSummary(v string) {
	o.Summary = &v
}

// GetTermsUrl returns the TermsUrl field value if set, zero value otherwise.
func (o *V0AppAddOnResponseItemModel) GetTermsUrl() string {
	if o == nil || IsNil(o.TermsUrl) {
		var ret string
		return ret
	}
	return *o.TermsUrl
}

// GetTermsUrlOk returns a tuple with the TermsUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0AppAddOnResponseItemModel) GetTermsUrlOk() (*string, bool) {
	if o == nil || IsNil(o.TermsUrl) {
		return nil, false
	}
	return o.TermsUrl, true
}

// HasTermsUrl returns a boolean if a field has been set.
func (o *V0AppAddOnResponseItemModel) HasTermsUrl() bool {
	if o != nil && !IsNil(o.TermsUrl) {
		return true
	}

	return false
}

// SetTermsUrl gets a reference to the given string and assigns it to the TermsUrl field.
func (o *V0AppAddOnResponseItemModel) SetTermsUrl(v string) {
	o.TermsUrl = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *V0AppAddOnResponseItemModel) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0AppAddOnResponseItemModel) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *V0AppAddOnResponseItemModel) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *V0AppAddOnResponseItemModel) SetTitle(v string) {
	o.Title = &v
}

func (o V0AppAddOnResponseItemModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0AppAddOnResponseItemModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
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
	if !IsNil(o.IsEnabled) {
		toSerialize["is_enabled"] = o.IsEnabled
	}
	if !IsNil(o.LoginUrl) {
		toSerialize["login_url"] = o.LoginUrl
	}
	if !IsNil(o.Plan) {
		toSerialize["plan"] = o.Plan
	}
	if !IsNil(o.Scopes) {
		toSerialize["scopes"] = o.Scopes
	}
	if !IsNil(o.SetupGuide) {
		toSerialize["setup_guide"] = o.SetupGuide
	}
	if !IsNil(o.Summary) {
		toSerialize["summary"] = o.Summary
	}
	if !IsNil(o.TermsUrl) {
		toSerialize["terms_url"] = o.TermsUrl
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	return toSerialize, nil
}

type NullableV0AppAddOnResponseItemModel struct {
	value *V0AppAddOnResponseItemModel
	isSet bool
}

func (v NullableV0AppAddOnResponseItemModel) Get() *V0AppAddOnResponseItemModel {
	return v.value
}

func (v *NullableV0AppAddOnResponseItemModel) Set(val *V0AppAddOnResponseItemModel) {
	v.value = val
	v.isSet = true
}

func (v NullableV0AppAddOnResponseItemModel) IsSet() bool {
	return v.isSet
}

func (v *NullableV0AppAddOnResponseItemModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0AppAddOnResponseItemModel(val *V0AppAddOnResponseItemModel) *NullableV0AppAddOnResponseItemModel {
	return &NullableV0AppAddOnResponseItemModel{value: val, isSet: true}
}

func (v NullableV0AppAddOnResponseItemModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0AppAddOnResponseItemModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


