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

// AppWebhookCreateReader is a Reader for the AppWebhookCreate structure.
type AppWebhookCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AppWebhookCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAppWebhookCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAppWebhookCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAppWebhookCreateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAppWebhookCreateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAppWebhookCreateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAppWebhookCreateOK creates a AppWebhookCreateOK with default headers values
func NewAppWebhookCreateOK() *AppWebhookCreateOK {
	return &AppWebhookCreateOK{}
}

/* AppWebhookCreateOK describes a response with status code 200, with default header values.

OK
*/
type AppWebhookCreateOK struct {
	Payload *models.V0WebhookRespModel
}

func (o *AppWebhookCreateOK) Error() string {
	return fmt.Sprintf("[POST /apps/{app-slug}/register-webhook][%d] appWebhookCreateOK  %+v", 200, o.Payload)
}
func (o *AppWebhookCreateOK) GetPayload() *models.V0WebhookRespModel {
	return o.Payload
}

func (o *AppWebhookCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V0WebhookRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAppWebhookCreateBadRequest creates a AppWebhookCreateBadRequest with default headers values
func NewAppWebhookCreateBadRequest() *AppWebhookCreateBadRequest {
	return &AppWebhookCreateBadRequest{}
}

/* AppWebhookCreateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type AppWebhookCreateBadRequest struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *AppWebhookCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /apps/{app-slug}/register-webhook][%d] appWebhookCreateBadRequest  %+v", 400, o.Payload)
}
func (o *AppWebhookCreateBadRequest) GetPayload() *models.ServiceStandardErrorRespModel {
	return o.Payload
}

func (o *AppWebhookCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAppWebhookCreateUnauthorized creates a AppWebhookCreateUnauthorized with default headers values
func NewAppWebhookCreateUnauthorized() *AppWebhookCreateUnauthorized {
	return &AppWebhookCreateUnauthorized{}
}

/* AppWebhookCreateUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type AppWebhookCreateUnauthorized struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *AppWebhookCreateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /apps/{app-slug}/register-webhook][%d] appWebhookCreateUnauthorized  %+v", 401, o.Payload)
}
func (o *AppWebhookCreateUnauthorized) GetPayload() *models.ServiceStandardErrorRespModel {
	return o.Payload
}

func (o *AppWebhookCreateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAppWebhookCreateNotFound creates a AppWebhookCreateNotFound with default headers values
func NewAppWebhookCreateNotFound() *AppWebhookCreateNotFound {
	return &AppWebhookCreateNotFound{}
}

/* AppWebhookCreateNotFound describes a response with status code 404, with default header values.

Not Found
*/
type AppWebhookCreateNotFound struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *AppWebhookCreateNotFound) Error() string {
	return fmt.Sprintf("[POST /apps/{app-slug}/register-webhook][%d] appWebhookCreateNotFound  %+v", 404, o.Payload)
}
func (o *AppWebhookCreateNotFound) GetPayload() *models.ServiceStandardErrorRespModel {
	return o.Payload
}

func (o *AppWebhookCreateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAppWebhookCreateInternalServerError creates a AppWebhookCreateInternalServerError with default headers values
func NewAppWebhookCreateInternalServerError() *AppWebhookCreateInternalServerError {
	return &AppWebhookCreateInternalServerError{}
}

/* AppWebhookCreateInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type AppWebhookCreateInternalServerError struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *AppWebhookCreateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /apps/{app-slug}/register-webhook][%d] appWebhookCreateInternalServerError  %+v", 500, o.Payload)
}
func (o *AppWebhookCreateInternalServerError) GetPayload() *models.ServiceStandardErrorRespModel {
	return o.Payload
}

func (o *AppWebhookCreateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
