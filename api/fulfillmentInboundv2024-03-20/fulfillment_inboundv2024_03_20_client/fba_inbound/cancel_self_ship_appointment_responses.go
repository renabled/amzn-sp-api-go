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

// CancelSelfShipAppointmentReader is a Reader for the CancelSelfShipAppointment structure.
type CancelSelfShipAppointmentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CancelSelfShipAppointmentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewCancelSelfShipAppointmentAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCancelSelfShipAppointmentBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCancelSelfShipAppointmentForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCancelSelfShipAppointmentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewCancelSelfShipAppointmentRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewCancelSelfShipAppointmentUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewCancelSelfShipAppointmentTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCancelSelfShipAppointmentInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewCancelSelfShipAppointmentServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCancelSelfShipAppointmentAccepted creates a CancelSelfShipAppointmentAccepted with default headers values
func NewCancelSelfShipAppointmentAccepted() *CancelSelfShipAppointmentAccepted {
	return &CancelSelfShipAppointmentAccepted{}
}

/*
CancelSelfShipAppointmentAccepted describes a response with status code 202, with default header values.

CancelSelfShipAppointment 202 response
*/
type CancelSelfShipAppointmentAccepted struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inboundv2024_03_20_models.CancelSelfShipAppointmentResponse
}

// IsSuccess returns true when this cancel self ship appointment accepted response has a 2xx status code
func (o *CancelSelfShipAppointmentAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cancel self ship appointment accepted response has a 3xx status code
func (o *CancelSelfShipAppointmentAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cancel self ship appointment accepted response has a 4xx status code
func (o *CancelSelfShipAppointmentAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this cancel self ship appointment accepted response has a 5xx status code
func (o *CancelSelfShipAppointmentAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this cancel self ship appointment accepted response a status code equal to that given
func (o *CancelSelfShipAppointmentAccepted) IsCode(code int) bool {
	return code == 202
}

func (o *CancelSelfShipAppointmentAccepted) Error() string {
	return fmt.Sprintf("[PUT /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/selfShipAppointmentSlots/{slotId}/cancellation][%d] cancelSelfShipAppointmentAccepted  %+v", 202, o.Payload)
}

func (o *CancelSelfShipAppointmentAccepted) String() string {
	return fmt.Sprintf("[PUT /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/selfShipAppointmentSlots/{slotId}/cancellation][%d] cancelSelfShipAppointmentAccepted  %+v", 202, o.Payload)
}

func (o *CancelSelfShipAppointmentAccepted) GetPayload() *fulfillment_inboundv2024_03_20_models.CancelSelfShipAppointmentResponse {
	return o.Payload
}

func (o *CancelSelfShipAppointmentAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fulfillment_inboundv2024_03_20_models.CancelSelfShipAppointmentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCancelSelfShipAppointmentBadRequest creates a CancelSelfShipAppointmentBadRequest with default headers values
func NewCancelSelfShipAppointmentBadRequest() *CancelSelfShipAppointmentBadRequest {
	return &CancelSelfShipAppointmentBadRequest{}
}

/*
CancelSelfShipAppointmentBadRequest describes a response with status code 400, with default header values.

Request has missing or invalid parameters and cannot be parsed.
*/
type CancelSelfShipAppointmentBadRequest struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inboundv2024_03_20_models.ErrorList
}

// IsSuccess returns true when this cancel self ship appointment bad request response has a 2xx status code
func (o *CancelSelfShipAppointmentBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cancel self ship appointment bad request response has a 3xx status code
func (o *CancelSelfShipAppointmentBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cancel self ship appointment bad request response has a 4xx status code
func (o *CancelSelfShipAppointmentBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this cancel self ship appointment bad request response has a 5xx status code
func (o *CancelSelfShipAppointmentBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this cancel self ship appointment bad request response a status code equal to that given
func (o *CancelSelfShipAppointmentBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *CancelSelfShipAppointmentBadRequest) Error() string {
	return fmt.Sprintf("[PUT /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/selfShipAppointmentSlots/{slotId}/cancellation][%d] cancelSelfShipAppointmentBadRequest  %+v", 400, o.Payload)
}

func (o *CancelSelfShipAppointmentBadRequest) String() string {
	return fmt.Sprintf("[PUT /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/selfShipAppointmentSlots/{slotId}/cancellation][%d] cancelSelfShipAppointmentBadRequest  %+v", 400, o.Payload)
}

func (o *CancelSelfShipAppointmentBadRequest) GetPayload() *fulfillment_inboundv2024_03_20_models.ErrorList {
	return o.Payload
}

func (o *CancelSelfShipAppointmentBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCancelSelfShipAppointmentForbidden creates a CancelSelfShipAppointmentForbidden with default headers values
func NewCancelSelfShipAppointmentForbidden() *CancelSelfShipAppointmentForbidden {
	return &CancelSelfShipAppointmentForbidden{}
}

/*
CancelSelfShipAppointmentForbidden describes a response with status code 403, with default header values.

Indicates that access to the resource is forbidden. Possible reasons include Access Denied, Unauthorized, Expired Token, or Invalid Signature.
*/
type CancelSelfShipAppointmentForbidden struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inboundv2024_03_20_models.ErrorList
}

// IsSuccess returns true when this cancel self ship appointment forbidden response has a 2xx status code
func (o *CancelSelfShipAppointmentForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cancel self ship appointment forbidden response has a 3xx status code
func (o *CancelSelfShipAppointmentForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cancel self ship appointment forbidden response has a 4xx status code
func (o *CancelSelfShipAppointmentForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this cancel self ship appointment forbidden response has a 5xx status code
func (o *CancelSelfShipAppointmentForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this cancel self ship appointment forbidden response a status code equal to that given
func (o *CancelSelfShipAppointmentForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *CancelSelfShipAppointmentForbidden) Error() string {
	return fmt.Sprintf("[PUT /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/selfShipAppointmentSlots/{slotId}/cancellation][%d] cancelSelfShipAppointmentForbidden  %+v", 403, o.Payload)
}

func (o *CancelSelfShipAppointmentForbidden) String() string {
	return fmt.Sprintf("[PUT /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/selfShipAppointmentSlots/{slotId}/cancellation][%d] cancelSelfShipAppointmentForbidden  %+v", 403, o.Payload)
}

func (o *CancelSelfShipAppointmentForbidden) GetPayload() *fulfillment_inboundv2024_03_20_models.ErrorList {
	return o.Payload
}

func (o *CancelSelfShipAppointmentForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCancelSelfShipAppointmentNotFound creates a CancelSelfShipAppointmentNotFound with default headers values
func NewCancelSelfShipAppointmentNotFound() *CancelSelfShipAppointmentNotFound {
	return &CancelSelfShipAppointmentNotFound{}
}

/*
CancelSelfShipAppointmentNotFound describes a response with status code 404, with default header values.

The resource specified does not exist.
*/
type CancelSelfShipAppointmentNotFound struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inboundv2024_03_20_models.ErrorList
}

// IsSuccess returns true when this cancel self ship appointment not found response has a 2xx status code
func (o *CancelSelfShipAppointmentNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cancel self ship appointment not found response has a 3xx status code
func (o *CancelSelfShipAppointmentNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cancel self ship appointment not found response has a 4xx status code
func (o *CancelSelfShipAppointmentNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this cancel self ship appointment not found response has a 5xx status code
func (o *CancelSelfShipAppointmentNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this cancel self ship appointment not found response a status code equal to that given
func (o *CancelSelfShipAppointmentNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *CancelSelfShipAppointmentNotFound) Error() string {
	return fmt.Sprintf("[PUT /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/selfShipAppointmentSlots/{slotId}/cancellation][%d] cancelSelfShipAppointmentNotFound  %+v", 404, o.Payload)
}

func (o *CancelSelfShipAppointmentNotFound) String() string {
	return fmt.Sprintf("[PUT /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/selfShipAppointmentSlots/{slotId}/cancellation][%d] cancelSelfShipAppointmentNotFound  %+v", 404, o.Payload)
}

func (o *CancelSelfShipAppointmentNotFound) GetPayload() *fulfillment_inboundv2024_03_20_models.ErrorList {
	return o.Payload
}

func (o *CancelSelfShipAppointmentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCancelSelfShipAppointmentRequestEntityTooLarge creates a CancelSelfShipAppointmentRequestEntityTooLarge with default headers values
func NewCancelSelfShipAppointmentRequestEntityTooLarge() *CancelSelfShipAppointmentRequestEntityTooLarge {
	return &CancelSelfShipAppointmentRequestEntityTooLarge{}
}

/*
CancelSelfShipAppointmentRequestEntityTooLarge describes a response with status code 413, with default header values.

The request size exceeded the maximum accepted size.
*/
type CancelSelfShipAppointmentRequestEntityTooLarge struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inboundv2024_03_20_models.ErrorList
}

// IsSuccess returns true when this cancel self ship appointment request entity too large response has a 2xx status code
func (o *CancelSelfShipAppointmentRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cancel self ship appointment request entity too large response has a 3xx status code
func (o *CancelSelfShipAppointmentRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cancel self ship appointment request entity too large response has a 4xx status code
func (o *CancelSelfShipAppointmentRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this cancel self ship appointment request entity too large response has a 5xx status code
func (o *CancelSelfShipAppointmentRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this cancel self ship appointment request entity too large response a status code equal to that given
func (o *CancelSelfShipAppointmentRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *CancelSelfShipAppointmentRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PUT /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/selfShipAppointmentSlots/{slotId}/cancellation][%d] cancelSelfShipAppointmentRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *CancelSelfShipAppointmentRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[PUT /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/selfShipAppointmentSlots/{slotId}/cancellation][%d] cancelSelfShipAppointmentRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *CancelSelfShipAppointmentRequestEntityTooLarge) GetPayload() *fulfillment_inboundv2024_03_20_models.ErrorList {
	return o.Payload
}

func (o *CancelSelfShipAppointmentRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCancelSelfShipAppointmentUnsupportedMediaType creates a CancelSelfShipAppointmentUnsupportedMediaType with default headers values
func NewCancelSelfShipAppointmentUnsupportedMediaType() *CancelSelfShipAppointmentUnsupportedMediaType {
	return &CancelSelfShipAppointmentUnsupportedMediaType{}
}

/*
CancelSelfShipAppointmentUnsupportedMediaType describes a response with status code 415, with default header values.

The request payload is in an unsupported format.
*/
type CancelSelfShipAppointmentUnsupportedMediaType struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inboundv2024_03_20_models.ErrorList
}

// IsSuccess returns true when this cancel self ship appointment unsupported media type response has a 2xx status code
func (o *CancelSelfShipAppointmentUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cancel self ship appointment unsupported media type response has a 3xx status code
func (o *CancelSelfShipAppointmentUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cancel self ship appointment unsupported media type response has a 4xx status code
func (o *CancelSelfShipAppointmentUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this cancel self ship appointment unsupported media type response has a 5xx status code
func (o *CancelSelfShipAppointmentUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this cancel self ship appointment unsupported media type response a status code equal to that given
func (o *CancelSelfShipAppointmentUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *CancelSelfShipAppointmentUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PUT /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/selfShipAppointmentSlots/{slotId}/cancellation][%d] cancelSelfShipAppointmentUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *CancelSelfShipAppointmentUnsupportedMediaType) String() string {
	return fmt.Sprintf("[PUT /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/selfShipAppointmentSlots/{slotId}/cancellation][%d] cancelSelfShipAppointmentUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *CancelSelfShipAppointmentUnsupportedMediaType) GetPayload() *fulfillment_inboundv2024_03_20_models.ErrorList {
	return o.Payload
}

func (o *CancelSelfShipAppointmentUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCancelSelfShipAppointmentTooManyRequests creates a CancelSelfShipAppointmentTooManyRequests with default headers values
func NewCancelSelfShipAppointmentTooManyRequests() *CancelSelfShipAppointmentTooManyRequests {
	return &CancelSelfShipAppointmentTooManyRequests{}
}

/*
CancelSelfShipAppointmentTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type CancelSelfShipAppointmentTooManyRequests struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inboundv2024_03_20_models.ErrorList
}

// IsSuccess returns true when this cancel self ship appointment too many requests response has a 2xx status code
func (o *CancelSelfShipAppointmentTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cancel self ship appointment too many requests response has a 3xx status code
func (o *CancelSelfShipAppointmentTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cancel self ship appointment too many requests response has a 4xx status code
func (o *CancelSelfShipAppointmentTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this cancel self ship appointment too many requests response has a 5xx status code
func (o *CancelSelfShipAppointmentTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this cancel self ship appointment too many requests response a status code equal to that given
func (o *CancelSelfShipAppointmentTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *CancelSelfShipAppointmentTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/selfShipAppointmentSlots/{slotId}/cancellation][%d] cancelSelfShipAppointmentTooManyRequests  %+v", 429, o.Payload)
}

func (o *CancelSelfShipAppointmentTooManyRequests) String() string {
	return fmt.Sprintf("[PUT /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/selfShipAppointmentSlots/{slotId}/cancellation][%d] cancelSelfShipAppointmentTooManyRequests  %+v", 429, o.Payload)
}

func (o *CancelSelfShipAppointmentTooManyRequests) GetPayload() *fulfillment_inboundv2024_03_20_models.ErrorList {
	return o.Payload
}

func (o *CancelSelfShipAppointmentTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCancelSelfShipAppointmentInternalServerError creates a CancelSelfShipAppointmentInternalServerError with default headers values
func NewCancelSelfShipAppointmentInternalServerError() *CancelSelfShipAppointmentInternalServerError {
	return &CancelSelfShipAppointmentInternalServerError{}
}

/*
CancelSelfShipAppointmentInternalServerError describes a response with status code 500, with default header values.

An unexpected condition occurred that prevented the server from fulfilling the request.
*/
type CancelSelfShipAppointmentInternalServerError struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inboundv2024_03_20_models.ErrorList
}

// IsSuccess returns true when this cancel self ship appointment internal server error response has a 2xx status code
func (o *CancelSelfShipAppointmentInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cancel self ship appointment internal server error response has a 3xx status code
func (o *CancelSelfShipAppointmentInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cancel self ship appointment internal server error response has a 4xx status code
func (o *CancelSelfShipAppointmentInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this cancel self ship appointment internal server error response has a 5xx status code
func (o *CancelSelfShipAppointmentInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this cancel self ship appointment internal server error response a status code equal to that given
func (o *CancelSelfShipAppointmentInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *CancelSelfShipAppointmentInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/selfShipAppointmentSlots/{slotId}/cancellation][%d] cancelSelfShipAppointmentInternalServerError  %+v", 500, o.Payload)
}

func (o *CancelSelfShipAppointmentInternalServerError) String() string {
	return fmt.Sprintf("[PUT /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/selfShipAppointmentSlots/{slotId}/cancellation][%d] cancelSelfShipAppointmentInternalServerError  %+v", 500, o.Payload)
}

func (o *CancelSelfShipAppointmentInternalServerError) GetPayload() *fulfillment_inboundv2024_03_20_models.ErrorList {
	return o.Payload
}

func (o *CancelSelfShipAppointmentInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCancelSelfShipAppointmentServiceUnavailable creates a CancelSelfShipAppointmentServiceUnavailable with default headers values
func NewCancelSelfShipAppointmentServiceUnavailable() *CancelSelfShipAppointmentServiceUnavailable {
	return &CancelSelfShipAppointmentServiceUnavailable{}
}

/*
CancelSelfShipAppointmentServiceUnavailable describes a response with status code 503, with default header values.

Temporary overloading or maintenance of the server.
*/
type CancelSelfShipAppointmentServiceUnavailable struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inboundv2024_03_20_models.ErrorList
}

// IsSuccess returns true when this cancel self ship appointment service unavailable response has a 2xx status code
func (o *CancelSelfShipAppointmentServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cancel self ship appointment service unavailable response has a 3xx status code
func (o *CancelSelfShipAppointmentServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cancel self ship appointment service unavailable response has a 4xx status code
func (o *CancelSelfShipAppointmentServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this cancel self ship appointment service unavailable response has a 5xx status code
func (o *CancelSelfShipAppointmentServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this cancel self ship appointment service unavailable response a status code equal to that given
func (o *CancelSelfShipAppointmentServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *CancelSelfShipAppointmentServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/selfShipAppointmentSlots/{slotId}/cancellation][%d] cancelSelfShipAppointmentServiceUnavailable  %+v", 503, o.Payload)
}

func (o *CancelSelfShipAppointmentServiceUnavailable) String() string {
	return fmt.Sprintf("[PUT /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/selfShipAppointmentSlots/{slotId}/cancellation][%d] cancelSelfShipAppointmentServiceUnavailable  %+v", 503, o.Payload)
}

func (o *CancelSelfShipAppointmentServiceUnavailable) GetPayload() *fulfillment_inboundv2024_03_20_models.ErrorList {
	return o.Payload
}

func (o *CancelSelfShipAppointmentServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
