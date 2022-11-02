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

// NewOutgoingWebhookListParams creates a new OutgoingWebhookListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewOutgoingWebhookListParams() *OutgoingWebhookListParams {
	return &OutgoingWebhookListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewOutgoingWebhookListParamsWithTimeout creates a new OutgoingWebhookListParams object
// with the ability to set a timeout on a request.
func NewOutgoingWebhookListParamsWithTimeout(timeout time.Duration) *OutgoingWebhookListParams {
	return &OutgoingWebhookListParams{
		timeout: timeout,
	}
}

// NewOutgoingWebhookListParamsWithContext creates a new OutgoingWebhookListParams object
// with the ability to set a context for a request.
func NewOutgoingWebhookListParamsWithContext(ctx context.Context) *OutgoingWebhookListParams {
	return &OutgoingWebhookListParams{
		Context: ctx,
	}
}

// NewOutgoingWebhookListParamsWithHTTPClient creates a new OutgoingWebhookListParams object
// with the ability to set a custom HTTPClient for a request.
func NewOutgoingWebhookListParamsWithHTTPClient(client *http.Client) *OutgoingWebhookListParams {
	return &OutgoingWebhookListParams{
		HTTPClient: client,
	}
}

/*
OutgoingWebhookListParams contains all the parameters to send to the API endpoint

	for the outgoing webhook list operation.

	Typically these are written to a http.Request.
*/
type OutgoingWebhookListParams struct {

	/* AppSlug.

	   App slug
	*/
	AppSlug string

	/* Limit.

	   Max number of elements per page (default: 50)
	*/
	Limit *int64

	/* Next.

	   Slug of the first webhook in the response
	*/
	Next *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the outgoing webhook list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OutgoingWebhookListParams) WithDefaults() *OutgoingWebhookListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the outgoing webhook list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OutgoingWebhookListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the outgoing webhook list params
func (o *OutgoingWebhookListParams) WithTimeout(timeout time.Duration) *OutgoingWebhookListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the outgoing webhook list params
func (o *OutgoingWebhookListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the outgoing webhook list params
func (o *OutgoingWebhookListParams) WithContext(ctx context.Context) *OutgoingWebhookListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the outgoing webhook list params
func (o *OutgoingWebhookListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the outgoing webhook list params
func (o *OutgoingWebhookListParams) WithHTTPClient(client *http.Client) *OutgoingWebhookListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the outgoing webhook list params
func (o *OutgoingWebhookListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppSlug adds the appSlug to the outgoing webhook list params
func (o *OutgoingWebhookListParams) WithAppSlug(appSlug string) *OutgoingWebhookListParams {
	o.SetAppSlug(appSlug)
	return o
}

// SetAppSlug adds the appSlug to the outgoing webhook list params
func (o *OutgoingWebhookListParams) SetAppSlug(appSlug string) {
	o.AppSlug = appSlug
}

// WithLimit adds the limit to the outgoing webhook list params
func (o *OutgoingWebhookListParams) WithLimit(limit *int64) *OutgoingWebhookListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the outgoing webhook list params
func (o *OutgoingWebhookListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithNext adds the next to the outgoing webhook list params
func (o *OutgoingWebhookListParams) WithNext(next *string) *OutgoingWebhookListParams {
	o.SetNext(next)
	return o
}

// SetNext adds the next to the outgoing webhook list params
func (o *OutgoingWebhookListParams) SetNext(next *string) {
	o.Next = next
}

// WriteToRequest writes these params to a swagger request
func (o *OutgoingWebhookListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param app-slug
	if err := r.SetPathParam("app-slug", o.AppSlug); err != nil {
		return err
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
