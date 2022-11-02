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

// V0OwnerAddOnResponseItemModel v0 owner add on response item model
//
// swagger:model v0.OwnerAddOnResponseItemModel
type V0OwnerAddOnResponseItemModel struct {

	// apps
	Apps []*V0AddOnAppResponseItemModel `json:"apps"`

	// documentation url
	DocumentationURL string `json:"documentation_url,omitempty"`

	// has ui
	HasUI bool `json:"has_ui,omitempty"`

	// icon
	Icon string `json:"icon,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// is beta
	IsBeta bool `json:"is_beta,omitempty"`

	// summary
	Summary string `json:"summary,omitempty"`

	// title
	Title string `json:"title,omitempty"`
}

// Validate validates this v0 owner add on response item model
func (m *V0OwnerAddOnResponseItemModel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateApps(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V0OwnerAddOnResponseItemModel) validateApps(formats strfmt.Registry) error {
	if swag.IsZero(m.Apps) { // not required
		return nil
	}

	for i := 0; i < len(m.Apps); i++ {
		if swag.IsZero(m.Apps[i]) { // not required
			continue
		}

		if m.Apps[i] != nil {
			if err := m.Apps[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("apps" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("apps" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this v0 owner add on response item model based on the context it is used
func (m *V0OwnerAddOnResponseItemModel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateApps(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V0OwnerAddOnResponseItemModel) contextValidateApps(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Apps); i++ {

		if m.Apps[i] != nil {
			if err := m.Apps[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("apps" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("apps" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V0OwnerAddOnResponseItemModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V0OwnerAddOnResponseItemModel) UnmarshalBinary(b []byte) error {
	var res V0OwnerAddOnResponseItemModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
