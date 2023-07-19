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

// AppRolesUpdateParamsBody app roles update params body
//
// swagger:model appRolesUpdateParamsBody
type AppRolesUpdateParamsBody struct {
	AppRolesUpdateParamsBodyAllOf0

	AppRolesUpdateParamsBodyAllOf1
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *AppRolesUpdateParamsBody) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 AppRolesUpdateParamsBodyAllOf0
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.AppRolesUpdateParamsBodyAllOf0 = aO0

	// AO1
	var aO1 AppRolesUpdateParamsBodyAllOf1
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.AppRolesUpdateParamsBodyAllOf1 = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m AppRolesUpdateParamsBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.AppRolesUpdateParamsBodyAllOf0)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.AppRolesUpdateParamsBodyAllOf1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this app roles update params body
func (m *AppRolesUpdateParamsBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with AppRolesUpdateParamsBodyAllOf0
	// validation for a type composition with AppRolesUpdateParamsBodyAllOf1
	if err := m.AppRolesUpdateParamsBodyAllOf1.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this app roles update params body based on the context it is used
func (m *AppRolesUpdateParamsBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with AppRolesUpdateParamsBodyAllOf0
	// validation for a type composition with AppRolesUpdateParamsBodyAllOf1
	if err := m.AppRolesUpdateParamsBodyAllOf1.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *AppRolesUpdateParamsBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AppRolesUpdateParamsBody) UnmarshalBinary(b []byte) error {
	var res AppRolesUpdateParamsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// AppRolesUpdateParamsBodyAllOf0 app roles update params body all of0
//
// swagger:model AppRolesUpdateParamsBodyAllOf0
type AppRolesUpdateParamsBodyAllOf0 interface{}
