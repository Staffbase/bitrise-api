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

// V0PlanDataModel v0 plan data model
//
// swagger:model v0.PlanDataModel
type V0PlanDataModel struct {

	// container count
	ContainerCount int64 `json:"container_count,omitempty"`

	// expires at
	ExpiresAt string `json:"expires_at,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// price
	Price *NullsInt64 `json:"price,omitempty"`
}

// Validate validates this v0 plan data model
func (m *V0PlanDataModel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePrice(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V0PlanDataModel) validatePrice(formats strfmt.Registry) error {
	if swag.IsZero(m.Price) { // not required
		return nil
	}

	if m.Price != nil {
		if err := m.Price.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("price")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("price")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v0 plan data model based on the context it is used
func (m *V0PlanDataModel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePrice(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V0PlanDataModel) contextValidatePrice(ctx context.Context, formats strfmt.Registry) error {

	if m.Price != nil {

		if swag.IsZero(m.Price) { // not required
			return nil
		}

		if err := m.Price.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("price")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("price")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V0PlanDataModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V0PlanDataModel) UnmarshalBinary(b []byte) error {
	var res V0PlanDataModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
