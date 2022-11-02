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

	// The slug of the owner of the repository at the git provider
	// Required: true
	GitOwner *string `json:"git_owner"`

	// The slug of the repository at the git provider
	// Required: true
	GitRepoSlug *string `json:"git_repo_slug"`

	// If `true` then the repository visibility setting will be public, in case of `false` it will be private
	// Required: true
	IsPublic *bool `json:"is_public"`

	// The slug of the organization, who will be the owner of the application. If it's not specified, then the authenticated user will be the owner.
	OrganizationSlug string `json:"organization_slug,omitempty"`

	// The git provider you are using, it can be `github`, `bitbucket`, `gitlab`, `gitlab-self-hosted` or `custom`
	// Required: true
	Provider *string `json:"provider"`

	// The URL of your repository
	// Required: true
	RepoURL *string `json:"repo_url"`

	// The title of the application. If it's not specified, it will be the git repository's name.
	Title string `json:"title,omitempty"`

	// It has to be provided by legacy reasons and has to have the `git` value
	// Required: true
	Type *string `json:"type"`
}

// Validate validates this v0 app upload params
func (m *V0AppUploadParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGitOwner(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGitRepoSlug(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsPublic(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProvider(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRepoURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V0AppUploadParams) validateGitOwner(formats strfmt.Registry) error {

	if err := validate.Required("git_owner", "body", m.GitOwner); err != nil {
		return err
	}

	return nil
}

func (m *V0AppUploadParams) validateGitRepoSlug(formats strfmt.Registry) error {

	if err := validate.Required("git_repo_slug", "body", m.GitRepoSlug); err != nil {
		return err
	}

	return nil
}

func (m *V0AppUploadParams) validateIsPublic(formats strfmt.Registry) error {

	if err := validate.Required("is_public", "body", m.IsPublic); err != nil {
		return err
	}

	return nil
}

func (m *V0AppUploadParams) validateProvider(formats strfmt.Registry) error {

	if err := validate.Required("provider", "body", m.Provider); err != nil {
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

func (m *V0AppUploadParams) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
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
