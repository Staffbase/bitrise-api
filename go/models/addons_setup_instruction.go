// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AddonsSetupInstruction addons setup instruction
//
// swagger:model addons.SetupInstruction
type AddonsSetupInstruction struct {

	// card content
	CardContent string `json:"card_content,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this addons setup instruction
func (m *AddonsSetupInstruction) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this addons setup instruction based on context it is used
func (m *AddonsSetupInstruction) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AddonsSetupInstruction) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AddonsSetupInstruction) UnmarshalBinary(b []byte) error {
	var res AddonsSetupInstruction
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
