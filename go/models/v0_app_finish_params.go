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

// V0AppFinishParams v0 app finish params
//
// swagger:model v0.AppFinishParams
type V0AppFinishParams struct {

	// Which config to use. Specify a config that matches your project type (e. g. `default-android-config` for `android`, etc.). If not speficied, default value is `other-config`. The available values are `default-android-config`, `default-cordova-config`, `default-fastlane-android-config`, `default-fastlane-ios-config`, `flutter-config-notest-app-android`, `flutter-config-notest-app-both`, `flutter-config-notest-app-ios`, `flutter-config-test-app-android`, `flutter-config-test-app-both`, `flutter-config-test-app-ios`, `default-ionic-config`, `default-ios-config`, `default-macos-config`, `default-react-native-config`, `default-react-native-expo-config`, `other-config`.
	Config string `json:"config,omitempty"`

	// Environment variables for the application workflows, e.g. {"env1":"val1","env2":"val2"}
	Envs map[string]string `json:"envs,omitempty"`

	// config specification mode, has to be specified with `manual` value
	Mode string `json:"mode,omitempty"`

	// The slug of the organization, who will be the owner of the application, if it's not specified, then the authenticated user will be the owner
	OrganizationSlug string `json:"organization_slug,omitempty"`

	// The type of your project (`android`, `ios`, `cordova`, `other`, `xamarin`, `macos`, `ionic`, `react-native`, `fastlane`, null)
	// Required: true
	ProjectType *string `json:"project_type"`

	// The id of the stack the application will be built (these can be found in the [system report](https://github.com/bitrise-io/bitrise.io/tree/master/system_reports) file names)
	// Required: true
	StackID *string `json:"stack_id"`
}

// Validate validates this v0 app finish params
func (m *V0AppFinishParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProjectType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStackID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V0AppFinishParams) validateProjectType(formats strfmt.Registry) error {

	if err := validate.Required("project_type", "body", m.ProjectType); err != nil {
		return err
	}

	return nil
}

func (m *V0AppFinishParams) validateStackID(formats strfmt.Registry) error {

	if err := validate.Required("stack_id", "body", m.StackID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this v0 app finish params based on context it is used
func (m *V0AppFinishParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V0AppFinishParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V0AppFinishParams) UnmarshalBinary(b []byte) error {
	var res V0AppFinishParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
