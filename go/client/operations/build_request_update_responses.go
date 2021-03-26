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

// BuildRequestUpdateReader is a Reader for the BuildRequestUpdate structure.
type BuildRequestUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BuildRequestUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBuildRequestUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewBuildRequestUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewBuildRequestUpdateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewBuildRequestUpdateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewBuildRequestUpdateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewBuildRequestUpdateOK creates a BuildRequestUpdateOK with default headers values
func NewBuildRequestUpdateOK() *BuildRequestUpdateOK {
	return &BuildRequestUpdateOK{}
}

/* BuildRequestUpdateOK describes a response with status code 200, with default header values.

OK
*/
type BuildRequestUpdateOK struct {
	Payload *models.V0BuildRequestUpdateResponseModel
}

func (o *BuildRequestUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /apps/{app-slug}/build-requests/{build-request-slug}][%d] buildRequestUpdateOK  %+v", 200, o.Payload)
}
func (o *BuildRequestUpdateOK) GetPayload() *models.V0BuildRequestUpdateResponseModel {
	return o.Payload
}

func (o *BuildRequestUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V0BuildRequestUpdateResponseModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBuildRequestUpdateBadRequest creates a BuildRequestUpdateBadRequest with default headers values
func NewBuildRequestUpdateBadRequest() *BuildRequestUpdateBadRequest {
	return &BuildRequestUpdateBadRequest{}
}

/* BuildRequestUpdateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type BuildRequestUpdateBadRequest struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *BuildRequestUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /apps/{app-slug}/build-requests/{build-request-slug}][%d] buildRequestUpdateBadRequest  %+v", 400, o.Payload)
}
func (o *BuildRequestUpdateBadRequest) GetPayload() *models.ServiceStandardErrorRespModel {
	return o.Payload
}

func (o *BuildRequestUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBuildRequestUpdateUnauthorized creates a BuildRequestUpdateUnauthorized with default headers values
func NewBuildRequestUpdateUnauthorized() *BuildRequestUpdateUnauthorized {
	return &BuildRequestUpdateUnauthorized{}
}

/* BuildRequestUpdateUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type BuildRequestUpdateUnauthorized struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *BuildRequestUpdateUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /apps/{app-slug}/build-requests/{build-request-slug}][%d] buildRequestUpdateUnauthorized  %+v", 401, o.Payload)
}
func (o *BuildRequestUpdateUnauthorized) GetPayload() *models.ServiceStandardErrorRespModel {
	return o.Payload
}

func (o *BuildRequestUpdateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBuildRequestUpdateNotFound creates a BuildRequestUpdateNotFound with default headers values
func NewBuildRequestUpdateNotFound() *BuildRequestUpdateNotFound {
	return &BuildRequestUpdateNotFound{}
}

/* BuildRequestUpdateNotFound describes a response with status code 404, with default header values.

Not Found
*/
type BuildRequestUpdateNotFound struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *BuildRequestUpdateNotFound) Error() string {
	return fmt.Sprintf("[PATCH /apps/{app-slug}/build-requests/{build-request-slug}][%d] buildRequestUpdateNotFound  %+v", 404, o.Payload)
}
func (o *BuildRequestUpdateNotFound) GetPayload() *models.ServiceStandardErrorRespModel {
	return o.Payload
}

func (o *BuildRequestUpdateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBuildRequestUpdateInternalServerError creates a BuildRequestUpdateInternalServerError with default headers values
func NewBuildRequestUpdateInternalServerError() *BuildRequestUpdateInternalServerError {
	return &BuildRequestUpdateInternalServerError{}
}

/* BuildRequestUpdateInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type BuildRequestUpdateInternalServerError struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *BuildRequestUpdateInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /apps/{app-slug}/build-requests/{build-request-slug}][%d] buildRequestUpdateInternalServerError  %+v", 500, o.Payload)
}
func (o *BuildRequestUpdateInternalServerError) GetPayload() *models.ServiceStandardErrorRespModel {
	return o.Payload
}

func (o *BuildRequestUpdateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
