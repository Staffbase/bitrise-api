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

// BuildTriggerReader is a Reader for the BuildTrigger structure.
type BuildTriggerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BuildTriggerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewBuildTriggerCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewBuildTriggerBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewBuildTriggerUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewBuildTriggerNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewBuildTriggerInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewBuildTriggerCreated creates a BuildTriggerCreated with default headers values
func NewBuildTriggerCreated() *BuildTriggerCreated {
	return &BuildTriggerCreated{}
}

/* BuildTriggerCreated describes a response with status code 201, with default header values.

OK
*/
type BuildTriggerCreated struct {
	Payload *models.V0BuildTriggerRespModel
}

func (o *BuildTriggerCreated) Error() string {
	return fmt.Sprintf("[POST /apps/{app-slug}/builds][%d] buildTriggerCreated  %+v", 201, o.Payload)
}
func (o *BuildTriggerCreated) GetPayload() *models.V0BuildTriggerRespModel {
	return o.Payload
}

func (o *BuildTriggerCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V0BuildTriggerRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBuildTriggerBadRequest creates a BuildTriggerBadRequest with default headers values
func NewBuildTriggerBadRequest() *BuildTriggerBadRequest {
	return &BuildTriggerBadRequest{}
}

/* BuildTriggerBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type BuildTriggerBadRequest struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *BuildTriggerBadRequest) Error() string {
	return fmt.Sprintf("[POST /apps/{app-slug}/builds][%d] buildTriggerBadRequest  %+v", 400, o.Payload)
}
func (o *BuildTriggerBadRequest) GetPayload() *models.ServiceStandardErrorRespModel {
	return o.Payload
}

func (o *BuildTriggerBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBuildTriggerUnauthorized creates a BuildTriggerUnauthorized with default headers values
func NewBuildTriggerUnauthorized() *BuildTriggerUnauthorized {
	return &BuildTriggerUnauthorized{}
}

/* BuildTriggerUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type BuildTriggerUnauthorized struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *BuildTriggerUnauthorized) Error() string {
	return fmt.Sprintf("[POST /apps/{app-slug}/builds][%d] buildTriggerUnauthorized  %+v", 401, o.Payload)
}
func (o *BuildTriggerUnauthorized) GetPayload() *models.ServiceStandardErrorRespModel {
	return o.Payload
}

func (o *BuildTriggerUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBuildTriggerNotFound creates a BuildTriggerNotFound with default headers values
func NewBuildTriggerNotFound() *BuildTriggerNotFound {
	return &BuildTriggerNotFound{}
}

/* BuildTriggerNotFound describes a response with status code 404, with default header values.

Not Found
*/
type BuildTriggerNotFound struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *BuildTriggerNotFound) Error() string {
	return fmt.Sprintf("[POST /apps/{app-slug}/builds][%d] buildTriggerNotFound  %+v", 404, o.Payload)
}
func (o *BuildTriggerNotFound) GetPayload() *models.ServiceStandardErrorRespModel {
	return o.Payload
}

func (o *BuildTriggerNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBuildTriggerInternalServerError creates a BuildTriggerInternalServerError with default headers values
func NewBuildTriggerInternalServerError() *BuildTriggerInternalServerError {
	return &BuildTriggerInternalServerError{}
}

/* BuildTriggerInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type BuildTriggerInternalServerError struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *BuildTriggerInternalServerError) Error() string {
	return fmt.Sprintf("[POST /apps/{app-slug}/builds][%d] buildTriggerInternalServerError  %+v", 500, o.Payload)
}
func (o *BuildTriggerInternalServerError) GetPayload() *models.ServiceStandardErrorRespModel {
	return o.Payload
}

func (o *BuildTriggerInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
