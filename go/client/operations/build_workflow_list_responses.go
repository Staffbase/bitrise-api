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

// BuildWorkflowListReader is a Reader for the BuildWorkflowList structure.
type BuildWorkflowListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BuildWorkflowListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBuildWorkflowListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewBuildWorkflowListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewBuildWorkflowListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewBuildWorkflowListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewBuildWorkflowListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /apps/{app-slug}/build-workflows] build-workflow-list", response, response.Code())
	}
}

// NewBuildWorkflowListOK creates a BuildWorkflowListOK with default headers values
func NewBuildWorkflowListOK() *BuildWorkflowListOK {
	return &BuildWorkflowListOK{}
}

/*
BuildWorkflowListOK describes a response with status code 200, with default header values.

OK
*/
type BuildWorkflowListOK struct {
	Payload *models.V0BuildWorkflowListResponseModel
}

// IsSuccess returns true when this build workflow list o k response has a 2xx status code
func (o *BuildWorkflowListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this build workflow list o k response has a 3xx status code
func (o *BuildWorkflowListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this build workflow list o k response has a 4xx status code
func (o *BuildWorkflowListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this build workflow list o k response has a 5xx status code
func (o *BuildWorkflowListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this build workflow list o k response a status code equal to that given
func (o *BuildWorkflowListOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the build workflow list o k response
func (o *BuildWorkflowListOK) Code() int {
	return 200
}

func (o *BuildWorkflowListOK) Error() string {
	return fmt.Sprintf("[GET /apps/{app-slug}/build-workflows][%d] buildWorkflowListOK  %+v", 200, o.Payload)
}

func (o *BuildWorkflowListOK) String() string {
	return fmt.Sprintf("[GET /apps/{app-slug}/build-workflows][%d] buildWorkflowListOK  %+v", 200, o.Payload)
}

func (o *BuildWorkflowListOK) GetPayload() *models.V0BuildWorkflowListResponseModel {
	return o.Payload
}

func (o *BuildWorkflowListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V0BuildWorkflowListResponseModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBuildWorkflowListBadRequest creates a BuildWorkflowListBadRequest with default headers values
func NewBuildWorkflowListBadRequest() *BuildWorkflowListBadRequest {
	return &BuildWorkflowListBadRequest{}
}

/*
BuildWorkflowListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type BuildWorkflowListBadRequest struct {
	Payload *models.ServiceStandardErrorRespModel
}

// IsSuccess returns true when this build workflow list bad request response has a 2xx status code
func (o *BuildWorkflowListBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this build workflow list bad request response has a 3xx status code
func (o *BuildWorkflowListBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this build workflow list bad request response has a 4xx status code
func (o *BuildWorkflowListBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this build workflow list bad request response has a 5xx status code
func (o *BuildWorkflowListBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this build workflow list bad request response a status code equal to that given
func (o *BuildWorkflowListBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the build workflow list bad request response
func (o *BuildWorkflowListBadRequest) Code() int {
	return 400
}

func (o *BuildWorkflowListBadRequest) Error() string {
	return fmt.Sprintf("[GET /apps/{app-slug}/build-workflows][%d] buildWorkflowListBadRequest  %+v", 400, o.Payload)
}

func (o *BuildWorkflowListBadRequest) String() string {
	return fmt.Sprintf("[GET /apps/{app-slug}/build-workflows][%d] buildWorkflowListBadRequest  %+v", 400, o.Payload)
}

func (o *BuildWorkflowListBadRequest) GetPayload() *models.ServiceStandardErrorRespModel {
	return o.Payload
}

func (o *BuildWorkflowListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBuildWorkflowListUnauthorized creates a BuildWorkflowListUnauthorized with default headers values
func NewBuildWorkflowListUnauthorized() *BuildWorkflowListUnauthorized {
	return &BuildWorkflowListUnauthorized{}
}

/*
BuildWorkflowListUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type BuildWorkflowListUnauthorized struct {
	Payload *models.ServiceStandardErrorRespModel
}

// IsSuccess returns true when this build workflow list unauthorized response has a 2xx status code
func (o *BuildWorkflowListUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this build workflow list unauthorized response has a 3xx status code
func (o *BuildWorkflowListUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this build workflow list unauthorized response has a 4xx status code
func (o *BuildWorkflowListUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this build workflow list unauthorized response has a 5xx status code
func (o *BuildWorkflowListUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this build workflow list unauthorized response a status code equal to that given
func (o *BuildWorkflowListUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the build workflow list unauthorized response
func (o *BuildWorkflowListUnauthorized) Code() int {
	return 401
}

func (o *BuildWorkflowListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /apps/{app-slug}/build-workflows][%d] buildWorkflowListUnauthorized  %+v", 401, o.Payload)
}

func (o *BuildWorkflowListUnauthorized) String() string {
	return fmt.Sprintf("[GET /apps/{app-slug}/build-workflows][%d] buildWorkflowListUnauthorized  %+v", 401, o.Payload)
}

func (o *BuildWorkflowListUnauthorized) GetPayload() *models.ServiceStandardErrorRespModel {
	return o.Payload
}

func (o *BuildWorkflowListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBuildWorkflowListNotFound creates a BuildWorkflowListNotFound with default headers values
func NewBuildWorkflowListNotFound() *BuildWorkflowListNotFound {
	return &BuildWorkflowListNotFound{}
}

/*
BuildWorkflowListNotFound describes a response with status code 404, with default header values.

Not Found
*/
type BuildWorkflowListNotFound struct {
	Payload *models.ServiceStandardErrorRespModel
}

// IsSuccess returns true when this build workflow list not found response has a 2xx status code
func (o *BuildWorkflowListNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this build workflow list not found response has a 3xx status code
func (o *BuildWorkflowListNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this build workflow list not found response has a 4xx status code
func (o *BuildWorkflowListNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this build workflow list not found response has a 5xx status code
func (o *BuildWorkflowListNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this build workflow list not found response a status code equal to that given
func (o *BuildWorkflowListNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the build workflow list not found response
func (o *BuildWorkflowListNotFound) Code() int {
	return 404
}

func (o *BuildWorkflowListNotFound) Error() string {
	return fmt.Sprintf("[GET /apps/{app-slug}/build-workflows][%d] buildWorkflowListNotFound  %+v", 404, o.Payload)
}

func (o *BuildWorkflowListNotFound) String() string {
	return fmt.Sprintf("[GET /apps/{app-slug}/build-workflows][%d] buildWorkflowListNotFound  %+v", 404, o.Payload)
}

func (o *BuildWorkflowListNotFound) GetPayload() *models.ServiceStandardErrorRespModel {
	return o.Payload
}

func (o *BuildWorkflowListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBuildWorkflowListInternalServerError creates a BuildWorkflowListInternalServerError with default headers values
func NewBuildWorkflowListInternalServerError() *BuildWorkflowListInternalServerError {
	return &BuildWorkflowListInternalServerError{}
}

/*
BuildWorkflowListInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type BuildWorkflowListInternalServerError struct {
	Payload *models.ServiceStandardErrorRespModel
}

// IsSuccess returns true when this build workflow list internal server error response has a 2xx status code
func (o *BuildWorkflowListInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this build workflow list internal server error response has a 3xx status code
func (o *BuildWorkflowListInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this build workflow list internal server error response has a 4xx status code
func (o *BuildWorkflowListInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this build workflow list internal server error response has a 5xx status code
func (o *BuildWorkflowListInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this build workflow list internal server error response a status code equal to that given
func (o *BuildWorkflowListInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the build workflow list internal server error response
func (o *BuildWorkflowListInternalServerError) Code() int {
	return 500
}

func (o *BuildWorkflowListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /apps/{app-slug}/build-workflows][%d] buildWorkflowListInternalServerError  %+v", 500, o.Payload)
}

func (o *BuildWorkflowListInternalServerError) String() string {
	return fmt.Sprintf("[GET /apps/{app-slug}/build-workflows][%d] buildWorkflowListInternalServerError  %+v", 500, o.Payload)
}

func (o *BuildWorkflowListInternalServerError) GetPayload() *models.ServiceStandardErrorRespModel {
	return o.Payload
}

func (o *BuildWorkflowListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
