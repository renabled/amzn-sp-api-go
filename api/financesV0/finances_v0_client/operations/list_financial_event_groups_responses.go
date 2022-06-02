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

// ListFinancialEventGroupsReader is a Reader for the ListFinancialEventGroups structure.
type ListFinancialEventGroupsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListFinancialEventGroupsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListFinancialEventGroupsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListFinancialEventGroupsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListFinancialEventGroupsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListFinancialEventGroupsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewListFinancialEventGroupsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListFinancialEventGroupsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewListFinancialEventGroupsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListFinancialEventGroupsOK creates a ListFinancialEventGroupsOK with default headers values
func NewListFinancialEventGroupsOK() *ListFinancialEventGroupsOK {
	return &ListFinancialEventGroupsOK{}
}

/* ListFinancialEventGroupsOK describes a response with status code 200, with default header values.

Success.
*/
type ListFinancialEventGroupsOK struct {
	Payload *finances_v0_models.ListFinancialEventGroupsResponse
}

func (o *ListFinancialEventGroupsOK) Error() string {
	return fmt.Sprintf("[GET /finances/v0/financialEventGroups][%d] listFinancialEventGroupsOK  %+v", 200, o.Payload)
}
func (o *ListFinancialEventGroupsOK) GetPayload() *finances_v0_models.ListFinancialEventGroupsResponse {
	return o.Payload
}

func (o *ListFinancialEventGroupsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(finances_v0_models.ListFinancialEventGroupsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListFinancialEventGroupsBadRequest creates a ListFinancialEventGroupsBadRequest with default headers values
func NewListFinancialEventGroupsBadRequest() *ListFinancialEventGroupsBadRequest {
	return &ListFinancialEventGroupsBadRequest{}
}

/* ListFinancialEventGroupsBadRequest describes a response with status code 400, with default header values.

Request has missing or invalid parameters and cannot be parsed.
*/
type ListFinancialEventGroupsBadRequest struct {
	Payload *finances_v0_models.ListFinancialEventGroupsResponse
}

func (o *ListFinancialEventGroupsBadRequest) Error() string {
	return fmt.Sprintf("[GET /finances/v0/financialEventGroups][%d] listFinancialEventGroupsBadRequest  %+v", 400, o.Payload)
}
func (o *ListFinancialEventGroupsBadRequest) GetPayload() *finances_v0_models.ListFinancialEventGroupsResponse {
	return o.Payload
}

func (o *ListFinancialEventGroupsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(finances_v0_models.ListFinancialEventGroupsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListFinancialEventGroupsForbidden creates a ListFinancialEventGroupsForbidden with default headers values
func NewListFinancialEventGroupsForbidden() *ListFinancialEventGroupsForbidden {
	return &ListFinancialEventGroupsForbidden{}
}

/* ListFinancialEventGroupsForbidden describes a response with status code 403, with default header values.

Indicates that access to the resource is forbidden. Possible reasons include Access Denied, Unauthorized, Expired Token, or Invalid Signature.
*/
type ListFinancialEventGroupsForbidden struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *finances_v0_models.ListFinancialEventGroupsResponse
}

func (o *ListFinancialEventGroupsForbidden) Error() string {
	return fmt.Sprintf("[GET /finances/v0/financialEventGroups][%d] listFinancialEventGroupsForbidden  %+v", 403, o.Payload)
}
func (o *ListFinancialEventGroupsForbidden) GetPayload() *finances_v0_models.ListFinancialEventGroupsResponse {
	return o.Payload
}

func (o *ListFinancialEventGroupsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(finances_v0_models.ListFinancialEventGroupsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListFinancialEventGroupsNotFound creates a ListFinancialEventGroupsNotFound with default headers values
func NewListFinancialEventGroupsNotFound() *ListFinancialEventGroupsNotFound {
	return &ListFinancialEventGroupsNotFound{}
}

/* ListFinancialEventGroupsNotFound describes a response with status code 404, with default header values.

The resource specified does not exist.
*/
type ListFinancialEventGroupsNotFound struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *finances_v0_models.ListFinancialEventGroupsResponse
}

func (o *ListFinancialEventGroupsNotFound) Error() string {
	return fmt.Sprintf("[GET /finances/v0/financialEventGroups][%d] listFinancialEventGroupsNotFound  %+v", 404, o.Payload)
}
func (o *ListFinancialEventGroupsNotFound) GetPayload() *finances_v0_models.ListFinancialEventGroupsResponse {
	return o.Payload
}

func (o *ListFinancialEventGroupsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(finances_v0_models.ListFinancialEventGroupsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListFinancialEventGroupsTooManyRequests creates a ListFinancialEventGroupsTooManyRequests with default headers values
func NewListFinancialEventGroupsTooManyRequests() *ListFinancialEventGroupsTooManyRequests {
	return &ListFinancialEventGroupsTooManyRequests{}
}

/* ListFinancialEventGroupsTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type ListFinancialEventGroupsTooManyRequests struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *finances_v0_models.ListFinancialEventGroupsResponse
}

func (o *ListFinancialEventGroupsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /finances/v0/financialEventGroups][%d] listFinancialEventGroupsTooManyRequests  %+v", 429, o.Payload)
}
func (o *ListFinancialEventGroupsTooManyRequests) GetPayload() *finances_v0_models.ListFinancialEventGroupsResponse {
	return o.Payload
}

func (o *ListFinancialEventGroupsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(finances_v0_models.ListFinancialEventGroupsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListFinancialEventGroupsInternalServerError creates a ListFinancialEventGroupsInternalServerError with default headers values
func NewListFinancialEventGroupsInternalServerError() *ListFinancialEventGroupsInternalServerError {
	return &ListFinancialEventGroupsInternalServerError{}
}

/* ListFinancialEventGroupsInternalServerError describes a response with status code 500, with default header values.

An unexpected condition occurred that prevented the server from fulfilling the request.
*/
type ListFinancialEventGroupsInternalServerError struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *finances_v0_models.ListFinancialEventGroupsResponse
}

func (o *ListFinancialEventGroupsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /finances/v0/financialEventGroups][%d] listFinancialEventGroupsInternalServerError  %+v", 500, o.Payload)
}
func (o *ListFinancialEventGroupsInternalServerError) GetPayload() *finances_v0_models.ListFinancialEventGroupsResponse {
	return o.Payload
}

func (o *ListFinancialEventGroupsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(finances_v0_models.ListFinancialEventGroupsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListFinancialEventGroupsServiceUnavailable creates a ListFinancialEventGroupsServiceUnavailable with default headers values
func NewListFinancialEventGroupsServiceUnavailable() *ListFinancialEventGroupsServiceUnavailable {
	return &ListFinancialEventGroupsServiceUnavailable{}
}

/* ListFinancialEventGroupsServiceUnavailable describes a response with status code 503, with default header values.

Temporary overloading or maintenance of the server.
*/
type ListFinancialEventGroupsServiceUnavailable struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *finances_v0_models.ListFinancialEventGroupsResponse
}

func (o *ListFinancialEventGroupsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /finances/v0/financialEventGroups][%d] listFinancialEventGroupsServiceUnavailable  %+v", 503, o.Payload)
}
func (o *ListFinancialEventGroupsServiceUnavailable) GetPayload() *finances_v0_models.ListFinancialEventGroupsResponse {
	return o.Payload
}

func (o *ListFinancialEventGroupsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(finances_v0_models.ListFinancialEventGroupsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
