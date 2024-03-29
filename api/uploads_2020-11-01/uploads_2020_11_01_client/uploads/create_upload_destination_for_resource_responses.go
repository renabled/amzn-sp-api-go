// Code generated by go-swagger; DO NOT EDIT.

package uploads

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/renabled/amzn-sp-api-go/api/uploads_2020-11-01/uploads_2020_11_01_models"
)

// CreateUploadDestinationForResourceReader is a Reader for the CreateUploadDestinationForResource structure.
type CreateUploadDestinationForResourceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateUploadDestinationForResourceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateUploadDestinationForResourceCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateUploadDestinationForResourceBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateUploadDestinationForResourceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateUploadDestinationForResourceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewCreateUploadDestinationForResourceRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewCreateUploadDestinationForResourceUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewCreateUploadDestinationForResourceTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateUploadDestinationForResourceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewCreateUploadDestinationForResourceServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateUploadDestinationForResourceCreated creates a CreateUploadDestinationForResourceCreated with default headers values
func NewCreateUploadDestinationForResourceCreated() *CreateUploadDestinationForResourceCreated {
	return &CreateUploadDestinationForResourceCreated{}
}

/*
CreateUploadDestinationForResourceCreated describes a response with status code 201, with default header values.

Success.
*/
type CreateUploadDestinationForResourceCreated struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *uploads_2020_11_01_models.CreateUploadDestinationResponse
}

// IsSuccess returns true when this create upload destination for resource created response has a 2xx status code
func (o *CreateUploadDestinationForResourceCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create upload destination for resource created response has a 3xx status code
func (o *CreateUploadDestinationForResourceCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create upload destination for resource created response has a 4xx status code
func (o *CreateUploadDestinationForResourceCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create upload destination for resource created response has a 5xx status code
func (o *CreateUploadDestinationForResourceCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create upload destination for resource created response a status code equal to that given
func (o *CreateUploadDestinationForResourceCreated) IsCode(code int) bool {
	return code == 201
}

func (o *CreateUploadDestinationForResourceCreated) Error() string {
	return fmt.Sprintf("[POST /uploads/2020-11-01/uploadDestinations/{resource}][%d] createUploadDestinationForResourceCreated  %+v", 201, o.Payload)
}

func (o *CreateUploadDestinationForResourceCreated) String() string {
	return fmt.Sprintf("[POST /uploads/2020-11-01/uploadDestinations/{resource}][%d] createUploadDestinationForResourceCreated  %+v", 201, o.Payload)
}

func (o *CreateUploadDestinationForResourceCreated) GetPayload() *uploads_2020_11_01_models.CreateUploadDestinationResponse {
	return o.Payload
}

func (o *CreateUploadDestinationForResourceCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RateLimit-Limit
	hdrXAmznRateLimitLimit := response.GetHeader("x-amzn-RateLimit-Limit")

	if hdrXAmznRateLimitLimit != "" {
		o.XAmznRateLimitLimit = hdrXAmznRateLimitLimit
	}

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(uploads_2020_11_01_models.CreateUploadDestinationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateUploadDestinationForResourceBadRequest creates a CreateUploadDestinationForResourceBadRequest with default headers values
func NewCreateUploadDestinationForResourceBadRequest() *CreateUploadDestinationForResourceBadRequest {
	return &CreateUploadDestinationForResourceBadRequest{}
}

/*
CreateUploadDestinationForResourceBadRequest describes a response with status code 400, with default header values.

Request has missing or invalid parameters and cannot be parsed.
*/
type CreateUploadDestinationForResourceBadRequest struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *uploads_2020_11_01_models.CreateUploadDestinationResponse
}

// IsSuccess returns true when this create upload destination for resource bad request response has a 2xx status code
func (o *CreateUploadDestinationForResourceBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create upload destination for resource bad request response has a 3xx status code
func (o *CreateUploadDestinationForResourceBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create upload destination for resource bad request response has a 4xx status code
func (o *CreateUploadDestinationForResourceBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create upload destination for resource bad request response has a 5xx status code
func (o *CreateUploadDestinationForResourceBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create upload destination for resource bad request response a status code equal to that given
func (o *CreateUploadDestinationForResourceBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *CreateUploadDestinationForResourceBadRequest) Error() string {
	return fmt.Sprintf("[POST /uploads/2020-11-01/uploadDestinations/{resource}][%d] createUploadDestinationForResourceBadRequest  %+v", 400, o.Payload)
}

func (o *CreateUploadDestinationForResourceBadRequest) String() string {
	return fmt.Sprintf("[POST /uploads/2020-11-01/uploadDestinations/{resource}][%d] createUploadDestinationForResourceBadRequest  %+v", 400, o.Payload)
}

func (o *CreateUploadDestinationForResourceBadRequest) GetPayload() *uploads_2020_11_01_models.CreateUploadDestinationResponse {
	return o.Payload
}

func (o *CreateUploadDestinationForResourceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RateLimit-Limit
	hdrXAmznRateLimitLimit := response.GetHeader("x-amzn-RateLimit-Limit")

	if hdrXAmznRateLimitLimit != "" {
		o.XAmznRateLimitLimit = hdrXAmznRateLimitLimit
	}

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(uploads_2020_11_01_models.CreateUploadDestinationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateUploadDestinationForResourceForbidden creates a CreateUploadDestinationForResourceForbidden with default headers values
func NewCreateUploadDestinationForResourceForbidden() *CreateUploadDestinationForResourceForbidden {
	return &CreateUploadDestinationForResourceForbidden{}
}

/*
CreateUploadDestinationForResourceForbidden describes a response with status code 403, with default header values.

Indicates that access to the resource is forbidden. Possible reasons include Access Denied, Unauthorized, Expired Token, or Invalid Signature.
*/
type CreateUploadDestinationForResourceForbidden struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *uploads_2020_11_01_models.CreateUploadDestinationResponse
}

// IsSuccess returns true when this create upload destination for resource forbidden response has a 2xx status code
func (o *CreateUploadDestinationForResourceForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create upload destination for resource forbidden response has a 3xx status code
func (o *CreateUploadDestinationForResourceForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create upload destination for resource forbidden response has a 4xx status code
func (o *CreateUploadDestinationForResourceForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create upload destination for resource forbidden response has a 5xx status code
func (o *CreateUploadDestinationForResourceForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create upload destination for resource forbidden response a status code equal to that given
func (o *CreateUploadDestinationForResourceForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *CreateUploadDestinationForResourceForbidden) Error() string {
	return fmt.Sprintf("[POST /uploads/2020-11-01/uploadDestinations/{resource}][%d] createUploadDestinationForResourceForbidden  %+v", 403, o.Payload)
}

func (o *CreateUploadDestinationForResourceForbidden) String() string {
	return fmt.Sprintf("[POST /uploads/2020-11-01/uploadDestinations/{resource}][%d] createUploadDestinationForResourceForbidden  %+v", 403, o.Payload)
}

func (o *CreateUploadDestinationForResourceForbidden) GetPayload() *uploads_2020_11_01_models.CreateUploadDestinationResponse {
	return o.Payload
}

func (o *CreateUploadDestinationForResourceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(uploads_2020_11_01_models.CreateUploadDestinationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateUploadDestinationForResourceNotFound creates a CreateUploadDestinationForResourceNotFound with default headers values
func NewCreateUploadDestinationForResourceNotFound() *CreateUploadDestinationForResourceNotFound {
	return &CreateUploadDestinationForResourceNotFound{}
}

/*
CreateUploadDestinationForResourceNotFound describes a response with status code 404, with default header values.

The resource specified does not exist.
*/
type CreateUploadDestinationForResourceNotFound struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *uploads_2020_11_01_models.CreateUploadDestinationResponse
}

// IsSuccess returns true when this create upload destination for resource not found response has a 2xx status code
func (o *CreateUploadDestinationForResourceNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create upload destination for resource not found response has a 3xx status code
func (o *CreateUploadDestinationForResourceNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create upload destination for resource not found response has a 4xx status code
func (o *CreateUploadDestinationForResourceNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this create upload destination for resource not found response has a 5xx status code
func (o *CreateUploadDestinationForResourceNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this create upload destination for resource not found response a status code equal to that given
func (o *CreateUploadDestinationForResourceNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *CreateUploadDestinationForResourceNotFound) Error() string {
	return fmt.Sprintf("[POST /uploads/2020-11-01/uploadDestinations/{resource}][%d] createUploadDestinationForResourceNotFound  %+v", 404, o.Payload)
}

func (o *CreateUploadDestinationForResourceNotFound) String() string {
	return fmt.Sprintf("[POST /uploads/2020-11-01/uploadDestinations/{resource}][%d] createUploadDestinationForResourceNotFound  %+v", 404, o.Payload)
}

func (o *CreateUploadDestinationForResourceNotFound) GetPayload() *uploads_2020_11_01_models.CreateUploadDestinationResponse {
	return o.Payload
}

func (o *CreateUploadDestinationForResourceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RateLimit-Limit
	hdrXAmznRateLimitLimit := response.GetHeader("x-amzn-RateLimit-Limit")

	if hdrXAmznRateLimitLimit != "" {
		o.XAmznRateLimitLimit = hdrXAmznRateLimitLimit
	}

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(uploads_2020_11_01_models.CreateUploadDestinationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateUploadDestinationForResourceRequestEntityTooLarge creates a CreateUploadDestinationForResourceRequestEntityTooLarge with default headers values
func NewCreateUploadDestinationForResourceRequestEntityTooLarge() *CreateUploadDestinationForResourceRequestEntityTooLarge {
	return &CreateUploadDestinationForResourceRequestEntityTooLarge{}
}

/*
CreateUploadDestinationForResourceRequestEntityTooLarge describes a response with status code 413, with default header values.

The request size exceeded the maximum accepted size.
*/
type CreateUploadDestinationForResourceRequestEntityTooLarge struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *uploads_2020_11_01_models.CreateUploadDestinationResponse
}

// IsSuccess returns true when this create upload destination for resource request entity too large response has a 2xx status code
func (o *CreateUploadDestinationForResourceRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create upload destination for resource request entity too large response has a 3xx status code
func (o *CreateUploadDestinationForResourceRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create upload destination for resource request entity too large response has a 4xx status code
func (o *CreateUploadDestinationForResourceRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this create upload destination for resource request entity too large response has a 5xx status code
func (o *CreateUploadDestinationForResourceRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this create upload destination for resource request entity too large response a status code equal to that given
func (o *CreateUploadDestinationForResourceRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *CreateUploadDestinationForResourceRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /uploads/2020-11-01/uploadDestinations/{resource}][%d] createUploadDestinationForResourceRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *CreateUploadDestinationForResourceRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[POST /uploads/2020-11-01/uploadDestinations/{resource}][%d] createUploadDestinationForResourceRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *CreateUploadDestinationForResourceRequestEntityTooLarge) GetPayload() *uploads_2020_11_01_models.CreateUploadDestinationResponse {
	return o.Payload
}

func (o *CreateUploadDestinationForResourceRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RateLimit-Limit
	hdrXAmznRateLimitLimit := response.GetHeader("x-amzn-RateLimit-Limit")

	if hdrXAmznRateLimitLimit != "" {
		o.XAmznRateLimitLimit = hdrXAmznRateLimitLimit
	}

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(uploads_2020_11_01_models.CreateUploadDestinationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateUploadDestinationForResourceUnsupportedMediaType creates a CreateUploadDestinationForResourceUnsupportedMediaType with default headers values
func NewCreateUploadDestinationForResourceUnsupportedMediaType() *CreateUploadDestinationForResourceUnsupportedMediaType {
	return &CreateUploadDestinationForResourceUnsupportedMediaType{}
}

/*
CreateUploadDestinationForResourceUnsupportedMediaType describes a response with status code 415, with default header values.

The request payload is in an unsupported format.
*/
type CreateUploadDestinationForResourceUnsupportedMediaType struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *uploads_2020_11_01_models.CreateUploadDestinationResponse
}

// IsSuccess returns true when this create upload destination for resource unsupported media type response has a 2xx status code
func (o *CreateUploadDestinationForResourceUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create upload destination for resource unsupported media type response has a 3xx status code
func (o *CreateUploadDestinationForResourceUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create upload destination for resource unsupported media type response has a 4xx status code
func (o *CreateUploadDestinationForResourceUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this create upload destination for resource unsupported media type response has a 5xx status code
func (o *CreateUploadDestinationForResourceUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this create upload destination for resource unsupported media type response a status code equal to that given
func (o *CreateUploadDestinationForResourceUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *CreateUploadDestinationForResourceUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /uploads/2020-11-01/uploadDestinations/{resource}][%d] createUploadDestinationForResourceUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *CreateUploadDestinationForResourceUnsupportedMediaType) String() string {
	return fmt.Sprintf("[POST /uploads/2020-11-01/uploadDestinations/{resource}][%d] createUploadDestinationForResourceUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *CreateUploadDestinationForResourceUnsupportedMediaType) GetPayload() *uploads_2020_11_01_models.CreateUploadDestinationResponse {
	return o.Payload
}

func (o *CreateUploadDestinationForResourceUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RateLimit-Limit
	hdrXAmznRateLimitLimit := response.GetHeader("x-amzn-RateLimit-Limit")

	if hdrXAmznRateLimitLimit != "" {
		o.XAmznRateLimitLimit = hdrXAmznRateLimitLimit
	}

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(uploads_2020_11_01_models.CreateUploadDestinationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateUploadDestinationForResourceTooManyRequests creates a CreateUploadDestinationForResourceTooManyRequests with default headers values
func NewCreateUploadDestinationForResourceTooManyRequests() *CreateUploadDestinationForResourceTooManyRequests {
	return &CreateUploadDestinationForResourceTooManyRequests{}
}

/*
CreateUploadDestinationForResourceTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type CreateUploadDestinationForResourceTooManyRequests struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *uploads_2020_11_01_models.CreateUploadDestinationResponse
}

// IsSuccess returns true when this create upload destination for resource too many requests response has a 2xx status code
func (o *CreateUploadDestinationForResourceTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create upload destination for resource too many requests response has a 3xx status code
func (o *CreateUploadDestinationForResourceTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create upload destination for resource too many requests response has a 4xx status code
func (o *CreateUploadDestinationForResourceTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this create upload destination for resource too many requests response has a 5xx status code
func (o *CreateUploadDestinationForResourceTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this create upload destination for resource too many requests response a status code equal to that given
func (o *CreateUploadDestinationForResourceTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *CreateUploadDestinationForResourceTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /uploads/2020-11-01/uploadDestinations/{resource}][%d] createUploadDestinationForResourceTooManyRequests  %+v", 429, o.Payload)
}

func (o *CreateUploadDestinationForResourceTooManyRequests) String() string {
	return fmt.Sprintf("[POST /uploads/2020-11-01/uploadDestinations/{resource}][%d] createUploadDestinationForResourceTooManyRequests  %+v", 429, o.Payload)
}

func (o *CreateUploadDestinationForResourceTooManyRequests) GetPayload() *uploads_2020_11_01_models.CreateUploadDestinationResponse {
	return o.Payload
}

func (o *CreateUploadDestinationForResourceTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RateLimit-Limit
	hdrXAmznRateLimitLimit := response.GetHeader("x-amzn-RateLimit-Limit")

	if hdrXAmznRateLimitLimit != "" {
		o.XAmznRateLimitLimit = hdrXAmznRateLimitLimit
	}

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(uploads_2020_11_01_models.CreateUploadDestinationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateUploadDestinationForResourceInternalServerError creates a CreateUploadDestinationForResourceInternalServerError with default headers values
func NewCreateUploadDestinationForResourceInternalServerError() *CreateUploadDestinationForResourceInternalServerError {
	return &CreateUploadDestinationForResourceInternalServerError{}
}

/*
CreateUploadDestinationForResourceInternalServerError describes a response with status code 500, with default header values.

An unexpected condition occurred that prevented the server from fulfilling the request.
*/
type CreateUploadDestinationForResourceInternalServerError struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *uploads_2020_11_01_models.CreateUploadDestinationResponse
}

// IsSuccess returns true when this create upload destination for resource internal server error response has a 2xx status code
func (o *CreateUploadDestinationForResourceInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create upload destination for resource internal server error response has a 3xx status code
func (o *CreateUploadDestinationForResourceInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create upload destination for resource internal server error response has a 4xx status code
func (o *CreateUploadDestinationForResourceInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create upload destination for resource internal server error response has a 5xx status code
func (o *CreateUploadDestinationForResourceInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create upload destination for resource internal server error response a status code equal to that given
func (o *CreateUploadDestinationForResourceInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *CreateUploadDestinationForResourceInternalServerError) Error() string {
	return fmt.Sprintf("[POST /uploads/2020-11-01/uploadDestinations/{resource}][%d] createUploadDestinationForResourceInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateUploadDestinationForResourceInternalServerError) String() string {
	return fmt.Sprintf("[POST /uploads/2020-11-01/uploadDestinations/{resource}][%d] createUploadDestinationForResourceInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateUploadDestinationForResourceInternalServerError) GetPayload() *uploads_2020_11_01_models.CreateUploadDestinationResponse {
	return o.Payload
}

func (o *CreateUploadDestinationForResourceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RateLimit-Limit
	hdrXAmznRateLimitLimit := response.GetHeader("x-amzn-RateLimit-Limit")

	if hdrXAmznRateLimitLimit != "" {
		o.XAmznRateLimitLimit = hdrXAmznRateLimitLimit
	}

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(uploads_2020_11_01_models.CreateUploadDestinationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateUploadDestinationForResourceServiceUnavailable creates a CreateUploadDestinationForResourceServiceUnavailable with default headers values
func NewCreateUploadDestinationForResourceServiceUnavailable() *CreateUploadDestinationForResourceServiceUnavailable {
	return &CreateUploadDestinationForResourceServiceUnavailable{}
}

/*
CreateUploadDestinationForResourceServiceUnavailable describes a response with status code 503, with default header values.

Temporary overloading or maintenance of the server.
*/
type CreateUploadDestinationForResourceServiceUnavailable struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *uploads_2020_11_01_models.CreateUploadDestinationResponse
}

// IsSuccess returns true when this create upload destination for resource service unavailable response has a 2xx status code
func (o *CreateUploadDestinationForResourceServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create upload destination for resource service unavailable response has a 3xx status code
func (o *CreateUploadDestinationForResourceServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create upload destination for resource service unavailable response has a 4xx status code
func (o *CreateUploadDestinationForResourceServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this create upload destination for resource service unavailable response has a 5xx status code
func (o *CreateUploadDestinationForResourceServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this create upload destination for resource service unavailable response a status code equal to that given
func (o *CreateUploadDestinationForResourceServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *CreateUploadDestinationForResourceServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /uploads/2020-11-01/uploadDestinations/{resource}][%d] createUploadDestinationForResourceServiceUnavailable  %+v", 503, o.Payload)
}

func (o *CreateUploadDestinationForResourceServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /uploads/2020-11-01/uploadDestinations/{resource}][%d] createUploadDestinationForResourceServiceUnavailable  %+v", 503, o.Payload)
}

func (o *CreateUploadDestinationForResourceServiceUnavailable) GetPayload() *uploads_2020_11_01_models.CreateUploadDestinationResponse {
	return o.Payload
}

func (o *CreateUploadDestinationForResourceServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RateLimit-Limit
	hdrXAmznRateLimitLimit := response.GetHeader("x-amzn-RateLimit-Limit")

	if hdrXAmznRateLimitLimit != "" {
		o.XAmznRateLimitLimit = hdrXAmznRateLimitLimit
	}

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(uploads_2020_11_01_models.CreateUploadDestinationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
