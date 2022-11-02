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

// NewBuildRequestListParams creates a new BuildRequestListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewBuildRequestListParams() *BuildRequestListParams {
	return &BuildRequestListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewBuildRequestListParamsWithTimeout creates a new BuildRequestListParams object
// with the ability to set a timeout on a request.
func NewBuildRequestListParamsWithTimeout(timeout time.Duration) *BuildRequestListParams {
	return &BuildRequestListParams{
		timeout: timeout,
	}
}

// NewBuildRequestListParamsWithContext creates a new BuildRequestListParams object
// with the ability to set a context for a request.
func NewBuildRequestListParamsWithContext(ctx context.Context) *BuildRequestListParams {
	return &BuildRequestListParams{
		Context: ctx,
	}
}

// NewBuildRequestListParamsWithHTTPClient creates a new BuildRequestListParams object
// with the ability to set a custom HTTPClient for a request.
func NewBuildRequestListParamsWithHTTPClient(client *http.Client) *BuildRequestListParams {
	return &BuildRequestListParams{
		HTTPClient: client,
	}
}

/*
BuildRequestListParams contains all the parameters to send to the API endpoint

	for the build request list operation.

	Typically these are written to a http.Request.
*/
type BuildRequestListParams struct {

	/* AppSlug.

	   App slug
	*/
	AppSlug string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the build request list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BuildRequestListParams) WithDefaults() *BuildRequestListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the build request list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BuildRequestListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the build request list params
func (o *BuildRequestListParams) WithTimeout(timeout time.Duration) *BuildRequestListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the build request list params
func (o *BuildRequestListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the build request list params
func (o *BuildRequestListParams) WithContext(ctx context.Context) *BuildRequestListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the build request list params
func (o *BuildRequestListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the build request list params
func (o *BuildRequestListParams) WithHTTPClient(client *http.Client) *BuildRequestListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the build request list params
func (o *BuildRequestListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppSlug adds the appSlug to the build request list params
func (o *BuildRequestListParams) WithAppSlug(appSlug string) *BuildRequestListParams {
	o.SetAppSlug(appSlug)
	return o
}

// SetAppSlug adds the appSlug to the build request list params
func (o *BuildRequestListParams) SetAppSlug(appSlug string) {
	o.AppSlug = appSlug
}

// WriteToRequest writes these params to a swagger request
func (o *BuildRequestListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
