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

// UserPlanReader is a Reader for the UserPlan structure.
type UserPlanReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserPlanReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUserPlanOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUserPlanBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUserPlanUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUserPlanNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUserPlanInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /me/plan] user-plan", response, response.Code())
	}
}

// NewUserPlanOK creates a UserPlanOK with default headers values
func NewUserPlanOK() *UserPlanOK {
	return &UserPlanOK{}
}

/*
UserPlanOK describes a response with status code 200, with default header values.

OK
*/
type UserPlanOK struct {
	Payload *models.V0UserPlanRespModel
}

// IsSuccess returns true when this user plan o k response has a 2xx status code
func (o *UserPlanOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this user plan o k response has a 3xx status code
func (o *UserPlanOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user plan o k response has a 4xx status code
func (o *UserPlanOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this user plan o k response has a 5xx status code
func (o *UserPlanOK) IsServerError() bool {
	return false
}

// IsCode returns true when this user plan o k response a status code equal to that given
func (o *UserPlanOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the user plan o k response
func (o *UserPlanOK) Code() int {
	return 200
}

func (o *UserPlanOK) Error() string {
	return fmt.Sprintf("[GET /me/plan][%d] userPlanOK  %+v", 200, o.Payload)
}

func (o *UserPlanOK) String() string {
	return fmt.Sprintf("[GET /me/plan][%d] userPlanOK  %+v", 200, o.Payload)
}

func (o *UserPlanOK) GetPayload() *models.V0UserPlanRespModel {
	return o.Payload
}

func (o *UserPlanOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V0UserPlanRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserPlanBadRequest creates a UserPlanBadRequest with default headers values
func NewUserPlanBadRequest() *UserPlanBadRequest {
	return &UserPlanBadRequest{}
}

/*
UserPlanBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UserPlanBadRequest struct {
	Payload *models.ServiceStandardErrorRespModel
}

// IsSuccess returns true when this user plan bad request response has a 2xx status code
func (o *UserPlanBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user plan bad request response has a 3xx status code
func (o *UserPlanBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user plan bad request response has a 4xx status code
func (o *UserPlanBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this user plan bad request response has a 5xx status code
func (o *UserPlanBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this user plan bad request response a status code equal to that given
func (o *UserPlanBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the user plan bad request response
func (o *UserPlanBadRequest) Code() int {
	return 400
}

func (o *UserPlanBadRequest) Error() string {
	return fmt.Sprintf("[GET /me/plan][%d] userPlanBadRequest  %+v", 400, o.Payload)
}

func (o *UserPlanBadRequest) String() string {
	return fmt.Sprintf("[GET /me/plan][%d] userPlanBadRequest  %+v", 400, o.Payload)
}

func (o *UserPlanBadRequest) GetPayload() *models.ServiceStandardErrorRespModel {
	return o.Payload
}

func (o *UserPlanBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserPlanUnauthorized creates a UserPlanUnauthorized with default headers values
func NewUserPlanUnauthorized() *UserPlanUnauthorized {
	return &UserPlanUnauthorized{}
}

/*
UserPlanUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type UserPlanUnauthorized struct {
	Payload *models.ServiceStandardErrorRespModel
}

// IsSuccess returns true when this user plan unauthorized response has a 2xx status code
func (o *UserPlanUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user plan unauthorized response has a 3xx status code
func (o *UserPlanUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user plan unauthorized response has a 4xx status code
func (o *UserPlanUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this user plan unauthorized response has a 5xx status code
func (o *UserPlanUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this user plan unauthorized response a status code equal to that given
func (o *UserPlanUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the user plan unauthorized response
func (o *UserPlanUnauthorized) Code() int {
	return 401
}

func (o *UserPlanUnauthorized) Error() string {
	return fmt.Sprintf("[GET /me/plan][%d] userPlanUnauthorized  %+v", 401, o.Payload)
}

func (o *UserPlanUnauthorized) String() string {
	return fmt.Sprintf("[GET /me/plan][%d] userPlanUnauthorized  %+v", 401, o.Payload)
}

func (o *UserPlanUnauthorized) GetPayload() *models.ServiceStandardErrorRespModel {
	return o.Payload
}

func (o *UserPlanUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserPlanNotFound creates a UserPlanNotFound with default headers values
func NewUserPlanNotFound() *UserPlanNotFound {
	return &UserPlanNotFound{}
}

/*
UserPlanNotFound describes a response with status code 404, with default header values.

Not Found
*/
type UserPlanNotFound struct {
	Payload *models.ServiceStandardErrorRespModel
}

// IsSuccess returns true when this user plan not found response has a 2xx status code
func (o *UserPlanNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user plan not found response has a 3xx status code
func (o *UserPlanNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user plan not found response has a 4xx status code
func (o *UserPlanNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this user plan not found response has a 5xx status code
func (o *UserPlanNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this user plan not found response a status code equal to that given
func (o *UserPlanNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the user plan not found response
func (o *UserPlanNotFound) Code() int {
	return 404
}

func (o *UserPlanNotFound) Error() string {
	return fmt.Sprintf("[GET /me/plan][%d] userPlanNotFound  %+v", 404, o.Payload)
}

func (o *UserPlanNotFound) String() string {
	return fmt.Sprintf("[GET /me/plan][%d] userPlanNotFound  %+v", 404, o.Payload)
}

func (o *UserPlanNotFound) GetPayload() *models.ServiceStandardErrorRespModel {
	return o.Payload
}

func (o *UserPlanNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserPlanInternalServerError creates a UserPlanInternalServerError with default headers values
func NewUserPlanInternalServerError() *UserPlanInternalServerError {
	return &UserPlanInternalServerError{}
}

/*
UserPlanInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type UserPlanInternalServerError struct {
	Payload *models.ServiceStandardErrorRespModel
}

// IsSuccess returns true when this user plan internal server error response has a 2xx status code
func (o *UserPlanInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user plan internal server error response has a 3xx status code
func (o *UserPlanInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user plan internal server error response has a 4xx status code
func (o *UserPlanInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this user plan internal server error response has a 5xx status code
func (o *UserPlanInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this user plan internal server error response a status code equal to that given
func (o *UserPlanInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the user plan internal server error response
func (o *UserPlanInternalServerError) Code() int {
	return 500
}

func (o *UserPlanInternalServerError) Error() string {
	return fmt.Sprintf("[GET /me/plan][%d] userPlanInternalServerError  %+v", 500, o.Payload)
}

func (o *UserPlanInternalServerError) String() string {
	return fmt.Sprintf("[GET /me/plan][%d] userPlanInternalServerError  %+v", 500, o.Payload)
}

func (o *UserPlanInternalServerError) GetPayload() *models.ServiceStandardErrorRespModel {
	return o.Payload
}

func (o *UserPlanInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
