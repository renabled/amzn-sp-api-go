// Code generated by go-swagger; DO NOT EDIT.

package update_inventory

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/renabled/amzn-sp-api-go/api/vendorDirectFulfillmentInventoryV1/vendor_direct_fulfillment_inventory_v1_models"
)

// SubmitInventoryUpdateReader is a Reader for the SubmitInventoryUpdate structure.
type SubmitInventoryUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SubmitInventoryUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewSubmitInventoryUpdateAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSubmitInventoryUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSubmitInventoryUpdateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSubmitInventoryUpdateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewSubmitInventoryUpdateRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewSubmitInventoryUpdateUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewSubmitInventoryUpdateTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSubmitInventoryUpdateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewSubmitInventoryUpdateServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSubmitInventoryUpdateAccepted creates a SubmitInventoryUpdateAccepted with default headers values
func NewSubmitInventoryUpdateAccepted() *SubmitInventoryUpdateAccepted {
	return &SubmitInventoryUpdateAccepted{}
}

/*
SubmitInventoryUpdateAccepted describes a response with status code 202, with default header values.

Success.
*/
type SubmitInventoryUpdateAccepted struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_inventory_v1_models.SubmitInventoryUpdateResponse
}

// IsSuccess returns true when this submit inventory update accepted response has a 2xx status code
func (o *SubmitInventoryUpdateAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this submit inventory update accepted response has a 3xx status code
func (o *SubmitInventoryUpdateAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this submit inventory update accepted response has a 4xx status code
func (o *SubmitInventoryUpdateAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this submit inventory update accepted response has a 5xx status code
func (o *SubmitInventoryUpdateAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this submit inventory update accepted response a status code equal to that given
func (o *SubmitInventoryUpdateAccepted) IsCode(code int) bool {
	return code == 202
}

func (o *SubmitInventoryUpdateAccepted) Error() string {
	return fmt.Sprintf("[POST /vendor/directFulfillment/inventory/v1/warehouses/{warehouseId}/items][%d] submitInventoryUpdateAccepted  %+v", 202, o.Payload)
}

func (o *SubmitInventoryUpdateAccepted) String() string {
	return fmt.Sprintf("[POST /vendor/directFulfillment/inventory/v1/warehouses/{warehouseId}/items][%d] submitInventoryUpdateAccepted  %+v", 202, o.Payload)
}

func (o *SubmitInventoryUpdateAccepted) GetPayload() *vendor_direct_fulfillment_inventory_v1_models.SubmitInventoryUpdateResponse {
	return o.Payload
}

func (o *SubmitInventoryUpdateAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(vendor_direct_fulfillment_inventory_v1_models.SubmitInventoryUpdateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSubmitInventoryUpdateBadRequest creates a SubmitInventoryUpdateBadRequest with default headers values
func NewSubmitInventoryUpdateBadRequest() *SubmitInventoryUpdateBadRequest {
	return &SubmitInventoryUpdateBadRequest{}
}

/*
SubmitInventoryUpdateBadRequest describes a response with status code 400, with default header values.

Request has missing or invalid parameters and cannot be parsed.
*/
type SubmitInventoryUpdateBadRequest struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_inventory_v1_models.SubmitInventoryUpdateResponse
}

// IsSuccess returns true when this submit inventory update bad request response has a 2xx status code
func (o *SubmitInventoryUpdateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this submit inventory update bad request response has a 3xx status code
func (o *SubmitInventoryUpdateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this submit inventory update bad request response has a 4xx status code
func (o *SubmitInventoryUpdateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this submit inventory update bad request response has a 5xx status code
func (o *SubmitInventoryUpdateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this submit inventory update bad request response a status code equal to that given
func (o *SubmitInventoryUpdateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *SubmitInventoryUpdateBadRequest) Error() string {
	return fmt.Sprintf("[POST /vendor/directFulfillment/inventory/v1/warehouses/{warehouseId}/items][%d] submitInventoryUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *SubmitInventoryUpdateBadRequest) String() string {
	return fmt.Sprintf("[POST /vendor/directFulfillment/inventory/v1/warehouses/{warehouseId}/items][%d] submitInventoryUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *SubmitInventoryUpdateBadRequest) GetPayload() *vendor_direct_fulfillment_inventory_v1_models.SubmitInventoryUpdateResponse {
	return o.Payload
}

func (o *SubmitInventoryUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(vendor_direct_fulfillment_inventory_v1_models.SubmitInventoryUpdateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSubmitInventoryUpdateForbidden creates a SubmitInventoryUpdateForbidden with default headers values
func NewSubmitInventoryUpdateForbidden() *SubmitInventoryUpdateForbidden {
	return &SubmitInventoryUpdateForbidden{}
}

/*
SubmitInventoryUpdateForbidden describes a response with status code 403, with default header values.

Indicates that access to the resource is forbidden. Possible reasons include Access Denied, Unauthorized, Expired Token, or Invalid Signature.
*/
type SubmitInventoryUpdateForbidden struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_inventory_v1_models.SubmitInventoryUpdateResponse
}

// IsSuccess returns true when this submit inventory update forbidden response has a 2xx status code
func (o *SubmitInventoryUpdateForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this submit inventory update forbidden response has a 3xx status code
func (o *SubmitInventoryUpdateForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this submit inventory update forbidden response has a 4xx status code
func (o *SubmitInventoryUpdateForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this submit inventory update forbidden response has a 5xx status code
func (o *SubmitInventoryUpdateForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this submit inventory update forbidden response a status code equal to that given
func (o *SubmitInventoryUpdateForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *SubmitInventoryUpdateForbidden) Error() string {
	return fmt.Sprintf("[POST /vendor/directFulfillment/inventory/v1/warehouses/{warehouseId}/items][%d] submitInventoryUpdateForbidden  %+v", 403, o.Payload)
}

func (o *SubmitInventoryUpdateForbidden) String() string {
	return fmt.Sprintf("[POST /vendor/directFulfillment/inventory/v1/warehouses/{warehouseId}/items][%d] submitInventoryUpdateForbidden  %+v", 403, o.Payload)
}

func (o *SubmitInventoryUpdateForbidden) GetPayload() *vendor_direct_fulfillment_inventory_v1_models.SubmitInventoryUpdateResponse {
	return o.Payload
}

func (o *SubmitInventoryUpdateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(vendor_direct_fulfillment_inventory_v1_models.SubmitInventoryUpdateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSubmitInventoryUpdateNotFound creates a SubmitInventoryUpdateNotFound with default headers values
func NewSubmitInventoryUpdateNotFound() *SubmitInventoryUpdateNotFound {
	return &SubmitInventoryUpdateNotFound{}
}

/*
SubmitInventoryUpdateNotFound describes a response with status code 404, with default header values.

The resource specified does not exist.
*/
type SubmitInventoryUpdateNotFound struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_inventory_v1_models.SubmitInventoryUpdateResponse
}

// IsSuccess returns true when this submit inventory update not found response has a 2xx status code
func (o *SubmitInventoryUpdateNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this submit inventory update not found response has a 3xx status code
func (o *SubmitInventoryUpdateNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this submit inventory update not found response has a 4xx status code
func (o *SubmitInventoryUpdateNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this submit inventory update not found response has a 5xx status code
func (o *SubmitInventoryUpdateNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this submit inventory update not found response a status code equal to that given
func (o *SubmitInventoryUpdateNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *SubmitInventoryUpdateNotFound) Error() string {
	return fmt.Sprintf("[POST /vendor/directFulfillment/inventory/v1/warehouses/{warehouseId}/items][%d] submitInventoryUpdateNotFound  %+v", 404, o.Payload)
}

func (o *SubmitInventoryUpdateNotFound) String() string {
	return fmt.Sprintf("[POST /vendor/directFulfillment/inventory/v1/warehouses/{warehouseId}/items][%d] submitInventoryUpdateNotFound  %+v", 404, o.Payload)
}

func (o *SubmitInventoryUpdateNotFound) GetPayload() *vendor_direct_fulfillment_inventory_v1_models.SubmitInventoryUpdateResponse {
	return o.Payload
}

func (o *SubmitInventoryUpdateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(vendor_direct_fulfillment_inventory_v1_models.SubmitInventoryUpdateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSubmitInventoryUpdateRequestEntityTooLarge creates a SubmitInventoryUpdateRequestEntityTooLarge with default headers values
func NewSubmitInventoryUpdateRequestEntityTooLarge() *SubmitInventoryUpdateRequestEntityTooLarge {
	return &SubmitInventoryUpdateRequestEntityTooLarge{}
}

/*
SubmitInventoryUpdateRequestEntityTooLarge describes a response with status code 413, with default header values.

The request size exceeded the maximum accepted size.
*/
type SubmitInventoryUpdateRequestEntityTooLarge struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_inventory_v1_models.SubmitInventoryUpdateResponse
}

// IsSuccess returns true when this submit inventory update request entity too large response has a 2xx status code
func (o *SubmitInventoryUpdateRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this submit inventory update request entity too large response has a 3xx status code
func (o *SubmitInventoryUpdateRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this submit inventory update request entity too large response has a 4xx status code
func (o *SubmitInventoryUpdateRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this submit inventory update request entity too large response has a 5xx status code
func (o *SubmitInventoryUpdateRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this submit inventory update request entity too large response a status code equal to that given
func (o *SubmitInventoryUpdateRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *SubmitInventoryUpdateRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /vendor/directFulfillment/inventory/v1/warehouses/{warehouseId}/items][%d] submitInventoryUpdateRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *SubmitInventoryUpdateRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[POST /vendor/directFulfillment/inventory/v1/warehouses/{warehouseId}/items][%d] submitInventoryUpdateRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *SubmitInventoryUpdateRequestEntityTooLarge) GetPayload() *vendor_direct_fulfillment_inventory_v1_models.SubmitInventoryUpdateResponse {
	return o.Payload
}

func (o *SubmitInventoryUpdateRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(vendor_direct_fulfillment_inventory_v1_models.SubmitInventoryUpdateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSubmitInventoryUpdateUnsupportedMediaType creates a SubmitInventoryUpdateUnsupportedMediaType with default headers values
func NewSubmitInventoryUpdateUnsupportedMediaType() *SubmitInventoryUpdateUnsupportedMediaType {
	return &SubmitInventoryUpdateUnsupportedMediaType{}
}

/*
SubmitInventoryUpdateUnsupportedMediaType describes a response with status code 415, with default header values.

The request payload is in an unsupported format.
*/
type SubmitInventoryUpdateUnsupportedMediaType struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_inventory_v1_models.SubmitInventoryUpdateResponse
}

// IsSuccess returns true when this submit inventory update unsupported media type response has a 2xx status code
func (o *SubmitInventoryUpdateUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this submit inventory update unsupported media type response has a 3xx status code
func (o *SubmitInventoryUpdateUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this submit inventory update unsupported media type response has a 4xx status code
func (o *SubmitInventoryUpdateUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this submit inventory update unsupported media type response has a 5xx status code
func (o *SubmitInventoryUpdateUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this submit inventory update unsupported media type response a status code equal to that given
func (o *SubmitInventoryUpdateUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *SubmitInventoryUpdateUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /vendor/directFulfillment/inventory/v1/warehouses/{warehouseId}/items][%d] submitInventoryUpdateUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *SubmitInventoryUpdateUnsupportedMediaType) String() string {
	return fmt.Sprintf("[POST /vendor/directFulfillment/inventory/v1/warehouses/{warehouseId}/items][%d] submitInventoryUpdateUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *SubmitInventoryUpdateUnsupportedMediaType) GetPayload() *vendor_direct_fulfillment_inventory_v1_models.SubmitInventoryUpdateResponse {
	return o.Payload
}

func (o *SubmitInventoryUpdateUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(vendor_direct_fulfillment_inventory_v1_models.SubmitInventoryUpdateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSubmitInventoryUpdateTooManyRequests creates a SubmitInventoryUpdateTooManyRequests with default headers values
func NewSubmitInventoryUpdateTooManyRequests() *SubmitInventoryUpdateTooManyRequests {
	return &SubmitInventoryUpdateTooManyRequests{}
}

/*
SubmitInventoryUpdateTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type SubmitInventoryUpdateTooManyRequests struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_inventory_v1_models.SubmitInventoryUpdateResponse
}

// IsSuccess returns true when this submit inventory update too many requests response has a 2xx status code
func (o *SubmitInventoryUpdateTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this submit inventory update too many requests response has a 3xx status code
func (o *SubmitInventoryUpdateTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this submit inventory update too many requests response has a 4xx status code
func (o *SubmitInventoryUpdateTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this submit inventory update too many requests response has a 5xx status code
func (o *SubmitInventoryUpdateTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this submit inventory update too many requests response a status code equal to that given
func (o *SubmitInventoryUpdateTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *SubmitInventoryUpdateTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /vendor/directFulfillment/inventory/v1/warehouses/{warehouseId}/items][%d] submitInventoryUpdateTooManyRequests  %+v", 429, o.Payload)
}

func (o *SubmitInventoryUpdateTooManyRequests) String() string {
	return fmt.Sprintf("[POST /vendor/directFulfillment/inventory/v1/warehouses/{warehouseId}/items][%d] submitInventoryUpdateTooManyRequests  %+v", 429, o.Payload)
}

func (o *SubmitInventoryUpdateTooManyRequests) GetPayload() *vendor_direct_fulfillment_inventory_v1_models.SubmitInventoryUpdateResponse {
	return o.Payload
}

func (o *SubmitInventoryUpdateTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(vendor_direct_fulfillment_inventory_v1_models.SubmitInventoryUpdateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSubmitInventoryUpdateInternalServerError creates a SubmitInventoryUpdateInternalServerError with default headers values
func NewSubmitInventoryUpdateInternalServerError() *SubmitInventoryUpdateInternalServerError {
	return &SubmitInventoryUpdateInternalServerError{}
}

/*
SubmitInventoryUpdateInternalServerError describes a response with status code 500, with default header values.

An unexpected condition occurred that prevented the server from fulfilling the request.
*/
type SubmitInventoryUpdateInternalServerError struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_inventory_v1_models.SubmitInventoryUpdateResponse
}

// IsSuccess returns true when this submit inventory update internal server error response has a 2xx status code
func (o *SubmitInventoryUpdateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this submit inventory update internal server error response has a 3xx status code
func (o *SubmitInventoryUpdateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this submit inventory update internal server error response has a 4xx status code
func (o *SubmitInventoryUpdateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this submit inventory update internal server error response has a 5xx status code
func (o *SubmitInventoryUpdateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this submit inventory update internal server error response a status code equal to that given
func (o *SubmitInventoryUpdateInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *SubmitInventoryUpdateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /vendor/directFulfillment/inventory/v1/warehouses/{warehouseId}/items][%d] submitInventoryUpdateInternalServerError  %+v", 500, o.Payload)
}

func (o *SubmitInventoryUpdateInternalServerError) String() string {
	return fmt.Sprintf("[POST /vendor/directFulfillment/inventory/v1/warehouses/{warehouseId}/items][%d] submitInventoryUpdateInternalServerError  %+v", 500, o.Payload)
}

func (o *SubmitInventoryUpdateInternalServerError) GetPayload() *vendor_direct_fulfillment_inventory_v1_models.SubmitInventoryUpdateResponse {
	return o.Payload
}

func (o *SubmitInventoryUpdateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(vendor_direct_fulfillment_inventory_v1_models.SubmitInventoryUpdateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSubmitInventoryUpdateServiceUnavailable creates a SubmitInventoryUpdateServiceUnavailable with default headers values
func NewSubmitInventoryUpdateServiceUnavailable() *SubmitInventoryUpdateServiceUnavailable {
	return &SubmitInventoryUpdateServiceUnavailable{}
}

/*
SubmitInventoryUpdateServiceUnavailable describes a response with status code 503, with default header values.

Temporary overloading or maintenance of the server.
*/
type SubmitInventoryUpdateServiceUnavailable struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_inventory_v1_models.SubmitInventoryUpdateResponse
}

// IsSuccess returns true when this submit inventory update service unavailable response has a 2xx status code
func (o *SubmitInventoryUpdateServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this submit inventory update service unavailable response has a 3xx status code
func (o *SubmitInventoryUpdateServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this submit inventory update service unavailable response has a 4xx status code
func (o *SubmitInventoryUpdateServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this submit inventory update service unavailable response has a 5xx status code
func (o *SubmitInventoryUpdateServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this submit inventory update service unavailable response a status code equal to that given
func (o *SubmitInventoryUpdateServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *SubmitInventoryUpdateServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /vendor/directFulfillment/inventory/v1/warehouses/{warehouseId}/items][%d] submitInventoryUpdateServiceUnavailable  %+v", 503, o.Payload)
}

func (o *SubmitInventoryUpdateServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /vendor/directFulfillment/inventory/v1/warehouses/{warehouseId}/items][%d] submitInventoryUpdateServiceUnavailable  %+v", 503, o.Payload)
}

func (o *SubmitInventoryUpdateServiceUnavailable) GetPayload() *vendor_direct_fulfillment_inventory_v1_models.SubmitInventoryUpdateResponse {
	return o.Payload
}

func (o *SubmitInventoryUpdateServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(vendor_direct_fulfillment_inventory_v1_models.SubmitInventoryUpdateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
