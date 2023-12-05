# V0ReleaseCreateAppStoreParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutomaticTestflightUpload** | Pointer to **bool** | @Deprecated use &#x60;automatic_store_upload&#x60; instead. Indicates whether or not to upload every release candidate build automatically to TestFlight (default: &#x60;false&#x60;) | [optional] 
**BundleId** | **string** | The bundle ID of the app to be released | 
**Description** | Pointer to **string** | An internal description of the release; it won&#39;t be propagated to the App Store (default: empty) | [optional] 
**Name** | **string** | The name/version of the release (e.g. &#x60;1.2&#x60;) | 
**ReleaseBranch** | Pointer to **string** | The branch used for building the release candidate (default: empty) | [optional] 
**SlackWebhookUrl** | Pointer to **string** | The Slack webhook URL to use for sending Slack notifications (default: empty) | [optional] 
**TeamsWebhookUrl** | Pointer to **string** | The Teams webhook URL to use for sending MS Teams notifications (default: empty) | [optional] 
**Workflow** | Pointer to **string** | The workflow used for building the release candidate (default: empty) | [optional] 

## Methods

### NewV0ReleaseCreateAppStoreParams

`func NewV0ReleaseCreateAppStoreParams(bundleId string, name string, ) *V0ReleaseCreateAppStoreParams`

NewV0ReleaseCreateAppStoreParams instantiates a new V0ReleaseCreateAppStoreParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0ReleaseCreateAppStoreParamsWithDefaults

`func NewV0ReleaseCreateAppStoreParamsWithDefaults() *V0ReleaseCreateAppStoreParams`

NewV0ReleaseCreateAppStoreParamsWithDefaults instantiates a new V0ReleaseCreateAppStoreParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutomaticTestflightUpload

`func (o *V0ReleaseCreateAppStoreParams) GetAutomaticTestflightUpload() bool`

GetAutomaticTestflightUpload returns the AutomaticTestflightUpload field if non-nil, zero value otherwise.

### GetAutomaticTestflightUploadOk

`func (o *V0ReleaseCreateAppStoreParams) GetAutomaticTestflightUploadOk() (*bool, bool)`

GetAutomaticTestflightUploadOk returns a tuple with the AutomaticTestflightUpload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomaticTestflightUpload

`func (o *V0ReleaseCreateAppStoreParams) SetAutomaticTestflightUpload(v bool)`

SetAutomaticTestflightUpload sets AutomaticTestflightUpload field to given value.

### HasAutomaticTestflightUpload

`func (o *V0ReleaseCreateAppStoreParams) HasAutomaticTestflightUpload() bool`

HasAutomaticTestflightUpload returns a boolean if a field has been set.

### GetBundleId

`func (o *V0ReleaseCreateAppStoreParams) GetBundleId() string`

GetBundleId returns the BundleId field if non-nil, zero value otherwise.

### GetBundleIdOk

`func (o *V0ReleaseCreateAppStoreParams) GetBundleIdOk() (*string, bool)`

GetBundleIdOk returns a tuple with the BundleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleId

`func (o *V0ReleaseCreateAppStoreParams) SetBundleId(v string)`

SetBundleId sets BundleId field to given value.


### GetDescription

`func (o *V0ReleaseCreateAppStoreParams) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *V0ReleaseCreateAppStoreParams) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *V0ReleaseCreateAppStoreParams) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *V0ReleaseCreateAppStoreParams) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *V0ReleaseCreateAppStoreParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V0ReleaseCreateAppStoreParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V0ReleaseCreateAppStoreParams) SetName(v string)`

SetName sets Name field to given value.


### GetReleaseBranch

`func (o *V0ReleaseCreateAppStoreParams) GetReleaseBranch() string`

GetReleaseBranch returns the ReleaseBranch field if non-nil, zero value otherwise.

### GetReleaseBranchOk

`func (o *V0ReleaseCreateAppStoreParams) GetReleaseBranchOk() (*string, bool)`

GetReleaseBranchOk returns a tuple with the ReleaseBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseBranch

`func (o *V0ReleaseCreateAppStoreParams) SetReleaseBranch(v string)`

SetReleaseBranch sets ReleaseBranch field to given value.

### HasReleaseBranch

`func (o *V0ReleaseCreateAppStoreParams) HasReleaseBranch() bool`

HasReleaseBranch returns a boolean if a field has been set.

### GetSlackWebhookUrl

`func (o *V0ReleaseCreateAppStoreParams) GetSlackWebhookUrl() string`

GetSlackWebhookUrl returns the SlackWebhookUrl field if non-nil, zero value otherwise.

### GetSlackWebhookUrlOk

`func (o *V0ReleaseCreateAppStoreParams) GetSlackWebhookUrlOk() (*string, bool)`

GetSlackWebhookUrlOk returns a tuple with the SlackWebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlackWebhookUrl

`func (o *V0ReleaseCreateAppStoreParams) SetSlackWebhookUrl(v string)`

SetSlackWebhookUrl sets SlackWebhookUrl field to given value.

### HasSlackWebhookUrl

`func (o *V0ReleaseCreateAppStoreParams) HasSlackWebhookUrl() bool`

HasSlackWebhookUrl returns a boolean if a field has been set.

### GetTeamsWebhookUrl

`func (o *V0ReleaseCreateAppStoreParams) GetTeamsWebhookUrl() string`

GetTeamsWebhookUrl returns the TeamsWebhookUrl field if non-nil, zero value otherwise.

### GetTeamsWebhookUrlOk

`func (o *V0ReleaseCreateAppStoreParams) GetTeamsWebhookUrlOk() (*string, bool)`

GetTeamsWebhookUrlOk returns a tuple with the TeamsWebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamsWebhookUrl

`func (o *V0ReleaseCreateAppStoreParams) SetTeamsWebhookUrl(v string)`

SetTeamsWebhookUrl sets TeamsWebhookUrl field to given value.

### HasTeamsWebhookUrl

`func (o *V0ReleaseCreateAppStoreParams) HasTeamsWebhookUrl() bool`

HasTeamsWebhookUrl returns a boolean if a field has been set.

### GetWorkflow

`func (o *V0ReleaseCreateAppStoreParams) GetWorkflow() string`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *V0ReleaseCreateAppStoreParams) GetWorkflowOk() (*string, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *V0ReleaseCreateAppStoreParams) SetWorkflow(v string)`

SetWorkflow sets Workflow field to given value.

### HasWorkflow

`func (o *V0ReleaseCreateAppStoreParams) HasWorkflow() bool`

HasWorkflow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


