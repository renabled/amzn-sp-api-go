// Code generated by go-swagger; DO NOT EDIT.

package vendor_shipping_labels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/renabled/amzn-sp-api-go/api/vendorDirectFulfillmentShipping_2021-12-28/vendor_direct_fulfillment_shipping_2021_12_28_models"
)

// GetShippingLabelReader is a Reader for the GetShippingLabel structure.
type GetShippingLabelReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetShippingLabelReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetShippingLabelOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetShippingLabelBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetShippingLabelUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetShippingLabelForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetShippingLabelNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetShippingLabelUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetShippingLabelTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetShippingLabelInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetShippingLabelServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetShippingLabelOK creates a GetShippingLabelOK with default headers values
func NewGetShippingLabelOK() *GetShippingLabelOK {
	return &GetShippingLabelOK{}
}

/*
GetShippingLabelOK describes a response with status code 200, with default header values.

Success.
*/
type GetShippingLabelOK struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_shipping_2021_12_28_models.ShippingLabel
}

// IsSuccess returns true when this get shipping label o k response has a 2xx status code
func (o *GetShippingLabelOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get shipping label o k response has a 3xx status code
func (o *GetShippingLabelOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get shipping label o k response has a 4xx status code
func (o *GetShippingLabelOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get shipping label o k response has a 5xx status code
func (o *GetShippingLabelOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get shipping label o k response a status code equal to that given
func (o *GetShippingLabelOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetShippingLabelOK) Error() string {
	return fmt.Sprintf("[GET /vendor/directFulfillment/shipping/2021-12-28/shippingLabels/{purchaseOrderNumber}][%d] getShippingLabelOK  %+v", 200, o.Payload)
}

func (o *GetShippingLabelOK) String() string {
	return fmt.Sprintf("[GET /vendor/directFulfillment/shipping/2021-12-28/shippingLabels/{purchaseOrderNumber}][%d] getShippingLabelOK  %+v", 200, o.Payload)
}

func (o *GetShippingLabelOK) GetPayload() *vendor_direct_fulfillment_shipping_2021_12_28_models.ShippingLabel {
	return o.Payload
}

func (o *GetShippingLabelOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(vendor_direct_fulfillment_shipping_2021_12_28_models.ShippingLabel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetShippingLabelBadRequest creates a GetShippingLabelBadRequest with default headers values
func NewGetShippingLabelBadRequest() *GetShippingLabelBadRequest {
	return &GetShippingLabelBadRequest{}
}

/*
GetShippingLabelBadRequest describes a response with status code 400, with default header values.

Request has missing or invalid parameters and cannot be parsed.
*/
type GetShippingLabelBadRequest struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_shipping_2021_12_28_models.ErrorList
}

// IsSuccess returns true when this get shipping label bad request response has a 2xx status code
func (o *GetShippingLabelBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get shipping label bad request response has a 3xx status code
func (o *GetShippingLabelBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get shipping label bad request response has a 4xx status code
func (o *GetShippingLabelBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get shipping label bad request response has a 5xx status code
func (o *GetShippingLabelBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get shipping label bad request response a status code equal to that given
func (o *GetShippingLabelBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetShippingLabelBadRequest) Error() string {
	return fmt.Sprintf("[GET /vendor/directFulfillment/shipping/2021-12-28/shippingLabels/{purchaseOrderNumber}][%d] getShippingLabelBadRequest  %+v", 400, o.Payload)
}

func (o *GetShippingLabelBadRequest) String() string {
	return fmt.Sprintf("[GET /vendor/directFulfillment/shipping/2021-12-28/shippingLabels/{purchaseOrderNumber}][%d] getShippingLabelBadRequest  %+v", 400, o.Payload)
}

func (o *GetShippingLabelBadRequest) GetPayload() *vendor_direct_fulfillment_shipping_2021_12_28_models.ErrorList {
	return o.Payload
}

func (o *GetShippingLabelBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetShippingLabelUnauthorized creates a GetShippingLabelUnauthorized with default headers values
func NewGetShippingLabelUnauthorized() *GetShippingLabelUnauthorized {
	return &GetShippingLabelUnauthorized{}
}

/*
GetShippingLabelUnauthorized describes a response with status code 401, with default header values.

The request's Authorization header is not formatted correctly or does not contain a valid token.
*/
type GetShippingLabelUnauthorized struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_shipping_2021_12_28_models.ErrorList
}

// IsSuccess returns true when this get shipping label unauthorized response has a 2xx status code
func (o *GetShippingLabelUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get shipping label unauthorized response has a 3xx status code
func (o *GetShippingLabelUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get shipping label unauthorized response has a 4xx status code
func (o *GetShippingLabelUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get shipping label unauthorized response has a 5xx status code
func (o *GetShippingLabelUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get shipping label unauthorized response a status code equal to that given
func (o *GetShippingLabelUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetShippingLabelUnauthorized) Error() string {
	return fmt.Sprintf("[GET /vendor/directFulfillment/shipping/2021-12-28/shippingLabels/{purchaseOrderNumber}][%d] getShippingLabelUnauthorized  %+v", 401, o.Payload)
}

func (o *GetShippingLabelUnauthorized) String() string {
	return fmt.Sprintf("[GET /vendor/directFulfillment/shipping/2021-12-28/shippingLabels/{purchaseOrderNumber}][%d] getShippingLabelUnauthorized  %+v", 401, o.Payload)
}

func (o *GetShippingLabelUnauthorized) GetPayload() *vendor_direct_fulfillment_shipping_2021_12_28_models.ErrorList {
	return o.Payload
}

func (o *GetShippingLabelUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetShippingLabelForbidden creates a GetShippingLabelForbidden with default headers values
func NewGetShippingLabelForbidden() *GetShippingLabelForbidden {
	return &GetShippingLabelForbidden{}
}

/*
GetShippingLabelForbidden describes a response with status code 403, with default header values.

Indicates that access to the resource is forbidden. Possible reasons include Access Denied, Unauthorized, Expired Token, or Invalid Signature.
*/
type GetShippingLabelForbidden struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_shipping_2021_12_28_models.ErrorList
}

// IsSuccess returns true when this get shipping label forbidden response has a 2xx status code
func (o *GetShippingLabelForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get shipping label forbidden response has a 3xx status code
func (o *GetShippingLabelForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get shipping label forbidden response has a 4xx status code
func (o *GetShippingLabelForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get shipping label forbidden response has a 5xx status code
func (o *GetShippingLabelForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get shipping label forbidden response a status code equal to that given
func (o *GetShippingLabelForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetShippingLabelForbidden) Error() string {
	return fmt.Sprintf("[GET /vendor/directFulfillment/shipping/2021-12-28/shippingLabels/{purchaseOrderNumber}][%d] getShippingLabelForbidden  %+v", 403, o.Payload)
}

func (o *GetShippingLabelForbidden) String() string {
	return fmt.Sprintf("[GET /vendor/directFulfillment/shipping/2021-12-28/shippingLabels/{purchaseOrderNumber}][%d] getShippingLabelForbidden  %+v", 403, o.Payload)
}

func (o *GetShippingLabelForbidden) GetPayload() *vendor_direct_fulfillment_shipping_2021_12_28_models.ErrorList {
	return o.Payload
}

func (o *GetShippingLabelForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetShippingLabelNotFound creates a GetShippingLabelNotFound with default headers values
func NewGetShippingLabelNotFound() *GetShippingLabelNotFound {
	return &GetShippingLabelNotFound{}
}

/*
GetShippingLabelNotFound describes a response with status code 404, with default header values.

The resource specified does not exist.
*/
type GetShippingLabelNotFound struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_shipping_2021_12_28_models.ErrorList
}

// IsSuccess returns true when this get shipping label not found response has a 2xx status code
func (o *GetShippingLabelNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get shipping label not found response has a 3xx status code
func (o *GetShippingLabelNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get shipping label not found response has a 4xx status code
func (o *GetShippingLabelNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get shipping label not found response has a 5xx status code
func (o *GetShippingLabelNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get shipping label not found response a status code equal to that given
func (o *GetShippingLabelNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetShippingLabelNotFound) Error() string {
	return fmt.Sprintf("[GET /vendor/directFulfillment/shipping/2021-12-28/shippingLabels/{purchaseOrderNumber}][%d] getShippingLabelNotFound  %+v", 404, o.Payload)
}

func (o *GetShippingLabelNotFound) String() string {
	return fmt.Sprintf("[GET /vendor/directFulfillment/shipping/2021-12-28/shippingLabels/{purchaseOrderNumber}][%d] getShippingLabelNotFound  %+v", 404, o.Payload)
}

func (o *GetShippingLabelNotFound) GetPayload() *vendor_direct_fulfillment_shipping_2021_12_28_models.ErrorList {
	return o.Payload
}

func (o *GetShippingLabelNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetShippingLabelUnsupportedMediaType creates a GetShippingLabelUnsupportedMediaType with default headers values
func NewGetShippingLabelUnsupportedMediaType() *GetShippingLabelUnsupportedMediaType {
	return &GetShippingLabelUnsupportedMediaType{}
}

/*
GetShippingLabelUnsupportedMediaType describes a response with status code 415, with default header values.

The request payload is in an unsupported format.
*/
type GetShippingLabelUnsupportedMediaType struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_shipping_2021_12_28_models.ErrorList
}

// IsSuccess returns true when this get shipping label unsupported media type response has a 2xx status code
func (o *GetShippingLabelUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get shipping label unsupported media type response has a 3xx status code
func (o *GetShippingLabelUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get shipping label unsupported media type response has a 4xx status code
func (o *GetShippingLabelUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get shipping label unsupported media type response has a 5xx status code
func (o *GetShippingLabelUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get shipping label unsupported media type response a status code equal to that given
func (o *GetShippingLabelUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetShippingLabelUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /vendor/directFulfillment/shipping/2021-12-28/shippingLabels/{purchaseOrderNumber}][%d] getShippingLabelUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetShippingLabelUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /vendor/directFulfillment/shipping/2021-12-28/shippingLabels/{purchaseOrderNumber}][%d] getShippingLabelUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetShippingLabelUnsupportedMediaType) GetPayload() *vendor_direct_fulfillment_shipping_2021_12_28_models.ErrorList {
	return o.Payload
}

func (o *GetShippingLabelUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetShippingLabelTooManyRequests creates a GetShippingLabelTooManyRequests with default headers values
func NewGetShippingLabelTooManyRequests() *GetShippingLabelTooManyRequests {
	return &GetShippingLabelTooManyRequests{}
}

/*
GetShippingLabelTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type GetShippingLabelTooManyRequests struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_shipping_2021_12_28_models.ErrorList
}

// IsSuccess returns true when this get shipping label too many requests response has a 2xx status code
func (o *GetShippingLabelTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get shipping label too many requests response has a 3xx status code
func (o *GetShippingLabelTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get shipping label too many requests response has a 4xx status code
func (o *GetShippingLabelTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get shipping label too many requests response has a 5xx status code
func (o *GetShippingLabelTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get shipping label too many requests response a status code equal to that given
func (o *GetShippingLabelTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetShippingLabelTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /vendor/directFulfillment/shipping/2021-12-28/shippingLabels/{purchaseOrderNumber}][%d] getShippingLabelTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetShippingLabelTooManyRequests) String() string {
	return fmt.Sprintf("[GET /vendor/directFulfillment/shipping/2021-12-28/shippingLabels/{purchaseOrderNumber}][%d] getShippingLabelTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetShippingLabelTooManyRequests) GetPayload() *vendor_direct_fulfillment_shipping_2021_12_28_models.ErrorList {
	return o.Payload
}

func (o *GetShippingLabelTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetShippingLabelInternalServerError creates a GetShippingLabelInternalServerError with default headers values
func NewGetShippingLabelInternalServerError() *GetShippingLabelInternalServerError {
	return &GetShippingLabelInternalServerError{}
}

/*
GetShippingLabelInternalServerError describes a response with status code 500, with default header values.

An unexpected condition occurred that prevented the server from fulfilling the request.
*/
type GetShippingLabelInternalServerError struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_shipping_2021_12_28_models.ErrorList
}

// IsSuccess returns true when this get shipping label internal server error response has a 2xx status code
func (o *GetShippingLabelInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get shipping label internal server error response has a 3xx status code
func (o *GetShippingLabelInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get shipping label internal server error response has a 4xx status code
func (o *GetShippingLabelInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get shipping label internal server error response has a 5xx status code
func (o *GetShippingLabelInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get shipping label internal server error response a status code equal to that given
func (o *GetShippingLabelInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetShippingLabelInternalServerError) Error() string {
	return fmt.Sprintf("[GET /vendor/directFulfillment/shipping/2021-12-28/shippingLabels/{purchaseOrderNumber}][%d] getShippingLabelInternalServerError  %+v", 500, o.Payload)
}

func (o *GetShippingLabelInternalServerError) String() string {
	return fmt.Sprintf("[GET /vendor/directFulfillment/shipping/2021-12-28/shippingLabels/{purchaseOrderNumber}][%d] getShippingLabelInternalServerError  %+v", 500, o.Payload)
}

func (o *GetShippingLabelInternalServerError) GetPayload() *vendor_direct_fulfillment_shipping_2021_12_28_models.ErrorList {
	return o.Payload
}

func (o *GetShippingLabelInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetShippingLabelServiceUnavailable creates a GetShippingLabelServiceUnavailable with default headers values
func NewGetShippingLabelServiceUnavailable() *GetShippingLabelServiceUnavailable {
	return &GetShippingLabelServiceUnavailable{}
}

/*
GetShippingLabelServiceUnavailable describes a response with status code 503, with default header values.

Temporary overloading or maintenance of the server.
*/
type GetShippingLabelServiceUnavailable struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_shipping_2021_12_28_models.ErrorList
}

// IsSuccess returns true when this get shipping label service unavailable response has a 2xx status code
func (o *GetShippingLabelServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get shipping label service unavailable response has a 3xx status code
func (o *GetShippingLabelServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get shipping label service unavailable response has a 4xx status code
func (o *GetShippingLabelServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get shipping label service unavailable response has a 5xx status code
func (o *GetShippingLabelServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get shipping label service unavailable response a status code equal to that given
func (o *GetShippingLabelServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetShippingLabelServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /vendor/directFulfillment/shipping/2021-12-28/shippingLabels/{purchaseOrderNumber}][%d] getShippingLabelServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetShippingLabelServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /vendor/directFulfillment/shipping/2021-12-28/shippingLabels/{purchaseOrderNumber}][%d] getShippingLabelServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetShippingLabelServiceUnavailable) GetPayload() *vendor_direct_fulfillment_shipping_2021_12_28_models.ErrorList {
	return o.Payload
}

func (o *GetShippingLabelServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
