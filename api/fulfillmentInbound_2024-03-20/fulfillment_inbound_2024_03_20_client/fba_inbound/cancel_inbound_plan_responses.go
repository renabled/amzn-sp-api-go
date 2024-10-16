// Code generated by go-swagger; DO NOT EDIT.

package fba_inbound

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/renabled/amzn-sp-api-go/api/fulfillmentInbound_2024-03-20/fulfillment_inbound_2024_03_20_models"
)

// CancelInboundPlanReader is a Reader for the CancelInboundPlan structure.
type CancelInboundPlanReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CancelInboundPlanReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewCancelInboundPlanAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCancelInboundPlanBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCancelInboundPlanForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCancelInboundPlanNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewCancelInboundPlanRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewCancelInboundPlanUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewCancelInboundPlanTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCancelInboundPlanInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewCancelInboundPlanServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCancelInboundPlanAccepted creates a CancelInboundPlanAccepted with default headers values
func NewCancelInboundPlanAccepted() *CancelInboundPlanAccepted {
	return &CancelInboundPlanAccepted{}
}

/*
CancelInboundPlanAccepted describes a response with status code 202, with default header values.

CancelInboundPlan 202 response
*/
type CancelInboundPlanAccepted struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inbound_2024_03_20_models.CancelInboundPlanResponse
}

// IsSuccess returns true when this cancel inbound plan accepted response has a 2xx status code
func (o *CancelInboundPlanAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cancel inbound plan accepted response has a 3xx status code
func (o *CancelInboundPlanAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cancel inbound plan accepted response has a 4xx status code
func (o *CancelInboundPlanAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this cancel inbound plan accepted response has a 5xx status code
func (o *CancelInboundPlanAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this cancel inbound plan accepted response a status code equal to that given
func (o *CancelInboundPlanAccepted) IsCode(code int) bool {
	return code == 202
}

func (o *CancelInboundPlanAccepted) Error() string {
	return fmt.Sprintf("[PUT /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/cancellation][%d] cancelInboundPlanAccepted  %+v", 202, o.Payload)
}

func (o *CancelInboundPlanAccepted) String() string {
	return fmt.Sprintf("[PUT /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/cancellation][%d] cancelInboundPlanAccepted  %+v", 202, o.Payload)
}

func (o *CancelInboundPlanAccepted) GetPayload() *fulfillment_inbound_2024_03_20_models.CancelInboundPlanResponse {
	return o.Payload
}

func (o *CancelInboundPlanAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fulfillment_inbound_2024_03_20_models.CancelInboundPlanResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCancelInboundPlanBadRequest creates a CancelInboundPlanBadRequest with default headers values
func NewCancelInboundPlanBadRequest() *CancelInboundPlanBadRequest {
	return &CancelInboundPlanBadRequest{}
}

/*
CancelInboundPlanBadRequest describes a response with status code 400, with default header values.

Request has missing or invalid parameters and cannot be parsed.
*/
type CancelInboundPlanBadRequest struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inbound_2024_03_20_models.ErrorList
}

// IsSuccess returns true when this cancel inbound plan bad request response has a 2xx status code
func (o *CancelInboundPlanBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cancel inbound plan bad request response has a 3xx status code
func (o *CancelInboundPlanBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cancel inbound plan bad request response has a 4xx status code
func (o *CancelInboundPlanBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this cancel inbound plan bad request response has a 5xx status code
func (o *CancelInboundPlanBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this cancel inbound plan bad request response a status code equal to that given
func (o *CancelInboundPlanBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *CancelInboundPlanBadRequest) Error() string {
	return fmt.Sprintf("[PUT /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/cancellation][%d] cancelInboundPlanBadRequest  %+v", 400, o.Payload)
}

func (o *CancelInboundPlanBadRequest) String() string {
	return fmt.Sprintf("[PUT /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/cancellation][%d] cancelInboundPlanBadRequest  %+v", 400, o.Payload)
}

func (o *CancelInboundPlanBadRequest) GetPayload() *fulfillment_inbound_2024_03_20_models.ErrorList {
	return o.Payload
}

func (o *CancelInboundPlanBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fulfillment_inbound_2024_03_20_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCancelInboundPlanForbidden creates a CancelInboundPlanForbidden with default headers values
func NewCancelInboundPlanForbidden() *CancelInboundPlanForbidden {
	return &CancelInboundPlanForbidden{}
}

/*
CancelInboundPlanForbidden describes a response with status code 403, with default header values.

Indicates that access to the resource is forbidden. Possible reasons include Access Denied, Unauthorized, Expired Token, or Invalid Signature.
*/
type CancelInboundPlanForbidden struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inbound_2024_03_20_models.ErrorList
}

// IsSuccess returns true when this cancel inbound plan forbidden response has a 2xx status code
func (o *CancelInboundPlanForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cancel inbound plan forbidden response has a 3xx status code
func (o *CancelInboundPlanForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cancel inbound plan forbidden response has a 4xx status code
func (o *CancelInboundPlanForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this cancel inbound plan forbidden response has a 5xx status code
func (o *CancelInboundPlanForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this cancel inbound plan forbidden response a status code equal to that given
func (o *CancelInboundPlanForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *CancelInboundPlanForbidden) Error() string {
	return fmt.Sprintf("[PUT /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/cancellation][%d] cancelInboundPlanForbidden  %+v", 403, o.Payload)
}

func (o *CancelInboundPlanForbidden) String() string {
	return fmt.Sprintf("[PUT /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/cancellation][%d] cancelInboundPlanForbidden  %+v", 403, o.Payload)
}

func (o *CancelInboundPlanForbidden) GetPayload() *fulfillment_inbound_2024_03_20_models.ErrorList {
	return o.Payload
}

func (o *CancelInboundPlanForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(fulfillment_inbound_2024_03_20_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCancelInboundPlanNotFound creates a CancelInboundPlanNotFound with default headers values
func NewCancelInboundPlanNotFound() *CancelInboundPlanNotFound {
	return &CancelInboundPlanNotFound{}
}

/*
CancelInboundPlanNotFound describes a response with status code 404, with default header values.

The resource specified does not exist.
*/
type CancelInboundPlanNotFound struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inbound_2024_03_20_models.ErrorList
}

// IsSuccess returns true when this cancel inbound plan not found response has a 2xx status code
func (o *CancelInboundPlanNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cancel inbound plan not found response has a 3xx status code
func (o *CancelInboundPlanNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cancel inbound plan not found response has a 4xx status code
func (o *CancelInboundPlanNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this cancel inbound plan not found response has a 5xx status code
func (o *CancelInboundPlanNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this cancel inbound plan not found response a status code equal to that given
func (o *CancelInboundPlanNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *CancelInboundPlanNotFound) Error() string {
	return fmt.Sprintf("[PUT /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/cancellation][%d] cancelInboundPlanNotFound  %+v", 404, o.Payload)
}

func (o *CancelInboundPlanNotFound) String() string {
	return fmt.Sprintf("[PUT /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/cancellation][%d] cancelInboundPlanNotFound  %+v", 404, o.Payload)
}

func (o *CancelInboundPlanNotFound) GetPayload() *fulfillment_inbound_2024_03_20_models.ErrorList {
	return o.Payload
}

func (o *CancelInboundPlanNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fulfillment_inbound_2024_03_20_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCancelInboundPlanRequestEntityTooLarge creates a CancelInboundPlanRequestEntityTooLarge with default headers values
func NewCancelInboundPlanRequestEntityTooLarge() *CancelInboundPlanRequestEntityTooLarge {
	return &CancelInboundPlanRequestEntityTooLarge{}
}

/*
CancelInboundPlanRequestEntityTooLarge describes a response with status code 413, with default header values.

The request size exceeded the maximum accepted size.
*/
type CancelInboundPlanRequestEntityTooLarge struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inbound_2024_03_20_models.ErrorList
}

// IsSuccess returns true when this cancel inbound plan request entity too large response has a 2xx status code
func (o *CancelInboundPlanRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cancel inbound plan request entity too large response has a 3xx status code
func (o *CancelInboundPlanRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cancel inbound plan request entity too large response has a 4xx status code
func (o *CancelInboundPlanRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this cancel inbound plan request entity too large response has a 5xx status code
func (o *CancelInboundPlanRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this cancel inbound plan request entity too large response a status code equal to that given
func (o *CancelInboundPlanRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *CancelInboundPlanRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PUT /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/cancellation][%d] cancelInboundPlanRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *CancelInboundPlanRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[PUT /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/cancellation][%d] cancelInboundPlanRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *CancelInboundPlanRequestEntityTooLarge) GetPayload() *fulfillment_inbound_2024_03_20_models.ErrorList {
	return o.Payload
}

func (o *CancelInboundPlanRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(fulfillment_inbound_2024_03_20_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCancelInboundPlanUnsupportedMediaType creates a CancelInboundPlanUnsupportedMediaType with default headers values
func NewCancelInboundPlanUnsupportedMediaType() *CancelInboundPlanUnsupportedMediaType {
	return &CancelInboundPlanUnsupportedMediaType{}
}

/*
CancelInboundPlanUnsupportedMediaType describes a response with status code 415, with default header values.

The request payload is in an unsupported format.
*/
type CancelInboundPlanUnsupportedMediaType struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inbound_2024_03_20_models.ErrorList
}

// IsSuccess returns true when this cancel inbound plan unsupported media type response has a 2xx status code
func (o *CancelInboundPlanUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cancel inbound plan unsupported media type response has a 3xx status code
func (o *CancelInboundPlanUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cancel inbound plan unsupported media type response has a 4xx status code
func (o *CancelInboundPlanUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this cancel inbound plan unsupported media type response has a 5xx status code
func (o *CancelInboundPlanUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this cancel inbound plan unsupported media type response a status code equal to that given
func (o *CancelInboundPlanUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *CancelInboundPlanUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PUT /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/cancellation][%d] cancelInboundPlanUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *CancelInboundPlanUnsupportedMediaType) String() string {
	return fmt.Sprintf("[PUT /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/cancellation][%d] cancelInboundPlanUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *CancelInboundPlanUnsupportedMediaType) GetPayload() *fulfillment_inbound_2024_03_20_models.ErrorList {
	return o.Payload
}

func (o *CancelInboundPlanUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(fulfillment_inbound_2024_03_20_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCancelInboundPlanTooManyRequests creates a CancelInboundPlanTooManyRequests with default headers values
func NewCancelInboundPlanTooManyRequests() *CancelInboundPlanTooManyRequests {
	return &CancelInboundPlanTooManyRequests{}
}

/*
CancelInboundPlanTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type CancelInboundPlanTooManyRequests struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inbound_2024_03_20_models.ErrorList
}

// IsSuccess returns true when this cancel inbound plan too many requests response has a 2xx status code
func (o *CancelInboundPlanTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cancel inbound plan too many requests response has a 3xx status code
func (o *CancelInboundPlanTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cancel inbound plan too many requests response has a 4xx status code
func (o *CancelInboundPlanTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this cancel inbound plan too many requests response has a 5xx status code
func (o *CancelInboundPlanTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this cancel inbound plan too many requests response a status code equal to that given
func (o *CancelInboundPlanTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *CancelInboundPlanTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/cancellation][%d] cancelInboundPlanTooManyRequests  %+v", 429, o.Payload)
}

func (o *CancelInboundPlanTooManyRequests) String() string {
	return fmt.Sprintf("[PUT /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/cancellation][%d] cancelInboundPlanTooManyRequests  %+v", 429, o.Payload)
}

func (o *CancelInboundPlanTooManyRequests) GetPayload() *fulfillment_inbound_2024_03_20_models.ErrorList {
	return o.Payload
}

func (o *CancelInboundPlanTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(fulfillment_inbound_2024_03_20_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCancelInboundPlanInternalServerError creates a CancelInboundPlanInternalServerError with default headers values
func NewCancelInboundPlanInternalServerError() *CancelInboundPlanInternalServerError {
	return &CancelInboundPlanInternalServerError{}
}

/*
CancelInboundPlanInternalServerError describes a response with status code 500, with default header values.

An unexpected condition occurred that prevented the server from fulfilling the request.
*/
type CancelInboundPlanInternalServerError struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inbound_2024_03_20_models.ErrorList
}

// IsSuccess returns true when this cancel inbound plan internal server error response has a 2xx status code
func (o *CancelInboundPlanInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cancel inbound plan internal server error response has a 3xx status code
func (o *CancelInboundPlanInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cancel inbound plan internal server error response has a 4xx status code
func (o *CancelInboundPlanInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this cancel inbound plan internal server error response has a 5xx status code
func (o *CancelInboundPlanInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this cancel inbound plan internal server error response a status code equal to that given
func (o *CancelInboundPlanInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *CancelInboundPlanInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/cancellation][%d] cancelInboundPlanInternalServerError  %+v", 500, o.Payload)
}

func (o *CancelInboundPlanInternalServerError) String() string {
	return fmt.Sprintf("[PUT /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/cancellation][%d] cancelInboundPlanInternalServerError  %+v", 500, o.Payload)
}

func (o *CancelInboundPlanInternalServerError) GetPayload() *fulfillment_inbound_2024_03_20_models.ErrorList {
	return o.Payload
}

func (o *CancelInboundPlanInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(fulfillment_inbound_2024_03_20_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCancelInboundPlanServiceUnavailable creates a CancelInboundPlanServiceUnavailable with default headers values
func NewCancelInboundPlanServiceUnavailable() *CancelInboundPlanServiceUnavailable {
	return &CancelInboundPlanServiceUnavailable{}
}

/*
CancelInboundPlanServiceUnavailable describes a response with status code 503, with default header values.

Temporary overloading or maintenance of the server.
*/
type CancelInboundPlanServiceUnavailable struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inbound_2024_03_20_models.ErrorList
}

// IsSuccess returns true when this cancel inbound plan service unavailable response has a 2xx status code
func (o *CancelInboundPlanServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cancel inbound plan service unavailable response has a 3xx status code
func (o *CancelInboundPlanServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cancel inbound plan service unavailable response has a 4xx status code
func (o *CancelInboundPlanServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this cancel inbound plan service unavailable response has a 5xx status code
func (o *CancelInboundPlanServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this cancel inbound plan service unavailable response a status code equal to that given
func (o *CancelInboundPlanServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *CancelInboundPlanServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/cancellation][%d] cancelInboundPlanServiceUnavailable  %+v", 503, o.Payload)
}

func (o *CancelInboundPlanServiceUnavailable) String() string {
	return fmt.Sprintf("[PUT /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/cancellation][%d] cancelInboundPlanServiceUnavailable  %+v", 503, o.Payload)
}

func (o *CancelInboundPlanServiceUnavailable) GetPayload() *fulfillment_inbound_2024_03_20_models.ErrorList {
	return o.Payload
}

func (o *CancelInboundPlanServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(fulfillment_inbound_2024_03_20_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
