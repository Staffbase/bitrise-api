# AddonsPlan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Features** | Pointer to [**[]AddonsFeature**](AddonsFeature.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Price** | Pointer to **int32** |  | [optional] 

## Methods

### NewAddonsPlan

`func NewAddonsPlan() *AddonsPlan`

NewAddonsPlan instantiates a new AddonsPlan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddonsPlanWithDefaults

`func NewAddonsPlanWithDefaults() *AddonsPlan`

NewAddonsPlanWithDefaults instantiates a new AddonsPlan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeatures

`func (o *AddonsPlan) GetFeatures() []AddonsFeature`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *AddonsPlan) GetFeaturesOk() (*[]AddonsFeature, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *AddonsPlan) SetFeatures(v []AddonsFeature)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *AddonsPlan) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetId

`func (o *AddonsPlan) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddonsPlan) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddonsPlan) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AddonsPlan) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AddonsPlan) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddonsPlan) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddonsPlan) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AddonsPlan) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPrice

`func (o *AddonsPlan) GetPrice() int32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *AddonsPlan) GetPriceOk() (*int32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *AddonsPlan) SetPrice(v int32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *AddonsPlan) HasPrice() bool`

HasPrice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


