// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CustomBitriseYmlModel custom bitrise yml model
//
// swagger:model custom.BitriseYmlModel
type CustomBitriseYmlModel struct {

	// app
	App *CustomBitriseYmlModelApp `json:"app,omitempty"`
}

// Validate validates this custom bitrise yml model
func (m *CustomBitriseYmlModel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateApp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CustomBitriseYmlModel) validateApp(formats strfmt.Registry) error {
	if swag.IsZero(m.App) { // not required
		return nil
	}

	if m.App != nil {
		if err := m.App.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("app")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this custom bitrise yml model based on the context it is used
func (m *CustomBitriseYmlModel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateApp(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CustomBitriseYmlModel) contextValidateApp(ctx context.Context, formats strfmt.Registry) error {

	if m.App != nil {
		if err := m.App.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("app")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CustomBitriseYmlModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CustomBitriseYmlModel) UnmarshalBinary(b []byte) error {
	var res CustomBitriseYmlModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
