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

// NewSecretValueGetParams creates a new SecretValueGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSecretValueGetParams() *SecretValueGetParams {
	return &SecretValueGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSecretValueGetParamsWithTimeout creates a new SecretValueGetParams object
// with the ability to set a timeout on a request.
func NewSecretValueGetParamsWithTimeout(timeout time.Duration) *SecretValueGetParams {
	return &SecretValueGetParams{
		timeout: timeout,
	}
}

// NewSecretValueGetParamsWithContext creates a new SecretValueGetParams object
// with the ability to set a context for a request.
func NewSecretValueGetParamsWithContext(ctx context.Context) *SecretValueGetParams {
	return &SecretValueGetParams{
		Context: ctx,
	}
}

// NewSecretValueGetParamsWithHTTPClient creates a new SecretValueGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewSecretValueGetParamsWithHTTPClient(client *http.Client) *SecretValueGetParams {
	return &SecretValueGetParams{
		HTTPClient: client,
	}
}

/*
SecretValueGetParams contains all the parameters to send to the API endpoint

	for the secret value get operation.

	Typically these are written to a http.Request.
*/
type SecretValueGetParams struct {

	/* AppSlug.

	   App slug
	*/
	AppSlug string

	/* SecretName.

	   Secret name
	*/
	SecretName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the secret value get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SecretValueGetParams) WithDefaults() *SecretValueGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the secret value get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SecretValueGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the secret value get params
func (o *SecretValueGetParams) WithTimeout(timeout time.Duration) *SecretValueGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the secret value get params
func (o *SecretValueGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the secret value get params
func (o *SecretValueGetParams) WithContext(ctx context.Context) *SecretValueGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the secret value get params
func (o *SecretValueGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the secret value get params
func (o *SecretValueGetParams) WithHTTPClient(client *http.Client) *SecretValueGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the secret value get params
func (o *SecretValueGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppSlug adds the appSlug to the secret value get params
func (o *SecretValueGetParams) WithAppSlug(appSlug string) *SecretValueGetParams {
	o.SetAppSlug(appSlug)
	return o
}

// SetAppSlug adds the appSlug to the secret value get params
func (o *SecretValueGetParams) SetAppSlug(appSlug string) {
	o.AppSlug = appSlug
}

// WithSecretName adds the secretName to the secret value get params
func (o *SecretValueGetParams) WithSecretName(secretName string) *SecretValueGetParams {
	o.SetSecretName(secretName)
	return o
}

// SetSecretName adds the secretName to the secret value get params
func (o *SecretValueGetParams) SetSecretName(secretName string) {
	o.SecretName = secretName
}

// WriteToRequest writes these params to a swagger request
func (o *SecretValueGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param app-slug
	if err := r.SetPathParam("app-slug", o.AppSlug); err != nil {
		return err
	}

	// path param secret-name
	if err := r.SetPathParam("secret-name", o.SecretName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
