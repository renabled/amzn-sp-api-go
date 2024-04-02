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

// ScheduleSelfShipAppointmentReader is a Reader for the ScheduleSelfShipAppointment structure.
type ScheduleSelfShipAppointmentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ScheduleSelfShipAppointmentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewScheduleSelfShipAppointmentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewScheduleSelfShipAppointmentBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewScheduleSelfShipAppointmentForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewScheduleSelfShipAppointmentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewScheduleSelfShipAppointmentRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewScheduleSelfShipAppointmentUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewScheduleSelfShipAppointmentTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewScheduleSelfShipAppointmentInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewScheduleSelfShipAppointmentServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewScheduleSelfShipAppointmentOK creates a ScheduleSelfShipAppointmentOK with default headers values
func NewScheduleSelfShipAppointmentOK() *ScheduleSelfShipAppointmentOK {
	return &ScheduleSelfShipAppointmentOK{}
}

/*
ScheduleSelfShipAppointmentOK describes a response with status code 200, with default header values.

ScheduleSelfShipAppointment 200 response
*/
type ScheduleSelfShipAppointmentOK struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inboundv2024_03_20_models.ScheduleSelfShipAppointmentResponse
}

// IsSuccess returns true when this schedule self ship appointment o k response has a 2xx status code
func (o *ScheduleSelfShipAppointmentOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this schedule self ship appointment o k response has a 3xx status code
func (o *ScheduleSelfShipAppointmentOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this schedule self ship appointment o k response has a 4xx status code
func (o *ScheduleSelfShipAppointmentOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this schedule self ship appointment o k response has a 5xx status code
func (o *ScheduleSelfShipAppointmentOK) IsServerError() bool {
	return false
}

// IsCode returns true when this schedule self ship appointment o k response a status code equal to that given
func (o *ScheduleSelfShipAppointmentOK) IsCode(code int) bool {
	return code == 200
}

func (o *ScheduleSelfShipAppointmentOK) Error() string {
	return fmt.Sprintf("[POST /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/selfShipAppointmentSlots/{slotId}/schedule][%d] scheduleSelfShipAppointmentOK  %+v", 200, o.Payload)
}

func (o *ScheduleSelfShipAppointmentOK) String() string {
	return fmt.Sprintf("[POST /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/selfShipAppointmentSlots/{slotId}/schedule][%d] scheduleSelfShipAppointmentOK  %+v", 200, o.Payload)
}

func (o *ScheduleSelfShipAppointmentOK) GetPayload() *fulfillment_inboundv2024_03_20_models.ScheduleSelfShipAppointmentResponse {
	return o.Payload
}

func (o *ScheduleSelfShipAppointmentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fulfillment_inboundv2024_03_20_models.ScheduleSelfShipAppointmentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewScheduleSelfShipAppointmentBadRequest creates a ScheduleSelfShipAppointmentBadRequest with default headers values
func NewScheduleSelfShipAppointmentBadRequest() *ScheduleSelfShipAppointmentBadRequest {
	return &ScheduleSelfShipAppointmentBadRequest{}
}

/*
ScheduleSelfShipAppointmentBadRequest describes a response with status code 400, with default header values.

Request has missing or invalid parameters and cannot be parsed.
*/
type ScheduleSelfShipAppointmentBadRequest struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inboundv2024_03_20_models.ErrorList
}

// IsSuccess returns true when this schedule self ship appointment bad request response has a 2xx status code
func (o *ScheduleSelfShipAppointmentBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this schedule self ship appointment bad request response has a 3xx status code
func (o *ScheduleSelfShipAppointmentBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this schedule self ship appointment bad request response has a 4xx status code
func (o *ScheduleSelfShipAppointmentBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this schedule self ship appointment bad request response has a 5xx status code
func (o *ScheduleSelfShipAppointmentBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this schedule self ship appointment bad request response a status code equal to that given
func (o *ScheduleSelfShipAppointmentBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *ScheduleSelfShipAppointmentBadRequest) Error() string {
	return fmt.Sprintf("[POST /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/selfShipAppointmentSlots/{slotId}/schedule][%d] scheduleSelfShipAppointmentBadRequest  %+v", 400, o.Payload)
}

func (o *ScheduleSelfShipAppointmentBadRequest) String() string {
	return fmt.Sprintf("[POST /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/selfShipAppointmentSlots/{slotId}/schedule][%d] scheduleSelfShipAppointmentBadRequest  %+v", 400, o.Payload)
}

func (o *ScheduleSelfShipAppointmentBadRequest) GetPayload() *fulfillment_inboundv2024_03_20_models.ErrorList {
	return o.Payload
}

func (o *ScheduleSelfShipAppointmentBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewScheduleSelfShipAppointmentForbidden creates a ScheduleSelfShipAppointmentForbidden with default headers values
func NewScheduleSelfShipAppointmentForbidden() *ScheduleSelfShipAppointmentForbidden {
	return &ScheduleSelfShipAppointmentForbidden{}
}

/*
ScheduleSelfShipAppointmentForbidden describes a response with status code 403, with default header values.

Indicates that access to the resource is forbidden. Possible reasons include Access Denied, Unauthorized, Expired Token, or Invalid Signature.
*/
type ScheduleSelfShipAppointmentForbidden struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inboundv2024_03_20_models.ErrorList
}

// IsSuccess returns true when this schedule self ship appointment forbidden response has a 2xx status code
func (o *ScheduleSelfShipAppointmentForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this schedule self ship appointment forbidden response has a 3xx status code
func (o *ScheduleSelfShipAppointmentForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this schedule self ship appointment forbidden response has a 4xx status code
func (o *ScheduleSelfShipAppointmentForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this schedule self ship appointment forbidden response has a 5xx status code
func (o *ScheduleSelfShipAppointmentForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this schedule self ship appointment forbidden response a status code equal to that given
func (o *ScheduleSelfShipAppointmentForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ScheduleSelfShipAppointmentForbidden) Error() string {
	return fmt.Sprintf("[POST /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/selfShipAppointmentSlots/{slotId}/schedule][%d] scheduleSelfShipAppointmentForbidden  %+v", 403, o.Payload)
}

func (o *ScheduleSelfShipAppointmentForbidden) String() string {
	return fmt.Sprintf("[POST /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/selfShipAppointmentSlots/{slotId}/schedule][%d] scheduleSelfShipAppointmentForbidden  %+v", 403, o.Payload)
}

func (o *ScheduleSelfShipAppointmentForbidden) GetPayload() *fulfillment_inboundv2024_03_20_models.ErrorList {
	return o.Payload
}

func (o *ScheduleSelfShipAppointmentForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewScheduleSelfShipAppointmentNotFound creates a ScheduleSelfShipAppointmentNotFound with default headers values
func NewScheduleSelfShipAppointmentNotFound() *ScheduleSelfShipAppointmentNotFound {
	return &ScheduleSelfShipAppointmentNotFound{}
}

/*
ScheduleSelfShipAppointmentNotFound describes a response with status code 404, with default header values.

The resource specified does not exist.
*/
type ScheduleSelfShipAppointmentNotFound struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inboundv2024_03_20_models.ErrorList
}

// IsSuccess returns true when this schedule self ship appointment not found response has a 2xx status code
func (o *ScheduleSelfShipAppointmentNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this schedule self ship appointment not found response has a 3xx status code
func (o *ScheduleSelfShipAppointmentNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this schedule self ship appointment not found response has a 4xx status code
func (o *ScheduleSelfShipAppointmentNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this schedule self ship appointment not found response has a 5xx status code
func (o *ScheduleSelfShipAppointmentNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this schedule self ship appointment not found response a status code equal to that given
func (o *ScheduleSelfShipAppointmentNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *ScheduleSelfShipAppointmentNotFound) Error() string {
	return fmt.Sprintf("[POST /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/selfShipAppointmentSlots/{slotId}/schedule][%d] scheduleSelfShipAppointmentNotFound  %+v", 404, o.Payload)
}

func (o *ScheduleSelfShipAppointmentNotFound) String() string {
	return fmt.Sprintf("[POST /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/selfShipAppointmentSlots/{slotId}/schedule][%d] scheduleSelfShipAppointmentNotFound  %+v", 404, o.Payload)
}

func (o *ScheduleSelfShipAppointmentNotFound) GetPayload() *fulfillment_inboundv2024_03_20_models.ErrorList {
	return o.Payload
}

func (o *ScheduleSelfShipAppointmentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewScheduleSelfShipAppointmentRequestEntityTooLarge creates a ScheduleSelfShipAppointmentRequestEntityTooLarge with default headers values
func NewScheduleSelfShipAppointmentRequestEntityTooLarge() *ScheduleSelfShipAppointmentRequestEntityTooLarge {
	return &ScheduleSelfShipAppointmentRequestEntityTooLarge{}
}

/*
ScheduleSelfShipAppointmentRequestEntityTooLarge describes a response with status code 413, with default header values.

The request size exceeded the maximum accepted size.
*/
type ScheduleSelfShipAppointmentRequestEntityTooLarge struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inboundv2024_03_20_models.ErrorList
}

// IsSuccess returns true when this schedule self ship appointment request entity too large response has a 2xx status code
func (o *ScheduleSelfShipAppointmentRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this schedule self ship appointment request entity too large response has a 3xx status code
func (o *ScheduleSelfShipAppointmentRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this schedule self ship appointment request entity too large response has a 4xx status code
func (o *ScheduleSelfShipAppointmentRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this schedule self ship appointment request entity too large response has a 5xx status code
func (o *ScheduleSelfShipAppointmentRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this schedule self ship appointment request entity too large response a status code equal to that given
func (o *ScheduleSelfShipAppointmentRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *ScheduleSelfShipAppointmentRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/selfShipAppointmentSlots/{slotId}/schedule][%d] scheduleSelfShipAppointmentRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *ScheduleSelfShipAppointmentRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[POST /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/selfShipAppointmentSlots/{slotId}/schedule][%d] scheduleSelfShipAppointmentRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *ScheduleSelfShipAppointmentRequestEntityTooLarge) GetPayload() *fulfillment_inboundv2024_03_20_models.ErrorList {
	return o.Payload
}

func (o *ScheduleSelfShipAppointmentRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewScheduleSelfShipAppointmentUnsupportedMediaType creates a ScheduleSelfShipAppointmentUnsupportedMediaType with default headers values
func NewScheduleSelfShipAppointmentUnsupportedMediaType() *ScheduleSelfShipAppointmentUnsupportedMediaType {
	return &ScheduleSelfShipAppointmentUnsupportedMediaType{}
}

/*
ScheduleSelfShipAppointmentUnsupportedMediaType describes a response with status code 415, with default header values.

The request payload is in an unsupported format.
*/
type ScheduleSelfShipAppointmentUnsupportedMediaType struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inboundv2024_03_20_models.ErrorList
}

// IsSuccess returns true when this schedule self ship appointment unsupported media type response has a 2xx status code
func (o *ScheduleSelfShipAppointmentUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this schedule self ship appointment unsupported media type response has a 3xx status code
func (o *ScheduleSelfShipAppointmentUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this schedule self ship appointment unsupported media type response has a 4xx status code
func (o *ScheduleSelfShipAppointmentUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this schedule self ship appointment unsupported media type response has a 5xx status code
func (o *ScheduleSelfShipAppointmentUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this schedule self ship appointment unsupported media type response a status code equal to that given
func (o *ScheduleSelfShipAppointmentUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *ScheduleSelfShipAppointmentUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/selfShipAppointmentSlots/{slotId}/schedule][%d] scheduleSelfShipAppointmentUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *ScheduleSelfShipAppointmentUnsupportedMediaType) String() string {
	return fmt.Sprintf("[POST /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/selfShipAppointmentSlots/{slotId}/schedule][%d] scheduleSelfShipAppointmentUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *ScheduleSelfShipAppointmentUnsupportedMediaType) GetPayload() *fulfillment_inboundv2024_03_20_models.ErrorList {
	return o.Payload
}

func (o *ScheduleSelfShipAppointmentUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewScheduleSelfShipAppointmentTooManyRequests creates a ScheduleSelfShipAppointmentTooManyRequests with default headers values
func NewScheduleSelfShipAppointmentTooManyRequests() *ScheduleSelfShipAppointmentTooManyRequests {
	return &ScheduleSelfShipAppointmentTooManyRequests{}
}

/*
ScheduleSelfShipAppointmentTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type ScheduleSelfShipAppointmentTooManyRequests struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inboundv2024_03_20_models.ErrorList
}

// IsSuccess returns true when this schedule self ship appointment too many requests response has a 2xx status code
func (o *ScheduleSelfShipAppointmentTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this schedule self ship appointment too many requests response has a 3xx status code
func (o *ScheduleSelfShipAppointmentTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this schedule self ship appointment too many requests response has a 4xx status code
func (o *ScheduleSelfShipAppointmentTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this schedule self ship appointment too many requests response has a 5xx status code
func (o *ScheduleSelfShipAppointmentTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this schedule self ship appointment too many requests response a status code equal to that given
func (o *ScheduleSelfShipAppointmentTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *ScheduleSelfShipAppointmentTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/selfShipAppointmentSlots/{slotId}/schedule][%d] scheduleSelfShipAppointmentTooManyRequests  %+v", 429, o.Payload)
}

func (o *ScheduleSelfShipAppointmentTooManyRequests) String() string {
	return fmt.Sprintf("[POST /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/selfShipAppointmentSlots/{slotId}/schedule][%d] scheduleSelfShipAppointmentTooManyRequests  %+v", 429, o.Payload)
}

func (o *ScheduleSelfShipAppointmentTooManyRequests) GetPayload() *fulfillment_inboundv2024_03_20_models.ErrorList {
	return o.Payload
}

func (o *ScheduleSelfShipAppointmentTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewScheduleSelfShipAppointmentInternalServerError creates a ScheduleSelfShipAppointmentInternalServerError with default headers values
func NewScheduleSelfShipAppointmentInternalServerError() *ScheduleSelfShipAppointmentInternalServerError {
	return &ScheduleSelfShipAppointmentInternalServerError{}
}

/*
ScheduleSelfShipAppointmentInternalServerError describes a response with status code 500, with default header values.

An unexpected condition occurred that prevented the server from fulfilling the request.
*/
type ScheduleSelfShipAppointmentInternalServerError struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inboundv2024_03_20_models.ErrorList
}

// IsSuccess returns true when this schedule self ship appointment internal server error response has a 2xx status code
func (o *ScheduleSelfShipAppointmentInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this schedule self ship appointment internal server error response has a 3xx status code
func (o *ScheduleSelfShipAppointmentInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this schedule self ship appointment internal server error response has a 4xx status code
func (o *ScheduleSelfShipAppointmentInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this schedule self ship appointment internal server error response has a 5xx status code
func (o *ScheduleSelfShipAppointmentInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this schedule self ship appointment internal server error response a status code equal to that given
func (o *ScheduleSelfShipAppointmentInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ScheduleSelfShipAppointmentInternalServerError) Error() string {
	return fmt.Sprintf("[POST /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/selfShipAppointmentSlots/{slotId}/schedule][%d] scheduleSelfShipAppointmentInternalServerError  %+v", 500, o.Payload)
}

func (o *ScheduleSelfShipAppointmentInternalServerError) String() string {
	return fmt.Sprintf("[POST /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/selfShipAppointmentSlots/{slotId}/schedule][%d] scheduleSelfShipAppointmentInternalServerError  %+v", 500, o.Payload)
}

func (o *ScheduleSelfShipAppointmentInternalServerError) GetPayload() *fulfillment_inboundv2024_03_20_models.ErrorList {
	return o.Payload
}

func (o *ScheduleSelfShipAppointmentInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewScheduleSelfShipAppointmentServiceUnavailable creates a ScheduleSelfShipAppointmentServiceUnavailable with default headers values
func NewScheduleSelfShipAppointmentServiceUnavailable() *ScheduleSelfShipAppointmentServiceUnavailable {
	return &ScheduleSelfShipAppointmentServiceUnavailable{}
}

/*
ScheduleSelfShipAppointmentServiceUnavailable describes a response with status code 503, with default header values.

Temporary overloading or maintenance of the server.
*/
type ScheduleSelfShipAppointmentServiceUnavailable struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inboundv2024_03_20_models.ErrorList
}

// IsSuccess returns true when this schedule self ship appointment service unavailable response has a 2xx status code
func (o *ScheduleSelfShipAppointmentServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this schedule self ship appointment service unavailable response has a 3xx status code
func (o *ScheduleSelfShipAppointmentServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this schedule self ship appointment service unavailable response has a 4xx status code
func (o *ScheduleSelfShipAppointmentServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this schedule self ship appointment service unavailable response has a 5xx status code
func (o *ScheduleSelfShipAppointmentServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this schedule self ship appointment service unavailable response a status code equal to that given
func (o *ScheduleSelfShipAppointmentServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *ScheduleSelfShipAppointmentServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/selfShipAppointmentSlots/{slotId}/schedule][%d] scheduleSelfShipAppointmentServiceUnavailable  %+v", 503, o.Payload)
}

func (o *ScheduleSelfShipAppointmentServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/selfShipAppointmentSlots/{slotId}/schedule][%d] scheduleSelfShipAppointmentServiceUnavailable  %+v", 503, o.Payload)
}

func (o *ScheduleSelfShipAppointmentServiceUnavailable) GetPayload() *fulfillment_inboundv2024_03_20_models.ErrorList {
	return o.Payload
}

func (o *ScheduleSelfShipAppointmentServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
