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

// ArtifactListReader is a Reader for the ArtifactList structure.
type ArtifactListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ArtifactListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewArtifactListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewArtifactListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewArtifactListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewArtifactListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewArtifactListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /apps/{app-slug}/builds/{build-slug}/artifacts] artifact-list", response, response.Code())
	}
}

// NewArtifactListOK creates a ArtifactListOK with default headers values
func NewArtifactListOK() *ArtifactListOK {
	return &ArtifactListOK{}
}

/*
ArtifactListOK describes a response with status code 200, with default header values.

OK
*/
type ArtifactListOK struct {
	Payload *models.V0ArtifactListResponseModel
}

// IsSuccess returns true when this artifact list o k response has a 2xx status code
func (o *ArtifactListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this artifact list o k response has a 3xx status code
func (o *ArtifactListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this artifact list o k response has a 4xx status code
func (o *ArtifactListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this artifact list o k response has a 5xx status code
func (o *ArtifactListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this artifact list o k response a status code equal to that given
func (o *ArtifactListOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the artifact list o k response
func (o *ArtifactListOK) Code() int {
	return 200
}

func (o *ArtifactListOK) Error() string {
	return fmt.Sprintf("[GET /apps/{app-slug}/builds/{build-slug}/artifacts][%d] artifactListOK  %+v", 200, o.Payload)
}

func (o *ArtifactListOK) String() string {
	return fmt.Sprintf("[GET /apps/{app-slug}/builds/{build-slug}/artifacts][%d] artifactListOK  %+v", 200, o.Payload)
}

func (o *ArtifactListOK) GetPayload() *models.V0ArtifactListResponseModel {
	return o.Payload
}

func (o *ArtifactListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V0ArtifactListResponseModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewArtifactListBadRequest creates a ArtifactListBadRequest with default headers values
func NewArtifactListBadRequest() *ArtifactListBadRequest {
	return &ArtifactListBadRequest{}
}

/*
ArtifactListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ArtifactListBadRequest struct {
	Payload *models.ServiceStandardErrorRespModel
}

// IsSuccess returns true when this artifact list bad request response has a 2xx status code
func (o *ArtifactListBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this artifact list bad request response has a 3xx status code
func (o *ArtifactListBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this artifact list bad request response has a 4xx status code
func (o *ArtifactListBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this artifact list bad request response has a 5xx status code
func (o *ArtifactListBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this artifact list bad request response a status code equal to that given
func (o *ArtifactListBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the artifact list bad request response
func (o *ArtifactListBadRequest) Code() int {
	return 400
}

func (o *ArtifactListBadRequest) Error() string {
	return fmt.Sprintf("[GET /apps/{app-slug}/builds/{build-slug}/artifacts][%d] artifactListBadRequest  %+v", 400, o.Payload)
}

func (o *ArtifactListBadRequest) String() string {
	return fmt.Sprintf("[GET /apps/{app-slug}/builds/{build-slug}/artifacts][%d] artifactListBadRequest  %+v", 400, o.Payload)
}

func (o *ArtifactListBadRequest) GetPayload() *models.ServiceStandardErrorRespModel {
	return o.Payload
}

func (o *ArtifactListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewArtifactListUnauthorized creates a ArtifactListUnauthorized with default headers values
func NewArtifactListUnauthorized() *ArtifactListUnauthorized {
	return &ArtifactListUnauthorized{}
}

/*
ArtifactListUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ArtifactListUnauthorized struct {
	Payload *models.ServiceStandardErrorRespModel
}

// IsSuccess returns true when this artifact list unauthorized response has a 2xx status code
func (o *ArtifactListUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this artifact list unauthorized response has a 3xx status code
func (o *ArtifactListUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this artifact list unauthorized response has a 4xx status code
func (o *ArtifactListUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this artifact list unauthorized response has a 5xx status code
func (o *ArtifactListUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this artifact list unauthorized response a status code equal to that given
func (o *ArtifactListUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the artifact list unauthorized response
func (o *ArtifactListUnauthorized) Code() int {
	return 401
}

func (o *ArtifactListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /apps/{app-slug}/builds/{build-slug}/artifacts][%d] artifactListUnauthorized  %+v", 401, o.Payload)
}

func (o *ArtifactListUnauthorized) String() string {
	return fmt.Sprintf("[GET /apps/{app-slug}/builds/{build-slug}/artifacts][%d] artifactListUnauthorized  %+v", 401, o.Payload)
}

func (o *ArtifactListUnauthorized) GetPayload() *models.ServiceStandardErrorRespModel {
	return o.Payload
}

func (o *ArtifactListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewArtifactListNotFound creates a ArtifactListNotFound with default headers values
func NewArtifactListNotFound() *ArtifactListNotFound {
	return &ArtifactListNotFound{}
}

/*
ArtifactListNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ArtifactListNotFound struct {
	Payload *models.ServiceStandardErrorRespModel
}

// IsSuccess returns true when this artifact list not found response has a 2xx status code
func (o *ArtifactListNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this artifact list not found response has a 3xx status code
func (o *ArtifactListNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this artifact list not found response has a 4xx status code
func (o *ArtifactListNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this artifact list not found response has a 5xx status code
func (o *ArtifactListNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this artifact list not found response a status code equal to that given
func (o *ArtifactListNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the artifact list not found response
func (o *ArtifactListNotFound) Code() int {
	return 404
}

func (o *ArtifactListNotFound) Error() string {
	return fmt.Sprintf("[GET /apps/{app-slug}/builds/{build-slug}/artifacts][%d] artifactListNotFound  %+v", 404, o.Payload)
}

func (o *ArtifactListNotFound) String() string {
	return fmt.Sprintf("[GET /apps/{app-slug}/builds/{build-slug}/artifacts][%d] artifactListNotFound  %+v", 404, o.Payload)
}

func (o *ArtifactListNotFound) GetPayload() *models.ServiceStandardErrorRespModel {
	return o.Payload
}

func (o *ArtifactListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewArtifactListInternalServerError creates a ArtifactListInternalServerError with default headers values
func NewArtifactListInternalServerError() *ArtifactListInternalServerError {
	return &ArtifactListInternalServerError{}
}

/*
ArtifactListInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type ArtifactListInternalServerError struct {
	Payload *models.ServiceStandardErrorRespModel
}

// IsSuccess returns true when this artifact list internal server error response has a 2xx status code
func (o *ArtifactListInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this artifact list internal server error response has a 3xx status code
func (o *ArtifactListInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this artifact list internal server error response has a 4xx status code
func (o *ArtifactListInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this artifact list internal server error response has a 5xx status code
func (o *ArtifactListInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this artifact list internal server error response a status code equal to that given
func (o *ArtifactListInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the artifact list internal server error response
func (o *ArtifactListInternalServerError) Code() int {
	return 500
}

func (o *ArtifactListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /apps/{app-slug}/builds/{build-slug}/artifacts][%d] artifactListInternalServerError  %+v", 500, o.Payload)
}

func (o *ArtifactListInternalServerError) String() string {
	return fmt.Sprintf("[GET /apps/{app-slug}/builds/{build-slug}/artifacts][%d] artifactListInternalServerError  %+v", 500, o.Payload)
}

func (o *ArtifactListInternalServerError) GetPayload() *models.ServiceStandardErrorRespModel {
	return o.Payload
}

func (o *ArtifactListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
