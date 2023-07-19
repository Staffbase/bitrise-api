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

// WebhookDeliveryItemListReader is a Reader for the WebhookDeliveryItemList structure.
type WebhookDeliveryItemListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WebhookDeliveryItemListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWebhookDeliveryItemListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewWebhookDeliveryItemListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewWebhookDeliveryItemListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewWebhookDeliveryItemListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /apps/{app-slug}/outgoing-webhooks/{app-webhook-slug}/delivery-items] webhook-delivery-item-list", response, response.Code())
	}
}

// NewWebhookDeliveryItemListOK creates a WebhookDeliveryItemListOK with default headers values
func NewWebhookDeliveryItemListOK() *WebhookDeliveryItemListOK {
	return &WebhookDeliveryItemListOK{}
}

/*
WebhookDeliveryItemListOK describes a response with status code 200, with default header values.

OK
*/
type WebhookDeliveryItemListOK struct {
	Payload *models.V0WebhookDeliveryItemShowResponseModel
}

// IsSuccess returns true when this webhook delivery item list o k response has a 2xx status code
func (o *WebhookDeliveryItemListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this webhook delivery item list o k response has a 3xx status code
func (o *WebhookDeliveryItemListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this webhook delivery item list o k response has a 4xx status code
func (o *WebhookDeliveryItemListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this webhook delivery item list o k response has a 5xx status code
func (o *WebhookDeliveryItemListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this webhook delivery item list o k response a status code equal to that given
func (o *WebhookDeliveryItemListOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the webhook delivery item list o k response
func (o *WebhookDeliveryItemListOK) Code() int {
	return 200
}

func (o *WebhookDeliveryItemListOK) Error() string {
	return fmt.Sprintf("[GET /apps/{app-slug}/outgoing-webhooks/{app-webhook-slug}/delivery-items][%d] webhookDeliveryItemListOK  %+v", 200, o.Payload)
}

func (o *WebhookDeliveryItemListOK) String() string {
	return fmt.Sprintf("[GET /apps/{app-slug}/outgoing-webhooks/{app-webhook-slug}/delivery-items][%d] webhookDeliveryItemListOK  %+v", 200, o.Payload)
}

func (o *WebhookDeliveryItemListOK) GetPayload() *models.V0WebhookDeliveryItemShowResponseModel {
	return o.Payload
}

func (o *WebhookDeliveryItemListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V0WebhookDeliveryItemShowResponseModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWebhookDeliveryItemListBadRequest creates a WebhookDeliveryItemListBadRequest with default headers values
func NewWebhookDeliveryItemListBadRequest() *WebhookDeliveryItemListBadRequest {
	return &WebhookDeliveryItemListBadRequest{}
}

/*
WebhookDeliveryItemListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type WebhookDeliveryItemListBadRequest struct {
	Payload *models.ServiceStandardErrorRespModel
}

// IsSuccess returns true when this webhook delivery item list bad request response has a 2xx status code
func (o *WebhookDeliveryItemListBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this webhook delivery item list bad request response has a 3xx status code
func (o *WebhookDeliveryItemListBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this webhook delivery item list bad request response has a 4xx status code
func (o *WebhookDeliveryItemListBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this webhook delivery item list bad request response has a 5xx status code
func (o *WebhookDeliveryItemListBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this webhook delivery item list bad request response a status code equal to that given
func (o *WebhookDeliveryItemListBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the webhook delivery item list bad request response
func (o *WebhookDeliveryItemListBadRequest) Code() int {
	return 400
}

func (o *WebhookDeliveryItemListBadRequest) Error() string {
	return fmt.Sprintf("[GET /apps/{app-slug}/outgoing-webhooks/{app-webhook-slug}/delivery-items][%d] webhookDeliveryItemListBadRequest  %+v", 400, o.Payload)
}

func (o *WebhookDeliveryItemListBadRequest) String() string {
	return fmt.Sprintf("[GET /apps/{app-slug}/outgoing-webhooks/{app-webhook-slug}/delivery-items][%d] webhookDeliveryItemListBadRequest  %+v", 400, o.Payload)
}

func (o *WebhookDeliveryItemListBadRequest) GetPayload() *models.ServiceStandardErrorRespModel {
	return o.Payload
}

func (o *WebhookDeliveryItemListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWebhookDeliveryItemListUnauthorized creates a WebhookDeliveryItemListUnauthorized with default headers values
func NewWebhookDeliveryItemListUnauthorized() *WebhookDeliveryItemListUnauthorized {
	return &WebhookDeliveryItemListUnauthorized{}
}

/*
WebhookDeliveryItemListUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type WebhookDeliveryItemListUnauthorized struct {
	Payload *models.ServiceStandardErrorRespModel
}

// IsSuccess returns true when this webhook delivery item list unauthorized response has a 2xx status code
func (o *WebhookDeliveryItemListUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this webhook delivery item list unauthorized response has a 3xx status code
func (o *WebhookDeliveryItemListUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this webhook delivery item list unauthorized response has a 4xx status code
func (o *WebhookDeliveryItemListUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this webhook delivery item list unauthorized response has a 5xx status code
func (o *WebhookDeliveryItemListUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this webhook delivery item list unauthorized response a status code equal to that given
func (o *WebhookDeliveryItemListUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the webhook delivery item list unauthorized response
func (o *WebhookDeliveryItemListUnauthorized) Code() int {
	return 401
}

func (o *WebhookDeliveryItemListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /apps/{app-slug}/outgoing-webhooks/{app-webhook-slug}/delivery-items][%d] webhookDeliveryItemListUnauthorized  %+v", 401, o.Payload)
}

func (o *WebhookDeliveryItemListUnauthorized) String() string {
	return fmt.Sprintf("[GET /apps/{app-slug}/outgoing-webhooks/{app-webhook-slug}/delivery-items][%d] webhookDeliveryItemListUnauthorized  %+v", 401, o.Payload)
}

func (o *WebhookDeliveryItemListUnauthorized) GetPayload() *models.ServiceStandardErrorRespModel {
	return o.Payload
}

func (o *WebhookDeliveryItemListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWebhookDeliveryItemListInternalServerError creates a WebhookDeliveryItemListInternalServerError with default headers values
func NewWebhookDeliveryItemListInternalServerError() *WebhookDeliveryItemListInternalServerError {
	return &WebhookDeliveryItemListInternalServerError{}
}

/*
WebhookDeliveryItemListInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type WebhookDeliveryItemListInternalServerError struct {
	Payload *models.ServiceStandardErrorRespModel
}

// IsSuccess returns true when this webhook delivery item list internal server error response has a 2xx status code
func (o *WebhookDeliveryItemListInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this webhook delivery item list internal server error response has a 3xx status code
func (o *WebhookDeliveryItemListInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this webhook delivery item list internal server error response has a 4xx status code
func (o *WebhookDeliveryItemListInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this webhook delivery item list internal server error response has a 5xx status code
func (o *WebhookDeliveryItemListInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this webhook delivery item list internal server error response a status code equal to that given
func (o *WebhookDeliveryItemListInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the webhook delivery item list internal server error response
func (o *WebhookDeliveryItemListInternalServerError) Code() int {
	return 500
}

func (o *WebhookDeliveryItemListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /apps/{app-slug}/outgoing-webhooks/{app-webhook-slug}/delivery-items][%d] webhookDeliveryItemListInternalServerError  %+v", 500, o.Payload)
}

func (o *WebhookDeliveryItemListInternalServerError) String() string {
	return fmt.Sprintf("[GET /apps/{app-slug}/outgoing-webhooks/{app-webhook-slug}/delivery-items][%d] webhookDeliveryItemListInternalServerError  %+v", 500, o.Payload)
}

func (o *WebhookDeliveryItemListInternalServerError) GetPayload() *models.ServiceStandardErrorRespModel {
	return o.Payload
}

func (o *WebhookDeliveryItemListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
