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

// NewAppUpdateParams creates a new AppUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAppUpdateParams() *AppUpdateParams {
	return &AppUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAppUpdateParamsWithTimeout creates a new AppUpdateParams object
// with the ability to set a timeout on a request.
func NewAppUpdateParamsWithTimeout(timeout time.Duration) *AppUpdateParams {
	return &AppUpdateParams{
		timeout: timeout,
	}
}

// NewAppUpdateParamsWithContext creates a new AppUpdateParams object
// with the ability to set a context for a request.
func NewAppUpdateParamsWithContext(ctx context.Context) *AppUpdateParams {
	return &AppUpdateParams{
		Context: ctx,
	}
}

// NewAppUpdateParamsWithHTTPClient creates a new AppUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewAppUpdateParamsWithHTTPClient(client *http.Client) *AppUpdateParams {
	return &AppUpdateParams{
		HTTPClient: client,
	}
}

/*
AppUpdateParams contains all the parameters to send to the API endpoint

	for the app update operation.

	Typically these are written to a http.Request.
*/
type AppUpdateParams struct {

	/* App.

	   App update params. All fields are optional, omit the fields you don't wish to update.
	*/
	App *models.V0AppUpdateParams

	/* AppSlug.

	   App slug
	*/
	AppSlug string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the app update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AppUpdateParams) WithDefaults() *AppUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the app update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AppUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the app update params
func (o *AppUpdateParams) WithTimeout(timeout time.Duration) *AppUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the app update params
func (o *AppUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the app update params
func (o *AppUpdateParams) WithContext(ctx context.Context) *AppUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the app update params
func (o *AppUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the app update params
func (o *AppUpdateParams) WithHTTPClient(client *http.Client) *AppUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the app update params
func (o *AppUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithApp adds the app to the app update params
func (o *AppUpdateParams) WithApp(app *models.V0AppUpdateParams) *AppUpdateParams {
	o.SetApp(app)
	return o
}

// SetApp adds the app to the app update params
func (o *AppUpdateParams) SetApp(app *models.V0AppUpdateParams) {
	o.App = app
}

// WithAppSlug adds the appSlug to the app update params
func (o *AppUpdateParams) WithAppSlug(appSlug string) *AppUpdateParams {
	o.SetAppSlug(appSlug)
	return o
}

// SetAppSlug adds the appSlug to the app update params
func (o *AppUpdateParams) SetAppSlug(appSlug string) {
	o.AppSlug = appSlug
}

// WriteToRequest writes these params to a swagger request
func (o *AppUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.App != nil {
		if err := r.SetBodyParam(o.App); err != nil {
			return err
		}
	}

	// path param app-slug
	if err := r.SetPathParam("app-slug", o.AppSlug); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
