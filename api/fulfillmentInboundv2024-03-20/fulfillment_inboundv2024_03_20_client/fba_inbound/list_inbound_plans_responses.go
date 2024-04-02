// Code generated by go-swagger; DO NOT EDIT.

package fba_inbound

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/renabled/amzn-sp-api-go/api/fulfillmentInboundv2024-03-20/fulfillment_inboundv2024_03_20_models"
)

// ListInboundPlansReader is a Reader for the ListInboundPlans structure.
type ListInboundPlansReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListInboundPlansReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListInboundPlansOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListInboundPlansBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListInboundPlansForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListInboundPlansNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewListInboundPlansRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewListInboundPlansUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewListInboundPlansTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListInboundPlansInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewListInboundPlansServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListInboundPlansOK creates a ListInboundPlansOK with default headers values
func NewListInboundPlansOK() *ListInboundPlansOK {
	return &ListInboundPlansOK{}
}

/*
ListInboundPlansOK describes a response with status code 200, with default header values.

ListInboundPlans 200 response
*/
type ListInboundPlansOK struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inboundv2024_03_20_models.ListInboundPlansResponse
}

// IsSuccess returns true when this list inbound plans o k response has a 2xx status code
func (o *ListInboundPlansOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list inbound plans o k response has a 3xx status code
func (o *ListInboundPlansOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list inbound plans o k response has a 4xx status code
func (o *ListInboundPlansOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list inbound plans o k response has a 5xx status code
func (o *ListInboundPlansOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list inbound plans o k response a status code equal to that given
func (o *ListInboundPlansOK) IsCode(code int) bool {
	return code == 200
}

func (o *ListInboundPlansOK) Error() string {
	return fmt.Sprintf("[GET /inbound/fba/2024-03-20/inboundPlans][%d] listInboundPlansOK  %+v", 200, o.Payload)
}

func (o *ListInboundPlansOK) String() string {
	return fmt.Sprintf("[GET /inbound/fba/2024-03-20/inboundPlans][%d] listInboundPlansOK  %+v", 200, o.Payload)
}

func (o *ListInboundPlansOK) GetPayload() *fulfillment_inboundv2024_03_20_models.ListInboundPlansResponse {
	return o.Payload
}

func (o *ListInboundPlansOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fulfillment_inboundv2024_03_20_models.ListInboundPlansResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListInboundPlansBadRequest creates a ListInboundPlansBadRequest with default headers values
func NewListInboundPlansBadRequest() *ListInboundPlansBadRequest {
	return &ListInboundPlansBadRequest{}
}

/*
ListInboundPlansBadRequest describes a response with status code 400, with default header values.

Request has missing or invalid parameters and cannot be parsed.
*/
type ListInboundPlansBadRequest struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inboundv2024_03_20_models.ErrorList
}

// IsSuccess returns true when this list inbound plans bad request response has a 2xx status code
func (o *ListInboundPlansBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list inbound plans bad request response has a 3xx status code
func (o *ListInboundPlansBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list inbound plans bad request response has a 4xx status code
func (o *ListInboundPlansBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this list inbound plans bad request response has a 5xx status code
func (o *ListInboundPlansBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this list inbound plans bad request response a status code equal to that given
func (o *ListInboundPlansBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *ListInboundPlansBadRequest) Error() string {
	return fmt.Sprintf("[GET /inbound/fba/2024-03-20/inboundPlans][%d] listInboundPlansBadRequest  %+v", 400, o.Payload)
}

func (o *ListInboundPlansBadRequest) String() string {
	return fmt.Sprintf("[GET /inbound/fba/2024-03-20/inboundPlans][%d] listInboundPlansBadRequest  %+v", 400, o.Payload)
}

func (o *ListInboundPlansBadRequest) GetPayload() *fulfillment_inboundv2024_03_20_models.ErrorList {
	return o.Payload
}

func (o *ListInboundPlansBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fulfillment_inboundv2024_03_20_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListInboundPlansForbidden creates a ListInboundPlansForbidden with default headers values
func NewListInboundPlansForbidden() *ListInboundPlansForbidden {
	return &ListInboundPlansForbidden{}
}

/*
ListInboundPlansForbidden describes a response with status code 403, with default header values.

Indicates that access to the resource is forbidden. Possible reasons include Access Denied, Unauthorized, Expired Token, or Invalid Signature.
*/
type ListInboundPlansForbidden struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inboundv2024_03_20_models.ErrorList
}

// IsSuccess returns true when this list inbound plans forbidden response has a 2xx status code
func (o *ListInboundPlansForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list inbound plans forbidden response has a 3xx status code
func (o *ListInboundPlansForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list inbound plans forbidden response has a 4xx status code
func (o *ListInboundPlansForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this list inbound plans forbidden response has a 5xx status code
func (o *ListInboundPlansForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this list inbound plans forbidden response a status code equal to that given
func (o *ListInboundPlansForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ListInboundPlansForbidden) Error() string {
	return fmt.Sprintf("[GET /inbound/fba/2024-03-20/inboundPlans][%d] listInboundPlansForbidden  %+v", 403, o.Payload)
}

func (o *ListInboundPlansForbidden) String() string {
	return fmt.Sprintf("[GET /inbound/fba/2024-03-20/inboundPlans][%d] listInboundPlansForbidden  %+v", 403, o.Payload)
}

func (o *ListInboundPlansForbidden) GetPayload() *fulfillment_inboundv2024_03_20_models.ErrorList {
	return o.Payload
}

func (o *ListInboundPlansForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(fulfillment_inboundv2024_03_20_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListInboundPlansNotFound creates a ListInboundPlansNotFound with default headers values
func NewListInboundPlansNotFound() *ListInboundPlansNotFound {
	return &ListInboundPlansNotFound{}
}

/*
ListInboundPlansNotFound describes a response with status code 404, with default header values.

The resource specified does not exist.
*/
type ListInboundPlansNotFound struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inboundv2024_03_20_models.ErrorList
}

// IsSuccess returns true when this list inbound plans not found response has a 2xx status code
func (o *ListInboundPlansNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list inbound plans not found response has a 3xx status code
func (o *ListInboundPlansNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list inbound plans not found response has a 4xx status code
func (o *ListInboundPlansNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this list inbound plans not found response has a 5xx status code
func (o *ListInboundPlansNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this list inbound plans not found response a status code equal to that given
func (o *ListInboundPlansNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *ListInboundPlansNotFound) Error() string {
	return fmt.Sprintf("[GET /inbound/fba/2024-03-20/inboundPlans][%d] listInboundPlansNotFound  %+v", 404, o.Payload)
}

func (o *ListInboundPlansNotFound) String() string {
	return fmt.Sprintf("[GET /inbound/fba/2024-03-20/inboundPlans][%d] listInboundPlansNotFound  %+v", 404, o.Payload)
}

func (o *ListInboundPlansNotFound) GetPayload() *fulfillment_inboundv2024_03_20_models.ErrorList {
	return o.Payload
}

func (o *ListInboundPlansNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fulfillment_inboundv2024_03_20_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListInboundPlansRequestEntityTooLarge creates a ListInboundPlansRequestEntityTooLarge with default headers values
func NewListInboundPlansRequestEntityTooLarge() *ListInboundPlansRequestEntityTooLarge {
	return &ListInboundPlansRequestEntityTooLarge{}
}

/*
ListInboundPlansRequestEntityTooLarge describes a response with status code 413, with default header values.

The request size exceeded the maximum accepted size.
*/
type ListInboundPlansRequestEntityTooLarge struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inboundv2024_03_20_models.ErrorList
}

// IsSuccess returns true when this list inbound plans request entity too large response has a 2xx status code
func (o *ListInboundPlansRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list inbound plans request entity too large response has a 3xx status code
func (o *ListInboundPlansRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list inbound plans request entity too large response has a 4xx status code
func (o *ListInboundPlansRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this list inbound plans request entity too large response has a 5xx status code
func (o *ListInboundPlansRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this list inbound plans request entity too large response a status code equal to that given
func (o *ListInboundPlansRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *ListInboundPlansRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /inbound/fba/2024-03-20/inboundPlans][%d] listInboundPlansRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *ListInboundPlansRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /inbound/fba/2024-03-20/inboundPlans][%d] listInboundPlansRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *ListInboundPlansRequestEntityTooLarge) GetPayload() *fulfillment_inboundv2024_03_20_models.ErrorList {
	return o.Payload
}

func (o *ListInboundPlansRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(fulfillment_inboundv2024_03_20_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListInboundPlansUnsupportedMediaType creates a ListInboundPlansUnsupportedMediaType with default headers values
func NewListInboundPlansUnsupportedMediaType() *ListInboundPlansUnsupportedMediaType {
	return &ListInboundPlansUnsupportedMediaType{}
}

/*
ListInboundPlansUnsupportedMediaType describes a response with status code 415, with default header values.

The request payload is in an unsupported format.
*/
type ListInboundPlansUnsupportedMediaType struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inboundv2024_03_20_models.ErrorList
}

// IsSuccess returns true when this list inbound plans unsupported media type response has a 2xx status code
func (o *ListInboundPlansUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list inbound plans unsupported media type response has a 3xx status code
func (o *ListInboundPlansUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list inbound plans unsupported media type response has a 4xx status code
func (o *ListInboundPlansUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this list inbound plans unsupported media type response has a 5xx status code
func (o *ListInboundPlansUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this list inbound plans unsupported media type response a status code equal to that given
func (o *ListInboundPlansUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *ListInboundPlansUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /inbound/fba/2024-03-20/inboundPlans][%d] listInboundPlansUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *ListInboundPlansUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /inbound/fba/2024-03-20/inboundPlans][%d] listInboundPlansUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *ListInboundPlansUnsupportedMediaType) GetPayload() *fulfillment_inboundv2024_03_20_models.ErrorList {
	return o.Payload
}

func (o *ListInboundPlansUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(fulfillment_inboundv2024_03_20_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListInboundPlansTooManyRequests creates a ListInboundPlansTooManyRequests with default headers values
func NewListInboundPlansTooManyRequests() *ListInboundPlansTooManyRequests {
	return &ListInboundPlansTooManyRequests{}
}

/*
ListInboundPlansTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type ListInboundPlansTooManyRequests struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inboundv2024_03_20_models.ErrorList
}

// IsSuccess returns true when this list inbound plans too many requests response has a 2xx status code
func (o *ListInboundPlansTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list inbound plans too many requests response has a 3xx status code
func (o *ListInboundPlansTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list inbound plans too many requests response has a 4xx status code
func (o *ListInboundPlansTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this list inbound plans too many requests response has a 5xx status code
func (o *ListInboundPlansTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this list inbound plans too many requests response a status code equal to that given
func (o *ListInboundPlansTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *ListInboundPlansTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /inbound/fba/2024-03-20/inboundPlans][%d] listInboundPlansTooManyRequests  %+v", 429, o.Payload)
}

func (o *ListInboundPlansTooManyRequests) String() string {
	return fmt.Sprintf("[GET /inbound/fba/2024-03-20/inboundPlans][%d] listInboundPlansTooManyRequests  %+v", 429, o.Payload)
}

func (o *ListInboundPlansTooManyRequests) GetPayload() *fulfillment_inboundv2024_03_20_models.ErrorList {
	return o.Payload
}

func (o *ListInboundPlansTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(fulfillment_inboundv2024_03_20_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListInboundPlansInternalServerError creates a ListInboundPlansInternalServerError with default headers values
func NewListInboundPlansInternalServerError() *ListInboundPlansInternalServerError {
	return &ListInboundPlansInternalServerError{}
}

/*
ListInboundPlansInternalServerError describes a response with status code 500, with default header values.

An unexpected condition occurred that prevented the server from fulfilling the request.
*/
type ListInboundPlansInternalServerError struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inboundv2024_03_20_models.ErrorList
}

// IsSuccess returns true when this list inbound plans internal server error response has a 2xx status code
func (o *ListInboundPlansInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list inbound plans internal server error response has a 3xx status code
func (o *ListInboundPlansInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list inbound plans internal server error response has a 4xx status code
func (o *ListInboundPlansInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this list inbound plans internal server error response has a 5xx status code
func (o *ListInboundPlansInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this list inbound plans internal server error response a status code equal to that given
func (o *ListInboundPlansInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ListInboundPlansInternalServerError) Error() string {
	return fmt.Sprintf("[GET /inbound/fba/2024-03-20/inboundPlans][%d] listInboundPlansInternalServerError  %+v", 500, o.Payload)
}

func (o *ListInboundPlansInternalServerError) String() string {
	return fmt.Sprintf("[GET /inbound/fba/2024-03-20/inboundPlans][%d] listInboundPlansInternalServerError  %+v", 500, o.Payload)
}

func (o *ListInboundPlansInternalServerError) GetPayload() *fulfillment_inboundv2024_03_20_models.ErrorList {
	return o.Payload
}

func (o *ListInboundPlansInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(fulfillment_inboundv2024_03_20_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListInboundPlansServiceUnavailable creates a ListInboundPlansServiceUnavailable with default headers values
func NewListInboundPlansServiceUnavailable() *ListInboundPlansServiceUnavailable {
	return &ListInboundPlansServiceUnavailable{}
}

/*
ListInboundPlansServiceUnavailable describes a response with status code 503, with default header values.

Temporary overloading or maintenance of the server.
*/
type ListInboundPlansServiceUnavailable struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inboundv2024_03_20_models.ErrorList
}

// IsSuccess returns true when this list inbound plans service unavailable response has a 2xx status code
func (o *ListInboundPlansServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list inbound plans service unavailable response has a 3xx status code
func (o *ListInboundPlansServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list inbound plans service unavailable response has a 4xx status code
func (o *ListInboundPlansServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this list inbound plans service unavailable response has a 5xx status code
func (o *ListInboundPlansServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this list inbound plans service unavailable response a status code equal to that given
func (o *ListInboundPlansServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *ListInboundPlansServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /inbound/fba/2024-03-20/inboundPlans][%d] listInboundPlansServiceUnavailable  %+v", 503, o.Payload)
}

func (o *ListInboundPlansServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /inbound/fba/2024-03-20/inboundPlans][%d] listInboundPlansServiceUnavailable  %+v", 503, o.Payload)
}

func (o *ListInboundPlansServiceUnavailable) GetPayload() *fulfillment_inboundv2024_03_20_models.ErrorList {
	return o.Payload
}

func (o *ListInboundPlansServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(fulfillment_inboundv2024_03_20_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
