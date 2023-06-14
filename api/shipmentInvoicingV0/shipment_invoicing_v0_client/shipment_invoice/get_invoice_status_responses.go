// Code generated by go-swagger; DO NOT EDIT.

package shipment_invoice

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/renabled/amzn-sp-api-go/api/shipmentInvoicingV0/shipment_invoicing_v0_models"
)

// GetInvoiceStatusReader is a Reader for the GetInvoiceStatus structure.
type GetInvoiceStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInvoiceStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetInvoiceStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetInvoiceStatusBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetInvoiceStatusUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetInvoiceStatusForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetInvoiceStatusNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetInvoiceStatusUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetInvoiceStatusTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetInvoiceStatusInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetInvoiceStatusServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetInvoiceStatusOK creates a GetInvoiceStatusOK with default headers values
func NewGetInvoiceStatusOK() *GetInvoiceStatusOK {
	return &GetInvoiceStatusOK{}
}

/*
GetInvoiceStatusOK describes a response with status code 200, with default header values.

Success.
*/
type GetInvoiceStatusOK struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *shipment_invoicing_v0_models.GetInvoiceStatusResponse
}

// IsSuccess returns true when this get invoice status o k response has a 2xx status code
func (o *GetInvoiceStatusOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get invoice status o k response has a 3xx status code
func (o *GetInvoiceStatusOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get invoice status o k response has a 4xx status code
func (o *GetInvoiceStatusOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get invoice status o k response has a 5xx status code
func (o *GetInvoiceStatusOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get invoice status o k response a status code equal to that given
func (o *GetInvoiceStatusOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetInvoiceStatusOK) Error() string {
	return fmt.Sprintf("[GET /fba/outbound/brazil/v0/shipments/{shipmentId}/invoice/status][%d] getInvoiceStatusOK  %+v", 200, o.Payload)
}

func (o *GetInvoiceStatusOK) String() string {
	return fmt.Sprintf("[GET /fba/outbound/brazil/v0/shipments/{shipmentId}/invoice/status][%d] getInvoiceStatusOK  %+v", 200, o.Payload)
}

func (o *GetInvoiceStatusOK) GetPayload() *shipment_invoicing_v0_models.GetInvoiceStatusResponse {
	return o.Payload
}

func (o *GetInvoiceStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(shipment_invoicing_v0_models.GetInvoiceStatusResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInvoiceStatusBadRequest creates a GetInvoiceStatusBadRequest with default headers values
func NewGetInvoiceStatusBadRequest() *GetInvoiceStatusBadRequest {
	return &GetInvoiceStatusBadRequest{}
}

/*
GetInvoiceStatusBadRequest describes a response with status code 400, with default header values.

Request has missing or invalid parameters and cannot be parsed.
*/
type GetInvoiceStatusBadRequest struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *shipment_invoicing_v0_models.GetInvoiceStatusResponse
}

// IsSuccess returns true when this get invoice status bad request response has a 2xx status code
func (o *GetInvoiceStatusBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get invoice status bad request response has a 3xx status code
func (o *GetInvoiceStatusBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get invoice status bad request response has a 4xx status code
func (o *GetInvoiceStatusBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get invoice status bad request response has a 5xx status code
func (o *GetInvoiceStatusBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get invoice status bad request response a status code equal to that given
func (o *GetInvoiceStatusBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetInvoiceStatusBadRequest) Error() string {
	return fmt.Sprintf("[GET /fba/outbound/brazil/v0/shipments/{shipmentId}/invoice/status][%d] getInvoiceStatusBadRequest  %+v", 400, o.Payload)
}

func (o *GetInvoiceStatusBadRequest) String() string {
	return fmt.Sprintf("[GET /fba/outbound/brazil/v0/shipments/{shipmentId}/invoice/status][%d] getInvoiceStatusBadRequest  %+v", 400, o.Payload)
}

func (o *GetInvoiceStatusBadRequest) GetPayload() *shipment_invoicing_v0_models.GetInvoiceStatusResponse {
	return o.Payload
}

func (o *GetInvoiceStatusBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(shipment_invoicing_v0_models.GetInvoiceStatusResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInvoiceStatusUnauthorized creates a GetInvoiceStatusUnauthorized with default headers values
func NewGetInvoiceStatusUnauthorized() *GetInvoiceStatusUnauthorized {
	return &GetInvoiceStatusUnauthorized{}
}

/*
GetInvoiceStatusUnauthorized describes a response with status code 401, with default header values.

The request's Authorization header is not formatted correctly or does not contain a valid token.
*/
type GetInvoiceStatusUnauthorized struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *shipment_invoicing_v0_models.GetInvoiceStatusResponse
}

// IsSuccess returns true when this get invoice status unauthorized response has a 2xx status code
func (o *GetInvoiceStatusUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get invoice status unauthorized response has a 3xx status code
func (o *GetInvoiceStatusUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get invoice status unauthorized response has a 4xx status code
func (o *GetInvoiceStatusUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get invoice status unauthorized response has a 5xx status code
func (o *GetInvoiceStatusUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get invoice status unauthorized response a status code equal to that given
func (o *GetInvoiceStatusUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetInvoiceStatusUnauthorized) Error() string {
	return fmt.Sprintf("[GET /fba/outbound/brazil/v0/shipments/{shipmentId}/invoice/status][%d] getInvoiceStatusUnauthorized  %+v", 401, o.Payload)
}

func (o *GetInvoiceStatusUnauthorized) String() string {
	return fmt.Sprintf("[GET /fba/outbound/brazil/v0/shipments/{shipmentId}/invoice/status][%d] getInvoiceStatusUnauthorized  %+v", 401, o.Payload)
}

func (o *GetInvoiceStatusUnauthorized) GetPayload() *shipment_invoicing_v0_models.GetInvoiceStatusResponse {
	return o.Payload
}

func (o *GetInvoiceStatusUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(shipment_invoicing_v0_models.GetInvoiceStatusResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInvoiceStatusForbidden creates a GetInvoiceStatusForbidden with default headers values
func NewGetInvoiceStatusForbidden() *GetInvoiceStatusForbidden {
	return &GetInvoiceStatusForbidden{}
}

/*
GetInvoiceStatusForbidden describes a response with status code 403, with default header values.

Indicates that access to the resource is forbidden. Possible reasons include Access Denied, Unauthorized, Expired Token, or Invalid Signature.
*/
type GetInvoiceStatusForbidden struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *shipment_invoicing_v0_models.GetInvoiceStatusResponse
}

// IsSuccess returns true when this get invoice status forbidden response has a 2xx status code
func (o *GetInvoiceStatusForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get invoice status forbidden response has a 3xx status code
func (o *GetInvoiceStatusForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get invoice status forbidden response has a 4xx status code
func (o *GetInvoiceStatusForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get invoice status forbidden response has a 5xx status code
func (o *GetInvoiceStatusForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get invoice status forbidden response a status code equal to that given
func (o *GetInvoiceStatusForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetInvoiceStatusForbidden) Error() string {
	return fmt.Sprintf("[GET /fba/outbound/brazil/v0/shipments/{shipmentId}/invoice/status][%d] getInvoiceStatusForbidden  %+v", 403, o.Payload)
}

func (o *GetInvoiceStatusForbidden) String() string {
	return fmt.Sprintf("[GET /fba/outbound/brazil/v0/shipments/{shipmentId}/invoice/status][%d] getInvoiceStatusForbidden  %+v", 403, o.Payload)
}

func (o *GetInvoiceStatusForbidden) GetPayload() *shipment_invoicing_v0_models.GetInvoiceStatusResponse {
	return o.Payload
}

func (o *GetInvoiceStatusForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(shipment_invoicing_v0_models.GetInvoiceStatusResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInvoiceStatusNotFound creates a GetInvoiceStatusNotFound with default headers values
func NewGetInvoiceStatusNotFound() *GetInvoiceStatusNotFound {
	return &GetInvoiceStatusNotFound{}
}

/*
GetInvoiceStatusNotFound describes a response with status code 404, with default header values.

The resource specified does not exist.
*/
type GetInvoiceStatusNotFound struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *shipment_invoicing_v0_models.GetInvoiceStatusResponse
}

// IsSuccess returns true when this get invoice status not found response has a 2xx status code
func (o *GetInvoiceStatusNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get invoice status not found response has a 3xx status code
func (o *GetInvoiceStatusNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get invoice status not found response has a 4xx status code
func (o *GetInvoiceStatusNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get invoice status not found response has a 5xx status code
func (o *GetInvoiceStatusNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get invoice status not found response a status code equal to that given
func (o *GetInvoiceStatusNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetInvoiceStatusNotFound) Error() string {
	return fmt.Sprintf("[GET /fba/outbound/brazil/v0/shipments/{shipmentId}/invoice/status][%d] getInvoiceStatusNotFound  %+v", 404, o.Payload)
}

func (o *GetInvoiceStatusNotFound) String() string {
	return fmt.Sprintf("[GET /fba/outbound/brazil/v0/shipments/{shipmentId}/invoice/status][%d] getInvoiceStatusNotFound  %+v", 404, o.Payload)
}

func (o *GetInvoiceStatusNotFound) GetPayload() *shipment_invoicing_v0_models.GetInvoiceStatusResponse {
	return o.Payload
}

func (o *GetInvoiceStatusNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(shipment_invoicing_v0_models.GetInvoiceStatusResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInvoiceStatusUnsupportedMediaType creates a GetInvoiceStatusUnsupportedMediaType with default headers values
func NewGetInvoiceStatusUnsupportedMediaType() *GetInvoiceStatusUnsupportedMediaType {
	return &GetInvoiceStatusUnsupportedMediaType{}
}

/*
GetInvoiceStatusUnsupportedMediaType describes a response with status code 415, with default header values.

The request payload is in an unsupported format.
*/
type GetInvoiceStatusUnsupportedMediaType struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *shipment_invoicing_v0_models.GetInvoiceStatusResponse
}

// IsSuccess returns true when this get invoice status unsupported media type response has a 2xx status code
func (o *GetInvoiceStatusUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get invoice status unsupported media type response has a 3xx status code
func (o *GetInvoiceStatusUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get invoice status unsupported media type response has a 4xx status code
func (o *GetInvoiceStatusUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get invoice status unsupported media type response has a 5xx status code
func (o *GetInvoiceStatusUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get invoice status unsupported media type response a status code equal to that given
func (o *GetInvoiceStatusUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetInvoiceStatusUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /fba/outbound/brazil/v0/shipments/{shipmentId}/invoice/status][%d] getInvoiceStatusUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetInvoiceStatusUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /fba/outbound/brazil/v0/shipments/{shipmentId}/invoice/status][%d] getInvoiceStatusUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetInvoiceStatusUnsupportedMediaType) GetPayload() *shipment_invoicing_v0_models.GetInvoiceStatusResponse {
	return o.Payload
}

func (o *GetInvoiceStatusUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(shipment_invoicing_v0_models.GetInvoiceStatusResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInvoiceStatusTooManyRequests creates a GetInvoiceStatusTooManyRequests with default headers values
func NewGetInvoiceStatusTooManyRequests() *GetInvoiceStatusTooManyRequests {
	return &GetInvoiceStatusTooManyRequests{}
}

/*
GetInvoiceStatusTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type GetInvoiceStatusTooManyRequests struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *shipment_invoicing_v0_models.GetInvoiceStatusResponse
}

// IsSuccess returns true when this get invoice status too many requests response has a 2xx status code
func (o *GetInvoiceStatusTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get invoice status too many requests response has a 3xx status code
func (o *GetInvoiceStatusTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get invoice status too many requests response has a 4xx status code
func (o *GetInvoiceStatusTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get invoice status too many requests response has a 5xx status code
func (o *GetInvoiceStatusTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get invoice status too many requests response a status code equal to that given
func (o *GetInvoiceStatusTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetInvoiceStatusTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /fba/outbound/brazil/v0/shipments/{shipmentId}/invoice/status][%d] getInvoiceStatusTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetInvoiceStatusTooManyRequests) String() string {
	return fmt.Sprintf("[GET /fba/outbound/brazil/v0/shipments/{shipmentId}/invoice/status][%d] getInvoiceStatusTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetInvoiceStatusTooManyRequests) GetPayload() *shipment_invoicing_v0_models.GetInvoiceStatusResponse {
	return o.Payload
}

func (o *GetInvoiceStatusTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(shipment_invoicing_v0_models.GetInvoiceStatusResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInvoiceStatusInternalServerError creates a GetInvoiceStatusInternalServerError with default headers values
func NewGetInvoiceStatusInternalServerError() *GetInvoiceStatusInternalServerError {
	return &GetInvoiceStatusInternalServerError{}
}

/*
GetInvoiceStatusInternalServerError describes a response with status code 500, with default header values.

An unexpected condition occurred that prevented the server from fulfilling the request.
*/
type GetInvoiceStatusInternalServerError struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *shipment_invoicing_v0_models.GetInvoiceStatusResponse
}

// IsSuccess returns true when this get invoice status internal server error response has a 2xx status code
func (o *GetInvoiceStatusInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get invoice status internal server error response has a 3xx status code
func (o *GetInvoiceStatusInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get invoice status internal server error response has a 4xx status code
func (o *GetInvoiceStatusInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get invoice status internal server error response has a 5xx status code
func (o *GetInvoiceStatusInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get invoice status internal server error response a status code equal to that given
func (o *GetInvoiceStatusInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetInvoiceStatusInternalServerError) Error() string {
	return fmt.Sprintf("[GET /fba/outbound/brazil/v0/shipments/{shipmentId}/invoice/status][%d] getInvoiceStatusInternalServerError  %+v", 500, o.Payload)
}

func (o *GetInvoiceStatusInternalServerError) String() string {
	return fmt.Sprintf("[GET /fba/outbound/brazil/v0/shipments/{shipmentId}/invoice/status][%d] getInvoiceStatusInternalServerError  %+v", 500, o.Payload)
}

func (o *GetInvoiceStatusInternalServerError) GetPayload() *shipment_invoicing_v0_models.GetInvoiceStatusResponse {
	return o.Payload
}

func (o *GetInvoiceStatusInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(shipment_invoicing_v0_models.GetInvoiceStatusResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInvoiceStatusServiceUnavailable creates a GetInvoiceStatusServiceUnavailable with default headers values
func NewGetInvoiceStatusServiceUnavailable() *GetInvoiceStatusServiceUnavailable {
	return &GetInvoiceStatusServiceUnavailable{}
}

/*
GetInvoiceStatusServiceUnavailable describes a response with status code 503, with default header values.

Temporary overloading or maintenance of the server.
*/
type GetInvoiceStatusServiceUnavailable struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *shipment_invoicing_v0_models.GetInvoiceStatusResponse
}

// IsSuccess returns true when this get invoice status service unavailable response has a 2xx status code
func (o *GetInvoiceStatusServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get invoice status service unavailable response has a 3xx status code
func (o *GetInvoiceStatusServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get invoice status service unavailable response has a 4xx status code
func (o *GetInvoiceStatusServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get invoice status service unavailable response has a 5xx status code
func (o *GetInvoiceStatusServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get invoice status service unavailable response a status code equal to that given
func (o *GetInvoiceStatusServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetInvoiceStatusServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /fba/outbound/brazil/v0/shipments/{shipmentId}/invoice/status][%d] getInvoiceStatusServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetInvoiceStatusServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /fba/outbound/brazil/v0/shipments/{shipmentId}/invoice/status][%d] getInvoiceStatusServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetInvoiceStatusServiceUnavailable) GetPayload() *shipment_invoicing_v0_models.GetInvoiceStatusResponse {
	return o.Payload
}

func (o *GetInvoiceStatusServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(shipment_invoicing_v0_models.GetInvoiceStatusResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
