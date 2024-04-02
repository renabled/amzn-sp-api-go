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

// UpdateShipmentTrackingDetailsReader is a Reader for the UpdateShipmentTrackingDetails structure.
type UpdateShipmentTrackingDetailsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateShipmentTrackingDetailsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewUpdateShipmentTrackingDetailsAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateShipmentTrackingDetailsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateShipmentTrackingDetailsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateShipmentTrackingDetailsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewUpdateShipmentTrackingDetailsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewUpdateShipmentTrackingDetailsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewUpdateShipmentTrackingDetailsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateShipmentTrackingDetailsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewUpdateShipmentTrackingDetailsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateShipmentTrackingDetailsAccepted creates a UpdateShipmentTrackingDetailsAccepted with default headers values
func NewUpdateShipmentTrackingDetailsAccepted() *UpdateShipmentTrackingDetailsAccepted {
	return &UpdateShipmentTrackingDetailsAccepted{}
}

/*
UpdateShipmentTrackingDetailsAccepted describes a response with status code 202, with default header values.

UpdateShipmentTrackingDetails 202 response
*/
type UpdateShipmentTrackingDetailsAccepted struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inboundv2024_03_20_models.UpdateShipmentTrackingDetailsResponse
}

// IsSuccess returns true when this update shipment tracking details accepted response has a 2xx status code
func (o *UpdateShipmentTrackingDetailsAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update shipment tracking details accepted response has a 3xx status code
func (o *UpdateShipmentTrackingDetailsAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update shipment tracking details accepted response has a 4xx status code
func (o *UpdateShipmentTrackingDetailsAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this update shipment tracking details accepted response has a 5xx status code
func (o *UpdateShipmentTrackingDetailsAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this update shipment tracking details accepted response a status code equal to that given
func (o *UpdateShipmentTrackingDetailsAccepted) IsCode(code int) bool {
	return code == 202
}

func (o *UpdateShipmentTrackingDetailsAccepted) Error() string {
	return fmt.Sprintf("[PUT /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/trackingDetails][%d] updateShipmentTrackingDetailsAccepted  %+v", 202, o.Payload)
}

func (o *UpdateShipmentTrackingDetailsAccepted) String() string {
	return fmt.Sprintf("[PUT /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/trackingDetails][%d] updateShipmentTrackingDetailsAccepted  %+v", 202, o.Payload)
}

func (o *UpdateShipmentTrackingDetailsAccepted) GetPayload() *fulfillment_inboundv2024_03_20_models.UpdateShipmentTrackingDetailsResponse {
	return o.Payload
}

func (o *UpdateShipmentTrackingDetailsAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fulfillment_inboundv2024_03_20_models.UpdateShipmentTrackingDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateShipmentTrackingDetailsBadRequest creates a UpdateShipmentTrackingDetailsBadRequest with default headers values
func NewUpdateShipmentTrackingDetailsBadRequest() *UpdateShipmentTrackingDetailsBadRequest {
	return &UpdateShipmentTrackingDetailsBadRequest{}
}

/*
UpdateShipmentTrackingDetailsBadRequest describes a response with status code 400, with default header values.

Request has missing or invalid parameters and cannot be parsed.
*/
type UpdateShipmentTrackingDetailsBadRequest struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inboundv2024_03_20_models.ErrorList
}

// IsSuccess returns true when this update shipment tracking details bad request response has a 2xx status code
func (o *UpdateShipmentTrackingDetailsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update shipment tracking details bad request response has a 3xx status code
func (o *UpdateShipmentTrackingDetailsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update shipment tracking details bad request response has a 4xx status code
func (o *UpdateShipmentTrackingDetailsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update shipment tracking details bad request response has a 5xx status code
func (o *UpdateShipmentTrackingDetailsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update shipment tracking details bad request response a status code equal to that given
func (o *UpdateShipmentTrackingDetailsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *UpdateShipmentTrackingDetailsBadRequest) Error() string {
	return fmt.Sprintf("[PUT /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/trackingDetails][%d] updateShipmentTrackingDetailsBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateShipmentTrackingDetailsBadRequest) String() string {
	return fmt.Sprintf("[PUT /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/trackingDetails][%d] updateShipmentTrackingDetailsBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateShipmentTrackingDetailsBadRequest) GetPayload() *fulfillment_inboundv2024_03_20_models.ErrorList {
	return o.Payload
}

func (o *UpdateShipmentTrackingDetailsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateShipmentTrackingDetailsForbidden creates a UpdateShipmentTrackingDetailsForbidden with default headers values
func NewUpdateShipmentTrackingDetailsForbidden() *UpdateShipmentTrackingDetailsForbidden {
	return &UpdateShipmentTrackingDetailsForbidden{}
}

/*
UpdateShipmentTrackingDetailsForbidden describes a response with status code 403, with default header values.

Indicates that access to the resource is forbidden. Possible reasons include Access Denied, Unauthorized, Expired Token, or Invalid Signature.
*/
type UpdateShipmentTrackingDetailsForbidden struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inboundv2024_03_20_models.ErrorList
}

// IsSuccess returns true when this update shipment tracking details forbidden response has a 2xx status code
func (o *UpdateShipmentTrackingDetailsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update shipment tracking details forbidden response has a 3xx status code
func (o *UpdateShipmentTrackingDetailsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update shipment tracking details forbidden response has a 4xx status code
func (o *UpdateShipmentTrackingDetailsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update shipment tracking details forbidden response has a 5xx status code
func (o *UpdateShipmentTrackingDetailsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update shipment tracking details forbidden response a status code equal to that given
func (o *UpdateShipmentTrackingDetailsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *UpdateShipmentTrackingDetailsForbidden) Error() string {
	return fmt.Sprintf("[PUT /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/trackingDetails][%d] updateShipmentTrackingDetailsForbidden  %+v", 403, o.Payload)
}

func (o *UpdateShipmentTrackingDetailsForbidden) String() string {
	return fmt.Sprintf("[PUT /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/trackingDetails][%d] updateShipmentTrackingDetailsForbidden  %+v", 403, o.Payload)
}

func (o *UpdateShipmentTrackingDetailsForbidden) GetPayload() *fulfillment_inboundv2024_03_20_models.ErrorList {
	return o.Payload
}

func (o *UpdateShipmentTrackingDetailsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateShipmentTrackingDetailsNotFound creates a UpdateShipmentTrackingDetailsNotFound with default headers values
func NewUpdateShipmentTrackingDetailsNotFound() *UpdateShipmentTrackingDetailsNotFound {
	return &UpdateShipmentTrackingDetailsNotFound{}
}

/*
UpdateShipmentTrackingDetailsNotFound describes a response with status code 404, with default header values.

The resource specified does not exist.
*/
type UpdateShipmentTrackingDetailsNotFound struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inboundv2024_03_20_models.ErrorList
}

// IsSuccess returns true when this update shipment tracking details not found response has a 2xx status code
func (o *UpdateShipmentTrackingDetailsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update shipment tracking details not found response has a 3xx status code
func (o *UpdateShipmentTrackingDetailsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update shipment tracking details not found response has a 4xx status code
func (o *UpdateShipmentTrackingDetailsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update shipment tracking details not found response has a 5xx status code
func (o *UpdateShipmentTrackingDetailsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update shipment tracking details not found response a status code equal to that given
func (o *UpdateShipmentTrackingDetailsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *UpdateShipmentTrackingDetailsNotFound) Error() string {
	return fmt.Sprintf("[PUT /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/trackingDetails][%d] updateShipmentTrackingDetailsNotFound  %+v", 404, o.Payload)
}

func (o *UpdateShipmentTrackingDetailsNotFound) String() string {
	return fmt.Sprintf("[PUT /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/trackingDetails][%d] updateShipmentTrackingDetailsNotFound  %+v", 404, o.Payload)
}

func (o *UpdateShipmentTrackingDetailsNotFound) GetPayload() *fulfillment_inboundv2024_03_20_models.ErrorList {
	return o.Payload
}

func (o *UpdateShipmentTrackingDetailsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateShipmentTrackingDetailsRequestEntityTooLarge creates a UpdateShipmentTrackingDetailsRequestEntityTooLarge with default headers values
func NewUpdateShipmentTrackingDetailsRequestEntityTooLarge() *UpdateShipmentTrackingDetailsRequestEntityTooLarge {
	return &UpdateShipmentTrackingDetailsRequestEntityTooLarge{}
}

/*
UpdateShipmentTrackingDetailsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request size exceeded the maximum accepted size.
*/
type UpdateShipmentTrackingDetailsRequestEntityTooLarge struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inboundv2024_03_20_models.ErrorList
}

// IsSuccess returns true when this update shipment tracking details request entity too large response has a 2xx status code
func (o *UpdateShipmentTrackingDetailsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update shipment tracking details request entity too large response has a 3xx status code
func (o *UpdateShipmentTrackingDetailsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update shipment tracking details request entity too large response has a 4xx status code
func (o *UpdateShipmentTrackingDetailsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this update shipment tracking details request entity too large response has a 5xx status code
func (o *UpdateShipmentTrackingDetailsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this update shipment tracking details request entity too large response a status code equal to that given
func (o *UpdateShipmentTrackingDetailsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *UpdateShipmentTrackingDetailsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PUT /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/trackingDetails][%d] updateShipmentTrackingDetailsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *UpdateShipmentTrackingDetailsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[PUT /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/trackingDetails][%d] updateShipmentTrackingDetailsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *UpdateShipmentTrackingDetailsRequestEntityTooLarge) GetPayload() *fulfillment_inboundv2024_03_20_models.ErrorList {
	return o.Payload
}

func (o *UpdateShipmentTrackingDetailsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateShipmentTrackingDetailsUnsupportedMediaType creates a UpdateShipmentTrackingDetailsUnsupportedMediaType with default headers values
func NewUpdateShipmentTrackingDetailsUnsupportedMediaType() *UpdateShipmentTrackingDetailsUnsupportedMediaType {
	return &UpdateShipmentTrackingDetailsUnsupportedMediaType{}
}

/*
UpdateShipmentTrackingDetailsUnsupportedMediaType describes a response with status code 415, with default header values.

The request payload is in an unsupported format.
*/
type UpdateShipmentTrackingDetailsUnsupportedMediaType struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inboundv2024_03_20_models.ErrorList
}

// IsSuccess returns true when this update shipment tracking details unsupported media type response has a 2xx status code
func (o *UpdateShipmentTrackingDetailsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update shipment tracking details unsupported media type response has a 3xx status code
func (o *UpdateShipmentTrackingDetailsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update shipment tracking details unsupported media type response has a 4xx status code
func (o *UpdateShipmentTrackingDetailsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this update shipment tracking details unsupported media type response has a 5xx status code
func (o *UpdateShipmentTrackingDetailsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this update shipment tracking details unsupported media type response a status code equal to that given
func (o *UpdateShipmentTrackingDetailsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *UpdateShipmentTrackingDetailsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PUT /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/trackingDetails][%d] updateShipmentTrackingDetailsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *UpdateShipmentTrackingDetailsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[PUT /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/trackingDetails][%d] updateShipmentTrackingDetailsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *UpdateShipmentTrackingDetailsUnsupportedMediaType) GetPayload() *fulfillment_inboundv2024_03_20_models.ErrorList {
	return o.Payload
}

func (o *UpdateShipmentTrackingDetailsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateShipmentTrackingDetailsTooManyRequests creates a UpdateShipmentTrackingDetailsTooManyRequests with default headers values
func NewUpdateShipmentTrackingDetailsTooManyRequests() *UpdateShipmentTrackingDetailsTooManyRequests {
	return &UpdateShipmentTrackingDetailsTooManyRequests{}
}

/*
UpdateShipmentTrackingDetailsTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type UpdateShipmentTrackingDetailsTooManyRequests struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inboundv2024_03_20_models.ErrorList
}

// IsSuccess returns true when this update shipment tracking details too many requests response has a 2xx status code
func (o *UpdateShipmentTrackingDetailsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update shipment tracking details too many requests response has a 3xx status code
func (o *UpdateShipmentTrackingDetailsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update shipment tracking details too many requests response has a 4xx status code
func (o *UpdateShipmentTrackingDetailsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this update shipment tracking details too many requests response has a 5xx status code
func (o *UpdateShipmentTrackingDetailsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this update shipment tracking details too many requests response a status code equal to that given
func (o *UpdateShipmentTrackingDetailsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *UpdateShipmentTrackingDetailsTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/trackingDetails][%d] updateShipmentTrackingDetailsTooManyRequests  %+v", 429, o.Payload)
}

func (o *UpdateShipmentTrackingDetailsTooManyRequests) String() string {
	return fmt.Sprintf("[PUT /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/trackingDetails][%d] updateShipmentTrackingDetailsTooManyRequests  %+v", 429, o.Payload)
}

func (o *UpdateShipmentTrackingDetailsTooManyRequests) GetPayload() *fulfillment_inboundv2024_03_20_models.ErrorList {
	return o.Payload
}

func (o *UpdateShipmentTrackingDetailsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateShipmentTrackingDetailsInternalServerError creates a UpdateShipmentTrackingDetailsInternalServerError with default headers values
func NewUpdateShipmentTrackingDetailsInternalServerError() *UpdateShipmentTrackingDetailsInternalServerError {
	return &UpdateShipmentTrackingDetailsInternalServerError{}
}

/*
UpdateShipmentTrackingDetailsInternalServerError describes a response with status code 500, with default header values.

An unexpected condition occurred that prevented the server from fulfilling the request.
*/
type UpdateShipmentTrackingDetailsInternalServerError struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inboundv2024_03_20_models.ErrorList
}

// IsSuccess returns true when this update shipment tracking details internal server error response has a 2xx status code
func (o *UpdateShipmentTrackingDetailsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update shipment tracking details internal server error response has a 3xx status code
func (o *UpdateShipmentTrackingDetailsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update shipment tracking details internal server error response has a 4xx status code
func (o *UpdateShipmentTrackingDetailsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this update shipment tracking details internal server error response has a 5xx status code
func (o *UpdateShipmentTrackingDetailsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this update shipment tracking details internal server error response a status code equal to that given
func (o *UpdateShipmentTrackingDetailsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *UpdateShipmentTrackingDetailsInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/trackingDetails][%d] updateShipmentTrackingDetailsInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateShipmentTrackingDetailsInternalServerError) String() string {
	return fmt.Sprintf("[PUT /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/trackingDetails][%d] updateShipmentTrackingDetailsInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateShipmentTrackingDetailsInternalServerError) GetPayload() *fulfillment_inboundv2024_03_20_models.ErrorList {
	return o.Payload
}

func (o *UpdateShipmentTrackingDetailsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateShipmentTrackingDetailsServiceUnavailable creates a UpdateShipmentTrackingDetailsServiceUnavailable with default headers values
func NewUpdateShipmentTrackingDetailsServiceUnavailable() *UpdateShipmentTrackingDetailsServiceUnavailable {
	return &UpdateShipmentTrackingDetailsServiceUnavailable{}
}

/*
UpdateShipmentTrackingDetailsServiceUnavailable describes a response with status code 503, with default header values.

Temporary overloading or maintenance of the server.
*/
type UpdateShipmentTrackingDetailsServiceUnavailable struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inboundv2024_03_20_models.ErrorList
}

// IsSuccess returns true when this update shipment tracking details service unavailable response has a 2xx status code
func (o *UpdateShipmentTrackingDetailsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update shipment tracking details service unavailable response has a 3xx status code
func (o *UpdateShipmentTrackingDetailsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update shipment tracking details service unavailable response has a 4xx status code
func (o *UpdateShipmentTrackingDetailsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this update shipment tracking details service unavailable response has a 5xx status code
func (o *UpdateShipmentTrackingDetailsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this update shipment tracking details service unavailable response a status code equal to that given
func (o *UpdateShipmentTrackingDetailsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *UpdateShipmentTrackingDetailsServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/trackingDetails][%d] updateShipmentTrackingDetailsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *UpdateShipmentTrackingDetailsServiceUnavailable) String() string {
	return fmt.Sprintf("[PUT /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/trackingDetails][%d] updateShipmentTrackingDetailsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *UpdateShipmentTrackingDetailsServiceUnavailable) GetPayload() *fulfillment_inboundv2024_03_20_models.ErrorList {
	return o.Payload
}

func (o *UpdateShipmentTrackingDetailsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
