// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PipelineabledomainBuildToolInvocation pipelineabledomain build tool invocation
//
// swagger:model pipelineabledomain.BuildToolInvocation
type PipelineabledomainBuildToolInvocation struct {

	// invocation id
	InvocationID string `json:"invocation_id,omitempty"`

	// tool
	Tool string `json:"tool,omitempty"`

	// tool version
	ToolVersion string `json:"tool_version,omitempty"`
}

// Validate validates this pipelineabledomain build tool invocation
func (m *PipelineabledomainBuildToolInvocation) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this pipelineabledomain build tool invocation based on context it is used
func (m *PipelineabledomainBuildToolInvocation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PipelineabledomainBuildToolInvocation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PipelineabledomainBuildToolInvocation) UnmarshalBinary(b []byte) error {
	var res PipelineabledomainBuildToolInvocation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
