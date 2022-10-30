// Code generated by go-swagger; DO NOT EDIT.

package vendor_shipping_labels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/xamandar/amzn-sp-api-go/api/vendorDirectFulfillmentShipping_2021-12-28/vendor_direct_fulfillment_shipping_2021_12_28_models"
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

/*
SubmitShippingLabelRequestAccepted describes a response with status code 202, with default header values.

Success.
*/
type SubmitShippingLabelRequestAccepted struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_shipping_2021_12_28_models.TransactionReference
}

// IsSuccess returns true when this submit shipping label request accepted response has a 2xx status code
func (o *SubmitShippingLabelRequestAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this submit shipping label request accepted response has a 3xx status code
func (o *SubmitShippingLabelRequestAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this submit shipping label request accepted response has a 4xx status code
func (o *SubmitShippingLabelRequestAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this submit shipping label request accepted response has a 5xx status code
func (o *SubmitShippingLabelRequestAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this submit shipping label request accepted response a status code equal to that given
func (o *SubmitShippingLabelRequestAccepted) IsCode(code int) bool {
	return code == 202
}

func (o *SubmitShippingLabelRequestAccepted) Error() string {
	return fmt.Sprintf("[POST /vendor/directFulfillment/shipping/2021-12-28/shippingLabels][%d] submitShippingLabelRequestAccepted  %+v", 202, o.Payload)
}

func (o *SubmitShippingLabelRequestAccepted) String() string {
	return fmt.Sprintf("[POST /vendor/directFulfillment/shipping/2021-12-28/shippingLabels][%d] submitShippingLabelRequestAccepted  %+v", 202, o.Payload)
}

func (o *SubmitShippingLabelRequestAccepted) GetPayload() *vendor_direct_fulfillment_shipping_2021_12_28_models.TransactionReference {
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

	o.Payload = new(vendor_direct_fulfillment_shipping_2021_12_28_models.TransactionReference)

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

/*
SubmitShippingLabelRequestBadRequest describes a response with status code 400, with default header values.

Request has missing or invalid parameters and cannot be parsed.
*/
type SubmitShippingLabelRequestBadRequest struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_shipping_2021_12_28_models.ErrorList
}

// IsSuccess returns true when this submit shipping label request bad request response has a 2xx status code
func (o *SubmitShippingLabelRequestBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this submit shipping label request bad request response has a 3xx status code
func (o *SubmitShippingLabelRequestBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this submit shipping label request bad request response has a 4xx status code
func (o *SubmitShippingLabelRequestBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this submit shipping label request bad request response has a 5xx status code
func (o *SubmitShippingLabelRequestBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this submit shipping label request bad request response a status code equal to that given
func (o *SubmitShippingLabelRequestBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *SubmitShippingLabelRequestBadRequest) Error() string {
	return fmt.Sprintf("[POST /vendor/directFulfillment/shipping/2021-12-28/shippingLabels][%d] submitShippingLabelRequestBadRequest  %+v", 400, o.Payload)
}

func (o *SubmitShippingLabelRequestBadRequest) String() string {
	return fmt.Sprintf("[POST /vendor/directFulfillment/shipping/2021-12-28/shippingLabels][%d] submitShippingLabelRequestBadRequest  %+v", 400, o.Payload)
}

func (o *SubmitShippingLabelRequestBadRequest) GetPayload() *vendor_direct_fulfillment_shipping_2021_12_28_models.ErrorList {
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

	o.Payload = new(vendor_direct_fulfillment_shipping_2021_12_28_models.ErrorList)

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

/*
SubmitShippingLabelRequestForbidden describes a response with status code 403, with default header values.

Indicates that access to the resource is forbidden. Possible reasons include Access Denied, Unauthorized, Expired Token, or Invalid Signature.
*/
type SubmitShippingLabelRequestForbidden struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_shipping_2021_12_28_models.ErrorList
}

// IsSuccess returns true when this submit shipping label request forbidden response has a 2xx status code
func (o *SubmitShippingLabelRequestForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this submit shipping label request forbidden response has a 3xx status code
func (o *SubmitShippingLabelRequestForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this submit shipping label request forbidden response has a 4xx status code
func (o *SubmitShippingLabelRequestForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this submit shipping label request forbidden response has a 5xx status code
func (o *SubmitShippingLabelRequestForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this submit shipping label request forbidden response a status code equal to that given
func (o *SubmitShippingLabelRequestForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *SubmitShippingLabelRequestForbidden) Error() string {
	return fmt.Sprintf("[POST /vendor/directFulfillment/shipping/2021-12-28/shippingLabels][%d] submitShippingLabelRequestForbidden  %+v", 403, o.Payload)
}

func (o *SubmitShippingLabelRequestForbidden) String() string {
	return fmt.Sprintf("[POST /vendor/directFulfillment/shipping/2021-12-28/shippingLabels][%d] submitShippingLabelRequestForbidden  %+v", 403, o.Payload)
}

func (o *SubmitShippingLabelRequestForbidden) GetPayload() *vendor_direct_fulfillment_shipping_2021_12_28_models.ErrorList {
	return o.Payload
}

func (o *SubmitShippingLabelRequestForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewSubmitShippingLabelRequestNotFound creates a SubmitShippingLabelRequestNotFound with default headers values
func NewSubmitShippingLabelRequestNotFound() *SubmitShippingLabelRequestNotFound {
	return &SubmitShippingLabelRequestNotFound{}
}

/*
SubmitShippingLabelRequestNotFound describes a response with status code 404, with default header values.

The resource specified does not exist.
*/
type SubmitShippingLabelRequestNotFound struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_shipping_2021_12_28_models.ErrorList
}

// IsSuccess returns true when this submit shipping label request not found response has a 2xx status code
func (o *SubmitShippingLabelRequestNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this submit shipping label request not found response has a 3xx status code
func (o *SubmitShippingLabelRequestNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this submit shipping label request not found response has a 4xx status code
func (o *SubmitShippingLabelRequestNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this submit shipping label request not found response has a 5xx status code
func (o *SubmitShippingLabelRequestNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this submit shipping label request not found response a status code equal to that given
func (o *SubmitShippingLabelRequestNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *SubmitShippingLabelRequestNotFound) Error() string {
	return fmt.Sprintf("[POST /vendor/directFulfillment/shipping/2021-12-28/shippingLabels][%d] submitShippingLabelRequestNotFound  %+v", 404, o.Payload)
}

func (o *SubmitShippingLabelRequestNotFound) String() string {
	return fmt.Sprintf("[POST /vendor/directFulfillment/shipping/2021-12-28/shippingLabels][%d] submitShippingLabelRequestNotFound  %+v", 404, o.Payload)
}

func (o *SubmitShippingLabelRequestNotFound) GetPayload() *vendor_direct_fulfillment_shipping_2021_12_28_models.ErrorList {
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

	o.Payload = new(vendor_direct_fulfillment_shipping_2021_12_28_models.ErrorList)

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

/*
SubmitShippingLabelRequestRequestEntityTooLarge describes a response with status code 413, with default header values.

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

	Payload *vendor_direct_fulfillment_shipping_2021_12_28_models.ErrorList
}

// IsSuccess returns true when this submit shipping label request request entity too large response has a 2xx status code
func (o *SubmitShippingLabelRequestRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this submit shipping label request request entity too large response has a 3xx status code
func (o *SubmitShippingLabelRequestRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this submit shipping label request request entity too large response has a 4xx status code
func (o *SubmitShippingLabelRequestRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this submit shipping label request request entity too large response has a 5xx status code
func (o *SubmitShippingLabelRequestRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this submit shipping label request request entity too large response a status code equal to that given
func (o *SubmitShippingLabelRequestRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *SubmitShippingLabelRequestRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /vendor/directFulfillment/shipping/2021-12-28/shippingLabels][%d] submitShippingLabelRequestRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *SubmitShippingLabelRequestRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[POST /vendor/directFulfillment/shipping/2021-12-28/shippingLabels][%d] submitShippingLabelRequestRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *SubmitShippingLabelRequestRequestEntityTooLarge) GetPayload() *vendor_direct_fulfillment_shipping_2021_12_28_models.ErrorList {
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

	o.Payload = new(vendor_direct_fulfillment_shipping_2021_12_28_models.ErrorList)

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

/*
SubmitShippingLabelRequestUnsupportedMediaType describes a response with status code 415, with default header values.

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

	Payload *vendor_direct_fulfillment_shipping_2021_12_28_models.ErrorList
}

// IsSuccess returns true when this submit shipping label request unsupported media type response has a 2xx status code
func (o *SubmitShippingLabelRequestUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this submit shipping label request unsupported media type response has a 3xx status code
func (o *SubmitShippingLabelRequestUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this submit shipping label request unsupported media type response has a 4xx status code
func (o *SubmitShippingLabelRequestUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this submit shipping label request unsupported media type response has a 5xx status code
func (o *SubmitShippingLabelRequestUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this submit shipping label request unsupported media type response a status code equal to that given
func (o *SubmitShippingLabelRequestUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *SubmitShippingLabelRequestUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /vendor/directFulfillment/shipping/2021-12-28/shippingLabels][%d] submitShippingLabelRequestUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *SubmitShippingLabelRequestUnsupportedMediaType) String() string {
	return fmt.Sprintf("[POST /vendor/directFulfillment/shipping/2021-12-28/shippingLabels][%d] submitShippingLabelRequestUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *SubmitShippingLabelRequestUnsupportedMediaType) GetPayload() *vendor_direct_fulfillment_shipping_2021_12_28_models.ErrorList {
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

	o.Payload = new(vendor_direct_fulfillment_shipping_2021_12_28_models.ErrorList)

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

/*
SubmitShippingLabelRequestTooManyRequests describes a response with status code 429, with default header values.

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

	Payload *vendor_direct_fulfillment_shipping_2021_12_28_models.ErrorList
}

// IsSuccess returns true when this submit shipping label request too many requests response has a 2xx status code
func (o *SubmitShippingLabelRequestTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this submit shipping label request too many requests response has a 3xx status code
func (o *SubmitShippingLabelRequestTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this submit shipping label request too many requests response has a 4xx status code
func (o *SubmitShippingLabelRequestTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this submit shipping label request too many requests response has a 5xx status code
func (o *SubmitShippingLabelRequestTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this submit shipping label request too many requests response a status code equal to that given
func (o *SubmitShippingLabelRequestTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *SubmitShippingLabelRequestTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /vendor/directFulfillment/shipping/2021-12-28/shippingLabels][%d] submitShippingLabelRequestTooManyRequests  %+v", 429, o.Payload)
}

func (o *SubmitShippingLabelRequestTooManyRequests) String() string {
	return fmt.Sprintf("[POST /vendor/directFulfillment/shipping/2021-12-28/shippingLabels][%d] submitShippingLabelRequestTooManyRequests  %+v", 429, o.Payload)
}

func (o *SubmitShippingLabelRequestTooManyRequests) GetPayload() *vendor_direct_fulfillment_shipping_2021_12_28_models.ErrorList {
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

	o.Payload = new(vendor_direct_fulfillment_shipping_2021_12_28_models.ErrorList)

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

/*
SubmitShippingLabelRequestInternalServerError describes a response with status code 500, with default header values.

An unexpected condition occurred that prevented the server from fulfilling the request.
*/
type SubmitShippingLabelRequestInternalServerError struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_shipping_2021_12_28_models.ErrorList
}

// IsSuccess returns true when this submit shipping label request internal server error response has a 2xx status code
func (o *SubmitShippingLabelRequestInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this submit shipping label request internal server error response has a 3xx status code
func (o *SubmitShippingLabelRequestInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this submit shipping label request internal server error response has a 4xx status code
func (o *SubmitShippingLabelRequestInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this submit shipping label request internal server error response has a 5xx status code
func (o *SubmitShippingLabelRequestInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this submit shipping label request internal server error response a status code equal to that given
func (o *SubmitShippingLabelRequestInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *SubmitShippingLabelRequestInternalServerError) Error() string {
	return fmt.Sprintf("[POST /vendor/directFulfillment/shipping/2021-12-28/shippingLabels][%d] submitShippingLabelRequestInternalServerError  %+v", 500, o.Payload)
}

func (o *SubmitShippingLabelRequestInternalServerError) String() string {
	return fmt.Sprintf("[POST /vendor/directFulfillment/shipping/2021-12-28/shippingLabels][%d] submitShippingLabelRequestInternalServerError  %+v", 500, o.Payload)
}

func (o *SubmitShippingLabelRequestInternalServerError) GetPayload() *vendor_direct_fulfillment_shipping_2021_12_28_models.ErrorList {
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

	o.Payload = new(vendor_direct_fulfillment_shipping_2021_12_28_models.ErrorList)

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

/*
SubmitShippingLabelRequestServiceUnavailable describes a response with status code 503, with default header values.

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

	Payload *vendor_direct_fulfillment_shipping_2021_12_28_models.ErrorList
}

// IsSuccess returns true when this submit shipping label request service unavailable response has a 2xx status code
func (o *SubmitShippingLabelRequestServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this submit shipping label request service unavailable response has a 3xx status code
func (o *SubmitShippingLabelRequestServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this submit shipping label request service unavailable response has a 4xx status code
func (o *SubmitShippingLabelRequestServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this submit shipping label request service unavailable response has a 5xx status code
func (o *SubmitShippingLabelRequestServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this submit shipping label request service unavailable response a status code equal to that given
func (o *SubmitShippingLabelRequestServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *SubmitShippingLabelRequestServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /vendor/directFulfillment/shipping/2021-12-28/shippingLabels][%d] submitShippingLabelRequestServiceUnavailable  %+v", 503, o.Payload)
}

func (o *SubmitShippingLabelRequestServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /vendor/directFulfillment/shipping/2021-12-28/shippingLabels][%d] submitShippingLabelRequestServiceUnavailable  %+v", 503, o.Payload)
}

func (o *SubmitShippingLabelRequestServiceUnavailable) GetPayload() *vendor_direct_fulfillment_shipping_2021_12_28_models.ErrorList {
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

	o.Payload = new(vendor_direct_fulfillment_shipping_2021_12_28_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
