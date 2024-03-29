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

// NewActivityListParams creates a new ActivityListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewActivityListParams() *ActivityListParams {
	return &ActivityListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewActivityListParamsWithTimeout creates a new ActivityListParams object
// with the ability to set a timeout on a request.
func NewActivityListParamsWithTimeout(timeout time.Duration) *ActivityListParams {
	return &ActivityListParams{
		timeout: timeout,
	}
}

// NewActivityListParamsWithContext creates a new ActivityListParams object
// with the ability to set a context for a request.
func NewActivityListParamsWithContext(ctx context.Context) *ActivityListParams {
	return &ActivityListParams{
		Context: ctx,
	}
}

// NewActivityListParamsWithHTTPClient creates a new ActivityListParams object
// with the ability to set a custom HTTPClient for a request.
func NewActivityListParamsWithHTTPClient(client *http.Client) *ActivityListParams {
	return &ActivityListParams{
		HTTPClient: client,
	}
}

/*
ActivityListParams contains all the parameters to send to the API endpoint

	for the activity list operation.

	Typically these are written to a http.Request.
*/
type ActivityListParams struct {

	/* Limit.

	   Max number of elements per page (default: 50)
	*/
	Limit *int64

	/* Next.

	   Slug of the first activity event in the response
	*/
	Next *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the activity list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ActivityListParams) WithDefaults() *ActivityListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the activity list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ActivityListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the activity list params
func (o *ActivityListParams) WithTimeout(timeout time.Duration) *ActivityListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the activity list params
func (o *ActivityListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the activity list params
func (o *ActivityListParams) WithContext(ctx context.Context) *ActivityListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the activity list params
func (o *ActivityListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the activity list params
func (o *ActivityListParams) WithHTTPClient(client *http.Client) *ActivityListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the activity list params
func (o *ActivityListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the activity list params
func (o *ActivityListParams) WithLimit(limit *int64) *ActivityListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the activity list params
func (o *ActivityListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithNext adds the next to the activity list params
func (o *ActivityListParams) WithNext(next *string) *ActivityListParams {
	o.SetNext(next)
	return o
}

// SetNext adds the next to the activity list params
func (o *ActivityListParams) SetNext(next *string) {
	o.Next = next
}

// WriteToRequest writes these params to a swagger request
func (o *ActivityListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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
