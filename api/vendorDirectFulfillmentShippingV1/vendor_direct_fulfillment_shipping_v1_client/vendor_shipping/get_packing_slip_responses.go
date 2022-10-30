// Code generated by go-swagger; DO NOT EDIT.

package vendor_shipping

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/xamandar/amzn-sp-api-go/api/vendorDirectFulfillmentShippingV1/vendor_direct_fulfillment_shipping_v1_models"
)

// GetPackingSlipReader is a Reader for the GetPackingSlip structure.
type GetPackingSlipReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPackingSlipReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPackingSlipOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetPackingSlipBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetPackingSlipUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetPackingSlipForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetPackingSlipNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetPackingSlipUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetPackingSlipTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetPackingSlipInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetPackingSlipServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetPackingSlipOK creates a GetPackingSlipOK with default headers values
func NewGetPackingSlipOK() *GetPackingSlipOK {
	return &GetPackingSlipOK{}
}

/*
GetPackingSlipOK describes a response with status code 200, with default header values.

Success.
*/
type GetPackingSlipOK struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_shipping_v1_models.GetPackingSlipResponse
}

// IsSuccess returns true when this get packing slip o k response has a 2xx status code
func (o *GetPackingSlipOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get packing slip o k response has a 3xx status code
func (o *GetPackingSlipOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get packing slip o k response has a 4xx status code
func (o *GetPackingSlipOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get packing slip o k response has a 5xx status code
func (o *GetPackingSlipOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get packing slip o k response a status code equal to that given
func (o *GetPackingSlipOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetPackingSlipOK) Error() string {
	return fmt.Sprintf("[GET /vendor/directFulfillment/shipping/v1/packingSlips/{purchaseOrderNumber}][%d] getPackingSlipOK  %+v", 200, o.Payload)
}

func (o *GetPackingSlipOK) String() string {
	return fmt.Sprintf("[GET /vendor/directFulfillment/shipping/v1/packingSlips/{purchaseOrderNumber}][%d] getPackingSlipOK  %+v", 200, o.Payload)
}

func (o *GetPackingSlipOK) GetPayload() *vendor_direct_fulfillment_shipping_v1_models.GetPackingSlipResponse {
	return o.Payload
}

func (o *GetPackingSlipOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(vendor_direct_fulfillment_shipping_v1_models.GetPackingSlipResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPackingSlipBadRequest creates a GetPackingSlipBadRequest with default headers values
func NewGetPackingSlipBadRequest() *GetPackingSlipBadRequest {
	return &GetPackingSlipBadRequest{}
}

/*
GetPackingSlipBadRequest describes a response with status code 400, with default header values.

Request has missing or invalid parameters and cannot be parsed.
*/
type GetPackingSlipBadRequest struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_shipping_v1_models.GetPackingSlipResponse
}

// IsSuccess returns true when this get packing slip bad request response has a 2xx status code
func (o *GetPackingSlipBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get packing slip bad request response has a 3xx status code
func (o *GetPackingSlipBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get packing slip bad request response has a 4xx status code
func (o *GetPackingSlipBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get packing slip bad request response has a 5xx status code
func (o *GetPackingSlipBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get packing slip bad request response a status code equal to that given
func (o *GetPackingSlipBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetPackingSlipBadRequest) Error() string {
	return fmt.Sprintf("[GET /vendor/directFulfillment/shipping/v1/packingSlips/{purchaseOrderNumber}][%d] getPackingSlipBadRequest  %+v", 400, o.Payload)
}

func (o *GetPackingSlipBadRequest) String() string {
	return fmt.Sprintf("[GET /vendor/directFulfillment/shipping/v1/packingSlips/{purchaseOrderNumber}][%d] getPackingSlipBadRequest  %+v", 400, o.Payload)
}

func (o *GetPackingSlipBadRequest) GetPayload() *vendor_direct_fulfillment_shipping_v1_models.GetPackingSlipResponse {
	return o.Payload
}

func (o *GetPackingSlipBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(vendor_direct_fulfillment_shipping_v1_models.GetPackingSlipResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPackingSlipUnauthorized creates a GetPackingSlipUnauthorized with default headers values
func NewGetPackingSlipUnauthorized() *GetPackingSlipUnauthorized {
	return &GetPackingSlipUnauthorized{}
}

/*
GetPackingSlipUnauthorized describes a response with status code 401, with default header values.

The request's Authorization header is not formatted correctly or does not contain a valid token.
*/
type GetPackingSlipUnauthorized struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_shipping_v1_models.GetPackingSlipResponse
}

// IsSuccess returns true when this get packing slip unauthorized response has a 2xx status code
func (o *GetPackingSlipUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get packing slip unauthorized response has a 3xx status code
func (o *GetPackingSlipUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get packing slip unauthorized response has a 4xx status code
func (o *GetPackingSlipUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get packing slip unauthorized response has a 5xx status code
func (o *GetPackingSlipUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get packing slip unauthorized response a status code equal to that given
func (o *GetPackingSlipUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetPackingSlipUnauthorized) Error() string {
	return fmt.Sprintf("[GET /vendor/directFulfillment/shipping/v1/packingSlips/{purchaseOrderNumber}][%d] getPackingSlipUnauthorized  %+v", 401, o.Payload)
}

func (o *GetPackingSlipUnauthorized) String() string {
	return fmt.Sprintf("[GET /vendor/directFulfillment/shipping/v1/packingSlips/{purchaseOrderNumber}][%d] getPackingSlipUnauthorized  %+v", 401, o.Payload)
}

func (o *GetPackingSlipUnauthorized) GetPayload() *vendor_direct_fulfillment_shipping_v1_models.GetPackingSlipResponse {
	return o.Payload
}

func (o *GetPackingSlipUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(vendor_direct_fulfillment_shipping_v1_models.GetPackingSlipResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPackingSlipForbidden creates a GetPackingSlipForbidden with default headers values
func NewGetPackingSlipForbidden() *GetPackingSlipForbidden {
	return &GetPackingSlipForbidden{}
}

/*
GetPackingSlipForbidden describes a response with status code 403, with default header values.

Indicates that access to the resource is forbidden. Possible reasons include Access Denied, Unauthorized, Expired Token, or Invalid Signature.
*/
type GetPackingSlipForbidden struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_shipping_v1_models.GetPackingSlipResponse
}

// IsSuccess returns true when this get packing slip forbidden response has a 2xx status code
func (o *GetPackingSlipForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get packing slip forbidden response has a 3xx status code
func (o *GetPackingSlipForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get packing slip forbidden response has a 4xx status code
func (o *GetPackingSlipForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get packing slip forbidden response has a 5xx status code
func (o *GetPackingSlipForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get packing slip forbidden response a status code equal to that given
func (o *GetPackingSlipForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetPackingSlipForbidden) Error() string {
	return fmt.Sprintf("[GET /vendor/directFulfillment/shipping/v1/packingSlips/{purchaseOrderNumber}][%d] getPackingSlipForbidden  %+v", 403, o.Payload)
}

func (o *GetPackingSlipForbidden) String() string {
	return fmt.Sprintf("[GET /vendor/directFulfillment/shipping/v1/packingSlips/{purchaseOrderNumber}][%d] getPackingSlipForbidden  %+v", 403, o.Payload)
}

func (o *GetPackingSlipForbidden) GetPayload() *vendor_direct_fulfillment_shipping_v1_models.GetPackingSlipResponse {
	return o.Payload
}

func (o *GetPackingSlipForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(vendor_direct_fulfillment_shipping_v1_models.GetPackingSlipResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPackingSlipNotFound creates a GetPackingSlipNotFound with default headers values
func NewGetPackingSlipNotFound() *GetPackingSlipNotFound {
	return &GetPackingSlipNotFound{}
}

/*
GetPackingSlipNotFound describes a response with status code 404, with default header values.

The resource specified does not exist.
*/
type GetPackingSlipNotFound struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_shipping_v1_models.GetPackingSlipResponse
}

// IsSuccess returns true when this get packing slip not found response has a 2xx status code
func (o *GetPackingSlipNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get packing slip not found response has a 3xx status code
func (o *GetPackingSlipNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get packing slip not found response has a 4xx status code
func (o *GetPackingSlipNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get packing slip not found response has a 5xx status code
func (o *GetPackingSlipNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get packing slip not found response a status code equal to that given
func (o *GetPackingSlipNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetPackingSlipNotFound) Error() string {
	return fmt.Sprintf("[GET /vendor/directFulfillment/shipping/v1/packingSlips/{purchaseOrderNumber}][%d] getPackingSlipNotFound  %+v", 404, o.Payload)
}

func (o *GetPackingSlipNotFound) String() string {
	return fmt.Sprintf("[GET /vendor/directFulfillment/shipping/v1/packingSlips/{purchaseOrderNumber}][%d] getPackingSlipNotFound  %+v", 404, o.Payload)
}

func (o *GetPackingSlipNotFound) GetPayload() *vendor_direct_fulfillment_shipping_v1_models.GetPackingSlipResponse {
	return o.Payload
}

func (o *GetPackingSlipNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(vendor_direct_fulfillment_shipping_v1_models.GetPackingSlipResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPackingSlipUnsupportedMediaType creates a GetPackingSlipUnsupportedMediaType with default headers values
func NewGetPackingSlipUnsupportedMediaType() *GetPackingSlipUnsupportedMediaType {
	return &GetPackingSlipUnsupportedMediaType{}
}

/*
GetPackingSlipUnsupportedMediaType describes a response with status code 415, with default header values.

The request payload is in an unsupported format.
*/
type GetPackingSlipUnsupportedMediaType struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_shipping_v1_models.GetPackingSlipResponse
}

// IsSuccess returns true when this get packing slip unsupported media type response has a 2xx status code
func (o *GetPackingSlipUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get packing slip unsupported media type response has a 3xx status code
func (o *GetPackingSlipUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get packing slip unsupported media type response has a 4xx status code
func (o *GetPackingSlipUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get packing slip unsupported media type response has a 5xx status code
func (o *GetPackingSlipUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get packing slip unsupported media type response a status code equal to that given
func (o *GetPackingSlipUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetPackingSlipUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /vendor/directFulfillment/shipping/v1/packingSlips/{purchaseOrderNumber}][%d] getPackingSlipUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetPackingSlipUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /vendor/directFulfillment/shipping/v1/packingSlips/{purchaseOrderNumber}][%d] getPackingSlipUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetPackingSlipUnsupportedMediaType) GetPayload() *vendor_direct_fulfillment_shipping_v1_models.GetPackingSlipResponse {
	return o.Payload
}

func (o *GetPackingSlipUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(vendor_direct_fulfillment_shipping_v1_models.GetPackingSlipResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPackingSlipTooManyRequests creates a GetPackingSlipTooManyRequests with default headers values
func NewGetPackingSlipTooManyRequests() *GetPackingSlipTooManyRequests {
	return &GetPackingSlipTooManyRequests{}
}

/*
GetPackingSlipTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type GetPackingSlipTooManyRequests struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_shipping_v1_models.GetPackingSlipResponse
}

// IsSuccess returns true when this get packing slip too many requests response has a 2xx status code
func (o *GetPackingSlipTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get packing slip too many requests response has a 3xx status code
func (o *GetPackingSlipTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get packing slip too many requests response has a 4xx status code
func (o *GetPackingSlipTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get packing slip too many requests response has a 5xx status code
func (o *GetPackingSlipTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get packing slip too many requests response a status code equal to that given
func (o *GetPackingSlipTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetPackingSlipTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /vendor/directFulfillment/shipping/v1/packingSlips/{purchaseOrderNumber}][%d] getPackingSlipTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetPackingSlipTooManyRequests) String() string {
	return fmt.Sprintf("[GET /vendor/directFulfillment/shipping/v1/packingSlips/{purchaseOrderNumber}][%d] getPackingSlipTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetPackingSlipTooManyRequests) GetPayload() *vendor_direct_fulfillment_shipping_v1_models.GetPackingSlipResponse {
	return o.Payload
}

func (o *GetPackingSlipTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(vendor_direct_fulfillment_shipping_v1_models.GetPackingSlipResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPackingSlipInternalServerError creates a GetPackingSlipInternalServerError with default headers values
func NewGetPackingSlipInternalServerError() *GetPackingSlipInternalServerError {
	return &GetPackingSlipInternalServerError{}
}

/*
GetPackingSlipInternalServerError describes a response with status code 500, with default header values.

Encountered an unexpected condition which prevented the server from fulfilling the request.
*/
type GetPackingSlipInternalServerError struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_shipping_v1_models.GetPackingSlipResponse
}

// IsSuccess returns true when this get packing slip internal server error response has a 2xx status code
func (o *GetPackingSlipInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get packing slip internal server error response has a 3xx status code
func (o *GetPackingSlipInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get packing slip internal server error response has a 4xx status code
func (o *GetPackingSlipInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get packing slip internal server error response has a 5xx status code
func (o *GetPackingSlipInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get packing slip internal server error response a status code equal to that given
func (o *GetPackingSlipInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetPackingSlipInternalServerError) Error() string {
	return fmt.Sprintf("[GET /vendor/directFulfillment/shipping/v1/packingSlips/{purchaseOrderNumber}][%d] getPackingSlipInternalServerError  %+v", 500, o.Payload)
}

func (o *GetPackingSlipInternalServerError) String() string {
	return fmt.Sprintf("[GET /vendor/directFulfillment/shipping/v1/packingSlips/{purchaseOrderNumber}][%d] getPackingSlipInternalServerError  %+v", 500, o.Payload)
}

func (o *GetPackingSlipInternalServerError) GetPayload() *vendor_direct_fulfillment_shipping_v1_models.GetPackingSlipResponse {
	return o.Payload
}

func (o *GetPackingSlipInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(vendor_direct_fulfillment_shipping_v1_models.GetPackingSlipResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPackingSlipServiceUnavailable creates a GetPackingSlipServiceUnavailable with default headers values
func NewGetPackingSlipServiceUnavailable() *GetPackingSlipServiceUnavailable {
	return &GetPackingSlipServiceUnavailable{}
}

/*
GetPackingSlipServiceUnavailable describes a response with status code 503, with default header values.

Temporary overloading or maintenance of the server.
*/
type GetPackingSlipServiceUnavailable struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_shipping_v1_models.GetPackingSlipResponse
}

// IsSuccess returns true when this get packing slip service unavailable response has a 2xx status code
func (o *GetPackingSlipServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get packing slip service unavailable response has a 3xx status code
func (o *GetPackingSlipServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get packing slip service unavailable response has a 4xx status code
func (o *GetPackingSlipServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get packing slip service unavailable response has a 5xx status code
func (o *GetPackingSlipServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get packing slip service unavailable response a status code equal to that given
func (o *GetPackingSlipServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetPackingSlipServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /vendor/directFulfillment/shipping/v1/packingSlips/{purchaseOrderNumber}][%d] getPackingSlipServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetPackingSlipServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /vendor/directFulfillment/shipping/v1/packingSlips/{purchaseOrderNumber}][%d] getPackingSlipServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetPackingSlipServiceUnavailable) GetPayload() *vendor_direct_fulfillment_shipping_v1_models.GetPackingSlipResponse {
	return o.Payload
}

func (o *GetPackingSlipServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(vendor_direct_fulfillment_shipping_v1_models.GetPackingSlipResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
