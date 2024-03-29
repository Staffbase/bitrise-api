// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewBuildListParams creates a new BuildListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewBuildListParams() *BuildListParams {
	return &BuildListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewBuildListParamsWithTimeout creates a new BuildListParams object
// with the ability to set a timeout on a request.
func NewBuildListParamsWithTimeout(timeout time.Duration) *BuildListParams {
	return &BuildListParams{
		timeout: timeout,
	}
}

// NewBuildListParamsWithContext creates a new BuildListParams object
// with the ability to set a context for a request.
func NewBuildListParamsWithContext(ctx context.Context) *BuildListParams {
	return &BuildListParams{
		Context: ctx,
	}
}

// NewBuildListParamsWithHTTPClient creates a new BuildListParams object
// with the ability to set a custom HTTPClient for a request.
func NewBuildListParamsWithHTTPClient(client *http.Client) *BuildListParams {
	return &BuildListParams{
		HTTPClient: client,
	}
}

/*
BuildListParams contains all the parameters to send to the API endpoint

	for the build list operation.

	Typically these are written to a http.Request.
*/
type BuildListParams struct {

	/* After.

	   List builds run after a given date (Unix Timestamp)
	*/
	After *int64

	/* AppSlug.

	   App slug
	*/
	AppSlug string

	/* Before.

	   List builds run before a given date (Unix Timestamp)
	*/
	Before *int64

	/* Branch.

	   The branch which was built
	*/
	Branch *string

	/* BuildNumber.

	   The build number
	*/
	BuildNumber *int64

	/* CommitMessage.

	   The commit message of the build
	*/
	CommitMessage *string

	/* IsPipelineBuild.

	   Whether the builds are part of a pipeline or not
	*/
	IsPipelineBuild *bool

	/* Limit.

	   Max number of elements per page (default: 50)
	*/
	Limit *int64

	/* Next.

	   Slug of the first build in the response
	*/
	Next *string

	/* PullRequestID.

	   The id of the pull request that triggered the build
	*/
	PullRequestID *int64

	/* SortBy.

	   Order of builds: sort them based on when they were created or the time when they were triggered
	*/
	SortBy *string

	/* Status.

	   The status of the build: not finished (0), successful (1), failed (2), aborted with failure (3), aborted with success (4)
	*/
	Status *int64

	/* TriggerEventType.

	   The event that triggered the build (push, pull-request, tag)
	*/
	TriggerEventType *string

	/* Workflow.

	   The name of the workflow used for the build
	*/
	Workflow *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the build list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BuildListParams) WithDefaults() *BuildListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the build list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BuildListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the build list params
func (o *BuildListParams) WithTimeout(timeout time.Duration) *BuildListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the build list params
func (o *BuildListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the build list params
func (o *BuildListParams) WithContext(ctx context.Context) *BuildListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the build list params
func (o *BuildListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the build list params
func (o *BuildListParams) WithHTTPClient(client *http.Client) *BuildListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the build list params
func (o *BuildListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAfter adds the after to the build list params
func (o *BuildListParams) WithAfter(after *int64) *BuildListParams {
	o.SetAfter(after)
	return o
}

// SetAfter adds the after to the build list params
func (o *BuildListParams) SetAfter(after *int64) {
	o.After = after
}

// WithAppSlug adds the appSlug to the build list params
func (o *BuildListParams) WithAppSlug(appSlug string) *BuildListParams {
	o.SetAppSlug(appSlug)
	return o
}

// SetAppSlug adds the appSlug to the build list params
func (o *BuildListParams) SetAppSlug(appSlug string) {
	o.AppSlug = appSlug
}

// WithBefore adds the before to the build list params
func (o *BuildListParams) WithBefore(before *int64) *BuildListParams {
	o.SetBefore(before)
	return o
}

// SetBefore adds the before to the build list params
func (o *BuildListParams) SetBefore(before *int64) {
	o.Before = before
}

// WithBranch adds the branch to the build list params
func (o *BuildListParams) WithBranch(branch *string) *BuildListParams {
	o.SetBranch(branch)
	return o
}

// SetBranch adds the branch to the build list params
func (o *BuildListParams) SetBranch(branch *string) {
	o.Branch = branch
}

// WithBuildNumber adds the buildNumber to the build list params
func (o *BuildListParams) WithBuildNumber(buildNumber *int64) *BuildListParams {
	o.SetBuildNumber(buildNumber)
	return o
}

// SetBuildNumber adds the buildNumber to the build list params
func (o *BuildListParams) SetBuildNumber(buildNumber *int64) {
	o.BuildNumber = buildNumber
}

// WithCommitMessage adds the commitMessage to the build list params
func (o *BuildListParams) WithCommitMessage(commitMessage *string) *BuildListParams {
	o.SetCommitMessage(commitMessage)
	return o
}

// SetCommitMessage adds the commitMessage to the build list params
func (o *BuildListParams) SetCommitMessage(commitMessage *string) {
	o.CommitMessage = commitMessage
}

// WithIsPipelineBuild adds the isPipelineBuild to the build list params
func (o *BuildListParams) WithIsPipelineBuild(isPipelineBuild *bool) *BuildListParams {
	o.SetIsPipelineBuild(isPipelineBuild)
	return o
}

// SetIsPipelineBuild adds the isPipelineBuild to the build list params
func (o *BuildListParams) SetIsPipelineBuild(isPipelineBuild *bool) {
	o.IsPipelineBuild = isPipelineBuild
}

// WithLimit adds the limit to the build list params
func (o *BuildListParams) WithLimit(limit *int64) *BuildListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the build list params
func (o *BuildListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithNext adds the next to the build list params
func (o *BuildListParams) WithNext(next *string) *BuildListParams {
	o.SetNext(next)
	return o
}

// SetNext adds the next to the build list params
func (o *BuildListParams) SetNext(next *string) {
	o.Next = next
}

// WithPullRequestID adds the pullRequestID to the build list params
func (o *BuildListParams) WithPullRequestID(pullRequestID *int64) *BuildListParams {
	o.SetPullRequestID(pullRequestID)
	return o
}

// SetPullRequestID adds the pullRequestId to the build list params
func (o *BuildListParams) SetPullRequestID(pullRequestID *int64) {
	o.PullRequestID = pullRequestID
}

// WithSortBy adds the sortBy to the build list params
func (o *BuildListParams) WithSortBy(sortBy *string) *BuildListParams {
	o.SetSortBy(sortBy)
	return o
}

// SetSortBy adds the sortBy to the build list params
func (o *BuildListParams) SetSortBy(sortBy *string) {
	o.SortBy = sortBy
}

// WithStatus adds the status to the build list params
func (o *BuildListParams) WithStatus(status *int64) *BuildListParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the build list params
func (o *BuildListParams) SetStatus(status *int64) {
	o.Status = status
}

// WithTriggerEventType adds the triggerEventType to the build list params
func (o *BuildListParams) WithTriggerEventType(triggerEventType *string) *BuildListParams {
	o.SetTriggerEventType(triggerEventType)
	return o
}

// SetTriggerEventType adds the triggerEventType to the build list params
func (o *BuildListParams) SetTriggerEventType(triggerEventType *string) {
	o.TriggerEventType = triggerEventType
}

// WithWorkflow adds the workflow to the build list params
func (o *BuildListParams) WithWorkflow(workflow *string) *BuildListParams {
	o.SetWorkflow(workflow)
	return o
}

// SetWorkflow adds the workflow to the build list params
func (o *BuildListParams) SetWorkflow(workflow *string) {
	o.Workflow = workflow
}

// WriteToRequest writes these params to a swagger request
func (o *BuildListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.After != nil {

		// query param after
		var qrAfter int64

		if o.After != nil {
			qrAfter = *o.After
		}
		qAfter := swag.FormatInt64(qrAfter)
		if qAfter != "" {

			if err := r.SetQueryParam("after", qAfter); err != nil {
				return err
			}
		}
	}

	// path param app-slug
	if err := r.SetPathParam("app-slug", o.AppSlug); err != nil {
		return err
	}

	if o.Before != nil {

		// query param before
		var qrBefore int64

		if o.Before != nil {
			qrBefore = *o.Before
		}
		qBefore := swag.FormatInt64(qrBefore)
		if qBefore != "" {

			if err := r.SetQueryParam("before", qBefore); err != nil {
				return err
			}
		}
	}

	if o.Branch != nil {

		// query param branch
		var qrBranch string

		if o.Branch != nil {
			qrBranch = *o.Branch
		}
		qBranch := qrBranch
		if qBranch != "" {

			if err := r.SetQueryParam("branch", qBranch); err != nil {
				return err
			}
		}
	}

	if o.BuildNumber != nil {

		// query param build_number
		var qrBuildNumber int64

		if o.BuildNumber != nil {
			qrBuildNumber = *o.BuildNumber
		}
		qBuildNumber := swag.FormatInt64(qrBuildNumber)
		if qBuildNumber != "" {

			if err := r.SetQueryParam("build_number", qBuildNumber); err != nil {
				return err
			}
		}
	}

	if o.CommitMessage != nil {

		// query param commit_message
		var qrCommitMessage string

		if o.CommitMessage != nil {
			qrCommitMessage = *o.CommitMessage
		}
		qCommitMessage := qrCommitMessage
		if qCommitMessage != "" {

			if err := r.SetQueryParam("commit_message", qCommitMessage); err != nil {
				return err
			}
		}
	}

	if o.IsPipelineBuild != nil {

		// query param is_pipeline_build
		var qrIsPipelineBuild bool

		if o.IsPipelineBuild != nil {
			qrIsPipelineBuild = *o.IsPipelineBuild
		}
		qIsPipelineBuild := swag.FormatBool(qrIsPipelineBuild)
		if qIsPipelineBuild != "" {

			if err := r.SetQueryParam("is_pipeline_build", qIsPipelineBuild); err != nil {
				return err
			}
		}
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {

			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}
	}

	if o.Next != nil {

		// query param next
		var qrNext string

		if o.Next != nil {
			qrNext = *o.Next
		}
		qNext := qrNext
		if qNext != "" {

			if err := r.SetQueryParam("next", qNext); err != nil {
				return err
			}
		}
	}

	if o.PullRequestID != nil {

		// query param pull_request_id
		var qrPullRequestID int64

		if o.PullRequestID != nil {
			qrPullRequestID = *o.PullRequestID
		}
		qPullRequestID := swag.FormatInt64(qrPullRequestID)
		if qPullRequestID != "" {

			if err := r.SetQueryParam("pull_request_id", qPullRequestID); err != nil {
				return err
			}
		}
	}

	if o.SortBy != nil {

		// query param sort_by
		var qrSortBy string

		if o.SortBy != nil {
			qrSortBy = *o.SortBy
		}
		qSortBy := qrSortBy
		if qSortBy != "" {

			if err := r.SetQueryParam("sort_by", qSortBy); err != nil {
				return err
			}
		}
	}

	if o.Status != nil {

		// query param status
		var qrStatus int64

		if o.Status != nil {
			qrStatus = *o.Status
		}
		qStatus := swag.FormatInt64(qrStatus)
		if qStatus != "" {

			if err := r.SetQueryParam("status", qStatus); err != nil {
				return err
			}
		}
	}

	if o.TriggerEventType != nil {

		// query param trigger_event_type
		var qrTriggerEventType string

		if o.TriggerEventType != nil {
			qrTriggerEventType = *o.TriggerEventType
		}
		qTriggerEventType := qrTriggerEventType
		if qTriggerEventType != "" {

			if err := r.SetQueryParam("trigger_event_type", qTriggerEventType); err != nil {
				return err
			}
		}
	}

	if o.Workflow != nil {

		// query param workflow
		var qrWorkflow string

		if o.Workflow != nil {
			qrWorkflow = *o.Workflow
		}
		qWorkflow := qrWorkflow
		if qWorkflow != "" {

			if err := r.SetQueryParam("workflow", qWorkflow); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
