// Code generated by go-swagger; DO NOT EDIT.

package fba_inbound

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/renabled/amzn-sp-api-go/api/fulfillmentInboundV0/fulfillment_inbound_v0_models"
)

// GetLabelsReader is a Reader for the GetLabels structure.
type GetLabelsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLabelsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLabelsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetLabelsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetLabelsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetLabelsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetLabelsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetLabelsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetLabelsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetLabelsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetLabelsOK creates a GetLabelsOK with default headers values
func NewGetLabelsOK() *GetLabelsOK {
	return &GetLabelsOK{}
}

/*
GetLabelsOK describes a response with status code 200, with default header values.

Success.
*/
type GetLabelsOK struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inbound_v0_models.GetLabelsResponse
}

// IsSuccess returns true when this get labels o k response has a 2xx status code
func (o *GetLabelsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get labels o k response has a 3xx status code
func (o *GetLabelsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get labels o k response has a 4xx status code
func (o *GetLabelsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get labels o k response has a 5xx status code
func (o *GetLabelsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get labels o k response a status code equal to that given
func (o *GetLabelsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetLabelsOK) Error() string {
	return fmt.Sprintf("[GET /fba/inbound/v0/shipments/{shipmentId}/labels][%d] getLabelsOK  %+v", 200, o.Payload)
}

func (o *GetLabelsOK) String() string {
	return fmt.Sprintf("[GET /fba/inbound/v0/shipments/{shipmentId}/labels][%d] getLabelsOK  %+v", 200, o.Payload)
}

func (o *GetLabelsOK) GetPayload() *fulfillment_inbound_v0_models.GetLabelsResponse {
	return o.Payload
}

func (o *GetLabelsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fulfillment_inbound_v0_models.GetLabelsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLabelsBadRequest creates a GetLabelsBadRequest with default headers values
func NewGetLabelsBadRequest() *GetLabelsBadRequest {
	return &GetLabelsBadRequest{}
}

/*
GetLabelsBadRequest describes a response with status code 400, with default header values.

Request has missing or invalid parameters and cannot be parsed.
*/
type GetLabelsBadRequest struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inbound_v0_models.GetLabelsResponse
}

// IsSuccess returns true when this get labels bad request response has a 2xx status code
func (o *GetLabelsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get labels bad request response has a 3xx status code
func (o *GetLabelsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get labels bad request response has a 4xx status code
func (o *GetLabelsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get labels bad request response has a 5xx status code
func (o *GetLabelsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get labels bad request response a status code equal to that given
func (o *GetLabelsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetLabelsBadRequest) Error() string {
	return fmt.Sprintf("[GET /fba/inbound/v0/shipments/{shipmentId}/labels][%d] getLabelsBadRequest  %+v", 400, o.Payload)
}

func (o *GetLabelsBadRequest) String() string {
	return fmt.Sprintf("[GET /fba/inbound/v0/shipments/{shipmentId}/labels][%d] getLabelsBadRequest  %+v", 400, o.Payload)
}

func (o *GetLabelsBadRequest) GetPayload() *fulfillment_inbound_v0_models.GetLabelsResponse {
	return o.Payload
}

func (o *GetLabelsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fulfillment_inbound_v0_models.GetLabelsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLabelsUnauthorized creates a GetLabelsUnauthorized with default headers values
func NewGetLabelsUnauthorized() *GetLabelsUnauthorized {
	return &GetLabelsUnauthorized{}
}

/*
GetLabelsUnauthorized describes a response with status code 401, with default header values.

The request's Authorization header is not formatted correctly or does not contain a valid token.
*/
type GetLabelsUnauthorized struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inbound_v0_models.GetLabelsResponse
}

// IsSuccess returns true when this get labels unauthorized response has a 2xx status code
func (o *GetLabelsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get labels unauthorized response has a 3xx status code
func (o *GetLabelsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get labels unauthorized response has a 4xx status code
func (o *GetLabelsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get labels unauthorized response has a 5xx status code
func (o *GetLabelsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get labels unauthorized response a status code equal to that given
func (o *GetLabelsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetLabelsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /fba/inbound/v0/shipments/{shipmentId}/labels][%d] getLabelsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetLabelsUnauthorized) String() string {
	return fmt.Sprintf("[GET /fba/inbound/v0/shipments/{shipmentId}/labels][%d] getLabelsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetLabelsUnauthorized) GetPayload() *fulfillment_inbound_v0_models.GetLabelsResponse {
	return o.Payload
}

func (o *GetLabelsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(fulfillment_inbound_v0_models.GetLabelsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLabelsForbidden creates a GetLabelsForbidden with default headers values
func NewGetLabelsForbidden() *GetLabelsForbidden {
	return &GetLabelsForbidden{}
}

/*
GetLabelsForbidden describes a response with status code 403, with default header values.

Indicates that access to the resource is forbidden. Possible reasons include Access Denied, Unauthorized, Expired Token, or Invalid Signature.
*/
type GetLabelsForbidden struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inbound_v0_models.GetLabelsResponse
}

// IsSuccess returns true when this get labels forbidden response has a 2xx status code
func (o *GetLabelsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get labels forbidden response has a 3xx status code
func (o *GetLabelsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get labels forbidden response has a 4xx status code
func (o *GetLabelsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get labels forbidden response has a 5xx status code
func (o *GetLabelsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get labels forbidden response a status code equal to that given
func (o *GetLabelsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetLabelsForbidden) Error() string {
	return fmt.Sprintf("[GET /fba/inbound/v0/shipments/{shipmentId}/labels][%d] getLabelsForbidden  %+v", 403, o.Payload)
}

func (o *GetLabelsForbidden) String() string {
	return fmt.Sprintf("[GET /fba/inbound/v0/shipments/{shipmentId}/labels][%d] getLabelsForbidden  %+v", 403, o.Payload)
}

func (o *GetLabelsForbidden) GetPayload() *fulfillment_inbound_v0_models.GetLabelsResponse {
	return o.Payload
}

func (o *GetLabelsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(fulfillment_inbound_v0_models.GetLabelsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLabelsNotFound creates a GetLabelsNotFound with default headers values
func NewGetLabelsNotFound() *GetLabelsNotFound {
	return &GetLabelsNotFound{}
}

/*
GetLabelsNotFound describes a response with status code 404, with default header values.

The specified resource does not exist.
*/
type GetLabelsNotFound struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inbound_v0_models.GetLabelsResponse
}

// IsSuccess returns true when this get labels not found response has a 2xx status code
func (o *GetLabelsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get labels not found response has a 3xx status code
func (o *GetLabelsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get labels not found response has a 4xx status code
func (o *GetLabelsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get labels not found response has a 5xx status code
func (o *GetLabelsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get labels not found response a status code equal to that given
func (o *GetLabelsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetLabelsNotFound) Error() string {
	return fmt.Sprintf("[GET /fba/inbound/v0/shipments/{shipmentId}/labels][%d] getLabelsNotFound  %+v", 404, o.Payload)
}

func (o *GetLabelsNotFound) String() string {
	return fmt.Sprintf("[GET /fba/inbound/v0/shipments/{shipmentId}/labels][%d] getLabelsNotFound  %+v", 404, o.Payload)
}

func (o *GetLabelsNotFound) GetPayload() *fulfillment_inbound_v0_models.GetLabelsResponse {
	return o.Payload
}

func (o *GetLabelsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fulfillment_inbound_v0_models.GetLabelsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLabelsTooManyRequests creates a GetLabelsTooManyRequests with default headers values
func NewGetLabelsTooManyRequests() *GetLabelsTooManyRequests {
	return &GetLabelsTooManyRequests{}
}

/*
GetLabelsTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type GetLabelsTooManyRequests struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inbound_v0_models.GetLabelsResponse
}

// IsSuccess returns true when this get labels too many requests response has a 2xx status code
func (o *GetLabelsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get labels too many requests response has a 3xx status code
func (o *GetLabelsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get labels too many requests response has a 4xx status code
func (o *GetLabelsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get labels too many requests response has a 5xx status code
func (o *GetLabelsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get labels too many requests response a status code equal to that given
func (o *GetLabelsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetLabelsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /fba/inbound/v0/shipments/{shipmentId}/labels][%d] getLabelsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetLabelsTooManyRequests) String() string {
	return fmt.Sprintf("[GET /fba/inbound/v0/shipments/{shipmentId}/labels][%d] getLabelsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetLabelsTooManyRequests) GetPayload() *fulfillment_inbound_v0_models.GetLabelsResponse {
	return o.Payload
}

func (o *GetLabelsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(fulfillment_inbound_v0_models.GetLabelsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLabelsInternalServerError creates a GetLabelsInternalServerError with default headers values
func NewGetLabelsInternalServerError() *GetLabelsInternalServerError {
	return &GetLabelsInternalServerError{}
}

/*
GetLabelsInternalServerError describes a response with status code 500, with default header values.

An unexpected condition occurred that prevented the server from fulfilling the request.
*/
type GetLabelsInternalServerError struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inbound_v0_models.GetLabelsResponse
}

// IsSuccess returns true when this get labels internal server error response has a 2xx status code
func (o *GetLabelsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get labels internal server error response has a 3xx status code
func (o *GetLabelsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get labels internal server error response has a 4xx status code
func (o *GetLabelsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get labels internal server error response has a 5xx status code
func (o *GetLabelsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get labels internal server error response a status code equal to that given
func (o *GetLabelsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetLabelsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /fba/inbound/v0/shipments/{shipmentId}/labels][%d] getLabelsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetLabelsInternalServerError) String() string {
	return fmt.Sprintf("[GET /fba/inbound/v0/shipments/{shipmentId}/labels][%d] getLabelsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetLabelsInternalServerError) GetPayload() *fulfillment_inbound_v0_models.GetLabelsResponse {
	return o.Payload
}

func (o *GetLabelsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(fulfillment_inbound_v0_models.GetLabelsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLabelsServiceUnavailable creates a GetLabelsServiceUnavailable with default headers values
func NewGetLabelsServiceUnavailable() *GetLabelsServiceUnavailable {
	return &GetLabelsServiceUnavailable{}
}

/*
GetLabelsServiceUnavailable describes a response with status code 503, with default header values.

Temporary overloading or maintenance of the server.
*/
type GetLabelsServiceUnavailable struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inbound_v0_models.GetLabelsResponse
}

// IsSuccess returns true when this get labels service unavailable response has a 2xx status code
func (o *GetLabelsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get labels service unavailable response has a 3xx status code
func (o *GetLabelsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get labels service unavailable response has a 4xx status code
func (o *GetLabelsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get labels service unavailable response has a 5xx status code
func (o *GetLabelsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get labels service unavailable response a status code equal to that given
func (o *GetLabelsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetLabelsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /fba/inbound/v0/shipments/{shipmentId}/labels][%d] getLabelsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetLabelsServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /fba/inbound/v0/shipments/{shipmentId}/labels][%d] getLabelsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetLabelsServiceUnavailable) GetPayload() *fulfillment_inbound_v0_models.GetLabelsResponse {
	return o.Payload
}

func (o *GetLabelsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(fulfillment_inbound_v0_models.GetLabelsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
