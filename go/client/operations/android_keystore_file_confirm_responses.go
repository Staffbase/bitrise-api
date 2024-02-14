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

// AndroidKeystoreFileConfirmReader is a Reader for the AndroidKeystoreFileConfirm structure.
type AndroidKeystoreFileConfirmReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AndroidKeystoreFileConfirmReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAndroidKeystoreFileConfirmOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAndroidKeystoreFileConfirmBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAndroidKeystoreFileConfirmUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAndroidKeystoreFileConfirmNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAndroidKeystoreFileConfirmInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /apps/{app-slug}/android-keystore-files/{android-keystore-file-slug}/uploaded] android-keystore-file-confirm", response, response.Code())
	}
}

// NewAndroidKeystoreFileConfirmOK creates a AndroidKeystoreFileConfirmOK with default headers values
func NewAndroidKeystoreFileConfirmOK() *AndroidKeystoreFileConfirmOK {
	return &AndroidKeystoreFileConfirmOK{}
}

/*
AndroidKeystoreFileConfirmOK describes a response with status code 200, with default header values.

OK
*/
type AndroidKeystoreFileConfirmOK struct {
	Payload *models.V0ProjectFileStorageResponseModel
}

// IsSuccess returns true when this android keystore file confirm o k response has a 2xx status code
func (o *AndroidKeystoreFileConfirmOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this android keystore file confirm o k response has a 3xx status code
func (o *AndroidKeystoreFileConfirmOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this android keystore file confirm o k response has a 4xx status code
func (o *AndroidKeystoreFileConfirmOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this android keystore file confirm o k response has a 5xx status code
func (o *AndroidKeystoreFileConfirmOK) IsServerError() bool {
	return false
}

// IsCode returns true when this android keystore file confirm o k response a status code equal to that given
func (o *AndroidKeystoreFileConfirmOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the android keystore file confirm o k response
func (o *AndroidKeystoreFileConfirmOK) Code() int {
	return 200
}

func (o *AndroidKeystoreFileConfirmOK) Error() string {
	return fmt.Sprintf("[POST /apps/{app-slug}/android-keystore-files/{android-keystore-file-slug}/uploaded][%d] androidKeystoreFileConfirmOK  %+v", 200, o.Payload)
}

func (o *AndroidKeystoreFileConfirmOK) String() string {
	return fmt.Sprintf("[POST /apps/{app-slug}/android-keystore-files/{android-keystore-file-slug}/uploaded][%d] androidKeystoreFileConfirmOK  %+v", 200, o.Payload)
}

func (o *AndroidKeystoreFileConfirmOK) GetPayload() *models.V0ProjectFileStorageResponseModel {
	return o.Payload
}

func (o *AndroidKeystoreFileConfirmOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V0ProjectFileStorageResponseModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAndroidKeystoreFileConfirmBadRequest creates a AndroidKeystoreFileConfirmBadRequest with default headers values
func NewAndroidKeystoreFileConfirmBadRequest() *AndroidKeystoreFileConfirmBadRequest {
	return &AndroidKeystoreFileConfirmBadRequest{}
}

/*
AndroidKeystoreFileConfirmBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type AndroidKeystoreFileConfirmBadRequest struct {
	Payload *models.ServiceStandardErrorRespModel
}

// IsSuccess returns true when this android keystore file confirm bad request response has a 2xx status code
func (o *AndroidKeystoreFileConfirmBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this android keystore file confirm bad request response has a 3xx status code
func (o *AndroidKeystoreFileConfirmBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this android keystore file confirm bad request response has a 4xx status code
func (o *AndroidKeystoreFileConfirmBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this android keystore file confirm bad request response has a 5xx status code
func (o *AndroidKeystoreFileConfirmBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this android keystore file confirm bad request response a status code equal to that given
func (o *AndroidKeystoreFileConfirmBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the android keystore file confirm bad request response
func (o *AndroidKeystoreFileConfirmBadRequest) Code() int {
	return 400
}

func (o *AndroidKeystoreFileConfirmBadRequest) Error() string {
	return fmt.Sprintf("[POST /apps/{app-slug}/android-keystore-files/{android-keystore-file-slug}/uploaded][%d] androidKeystoreFileConfirmBadRequest  %+v", 400, o.Payload)
}

func (o *AndroidKeystoreFileConfirmBadRequest) String() string {
	return fmt.Sprintf("[POST /apps/{app-slug}/android-keystore-files/{android-keystore-file-slug}/uploaded][%d] androidKeystoreFileConfirmBadRequest  %+v", 400, o.Payload)
}

func (o *AndroidKeystoreFileConfirmBadRequest) GetPayload() *models.ServiceStandardErrorRespModel {
	return o.Payload
}

func (o *AndroidKeystoreFileConfirmBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAndroidKeystoreFileConfirmUnauthorized creates a AndroidKeystoreFileConfirmUnauthorized with default headers values
func NewAndroidKeystoreFileConfirmUnauthorized() *AndroidKeystoreFileConfirmUnauthorized {
	return &AndroidKeystoreFileConfirmUnauthorized{}
}

/*
AndroidKeystoreFileConfirmUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type AndroidKeystoreFileConfirmUnauthorized struct {
	Payload *models.ServiceStandardErrorRespModel
}

// IsSuccess returns true when this android keystore file confirm unauthorized response has a 2xx status code
func (o *AndroidKeystoreFileConfirmUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this android keystore file confirm unauthorized response has a 3xx status code
func (o *AndroidKeystoreFileConfirmUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this android keystore file confirm unauthorized response has a 4xx status code
func (o *AndroidKeystoreFileConfirmUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this android keystore file confirm unauthorized response has a 5xx status code
func (o *AndroidKeystoreFileConfirmUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this android keystore file confirm unauthorized response a status code equal to that given
func (o *AndroidKeystoreFileConfirmUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the android keystore file confirm unauthorized response
func (o *AndroidKeystoreFileConfirmUnauthorized) Code() int {
	return 401
}

func (o *AndroidKeystoreFileConfirmUnauthorized) Error() string {
	return fmt.Sprintf("[POST /apps/{app-slug}/android-keystore-files/{android-keystore-file-slug}/uploaded][%d] androidKeystoreFileConfirmUnauthorized  %+v", 401, o.Payload)
}

func (o *AndroidKeystoreFileConfirmUnauthorized) String() string {
	return fmt.Sprintf("[POST /apps/{app-slug}/android-keystore-files/{android-keystore-file-slug}/uploaded][%d] androidKeystoreFileConfirmUnauthorized  %+v", 401, o.Payload)
}

func (o *AndroidKeystoreFileConfirmUnauthorized) GetPayload() *models.ServiceStandardErrorRespModel {
	return o.Payload
}

func (o *AndroidKeystoreFileConfirmUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAndroidKeystoreFileConfirmNotFound creates a AndroidKeystoreFileConfirmNotFound with default headers values
func NewAndroidKeystoreFileConfirmNotFound() *AndroidKeystoreFileConfirmNotFound {
	return &AndroidKeystoreFileConfirmNotFound{}
}

/*
AndroidKeystoreFileConfirmNotFound describes a response with status code 404, with default header values.

Not Found
*/
type AndroidKeystoreFileConfirmNotFound struct {
	Payload *models.ServiceStandardErrorRespModel
}

// IsSuccess returns true when this android keystore file confirm not found response has a 2xx status code
func (o *AndroidKeystoreFileConfirmNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this android keystore file confirm not found response has a 3xx status code
func (o *AndroidKeystoreFileConfirmNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this android keystore file confirm not found response has a 4xx status code
func (o *AndroidKeystoreFileConfirmNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this android keystore file confirm not found response has a 5xx status code
func (o *AndroidKeystoreFileConfirmNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this android keystore file confirm not found response a status code equal to that given
func (o *AndroidKeystoreFileConfirmNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the android keystore file confirm not found response
func (o *AndroidKeystoreFileConfirmNotFound) Code() int {
	return 404
}

func (o *AndroidKeystoreFileConfirmNotFound) Error() string {
	return fmt.Sprintf("[POST /apps/{app-slug}/android-keystore-files/{android-keystore-file-slug}/uploaded][%d] androidKeystoreFileConfirmNotFound  %+v", 404, o.Payload)
}

func (o *AndroidKeystoreFileConfirmNotFound) String() string {
	return fmt.Sprintf("[POST /apps/{app-slug}/android-keystore-files/{android-keystore-file-slug}/uploaded][%d] androidKeystoreFileConfirmNotFound  %+v", 404, o.Payload)
}

func (o *AndroidKeystoreFileConfirmNotFound) GetPayload() *models.ServiceStandardErrorRespModel {
	return o.Payload
}

func (o *AndroidKeystoreFileConfirmNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAndroidKeystoreFileConfirmInternalServerError creates a AndroidKeystoreFileConfirmInternalServerError with default headers values
func NewAndroidKeystoreFileConfirmInternalServerError() *AndroidKeystoreFileConfirmInternalServerError {
	return &AndroidKeystoreFileConfirmInternalServerError{}
}

/*
AndroidKeystoreFileConfirmInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type AndroidKeystoreFileConfirmInternalServerError struct {
	Payload *models.ServiceStandardErrorRespModel
}

// IsSuccess returns true when this android keystore file confirm internal server error response has a 2xx status code
func (o *AndroidKeystoreFileConfirmInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this android keystore file confirm internal server error response has a 3xx status code
func (o *AndroidKeystoreFileConfirmInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this android keystore file confirm internal server error response has a 4xx status code
func (o *AndroidKeystoreFileConfirmInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this android keystore file confirm internal server error response has a 5xx status code
func (o *AndroidKeystoreFileConfirmInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this android keystore file confirm internal server error response a status code equal to that given
func (o *AndroidKeystoreFileConfirmInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the android keystore file confirm internal server error response
func (o *AndroidKeystoreFileConfirmInternalServerError) Code() int {
	return 500
}

func (o *AndroidKeystoreFileConfirmInternalServerError) Error() string {
	return fmt.Sprintf("[POST /apps/{app-slug}/android-keystore-files/{android-keystore-file-slug}/uploaded][%d] androidKeystoreFileConfirmInternalServerError  %+v", 500, o.Payload)
}

func (o *AndroidKeystoreFileConfirmInternalServerError) String() string {
	return fmt.Sprintf("[POST /apps/{app-slug}/android-keystore-files/{android-keystore-file-slug}/uploaded][%d] androidKeystoreFileConfirmInternalServerError  %+v", 500, o.Payload)
}

func (o *AndroidKeystoreFileConfirmInternalServerError) GetPayload() *models.ServiceStandardErrorRespModel {
	return o.Payload
}

func (o *AndroidKeystoreFileConfirmInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
