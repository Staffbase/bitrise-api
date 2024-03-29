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
		return nil, runtime.NewAPIError("[POST /apps/{app-slug}/outgoing-webhooks] outgoing-webhook-create", response, response.Code())
	}
}

// NewOutgoingWebhookCreateOK creates a OutgoingWebhookCreateOK with default headers values
func NewOutgoingWebhookCreateOK() *OutgoingWebhookCreateOK {
	return &OutgoingWebhookCreateOK{}
}

/*
OutgoingWebhookCreateOK describes a response with status code 200, with default header values.

OK
*/
type OutgoingWebhookCreateOK struct {
	Payload *models.V0AppWebhookCreatedResponseModel
}

// IsSuccess returns true when this outgoing webhook create o k response has a 2xx status code
func (o *OutgoingWebhookCreateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this outgoing webhook create o k response has a 3xx status code
func (o *OutgoingWebhookCreateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this outgoing webhook create o k response has a 4xx status code
func (o *OutgoingWebhookCreateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this outgoing webhook create o k response has a 5xx status code
func (o *OutgoingWebhookCreateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this outgoing webhook create o k response a status code equal to that given
func (o *OutgoingWebhookCreateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the outgoing webhook create o k response
func (o *OutgoingWebhookCreateOK) Code() int {
	return 200
}

func (o *OutgoingWebhookCreateOK) Error() string {
	return fmt.Sprintf("[POST /apps/{app-slug}/outgoing-webhooks][%d] outgoingWebhookCreateOK  %+v", 200, o.Payload)
}

func (o *OutgoingWebhookCreateOK) String() string {
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

/*
OutgoingWebhookCreateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type OutgoingWebhookCreateBadRequest struct {
	Payload *models.ServiceStandardErrorRespModel
}

// IsSuccess returns true when this outgoing webhook create bad request response has a 2xx status code
func (o *OutgoingWebhookCreateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this outgoing webhook create bad request response has a 3xx status code
func (o *OutgoingWebhookCreateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this outgoing webhook create bad request response has a 4xx status code
func (o *OutgoingWebhookCreateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this outgoing webhook create bad request response has a 5xx status code
func (o *OutgoingWebhookCreateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this outgoing webhook create bad request response a status code equal to that given
func (o *OutgoingWebhookCreateBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the outgoing webhook create bad request response
func (o *OutgoingWebhookCreateBadRequest) Code() int {
	return 400
}

func (o *OutgoingWebhookCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /apps/{app-slug}/outgoing-webhooks][%d] outgoingWebhookCreateBadRequest  %+v", 400, o.Payload)
}

func (o *OutgoingWebhookCreateBadRequest) String() string {
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

/*
OutgoingWebhookCreateUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type OutgoingWebhookCreateUnauthorized struct {
	Payload *models.ServiceStandardErrorRespModel
}

// IsSuccess returns true when this outgoing webhook create unauthorized response has a 2xx status code
func (o *OutgoingWebhookCreateUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this outgoing webhook create unauthorized response has a 3xx status code
func (o *OutgoingWebhookCreateUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this outgoing webhook create unauthorized response has a 4xx status code
func (o *OutgoingWebhookCreateUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this outgoing webhook create unauthorized response has a 5xx status code
func (o *OutgoingWebhookCreateUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this outgoing webhook create unauthorized response a status code equal to that given
func (o *OutgoingWebhookCreateUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the outgoing webhook create unauthorized response
func (o *OutgoingWebhookCreateUnauthorized) Code() int {
	return 401
}

func (o *OutgoingWebhookCreateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /apps/{app-slug}/outgoing-webhooks][%d] outgoingWebhookCreateUnauthorized  %+v", 401, o.Payload)
}

func (o *OutgoingWebhookCreateUnauthorized) String() string {
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

/*
OutgoingWebhookCreateNotFound describes a response with status code 404, with default header values.

Not Found
*/
type OutgoingWebhookCreateNotFound struct {
	Payload *models.ServiceStandardErrorRespModel
}

// IsSuccess returns true when this outgoing webhook create not found response has a 2xx status code
func (o *OutgoingWebhookCreateNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this outgoing webhook create not found response has a 3xx status code
func (o *OutgoingWebhookCreateNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this outgoing webhook create not found response has a 4xx status code
func (o *OutgoingWebhookCreateNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this outgoing webhook create not found response has a 5xx status code
func (o *OutgoingWebhookCreateNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this outgoing webhook create not found response a status code equal to that given
func (o *OutgoingWebhookCreateNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the outgoing webhook create not found response
func (o *OutgoingWebhookCreateNotFound) Code() int {
	return 404
}

func (o *OutgoingWebhookCreateNotFound) Error() string {
	return fmt.Sprintf("[POST /apps/{app-slug}/outgoing-webhooks][%d] outgoingWebhookCreateNotFound  %+v", 404, o.Payload)
}

func (o *OutgoingWebhookCreateNotFound) String() string {
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

/*
OutgoingWebhookCreateInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type OutgoingWebhookCreateInternalServerError struct {
	Payload *models.ServiceStandardErrorRespModel
}

// IsSuccess returns true when this outgoing webhook create internal server error response has a 2xx status code
func (o *OutgoingWebhookCreateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this outgoing webhook create internal server error response has a 3xx status code
func (o *OutgoingWebhookCreateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this outgoing webhook create internal server error response has a 4xx status code
func (o *OutgoingWebhookCreateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this outgoing webhook create internal server error response has a 5xx status code
func (o *OutgoingWebhookCreateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this outgoing webhook create internal server error response a status code equal to that given
func (o *OutgoingWebhookCreateInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the outgoing webhook create internal server error response
func (o *OutgoingWebhookCreateInternalServerError) Code() int {
	return 500
}

func (o *OutgoingWebhookCreateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /apps/{app-slug}/outgoing-webhooks][%d] outgoingWebhookCreateInternalServerError  %+v", 500, o.Payload)
}

func (o *OutgoingWebhookCreateInternalServerError) String() string {
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
