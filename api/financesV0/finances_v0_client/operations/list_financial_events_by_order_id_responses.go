// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/renabled/amzn-sp-api-go/api/financesV0/finances_v0_models"
)

// ListFinancialEventsByOrderIDReader is a Reader for the ListFinancialEventsByOrderID structure.
type ListFinancialEventsByOrderIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListFinancialEventsByOrderIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListFinancialEventsByOrderIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListFinancialEventsByOrderIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListFinancialEventsByOrderIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListFinancialEventsByOrderIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewListFinancialEventsByOrderIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListFinancialEventsByOrderIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewListFinancialEventsByOrderIDServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListFinancialEventsByOrderIDOK creates a ListFinancialEventsByOrderIDOK with default headers values
func NewListFinancialEventsByOrderIDOK() *ListFinancialEventsByOrderIDOK {
	return &ListFinancialEventsByOrderIDOK{}
}

/*
ListFinancialEventsByOrderIDOK describes a response with status code 200, with default header values.

Financial Events successfully retrieved.
*/
type ListFinancialEventsByOrderIDOK struct {
	Payload *finances_v0_models.ListFinancialEventsResponse
}

// IsSuccess returns true when this list financial events by order Id o k response has a 2xx status code
func (o *ListFinancialEventsByOrderIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list financial events by order Id o k response has a 3xx status code
func (o *ListFinancialEventsByOrderIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list financial events by order Id o k response has a 4xx status code
func (o *ListFinancialEventsByOrderIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list financial events by order Id o k response has a 5xx status code
func (o *ListFinancialEventsByOrderIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list financial events by order Id o k response a status code equal to that given
func (o *ListFinancialEventsByOrderIDOK) IsCode(code int) bool {
	return code == 200
}

func (o *ListFinancialEventsByOrderIDOK) Error() string {
	return fmt.Sprintf("[GET /finances/v0/orders/{orderId}/financialEvents][%d] listFinancialEventsByOrderIdOK  %+v", 200, o.Payload)
}

func (o *ListFinancialEventsByOrderIDOK) String() string {
	return fmt.Sprintf("[GET /finances/v0/orders/{orderId}/financialEvents][%d] listFinancialEventsByOrderIdOK  %+v", 200, o.Payload)
}

func (o *ListFinancialEventsByOrderIDOK) GetPayload() *finances_v0_models.ListFinancialEventsResponse {
	return o.Payload
}

func (o *ListFinancialEventsByOrderIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(finances_v0_models.ListFinancialEventsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListFinancialEventsByOrderIDBadRequest creates a ListFinancialEventsByOrderIDBadRequest with default headers values
func NewListFinancialEventsByOrderIDBadRequest() *ListFinancialEventsByOrderIDBadRequest {
	return &ListFinancialEventsByOrderIDBadRequest{}
}

/*
ListFinancialEventsByOrderIDBadRequest describes a response with status code 400, with default header values.

Request has missing or invalid parameters and cannot be parsed.
*/
type ListFinancialEventsByOrderIDBadRequest struct {
	Payload *finances_v0_models.ListFinancialEventsResponse
}

// IsSuccess returns true when this list financial events by order Id bad request response has a 2xx status code
func (o *ListFinancialEventsByOrderIDBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list financial events by order Id bad request response has a 3xx status code
func (o *ListFinancialEventsByOrderIDBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list financial events by order Id bad request response has a 4xx status code
func (o *ListFinancialEventsByOrderIDBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this list financial events by order Id bad request response has a 5xx status code
func (o *ListFinancialEventsByOrderIDBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this list financial events by order Id bad request response a status code equal to that given
func (o *ListFinancialEventsByOrderIDBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *ListFinancialEventsByOrderIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /finances/v0/orders/{orderId}/financialEvents][%d] listFinancialEventsByOrderIdBadRequest  %+v", 400, o.Payload)
}

func (o *ListFinancialEventsByOrderIDBadRequest) String() string {
	return fmt.Sprintf("[GET /finances/v0/orders/{orderId}/financialEvents][%d] listFinancialEventsByOrderIdBadRequest  %+v", 400, o.Payload)
}

func (o *ListFinancialEventsByOrderIDBadRequest) GetPayload() *finances_v0_models.ListFinancialEventsResponse {
	return o.Payload
}

func (o *ListFinancialEventsByOrderIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(finances_v0_models.ListFinancialEventsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListFinancialEventsByOrderIDForbidden creates a ListFinancialEventsByOrderIDForbidden with default headers values
func NewListFinancialEventsByOrderIDForbidden() *ListFinancialEventsByOrderIDForbidden {
	return &ListFinancialEventsByOrderIDForbidden{}
}

/*
ListFinancialEventsByOrderIDForbidden describes a response with status code 403, with default header values.

Indicates that access to the resource is forbidden. Possible reasons include Access Denied, Unauthorized, Expired Token, or Invalid Signature.
*/
type ListFinancialEventsByOrderIDForbidden struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *finances_v0_models.ListFinancialEventsResponse
}

// IsSuccess returns true when this list financial events by order Id forbidden response has a 2xx status code
func (o *ListFinancialEventsByOrderIDForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list financial events by order Id forbidden response has a 3xx status code
func (o *ListFinancialEventsByOrderIDForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list financial events by order Id forbidden response has a 4xx status code
func (o *ListFinancialEventsByOrderIDForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this list financial events by order Id forbidden response has a 5xx status code
func (o *ListFinancialEventsByOrderIDForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this list financial events by order Id forbidden response a status code equal to that given
func (o *ListFinancialEventsByOrderIDForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ListFinancialEventsByOrderIDForbidden) Error() string {
	return fmt.Sprintf("[GET /finances/v0/orders/{orderId}/financialEvents][%d] listFinancialEventsByOrderIdForbidden  %+v", 403, o.Payload)
}

func (o *ListFinancialEventsByOrderIDForbidden) String() string {
	return fmt.Sprintf("[GET /finances/v0/orders/{orderId}/financialEvents][%d] listFinancialEventsByOrderIdForbidden  %+v", 403, o.Payload)
}

func (o *ListFinancialEventsByOrderIDForbidden) GetPayload() *finances_v0_models.ListFinancialEventsResponse {
	return o.Payload
}

func (o *ListFinancialEventsByOrderIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(finances_v0_models.ListFinancialEventsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListFinancialEventsByOrderIDNotFound creates a ListFinancialEventsByOrderIDNotFound with default headers values
func NewListFinancialEventsByOrderIDNotFound() *ListFinancialEventsByOrderIDNotFound {
	return &ListFinancialEventsByOrderIDNotFound{}
}

/*
ListFinancialEventsByOrderIDNotFound describes a response with status code 404, with default header values.

The resource specified does not exist.
*/
type ListFinancialEventsByOrderIDNotFound struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *finances_v0_models.ListFinancialEventsResponse
}

// IsSuccess returns true when this list financial events by order Id not found response has a 2xx status code
func (o *ListFinancialEventsByOrderIDNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list financial events by order Id not found response has a 3xx status code
func (o *ListFinancialEventsByOrderIDNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list financial events by order Id not found response has a 4xx status code
func (o *ListFinancialEventsByOrderIDNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this list financial events by order Id not found response has a 5xx status code
func (o *ListFinancialEventsByOrderIDNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this list financial events by order Id not found response a status code equal to that given
func (o *ListFinancialEventsByOrderIDNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *ListFinancialEventsByOrderIDNotFound) Error() string {
	return fmt.Sprintf("[GET /finances/v0/orders/{orderId}/financialEvents][%d] listFinancialEventsByOrderIdNotFound  %+v", 404, o.Payload)
}

func (o *ListFinancialEventsByOrderIDNotFound) String() string {
	return fmt.Sprintf("[GET /finances/v0/orders/{orderId}/financialEvents][%d] listFinancialEventsByOrderIdNotFound  %+v", 404, o.Payload)
}

func (o *ListFinancialEventsByOrderIDNotFound) GetPayload() *finances_v0_models.ListFinancialEventsResponse {
	return o.Payload
}

func (o *ListFinancialEventsByOrderIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(finances_v0_models.ListFinancialEventsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListFinancialEventsByOrderIDTooManyRequests creates a ListFinancialEventsByOrderIDTooManyRequests with default headers values
func NewListFinancialEventsByOrderIDTooManyRequests() *ListFinancialEventsByOrderIDTooManyRequests {
	return &ListFinancialEventsByOrderIDTooManyRequests{}
}

/*
ListFinancialEventsByOrderIDTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type ListFinancialEventsByOrderIDTooManyRequests struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *finances_v0_models.ListFinancialEventsResponse
}

// IsSuccess returns true when this list financial events by order Id too many requests response has a 2xx status code
func (o *ListFinancialEventsByOrderIDTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list financial events by order Id too many requests response has a 3xx status code
func (o *ListFinancialEventsByOrderIDTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list financial events by order Id too many requests response has a 4xx status code
func (o *ListFinancialEventsByOrderIDTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this list financial events by order Id too many requests response has a 5xx status code
func (o *ListFinancialEventsByOrderIDTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this list financial events by order Id too many requests response a status code equal to that given
func (o *ListFinancialEventsByOrderIDTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *ListFinancialEventsByOrderIDTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /finances/v0/orders/{orderId}/financialEvents][%d] listFinancialEventsByOrderIdTooManyRequests  %+v", 429, o.Payload)
}

func (o *ListFinancialEventsByOrderIDTooManyRequests) String() string {
	return fmt.Sprintf("[GET /finances/v0/orders/{orderId}/financialEvents][%d] listFinancialEventsByOrderIdTooManyRequests  %+v", 429, o.Payload)
}

func (o *ListFinancialEventsByOrderIDTooManyRequests) GetPayload() *finances_v0_models.ListFinancialEventsResponse {
	return o.Payload
}

func (o *ListFinancialEventsByOrderIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(finances_v0_models.ListFinancialEventsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListFinancialEventsByOrderIDInternalServerError creates a ListFinancialEventsByOrderIDInternalServerError with default headers values
func NewListFinancialEventsByOrderIDInternalServerError() *ListFinancialEventsByOrderIDInternalServerError {
	return &ListFinancialEventsByOrderIDInternalServerError{}
}

/*
ListFinancialEventsByOrderIDInternalServerError describes a response with status code 500, with default header values.

An unexpected condition occurred that prevented the server from fulfilling the request.
*/
type ListFinancialEventsByOrderIDInternalServerError struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *finances_v0_models.ListFinancialEventsResponse
}

// IsSuccess returns true when this list financial events by order Id internal server error response has a 2xx status code
func (o *ListFinancialEventsByOrderIDInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list financial events by order Id internal server error response has a 3xx status code
func (o *ListFinancialEventsByOrderIDInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list financial events by order Id internal server error response has a 4xx status code
func (o *ListFinancialEventsByOrderIDInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this list financial events by order Id internal server error response has a 5xx status code
func (o *ListFinancialEventsByOrderIDInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this list financial events by order Id internal server error response a status code equal to that given
func (o *ListFinancialEventsByOrderIDInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ListFinancialEventsByOrderIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /finances/v0/orders/{orderId}/financialEvents][%d] listFinancialEventsByOrderIdInternalServerError  %+v", 500, o.Payload)
}

func (o *ListFinancialEventsByOrderIDInternalServerError) String() string {
	return fmt.Sprintf("[GET /finances/v0/orders/{orderId}/financialEvents][%d] listFinancialEventsByOrderIdInternalServerError  %+v", 500, o.Payload)
}

func (o *ListFinancialEventsByOrderIDInternalServerError) GetPayload() *finances_v0_models.ListFinancialEventsResponse {
	return o.Payload
}

func (o *ListFinancialEventsByOrderIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(finances_v0_models.ListFinancialEventsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListFinancialEventsByOrderIDServiceUnavailable creates a ListFinancialEventsByOrderIDServiceUnavailable with default headers values
func NewListFinancialEventsByOrderIDServiceUnavailable() *ListFinancialEventsByOrderIDServiceUnavailable {
	return &ListFinancialEventsByOrderIDServiceUnavailable{}
}

/*
ListFinancialEventsByOrderIDServiceUnavailable describes a response with status code 503, with default header values.

Temporary overloading or maintenance of the server.
*/
type ListFinancialEventsByOrderIDServiceUnavailable struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *finances_v0_models.ListFinancialEventsResponse
}

// IsSuccess returns true when this list financial events by order Id service unavailable response has a 2xx status code
func (o *ListFinancialEventsByOrderIDServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list financial events by order Id service unavailable response has a 3xx status code
func (o *ListFinancialEventsByOrderIDServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list financial events by order Id service unavailable response has a 4xx status code
func (o *ListFinancialEventsByOrderIDServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this list financial events by order Id service unavailable response has a 5xx status code
func (o *ListFinancialEventsByOrderIDServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this list financial events by order Id service unavailable response a status code equal to that given
func (o *ListFinancialEventsByOrderIDServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *ListFinancialEventsByOrderIDServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /finances/v0/orders/{orderId}/financialEvents][%d] listFinancialEventsByOrderIdServiceUnavailable  %+v", 503, o.Payload)
}

func (o *ListFinancialEventsByOrderIDServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /finances/v0/orders/{orderId}/financialEvents][%d] listFinancialEventsByOrderIdServiceUnavailable  %+v", 503, o.Payload)
}

func (o *ListFinancialEventsByOrderIDServiceUnavailable) GetPayload() *finances_v0_models.ListFinancialEventsResponse {
	return o.Payload
}

func (o *ListFinancialEventsByOrderIDServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(finances_v0_models.ListFinancialEventsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
