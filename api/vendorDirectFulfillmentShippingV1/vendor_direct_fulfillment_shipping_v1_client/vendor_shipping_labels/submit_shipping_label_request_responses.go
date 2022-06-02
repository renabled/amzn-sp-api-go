// Code generated by go-swagger; DO NOT EDIT.

package vendor_shipping_labels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/xamandar/amzn-sp-api-go/api/vendorDirectFulfillmentShippingV1/vendor_direct_fulfillment_shipping_v1_models"
)

// SubmitShippingLabelRequestReader is a Reader for the SubmitShippingLabelRequest structure.
type SubmitShippingLabelRequestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SubmitShippingLabelRequestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewSubmitShippingLabelRequestAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSubmitShippingLabelRequestBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSubmitShippingLabelRequestForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSubmitShippingLabelRequestNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewSubmitShippingLabelRequestRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewSubmitShippingLabelRequestUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewSubmitShippingLabelRequestTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSubmitShippingLabelRequestInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewSubmitShippingLabelRequestServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSubmitShippingLabelRequestAccepted creates a SubmitShippingLabelRequestAccepted with default headers values
func NewSubmitShippingLabelRequestAccepted() *SubmitShippingLabelRequestAccepted {
	return &SubmitShippingLabelRequestAccepted{}
}

/* SubmitShippingLabelRequestAccepted describes a response with status code 202, with default header values.

Success.
*/
type SubmitShippingLabelRequestAccepted struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_shipping_v1_models.SubmitShippingLabelsResponse
}

func (o *SubmitShippingLabelRequestAccepted) Error() string {
	return fmt.Sprintf("[POST /vendor/directFulfillment/shipping/v1/shippingLabels][%d] submitShippingLabelRequestAccepted  %+v", 202, o.Payload)
}
func (o *SubmitShippingLabelRequestAccepted) GetPayload() *vendor_direct_fulfillment_shipping_v1_models.SubmitShippingLabelsResponse {
	return o.Payload
}

func (o *SubmitShippingLabelRequestAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(vendor_direct_fulfillment_shipping_v1_models.SubmitShippingLabelsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSubmitShippingLabelRequestBadRequest creates a SubmitShippingLabelRequestBadRequest with default headers values
func NewSubmitShippingLabelRequestBadRequest() *SubmitShippingLabelRequestBadRequest {
	return &SubmitShippingLabelRequestBadRequest{}
}

/* SubmitShippingLabelRequestBadRequest describes a response with status code 400, with default header values.

Request has missing or invalid parameters and cannot be parsed.
*/
type SubmitShippingLabelRequestBadRequest struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_shipping_v1_models.SubmitShippingLabelsResponse
}

func (o *SubmitShippingLabelRequestBadRequest) Error() string {
	return fmt.Sprintf("[POST /vendor/directFulfillment/shipping/v1/shippingLabels][%d] submitShippingLabelRequestBadRequest  %+v", 400, o.Payload)
}
func (o *SubmitShippingLabelRequestBadRequest) GetPayload() *vendor_direct_fulfillment_shipping_v1_models.SubmitShippingLabelsResponse {
	return o.Payload
}

func (o *SubmitShippingLabelRequestBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(vendor_direct_fulfillment_shipping_v1_models.SubmitShippingLabelsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSubmitShippingLabelRequestForbidden creates a SubmitShippingLabelRequestForbidden with default headers values
func NewSubmitShippingLabelRequestForbidden() *SubmitShippingLabelRequestForbidden {
	return &SubmitShippingLabelRequestForbidden{}
}

/* SubmitShippingLabelRequestForbidden describes a response with status code 403, with default header values.

Indicates that access to the resource is forbidden. Possible reasons include Access Denied, Unauthorized, Expired Token, or Invalid Signature.
*/
type SubmitShippingLabelRequestForbidden struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_shipping_v1_models.SubmitShippingLabelsResponse
}

func (o *SubmitShippingLabelRequestForbidden) Error() string {
	return fmt.Sprintf("[POST /vendor/directFulfillment/shipping/v1/shippingLabels][%d] submitShippingLabelRequestForbidden  %+v", 403, o.Payload)
}
func (o *SubmitShippingLabelRequestForbidden) GetPayload() *vendor_direct_fulfillment_shipping_v1_models.SubmitShippingLabelsResponse {
	return o.Payload
}

func (o *SubmitShippingLabelRequestForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(vendor_direct_fulfillment_shipping_v1_models.SubmitShippingLabelsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSubmitShippingLabelRequestNotFound creates a SubmitShippingLabelRequestNotFound with default headers values
func NewSubmitShippingLabelRequestNotFound() *SubmitShippingLabelRequestNotFound {
	return &SubmitShippingLabelRequestNotFound{}
}

/* SubmitShippingLabelRequestNotFound describes a response with status code 404, with default header values.

The resource specified does not exist.
*/
type SubmitShippingLabelRequestNotFound struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_shipping_v1_models.SubmitShippingLabelsResponse
}

func (o *SubmitShippingLabelRequestNotFound) Error() string {
	return fmt.Sprintf("[POST /vendor/directFulfillment/shipping/v1/shippingLabels][%d] submitShippingLabelRequestNotFound  %+v", 404, o.Payload)
}
func (o *SubmitShippingLabelRequestNotFound) GetPayload() *vendor_direct_fulfillment_shipping_v1_models.SubmitShippingLabelsResponse {
	return o.Payload
}

func (o *SubmitShippingLabelRequestNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(vendor_direct_fulfillment_shipping_v1_models.SubmitShippingLabelsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSubmitShippingLabelRequestRequestEntityTooLarge creates a SubmitShippingLabelRequestRequestEntityTooLarge with default headers values
func NewSubmitShippingLabelRequestRequestEntityTooLarge() *SubmitShippingLabelRequestRequestEntityTooLarge {
	return &SubmitShippingLabelRequestRequestEntityTooLarge{}
}

/* SubmitShippingLabelRequestRequestEntityTooLarge describes a response with status code 413, with default header values.

The request size exceeded the maximum accepted size.
*/
type SubmitShippingLabelRequestRequestEntityTooLarge struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_shipping_v1_models.SubmitShippingLabelsResponse
}

func (o *SubmitShippingLabelRequestRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /vendor/directFulfillment/shipping/v1/shippingLabels][%d] submitShippingLabelRequestRequestEntityTooLarge  %+v", 413, o.Payload)
}
func (o *SubmitShippingLabelRequestRequestEntityTooLarge) GetPayload() *vendor_direct_fulfillment_shipping_v1_models.SubmitShippingLabelsResponse {
	return o.Payload
}

func (o *SubmitShippingLabelRequestRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(vendor_direct_fulfillment_shipping_v1_models.SubmitShippingLabelsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSubmitShippingLabelRequestUnsupportedMediaType creates a SubmitShippingLabelRequestUnsupportedMediaType with default headers values
func NewSubmitShippingLabelRequestUnsupportedMediaType() *SubmitShippingLabelRequestUnsupportedMediaType {
	return &SubmitShippingLabelRequestUnsupportedMediaType{}
}

/* SubmitShippingLabelRequestUnsupportedMediaType describes a response with status code 415, with default header values.

The request payload is in an unsupported format.
*/
type SubmitShippingLabelRequestUnsupportedMediaType struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_shipping_v1_models.SubmitShippingLabelsResponse
}

func (o *SubmitShippingLabelRequestUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /vendor/directFulfillment/shipping/v1/shippingLabels][%d] submitShippingLabelRequestUnsupportedMediaType  %+v", 415, o.Payload)
}
func (o *SubmitShippingLabelRequestUnsupportedMediaType) GetPayload() *vendor_direct_fulfillment_shipping_v1_models.SubmitShippingLabelsResponse {
	return o.Payload
}

func (o *SubmitShippingLabelRequestUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(vendor_direct_fulfillment_shipping_v1_models.SubmitShippingLabelsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSubmitShippingLabelRequestTooManyRequests creates a SubmitShippingLabelRequestTooManyRequests with default headers values
func NewSubmitShippingLabelRequestTooManyRequests() *SubmitShippingLabelRequestTooManyRequests {
	return &SubmitShippingLabelRequestTooManyRequests{}
}

/* SubmitShippingLabelRequestTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type SubmitShippingLabelRequestTooManyRequests struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_shipping_v1_models.SubmitShippingLabelsResponse
}

func (o *SubmitShippingLabelRequestTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /vendor/directFulfillment/shipping/v1/shippingLabels][%d] submitShippingLabelRequestTooManyRequests  %+v", 429, o.Payload)
}
func (o *SubmitShippingLabelRequestTooManyRequests) GetPayload() *vendor_direct_fulfillment_shipping_v1_models.SubmitShippingLabelsResponse {
	return o.Payload
}

func (o *SubmitShippingLabelRequestTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(vendor_direct_fulfillment_shipping_v1_models.SubmitShippingLabelsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSubmitShippingLabelRequestInternalServerError creates a SubmitShippingLabelRequestInternalServerError with default headers values
func NewSubmitShippingLabelRequestInternalServerError() *SubmitShippingLabelRequestInternalServerError {
	return &SubmitShippingLabelRequestInternalServerError{}
}

/* SubmitShippingLabelRequestInternalServerError describes a response with status code 500, with default header values.

Encountered an unexpected condition which prevented the server from fulfilling the request.
*/
type SubmitShippingLabelRequestInternalServerError struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_shipping_v1_models.SubmitShippingLabelsResponse
}

func (o *SubmitShippingLabelRequestInternalServerError) Error() string {
	return fmt.Sprintf("[POST /vendor/directFulfillment/shipping/v1/shippingLabels][%d] submitShippingLabelRequestInternalServerError  %+v", 500, o.Payload)
}
func (o *SubmitShippingLabelRequestInternalServerError) GetPayload() *vendor_direct_fulfillment_shipping_v1_models.SubmitShippingLabelsResponse {
	return o.Payload
}

func (o *SubmitShippingLabelRequestInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(vendor_direct_fulfillment_shipping_v1_models.SubmitShippingLabelsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSubmitShippingLabelRequestServiceUnavailable creates a SubmitShippingLabelRequestServiceUnavailable with default headers values
func NewSubmitShippingLabelRequestServiceUnavailable() *SubmitShippingLabelRequestServiceUnavailable {
	return &SubmitShippingLabelRequestServiceUnavailable{}
}

/* SubmitShippingLabelRequestServiceUnavailable describes a response with status code 503, with default header values.

Temporary overloading or maintenance of the server.
*/
type SubmitShippingLabelRequestServiceUnavailable struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_shipping_v1_models.SubmitShippingLabelsResponse
}

func (o *SubmitShippingLabelRequestServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /vendor/directFulfillment/shipping/v1/shippingLabels][%d] submitShippingLabelRequestServiceUnavailable  %+v", 503, o.Payload)
}
func (o *SubmitShippingLabelRequestServiceUnavailable) GetPayload() *vendor_direct_fulfillment_shipping_v1_models.SubmitShippingLabelsResponse {
	return o.Payload
}

func (o *SubmitShippingLabelRequestServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(vendor_direct_fulfillment_shipping_v1_models.SubmitShippingLabelsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
