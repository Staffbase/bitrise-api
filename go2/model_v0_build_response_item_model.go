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

// checks if the V0BuildResponseItemModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0BuildResponseItemModel{}

// V0BuildResponseItemModel struct for V0BuildResponseItemModel
type V0BuildResponseItemModel struct {
	AbortReason *NullsString `json:"abort_reason,omitempty"`
	Branch *NullsString `json:"branch,omitempty"`
	BuildNumber *int32 `json:"build_number,omitempty"`
	CommitHash *NullsString `json:"commit_hash,omitempty"`
	CommitMessage *NullsString `json:"commit_message,omitempty"`
	CommitViewUrl *NullsString `json:"commit_view_url,omitempty"`
	CreditCost *NullsInt64 `json:"credit_cost,omitempty"`
	EnvironmentPrepareFinishedAt *string `json:"environment_prepare_finished_at,omitempty"`
	FinishedAt *string `json:"finished_at,omitempty"`
	IsOnHold *bool `json:"is_on_hold,omitempty"`
	IsProcessed *bool `json:"is_processed,omitempty"`
	IsStatusSent *bool `json:"is_status_sent,omitempty"`
	LogFormat *string `json:"log_format,omitempty"`
	MachineTypeId *NullsString `json:"machine_type_id,omitempty"`
	OriginalBuildParams []int32 `json:"original_build_params,omitempty"`
	PipelineWorkflowId *string `json:"pipeline_workflow_id,omitempty"`
	PullRequestId *int32 `json:"pull_request_id,omitempty"`
	PullRequestTargetBranch *NullsString `json:"pull_request_target_branch,omitempty"`
	PullRequestViewUrl *NullsString `json:"pull_request_view_url,omitempty"`
	Slug *string `json:"slug,omitempty"`
	StackIdentifier *NullsString `json:"stack_identifier,omitempty"`
	StartedOnWorkerAt *string `json:"started_on_worker_at,omitempty"`
	Status *int32 `json:"status,omitempty"`
	StatusText *string `json:"status_text,omitempty"`
	Tag *NullsString `json:"tag,omitempty"`
	TriggeredAt *string `json:"triggered_at,omitempty"`
	TriggeredBy *NullsString `json:"triggered_by,omitempty"`
	TriggeredWorkflow *string `json:"triggered_workflow,omitempty"`
}

// NewV0BuildResponseItemModel instantiates a new V0BuildResponseItemModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0BuildResponseItemModel() *V0BuildResponseItemModel {
	this := V0BuildResponseItemModel{}
	return &this
}

// NewV0BuildResponseItemModelWithDefaults instantiates a new V0BuildResponseItemModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0BuildResponseItemModelWithDefaults() *V0BuildResponseItemModel {
	this := V0BuildResponseItemModel{}
	return &this
}

// GetAbortReason returns the AbortReason field value if set, zero value otherwise.
func (o *V0BuildResponseItemModel) GetAbortReason() NullsString {
	if o == nil || IsNil(o.AbortReason) {
		var ret NullsString
		return ret
	}
	return *o.AbortReason
}

// GetAbortReasonOk returns a tuple with the AbortReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0BuildResponseItemModel) GetAbortReasonOk() (*NullsString, bool) {
	if o == nil || IsNil(o.AbortReason) {
		return nil, false
	}
	return o.AbortReason, true
}

// HasAbortReason returns a boolean if a field has been set.
func (o *V0BuildResponseItemModel) HasAbortReason() bool {
	if o != nil && !IsNil(o.AbortReason) {
		return true
	}

	return false
}

// SetAbortReason gets a reference to the given NullsString and assigns it to the AbortReason field.
func (o *V0BuildResponseItemModel) SetAbortReason(v NullsString) {
	o.AbortReason = &v
}

// GetBranch returns the Branch field value if set, zero value otherwise.
func (o *V0BuildResponseItemModel) GetBranch() NullsString {
	if o == nil || IsNil(o.Branch) {
		var ret NullsString
		return ret
	}
	return *o.Branch
}

// GetBranchOk returns a tuple with the Branch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0BuildResponseItemModel) GetBranchOk() (*NullsString, bool) {
	if o == nil || IsNil(o.Branch) {
		return nil, false
	}
	return o.Branch, true
}

// HasBranch returns a boolean if a field has been set.
func (o *V0BuildResponseItemModel) HasBranch() bool {
	if o != nil && !IsNil(o.Branch) {
		return true
	}

	return false
}

// SetBranch gets a reference to the given NullsString and assigns it to the Branch field.
func (o *V0BuildResponseItemModel) SetBranch(v NullsString) {
	o.Branch = &v
}

// GetBuildNumber returns the BuildNumber field value if set, zero value otherwise.
func (o *V0BuildResponseItemModel) GetBuildNumber() int32 {
	if o == nil || IsNil(o.BuildNumber) {
		var ret int32
		return ret
	}
	return *o.BuildNumber
}

// GetBuildNumberOk returns a tuple with the BuildNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0BuildResponseItemModel) GetBuildNumberOk() (*int32, bool) {
	if o == nil || IsNil(o.BuildNumber) {
		return nil, false
	}
	return o.BuildNumber, true
}

// HasBuildNumber returns a boolean if a field has been set.
func (o *V0BuildResponseItemModel) HasBuildNumber() bool {
	if o != nil && !IsNil(o.BuildNumber) {
		return true
	}

	return false
}

// SetBuildNumber gets a reference to the given int32 and assigns it to the BuildNumber field.
func (o *V0BuildResponseItemModel) SetBuildNumber(v int32) {
	o.BuildNumber = &v
}

// GetCommitHash returns the CommitHash field value if set, zero value otherwise.
func (o *V0BuildResponseItemModel) GetCommitHash() NullsString {
	if o == nil || IsNil(o.CommitHash) {
		var ret NullsString
		return ret
	}
	return *o.CommitHash
}

// GetCommitHashOk returns a tuple with the CommitHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0BuildResponseItemModel) GetCommitHashOk() (*NullsString, bool) {
	if o == nil || IsNil(o.CommitHash) {
		return nil, false
	}
	return o.CommitHash, true
}

// HasCommitHash returns a boolean if a field has been set.
func (o *V0BuildResponseItemModel) HasCommitHash() bool {
	if o != nil && !IsNil(o.CommitHash) {
		return true
	}

	return false
}

// SetCommitHash gets a reference to the given NullsString and assigns it to the CommitHash field.
func (o *V0BuildResponseItemModel) SetCommitHash(v NullsString) {
	o.CommitHash = &v
}

// GetCommitMessage returns the CommitMessage field value if set, zero value otherwise.
func (o *V0BuildResponseItemModel) GetCommitMessage() NullsString {
	if o == nil || IsNil(o.CommitMessage) {
		var ret NullsString
		return ret
	}
	return *o.CommitMessage
}

// GetCommitMessageOk returns a tuple with the CommitMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0BuildResponseItemModel) GetCommitMessageOk() (*NullsString, bool) {
	if o == nil || IsNil(o.CommitMessage) {
		return nil, false
	}
	return o.CommitMessage, true
}

// HasCommitMessage returns a boolean if a field has been set.
func (o *V0BuildResponseItemModel) HasCommitMessage() bool {
	if o != nil && !IsNil(o.CommitMessage) {
		return true
	}

	return false
}

// SetCommitMessage gets a reference to the given NullsString and assigns it to the CommitMessage field.
func (o *V0BuildResponseItemModel) SetCommitMessage(v NullsString) {
	o.CommitMessage = &v
}

// GetCommitViewUrl returns the CommitViewUrl field value if set, zero value otherwise.
func (o *V0BuildResponseItemModel) GetCommitViewUrl() NullsString {
	if o == nil || IsNil(o.CommitViewUrl) {
		var ret NullsString
		return ret
	}
	return *o.CommitViewUrl
}

// GetCommitViewUrlOk returns a tuple with the CommitViewUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0BuildResponseItemModel) GetCommitViewUrlOk() (*NullsString, bool) {
	if o == nil || IsNil(o.CommitViewUrl) {
		return nil, false
	}
	return o.CommitViewUrl, true
}

// HasCommitViewUrl returns a boolean if a field has been set.
func (o *V0BuildResponseItemModel) HasCommitViewUrl() bool {
	if o != nil && !IsNil(o.CommitViewUrl) {
		return true
	}

	return false
}

// SetCommitViewUrl gets a reference to the given NullsString and assigns it to the CommitViewUrl field.
func (o *V0BuildResponseItemModel) SetCommitViewUrl(v NullsString) {
	o.CommitViewUrl = &v
}

// GetCreditCost returns the CreditCost field value if set, zero value otherwise.
func (o *V0BuildResponseItemModel) GetCreditCost() NullsInt64 {
	if o == nil || IsNil(o.CreditCost) {
		var ret NullsInt64
		return ret
	}
	return *o.CreditCost
}

// GetCreditCostOk returns a tuple with the CreditCost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0BuildResponseItemModel) GetCreditCostOk() (*NullsInt64, bool) {
	if o == nil || IsNil(o.CreditCost) {
		return nil, false
	}
	return o.CreditCost, true
}

// HasCreditCost returns a boolean if a field has been set.
func (o *V0BuildResponseItemModel) HasCreditCost() bool {
	if o != nil && !IsNil(o.CreditCost) {
		return true
	}

	return false
}

// SetCreditCost gets a reference to the given NullsInt64 and assigns it to the CreditCost field.
func (o *V0BuildResponseItemModel) SetCreditCost(v NullsInt64) {
	o.CreditCost = &v
}

// GetEnvironmentPrepareFinishedAt returns the EnvironmentPrepareFinishedAt field value if set, zero value otherwise.
func (o *V0BuildResponseItemModel) GetEnvironmentPrepareFinishedAt() string {
	if o == nil || IsNil(o.EnvironmentPrepareFinishedAt) {
		var ret string
		return ret
	}
	return *o.EnvironmentPrepareFinishedAt
}

// GetEnvironmentPrepareFinishedAtOk returns a tuple with the EnvironmentPrepareFinishedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0BuildResponseItemModel) GetEnvironmentPrepareFinishedAtOk() (*string, bool) {
	if o == nil || IsNil(o.EnvironmentPrepareFinishedAt) {
		return nil, false
	}
	return o.EnvironmentPrepareFinishedAt, true
}

// HasEnvironmentPrepareFinishedAt returns a boolean if a field has been set.
func (o *V0BuildResponseItemModel) HasEnvironmentPrepareFinishedAt() bool {
	if o != nil && !IsNil(o.EnvironmentPrepareFinishedAt) {
		return true
	}

	return false
}

// SetEnvironmentPrepareFinishedAt gets a reference to the given string and assigns it to the EnvironmentPrepareFinishedAt field.
func (o *V0BuildResponseItemModel) SetEnvironmentPrepareFinishedAt(v string) {
	o.EnvironmentPrepareFinishedAt = &v
}

// GetFinishedAt returns the FinishedAt field value if set, zero value otherwise.
func (o *V0BuildResponseItemModel) GetFinishedAt() string {
	if o == nil || IsNil(o.FinishedAt) {
		var ret string
		return ret
	}
	return *o.FinishedAt
}

// GetFinishedAtOk returns a tuple with the FinishedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0BuildResponseItemModel) GetFinishedAtOk() (*string, bool) {
	if o == nil || IsNil(o.FinishedAt) {
		return nil, false
	}
	return o.FinishedAt, true
}

// HasFinishedAt returns a boolean if a field has been set.
func (o *V0BuildResponseItemModel) HasFinishedAt() bool {
	if o != nil && !IsNil(o.FinishedAt) {
		return true
	}

	return false
}

// SetFinishedAt gets a reference to the given string and assigns it to the FinishedAt field.
func (o *V0BuildResponseItemModel) SetFinishedAt(v string) {
	o.FinishedAt = &v
}

// GetIsOnHold returns the IsOnHold field value if set, zero value otherwise.
func (o *V0BuildResponseItemModel) GetIsOnHold() bool {
	if o == nil || IsNil(o.IsOnHold) {
		var ret bool
		return ret
	}
	return *o.IsOnHold
}

// GetIsOnHoldOk returns a tuple with the IsOnHold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0BuildResponseItemModel) GetIsOnHoldOk() (*bool, bool) {
	if o == nil || IsNil(o.IsOnHold) {
		return nil, false
	}
	return o.IsOnHold, true
}

// HasIsOnHold returns a boolean if a field has been set.
func (o *V0BuildResponseItemModel) HasIsOnHold() bool {
	if o != nil && !IsNil(o.IsOnHold) {
		return true
	}

	return false
}

// SetIsOnHold gets a reference to the given bool and assigns it to the IsOnHold field.
func (o *V0BuildResponseItemModel) SetIsOnHold(v bool) {
	o.IsOnHold = &v
}

// GetIsProcessed returns the IsProcessed field value if set, zero value otherwise.
func (o *V0BuildResponseItemModel) GetIsProcessed() bool {
	if o == nil || IsNil(o.IsProcessed) {
		var ret bool
		return ret
	}
	return *o.IsProcessed
}

// GetIsProcessedOk returns a tuple with the IsProcessed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0BuildResponseItemModel) GetIsProcessedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsProcessed) {
		return nil, false
	}
	return o.IsProcessed, true
}

// HasIsProcessed returns a boolean if a field has been set.
func (o *V0BuildResponseItemModel) HasIsProcessed() bool {
	if o != nil && !IsNil(o.IsProcessed) {
		return true
	}

	return false
}

// SetIsProcessed gets a reference to the given bool and assigns it to the IsProcessed field.
func (o *V0BuildResponseItemModel) SetIsProcessed(v bool) {
	o.IsProcessed = &v
}

// GetIsStatusSent returns the IsStatusSent field value if set, zero value otherwise.
func (o *V0BuildResponseItemModel) GetIsStatusSent() bool {
	if o == nil || IsNil(o.IsStatusSent) {
		var ret bool
		return ret
	}
	return *o.IsStatusSent
}

// GetIsStatusSentOk returns a tuple with the IsStatusSent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0BuildResponseItemModel) GetIsStatusSentOk() (*bool, bool) {
	if o == nil || IsNil(o.IsStatusSent) {
		return nil, false
	}
	return o.IsStatusSent, true
}

// HasIsStatusSent returns a boolean if a field has been set.
func (o *V0BuildResponseItemModel) HasIsStatusSent() bool {
	if o != nil && !IsNil(o.IsStatusSent) {
		return true
	}

	return false
}

// SetIsStatusSent gets a reference to the given bool and assigns it to the IsStatusSent field.
func (o *V0BuildResponseItemModel) SetIsStatusSent(v bool) {
	o.IsStatusSent = &v
}

// GetLogFormat returns the LogFormat field value if set, zero value otherwise.
func (o *V0BuildResponseItemModel) GetLogFormat() string {
	if o == nil || IsNil(o.LogFormat) {
		var ret string
		return ret
	}
	return *o.LogFormat
}

// GetLogFormatOk returns a tuple with the LogFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0BuildResponseItemModel) GetLogFormatOk() (*string, bool) {
	if o == nil || IsNil(o.LogFormat) {
		return nil, false
	}
	return o.LogFormat, true
}

// HasLogFormat returns a boolean if a field has been set.
func (o *V0BuildResponseItemModel) HasLogFormat() bool {
	if o != nil && !IsNil(o.LogFormat) {
		return true
	}

	return false
}

// SetLogFormat gets a reference to the given string and assigns it to the LogFormat field.
func (o *V0BuildResponseItemModel) SetLogFormat(v string) {
	o.LogFormat = &v
}

// GetMachineTypeId returns the MachineTypeId field value if set, zero value otherwise.
func (o *V0BuildResponseItemModel) GetMachineTypeId() NullsString {
	if o == nil || IsNil(o.MachineTypeId) {
		var ret NullsString
		return ret
	}
	return *o.MachineTypeId
}

// GetMachineTypeIdOk returns a tuple with the MachineTypeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0BuildResponseItemModel) GetMachineTypeIdOk() (*NullsString, bool) {
	if o == nil || IsNil(o.MachineTypeId) {
		return nil, false
	}
	return o.MachineTypeId, true
}

// HasMachineTypeId returns a boolean if a field has been set.
func (o *V0BuildResponseItemModel) HasMachineTypeId() bool {
	if o != nil && !IsNil(o.MachineTypeId) {
		return true
	}

	return false
}

// SetMachineTypeId gets a reference to the given NullsString and assigns it to the MachineTypeId field.
func (o *V0BuildResponseItemModel) SetMachineTypeId(v NullsString) {
	o.MachineTypeId = &v
}

// GetOriginalBuildParams returns the OriginalBuildParams field value if set, zero value otherwise.
func (o *V0BuildResponseItemModel) GetOriginalBuildParams() []int32 {
	if o == nil || IsNil(o.OriginalBuildParams) {
		var ret []int32
		return ret
	}
	return o.OriginalBuildParams
}

// GetOriginalBuildParamsOk returns a tuple with the OriginalBuildParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0BuildResponseItemModel) GetOriginalBuildParamsOk() ([]int32, bool) {
	if o == nil || IsNil(o.OriginalBuildParams) {
		return nil, false
	}
	return o.OriginalBuildParams, true
}

// HasOriginalBuildParams returns a boolean if a field has been set.
func (o *V0BuildResponseItemModel) HasOriginalBuildParams() bool {
	if o != nil && !IsNil(o.OriginalBuildParams) {
		return true
	}

	return false
}

// SetOriginalBuildParams gets a reference to the given []int32 and assigns it to the OriginalBuildParams field.
func (o *V0BuildResponseItemModel) SetOriginalBuildParams(v []int32) {
	o.OriginalBuildParams = v
}

// GetPipelineWorkflowId returns the PipelineWorkflowId field value if set, zero value otherwise.
func (o *V0BuildResponseItemModel) GetPipelineWorkflowId() string {
	if o == nil || IsNil(o.PipelineWorkflowId) {
		var ret string
		return ret
	}
	return *o.PipelineWorkflowId
}

// GetPipelineWorkflowIdOk returns a tuple with the PipelineWorkflowId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0BuildResponseItemModel) GetPipelineWorkflowIdOk() (*string, bool) {
	if o == nil || IsNil(o.PipelineWorkflowId) {
		return nil, false
	}
	return o.PipelineWorkflowId, true
}

// HasPipelineWorkflowId returns a boolean if a field has been set.
func (o *V0BuildResponseItemModel) HasPipelineWorkflowId() bool {
	if o != nil && !IsNil(o.PipelineWorkflowId) {
		return true
	}

	return false
}

// SetPipelineWorkflowId gets a reference to the given string and assigns it to the PipelineWorkflowId field.
func (o *V0BuildResponseItemModel) SetPipelineWorkflowId(v string) {
	o.PipelineWorkflowId = &v
}

// GetPullRequestId returns the PullRequestId field value if set, zero value otherwise.
func (o *V0BuildResponseItemModel) GetPullRequestId() int32 {
	if o == nil || IsNil(o.PullRequestId) {
		var ret int32
		return ret
	}
	return *o.PullRequestId
}

// GetPullRequestIdOk returns a tuple with the PullRequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0BuildResponseItemModel) GetPullRequestIdOk() (*int32, bool) {
	if o == nil || IsNil(o.PullRequestId) {
		return nil, false
	}
	return o.PullRequestId, true
}

// HasPullRequestId returns a boolean if a field has been set.
func (o *V0BuildResponseItemModel) HasPullRequestId() bool {
	if o != nil && !IsNil(o.PullRequestId) {
		return true
	}

	return false
}

// SetPullRequestId gets a reference to the given int32 and assigns it to the PullRequestId field.
func (o *V0BuildResponseItemModel) SetPullRequestId(v int32) {
	o.PullRequestId = &v
}

// GetPullRequestTargetBranch returns the PullRequestTargetBranch field value if set, zero value otherwise.
func (o *V0BuildResponseItemModel) GetPullRequestTargetBranch() NullsString {
	if o == nil || IsNil(o.PullRequestTargetBranch) {
		var ret NullsString
		return ret
	}
	return *o.PullRequestTargetBranch
}

// GetPullRequestTargetBranchOk returns a tuple with the PullRequestTargetBranch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0BuildResponseItemModel) GetPullRequestTargetBranchOk() (*NullsString, bool) {
	if o == nil || IsNil(o.PullRequestTargetBranch) {
		return nil, false
	}
	return o.PullRequestTargetBranch, true
}

// HasPullRequestTargetBranch returns a boolean if a field has been set.
func (o *V0BuildResponseItemModel) HasPullRequestTargetBranch() bool {
	if o != nil && !IsNil(o.PullRequestTargetBranch) {
		return true
	}

	return false
}

// SetPullRequestTargetBranch gets a reference to the given NullsString and assigns it to the PullRequestTargetBranch field.
func (o *V0BuildResponseItemModel) SetPullRequestTargetBranch(v NullsString) {
	o.PullRequestTargetBranch = &v
}

// GetPullRequestViewUrl returns the PullRequestViewUrl field value if set, zero value otherwise.
func (o *V0BuildResponseItemModel) GetPullRequestViewUrl() NullsString {
	if o == nil || IsNil(o.PullRequestViewUrl) {
		var ret NullsString
		return ret
	}
	return *o.PullRequestViewUrl
}

// GetPullRequestViewUrlOk returns a tuple with the PullRequestViewUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0BuildResponseItemModel) GetPullRequestViewUrlOk() (*NullsString, bool) {
	if o == nil || IsNil(o.PullRequestViewUrl) {
		return nil, false
	}
	return o.PullRequestViewUrl, true
}

// HasPullRequestViewUrl returns a boolean if a field has been set.
func (o *V0BuildResponseItemModel) HasPullRequestViewUrl() bool {
	if o != nil && !IsNil(o.PullRequestViewUrl) {
		return true
	}

	return false
}

// SetPullRequestViewUrl gets a reference to the given NullsString and assigns it to the PullRequestViewUrl field.
func (o *V0BuildResponseItemModel) SetPullRequestViewUrl(v NullsString) {
	o.PullRequestViewUrl = &v
}

// GetSlug returns the Slug field value if set, zero value otherwise.
func (o *V0BuildResponseItemModel) GetSlug() string {
	if o == nil || IsNil(o.Slug) {
		var ret string
		return ret
	}
	return *o.Slug
}

// GetSlugOk returns a tuple with the Slug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0BuildResponseItemModel) GetSlugOk() (*string, bool) {
	if o == nil || IsNil(o.Slug) {
		return nil, false
	}
	return o.Slug, true
}

// HasSlug returns a boolean if a field has been set.
func (o *V0BuildResponseItemModel) HasSlug() bool {
	if o != nil && !IsNil(o.Slug) {
		return true
	}

	return false
}

// SetSlug gets a reference to the given string and assigns it to the Slug field.
func (o *V0BuildResponseItemModel) SetSlug(v string) {
	o.Slug = &v
}

// GetStackIdentifier returns the StackIdentifier field value if set, zero value otherwise.
func (o *V0BuildResponseItemModel) GetStackIdentifier() NullsString {
	if o == nil || IsNil(o.StackIdentifier) {
		var ret NullsString
		return ret
	}
	return *o.StackIdentifier
}

// GetStackIdentifierOk returns a tuple with the StackIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0BuildResponseItemModel) GetStackIdentifierOk() (*NullsString, bool) {
	if o == nil || IsNil(o.StackIdentifier) {
		return nil, false
	}
	return o.StackIdentifier, true
}

// HasStackIdentifier returns a boolean if a field has been set.
func (o *V0BuildResponseItemModel) HasStackIdentifier() bool {
	if o != nil && !IsNil(o.StackIdentifier) {
		return true
	}

	return false
}

// SetStackIdentifier gets a reference to the given NullsString and assigns it to the StackIdentifier field.
func (o *V0BuildResponseItemModel) SetStackIdentifier(v NullsString) {
	o.StackIdentifier = &v
}

// GetStartedOnWorkerAt returns the StartedOnWorkerAt field value if set, zero value otherwise.
func (o *V0BuildResponseItemModel) GetStartedOnWorkerAt() string {
	if o == nil || IsNil(o.StartedOnWorkerAt) {
		var ret string
		return ret
	}
	return *o.StartedOnWorkerAt
}

// GetStartedOnWorkerAtOk returns a tuple with the StartedOnWorkerAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0BuildResponseItemModel) GetStartedOnWorkerAtOk() (*string, bool) {
	if o == nil || IsNil(o.StartedOnWorkerAt) {
		return nil, false
	}
	return o.StartedOnWorkerAt, true
}

// HasStartedOnWorkerAt returns a boolean if a field has been set.
func (o *V0BuildResponseItemModel) HasStartedOnWorkerAt() bool {
	if o != nil && !IsNil(o.StartedOnWorkerAt) {
		return true
	}

	return false
}

// SetStartedOnWorkerAt gets a reference to the given string and assigns it to the StartedOnWorkerAt field.
func (o *V0BuildResponseItemModel) SetStartedOnWorkerAt(v string) {
	o.StartedOnWorkerAt = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *V0BuildResponseItemModel) GetStatus() int32 {
	if o == nil || IsNil(o.Status) {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0BuildResponseItemModel) GetStatusOk() (*int32, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *V0BuildResponseItemModel) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *V0BuildResponseItemModel) SetStatus(v int32) {
	o.Status = &v
}

// GetStatusText returns the StatusText field value if set, zero value otherwise.
func (o *V0BuildResponseItemModel) GetStatusText() string {
	if o == nil || IsNil(o.StatusText) {
		var ret string
		return ret
	}
	return *o.StatusText
}

// GetStatusTextOk returns a tuple with the StatusText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0BuildResponseItemModel) GetStatusTextOk() (*string, bool) {
	if o == nil || IsNil(o.StatusText) {
		return nil, false
	}
	return o.StatusText, true
}

// HasStatusText returns a boolean if a field has been set.
func (o *V0BuildResponseItemModel) HasStatusText() bool {
	if o != nil && !IsNil(o.StatusText) {
		return true
	}

	return false
}

// SetStatusText gets a reference to the given string and assigns it to the StatusText field.
func (o *V0BuildResponseItemModel) SetStatusText(v string) {
	o.StatusText = &v
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *V0BuildResponseItemModel) GetTag() NullsString {
	if o == nil || IsNil(o.Tag) {
		var ret NullsString
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0BuildResponseItemModel) GetTagOk() (*NullsString, bool) {
	if o == nil || IsNil(o.Tag) {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *V0BuildResponseItemModel) HasTag() bool {
	if o != nil && !IsNil(o.Tag) {
		return true
	}

	return false
}

// SetTag gets a reference to the given NullsString and assigns it to the Tag field.
func (o *V0BuildResponseItemModel) SetTag(v NullsString) {
	o.Tag = &v
}

// GetTriggeredAt returns the TriggeredAt field value if set, zero value otherwise.
func (o *V0BuildResponseItemModel) GetTriggeredAt() string {
	if o == nil || IsNil(o.TriggeredAt) {
		var ret string
		return ret
	}
	return *o.TriggeredAt
}

// GetTriggeredAtOk returns a tuple with the TriggeredAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0BuildResponseItemModel) GetTriggeredAtOk() (*string, bool) {
	if o == nil || IsNil(o.TriggeredAt) {
		return nil, false
	}
	return o.TriggeredAt, true
}

// HasTriggeredAt returns a boolean if a field has been set.
func (o *V0BuildResponseItemModel) HasTriggeredAt() bool {
	if o != nil && !IsNil(o.TriggeredAt) {
		return true
	}

	return false
}

// SetTriggeredAt gets a reference to the given string and assigns it to the TriggeredAt field.
func (o *V0BuildResponseItemModel) SetTriggeredAt(v string) {
	o.TriggeredAt = &v
}

// GetTriggeredBy returns the TriggeredBy field value if set, zero value otherwise.
func (o *V0BuildResponseItemModel) GetTriggeredBy() NullsString {
	if o == nil || IsNil(o.TriggeredBy) {
		var ret NullsString
		return ret
	}
	return *o.TriggeredBy
}

// GetTriggeredByOk returns a tuple with the TriggeredBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0BuildResponseItemModel) GetTriggeredByOk() (*NullsString, bool) {
	if o == nil || IsNil(o.TriggeredBy) {
		return nil, false
	}
	return o.TriggeredBy, true
}

// HasTriggeredBy returns a boolean if a field has been set.
func (o *V0BuildResponseItemModel) HasTriggeredBy() bool {
	if o != nil && !IsNil(o.TriggeredBy) {
		return true
	}

	return false
}

// SetTriggeredBy gets a reference to the given NullsString and assigns it to the TriggeredBy field.
func (o *V0BuildResponseItemModel) SetTriggeredBy(v NullsString) {
	o.TriggeredBy = &v
}

// GetTriggeredWorkflow returns the TriggeredWorkflow field value if set, zero value otherwise.
func (o *V0BuildResponseItemModel) GetTriggeredWorkflow() string {
	if o == nil || IsNil(o.TriggeredWorkflow) {
		var ret string
		return ret
	}
	return *o.TriggeredWorkflow
}

// GetTriggeredWorkflowOk returns a tuple with the TriggeredWorkflow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0BuildResponseItemModel) GetTriggeredWorkflowOk() (*string, bool) {
	if o == nil || IsNil(o.TriggeredWorkflow) {
		return nil, false
	}
	return o.TriggeredWorkflow, true
}

// HasTriggeredWorkflow returns a boolean if a field has been set.
func (o *V0BuildResponseItemModel) HasTriggeredWorkflow() bool {
	if o != nil && !IsNil(o.TriggeredWorkflow) {
		return true
	}

	return false
}

// SetTriggeredWorkflow gets a reference to the given string and assigns it to the TriggeredWorkflow field.
func (o *V0BuildResponseItemModel) SetTriggeredWorkflow(v string) {
	o.TriggeredWorkflow = &v
}

func (o V0BuildResponseItemModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0BuildResponseItemModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AbortReason) {
		toSerialize["abort_reason"] = o.AbortReason
	}
	if !IsNil(o.Branch) {
		toSerialize["branch"] = o.Branch
	}
	if !IsNil(o.BuildNumber) {
		toSerialize["build_number"] = o.BuildNumber
	}
	if !IsNil(o.CommitHash) {
		toSerialize["commit_hash"] = o.CommitHash
	}
	if !IsNil(o.CommitMessage) {
		toSerialize["commit_message"] = o.CommitMessage
	}
	if !IsNil(o.CommitViewUrl) {
		toSerialize["commit_view_url"] = o.CommitViewUrl
	}
	if !IsNil(o.CreditCost) {
		toSerialize["credit_cost"] = o.CreditCost
	}
	if !IsNil(o.EnvironmentPrepareFinishedAt) {
		toSerialize["environment_prepare_finished_at"] = o.EnvironmentPrepareFinishedAt
	}
	if !IsNil(o.FinishedAt) {
		toSerialize["finished_at"] = o.FinishedAt
	}
	if !IsNil(o.IsOnHold) {
		toSerialize["is_on_hold"] = o.IsOnHold
	}
	if !IsNil(o.IsProcessed) {
		toSerialize["is_processed"] = o.IsProcessed
	}
	if !IsNil(o.IsStatusSent) {
		toSerialize["is_status_sent"] = o.IsStatusSent
	}
	if !IsNil(o.LogFormat) {
		toSerialize["log_format"] = o.LogFormat
	}
	if !IsNil(o.MachineTypeId) {
		toSerialize["machine_type_id"] = o.MachineTypeId
	}
	if !IsNil(o.OriginalBuildParams) {
		toSerialize["original_build_params"] = o.OriginalBuildParams
	}
	if !IsNil(o.PipelineWorkflowId) {
		toSerialize["pipeline_workflow_id"] = o.PipelineWorkflowId
	}
	if !IsNil(o.PullRequestId) {
		toSerialize["pull_request_id"] = o.PullRequestId
	}
	if !IsNil(o.PullRequestTargetBranch) {
		toSerialize["pull_request_target_branch"] = o.PullRequestTargetBranch
	}
	if !IsNil(o.PullRequestViewUrl) {
		toSerialize["pull_request_view_url"] = o.PullRequestViewUrl
	}
	if !IsNil(o.Slug) {
		toSerialize["slug"] = o.Slug
	}
	if !IsNil(o.StackIdentifier) {
		toSerialize["stack_identifier"] = o.StackIdentifier
	}
	if !IsNil(o.StartedOnWorkerAt) {
		toSerialize["started_on_worker_at"] = o.StartedOnWorkerAt
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.StatusText) {
		toSerialize["status_text"] = o.StatusText
	}
	if !IsNil(o.Tag) {
		toSerialize["tag"] = o.Tag
	}
	if !IsNil(o.TriggeredAt) {
		toSerialize["triggered_at"] = o.TriggeredAt
	}
	if !IsNil(o.TriggeredBy) {
		toSerialize["triggered_by"] = o.TriggeredBy
	}
	if !IsNil(o.TriggeredWorkflow) {
		toSerialize["triggered_workflow"] = o.TriggeredWorkflow
	}
	return toSerialize, nil
}

type NullableV0BuildResponseItemModel struct {
	value *V0BuildResponseItemModel
	isSet bool
}

func (v NullableV0BuildResponseItemModel) Get() *V0BuildResponseItemModel {
	return v.value
}

func (v *NullableV0BuildResponseItemModel) Set(val *V0BuildResponseItemModel) {
	v.value = val
	v.isSet = true
}

func (v NullableV0BuildResponseItemModel) IsSet() bool {
	return v.isSet
}

func (v *NullableV0BuildResponseItemModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0BuildResponseItemModel(val *V0BuildResponseItemModel) *NullableV0BuildResponseItemModel {
	return &NullableV0BuildResponseItemModel{value: val, isSet: true}
}

func (v NullableV0BuildResponseItemModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0BuildResponseItemModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


