// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V0PagingResponseModel v0 paging response model
//
// swagger:model v0.PagingResponseModel
type V0PagingResponseModel struct {

	// Next is the "anchor" for pagination. This value should be passed to the same endpoint
	// to get the next page. Empty/not included if there's no "next" page.
	// Stop paging when there's no "Next" item in the response!
	Next string `json:"next,omitempty"`

	// PageItemLimit - per-page item count. A given page might include
	// less items if there's not enough items. This value is the "max item count per page".
	PageItemLimit int64 `json:"page_item_limit,omitempty"`

	// TotalItemCount - total item count, through "all pages"
	TotalItemCount int64 `json:"total_item_count,omitempty"`
}

// Validate validates this v0 paging response model
func (m *V0PagingResponseModel) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v0 paging response model based on context it is used
func (m *V0PagingResponseModel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V0PagingResponseModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V0PagingResponseModel) UnmarshalBinary(b []byte) error {
	var res V0PagingResponseModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
