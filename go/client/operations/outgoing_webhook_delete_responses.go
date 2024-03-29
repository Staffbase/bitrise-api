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

// OutgoingWebhookDeleteReader is a Reader for the OutgoingWebhookDelete structure.
type OutgoingWebhookDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OutgoingWebhookDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOutgoingWebhookDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewOutgoingWebhookDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewOutgoingWebhookDeleteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewOutgoingWebhookDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewOutgoingWebhookDeleteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /apps/{app-slug}/outgoing-webhooks/{app-webhook-slug}] outgoing-webhook-delete", response, response.Code())
	}
}

// NewOutgoingWebhookDeleteOK creates a OutgoingWebhookDeleteOK with default headers values
func NewOutgoingWebhookDeleteOK() *OutgoingWebhookDeleteOK {
	return &OutgoingWebhookDeleteOK{}
}

/*
OutgoingWebhookDeleteOK describes a response with status code 200, with default header values.

OK
*/
type OutgoingWebhookDeleteOK struct {
	Payload *models.V0AppWebhookDeletedResponseModel
}

// IsSuccess returns true when this outgoing webhook delete o k response has a 2xx status code
func (o *OutgoingWebhookDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this outgoing webhook delete o k response has a 3xx status code
func (o *OutgoingWebhookDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this outgoing webhook delete o k response has a 4xx status code
func (o *OutgoingWebhookDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this outgoing webhook delete o k response has a 5xx status code
func (o *OutgoingWebhookDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this outgoing webhook delete o k response a status code equal to that given
func (o *OutgoingWebhookDeleteOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the outgoing webhook delete o k response
func (o *OutgoingWebhookDeleteOK) Code() int {
	return 200
}

func (o *OutgoingWebhookDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /apps/{app-slug}/outgoing-webhooks/{app-webhook-slug}][%d] outgoingWebhookDeleteOK  %+v", 200, o.Payload)
}

func (o *OutgoingWebhookDeleteOK) String() string {
	return fmt.Sprintf("[DELETE /apps/{app-slug}/outgoing-webhooks/{app-webhook-slug}][%d] outgoingWebhookDeleteOK  %+v", 200, o.Payload)
}

func (o *OutgoingWebhookDeleteOK) GetPayload() *models.V0AppWebhookDeletedResponseModel {
	return o.Payload
}

func (o *OutgoingWebhookDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V0AppWebhookDeletedResponseModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOutgoingWebhookDeleteBadRequest creates a OutgoingWebhookDeleteBadRequest with default headers values
func NewOutgoingWebhookDeleteBadRequest() *OutgoingWebhookDeleteBadRequest {
	return &OutgoingWebhookDeleteBadRequest{}
}

/*
OutgoingWebhookDeleteBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type OutgoingWebhookDeleteBadRequest struct {
	Payload *models.ServiceStandardErrorRespModel
}

// IsSuccess returns true when this outgoing webhook delete bad request response has a 2xx status code
func (o *OutgoingWebhookDeleteBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this outgoing webhook delete bad request response has a 3xx status code
func (o *OutgoingWebhookDeleteBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this outgoing webhook delete bad request response has a 4xx status code
func (o *OutgoingWebhookDeleteBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this outgoing webhook delete bad request response has a 5xx status code
func (o *OutgoingWebhookDeleteBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this outgoing webhook delete bad request response a status code equal to that given
func (o *OutgoingWebhookDeleteBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the outgoing webhook delete bad request response
func (o *OutgoingWebhookDeleteBadRequest) Code() int {
	return 400
}

func (o *OutgoingWebhookDeleteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /apps/{app-slug}/outgoing-webhooks/{app-webhook-slug}][%d] outgoingWebhookDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *OutgoingWebhookDeleteBadRequest) String() string {
	return fmt.Sprintf("[DELETE /apps/{app-slug}/outgoing-webhooks/{app-webhook-slug}][%d] outgoingWebhookDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *OutgoingWebhookDeleteBadRequest) GetPayload() *models.ServiceStandardErrorRespModel {
	return o.Payload
}

func (o *OutgoingWebhookDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOutgoingWebhookDeleteUnauthorized creates a OutgoingWebhookDeleteUnauthorized with default headers values
func NewOutgoingWebhookDeleteUnauthorized() *OutgoingWebhookDeleteUnauthorized {
	return &OutgoingWebhookDeleteUnauthorized{}
}

/*
OutgoingWebhookDeleteUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type OutgoingWebhookDeleteUnauthorized struct {
	Payload *models.ServiceStandardErrorRespModel
}

// IsSuccess returns true when this outgoing webhook delete unauthorized response has a 2xx status code
func (o *OutgoingWebhookDeleteUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this outgoing webhook delete unauthorized response has a 3xx status code
func (o *OutgoingWebhookDeleteUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this outgoing webhook delete unauthorized response has a 4xx status code
func (o *OutgoingWebhookDeleteUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this outgoing webhook delete unauthorized response has a 5xx status code
func (o *OutgoingWebhookDeleteUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this outgoing webhook delete unauthorized response a status code equal to that given
func (o *OutgoingWebhookDeleteUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the outgoing webhook delete unauthorized response
func (o *OutgoingWebhookDeleteUnauthorized) Code() int {
	return 401
}

func (o *OutgoingWebhookDeleteUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /apps/{app-slug}/outgoing-webhooks/{app-webhook-slug}][%d] outgoingWebhookDeleteUnauthorized  %+v", 401, o.Payload)
}

func (o *OutgoingWebhookDeleteUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /apps/{app-slug}/outgoing-webhooks/{app-webhook-slug}][%d] outgoingWebhookDeleteUnauthorized  %+v", 401, o.Payload)
}

func (o *OutgoingWebhookDeleteUnauthorized) GetPayload() *models.ServiceStandardErrorRespModel {
	return o.Payload
}

func (o *OutgoingWebhookDeleteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOutgoingWebhookDeleteNotFound creates a OutgoingWebhookDeleteNotFound with default headers values
func NewOutgoingWebhookDeleteNotFound() *OutgoingWebhookDeleteNotFound {
	return &OutgoingWebhookDeleteNotFound{}
}

/*
OutgoingWebhookDeleteNotFound describes a response with status code 404, with default header values.

Not Found
*/
type OutgoingWebhookDeleteNotFound struct {
	Payload *models.ServiceStandardErrorRespModel
}

// IsSuccess returns true when this outgoing webhook delete not found response has a 2xx status code
func (o *OutgoingWebhookDeleteNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this outgoing webhook delete not found response has a 3xx status code
func (o *OutgoingWebhookDeleteNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this outgoing webhook delete not found response has a 4xx status code
func (o *OutgoingWebhookDeleteNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this outgoing webhook delete not found response has a 5xx status code
func (o *OutgoingWebhookDeleteNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this outgoing webhook delete not found response a status code equal to that given
func (o *OutgoingWebhookDeleteNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the outgoing webhook delete not found response
func (o *OutgoingWebhookDeleteNotFound) Code() int {
	return 404
}

func (o *OutgoingWebhookDeleteNotFound) Error() string {
	return fmt.Sprintf("[DELETE /apps/{app-slug}/outgoing-webhooks/{app-webhook-slug}][%d] outgoingWebhookDeleteNotFound  %+v", 404, o.Payload)
}

func (o *OutgoingWebhookDeleteNotFound) String() string {
	return fmt.Sprintf("[DELETE /apps/{app-slug}/outgoing-webhooks/{app-webhook-slug}][%d] outgoingWebhookDeleteNotFound  %+v", 404, o.Payload)
}

func (o *OutgoingWebhookDeleteNotFound) GetPayload() *models.ServiceStandardErrorRespModel {
	return o.Payload
}

func (o *OutgoingWebhookDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOutgoingWebhookDeleteInternalServerError creates a OutgoingWebhookDeleteInternalServerError with default headers values
func NewOutgoingWebhookDeleteInternalServerError() *OutgoingWebhookDeleteInternalServerError {
	return &OutgoingWebhookDeleteInternalServerError{}
}

/*
OutgoingWebhookDeleteInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type OutgoingWebhookDeleteInternalServerError struct {
	Payload *models.ServiceStandardErrorRespModel
}

// IsSuccess returns true when this outgoing webhook delete internal server error response has a 2xx status code
func (o *OutgoingWebhookDeleteInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this outgoing webhook delete internal server error response has a 3xx status code
func (o *OutgoingWebhookDeleteInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this outgoing webhook delete internal server error response has a 4xx status code
func (o *OutgoingWebhookDeleteInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this outgoing webhook delete internal server error response has a 5xx status code
func (o *OutgoingWebhookDeleteInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this outgoing webhook delete internal server error response a status code equal to that given
func (o *OutgoingWebhookDeleteInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the outgoing webhook delete internal server error response
func (o *OutgoingWebhookDeleteInternalServerError) Code() int {
	return 500
}

func (o *OutgoingWebhookDeleteInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /apps/{app-slug}/outgoing-webhooks/{app-webhook-slug}][%d] outgoingWebhookDeleteInternalServerError  %+v", 500, o.Payload)
}

func (o *OutgoingWebhookDeleteInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /apps/{app-slug}/outgoing-webhooks/{app-webhook-slug}][%d] outgoingWebhookDeleteInternalServerError  %+v", 500, o.Payload)
}

func (o *OutgoingWebhookDeleteInternalServerError) GetPayload() *models.ServiceStandardErrorRespModel {
	return o.Payload
}

func (o *OutgoingWebhookDeleteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
