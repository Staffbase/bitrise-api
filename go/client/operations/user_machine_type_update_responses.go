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

// UserMachineTypeUpdateReader is a Reader for the UserMachineTypeUpdate structure.
type UserMachineTypeUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserMachineTypeUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUserMachineTypeUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUserMachineTypeUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUserMachineTypeUpdateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUserMachineTypeUpdateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUserMachineTypeUpdateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PATCH /user/{user-slug}/apps/machine_types] user-machine-type-update", response, response.Code())
	}
}

// NewUserMachineTypeUpdateOK creates a UserMachineTypeUpdateOK with default headers values
func NewUserMachineTypeUpdateOK() *UserMachineTypeUpdateOK {
	return &UserMachineTypeUpdateOK{}
}

/*
UserMachineTypeUpdateOK describes a response with status code 200, with default header values.

OK
*/
type UserMachineTypeUpdateOK struct {
	Payload *models.V0OrganizationUpdateMachineTypeResponse
}

// IsSuccess returns true when this user machine type update o k response has a 2xx status code
func (o *UserMachineTypeUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this user machine type update o k response has a 3xx status code
func (o *UserMachineTypeUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user machine type update o k response has a 4xx status code
func (o *UserMachineTypeUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this user machine type update o k response has a 5xx status code
func (o *UserMachineTypeUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this user machine type update o k response a status code equal to that given
func (o *UserMachineTypeUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the user machine type update o k response
func (o *UserMachineTypeUpdateOK) Code() int {
	return 200
}

func (o *UserMachineTypeUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /user/{user-slug}/apps/machine_types][%d] userMachineTypeUpdateOK  %+v", 200, o.Payload)
}

func (o *UserMachineTypeUpdateOK) String() string {
	return fmt.Sprintf("[PATCH /user/{user-slug}/apps/machine_types][%d] userMachineTypeUpdateOK  %+v", 200, o.Payload)
}

func (o *UserMachineTypeUpdateOK) GetPayload() *models.V0OrganizationUpdateMachineTypeResponse {
	return o.Payload
}

func (o *UserMachineTypeUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V0OrganizationUpdateMachineTypeResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserMachineTypeUpdateBadRequest creates a UserMachineTypeUpdateBadRequest with default headers values
func NewUserMachineTypeUpdateBadRequest() *UserMachineTypeUpdateBadRequest {
	return &UserMachineTypeUpdateBadRequest{}
}

/*
UserMachineTypeUpdateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UserMachineTypeUpdateBadRequest struct {
	Payload *models.ServiceStandardErrorRespModel
}

// IsSuccess returns true when this user machine type update bad request response has a 2xx status code
func (o *UserMachineTypeUpdateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user machine type update bad request response has a 3xx status code
func (o *UserMachineTypeUpdateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user machine type update bad request response has a 4xx status code
func (o *UserMachineTypeUpdateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this user machine type update bad request response has a 5xx status code
func (o *UserMachineTypeUpdateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this user machine type update bad request response a status code equal to that given
func (o *UserMachineTypeUpdateBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the user machine type update bad request response
func (o *UserMachineTypeUpdateBadRequest) Code() int {
	return 400
}

func (o *UserMachineTypeUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /user/{user-slug}/apps/machine_types][%d] userMachineTypeUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *UserMachineTypeUpdateBadRequest) String() string {
	return fmt.Sprintf("[PATCH /user/{user-slug}/apps/machine_types][%d] userMachineTypeUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *UserMachineTypeUpdateBadRequest) GetPayload() *models.ServiceStandardErrorRespModel {
	return o.Payload
}

func (o *UserMachineTypeUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserMachineTypeUpdateUnauthorized creates a UserMachineTypeUpdateUnauthorized with default headers values
func NewUserMachineTypeUpdateUnauthorized() *UserMachineTypeUpdateUnauthorized {
	return &UserMachineTypeUpdateUnauthorized{}
}

/*
UserMachineTypeUpdateUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type UserMachineTypeUpdateUnauthorized struct {
	Payload *models.ServiceStandardErrorRespModel
}

// IsSuccess returns true when this user machine type update unauthorized response has a 2xx status code
func (o *UserMachineTypeUpdateUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user machine type update unauthorized response has a 3xx status code
func (o *UserMachineTypeUpdateUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user machine type update unauthorized response has a 4xx status code
func (o *UserMachineTypeUpdateUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this user machine type update unauthorized response has a 5xx status code
func (o *UserMachineTypeUpdateUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this user machine type update unauthorized response a status code equal to that given
func (o *UserMachineTypeUpdateUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the user machine type update unauthorized response
func (o *UserMachineTypeUpdateUnauthorized) Code() int {
	return 401
}

func (o *UserMachineTypeUpdateUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /user/{user-slug}/apps/machine_types][%d] userMachineTypeUpdateUnauthorized  %+v", 401, o.Payload)
}

func (o *UserMachineTypeUpdateUnauthorized) String() string {
	return fmt.Sprintf("[PATCH /user/{user-slug}/apps/machine_types][%d] userMachineTypeUpdateUnauthorized  %+v", 401, o.Payload)
}

func (o *UserMachineTypeUpdateUnauthorized) GetPayload() *models.ServiceStandardErrorRespModel {
	return o.Payload
}

func (o *UserMachineTypeUpdateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserMachineTypeUpdateNotFound creates a UserMachineTypeUpdateNotFound with default headers values
func NewUserMachineTypeUpdateNotFound() *UserMachineTypeUpdateNotFound {
	return &UserMachineTypeUpdateNotFound{}
}

/*
UserMachineTypeUpdateNotFound describes a response with status code 404, with default header values.

Not Found
*/
type UserMachineTypeUpdateNotFound struct {
	Payload *models.ServiceStandardErrorRespModel
}

// IsSuccess returns true when this user machine type update not found response has a 2xx status code
func (o *UserMachineTypeUpdateNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user machine type update not found response has a 3xx status code
func (o *UserMachineTypeUpdateNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user machine type update not found response has a 4xx status code
func (o *UserMachineTypeUpdateNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this user machine type update not found response has a 5xx status code
func (o *UserMachineTypeUpdateNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this user machine type update not found response a status code equal to that given
func (o *UserMachineTypeUpdateNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the user machine type update not found response
func (o *UserMachineTypeUpdateNotFound) Code() int {
	return 404
}

func (o *UserMachineTypeUpdateNotFound) Error() string {
	return fmt.Sprintf("[PATCH /user/{user-slug}/apps/machine_types][%d] userMachineTypeUpdateNotFound  %+v", 404, o.Payload)
}

func (o *UserMachineTypeUpdateNotFound) String() string {
	return fmt.Sprintf("[PATCH /user/{user-slug}/apps/machine_types][%d] userMachineTypeUpdateNotFound  %+v", 404, o.Payload)
}

func (o *UserMachineTypeUpdateNotFound) GetPayload() *models.ServiceStandardErrorRespModel {
	return o.Payload
}

func (o *UserMachineTypeUpdateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserMachineTypeUpdateInternalServerError creates a UserMachineTypeUpdateInternalServerError with default headers values
func NewUserMachineTypeUpdateInternalServerError() *UserMachineTypeUpdateInternalServerError {
	return &UserMachineTypeUpdateInternalServerError{}
}

/*
UserMachineTypeUpdateInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type UserMachineTypeUpdateInternalServerError struct {
	Payload *models.ServiceStandardErrorRespModel
}

// IsSuccess returns true when this user machine type update internal server error response has a 2xx status code
func (o *UserMachineTypeUpdateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user machine type update internal server error response has a 3xx status code
func (o *UserMachineTypeUpdateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user machine type update internal server error response has a 4xx status code
func (o *UserMachineTypeUpdateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this user machine type update internal server error response has a 5xx status code
func (o *UserMachineTypeUpdateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this user machine type update internal server error response a status code equal to that given
func (o *UserMachineTypeUpdateInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the user machine type update internal server error response
func (o *UserMachineTypeUpdateInternalServerError) Code() int {
	return 500
}

func (o *UserMachineTypeUpdateInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /user/{user-slug}/apps/machine_types][%d] userMachineTypeUpdateInternalServerError  %+v", 500, o.Payload)
}

func (o *UserMachineTypeUpdateInternalServerError) String() string {
	return fmt.Sprintf("[PATCH /user/{user-slug}/apps/machine_types][%d] userMachineTypeUpdateInternalServerError  %+v", 500, o.Payload)
}

func (o *UserMachineTypeUpdateInternalServerError) GetPayload() *models.ServiceStandardErrorRespModel {
	return o.Payload
}

func (o *UserMachineTypeUpdateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
