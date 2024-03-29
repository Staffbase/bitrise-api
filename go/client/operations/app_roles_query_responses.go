// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/Staffbase/bitrise-api/go/models"
)

// AppRolesQueryReader is a Reader for the AppRolesQuery structure.
type AppRolesQueryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AppRolesQueryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAppRolesQueryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewAppRolesQueryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /apps/{app-slug}/roles/{role-name}] app-roles-query", response, response.Code())
	}
}

// NewAppRolesQueryOK creates a AppRolesQueryOK with default headers values
func NewAppRolesQueryOK() *AppRolesQueryOK {
	return &AppRolesQueryOK{}
}

/*
AppRolesQueryOK describes a response with status code 200, with default header values.

List of group slugs
*/
type AppRolesQueryOK struct {
	Payload *AppRolesQueryOKBody
}

// IsSuccess returns true when this app roles query o k response has a 2xx status code
func (o *AppRolesQueryOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this app roles query o k response has a 3xx status code
func (o *AppRolesQueryOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this app roles query o k response has a 4xx status code
func (o *AppRolesQueryOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this app roles query o k response has a 5xx status code
func (o *AppRolesQueryOK) IsServerError() bool {
	return false
}

// IsCode returns true when this app roles query o k response a status code equal to that given
func (o *AppRolesQueryOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the app roles query o k response
func (o *AppRolesQueryOK) Code() int {
	return 200
}

func (o *AppRolesQueryOK) Error() string {
	return fmt.Sprintf("[GET /apps/{app-slug}/roles/{role-name}][%d] appRolesQueryOK  %+v", 200, o.Payload)
}

func (o *AppRolesQueryOK) String() string {
	return fmt.Sprintf("[GET /apps/{app-slug}/roles/{role-name}][%d] appRolesQueryOK  %+v", 200, o.Payload)
}

func (o *AppRolesQueryOK) GetPayload() *AppRolesQueryOKBody {
	return o.Payload
}

func (o *AppRolesQueryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(AppRolesQueryOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAppRolesQueryNotFound creates a AppRolesQueryNotFound with default headers values
func NewAppRolesQueryNotFound() *AppRolesQueryNotFound {
	return &AppRolesQueryNotFound{}
}

/*
AppRolesQueryNotFound describes a response with status code 404, with default header values.

Not Found
*/
type AppRolesQueryNotFound struct {
	Payload *models.ServiceProxyErrorRespModel
}

// IsSuccess returns true when this app roles query not found response has a 2xx status code
func (o *AppRolesQueryNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this app roles query not found response has a 3xx status code
func (o *AppRolesQueryNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this app roles query not found response has a 4xx status code
func (o *AppRolesQueryNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this app roles query not found response has a 5xx status code
func (o *AppRolesQueryNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this app roles query not found response a status code equal to that given
func (o *AppRolesQueryNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the app roles query not found response
func (o *AppRolesQueryNotFound) Code() int {
	return 404
}

func (o *AppRolesQueryNotFound) Error() string {
	return fmt.Sprintf("[GET /apps/{app-slug}/roles/{role-name}][%d] appRolesQueryNotFound  %+v", 404, o.Payload)
}

func (o *AppRolesQueryNotFound) String() string {
	return fmt.Sprintf("[GET /apps/{app-slug}/roles/{role-name}][%d] appRolesQueryNotFound  %+v", 404, o.Payload)
}

func (o *AppRolesQueryNotFound) GetPayload() *models.ServiceProxyErrorRespModel {
	return o.Payload
}

func (o *AppRolesQueryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceProxyErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
AppRolesQueryOKBody app roles query o k body
swagger:model AppRolesQueryOKBody
*/
type AppRolesQueryOKBody struct {
	AppRolesQueryOKBodyAllOf0

	// groups
	Groups []string `json:"groups"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *AppRolesQueryOKBody) UnmarshalJSON(raw []byte) error {
	// AppRolesQueryOKBodyAO0
	var appRolesQueryOKBodyAO0 AppRolesQueryOKBodyAllOf0
	if err := swag.ReadJSON(raw, &appRolesQueryOKBodyAO0); err != nil {
		return err
	}
	o.AppRolesQueryOKBodyAllOf0 = appRolesQueryOKBodyAO0

	// AppRolesQueryOKBodyAO1
	var dataAppRolesQueryOKBodyAO1 struct {
		Groups []string `json:"groups"`
	}
	if err := swag.ReadJSON(raw, &dataAppRolesQueryOKBodyAO1); err != nil {
		return err
	}

	o.Groups = dataAppRolesQueryOKBodyAO1.Groups

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o AppRolesQueryOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	appRolesQueryOKBodyAO0, err := swag.WriteJSON(o.AppRolesQueryOKBodyAllOf0)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, appRolesQueryOKBodyAO0)
	var dataAppRolesQueryOKBodyAO1 struct {
		Groups []string `json:"groups"`
	}

	dataAppRolesQueryOKBodyAO1.Groups = o.Groups

	jsonDataAppRolesQueryOKBodyAO1, errAppRolesQueryOKBodyAO1 := swag.WriteJSON(dataAppRolesQueryOKBodyAO1)
	if errAppRolesQueryOKBodyAO1 != nil {
		return nil, errAppRolesQueryOKBodyAO1
	}
	_parts = append(_parts, jsonDataAppRolesQueryOKBodyAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this app roles query o k body
func (o *AppRolesQueryOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with AppRolesQueryOKBodyAllOf0

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this app roles query o k body based on context it is used
func (o *AppRolesQueryOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AppRolesQueryOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AppRolesQueryOKBody) UnmarshalBinary(b []byte) error {
	var res AppRolesQueryOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
AppRolesQueryOKBodyAllOf0 app roles query o k body all of0
swagger:model AppRolesQueryOKBodyAllOf0
*/
type AppRolesQueryOKBodyAllOf0 interface{}
