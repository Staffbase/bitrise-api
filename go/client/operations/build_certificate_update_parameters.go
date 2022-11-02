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

// NewBuildCertificateUpdateParams creates a new BuildCertificateUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewBuildCertificateUpdateParams() *BuildCertificateUpdateParams {
	return &BuildCertificateUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewBuildCertificateUpdateParamsWithTimeout creates a new BuildCertificateUpdateParams object
// with the ability to set a timeout on a request.
func NewBuildCertificateUpdateParamsWithTimeout(timeout time.Duration) *BuildCertificateUpdateParams {
	return &BuildCertificateUpdateParams{
		timeout: timeout,
	}
}

// NewBuildCertificateUpdateParamsWithContext creates a new BuildCertificateUpdateParams object
// with the ability to set a context for a request.
func NewBuildCertificateUpdateParamsWithContext(ctx context.Context) *BuildCertificateUpdateParams {
	return &BuildCertificateUpdateParams{
		Context: ctx,
	}
}

// NewBuildCertificateUpdateParamsWithHTTPClient creates a new BuildCertificateUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewBuildCertificateUpdateParamsWithHTTPClient(client *http.Client) *BuildCertificateUpdateParams {
	return &BuildCertificateUpdateParams{
		HTTPClient: client,
	}
}

/*
BuildCertificateUpdateParams contains all the parameters to send to the API endpoint

	for the build certificate update operation.

	Typically these are written to a http.Request.
*/
type BuildCertificateUpdateParams struct {

	/* AppSlug.

	   App slug
	*/
	AppSlug string

	/* BuildCertificate.

	   Build certificate parameters
	*/
	BuildCertificate *models.V0BuildCertificateUpdateParams

	/* BuildCertificateSlug.

	   Build certificate slug
	*/
	BuildCertificateSlug string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the build certificate update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BuildCertificateUpdateParams) WithDefaults() *BuildCertificateUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the build certificate update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BuildCertificateUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the build certificate update params
func (o *BuildCertificateUpdateParams) WithTimeout(timeout time.Duration) *BuildCertificateUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the build certificate update params
func (o *BuildCertificateUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the build certificate update params
func (o *BuildCertificateUpdateParams) WithContext(ctx context.Context) *BuildCertificateUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the build certificate update params
func (o *BuildCertificateUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the build certificate update params
func (o *BuildCertificateUpdateParams) WithHTTPClient(client *http.Client) *BuildCertificateUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the build certificate update params
func (o *BuildCertificateUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppSlug adds the appSlug to the build certificate update params
func (o *BuildCertificateUpdateParams) WithAppSlug(appSlug string) *BuildCertificateUpdateParams {
	o.SetAppSlug(appSlug)
	return o
}

// SetAppSlug adds the appSlug to the build certificate update params
func (o *BuildCertificateUpdateParams) SetAppSlug(appSlug string) {
	o.AppSlug = appSlug
}

// WithBuildCertificate adds the buildCertificate to the build certificate update params
func (o *BuildCertificateUpdateParams) WithBuildCertificate(buildCertificate *models.V0BuildCertificateUpdateParams) *BuildCertificateUpdateParams {
	o.SetBuildCertificate(buildCertificate)
	return o
}

// SetBuildCertificate adds the buildCertificate to the build certificate update params
func (o *BuildCertificateUpdateParams) SetBuildCertificate(buildCertificate *models.V0BuildCertificateUpdateParams) {
	o.BuildCertificate = buildCertificate
}

// WithBuildCertificateSlug adds the buildCertificateSlug to the build certificate update params
func (o *BuildCertificateUpdateParams) WithBuildCertificateSlug(buildCertificateSlug string) *BuildCertificateUpdateParams {
	o.SetBuildCertificateSlug(buildCertificateSlug)
	return o
}

// SetBuildCertificateSlug adds the buildCertificateSlug to the build certificate update params
func (o *BuildCertificateUpdateParams) SetBuildCertificateSlug(buildCertificateSlug string) {
	o.BuildCertificateSlug = buildCertificateSlug
}

// WriteToRequest writes these params to a swagger request
func (o *BuildCertificateUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param app-slug
	if err := r.SetPathParam("app-slug", o.AppSlug); err != nil {
		return err
	}
	if o.BuildCertificate != nil {
		if err := r.SetBodyParam(o.BuildCertificate); err != nil {
			return err
		}
	}

	// path param build-certificate-slug
	if err := r.SetPathParam("build-certificate-slug", o.BuildCertificateSlug); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
