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

// AppSetupBitriseYmlConfigUpdateReader is a Reader for the AppSetupBitriseYmlConfigUpdate structure.
type AppSetupBitriseYmlConfigUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AppSetupBitriseYmlConfigUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewAppSetupBitriseYmlConfigUpdateNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAppSetupBitriseYmlConfigUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAppSetupBitriseYmlConfigUpdateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAppSetupBitriseYmlConfigUpdateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAppSetupBitriseYmlConfigUpdateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /apps/{app-slug}/bitrise.yml/config] app-setup-bitrise-yml-config-update", response, response.Code())
	}
}

// NewAppSetupBitriseYmlConfigUpdateNoContent creates a AppSetupBitriseYmlConfigUpdateNoContent with default headers values
func NewAppSetupBitriseYmlConfigUpdateNoContent() *AppSetupBitriseYmlConfigUpdateNoContent {
	return &AppSetupBitriseYmlConfigUpdateNoContent{}
}

/*
AppSetupBitriseYmlConfigUpdateNoContent describes a response with status code 204, with default header values.

No Content
*/
type AppSetupBitriseYmlConfigUpdateNoContent struct {
}

// IsSuccess returns true when this app setup bitrise yml config update no content response has a 2xx status code
func (o *AppSetupBitriseYmlConfigUpdateNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this app setup bitrise yml config update no content response has a 3xx status code
func (o *AppSetupBitriseYmlConfigUpdateNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this app setup bitrise yml config update no content response has a 4xx status code
func (o *AppSetupBitriseYmlConfigUpdateNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this app setup bitrise yml config update no content response has a 5xx status code
func (o *AppSetupBitriseYmlConfigUpdateNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this app setup bitrise yml config update no content response a status code equal to that given
func (o *AppSetupBitriseYmlConfigUpdateNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the app setup bitrise yml config update no content response
func (o *AppSetupBitriseYmlConfigUpdateNoContent) Code() int {
	return 204
}

func (o *AppSetupBitriseYmlConfigUpdateNoContent) Error() string {
	return fmt.Sprintf("[PUT /apps/{app-slug}/bitrise.yml/config][%d] appSetupBitriseYmlConfigUpdateNoContent ", 204)
}

func (o *AppSetupBitriseYmlConfigUpdateNoContent) String() string {
	return fmt.Sprintf("[PUT /apps/{app-slug}/bitrise.yml/config][%d] appSetupBitriseYmlConfigUpdateNoContent ", 204)
}

func (o *AppSetupBitriseYmlConfigUpdateNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAppSetupBitriseYmlConfigUpdateBadRequest creates a AppSetupBitriseYmlConfigUpdateBadRequest with default headers values
func NewAppSetupBitriseYmlConfigUpdateBadRequest() *AppSetupBitriseYmlConfigUpdateBadRequest {
	return &AppSetupBitriseYmlConfigUpdateBadRequest{}
}

/*
AppSetupBitriseYmlConfigUpdateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type AppSetupBitriseYmlConfigUpdateBadRequest struct {
	Payload *models.ServiceStandardErrorRespModel
}

// IsSuccess returns true when this app setup bitrise yml config update bad request response has a 2xx status code
func (o *AppSetupBitriseYmlConfigUpdateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this app setup bitrise yml config update bad request response has a 3xx status code
func (o *AppSetupBitriseYmlConfigUpdateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this app setup bitrise yml config update bad request response has a 4xx status code
func (o *AppSetupBitriseYmlConfigUpdateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this app setup bitrise yml config update bad request response has a 5xx status code
func (o *AppSetupBitriseYmlConfigUpdateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this app setup bitrise yml config update bad request response a status code equal to that given
func (o *AppSetupBitriseYmlConfigUpdateBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the app setup bitrise yml config update bad request response
func (o *AppSetupBitriseYmlConfigUpdateBadRequest) Code() int {
	return 400
}

func (o *AppSetupBitriseYmlConfigUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PUT /apps/{app-slug}/bitrise.yml/config][%d] appSetupBitriseYmlConfigUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *AppSetupBitriseYmlConfigUpdateBadRequest) String() string {
	return fmt.Sprintf("[PUT /apps/{app-slug}/bitrise.yml/config][%d] appSetupBitriseYmlConfigUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *AppSetupBitriseYmlConfigUpdateBadRequest) GetPayload() *models.ServiceStandardErrorRespModel {
	return o.Payload
}

func (o *AppSetupBitriseYmlConfigUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAppSetupBitriseYmlConfigUpdateUnauthorized creates a AppSetupBitriseYmlConfigUpdateUnauthorized with default headers values
func NewAppSetupBitriseYmlConfigUpdateUnauthorized() *AppSetupBitriseYmlConfigUpdateUnauthorized {
	return &AppSetupBitriseYmlConfigUpdateUnauthorized{}
}

/*
AppSetupBitriseYmlConfigUpdateUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type AppSetupBitriseYmlConfigUpdateUnauthorized struct {
	Payload *models.ServiceStandardErrorRespModel
}

// IsSuccess returns true when this app setup bitrise yml config update unauthorized response has a 2xx status code
func (o *AppSetupBitriseYmlConfigUpdateUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this app setup bitrise yml config update unauthorized response has a 3xx status code
func (o *AppSetupBitriseYmlConfigUpdateUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this app setup bitrise yml config update unauthorized response has a 4xx status code
func (o *AppSetupBitriseYmlConfigUpdateUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this app setup bitrise yml config update unauthorized response has a 5xx status code
func (o *AppSetupBitriseYmlConfigUpdateUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this app setup bitrise yml config update unauthorized response a status code equal to that given
func (o *AppSetupBitriseYmlConfigUpdateUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the app setup bitrise yml config update unauthorized response
func (o *AppSetupBitriseYmlConfigUpdateUnauthorized) Code() int {
	return 401
}

func (o *AppSetupBitriseYmlConfigUpdateUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /apps/{app-slug}/bitrise.yml/config][%d] appSetupBitriseYmlConfigUpdateUnauthorized  %+v", 401, o.Payload)
}

func (o *AppSetupBitriseYmlConfigUpdateUnauthorized) String() string {
	return fmt.Sprintf("[PUT /apps/{app-slug}/bitrise.yml/config][%d] appSetupBitriseYmlConfigUpdateUnauthorized  %+v", 401, o.Payload)
}

func (o *AppSetupBitriseYmlConfigUpdateUnauthorized) GetPayload() *models.ServiceStandardErrorRespModel {
	return o.Payload
}

func (o *AppSetupBitriseYmlConfigUpdateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAppSetupBitriseYmlConfigUpdateNotFound creates a AppSetupBitriseYmlConfigUpdateNotFound with default headers values
func NewAppSetupBitriseYmlConfigUpdateNotFound() *AppSetupBitriseYmlConfigUpdateNotFound {
	return &AppSetupBitriseYmlConfigUpdateNotFound{}
}

/*
AppSetupBitriseYmlConfigUpdateNotFound describes a response with status code 404, with default header values.

Not Found
*/
type AppSetupBitriseYmlConfigUpdateNotFound struct {
	Payload *models.ServiceStandardErrorRespModel
}

// IsSuccess returns true when this app setup bitrise yml config update not found response has a 2xx status code
func (o *AppSetupBitriseYmlConfigUpdateNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this app setup bitrise yml config update not found response has a 3xx status code
func (o *AppSetupBitriseYmlConfigUpdateNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this app setup bitrise yml config update not found response has a 4xx status code
func (o *AppSetupBitriseYmlConfigUpdateNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this app setup bitrise yml config update not found response has a 5xx status code
func (o *AppSetupBitriseYmlConfigUpdateNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this app setup bitrise yml config update not found response a status code equal to that given
func (o *AppSetupBitriseYmlConfigUpdateNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the app setup bitrise yml config update not found response
func (o *AppSetupBitriseYmlConfigUpdateNotFound) Code() int {
	return 404
}

func (o *AppSetupBitriseYmlConfigUpdateNotFound) Error() string {
	return fmt.Sprintf("[PUT /apps/{app-slug}/bitrise.yml/config][%d] appSetupBitriseYmlConfigUpdateNotFound  %+v", 404, o.Payload)
}

func (o *AppSetupBitriseYmlConfigUpdateNotFound) String() string {
	return fmt.Sprintf("[PUT /apps/{app-slug}/bitrise.yml/config][%d] appSetupBitriseYmlConfigUpdateNotFound  %+v", 404, o.Payload)
}

func (o *AppSetupBitriseYmlConfigUpdateNotFound) GetPayload() *models.ServiceStandardErrorRespModel {
	return o.Payload
}

func (o *AppSetupBitriseYmlConfigUpdateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAppSetupBitriseYmlConfigUpdateInternalServerError creates a AppSetupBitriseYmlConfigUpdateInternalServerError with default headers values
func NewAppSetupBitriseYmlConfigUpdateInternalServerError() *AppSetupBitriseYmlConfigUpdateInternalServerError {
	return &AppSetupBitriseYmlConfigUpdateInternalServerError{}
}

/*
AppSetupBitriseYmlConfigUpdateInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type AppSetupBitriseYmlConfigUpdateInternalServerError struct {
	Payload *models.ServiceStandardErrorRespModel
}

// IsSuccess returns true when this app setup bitrise yml config update internal server error response has a 2xx status code
func (o *AppSetupBitriseYmlConfigUpdateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this app setup bitrise yml config update internal server error response has a 3xx status code
func (o *AppSetupBitriseYmlConfigUpdateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this app setup bitrise yml config update internal server error response has a 4xx status code
func (o *AppSetupBitriseYmlConfigUpdateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this app setup bitrise yml config update internal server error response has a 5xx status code
func (o *AppSetupBitriseYmlConfigUpdateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this app setup bitrise yml config update internal server error response a status code equal to that given
func (o *AppSetupBitriseYmlConfigUpdateInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the app setup bitrise yml config update internal server error response
func (o *AppSetupBitriseYmlConfigUpdateInternalServerError) Code() int {
	return 500
}

func (o *AppSetupBitriseYmlConfigUpdateInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /apps/{app-slug}/bitrise.yml/config][%d] appSetupBitriseYmlConfigUpdateInternalServerError  %+v", 500, o.Payload)
}

func (o *AppSetupBitriseYmlConfigUpdateInternalServerError) String() string {
	return fmt.Sprintf("[PUT /apps/{app-slug}/bitrise.yml/config][%d] appSetupBitriseYmlConfigUpdateInternalServerError  %+v", 500, o.Payload)
}

func (o *AppSetupBitriseYmlConfigUpdateInternalServerError) GetPayload() *models.ServiceStandardErrorRespModel {
	return o.Payload
}

func (o *AppSetupBitriseYmlConfigUpdateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
