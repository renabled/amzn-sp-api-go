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

// GetCustomerInvoicesReader is a Reader for the GetCustomerInvoices structure.
type GetCustomerInvoicesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCustomerInvoicesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCustomerInvoicesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetCustomerInvoicesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetCustomerInvoicesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetCustomerInvoicesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetCustomerInvoicesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetCustomerInvoicesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetCustomerInvoicesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetCustomerInvoicesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetCustomerInvoicesOK creates a GetCustomerInvoicesOK with default headers values
func NewGetCustomerInvoicesOK() *GetCustomerInvoicesOK {
	return &GetCustomerInvoicesOK{}
}

/* GetCustomerInvoicesOK describes a response with status code 200, with default header values.

Success.
*/
type GetCustomerInvoicesOK struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_shipping_v1_models.GetCustomerInvoicesResponse
}

func (o *GetCustomerInvoicesOK) Error() string {
	return fmt.Sprintf("[GET /vendor/directFulfillment/shipping/v1/customerInvoices][%d] getCustomerInvoicesOK  %+v", 200, o.Payload)
}
func (o *GetCustomerInvoicesOK) GetPayload() *vendor_direct_fulfillment_shipping_v1_models.GetCustomerInvoicesResponse {
	return o.Payload
}

func (o *GetCustomerInvoicesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(vendor_direct_fulfillment_shipping_v1_models.GetCustomerInvoicesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCustomerInvoicesBadRequest creates a GetCustomerInvoicesBadRequest with default headers values
func NewGetCustomerInvoicesBadRequest() *GetCustomerInvoicesBadRequest {
	return &GetCustomerInvoicesBadRequest{}
}

/* GetCustomerInvoicesBadRequest describes a response with status code 400, with default header values.

Request has missing or invalid parameters and cannot be parsed.
*/
type GetCustomerInvoicesBadRequest struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_shipping_v1_models.GetCustomerInvoiceResponse
}

func (o *GetCustomerInvoicesBadRequest) Error() string {
	return fmt.Sprintf("[GET /vendor/directFulfillment/shipping/v1/customerInvoices][%d] getCustomerInvoicesBadRequest  %+v", 400, o.Payload)
}
func (o *GetCustomerInvoicesBadRequest) GetPayload() *vendor_direct_fulfillment_shipping_v1_models.GetCustomerInvoiceResponse {
	return o.Payload
}

func (o *GetCustomerInvoicesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCustomerInvoicesForbidden creates a GetCustomerInvoicesForbidden with default headers values
func NewGetCustomerInvoicesForbidden() *GetCustomerInvoicesForbidden {
	return &GetCustomerInvoicesForbidden{}
}

/* GetCustomerInvoicesForbidden describes a response with status code 403, with default header values.

Indicates that access to the resource is forbidden. Possible reasons include Access Denied, Unauthorized, Expired Token, or Invalid Signature.
*/
type GetCustomerInvoicesForbidden struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_shipping_v1_models.GetCustomerInvoiceResponse
}

func (o *GetCustomerInvoicesForbidden) Error() string {
	return fmt.Sprintf("[GET /vendor/directFulfillment/shipping/v1/customerInvoices][%d] getCustomerInvoicesForbidden  %+v", 403, o.Payload)
}
func (o *GetCustomerInvoicesForbidden) GetPayload() *vendor_direct_fulfillment_shipping_v1_models.GetCustomerInvoiceResponse {
	return o.Payload
}

func (o *GetCustomerInvoicesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCustomerInvoicesNotFound creates a GetCustomerInvoicesNotFound with default headers values
func NewGetCustomerInvoicesNotFound() *GetCustomerInvoicesNotFound {
	return &GetCustomerInvoicesNotFound{}
}

/* GetCustomerInvoicesNotFound describes a response with status code 404, with default header values.

The resource specified does not exist.
*/
type GetCustomerInvoicesNotFound struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_shipping_v1_models.GetCustomerInvoiceResponse
}

func (o *GetCustomerInvoicesNotFound) Error() string {
	return fmt.Sprintf("[GET /vendor/directFulfillment/shipping/v1/customerInvoices][%d] getCustomerInvoicesNotFound  %+v", 404, o.Payload)
}
func (o *GetCustomerInvoicesNotFound) GetPayload() *vendor_direct_fulfillment_shipping_v1_models.GetCustomerInvoiceResponse {
	return o.Payload
}

func (o *GetCustomerInvoicesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCustomerInvoicesUnsupportedMediaType creates a GetCustomerInvoicesUnsupportedMediaType with default headers values
func NewGetCustomerInvoicesUnsupportedMediaType() *GetCustomerInvoicesUnsupportedMediaType {
	return &GetCustomerInvoicesUnsupportedMediaType{}
}

/* GetCustomerInvoicesUnsupportedMediaType describes a response with status code 415, with default header values.

The request payload is in an unsupported format.
*/
type GetCustomerInvoicesUnsupportedMediaType struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_shipping_v1_models.GetCustomerInvoiceResponse
}

func (o *GetCustomerInvoicesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /vendor/directFulfillment/shipping/v1/customerInvoices][%d] getCustomerInvoicesUnsupportedMediaType  %+v", 415, o.Payload)
}
func (o *GetCustomerInvoicesUnsupportedMediaType) GetPayload() *vendor_direct_fulfillment_shipping_v1_models.GetCustomerInvoiceResponse {
	return o.Payload
}

func (o *GetCustomerInvoicesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCustomerInvoicesTooManyRequests creates a GetCustomerInvoicesTooManyRequests with default headers values
func NewGetCustomerInvoicesTooManyRequests() *GetCustomerInvoicesTooManyRequests {
	return &GetCustomerInvoicesTooManyRequests{}
}

/* GetCustomerInvoicesTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type GetCustomerInvoicesTooManyRequests struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_shipping_v1_models.GetCustomerInvoiceResponse
}

func (o *GetCustomerInvoicesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /vendor/directFulfillment/shipping/v1/customerInvoices][%d] getCustomerInvoicesTooManyRequests  %+v", 429, o.Payload)
}
func (o *GetCustomerInvoicesTooManyRequests) GetPayload() *vendor_direct_fulfillment_shipping_v1_models.GetCustomerInvoiceResponse {
	return o.Payload
}

func (o *GetCustomerInvoicesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCustomerInvoicesInternalServerError creates a GetCustomerInvoicesInternalServerError with default headers values
func NewGetCustomerInvoicesInternalServerError() *GetCustomerInvoicesInternalServerError {
	return &GetCustomerInvoicesInternalServerError{}
}

/* GetCustomerInvoicesInternalServerError describes a response with status code 500, with default header values.

Encountered an unexpected condition which prevented the server from fulfilling the request.
*/
type GetCustomerInvoicesInternalServerError struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_shipping_v1_models.GetCustomerInvoiceResponse
}

func (o *GetCustomerInvoicesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /vendor/directFulfillment/shipping/v1/customerInvoices][%d] getCustomerInvoicesInternalServerError  %+v", 500, o.Payload)
}
func (o *GetCustomerInvoicesInternalServerError) GetPayload() *vendor_direct_fulfillment_shipping_v1_models.GetCustomerInvoiceResponse {
	return o.Payload
}

func (o *GetCustomerInvoicesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCustomerInvoicesServiceUnavailable creates a GetCustomerInvoicesServiceUnavailable with default headers values
func NewGetCustomerInvoicesServiceUnavailable() *GetCustomerInvoicesServiceUnavailable {
	return &GetCustomerInvoicesServiceUnavailable{}
}

/* GetCustomerInvoicesServiceUnavailable describes a response with status code 503, with default header values.

Temporary overloading or maintenance of the server.
*/
type GetCustomerInvoicesServiceUnavailable struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_shipping_v1_models.GetCustomerInvoiceResponse
}

func (o *GetCustomerInvoicesServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /vendor/directFulfillment/shipping/v1/customerInvoices][%d] getCustomerInvoicesServiceUnavailable  %+v", 503, o.Payload)
}
func (o *GetCustomerInvoicesServiceUnavailable) GetPayload() *vendor_direct_fulfillment_shipping_v1_models.GetCustomerInvoiceResponse {
	return o.Payload
}

func (o *GetCustomerInvoicesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
