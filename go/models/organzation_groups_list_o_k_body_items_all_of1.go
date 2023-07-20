// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// OrganzationGroupsListOKBodyItemsAllOf1 organzation groups list o k body items all of1
//
// swagger:model organzationGroupsListOKBodyItemsAllOf1
type OrganzationGroupsListOKBodyItemsAllOf1 struct {

	// name
	Name string `json:"name,omitempty"`

	// slug
	Slug string `json:"slug,omitempty"`
}

// Validate validates this organzation groups list o k body items all of1
func (m *OrganzationGroupsListOKBodyItemsAllOf1) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this organzation groups list o k body items all of1 based on context it is used
func (m *OrganzationGroupsListOKBodyItemsAllOf1) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OrganzationGroupsListOKBodyItemsAllOf1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OrganzationGroupsListOKBodyItemsAllOf1) UnmarshalBinary(b []byte) error {
	var res OrganzationGroupsListOKBodyItemsAllOf1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}