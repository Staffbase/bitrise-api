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

// NewBuildCertificateListParams creates a new BuildCertificateListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewBuildCertificateListParams() *BuildCertificateListParams {
	return &BuildCertificateListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewBuildCertificateListParamsWithTimeout creates a new BuildCertificateListParams object
// with the ability to set a timeout on a request.
func NewBuildCertificateListParamsWithTimeout(timeout time.Duration) *BuildCertificateListParams {
	return &BuildCertificateListParams{
		timeout: timeout,
	}
}

// NewBuildCertificateListParamsWithContext creates a new BuildCertificateListParams object
// with the ability to set a context for a request.
func NewBuildCertificateListParamsWithContext(ctx context.Context) *BuildCertificateListParams {
	return &BuildCertificateListParams{
		Context: ctx,
	}
}

// NewBuildCertificateListParamsWithHTTPClient creates a new BuildCertificateListParams object
// with the ability to set a custom HTTPClient for a request.
func NewBuildCertificateListParamsWithHTTPClient(client *http.Client) *BuildCertificateListParams {
	return &BuildCertificateListParams{
		HTTPClient: client,
	}
}

/*
BuildCertificateListParams contains all the parameters to send to the API endpoint

	for the build certificate list operation.

	Typically these are written to a http.Request.
*/
type BuildCertificateListParams struct {

	/* AppSlug.

	   App slug
	*/
	AppSlug string

	/* Limit.

	   Max number of build certificates per page is 50.
	*/
	Limit *int64

	/* Next.

	   Slug of the first build certificate in the response
	*/
	Next *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the build certificate list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BuildCertificateListParams) WithDefaults() *BuildCertificateListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the build certificate list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BuildCertificateListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the build certificate list params
func (o *BuildCertificateListParams) WithTimeout(timeout time.Duration) *BuildCertificateListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the build certificate list params
func (o *BuildCertificateListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the build certificate list params
func (o *BuildCertificateListParams) WithContext(ctx context.Context) *BuildCertificateListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the build certificate list params
func (o *BuildCertificateListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the build certificate list params
func (o *BuildCertificateListParams) WithHTTPClient(client *http.Client) *BuildCertificateListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the build certificate list params
func (o *BuildCertificateListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppSlug adds the appSlug to the build certificate list params
func (o *BuildCertificateListParams) WithAppSlug(appSlug string) *BuildCertificateListParams {
	o.SetAppSlug(appSlug)
	return o
}

// SetAppSlug adds the appSlug to the build certificate list params
func (o *BuildCertificateListParams) SetAppSlug(appSlug string) {
	o.AppSlug = appSlug
}

// WithLimit adds the limit to the build certificate list params
func (o *BuildCertificateListParams) WithLimit(limit *int64) *BuildCertificateListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the build certificate list params
func (o *BuildCertificateListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithNext adds the next to the build certificate list params
func (o *BuildCertificateListParams) WithNext(next *string) *BuildCertificateListParams {
	o.SetNext(next)
	return o
}

// SetNext adds the next to the build certificate list params
func (o *BuildCertificateListParams) SetNext(next *string) {
	o.Next = next
}

// WriteToRequest writes these params to a swagger request
func (o *BuildCertificateListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
