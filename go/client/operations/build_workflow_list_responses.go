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
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewBuildWorkflowListOK creates a BuildWorkflowListOK with default headers values
func NewBuildWorkflowListOK() *BuildWorkflowListOK {
	return &BuildWorkflowListOK{}
}

/* BuildWorkflowListOK describes a response with status code 200, with default header values.

OK
*/
type BuildWorkflowListOK struct {
	Payload *models.V0BuildWorkflowListResponseModel
}

func (o *BuildWorkflowListOK) Error() string {
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

/* BuildWorkflowListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type BuildWorkflowListBadRequest struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *BuildWorkflowListBadRequest) Error() string {
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

/* BuildWorkflowListUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type BuildWorkflowListUnauthorized struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *BuildWorkflowListUnauthorized) Error() string {
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

/* BuildWorkflowListNotFound describes a response with status code 404, with default header values.

Not Found
*/
type BuildWorkflowListNotFound struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *BuildWorkflowListNotFound) Error() string {
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

/* BuildWorkflowListInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type BuildWorkflowListInternalServerError struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *BuildWorkflowListInternalServerError) Error() string {
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
