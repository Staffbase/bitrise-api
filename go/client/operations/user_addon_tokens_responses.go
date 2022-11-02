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

// UserAddonTokensReader is a Reader for the UserAddonTokens structure.
type UserAddonTokensReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserAddonTokensReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUserAddonTokensOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUserAddonTokensBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUserAddonTokensUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUserAddonTokensNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUserAddonTokensInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUserAddonTokensOK creates a UserAddonTokensOK with default headers values
func NewUserAddonTokensOK() *UserAddonTokensOK {
	return &UserAddonTokensOK{}
}

/*
UserAddonTokensOK describes a response with status code 200, with default header values.

OK
*/
type UserAddonTokensOK struct {
	Payload []string
}

// IsSuccess returns true when this user addon tokens o k response has a 2xx status code
func (o *UserAddonTokensOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this user addon tokens o k response has a 3xx status code
func (o *UserAddonTokensOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user addon tokens o k response has a 4xx status code
func (o *UserAddonTokensOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this user addon tokens o k response has a 5xx status code
func (o *UserAddonTokensOK) IsServerError() bool {
	return false
}

// IsCode returns true when this user addon tokens o k response a status code equal to that given
func (o *UserAddonTokensOK) IsCode(code int) bool {
	return code == 200
}

func (o *UserAddonTokensOK) Error() string {
	return fmt.Sprintf("[GET /me/addon-tokens][%d] userAddonTokensOK  %+v", 200, o.Payload)
}

func (o *UserAddonTokensOK) String() string {
	return fmt.Sprintf("[GET /me/addon-tokens][%d] userAddonTokensOK  %+v", 200, o.Payload)
}

func (o *UserAddonTokensOK) GetPayload() []string {
	return o.Payload
}

func (o *UserAddonTokensOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserAddonTokensBadRequest creates a UserAddonTokensBadRequest with default headers values
func NewUserAddonTokensBadRequest() *UserAddonTokensBadRequest {
	return &UserAddonTokensBadRequest{}
}

/*
UserAddonTokensBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UserAddonTokensBadRequest struct {
	Payload *models.ServiceStandardErrorRespModel
}

// IsSuccess returns true when this user addon tokens bad request response has a 2xx status code
func (o *UserAddonTokensBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user addon tokens bad request response has a 3xx status code
func (o *UserAddonTokensBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user addon tokens bad request response has a 4xx status code
func (o *UserAddonTokensBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this user addon tokens bad request response has a 5xx status code
func (o *UserAddonTokensBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this user addon tokens bad request response a status code equal to that given
func (o *UserAddonTokensBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *UserAddonTokensBadRequest) Error() string {
	return fmt.Sprintf("[GET /me/addon-tokens][%d] userAddonTokensBadRequest  %+v", 400, o.Payload)
}

func (o *UserAddonTokensBadRequest) String() string {
	return fmt.Sprintf("[GET /me/addon-tokens][%d] userAddonTokensBadRequest  %+v", 400, o.Payload)
}

func (o *UserAddonTokensBadRequest) GetPayload() *models.ServiceStandardErrorRespModel {
	return o.Payload
}

func (o *UserAddonTokensBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserAddonTokensUnauthorized creates a UserAddonTokensUnauthorized with default headers values
func NewUserAddonTokensUnauthorized() *UserAddonTokensUnauthorized {
	return &UserAddonTokensUnauthorized{}
}

/*
UserAddonTokensUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type UserAddonTokensUnauthorized struct {
	Payload *models.ServiceStandardErrorRespModel
}

// IsSuccess returns true when this user addon tokens unauthorized response has a 2xx status code
func (o *UserAddonTokensUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user addon tokens unauthorized response has a 3xx status code
func (o *UserAddonTokensUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user addon tokens unauthorized response has a 4xx status code
func (o *UserAddonTokensUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this user addon tokens unauthorized response has a 5xx status code
func (o *UserAddonTokensUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this user addon tokens unauthorized response a status code equal to that given
func (o *UserAddonTokensUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *UserAddonTokensUnauthorized) Error() string {
	return fmt.Sprintf("[GET /me/addon-tokens][%d] userAddonTokensUnauthorized  %+v", 401, o.Payload)
}

func (o *UserAddonTokensUnauthorized) String() string {
	return fmt.Sprintf("[GET /me/addon-tokens][%d] userAddonTokensUnauthorized  %+v", 401, o.Payload)
}

func (o *UserAddonTokensUnauthorized) GetPayload() *models.ServiceStandardErrorRespModel {
	return o.Payload
}

func (o *UserAddonTokensUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserAddonTokensNotFound creates a UserAddonTokensNotFound with default headers values
func NewUserAddonTokensNotFound() *UserAddonTokensNotFound {
	return &UserAddonTokensNotFound{}
}

/*
UserAddonTokensNotFound describes a response with status code 404, with default header values.

Not Found
*/
type UserAddonTokensNotFound struct {
	Payload *models.ServiceStandardErrorRespModel
}

// IsSuccess returns true when this user addon tokens not found response has a 2xx status code
func (o *UserAddonTokensNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user addon tokens not found response has a 3xx status code
func (o *UserAddonTokensNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user addon tokens not found response has a 4xx status code
func (o *UserAddonTokensNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this user addon tokens not found response has a 5xx status code
func (o *UserAddonTokensNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this user addon tokens not found response a status code equal to that given
func (o *UserAddonTokensNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *UserAddonTokensNotFound) Error() string {
	return fmt.Sprintf("[GET /me/addon-tokens][%d] userAddonTokensNotFound  %+v", 404, o.Payload)
}

func (o *UserAddonTokensNotFound) String() string {
	return fmt.Sprintf("[GET /me/addon-tokens][%d] userAddonTokensNotFound  %+v", 404, o.Payload)
}

func (o *UserAddonTokensNotFound) GetPayload() *models.ServiceStandardErrorRespModel {
	return o.Payload
}

func (o *UserAddonTokensNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserAddonTokensInternalServerError creates a UserAddonTokensInternalServerError with default headers values
func NewUserAddonTokensInternalServerError() *UserAddonTokensInternalServerError {
	return &UserAddonTokensInternalServerError{}
}

/*
UserAddonTokensInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type UserAddonTokensInternalServerError struct {
	Payload *models.ServiceStandardErrorRespModel
}

// IsSuccess returns true when this user addon tokens internal server error response has a 2xx status code
func (o *UserAddonTokensInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user addon tokens internal server error response has a 3xx status code
func (o *UserAddonTokensInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user addon tokens internal server error response has a 4xx status code
func (o *UserAddonTokensInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this user addon tokens internal server error response has a 5xx status code
func (o *UserAddonTokensInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this user addon tokens internal server error response a status code equal to that given
func (o *UserAddonTokensInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *UserAddonTokensInternalServerError) Error() string {
	return fmt.Sprintf("[GET /me/addon-tokens][%d] userAddonTokensInternalServerError  %+v", 500, o.Payload)
}

func (o *UserAddonTokensInternalServerError) String() string {
	return fmt.Sprintf("[GET /me/addon-tokens][%d] userAddonTokensInternalServerError  %+v", 500, o.Payload)
}

func (o *UserAddonTokensInternalServerError) GetPayload() *models.ServiceStandardErrorRespModel {
	return o.Payload
}

func (o *UserAddonTokensInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}