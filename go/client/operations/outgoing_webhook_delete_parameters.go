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

// NewOutgoingWebhookDeleteParams creates a new OutgoingWebhookDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewOutgoingWebhookDeleteParams() *OutgoingWebhookDeleteParams {
	return &OutgoingWebhookDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewOutgoingWebhookDeleteParamsWithTimeout creates a new OutgoingWebhookDeleteParams object
// with the ability to set a timeout on a request.
func NewOutgoingWebhookDeleteParamsWithTimeout(timeout time.Duration) *OutgoingWebhookDeleteParams {
	return &OutgoingWebhookDeleteParams{
		timeout: timeout,
	}
}

// NewOutgoingWebhookDeleteParamsWithContext creates a new OutgoingWebhookDeleteParams object
// with the ability to set a context for a request.
func NewOutgoingWebhookDeleteParamsWithContext(ctx context.Context) *OutgoingWebhookDeleteParams {
	return &OutgoingWebhookDeleteParams{
		Context: ctx,
	}
}

// NewOutgoingWebhookDeleteParamsWithHTTPClient creates a new OutgoingWebhookDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewOutgoingWebhookDeleteParamsWithHTTPClient(client *http.Client) *OutgoingWebhookDeleteParams {
	return &OutgoingWebhookDeleteParams{
		HTTPClient: client,
	}
}

/* OutgoingWebhookDeleteParams contains all the parameters to send to the API endpoint
   for the outgoing webhook delete operation.

   Typically these are written to a http.Request.
*/
type OutgoingWebhookDeleteParams struct {

	/* AppSlug.

	   App slug
	*/
	AppSlug string

	/* AppWebhookSlug.

	   App webhook slug
	*/
	AppWebhookSlug string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the outgoing webhook delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OutgoingWebhookDeleteParams) WithDefaults() *OutgoingWebhookDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the outgoing webhook delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OutgoingWebhookDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the outgoing webhook delete params
func (o *OutgoingWebhookDeleteParams) WithTimeout(timeout time.Duration) *OutgoingWebhookDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the outgoing webhook delete params
func (o *OutgoingWebhookDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the outgoing webhook delete params
func (o *OutgoingWebhookDeleteParams) WithContext(ctx context.Context) *OutgoingWebhookDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the outgoing webhook delete params
func (o *OutgoingWebhookDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the outgoing webhook delete params
func (o *OutgoingWebhookDeleteParams) WithHTTPClient(client *http.Client) *OutgoingWebhookDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the outgoing webhook delete params
func (o *OutgoingWebhookDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppSlug adds the appSlug to the outgoing webhook delete params
func (o *OutgoingWebhookDeleteParams) WithAppSlug(appSlug string) *OutgoingWebhookDeleteParams {
	o.SetAppSlug(appSlug)
	return o
}

// SetAppSlug adds the appSlug to the outgoing webhook delete params
func (o *OutgoingWebhookDeleteParams) SetAppSlug(appSlug string) {
	o.AppSlug = appSlug
}

// WithAppWebhookSlug adds the appWebhookSlug to the outgoing webhook delete params
func (o *OutgoingWebhookDeleteParams) WithAppWebhookSlug(appWebhookSlug string) *OutgoingWebhookDeleteParams {
	o.SetAppWebhookSlug(appWebhookSlug)
	return o
}

// SetAppWebhookSlug adds the appWebhookSlug to the outgoing webhook delete params
func (o *OutgoingWebhookDeleteParams) SetAppWebhookSlug(appWebhookSlug string) {
	o.AppWebhookSlug = appWebhookSlug
}

// WriteToRequest writes these params to a swagger request
func (o *OutgoingWebhookDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param app-slug
	if err := r.SetPathParam("app-slug", o.AppSlug); err != nil {
		return err
	}

	// path param app-webhook-slug
	if err := r.SetPathParam("app-webhook-slug", o.AppWebhookSlug); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
