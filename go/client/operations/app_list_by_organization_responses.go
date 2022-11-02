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

// AppListByOrganizationReader is a Reader for the AppListByOrganization structure.
type AppListByOrganizationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AppListByOrganizationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAppListByOrganizationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAppListByOrganizationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAppListByOrganizationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAppListByOrganizationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAppListByOrganizationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAppListByOrganizationOK creates a AppListByOrganizationOK with default headers values
func NewAppListByOrganizationOK() *AppListByOrganizationOK {
	return &AppListByOrganizationOK{}
}

/*
AppListByOrganizationOK describes a response with status code 200, with default header values.

OK
*/
type AppListByOrganizationOK struct {
	Payload *models.V0AppListResponseModel
}

// IsSuccess returns true when this app list by organization o k response has a 2xx status code
func (o *AppListByOrganizationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this app list by organization o k response has a 3xx status code
func (o *AppListByOrganizationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this app list by organization o k response has a 4xx status code
func (o *AppListByOrganizationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this app list by organization o k response has a 5xx status code
func (o *AppListByOrganizationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this app list by organization o k response a status code equal to that given
func (o *AppListByOrganizationOK) IsCode(code int) bool {
	return code == 200
}

func (o *AppListByOrganizationOK) Error() string {
	return fmt.Sprintf("[GET /organizations/{org-slug}/apps][%d] appListByOrganizationOK  %+v", 200, o.Payload)
}

func (o *AppListByOrganizationOK) String() string {
	return fmt.Sprintf("[GET /organizations/{org-slug}/apps][%d] appListByOrganizationOK  %+v", 200, o.Payload)
}

func (o *AppListByOrganizationOK) GetPayload() *models.V0AppListResponseModel {
	return o.Payload
}

func (o *AppListByOrganizationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V0AppListResponseModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAppListByOrganizationBadRequest creates a AppListByOrganizationBadRequest with default headers values
func NewAppListByOrganizationBadRequest() *AppListByOrganizationBadRequest {
	return &AppListByOrganizationBadRequest{}
}

/*
AppListByOrganizationBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type AppListByOrganizationBadRequest struct {
	Payload *models.ServiceStandardErrorRespModel
}

// IsSuccess returns true when this app list by organization bad request response has a 2xx status code
func (o *AppListByOrganizationBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this app list by organization bad request response has a 3xx status code
func (o *AppListByOrganizationBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this app list by organization bad request response has a 4xx status code
func (o *AppListByOrganizationBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this app list by organization bad request response has a 5xx status code
func (o *AppListByOrganizationBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this app list by organization bad request response a status code equal to that given
func (o *AppListByOrganizationBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *AppListByOrganizationBadRequest) Error() string {
	return fmt.Sprintf("[GET /organizations/{org-slug}/apps][%d] appListByOrganizationBadRequest  %+v", 400, o.Payload)
}

func (o *AppListByOrganizationBadRequest) String() string {
	return fmt.Sprintf("[GET /organizations/{org-slug}/apps][%d] appListByOrganizationBadRequest  %+v", 400, o.Payload)
}

func (o *AppListByOrganizationBadRequest) GetPayload() *models.ServiceStandardErrorRespModel {
	return o.Payload
}

func (o *AppListByOrganizationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAppListByOrganizationUnauthorized creates a AppListByOrganizationUnauthorized with default headers values
func NewAppListByOrganizationUnauthorized() *AppListByOrganizationUnauthorized {
	return &AppListByOrganizationUnauthorized{}
}

/*
AppListByOrganizationUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type AppListByOrganizationUnauthorized struct {
	Payload *models.ServiceStandardErrorRespModel
}

// IsSuccess returns true when this app list by organization unauthorized response has a 2xx status code
func (o *AppListByOrganizationUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this app list by organization unauthorized response has a 3xx status code
func (o *AppListByOrganizationUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this app list by organization unauthorized response has a 4xx status code
func (o *AppListByOrganizationUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this app list by organization unauthorized response has a 5xx status code
func (o *AppListByOrganizationUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this app list by organization unauthorized response a status code equal to that given
func (o *AppListByOrganizationUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *AppListByOrganizationUnauthorized) Error() string {
	return fmt.Sprintf("[GET /organizations/{org-slug}/apps][%d] appListByOrganizationUnauthorized  %+v", 401, o.Payload)
}

func (o *AppListByOrganizationUnauthorized) String() string {
	return fmt.Sprintf("[GET /organizations/{org-slug}/apps][%d] appListByOrganizationUnauthorized  %+v", 401, o.Payload)
}

func (o *AppListByOrganizationUnauthorized) GetPayload() *models.ServiceStandardErrorRespModel {
	return o.Payload
}

func (o *AppListByOrganizationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAppListByOrganizationNotFound creates a AppListByOrganizationNotFound with default headers values
func NewAppListByOrganizationNotFound() *AppListByOrganizationNotFound {
	return &AppListByOrganizationNotFound{}
}

/*
AppListByOrganizationNotFound describes a response with status code 404, with default header values.

Not Found
*/
type AppListByOrganizationNotFound struct {
	Payload *models.ServiceStandardErrorRespModel
}

// IsSuccess returns true when this app list by organization not found response has a 2xx status code
func (o *AppListByOrganizationNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this app list by organization not found response has a 3xx status code
func (o *AppListByOrganizationNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this app list by organization not found response has a 4xx status code
func (o *AppListByOrganizationNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this app list by organization not found response has a 5xx status code
func (o *AppListByOrganizationNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this app list by organization not found response a status code equal to that given
func (o *AppListByOrganizationNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *AppListByOrganizationNotFound) Error() string {
	return fmt.Sprintf("[GET /organizations/{org-slug}/apps][%d] appListByOrganizationNotFound  %+v", 404, o.Payload)
}

func (o *AppListByOrganizationNotFound) String() string {
	return fmt.Sprintf("[GET /organizations/{org-slug}/apps][%d] appListByOrganizationNotFound  %+v", 404, o.Payload)
}

func (o *AppListByOrganizationNotFound) GetPayload() *models.ServiceStandardErrorRespModel {
	return o.Payload
}

func (o *AppListByOrganizationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAppListByOrganizationInternalServerError creates a AppListByOrganizationInternalServerError with default headers values
func NewAppListByOrganizationInternalServerError() *AppListByOrganizationInternalServerError {
	return &AppListByOrganizationInternalServerError{}
}

/*
AppListByOrganizationInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type AppListByOrganizationInternalServerError struct {
	Payload *models.ServiceStandardErrorRespModel
}

// IsSuccess returns true when this app list by organization internal server error response has a 2xx status code
func (o *AppListByOrganizationInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this app list by organization internal server error response has a 3xx status code
func (o *AppListByOrganizationInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this app list by organization internal server error response has a 4xx status code
func (o *AppListByOrganizationInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this app list by organization internal server error response has a 5xx status code
func (o *AppListByOrganizationInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this app list by organization internal server error response a status code equal to that given
func (o *AppListByOrganizationInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *AppListByOrganizationInternalServerError) Error() string {
	return fmt.Sprintf("[GET /organizations/{org-slug}/apps][%d] appListByOrganizationInternalServerError  %+v", 500, o.Payload)
}

func (o *AppListByOrganizationInternalServerError) String() string {
	return fmt.Sprintf("[GET /organizations/{org-slug}/apps][%d] appListByOrganizationInternalServerError  %+v", 500, o.Payload)
}

func (o *AppListByOrganizationInternalServerError) GetPayload() *models.ServiceStandardErrorRespModel {
	return o.Payload
}

func (o *AppListByOrganizationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
