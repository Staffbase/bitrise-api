# AddonsAddon

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BannerImage** | Pointer to **string** |  | [optional] 
**CardHeaderColors** | Pointer to **[]string** |  | [optional] 
**Categories** | Pointer to **[]string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DeveloperLinks** | Pointer to [**[]AddonsDeveloperLink**](AddonsDeveloperLink.md) |  | [optional] 
**DocumentationUrl** | Pointer to **string** |  | [optional] 
**HasUi** | Pointer to **bool** |  | [optional] 
**Icon** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**IsBeta** | Pointer to **bool** |  | [optional] 
**Plans** | Pointer to [**[]AddonsPlan**](AddonsPlan.md) |  | [optional] 
**Platforms** | Pointer to **[]string** |  | [optional] 
**PreviewImages** | Pointer to **[]string** |  | [optional] 
**SetupGuide** | Pointer to [**AddonsSetupGuide**](AddonsSetupGuide.md) |  | [optional] 
**Subtitle** | Pointer to **string** |  | [optional] 
**Summary** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 

## Methods

### NewAddonsAddon

`func NewAddonsAddon() *AddonsAddon`

NewAddonsAddon instantiates a new AddonsAddon object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddonsAddonWithDefaults

`func NewAddonsAddonWithDefaults() *AddonsAddon`

NewAddonsAddonWithDefaults instantiates a new AddonsAddon object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBannerImage

`func (o *AddonsAddon) GetBannerImage() string`

GetBannerImage returns the BannerImage field if non-nil, zero value otherwise.

### GetBannerImageOk

`func (o *AddonsAddon) GetBannerImageOk() (*string, bool)`

GetBannerImageOk returns a tuple with the BannerImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBannerImage

`func (o *AddonsAddon) SetBannerImage(v string)`

SetBannerImage sets BannerImage field to given value.

### HasBannerImage

`func (o *AddonsAddon) HasBannerImage() bool`

HasBannerImage returns a boolean if a field has been set.

### GetCardHeaderColors

`func (o *AddonsAddon) GetCardHeaderColors() []string`

GetCardHeaderColors returns the CardHeaderColors field if non-nil, zero value otherwise.

### GetCardHeaderColorsOk

`func (o *AddonsAddon) GetCardHeaderColorsOk() (*[]string, bool)`

GetCardHeaderColorsOk returns a tuple with the CardHeaderColors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardHeaderColors

`func (o *AddonsAddon) SetCardHeaderColors(v []string)`

SetCardHeaderColors sets CardHeaderColors field to given value.

### HasCardHeaderColors

`func (o *AddonsAddon) HasCardHeaderColors() bool`

HasCardHeaderColors returns a boolean if a field has been set.

### GetCategories

`func (o *AddonsAddon) GetCategories() []string`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *AddonsAddon) GetCategoriesOk() (*[]string, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *AddonsAddon) SetCategories(v []string)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *AddonsAddon) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetDescription

`func (o *AddonsAddon) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddonsAddon) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddonsAddon) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddonsAddon) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDeveloperLinks

`func (o *AddonsAddon) GetDeveloperLinks() []AddonsDeveloperLink`

GetDeveloperLinks returns the DeveloperLinks field if non-nil, zero value otherwise.

### GetDeveloperLinksOk

`func (o *AddonsAddon) GetDeveloperLinksOk() (*[]AddonsDeveloperLink, bool)`

GetDeveloperLinksOk returns a tuple with the DeveloperLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeveloperLinks

`func (o *AddonsAddon) SetDeveloperLinks(v []AddonsDeveloperLink)`

SetDeveloperLinks sets DeveloperLinks field to given value.

### HasDeveloperLinks

`func (o *AddonsAddon) HasDeveloperLinks() bool`

HasDeveloperLinks returns a boolean if a field has been set.

### GetDocumentationUrl

`func (o *AddonsAddon) GetDocumentationUrl() string`

GetDocumentationUrl returns the DocumentationUrl field if non-nil, zero value otherwise.

### GetDocumentationUrlOk

`func (o *AddonsAddon) GetDocumentationUrlOk() (*string, bool)`

GetDocumentationUrlOk returns a tuple with the DocumentationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentationUrl

`func (o *AddonsAddon) SetDocumentationUrl(v string)`

SetDocumentationUrl sets DocumentationUrl field to given value.

### HasDocumentationUrl

`func (o *AddonsAddon) HasDocumentationUrl() bool`

HasDocumentationUrl returns a boolean if a field has been set.

### GetHasUi

`func (o *AddonsAddon) GetHasUi() bool`

GetHasUi returns the HasUi field if non-nil, zero value otherwise.

### GetHasUiOk

`func (o *AddonsAddon) GetHasUiOk() (*bool, bool)`

GetHasUiOk returns a tuple with the HasUi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasUi

`func (o *AddonsAddon) SetHasUi(v bool)`

SetHasUi sets HasUi field to given value.

### HasHasUi

`func (o *AddonsAddon) HasHasUi() bool`

HasHasUi returns a boolean if a field has been set.

### GetIcon

`func (o *AddonsAddon) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *AddonsAddon) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *AddonsAddon) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *AddonsAddon) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetId

`func (o *AddonsAddon) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddonsAddon) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddonsAddon) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AddonsAddon) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsBeta

`func (o *AddonsAddon) GetIsBeta() bool`

GetIsBeta returns the IsBeta field if non-nil, zero value otherwise.

### GetIsBetaOk

`func (o *AddonsAddon) GetIsBetaOk() (*bool, bool)`

GetIsBetaOk returns a tuple with the IsBeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBeta

`func (o *AddonsAddon) SetIsBeta(v bool)`

SetIsBeta sets IsBeta field to given value.

### HasIsBeta

`func (o *AddonsAddon) HasIsBeta() bool`

HasIsBeta returns a boolean if a field has been set.

### GetPlans

`func (o *AddonsAddon) GetPlans() []AddonsPlan`

GetPlans returns the Plans field if non-nil, zero value otherwise.

### GetPlansOk

`func (o *AddonsAddon) GetPlansOk() (*[]AddonsPlan, bool)`

GetPlansOk returns a tuple with the Plans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlans

`func (o *AddonsAddon) SetPlans(v []AddonsPlan)`

SetPlans sets Plans field to given value.

### HasPlans

`func (o *AddonsAddon) HasPlans() bool`

HasPlans returns a boolean if a field has been set.

### GetPlatforms

`func (o *AddonsAddon) GetPlatforms() []string`

GetPlatforms returns the Platforms field if non-nil, zero value otherwise.

### GetPlatformsOk

`func (o *AddonsAddon) GetPlatformsOk() (*[]string, bool)`

GetPlatformsOk returns a tuple with the Platforms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatforms

`func (o *AddonsAddon) SetPlatforms(v []string)`

SetPlatforms sets Platforms field to given value.

### HasPlatforms

`func (o *AddonsAddon) HasPlatforms() bool`

HasPlatforms returns a boolean if a field has been set.

### GetPreviewImages

`func (o *AddonsAddon) GetPreviewImages() []string`

GetPreviewImages returns the PreviewImages field if non-nil, zero value otherwise.

### GetPreviewImagesOk

`func (o *AddonsAddon) GetPreviewImagesOk() (*[]string, bool)`

GetPreviewImagesOk returns a tuple with the PreviewImages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviewImages

`func (o *AddonsAddon) SetPreviewImages(v []string)`

SetPreviewImages sets PreviewImages field to given value.

### HasPreviewImages

`func (o *AddonsAddon) HasPreviewImages() bool`

HasPreviewImages returns a boolean if a field has been set.

### GetSetupGuide

`func (o *AddonsAddon) GetSetupGuide() AddonsSetupGuide`

GetSetupGuide returns the SetupGuide field if non-nil, zero value otherwise.

### GetSetupGuideOk

`func (o *AddonsAddon) GetSetupGuideOk() (*AddonsSetupGuide, bool)`

GetSetupGuideOk returns a tuple with the SetupGuide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetupGuide

`func (o *AddonsAddon) SetSetupGuide(v AddonsSetupGuide)`

SetSetupGuide sets SetupGuide field to given value.

### HasSetupGuide

`func (o *AddonsAddon) HasSetupGuide() bool`

HasSetupGuide returns a boolean if a field has been set.

### GetSubtitle

`func (o *AddonsAddon) GetSubtitle() string`

GetSubtitle returns the Subtitle field if non-nil, zero value otherwise.

### GetSubtitleOk

`func (o *AddonsAddon) GetSubtitleOk() (*string, bool)`

GetSubtitleOk returns a tuple with the Subtitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtitle

`func (o *AddonsAddon) SetSubtitle(v string)`

SetSubtitle sets Subtitle field to given value.

### HasSubtitle

`func (o *AddonsAddon) HasSubtitle() bool`

HasSubtitle returns a boolean if a field has been set.

### GetSummary

`func (o *AddonsAddon) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *AddonsAddon) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *AddonsAddon) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *AddonsAddon) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetTitle

`func (o *AddonsAddon) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AddonsAddon) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AddonsAddon) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AddonsAddon) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


