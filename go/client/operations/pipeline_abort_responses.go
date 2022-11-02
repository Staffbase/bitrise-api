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

// PipelineAbortReader is a Reader for the PipelineAbort structure.
type PipelineAbortReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PipelineAbortReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPipelineAbortOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPipelineAbortBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPipelineAbortUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPipelineAbortNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPipelineAbortInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPipelineAbortOK creates a PipelineAbortOK with default headers values
func NewPipelineAbortOK() *PipelineAbortOK {
	return &PipelineAbortOK{}
}

/*
PipelineAbortOK describes a response with status code 200, with default header values.

OK
*/
type PipelineAbortOK struct {
}

// IsSuccess returns true when this pipeline abort o k response has a 2xx status code
func (o *PipelineAbortOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pipeline abort o k response has a 3xx status code
func (o *PipelineAbortOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pipeline abort o k response has a 4xx status code
func (o *PipelineAbortOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this pipeline abort o k response has a 5xx status code
func (o *PipelineAbortOK) IsServerError() bool {
	return false
}

// IsCode returns true when this pipeline abort o k response a status code equal to that given
func (o *PipelineAbortOK) IsCode(code int) bool {
	return code == 200
}

func (o *PipelineAbortOK) Error() string {
	return fmt.Sprintf("[POST /apps/{app-slug}/pipelines/{pipeline-id}/abort][%d] pipelineAbortOK ", 200)
}

func (o *PipelineAbortOK) String() string {
	return fmt.Sprintf("[POST /apps/{app-slug}/pipelines/{pipeline-id}/abort][%d] pipelineAbortOK ", 200)
}

func (o *PipelineAbortOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPipelineAbortBadRequest creates a PipelineAbortBadRequest with default headers values
func NewPipelineAbortBadRequest() *PipelineAbortBadRequest {
	return &PipelineAbortBadRequest{}
}

/*
PipelineAbortBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PipelineAbortBadRequest struct {
	Payload *models.ServiceStandardErrorRespModel
}

// IsSuccess returns true when this pipeline abort bad request response has a 2xx status code
func (o *PipelineAbortBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pipeline abort bad request response has a 3xx status code
func (o *PipelineAbortBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pipeline abort bad request response has a 4xx status code
func (o *PipelineAbortBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this pipeline abort bad request response has a 5xx status code
func (o *PipelineAbortBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this pipeline abort bad request response a status code equal to that given
func (o *PipelineAbortBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PipelineAbortBadRequest) Error() string {
	return fmt.Sprintf("[POST /apps/{app-slug}/pipelines/{pipeline-id}/abort][%d] pipelineAbortBadRequest  %+v", 400, o.Payload)
}

func (o *PipelineAbortBadRequest) String() string {
	return fmt.Sprintf("[POST /apps/{app-slug}/pipelines/{pipeline-id}/abort][%d] pipelineAbortBadRequest  %+v", 400, o.Payload)
}

func (o *PipelineAbortBadRequest) GetPayload() *models.ServiceStandardErrorRespModel {
	return o.Payload
}

func (o *PipelineAbortBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPipelineAbortUnauthorized creates a PipelineAbortUnauthorized with default headers values
func NewPipelineAbortUnauthorized() *PipelineAbortUnauthorized {
	return &PipelineAbortUnauthorized{}
}

/*
PipelineAbortUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PipelineAbortUnauthorized struct {
	Payload *models.ServiceStandardErrorRespModel
}

// IsSuccess returns true when this pipeline abort unauthorized response has a 2xx status code
func (o *PipelineAbortUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pipeline abort unauthorized response has a 3xx status code
func (o *PipelineAbortUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pipeline abort unauthorized response has a 4xx status code
func (o *PipelineAbortUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this pipeline abort unauthorized response has a 5xx status code
func (o *PipelineAbortUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this pipeline abort unauthorized response a status code equal to that given
func (o *PipelineAbortUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PipelineAbortUnauthorized) Error() string {
	return fmt.Sprintf("[POST /apps/{app-slug}/pipelines/{pipeline-id}/abort][%d] pipelineAbortUnauthorized  %+v", 401, o.Payload)
}

func (o *PipelineAbortUnauthorized) String() string {
	return fmt.Sprintf("[POST /apps/{app-slug}/pipelines/{pipeline-id}/abort][%d] pipelineAbortUnauthorized  %+v", 401, o.Payload)
}

func (o *PipelineAbortUnauthorized) GetPayload() *models.ServiceStandardErrorRespModel {
	return o.Payload
}

func (o *PipelineAbortUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPipelineAbortNotFound creates a PipelineAbortNotFound with default headers values
func NewPipelineAbortNotFound() *PipelineAbortNotFound {
	return &PipelineAbortNotFound{}
}

/*
PipelineAbortNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PipelineAbortNotFound struct {
	Payload *models.ServiceStandardErrorRespModel
}

// IsSuccess returns true when this pipeline abort not found response has a 2xx status code
func (o *PipelineAbortNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pipeline abort not found response has a 3xx status code
func (o *PipelineAbortNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pipeline abort not found response has a 4xx status code
func (o *PipelineAbortNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this pipeline abort not found response has a 5xx status code
func (o *PipelineAbortNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this pipeline abort not found response a status code equal to that given
func (o *PipelineAbortNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PipelineAbortNotFound) Error() string {
	return fmt.Sprintf("[POST /apps/{app-slug}/pipelines/{pipeline-id}/abort][%d] pipelineAbortNotFound  %+v", 404, o.Payload)
}

func (o *PipelineAbortNotFound) String() string {
	return fmt.Sprintf("[POST /apps/{app-slug}/pipelines/{pipeline-id}/abort][%d] pipelineAbortNotFound  %+v", 404, o.Payload)
}

func (o *PipelineAbortNotFound) GetPayload() *models.ServiceStandardErrorRespModel {
	return o.Payload
}

func (o *PipelineAbortNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPipelineAbortInternalServerError creates a PipelineAbortInternalServerError with default headers values
func NewPipelineAbortInternalServerError() *PipelineAbortInternalServerError {
	return &PipelineAbortInternalServerError{}
}

/*
PipelineAbortInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PipelineAbortInternalServerError struct {
	Payload *models.ServiceStandardErrorRespModel
}

// IsSuccess returns true when this pipeline abort internal server error response has a 2xx status code
func (o *PipelineAbortInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pipeline abort internal server error response has a 3xx status code
func (o *PipelineAbortInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pipeline abort internal server error response has a 4xx status code
func (o *PipelineAbortInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this pipeline abort internal server error response has a 5xx status code
func (o *PipelineAbortInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this pipeline abort internal server error response a status code equal to that given
func (o *PipelineAbortInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PipelineAbortInternalServerError) Error() string {
	return fmt.Sprintf("[POST /apps/{app-slug}/pipelines/{pipeline-id}/abort][%d] pipelineAbortInternalServerError  %+v", 500, o.Payload)
}

func (o *PipelineAbortInternalServerError) String() string {
	return fmt.Sprintf("[POST /apps/{app-slug}/pipelines/{pipeline-id}/abort][%d] pipelineAbortInternalServerError  %+v", 500, o.Payload)
}

func (o *PipelineAbortInternalServerError) GetPayload() *models.ServiceStandardErrorRespModel {
	return o.Payload
}

func (o *PipelineAbortInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
