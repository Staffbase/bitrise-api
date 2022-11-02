// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V0BuildResponseItemModel v0 build response item model
//
// swagger:model v0.BuildResponseItemModel
type V0BuildResponseItemModel struct {

	// abort reason
	AbortReason *GithubComMarkbatesPopNullsString `json:"abort_reason,omitempty"`

	// branch
	Branch *GithubComMarkbatesPopNullsString `json:"branch,omitempty"`

	// build number
	BuildNumber int64 `json:"build_number,omitempty"`

	// commit hash
	CommitHash *GithubComMarkbatesPopNullsString `json:"commit_hash,omitempty"`

	// commit message
	CommitMessage *GithubComMarkbatesPopNullsString `json:"commit_message,omitempty"`

	// commit view url
	CommitViewURL *GithubComMarkbatesPopNullsString `json:"commit_view_url,omitempty"`

	// credit cost
	CreditCost *GithubComMarkbatesPopNullsInt64 `json:"credit_cost,omitempty"`

	// environment prepare finished at
	EnvironmentPrepareFinishedAt string `json:"environment_prepare_finished_at,omitempty"`

	// finished at
	FinishedAt string `json:"finished_at,omitempty"`

	// is on hold
	IsOnHold bool `json:"is_on_hold,omitempty"`

	// is processed
	IsProcessed bool `json:"is_processed,omitempty"`

	// is status sent
	IsStatusSent bool `json:"is_status_sent,omitempty"`

	// machine type id
	MachineTypeID *GithubComMarkbatesPopNullsString `json:"machine_type_id,omitempty"`

	// original build params
	OriginalBuildParams []int64 `json:"original_build_params"`

	// pipeline workflow id
	PipelineWorkflowID string `json:"pipeline_workflow_id,omitempty"`

	// pull request id
	PullRequestID int64 `json:"pull_request_id,omitempty"`

	// pull request target branch
	PullRequestTargetBranch *GithubComMarkbatesPopNullsString `json:"pull_request_target_branch,omitempty"`

	// pull request view url
	PullRequestViewURL *GithubComMarkbatesPopNullsString `json:"pull_request_view_url,omitempty"`

	// slug
	Slug string `json:"slug,omitempty"`

	// stack identifier
	StackIdentifier *GithubComMarkbatesPopNullsString `json:"stack_identifier,omitempty"`

	// started on worker at
	StartedOnWorkerAt string `json:"started_on_worker_at,omitempty"`

	// status
	Status int64 `json:"status,omitempty"`

	// status text
	StatusText string `json:"status_text,omitempty"`

	// tag
	Tag *GithubComMarkbatesPopNullsString `json:"tag,omitempty"`

	// triggered at
	TriggeredAt string `json:"triggered_at,omitempty"`

	// triggered by
	TriggeredBy *GithubComMarkbatesPopNullsString `json:"triggered_by,omitempty"`

	// triggered workflow
	TriggeredWorkflow string `json:"triggered_workflow,omitempty"`
}

// Validate validates this v0 build response item model
func (m *V0BuildResponseItemModel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAbortReason(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBranch(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCommitHash(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCommitMessage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCommitViewURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreditCost(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMachineTypeID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePullRequestTargetBranch(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePullRequestViewURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStackIdentifier(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTag(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTriggeredBy(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V0BuildResponseItemModel) validateAbortReason(formats strfmt.Registry) error {
	if swag.IsZero(m.AbortReason) { // not required
		return nil
	}

	if m.AbortReason != nil {
		if err := m.AbortReason.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("abort_reason")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("abort_reason")
			}
			return err
		}
	}

	return nil
}

func (m *V0BuildResponseItemModel) validateBranch(formats strfmt.Registry) error {
	if swag.IsZero(m.Branch) { // not required
		return nil
	}

	if m.Branch != nil {
		if err := m.Branch.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("branch")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("branch")
			}
			return err
		}
	}

	return nil
}

func (m *V0BuildResponseItemModel) validateCommitHash(formats strfmt.Registry) error {
	if swag.IsZero(m.CommitHash) { // not required
		return nil
	}

	if m.CommitHash != nil {
		if err := m.CommitHash.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("commit_hash")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("commit_hash")
			}
			return err
		}
	}

	return nil
}

func (m *V0BuildResponseItemModel) validateCommitMessage(formats strfmt.Registry) error {
	if swag.IsZero(m.CommitMessage) { // not required
		return nil
	}

	if m.CommitMessage != nil {
		if err := m.CommitMessage.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("commit_message")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("commit_message")
			}
			return err
		}
	}

	return nil
}

func (m *V0BuildResponseItemModel) validateCommitViewURL(formats strfmt.Registry) error {
	if swag.IsZero(m.CommitViewURL) { // not required
		return nil
	}

	if m.CommitViewURL != nil {
		if err := m.CommitViewURL.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("commit_view_url")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("commit_view_url")
			}
			return err
		}
	}

	return nil
}

func (m *V0BuildResponseItemModel) validateCreditCost(formats strfmt.Registry) error {
	if swag.IsZero(m.CreditCost) { // not required
		return nil
	}

	if m.CreditCost != nil {
		if err := m.CreditCost.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("credit_cost")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("credit_cost")
			}
			return err
		}
	}

	return nil
}

func (m *V0BuildResponseItemModel) validateMachineTypeID(formats strfmt.Registry) error {
	if swag.IsZero(m.MachineTypeID) { // not required
		return nil
	}

	if m.MachineTypeID != nil {
		if err := m.MachineTypeID.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("machine_type_id")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("machine_type_id")
			}
			return err
		}
	}

	return nil
}

func (m *V0BuildResponseItemModel) validatePullRequestTargetBranch(formats strfmt.Registry) error {
	if swag.IsZero(m.PullRequestTargetBranch) { // not required
		return nil
	}

	if m.PullRequestTargetBranch != nil {
		if err := m.PullRequestTargetBranch.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pull_request_target_branch")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pull_request_target_branch")
			}
			return err
		}
	}

	return nil
}

func (m *V0BuildResponseItemModel) validatePullRequestViewURL(formats strfmt.Registry) error {
	if swag.IsZero(m.PullRequestViewURL) { // not required
		return nil
	}

	if m.PullRequestViewURL != nil {
		if err := m.PullRequestViewURL.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pull_request_view_url")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pull_request_view_url")
			}
			return err
		}
	}

	return nil
}

func (m *V0BuildResponseItemModel) validateStackIdentifier(formats strfmt.Registry) error {
	if swag.IsZero(m.StackIdentifier) { // not required
		return nil
	}

	if m.StackIdentifier != nil {
		if err := m.StackIdentifier.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("stack_identifier")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("stack_identifier")
			}
			return err
		}
	}

	return nil
}

func (m *V0BuildResponseItemModel) validateTag(formats strfmt.Registry) error {
	if swag.IsZero(m.Tag) { // not required
		return nil
	}

	if m.Tag != nil {
		if err := m.Tag.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tag")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tag")
			}
			return err
		}
	}

	return nil
}

func (m *V0BuildResponseItemModel) validateTriggeredBy(formats strfmt.Registry) error {
	if swag.IsZero(m.TriggeredBy) { // not required
		return nil
	}

	if m.TriggeredBy != nil {
		if err := m.TriggeredBy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("triggered_by")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("triggered_by")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v0 build response item model based on the context it is used
func (m *V0BuildResponseItemModel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAbortReason(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateBranch(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCommitHash(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCommitMessage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCommitViewURL(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCreditCost(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMachineTypeID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePullRequestTargetBranch(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePullRequestViewURL(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStackIdentifier(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTag(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTriggeredBy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V0BuildResponseItemModel) contextValidateAbortReason(ctx context.Context, formats strfmt.Registry) error {

	if m.AbortReason != nil {
		if err := m.AbortReason.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("abort_reason")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("abort_reason")
			}
			return err
		}
	}

	return nil
}

func (m *V0BuildResponseItemModel) contextValidateBranch(ctx context.Context, formats strfmt.Registry) error {

	if m.Branch != nil {
		if err := m.Branch.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("branch")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("branch")
			}
			return err
		}
	}

	return nil
}

func (m *V0BuildResponseItemModel) contextValidateCommitHash(ctx context.Context, formats strfmt.Registry) error {

	if m.CommitHash != nil {
		if err := m.CommitHash.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("commit_hash")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("commit_hash")
			}
			return err
		}
	}

	return nil
}

func (m *V0BuildResponseItemModel) contextValidateCommitMessage(ctx context.Context, formats strfmt.Registry) error {

	if m.CommitMessage != nil {
		if err := m.CommitMessage.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("commit_message")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("commit_message")
			}
			return err
		}
	}

	return nil
}

func (m *V0BuildResponseItemModel) contextValidateCommitViewURL(ctx context.Context, formats strfmt.Registry) error {

	if m.CommitViewURL != nil {
		if err := m.CommitViewURL.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("commit_view_url")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("commit_view_url")
			}
			return err
		}
	}

	return nil
}

func (m *V0BuildResponseItemModel) contextValidateCreditCost(ctx context.Context, formats strfmt.Registry) error {

	if m.CreditCost != nil {
		if err := m.CreditCost.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("credit_cost")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("credit_cost")
			}
			return err
		}
	}

	return nil
}

func (m *V0BuildResponseItemModel) contextValidateMachineTypeID(ctx context.Context, formats strfmt.Registry) error {

	if m.MachineTypeID != nil {
		if err := m.MachineTypeID.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("machine_type_id")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("machine_type_id")
			}
			return err
		}
	}

	return nil
}

func (m *V0BuildResponseItemModel) contextValidatePullRequestTargetBranch(ctx context.Context, formats strfmt.Registry) error {

	if m.PullRequestTargetBranch != nil {
		if err := m.PullRequestTargetBranch.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pull_request_target_branch")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pull_request_target_branch")
			}
			return err
		}
	}

	return nil
}

func (m *V0BuildResponseItemModel) contextValidatePullRequestViewURL(ctx context.Context, formats strfmt.Registry) error {

	if m.PullRequestViewURL != nil {
		if err := m.PullRequestViewURL.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pull_request_view_url")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pull_request_view_url")
			}
			return err
		}
	}

	return nil
}

func (m *V0BuildResponseItemModel) contextValidateStackIdentifier(ctx context.Context, formats strfmt.Registry) error {

	if m.StackIdentifier != nil {
		if err := m.StackIdentifier.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("stack_identifier")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("stack_identifier")
			}
			return err
		}
	}

	return nil
}

func (m *V0BuildResponseItemModel) contextValidateTag(ctx context.Context, formats strfmt.Registry) error {

	if m.Tag != nil {
		if err := m.Tag.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tag")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tag")
			}
			return err
		}
	}

	return nil
}

func (m *V0BuildResponseItemModel) contextValidateTriggeredBy(ctx context.Context, formats strfmt.Registry) error {

	if m.TriggeredBy != nil {
		if err := m.TriggeredBy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("triggered_by")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("triggered_by")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V0BuildResponseItemModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V0BuildResponseItemModel) UnmarshalBinary(b []byte) error {
	var res V0BuildResponseItemModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
