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

// NewWebhookDeliveryItemRedeliverParams creates a new WebhookDeliveryItemRedeliverParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWebhookDeliveryItemRedeliverParams() *WebhookDeliveryItemRedeliverParams {
	return &WebhookDeliveryItemRedeliverParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewWebhookDeliveryItemRedeliverParamsWithTimeout creates a new WebhookDeliveryItemRedeliverParams object
// with the ability to set a timeout on a request.
func NewWebhookDeliveryItemRedeliverParamsWithTimeout(timeout time.Duration) *WebhookDeliveryItemRedeliverParams {
	return &WebhookDeliveryItemRedeliverParams{
		timeout: timeout,
	}
}

// NewWebhookDeliveryItemRedeliverParamsWithContext creates a new WebhookDeliveryItemRedeliverParams object
// with the ability to set a context for a request.
func NewWebhookDeliveryItemRedeliverParamsWithContext(ctx context.Context) *WebhookDeliveryItemRedeliverParams {
	return &WebhookDeliveryItemRedeliverParams{
		Context: ctx,
	}
}

// NewWebhookDeliveryItemRedeliverParamsWithHTTPClient creates a new WebhookDeliveryItemRedeliverParams object
// with the ability to set a custom HTTPClient for a request.
func NewWebhookDeliveryItemRedeliverParamsWithHTTPClient(client *http.Client) *WebhookDeliveryItemRedeliverParams {
	return &WebhookDeliveryItemRedeliverParams{
		HTTPClient: client,
	}
}

/*
WebhookDeliveryItemRedeliverParams contains all the parameters to send to the API endpoint

	for the webhook delivery item redeliver operation.

	Typically these are written to a http.Request.
*/
type WebhookDeliveryItemRedeliverParams struct {

	/* AppSlug.

	   App slug
	*/
	AppSlug string

	/* AppWebhookSlug.

	   App webhook slug
	*/
	AppWebhookSlug string

	/* WebhookDeliveryItemSlug.

	   Webhook delivery item slug
	*/
	WebhookDeliveryItemSlug string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the webhook delivery item redeliver params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WebhookDeliveryItemRedeliverParams) WithDefaults() *WebhookDeliveryItemRedeliverParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the webhook delivery item redeliver params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WebhookDeliveryItemRedeliverParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the webhook delivery item redeliver params
func (o *WebhookDeliveryItemRedeliverParams) WithTimeout(timeout time.Duration) *WebhookDeliveryItemRedeliverParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the webhook delivery item redeliver params
func (o *WebhookDeliveryItemRedeliverParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the webhook delivery item redeliver params
func (o *WebhookDeliveryItemRedeliverParams) WithContext(ctx context.Context) *WebhookDeliveryItemRedeliverParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the webhook delivery item redeliver params
func (o *WebhookDeliveryItemRedeliverParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the webhook delivery item redeliver params
func (o *WebhookDeliveryItemRedeliverParams) WithHTTPClient(client *http.Client) *WebhookDeliveryItemRedeliverParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the webhook delivery item redeliver params
func (o *WebhookDeliveryItemRedeliverParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppSlug adds the appSlug to the webhook delivery item redeliver params
func (o *WebhookDeliveryItemRedeliverParams) WithAppSlug(appSlug string) *WebhookDeliveryItemRedeliverParams {
	o.SetAppSlug(appSlug)
	return o
}

// SetAppSlug adds the appSlug to the webhook delivery item redeliver params
func (o *WebhookDeliveryItemRedeliverParams) SetAppSlug(appSlug string) {
	o.AppSlug = appSlug
}

// WithAppWebhookSlug adds the appWebhookSlug to the webhook delivery item redeliver params
func (o *WebhookDeliveryItemRedeliverParams) WithAppWebhookSlug(appWebhookSlug string) *WebhookDeliveryItemRedeliverParams {
	o.SetAppWebhookSlug(appWebhookSlug)
	return o
}

// SetAppWebhookSlug adds the appWebhookSlug to the webhook delivery item redeliver params
func (o *WebhookDeliveryItemRedeliverParams) SetAppWebhookSlug(appWebhookSlug string) {
	o.AppWebhookSlug = appWebhookSlug
}

// WithWebhookDeliveryItemSlug adds the webhookDeliveryItemSlug to the webhook delivery item redeliver params
func (o *WebhookDeliveryItemRedeliverParams) WithWebhookDeliveryItemSlug(webhookDeliveryItemSlug string) *WebhookDeliveryItemRedeliverParams {
	o.SetWebhookDeliveryItemSlug(webhookDeliveryItemSlug)
	return o
}

// SetWebhookDeliveryItemSlug adds the webhookDeliveryItemSlug to the webhook delivery item redeliver params
func (o *WebhookDeliveryItemRedeliverParams) SetWebhookDeliveryItemSlug(webhookDeliveryItemSlug string) {
	o.WebhookDeliveryItemSlug = webhookDeliveryItemSlug
}

// WriteToRequest writes these params to a swagger request
func (o *WebhookDeliveryItemRedeliverParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param webhook-delivery-item-slug
	if err := r.SetPathParam("webhook-delivery-item-slug", o.WebhookDeliveryItemSlug); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
