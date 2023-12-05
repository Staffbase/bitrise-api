# V0ArtifactListElementResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArtifactMeta** | Pointer to **[]int32** |  | [optional] 
**ArtifactType** | Pointer to [**NullsString**](NullsString.md) |  | [optional] 
**FileSizeBytes** | Pointer to **int32** |  | [optional] 
**IntermediateFileInfo** | Pointer to **[]int32** |  | [optional] 
**IsPublicPageEnabled** | Pointer to **bool** |  | [optional] 
**Slug** | Pointer to **string** |  | [optional] 
**Title** | Pointer to [**NullsString**](NullsString.md) |  | [optional] 

## Methods

### NewV0ArtifactListElementResponseModel

`func NewV0ArtifactListElementResponseModel() *V0ArtifactListElementResponseModel`

NewV0ArtifactListElementResponseModel instantiates a new V0ArtifactListElementResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0ArtifactListElementResponseModelWithDefaults

`func NewV0ArtifactListElementResponseModelWithDefaults() *V0ArtifactListElementResponseModel`

NewV0ArtifactListElementResponseModelWithDefaults instantiates a new V0ArtifactListElementResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArtifactMeta

`func (o *V0ArtifactListElementResponseModel) GetArtifactMeta() []int32`

GetArtifactMeta returns the ArtifactMeta field if non-nil, zero value otherwise.

### GetArtifactMetaOk

`func (o *V0ArtifactListElementResponseModel) GetArtifactMetaOk() (*[]int32, bool)`

GetArtifactMetaOk returns a tuple with the ArtifactMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactMeta

`func (o *V0ArtifactListElementResponseModel) SetArtifactMeta(v []int32)`

SetArtifactMeta sets ArtifactMeta field to given value.

### HasArtifactMeta

`func (o *V0ArtifactListElementResponseModel) HasArtifactMeta() bool`

HasArtifactMeta returns a boolean if a field has been set.

### GetArtifactType

`func (o *V0ArtifactListElementResponseModel) GetArtifactType() NullsString`

GetArtifactType returns the ArtifactType field if non-nil, zero value otherwise.

### GetArtifactTypeOk

`func (o *V0ArtifactListElementResponseModel) GetArtifactTypeOk() (*NullsString, bool)`

GetArtifactTypeOk returns a tuple with the ArtifactType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactType

`func (o *V0ArtifactListElementResponseModel) SetArtifactType(v NullsString)`

SetArtifactType sets ArtifactType field to given value.

### HasArtifactType

`func (o *V0ArtifactListElementResponseModel) HasArtifactType() bool`

HasArtifactType returns a boolean if a field has been set.

### GetFileSizeBytes

`func (o *V0ArtifactListElementResponseModel) GetFileSizeBytes() int32`

GetFileSizeBytes returns the FileSizeBytes field if non-nil, zero value otherwise.

### GetFileSizeBytesOk

`func (o *V0ArtifactListElementResponseModel) GetFileSizeBytesOk() (*int32, bool)`

GetFileSizeBytesOk returns a tuple with the FileSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSizeBytes

`func (o *V0ArtifactListElementResponseModel) SetFileSizeBytes(v int32)`

SetFileSizeBytes sets FileSizeBytes field to given value.

### HasFileSizeBytes

`func (o *V0ArtifactListElementResponseModel) HasFileSizeBytes() bool`

HasFileSizeBytes returns a boolean if a field has been set.

### GetIntermediateFileInfo

`func (o *V0ArtifactListElementResponseModel) GetIntermediateFileInfo() []int32`

GetIntermediateFileInfo returns the IntermediateFileInfo field if non-nil, zero value otherwise.

### GetIntermediateFileInfoOk

`func (o *V0ArtifactListElementResponseModel) GetIntermediateFileInfoOk() (*[]int32, bool)`

GetIntermediateFileInfoOk returns a tuple with the IntermediateFileInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntermediateFileInfo

`func (o *V0ArtifactListElementResponseModel) SetIntermediateFileInfo(v []int32)`

SetIntermediateFileInfo sets IntermediateFileInfo field to given value.

### HasIntermediateFileInfo

`func (o *V0ArtifactListElementResponseModel) HasIntermediateFileInfo() bool`

HasIntermediateFileInfo returns a boolean if a field has been set.

### GetIsPublicPageEnabled

`func (o *V0ArtifactListElementResponseModel) GetIsPublicPageEnabled() bool`

GetIsPublicPageEnabled returns the IsPublicPageEnabled field if non-nil, zero value otherwise.

### GetIsPublicPageEnabledOk

`func (o *V0ArtifactListElementResponseModel) GetIsPublicPageEnabledOk() (*bool, bool)`

GetIsPublicPageEnabledOk returns a tuple with the IsPublicPageEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublicPageEnabled

`func (o *V0ArtifactListElementResponseModel) SetIsPublicPageEnabled(v bool)`

SetIsPublicPageEnabled sets IsPublicPageEnabled field to given value.

### HasIsPublicPageEnabled

`func (o *V0ArtifactListElementResponseModel) HasIsPublicPageEnabled() bool`

HasIsPublicPageEnabled returns a boolean if a field has been set.

### GetSlug

`func (o *V0ArtifactListElementResponseModel) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *V0ArtifactListElementResponseModel) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *V0ArtifactListElementResponseModel) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *V0ArtifactListElementResponseModel) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetTitle

`func (o *V0ArtifactListElementResponseModel) GetTitle() NullsString`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *V0ArtifactListElementResponseModel) GetTitleOk() (*NullsString, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *V0ArtifactListElementResponseModel) SetTitle(v NullsString)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *V0ArtifactListElementResponseModel) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


