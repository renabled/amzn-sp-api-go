// Code generated by go-swagger; DO NOT EDIT.

package aplus_content

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/renabled/amzn-sp-api-go/api/aplusContent_2020-11-01/aplus_content_2020_11_01_models"
)

// ValidateContentDocumentAsinRelationsReader is a Reader for the ValidateContentDocumentAsinRelations structure.
type ValidateContentDocumentAsinRelationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ValidateContentDocumentAsinRelationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewValidateContentDocumentAsinRelationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewValidateContentDocumentAsinRelationsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewValidateContentDocumentAsinRelationsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewValidateContentDocumentAsinRelationsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewValidateContentDocumentAsinRelationsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewValidateContentDocumentAsinRelationsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewValidateContentDocumentAsinRelationsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewValidateContentDocumentAsinRelationsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewValidateContentDocumentAsinRelationsOK creates a ValidateContentDocumentAsinRelationsOK with default headers values
func NewValidateContentDocumentAsinRelationsOK() *ValidateContentDocumentAsinRelationsOK {
	return &ValidateContentDocumentAsinRelationsOK{}
}

/*
ValidateContentDocumentAsinRelationsOK describes a response with status code 200, with default header values.

Success.
*/
type ValidateContentDocumentAsinRelationsOK struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *aplus_content_2020_11_01_models.ValidateContentDocumentAsinRelationsResponse
}

// IsSuccess returns true when this validate content document asin relations o k response has a 2xx status code
func (o *ValidateContentDocumentAsinRelationsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this validate content document asin relations o k response has a 3xx status code
func (o *ValidateContentDocumentAsinRelationsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this validate content document asin relations o k response has a 4xx status code
func (o *ValidateContentDocumentAsinRelationsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this validate content document asin relations o k response has a 5xx status code
func (o *ValidateContentDocumentAsinRelationsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this validate content document asin relations o k response a status code equal to that given
func (o *ValidateContentDocumentAsinRelationsOK) IsCode(code int) bool {
	return code == 200
}

func (o *ValidateContentDocumentAsinRelationsOK) Error() string {
	return fmt.Sprintf("[POST /aplus/2020-11-01/contentAsinValidations][%d] validateContentDocumentAsinRelationsOK  %+v", 200, o.Payload)
}

func (o *ValidateContentDocumentAsinRelationsOK) String() string {
	return fmt.Sprintf("[POST /aplus/2020-11-01/contentAsinValidations][%d] validateContentDocumentAsinRelationsOK  %+v", 200, o.Payload)
}

func (o *ValidateContentDocumentAsinRelationsOK) GetPayload() *aplus_content_2020_11_01_models.ValidateContentDocumentAsinRelationsResponse {
	return o.Payload
}

func (o *ValidateContentDocumentAsinRelationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(aplus_content_2020_11_01_models.ValidateContentDocumentAsinRelationsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewValidateContentDocumentAsinRelationsBadRequest creates a ValidateContentDocumentAsinRelationsBadRequest with default headers values
func NewValidateContentDocumentAsinRelationsBadRequest() *ValidateContentDocumentAsinRelationsBadRequest {
	return &ValidateContentDocumentAsinRelationsBadRequest{}
}

/*
ValidateContentDocumentAsinRelationsBadRequest describes a response with status code 400, with default header values.

Request has missing or invalid parameters and cannot be parsed.
*/
type ValidateContentDocumentAsinRelationsBadRequest struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *aplus_content_2020_11_01_models.ErrorList
}

// IsSuccess returns true when this validate content document asin relations bad request response has a 2xx status code
func (o *ValidateContentDocumentAsinRelationsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this validate content document asin relations bad request response has a 3xx status code
func (o *ValidateContentDocumentAsinRelationsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this validate content document asin relations bad request response has a 4xx status code
func (o *ValidateContentDocumentAsinRelationsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this validate content document asin relations bad request response has a 5xx status code
func (o *ValidateContentDocumentAsinRelationsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this validate content document asin relations bad request response a status code equal to that given
func (o *ValidateContentDocumentAsinRelationsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *ValidateContentDocumentAsinRelationsBadRequest) Error() string {
	return fmt.Sprintf("[POST /aplus/2020-11-01/contentAsinValidations][%d] validateContentDocumentAsinRelationsBadRequest  %+v", 400, o.Payload)
}

func (o *ValidateContentDocumentAsinRelationsBadRequest) String() string {
	return fmt.Sprintf("[POST /aplus/2020-11-01/contentAsinValidations][%d] validateContentDocumentAsinRelationsBadRequest  %+v", 400, o.Payload)
}

func (o *ValidateContentDocumentAsinRelationsBadRequest) GetPayload() *aplus_content_2020_11_01_models.ErrorList {
	return o.Payload
}

func (o *ValidateContentDocumentAsinRelationsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(aplus_content_2020_11_01_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewValidateContentDocumentAsinRelationsUnauthorized creates a ValidateContentDocumentAsinRelationsUnauthorized with default headers values
func NewValidateContentDocumentAsinRelationsUnauthorized() *ValidateContentDocumentAsinRelationsUnauthorized {
	return &ValidateContentDocumentAsinRelationsUnauthorized{}
}

/*
ValidateContentDocumentAsinRelationsUnauthorized describes a response with status code 401, with default header values.

The request's Authorization header is not formatted correctly or does not contain a valid token.
*/
type ValidateContentDocumentAsinRelationsUnauthorized struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *aplus_content_2020_11_01_models.ErrorList
}

// IsSuccess returns true when this validate content document asin relations unauthorized response has a 2xx status code
func (o *ValidateContentDocumentAsinRelationsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this validate content document asin relations unauthorized response has a 3xx status code
func (o *ValidateContentDocumentAsinRelationsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this validate content document asin relations unauthorized response has a 4xx status code
func (o *ValidateContentDocumentAsinRelationsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this validate content document asin relations unauthorized response has a 5xx status code
func (o *ValidateContentDocumentAsinRelationsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this validate content document asin relations unauthorized response a status code equal to that given
func (o *ValidateContentDocumentAsinRelationsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *ValidateContentDocumentAsinRelationsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /aplus/2020-11-01/contentAsinValidations][%d] validateContentDocumentAsinRelationsUnauthorized  %+v", 401, o.Payload)
}

func (o *ValidateContentDocumentAsinRelationsUnauthorized) String() string {
	return fmt.Sprintf("[POST /aplus/2020-11-01/contentAsinValidations][%d] validateContentDocumentAsinRelationsUnauthorized  %+v", 401, o.Payload)
}

func (o *ValidateContentDocumentAsinRelationsUnauthorized) GetPayload() *aplus_content_2020_11_01_models.ErrorList {
	return o.Payload
}

func (o *ValidateContentDocumentAsinRelationsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(aplus_content_2020_11_01_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewValidateContentDocumentAsinRelationsForbidden creates a ValidateContentDocumentAsinRelationsForbidden with default headers values
func NewValidateContentDocumentAsinRelationsForbidden() *ValidateContentDocumentAsinRelationsForbidden {
	return &ValidateContentDocumentAsinRelationsForbidden{}
}

/*
ValidateContentDocumentAsinRelationsForbidden describes a response with status code 403, with default header values.

Indicates that access to the resource is forbidden. Possible reasons include Access Denied, Unauthorized, Expired Token, or Invalid Signature.
*/
type ValidateContentDocumentAsinRelationsForbidden struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *aplus_content_2020_11_01_models.ErrorList
}

// IsSuccess returns true when this validate content document asin relations forbidden response has a 2xx status code
func (o *ValidateContentDocumentAsinRelationsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this validate content document asin relations forbidden response has a 3xx status code
func (o *ValidateContentDocumentAsinRelationsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this validate content document asin relations forbidden response has a 4xx status code
func (o *ValidateContentDocumentAsinRelationsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this validate content document asin relations forbidden response has a 5xx status code
func (o *ValidateContentDocumentAsinRelationsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this validate content document asin relations forbidden response a status code equal to that given
func (o *ValidateContentDocumentAsinRelationsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ValidateContentDocumentAsinRelationsForbidden) Error() string {
	return fmt.Sprintf("[POST /aplus/2020-11-01/contentAsinValidations][%d] validateContentDocumentAsinRelationsForbidden  %+v", 403, o.Payload)
}

func (o *ValidateContentDocumentAsinRelationsForbidden) String() string {
	return fmt.Sprintf("[POST /aplus/2020-11-01/contentAsinValidations][%d] validateContentDocumentAsinRelationsForbidden  %+v", 403, o.Payload)
}

func (o *ValidateContentDocumentAsinRelationsForbidden) GetPayload() *aplus_content_2020_11_01_models.ErrorList {
	return o.Payload
}

func (o *ValidateContentDocumentAsinRelationsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(aplus_content_2020_11_01_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewValidateContentDocumentAsinRelationsNotFound creates a ValidateContentDocumentAsinRelationsNotFound with default headers values
func NewValidateContentDocumentAsinRelationsNotFound() *ValidateContentDocumentAsinRelationsNotFound {
	return &ValidateContentDocumentAsinRelationsNotFound{}
}

/*
ValidateContentDocumentAsinRelationsNotFound describes a response with status code 404, with default header values.

The specified resource does not exist.
*/
type ValidateContentDocumentAsinRelationsNotFound struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *aplus_content_2020_11_01_models.ErrorList
}

// IsSuccess returns true when this validate content document asin relations not found response has a 2xx status code
func (o *ValidateContentDocumentAsinRelationsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this validate content document asin relations not found response has a 3xx status code
func (o *ValidateContentDocumentAsinRelationsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this validate content document asin relations not found response has a 4xx status code
func (o *ValidateContentDocumentAsinRelationsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this validate content document asin relations not found response has a 5xx status code
func (o *ValidateContentDocumentAsinRelationsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this validate content document asin relations not found response a status code equal to that given
func (o *ValidateContentDocumentAsinRelationsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *ValidateContentDocumentAsinRelationsNotFound) Error() string {
	return fmt.Sprintf("[POST /aplus/2020-11-01/contentAsinValidations][%d] validateContentDocumentAsinRelationsNotFound  %+v", 404, o.Payload)
}

func (o *ValidateContentDocumentAsinRelationsNotFound) String() string {
	return fmt.Sprintf("[POST /aplus/2020-11-01/contentAsinValidations][%d] validateContentDocumentAsinRelationsNotFound  %+v", 404, o.Payload)
}

func (o *ValidateContentDocumentAsinRelationsNotFound) GetPayload() *aplus_content_2020_11_01_models.ErrorList {
	return o.Payload
}

func (o *ValidateContentDocumentAsinRelationsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(aplus_content_2020_11_01_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewValidateContentDocumentAsinRelationsTooManyRequests creates a ValidateContentDocumentAsinRelationsTooManyRequests with default headers values
func NewValidateContentDocumentAsinRelationsTooManyRequests() *ValidateContentDocumentAsinRelationsTooManyRequests {
	return &ValidateContentDocumentAsinRelationsTooManyRequests{}
}

/*
ValidateContentDocumentAsinRelationsTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type ValidateContentDocumentAsinRelationsTooManyRequests struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *aplus_content_2020_11_01_models.ErrorList
}

// IsSuccess returns true when this validate content document asin relations too many requests response has a 2xx status code
func (o *ValidateContentDocumentAsinRelationsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this validate content document asin relations too many requests response has a 3xx status code
func (o *ValidateContentDocumentAsinRelationsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this validate content document asin relations too many requests response has a 4xx status code
func (o *ValidateContentDocumentAsinRelationsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this validate content document asin relations too many requests response has a 5xx status code
func (o *ValidateContentDocumentAsinRelationsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this validate content document asin relations too many requests response a status code equal to that given
func (o *ValidateContentDocumentAsinRelationsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *ValidateContentDocumentAsinRelationsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /aplus/2020-11-01/contentAsinValidations][%d] validateContentDocumentAsinRelationsTooManyRequests  %+v", 429, o.Payload)
}

func (o *ValidateContentDocumentAsinRelationsTooManyRequests) String() string {
	return fmt.Sprintf("[POST /aplus/2020-11-01/contentAsinValidations][%d] validateContentDocumentAsinRelationsTooManyRequests  %+v", 429, o.Payload)
}

func (o *ValidateContentDocumentAsinRelationsTooManyRequests) GetPayload() *aplus_content_2020_11_01_models.ErrorList {
	return o.Payload
}

func (o *ValidateContentDocumentAsinRelationsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(aplus_content_2020_11_01_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewValidateContentDocumentAsinRelationsInternalServerError creates a ValidateContentDocumentAsinRelationsInternalServerError with default headers values
func NewValidateContentDocumentAsinRelationsInternalServerError() *ValidateContentDocumentAsinRelationsInternalServerError {
	return &ValidateContentDocumentAsinRelationsInternalServerError{}
}

/*
ValidateContentDocumentAsinRelationsInternalServerError describes a response with status code 500, with default header values.

An unexpected condition occurred that prevented the server from fulfilling the request.
*/
type ValidateContentDocumentAsinRelationsInternalServerError struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *aplus_content_2020_11_01_models.ErrorList
}

// IsSuccess returns true when this validate content document asin relations internal server error response has a 2xx status code
func (o *ValidateContentDocumentAsinRelationsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this validate content document asin relations internal server error response has a 3xx status code
func (o *ValidateContentDocumentAsinRelationsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this validate content document asin relations internal server error response has a 4xx status code
func (o *ValidateContentDocumentAsinRelationsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this validate content document asin relations internal server error response has a 5xx status code
func (o *ValidateContentDocumentAsinRelationsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this validate content document asin relations internal server error response a status code equal to that given
func (o *ValidateContentDocumentAsinRelationsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ValidateContentDocumentAsinRelationsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /aplus/2020-11-01/contentAsinValidations][%d] validateContentDocumentAsinRelationsInternalServerError  %+v", 500, o.Payload)
}

func (o *ValidateContentDocumentAsinRelationsInternalServerError) String() string {
	return fmt.Sprintf("[POST /aplus/2020-11-01/contentAsinValidations][%d] validateContentDocumentAsinRelationsInternalServerError  %+v", 500, o.Payload)
}

func (o *ValidateContentDocumentAsinRelationsInternalServerError) GetPayload() *aplus_content_2020_11_01_models.ErrorList {
	return o.Payload
}

func (o *ValidateContentDocumentAsinRelationsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(aplus_content_2020_11_01_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewValidateContentDocumentAsinRelationsServiceUnavailable creates a ValidateContentDocumentAsinRelationsServiceUnavailable with default headers values
func NewValidateContentDocumentAsinRelationsServiceUnavailable() *ValidateContentDocumentAsinRelationsServiceUnavailable {
	return &ValidateContentDocumentAsinRelationsServiceUnavailable{}
}

/*
ValidateContentDocumentAsinRelationsServiceUnavailable describes a response with status code 503, with default header values.

Temporary overloading or maintenance of the server.
*/
type ValidateContentDocumentAsinRelationsServiceUnavailable struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *aplus_content_2020_11_01_models.ErrorList
}

// IsSuccess returns true when this validate content document asin relations service unavailable response has a 2xx status code
func (o *ValidateContentDocumentAsinRelationsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this validate content document asin relations service unavailable response has a 3xx status code
func (o *ValidateContentDocumentAsinRelationsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this validate content document asin relations service unavailable response has a 4xx status code
func (o *ValidateContentDocumentAsinRelationsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this validate content document asin relations service unavailable response has a 5xx status code
func (o *ValidateContentDocumentAsinRelationsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this validate content document asin relations service unavailable response a status code equal to that given
func (o *ValidateContentDocumentAsinRelationsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *ValidateContentDocumentAsinRelationsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /aplus/2020-11-01/contentAsinValidations][%d] validateContentDocumentAsinRelationsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *ValidateContentDocumentAsinRelationsServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /aplus/2020-11-01/contentAsinValidations][%d] validateContentDocumentAsinRelationsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *ValidateContentDocumentAsinRelationsServiceUnavailable) GetPayload() *aplus_content_2020_11_01_models.ErrorList {
	return o.Payload
}

func (o *ValidateContentDocumentAsinRelationsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(aplus_content_2020_11_01_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
