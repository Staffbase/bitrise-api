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

// GenericProjectFileDeleteReader is a Reader for the GenericProjectFileDelete structure.
type GenericProjectFileDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GenericProjectFileDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGenericProjectFileDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGenericProjectFileDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGenericProjectFileDeleteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGenericProjectFileDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGenericProjectFileDeleteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGenericProjectFileDeleteOK creates a GenericProjectFileDeleteOK with default headers values
func NewGenericProjectFileDeleteOK() *GenericProjectFileDeleteOK {
	return &GenericProjectFileDeleteOK{}
}

/* GenericProjectFileDeleteOK describes a response with status code 200, with default header values.

OK
*/
type GenericProjectFileDeleteOK struct {
	Payload *models.V0ProjectFileStorageResponseModel
}

func (o *GenericProjectFileDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /apps/{app-slug}/generic-project-files/{generic-project-file-slug}][%d] genericProjectFileDeleteOK  %+v", 200, o.Payload)
}
func (o *GenericProjectFileDeleteOK) GetPayload() *models.V0ProjectFileStorageResponseModel {
	return o.Payload
}

func (o *GenericProjectFileDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V0ProjectFileStorageResponseModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGenericProjectFileDeleteBadRequest creates a GenericProjectFileDeleteBadRequest with default headers values
func NewGenericProjectFileDeleteBadRequest() *GenericProjectFileDeleteBadRequest {
	return &GenericProjectFileDeleteBadRequest{}
}

/* GenericProjectFileDeleteBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GenericProjectFileDeleteBadRequest struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *GenericProjectFileDeleteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /apps/{app-slug}/generic-project-files/{generic-project-file-slug}][%d] genericProjectFileDeleteBadRequest  %+v", 400, o.Payload)
}
func (o *GenericProjectFileDeleteBadRequest) GetPayload() *models.ServiceStandardErrorRespModel {
	return o.Payload
}

func (o *GenericProjectFileDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGenericProjectFileDeleteUnauthorized creates a GenericProjectFileDeleteUnauthorized with default headers values
func NewGenericProjectFileDeleteUnauthorized() *GenericProjectFileDeleteUnauthorized {
	return &GenericProjectFileDeleteUnauthorized{}
}

/* GenericProjectFileDeleteUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GenericProjectFileDeleteUnauthorized struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *GenericProjectFileDeleteUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /apps/{app-slug}/generic-project-files/{generic-project-file-slug}][%d] genericProjectFileDeleteUnauthorized  %+v", 401, o.Payload)
}
func (o *GenericProjectFileDeleteUnauthorized) GetPayload() *models.ServiceStandardErrorRespModel {
	return o.Payload
}

func (o *GenericProjectFileDeleteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGenericProjectFileDeleteNotFound creates a GenericProjectFileDeleteNotFound with default headers values
func NewGenericProjectFileDeleteNotFound() *GenericProjectFileDeleteNotFound {
	return &GenericProjectFileDeleteNotFound{}
}

/* GenericProjectFileDeleteNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GenericProjectFileDeleteNotFound struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *GenericProjectFileDeleteNotFound) Error() string {
	return fmt.Sprintf("[DELETE /apps/{app-slug}/generic-project-files/{generic-project-file-slug}][%d] genericProjectFileDeleteNotFound  %+v", 404, o.Payload)
}
func (o *GenericProjectFileDeleteNotFound) GetPayload() *models.ServiceStandardErrorRespModel {
	return o.Payload
}

func (o *GenericProjectFileDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGenericProjectFileDeleteInternalServerError creates a GenericProjectFileDeleteInternalServerError with default headers values
func NewGenericProjectFileDeleteInternalServerError() *GenericProjectFileDeleteInternalServerError {
	return &GenericProjectFileDeleteInternalServerError{}
}

/* GenericProjectFileDeleteInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GenericProjectFileDeleteInternalServerError struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *GenericProjectFileDeleteInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /apps/{app-slug}/generic-project-files/{generic-project-file-slug}][%d] genericProjectFileDeleteInternalServerError  %+v", 500, o.Payload)
}
func (o *GenericProjectFileDeleteInternalServerError) GetPayload() *models.ServiceStandardErrorRespModel {
	return o.Payload
}

func (o *GenericProjectFileDeleteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
