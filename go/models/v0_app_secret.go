// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V0AppSecret v0 app secret
//
// swagger:model v0.AppSecret
type V0AppSecret struct {

	// Replace variable in inputs? https://devcenter.bitrise.io/en/references/steps-reference/step-inputs-reference.html#step-input-properties
	ExpandInStepInputs bool `json:"expand_in_step_inputs,omitempty"`

	// Expose for Pull Requests?
	IsExposedForPullRequests bool `json:"is_exposed_for_pull_requests,omitempty"`

	// Secret marked as protected?
	IsProtected bool `json:"is_protected,omitempty"`

	// Name of the secret
	Name string `json:"name,omitempty"`
}

// Validate validates this v0 app secret
func (m *V0AppSecret) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v0 app secret based on context it is used
func (m *V0AppSecret) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V0AppSecret) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V0AppSecret) UnmarshalBinary(b []byte) error {
	var res V0AppSecret
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}