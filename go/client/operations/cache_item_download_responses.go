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

// CacheItemDownloadReader is a Reader for the CacheItemDownload structure.
type CacheItemDownloadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CacheItemDownloadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCacheItemDownloadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCacheItemDownloadBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCacheItemDownloadUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCacheItemDownloadNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCacheItemDownloadInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /apps/{app-slug}/cache-items/{cache-item-id}/download] cache-item-download", response, response.Code())
	}
}

// NewCacheItemDownloadOK creates a CacheItemDownloadOK with default headers values
func NewCacheItemDownloadOK() *CacheItemDownloadOK {
	return &CacheItemDownloadOK{}
}

/*
CacheItemDownloadOK describes a response with status code 200, with default header values.

OK
*/
type CacheItemDownloadOK struct {
	Payload *models.V0CacheItemDownloadResponseModel
}

// IsSuccess returns true when this cache item download o k response has a 2xx status code
func (o *CacheItemDownloadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cache item download o k response has a 3xx status code
func (o *CacheItemDownloadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cache item download o k response has a 4xx status code
func (o *CacheItemDownloadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this cache item download o k response has a 5xx status code
func (o *CacheItemDownloadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this cache item download o k response a status code equal to that given
func (o *CacheItemDownloadOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the cache item download o k response
func (o *CacheItemDownloadOK) Code() int {
	return 200
}

func (o *CacheItemDownloadOK) Error() string {
	return fmt.Sprintf("[GET /apps/{app-slug}/cache-items/{cache-item-id}/download][%d] cacheItemDownloadOK  %+v", 200, o.Payload)
}

func (o *CacheItemDownloadOK) String() string {
	return fmt.Sprintf("[GET /apps/{app-slug}/cache-items/{cache-item-id}/download][%d] cacheItemDownloadOK  %+v", 200, o.Payload)
}

func (o *CacheItemDownloadOK) GetPayload() *models.V0CacheItemDownloadResponseModel {
	return o.Payload
}

func (o *CacheItemDownloadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V0CacheItemDownloadResponseModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCacheItemDownloadBadRequest creates a CacheItemDownloadBadRequest with default headers values
func NewCacheItemDownloadBadRequest() *CacheItemDownloadBadRequest {
	return &CacheItemDownloadBadRequest{}
}

/*
CacheItemDownloadBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CacheItemDownloadBadRequest struct {
	Payload *models.ServiceStandardErrorRespModel
}

// IsSuccess returns true when this cache item download bad request response has a 2xx status code
func (o *CacheItemDownloadBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cache item download bad request response has a 3xx status code
func (o *CacheItemDownloadBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cache item download bad request response has a 4xx status code
func (o *CacheItemDownloadBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this cache item download bad request response has a 5xx status code
func (o *CacheItemDownloadBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this cache item download bad request response a status code equal to that given
func (o *CacheItemDownloadBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the cache item download bad request response
func (o *CacheItemDownloadBadRequest) Code() int {
	return 400
}

func (o *CacheItemDownloadBadRequest) Error() string {
	return fmt.Sprintf("[GET /apps/{app-slug}/cache-items/{cache-item-id}/download][%d] cacheItemDownloadBadRequest  %+v", 400, o.Payload)
}

func (o *CacheItemDownloadBadRequest) String() string {
	return fmt.Sprintf("[GET /apps/{app-slug}/cache-items/{cache-item-id}/download][%d] cacheItemDownloadBadRequest  %+v", 400, o.Payload)
}

func (o *CacheItemDownloadBadRequest) GetPayload() *models.ServiceStandardErrorRespModel {
	return o.Payload
}

func (o *CacheItemDownloadBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCacheItemDownloadUnauthorized creates a CacheItemDownloadUnauthorized with default headers values
func NewCacheItemDownloadUnauthorized() *CacheItemDownloadUnauthorized {
	return &CacheItemDownloadUnauthorized{}
}

/*
CacheItemDownloadUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CacheItemDownloadUnauthorized struct {
	Payload *models.ServiceStandardErrorRespModel
}

// IsSuccess returns true when this cache item download unauthorized response has a 2xx status code
func (o *CacheItemDownloadUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cache item download unauthorized response has a 3xx status code
func (o *CacheItemDownloadUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cache item download unauthorized response has a 4xx status code
func (o *CacheItemDownloadUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this cache item download unauthorized response has a 5xx status code
func (o *CacheItemDownloadUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this cache item download unauthorized response a status code equal to that given
func (o *CacheItemDownloadUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the cache item download unauthorized response
func (o *CacheItemDownloadUnauthorized) Code() int {
	return 401
}

func (o *CacheItemDownloadUnauthorized) Error() string {
	return fmt.Sprintf("[GET /apps/{app-slug}/cache-items/{cache-item-id}/download][%d] cacheItemDownloadUnauthorized  %+v", 401, o.Payload)
}

func (o *CacheItemDownloadUnauthorized) String() string {
	return fmt.Sprintf("[GET /apps/{app-slug}/cache-items/{cache-item-id}/download][%d] cacheItemDownloadUnauthorized  %+v", 401, o.Payload)
}

func (o *CacheItemDownloadUnauthorized) GetPayload() *models.ServiceStandardErrorRespModel {
	return o.Payload
}

func (o *CacheItemDownloadUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCacheItemDownloadNotFound creates a CacheItemDownloadNotFound with default headers values
func NewCacheItemDownloadNotFound() *CacheItemDownloadNotFound {
	return &CacheItemDownloadNotFound{}
}

/*
CacheItemDownloadNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CacheItemDownloadNotFound struct {
	Payload *models.ServiceStandardErrorRespModel
}

// IsSuccess returns true when this cache item download not found response has a 2xx status code
func (o *CacheItemDownloadNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cache item download not found response has a 3xx status code
func (o *CacheItemDownloadNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cache item download not found response has a 4xx status code
func (o *CacheItemDownloadNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this cache item download not found response has a 5xx status code
func (o *CacheItemDownloadNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this cache item download not found response a status code equal to that given
func (o *CacheItemDownloadNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the cache item download not found response
func (o *CacheItemDownloadNotFound) Code() int {
	return 404
}

func (o *CacheItemDownloadNotFound) Error() string {
	return fmt.Sprintf("[GET /apps/{app-slug}/cache-items/{cache-item-id}/download][%d] cacheItemDownloadNotFound  %+v", 404, o.Payload)
}

func (o *CacheItemDownloadNotFound) String() string {
	return fmt.Sprintf("[GET /apps/{app-slug}/cache-items/{cache-item-id}/download][%d] cacheItemDownloadNotFound  %+v", 404, o.Payload)
}

func (o *CacheItemDownloadNotFound) GetPayload() *models.ServiceStandardErrorRespModel {
	return o.Payload
}

func (o *CacheItemDownloadNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCacheItemDownloadInternalServerError creates a CacheItemDownloadInternalServerError with default headers values
func NewCacheItemDownloadInternalServerError() *CacheItemDownloadInternalServerError {
	return &CacheItemDownloadInternalServerError{}
}

/*
CacheItemDownloadInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type CacheItemDownloadInternalServerError struct {
	Payload *models.ServiceStandardErrorRespModel
}

// IsSuccess returns true when this cache item download internal server error response has a 2xx status code
func (o *CacheItemDownloadInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cache item download internal server error response has a 3xx status code
func (o *CacheItemDownloadInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cache item download internal server error response has a 4xx status code
func (o *CacheItemDownloadInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this cache item download internal server error response has a 5xx status code
func (o *CacheItemDownloadInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this cache item download internal server error response a status code equal to that given
func (o *CacheItemDownloadInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the cache item download internal server error response
func (o *CacheItemDownloadInternalServerError) Code() int {
	return 500
}

func (o *CacheItemDownloadInternalServerError) Error() string {
	return fmt.Sprintf("[GET /apps/{app-slug}/cache-items/{cache-item-id}/download][%d] cacheItemDownloadInternalServerError  %+v", 500, o.Payload)
}

func (o *CacheItemDownloadInternalServerError) String() string {
	return fmt.Sprintf("[GET /apps/{app-slug}/cache-items/{cache-item-id}/download][%d] cacheItemDownloadInternalServerError  %+v", 500, o.Payload)
}

func (o *CacheItemDownloadInternalServerError) GetPayload() *models.ServiceStandardErrorRespModel {
	return o.Payload
}

func (o *CacheItemDownloadInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
