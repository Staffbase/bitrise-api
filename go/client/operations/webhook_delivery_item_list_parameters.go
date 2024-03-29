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

// NewWebhookDeliveryItemListParams creates a new WebhookDeliveryItemListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWebhookDeliveryItemListParams() *WebhookDeliveryItemListParams {
	return &WebhookDeliveryItemListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewWebhookDeliveryItemListParamsWithTimeout creates a new WebhookDeliveryItemListParams object
// with the ability to set a timeout on a request.
func NewWebhookDeliveryItemListParamsWithTimeout(timeout time.Duration) *WebhookDeliveryItemListParams {
	return &WebhookDeliveryItemListParams{
		timeout: timeout,
	}
}

// NewWebhookDeliveryItemListParamsWithContext creates a new WebhookDeliveryItemListParams object
// with the ability to set a context for a request.
func NewWebhookDeliveryItemListParamsWithContext(ctx context.Context) *WebhookDeliveryItemListParams {
	return &WebhookDeliveryItemListParams{
		Context: ctx,
	}
}

// NewWebhookDeliveryItemListParamsWithHTTPClient creates a new WebhookDeliveryItemListParams object
// with the ability to set a custom HTTPClient for a request.
func NewWebhookDeliveryItemListParamsWithHTTPClient(client *http.Client) *WebhookDeliveryItemListParams {
	return &WebhookDeliveryItemListParams{
		HTTPClient: client,
	}
}

/*
WebhookDeliveryItemListParams contains all the parameters to send to the API endpoint

	for the webhook delivery item list operation.

	Typically these are written to a http.Request.
*/
type WebhookDeliveryItemListParams struct {

	/* AppSlug.

	   App slug
	*/
	AppSlug string

	/* AppWebhookSlug.

	   App webhook slug
	*/
	AppWebhookSlug string

	/* Limit.

	   Max number of elements per page (default: 50)
	*/
	Limit *int64

	/* Next.

	   Slug of the first delivery item in the response
	*/
	Next *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the webhook delivery item list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WebhookDeliveryItemListParams) WithDefaults() *WebhookDeliveryItemListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the webhook delivery item list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WebhookDeliveryItemListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the webhook delivery item list params
func (o *WebhookDeliveryItemListParams) WithTimeout(timeout time.Duration) *WebhookDeliveryItemListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the webhook delivery item list params
func (o *WebhookDeliveryItemListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the webhook delivery item list params
func (o *WebhookDeliveryItemListParams) WithContext(ctx context.Context) *WebhookDeliveryItemListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the webhook delivery item list params
func (o *WebhookDeliveryItemListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the webhook delivery item list params
func (o *WebhookDeliveryItemListParams) WithHTTPClient(client *http.Client) *WebhookDeliveryItemListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the webhook delivery item list params
func (o *WebhookDeliveryItemListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppSlug adds the appSlug to the webhook delivery item list params
func (o *WebhookDeliveryItemListParams) WithAppSlug(appSlug string) *WebhookDeliveryItemListParams {
	o.SetAppSlug(appSlug)
	return o
}

// SetAppSlug adds the appSlug to the webhook delivery item list params
func (o *WebhookDeliveryItemListParams) SetAppSlug(appSlug string) {
	o.AppSlug = appSlug
}

// WithAppWebhookSlug adds the appWebhookSlug to the webhook delivery item list params
func (o *WebhookDeliveryItemListParams) WithAppWebhookSlug(appWebhookSlug string) *WebhookDeliveryItemListParams {
	o.SetAppWebhookSlug(appWebhookSlug)
	return o
}

// SetAppWebhookSlug adds the appWebhookSlug to the webhook delivery item list params
func (o *WebhookDeliveryItemListParams) SetAppWebhookSlug(appWebhookSlug string) {
	o.AppWebhookSlug = appWebhookSlug
}

// WithLimit adds the limit to the webhook delivery item list params
func (o *WebhookDeliveryItemListParams) WithLimit(limit *int64) *WebhookDeliveryItemListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the webhook delivery item list params
func (o *WebhookDeliveryItemListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithNext adds the next to the webhook delivery item list params
func (o *WebhookDeliveryItemListParams) WithNext(next *string) *WebhookDeliveryItemListParams {
	o.SetNext(next)
	return o
}

// SetNext adds the next to the webhook delivery item list params
func (o *WebhookDeliveryItemListParams) SetNext(next *string) {
	o.Next = next
}

// WriteToRequest writes these params to a swagger request
func (o *WebhookDeliveryItemListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
