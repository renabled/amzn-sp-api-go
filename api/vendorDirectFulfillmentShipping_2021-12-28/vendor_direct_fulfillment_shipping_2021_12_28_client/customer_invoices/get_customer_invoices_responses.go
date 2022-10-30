// Code generated by go-swagger; DO NOT EDIT.

package customer_invoices

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/xamandar/amzn-sp-api-go/api/vendorDirectFulfillmentShipping_2021-12-28/vendor_direct_fulfillment_shipping_2021_12_28_models"
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

/*
GetCustomerInvoicesOK describes a response with status code 200, with default header values.

Success.
*/
type GetCustomerInvoicesOK struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_shipping_2021_12_28_models.CustomerInvoiceList
}

// IsSuccess returns true when this get customer invoices o k response has a 2xx status code
func (o *GetCustomerInvoicesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get customer invoices o k response has a 3xx status code
func (o *GetCustomerInvoicesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get customer invoices o k response has a 4xx status code
func (o *GetCustomerInvoicesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get customer invoices o k response has a 5xx status code
func (o *GetCustomerInvoicesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get customer invoices o k response a status code equal to that given
func (o *GetCustomerInvoicesOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetCustomerInvoicesOK) Error() string {
	return fmt.Sprintf("[GET /vendor/directFulfillment/shipping/2021-12-28/customerInvoices][%d] getCustomerInvoicesOK  %+v", 200, o.Payload)
}

func (o *GetCustomerInvoicesOK) String() string {
	return fmt.Sprintf("[GET /vendor/directFulfillment/shipping/2021-12-28/customerInvoices][%d] getCustomerInvoicesOK  %+v", 200, o.Payload)
}

func (o *GetCustomerInvoicesOK) GetPayload() *vendor_direct_fulfillment_shipping_2021_12_28_models.CustomerInvoiceList {
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

	o.Payload = new(vendor_direct_fulfillment_shipping_2021_12_28_models.CustomerInvoiceList)

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

/*
GetCustomerInvoicesBadRequest describes a response with status code 400, with default header values.

Request has missing or invalid parameters and cannot be parsed.
*/
type GetCustomerInvoicesBadRequest struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_shipping_2021_12_28_models.ErrorList
}

// IsSuccess returns true when this get customer invoices bad request response has a 2xx status code
func (o *GetCustomerInvoicesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get customer invoices bad request response has a 3xx status code
func (o *GetCustomerInvoicesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get customer invoices bad request response has a 4xx status code
func (o *GetCustomerInvoicesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get customer invoices bad request response has a 5xx status code
func (o *GetCustomerInvoicesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get customer invoices bad request response a status code equal to that given
func (o *GetCustomerInvoicesBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetCustomerInvoicesBadRequest) Error() string {
	return fmt.Sprintf("[GET /vendor/directFulfillment/shipping/2021-12-28/customerInvoices][%d] getCustomerInvoicesBadRequest  %+v", 400, o.Payload)
}

func (o *GetCustomerInvoicesBadRequest) String() string {
	return fmt.Sprintf("[GET /vendor/directFulfillment/shipping/2021-12-28/customerInvoices][%d] getCustomerInvoicesBadRequest  %+v", 400, o.Payload)
}

func (o *GetCustomerInvoicesBadRequest) GetPayload() *vendor_direct_fulfillment_shipping_2021_12_28_models.ErrorList {
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

	o.Payload = new(vendor_direct_fulfillment_shipping_2021_12_28_models.ErrorList)

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

/*
GetCustomerInvoicesForbidden describes a response with status code 403, with default header values.

Indicates that access to the resource is forbidden. Possible reasons include Access Denied, Unauthorized, Expired Token, or Invalid Signature.
*/
type GetCustomerInvoicesForbidden struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_shipping_2021_12_28_models.ErrorList
}

// IsSuccess returns true when this get customer invoices forbidden response has a 2xx status code
func (o *GetCustomerInvoicesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get customer invoices forbidden response has a 3xx status code
func (o *GetCustomerInvoicesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get customer invoices forbidden response has a 4xx status code
func (o *GetCustomerInvoicesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get customer invoices forbidden response has a 5xx status code
func (o *GetCustomerInvoicesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get customer invoices forbidden response a status code equal to that given
func (o *GetCustomerInvoicesForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetCustomerInvoicesForbidden) Error() string {
	return fmt.Sprintf("[GET /vendor/directFulfillment/shipping/2021-12-28/customerInvoices][%d] getCustomerInvoicesForbidden  %+v", 403, o.Payload)
}

func (o *GetCustomerInvoicesForbidden) String() string {
	return fmt.Sprintf("[GET /vendor/directFulfillment/shipping/2021-12-28/customerInvoices][%d] getCustomerInvoicesForbidden  %+v", 403, o.Payload)
}

func (o *GetCustomerInvoicesForbidden) GetPayload() *vendor_direct_fulfillment_shipping_2021_12_28_models.ErrorList {
	return o.Payload
}

func (o *GetCustomerInvoicesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCustomerInvoicesNotFound creates a GetCustomerInvoicesNotFound with default headers values
func NewGetCustomerInvoicesNotFound() *GetCustomerInvoicesNotFound {
	return &GetCustomerInvoicesNotFound{}
}

/*
GetCustomerInvoicesNotFound describes a response with status code 404, with default header values.

The resource specified does not exist.
*/
type GetCustomerInvoicesNotFound struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_shipping_2021_12_28_models.ErrorList
}

// IsSuccess returns true when this get customer invoices not found response has a 2xx status code
func (o *GetCustomerInvoicesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get customer invoices not found response has a 3xx status code
func (o *GetCustomerInvoicesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get customer invoices not found response has a 4xx status code
func (o *GetCustomerInvoicesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get customer invoices not found response has a 5xx status code
func (o *GetCustomerInvoicesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get customer invoices not found response a status code equal to that given
func (o *GetCustomerInvoicesNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetCustomerInvoicesNotFound) Error() string {
	return fmt.Sprintf("[GET /vendor/directFulfillment/shipping/2021-12-28/customerInvoices][%d] getCustomerInvoicesNotFound  %+v", 404, o.Payload)
}

func (o *GetCustomerInvoicesNotFound) String() string {
	return fmt.Sprintf("[GET /vendor/directFulfillment/shipping/2021-12-28/customerInvoices][%d] getCustomerInvoicesNotFound  %+v", 404, o.Payload)
}

func (o *GetCustomerInvoicesNotFound) GetPayload() *vendor_direct_fulfillment_shipping_2021_12_28_models.ErrorList {
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

	o.Payload = new(vendor_direct_fulfillment_shipping_2021_12_28_models.ErrorList)

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

/*
GetCustomerInvoicesUnsupportedMediaType describes a response with status code 415, with default header values.

The request payload is in an unsupported format.
*/
type GetCustomerInvoicesUnsupportedMediaType struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_shipping_2021_12_28_models.ErrorList
}

// IsSuccess returns true when this get customer invoices unsupported media type response has a 2xx status code
func (o *GetCustomerInvoicesUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get customer invoices unsupported media type response has a 3xx status code
func (o *GetCustomerInvoicesUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get customer invoices unsupported media type response has a 4xx status code
func (o *GetCustomerInvoicesUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get customer invoices unsupported media type response has a 5xx status code
func (o *GetCustomerInvoicesUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get customer invoices unsupported media type response a status code equal to that given
func (o *GetCustomerInvoicesUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetCustomerInvoicesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /vendor/directFulfillment/shipping/2021-12-28/customerInvoices][%d] getCustomerInvoicesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetCustomerInvoicesUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /vendor/directFulfillment/shipping/2021-12-28/customerInvoices][%d] getCustomerInvoicesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetCustomerInvoicesUnsupportedMediaType) GetPayload() *vendor_direct_fulfillment_shipping_2021_12_28_models.ErrorList {
	return o.Payload
}

func (o *GetCustomerInvoicesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCustomerInvoicesTooManyRequests creates a GetCustomerInvoicesTooManyRequests with default headers values
func NewGetCustomerInvoicesTooManyRequests() *GetCustomerInvoicesTooManyRequests {
	return &GetCustomerInvoicesTooManyRequests{}
}

/*
GetCustomerInvoicesTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type GetCustomerInvoicesTooManyRequests struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_shipping_2021_12_28_models.ErrorList
}

// IsSuccess returns true when this get customer invoices too many requests response has a 2xx status code
func (o *GetCustomerInvoicesTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get customer invoices too many requests response has a 3xx status code
func (o *GetCustomerInvoicesTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get customer invoices too many requests response has a 4xx status code
func (o *GetCustomerInvoicesTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get customer invoices too many requests response has a 5xx status code
func (o *GetCustomerInvoicesTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get customer invoices too many requests response a status code equal to that given
func (o *GetCustomerInvoicesTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetCustomerInvoicesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /vendor/directFulfillment/shipping/2021-12-28/customerInvoices][%d] getCustomerInvoicesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetCustomerInvoicesTooManyRequests) String() string {
	return fmt.Sprintf("[GET /vendor/directFulfillment/shipping/2021-12-28/customerInvoices][%d] getCustomerInvoicesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetCustomerInvoicesTooManyRequests) GetPayload() *vendor_direct_fulfillment_shipping_2021_12_28_models.ErrorList {
	return o.Payload
}

func (o *GetCustomerInvoicesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCustomerInvoicesInternalServerError creates a GetCustomerInvoicesInternalServerError with default headers values
func NewGetCustomerInvoicesInternalServerError() *GetCustomerInvoicesInternalServerError {
	return &GetCustomerInvoicesInternalServerError{}
}

/*
GetCustomerInvoicesInternalServerError describes a response with status code 500, with default header values.

An unexpected condition occurred that prevented the server from fulfilling the request.
*/
type GetCustomerInvoicesInternalServerError struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_shipping_2021_12_28_models.ErrorList
}

// IsSuccess returns true when this get customer invoices internal server error response has a 2xx status code
func (o *GetCustomerInvoicesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get customer invoices internal server error response has a 3xx status code
func (o *GetCustomerInvoicesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get customer invoices internal server error response has a 4xx status code
func (o *GetCustomerInvoicesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get customer invoices internal server error response has a 5xx status code
func (o *GetCustomerInvoicesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get customer invoices internal server error response a status code equal to that given
func (o *GetCustomerInvoicesInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetCustomerInvoicesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /vendor/directFulfillment/shipping/2021-12-28/customerInvoices][%d] getCustomerInvoicesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCustomerInvoicesInternalServerError) String() string {
	return fmt.Sprintf("[GET /vendor/directFulfillment/shipping/2021-12-28/customerInvoices][%d] getCustomerInvoicesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCustomerInvoicesInternalServerError) GetPayload() *vendor_direct_fulfillment_shipping_2021_12_28_models.ErrorList {
	return o.Payload
}

func (o *GetCustomerInvoicesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCustomerInvoicesServiceUnavailable creates a GetCustomerInvoicesServiceUnavailable with default headers values
func NewGetCustomerInvoicesServiceUnavailable() *GetCustomerInvoicesServiceUnavailable {
	return &GetCustomerInvoicesServiceUnavailable{}
}

/*
GetCustomerInvoicesServiceUnavailable describes a response with status code 503, with default header values.

Temporary overloading or maintenance of the server.
*/
type GetCustomerInvoicesServiceUnavailable struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_shipping_2021_12_28_models.ErrorList
}

// IsSuccess returns true when this get customer invoices service unavailable response has a 2xx status code
func (o *GetCustomerInvoicesServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get customer invoices service unavailable response has a 3xx status code
func (o *GetCustomerInvoicesServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get customer invoices service unavailable response has a 4xx status code
func (o *GetCustomerInvoicesServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get customer invoices service unavailable response has a 5xx status code
func (o *GetCustomerInvoicesServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get customer invoices service unavailable response a status code equal to that given
func (o *GetCustomerInvoicesServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetCustomerInvoicesServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /vendor/directFulfillment/shipping/2021-12-28/customerInvoices][%d] getCustomerInvoicesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetCustomerInvoicesServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /vendor/directFulfillment/shipping/2021-12-28/customerInvoices][%d] getCustomerInvoicesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetCustomerInvoicesServiceUnavailable) GetPayload() *vendor_direct_fulfillment_shipping_2021_12_28_models.ErrorList {
	return o.Payload
}

func (o *GetCustomerInvoicesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
