// Code generated by go-swagger; DO NOT EDIT.

package vendor_orders

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/renabled/amzn-sp-api-go/api/vendorOrders/vendor_orders_models"
)

// GetPurchaseOrdersReader is a Reader for the GetPurchaseOrders structure.
type GetPurchaseOrdersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPurchaseOrdersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPurchaseOrdersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetPurchaseOrdersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetPurchaseOrdersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetPurchaseOrdersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetPurchaseOrdersUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetPurchaseOrdersTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetPurchaseOrdersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetPurchaseOrdersServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetPurchaseOrdersOK creates a GetPurchaseOrdersOK with default headers values
func NewGetPurchaseOrdersOK() *GetPurchaseOrdersOK {
	return &GetPurchaseOrdersOK{}
}

/*
GetPurchaseOrdersOK describes a response with status code 200, with default header values.

Success.
*/
type GetPurchaseOrdersOK struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_orders_models.GetPurchaseOrdersResponse
}

// IsSuccess returns true when this get purchase orders o k response has a 2xx status code
func (o *GetPurchaseOrdersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get purchase orders o k response has a 3xx status code
func (o *GetPurchaseOrdersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get purchase orders o k response has a 4xx status code
func (o *GetPurchaseOrdersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get purchase orders o k response has a 5xx status code
func (o *GetPurchaseOrdersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get purchase orders o k response a status code equal to that given
func (o *GetPurchaseOrdersOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetPurchaseOrdersOK) Error() string {
	return fmt.Sprintf("[GET /vendor/orders/v1/purchaseOrders][%d] getPurchaseOrdersOK  %+v", 200, o.Payload)
}

func (o *GetPurchaseOrdersOK) String() string {
	return fmt.Sprintf("[GET /vendor/orders/v1/purchaseOrders][%d] getPurchaseOrdersOK  %+v", 200, o.Payload)
}

func (o *GetPurchaseOrdersOK) GetPayload() *vendor_orders_models.GetPurchaseOrdersResponse {
	return o.Payload
}

func (o *GetPurchaseOrdersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(vendor_orders_models.GetPurchaseOrdersResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPurchaseOrdersBadRequest creates a GetPurchaseOrdersBadRequest with default headers values
func NewGetPurchaseOrdersBadRequest() *GetPurchaseOrdersBadRequest {
	return &GetPurchaseOrdersBadRequest{}
}

/*
GetPurchaseOrdersBadRequest describes a response with status code 400, with default header values.

Request has missing or invalid parameters and cannot be parsed.
*/
type GetPurchaseOrdersBadRequest struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_orders_models.GetPurchaseOrdersResponse
}

// IsSuccess returns true when this get purchase orders bad request response has a 2xx status code
func (o *GetPurchaseOrdersBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get purchase orders bad request response has a 3xx status code
func (o *GetPurchaseOrdersBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get purchase orders bad request response has a 4xx status code
func (o *GetPurchaseOrdersBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get purchase orders bad request response has a 5xx status code
func (o *GetPurchaseOrdersBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get purchase orders bad request response a status code equal to that given
func (o *GetPurchaseOrdersBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetPurchaseOrdersBadRequest) Error() string {
	return fmt.Sprintf("[GET /vendor/orders/v1/purchaseOrders][%d] getPurchaseOrdersBadRequest  %+v", 400, o.Payload)
}

func (o *GetPurchaseOrdersBadRequest) String() string {
	return fmt.Sprintf("[GET /vendor/orders/v1/purchaseOrders][%d] getPurchaseOrdersBadRequest  %+v", 400, o.Payload)
}

func (o *GetPurchaseOrdersBadRequest) GetPayload() *vendor_orders_models.GetPurchaseOrdersResponse {
	return o.Payload
}

func (o *GetPurchaseOrdersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(vendor_orders_models.GetPurchaseOrdersResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPurchaseOrdersForbidden creates a GetPurchaseOrdersForbidden with default headers values
func NewGetPurchaseOrdersForbidden() *GetPurchaseOrdersForbidden {
	return &GetPurchaseOrdersForbidden{}
}

/*
GetPurchaseOrdersForbidden describes a response with status code 403, with default header values.

Indicates that access to the resource is forbidden. Possible reasons include `Access Denied`, `Unauthorized`, `Expired Token`, or `Invalid Signature`.
*/
type GetPurchaseOrdersForbidden struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_orders_models.GetPurchaseOrdersResponse
}

// IsSuccess returns true when this get purchase orders forbidden response has a 2xx status code
func (o *GetPurchaseOrdersForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get purchase orders forbidden response has a 3xx status code
func (o *GetPurchaseOrdersForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get purchase orders forbidden response has a 4xx status code
func (o *GetPurchaseOrdersForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get purchase orders forbidden response has a 5xx status code
func (o *GetPurchaseOrdersForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get purchase orders forbidden response a status code equal to that given
func (o *GetPurchaseOrdersForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetPurchaseOrdersForbidden) Error() string {
	return fmt.Sprintf("[GET /vendor/orders/v1/purchaseOrders][%d] getPurchaseOrdersForbidden  %+v", 403, o.Payload)
}

func (o *GetPurchaseOrdersForbidden) String() string {
	return fmt.Sprintf("[GET /vendor/orders/v1/purchaseOrders][%d] getPurchaseOrdersForbidden  %+v", 403, o.Payload)
}

func (o *GetPurchaseOrdersForbidden) GetPayload() *vendor_orders_models.GetPurchaseOrdersResponse {
	return o.Payload
}

func (o *GetPurchaseOrdersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(vendor_orders_models.GetPurchaseOrdersResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPurchaseOrdersNotFound creates a GetPurchaseOrdersNotFound with default headers values
func NewGetPurchaseOrdersNotFound() *GetPurchaseOrdersNotFound {
	return &GetPurchaseOrdersNotFound{}
}

/*
GetPurchaseOrdersNotFound describes a response with status code 404, with default header values.

The resource specified does not exist.
*/
type GetPurchaseOrdersNotFound struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_orders_models.GetPurchaseOrdersResponse
}

// IsSuccess returns true when this get purchase orders not found response has a 2xx status code
func (o *GetPurchaseOrdersNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get purchase orders not found response has a 3xx status code
func (o *GetPurchaseOrdersNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get purchase orders not found response has a 4xx status code
func (o *GetPurchaseOrdersNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get purchase orders not found response has a 5xx status code
func (o *GetPurchaseOrdersNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get purchase orders not found response a status code equal to that given
func (o *GetPurchaseOrdersNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetPurchaseOrdersNotFound) Error() string {
	return fmt.Sprintf("[GET /vendor/orders/v1/purchaseOrders][%d] getPurchaseOrdersNotFound  %+v", 404, o.Payload)
}

func (o *GetPurchaseOrdersNotFound) String() string {
	return fmt.Sprintf("[GET /vendor/orders/v1/purchaseOrders][%d] getPurchaseOrdersNotFound  %+v", 404, o.Payload)
}

func (o *GetPurchaseOrdersNotFound) GetPayload() *vendor_orders_models.GetPurchaseOrdersResponse {
	return o.Payload
}

func (o *GetPurchaseOrdersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(vendor_orders_models.GetPurchaseOrdersResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPurchaseOrdersUnsupportedMediaType creates a GetPurchaseOrdersUnsupportedMediaType with default headers values
func NewGetPurchaseOrdersUnsupportedMediaType() *GetPurchaseOrdersUnsupportedMediaType {
	return &GetPurchaseOrdersUnsupportedMediaType{}
}

/*
GetPurchaseOrdersUnsupportedMediaType describes a response with status code 415, with default header values.

The request payload is in an unsupported format.
*/
type GetPurchaseOrdersUnsupportedMediaType struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_orders_models.GetPurchaseOrdersResponse
}

// IsSuccess returns true when this get purchase orders unsupported media type response has a 2xx status code
func (o *GetPurchaseOrdersUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get purchase orders unsupported media type response has a 3xx status code
func (o *GetPurchaseOrdersUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get purchase orders unsupported media type response has a 4xx status code
func (o *GetPurchaseOrdersUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get purchase orders unsupported media type response has a 5xx status code
func (o *GetPurchaseOrdersUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get purchase orders unsupported media type response a status code equal to that given
func (o *GetPurchaseOrdersUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetPurchaseOrdersUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /vendor/orders/v1/purchaseOrders][%d] getPurchaseOrdersUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetPurchaseOrdersUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /vendor/orders/v1/purchaseOrders][%d] getPurchaseOrdersUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetPurchaseOrdersUnsupportedMediaType) GetPayload() *vendor_orders_models.GetPurchaseOrdersResponse {
	return o.Payload
}

func (o *GetPurchaseOrdersUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(vendor_orders_models.GetPurchaseOrdersResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPurchaseOrdersTooManyRequests creates a GetPurchaseOrdersTooManyRequests with default headers values
func NewGetPurchaseOrdersTooManyRequests() *GetPurchaseOrdersTooManyRequests {
	return &GetPurchaseOrdersTooManyRequests{}
}

/*
GetPurchaseOrdersTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type GetPurchaseOrdersTooManyRequests struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_orders_models.GetPurchaseOrdersResponse
}

// IsSuccess returns true when this get purchase orders too many requests response has a 2xx status code
func (o *GetPurchaseOrdersTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get purchase orders too many requests response has a 3xx status code
func (o *GetPurchaseOrdersTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get purchase orders too many requests response has a 4xx status code
func (o *GetPurchaseOrdersTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get purchase orders too many requests response has a 5xx status code
func (o *GetPurchaseOrdersTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get purchase orders too many requests response a status code equal to that given
func (o *GetPurchaseOrdersTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetPurchaseOrdersTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /vendor/orders/v1/purchaseOrders][%d] getPurchaseOrdersTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetPurchaseOrdersTooManyRequests) String() string {
	return fmt.Sprintf("[GET /vendor/orders/v1/purchaseOrders][%d] getPurchaseOrdersTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetPurchaseOrdersTooManyRequests) GetPayload() *vendor_orders_models.GetPurchaseOrdersResponse {
	return o.Payload
}

func (o *GetPurchaseOrdersTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(vendor_orders_models.GetPurchaseOrdersResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPurchaseOrdersInternalServerError creates a GetPurchaseOrdersInternalServerError with default headers values
func NewGetPurchaseOrdersInternalServerError() *GetPurchaseOrdersInternalServerError {
	return &GetPurchaseOrdersInternalServerError{}
}

/*
GetPurchaseOrdersInternalServerError describes a response with status code 500, with default header values.

An unexpected condition occurred that prevented the server from fulfilling the request.
*/
type GetPurchaseOrdersInternalServerError struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_orders_models.GetPurchaseOrdersResponse
}

// IsSuccess returns true when this get purchase orders internal server error response has a 2xx status code
func (o *GetPurchaseOrdersInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get purchase orders internal server error response has a 3xx status code
func (o *GetPurchaseOrdersInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get purchase orders internal server error response has a 4xx status code
func (o *GetPurchaseOrdersInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get purchase orders internal server error response has a 5xx status code
func (o *GetPurchaseOrdersInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get purchase orders internal server error response a status code equal to that given
func (o *GetPurchaseOrdersInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetPurchaseOrdersInternalServerError) Error() string {
	return fmt.Sprintf("[GET /vendor/orders/v1/purchaseOrders][%d] getPurchaseOrdersInternalServerError  %+v", 500, o.Payload)
}

func (o *GetPurchaseOrdersInternalServerError) String() string {
	return fmt.Sprintf("[GET /vendor/orders/v1/purchaseOrders][%d] getPurchaseOrdersInternalServerError  %+v", 500, o.Payload)
}

func (o *GetPurchaseOrdersInternalServerError) GetPayload() *vendor_orders_models.GetPurchaseOrdersResponse {
	return o.Payload
}

func (o *GetPurchaseOrdersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(vendor_orders_models.GetPurchaseOrdersResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPurchaseOrdersServiceUnavailable creates a GetPurchaseOrdersServiceUnavailable with default headers values
func NewGetPurchaseOrdersServiceUnavailable() *GetPurchaseOrdersServiceUnavailable {
	return &GetPurchaseOrdersServiceUnavailable{}
}

/*
GetPurchaseOrdersServiceUnavailable describes a response with status code 503, with default header values.

Temporary overloading or maintenance of the server.
*/
type GetPurchaseOrdersServiceUnavailable struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_orders_models.GetPurchaseOrdersResponse
}

// IsSuccess returns true when this get purchase orders service unavailable response has a 2xx status code
func (o *GetPurchaseOrdersServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get purchase orders service unavailable response has a 3xx status code
func (o *GetPurchaseOrdersServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get purchase orders service unavailable response has a 4xx status code
func (o *GetPurchaseOrdersServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get purchase orders service unavailable response has a 5xx status code
func (o *GetPurchaseOrdersServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get purchase orders service unavailable response a status code equal to that given
func (o *GetPurchaseOrdersServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetPurchaseOrdersServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /vendor/orders/v1/purchaseOrders][%d] getPurchaseOrdersServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetPurchaseOrdersServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /vendor/orders/v1/purchaseOrders][%d] getPurchaseOrdersServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetPurchaseOrdersServiceUnavailable) GetPayload() *vendor_orders_models.GetPurchaseOrdersResponse {
	return o.Payload
}

func (o *GetPurchaseOrdersServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(vendor_orders_models.GetPurchaseOrdersResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
