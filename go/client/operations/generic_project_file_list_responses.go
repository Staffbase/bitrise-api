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

// GenericProjectFileListReader is a Reader for the GenericProjectFileList structure.
type GenericProjectFileListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GenericProjectFileListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGenericProjectFileListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGenericProjectFileListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGenericProjectFileListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGenericProjectFileListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGenericProjectFileListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGenericProjectFileListOK creates a GenericProjectFileListOK with default headers values
func NewGenericProjectFileListOK() *GenericProjectFileListOK {
	return &GenericProjectFileListOK{}
}

/* GenericProjectFileListOK describes a response with status code 200, with default header values.

OK
*/
type GenericProjectFileListOK struct {
	Payload *models.V0ProjectFileStorageListResponseModel
}

func (o *GenericProjectFileListOK) Error() string {
	return fmt.Sprintf("[GET /apps/{app-slug}/generic-project-files][%d] genericProjectFileListOK  %+v", 200, o.Payload)
}
func (o *GenericProjectFileListOK) GetPayload() *models.V0ProjectFileStorageListResponseModel {
	return o.Payload
}

func (o *GenericProjectFileListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V0ProjectFileStorageListResponseModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGenericProjectFileListBadRequest creates a GenericProjectFileListBadRequest with default headers values
func NewGenericProjectFileListBadRequest() *GenericProjectFileListBadRequest {
	return &GenericProjectFileListBadRequest{}
}

/* GenericProjectFileListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GenericProjectFileListBadRequest struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *GenericProjectFileListBadRequest) Error() string {
	return fmt.Sprintf("[GET /apps/{app-slug}/generic-project-files][%d] genericProjectFileListBadRequest  %+v", 400, o.Payload)
}
func (o *GenericProjectFileListBadRequest) GetPayload() *models.ServiceStandardErrorRespModel {
	return o.Payload
}

func (o *GenericProjectFileListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGenericProjectFileListUnauthorized creates a GenericProjectFileListUnauthorized with default headers values
func NewGenericProjectFileListUnauthorized() *GenericProjectFileListUnauthorized {
	return &GenericProjectFileListUnauthorized{}
}

/* GenericProjectFileListUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GenericProjectFileListUnauthorized struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *GenericProjectFileListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /apps/{app-slug}/generic-project-files][%d] genericProjectFileListUnauthorized  %+v", 401, o.Payload)
}
func (o *GenericProjectFileListUnauthorized) GetPayload() *models.ServiceStandardErrorRespModel {
	return o.Payload
}

func (o *GenericProjectFileListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGenericProjectFileListNotFound creates a GenericProjectFileListNotFound with default headers values
func NewGenericProjectFileListNotFound() *GenericProjectFileListNotFound {
	return &GenericProjectFileListNotFound{}
}

/* GenericProjectFileListNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GenericProjectFileListNotFound struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *GenericProjectFileListNotFound) Error() string {
	return fmt.Sprintf("[GET /apps/{app-slug}/generic-project-files][%d] genericProjectFileListNotFound  %+v", 404, o.Payload)
}
func (o *GenericProjectFileListNotFound) GetPayload() *models.ServiceStandardErrorRespModel {
	return o.Payload
}

func (o *GenericProjectFileListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGenericProjectFileListInternalServerError creates a GenericProjectFileListInternalServerError with default headers values
func NewGenericProjectFileListInternalServerError() *GenericProjectFileListInternalServerError {
	return &GenericProjectFileListInternalServerError{}
}

/* GenericProjectFileListInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GenericProjectFileListInternalServerError struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *GenericProjectFileListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /apps/{app-slug}/generic-project-files][%d] genericProjectFileListInternalServerError  %+v", 500, o.Payload)
}
func (o *GenericProjectFileListInternalServerError) GetPayload() *models.ServiceStandardErrorRespModel {
	return o.Payload
}

func (o *GenericProjectFileListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
