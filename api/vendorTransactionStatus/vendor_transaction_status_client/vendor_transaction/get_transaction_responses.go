// Code generated by go-swagger; DO NOT EDIT.

package vendor_transaction

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/renabled/amzn-sp-api-go/api/vendorTransactionStatus/vendor_transaction_status_models"
)

// GetTransactionReader is a Reader for the GetTransaction structure.
type GetTransactionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTransactionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTransactionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetTransactionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetTransactionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetTransactionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetTransactionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetTransactionUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetTransactionTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetTransactionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetTransactionServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetTransactionOK creates a GetTransactionOK with default headers values
func NewGetTransactionOK() *GetTransactionOK {
	return &GetTransactionOK{}
}

/*
GetTransactionOK describes a response with status code 200, with default header values.

Success.
*/
type GetTransactionOK struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_transaction_status_models.GetTransactionResponse
}

// IsSuccess returns true when this get transaction o k response has a 2xx status code
func (o *GetTransactionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get transaction o k response has a 3xx status code
func (o *GetTransactionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get transaction o k response has a 4xx status code
func (o *GetTransactionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get transaction o k response has a 5xx status code
func (o *GetTransactionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get transaction o k response a status code equal to that given
func (o *GetTransactionOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetTransactionOK) Error() string {
	return fmt.Sprintf("[GET /vendor/transactions/v1/transactions/{transactionId}][%d] getTransactionOK  %+v", 200, o.Payload)
}

func (o *GetTransactionOK) String() string {
	return fmt.Sprintf("[GET /vendor/transactions/v1/transactions/{transactionId}][%d] getTransactionOK  %+v", 200, o.Payload)
}

func (o *GetTransactionOK) GetPayload() *vendor_transaction_status_models.GetTransactionResponse {
	return o.Payload
}

func (o *GetTransactionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(vendor_transaction_status_models.GetTransactionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTransactionBadRequest creates a GetTransactionBadRequest with default headers values
func NewGetTransactionBadRequest() *GetTransactionBadRequest {
	return &GetTransactionBadRequest{}
}

/*
GetTransactionBadRequest describes a response with status code 400, with default header values.

Request has missing or not valid parameters and cannot be parsed.
*/
type GetTransactionBadRequest struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_transaction_status_models.GetTransactionResponse
}

// IsSuccess returns true when this get transaction bad request response has a 2xx status code
func (o *GetTransactionBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get transaction bad request response has a 3xx status code
func (o *GetTransactionBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get transaction bad request response has a 4xx status code
func (o *GetTransactionBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get transaction bad request response has a 5xx status code
func (o *GetTransactionBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get transaction bad request response a status code equal to that given
func (o *GetTransactionBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetTransactionBadRequest) Error() string {
	return fmt.Sprintf("[GET /vendor/transactions/v1/transactions/{transactionId}][%d] getTransactionBadRequest  %+v", 400, o.Payload)
}

func (o *GetTransactionBadRequest) String() string {
	return fmt.Sprintf("[GET /vendor/transactions/v1/transactions/{transactionId}][%d] getTransactionBadRequest  %+v", 400, o.Payload)
}

func (o *GetTransactionBadRequest) GetPayload() *vendor_transaction_status_models.GetTransactionResponse {
	return o.Payload
}

func (o *GetTransactionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(vendor_transaction_status_models.GetTransactionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTransactionUnauthorized creates a GetTransactionUnauthorized with default headers values
func NewGetTransactionUnauthorized() *GetTransactionUnauthorized {
	return &GetTransactionUnauthorized{}
}

/*
GetTransactionUnauthorized describes a response with status code 401, with default header values.

The request's authorization header is not formatted correctly or does not contain a valid token.
*/
type GetTransactionUnauthorized struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_transaction_status_models.GetTransactionResponse
}

// IsSuccess returns true when this get transaction unauthorized response has a 2xx status code
func (o *GetTransactionUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get transaction unauthorized response has a 3xx status code
func (o *GetTransactionUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get transaction unauthorized response has a 4xx status code
func (o *GetTransactionUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get transaction unauthorized response has a 5xx status code
func (o *GetTransactionUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get transaction unauthorized response a status code equal to that given
func (o *GetTransactionUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetTransactionUnauthorized) Error() string {
	return fmt.Sprintf("[GET /vendor/transactions/v1/transactions/{transactionId}][%d] getTransactionUnauthorized  %+v", 401, o.Payload)
}

func (o *GetTransactionUnauthorized) String() string {
	return fmt.Sprintf("[GET /vendor/transactions/v1/transactions/{transactionId}][%d] getTransactionUnauthorized  %+v", 401, o.Payload)
}

func (o *GetTransactionUnauthorized) GetPayload() *vendor_transaction_status_models.GetTransactionResponse {
	return o.Payload
}

func (o *GetTransactionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(vendor_transaction_status_models.GetTransactionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTransactionForbidden creates a GetTransactionForbidden with default headers values
func NewGetTransactionForbidden() *GetTransactionForbidden {
	return &GetTransactionForbidden{}
}

/*
GetTransactionForbidden describes a response with status code 403, with default header values.

Indicates that access to the resource is forbidden. Possible reasons include Access Denied, Unauthorized, Expired Token, or Invalid Signature.
*/
type GetTransactionForbidden struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_transaction_status_models.GetTransactionResponse
}

// IsSuccess returns true when this get transaction forbidden response has a 2xx status code
func (o *GetTransactionForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get transaction forbidden response has a 3xx status code
func (o *GetTransactionForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get transaction forbidden response has a 4xx status code
func (o *GetTransactionForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get transaction forbidden response has a 5xx status code
func (o *GetTransactionForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get transaction forbidden response a status code equal to that given
func (o *GetTransactionForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetTransactionForbidden) Error() string {
	return fmt.Sprintf("[GET /vendor/transactions/v1/transactions/{transactionId}][%d] getTransactionForbidden  %+v", 403, o.Payload)
}

func (o *GetTransactionForbidden) String() string {
	return fmt.Sprintf("[GET /vendor/transactions/v1/transactions/{transactionId}][%d] getTransactionForbidden  %+v", 403, o.Payload)
}

func (o *GetTransactionForbidden) GetPayload() *vendor_transaction_status_models.GetTransactionResponse {
	return o.Payload
}

func (o *GetTransactionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(vendor_transaction_status_models.GetTransactionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTransactionNotFound creates a GetTransactionNotFound with default headers values
func NewGetTransactionNotFound() *GetTransactionNotFound {
	return &GetTransactionNotFound{}
}

/*
GetTransactionNotFound describes a response with status code 404, with default header values.

The specified resource does not exist.
*/
type GetTransactionNotFound struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_transaction_status_models.GetTransactionResponse
}

// IsSuccess returns true when this get transaction not found response has a 2xx status code
func (o *GetTransactionNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get transaction not found response has a 3xx status code
func (o *GetTransactionNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get transaction not found response has a 4xx status code
func (o *GetTransactionNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get transaction not found response has a 5xx status code
func (o *GetTransactionNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get transaction not found response a status code equal to that given
func (o *GetTransactionNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetTransactionNotFound) Error() string {
	return fmt.Sprintf("[GET /vendor/transactions/v1/transactions/{transactionId}][%d] getTransactionNotFound  %+v", 404, o.Payload)
}

func (o *GetTransactionNotFound) String() string {
	return fmt.Sprintf("[GET /vendor/transactions/v1/transactions/{transactionId}][%d] getTransactionNotFound  %+v", 404, o.Payload)
}

func (o *GetTransactionNotFound) GetPayload() *vendor_transaction_status_models.GetTransactionResponse {
	return o.Payload
}

func (o *GetTransactionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(vendor_transaction_status_models.GetTransactionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTransactionUnsupportedMediaType creates a GetTransactionUnsupportedMediaType with default headers values
func NewGetTransactionUnsupportedMediaType() *GetTransactionUnsupportedMediaType {
	return &GetTransactionUnsupportedMediaType{}
}

/*
GetTransactionUnsupportedMediaType describes a response with status code 415, with default header values.

The request payload is in an unsupported format.
*/
type GetTransactionUnsupportedMediaType struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_transaction_status_models.GetTransactionResponse
}

// IsSuccess returns true when this get transaction unsupported media type response has a 2xx status code
func (o *GetTransactionUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get transaction unsupported media type response has a 3xx status code
func (o *GetTransactionUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get transaction unsupported media type response has a 4xx status code
func (o *GetTransactionUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get transaction unsupported media type response has a 5xx status code
func (o *GetTransactionUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get transaction unsupported media type response a status code equal to that given
func (o *GetTransactionUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetTransactionUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /vendor/transactions/v1/transactions/{transactionId}][%d] getTransactionUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetTransactionUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /vendor/transactions/v1/transactions/{transactionId}][%d] getTransactionUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetTransactionUnsupportedMediaType) GetPayload() *vendor_transaction_status_models.GetTransactionResponse {
	return o.Payload
}

func (o *GetTransactionUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(vendor_transaction_status_models.GetTransactionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTransactionTooManyRequests creates a GetTransactionTooManyRequests with default headers values
func NewGetTransactionTooManyRequests() *GetTransactionTooManyRequests {
	return &GetTransactionTooManyRequests{}
}

/*
GetTransactionTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type GetTransactionTooManyRequests struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_transaction_status_models.GetTransactionResponse
}

// IsSuccess returns true when this get transaction too many requests response has a 2xx status code
func (o *GetTransactionTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get transaction too many requests response has a 3xx status code
func (o *GetTransactionTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get transaction too many requests response has a 4xx status code
func (o *GetTransactionTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get transaction too many requests response has a 5xx status code
func (o *GetTransactionTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get transaction too many requests response a status code equal to that given
func (o *GetTransactionTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetTransactionTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /vendor/transactions/v1/transactions/{transactionId}][%d] getTransactionTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetTransactionTooManyRequests) String() string {
	return fmt.Sprintf("[GET /vendor/transactions/v1/transactions/{transactionId}][%d] getTransactionTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetTransactionTooManyRequests) GetPayload() *vendor_transaction_status_models.GetTransactionResponse {
	return o.Payload
}

func (o *GetTransactionTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(vendor_transaction_status_models.GetTransactionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTransactionInternalServerError creates a GetTransactionInternalServerError with default headers values
func NewGetTransactionInternalServerError() *GetTransactionInternalServerError {
	return &GetTransactionInternalServerError{}
}

/*
GetTransactionInternalServerError describes a response with status code 500, with default header values.

An unexpected condition occurred that prevented the server from fulfilling the request.
*/
type GetTransactionInternalServerError struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_transaction_status_models.GetTransactionResponse
}

// IsSuccess returns true when this get transaction internal server error response has a 2xx status code
func (o *GetTransactionInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get transaction internal server error response has a 3xx status code
func (o *GetTransactionInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get transaction internal server error response has a 4xx status code
func (o *GetTransactionInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get transaction internal server error response has a 5xx status code
func (o *GetTransactionInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get transaction internal server error response a status code equal to that given
func (o *GetTransactionInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetTransactionInternalServerError) Error() string {
	return fmt.Sprintf("[GET /vendor/transactions/v1/transactions/{transactionId}][%d] getTransactionInternalServerError  %+v", 500, o.Payload)
}

func (o *GetTransactionInternalServerError) String() string {
	return fmt.Sprintf("[GET /vendor/transactions/v1/transactions/{transactionId}][%d] getTransactionInternalServerError  %+v", 500, o.Payload)
}

func (o *GetTransactionInternalServerError) GetPayload() *vendor_transaction_status_models.GetTransactionResponse {
	return o.Payload
}

func (o *GetTransactionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(vendor_transaction_status_models.GetTransactionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTransactionServiceUnavailable creates a GetTransactionServiceUnavailable with default headers values
func NewGetTransactionServiceUnavailable() *GetTransactionServiceUnavailable {
	return &GetTransactionServiceUnavailable{}
}

/*
GetTransactionServiceUnavailable describes a response with status code 503, with default header values.

Temporary overloading or maintenance of the server.
*/
type GetTransactionServiceUnavailable struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_transaction_status_models.GetTransactionResponse
}

// IsSuccess returns true when this get transaction service unavailable response has a 2xx status code
func (o *GetTransactionServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get transaction service unavailable response has a 3xx status code
func (o *GetTransactionServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get transaction service unavailable response has a 4xx status code
func (o *GetTransactionServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get transaction service unavailable response has a 5xx status code
func (o *GetTransactionServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get transaction service unavailable response a status code equal to that given
func (o *GetTransactionServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetTransactionServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /vendor/transactions/v1/transactions/{transactionId}][%d] getTransactionServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetTransactionServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /vendor/transactions/v1/transactions/{transactionId}][%d] getTransactionServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetTransactionServiceUnavailable) GetPayload() *vendor_transaction_status_models.GetTransactionResponse {
	return o.Payload
}

func (o *GetTransactionServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(vendor_transaction_status_models.GetTransactionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
