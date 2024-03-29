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

// V0AppUploadParams v0 app upload params
//
// swagger:model v0.AppUploadParams
type V0AppUploadParams struct {

	// The default branch of the repository. If it's not specified, it will be `master`.
	DefaultBranchName string `json:"default_branch_name,omitempty"`

	// [Deprecated] You no longer need to provide this field.
	GitOwner string `json:"git_owner,omitempty"`

	// [Deprecated] You no longer need to provide this field.
	GitRepoSlug string `json:"git_repo_slug,omitempty"`

	// If `true` then the repository visibility setting will be public, in case of `false` it will be private
	// Required: true
	IsPublic *bool `json:"is_public"`

	// Toggles whether manual approval should be enabled for the app's builds. If it's not specified, it will be `true`.
	ManualApprovalEnabled bool `json:"manual_approval_enabled,omitempty"`

	// The slug of the organization, who will be the owner of the application. If it's not specified, then the authenticated user will be the owner.
	OrganizationSlug string `json:"organization_slug,omitempty"`

	// The git provider you are using, it can be `github`, `bitbucket`, `gitlab`, `gitlab-self-hosted` or `custom`
	Provider string `json:"provider,omitempty"`

	// The URL of your repository
	// Required: true
	RepoURL *string `json:"repo_url"`

	// The title of the application. If it's not specified, it will be the git repository's name.
	Title string `json:"title,omitempty"`

	// [Deprecated] You no longer need to provide this field.
	Type string `json:"type,omitempty"`
}

// Validate validates this v0 app upload params
func (m *V0AppUploadParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIsPublic(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRepoURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V0AppUploadParams) validateIsPublic(formats strfmt.Registry) error {

	if err := validate.Required("is_public", "body", m.IsPublic); err != nil {
		return err
	}

	return nil
}

func (m *V0AppUploadParams) validateRepoURL(formats strfmt.Registry) error {

	if err := validate.Required("repo_url", "body", m.RepoURL); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this v0 app upload params based on context it is used
func (m *V0AppUploadParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V0AppUploadParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V0AppUploadParams) UnmarshalBinary(b []byte) error {
	var res V0AppUploadParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
