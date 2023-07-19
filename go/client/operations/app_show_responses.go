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

// AppShowReader is a Reader for the AppShow structure.
type AppShowReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AppShowReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAppShowOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAppShowBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAppShowUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAppShowNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAppShowInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /apps/{app-slug}] app-show", response, response.Code())
	}
}

// NewAppShowOK creates a AppShowOK with default headers values
func NewAppShowOK() *AppShowOK {
	return &AppShowOK{}
}

/*
AppShowOK describes a response with status code 200, with default header values.

OK
*/
type AppShowOK struct {
	Payload *models.V0AppShowResponseModel
}

// IsSuccess returns true when this app show o k response has a 2xx status code
func (o *AppShowOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this app show o k response has a 3xx status code
func (o *AppShowOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this app show o k response has a 4xx status code
func (o *AppShowOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this app show o k response has a 5xx status code
func (o *AppShowOK) IsServerError() bool {
	return false
}

// IsCode returns true when this app show o k response a status code equal to that given
func (o *AppShowOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the app show o k response
func (o *AppShowOK) Code() int {
	return 200
}

func (o *AppShowOK) Error() string {
	return fmt.Sprintf("[GET /apps/{app-slug}][%d] appShowOK  %+v", 200, o.Payload)
}

func (o *AppShowOK) String() string {
	return fmt.Sprintf("[GET /apps/{app-slug}][%d] appShowOK  %+v", 200, o.Payload)
}

func (o *AppShowOK) GetPayload() *models.V0AppShowResponseModel {
	return o.Payload
}

func (o *AppShowOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V0AppShowResponseModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAppShowBadRequest creates a AppShowBadRequest with default headers values
func NewAppShowBadRequest() *AppShowBadRequest {
	return &AppShowBadRequest{}
}

/*
AppShowBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type AppShowBadRequest struct {
	Payload *models.ServiceStandardErrorRespModel
}

// IsSuccess returns true when this app show bad request response has a 2xx status code
func (o *AppShowBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this app show bad request response has a 3xx status code
func (o *AppShowBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this app show bad request response has a 4xx status code
func (o *AppShowBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this app show bad request response has a 5xx status code
func (o *AppShowBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this app show bad request response a status code equal to that given
func (o *AppShowBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the app show bad request response
func (o *AppShowBadRequest) Code() int {
	return 400
}

func (o *AppShowBadRequest) Error() string {
	return fmt.Sprintf("[GET /apps/{app-slug}][%d] appShowBadRequest  %+v", 400, o.Payload)
}

func (o *AppShowBadRequest) String() string {
	return fmt.Sprintf("[GET /apps/{app-slug}][%d] appShowBadRequest  %+v", 400, o.Payload)
}

func (o *AppShowBadRequest) GetPayload() *models.ServiceStandardErrorRespModel {
	return o.Payload
}

func (o *AppShowBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAppShowUnauthorized creates a AppShowUnauthorized with default headers values
func NewAppShowUnauthorized() *AppShowUnauthorized {
	return &AppShowUnauthorized{}
}

/*
AppShowUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type AppShowUnauthorized struct {
	Payload *models.ServiceStandardErrorRespModel
}

// IsSuccess returns true when this app show unauthorized response has a 2xx status code
func (o *AppShowUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this app show unauthorized response has a 3xx status code
func (o *AppShowUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this app show unauthorized response has a 4xx status code
func (o *AppShowUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this app show unauthorized response has a 5xx status code
func (o *AppShowUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this app show unauthorized response a status code equal to that given
func (o *AppShowUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the app show unauthorized response
func (o *AppShowUnauthorized) Code() int {
	return 401
}

func (o *AppShowUnauthorized) Error() string {
	return fmt.Sprintf("[GET /apps/{app-slug}][%d] appShowUnauthorized  %+v", 401, o.Payload)
}

func (o *AppShowUnauthorized) String() string {
	return fmt.Sprintf("[GET /apps/{app-slug}][%d] appShowUnauthorized  %+v", 401, o.Payload)
}

func (o *AppShowUnauthorized) GetPayload() *models.ServiceStandardErrorRespModel {
	return o.Payload
}

func (o *AppShowUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAppShowNotFound creates a AppShowNotFound with default headers values
func NewAppShowNotFound() *AppShowNotFound {
	return &AppShowNotFound{}
}

/*
AppShowNotFound describes a response with status code 404, with default header values.

Not Found
*/
type AppShowNotFound struct {
	Payload *models.ServiceStandardErrorRespModel
}

// IsSuccess returns true when this app show not found response has a 2xx status code
func (o *AppShowNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this app show not found response has a 3xx status code
func (o *AppShowNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this app show not found response has a 4xx status code
func (o *AppShowNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this app show not found response has a 5xx status code
func (o *AppShowNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this app show not found response a status code equal to that given
func (o *AppShowNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the app show not found response
func (o *AppShowNotFound) Code() int {
	return 404
}

func (o *AppShowNotFound) Error() string {
	return fmt.Sprintf("[GET /apps/{app-slug}][%d] appShowNotFound  %+v", 404, o.Payload)
}

func (o *AppShowNotFound) String() string {
	return fmt.Sprintf("[GET /apps/{app-slug}][%d] appShowNotFound  %+v", 404, o.Payload)
}

func (o *AppShowNotFound) GetPayload() *models.ServiceStandardErrorRespModel {
	return o.Payload
}

func (o *AppShowNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAppShowInternalServerError creates a AppShowInternalServerError with default headers values
func NewAppShowInternalServerError() *AppShowInternalServerError {
	return &AppShowInternalServerError{}
}

/*
AppShowInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type AppShowInternalServerError struct {
	Payload *models.ServiceStandardErrorRespModel
}

// IsSuccess returns true when this app show internal server error response has a 2xx status code
func (o *AppShowInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this app show internal server error response has a 3xx status code
func (o *AppShowInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this app show internal server error response has a 4xx status code
func (o *AppShowInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this app show internal server error response has a 5xx status code
func (o *AppShowInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this app show internal server error response a status code equal to that given
func (o *AppShowInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the app show internal server error response
func (o *AppShowInternalServerError) Code() int {
	return 500
}

func (o *AppShowInternalServerError) Error() string {
	return fmt.Sprintf("[GET /apps/{app-slug}][%d] appShowInternalServerError  %+v", 500, o.Payload)
}

func (o *AppShowInternalServerError) String() string {
	return fmt.Sprintf("[GET /apps/{app-slug}][%d] appShowInternalServerError  %+v", 500, o.Payload)
}

func (o *AppShowInternalServerError) GetPayload() *models.ServiceStandardErrorRespModel {
	return o.Payload
}

func (o *AppShowInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
