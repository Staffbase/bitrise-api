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

// V0AndroidKeystoreFileUploadParams v0 android keystore file upload params
//
// swagger:model v0.AndroidKeystoreFileUploadParams
type V0AndroidKeystoreFileUploadParams struct {

	// alias
	// Required: true
	Alias *string `json:"alias"`

	// keystore file name
	KeystoreFileName string `json:"keystore_file_name,omitempty"`

	// password
	// Required: true
	Password *string `json:"password"`

	// private key password
	// Required: true
	PrivateKeyPassword *string `json:"private_key_password"`

	// upload file name
	// Required: true
	UploadFileName *string `json:"upload_file_name"`

	// upload file size
	// Required: true
	UploadFileSize *int64 `json:"upload_file_size"`
}

// Validate validates this v0 android keystore file upload params
func (m *V0AndroidKeystoreFileUploadParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAlias(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePassword(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrivateKeyPassword(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUploadFileName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUploadFileSize(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V0AndroidKeystoreFileUploadParams) validateAlias(formats strfmt.Registry) error {

	if err := validate.Required("alias", "body", m.Alias); err != nil {
		return err
	}

	return nil
}

func (m *V0AndroidKeystoreFileUploadParams) validatePassword(formats strfmt.Registry) error {

	if err := validate.Required("password", "body", m.Password); err != nil {
		return err
	}

	return nil
}

func (m *V0AndroidKeystoreFileUploadParams) validatePrivateKeyPassword(formats strfmt.Registry) error {

	if err := validate.Required("private_key_password", "body", m.PrivateKeyPassword); err != nil {
		return err
	}

	return nil
}

func (m *V0AndroidKeystoreFileUploadParams) validateUploadFileName(formats strfmt.Registry) error {

	if err := validate.Required("upload_file_name", "body", m.UploadFileName); err != nil {
		return err
	}

	return nil
}

func (m *V0AndroidKeystoreFileUploadParams) validateUploadFileSize(formats strfmt.Registry) error {

	if err := validate.Required("upload_file_size", "body", m.UploadFileSize); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this v0 android keystore file upload params based on context it is used
func (m *V0AndroidKeystoreFileUploadParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V0AndroidKeystoreFileUploadParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V0AndroidKeystoreFileUploadParams) UnmarshalBinary(b []byte) error {
	var res V0AndroidKeystoreFileUploadParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
