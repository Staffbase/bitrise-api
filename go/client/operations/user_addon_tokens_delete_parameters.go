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

// NewUserAddonTokensDeleteParams creates a new UserAddonTokensDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUserAddonTokensDeleteParams() *UserAddonTokensDeleteParams {
	return &UserAddonTokensDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUserAddonTokensDeleteParamsWithTimeout creates a new UserAddonTokensDeleteParams object
// with the ability to set a timeout on a request.
func NewUserAddonTokensDeleteParamsWithTimeout(timeout time.Duration) *UserAddonTokensDeleteParams {
	return &UserAddonTokensDeleteParams{
		timeout: timeout,
	}
}

// NewUserAddonTokensDeleteParamsWithContext creates a new UserAddonTokensDeleteParams object
// with the ability to set a context for a request.
func NewUserAddonTokensDeleteParamsWithContext(ctx context.Context) *UserAddonTokensDeleteParams {
	return &UserAddonTokensDeleteParams{
		Context: ctx,
	}
}

// NewUserAddonTokensDeleteParamsWithHTTPClient creates a new UserAddonTokensDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewUserAddonTokensDeleteParamsWithHTTPClient(client *http.Client) *UserAddonTokensDeleteParams {
	return &UserAddonTokensDeleteParams{
		HTTPClient: client,
	}
}

/*
UserAddonTokensDeleteParams contains all the parameters to send to the API endpoint

	for the user addon tokens delete operation.

	Typically these are written to a http.Request.
*/
type UserAddonTokensDeleteParams struct {

	/* AddonID.

	   Addon ID
	*/
	AddonID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the user addon tokens delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UserAddonTokensDeleteParams) WithDefaults() *UserAddonTokensDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the user addon tokens delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UserAddonTokensDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the user addon tokens delete params
func (o *UserAddonTokensDeleteParams) WithTimeout(timeout time.Duration) *UserAddonTokensDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the user addon tokens delete params
func (o *UserAddonTokensDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the user addon tokens delete params
func (o *UserAddonTokensDeleteParams) WithContext(ctx context.Context) *UserAddonTokensDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the user addon tokens delete params
func (o *UserAddonTokensDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the user addon tokens delete params
func (o *UserAddonTokensDeleteParams) WithHTTPClient(client *http.Client) *UserAddonTokensDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the user addon tokens delete params
func (o *UserAddonTokensDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAddonID adds the addonID to the user addon tokens delete params
func (o *UserAddonTokensDeleteParams) WithAddonID(addonID string) *UserAddonTokensDeleteParams {
	o.SetAddonID(addonID)
	return o
}

// SetAddonID adds the addonId to the user addon tokens delete params
func (o *UserAddonTokensDeleteParams) SetAddonID(addonID string) {
	o.AddonID = addonID
}

// WriteToRequest writes these params to a swagger request
func (o *UserAddonTokensDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param addon-id
	if err := r.SetPathParam("addon-id", o.AddonID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}