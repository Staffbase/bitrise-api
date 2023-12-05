# V0ArtifactResponseItemModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArtifactMeta** | Pointer to **[]int32** |  | [optional] 
**ArtifactType** | Pointer to [**NullsString**](NullsString.md) |  | [optional] 
**ExpiringDownloadUrl** | Pointer to **string** |  | [optional] 
**FileSizeBytes** | Pointer to **int32** |  | [optional] 
**IntermediateFileInfo** | Pointer to **[]int32** |  | [optional] 
**IsPublicPageEnabled** | Pointer to **bool** |  | [optional] 
**PublicInstallPageUrl** | Pointer to **string** |  | [optional] 
**Slug** | Pointer to **string** |  | [optional] 
**Title** | Pointer to [**NullsString**](NullsString.md) |  | [optional] 

## Methods

### NewV0ArtifactResponseItemModel

`func NewV0ArtifactResponseItemModel() *V0ArtifactResponseItemModel`

NewV0ArtifactResponseItemModel instantiates a new V0ArtifactResponseItemModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0ArtifactResponseItemModelWithDefaults

`func NewV0ArtifactResponseItemModelWithDefaults() *V0ArtifactResponseItemModel`

NewV0ArtifactResponseItemModelWithDefaults instantiates a new V0ArtifactResponseItemModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArtifactMeta

`func (o *V0ArtifactResponseItemModel) GetArtifactMeta() []int32`

GetArtifactMeta returns the ArtifactMeta field if non-nil, zero value otherwise.

### GetArtifactMetaOk

`func (o *V0ArtifactResponseItemModel) GetArtifactMetaOk() (*[]int32, bool)`

GetArtifactMetaOk returns a tuple with the ArtifactMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactMeta

`func (o *V0ArtifactResponseItemModel) SetArtifactMeta(v []int32)`

SetArtifactMeta sets ArtifactMeta field to given value.

### HasArtifactMeta

`func (o *V0ArtifactResponseItemModel) HasArtifactMeta() bool`

HasArtifactMeta returns a boolean if a field has been set.

### GetArtifactType

`func (o *V0ArtifactResponseItemModel) GetArtifactType() NullsString`

GetArtifactType returns the ArtifactType field if non-nil, zero value otherwise.

### GetArtifactTypeOk

`func (o *V0ArtifactResponseItemModel) GetArtifactTypeOk() (*NullsString, bool)`

GetArtifactTypeOk returns a tuple with the ArtifactType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactType

`func (o *V0ArtifactResponseItemModel) SetArtifactType(v NullsString)`

SetArtifactType sets ArtifactType field to given value.

### HasArtifactType

`func (o *V0ArtifactResponseItemModel) HasArtifactType() bool`

HasArtifactType returns a boolean if a field has been set.

### GetExpiringDownloadUrl

`func (o *V0ArtifactResponseItemModel) GetExpiringDownloadUrl() string`

GetExpiringDownloadUrl returns the ExpiringDownloadUrl field if non-nil, zero value otherwise.

### GetExpiringDownloadUrlOk

`func (o *V0ArtifactResponseItemModel) GetExpiringDownloadUrlOk() (*string, bool)`

GetExpiringDownloadUrlOk returns a tuple with the ExpiringDownloadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiringDownloadUrl

`func (o *V0ArtifactResponseItemModel) SetExpiringDownloadUrl(v string)`

SetExpiringDownloadUrl sets ExpiringDownloadUrl field to given value.

### HasExpiringDownloadUrl

`func (o *V0ArtifactResponseItemModel) HasExpiringDownloadUrl() bool`

HasExpiringDownloadUrl returns a boolean if a field has been set.

### GetFileSizeBytes

`func (o *V0ArtifactResponseItemModel) GetFileSizeBytes() int32`

GetFileSizeBytes returns the FileSizeBytes field if non-nil, zero value otherwise.

### GetFileSizeBytesOk

`func (o *V0ArtifactResponseItemModel) GetFileSizeBytesOk() (*int32, bool)`

GetFileSizeBytesOk returns a tuple with the FileSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSizeBytes

`func (o *V0ArtifactResponseItemModel) SetFileSizeBytes(v int32)`

SetFileSizeBytes sets FileSizeBytes field to given value.

### HasFileSizeBytes

`func (o *V0ArtifactResponseItemModel) HasFileSizeBytes() bool`

HasFileSizeBytes returns a boolean if a field has been set.

### GetIntermediateFileInfo

`func (o *V0ArtifactResponseItemModel) GetIntermediateFileInfo() []int32`

GetIntermediateFileInfo returns the IntermediateFileInfo field if non-nil, zero value otherwise.

### GetIntermediateFileInfoOk

`func (o *V0ArtifactResponseItemModel) GetIntermediateFileInfoOk() (*[]int32, bool)`

GetIntermediateFileInfoOk returns a tuple with the IntermediateFileInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntermediateFileInfo

`func (o *V0ArtifactResponseItemModel) SetIntermediateFileInfo(v []int32)`

SetIntermediateFileInfo sets IntermediateFileInfo field to given value.

### HasIntermediateFileInfo

`func (o *V0ArtifactResponseItemModel) HasIntermediateFileInfo() bool`

HasIntermediateFileInfo returns a boolean if a field has been set.

### GetIsPublicPageEnabled

`func (o *V0ArtifactResponseItemModel) GetIsPublicPageEnabled() bool`

GetIsPublicPageEnabled returns the IsPublicPageEnabled field if non-nil, zero value otherwise.

### GetIsPublicPageEnabledOk

`func (o *V0ArtifactResponseItemModel) GetIsPublicPageEnabledOk() (*bool, bool)`

GetIsPublicPageEnabledOk returns a tuple with the IsPublicPageEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublicPageEnabled

`func (o *V0ArtifactResponseItemModel) SetIsPublicPageEnabled(v bool)`

SetIsPublicPageEnabled sets IsPublicPageEnabled field to given value.

### HasIsPublicPageEnabled

`func (o *V0ArtifactResponseItemModel) HasIsPublicPageEnabled() bool`

HasIsPublicPageEnabled returns a boolean if a field has been set.

### GetPublicInstallPageUrl

`func (o *V0ArtifactResponseItemModel) GetPublicInstallPageUrl() string`

GetPublicInstallPageUrl returns the PublicInstallPageUrl field if non-nil, zero value otherwise.

### GetPublicInstallPageUrlOk

`func (o *V0ArtifactResponseItemModel) GetPublicInstallPageUrlOk() (*string, bool)`

GetPublicInstallPageUrlOk returns a tuple with the PublicInstallPageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicInstallPageUrl

`func (o *V0ArtifactResponseItemModel) SetPublicInstallPageUrl(v string)`

SetPublicInstallPageUrl sets PublicInstallPageUrl field to given value.

### HasPublicInstallPageUrl

`func (o *V0ArtifactResponseItemModel) HasPublicInstallPageUrl() bool`

HasPublicInstallPageUrl returns a boolean if a field has been set.

### GetSlug

`func (o *V0ArtifactResponseItemModel) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *V0ArtifactResponseItemModel) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *V0ArtifactResponseItemModel) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *V0ArtifactResponseItemModel) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetTitle

`func (o *V0ArtifactResponseItemModel) GetTitle() NullsString`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *V0ArtifactResponseItemModel) GetTitleOk() (*NullsString, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *V0ArtifactResponseItemModel) SetTitle(v NullsString)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *V0ArtifactResponseItemModel) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


