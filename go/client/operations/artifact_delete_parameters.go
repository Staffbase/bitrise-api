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

// NewArtifactDeleteParams creates a new ArtifactDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewArtifactDeleteParams() *ArtifactDeleteParams {
	return &ArtifactDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewArtifactDeleteParamsWithTimeout creates a new ArtifactDeleteParams object
// with the ability to set a timeout on a request.
func NewArtifactDeleteParamsWithTimeout(timeout time.Duration) *ArtifactDeleteParams {
	return &ArtifactDeleteParams{
		timeout: timeout,
	}
}

// NewArtifactDeleteParamsWithContext creates a new ArtifactDeleteParams object
// with the ability to set a context for a request.
func NewArtifactDeleteParamsWithContext(ctx context.Context) *ArtifactDeleteParams {
	return &ArtifactDeleteParams{
		Context: ctx,
	}
}

// NewArtifactDeleteParamsWithHTTPClient creates a new ArtifactDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewArtifactDeleteParamsWithHTTPClient(client *http.Client) *ArtifactDeleteParams {
	return &ArtifactDeleteParams{
		HTTPClient: client,
	}
}

/*
ArtifactDeleteParams contains all the parameters to send to the API endpoint

	for the artifact delete operation.

	Typically these are written to a http.Request.
*/
type ArtifactDeleteParams struct {

	/* AppSlug.

	   App slug
	*/
	AppSlug string

	/* ArtifactSlug.

	   Artifact slug
	*/
	ArtifactSlug string

	/* BuildSlug.

	   Build slug
	*/
	BuildSlug string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the artifact delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ArtifactDeleteParams) WithDefaults() *ArtifactDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the artifact delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ArtifactDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the artifact delete params
func (o *ArtifactDeleteParams) WithTimeout(timeout time.Duration) *ArtifactDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the artifact delete params
func (o *ArtifactDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the artifact delete params
func (o *ArtifactDeleteParams) WithContext(ctx context.Context) *ArtifactDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the artifact delete params
func (o *ArtifactDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the artifact delete params
func (o *ArtifactDeleteParams) WithHTTPClient(client *http.Client) *ArtifactDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the artifact delete params
func (o *ArtifactDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppSlug adds the appSlug to the artifact delete params
func (o *ArtifactDeleteParams) WithAppSlug(appSlug string) *ArtifactDeleteParams {
	o.SetAppSlug(appSlug)
	return o
}

// SetAppSlug adds the appSlug to the artifact delete params
func (o *ArtifactDeleteParams) SetAppSlug(appSlug string) {
	o.AppSlug = appSlug
}

// WithArtifactSlug adds the artifactSlug to the artifact delete params
func (o *ArtifactDeleteParams) WithArtifactSlug(artifactSlug string) *ArtifactDeleteParams {
	o.SetArtifactSlug(artifactSlug)
	return o
}

// SetArtifactSlug adds the artifactSlug to the artifact delete params
func (o *ArtifactDeleteParams) SetArtifactSlug(artifactSlug string) {
	o.ArtifactSlug = artifactSlug
}

// WithBuildSlug adds the buildSlug to the artifact delete params
func (o *ArtifactDeleteParams) WithBuildSlug(buildSlug string) *ArtifactDeleteParams {
	o.SetBuildSlug(buildSlug)
	return o
}

// SetBuildSlug adds the buildSlug to the artifact delete params
func (o *ArtifactDeleteParams) SetBuildSlug(buildSlug string) {
	o.BuildSlug = buildSlug
}

// WriteToRequest writes these params to a swagger request
func (o *ArtifactDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param app-slug
	if err := r.SetPathParam("app-slug", o.AppSlug); err != nil {
		return err
	}

	// path param artifact-slug
	if err := r.SetPathParam("artifact-slug", o.ArtifactSlug); err != nil {
		return err
	}

	// path param build-slug
	if err := r.SetPathParam("build-slug", o.BuildSlug); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
