// Code generated by go-swagger; DO NOT EDIT.

package vendor_shipping

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/renabled/amzn-sp-api-go/api/vendorDirectFulfillmentShippingV1/vendor_direct_fulfillment_shipping_v1_models"
)

// SubmitShipmentConfirmationsReader is a Reader for the SubmitShipmentConfirmations structure.
type SubmitShipmentConfirmationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SubmitShipmentConfirmationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewSubmitShipmentConfirmationsAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSubmitShipmentConfirmationsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSubmitShipmentConfirmationsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSubmitShipmentConfirmationsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewSubmitShipmentConfirmationsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewSubmitShipmentConfirmationsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewSubmitShipmentConfirmationsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSubmitShipmentConfirmationsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewSubmitShipmentConfirmationsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSubmitShipmentConfirmationsAccepted creates a SubmitShipmentConfirmationsAccepted with default headers values
func NewSubmitShipmentConfirmationsAccepted() *SubmitShipmentConfirmationsAccepted {
	return &SubmitShipmentConfirmationsAccepted{}
}

/*
SubmitShipmentConfirmationsAccepted describes a response with status code 202, with default header values.

Success.
*/
type SubmitShipmentConfirmationsAccepted struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_shipping_v1_models.SubmitShipmentConfirmationsResponse
}

// IsSuccess returns true when this submit shipment confirmations accepted response has a 2xx status code
func (o *SubmitShipmentConfirmationsAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this submit shipment confirmations accepted response has a 3xx status code
func (o *SubmitShipmentConfirmationsAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this submit shipment confirmations accepted response has a 4xx status code
func (o *SubmitShipmentConfirmationsAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this submit shipment confirmations accepted response has a 5xx status code
func (o *SubmitShipmentConfirmationsAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this submit shipment confirmations accepted response a status code equal to that given
func (o *SubmitShipmentConfirmationsAccepted) IsCode(code int) bool {
	return code == 202
}

func (o *SubmitShipmentConfirmationsAccepted) Error() string {
	return fmt.Sprintf("[POST /vendor/directFulfillment/shipping/v1/shipmentConfirmations][%d] submitShipmentConfirmationsAccepted  %+v", 202, o.Payload)
}

func (o *SubmitShipmentConfirmationsAccepted) String() string {
	return fmt.Sprintf("[POST /vendor/directFulfillment/shipping/v1/shipmentConfirmations][%d] submitShipmentConfirmationsAccepted  %+v", 202, o.Payload)
}

func (o *SubmitShipmentConfirmationsAccepted) GetPayload() *vendor_direct_fulfillment_shipping_v1_models.SubmitShipmentConfirmationsResponse {
	return o.Payload
}

func (o *SubmitShipmentConfirmationsAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(vendor_direct_fulfillment_shipping_v1_models.SubmitShipmentConfirmationsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSubmitShipmentConfirmationsBadRequest creates a SubmitShipmentConfirmationsBadRequest with default headers values
func NewSubmitShipmentConfirmationsBadRequest() *SubmitShipmentConfirmationsBadRequest {
	return &SubmitShipmentConfirmationsBadRequest{}
}

/*
SubmitShipmentConfirmationsBadRequest describes a response with status code 400, with default header values.

Request has missing or invalid parameters and cannot be parsed.
*/
type SubmitShipmentConfirmationsBadRequest struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_shipping_v1_models.SubmitShipmentConfirmationsResponse
}

// IsSuccess returns true when this submit shipment confirmations bad request response has a 2xx status code
func (o *SubmitShipmentConfirmationsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this submit shipment confirmations bad request response has a 3xx status code
func (o *SubmitShipmentConfirmationsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this submit shipment confirmations bad request response has a 4xx status code
func (o *SubmitShipmentConfirmationsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this submit shipment confirmations bad request response has a 5xx status code
func (o *SubmitShipmentConfirmationsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this submit shipment confirmations bad request response a status code equal to that given
func (o *SubmitShipmentConfirmationsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *SubmitShipmentConfirmationsBadRequest) Error() string {
	return fmt.Sprintf("[POST /vendor/directFulfillment/shipping/v1/shipmentConfirmations][%d] submitShipmentConfirmationsBadRequest  %+v", 400, o.Payload)
}

func (o *SubmitShipmentConfirmationsBadRequest) String() string {
	return fmt.Sprintf("[POST /vendor/directFulfillment/shipping/v1/shipmentConfirmations][%d] submitShipmentConfirmationsBadRequest  %+v", 400, o.Payload)
}

func (o *SubmitShipmentConfirmationsBadRequest) GetPayload() *vendor_direct_fulfillment_shipping_v1_models.SubmitShipmentConfirmationsResponse {
	return o.Payload
}

func (o *SubmitShipmentConfirmationsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(vendor_direct_fulfillment_shipping_v1_models.SubmitShipmentConfirmationsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSubmitShipmentConfirmationsForbidden creates a SubmitShipmentConfirmationsForbidden with default headers values
func NewSubmitShipmentConfirmationsForbidden() *SubmitShipmentConfirmationsForbidden {
	return &SubmitShipmentConfirmationsForbidden{}
}

/*
SubmitShipmentConfirmationsForbidden describes a response with status code 403, with default header values.

Indicates that access to the resource is forbidden. Possible reasons include Access Denied, Unauthorized, Expired Token, or Invalid Signature.
*/
type SubmitShipmentConfirmationsForbidden struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_shipping_v1_models.SubmitShipmentConfirmationsResponse
}

// IsSuccess returns true when this submit shipment confirmations forbidden response has a 2xx status code
func (o *SubmitShipmentConfirmationsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this submit shipment confirmations forbidden response has a 3xx status code
func (o *SubmitShipmentConfirmationsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this submit shipment confirmations forbidden response has a 4xx status code
func (o *SubmitShipmentConfirmationsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this submit shipment confirmations forbidden response has a 5xx status code
func (o *SubmitShipmentConfirmationsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this submit shipment confirmations forbidden response a status code equal to that given
func (o *SubmitShipmentConfirmationsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *SubmitShipmentConfirmationsForbidden) Error() string {
	return fmt.Sprintf("[POST /vendor/directFulfillment/shipping/v1/shipmentConfirmations][%d] submitShipmentConfirmationsForbidden  %+v", 403, o.Payload)
}

func (o *SubmitShipmentConfirmationsForbidden) String() string {
	return fmt.Sprintf("[POST /vendor/directFulfillment/shipping/v1/shipmentConfirmations][%d] submitShipmentConfirmationsForbidden  %+v", 403, o.Payload)
}

func (o *SubmitShipmentConfirmationsForbidden) GetPayload() *vendor_direct_fulfillment_shipping_v1_models.SubmitShipmentConfirmationsResponse {
	return o.Payload
}

func (o *SubmitShipmentConfirmationsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(vendor_direct_fulfillment_shipping_v1_models.SubmitShipmentConfirmationsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSubmitShipmentConfirmationsNotFound creates a SubmitShipmentConfirmationsNotFound with default headers values
func NewSubmitShipmentConfirmationsNotFound() *SubmitShipmentConfirmationsNotFound {
	return &SubmitShipmentConfirmationsNotFound{}
}

/*
SubmitShipmentConfirmationsNotFound describes a response with status code 404, with default header values.

The resource specified does not exist.
*/
type SubmitShipmentConfirmationsNotFound struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_shipping_v1_models.SubmitShipmentConfirmationsResponse
}

// IsSuccess returns true when this submit shipment confirmations not found response has a 2xx status code
func (o *SubmitShipmentConfirmationsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this submit shipment confirmations not found response has a 3xx status code
func (o *SubmitShipmentConfirmationsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this submit shipment confirmations not found response has a 4xx status code
func (o *SubmitShipmentConfirmationsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this submit shipment confirmations not found response has a 5xx status code
func (o *SubmitShipmentConfirmationsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this submit shipment confirmations not found response a status code equal to that given
func (o *SubmitShipmentConfirmationsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *SubmitShipmentConfirmationsNotFound) Error() string {
	return fmt.Sprintf("[POST /vendor/directFulfillment/shipping/v1/shipmentConfirmations][%d] submitShipmentConfirmationsNotFound  %+v", 404, o.Payload)
}

func (o *SubmitShipmentConfirmationsNotFound) String() string {
	return fmt.Sprintf("[POST /vendor/directFulfillment/shipping/v1/shipmentConfirmations][%d] submitShipmentConfirmationsNotFound  %+v", 404, o.Payload)
}

func (o *SubmitShipmentConfirmationsNotFound) GetPayload() *vendor_direct_fulfillment_shipping_v1_models.SubmitShipmentConfirmationsResponse {
	return o.Payload
}

func (o *SubmitShipmentConfirmationsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(vendor_direct_fulfillment_shipping_v1_models.SubmitShipmentConfirmationsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSubmitShipmentConfirmationsRequestEntityTooLarge creates a SubmitShipmentConfirmationsRequestEntityTooLarge with default headers values
func NewSubmitShipmentConfirmationsRequestEntityTooLarge() *SubmitShipmentConfirmationsRequestEntityTooLarge {
	return &SubmitShipmentConfirmationsRequestEntityTooLarge{}
}

/*
SubmitShipmentConfirmationsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request size exceeded the maximum accepted size.
*/
type SubmitShipmentConfirmationsRequestEntityTooLarge struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_shipping_v1_models.SubmitShipmentConfirmationsResponse
}

// IsSuccess returns true when this submit shipment confirmations request entity too large response has a 2xx status code
func (o *SubmitShipmentConfirmationsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this submit shipment confirmations request entity too large response has a 3xx status code
func (o *SubmitShipmentConfirmationsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this submit shipment confirmations request entity too large response has a 4xx status code
func (o *SubmitShipmentConfirmationsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this submit shipment confirmations request entity too large response has a 5xx status code
func (o *SubmitShipmentConfirmationsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this submit shipment confirmations request entity too large response a status code equal to that given
func (o *SubmitShipmentConfirmationsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *SubmitShipmentConfirmationsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /vendor/directFulfillment/shipping/v1/shipmentConfirmations][%d] submitShipmentConfirmationsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *SubmitShipmentConfirmationsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[POST /vendor/directFulfillment/shipping/v1/shipmentConfirmations][%d] submitShipmentConfirmationsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *SubmitShipmentConfirmationsRequestEntityTooLarge) GetPayload() *vendor_direct_fulfillment_shipping_v1_models.SubmitShipmentConfirmationsResponse {
	return o.Payload
}

func (o *SubmitShipmentConfirmationsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(vendor_direct_fulfillment_shipping_v1_models.SubmitShipmentConfirmationsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSubmitShipmentConfirmationsUnsupportedMediaType creates a SubmitShipmentConfirmationsUnsupportedMediaType with default headers values
func NewSubmitShipmentConfirmationsUnsupportedMediaType() *SubmitShipmentConfirmationsUnsupportedMediaType {
	return &SubmitShipmentConfirmationsUnsupportedMediaType{}
}

/*
SubmitShipmentConfirmationsUnsupportedMediaType describes a response with status code 415, with default header values.

The request payload is in an unsupported format.
*/
type SubmitShipmentConfirmationsUnsupportedMediaType struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_shipping_v1_models.SubmitShipmentConfirmationsResponse
}

// IsSuccess returns true when this submit shipment confirmations unsupported media type response has a 2xx status code
func (o *SubmitShipmentConfirmationsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this submit shipment confirmations unsupported media type response has a 3xx status code
func (o *SubmitShipmentConfirmationsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this submit shipment confirmations unsupported media type response has a 4xx status code
func (o *SubmitShipmentConfirmationsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this submit shipment confirmations unsupported media type response has a 5xx status code
func (o *SubmitShipmentConfirmationsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this submit shipment confirmations unsupported media type response a status code equal to that given
func (o *SubmitShipmentConfirmationsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *SubmitShipmentConfirmationsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /vendor/directFulfillment/shipping/v1/shipmentConfirmations][%d] submitShipmentConfirmationsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *SubmitShipmentConfirmationsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[POST /vendor/directFulfillment/shipping/v1/shipmentConfirmations][%d] submitShipmentConfirmationsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *SubmitShipmentConfirmationsUnsupportedMediaType) GetPayload() *vendor_direct_fulfillment_shipping_v1_models.SubmitShipmentConfirmationsResponse {
	return o.Payload
}

func (o *SubmitShipmentConfirmationsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(vendor_direct_fulfillment_shipping_v1_models.SubmitShipmentConfirmationsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSubmitShipmentConfirmationsTooManyRequests creates a SubmitShipmentConfirmationsTooManyRequests with default headers values
func NewSubmitShipmentConfirmationsTooManyRequests() *SubmitShipmentConfirmationsTooManyRequests {
	return &SubmitShipmentConfirmationsTooManyRequests{}
}

/*
SubmitShipmentConfirmationsTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type SubmitShipmentConfirmationsTooManyRequests struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_shipping_v1_models.SubmitShipmentConfirmationsResponse
}

// IsSuccess returns true when this submit shipment confirmations too many requests response has a 2xx status code
func (o *SubmitShipmentConfirmationsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this submit shipment confirmations too many requests response has a 3xx status code
func (o *SubmitShipmentConfirmationsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this submit shipment confirmations too many requests response has a 4xx status code
func (o *SubmitShipmentConfirmationsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this submit shipment confirmations too many requests response has a 5xx status code
func (o *SubmitShipmentConfirmationsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this submit shipment confirmations too many requests response a status code equal to that given
func (o *SubmitShipmentConfirmationsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *SubmitShipmentConfirmationsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /vendor/directFulfillment/shipping/v1/shipmentConfirmations][%d] submitShipmentConfirmationsTooManyRequests  %+v", 429, o.Payload)
}

func (o *SubmitShipmentConfirmationsTooManyRequests) String() string {
	return fmt.Sprintf("[POST /vendor/directFulfillment/shipping/v1/shipmentConfirmations][%d] submitShipmentConfirmationsTooManyRequests  %+v", 429, o.Payload)
}

func (o *SubmitShipmentConfirmationsTooManyRequests) GetPayload() *vendor_direct_fulfillment_shipping_v1_models.SubmitShipmentConfirmationsResponse {
	return o.Payload
}

func (o *SubmitShipmentConfirmationsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(vendor_direct_fulfillment_shipping_v1_models.SubmitShipmentConfirmationsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSubmitShipmentConfirmationsInternalServerError creates a SubmitShipmentConfirmationsInternalServerError with default headers values
func NewSubmitShipmentConfirmationsInternalServerError() *SubmitShipmentConfirmationsInternalServerError {
	return &SubmitShipmentConfirmationsInternalServerError{}
}

/*
SubmitShipmentConfirmationsInternalServerError describes a response with status code 500, with default header values.

Encountered an unexpected condition which prevented the server from fulfilling the request.
*/
type SubmitShipmentConfirmationsInternalServerError struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_shipping_v1_models.SubmitShipmentConfirmationsResponse
}

// IsSuccess returns true when this submit shipment confirmations internal server error response has a 2xx status code
func (o *SubmitShipmentConfirmationsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this submit shipment confirmations internal server error response has a 3xx status code
func (o *SubmitShipmentConfirmationsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this submit shipment confirmations internal server error response has a 4xx status code
func (o *SubmitShipmentConfirmationsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this submit shipment confirmations internal server error response has a 5xx status code
func (o *SubmitShipmentConfirmationsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this submit shipment confirmations internal server error response a status code equal to that given
func (o *SubmitShipmentConfirmationsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *SubmitShipmentConfirmationsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /vendor/directFulfillment/shipping/v1/shipmentConfirmations][%d] submitShipmentConfirmationsInternalServerError  %+v", 500, o.Payload)
}

func (o *SubmitShipmentConfirmationsInternalServerError) String() string {
	return fmt.Sprintf("[POST /vendor/directFulfillment/shipping/v1/shipmentConfirmations][%d] submitShipmentConfirmationsInternalServerError  %+v", 500, o.Payload)
}

func (o *SubmitShipmentConfirmationsInternalServerError) GetPayload() *vendor_direct_fulfillment_shipping_v1_models.SubmitShipmentConfirmationsResponse {
	return o.Payload
}

func (o *SubmitShipmentConfirmationsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(vendor_direct_fulfillment_shipping_v1_models.SubmitShipmentConfirmationsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSubmitShipmentConfirmationsServiceUnavailable creates a SubmitShipmentConfirmationsServiceUnavailable with default headers values
func NewSubmitShipmentConfirmationsServiceUnavailable() *SubmitShipmentConfirmationsServiceUnavailable {
	return &SubmitShipmentConfirmationsServiceUnavailable{}
}

/*
SubmitShipmentConfirmationsServiceUnavailable describes a response with status code 503, with default header values.

Temporary overloading or maintenance of the server.
*/
type SubmitShipmentConfirmationsServiceUnavailable struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_shipping_v1_models.SubmitShipmentConfirmationsResponse
}

// IsSuccess returns true when this submit shipment confirmations service unavailable response has a 2xx status code
func (o *SubmitShipmentConfirmationsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this submit shipment confirmations service unavailable response has a 3xx status code
func (o *SubmitShipmentConfirmationsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this submit shipment confirmations service unavailable response has a 4xx status code
func (o *SubmitShipmentConfirmationsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this submit shipment confirmations service unavailable response has a 5xx status code
func (o *SubmitShipmentConfirmationsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this submit shipment confirmations service unavailable response a status code equal to that given
func (o *SubmitShipmentConfirmationsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *SubmitShipmentConfirmationsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /vendor/directFulfillment/shipping/v1/shipmentConfirmations][%d] submitShipmentConfirmationsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *SubmitShipmentConfirmationsServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /vendor/directFulfillment/shipping/v1/shipmentConfirmations][%d] submitShipmentConfirmationsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *SubmitShipmentConfirmationsServiceUnavailable) GetPayload() *vendor_direct_fulfillment_shipping_v1_models.SubmitShipmentConfirmationsResponse {
	return o.Payload
}

func (o *SubmitShipmentConfirmationsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(vendor_direct_fulfillment_shipping_v1_models.SubmitShipmentConfirmationsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
