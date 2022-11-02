// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V0CacheItemDownloadURLResponseModel v0 cache item download URL response model
//
// swagger:model v0.CacheItemDownloadURLResponseModel
type V0CacheItemDownloadURLResponseModel struct {

	// download url
	DownloadURL string `json:"download_url,omitempty"`
}

// Validate validates this v0 cache item download URL response model
func (m *V0CacheItemDownloadURLResponseModel) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v0 cache item download URL response model based on context it is used
func (m *V0CacheItemDownloadURLResponseModel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V0CacheItemDownloadURLResponseModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V0CacheItemDownloadURLResponseModel) UnmarshalBinary(b []byte) error {
	var res V0CacheItemDownloadURLResponseModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}