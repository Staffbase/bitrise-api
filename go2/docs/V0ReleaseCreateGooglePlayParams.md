# V0ReleaseCreateGooglePlayParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutomaticPlayConsoleUpload** | Pointer to **bool** | Indicates whether or not to upload every release candidate build automatically to the Google Play Console (default: &#x60;false&#x60;) | [optional] 
**Description** | Pointer to **string** | An internal description of the release; it won&#39;t be propagated to the App Store (default: empty) | [optional] 
**Name** | **string** | The name/version of the release (e.g. &#x60;1.2&#x60;) | 
**PackageName** | **string** | The package name of the app to be released | 
**ReleaseBranch** | Pointer to **string** | The branch used for building the release candidate (default: empty) | [optional] 
**SlackWebhookUrl** | Pointer to **string** | The Slack webhook URL to use for sending Slack notifications (default: empty) | [optional] 
**TeamsWebhookUrl** | Pointer to **string** | The Teams webhook URL to use for sending MS Teams notifications (default: empty) | [optional] 
**Workflow** | Pointer to **string** | The workflow used for building the release candidate (default: empty) | [optional] 

## Methods

### NewV0ReleaseCreateGooglePlayParams

`func NewV0ReleaseCreateGooglePlayParams(name string, packageName string, ) *V0ReleaseCreateGooglePlayParams`

NewV0ReleaseCreateGooglePlayParams instantiates a new V0ReleaseCreateGooglePlayParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0ReleaseCreateGooglePlayParamsWithDefaults

`func NewV0ReleaseCreateGooglePlayParamsWithDefaults() *V0ReleaseCreateGooglePlayParams`

NewV0ReleaseCreateGooglePlayParamsWithDefaults instantiates a new V0ReleaseCreateGooglePlayParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutomaticPlayConsoleUpload

`func (o *V0ReleaseCreateGooglePlayParams) GetAutomaticPlayConsoleUpload() bool`

GetAutomaticPlayConsoleUpload returns the AutomaticPlayConsoleUpload field if non-nil, zero value otherwise.

### GetAutomaticPlayConsoleUploadOk

`func (o *V0ReleaseCreateGooglePlayParams) GetAutomaticPlayConsoleUploadOk() (*bool, bool)`

GetAutomaticPlayConsoleUploadOk returns a tuple with the AutomaticPlayConsoleUpload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomaticPlayConsoleUpload

`func (o *V0ReleaseCreateGooglePlayParams) SetAutomaticPlayConsoleUpload(v bool)`

SetAutomaticPlayConsoleUpload sets AutomaticPlayConsoleUpload field to given value.

### HasAutomaticPlayConsoleUpload

`func (o *V0ReleaseCreateGooglePlayParams) HasAutomaticPlayConsoleUpload() bool`

HasAutomaticPlayConsoleUpload returns a boolean if a field has been set.

### GetDescription

`func (o *V0ReleaseCreateGooglePlayParams) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *V0ReleaseCreateGooglePlayParams) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *V0ReleaseCreateGooglePlayParams) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *V0ReleaseCreateGooglePlayParams) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *V0ReleaseCreateGooglePlayParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V0ReleaseCreateGooglePlayParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V0ReleaseCreateGooglePlayParams) SetName(v string)`

SetName sets Name field to given value.


### GetPackageName

`func (o *V0ReleaseCreateGooglePlayParams) GetPackageName() string`

GetPackageName returns the PackageName field if non-nil, zero value otherwise.

### GetPackageNameOk

`func (o *V0ReleaseCreateGooglePlayParams) GetPackageNameOk() (*string, bool)`

GetPackageNameOk returns a tuple with the PackageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageName

`func (o *V0ReleaseCreateGooglePlayParams) SetPackageName(v string)`

SetPackageName sets PackageName field to given value.


### GetReleaseBranch

`func (o *V0ReleaseCreateGooglePlayParams) GetReleaseBranch() string`

GetReleaseBranch returns the ReleaseBranch field if non-nil, zero value otherwise.

### GetReleaseBranchOk

`func (o *V0ReleaseCreateGooglePlayParams) GetReleaseBranchOk() (*string, bool)`

GetReleaseBranchOk returns a tuple with the ReleaseBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseBranch

`func (o *V0ReleaseCreateGooglePlayParams) SetReleaseBranch(v string)`

SetReleaseBranch sets ReleaseBranch field to given value.

### HasReleaseBranch

`func (o *V0ReleaseCreateGooglePlayParams) HasReleaseBranch() bool`

HasReleaseBranch returns a boolean if a field has been set.

### GetSlackWebhookUrl

`func (o *V0ReleaseCreateGooglePlayParams) GetSlackWebhookUrl() string`

GetSlackWebhookUrl returns the SlackWebhookUrl field if non-nil, zero value otherwise.

### GetSlackWebhookUrlOk

`func (o *V0ReleaseCreateGooglePlayParams) GetSlackWebhookUrlOk() (*string, bool)`

GetSlackWebhookUrlOk returns a tuple with the SlackWebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlackWebhookUrl

`func (o *V0ReleaseCreateGooglePlayParams) SetSlackWebhookUrl(v string)`

SetSlackWebhookUrl sets SlackWebhookUrl field to given value.

### HasSlackWebhookUrl

`func (o *V0ReleaseCreateGooglePlayParams) HasSlackWebhookUrl() bool`

HasSlackWebhookUrl returns a boolean if a field has been set.

### GetTeamsWebhookUrl

`func (o *V0ReleaseCreateGooglePlayParams) GetTeamsWebhookUrl() string`

GetTeamsWebhookUrl returns the TeamsWebhookUrl field if non-nil, zero value otherwise.

### GetTeamsWebhookUrlOk

`func (o *V0ReleaseCreateGooglePlayParams) GetTeamsWebhookUrlOk() (*string, bool)`

GetTeamsWebhookUrlOk returns a tuple with the TeamsWebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamsWebhookUrl

`func (o *V0ReleaseCreateGooglePlayParams) SetTeamsWebhookUrl(v string)`

SetTeamsWebhookUrl sets TeamsWebhookUrl field to given value.

### HasTeamsWebhookUrl

`func (o *V0ReleaseCreateGooglePlayParams) HasTeamsWebhookUrl() bool`

HasTeamsWebhookUrl returns a boolean if a field has been set.

### GetWorkflow

`func (o *V0ReleaseCreateGooglePlayParams) GetWorkflow() string`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *V0ReleaseCreateGooglePlayParams) GetWorkflowOk() (*string, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *V0ReleaseCreateGooglePlayParams) SetWorkflow(v string)`

SetWorkflow sets Workflow field to given value.

### HasWorkflow

`func (o *V0ReleaseCreateGooglePlayParams) HasWorkflow() bool`

HasWorkflow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


