// Code generated by go-swagger; DO NOT EDIT.

package fba_inbound

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/renabled/amzn-sp-api-go/api/fulfillmentInboundV0/fulfillment_inbound_v0_models"
)

// UpdateInboundShipmentReader is a Reader for the UpdateInboundShipment structure.
type UpdateInboundShipmentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateInboundShipmentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateInboundShipmentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateInboundShipmentBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateInboundShipmentUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateInboundShipmentForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateInboundShipmentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewUpdateInboundShipmentTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateInboundShipmentInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewUpdateInboundShipmentServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateInboundShipmentOK creates a UpdateInboundShipmentOK with default headers values
func NewUpdateInboundShipmentOK() *UpdateInboundShipmentOK {
	return &UpdateInboundShipmentOK{}
}

/*
UpdateInboundShipmentOK describes a response with status code 200, with default header values.

Success.
*/
type UpdateInboundShipmentOK struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inbound_v0_models.InboundShipmentResponse
}

// IsSuccess returns true when this update inbound shipment o k response has a 2xx status code
func (o *UpdateInboundShipmentOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update inbound shipment o k response has a 3xx status code
func (o *UpdateInboundShipmentOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update inbound shipment o k response has a 4xx status code
func (o *UpdateInboundShipmentOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update inbound shipment o k response has a 5xx status code
func (o *UpdateInboundShipmentOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update inbound shipment o k response a status code equal to that given
func (o *UpdateInboundShipmentOK) IsCode(code int) bool {
	return code == 200
}

func (o *UpdateInboundShipmentOK) Error() string {
	return fmt.Sprintf("[PUT /fba/inbound/v0/shipments/{shipmentId}][%d] updateInboundShipmentOK  %+v", 200, o.Payload)
}

func (o *UpdateInboundShipmentOK) String() string {
	return fmt.Sprintf("[PUT /fba/inbound/v0/shipments/{shipmentId}][%d] updateInboundShipmentOK  %+v", 200, o.Payload)
}

func (o *UpdateInboundShipmentOK) GetPayload() *fulfillment_inbound_v0_models.InboundShipmentResponse {
	return o.Payload
}

func (o *UpdateInboundShipmentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fulfillment_inbound_v0_models.InboundShipmentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateInboundShipmentBadRequest creates a UpdateInboundShipmentBadRequest with default headers values
func NewUpdateInboundShipmentBadRequest() *UpdateInboundShipmentBadRequest {
	return &UpdateInboundShipmentBadRequest{}
}

/*
UpdateInboundShipmentBadRequest describes a response with status code 400, with default header values.

Request has missing or invalid parameters and cannot be parsed.
*/
type UpdateInboundShipmentBadRequest struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inbound_v0_models.InboundShipmentResponse
}

// IsSuccess returns true when this update inbound shipment bad request response has a 2xx status code
func (o *UpdateInboundShipmentBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update inbound shipment bad request response has a 3xx status code
func (o *UpdateInboundShipmentBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update inbound shipment bad request response has a 4xx status code
func (o *UpdateInboundShipmentBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update inbound shipment bad request response has a 5xx status code
func (o *UpdateInboundShipmentBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update inbound shipment bad request response a status code equal to that given
func (o *UpdateInboundShipmentBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *UpdateInboundShipmentBadRequest) Error() string {
	return fmt.Sprintf("[PUT /fba/inbound/v0/shipments/{shipmentId}][%d] updateInboundShipmentBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateInboundShipmentBadRequest) String() string {
	return fmt.Sprintf("[PUT /fba/inbound/v0/shipments/{shipmentId}][%d] updateInboundShipmentBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateInboundShipmentBadRequest) GetPayload() *fulfillment_inbound_v0_models.InboundShipmentResponse {
	return o.Payload
}

func (o *UpdateInboundShipmentBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fulfillment_inbound_v0_models.InboundShipmentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateInboundShipmentUnauthorized creates a UpdateInboundShipmentUnauthorized with default headers values
func NewUpdateInboundShipmentUnauthorized() *UpdateInboundShipmentUnauthorized {
	return &UpdateInboundShipmentUnauthorized{}
}

/*
UpdateInboundShipmentUnauthorized describes a response with status code 401, with default header values.

The request's Authorization header is not formatted correctly or does not contain a valid token.
*/
type UpdateInboundShipmentUnauthorized struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inbound_v0_models.InboundShipmentResponse
}

// IsSuccess returns true when this update inbound shipment unauthorized response has a 2xx status code
func (o *UpdateInboundShipmentUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update inbound shipment unauthorized response has a 3xx status code
func (o *UpdateInboundShipmentUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update inbound shipment unauthorized response has a 4xx status code
func (o *UpdateInboundShipmentUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this update inbound shipment unauthorized response has a 5xx status code
func (o *UpdateInboundShipmentUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this update inbound shipment unauthorized response a status code equal to that given
func (o *UpdateInboundShipmentUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *UpdateInboundShipmentUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /fba/inbound/v0/shipments/{shipmentId}][%d] updateInboundShipmentUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdateInboundShipmentUnauthorized) String() string {
	return fmt.Sprintf("[PUT /fba/inbound/v0/shipments/{shipmentId}][%d] updateInboundShipmentUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdateInboundShipmentUnauthorized) GetPayload() *fulfillment_inbound_v0_models.InboundShipmentResponse {
	return o.Payload
}

func (o *UpdateInboundShipmentUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(fulfillment_inbound_v0_models.InboundShipmentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateInboundShipmentForbidden creates a UpdateInboundShipmentForbidden with default headers values
func NewUpdateInboundShipmentForbidden() *UpdateInboundShipmentForbidden {
	return &UpdateInboundShipmentForbidden{}
}

/*
UpdateInboundShipmentForbidden describes a response with status code 403, with default header values.

Indicates that access to the resource is forbidden. Possible reasons include Access Denied, Unauthorized, Expired Token, or Invalid Signature.
*/
type UpdateInboundShipmentForbidden struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inbound_v0_models.InboundShipmentResponse
}

// IsSuccess returns true when this update inbound shipment forbidden response has a 2xx status code
func (o *UpdateInboundShipmentForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update inbound shipment forbidden response has a 3xx status code
func (o *UpdateInboundShipmentForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update inbound shipment forbidden response has a 4xx status code
func (o *UpdateInboundShipmentForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update inbound shipment forbidden response has a 5xx status code
func (o *UpdateInboundShipmentForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update inbound shipment forbidden response a status code equal to that given
func (o *UpdateInboundShipmentForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *UpdateInboundShipmentForbidden) Error() string {
	return fmt.Sprintf("[PUT /fba/inbound/v0/shipments/{shipmentId}][%d] updateInboundShipmentForbidden  %+v", 403, o.Payload)
}

func (o *UpdateInboundShipmentForbidden) String() string {
	return fmt.Sprintf("[PUT /fba/inbound/v0/shipments/{shipmentId}][%d] updateInboundShipmentForbidden  %+v", 403, o.Payload)
}

func (o *UpdateInboundShipmentForbidden) GetPayload() *fulfillment_inbound_v0_models.InboundShipmentResponse {
	return o.Payload
}

func (o *UpdateInboundShipmentForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(fulfillment_inbound_v0_models.InboundShipmentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateInboundShipmentNotFound creates a UpdateInboundShipmentNotFound with default headers values
func NewUpdateInboundShipmentNotFound() *UpdateInboundShipmentNotFound {
	return &UpdateInboundShipmentNotFound{}
}

/*
UpdateInboundShipmentNotFound describes a response with status code 404, with default header values.

The specified resource does not exist.
*/
type UpdateInboundShipmentNotFound struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inbound_v0_models.InboundShipmentResponse
}

// IsSuccess returns true when this update inbound shipment not found response has a 2xx status code
func (o *UpdateInboundShipmentNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update inbound shipment not found response has a 3xx status code
func (o *UpdateInboundShipmentNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update inbound shipment not found response has a 4xx status code
func (o *UpdateInboundShipmentNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update inbound shipment not found response has a 5xx status code
func (o *UpdateInboundShipmentNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update inbound shipment not found response a status code equal to that given
func (o *UpdateInboundShipmentNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *UpdateInboundShipmentNotFound) Error() string {
	return fmt.Sprintf("[PUT /fba/inbound/v0/shipments/{shipmentId}][%d] updateInboundShipmentNotFound  %+v", 404, o.Payload)
}

func (o *UpdateInboundShipmentNotFound) String() string {
	return fmt.Sprintf("[PUT /fba/inbound/v0/shipments/{shipmentId}][%d] updateInboundShipmentNotFound  %+v", 404, o.Payload)
}

func (o *UpdateInboundShipmentNotFound) GetPayload() *fulfillment_inbound_v0_models.InboundShipmentResponse {
	return o.Payload
}

func (o *UpdateInboundShipmentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fulfillment_inbound_v0_models.InboundShipmentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateInboundShipmentTooManyRequests creates a UpdateInboundShipmentTooManyRequests with default headers values
func NewUpdateInboundShipmentTooManyRequests() *UpdateInboundShipmentTooManyRequests {
	return &UpdateInboundShipmentTooManyRequests{}
}

/*
UpdateInboundShipmentTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type UpdateInboundShipmentTooManyRequests struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inbound_v0_models.InboundShipmentResponse
}

// IsSuccess returns true when this update inbound shipment too many requests response has a 2xx status code
func (o *UpdateInboundShipmentTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update inbound shipment too many requests response has a 3xx status code
func (o *UpdateInboundShipmentTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update inbound shipment too many requests response has a 4xx status code
func (o *UpdateInboundShipmentTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this update inbound shipment too many requests response has a 5xx status code
func (o *UpdateInboundShipmentTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this update inbound shipment too many requests response a status code equal to that given
func (o *UpdateInboundShipmentTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *UpdateInboundShipmentTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /fba/inbound/v0/shipments/{shipmentId}][%d] updateInboundShipmentTooManyRequests  %+v", 429, o.Payload)
}

func (o *UpdateInboundShipmentTooManyRequests) String() string {
	return fmt.Sprintf("[PUT /fba/inbound/v0/shipments/{shipmentId}][%d] updateInboundShipmentTooManyRequests  %+v", 429, o.Payload)
}

func (o *UpdateInboundShipmentTooManyRequests) GetPayload() *fulfillment_inbound_v0_models.InboundShipmentResponse {
	return o.Payload
}

func (o *UpdateInboundShipmentTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(fulfillment_inbound_v0_models.InboundShipmentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateInboundShipmentInternalServerError creates a UpdateInboundShipmentInternalServerError with default headers values
func NewUpdateInboundShipmentInternalServerError() *UpdateInboundShipmentInternalServerError {
	return &UpdateInboundShipmentInternalServerError{}
}

/*
UpdateInboundShipmentInternalServerError describes a response with status code 500, with default header values.

An unexpected condition occurred that prevented the server from fulfilling the request.
*/
type UpdateInboundShipmentInternalServerError struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inbound_v0_models.InboundShipmentResponse
}

// IsSuccess returns true when this update inbound shipment internal server error response has a 2xx status code
func (o *UpdateInboundShipmentInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update inbound shipment internal server error response has a 3xx status code
func (o *UpdateInboundShipmentInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update inbound shipment internal server error response has a 4xx status code
func (o *UpdateInboundShipmentInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this update inbound shipment internal server error response has a 5xx status code
func (o *UpdateInboundShipmentInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this update inbound shipment internal server error response a status code equal to that given
func (o *UpdateInboundShipmentInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *UpdateInboundShipmentInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /fba/inbound/v0/shipments/{shipmentId}][%d] updateInboundShipmentInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateInboundShipmentInternalServerError) String() string {
	return fmt.Sprintf("[PUT /fba/inbound/v0/shipments/{shipmentId}][%d] updateInboundShipmentInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateInboundShipmentInternalServerError) GetPayload() *fulfillment_inbound_v0_models.InboundShipmentResponse {
	return o.Payload
}

func (o *UpdateInboundShipmentInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(fulfillment_inbound_v0_models.InboundShipmentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateInboundShipmentServiceUnavailable creates a UpdateInboundShipmentServiceUnavailable with default headers values
func NewUpdateInboundShipmentServiceUnavailable() *UpdateInboundShipmentServiceUnavailable {
	return &UpdateInboundShipmentServiceUnavailable{}
}

/*
UpdateInboundShipmentServiceUnavailable describes a response with status code 503, with default header values.

Temporary overloading or maintenance of the server.
*/
type UpdateInboundShipmentServiceUnavailable struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inbound_v0_models.InboundShipmentResponse
}

// IsSuccess returns true when this update inbound shipment service unavailable response has a 2xx status code
func (o *UpdateInboundShipmentServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update inbound shipment service unavailable response has a 3xx status code
func (o *UpdateInboundShipmentServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update inbound shipment service unavailable response has a 4xx status code
func (o *UpdateInboundShipmentServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this update inbound shipment service unavailable response has a 5xx status code
func (o *UpdateInboundShipmentServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this update inbound shipment service unavailable response a status code equal to that given
func (o *UpdateInboundShipmentServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *UpdateInboundShipmentServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /fba/inbound/v0/shipments/{shipmentId}][%d] updateInboundShipmentServiceUnavailable  %+v", 503, o.Payload)
}

func (o *UpdateInboundShipmentServiceUnavailable) String() string {
	return fmt.Sprintf("[PUT /fba/inbound/v0/shipments/{shipmentId}][%d] updateInboundShipmentServiceUnavailable  %+v", 503, o.Payload)
}

func (o *UpdateInboundShipmentServiceUnavailable) GetPayload() *fulfillment_inbound_v0_models.InboundShipmentResponse {
	return o.Payload
}

func (o *UpdateInboundShipmentServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(fulfillment_inbound_v0_models.InboundShipmentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
