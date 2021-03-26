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

// AvatarCandidatePromoteReader is a Reader for the AvatarCandidatePromote structure.
type AvatarCandidatePromoteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AvatarCandidatePromoteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAvatarCandidatePromoteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAvatarCandidatePromoteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAvatarCandidatePromoteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAvatarCandidatePromoteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAvatarCandidatePromoteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAvatarCandidatePromoteOK creates a AvatarCandidatePromoteOK with default headers values
func NewAvatarCandidatePromoteOK() *AvatarCandidatePromoteOK {
	return &AvatarCandidatePromoteOK{}
}

/* AvatarCandidatePromoteOK describes a response with status code 200, with default header values.

OK
*/
type AvatarCandidatePromoteOK struct {
	Payload *models.V0AvatarPromoteResponseModel
}

func (o *AvatarCandidatePromoteOK) Error() string {
	return fmt.Sprintf("[PATCH /apps/{app-slug}/avatar-candidates/{avatar-slug}][%d] avatarCandidatePromoteOK  %+v", 200, o.Payload)
}
func (o *AvatarCandidatePromoteOK) GetPayload() *models.V0AvatarPromoteResponseModel {
	return o.Payload
}

func (o *AvatarCandidatePromoteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V0AvatarPromoteResponseModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAvatarCandidatePromoteBadRequest creates a AvatarCandidatePromoteBadRequest with default headers values
func NewAvatarCandidatePromoteBadRequest() *AvatarCandidatePromoteBadRequest {
	return &AvatarCandidatePromoteBadRequest{}
}

/* AvatarCandidatePromoteBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type AvatarCandidatePromoteBadRequest struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *AvatarCandidatePromoteBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /apps/{app-slug}/avatar-candidates/{avatar-slug}][%d] avatarCandidatePromoteBadRequest  %+v", 400, o.Payload)
}
func (o *AvatarCandidatePromoteBadRequest) GetPayload() *models.ServiceStandardErrorRespModel {
	return o.Payload
}

func (o *AvatarCandidatePromoteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAvatarCandidatePromoteUnauthorized creates a AvatarCandidatePromoteUnauthorized with default headers values
func NewAvatarCandidatePromoteUnauthorized() *AvatarCandidatePromoteUnauthorized {
	return &AvatarCandidatePromoteUnauthorized{}
}

/* AvatarCandidatePromoteUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type AvatarCandidatePromoteUnauthorized struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *AvatarCandidatePromoteUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /apps/{app-slug}/avatar-candidates/{avatar-slug}][%d] avatarCandidatePromoteUnauthorized  %+v", 401, o.Payload)
}
func (o *AvatarCandidatePromoteUnauthorized) GetPayload() *models.ServiceStandardErrorRespModel {
	return o.Payload
}

func (o *AvatarCandidatePromoteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAvatarCandidatePromoteNotFound creates a AvatarCandidatePromoteNotFound with default headers values
func NewAvatarCandidatePromoteNotFound() *AvatarCandidatePromoteNotFound {
	return &AvatarCandidatePromoteNotFound{}
}

/* AvatarCandidatePromoteNotFound describes a response with status code 404, with default header values.

Not Found
*/
type AvatarCandidatePromoteNotFound struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *AvatarCandidatePromoteNotFound) Error() string {
	return fmt.Sprintf("[PATCH /apps/{app-slug}/avatar-candidates/{avatar-slug}][%d] avatarCandidatePromoteNotFound  %+v", 404, o.Payload)
}
func (o *AvatarCandidatePromoteNotFound) GetPayload() *models.ServiceStandardErrorRespModel {
	return o.Payload
}

func (o *AvatarCandidatePromoteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAvatarCandidatePromoteInternalServerError creates a AvatarCandidatePromoteInternalServerError with default headers values
func NewAvatarCandidatePromoteInternalServerError() *AvatarCandidatePromoteInternalServerError {
	return &AvatarCandidatePromoteInternalServerError{}
}

/* AvatarCandidatePromoteInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type AvatarCandidatePromoteInternalServerError struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *AvatarCandidatePromoteInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /apps/{app-slug}/avatar-candidates/{avatar-slug}][%d] avatarCandidatePromoteInternalServerError  %+v", 500, o.Payload)
}
func (o *AvatarCandidatePromoteInternalServerError) GetPayload() *models.ServiceStandardErrorRespModel {
	return o.Payload
}

func (o *AvatarCandidatePromoteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
