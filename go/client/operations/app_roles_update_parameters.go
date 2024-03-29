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

// NewAppRolesUpdateParams creates a new AppRolesUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAppRolesUpdateParams() *AppRolesUpdateParams {
	return &AppRolesUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAppRolesUpdateParamsWithTimeout creates a new AppRolesUpdateParams object
// with the ability to set a timeout on a request.
func NewAppRolesUpdateParamsWithTimeout(timeout time.Duration) *AppRolesUpdateParams {
	return &AppRolesUpdateParams{
		timeout: timeout,
	}
}

// NewAppRolesUpdateParamsWithContext creates a new AppRolesUpdateParams object
// with the ability to set a context for a request.
func NewAppRolesUpdateParamsWithContext(ctx context.Context) *AppRolesUpdateParams {
	return &AppRolesUpdateParams{
		Context: ctx,
	}
}

// NewAppRolesUpdateParamsWithHTTPClient creates a new AppRolesUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewAppRolesUpdateParamsWithHTTPClient(client *http.Client) *AppRolesUpdateParams {
	return &AppRolesUpdateParams{
		HTTPClient: client,
	}
}

/*
AppRolesUpdateParams contains all the parameters to send to the API endpoint

	for the app roles update operation.

	Typically these are written to a http.Request.
*/
type AppRolesUpdateParams struct {

	/* AppSlug.

	   Slug of the app
	*/
	AppSlug string

	/* Groups.

	   List of group slugs
	*/
	Groups AppRolesUpdateBody

	/* RoleName.

	   Name of the role being modified, supported values are: admin, manager (equals developer), and member (equals tester/qa)
	*/
	RoleName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the app roles update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AppRolesUpdateParams) WithDefaults() *AppRolesUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the app roles update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AppRolesUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the app roles update params
func (o *AppRolesUpdateParams) WithTimeout(timeout time.Duration) *AppRolesUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the app roles update params
func (o *AppRolesUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the app roles update params
func (o *AppRolesUpdateParams) WithContext(ctx context.Context) *AppRolesUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the app roles update params
func (o *AppRolesUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the app roles update params
func (o *AppRolesUpdateParams) WithHTTPClient(client *http.Client) *AppRolesUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the app roles update params
func (o *AppRolesUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppSlug adds the appSlug to the app roles update params
func (o *AppRolesUpdateParams) WithAppSlug(appSlug string) *AppRolesUpdateParams {
	o.SetAppSlug(appSlug)
	return o
}

// SetAppSlug adds the appSlug to the app roles update params
func (o *AppRolesUpdateParams) SetAppSlug(appSlug string) {
	o.AppSlug = appSlug
}

// WithGroups adds the groups to the app roles update params
func (o *AppRolesUpdateParams) WithGroups(groups AppRolesUpdateBody) *AppRolesUpdateParams {
	o.SetGroups(groups)
	return o
}

// SetGroups adds the groups to the app roles update params
func (o *AppRolesUpdateParams) SetGroups(groups AppRolesUpdateBody) {
	o.Groups = groups
}

// WithRoleName adds the roleName to the app roles update params
func (o *AppRolesUpdateParams) WithRoleName(roleName string) *AppRolesUpdateParams {
	o.SetRoleName(roleName)
	return o
}

// SetRoleName adds the roleName to the app roles update params
func (o *AppRolesUpdateParams) SetRoleName(roleName string) {
	o.RoleName = roleName
}

// WriteToRequest writes these params to a swagger request
func (o *AppRolesUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param app-slug
	if err := r.SetPathParam("app-slug", o.AppSlug); err != nil {
		return err
	}
	if err := r.SetBodyParam(o.Groups); err != nil {
		return err
	}

	// path param role-name
	if err := r.SetPathParam("role-name", o.RoleName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
