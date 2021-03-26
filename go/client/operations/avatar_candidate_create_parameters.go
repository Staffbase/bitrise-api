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

	"github.com/Staffbase/bitrise-api/go/models"
)

// NewAvatarCandidateCreateParams creates a new AvatarCandidateCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAvatarCandidateCreateParams() *AvatarCandidateCreateParams {
	return &AvatarCandidateCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAvatarCandidateCreateParamsWithTimeout creates a new AvatarCandidateCreateParams object
// with the ability to set a timeout on a request.
func NewAvatarCandidateCreateParamsWithTimeout(timeout time.Duration) *AvatarCandidateCreateParams {
	return &AvatarCandidateCreateParams{
		timeout: timeout,
	}
}

// NewAvatarCandidateCreateParamsWithContext creates a new AvatarCandidateCreateParams object
// with the ability to set a context for a request.
func NewAvatarCandidateCreateParamsWithContext(ctx context.Context) *AvatarCandidateCreateParams {
	return &AvatarCandidateCreateParams{
		Context: ctx,
	}
}

// NewAvatarCandidateCreateParamsWithHTTPClient creates a new AvatarCandidateCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewAvatarCandidateCreateParamsWithHTTPClient(client *http.Client) *AvatarCandidateCreateParams {
	return &AvatarCandidateCreateParams{
		HTTPClient: client,
	}
}

/* AvatarCandidateCreateParams contains all the parameters to send to the API endpoint
   for the avatar candidate create operation.

   Typically these are written to a http.Request.
*/
type AvatarCandidateCreateParams struct {

	/* AppSlug.

	   App slug
	*/
	AppSlug string

	/* AvatarCandidate.

	   Avatar candidate parameters
	*/
	AvatarCandidate []*models.V0AvatarCandidateCreateParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the avatar candidate create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AvatarCandidateCreateParams) WithDefaults() *AvatarCandidateCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the avatar candidate create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AvatarCandidateCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the avatar candidate create params
func (o *AvatarCandidateCreateParams) WithTimeout(timeout time.Duration) *AvatarCandidateCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the avatar candidate create params
func (o *AvatarCandidateCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the avatar candidate create params
func (o *AvatarCandidateCreateParams) WithContext(ctx context.Context) *AvatarCandidateCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the avatar candidate create params
func (o *AvatarCandidateCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the avatar candidate create params
func (o *AvatarCandidateCreateParams) WithHTTPClient(client *http.Client) *AvatarCandidateCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the avatar candidate create params
func (o *AvatarCandidateCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppSlug adds the appSlug to the avatar candidate create params
func (o *AvatarCandidateCreateParams) WithAppSlug(appSlug string) *AvatarCandidateCreateParams {
	o.SetAppSlug(appSlug)
	return o
}

// SetAppSlug adds the appSlug to the avatar candidate create params
func (o *AvatarCandidateCreateParams) SetAppSlug(appSlug string) {
	o.AppSlug = appSlug
}

// WithAvatarCandidate adds the avatarCandidate to the avatar candidate create params
func (o *AvatarCandidateCreateParams) WithAvatarCandidate(avatarCandidate []*models.V0AvatarCandidateCreateParams) *AvatarCandidateCreateParams {
	o.SetAvatarCandidate(avatarCandidate)
	return o
}

// SetAvatarCandidate adds the avatarCandidate to the avatar candidate create params
func (o *AvatarCandidateCreateParams) SetAvatarCandidate(avatarCandidate []*models.V0AvatarCandidateCreateParams) {
	o.AvatarCandidate = avatarCandidate
}

// WriteToRequest writes these params to a swagger request
func (o *AvatarCandidateCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param app-slug
	if err := r.SetPathParam("app-slug", o.AppSlug); err != nil {
		return err
	}
	if o.AvatarCandidate != nil {
		if err := r.SetBodyParam(o.AvatarCandidate); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
