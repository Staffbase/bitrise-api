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

// NewBuildCertificateShowParams creates a new BuildCertificateShowParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewBuildCertificateShowParams() *BuildCertificateShowParams {
	return &BuildCertificateShowParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewBuildCertificateShowParamsWithTimeout creates a new BuildCertificateShowParams object
// with the ability to set a timeout on a request.
func NewBuildCertificateShowParamsWithTimeout(timeout time.Duration) *BuildCertificateShowParams {
	return &BuildCertificateShowParams{
		timeout: timeout,
	}
}

// NewBuildCertificateShowParamsWithContext creates a new BuildCertificateShowParams object
// with the ability to set a context for a request.
func NewBuildCertificateShowParamsWithContext(ctx context.Context) *BuildCertificateShowParams {
	return &BuildCertificateShowParams{
		Context: ctx,
	}
}

// NewBuildCertificateShowParamsWithHTTPClient creates a new BuildCertificateShowParams object
// with the ability to set a custom HTTPClient for a request.
func NewBuildCertificateShowParamsWithHTTPClient(client *http.Client) *BuildCertificateShowParams {
	return &BuildCertificateShowParams{
		HTTPClient: client,
	}
}

/* BuildCertificateShowParams contains all the parameters to send to the API endpoint
   for the build certificate show operation.

   Typically these are written to a http.Request.
*/
type BuildCertificateShowParams struct {

	/* AppSlug.

	   App slug
	*/
	AppSlug string

	/* BuildCertificateSlug.

	   Build certificate slug
	*/
	BuildCertificateSlug string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the build certificate show params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BuildCertificateShowParams) WithDefaults() *BuildCertificateShowParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the build certificate show params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BuildCertificateShowParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the build certificate show params
func (o *BuildCertificateShowParams) WithTimeout(timeout time.Duration) *BuildCertificateShowParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the build certificate show params
func (o *BuildCertificateShowParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the build certificate show params
func (o *BuildCertificateShowParams) WithContext(ctx context.Context) *BuildCertificateShowParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the build certificate show params
func (o *BuildCertificateShowParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the build certificate show params
func (o *BuildCertificateShowParams) WithHTTPClient(client *http.Client) *BuildCertificateShowParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the build certificate show params
func (o *BuildCertificateShowParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppSlug adds the appSlug to the build certificate show params
func (o *BuildCertificateShowParams) WithAppSlug(appSlug string) *BuildCertificateShowParams {
	o.SetAppSlug(appSlug)
	return o
}

// SetAppSlug adds the appSlug to the build certificate show params
func (o *BuildCertificateShowParams) SetAppSlug(appSlug string) {
	o.AppSlug = appSlug
}

// WithBuildCertificateSlug adds the buildCertificateSlug to the build certificate show params
func (o *BuildCertificateShowParams) WithBuildCertificateSlug(buildCertificateSlug string) *BuildCertificateShowParams {
	o.SetBuildCertificateSlug(buildCertificateSlug)
	return o
}

// SetBuildCertificateSlug adds the buildCertificateSlug to the build certificate show params
func (o *BuildCertificateShowParams) SetBuildCertificateSlug(buildCertificateSlug string) {
	o.BuildCertificateSlug = buildCertificateSlug
}

// WriteToRequest writes these params to a swagger request
func (o *BuildCertificateShowParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param app-slug
	if err := r.SetPathParam("app-slug", o.AppSlug); err != nil {
		return err
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
