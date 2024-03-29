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

	"github.com/Staffbase/bitrise-api/go/models"
)

// NewPipelineAbortParams creates a new PipelineAbortParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPipelineAbortParams() *PipelineAbortParams {
	return &PipelineAbortParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPipelineAbortParamsWithTimeout creates a new PipelineAbortParams object
// with the ability to set a timeout on a request.
func NewPipelineAbortParamsWithTimeout(timeout time.Duration) *PipelineAbortParams {
	return &PipelineAbortParams{
		timeout: timeout,
	}
}

// NewPipelineAbortParamsWithContext creates a new PipelineAbortParams object
// with the ability to set a context for a request.
func NewPipelineAbortParamsWithContext(ctx context.Context) *PipelineAbortParams {
	return &PipelineAbortParams{
		Context: ctx,
	}
}

// NewPipelineAbortParamsWithHTTPClient creates a new PipelineAbortParams object
// with the ability to set a custom HTTPClient for a request.
func NewPipelineAbortParamsWithHTTPClient(client *http.Client) *PipelineAbortParams {
	return &PipelineAbortParams{
		HTTPClient: client,
	}
}

/*
PipelineAbortParams contains all the parameters to send to the API endpoint

	for the pipeline abort operation.

	Typically these are written to a http.Request.
*/
type PipelineAbortParams struct {

	/* AppSlug.

	   App slug
	*/
	AppSlug string

	/* PipelineAbortParams.

	   Pipeline abort parameters
	*/
	PipelineAbortParams *models.V0PipelineAbortParams

	/* PipelineID.

	   Pipeline id
	*/
	PipelineID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the pipeline abort params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PipelineAbortParams) WithDefaults() *PipelineAbortParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the pipeline abort params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PipelineAbortParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the pipeline abort params
func (o *PipelineAbortParams) WithTimeout(timeout time.Duration) *PipelineAbortParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pipeline abort params
func (o *PipelineAbortParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pipeline abort params
func (o *PipelineAbortParams) WithContext(ctx context.Context) *PipelineAbortParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pipeline abort params
func (o *PipelineAbortParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pipeline abort params
func (o *PipelineAbortParams) WithHTTPClient(client *http.Client) *PipelineAbortParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pipeline abort params
func (o *PipelineAbortParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppSlug adds the appSlug to the pipeline abort params
func (o *PipelineAbortParams) WithAppSlug(appSlug string) *PipelineAbortParams {
	o.SetAppSlug(appSlug)
	return o
}

// SetAppSlug adds the appSlug to the pipeline abort params
func (o *PipelineAbortParams) SetAppSlug(appSlug string) {
	o.AppSlug = appSlug
}

// WithPipelineAbortParams adds the pipelineAbortParams to the pipeline abort params
func (o *PipelineAbortParams) WithPipelineAbortParams(pipelineAbortParams *models.V0PipelineAbortParams) *PipelineAbortParams {
	o.SetPipelineAbortParams(pipelineAbortParams)
	return o
}

// SetPipelineAbortParams adds the pipelineAbortParams to the pipeline abort params
func (o *PipelineAbortParams) SetPipelineAbortParams(pipelineAbortParams *models.V0PipelineAbortParams) {
	o.PipelineAbortParams = pipelineAbortParams
}

// WithPipelineID adds the pipelineID to the pipeline abort params
func (o *PipelineAbortParams) WithPipelineID(pipelineID string) *PipelineAbortParams {
	o.SetPipelineID(pipelineID)
	return o
}

// SetPipelineID adds the pipelineId to the pipeline abort params
func (o *PipelineAbortParams) SetPipelineID(pipelineID string) {
	o.PipelineID = pipelineID
}

// WriteToRequest writes these params to a swagger request
func (o *PipelineAbortParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param app-slug
	if err := r.SetPathParam("app-slug", o.AppSlug); err != nil {
		return err
	}
	if o.PipelineAbortParams != nil {
		if err := r.SetBodyParam(o.PipelineAbortParams); err != nil {
			return err
		}
	}

	// path param pipeline-id
	if err := r.SetPathParam("pipeline-id", o.PipelineID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
