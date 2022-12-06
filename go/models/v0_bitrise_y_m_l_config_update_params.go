// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V0BitriseYMLConfigUpdateParams v0 bitrise y m l config update params
//
// swagger:model v0.BitriseYMLConfigUpdateParams
type V0BitriseYMLConfigUpdateParams struct {

	// Location of bitrise.yml file. Enums(bitrise.io, repository)
	// Example: repository
	Location string `json:"location,omitempty"`
}

// Validate validates this v0 bitrise y m l config update params
func (m *V0BitriseYMLConfigUpdateParams) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v0 bitrise y m l config update params based on context it is used
func (m *V0BitriseYMLConfigUpdateParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V0BitriseYMLConfigUpdateParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V0BitriseYMLConfigUpdateParams) UnmarshalBinary(b []byte) error {
	var res V0BitriseYMLConfigUpdateParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
