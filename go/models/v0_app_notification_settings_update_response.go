// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V0AppNotificationSettingsUpdateResponse v0 app notification settings update response
//
// swagger:model v0.AppNotificationSettingsUpdateResponse
type V0AppNotificationSettingsUpdateResponse struct {

	// msg
	Msg string `json:"msg,omitempty"`
}

// Validate validates this v0 app notification settings update response
func (m *V0AppNotificationSettingsUpdateResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v0 app notification settings update response based on context it is used
func (m *V0AppNotificationSettingsUpdateResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V0AppNotificationSettingsUpdateResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V0AppNotificationSettingsUpdateResponse) UnmarshalBinary(b []byte) error {
	var res V0AppNotificationSettingsUpdateResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
