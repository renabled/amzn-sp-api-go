// Code generated by go-swagger; DO NOT EDIT.

package messaging

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/renabled/amzn-sp-api-go/api/messaging/messaging_models"
)

// SendInvoiceReader is a Reader for the SendInvoice structure.
type SendInvoiceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SendInvoiceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewSendInvoiceCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSendInvoiceBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSendInvoiceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSendInvoiceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewSendInvoiceRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewSendInvoiceUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewSendInvoiceTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSendInvoiceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewSendInvoiceServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSendInvoiceCreated creates a SendInvoiceCreated with default headers values
func NewSendInvoiceCreated() *SendInvoiceCreated {
	return &SendInvoiceCreated{}
}

/*
SendInvoiceCreated describes a response with status code 201, with default header values.

The message was created for the order.
*/
type SendInvoiceCreated struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference id.
	 */
	XAmznRequestID string

	Payload *messaging_models.InvoiceResponse
}

// IsSuccess returns true when this send invoice created response has a 2xx status code
func (o *SendInvoiceCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this send invoice created response has a 3xx status code
func (o *SendInvoiceCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this send invoice created response has a 4xx status code
func (o *SendInvoiceCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this send invoice created response has a 5xx status code
func (o *SendInvoiceCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this send invoice created response a status code equal to that given
func (o *SendInvoiceCreated) IsCode(code int) bool {
	return code == 201
}

func (o *SendInvoiceCreated) Error() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/invoice][%d] sendInvoiceCreated  %+v", 201, o.Payload)
}

func (o *SendInvoiceCreated) String() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/invoice][%d] sendInvoiceCreated  %+v", 201, o.Payload)
}

func (o *SendInvoiceCreated) GetPayload() *messaging_models.InvoiceResponse {
	return o.Payload
}

func (o *SendInvoiceCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(messaging_models.InvoiceResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSendInvoiceBadRequest creates a SendInvoiceBadRequest with default headers values
func NewSendInvoiceBadRequest() *SendInvoiceBadRequest {
	return &SendInvoiceBadRequest{}
}

/*
SendInvoiceBadRequest describes a response with status code 400, with default header values.

Request has missing or invalid parameters and cannot be parsed.
*/
type SendInvoiceBadRequest struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference id.
	 */
	XAmznRequestID string

	Payload *messaging_models.InvoiceResponse
}

// IsSuccess returns true when this send invoice bad request response has a 2xx status code
func (o *SendInvoiceBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this send invoice bad request response has a 3xx status code
func (o *SendInvoiceBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this send invoice bad request response has a 4xx status code
func (o *SendInvoiceBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this send invoice bad request response has a 5xx status code
func (o *SendInvoiceBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this send invoice bad request response a status code equal to that given
func (o *SendInvoiceBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *SendInvoiceBadRequest) Error() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/invoice][%d] sendInvoiceBadRequest  %+v", 400, o.Payload)
}

func (o *SendInvoiceBadRequest) String() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/invoice][%d] sendInvoiceBadRequest  %+v", 400, o.Payload)
}

func (o *SendInvoiceBadRequest) GetPayload() *messaging_models.InvoiceResponse {
	return o.Payload
}

func (o *SendInvoiceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(messaging_models.InvoiceResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSendInvoiceForbidden creates a SendInvoiceForbidden with default headers values
func NewSendInvoiceForbidden() *SendInvoiceForbidden {
	return &SendInvoiceForbidden{}
}

/*
SendInvoiceForbidden describes a response with status code 403, with default header values.

403 can be caused for reasons like Access Denied, Unauthorized, Expired Token, Invalid Signature or Resource Not Found.
*/
type SendInvoiceForbidden struct {

	/* Unique request reference id.
	 */
	XAmznRequestID string

	Payload *messaging_models.InvoiceResponse
}

// IsSuccess returns true when this send invoice forbidden response has a 2xx status code
func (o *SendInvoiceForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this send invoice forbidden response has a 3xx status code
func (o *SendInvoiceForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this send invoice forbidden response has a 4xx status code
func (o *SendInvoiceForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this send invoice forbidden response has a 5xx status code
func (o *SendInvoiceForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this send invoice forbidden response a status code equal to that given
func (o *SendInvoiceForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *SendInvoiceForbidden) Error() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/invoice][%d] sendInvoiceForbidden  %+v", 403, o.Payload)
}

func (o *SendInvoiceForbidden) String() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/invoice][%d] sendInvoiceForbidden  %+v", 403, o.Payload)
}

func (o *SendInvoiceForbidden) GetPayload() *messaging_models.InvoiceResponse {
	return o.Payload
}

func (o *SendInvoiceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(messaging_models.InvoiceResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSendInvoiceNotFound creates a SendInvoiceNotFound with default headers values
func NewSendInvoiceNotFound() *SendInvoiceNotFound {
	return &SendInvoiceNotFound{}
}

/*
SendInvoiceNotFound describes a response with status code 404, with default header values.

The resource specified does not exist.
*/
type SendInvoiceNotFound struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference id.
	 */
	XAmznRequestID string

	Payload *messaging_models.InvoiceResponse
}

// IsSuccess returns true when this send invoice not found response has a 2xx status code
func (o *SendInvoiceNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this send invoice not found response has a 3xx status code
func (o *SendInvoiceNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this send invoice not found response has a 4xx status code
func (o *SendInvoiceNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this send invoice not found response has a 5xx status code
func (o *SendInvoiceNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this send invoice not found response a status code equal to that given
func (o *SendInvoiceNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *SendInvoiceNotFound) Error() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/invoice][%d] sendInvoiceNotFound  %+v", 404, o.Payload)
}

func (o *SendInvoiceNotFound) String() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/invoice][%d] sendInvoiceNotFound  %+v", 404, o.Payload)
}

func (o *SendInvoiceNotFound) GetPayload() *messaging_models.InvoiceResponse {
	return o.Payload
}

func (o *SendInvoiceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(messaging_models.InvoiceResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSendInvoiceRequestEntityTooLarge creates a SendInvoiceRequestEntityTooLarge with default headers values
func NewSendInvoiceRequestEntityTooLarge() *SendInvoiceRequestEntityTooLarge {
	return &SendInvoiceRequestEntityTooLarge{}
}

/*
SendInvoiceRequestEntityTooLarge describes a response with status code 413, with default header values.

The request size exceeded the maximum accepted size.
*/
type SendInvoiceRequestEntityTooLarge struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference id.
	 */
	XAmznRequestID string

	Payload *messaging_models.InvoiceResponse
}

// IsSuccess returns true when this send invoice request entity too large response has a 2xx status code
func (o *SendInvoiceRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this send invoice request entity too large response has a 3xx status code
func (o *SendInvoiceRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this send invoice request entity too large response has a 4xx status code
func (o *SendInvoiceRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this send invoice request entity too large response has a 5xx status code
func (o *SendInvoiceRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this send invoice request entity too large response a status code equal to that given
func (o *SendInvoiceRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *SendInvoiceRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/invoice][%d] sendInvoiceRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *SendInvoiceRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/invoice][%d] sendInvoiceRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *SendInvoiceRequestEntityTooLarge) GetPayload() *messaging_models.InvoiceResponse {
	return o.Payload
}

func (o *SendInvoiceRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(messaging_models.InvoiceResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSendInvoiceUnsupportedMediaType creates a SendInvoiceUnsupportedMediaType with default headers values
func NewSendInvoiceUnsupportedMediaType() *SendInvoiceUnsupportedMediaType {
	return &SendInvoiceUnsupportedMediaType{}
}

/*
SendInvoiceUnsupportedMediaType describes a response with status code 415, with default header values.

The entity of the request is in a format not supported by the requested resource.
*/
type SendInvoiceUnsupportedMediaType struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference id.
	 */
	XAmznRequestID string

	Payload *messaging_models.InvoiceResponse
}

// IsSuccess returns true when this send invoice unsupported media type response has a 2xx status code
func (o *SendInvoiceUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this send invoice unsupported media type response has a 3xx status code
func (o *SendInvoiceUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this send invoice unsupported media type response has a 4xx status code
func (o *SendInvoiceUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this send invoice unsupported media type response has a 5xx status code
func (o *SendInvoiceUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this send invoice unsupported media type response a status code equal to that given
func (o *SendInvoiceUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *SendInvoiceUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/invoice][%d] sendInvoiceUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *SendInvoiceUnsupportedMediaType) String() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/invoice][%d] sendInvoiceUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *SendInvoiceUnsupportedMediaType) GetPayload() *messaging_models.InvoiceResponse {
	return o.Payload
}

func (o *SendInvoiceUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(messaging_models.InvoiceResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSendInvoiceTooManyRequests creates a SendInvoiceTooManyRequests with default headers values
func NewSendInvoiceTooManyRequests() *SendInvoiceTooManyRequests {
	return &SendInvoiceTooManyRequests{}
}

/*
SendInvoiceTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type SendInvoiceTooManyRequests struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference id.
	 */
	XAmznRequestID string

	Payload *messaging_models.InvoiceResponse
}

// IsSuccess returns true when this send invoice too many requests response has a 2xx status code
func (o *SendInvoiceTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this send invoice too many requests response has a 3xx status code
func (o *SendInvoiceTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this send invoice too many requests response has a 4xx status code
func (o *SendInvoiceTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this send invoice too many requests response has a 5xx status code
func (o *SendInvoiceTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this send invoice too many requests response a status code equal to that given
func (o *SendInvoiceTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *SendInvoiceTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/invoice][%d] sendInvoiceTooManyRequests  %+v", 429, o.Payload)
}

func (o *SendInvoiceTooManyRequests) String() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/invoice][%d] sendInvoiceTooManyRequests  %+v", 429, o.Payload)
}

func (o *SendInvoiceTooManyRequests) GetPayload() *messaging_models.InvoiceResponse {
	return o.Payload
}

func (o *SendInvoiceTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(messaging_models.InvoiceResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSendInvoiceInternalServerError creates a SendInvoiceInternalServerError with default headers values
func NewSendInvoiceInternalServerError() *SendInvoiceInternalServerError {
	return &SendInvoiceInternalServerError{}
}

/*
SendInvoiceInternalServerError describes a response with status code 500, with default header values.

Encountered an unexpected condition which prevented the server from fulfilling the request.
*/
type SendInvoiceInternalServerError struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference id.
	 */
	XAmznRequestID string

	Payload *messaging_models.InvoiceResponse
}

// IsSuccess returns true when this send invoice internal server error response has a 2xx status code
func (o *SendInvoiceInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this send invoice internal server error response has a 3xx status code
func (o *SendInvoiceInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this send invoice internal server error response has a 4xx status code
func (o *SendInvoiceInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this send invoice internal server error response has a 5xx status code
func (o *SendInvoiceInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this send invoice internal server error response a status code equal to that given
func (o *SendInvoiceInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *SendInvoiceInternalServerError) Error() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/invoice][%d] sendInvoiceInternalServerError  %+v", 500, o.Payload)
}

func (o *SendInvoiceInternalServerError) String() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/invoice][%d] sendInvoiceInternalServerError  %+v", 500, o.Payload)
}

func (o *SendInvoiceInternalServerError) GetPayload() *messaging_models.InvoiceResponse {
	return o.Payload
}

func (o *SendInvoiceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(messaging_models.InvoiceResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSendInvoiceServiceUnavailable creates a SendInvoiceServiceUnavailable with default headers values
func NewSendInvoiceServiceUnavailable() *SendInvoiceServiceUnavailable {
	return &SendInvoiceServiceUnavailable{}
}

/*
SendInvoiceServiceUnavailable describes a response with status code 503, with default header values.

Temporary overloading or maintenance of the server.
*/
type SendInvoiceServiceUnavailable struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference id.
	 */
	XAmznRequestID string

	Payload *messaging_models.InvoiceResponse
}

// IsSuccess returns true when this send invoice service unavailable response has a 2xx status code
func (o *SendInvoiceServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this send invoice service unavailable response has a 3xx status code
func (o *SendInvoiceServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this send invoice service unavailable response has a 4xx status code
func (o *SendInvoiceServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this send invoice service unavailable response has a 5xx status code
func (o *SendInvoiceServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this send invoice service unavailable response a status code equal to that given
func (o *SendInvoiceServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *SendInvoiceServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/invoice][%d] sendInvoiceServiceUnavailable  %+v", 503, o.Payload)
}

func (o *SendInvoiceServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/invoice][%d] sendInvoiceServiceUnavailable  %+v", 503, o.Payload)
}

func (o *SendInvoiceServiceUnavailable) GetPayload() *messaging_models.InvoiceResponse {
	return o.Payload
}

func (o *SendInvoiceServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(messaging_models.InvoiceResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
