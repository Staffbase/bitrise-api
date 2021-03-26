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

// BuildBitriseYmlShowReader is a Reader for the BuildBitriseYmlShow structure.
type BuildBitriseYmlShowReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BuildBitriseYmlShowReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBuildBitriseYmlShowOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewBuildBitriseYmlShowBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewBuildBitriseYmlShowUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewBuildBitriseYmlShowNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewBuildBitriseYmlShowInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewBuildBitriseYmlShowOK creates a BuildBitriseYmlShowOK with default headers values
func NewBuildBitriseYmlShowOK() *BuildBitriseYmlShowOK {
	return &BuildBitriseYmlShowOK{}
}

/* BuildBitriseYmlShowOK describes a response with status code 200, with default header values.

{the bitrise.yml in yml format}
*/
type BuildBitriseYmlShowOK struct {
	Payload string
}

func (o *BuildBitriseYmlShowOK) Error() string {
	return fmt.Sprintf("[GET /apps/{app-slug}/builds/{build-slug}/bitrise.yml][%d] buildBitriseYmlShowOK  %+v", 200, o.Payload)
}
func (o *BuildBitriseYmlShowOK) GetPayload() string {
	return o.Payload
}

func (o *BuildBitriseYmlShowOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBuildBitriseYmlShowBadRequest creates a BuildBitriseYmlShowBadRequest with default headers values
func NewBuildBitriseYmlShowBadRequest() *BuildBitriseYmlShowBadRequest {
	return &BuildBitriseYmlShowBadRequest{}
}

/* BuildBitriseYmlShowBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type BuildBitriseYmlShowBadRequest struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *BuildBitriseYmlShowBadRequest) Error() string {
	return fmt.Sprintf("[GET /apps/{app-slug}/builds/{build-slug}/bitrise.yml][%d] buildBitriseYmlShowBadRequest  %+v", 400, o.Payload)
}
func (o *BuildBitriseYmlShowBadRequest) GetPayload() *models.ServiceStandardErrorRespModel {
	return o.Payload
}

func (o *BuildBitriseYmlShowBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBuildBitriseYmlShowUnauthorized creates a BuildBitriseYmlShowUnauthorized with default headers values
func NewBuildBitriseYmlShowUnauthorized() *BuildBitriseYmlShowUnauthorized {
	return &BuildBitriseYmlShowUnauthorized{}
}

/* BuildBitriseYmlShowUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type BuildBitriseYmlShowUnauthorized struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *BuildBitriseYmlShowUnauthorized) Error() string {
	return fmt.Sprintf("[GET /apps/{app-slug}/builds/{build-slug}/bitrise.yml][%d] buildBitriseYmlShowUnauthorized  %+v", 401, o.Payload)
}
func (o *BuildBitriseYmlShowUnauthorized) GetPayload() *models.ServiceStandardErrorRespModel {
	return o.Payload
}

func (o *BuildBitriseYmlShowUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBuildBitriseYmlShowNotFound creates a BuildBitriseYmlShowNotFound with default headers values
func NewBuildBitriseYmlShowNotFound() *BuildBitriseYmlShowNotFound {
	return &BuildBitriseYmlShowNotFound{}
}

/* BuildBitriseYmlShowNotFound describes a response with status code 404, with default header values.

Not Found
*/
type BuildBitriseYmlShowNotFound struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *BuildBitriseYmlShowNotFound) Error() string {
	return fmt.Sprintf("[GET /apps/{app-slug}/builds/{build-slug}/bitrise.yml][%d] buildBitriseYmlShowNotFound  %+v", 404, o.Payload)
}
func (o *BuildBitriseYmlShowNotFound) GetPayload() *models.ServiceStandardErrorRespModel {
	return o.Payload
}

func (o *BuildBitriseYmlShowNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBuildBitriseYmlShowInternalServerError creates a BuildBitriseYmlShowInternalServerError with default headers values
func NewBuildBitriseYmlShowInternalServerError() *BuildBitriseYmlShowInternalServerError {
	return &BuildBitriseYmlShowInternalServerError{}
}

/* BuildBitriseYmlShowInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type BuildBitriseYmlShowInternalServerError struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *BuildBitriseYmlShowInternalServerError) Error() string {
	return fmt.Sprintf("[GET /apps/{app-slug}/builds/{build-slug}/bitrise.yml][%d] buildBitriseYmlShowInternalServerError  %+v", 500, o.Payload)
}
func (o *BuildBitriseYmlShowInternalServerError) GetPayload() *models.ServiceStandardErrorRespModel {
	return o.Payload
}

func (o *BuildBitriseYmlShowInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
