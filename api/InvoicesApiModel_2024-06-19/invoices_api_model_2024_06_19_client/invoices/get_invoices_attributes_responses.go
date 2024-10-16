// Code generated by go-swagger; DO NOT EDIT.

package invoices

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/renabled/amzn-sp-api-go/api/InvoicesApiModel_2024-06-19/invoices_api_model_2024_06_19_models"
)

// GetInvoicesAttributesReader is a Reader for the GetInvoicesAttributes structure.
type GetInvoicesAttributesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInvoicesAttributesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetInvoicesAttributesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetInvoicesAttributesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetInvoicesAttributesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetInvoicesAttributesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetInvoicesAttributesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetInvoicesAttributesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetInvoicesAttributesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetInvoicesAttributesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetInvoicesAttributesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetInvoicesAttributesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetInvoicesAttributesOK creates a GetInvoicesAttributesOK with default headers values
func NewGetInvoicesAttributesOK() *GetInvoicesAttributesOK {
	return &GetInvoicesAttributesOK{}
}

/*
GetInvoicesAttributesOK describes a response with status code 200, with default header values.

Success.
*/
type GetInvoicesAttributesOK struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *invoices_api_model_2024_06_19_models.GetInvoicesAttributesResponse
}

// IsSuccess returns true when this get invoices attributes o k response has a 2xx status code
func (o *GetInvoicesAttributesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get invoices attributes o k response has a 3xx status code
func (o *GetInvoicesAttributesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get invoices attributes o k response has a 4xx status code
func (o *GetInvoicesAttributesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get invoices attributes o k response has a 5xx status code
func (o *GetInvoicesAttributesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get invoices attributes o k response a status code equal to that given
func (o *GetInvoicesAttributesOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetInvoicesAttributesOK) Error() string {
	return fmt.Sprintf("[GET /tax/invoices/2024-06-19/attributes][%d] getInvoicesAttributesOK  %+v", 200, o.Payload)
}

func (o *GetInvoicesAttributesOK) String() string {
	return fmt.Sprintf("[GET /tax/invoices/2024-06-19/attributes][%d] getInvoicesAttributesOK  %+v", 200, o.Payload)
}

func (o *GetInvoicesAttributesOK) GetPayload() *invoices_api_model_2024_06_19_models.GetInvoicesAttributesResponse {
	return o.Payload
}

func (o *GetInvoicesAttributesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(invoices_api_model_2024_06_19_models.GetInvoicesAttributesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInvoicesAttributesBadRequest creates a GetInvoicesAttributesBadRequest with default headers values
func NewGetInvoicesAttributesBadRequest() *GetInvoicesAttributesBadRequest {
	return &GetInvoicesAttributesBadRequest{}
}

/*
GetInvoicesAttributesBadRequest describes a response with status code 400, with default header values.

Request has missing or invalid parameters and cannot be parsed.
*/
type GetInvoicesAttributesBadRequest struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *invoices_api_model_2024_06_19_models.ErrorList
}

// IsSuccess returns true when this get invoices attributes bad request response has a 2xx status code
func (o *GetInvoicesAttributesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get invoices attributes bad request response has a 3xx status code
func (o *GetInvoicesAttributesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get invoices attributes bad request response has a 4xx status code
func (o *GetInvoicesAttributesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get invoices attributes bad request response has a 5xx status code
func (o *GetInvoicesAttributesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get invoices attributes bad request response a status code equal to that given
func (o *GetInvoicesAttributesBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetInvoicesAttributesBadRequest) Error() string {
	return fmt.Sprintf("[GET /tax/invoices/2024-06-19/attributes][%d] getInvoicesAttributesBadRequest  %+v", 400, o.Payload)
}

func (o *GetInvoicesAttributesBadRequest) String() string {
	return fmt.Sprintf("[GET /tax/invoices/2024-06-19/attributes][%d] getInvoicesAttributesBadRequest  %+v", 400, o.Payload)
}

func (o *GetInvoicesAttributesBadRequest) GetPayload() *invoices_api_model_2024_06_19_models.ErrorList {
	return o.Payload
}

func (o *GetInvoicesAttributesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(invoices_api_model_2024_06_19_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInvoicesAttributesUnauthorized creates a GetInvoicesAttributesUnauthorized with default headers values
func NewGetInvoicesAttributesUnauthorized() *GetInvoicesAttributesUnauthorized {
	return &GetInvoicesAttributesUnauthorized{}
}

/*
GetInvoicesAttributesUnauthorized describes a response with status code 401, with default header values.

A list of error responses returned when a request is unsuccessful.
*/
type GetInvoicesAttributesUnauthorized struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *invoices_api_model_2024_06_19_models.ErrorList
}

// IsSuccess returns true when this get invoices attributes unauthorized response has a 2xx status code
func (o *GetInvoicesAttributesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get invoices attributes unauthorized response has a 3xx status code
func (o *GetInvoicesAttributesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get invoices attributes unauthorized response has a 4xx status code
func (o *GetInvoicesAttributesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get invoices attributes unauthorized response has a 5xx status code
func (o *GetInvoicesAttributesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get invoices attributes unauthorized response a status code equal to that given
func (o *GetInvoicesAttributesUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetInvoicesAttributesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /tax/invoices/2024-06-19/attributes][%d] getInvoicesAttributesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetInvoicesAttributesUnauthorized) String() string {
	return fmt.Sprintf("[GET /tax/invoices/2024-06-19/attributes][%d] getInvoicesAttributesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetInvoicesAttributesUnauthorized) GetPayload() *invoices_api_model_2024_06_19_models.ErrorList {
	return o.Payload
}

func (o *GetInvoicesAttributesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(invoices_api_model_2024_06_19_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInvoicesAttributesForbidden creates a GetInvoicesAttributesForbidden with default headers values
func NewGetInvoicesAttributesForbidden() *GetInvoicesAttributesForbidden {
	return &GetInvoicesAttributesForbidden{}
}

/*
GetInvoicesAttributesForbidden describes a response with status code 403, with default header values.

Indicates that access to the resource is forbidden. Possible reasons include Access Denied, Unauthorized, Expired Token, or Invalid Signature.
*/
type GetInvoicesAttributesForbidden struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *invoices_api_model_2024_06_19_models.ErrorList
}

// IsSuccess returns true when this get invoices attributes forbidden response has a 2xx status code
func (o *GetInvoicesAttributesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get invoices attributes forbidden response has a 3xx status code
func (o *GetInvoicesAttributesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get invoices attributes forbidden response has a 4xx status code
func (o *GetInvoicesAttributesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get invoices attributes forbidden response has a 5xx status code
func (o *GetInvoicesAttributesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get invoices attributes forbidden response a status code equal to that given
func (o *GetInvoicesAttributesForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetInvoicesAttributesForbidden) Error() string {
	return fmt.Sprintf("[GET /tax/invoices/2024-06-19/attributes][%d] getInvoicesAttributesForbidden  %+v", 403, o.Payload)
}

func (o *GetInvoicesAttributesForbidden) String() string {
	return fmt.Sprintf("[GET /tax/invoices/2024-06-19/attributes][%d] getInvoicesAttributesForbidden  %+v", 403, o.Payload)
}

func (o *GetInvoicesAttributesForbidden) GetPayload() *invoices_api_model_2024_06_19_models.ErrorList {
	return o.Payload
}

func (o *GetInvoicesAttributesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(invoices_api_model_2024_06_19_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInvoicesAttributesNotFound creates a GetInvoicesAttributesNotFound with default headers values
func NewGetInvoicesAttributesNotFound() *GetInvoicesAttributesNotFound {
	return &GetInvoicesAttributesNotFound{}
}

/*
GetInvoicesAttributesNotFound describes a response with status code 404, with default header values.

The resource specified does not exist.
*/
type GetInvoicesAttributesNotFound struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *invoices_api_model_2024_06_19_models.ErrorList
}

// IsSuccess returns true when this get invoices attributes not found response has a 2xx status code
func (o *GetInvoicesAttributesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get invoices attributes not found response has a 3xx status code
func (o *GetInvoicesAttributesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get invoices attributes not found response has a 4xx status code
func (o *GetInvoicesAttributesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get invoices attributes not found response has a 5xx status code
func (o *GetInvoicesAttributesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get invoices attributes not found response a status code equal to that given
func (o *GetInvoicesAttributesNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetInvoicesAttributesNotFound) Error() string {
	return fmt.Sprintf("[GET /tax/invoices/2024-06-19/attributes][%d] getInvoicesAttributesNotFound  %+v", 404, o.Payload)
}

func (o *GetInvoicesAttributesNotFound) String() string {
	return fmt.Sprintf("[GET /tax/invoices/2024-06-19/attributes][%d] getInvoicesAttributesNotFound  %+v", 404, o.Payload)
}

func (o *GetInvoicesAttributesNotFound) GetPayload() *invoices_api_model_2024_06_19_models.ErrorList {
	return o.Payload
}

func (o *GetInvoicesAttributesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(invoices_api_model_2024_06_19_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInvoicesAttributesRequestEntityTooLarge creates a GetInvoicesAttributesRequestEntityTooLarge with default headers values
func NewGetInvoicesAttributesRequestEntityTooLarge() *GetInvoicesAttributesRequestEntityTooLarge {
	return &GetInvoicesAttributesRequestEntityTooLarge{}
}

/*
GetInvoicesAttributesRequestEntityTooLarge describes a response with status code 413, with default header values.

The request size exceeded the maximum accepted size.
*/
type GetInvoicesAttributesRequestEntityTooLarge struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *invoices_api_model_2024_06_19_models.ErrorList
}

// IsSuccess returns true when this get invoices attributes request entity too large response has a 2xx status code
func (o *GetInvoicesAttributesRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get invoices attributes request entity too large response has a 3xx status code
func (o *GetInvoicesAttributesRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get invoices attributes request entity too large response has a 4xx status code
func (o *GetInvoicesAttributesRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get invoices attributes request entity too large response has a 5xx status code
func (o *GetInvoicesAttributesRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get invoices attributes request entity too large response a status code equal to that given
func (o *GetInvoicesAttributesRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetInvoicesAttributesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /tax/invoices/2024-06-19/attributes][%d] getInvoicesAttributesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetInvoicesAttributesRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /tax/invoices/2024-06-19/attributes][%d] getInvoicesAttributesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetInvoicesAttributesRequestEntityTooLarge) GetPayload() *invoices_api_model_2024_06_19_models.ErrorList {
	return o.Payload
}

func (o *GetInvoicesAttributesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(invoices_api_model_2024_06_19_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInvoicesAttributesUnsupportedMediaType creates a GetInvoicesAttributesUnsupportedMediaType with default headers values
func NewGetInvoicesAttributesUnsupportedMediaType() *GetInvoicesAttributesUnsupportedMediaType {
	return &GetInvoicesAttributesUnsupportedMediaType{}
}

/*
GetInvoicesAttributesUnsupportedMediaType describes a response with status code 415, with default header values.

The request payload is in an unsupported format.
*/
type GetInvoicesAttributesUnsupportedMediaType struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *invoices_api_model_2024_06_19_models.ErrorList
}

// IsSuccess returns true when this get invoices attributes unsupported media type response has a 2xx status code
func (o *GetInvoicesAttributesUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get invoices attributes unsupported media type response has a 3xx status code
func (o *GetInvoicesAttributesUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get invoices attributes unsupported media type response has a 4xx status code
func (o *GetInvoicesAttributesUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get invoices attributes unsupported media type response has a 5xx status code
func (o *GetInvoicesAttributesUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get invoices attributes unsupported media type response a status code equal to that given
func (o *GetInvoicesAttributesUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetInvoicesAttributesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /tax/invoices/2024-06-19/attributes][%d] getInvoicesAttributesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetInvoicesAttributesUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /tax/invoices/2024-06-19/attributes][%d] getInvoicesAttributesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetInvoicesAttributesUnsupportedMediaType) GetPayload() *invoices_api_model_2024_06_19_models.ErrorList {
	return o.Payload
}

func (o *GetInvoicesAttributesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(invoices_api_model_2024_06_19_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInvoicesAttributesTooManyRequests creates a GetInvoicesAttributesTooManyRequests with default headers values
func NewGetInvoicesAttributesTooManyRequests() *GetInvoicesAttributesTooManyRequests {
	return &GetInvoicesAttributesTooManyRequests{}
}

/*
GetInvoicesAttributesTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type GetInvoicesAttributesTooManyRequests struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *invoices_api_model_2024_06_19_models.ErrorList
}

// IsSuccess returns true when this get invoices attributes too many requests response has a 2xx status code
func (o *GetInvoicesAttributesTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get invoices attributes too many requests response has a 3xx status code
func (o *GetInvoicesAttributesTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get invoices attributes too many requests response has a 4xx status code
func (o *GetInvoicesAttributesTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get invoices attributes too many requests response has a 5xx status code
func (o *GetInvoicesAttributesTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get invoices attributes too many requests response a status code equal to that given
func (o *GetInvoicesAttributesTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetInvoicesAttributesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /tax/invoices/2024-06-19/attributes][%d] getInvoicesAttributesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetInvoicesAttributesTooManyRequests) String() string {
	return fmt.Sprintf("[GET /tax/invoices/2024-06-19/attributes][%d] getInvoicesAttributesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetInvoicesAttributesTooManyRequests) GetPayload() *invoices_api_model_2024_06_19_models.ErrorList {
	return o.Payload
}

func (o *GetInvoicesAttributesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(invoices_api_model_2024_06_19_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInvoicesAttributesInternalServerError creates a GetInvoicesAttributesInternalServerError with default headers values
func NewGetInvoicesAttributesInternalServerError() *GetInvoicesAttributesInternalServerError {
	return &GetInvoicesAttributesInternalServerError{}
}

/*
GetInvoicesAttributesInternalServerError describes a response with status code 500, with default header values.

An unexpected condition occurred that prevented the server from fulfilling the request.
*/
type GetInvoicesAttributesInternalServerError struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *invoices_api_model_2024_06_19_models.ErrorList
}

// IsSuccess returns true when this get invoices attributes internal server error response has a 2xx status code
func (o *GetInvoicesAttributesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get invoices attributes internal server error response has a 3xx status code
func (o *GetInvoicesAttributesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get invoices attributes internal server error response has a 4xx status code
func (o *GetInvoicesAttributesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get invoices attributes internal server error response has a 5xx status code
func (o *GetInvoicesAttributesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get invoices attributes internal server error response a status code equal to that given
func (o *GetInvoicesAttributesInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetInvoicesAttributesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /tax/invoices/2024-06-19/attributes][%d] getInvoicesAttributesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetInvoicesAttributesInternalServerError) String() string {
	return fmt.Sprintf("[GET /tax/invoices/2024-06-19/attributes][%d] getInvoicesAttributesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetInvoicesAttributesInternalServerError) GetPayload() *invoices_api_model_2024_06_19_models.ErrorList {
	return o.Payload
}

func (o *GetInvoicesAttributesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(invoices_api_model_2024_06_19_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInvoicesAttributesServiceUnavailable creates a GetInvoicesAttributesServiceUnavailable with default headers values
func NewGetInvoicesAttributesServiceUnavailable() *GetInvoicesAttributesServiceUnavailable {
	return &GetInvoicesAttributesServiceUnavailable{}
}

/*
GetInvoicesAttributesServiceUnavailable describes a response with status code 503, with default header values.

Temporary overloading or maintenance of the server.
*/
type GetInvoicesAttributesServiceUnavailable struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *invoices_api_model_2024_06_19_models.ErrorList
}

// IsSuccess returns true when this get invoices attributes service unavailable response has a 2xx status code
func (o *GetInvoicesAttributesServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get invoices attributes service unavailable response has a 3xx status code
func (o *GetInvoicesAttributesServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get invoices attributes service unavailable response has a 4xx status code
func (o *GetInvoicesAttributesServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get invoices attributes service unavailable response has a 5xx status code
func (o *GetInvoicesAttributesServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get invoices attributes service unavailable response a status code equal to that given
func (o *GetInvoicesAttributesServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetInvoicesAttributesServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /tax/invoices/2024-06-19/attributes][%d] getInvoicesAttributesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetInvoicesAttributesServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /tax/invoices/2024-06-19/attributes][%d] getInvoicesAttributesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetInvoicesAttributesServiceUnavailable) GetPayload() *invoices_api_model_2024_06_19_models.ErrorList {
	return o.Payload
}

func (o *GetInvoicesAttributesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(invoices_api_model_2024_06_19_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
