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

// SecretUpsertReader is a Reader for the SecretUpsert structure.
type SecretUpsertReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SecretUpsertReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewSecretUpsertCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewSecretUpsertNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSecretUpsertBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewSecretUpsertUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSecretUpsertNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSecretUpsertInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /apps/{app-slug}/secrets/{secret-name}] secret-upsert", response, response.Code())
	}
}

// NewSecretUpsertCreated creates a SecretUpsertCreated with default headers values
func NewSecretUpsertCreated() *SecretUpsertCreated {
	return &SecretUpsertCreated{}
}

/*
SecretUpsertCreated describes a response with status code 201, with default header values.

Created
*/
type SecretUpsertCreated struct {
}

// IsSuccess returns true when this secret upsert created response has a 2xx status code
func (o *SecretUpsertCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this secret upsert created response has a 3xx status code
func (o *SecretUpsertCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this secret upsert created response has a 4xx status code
func (o *SecretUpsertCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this secret upsert created response has a 5xx status code
func (o *SecretUpsertCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this secret upsert created response a status code equal to that given
func (o *SecretUpsertCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the secret upsert created response
func (o *SecretUpsertCreated) Code() int {
	return 201
}

func (o *SecretUpsertCreated) Error() string {
	return fmt.Sprintf("[PUT /apps/{app-slug}/secrets/{secret-name}][%d] secretUpsertCreated ", 201)
}

func (o *SecretUpsertCreated) String() string {
	return fmt.Sprintf("[PUT /apps/{app-slug}/secrets/{secret-name}][%d] secretUpsertCreated ", 201)
}

func (o *SecretUpsertCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSecretUpsertNoContent creates a SecretUpsertNoContent with default headers values
func NewSecretUpsertNoContent() *SecretUpsertNoContent {
	return &SecretUpsertNoContent{}
}

/*
SecretUpsertNoContent describes a response with status code 204, with default header values.

No Content
*/
type SecretUpsertNoContent struct {
}

// IsSuccess returns true when this secret upsert no content response has a 2xx status code
func (o *SecretUpsertNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this secret upsert no content response has a 3xx status code
func (o *SecretUpsertNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this secret upsert no content response has a 4xx status code
func (o *SecretUpsertNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this secret upsert no content response has a 5xx status code
func (o *SecretUpsertNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this secret upsert no content response a status code equal to that given
func (o *SecretUpsertNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the secret upsert no content response
func (o *SecretUpsertNoContent) Code() int {
	return 204
}

func (o *SecretUpsertNoContent) Error() string {
	return fmt.Sprintf("[PUT /apps/{app-slug}/secrets/{secret-name}][%d] secretUpsertNoContent ", 204)
}

func (o *SecretUpsertNoContent) String() string {
	return fmt.Sprintf("[PUT /apps/{app-slug}/secrets/{secret-name}][%d] secretUpsertNoContent ", 204)
}

func (o *SecretUpsertNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSecretUpsertBadRequest creates a SecretUpsertBadRequest with default headers values
func NewSecretUpsertBadRequest() *SecretUpsertBadRequest {
	return &SecretUpsertBadRequest{}
}

/*
SecretUpsertBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type SecretUpsertBadRequest struct {
	Payload *models.ServiceStandardErrorRespModel
}

// IsSuccess returns true when this secret upsert bad request response has a 2xx status code
func (o *SecretUpsertBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this secret upsert bad request response has a 3xx status code
func (o *SecretUpsertBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this secret upsert bad request response has a 4xx status code
func (o *SecretUpsertBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this secret upsert bad request response has a 5xx status code
func (o *SecretUpsertBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this secret upsert bad request response a status code equal to that given
func (o *SecretUpsertBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the secret upsert bad request response
func (o *SecretUpsertBadRequest) Code() int {
	return 400
}

func (o *SecretUpsertBadRequest) Error() string {
	return fmt.Sprintf("[PUT /apps/{app-slug}/secrets/{secret-name}][%d] secretUpsertBadRequest  %+v", 400, o.Payload)
}

func (o *SecretUpsertBadRequest) String() string {
	return fmt.Sprintf("[PUT /apps/{app-slug}/secrets/{secret-name}][%d] secretUpsertBadRequest  %+v", 400, o.Payload)
}

func (o *SecretUpsertBadRequest) GetPayload() *models.ServiceStandardErrorRespModel {
	return o.Payload
}

func (o *SecretUpsertBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSecretUpsertUnauthorized creates a SecretUpsertUnauthorized with default headers values
func NewSecretUpsertUnauthorized() *SecretUpsertUnauthorized {
	return &SecretUpsertUnauthorized{}
}

/*
SecretUpsertUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type SecretUpsertUnauthorized struct {
	Payload *models.ServiceStandardErrorRespModel
}

// IsSuccess returns true when this secret upsert unauthorized response has a 2xx status code
func (o *SecretUpsertUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this secret upsert unauthorized response has a 3xx status code
func (o *SecretUpsertUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this secret upsert unauthorized response has a 4xx status code
func (o *SecretUpsertUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this secret upsert unauthorized response has a 5xx status code
func (o *SecretUpsertUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this secret upsert unauthorized response a status code equal to that given
func (o *SecretUpsertUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the secret upsert unauthorized response
func (o *SecretUpsertUnauthorized) Code() int {
	return 401
}

func (o *SecretUpsertUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /apps/{app-slug}/secrets/{secret-name}][%d] secretUpsertUnauthorized  %+v", 401, o.Payload)
}

func (o *SecretUpsertUnauthorized) String() string {
	return fmt.Sprintf("[PUT /apps/{app-slug}/secrets/{secret-name}][%d] secretUpsertUnauthorized  %+v", 401, o.Payload)
}

func (o *SecretUpsertUnauthorized) GetPayload() *models.ServiceStandardErrorRespModel {
	return o.Payload
}

func (o *SecretUpsertUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSecretUpsertNotFound creates a SecretUpsertNotFound with default headers values
func NewSecretUpsertNotFound() *SecretUpsertNotFound {
	return &SecretUpsertNotFound{}
}

/*
SecretUpsertNotFound describes a response with status code 404, with default header values.

Not Found
*/
type SecretUpsertNotFound struct {
	Payload *models.ServiceStandardErrorRespModel
}

// IsSuccess returns true when this secret upsert not found response has a 2xx status code
func (o *SecretUpsertNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this secret upsert not found response has a 3xx status code
func (o *SecretUpsertNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this secret upsert not found response has a 4xx status code
func (o *SecretUpsertNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this secret upsert not found response has a 5xx status code
func (o *SecretUpsertNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this secret upsert not found response a status code equal to that given
func (o *SecretUpsertNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the secret upsert not found response
func (o *SecretUpsertNotFound) Code() int {
	return 404
}

func (o *SecretUpsertNotFound) Error() string {
	return fmt.Sprintf("[PUT /apps/{app-slug}/secrets/{secret-name}][%d] secretUpsertNotFound  %+v", 404, o.Payload)
}

func (o *SecretUpsertNotFound) String() string {
	return fmt.Sprintf("[PUT /apps/{app-slug}/secrets/{secret-name}][%d] secretUpsertNotFound  %+v", 404, o.Payload)
}

func (o *SecretUpsertNotFound) GetPayload() *models.ServiceStandardErrorRespModel {
	return o.Payload
}

func (o *SecretUpsertNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSecretUpsertInternalServerError creates a SecretUpsertInternalServerError with default headers values
func NewSecretUpsertInternalServerError() *SecretUpsertInternalServerError {
	return &SecretUpsertInternalServerError{}
}

/*
SecretUpsertInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type SecretUpsertInternalServerError struct {
	Payload *models.ServiceStandardErrorRespModel
}

// IsSuccess returns true when this secret upsert internal server error response has a 2xx status code
func (o *SecretUpsertInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this secret upsert internal server error response has a 3xx status code
func (o *SecretUpsertInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this secret upsert internal server error response has a 4xx status code
func (o *SecretUpsertInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this secret upsert internal server error response has a 5xx status code
func (o *SecretUpsertInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this secret upsert internal server error response a status code equal to that given
func (o *SecretUpsertInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the secret upsert internal server error response
func (o *SecretUpsertInternalServerError) Code() int {
	return 500
}

func (o *SecretUpsertInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /apps/{app-slug}/secrets/{secret-name}][%d] secretUpsertInternalServerError  %+v", 500, o.Payload)
}

func (o *SecretUpsertInternalServerError) String() string {
	return fmt.Sprintf("[PUT /apps/{app-slug}/secrets/{secret-name}][%d] secretUpsertInternalServerError  %+v", 500, o.Payload)
}

func (o *SecretUpsertInternalServerError) GetPayload() *models.ServiceStandardErrorRespModel {
	return o.Payload
}

func (o *SecretUpsertInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
