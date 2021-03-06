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

// V0ProvisionProfileUploadParams v0 provision profile upload params
//
// swagger:model v0.ProvisionProfileUploadParams
type V0ProvisionProfileUploadParams struct {

	// upload file name
	// Required: true
	UploadFileName *string `json:"upload_file_name"`

	// upload file size
	// Required: true
	UploadFileSize *int64 `json:"upload_file_size"`
}

// Validate validates this v0 provision profile upload params
func (m *V0ProvisionProfileUploadParams) Validate(formats strfmt.Registry) error {
	var res []error

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

func (m *V0ProvisionProfileUploadParams) validateUploadFileName(formats strfmt.Registry) error {

	if err := validate.Required("upload_file_name", "body", m.UploadFileName); err != nil {
		return err
	}

	return nil
}

func (m *V0ProvisionProfileUploadParams) validateUploadFileSize(formats strfmt.Registry) error {

	if err := validate.Required("upload_file_size", "body", m.UploadFileSize); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this v0 provision profile upload params based on context it is used
func (m *V0ProvisionProfileUploadParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V0ProvisionProfileUploadParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V0ProvisionProfileUploadParams) UnmarshalBinary(b []byte) error {
	var res V0ProvisionProfileUploadParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
