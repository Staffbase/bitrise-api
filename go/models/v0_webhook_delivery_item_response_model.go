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

// V0WebhookDeliveryItemResponseModel v0 webhook delivery item response model
//
// swagger:model v0.WebhookDeliveryItemResponseModel
type V0WebhookDeliveryItemResponseModel struct {

	// created at
	CreatedAt string `json:"created_at,omitempty"`

	// request headers
	RequestHeaders string `json:"request_headers,omitempty"`

	// request payload
	RequestPayload string `json:"request_payload,omitempty"`

	// request url
	RequestURL string `json:"request_url,omitempty"`

	// response body
	ResponseBody *NullsString `json:"response_body,omitempty"`

	// response headers
	ResponseHeaders *NullsString `json:"response_headers,omitempty"`

	// response http status
	ResponseHTTPStatus NullsInt64 `json:"response_http_status,omitempty"`

	// response seconds
	ResponseSeconds NullsInt64 `json:"response_seconds,omitempty"`

	// slug
	Slug string `json:"slug,omitempty"`

	// updated at
	UpdatedAt string `json:"updated_at,omitempty"`
}

// Validate validates this v0 webhook delivery item response model
func (m *V0WebhookDeliveryItemResponseModel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResponseBody(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResponseHeaders(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V0WebhookDeliveryItemResponseModel) validateResponseBody(formats strfmt.Registry) error {
	if swag.IsZero(m.ResponseBody) { // not required
		return nil
	}

	if m.ResponseBody != nil {
		if err := m.ResponseBody.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("response_body")
			}
			return err
		}
	}

	return nil
}

func (m *V0WebhookDeliveryItemResponseModel) validateResponseHeaders(formats strfmt.Registry) error {
	if swag.IsZero(m.ResponseHeaders) { // not required
		return nil
	}

	if m.ResponseHeaders != nil {
		if err := m.ResponseHeaders.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("response_headers")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v0 webhook delivery item response model based on the context it is used
func (m *V0WebhookDeliveryItemResponseModel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateResponseBody(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateResponseHeaders(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V0WebhookDeliveryItemResponseModel) contextValidateResponseBody(ctx context.Context, formats strfmt.Registry) error {

	if m.ResponseBody != nil {
		if err := m.ResponseBody.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("response_body")
			}
			return err
		}
	}

	return nil
}

func (m *V0WebhookDeliveryItemResponseModel) contextValidateResponseHeaders(ctx context.Context, formats strfmt.Registry) error {

	if m.ResponseHeaders != nil {
		if err := m.ResponseHeaders.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("response_headers")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V0WebhookDeliveryItemResponseModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V0WebhookDeliveryItemResponseModel) UnmarshalBinary(b []byte) error {
	var res V0WebhookDeliveryItemResponseModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
