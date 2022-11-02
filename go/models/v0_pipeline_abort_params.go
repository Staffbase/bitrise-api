// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V0PipelineAbortParams v0 pipeline abort params
//
// swagger:model v0.PipelineAbortParams
type V0PipelineAbortParams struct {

	// abort reason
	// Required: true
	AbortReason *string `json:"abort_reason"`

	// abort with success
	// Required: true
	AbortWithSuccess *bool `json:"abort_with_success"`

	// skip notifications
	// Required: true
	SkipNotifications *bool `json:"skip_notifications"`
}

// Validate validates this v0 pipeline abort params
func (m *V0PipelineAbortParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAbortReason(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAbortWithSuccess(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSkipNotifications(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V0PipelineAbortParams) validateAbortReason(formats strfmt.Registry) error {

	if err := validate.Required("abort_reason", "body", m.AbortReason); err != nil {
		return err
	}

	return nil
}

func (m *V0PipelineAbortParams) validateAbortWithSuccess(formats strfmt.Registry) error {

	if err := validate.Required("abort_with_success", "body", m.AbortWithSuccess); err != nil {
		return err
	}

	return nil
}

func (m *V0PipelineAbortParams) validateSkipNotifications(formats strfmt.Registry) error {

	if err := validate.Required("skip_notifications", "body", m.SkipNotifications); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this v0 pipeline abort params based on context it is used
func (m *V0PipelineAbortParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V0PipelineAbortParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V0PipelineAbortParams) UnmarshalBinary(b []byte) error {
	var res V0PipelineAbortParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
