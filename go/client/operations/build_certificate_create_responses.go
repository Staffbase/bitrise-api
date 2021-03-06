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

// BuildCertificateCreateReader is a Reader for the BuildCertificateCreate structure.
type BuildCertificateCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BuildCertificateCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewBuildCertificateCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewBuildCertificateCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewBuildCertificateCreateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewBuildCertificateCreateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewBuildCertificateCreateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewBuildCertificateCreateCreated creates a BuildCertificateCreateCreated with default headers values
func NewBuildCertificateCreateCreated() *BuildCertificateCreateCreated {
	return &BuildCertificateCreateCreated{}
}

/* BuildCertificateCreateCreated describes a response with status code 201, with default header values.

Created
*/
type BuildCertificateCreateCreated struct {
	Payload *models.V0BuildCertificateResponseModel
}

func (o *BuildCertificateCreateCreated) Error() string {
	return fmt.Sprintf("[POST /apps/{app-slug}/build-certificates][%d] buildCertificateCreateCreated  %+v", 201, o.Payload)
}
func (o *BuildCertificateCreateCreated) GetPayload() *models.V0BuildCertificateResponseModel {
	return o.Payload
}

func (o *BuildCertificateCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V0BuildCertificateResponseModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBuildCertificateCreateBadRequest creates a BuildCertificateCreateBadRequest with default headers values
func NewBuildCertificateCreateBadRequest() *BuildCertificateCreateBadRequest {
	return &BuildCertificateCreateBadRequest{}
}

/* BuildCertificateCreateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type BuildCertificateCreateBadRequest struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *BuildCertificateCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /apps/{app-slug}/build-certificates][%d] buildCertificateCreateBadRequest  %+v", 400, o.Payload)
}
func (o *BuildCertificateCreateBadRequest) GetPayload() *models.ServiceStandardErrorRespModel {
	return o.Payload
}

func (o *BuildCertificateCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBuildCertificateCreateUnauthorized creates a BuildCertificateCreateUnauthorized with default headers values
func NewBuildCertificateCreateUnauthorized() *BuildCertificateCreateUnauthorized {
	return &BuildCertificateCreateUnauthorized{}
}

/* BuildCertificateCreateUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type BuildCertificateCreateUnauthorized struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *BuildCertificateCreateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /apps/{app-slug}/build-certificates][%d] buildCertificateCreateUnauthorized  %+v", 401, o.Payload)
}
func (o *BuildCertificateCreateUnauthorized) GetPayload() *models.ServiceStandardErrorRespModel {
	return o.Payload
}

func (o *BuildCertificateCreateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBuildCertificateCreateNotFound creates a BuildCertificateCreateNotFound with default headers values
func NewBuildCertificateCreateNotFound() *BuildCertificateCreateNotFound {
	return &BuildCertificateCreateNotFound{}
}

/* BuildCertificateCreateNotFound describes a response with status code 404, with default header values.

Not Found
*/
type BuildCertificateCreateNotFound struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *BuildCertificateCreateNotFound) Error() string {
	return fmt.Sprintf("[POST /apps/{app-slug}/build-certificates][%d] buildCertificateCreateNotFound  %+v", 404, o.Payload)
}
func (o *BuildCertificateCreateNotFound) GetPayload() *models.ServiceStandardErrorRespModel {
	return o.Payload
}

func (o *BuildCertificateCreateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBuildCertificateCreateInternalServerError creates a BuildCertificateCreateInternalServerError with default headers values
func NewBuildCertificateCreateInternalServerError() *BuildCertificateCreateInternalServerError {
	return &BuildCertificateCreateInternalServerError{}
}

/* BuildCertificateCreateInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type BuildCertificateCreateInternalServerError struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *BuildCertificateCreateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /apps/{app-slug}/build-certificates][%d] buildCertificateCreateInternalServerError  %+v", 500, o.Payload)
}
func (o *BuildCertificateCreateInternalServerError) GetPayload() *models.ServiceStandardErrorRespModel {
	return o.Payload
}

func (o *BuildCertificateCreateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
