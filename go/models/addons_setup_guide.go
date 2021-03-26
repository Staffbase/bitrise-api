// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AddonsSetupGuide addons setup guide
//
// swagger:model addons.SetupGuide
type AddonsSetupGuide struct {

	// instructions
	Instructions []*AddonsSetupInstruction `json:"instructions"`

	// notification
	Notification string `json:"notification,omitempty"`
}

// Validate validates this addons setup guide
func (m *AddonsSetupGuide) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInstructions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AddonsSetupGuide) validateInstructions(formats strfmt.Registry) error {
	if swag.IsZero(m.Instructions) { // not required
		return nil
	}

	for i := 0; i < len(m.Instructions); i++ {
		if swag.IsZero(m.Instructions[i]) { // not required
			continue
		}

		if m.Instructions[i] != nil {
			if err := m.Instructions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("instructions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this addons setup guide based on the context it is used
func (m *AddonsSetupGuide) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateInstructions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AddonsSetupGuide) contextValidateInstructions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Instructions); i++ {

		if m.Instructions[i] != nil {
			if err := m.Instructions[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("instructions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *AddonsSetupGuide) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AddonsSetupGuide) UnmarshalBinary(b []byte) error {
	var res AddonsSetupGuide
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
