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

// NewUserPlanParams creates a new UserPlanParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUserPlanParams() *UserPlanParams {
	return &UserPlanParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUserPlanParamsWithTimeout creates a new UserPlanParams object
// with the ability to set a timeout on a request.
func NewUserPlanParamsWithTimeout(timeout time.Duration) *UserPlanParams {
	return &UserPlanParams{
		timeout: timeout,
	}
}

// NewUserPlanParamsWithContext creates a new UserPlanParams object
// with the ability to set a context for a request.
func NewUserPlanParamsWithContext(ctx context.Context) *UserPlanParams {
	return &UserPlanParams{
		Context: ctx,
	}
}

// NewUserPlanParamsWithHTTPClient creates a new UserPlanParams object
// with the ability to set a custom HTTPClient for a request.
func NewUserPlanParamsWithHTTPClient(client *http.Client) *UserPlanParams {
	return &UserPlanParams{
		HTTPClient: client,
	}
}

/*
UserPlanParams contains all the parameters to send to the API endpoint

	for the user plan operation.

	Typically these are written to a http.Request.
*/
type UserPlanParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the user plan params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UserPlanParams) WithDefaults() *UserPlanParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the user plan params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UserPlanParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the user plan params
func (o *UserPlanParams) WithTimeout(timeout time.Duration) *UserPlanParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the user plan params
func (o *UserPlanParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the user plan params
func (o *UserPlanParams) WithContext(ctx context.Context) *UserPlanParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the user plan params
func (o *UserPlanParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the user plan params
func (o *UserPlanParams) WithHTTPClient(client *http.Client) *UserPlanParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the user plan params
func (o *UserPlanParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *UserPlanParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
