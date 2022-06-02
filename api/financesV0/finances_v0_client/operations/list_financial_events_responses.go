// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/xamandar/amzn-sp-api-go/api/financesV0/finances_v0_models"
)

// ListFinancialEventsReader is a Reader for the ListFinancialEvents structure.
type ListFinancialEventsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListFinancialEventsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListFinancialEventsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListFinancialEventsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListFinancialEventsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListFinancialEventsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewListFinancialEventsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListFinancialEventsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewListFinancialEventsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListFinancialEventsOK creates a ListFinancialEventsOK with default headers values
func NewListFinancialEventsOK() *ListFinancialEventsOK {
	return &ListFinancialEventsOK{}
}

/* ListFinancialEventsOK describes a response with status code 200, with default header values.

Success.
*/
type ListFinancialEventsOK struct {
	Payload *finances_v0_models.ListFinancialEventsResponse
}

func (o *ListFinancialEventsOK) Error() string {
	return fmt.Sprintf("[GET /finances/v0/financialEvents][%d] listFinancialEventsOK  %+v", 200, o.Payload)
}
func (o *ListFinancialEventsOK) GetPayload() *finances_v0_models.ListFinancialEventsResponse {
	return o.Payload
}

func (o *ListFinancialEventsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(finances_v0_models.ListFinancialEventsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListFinancialEventsBadRequest creates a ListFinancialEventsBadRequest with default headers values
func NewListFinancialEventsBadRequest() *ListFinancialEventsBadRequest {
	return &ListFinancialEventsBadRequest{}
}

/* ListFinancialEventsBadRequest describes a response with status code 400, with default header values.

Request has missing or invalid parameters and cannot be parsed.
*/
type ListFinancialEventsBadRequest struct {
	Payload *finances_v0_models.ListFinancialEventsResponse
}

func (o *ListFinancialEventsBadRequest) Error() string {
	return fmt.Sprintf("[GET /finances/v0/financialEvents][%d] listFinancialEventsBadRequest  %+v", 400, o.Payload)
}
func (o *ListFinancialEventsBadRequest) GetPayload() *finances_v0_models.ListFinancialEventsResponse {
	return o.Payload
}

func (o *ListFinancialEventsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(finances_v0_models.ListFinancialEventsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListFinancialEventsForbidden creates a ListFinancialEventsForbidden with default headers values
func NewListFinancialEventsForbidden() *ListFinancialEventsForbidden {
	return &ListFinancialEventsForbidden{}
}

/* ListFinancialEventsForbidden describes a response with status code 403, with default header values.

Indicates that access to the resource is forbidden. Possible reasons include Access Denied, Unauthorized, Expired Token, or Invalid Signature.
*/
type ListFinancialEventsForbidden struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *finances_v0_models.ListFinancialEventsResponse
}

func (o *ListFinancialEventsForbidden) Error() string {
	return fmt.Sprintf("[GET /finances/v0/financialEvents][%d] listFinancialEventsForbidden  %+v", 403, o.Payload)
}
func (o *ListFinancialEventsForbidden) GetPayload() *finances_v0_models.ListFinancialEventsResponse {
	return o.Payload
}

func (o *ListFinancialEventsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewListFinancialEventsNotFound creates a ListFinancialEventsNotFound with default headers values
func NewListFinancialEventsNotFound() *ListFinancialEventsNotFound {
	return &ListFinancialEventsNotFound{}
}

/* ListFinancialEventsNotFound describes a response with status code 404, with default header values.

The resource specified does not exist.
*/
type ListFinancialEventsNotFound struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *finances_v0_models.ListFinancialEventsResponse
}

func (o *ListFinancialEventsNotFound) Error() string {
	return fmt.Sprintf("[GET /finances/v0/financialEvents][%d] listFinancialEventsNotFound  %+v", 404, o.Payload)
}
func (o *ListFinancialEventsNotFound) GetPayload() *finances_v0_models.ListFinancialEventsResponse {
	return o.Payload
}

func (o *ListFinancialEventsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewListFinancialEventsTooManyRequests creates a ListFinancialEventsTooManyRequests with default headers values
func NewListFinancialEventsTooManyRequests() *ListFinancialEventsTooManyRequests {
	return &ListFinancialEventsTooManyRequests{}
}

/* ListFinancialEventsTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type ListFinancialEventsTooManyRequests struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *finances_v0_models.ListFinancialEventsResponse
}

func (o *ListFinancialEventsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /finances/v0/financialEvents][%d] listFinancialEventsTooManyRequests  %+v", 429, o.Payload)
}
func (o *ListFinancialEventsTooManyRequests) GetPayload() *finances_v0_models.ListFinancialEventsResponse {
	return o.Payload
}

func (o *ListFinancialEventsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewListFinancialEventsInternalServerError creates a ListFinancialEventsInternalServerError with default headers values
func NewListFinancialEventsInternalServerError() *ListFinancialEventsInternalServerError {
	return &ListFinancialEventsInternalServerError{}
}

/* ListFinancialEventsInternalServerError describes a response with status code 500, with default header values.

An unexpected condition occurred that prevented the server from fulfilling the request.
*/
type ListFinancialEventsInternalServerError struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *finances_v0_models.ListFinancialEventsResponse
}

func (o *ListFinancialEventsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /finances/v0/financialEvents][%d] listFinancialEventsInternalServerError  %+v", 500, o.Payload)
}
func (o *ListFinancialEventsInternalServerError) GetPayload() *finances_v0_models.ListFinancialEventsResponse {
	return o.Payload
}

func (o *ListFinancialEventsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewListFinancialEventsServiceUnavailable creates a ListFinancialEventsServiceUnavailable with default headers values
func NewListFinancialEventsServiceUnavailable() *ListFinancialEventsServiceUnavailable {
	return &ListFinancialEventsServiceUnavailable{}
}

/* ListFinancialEventsServiceUnavailable describes a response with status code 503, with default header values.

Temporary overloading or maintenance of the server.
*/
type ListFinancialEventsServiceUnavailable struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *finances_v0_models.ListFinancialEventsResponse
}

func (o *ListFinancialEventsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /finances/v0/financialEvents][%d] listFinancialEventsServiceUnavailable  %+v", 503, o.Payload)
}
func (o *ListFinancialEventsServiceUnavailable) GetPayload() *finances_v0_models.ListFinancialEventsResponse {
	return o.Payload
}

func (o *ListFinancialEventsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
