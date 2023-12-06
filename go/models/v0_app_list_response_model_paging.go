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

// V0AppListResponseModelPaging pagination
//
// swagger:model v0AppListResponseModelPaging
type V0AppListResponseModelPaging struct {
	V0PagingResponseModel
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *V0AppListResponseModelPaging) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 V0PagingResponseModel
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.V0PagingResponseModel = aO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m V0AppListResponseModelPaging) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.V0PagingResponseModel)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this v0 app list response model paging
func (m *V0AppListResponseModelPaging) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with V0PagingResponseModel
	if err := m.V0PagingResponseModel.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this v0 app list response model paging based on the context it is used
func (m *V0AppListResponseModelPaging) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with V0PagingResponseModel
	if err := m.V0PagingResponseModel.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *V0AppListResponseModelPaging) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V0AppListResponseModelPaging) UnmarshalBinary(b []byte) error {
	var res V0AppListResponseModelPaging
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
