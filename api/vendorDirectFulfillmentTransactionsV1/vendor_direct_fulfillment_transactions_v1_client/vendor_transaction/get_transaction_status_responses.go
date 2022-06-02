// Code generated by go-swagger; DO NOT EDIT.

package vendor_transaction

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/xamandar/amzn-sp-api-go/api/vendorDirectFulfillmentTransactionsV1/vendor_direct_fulfillment_transactions_v1_models"
)

// GetTransactionStatusReader is a Reader for the GetTransactionStatus structure.
type GetTransactionStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTransactionStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTransactionStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetTransactionStatusBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetTransactionStatusUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetTransactionStatusForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetTransactionStatusNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetTransactionStatusUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetTransactionStatusTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetTransactionStatusInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetTransactionStatusServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetTransactionStatusOK creates a GetTransactionStatusOK with default headers values
func NewGetTransactionStatusOK() *GetTransactionStatusOK {
	return &GetTransactionStatusOK{}
}

/* GetTransactionStatusOK describes a response with status code 200, with default header values.

Success.
*/
type GetTransactionStatusOK struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_transactions_v1_models.GetTransactionResponse
}

func (o *GetTransactionStatusOK) Error() string {
	return fmt.Sprintf("[GET /vendor/directFulfillment/transactions/v1/transactions/{transactionId}][%d] getTransactionStatusOK  %+v", 200, o.Payload)
}
func (o *GetTransactionStatusOK) GetPayload() *vendor_direct_fulfillment_transactions_v1_models.GetTransactionResponse {
	return o.Payload
}

func (o *GetTransactionStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(vendor_direct_fulfillment_transactions_v1_models.GetTransactionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTransactionStatusBadRequest creates a GetTransactionStatusBadRequest with default headers values
func NewGetTransactionStatusBadRequest() *GetTransactionStatusBadRequest {
	return &GetTransactionStatusBadRequest{}
}

/* GetTransactionStatusBadRequest describes a response with status code 400, with default header values.

Request has missing or invalid parameters and cannot be parsed.
*/
type GetTransactionStatusBadRequest struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_transactions_v1_models.GetTransactionResponse
}

func (o *GetTransactionStatusBadRequest) Error() string {
	return fmt.Sprintf("[GET /vendor/directFulfillment/transactions/v1/transactions/{transactionId}][%d] getTransactionStatusBadRequest  %+v", 400, o.Payload)
}
func (o *GetTransactionStatusBadRequest) GetPayload() *vendor_direct_fulfillment_transactions_v1_models.GetTransactionResponse {
	return o.Payload
}

func (o *GetTransactionStatusBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(vendor_direct_fulfillment_transactions_v1_models.GetTransactionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTransactionStatusUnauthorized creates a GetTransactionStatusUnauthorized with default headers values
func NewGetTransactionStatusUnauthorized() *GetTransactionStatusUnauthorized {
	return &GetTransactionStatusUnauthorized{}
}

/* GetTransactionStatusUnauthorized describes a response with status code 401, with default header values.

The request's Authorization header is not formatted correctly or does not contain a valid token.
*/
type GetTransactionStatusUnauthorized struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_transactions_v1_models.GetTransactionResponse
}

func (o *GetTransactionStatusUnauthorized) Error() string {
	return fmt.Sprintf("[GET /vendor/directFulfillment/transactions/v1/transactions/{transactionId}][%d] getTransactionStatusUnauthorized  %+v", 401, o.Payload)
}
func (o *GetTransactionStatusUnauthorized) GetPayload() *vendor_direct_fulfillment_transactions_v1_models.GetTransactionResponse {
	return o.Payload
}

func (o *GetTransactionStatusUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(vendor_direct_fulfillment_transactions_v1_models.GetTransactionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTransactionStatusForbidden creates a GetTransactionStatusForbidden with default headers values
func NewGetTransactionStatusForbidden() *GetTransactionStatusForbidden {
	return &GetTransactionStatusForbidden{}
}

/* GetTransactionStatusForbidden describes a response with status code 403, with default header values.

Indicates that access to the resource is forbidden. Possible reasons include Access Denied, Unauthorized, Expired Token, or Invalid Signature.
*/
type GetTransactionStatusForbidden struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_transactions_v1_models.GetTransactionResponse
}

func (o *GetTransactionStatusForbidden) Error() string {
	return fmt.Sprintf("[GET /vendor/directFulfillment/transactions/v1/transactions/{transactionId}][%d] getTransactionStatusForbidden  %+v", 403, o.Payload)
}
func (o *GetTransactionStatusForbidden) GetPayload() *vendor_direct_fulfillment_transactions_v1_models.GetTransactionResponse {
	return o.Payload
}

func (o *GetTransactionStatusForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(vendor_direct_fulfillment_transactions_v1_models.GetTransactionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTransactionStatusNotFound creates a GetTransactionStatusNotFound with default headers values
func NewGetTransactionStatusNotFound() *GetTransactionStatusNotFound {
	return &GetTransactionStatusNotFound{}
}

/* GetTransactionStatusNotFound describes a response with status code 404, with default header values.

The resource specified does not exist.
*/
type GetTransactionStatusNotFound struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_transactions_v1_models.GetTransactionResponse
}

func (o *GetTransactionStatusNotFound) Error() string {
	return fmt.Sprintf("[GET /vendor/directFulfillment/transactions/v1/transactions/{transactionId}][%d] getTransactionStatusNotFound  %+v", 404, o.Payload)
}
func (o *GetTransactionStatusNotFound) GetPayload() *vendor_direct_fulfillment_transactions_v1_models.GetTransactionResponse {
	return o.Payload
}

func (o *GetTransactionStatusNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(vendor_direct_fulfillment_transactions_v1_models.GetTransactionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTransactionStatusUnsupportedMediaType creates a GetTransactionStatusUnsupportedMediaType with default headers values
func NewGetTransactionStatusUnsupportedMediaType() *GetTransactionStatusUnsupportedMediaType {
	return &GetTransactionStatusUnsupportedMediaType{}
}

/* GetTransactionStatusUnsupportedMediaType describes a response with status code 415, with default header values.

The request payload is in an unsupported format.
*/
type GetTransactionStatusUnsupportedMediaType struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_transactions_v1_models.GetTransactionResponse
}

func (o *GetTransactionStatusUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /vendor/directFulfillment/transactions/v1/transactions/{transactionId}][%d] getTransactionStatusUnsupportedMediaType  %+v", 415, o.Payload)
}
func (o *GetTransactionStatusUnsupportedMediaType) GetPayload() *vendor_direct_fulfillment_transactions_v1_models.GetTransactionResponse {
	return o.Payload
}

func (o *GetTransactionStatusUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(vendor_direct_fulfillment_transactions_v1_models.GetTransactionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTransactionStatusTooManyRequests creates a GetTransactionStatusTooManyRequests with default headers values
func NewGetTransactionStatusTooManyRequests() *GetTransactionStatusTooManyRequests {
	return &GetTransactionStatusTooManyRequests{}
}

/* GetTransactionStatusTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type GetTransactionStatusTooManyRequests struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_transactions_v1_models.GetTransactionResponse
}

func (o *GetTransactionStatusTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /vendor/directFulfillment/transactions/v1/transactions/{transactionId}][%d] getTransactionStatusTooManyRequests  %+v", 429, o.Payload)
}
func (o *GetTransactionStatusTooManyRequests) GetPayload() *vendor_direct_fulfillment_transactions_v1_models.GetTransactionResponse {
	return o.Payload
}

func (o *GetTransactionStatusTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(vendor_direct_fulfillment_transactions_v1_models.GetTransactionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTransactionStatusInternalServerError creates a GetTransactionStatusInternalServerError with default headers values
func NewGetTransactionStatusInternalServerError() *GetTransactionStatusInternalServerError {
	return &GetTransactionStatusInternalServerError{}
}

/* GetTransactionStatusInternalServerError describes a response with status code 500, with default header values.

An unexpected condition occurred that prevented the server from fulfilling the request.
*/
type GetTransactionStatusInternalServerError struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_transactions_v1_models.GetTransactionResponse
}

func (o *GetTransactionStatusInternalServerError) Error() string {
	return fmt.Sprintf("[GET /vendor/directFulfillment/transactions/v1/transactions/{transactionId}][%d] getTransactionStatusInternalServerError  %+v", 500, o.Payload)
}
func (o *GetTransactionStatusInternalServerError) GetPayload() *vendor_direct_fulfillment_transactions_v1_models.GetTransactionResponse {
	return o.Payload
}

func (o *GetTransactionStatusInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(vendor_direct_fulfillment_transactions_v1_models.GetTransactionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTransactionStatusServiceUnavailable creates a GetTransactionStatusServiceUnavailable with default headers values
func NewGetTransactionStatusServiceUnavailable() *GetTransactionStatusServiceUnavailable {
	return &GetTransactionStatusServiceUnavailable{}
}

/* GetTransactionStatusServiceUnavailable describes a response with status code 503, with default header values.

Temporary overloading or maintenance of the server.
*/
type GetTransactionStatusServiceUnavailable struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_transactions_v1_models.GetTransactionResponse
}

func (o *GetTransactionStatusServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /vendor/directFulfillment/transactions/v1/transactions/{transactionId}][%d] getTransactionStatusServiceUnavailable  %+v", 503, o.Payload)
}
func (o *GetTransactionStatusServiceUnavailable) GetPayload() *vendor_direct_fulfillment_transactions_v1_models.GetTransactionResponse {
	return o.Payload
}

func (o *GetTransactionStatusServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(vendor_direct_fulfillment_transactions_v1_models.GetTransactionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
