// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V0AppUpdateParams v0 app update params
//
// swagger:model v0.AppUpdateParams
type V0AppUpdateParams struct {

	// The new apple credential user ID (recommendation: use the UI to set this)
	AppleCredentialUserID int64 `json:"apple_credential_user_id,omitempty"`

	// The new apple credential user slug (recommendation: use the UI to set this)
	AppleCredentialUserSlug string `json:"apple_credential_user_slug,omitempty"`

	// The new default branch for the application.
	DefaultBranch string `json:"default_branch,omitempty"`

	// The new value of whether an application should be publicly visible
	IsPublic bool `json:"is_public,omitempty"`

	// The new repository URL for the application.
	RepositoryURL string `json:"repository_url,omitempty"`

	// The new service credential user ID (recommendation: use the UI to set this)
	ServicesCredentialUserID int64 `json:"services_credential_user_id,omitempty"`

	// The new title of the application.
	Title string `json:"title,omitempty"`
}

// Validate validates this v0 app update params
func (m *V0AppUpdateParams) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v0 app update params based on context it is used
func (m *V0AppUpdateParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V0AppUpdateParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V0AppUpdateParams) UnmarshalBinary(b []byte) error {
	var res V0AppUpdateParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
