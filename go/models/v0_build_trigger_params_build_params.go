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

// V0BuildTriggerParamsBuildParams v0 build trigger params build params
//
// swagger:model v0.BuildTriggerParamsBuildParams
type V0BuildTriggerParamsBuildParams struct {

	// base repository url
	BaseRepositoryURL string `json:"base_repository_url,omitempty"`

	// branch
	Branch string `json:"branch,omitempty"`

	// branch dest
	BranchDest string `json:"branch_dest,omitempty"`

	// branch dest repo owner
	BranchDestRepoOwner string `json:"branch_dest_repo_owner,omitempty"`

	// branch repo owner
	BranchRepoOwner string `json:"branch_repo_owner,omitempty"`

	// build request slug
	BuildRequestSlug string `json:"build_request_slug,omitempty"`

	// commit hash
	CommitHash string `json:"commit_hash,omitempty"`

	// commit message
	CommitMessage string `json:"commit_message,omitempty"`

	// commit paths
	CommitPaths []*V0CommitPaths `json:"commit_paths"`

	// diff url
	DiffURL string `json:"diff_url,omitempty"`

	// environments
	Environments []*V0BuildParamsEnvironment `json:"environments"`

	// head repository url
	HeadRepositoryURL string `json:"head_repository_url,omitempty"`

	// pipeline id
	PipelineID string `json:"pipeline_id,omitempty"`

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

	// skip git status report
	SkipGitStatusReport bool `json:"skip_git_status_report,omitempty"`

	// tag
	Tag string `json:"tag,omitempty"`

	// workflow id
	WorkflowID string `json:"workflow_id,omitempty"`
}

// Validate validates this v0 build trigger params build params
func (m *V0BuildTriggerParamsBuildParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCommitPaths(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnvironments(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V0BuildTriggerParamsBuildParams) validateCommitPaths(formats strfmt.Registry) error {
	if swag.IsZero(m.CommitPaths) { // not required
		return nil
	}

	for i := 0; i < len(m.CommitPaths); i++ {
		if swag.IsZero(m.CommitPaths[i]) { // not required
			continue
		}

		if m.CommitPaths[i] != nil {
			if err := m.CommitPaths[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("commit_paths" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("commit_paths" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V0BuildTriggerParamsBuildParams) validateEnvironments(formats strfmt.Registry) error {
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

// ContextValidate validate this v0 build trigger params build params based on the context it is used
func (m *V0BuildTriggerParamsBuildParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCommitPaths(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEnvironments(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V0BuildTriggerParamsBuildParams) contextValidateCommitPaths(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.CommitPaths); i++ {

		if m.CommitPaths[i] != nil {

			if swag.IsZero(m.CommitPaths[i]) { // not required
				return nil
			}

			if err := m.CommitPaths[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("commit_paths" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("commit_paths" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V0BuildTriggerParamsBuildParams) contextValidateEnvironments(ctx context.Context, formats strfmt.Registry) error {

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
func (m *V0BuildTriggerParamsBuildParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V0BuildTriggerParamsBuildParams) UnmarshalBinary(b []byte) error {
	var res V0BuildTriggerParamsBuildParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
