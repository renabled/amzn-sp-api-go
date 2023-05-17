// Code generated by go-swagger; DO NOT EDIT.

package shipment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/xamandar/amzn-sp-api-go/api/ordersV0/orders_v0_models"
)

// UpdateShipmentStatusReader is a Reader for the UpdateShipmentStatus structure.
type UpdateShipmentStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateShipmentStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewUpdateShipmentStatusNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateShipmentStatusBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateShipmentStatusForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateShipmentStatusNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewUpdateShipmentStatusRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewUpdateShipmentStatusUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewUpdateShipmentStatusTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateShipmentStatusInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewUpdateShipmentStatusServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateShipmentStatusNoContent creates a UpdateShipmentStatusNoContent with default headers values
func NewUpdateShipmentStatusNoContent() *UpdateShipmentStatusNoContent {
	return &UpdateShipmentStatusNoContent{}
}

/*
UpdateShipmentStatusNoContent describes a response with status code 204, with default header values.

Success.
*/
type UpdateShipmentStatusNoContent struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string
}

// IsSuccess returns true when this update shipment status no content response has a 2xx status code
func (o *UpdateShipmentStatusNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update shipment status no content response has a 3xx status code
func (o *UpdateShipmentStatusNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update shipment status no content response has a 4xx status code
func (o *UpdateShipmentStatusNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this update shipment status no content response has a 5xx status code
func (o *UpdateShipmentStatusNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this update shipment status no content response a status code equal to that given
func (o *UpdateShipmentStatusNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *UpdateShipmentStatusNoContent) Error() string {
	return fmt.Sprintf("[POST /orders/v0/orders/{orderId}/shipment][%d] updateShipmentStatusNoContent ", 204)
}

func (o *UpdateShipmentStatusNoContent) String() string {
	return fmt.Sprintf("[POST /orders/v0/orders/{orderId}/shipment][%d] updateShipmentStatusNoContent ", 204)
}

func (o *UpdateShipmentStatusNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	return nil
}

// NewUpdateShipmentStatusBadRequest creates a UpdateShipmentStatusBadRequest with default headers values
func NewUpdateShipmentStatusBadRequest() *UpdateShipmentStatusBadRequest {
	return &UpdateShipmentStatusBadRequest{}
}

/*
UpdateShipmentStatusBadRequest describes a response with status code 400, with default header values.

Request has missing or invalid parameters and cannot be parsed.
*/
type UpdateShipmentStatusBadRequest struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *orders_v0_models.UpdateShipmentStatusErrorResponse
}

// IsSuccess returns true when this update shipment status bad request response has a 2xx status code
func (o *UpdateShipmentStatusBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update shipment status bad request response has a 3xx status code
func (o *UpdateShipmentStatusBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update shipment status bad request response has a 4xx status code
func (o *UpdateShipmentStatusBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update shipment status bad request response has a 5xx status code
func (o *UpdateShipmentStatusBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update shipment status bad request response a status code equal to that given
func (o *UpdateShipmentStatusBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *UpdateShipmentStatusBadRequest) Error() string {
	return fmt.Sprintf("[POST /orders/v0/orders/{orderId}/shipment][%d] updateShipmentStatusBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateShipmentStatusBadRequest) String() string {
	return fmt.Sprintf("[POST /orders/v0/orders/{orderId}/shipment][%d] updateShipmentStatusBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateShipmentStatusBadRequest) GetPayload() *orders_v0_models.UpdateShipmentStatusErrorResponse {
	return o.Payload
}

func (o *UpdateShipmentStatusBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(orders_v0_models.UpdateShipmentStatusErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateShipmentStatusForbidden creates a UpdateShipmentStatusForbidden with default headers values
func NewUpdateShipmentStatusForbidden() *UpdateShipmentStatusForbidden {
	return &UpdateShipmentStatusForbidden{}
}

/*
UpdateShipmentStatusForbidden describes a response with status code 403, with default header values.

Indicates that access to the resource is forbidden. Possible reasons include Access Denied, Unauthorized, Expired Token, or Invalid Signature.
*/
type UpdateShipmentStatusForbidden struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *orders_v0_models.UpdateShipmentStatusErrorResponse
}

// IsSuccess returns true when this update shipment status forbidden response has a 2xx status code
func (o *UpdateShipmentStatusForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update shipment status forbidden response has a 3xx status code
func (o *UpdateShipmentStatusForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update shipment status forbidden response has a 4xx status code
func (o *UpdateShipmentStatusForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update shipment status forbidden response has a 5xx status code
func (o *UpdateShipmentStatusForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update shipment status forbidden response a status code equal to that given
func (o *UpdateShipmentStatusForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *UpdateShipmentStatusForbidden) Error() string {
	return fmt.Sprintf("[POST /orders/v0/orders/{orderId}/shipment][%d] updateShipmentStatusForbidden  %+v", 403, o.Payload)
}

func (o *UpdateShipmentStatusForbidden) String() string {
	return fmt.Sprintf("[POST /orders/v0/orders/{orderId}/shipment][%d] updateShipmentStatusForbidden  %+v", 403, o.Payload)
}

func (o *UpdateShipmentStatusForbidden) GetPayload() *orders_v0_models.UpdateShipmentStatusErrorResponse {
	return o.Payload
}

func (o *UpdateShipmentStatusForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(orders_v0_models.UpdateShipmentStatusErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateShipmentStatusNotFound creates a UpdateShipmentStatusNotFound with default headers values
func NewUpdateShipmentStatusNotFound() *UpdateShipmentStatusNotFound {
	return &UpdateShipmentStatusNotFound{}
}

/*
UpdateShipmentStatusNotFound describes a response with status code 404, with default header values.

The resource specified does not exist.
*/
type UpdateShipmentStatusNotFound struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *orders_v0_models.UpdateShipmentStatusErrorResponse
}

// IsSuccess returns true when this update shipment status not found response has a 2xx status code
func (o *UpdateShipmentStatusNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update shipment status not found response has a 3xx status code
func (o *UpdateShipmentStatusNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update shipment status not found response has a 4xx status code
func (o *UpdateShipmentStatusNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update shipment status not found response has a 5xx status code
func (o *UpdateShipmentStatusNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update shipment status not found response a status code equal to that given
func (o *UpdateShipmentStatusNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *UpdateShipmentStatusNotFound) Error() string {
	return fmt.Sprintf("[POST /orders/v0/orders/{orderId}/shipment][%d] updateShipmentStatusNotFound  %+v", 404, o.Payload)
}

func (o *UpdateShipmentStatusNotFound) String() string {
	return fmt.Sprintf("[POST /orders/v0/orders/{orderId}/shipment][%d] updateShipmentStatusNotFound  %+v", 404, o.Payload)
}

func (o *UpdateShipmentStatusNotFound) GetPayload() *orders_v0_models.UpdateShipmentStatusErrorResponse {
	return o.Payload
}

func (o *UpdateShipmentStatusNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(orders_v0_models.UpdateShipmentStatusErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateShipmentStatusRequestEntityTooLarge creates a UpdateShipmentStatusRequestEntityTooLarge with default headers values
func NewUpdateShipmentStatusRequestEntityTooLarge() *UpdateShipmentStatusRequestEntityTooLarge {
	return &UpdateShipmentStatusRequestEntityTooLarge{}
}

/*
UpdateShipmentStatusRequestEntityTooLarge describes a response with status code 413, with default header values.

The request size exceeded the maximum accepted size.
*/
type UpdateShipmentStatusRequestEntityTooLarge struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *orders_v0_models.UpdateShipmentStatusErrorResponse
}

// IsSuccess returns true when this update shipment status request entity too large response has a 2xx status code
func (o *UpdateShipmentStatusRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update shipment status request entity too large response has a 3xx status code
func (o *UpdateShipmentStatusRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update shipment status request entity too large response has a 4xx status code
func (o *UpdateShipmentStatusRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this update shipment status request entity too large response has a 5xx status code
func (o *UpdateShipmentStatusRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this update shipment status request entity too large response a status code equal to that given
func (o *UpdateShipmentStatusRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *UpdateShipmentStatusRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /orders/v0/orders/{orderId}/shipment][%d] updateShipmentStatusRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *UpdateShipmentStatusRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[POST /orders/v0/orders/{orderId}/shipment][%d] updateShipmentStatusRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *UpdateShipmentStatusRequestEntityTooLarge) GetPayload() *orders_v0_models.UpdateShipmentStatusErrorResponse {
	return o.Payload
}

func (o *UpdateShipmentStatusRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(orders_v0_models.UpdateShipmentStatusErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateShipmentStatusUnsupportedMediaType creates a UpdateShipmentStatusUnsupportedMediaType with default headers values
func NewUpdateShipmentStatusUnsupportedMediaType() *UpdateShipmentStatusUnsupportedMediaType {
	return &UpdateShipmentStatusUnsupportedMediaType{}
}

/*
UpdateShipmentStatusUnsupportedMediaType describes a response with status code 415, with default header values.

The request payload is in an unsupported format.
*/
type UpdateShipmentStatusUnsupportedMediaType struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *orders_v0_models.UpdateShipmentStatusErrorResponse
}

// IsSuccess returns true when this update shipment status unsupported media type response has a 2xx status code
func (o *UpdateShipmentStatusUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update shipment status unsupported media type response has a 3xx status code
func (o *UpdateShipmentStatusUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update shipment status unsupported media type response has a 4xx status code
func (o *UpdateShipmentStatusUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this update shipment status unsupported media type response has a 5xx status code
func (o *UpdateShipmentStatusUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this update shipment status unsupported media type response a status code equal to that given
func (o *UpdateShipmentStatusUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *UpdateShipmentStatusUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /orders/v0/orders/{orderId}/shipment][%d] updateShipmentStatusUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *UpdateShipmentStatusUnsupportedMediaType) String() string {
	return fmt.Sprintf("[POST /orders/v0/orders/{orderId}/shipment][%d] updateShipmentStatusUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *UpdateShipmentStatusUnsupportedMediaType) GetPayload() *orders_v0_models.UpdateShipmentStatusErrorResponse {
	return o.Payload
}

func (o *UpdateShipmentStatusUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(orders_v0_models.UpdateShipmentStatusErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateShipmentStatusTooManyRequests creates a UpdateShipmentStatusTooManyRequests with default headers values
func NewUpdateShipmentStatusTooManyRequests() *UpdateShipmentStatusTooManyRequests {
	return &UpdateShipmentStatusTooManyRequests{}
}

/*
UpdateShipmentStatusTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type UpdateShipmentStatusTooManyRequests struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *orders_v0_models.UpdateShipmentStatusErrorResponse
}

// IsSuccess returns true when this update shipment status too many requests response has a 2xx status code
func (o *UpdateShipmentStatusTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update shipment status too many requests response has a 3xx status code
func (o *UpdateShipmentStatusTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update shipment status too many requests response has a 4xx status code
func (o *UpdateShipmentStatusTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this update shipment status too many requests response has a 5xx status code
func (o *UpdateShipmentStatusTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this update shipment status too many requests response a status code equal to that given
func (o *UpdateShipmentStatusTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *UpdateShipmentStatusTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /orders/v0/orders/{orderId}/shipment][%d] updateShipmentStatusTooManyRequests  %+v", 429, o.Payload)
}

func (o *UpdateShipmentStatusTooManyRequests) String() string {
	return fmt.Sprintf("[POST /orders/v0/orders/{orderId}/shipment][%d] updateShipmentStatusTooManyRequests  %+v", 429, o.Payload)
}

func (o *UpdateShipmentStatusTooManyRequests) GetPayload() *orders_v0_models.UpdateShipmentStatusErrorResponse {
	return o.Payload
}

func (o *UpdateShipmentStatusTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(orders_v0_models.UpdateShipmentStatusErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateShipmentStatusInternalServerError creates a UpdateShipmentStatusInternalServerError with default headers values
func NewUpdateShipmentStatusInternalServerError() *UpdateShipmentStatusInternalServerError {
	return &UpdateShipmentStatusInternalServerError{}
}

/*
UpdateShipmentStatusInternalServerError describes a response with status code 500, with default header values.

An unexpected condition occurred that prevented the server from fulfilling the request.
*/
type UpdateShipmentStatusInternalServerError struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *orders_v0_models.UpdateShipmentStatusErrorResponse
}

// IsSuccess returns true when this update shipment status internal server error response has a 2xx status code
func (o *UpdateShipmentStatusInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update shipment status internal server error response has a 3xx status code
func (o *UpdateShipmentStatusInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update shipment status internal server error response has a 4xx status code
func (o *UpdateShipmentStatusInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this update shipment status internal server error response has a 5xx status code
func (o *UpdateShipmentStatusInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this update shipment status internal server error response a status code equal to that given
func (o *UpdateShipmentStatusInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *UpdateShipmentStatusInternalServerError) Error() string {
	return fmt.Sprintf("[POST /orders/v0/orders/{orderId}/shipment][%d] updateShipmentStatusInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateShipmentStatusInternalServerError) String() string {
	return fmt.Sprintf("[POST /orders/v0/orders/{orderId}/shipment][%d] updateShipmentStatusInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateShipmentStatusInternalServerError) GetPayload() *orders_v0_models.UpdateShipmentStatusErrorResponse {
	return o.Payload
}

func (o *UpdateShipmentStatusInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(orders_v0_models.UpdateShipmentStatusErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateShipmentStatusServiceUnavailable creates a UpdateShipmentStatusServiceUnavailable with default headers values
func NewUpdateShipmentStatusServiceUnavailable() *UpdateShipmentStatusServiceUnavailable {
	return &UpdateShipmentStatusServiceUnavailable{}
}

/*
UpdateShipmentStatusServiceUnavailable describes a response with status code 503, with default header values.

Temporary overloading or maintenance of the server.
*/
type UpdateShipmentStatusServiceUnavailable struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *orders_v0_models.UpdateShipmentStatusErrorResponse
}

// IsSuccess returns true when this update shipment status service unavailable response has a 2xx status code
func (o *UpdateShipmentStatusServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update shipment status service unavailable response has a 3xx status code
func (o *UpdateShipmentStatusServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update shipment status service unavailable response has a 4xx status code
func (o *UpdateShipmentStatusServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this update shipment status service unavailable response has a 5xx status code
func (o *UpdateShipmentStatusServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this update shipment status service unavailable response a status code equal to that given
func (o *UpdateShipmentStatusServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *UpdateShipmentStatusServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /orders/v0/orders/{orderId}/shipment][%d] updateShipmentStatusServiceUnavailable  %+v", 503, o.Payload)
}

func (o *UpdateShipmentStatusServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /orders/v0/orders/{orderId}/shipment][%d] updateShipmentStatusServiceUnavailable  %+v", 503, o.Payload)
}

func (o *UpdateShipmentStatusServiceUnavailable) GetPayload() *orders_v0_models.UpdateShipmentStatusErrorResponse {
	return o.Payload
}

func (o *UpdateShipmentStatusServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(orders_v0_models.UpdateShipmentStatusErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}