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

// OutgoingWebhookListReader is a Reader for the OutgoingWebhookList structure.
type OutgoingWebhookListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OutgoingWebhookListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOutgoingWebhookListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewOutgoingWebhookListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewOutgoingWebhookListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewOutgoingWebhookListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /apps/{app-slug}/outgoing-webhooks] outgoing-webhook-list", response, response.Code())
	}
}

// NewOutgoingWebhookListOK creates a OutgoingWebhookListOK with default headers values
func NewOutgoingWebhookListOK() *OutgoingWebhookListOK {
	return &OutgoingWebhookListOK{}
}

/*
OutgoingWebhookListOK describes a response with status code 200, with default header values.

OK
*/
type OutgoingWebhookListOK struct {
	Payload *models.V0AppWebhookListResponseModel
}

// IsSuccess returns true when this outgoing webhook list o k response has a 2xx status code
func (o *OutgoingWebhookListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this outgoing webhook list o k response has a 3xx status code
func (o *OutgoingWebhookListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this outgoing webhook list o k response has a 4xx status code
func (o *OutgoingWebhookListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this outgoing webhook list o k response has a 5xx status code
func (o *OutgoingWebhookListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this outgoing webhook list o k response a status code equal to that given
func (o *OutgoingWebhookListOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the outgoing webhook list o k response
func (o *OutgoingWebhookListOK) Code() int {
	return 200
}

func (o *OutgoingWebhookListOK) Error() string {
	return fmt.Sprintf("[GET /apps/{app-slug}/outgoing-webhooks][%d] outgoingWebhookListOK  %+v", 200, o.Payload)
}

func (o *OutgoingWebhookListOK) String() string {
	return fmt.Sprintf("[GET /apps/{app-slug}/outgoing-webhooks][%d] outgoingWebhookListOK  %+v", 200, o.Payload)
}

func (o *OutgoingWebhookListOK) GetPayload() *models.V0AppWebhookListResponseModel {
	return o.Payload
}

func (o *OutgoingWebhookListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V0AppWebhookListResponseModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOutgoingWebhookListBadRequest creates a OutgoingWebhookListBadRequest with default headers values
func NewOutgoingWebhookListBadRequest() *OutgoingWebhookListBadRequest {
	return &OutgoingWebhookListBadRequest{}
}

/*
OutgoingWebhookListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type OutgoingWebhookListBadRequest struct {
	Payload *models.ServiceStandardErrorRespModel
}

// IsSuccess returns true when this outgoing webhook list bad request response has a 2xx status code
func (o *OutgoingWebhookListBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this outgoing webhook list bad request response has a 3xx status code
func (o *OutgoingWebhookListBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this outgoing webhook list bad request response has a 4xx status code
func (o *OutgoingWebhookListBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this outgoing webhook list bad request response has a 5xx status code
func (o *OutgoingWebhookListBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this outgoing webhook list bad request response a status code equal to that given
func (o *OutgoingWebhookListBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the outgoing webhook list bad request response
func (o *OutgoingWebhookListBadRequest) Code() int {
	return 400
}

func (o *OutgoingWebhookListBadRequest) Error() string {
	return fmt.Sprintf("[GET /apps/{app-slug}/outgoing-webhooks][%d] outgoingWebhookListBadRequest  %+v", 400, o.Payload)
}

func (o *OutgoingWebhookListBadRequest) String() string {
	return fmt.Sprintf("[GET /apps/{app-slug}/outgoing-webhooks][%d] outgoingWebhookListBadRequest  %+v", 400, o.Payload)
}

func (o *OutgoingWebhookListBadRequest) GetPayload() *models.ServiceStandardErrorRespModel {
	return o.Payload
}

func (o *OutgoingWebhookListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOutgoingWebhookListUnauthorized creates a OutgoingWebhookListUnauthorized with default headers values
func NewOutgoingWebhookListUnauthorized() *OutgoingWebhookListUnauthorized {
	return &OutgoingWebhookListUnauthorized{}
}

/*
OutgoingWebhookListUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type OutgoingWebhookListUnauthorized struct {
	Payload *models.ServiceStandardErrorRespModel
}

// IsSuccess returns true when this outgoing webhook list unauthorized response has a 2xx status code
func (o *OutgoingWebhookListUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this outgoing webhook list unauthorized response has a 3xx status code
func (o *OutgoingWebhookListUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this outgoing webhook list unauthorized response has a 4xx status code
func (o *OutgoingWebhookListUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this outgoing webhook list unauthorized response has a 5xx status code
func (o *OutgoingWebhookListUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this outgoing webhook list unauthorized response a status code equal to that given
func (o *OutgoingWebhookListUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the outgoing webhook list unauthorized response
func (o *OutgoingWebhookListUnauthorized) Code() int {
	return 401
}

func (o *OutgoingWebhookListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /apps/{app-slug}/outgoing-webhooks][%d] outgoingWebhookListUnauthorized  %+v", 401, o.Payload)
}

func (o *OutgoingWebhookListUnauthorized) String() string {
	return fmt.Sprintf("[GET /apps/{app-slug}/outgoing-webhooks][%d] outgoingWebhookListUnauthorized  %+v", 401, o.Payload)
}

func (o *OutgoingWebhookListUnauthorized) GetPayload() *models.ServiceStandardErrorRespModel {
	return o.Payload
}

func (o *OutgoingWebhookListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOutgoingWebhookListInternalServerError creates a OutgoingWebhookListInternalServerError with default headers values
func NewOutgoingWebhookListInternalServerError() *OutgoingWebhookListInternalServerError {
	return &OutgoingWebhookListInternalServerError{}
}

/*
OutgoingWebhookListInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type OutgoingWebhookListInternalServerError struct {
	Payload *models.ServiceStandardErrorRespModel
}

// IsSuccess returns true when this outgoing webhook list internal server error response has a 2xx status code
func (o *OutgoingWebhookListInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this outgoing webhook list internal server error response has a 3xx status code
func (o *OutgoingWebhookListInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this outgoing webhook list internal server error response has a 4xx status code
func (o *OutgoingWebhookListInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this outgoing webhook list internal server error response has a 5xx status code
func (o *OutgoingWebhookListInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this outgoing webhook list internal server error response a status code equal to that given
func (o *OutgoingWebhookListInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the outgoing webhook list internal server error response
func (o *OutgoingWebhookListInternalServerError) Code() int {
	return 500
}

func (o *OutgoingWebhookListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /apps/{app-slug}/outgoing-webhooks][%d] outgoingWebhookListInternalServerError  %+v", 500, o.Payload)
}

func (o *OutgoingWebhookListInternalServerError) String() string {
	return fmt.Sprintf("[GET /apps/{app-slug}/outgoing-webhooks][%d] outgoingWebhookListInternalServerError  %+v", 500, o.Payload)
}

func (o *OutgoingWebhookListInternalServerError) GetPayload() *models.ServiceStandardErrorRespModel {
	return o.Payload
}

func (o *OutgoingWebhookListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
