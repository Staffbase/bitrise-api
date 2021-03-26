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

// OutgoingWebhookCreateReader is a Reader for the OutgoingWebhookCreate structure.
type OutgoingWebhookCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OutgoingWebhookCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOutgoingWebhookCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewOutgoingWebhookCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewOutgoingWebhookCreateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewOutgoingWebhookCreateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewOutgoingWebhookCreateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewOutgoingWebhookCreateOK creates a OutgoingWebhookCreateOK with default headers values
func NewOutgoingWebhookCreateOK() *OutgoingWebhookCreateOK {
	return &OutgoingWebhookCreateOK{}
}

/* OutgoingWebhookCreateOK describes a response with status code 200, with default header values.

OK
*/
type OutgoingWebhookCreateOK struct {
	Payload *models.V0AppWebhookCreatedResponseModel
}

func (o *OutgoingWebhookCreateOK) Error() string {
	return fmt.Sprintf("[POST /apps/{app-slug}/outgoing-webhooks][%d] outgoingWebhookCreateOK  %+v", 200, o.Payload)
}
func (o *OutgoingWebhookCreateOK) GetPayload() *models.V0AppWebhookCreatedResponseModel {
	return o.Payload
}

func (o *OutgoingWebhookCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V0AppWebhookCreatedResponseModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOutgoingWebhookCreateBadRequest creates a OutgoingWebhookCreateBadRequest with default headers values
func NewOutgoingWebhookCreateBadRequest() *OutgoingWebhookCreateBadRequest {
	return &OutgoingWebhookCreateBadRequest{}
}

/* OutgoingWebhookCreateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type OutgoingWebhookCreateBadRequest struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *OutgoingWebhookCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /apps/{app-slug}/outgoing-webhooks][%d] outgoingWebhookCreateBadRequest  %+v", 400, o.Payload)
}
func (o *OutgoingWebhookCreateBadRequest) GetPayload() *models.ServiceStandardErrorRespModel {
	return o.Payload
}

func (o *OutgoingWebhookCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOutgoingWebhookCreateUnauthorized creates a OutgoingWebhookCreateUnauthorized with default headers values
func NewOutgoingWebhookCreateUnauthorized() *OutgoingWebhookCreateUnauthorized {
	return &OutgoingWebhookCreateUnauthorized{}
}

/* OutgoingWebhookCreateUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type OutgoingWebhookCreateUnauthorized struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *OutgoingWebhookCreateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /apps/{app-slug}/outgoing-webhooks][%d] outgoingWebhookCreateUnauthorized  %+v", 401, o.Payload)
}
func (o *OutgoingWebhookCreateUnauthorized) GetPayload() *models.ServiceStandardErrorRespModel {
	return o.Payload
}

func (o *OutgoingWebhookCreateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOutgoingWebhookCreateNotFound creates a OutgoingWebhookCreateNotFound with default headers values
func NewOutgoingWebhookCreateNotFound() *OutgoingWebhookCreateNotFound {
	return &OutgoingWebhookCreateNotFound{}
}

/* OutgoingWebhookCreateNotFound describes a response with status code 404, with default header values.

Not Found
*/
type OutgoingWebhookCreateNotFound struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *OutgoingWebhookCreateNotFound) Error() string {
	return fmt.Sprintf("[POST /apps/{app-slug}/outgoing-webhooks][%d] outgoingWebhookCreateNotFound  %+v", 404, o.Payload)
}
func (o *OutgoingWebhookCreateNotFound) GetPayload() *models.ServiceStandardErrorRespModel {
	return o.Payload
}

func (o *OutgoingWebhookCreateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOutgoingWebhookCreateInternalServerError creates a OutgoingWebhookCreateInternalServerError with default headers values
func NewOutgoingWebhookCreateInternalServerError() *OutgoingWebhookCreateInternalServerError {
	return &OutgoingWebhookCreateInternalServerError{}
}

/* OutgoingWebhookCreateInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type OutgoingWebhookCreateInternalServerError struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *OutgoingWebhookCreateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /apps/{app-slug}/outgoing-webhooks][%d] outgoingWebhookCreateInternalServerError  %+v", 500, o.Payload)
}
func (o *OutgoingWebhookCreateInternalServerError) GetPayload() *models.ServiceStandardErrorRespModel {
	return o.Payload
}

func (o *OutgoingWebhookCreateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
