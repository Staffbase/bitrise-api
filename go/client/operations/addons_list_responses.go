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

// AddonsListReader is a Reader for the AddonsList structure.
type AddonsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddonsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddonsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAddonsListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAddonsListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAddonsListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAddonsListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAddonsListOK creates a AddonsListOK with default headers values
func NewAddonsListOK() *AddonsListOK {
	return &AddonsListOK{}
}

/*
AddonsListOK describes a response with status code 200, with default header values.

OK
*/
type AddonsListOK struct {
	Payload *models.V0AddonsListResponseModel
}

// IsSuccess returns true when this addons list o k response has a 2xx status code
func (o *AddonsListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this addons list o k response has a 3xx status code
func (o *AddonsListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this addons list o k response has a 4xx status code
func (o *AddonsListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this addons list o k response has a 5xx status code
func (o *AddonsListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this addons list o k response a status code equal to that given
func (o *AddonsListOK) IsCode(code int) bool {
	return code == 200
}

func (o *AddonsListOK) Error() string {
	return fmt.Sprintf("[GET /addons][%d] addonsListOK  %+v", 200, o.Payload)
}

func (o *AddonsListOK) String() string {
	return fmt.Sprintf("[GET /addons][%d] addonsListOK  %+v", 200, o.Payload)
}

func (o *AddonsListOK) GetPayload() *models.V0AddonsListResponseModel {
	return o.Payload
}

func (o *AddonsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V0AddonsListResponseModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddonsListBadRequest creates a AddonsListBadRequest with default headers values
func NewAddonsListBadRequest() *AddonsListBadRequest {
	return &AddonsListBadRequest{}
}

/*
AddonsListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type AddonsListBadRequest struct {
	Payload *models.ServiceStandardErrorRespModel
}

// IsSuccess returns true when this addons list bad request response has a 2xx status code
func (o *AddonsListBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this addons list bad request response has a 3xx status code
func (o *AddonsListBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this addons list bad request response has a 4xx status code
func (o *AddonsListBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this addons list bad request response has a 5xx status code
func (o *AddonsListBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this addons list bad request response a status code equal to that given
func (o *AddonsListBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *AddonsListBadRequest) Error() string {
	return fmt.Sprintf("[GET /addons][%d] addonsListBadRequest  %+v", 400, o.Payload)
}

func (o *AddonsListBadRequest) String() string {
	return fmt.Sprintf("[GET /addons][%d] addonsListBadRequest  %+v", 400, o.Payload)
}

func (o *AddonsListBadRequest) GetPayload() *models.ServiceStandardErrorRespModel {
	return o.Payload
}

func (o *AddonsListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddonsListUnauthorized creates a AddonsListUnauthorized with default headers values
func NewAddonsListUnauthorized() *AddonsListUnauthorized {
	return &AddonsListUnauthorized{}
}

/*
AddonsListUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type AddonsListUnauthorized struct {
	Payload *models.ServiceStandardErrorRespModel
}

// IsSuccess returns true when this addons list unauthorized response has a 2xx status code
func (o *AddonsListUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this addons list unauthorized response has a 3xx status code
func (o *AddonsListUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this addons list unauthorized response has a 4xx status code
func (o *AddonsListUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this addons list unauthorized response has a 5xx status code
func (o *AddonsListUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this addons list unauthorized response a status code equal to that given
func (o *AddonsListUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *AddonsListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /addons][%d] addonsListUnauthorized  %+v", 401, o.Payload)
}

func (o *AddonsListUnauthorized) String() string {
	return fmt.Sprintf("[GET /addons][%d] addonsListUnauthorized  %+v", 401, o.Payload)
}

func (o *AddonsListUnauthorized) GetPayload() *models.ServiceStandardErrorRespModel {
	return o.Payload
}

func (o *AddonsListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddonsListNotFound creates a AddonsListNotFound with default headers values
func NewAddonsListNotFound() *AddonsListNotFound {
	return &AddonsListNotFound{}
}

/*
AddonsListNotFound describes a response with status code 404, with default header values.

Not Found
*/
type AddonsListNotFound struct {
	Payload *models.ServiceStandardErrorRespModel
}

// IsSuccess returns true when this addons list not found response has a 2xx status code
func (o *AddonsListNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this addons list not found response has a 3xx status code
func (o *AddonsListNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this addons list not found response has a 4xx status code
func (o *AddonsListNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this addons list not found response has a 5xx status code
func (o *AddonsListNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this addons list not found response a status code equal to that given
func (o *AddonsListNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *AddonsListNotFound) Error() string {
	return fmt.Sprintf("[GET /addons][%d] addonsListNotFound  %+v", 404, o.Payload)
}

func (o *AddonsListNotFound) String() string {
	return fmt.Sprintf("[GET /addons][%d] addonsListNotFound  %+v", 404, o.Payload)
}

func (o *AddonsListNotFound) GetPayload() *models.ServiceStandardErrorRespModel {
	return o.Payload
}

func (o *AddonsListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddonsListInternalServerError creates a AddonsListInternalServerError with default headers values
func NewAddonsListInternalServerError() *AddonsListInternalServerError {
	return &AddonsListInternalServerError{}
}

/*
AddonsListInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type AddonsListInternalServerError struct {
	Payload *models.ServiceStandardErrorRespModel
}

// IsSuccess returns true when this addons list internal server error response has a 2xx status code
func (o *AddonsListInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this addons list internal server error response has a 3xx status code
func (o *AddonsListInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this addons list internal server error response has a 4xx status code
func (o *AddonsListInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this addons list internal server error response has a 5xx status code
func (o *AddonsListInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this addons list internal server error response a status code equal to that given
func (o *AddonsListInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *AddonsListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /addons][%d] addonsListInternalServerError  %+v", 500, o.Payload)
}

func (o *AddonsListInternalServerError) String() string {
	return fmt.Sprintf("[GET /addons][%d] addonsListInternalServerError  %+v", 500, o.Payload)
}

func (o *AddonsListInternalServerError) GetPayload() *models.ServiceStandardErrorRespModel {
	return o.Payload
}

func (o *AddonsListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
