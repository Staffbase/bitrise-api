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

// ArtifactShowReader is a Reader for the ArtifactShow structure.
type ArtifactShowReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ArtifactShowReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewArtifactShowOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewArtifactShowBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewArtifactShowUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewArtifactShowNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewArtifactShowInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewArtifactShowOK creates a ArtifactShowOK with default headers values
func NewArtifactShowOK() *ArtifactShowOK {
	return &ArtifactShowOK{}
}

/*
ArtifactShowOK describes a response with status code 200, with default header values.

OK
*/
type ArtifactShowOK struct {
	Payload *models.V0ArtifactShowResponseModel
}

// IsSuccess returns true when this artifact show o k response has a 2xx status code
func (o *ArtifactShowOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this artifact show o k response has a 3xx status code
func (o *ArtifactShowOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this artifact show o k response has a 4xx status code
func (o *ArtifactShowOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this artifact show o k response has a 5xx status code
func (o *ArtifactShowOK) IsServerError() bool {
	return false
}

// IsCode returns true when this artifact show o k response a status code equal to that given
func (o *ArtifactShowOK) IsCode(code int) bool {
	return code == 200
}

func (o *ArtifactShowOK) Error() string {
	return fmt.Sprintf("[GET /apps/{app-slug}/builds/{build-slug}/artifacts/{artifact-slug}][%d] artifactShowOK  %+v", 200, o.Payload)
}

func (o *ArtifactShowOK) String() string {
	return fmt.Sprintf("[GET /apps/{app-slug}/builds/{build-slug}/artifacts/{artifact-slug}][%d] artifactShowOK  %+v", 200, o.Payload)
}

func (o *ArtifactShowOK) GetPayload() *models.V0ArtifactShowResponseModel {
	return o.Payload
}

func (o *ArtifactShowOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V0ArtifactShowResponseModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewArtifactShowBadRequest creates a ArtifactShowBadRequest with default headers values
func NewArtifactShowBadRequest() *ArtifactShowBadRequest {
	return &ArtifactShowBadRequest{}
}

/*
ArtifactShowBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ArtifactShowBadRequest struct {
	Payload *models.ServiceStandardErrorRespModel
}

// IsSuccess returns true when this artifact show bad request response has a 2xx status code
func (o *ArtifactShowBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this artifact show bad request response has a 3xx status code
func (o *ArtifactShowBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this artifact show bad request response has a 4xx status code
func (o *ArtifactShowBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this artifact show bad request response has a 5xx status code
func (o *ArtifactShowBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this artifact show bad request response a status code equal to that given
func (o *ArtifactShowBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *ArtifactShowBadRequest) Error() string {
	return fmt.Sprintf("[GET /apps/{app-slug}/builds/{build-slug}/artifacts/{artifact-slug}][%d] artifactShowBadRequest  %+v", 400, o.Payload)
}

func (o *ArtifactShowBadRequest) String() string {
	return fmt.Sprintf("[GET /apps/{app-slug}/builds/{build-slug}/artifacts/{artifact-slug}][%d] artifactShowBadRequest  %+v", 400, o.Payload)
}

func (o *ArtifactShowBadRequest) GetPayload() *models.ServiceStandardErrorRespModel {
	return o.Payload
}

func (o *ArtifactShowBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewArtifactShowUnauthorized creates a ArtifactShowUnauthorized with default headers values
func NewArtifactShowUnauthorized() *ArtifactShowUnauthorized {
	return &ArtifactShowUnauthorized{}
}

/*
ArtifactShowUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ArtifactShowUnauthorized struct {
	Payload *models.ServiceStandardErrorRespModel
}

// IsSuccess returns true when this artifact show unauthorized response has a 2xx status code
func (o *ArtifactShowUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this artifact show unauthorized response has a 3xx status code
func (o *ArtifactShowUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this artifact show unauthorized response has a 4xx status code
func (o *ArtifactShowUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this artifact show unauthorized response has a 5xx status code
func (o *ArtifactShowUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this artifact show unauthorized response a status code equal to that given
func (o *ArtifactShowUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *ArtifactShowUnauthorized) Error() string {
	return fmt.Sprintf("[GET /apps/{app-slug}/builds/{build-slug}/artifacts/{artifact-slug}][%d] artifactShowUnauthorized  %+v", 401, o.Payload)
}

func (o *ArtifactShowUnauthorized) String() string {
	return fmt.Sprintf("[GET /apps/{app-slug}/builds/{build-slug}/artifacts/{artifact-slug}][%d] artifactShowUnauthorized  %+v", 401, o.Payload)
}

func (o *ArtifactShowUnauthorized) GetPayload() *models.ServiceStandardErrorRespModel {
	return o.Payload
}

func (o *ArtifactShowUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewArtifactShowNotFound creates a ArtifactShowNotFound with default headers values
func NewArtifactShowNotFound() *ArtifactShowNotFound {
	return &ArtifactShowNotFound{}
}

/*
ArtifactShowNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ArtifactShowNotFound struct {
	Payload *models.ServiceStandardErrorRespModel
}

// IsSuccess returns true when this artifact show not found response has a 2xx status code
func (o *ArtifactShowNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this artifact show not found response has a 3xx status code
func (o *ArtifactShowNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this artifact show not found response has a 4xx status code
func (o *ArtifactShowNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this artifact show not found response has a 5xx status code
func (o *ArtifactShowNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this artifact show not found response a status code equal to that given
func (o *ArtifactShowNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *ArtifactShowNotFound) Error() string {
	return fmt.Sprintf("[GET /apps/{app-slug}/builds/{build-slug}/artifacts/{artifact-slug}][%d] artifactShowNotFound  %+v", 404, o.Payload)
}

func (o *ArtifactShowNotFound) String() string {
	return fmt.Sprintf("[GET /apps/{app-slug}/builds/{build-slug}/artifacts/{artifact-slug}][%d] artifactShowNotFound  %+v", 404, o.Payload)
}

func (o *ArtifactShowNotFound) GetPayload() *models.ServiceStandardErrorRespModel {
	return o.Payload
}

func (o *ArtifactShowNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewArtifactShowInternalServerError creates a ArtifactShowInternalServerError with default headers values
func NewArtifactShowInternalServerError() *ArtifactShowInternalServerError {
	return &ArtifactShowInternalServerError{}
}

/*
ArtifactShowInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type ArtifactShowInternalServerError struct {
	Payload *models.ServiceStandardErrorRespModel
}

// IsSuccess returns true when this artifact show internal server error response has a 2xx status code
func (o *ArtifactShowInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this artifact show internal server error response has a 3xx status code
func (o *ArtifactShowInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this artifact show internal server error response has a 4xx status code
func (o *ArtifactShowInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this artifact show internal server error response has a 5xx status code
func (o *ArtifactShowInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this artifact show internal server error response a status code equal to that given
func (o *ArtifactShowInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ArtifactShowInternalServerError) Error() string {
	return fmt.Sprintf("[GET /apps/{app-slug}/builds/{build-slug}/artifacts/{artifact-slug}][%d] artifactShowInternalServerError  %+v", 500, o.Payload)
}

func (o *ArtifactShowInternalServerError) String() string {
	return fmt.Sprintf("[GET /apps/{app-slug}/builds/{build-slug}/artifacts/{artifact-slug}][%d] artifactShowInternalServerError  %+v", 500, o.Payload)
}

func (o *ArtifactShowInternalServerError) GetPayload() *models.ServiceStandardErrorRespModel {
	return o.Payload
}

func (o *ArtifactShowInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
