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

// PipelineabledomainTriggerParams pipelineabledomain trigger params
//
// swagger:model pipelineabledomain.TriggerParams
type PipelineabledomainTriggerParams struct {

	// branch
	Branch string `json:"branch,omitempty"`

	// branch dest
	BranchDest string `json:"branch_dest,omitempty"`

	// branch dest repo owner
	BranchDestRepoOwner string `json:"branch_dest_repo_owner,omitempty"`

	// branch repo owner
	BranchRepoOwner string `json:"branch_repo_owner,omitempty"`

	// commit hash
	CommitHash string `json:"commit_hash,omitempty"`

	// commit message
	CommitMessage string `json:"commit_message,omitempty"`

	// environments
	Environments []*PipelineabledomainEnvironments `json:"environments"`

	// pull request author
	PullRequestAuthor string `json:"pull_request_author,omitempty"`

	// pull request head branch
	PullRequestHeadBranch string `json:"pull_request_head_branch,omitempty"`

	// pull request id
	PullRequestID interface{} `json:"pull_request_id,omitempty"`

	// pull request merge branch
	PullRequestMergeBranch string `json:"pull_request_merge_branch,omitempty"`

	// pull request repository url
	PullRequestRepositoryURL string `json:"pull_request_repository_url,omitempty"`

	// pull request unverified merge branch
	PullRequestUnverifiedMergeBranch string `json:"pull_request_unverified_merge_branch,omitempty"`

	// tag
	Tag interface{} `json:"tag,omitempty"`
}

// Validate validates this pipelineabledomain trigger params
func (m *PipelineabledomainTriggerParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnvironments(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PipelineabledomainTriggerParams) validateEnvironments(formats strfmt.Registry) error {
	if swag.IsZero(m.Environments) { // not required
		return nil
	}

	for i := 0; i < len(m.Environments); i++ {
		if swag.IsZero(m.Environments[i]) { // not required
			continue
		}

		if m.Environments[i] != nil {
			if err := m.Environments[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("environments" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("environments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this pipelineabledomain trigger params based on the context it is used
func (m *PipelineabledomainTriggerParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEnvironments(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PipelineabledomainTriggerParams) contextValidateEnvironments(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Environments); i++ {

		if m.Environments[i] != nil {

			if swag.IsZero(m.Environments[i]) { // not required
				return nil
			}

			if err := m.Environments[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("environments" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("environments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PipelineabledomainTriggerParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PipelineabledomainTriggerParams) UnmarshalBinary(b []byte) error {
	var res PipelineabledomainTriggerParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
