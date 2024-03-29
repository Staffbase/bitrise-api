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
)

// NewBuildWorkflowListParams creates a new BuildWorkflowListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewBuildWorkflowListParams() *BuildWorkflowListParams {
	return &BuildWorkflowListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewBuildWorkflowListParamsWithTimeout creates a new BuildWorkflowListParams object
// with the ability to set a timeout on a request.
func NewBuildWorkflowListParamsWithTimeout(timeout time.Duration) *BuildWorkflowListParams {
	return &BuildWorkflowListParams{
		timeout: timeout,
	}
}

// NewBuildWorkflowListParamsWithContext creates a new BuildWorkflowListParams object
// with the ability to set a context for a request.
func NewBuildWorkflowListParamsWithContext(ctx context.Context) *BuildWorkflowListParams {
	return &BuildWorkflowListParams{
		Context: ctx,
	}
}

// NewBuildWorkflowListParamsWithHTTPClient creates a new BuildWorkflowListParams object
// with the ability to set a custom HTTPClient for a request.
func NewBuildWorkflowListParamsWithHTTPClient(client *http.Client) *BuildWorkflowListParams {
	return &BuildWorkflowListParams{
		HTTPClient: client,
	}
}

/*
BuildWorkflowListParams contains all the parameters to send to the API endpoint

	for the build workflow list operation.

	Typically these are written to a http.Request.
*/
type BuildWorkflowListParams struct {

	/* AppSlug.

	   App slug
	*/
	AppSlug string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the build workflow list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BuildWorkflowListParams) WithDefaults() *BuildWorkflowListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the build workflow list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BuildWorkflowListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the build workflow list params
func (o *BuildWorkflowListParams) WithTimeout(timeout time.Duration) *BuildWorkflowListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the build workflow list params
func (o *BuildWorkflowListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the build workflow list params
func (o *BuildWorkflowListParams) WithContext(ctx context.Context) *BuildWorkflowListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the build workflow list params
func (o *BuildWorkflowListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the build workflow list params
func (o *BuildWorkflowListParams) WithHTTPClient(client *http.Client) *BuildWorkflowListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the build workflow list params
func (o *BuildWorkflowListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppSlug adds the appSlug to the build workflow list params
func (o *BuildWorkflowListParams) WithAppSlug(appSlug string) *BuildWorkflowListParams {
	o.SetAppSlug(appSlug)
	return o
}

// SetAppSlug adds the appSlug to the build workflow list params
func (o *BuildWorkflowListParams) SetAppSlug(appSlug string) {
	o.AppSlug = appSlug
}

// WriteToRequest writes these params to a swagger request
func (o *BuildWorkflowListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param app-slug
	if err := r.SetPathParam("app-slug", o.AppSlug); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
