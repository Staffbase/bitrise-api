// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V0PipelineRebuildParams v0 pipeline rebuild params
//
// swagger:model v0.PipelineRebuildParams
type V0PipelineRebuildParams struct {

	// partial
	Partial bool `json:"partial,omitempty"`

	// triggered by
	TriggeredBy string `json:"triggered_by,omitempty"`
}

// Validate validates this v0 pipeline rebuild params
func (m *V0PipelineRebuildParams) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v0 pipeline rebuild params based on context it is used
func (m *V0PipelineRebuildParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V0PipelineRebuildParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V0PipelineRebuildParams) UnmarshalBinary(b []byte) error {
	var res V0PipelineRebuildParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}