// Code generated by go-swagger; DO NOT EDIT.

package customer_invoices

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/xamandar/amzn-sp-api-go/api/vendorDirectFulfillmentShippingV1/vendor_direct_fulfillment_shipping_v1_models"
)

// GetCustomerInvoiceReader is a Reader for the GetCustomerInvoice structure.
type GetCustomerInvoiceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCustomerInvoiceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCustomerInvoiceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetCustomerInvoiceBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetCustomerInvoiceUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetCustomerInvoiceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetCustomerInvoiceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetCustomerInvoiceUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetCustomerInvoiceTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetCustomerInvoiceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetCustomerInvoiceServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetCustomerInvoiceOK creates a GetCustomerInvoiceOK with default headers values
func NewGetCustomerInvoiceOK() *GetCustomerInvoiceOK {
	return &GetCustomerInvoiceOK{}
}

/*
GetCustomerInvoiceOK describes a response with status code 200, with default header values.

Success.
*/
type GetCustomerInvoiceOK struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_shipping_v1_models.GetCustomerInvoiceResponse
}

// IsSuccess returns true when this get customer invoice o k response has a 2xx status code
func (o *GetCustomerInvoiceOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get customer invoice o k response has a 3xx status code
func (o *GetCustomerInvoiceOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get customer invoice o k response has a 4xx status code
func (o *GetCustomerInvoiceOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get customer invoice o k response has a 5xx status code
func (o *GetCustomerInvoiceOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get customer invoice o k response a status code equal to that given
func (o *GetCustomerInvoiceOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetCustomerInvoiceOK) Error() string {
	return fmt.Sprintf("[GET /vendor/directFulfillment/shipping/v1/customerInvoices/{purchaseOrderNumber}][%d] getCustomerInvoiceOK  %+v", 200, o.Payload)
}

func (o *GetCustomerInvoiceOK) String() string {
	return fmt.Sprintf("[GET /vendor/directFulfillment/shipping/v1/customerInvoices/{purchaseOrderNumber}][%d] getCustomerInvoiceOK  %+v", 200, o.Payload)
}

func (o *GetCustomerInvoiceOK) GetPayload() *vendor_direct_fulfillment_shipping_v1_models.GetCustomerInvoiceResponse {
	return o.Payload
}

func (o *GetCustomerInvoiceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(vendor_direct_fulfillment_shipping_v1_models.GetCustomerInvoiceResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCustomerInvoiceBadRequest creates a GetCustomerInvoiceBadRequest with default headers values
func NewGetCustomerInvoiceBadRequest() *GetCustomerInvoiceBadRequest {
	return &GetCustomerInvoiceBadRequest{}
}

/*
GetCustomerInvoiceBadRequest describes a response with status code 400, with default header values.

Request has missing or invalid parameters and cannot be parsed.
*/
type GetCustomerInvoiceBadRequest struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_shipping_v1_models.GetCustomerInvoiceResponse
}

// IsSuccess returns true when this get customer invoice bad request response has a 2xx status code
func (o *GetCustomerInvoiceBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get customer invoice bad request response has a 3xx status code
func (o *GetCustomerInvoiceBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get customer invoice bad request response has a 4xx status code
func (o *GetCustomerInvoiceBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get customer invoice bad request response has a 5xx status code
func (o *GetCustomerInvoiceBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get customer invoice bad request response a status code equal to that given
func (o *GetCustomerInvoiceBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetCustomerInvoiceBadRequest) Error() string {
	return fmt.Sprintf("[GET /vendor/directFulfillment/shipping/v1/customerInvoices/{purchaseOrderNumber}][%d] getCustomerInvoiceBadRequest  %+v", 400, o.Payload)
}

func (o *GetCustomerInvoiceBadRequest) String() string {
	return fmt.Sprintf("[GET /vendor/directFulfillment/shipping/v1/customerInvoices/{purchaseOrderNumber}][%d] getCustomerInvoiceBadRequest  %+v", 400, o.Payload)
}

func (o *GetCustomerInvoiceBadRequest) GetPayload() *vendor_direct_fulfillment_shipping_v1_models.GetCustomerInvoiceResponse {
	return o.Payload
}

func (o *GetCustomerInvoiceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(vendor_direct_fulfillment_shipping_v1_models.GetCustomerInvoiceResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCustomerInvoiceUnauthorized creates a GetCustomerInvoiceUnauthorized with default headers values
func NewGetCustomerInvoiceUnauthorized() *GetCustomerInvoiceUnauthorized {
	return &GetCustomerInvoiceUnauthorized{}
}

/*
GetCustomerInvoiceUnauthorized describes a response with status code 401, with default header values.

The request's Authorization header is not formatted correctly or does not contain a valid token.
*/
type GetCustomerInvoiceUnauthorized struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_shipping_v1_models.GetCustomerInvoiceResponse
}

// IsSuccess returns true when this get customer invoice unauthorized response has a 2xx status code
func (o *GetCustomerInvoiceUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get customer invoice unauthorized response has a 3xx status code
func (o *GetCustomerInvoiceUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get customer invoice unauthorized response has a 4xx status code
func (o *GetCustomerInvoiceUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get customer invoice unauthorized response has a 5xx status code
func (o *GetCustomerInvoiceUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get customer invoice unauthorized response a status code equal to that given
func (o *GetCustomerInvoiceUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetCustomerInvoiceUnauthorized) Error() string {
	return fmt.Sprintf("[GET /vendor/directFulfillment/shipping/v1/customerInvoices/{purchaseOrderNumber}][%d] getCustomerInvoiceUnauthorized  %+v", 401, o.Payload)
}

func (o *GetCustomerInvoiceUnauthorized) String() string {
	return fmt.Sprintf("[GET /vendor/directFulfillment/shipping/v1/customerInvoices/{purchaseOrderNumber}][%d] getCustomerInvoiceUnauthorized  %+v", 401, o.Payload)
}

func (o *GetCustomerInvoiceUnauthorized) GetPayload() *vendor_direct_fulfillment_shipping_v1_models.GetCustomerInvoiceResponse {
	return o.Payload
}

func (o *GetCustomerInvoiceUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(vendor_direct_fulfillment_shipping_v1_models.GetCustomerInvoiceResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCustomerInvoiceForbidden creates a GetCustomerInvoiceForbidden with default headers values
func NewGetCustomerInvoiceForbidden() *GetCustomerInvoiceForbidden {
	return &GetCustomerInvoiceForbidden{}
}

/*
GetCustomerInvoiceForbidden describes a response with status code 403, with default header values.

Indicates that access to the resource is forbidden. Possible reasons include Access Denied, Unauthorized, Expired Token, or Invalid Signature.
*/
type GetCustomerInvoiceForbidden struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_shipping_v1_models.GetCustomerInvoiceResponse
}

// IsSuccess returns true when this get customer invoice forbidden response has a 2xx status code
func (o *GetCustomerInvoiceForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get customer invoice forbidden response has a 3xx status code
func (o *GetCustomerInvoiceForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get customer invoice forbidden response has a 4xx status code
func (o *GetCustomerInvoiceForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get customer invoice forbidden response has a 5xx status code
func (o *GetCustomerInvoiceForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get customer invoice forbidden response a status code equal to that given
func (o *GetCustomerInvoiceForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetCustomerInvoiceForbidden) Error() string {
	return fmt.Sprintf("[GET /vendor/directFulfillment/shipping/v1/customerInvoices/{purchaseOrderNumber}][%d] getCustomerInvoiceForbidden  %+v", 403, o.Payload)
}

func (o *GetCustomerInvoiceForbidden) String() string {
	return fmt.Sprintf("[GET /vendor/directFulfillment/shipping/v1/customerInvoices/{purchaseOrderNumber}][%d] getCustomerInvoiceForbidden  %+v", 403, o.Payload)
}

func (o *GetCustomerInvoiceForbidden) GetPayload() *vendor_direct_fulfillment_shipping_v1_models.GetCustomerInvoiceResponse {
	return o.Payload
}

func (o *GetCustomerInvoiceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(vendor_direct_fulfillment_shipping_v1_models.GetCustomerInvoiceResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCustomerInvoiceNotFound creates a GetCustomerInvoiceNotFound with default headers values
func NewGetCustomerInvoiceNotFound() *GetCustomerInvoiceNotFound {
	return &GetCustomerInvoiceNotFound{}
}

/*
GetCustomerInvoiceNotFound describes a response with status code 404, with default header values.

The resource specified does not exist.
*/
type GetCustomerInvoiceNotFound struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_shipping_v1_models.GetCustomerInvoiceResponse
}

// IsSuccess returns true when this get customer invoice not found response has a 2xx status code
func (o *GetCustomerInvoiceNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get customer invoice not found response has a 3xx status code
func (o *GetCustomerInvoiceNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get customer invoice not found response has a 4xx status code
func (o *GetCustomerInvoiceNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get customer invoice not found response has a 5xx status code
func (o *GetCustomerInvoiceNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get customer invoice not found response a status code equal to that given
func (o *GetCustomerInvoiceNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetCustomerInvoiceNotFound) Error() string {
	return fmt.Sprintf("[GET /vendor/directFulfillment/shipping/v1/customerInvoices/{purchaseOrderNumber}][%d] getCustomerInvoiceNotFound  %+v", 404, o.Payload)
}

func (o *GetCustomerInvoiceNotFound) String() string {
	return fmt.Sprintf("[GET /vendor/directFulfillment/shipping/v1/customerInvoices/{purchaseOrderNumber}][%d] getCustomerInvoiceNotFound  %+v", 404, o.Payload)
}

func (o *GetCustomerInvoiceNotFound) GetPayload() *vendor_direct_fulfillment_shipping_v1_models.GetCustomerInvoiceResponse {
	return o.Payload
}

func (o *GetCustomerInvoiceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(vendor_direct_fulfillment_shipping_v1_models.GetCustomerInvoiceResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCustomerInvoiceUnsupportedMediaType creates a GetCustomerInvoiceUnsupportedMediaType with default headers values
func NewGetCustomerInvoiceUnsupportedMediaType() *GetCustomerInvoiceUnsupportedMediaType {
	return &GetCustomerInvoiceUnsupportedMediaType{}
}

/*
GetCustomerInvoiceUnsupportedMediaType describes a response with status code 415, with default header values.

The request payload is in an unsupported format.
*/
type GetCustomerInvoiceUnsupportedMediaType struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_shipping_v1_models.GetCustomerInvoiceResponse
}

// IsSuccess returns true when this get customer invoice unsupported media type response has a 2xx status code
func (o *GetCustomerInvoiceUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get customer invoice unsupported media type response has a 3xx status code
func (o *GetCustomerInvoiceUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get customer invoice unsupported media type response has a 4xx status code
func (o *GetCustomerInvoiceUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get customer invoice unsupported media type response has a 5xx status code
func (o *GetCustomerInvoiceUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get customer invoice unsupported media type response a status code equal to that given
func (o *GetCustomerInvoiceUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetCustomerInvoiceUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /vendor/directFulfillment/shipping/v1/customerInvoices/{purchaseOrderNumber}][%d] getCustomerInvoiceUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetCustomerInvoiceUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /vendor/directFulfillment/shipping/v1/customerInvoices/{purchaseOrderNumber}][%d] getCustomerInvoiceUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetCustomerInvoiceUnsupportedMediaType) GetPayload() *vendor_direct_fulfillment_shipping_v1_models.GetCustomerInvoiceResponse {
	return o.Payload
}

func (o *GetCustomerInvoiceUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(vendor_direct_fulfillment_shipping_v1_models.GetCustomerInvoiceResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCustomerInvoiceTooManyRequests creates a GetCustomerInvoiceTooManyRequests with default headers values
func NewGetCustomerInvoiceTooManyRequests() *GetCustomerInvoiceTooManyRequests {
	return &GetCustomerInvoiceTooManyRequests{}
}

/*
GetCustomerInvoiceTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type GetCustomerInvoiceTooManyRequests struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_shipping_v1_models.GetCustomerInvoiceResponse
}

// IsSuccess returns true when this get customer invoice too many requests response has a 2xx status code
func (o *GetCustomerInvoiceTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get customer invoice too many requests response has a 3xx status code
func (o *GetCustomerInvoiceTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get customer invoice too many requests response has a 4xx status code
func (o *GetCustomerInvoiceTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get customer invoice too many requests response has a 5xx status code
func (o *GetCustomerInvoiceTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get customer invoice too many requests response a status code equal to that given
func (o *GetCustomerInvoiceTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetCustomerInvoiceTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /vendor/directFulfillment/shipping/v1/customerInvoices/{purchaseOrderNumber}][%d] getCustomerInvoiceTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetCustomerInvoiceTooManyRequests) String() string {
	return fmt.Sprintf("[GET /vendor/directFulfillment/shipping/v1/customerInvoices/{purchaseOrderNumber}][%d] getCustomerInvoiceTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetCustomerInvoiceTooManyRequests) GetPayload() *vendor_direct_fulfillment_shipping_v1_models.GetCustomerInvoiceResponse {
	return o.Payload
}

func (o *GetCustomerInvoiceTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(vendor_direct_fulfillment_shipping_v1_models.GetCustomerInvoiceResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCustomerInvoiceInternalServerError creates a GetCustomerInvoiceInternalServerError with default headers values
func NewGetCustomerInvoiceInternalServerError() *GetCustomerInvoiceInternalServerError {
	return &GetCustomerInvoiceInternalServerError{}
}

/*
GetCustomerInvoiceInternalServerError describes a response with status code 500, with default header values.

Encountered an unexpected condition which prevented the server from fulfilling the request.
*/
type GetCustomerInvoiceInternalServerError struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_shipping_v1_models.GetCustomerInvoiceResponse
}

// IsSuccess returns true when this get customer invoice internal server error response has a 2xx status code
func (o *GetCustomerInvoiceInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get customer invoice internal server error response has a 3xx status code
func (o *GetCustomerInvoiceInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get customer invoice internal server error response has a 4xx status code
func (o *GetCustomerInvoiceInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get customer invoice internal server error response has a 5xx status code
func (o *GetCustomerInvoiceInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get customer invoice internal server error response a status code equal to that given
func (o *GetCustomerInvoiceInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetCustomerInvoiceInternalServerError) Error() string {
	return fmt.Sprintf("[GET /vendor/directFulfillment/shipping/v1/customerInvoices/{purchaseOrderNumber}][%d] getCustomerInvoiceInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCustomerInvoiceInternalServerError) String() string {
	return fmt.Sprintf("[GET /vendor/directFulfillment/shipping/v1/customerInvoices/{purchaseOrderNumber}][%d] getCustomerInvoiceInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCustomerInvoiceInternalServerError) GetPayload() *vendor_direct_fulfillment_shipping_v1_models.GetCustomerInvoiceResponse {
	return o.Payload
}

func (o *GetCustomerInvoiceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(vendor_direct_fulfillment_shipping_v1_models.GetCustomerInvoiceResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCustomerInvoiceServiceUnavailable creates a GetCustomerInvoiceServiceUnavailable with default headers values
func NewGetCustomerInvoiceServiceUnavailable() *GetCustomerInvoiceServiceUnavailable {
	return &GetCustomerInvoiceServiceUnavailable{}
}

/*
GetCustomerInvoiceServiceUnavailable describes a response with status code 503, with default header values.

Temporary overloading or maintenance of the server.
*/
type GetCustomerInvoiceServiceUnavailable struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_shipping_v1_models.GetCustomerInvoiceResponse
}

// IsSuccess returns true when this get customer invoice service unavailable response has a 2xx status code
func (o *GetCustomerInvoiceServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get customer invoice service unavailable response has a 3xx status code
func (o *GetCustomerInvoiceServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get customer invoice service unavailable response has a 4xx status code
func (o *GetCustomerInvoiceServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get customer invoice service unavailable response has a 5xx status code
func (o *GetCustomerInvoiceServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get customer invoice service unavailable response a status code equal to that given
func (o *GetCustomerInvoiceServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetCustomerInvoiceServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /vendor/directFulfillment/shipping/v1/customerInvoices/{purchaseOrderNumber}][%d] getCustomerInvoiceServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetCustomerInvoiceServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /vendor/directFulfillment/shipping/v1/customerInvoices/{purchaseOrderNumber}][%d] getCustomerInvoiceServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetCustomerInvoiceServiceUnavailable) GetPayload() *vendor_direct_fulfillment_shipping_v1_models.GetCustomerInvoiceResponse {
	return o.Payload
}

func (o *GetCustomerInvoiceServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(vendor_direct_fulfillment_shipping_v1_models.GetCustomerInvoiceResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
