// Code generated by go-swagger; DO NOT EDIT.

package vendor_shipping

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/xamandar/amzn-sp-api-go/api/vendorDirectFulfillmentShipping_2021-12-28/vendor_direct_fulfillment_shipping_2021_12_28_models"
)

// SubmitShipmentStatusUpdatesReader is a Reader for the SubmitShipmentStatusUpdates structure.
type SubmitShipmentStatusUpdatesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SubmitShipmentStatusUpdatesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewSubmitShipmentStatusUpdatesAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSubmitShipmentStatusUpdatesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSubmitShipmentStatusUpdatesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSubmitShipmentStatusUpdatesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewSubmitShipmentStatusUpdatesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewSubmitShipmentStatusUpdatesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewSubmitShipmentStatusUpdatesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSubmitShipmentStatusUpdatesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewSubmitShipmentStatusUpdatesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSubmitShipmentStatusUpdatesAccepted creates a SubmitShipmentStatusUpdatesAccepted with default headers values
func NewSubmitShipmentStatusUpdatesAccepted() *SubmitShipmentStatusUpdatesAccepted {
	return &SubmitShipmentStatusUpdatesAccepted{}
}

/*
SubmitShipmentStatusUpdatesAccepted describes a response with status code 202, with default header values.

Success.
*/
type SubmitShipmentStatusUpdatesAccepted struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_shipping_2021_12_28_models.TransactionReference
}

// IsSuccess returns true when this submit shipment status updates accepted response has a 2xx status code
func (o *SubmitShipmentStatusUpdatesAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this submit shipment status updates accepted response has a 3xx status code
func (o *SubmitShipmentStatusUpdatesAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this submit shipment status updates accepted response has a 4xx status code
func (o *SubmitShipmentStatusUpdatesAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this submit shipment status updates accepted response has a 5xx status code
func (o *SubmitShipmentStatusUpdatesAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this submit shipment status updates accepted response a status code equal to that given
func (o *SubmitShipmentStatusUpdatesAccepted) IsCode(code int) bool {
	return code == 202
}

func (o *SubmitShipmentStatusUpdatesAccepted) Error() string {
	return fmt.Sprintf("[POST /vendor/directFulfillment/shipping/2021-12-28/shipmentStatusUpdates][%d] submitShipmentStatusUpdatesAccepted  %+v", 202, o.Payload)
}

func (o *SubmitShipmentStatusUpdatesAccepted) String() string {
	return fmt.Sprintf("[POST /vendor/directFulfillment/shipping/2021-12-28/shipmentStatusUpdates][%d] submitShipmentStatusUpdatesAccepted  %+v", 202, o.Payload)
}

func (o *SubmitShipmentStatusUpdatesAccepted) GetPayload() *vendor_direct_fulfillment_shipping_2021_12_28_models.TransactionReference {
	return o.Payload
}

func (o *SubmitShipmentStatusUpdatesAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(vendor_direct_fulfillment_shipping_2021_12_28_models.TransactionReference)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSubmitShipmentStatusUpdatesBadRequest creates a SubmitShipmentStatusUpdatesBadRequest with default headers values
func NewSubmitShipmentStatusUpdatesBadRequest() *SubmitShipmentStatusUpdatesBadRequest {
	return &SubmitShipmentStatusUpdatesBadRequest{}
}

/*
SubmitShipmentStatusUpdatesBadRequest describes a response with status code 400, with default header values.

Request has missing or invalid parameters and cannot be parsed.
*/
type SubmitShipmentStatusUpdatesBadRequest struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_shipping_2021_12_28_models.ErrorList
}

// IsSuccess returns true when this submit shipment status updates bad request response has a 2xx status code
func (o *SubmitShipmentStatusUpdatesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this submit shipment status updates bad request response has a 3xx status code
func (o *SubmitShipmentStatusUpdatesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this submit shipment status updates bad request response has a 4xx status code
func (o *SubmitShipmentStatusUpdatesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this submit shipment status updates bad request response has a 5xx status code
func (o *SubmitShipmentStatusUpdatesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this submit shipment status updates bad request response a status code equal to that given
func (o *SubmitShipmentStatusUpdatesBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *SubmitShipmentStatusUpdatesBadRequest) Error() string {
	return fmt.Sprintf("[POST /vendor/directFulfillment/shipping/2021-12-28/shipmentStatusUpdates][%d] submitShipmentStatusUpdatesBadRequest  %+v", 400, o.Payload)
}

func (o *SubmitShipmentStatusUpdatesBadRequest) String() string {
	return fmt.Sprintf("[POST /vendor/directFulfillment/shipping/2021-12-28/shipmentStatusUpdates][%d] submitShipmentStatusUpdatesBadRequest  %+v", 400, o.Payload)
}

func (o *SubmitShipmentStatusUpdatesBadRequest) GetPayload() *vendor_direct_fulfillment_shipping_2021_12_28_models.ErrorList {
	return o.Payload
}

func (o *SubmitShipmentStatusUpdatesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(vendor_direct_fulfillment_shipping_2021_12_28_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSubmitShipmentStatusUpdatesForbidden creates a SubmitShipmentStatusUpdatesForbidden with default headers values
func NewSubmitShipmentStatusUpdatesForbidden() *SubmitShipmentStatusUpdatesForbidden {
	return &SubmitShipmentStatusUpdatesForbidden{}
}

/*
SubmitShipmentStatusUpdatesForbidden describes a response with status code 403, with default header values.

Indicates that access to the resource is forbidden. Possible reasons include Access Denied, Unauthorized, Expired Token, or Invalid Signature.
*/
type SubmitShipmentStatusUpdatesForbidden struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_shipping_2021_12_28_models.ErrorList
}

// IsSuccess returns true when this submit shipment status updates forbidden response has a 2xx status code
func (o *SubmitShipmentStatusUpdatesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this submit shipment status updates forbidden response has a 3xx status code
func (o *SubmitShipmentStatusUpdatesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this submit shipment status updates forbidden response has a 4xx status code
func (o *SubmitShipmentStatusUpdatesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this submit shipment status updates forbidden response has a 5xx status code
func (o *SubmitShipmentStatusUpdatesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this submit shipment status updates forbidden response a status code equal to that given
func (o *SubmitShipmentStatusUpdatesForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *SubmitShipmentStatusUpdatesForbidden) Error() string {
	return fmt.Sprintf("[POST /vendor/directFulfillment/shipping/2021-12-28/shipmentStatusUpdates][%d] submitShipmentStatusUpdatesForbidden  %+v", 403, o.Payload)
}

func (o *SubmitShipmentStatusUpdatesForbidden) String() string {
	return fmt.Sprintf("[POST /vendor/directFulfillment/shipping/2021-12-28/shipmentStatusUpdates][%d] submitShipmentStatusUpdatesForbidden  %+v", 403, o.Payload)
}

func (o *SubmitShipmentStatusUpdatesForbidden) GetPayload() *vendor_direct_fulfillment_shipping_2021_12_28_models.ErrorList {
	return o.Payload
}

func (o *SubmitShipmentStatusUpdatesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(vendor_direct_fulfillment_shipping_2021_12_28_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSubmitShipmentStatusUpdatesNotFound creates a SubmitShipmentStatusUpdatesNotFound with default headers values
func NewSubmitShipmentStatusUpdatesNotFound() *SubmitShipmentStatusUpdatesNotFound {
	return &SubmitShipmentStatusUpdatesNotFound{}
}

/*
SubmitShipmentStatusUpdatesNotFound describes a response with status code 404, with default header values.

The resource specified does not exist.
*/
type SubmitShipmentStatusUpdatesNotFound struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_shipping_2021_12_28_models.ErrorList
}

// IsSuccess returns true when this submit shipment status updates not found response has a 2xx status code
func (o *SubmitShipmentStatusUpdatesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this submit shipment status updates not found response has a 3xx status code
func (o *SubmitShipmentStatusUpdatesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this submit shipment status updates not found response has a 4xx status code
func (o *SubmitShipmentStatusUpdatesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this submit shipment status updates not found response has a 5xx status code
func (o *SubmitShipmentStatusUpdatesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this submit shipment status updates not found response a status code equal to that given
func (o *SubmitShipmentStatusUpdatesNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *SubmitShipmentStatusUpdatesNotFound) Error() string {
	return fmt.Sprintf("[POST /vendor/directFulfillment/shipping/2021-12-28/shipmentStatusUpdates][%d] submitShipmentStatusUpdatesNotFound  %+v", 404, o.Payload)
}

func (o *SubmitShipmentStatusUpdatesNotFound) String() string {
	return fmt.Sprintf("[POST /vendor/directFulfillment/shipping/2021-12-28/shipmentStatusUpdates][%d] submitShipmentStatusUpdatesNotFound  %+v", 404, o.Payload)
}

func (o *SubmitShipmentStatusUpdatesNotFound) GetPayload() *vendor_direct_fulfillment_shipping_2021_12_28_models.ErrorList {
	return o.Payload
}

func (o *SubmitShipmentStatusUpdatesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(vendor_direct_fulfillment_shipping_2021_12_28_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSubmitShipmentStatusUpdatesRequestEntityTooLarge creates a SubmitShipmentStatusUpdatesRequestEntityTooLarge with default headers values
func NewSubmitShipmentStatusUpdatesRequestEntityTooLarge() *SubmitShipmentStatusUpdatesRequestEntityTooLarge {
	return &SubmitShipmentStatusUpdatesRequestEntityTooLarge{}
}

/*
SubmitShipmentStatusUpdatesRequestEntityTooLarge describes a response with status code 413, with default header values.

The request size exceeded the maximum accepted size.
*/
type SubmitShipmentStatusUpdatesRequestEntityTooLarge struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_shipping_2021_12_28_models.ErrorList
}

// IsSuccess returns true when this submit shipment status updates request entity too large response has a 2xx status code
func (o *SubmitShipmentStatusUpdatesRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this submit shipment status updates request entity too large response has a 3xx status code
func (o *SubmitShipmentStatusUpdatesRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this submit shipment status updates request entity too large response has a 4xx status code
func (o *SubmitShipmentStatusUpdatesRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this submit shipment status updates request entity too large response has a 5xx status code
func (o *SubmitShipmentStatusUpdatesRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this submit shipment status updates request entity too large response a status code equal to that given
func (o *SubmitShipmentStatusUpdatesRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *SubmitShipmentStatusUpdatesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /vendor/directFulfillment/shipping/2021-12-28/shipmentStatusUpdates][%d] submitShipmentStatusUpdatesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *SubmitShipmentStatusUpdatesRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[POST /vendor/directFulfillment/shipping/2021-12-28/shipmentStatusUpdates][%d] submitShipmentStatusUpdatesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *SubmitShipmentStatusUpdatesRequestEntityTooLarge) GetPayload() *vendor_direct_fulfillment_shipping_2021_12_28_models.ErrorList {
	return o.Payload
}

func (o *SubmitShipmentStatusUpdatesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(vendor_direct_fulfillment_shipping_2021_12_28_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSubmitShipmentStatusUpdatesUnsupportedMediaType creates a SubmitShipmentStatusUpdatesUnsupportedMediaType with default headers values
func NewSubmitShipmentStatusUpdatesUnsupportedMediaType() *SubmitShipmentStatusUpdatesUnsupportedMediaType {
	return &SubmitShipmentStatusUpdatesUnsupportedMediaType{}
}

/*
SubmitShipmentStatusUpdatesUnsupportedMediaType describes a response with status code 415, with default header values.

The request payload is in an unsupported format.
*/
type SubmitShipmentStatusUpdatesUnsupportedMediaType struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_shipping_2021_12_28_models.ErrorList
}

// IsSuccess returns true when this submit shipment status updates unsupported media type response has a 2xx status code
func (o *SubmitShipmentStatusUpdatesUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this submit shipment status updates unsupported media type response has a 3xx status code
func (o *SubmitShipmentStatusUpdatesUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this submit shipment status updates unsupported media type response has a 4xx status code
func (o *SubmitShipmentStatusUpdatesUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this submit shipment status updates unsupported media type response has a 5xx status code
func (o *SubmitShipmentStatusUpdatesUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this submit shipment status updates unsupported media type response a status code equal to that given
func (o *SubmitShipmentStatusUpdatesUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *SubmitShipmentStatusUpdatesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /vendor/directFulfillment/shipping/2021-12-28/shipmentStatusUpdates][%d] submitShipmentStatusUpdatesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *SubmitShipmentStatusUpdatesUnsupportedMediaType) String() string {
	return fmt.Sprintf("[POST /vendor/directFulfillment/shipping/2021-12-28/shipmentStatusUpdates][%d] submitShipmentStatusUpdatesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *SubmitShipmentStatusUpdatesUnsupportedMediaType) GetPayload() *vendor_direct_fulfillment_shipping_2021_12_28_models.ErrorList {
	return o.Payload
}

func (o *SubmitShipmentStatusUpdatesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(vendor_direct_fulfillment_shipping_2021_12_28_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSubmitShipmentStatusUpdatesTooManyRequests creates a SubmitShipmentStatusUpdatesTooManyRequests with default headers values
func NewSubmitShipmentStatusUpdatesTooManyRequests() *SubmitShipmentStatusUpdatesTooManyRequests {
	return &SubmitShipmentStatusUpdatesTooManyRequests{}
}

/*
SubmitShipmentStatusUpdatesTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type SubmitShipmentStatusUpdatesTooManyRequests struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_shipping_2021_12_28_models.ErrorList
}

// IsSuccess returns true when this submit shipment status updates too many requests response has a 2xx status code
func (o *SubmitShipmentStatusUpdatesTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this submit shipment status updates too many requests response has a 3xx status code
func (o *SubmitShipmentStatusUpdatesTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this submit shipment status updates too many requests response has a 4xx status code
func (o *SubmitShipmentStatusUpdatesTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this submit shipment status updates too many requests response has a 5xx status code
func (o *SubmitShipmentStatusUpdatesTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this submit shipment status updates too many requests response a status code equal to that given
func (o *SubmitShipmentStatusUpdatesTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *SubmitShipmentStatusUpdatesTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /vendor/directFulfillment/shipping/2021-12-28/shipmentStatusUpdates][%d] submitShipmentStatusUpdatesTooManyRequests  %+v", 429, o.Payload)
}

func (o *SubmitShipmentStatusUpdatesTooManyRequests) String() string {
	return fmt.Sprintf("[POST /vendor/directFulfillment/shipping/2021-12-28/shipmentStatusUpdates][%d] submitShipmentStatusUpdatesTooManyRequests  %+v", 429, o.Payload)
}

func (o *SubmitShipmentStatusUpdatesTooManyRequests) GetPayload() *vendor_direct_fulfillment_shipping_2021_12_28_models.ErrorList {
	return o.Payload
}

func (o *SubmitShipmentStatusUpdatesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(vendor_direct_fulfillment_shipping_2021_12_28_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSubmitShipmentStatusUpdatesInternalServerError creates a SubmitShipmentStatusUpdatesInternalServerError with default headers values
func NewSubmitShipmentStatusUpdatesInternalServerError() *SubmitShipmentStatusUpdatesInternalServerError {
	return &SubmitShipmentStatusUpdatesInternalServerError{}
}

/*
SubmitShipmentStatusUpdatesInternalServerError describes a response with status code 500, with default header values.

An unexpected condition occurred that prevented the server from fulfilling the request.
*/
type SubmitShipmentStatusUpdatesInternalServerError struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_shipping_2021_12_28_models.ErrorList
}

// IsSuccess returns true when this submit shipment status updates internal server error response has a 2xx status code
func (o *SubmitShipmentStatusUpdatesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this submit shipment status updates internal server error response has a 3xx status code
func (o *SubmitShipmentStatusUpdatesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this submit shipment status updates internal server error response has a 4xx status code
func (o *SubmitShipmentStatusUpdatesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this submit shipment status updates internal server error response has a 5xx status code
func (o *SubmitShipmentStatusUpdatesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this submit shipment status updates internal server error response a status code equal to that given
func (o *SubmitShipmentStatusUpdatesInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *SubmitShipmentStatusUpdatesInternalServerError) Error() string {
	return fmt.Sprintf("[POST /vendor/directFulfillment/shipping/2021-12-28/shipmentStatusUpdates][%d] submitShipmentStatusUpdatesInternalServerError  %+v", 500, o.Payload)
}

func (o *SubmitShipmentStatusUpdatesInternalServerError) String() string {
	return fmt.Sprintf("[POST /vendor/directFulfillment/shipping/2021-12-28/shipmentStatusUpdates][%d] submitShipmentStatusUpdatesInternalServerError  %+v", 500, o.Payload)
}

func (o *SubmitShipmentStatusUpdatesInternalServerError) GetPayload() *vendor_direct_fulfillment_shipping_2021_12_28_models.ErrorList {
	return o.Payload
}

func (o *SubmitShipmentStatusUpdatesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(vendor_direct_fulfillment_shipping_2021_12_28_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSubmitShipmentStatusUpdatesServiceUnavailable creates a SubmitShipmentStatusUpdatesServiceUnavailable with default headers values
func NewSubmitShipmentStatusUpdatesServiceUnavailable() *SubmitShipmentStatusUpdatesServiceUnavailable {
	return &SubmitShipmentStatusUpdatesServiceUnavailable{}
}

/*
SubmitShipmentStatusUpdatesServiceUnavailable describes a response with status code 503, with default header values.

Temporary overloading or maintenance of the server.
*/
type SubmitShipmentStatusUpdatesServiceUnavailable struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_shipping_2021_12_28_models.ErrorList
}

// IsSuccess returns true when this submit shipment status updates service unavailable response has a 2xx status code
func (o *SubmitShipmentStatusUpdatesServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this submit shipment status updates service unavailable response has a 3xx status code
func (o *SubmitShipmentStatusUpdatesServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this submit shipment status updates service unavailable response has a 4xx status code
func (o *SubmitShipmentStatusUpdatesServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this submit shipment status updates service unavailable response has a 5xx status code
func (o *SubmitShipmentStatusUpdatesServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this submit shipment status updates service unavailable response a status code equal to that given
func (o *SubmitShipmentStatusUpdatesServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *SubmitShipmentStatusUpdatesServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /vendor/directFulfillment/shipping/2021-12-28/shipmentStatusUpdates][%d] submitShipmentStatusUpdatesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *SubmitShipmentStatusUpdatesServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /vendor/directFulfillment/shipping/2021-12-28/shipmentStatusUpdates][%d] submitShipmentStatusUpdatesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *SubmitShipmentStatusUpdatesServiceUnavailable) GetPayload() *vendor_direct_fulfillment_shipping_2021_12_28_models.ErrorList {
	return o.Payload
}

func (o *SubmitShipmentStatusUpdatesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(vendor_direct_fulfillment_shipping_2021_12_28_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
