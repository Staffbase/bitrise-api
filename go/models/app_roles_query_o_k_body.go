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

// AppRolesQueryOKBody app roles query o k body
//
// swagger:model appRolesQueryOKBody
type AppRolesQueryOKBody struct {
	AppRolesQueryOKBodyAllOf0

	AppRolesQueryOKBodyAllOf1
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *AppRolesQueryOKBody) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 AppRolesQueryOKBodyAllOf0
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.AppRolesQueryOKBodyAllOf0 = aO0

	// AO1
	var aO1 AppRolesQueryOKBodyAllOf1
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.AppRolesQueryOKBodyAllOf1 = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m AppRolesQueryOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.AppRolesQueryOKBodyAllOf0)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.AppRolesQueryOKBodyAllOf1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this app roles query o k body
func (m *AppRolesQueryOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with AppRolesQueryOKBodyAllOf0
	// validation for a type composition with AppRolesQueryOKBodyAllOf1
	if err := m.AppRolesQueryOKBodyAllOf1.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this app roles query o k body based on the context it is used
func (m *AppRolesQueryOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with AppRolesQueryOKBodyAllOf0
	// validation for a type composition with AppRolesQueryOKBodyAllOf1
	if err := m.AppRolesQueryOKBodyAllOf1.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *AppRolesQueryOKBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AppRolesQueryOKBody) UnmarshalBinary(b []byte) error {
	var res AppRolesQueryOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// AppRolesQueryOKBodyAllOf0 app roles query o k body all of0
//
// swagger:model AppRolesQueryOKBodyAllOf0
type AppRolesQueryOKBodyAllOf0 interface{}
