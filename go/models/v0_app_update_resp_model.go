// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V0AppUpdateRespModel v0 app update resp model
//
// swagger:model v0.AppUpdateRespModel
type V0AppUpdateRespModel struct {

	// status
	Status string `json:"status,omitempty"`
}

// Validate validates this v0 app update resp model
func (m *V0AppUpdateRespModel) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v0 app update resp model based on context it is used
func (m *V0AppUpdateRespModel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V0AppUpdateRespModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V0AppUpdateRespModel) UnmarshalBinary(b []byte) error {
	var res V0AppUpdateRespModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
