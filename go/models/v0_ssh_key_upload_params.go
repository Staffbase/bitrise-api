// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V0SSHKeyUploadParams v0 SSH key upload params
//
// swagger:model v0.SSHKeyUploadParams
type V0SSHKeyUploadParams struct {

	// The private part of the SSH key you would like to use
	// Required: true
	AuthSSHPrivateKey *string `json:"auth_ssh_private_key"`

	// The public part of the SSH key you would like to use
	// Required: true
	AuthSSHPublicKey *string `json:"auth_ssh_public_key"`

	// If it's set to true, the provided SSH key will be registered at the provider of the application
	IsRegisterKeyIntoProviderService bool `json:"is_register_key_into_provider_service,omitempty"`
}

// Validate validates this v0 SSH key upload params
func (m *V0SSHKeyUploadParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthSSHPrivateKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAuthSSHPublicKey(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V0SSHKeyUploadParams) validateAuthSSHPrivateKey(formats strfmt.Registry) error {

	if err := validate.Required("auth_ssh_private_key", "body", m.AuthSSHPrivateKey); err != nil {
		return err
	}

	return nil
}

func (m *V0SSHKeyUploadParams) validateAuthSSHPublicKey(formats strfmt.Registry) error {

	if err := validate.Required("auth_ssh_public_key", "body", m.AuthSSHPublicKey); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this v0 SSH key upload params based on context it is used
func (m *V0SSHKeyUploadParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V0SSHKeyUploadParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V0SSHKeyUploadParams) UnmarshalBinary(b []byte) error {
	var res V0SSHKeyUploadParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
