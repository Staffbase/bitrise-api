// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Staffbase/bitrise-api/go/models"
)

// AddonListByUserReader is a Reader for the AddonListByUser structure.
type AddonListByUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddonListByUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddonListByUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAddonListByUserBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAddonListByUserUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAddonListByUserNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAddonListByUserInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAddonListByUserOK creates a AddonListByUserOK with default headers values
func NewAddonListByUserOK() *AddonListByUserOK {
	return &AddonListByUserOK{}
}

/* AddonListByUserOK describes a response with status code 200, with default header values.

OK
*/
type AddonListByUserOK struct {
	Payload *models.V0OwnerAddOnsListResponseModel
}

func (o *AddonListByUserOK) Error() string {
	return fmt.Sprintf("[GET /users/{user-slug}/addons][%d] addonListByUserOK  %+v", 200, o.Payload)
}
func (o *AddonListByUserOK) GetPayload() *models.V0OwnerAddOnsListResponseModel {
	return o.Payload
}

func (o *AddonListByUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V0OwnerAddOnsListResponseModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddonListByUserBadRequest creates a AddonListByUserBadRequest with default headers values
func NewAddonListByUserBadRequest() *AddonListByUserBadRequest {
	return &AddonListByUserBadRequest{}
}

/* AddonListByUserBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type AddonListByUserBadRequest struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *AddonListByUserBadRequest) Error() string {
	return fmt.Sprintf("[GET /users/{user-slug}/addons][%d] addonListByUserBadRequest  %+v", 400, o.Payload)
}
func (o *AddonListByUserBadRequest) GetPayload() *models.ServiceStandardErrorRespModel {
	return o.Payload
}

func (o *AddonListByUserBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddonListByUserUnauthorized creates a AddonListByUserUnauthorized with default headers values
func NewAddonListByUserUnauthorized() *AddonListByUserUnauthorized {
	return &AddonListByUserUnauthorized{}
}

/* AddonListByUserUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type AddonListByUserUnauthorized struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *AddonListByUserUnauthorized) Error() string {
	return fmt.Sprintf("[GET /users/{user-slug}/addons][%d] addonListByUserUnauthorized  %+v", 401, o.Payload)
}
func (o *AddonListByUserUnauthorized) GetPayload() *models.ServiceStandardErrorRespModel {
	return o.Payload
}

func (o *AddonListByUserUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddonListByUserNotFound creates a AddonListByUserNotFound with default headers values
func NewAddonListByUserNotFound() *AddonListByUserNotFound {
	return &AddonListByUserNotFound{}
}

/* AddonListByUserNotFound describes a response with status code 404, with default header values.

Not Found
*/
type AddonListByUserNotFound struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *AddonListByUserNotFound) Error() string {
	return fmt.Sprintf("[GET /users/{user-slug}/addons][%d] addonListByUserNotFound  %+v", 404, o.Payload)
}
func (o *AddonListByUserNotFound) GetPayload() *models.ServiceStandardErrorRespModel {
	return o.Payload
}

func (o *AddonListByUserNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddonListByUserInternalServerError creates a AddonListByUserInternalServerError with default headers values
func NewAddonListByUserInternalServerError() *AddonListByUserInternalServerError {
	return &AddonListByUserInternalServerError{}
}

/* AddonListByUserInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type AddonListByUserInternalServerError struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *AddonListByUserInternalServerError) Error() string {
	return fmt.Sprintf("[GET /users/{user-slug}/addons][%d] addonListByUserInternalServerError  %+v", 500, o.Payload)
}
func (o *AddonListByUserInternalServerError) GetPayload() *models.ServiceStandardErrorRespModel {
	return o.Payload
}

func (o *AddonListByUserInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
