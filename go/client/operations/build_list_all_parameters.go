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

// NewBuildListAllParams creates a new BuildListAllParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewBuildListAllParams() *BuildListAllParams {
	return &BuildListAllParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewBuildListAllParamsWithTimeout creates a new BuildListAllParams object
// with the ability to set a timeout on a request.
func NewBuildListAllParamsWithTimeout(timeout time.Duration) *BuildListAllParams {
	return &BuildListAllParams{
		timeout: timeout,
	}
}

// NewBuildListAllParamsWithContext creates a new BuildListAllParams object
// with the ability to set a context for a request.
func NewBuildListAllParamsWithContext(ctx context.Context) *BuildListAllParams {
	return &BuildListAllParams{
		Context: ctx,
	}
}

// NewBuildListAllParamsWithHTTPClient creates a new BuildListAllParams object
// with the ability to set a custom HTTPClient for a request.
func NewBuildListAllParamsWithHTTPClient(client *http.Client) *BuildListAllParams {
	return &BuildListAllParams{
		HTTPClient: client,
	}
}

/*
BuildListAllParams contains all the parameters to send to the API endpoint

	for the build list all operation.

	Typically these are written to a http.Request.
*/
type BuildListAllParams struct {

	/* IsOnHold.

	   Indicates whether the build has started yet (true: the build hasn't started)
	*/
	IsOnHold *bool

	/* Limit.

	   Max number of elements per page (default: 50)
	*/
	Limit *int64

	/* Next.

	   Slug of the first build in the response
	*/
	Next *string

	/* OwnerSlug.

	   The slug of the owner of the app or apps
	*/
	OwnerSlug *string

	/* Status.

	   The status of the build: not finished (0), successful (1), failed (2), aborted with failure (3), aborted with success (4)
	*/
	Status *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the build list all params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BuildListAllParams) WithDefaults() *BuildListAllParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the build list all params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BuildListAllParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the build list all params
func (o *BuildListAllParams) WithTimeout(timeout time.Duration) *BuildListAllParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the build list all params
func (o *BuildListAllParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the build list all params
func (o *BuildListAllParams) WithContext(ctx context.Context) *BuildListAllParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the build list all params
func (o *BuildListAllParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the build list all params
func (o *BuildListAllParams) WithHTTPClient(client *http.Client) *BuildListAllParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the build list all params
func (o *BuildListAllParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIsOnHold adds the isOnHold to the build list all params
func (o *BuildListAllParams) WithIsOnHold(isOnHold *bool) *BuildListAllParams {
	o.SetIsOnHold(isOnHold)
	return o
}

// SetIsOnHold adds the isOnHold to the build list all params
func (o *BuildListAllParams) SetIsOnHold(isOnHold *bool) {
	o.IsOnHold = isOnHold
}

// WithLimit adds the limit to the build list all params
func (o *BuildListAllParams) WithLimit(limit *int64) *BuildListAllParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the build list all params
func (o *BuildListAllParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithNext adds the next to the build list all params
func (o *BuildListAllParams) WithNext(next *string) *BuildListAllParams {
	o.SetNext(next)
	return o
}

// SetNext adds the next to the build list all params
func (o *BuildListAllParams) SetNext(next *string) {
	o.Next = next
}

// WithOwnerSlug adds the ownerSlug to the build list all params
func (o *BuildListAllParams) WithOwnerSlug(ownerSlug *string) *BuildListAllParams {
	o.SetOwnerSlug(ownerSlug)
	return o
}

// SetOwnerSlug adds the ownerSlug to the build list all params
func (o *BuildListAllParams) SetOwnerSlug(ownerSlug *string) {
	o.OwnerSlug = ownerSlug
}

// WithStatus adds the status to the build list all params
func (o *BuildListAllParams) WithStatus(status *int64) *BuildListAllParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the build list all params
func (o *BuildListAllParams) SetStatus(status *int64) {
	o.Status = status
}

// WriteToRequest writes these params to a swagger request
func (o *BuildListAllParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.IsOnHold != nil {

		// query param is_on_hold
		var qrIsOnHold bool

		if o.IsOnHold != nil {
			qrIsOnHold = *o.IsOnHold
		}
		qIsOnHold := swag.FormatBool(qrIsOnHold)
		if qIsOnHold != "" {

			if err := r.SetQueryParam("is_on_hold", qIsOnHold); err != nil {
				return err
			}
		}
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

	if o.OwnerSlug != nil {

		// query param owner_slug
		var qrOwnerSlug string

		if o.OwnerSlug != nil {
			qrOwnerSlug = *o.OwnerSlug
		}
		qOwnerSlug := qrOwnerSlug
		if qOwnerSlug != "" {

			if err := r.SetQueryParam("owner_slug", qOwnerSlug); err != nil {
				return err
			}
		}
	}

	if o.Status != nil {

		// query param status
		var qrStatus int64

		if o.Status != nil {
			qrStatus = *o.Status
		}
		qStatus := swag.FormatInt64(qrStatus)
		if qStatus != "" {

			if err := r.SetQueryParam("status", qStatus); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
