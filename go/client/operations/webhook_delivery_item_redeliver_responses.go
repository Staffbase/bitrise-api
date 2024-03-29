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

// WebhookDeliveryItemRedeliverReader is a Reader for the WebhookDeliveryItemRedeliver structure.
type WebhookDeliveryItemRedeliverReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WebhookDeliveryItemRedeliverReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWebhookDeliveryItemRedeliverOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewWebhookDeliveryItemRedeliverBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewWebhookDeliveryItemRedeliverUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewWebhookDeliveryItemRedeliverInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /apps/{app-slug}/outgoing-webhooks/{app-webhook-slug}/delivery-items/{webhook-delivery-item-slug}/redeliver] webhook-delivery-item-redeliver", response, response.Code())
	}
}

// NewWebhookDeliveryItemRedeliverOK creates a WebhookDeliveryItemRedeliverOK with default headers values
func NewWebhookDeliveryItemRedeliverOK() *WebhookDeliveryItemRedeliverOK {
	return &WebhookDeliveryItemRedeliverOK{}
}

/*
WebhookDeliveryItemRedeliverOK describes a response with status code 200, with default header values.

OK
*/
type WebhookDeliveryItemRedeliverOK struct {
	Payload *models.ServiceStandardErrorRespModel
}

// IsSuccess returns true when this webhook delivery item redeliver o k response has a 2xx status code
func (o *WebhookDeliveryItemRedeliverOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this webhook delivery item redeliver o k response has a 3xx status code
func (o *WebhookDeliveryItemRedeliverOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this webhook delivery item redeliver o k response has a 4xx status code
func (o *WebhookDeliveryItemRedeliverOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this webhook delivery item redeliver o k response has a 5xx status code
func (o *WebhookDeliveryItemRedeliverOK) IsServerError() bool {
	return false
}

// IsCode returns true when this webhook delivery item redeliver o k response a status code equal to that given
func (o *WebhookDeliveryItemRedeliverOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the webhook delivery item redeliver o k response
func (o *WebhookDeliveryItemRedeliverOK) Code() int {
	return 200
}

func (o *WebhookDeliveryItemRedeliverOK) Error() string {
	return fmt.Sprintf("[POST /apps/{app-slug}/outgoing-webhooks/{app-webhook-slug}/delivery-items/{webhook-delivery-item-slug}/redeliver][%d] webhookDeliveryItemRedeliverOK  %+v", 200, o.Payload)
}

func (o *WebhookDeliveryItemRedeliverOK) String() string {
	return fmt.Sprintf("[POST /apps/{app-slug}/outgoing-webhooks/{app-webhook-slug}/delivery-items/{webhook-delivery-item-slug}/redeliver][%d] webhookDeliveryItemRedeliverOK  %+v", 200, o.Payload)
}

func (o *WebhookDeliveryItemRedeliverOK) GetPayload() *models.ServiceStandardErrorRespModel {
	return o.Payload
}

func (o *WebhookDeliveryItemRedeliverOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWebhookDeliveryItemRedeliverBadRequest creates a WebhookDeliveryItemRedeliverBadRequest with default headers values
func NewWebhookDeliveryItemRedeliverBadRequest() *WebhookDeliveryItemRedeliverBadRequest {
	return &WebhookDeliveryItemRedeliverBadRequest{}
}

/*
WebhookDeliveryItemRedeliverBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type WebhookDeliveryItemRedeliverBadRequest struct {
	Payload *models.ServiceStandardErrorRespModel
}

// IsSuccess returns true when this webhook delivery item redeliver bad request response has a 2xx status code
func (o *WebhookDeliveryItemRedeliverBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this webhook delivery item redeliver bad request response has a 3xx status code
func (o *WebhookDeliveryItemRedeliverBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this webhook delivery item redeliver bad request response has a 4xx status code
func (o *WebhookDeliveryItemRedeliverBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this webhook delivery item redeliver bad request response has a 5xx status code
func (o *WebhookDeliveryItemRedeliverBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this webhook delivery item redeliver bad request response a status code equal to that given
func (o *WebhookDeliveryItemRedeliverBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the webhook delivery item redeliver bad request response
func (o *WebhookDeliveryItemRedeliverBadRequest) Code() int {
	return 400
}

func (o *WebhookDeliveryItemRedeliverBadRequest) Error() string {
	return fmt.Sprintf("[POST /apps/{app-slug}/outgoing-webhooks/{app-webhook-slug}/delivery-items/{webhook-delivery-item-slug}/redeliver][%d] webhookDeliveryItemRedeliverBadRequest  %+v", 400, o.Payload)
}

func (o *WebhookDeliveryItemRedeliverBadRequest) String() string {
	return fmt.Sprintf("[POST /apps/{app-slug}/outgoing-webhooks/{app-webhook-slug}/delivery-items/{webhook-delivery-item-slug}/redeliver][%d] webhookDeliveryItemRedeliverBadRequest  %+v", 400, o.Payload)
}

func (o *WebhookDeliveryItemRedeliverBadRequest) GetPayload() *models.ServiceStandardErrorRespModel {
	return o.Payload
}

func (o *WebhookDeliveryItemRedeliverBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWebhookDeliveryItemRedeliverUnauthorized creates a WebhookDeliveryItemRedeliverUnauthorized with default headers values
func NewWebhookDeliveryItemRedeliverUnauthorized() *WebhookDeliveryItemRedeliverUnauthorized {
	return &WebhookDeliveryItemRedeliverUnauthorized{}
}

/*
WebhookDeliveryItemRedeliverUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type WebhookDeliveryItemRedeliverUnauthorized struct {
	Payload *models.ServiceStandardErrorRespModel
}

// IsSuccess returns true when this webhook delivery item redeliver unauthorized response has a 2xx status code
func (o *WebhookDeliveryItemRedeliverUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this webhook delivery item redeliver unauthorized response has a 3xx status code
func (o *WebhookDeliveryItemRedeliverUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this webhook delivery item redeliver unauthorized response has a 4xx status code
func (o *WebhookDeliveryItemRedeliverUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this webhook delivery item redeliver unauthorized response has a 5xx status code
func (o *WebhookDeliveryItemRedeliverUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this webhook delivery item redeliver unauthorized response a status code equal to that given
func (o *WebhookDeliveryItemRedeliverUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the webhook delivery item redeliver unauthorized response
func (o *WebhookDeliveryItemRedeliverUnauthorized) Code() int {
	return 401
}

func (o *WebhookDeliveryItemRedeliverUnauthorized) Error() string {
	return fmt.Sprintf("[POST /apps/{app-slug}/outgoing-webhooks/{app-webhook-slug}/delivery-items/{webhook-delivery-item-slug}/redeliver][%d] webhookDeliveryItemRedeliverUnauthorized  %+v", 401, o.Payload)
}

func (o *WebhookDeliveryItemRedeliverUnauthorized) String() string {
	return fmt.Sprintf("[POST /apps/{app-slug}/outgoing-webhooks/{app-webhook-slug}/delivery-items/{webhook-delivery-item-slug}/redeliver][%d] webhookDeliveryItemRedeliverUnauthorized  %+v", 401, o.Payload)
}

func (o *WebhookDeliveryItemRedeliverUnauthorized) GetPayload() *models.ServiceStandardErrorRespModel {
	return o.Payload
}

func (o *WebhookDeliveryItemRedeliverUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWebhookDeliveryItemRedeliverInternalServerError creates a WebhookDeliveryItemRedeliverInternalServerError with default headers values
func NewWebhookDeliveryItemRedeliverInternalServerError() *WebhookDeliveryItemRedeliverInternalServerError {
	return &WebhookDeliveryItemRedeliverInternalServerError{}
}

/*
WebhookDeliveryItemRedeliverInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type WebhookDeliveryItemRedeliverInternalServerError struct {
	Payload *models.ServiceStandardErrorRespModel
}

// IsSuccess returns true when this webhook delivery item redeliver internal server error response has a 2xx status code
func (o *WebhookDeliveryItemRedeliverInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this webhook delivery item redeliver internal server error response has a 3xx status code
func (o *WebhookDeliveryItemRedeliverInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this webhook delivery item redeliver internal server error response has a 4xx status code
func (o *WebhookDeliveryItemRedeliverInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this webhook delivery item redeliver internal server error response has a 5xx status code
func (o *WebhookDeliveryItemRedeliverInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this webhook delivery item redeliver internal server error response a status code equal to that given
func (o *WebhookDeliveryItemRedeliverInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the webhook delivery item redeliver internal server error response
func (o *WebhookDeliveryItemRedeliverInternalServerError) Code() int {
	return 500
}

func (o *WebhookDeliveryItemRedeliverInternalServerError) Error() string {
	return fmt.Sprintf("[POST /apps/{app-slug}/outgoing-webhooks/{app-webhook-slug}/delivery-items/{webhook-delivery-item-slug}/redeliver][%d] webhookDeliveryItemRedeliverInternalServerError  %+v", 500, o.Payload)
}

func (o *WebhookDeliveryItemRedeliverInternalServerError) String() string {
	return fmt.Sprintf("[POST /apps/{app-slug}/outgoing-webhooks/{app-webhook-slug}/delivery-items/{webhook-delivery-item-slug}/redeliver][%d] webhookDeliveryItemRedeliverInternalServerError  %+v", 500, o.Payload)
}

func (o *WebhookDeliveryItemRedeliverInternalServerError) GetPayload() *models.ServiceStandardErrorRespModel {
	return o.Payload
}

func (o *WebhookDeliveryItemRedeliverInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
