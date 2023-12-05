# V0OrganizationDataModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvatarIconUrl** | Pointer to [**NullsString**](NullsString.md) |  | [optional] 
**ConcurrencyCount** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Owners** | Pointer to [**[]V0OrganizationOwner**](V0OrganizationOwner.md) |  | [optional] 
**Slug** | Pointer to **string** |  | [optional] 

## Methods

### NewV0OrganizationDataModel

`func NewV0OrganizationDataModel() *V0OrganizationDataModel`

NewV0OrganizationDataModel instantiates a new V0OrganizationDataModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0OrganizationDataModelWithDefaults

`func NewV0OrganizationDataModelWithDefaults() *V0OrganizationDataModel`

NewV0OrganizationDataModelWithDefaults instantiates a new V0OrganizationDataModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvatarIconUrl

`func (o *V0OrganizationDataModel) GetAvatarIconUrl() NullsString`

GetAvatarIconUrl returns the AvatarIconUrl field if non-nil, zero value otherwise.

### GetAvatarIconUrlOk

`func (o *V0OrganizationDataModel) GetAvatarIconUrlOk() (*NullsString, bool)`

GetAvatarIconUrlOk returns a tuple with the AvatarIconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarIconUrl

`func (o *V0OrganizationDataModel) SetAvatarIconUrl(v NullsString)`

SetAvatarIconUrl sets AvatarIconUrl field to given value.

### HasAvatarIconUrl

`func (o *V0OrganizationDataModel) HasAvatarIconUrl() bool`

HasAvatarIconUrl returns a boolean if a field has been set.

### GetConcurrencyCount

`func (o *V0OrganizationDataModel) GetConcurrencyCount() int32`

GetConcurrencyCount returns the ConcurrencyCount field if non-nil, zero value otherwise.

### GetConcurrencyCountOk

`func (o *V0OrganizationDataModel) GetConcurrencyCountOk() (*int32, bool)`

GetConcurrencyCountOk returns a tuple with the ConcurrencyCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcurrencyCount

`func (o *V0OrganizationDataModel) SetConcurrencyCount(v int32)`

SetConcurrencyCount sets ConcurrencyCount field to given value.

### HasConcurrencyCount

`func (o *V0OrganizationDataModel) HasConcurrencyCount() bool`

HasConcurrencyCount returns a boolean if a field has been set.

### GetName

`func (o *V0OrganizationDataModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V0OrganizationDataModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V0OrganizationDataModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V0OrganizationDataModel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOwners

`func (o *V0OrganizationDataModel) GetOwners() []V0OrganizationOwner`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *V0OrganizationDataModel) GetOwnersOk() (*[]V0OrganizationOwner, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *V0OrganizationDataModel) SetOwners(v []V0OrganizationOwner)`

SetOwners sets Owners field to given value.

### HasOwners

`func (o *V0OrganizationDataModel) HasOwners() bool`

HasOwners returns a boolean if a field has been set.

### GetSlug

`func (o *V0OrganizationDataModel) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *V0OrganizationDataModel) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *V0OrganizationDataModel) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *V0OrganizationDataModel) HasSlug() bool`

HasSlug returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


