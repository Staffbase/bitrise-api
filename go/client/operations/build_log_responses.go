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

// BuildLogReader is a Reader for the BuildLog structure.
type BuildLogReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BuildLogReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 400:
		result := NewBuildLogBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewBuildLogUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewBuildLogNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewBuildLogInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewBuildLogBadRequest creates a BuildLogBadRequest with default headers values
func NewBuildLogBadRequest() *BuildLogBadRequest {
	return &BuildLogBadRequest{}
}

/*
BuildLogBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type BuildLogBadRequest struct {
	Payload *models.ServiceStandardErrorRespModel
}

// IsSuccess returns true when this build log bad request response has a 2xx status code
func (o *BuildLogBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this build log bad request response has a 3xx status code
func (o *BuildLogBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this build log bad request response has a 4xx status code
func (o *BuildLogBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this build log bad request response has a 5xx status code
func (o *BuildLogBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this build log bad request response a status code equal to that given
func (o *BuildLogBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *BuildLogBadRequest) Error() string {
	return fmt.Sprintf("[GET /apps/{app-slug}/builds/{build-slug}/log][%d] buildLogBadRequest  %+v", 400, o.Payload)
}

func (o *BuildLogBadRequest) String() string {
	return fmt.Sprintf("[GET /apps/{app-slug}/builds/{build-slug}/log][%d] buildLogBadRequest  %+v", 400, o.Payload)
}

func (o *BuildLogBadRequest) GetPayload() *models.ServiceStandardErrorRespModel {
	return o.Payload
}

func (o *BuildLogBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBuildLogUnauthorized creates a BuildLogUnauthorized with default headers values
func NewBuildLogUnauthorized() *BuildLogUnauthorized {
	return &BuildLogUnauthorized{}
}

/*
BuildLogUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type BuildLogUnauthorized struct {
	Payload *models.ServiceStandardErrorRespModel
}

// IsSuccess returns true when this build log unauthorized response has a 2xx status code
func (o *BuildLogUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this build log unauthorized response has a 3xx status code
func (o *BuildLogUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this build log unauthorized response has a 4xx status code
func (o *BuildLogUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this build log unauthorized response has a 5xx status code
func (o *BuildLogUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this build log unauthorized response a status code equal to that given
func (o *BuildLogUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *BuildLogUnauthorized) Error() string {
	return fmt.Sprintf("[GET /apps/{app-slug}/builds/{build-slug}/log][%d] buildLogUnauthorized  %+v", 401, o.Payload)
}

func (o *BuildLogUnauthorized) String() string {
	return fmt.Sprintf("[GET /apps/{app-slug}/builds/{build-slug}/log][%d] buildLogUnauthorized  %+v", 401, o.Payload)
}

func (o *BuildLogUnauthorized) GetPayload() *models.ServiceStandardErrorRespModel {
	return o.Payload
}

func (o *BuildLogUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBuildLogNotFound creates a BuildLogNotFound with default headers values
func NewBuildLogNotFound() *BuildLogNotFound {
	return &BuildLogNotFound{}
}

/*
BuildLogNotFound describes a response with status code 404, with default header values.

Not Found
*/
type BuildLogNotFound struct {
	Payload *models.ServiceStandardErrorRespModel
}

// IsSuccess returns true when this build log not found response has a 2xx status code
func (o *BuildLogNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this build log not found response has a 3xx status code
func (o *BuildLogNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this build log not found response has a 4xx status code
func (o *BuildLogNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this build log not found response has a 5xx status code
func (o *BuildLogNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this build log not found response a status code equal to that given
func (o *BuildLogNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *BuildLogNotFound) Error() string {
	return fmt.Sprintf("[GET /apps/{app-slug}/builds/{build-slug}/log][%d] buildLogNotFound  %+v", 404, o.Payload)
}

func (o *BuildLogNotFound) String() string {
	return fmt.Sprintf("[GET /apps/{app-slug}/builds/{build-slug}/log][%d] buildLogNotFound  %+v", 404, o.Payload)
}

func (o *BuildLogNotFound) GetPayload() *models.ServiceStandardErrorRespModel {
	return o.Payload
}

func (o *BuildLogNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBuildLogInternalServerError creates a BuildLogInternalServerError with default headers values
func NewBuildLogInternalServerError() *BuildLogInternalServerError {
	return &BuildLogInternalServerError{}
}

/*
BuildLogInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type BuildLogInternalServerError struct {
	Payload *models.ServiceStandardErrorRespModel
}

// IsSuccess returns true when this build log internal server error response has a 2xx status code
func (o *BuildLogInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this build log internal server error response has a 3xx status code
func (o *BuildLogInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this build log internal server error response has a 4xx status code
func (o *BuildLogInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this build log internal server error response has a 5xx status code
func (o *BuildLogInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this build log internal server error response a status code equal to that given
func (o *BuildLogInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *BuildLogInternalServerError) Error() string {
	return fmt.Sprintf("[GET /apps/{app-slug}/builds/{build-slug}/log][%d] buildLogInternalServerError  %+v", 500, o.Payload)
}

func (o *BuildLogInternalServerError) String() string {
	return fmt.Sprintf("[GET /apps/{app-slug}/builds/{build-slug}/log][%d] buildLogInternalServerError  %+v", 500, o.Payload)
}

func (o *BuildLogInternalServerError) GetPayload() *models.ServiceStandardErrorRespModel {
	return o.Payload
}

func (o *BuildLogInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
