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

// GetPackingSlipsReader is a Reader for the GetPackingSlips structure.
type GetPackingSlipsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPackingSlipsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPackingSlipsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetPackingSlipsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetPackingSlipsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetPackingSlipsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetPackingSlipsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetPackingSlipsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetPackingSlipsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetPackingSlipsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetPackingSlipsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetPackingSlipsOK creates a GetPackingSlipsOK with default headers values
func NewGetPackingSlipsOK() *GetPackingSlipsOK {
	return &GetPackingSlipsOK{}
}

/*
GetPackingSlipsOK describes a response with status code 200, with default header values.

Success.
*/
type GetPackingSlipsOK struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_shipping_2021_12_28_models.PackingSlipList
}

// IsSuccess returns true when this get packing slips o k response has a 2xx status code
func (o *GetPackingSlipsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get packing slips o k response has a 3xx status code
func (o *GetPackingSlipsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get packing slips o k response has a 4xx status code
func (o *GetPackingSlipsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get packing slips o k response has a 5xx status code
func (o *GetPackingSlipsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get packing slips o k response a status code equal to that given
func (o *GetPackingSlipsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetPackingSlipsOK) Error() string {
	return fmt.Sprintf("[GET /vendor/directFulfillment/shipping/2021-12-28/packingSlips][%d] getPackingSlipsOK  %+v", 200, o.Payload)
}

func (o *GetPackingSlipsOK) String() string {
	return fmt.Sprintf("[GET /vendor/directFulfillment/shipping/2021-12-28/packingSlips][%d] getPackingSlipsOK  %+v", 200, o.Payload)
}

func (o *GetPackingSlipsOK) GetPayload() *vendor_direct_fulfillment_shipping_2021_12_28_models.PackingSlipList {
	return o.Payload
}

func (o *GetPackingSlipsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(vendor_direct_fulfillment_shipping_2021_12_28_models.PackingSlipList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPackingSlipsBadRequest creates a GetPackingSlipsBadRequest with default headers values
func NewGetPackingSlipsBadRequest() *GetPackingSlipsBadRequest {
	return &GetPackingSlipsBadRequest{}
}

/*
GetPackingSlipsBadRequest describes a response with status code 400, with default header values.

Request has missing or invalid parameters and cannot be parsed.
*/
type GetPackingSlipsBadRequest struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_shipping_2021_12_28_models.ErrorList
}

// IsSuccess returns true when this get packing slips bad request response has a 2xx status code
func (o *GetPackingSlipsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get packing slips bad request response has a 3xx status code
func (o *GetPackingSlipsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get packing slips bad request response has a 4xx status code
func (o *GetPackingSlipsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get packing slips bad request response has a 5xx status code
func (o *GetPackingSlipsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get packing slips bad request response a status code equal to that given
func (o *GetPackingSlipsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetPackingSlipsBadRequest) Error() string {
	return fmt.Sprintf("[GET /vendor/directFulfillment/shipping/2021-12-28/packingSlips][%d] getPackingSlipsBadRequest  %+v", 400, o.Payload)
}

func (o *GetPackingSlipsBadRequest) String() string {
	return fmt.Sprintf("[GET /vendor/directFulfillment/shipping/2021-12-28/packingSlips][%d] getPackingSlipsBadRequest  %+v", 400, o.Payload)
}

func (o *GetPackingSlipsBadRequest) GetPayload() *vendor_direct_fulfillment_shipping_2021_12_28_models.ErrorList {
	return o.Payload
}

func (o *GetPackingSlipsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetPackingSlipsUnauthorized creates a GetPackingSlipsUnauthorized with default headers values
func NewGetPackingSlipsUnauthorized() *GetPackingSlipsUnauthorized {
	return &GetPackingSlipsUnauthorized{}
}

/*
GetPackingSlipsUnauthorized describes a response with status code 401, with default header values.

The request's Authorization header is not formatted correctly or does not contain a valid token.
*/
type GetPackingSlipsUnauthorized struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_shipping_2021_12_28_models.ErrorList
}

// IsSuccess returns true when this get packing slips unauthorized response has a 2xx status code
func (o *GetPackingSlipsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get packing slips unauthorized response has a 3xx status code
func (o *GetPackingSlipsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get packing slips unauthorized response has a 4xx status code
func (o *GetPackingSlipsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get packing slips unauthorized response has a 5xx status code
func (o *GetPackingSlipsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get packing slips unauthorized response a status code equal to that given
func (o *GetPackingSlipsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetPackingSlipsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /vendor/directFulfillment/shipping/2021-12-28/packingSlips][%d] getPackingSlipsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetPackingSlipsUnauthorized) String() string {
	return fmt.Sprintf("[GET /vendor/directFulfillment/shipping/2021-12-28/packingSlips][%d] getPackingSlipsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetPackingSlipsUnauthorized) GetPayload() *vendor_direct_fulfillment_shipping_2021_12_28_models.ErrorList {
	return o.Payload
}

func (o *GetPackingSlipsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetPackingSlipsForbidden creates a GetPackingSlipsForbidden with default headers values
func NewGetPackingSlipsForbidden() *GetPackingSlipsForbidden {
	return &GetPackingSlipsForbidden{}
}

/*
GetPackingSlipsForbidden describes a response with status code 403, with default header values.

Indicates that access to the resource is forbidden. Possible reasons include Access Denied, Unauthorized, Expired Token, or Invalid Signature.
*/
type GetPackingSlipsForbidden struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_shipping_2021_12_28_models.ErrorList
}

// IsSuccess returns true when this get packing slips forbidden response has a 2xx status code
func (o *GetPackingSlipsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get packing slips forbidden response has a 3xx status code
func (o *GetPackingSlipsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get packing slips forbidden response has a 4xx status code
func (o *GetPackingSlipsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get packing slips forbidden response has a 5xx status code
func (o *GetPackingSlipsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get packing slips forbidden response a status code equal to that given
func (o *GetPackingSlipsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetPackingSlipsForbidden) Error() string {
	return fmt.Sprintf("[GET /vendor/directFulfillment/shipping/2021-12-28/packingSlips][%d] getPackingSlipsForbidden  %+v", 403, o.Payload)
}

func (o *GetPackingSlipsForbidden) String() string {
	return fmt.Sprintf("[GET /vendor/directFulfillment/shipping/2021-12-28/packingSlips][%d] getPackingSlipsForbidden  %+v", 403, o.Payload)
}

func (o *GetPackingSlipsForbidden) GetPayload() *vendor_direct_fulfillment_shipping_2021_12_28_models.ErrorList {
	return o.Payload
}

func (o *GetPackingSlipsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetPackingSlipsNotFound creates a GetPackingSlipsNotFound with default headers values
func NewGetPackingSlipsNotFound() *GetPackingSlipsNotFound {
	return &GetPackingSlipsNotFound{}
}

/*
GetPackingSlipsNotFound describes a response with status code 404, with default header values.

The resource specified does not exist.
*/
type GetPackingSlipsNotFound struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_shipping_2021_12_28_models.ErrorList
}

// IsSuccess returns true when this get packing slips not found response has a 2xx status code
func (o *GetPackingSlipsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get packing slips not found response has a 3xx status code
func (o *GetPackingSlipsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get packing slips not found response has a 4xx status code
func (o *GetPackingSlipsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get packing slips not found response has a 5xx status code
func (o *GetPackingSlipsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get packing slips not found response a status code equal to that given
func (o *GetPackingSlipsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetPackingSlipsNotFound) Error() string {
	return fmt.Sprintf("[GET /vendor/directFulfillment/shipping/2021-12-28/packingSlips][%d] getPackingSlipsNotFound  %+v", 404, o.Payload)
}

func (o *GetPackingSlipsNotFound) String() string {
	return fmt.Sprintf("[GET /vendor/directFulfillment/shipping/2021-12-28/packingSlips][%d] getPackingSlipsNotFound  %+v", 404, o.Payload)
}

func (o *GetPackingSlipsNotFound) GetPayload() *vendor_direct_fulfillment_shipping_2021_12_28_models.ErrorList {
	return o.Payload
}

func (o *GetPackingSlipsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetPackingSlipsUnsupportedMediaType creates a GetPackingSlipsUnsupportedMediaType with default headers values
func NewGetPackingSlipsUnsupportedMediaType() *GetPackingSlipsUnsupportedMediaType {
	return &GetPackingSlipsUnsupportedMediaType{}
}

/*
GetPackingSlipsUnsupportedMediaType describes a response with status code 415, with default header values.

The request payload is in an unsupported format.
*/
type GetPackingSlipsUnsupportedMediaType struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_shipping_2021_12_28_models.ErrorList
}

// IsSuccess returns true when this get packing slips unsupported media type response has a 2xx status code
func (o *GetPackingSlipsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get packing slips unsupported media type response has a 3xx status code
func (o *GetPackingSlipsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get packing slips unsupported media type response has a 4xx status code
func (o *GetPackingSlipsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get packing slips unsupported media type response has a 5xx status code
func (o *GetPackingSlipsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get packing slips unsupported media type response a status code equal to that given
func (o *GetPackingSlipsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetPackingSlipsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /vendor/directFulfillment/shipping/2021-12-28/packingSlips][%d] getPackingSlipsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetPackingSlipsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /vendor/directFulfillment/shipping/2021-12-28/packingSlips][%d] getPackingSlipsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetPackingSlipsUnsupportedMediaType) GetPayload() *vendor_direct_fulfillment_shipping_2021_12_28_models.ErrorList {
	return o.Payload
}

func (o *GetPackingSlipsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetPackingSlipsTooManyRequests creates a GetPackingSlipsTooManyRequests with default headers values
func NewGetPackingSlipsTooManyRequests() *GetPackingSlipsTooManyRequests {
	return &GetPackingSlipsTooManyRequests{}
}

/*
GetPackingSlipsTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type GetPackingSlipsTooManyRequests struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_shipping_2021_12_28_models.ErrorList
}

// IsSuccess returns true when this get packing slips too many requests response has a 2xx status code
func (o *GetPackingSlipsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get packing slips too many requests response has a 3xx status code
func (o *GetPackingSlipsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get packing slips too many requests response has a 4xx status code
func (o *GetPackingSlipsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get packing slips too many requests response has a 5xx status code
func (o *GetPackingSlipsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get packing slips too many requests response a status code equal to that given
func (o *GetPackingSlipsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetPackingSlipsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /vendor/directFulfillment/shipping/2021-12-28/packingSlips][%d] getPackingSlipsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetPackingSlipsTooManyRequests) String() string {
	return fmt.Sprintf("[GET /vendor/directFulfillment/shipping/2021-12-28/packingSlips][%d] getPackingSlipsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetPackingSlipsTooManyRequests) GetPayload() *vendor_direct_fulfillment_shipping_2021_12_28_models.ErrorList {
	return o.Payload
}

func (o *GetPackingSlipsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetPackingSlipsInternalServerError creates a GetPackingSlipsInternalServerError with default headers values
func NewGetPackingSlipsInternalServerError() *GetPackingSlipsInternalServerError {
	return &GetPackingSlipsInternalServerError{}
}

/*
GetPackingSlipsInternalServerError describes a response with status code 500, with default header values.

An unexpected condition occurred that prevented the server from fulfilling the request.
*/
type GetPackingSlipsInternalServerError struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_shipping_2021_12_28_models.ErrorList
}

// IsSuccess returns true when this get packing slips internal server error response has a 2xx status code
func (o *GetPackingSlipsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get packing slips internal server error response has a 3xx status code
func (o *GetPackingSlipsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get packing slips internal server error response has a 4xx status code
func (o *GetPackingSlipsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get packing slips internal server error response has a 5xx status code
func (o *GetPackingSlipsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get packing slips internal server error response a status code equal to that given
func (o *GetPackingSlipsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetPackingSlipsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /vendor/directFulfillment/shipping/2021-12-28/packingSlips][%d] getPackingSlipsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetPackingSlipsInternalServerError) String() string {
	return fmt.Sprintf("[GET /vendor/directFulfillment/shipping/2021-12-28/packingSlips][%d] getPackingSlipsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetPackingSlipsInternalServerError) GetPayload() *vendor_direct_fulfillment_shipping_2021_12_28_models.ErrorList {
	return o.Payload
}

func (o *GetPackingSlipsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetPackingSlipsServiceUnavailable creates a GetPackingSlipsServiceUnavailable with default headers values
func NewGetPackingSlipsServiceUnavailable() *GetPackingSlipsServiceUnavailable {
	return &GetPackingSlipsServiceUnavailable{}
}

/*
GetPackingSlipsServiceUnavailable describes a response with status code 503, with default header values.

Temporary overloading or maintenance of the server.
*/
type GetPackingSlipsServiceUnavailable struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_shipping_2021_12_28_models.ErrorList
}

// IsSuccess returns true when this get packing slips service unavailable response has a 2xx status code
func (o *GetPackingSlipsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get packing slips service unavailable response has a 3xx status code
func (o *GetPackingSlipsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get packing slips service unavailable response has a 4xx status code
func (o *GetPackingSlipsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get packing slips service unavailable response has a 5xx status code
func (o *GetPackingSlipsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get packing slips service unavailable response a status code equal to that given
func (o *GetPackingSlipsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetPackingSlipsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /vendor/directFulfillment/shipping/2021-12-28/packingSlips][%d] getPackingSlipsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetPackingSlipsServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /vendor/directFulfillment/shipping/2021-12-28/packingSlips][%d] getPackingSlipsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetPackingSlipsServiceUnavailable) GetPayload() *vendor_direct_fulfillment_shipping_2021_12_28_models.ErrorList {
	return o.Payload
}

func (o *GetPackingSlipsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
