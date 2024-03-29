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

// AndroidKeystoreFileListReader is a Reader for the AndroidKeystoreFileList structure.
type AndroidKeystoreFileListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AndroidKeystoreFileListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAndroidKeystoreFileListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAndroidKeystoreFileListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAndroidKeystoreFileListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAndroidKeystoreFileListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAndroidKeystoreFileListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /apps/{app-slug}/android-keystore-files] android-keystore-file-list", response, response.Code())
	}
}

// NewAndroidKeystoreFileListOK creates a AndroidKeystoreFileListOK with default headers values
func NewAndroidKeystoreFileListOK() *AndroidKeystoreFileListOK {
	return &AndroidKeystoreFileListOK{}
}

/*
AndroidKeystoreFileListOK describes a response with status code 200, with default header values.

OK
*/
type AndroidKeystoreFileListOK struct {
	Payload *models.V0ProjectFileStorageListResponseModel
}

// IsSuccess returns true when this android keystore file list o k response has a 2xx status code
func (o *AndroidKeystoreFileListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this android keystore file list o k response has a 3xx status code
func (o *AndroidKeystoreFileListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this android keystore file list o k response has a 4xx status code
func (o *AndroidKeystoreFileListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this android keystore file list o k response has a 5xx status code
func (o *AndroidKeystoreFileListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this android keystore file list o k response a status code equal to that given
func (o *AndroidKeystoreFileListOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the android keystore file list o k response
func (o *AndroidKeystoreFileListOK) Code() int {
	return 200
}

func (o *AndroidKeystoreFileListOK) Error() string {
	return fmt.Sprintf("[GET /apps/{app-slug}/android-keystore-files][%d] androidKeystoreFileListOK  %+v", 200, o.Payload)
}

func (o *AndroidKeystoreFileListOK) String() string {
	return fmt.Sprintf("[GET /apps/{app-slug}/android-keystore-files][%d] androidKeystoreFileListOK  %+v", 200, o.Payload)
}

func (o *AndroidKeystoreFileListOK) GetPayload() *models.V0ProjectFileStorageListResponseModel {
	return o.Payload
}

func (o *AndroidKeystoreFileListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V0ProjectFileStorageListResponseModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAndroidKeystoreFileListBadRequest creates a AndroidKeystoreFileListBadRequest with default headers values
func NewAndroidKeystoreFileListBadRequest() *AndroidKeystoreFileListBadRequest {
	return &AndroidKeystoreFileListBadRequest{}
}

/*
AndroidKeystoreFileListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type AndroidKeystoreFileListBadRequest struct {
	Payload *models.ServiceStandardErrorRespModel
}

// IsSuccess returns true when this android keystore file list bad request response has a 2xx status code
func (o *AndroidKeystoreFileListBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this android keystore file list bad request response has a 3xx status code
func (o *AndroidKeystoreFileListBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this android keystore file list bad request response has a 4xx status code
func (o *AndroidKeystoreFileListBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this android keystore file list bad request response has a 5xx status code
func (o *AndroidKeystoreFileListBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this android keystore file list bad request response a status code equal to that given
func (o *AndroidKeystoreFileListBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the android keystore file list bad request response
func (o *AndroidKeystoreFileListBadRequest) Code() int {
	return 400
}

func (o *AndroidKeystoreFileListBadRequest) Error() string {
	return fmt.Sprintf("[GET /apps/{app-slug}/android-keystore-files][%d] androidKeystoreFileListBadRequest  %+v", 400, o.Payload)
}

func (o *AndroidKeystoreFileListBadRequest) String() string {
	return fmt.Sprintf("[GET /apps/{app-slug}/android-keystore-files][%d] androidKeystoreFileListBadRequest  %+v", 400, o.Payload)
}

func (o *AndroidKeystoreFileListBadRequest) GetPayload() *models.ServiceStandardErrorRespModel {
	return o.Payload
}

func (o *AndroidKeystoreFileListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAndroidKeystoreFileListUnauthorized creates a AndroidKeystoreFileListUnauthorized with default headers values
func NewAndroidKeystoreFileListUnauthorized() *AndroidKeystoreFileListUnauthorized {
	return &AndroidKeystoreFileListUnauthorized{}
}

/*
AndroidKeystoreFileListUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type AndroidKeystoreFileListUnauthorized struct {
	Payload *models.ServiceStandardErrorRespModel
}

// IsSuccess returns true when this android keystore file list unauthorized response has a 2xx status code
func (o *AndroidKeystoreFileListUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this android keystore file list unauthorized response has a 3xx status code
func (o *AndroidKeystoreFileListUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this android keystore file list unauthorized response has a 4xx status code
func (o *AndroidKeystoreFileListUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this android keystore file list unauthorized response has a 5xx status code
func (o *AndroidKeystoreFileListUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this android keystore file list unauthorized response a status code equal to that given
func (o *AndroidKeystoreFileListUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the android keystore file list unauthorized response
func (o *AndroidKeystoreFileListUnauthorized) Code() int {
	return 401
}

func (o *AndroidKeystoreFileListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /apps/{app-slug}/android-keystore-files][%d] androidKeystoreFileListUnauthorized  %+v", 401, o.Payload)
}

func (o *AndroidKeystoreFileListUnauthorized) String() string {
	return fmt.Sprintf("[GET /apps/{app-slug}/android-keystore-files][%d] androidKeystoreFileListUnauthorized  %+v", 401, o.Payload)
}

func (o *AndroidKeystoreFileListUnauthorized) GetPayload() *models.ServiceStandardErrorRespModel {
	return o.Payload
}

func (o *AndroidKeystoreFileListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAndroidKeystoreFileListNotFound creates a AndroidKeystoreFileListNotFound with default headers values
func NewAndroidKeystoreFileListNotFound() *AndroidKeystoreFileListNotFound {
	return &AndroidKeystoreFileListNotFound{}
}

/*
AndroidKeystoreFileListNotFound describes a response with status code 404, with default header values.

Not Found
*/
type AndroidKeystoreFileListNotFound struct {
	Payload *models.ServiceStandardErrorRespModel
}

// IsSuccess returns true when this android keystore file list not found response has a 2xx status code
func (o *AndroidKeystoreFileListNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this android keystore file list not found response has a 3xx status code
func (o *AndroidKeystoreFileListNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this android keystore file list not found response has a 4xx status code
func (o *AndroidKeystoreFileListNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this android keystore file list not found response has a 5xx status code
func (o *AndroidKeystoreFileListNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this android keystore file list not found response a status code equal to that given
func (o *AndroidKeystoreFileListNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the android keystore file list not found response
func (o *AndroidKeystoreFileListNotFound) Code() int {
	return 404
}

func (o *AndroidKeystoreFileListNotFound) Error() string {
	return fmt.Sprintf("[GET /apps/{app-slug}/android-keystore-files][%d] androidKeystoreFileListNotFound  %+v", 404, o.Payload)
}

func (o *AndroidKeystoreFileListNotFound) String() string {
	return fmt.Sprintf("[GET /apps/{app-slug}/android-keystore-files][%d] androidKeystoreFileListNotFound  %+v", 404, o.Payload)
}

func (o *AndroidKeystoreFileListNotFound) GetPayload() *models.ServiceStandardErrorRespModel {
	return o.Payload
}

func (o *AndroidKeystoreFileListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAndroidKeystoreFileListInternalServerError creates a AndroidKeystoreFileListInternalServerError with default headers values
func NewAndroidKeystoreFileListInternalServerError() *AndroidKeystoreFileListInternalServerError {
	return &AndroidKeystoreFileListInternalServerError{}
}

/*
AndroidKeystoreFileListInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type AndroidKeystoreFileListInternalServerError struct {
	Payload *models.ServiceStandardErrorRespModel
}

// IsSuccess returns true when this android keystore file list internal server error response has a 2xx status code
func (o *AndroidKeystoreFileListInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this android keystore file list internal server error response has a 3xx status code
func (o *AndroidKeystoreFileListInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this android keystore file list internal server error response has a 4xx status code
func (o *AndroidKeystoreFileListInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this android keystore file list internal server error response has a 5xx status code
func (o *AndroidKeystoreFileListInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this android keystore file list internal server error response a status code equal to that given
func (o *AndroidKeystoreFileListInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the android keystore file list internal server error response
func (o *AndroidKeystoreFileListInternalServerError) Code() int {
	return 500
}

func (o *AndroidKeystoreFileListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /apps/{app-slug}/android-keystore-files][%d] androidKeystoreFileListInternalServerError  %+v", 500, o.Payload)
}

func (o *AndroidKeystoreFileListInternalServerError) String() string {
	return fmt.Sprintf("[GET /apps/{app-slug}/android-keystore-files][%d] androidKeystoreFileListInternalServerError  %+v", 500, o.Payload)
}

func (o *AndroidKeystoreFileListInternalServerError) GetPayload() *models.ServiceStandardErrorRespModel {
	return o.Payload
}

func (o *AndroidKeystoreFileListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
