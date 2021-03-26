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

// NewAvatarCandidateListParams creates a new AvatarCandidateListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAvatarCandidateListParams() *AvatarCandidateListParams {
	return &AvatarCandidateListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAvatarCandidateListParamsWithTimeout creates a new AvatarCandidateListParams object
// with the ability to set a timeout on a request.
func NewAvatarCandidateListParamsWithTimeout(timeout time.Duration) *AvatarCandidateListParams {
	return &AvatarCandidateListParams{
		timeout: timeout,
	}
}

// NewAvatarCandidateListParamsWithContext creates a new AvatarCandidateListParams object
// with the ability to set a context for a request.
func NewAvatarCandidateListParamsWithContext(ctx context.Context) *AvatarCandidateListParams {
	return &AvatarCandidateListParams{
		Context: ctx,
	}
}

// NewAvatarCandidateListParamsWithHTTPClient creates a new AvatarCandidateListParams object
// with the ability to set a custom HTTPClient for a request.
func NewAvatarCandidateListParamsWithHTTPClient(client *http.Client) *AvatarCandidateListParams {
	return &AvatarCandidateListParams{
		HTTPClient: client,
	}
}

/* AvatarCandidateListParams contains all the parameters to send to the API endpoint
   for the avatar candidate list operation.

   Typically these are written to a http.Request.
*/
type AvatarCandidateListParams struct {

	/* AppSlug.

	   App slug
	*/
	AppSlug string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the avatar candidate list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AvatarCandidateListParams) WithDefaults() *AvatarCandidateListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the avatar candidate list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AvatarCandidateListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the avatar candidate list params
func (o *AvatarCandidateListParams) WithTimeout(timeout time.Duration) *AvatarCandidateListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the avatar candidate list params
func (o *AvatarCandidateListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the avatar candidate list params
func (o *AvatarCandidateListParams) WithContext(ctx context.Context) *AvatarCandidateListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the avatar candidate list params
func (o *AvatarCandidateListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the avatar candidate list params
func (o *AvatarCandidateListParams) WithHTTPClient(client *http.Client) *AvatarCandidateListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the avatar candidate list params
func (o *AvatarCandidateListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppSlug adds the appSlug to the avatar candidate list params
func (o *AvatarCandidateListParams) WithAppSlug(appSlug string) *AvatarCandidateListParams {
	o.SetAppSlug(appSlug)
	return o
}

// SetAppSlug adds the appSlug to the avatar candidate list params
func (o *AvatarCandidateListParams) SetAppSlug(appSlug string) {
	o.AppSlug = appSlug
}

// WriteToRequest writes these params to a swagger request
func (o *AvatarCandidateListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param app-slug
	if err := r.SetPathParam("app-slug", o.AppSlug); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
