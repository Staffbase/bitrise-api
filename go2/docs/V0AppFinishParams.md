# V0AppFinishParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | Pointer to **string** | Which config to use. Specify a config that matches your project type (e. g. &#x60;default-android-config&#x60; for &#x60;android&#x60;, etc.). If not speficied, default value is &#x60;other-config&#x60;. The available values are &#x60;default-android-config&#x60;, &#x60;default-cordova-config&#x60;, &#x60;default-fastlane-android-config&#x60;, &#x60;default-fastlane-ios-config&#x60;, &#x60;flutter-config-notest-app-android&#x60;, &#x60;flutter-config-notest-app-both&#x60;, &#x60;flutter-config-notest-app-ios&#x60;, &#x60;flutter-config-test-app-android&#x60;, &#x60;flutter-config-test-app-both&#x60;, &#x60;flutter-config-test-app-ios&#x60;, &#x60;default-ionic-config&#x60;, &#x60;default-ios-config&#x60;, &#x60;default-macos-config&#x60;, &#x60;default-react-native-config&#x60;, &#x60;default-react-native-expo-config&#x60;, &#x60;other-config&#x60;. | [optional] 
**Envs** | Pointer to **map[string]string** | Environment variables for the application workflows, e.g. {\&quot;env1\&quot;:\&quot;val1\&quot;,\&quot;env2\&quot;:\&quot;val2\&quot;} | [optional] 
**Mode** | Pointer to **string** | config specification mode, has to be specified with &#x60;manual&#x60; value | [optional] 
**OrganizationSlug** | Pointer to **string** | The slug of the organization, who will be the owner of the application, if it&#39;s not specified, then the authenticated user will be the owner | [optional] 
**ProjectType** | **string** | The type of your project (&#x60;android&#x60;, &#x60;ios&#x60;, &#x60;cordova&#x60;, &#x60;other&#x60;, &#x60;xamarin&#x60;, &#x60;macos&#x60;, &#x60;ionic&#x60;, &#x60;react-native&#x60;, &#x60;fastlane&#x60;, null) | 
**StackId** | **string** | The id of the stack the application will be built (these can be found in the [system report](https://github.com/bitrise-io/bitrise.io/tree/master/system_reports) file names) | 

## Methods

### NewV0AppFinishParams

`func NewV0AppFinishParams(projectType string, stackId string, ) *V0AppFinishParams`

NewV0AppFinishParams instantiates a new V0AppFinishParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0AppFinishParamsWithDefaults

`func NewV0AppFinishParamsWithDefaults() *V0AppFinishParams`

NewV0AppFinishParamsWithDefaults instantiates a new V0AppFinishParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *V0AppFinishParams) GetConfig() string`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *V0AppFinishParams) GetConfigOk() (*string, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *V0AppFinishParams) SetConfig(v string)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *V0AppFinishParams) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetEnvs

`func (o *V0AppFinishParams) GetEnvs() map[string]string`

GetEnvs returns the Envs field if non-nil, zero value otherwise.

### GetEnvsOk

`func (o *V0AppFinishParams) GetEnvsOk() (*map[string]string, bool)`

GetEnvsOk returns a tuple with the Envs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvs

`func (o *V0AppFinishParams) SetEnvs(v map[string]string)`

SetEnvs sets Envs field to given value.

### HasEnvs

`func (o *V0AppFinishParams) HasEnvs() bool`

HasEnvs returns a boolean if a field has been set.

### GetMode

`func (o *V0AppFinishParams) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *V0AppFinishParams) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *V0AppFinishParams) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *V0AppFinishParams) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetOrganizationSlug

`func (o *V0AppFinishParams) GetOrganizationSlug() string`

GetOrganizationSlug returns the OrganizationSlug field if non-nil, zero value otherwise.

### GetOrganizationSlugOk

`func (o *V0AppFinishParams) GetOrganizationSlugOk() (*string, bool)`

GetOrganizationSlugOk returns a tuple with the OrganizationSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationSlug

`func (o *V0AppFinishParams) SetOrganizationSlug(v string)`

SetOrganizationSlug sets OrganizationSlug field to given value.

### HasOrganizationSlug

`func (o *V0AppFinishParams) HasOrganizationSlug() bool`

HasOrganizationSlug returns a boolean if a field has been set.

### GetProjectType

`func (o *V0AppFinishParams) GetProjectType() string`

GetProjectType returns the ProjectType field if non-nil, zero value otherwise.

### GetProjectTypeOk

`func (o *V0AppFinishParams) GetProjectTypeOk() (*string, bool)`

GetProjectTypeOk returns a tuple with the ProjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectType

`func (o *V0AppFinishParams) SetProjectType(v string)`

SetProjectType sets ProjectType field to given value.


### GetStackId

`func (o *V0AppFinishParams) GetStackId() string`

GetStackId returns the StackId field if non-nil, zero value otherwise.

### GetStackIdOk

`func (o *V0AppFinishParams) GetStackIdOk() (*string, bool)`

GetStackIdOk returns a tuple with the StackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackId

`func (o *V0AppFinishParams) SetStackId(v string)`

SetStackId sets StackId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


