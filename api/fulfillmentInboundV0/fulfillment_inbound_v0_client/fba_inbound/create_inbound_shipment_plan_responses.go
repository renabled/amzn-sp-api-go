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

// CreateInboundShipmentPlanReader is a Reader for the CreateInboundShipmentPlan structure.
type CreateInboundShipmentPlanReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateInboundShipmentPlanReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateInboundShipmentPlanOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateInboundShipmentPlanBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateInboundShipmentPlanUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateInboundShipmentPlanForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateInboundShipmentPlanNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewCreateInboundShipmentPlanTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateInboundShipmentPlanInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewCreateInboundShipmentPlanServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateInboundShipmentPlanOK creates a CreateInboundShipmentPlanOK with default headers values
func NewCreateInboundShipmentPlanOK() *CreateInboundShipmentPlanOK {
	return &CreateInboundShipmentPlanOK{}
}

/*
CreateInboundShipmentPlanOK describes a response with status code 200, with default header values.

Success.
*/
type CreateInboundShipmentPlanOK struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inbound_v0_models.CreateInboundShipmentPlanResponse
}

// IsSuccess returns true when this create inbound shipment plan o k response has a 2xx status code
func (o *CreateInboundShipmentPlanOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create inbound shipment plan o k response has a 3xx status code
func (o *CreateInboundShipmentPlanOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create inbound shipment plan o k response has a 4xx status code
func (o *CreateInboundShipmentPlanOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create inbound shipment plan o k response has a 5xx status code
func (o *CreateInboundShipmentPlanOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create inbound shipment plan o k response a status code equal to that given
func (o *CreateInboundShipmentPlanOK) IsCode(code int) bool {
	return code == 200
}

func (o *CreateInboundShipmentPlanOK) Error() string {
	return fmt.Sprintf("[POST /fba/inbound/v0/plans][%d] createInboundShipmentPlanOK  %+v", 200, o.Payload)
}

func (o *CreateInboundShipmentPlanOK) String() string {
	return fmt.Sprintf("[POST /fba/inbound/v0/plans][%d] createInboundShipmentPlanOK  %+v", 200, o.Payload)
}

func (o *CreateInboundShipmentPlanOK) GetPayload() *fulfillment_inbound_v0_models.CreateInboundShipmentPlanResponse {
	return o.Payload
}

func (o *CreateInboundShipmentPlanOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fulfillment_inbound_v0_models.CreateInboundShipmentPlanResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateInboundShipmentPlanBadRequest creates a CreateInboundShipmentPlanBadRequest with default headers values
func NewCreateInboundShipmentPlanBadRequest() *CreateInboundShipmentPlanBadRequest {
	return &CreateInboundShipmentPlanBadRequest{}
}

/*
CreateInboundShipmentPlanBadRequest describes a response with status code 400, with default header values.

Request has missing or invalid parameters and cannot be parsed.
*/
type CreateInboundShipmentPlanBadRequest struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inbound_v0_models.CreateInboundShipmentPlanResponse
}

// IsSuccess returns true when this create inbound shipment plan bad request response has a 2xx status code
func (o *CreateInboundShipmentPlanBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create inbound shipment plan bad request response has a 3xx status code
func (o *CreateInboundShipmentPlanBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create inbound shipment plan bad request response has a 4xx status code
func (o *CreateInboundShipmentPlanBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create inbound shipment plan bad request response has a 5xx status code
func (o *CreateInboundShipmentPlanBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create inbound shipment plan bad request response a status code equal to that given
func (o *CreateInboundShipmentPlanBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *CreateInboundShipmentPlanBadRequest) Error() string {
	return fmt.Sprintf("[POST /fba/inbound/v0/plans][%d] createInboundShipmentPlanBadRequest  %+v", 400, o.Payload)
}

func (o *CreateInboundShipmentPlanBadRequest) String() string {
	return fmt.Sprintf("[POST /fba/inbound/v0/plans][%d] createInboundShipmentPlanBadRequest  %+v", 400, o.Payload)
}

func (o *CreateInboundShipmentPlanBadRequest) GetPayload() *fulfillment_inbound_v0_models.CreateInboundShipmentPlanResponse {
	return o.Payload
}

func (o *CreateInboundShipmentPlanBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fulfillment_inbound_v0_models.CreateInboundShipmentPlanResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateInboundShipmentPlanUnauthorized creates a CreateInboundShipmentPlanUnauthorized with default headers values
func NewCreateInboundShipmentPlanUnauthorized() *CreateInboundShipmentPlanUnauthorized {
	return &CreateInboundShipmentPlanUnauthorized{}
}

/*
CreateInboundShipmentPlanUnauthorized describes a response with status code 401, with default header values.

The request's Authorization header is not formatted correctly or does not contain a valid token.
*/
type CreateInboundShipmentPlanUnauthorized struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inbound_v0_models.CreateInboundShipmentPlanResponse
}

// IsSuccess returns true when this create inbound shipment plan unauthorized response has a 2xx status code
func (o *CreateInboundShipmentPlanUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create inbound shipment plan unauthorized response has a 3xx status code
func (o *CreateInboundShipmentPlanUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create inbound shipment plan unauthorized response has a 4xx status code
func (o *CreateInboundShipmentPlanUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this create inbound shipment plan unauthorized response has a 5xx status code
func (o *CreateInboundShipmentPlanUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this create inbound shipment plan unauthorized response a status code equal to that given
func (o *CreateInboundShipmentPlanUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *CreateInboundShipmentPlanUnauthorized) Error() string {
	return fmt.Sprintf("[POST /fba/inbound/v0/plans][%d] createInboundShipmentPlanUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateInboundShipmentPlanUnauthorized) String() string {
	return fmt.Sprintf("[POST /fba/inbound/v0/plans][%d] createInboundShipmentPlanUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateInboundShipmentPlanUnauthorized) GetPayload() *fulfillment_inbound_v0_models.CreateInboundShipmentPlanResponse {
	return o.Payload
}

func (o *CreateInboundShipmentPlanUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(fulfillment_inbound_v0_models.CreateInboundShipmentPlanResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateInboundShipmentPlanForbidden creates a CreateInboundShipmentPlanForbidden with default headers values
func NewCreateInboundShipmentPlanForbidden() *CreateInboundShipmentPlanForbidden {
	return &CreateInboundShipmentPlanForbidden{}
}

/*
CreateInboundShipmentPlanForbidden describes a response with status code 403, with default header values.

Indicates that access to the resource is forbidden. Possible reasons include Access Denied, Unauthorized, Expired Token, or Invalid Signature.
*/
type CreateInboundShipmentPlanForbidden struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inbound_v0_models.CreateInboundShipmentPlanResponse
}

// IsSuccess returns true when this create inbound shipment plan forbidden response has a 2xx status code
func (o *CreateInboundShipmentPlanForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create inbound shipment plan forbidden response has a 3xx status code
func (o *CreateInboundShipmentPlanForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create inbound shipment plan forbidden response has a 4xx status code
func (o *CreateInboundShipmentPlanForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create inbound shipment plan forbidden response has a 5xx status code
func (o *CreateInboundShipmentPlanForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create inbound shipment plan forbidden response a status code equal to that given
func (o *CreateInboundShipmentPlanForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *CreateInboundShipmentPlanForbidden) Error() string {
	return fmt.Sprintf("[POST /fba/inbound/v0/plans][%d] createInboundShipmentPlanForbidden  %+v", 403, o.Payload)
}

func (o *CreateInboundShipmentPlanForbidden) String() string {
	return fmt.Sprintf("[POST /fba/inbound/v0/plans][%d] createInboundShipmentPlanForbidden  %+v", 403, o.Payload)
}

func (o *CreateInboundShipmentPlanForbidden) GetPayload() *fulfillment_inbound_v0_models.CreateInboundShipmentPlanResponse {
	return o.Payload
}

func (o *CreateInboundShipmentPlanForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(fulfillment_inbound_v0_models.CreateInboundShipmentPlanResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateInboundShipmentPlanNotFound creates a CreateInboundShipmentPlanNotFound with default headers values
func NewCreateInboundShipmentPlanNotFound() *CreateInboundShipmentPlanNotFound {
	return &CreateInboundShipmentPlanNotFound{}
}

/*
CreateInboundShipmentPlanNotFound describes a response with status code 404, with default header values.

The specified resource does not exist.
*/
type CreateInboundShipmentPlanNotFound struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inbound_v0_models.CreateInboundShipmentPlanResponse
}

// IsSuccess returns true when this create inbound shipment plan not found response has a 2xx status code
func (o *CreateInboundShipmentPlanNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create inbound shipment plan not found response has a 3xx status code
func (o *CreateInboundShipmentPlanNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create inbound shipment plan not found response has a 4xx status code
func (o *CreateInboundShipmentPlanNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this create inbound shipment plan not found response has a 5xx status code
func (o *CreateInboundShipmentPlanNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this create inbound shipment plan not found response a status code equal to that given
func (o *CreateInboundShipmentPlanNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *CreateInboundShipmentPlanNotFound) Error() string {
	return fmt.Sprintf("[POST /fba/inbound/v0/plans][%d] createInboundShipmentPlanNotFound  %+v", 404, o.Payload)
}

func (o *CreateInboundShipmentPlanNotFound) String() string {
	return fmt.Sprintf("[POST /fba/inbound/v0/plans][%d] createInboundShipmentPlanNotFound  %+v", 404, o.Payload)
}

func (o *CreateInboundShipmentPlanNotFound) GetPayload() *fulfillment_inbound_v0_models.CreateInboundShipmentPlanResponse {
	return o.Payload
}

func (o *CreateInboundShipmentPlanNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fulfillment_inbound_v0_models.CreateInboundShipmentPlanResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateInboundShipmentPlanTooManyRequests creates a CreateInboundShipmentPlanTooManyRequests with default headers values
func NewCreateInboundShipmentPlanTooManyRequests() *CreateInboundShipmentPlanTooManyRequests {
	return &CreateInboundShipmentPlanTooManyRequests{}
}

/*
CreateInboundShipmentPlanTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type CreateInboundShipmentPlanTooManyRequests struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inbound_v0_models.CreateInboundShipmentPlanResponse
}

// IsSuccess returns true when this create inbound shipment plan too many requests response has a 2xx status code
func (o *CreateInboundShipmentPlanTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create inbound shipment plan too many requests response has a 3xx status code
func (o *CreateInboundShipmentPlanTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create inbound shipment plan too many requests response has a 4xx status code
func (o *CreateInboundShipmentPlanTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this create inbound shipment plan too many requests response has a 5xx status code
func (o *CreateInboundShipmentPlanTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this create inbound shipment plan too many requests response a status code equal to that given
func (o *CreateInboundShipmentPlanTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *CreateInboundShipmentPlanTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /fba/inbound/v0/plans][%d] createInboundShipmentPlanTooManyRequests  %+v", 429, o.Payload)
}

func (o *CreateInboundShipmentPlanTooManyRequests) String() string {
	return fmt.Sprintf("[POST /fba/inbound/v0/plans][%d] createInboundShipmentPlanTooManyRequests  %+v", 429, o.Payload)
}

func (o *CreateInboundShipmentPlanTooManyRequests) GetPayload() *fulfillment_inbound_v0_models.CreateInboundShipmentPlanResponse {
	return o.Payload
}

func (o *CreateInboundShipmentPlanTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(fulfillment_inbound_v0_models.CreateInboundShipmentPlanResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateInboundShipmentPlanInternalServerError creates a CreateInboundShipmentPlanInternalServerError with default headers values
func NewCreateInboundShipmentPlanInternalServerError() *CreateInboundShipmentPlanInternalServerError {
	return &CreateInboundShipmentPlanInternalServerError{}
}

/*
CreateInboundShipmentPlanInternalServerError describes a response with status code 500, with default header values.

An unexpected condition occurred that prevented the server from fulfilling the request.
*/
type CreateInboundShipmentPlanInternalServerError struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inbound_v0_models.CreateInboundShipmentPlanResponse
}

// IsSuccess returns true when this create inbound shipment plan internal server error response has a 2xx status code
func (o *CreateInboundShipmentPlanInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create inbound shipment plan internal server error response has a 3xx status code
func (o *CreateInboundShipmentPlanInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create inbound shipment plan internal server error response has a 4xx status code
func (o *CreateInboundShipmentPlanInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create inbound shipment plan internal server error response has a 5xx status code
func (o *CreateInboundShipmentPlanInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create inbound shipment plan internal server error response a status code equal to that given
func (o *CreateInboundShipmentPlanInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *CreateInboundShipmentPlanInternalServerError) Error() string {
	return fmt.Sprintf("[POST /fba/inbound/v0/plans][%d] createInboundShipmentPlanInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateInboundShipmentPlanInternalServerError) String() string {
	return fmt.Sprintf("[POST /fba/inbound/v0/plans][%d] createInboundShipmentPlanInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateInboundShipmentPlanInternalServerError) GetPayload() *fulfillment_inbound_v0_models.CreateInboundShipmentPlanResponse {
	return o.Payload
}

func (o *CreateInboundShipmentPlanInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(fulfillment_inbound_v0_models.CreateInboundShipmentPlanResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateInboundShipmentPlanServiceUnavailable creates a CreateInboundShipmentPlanServiceUnavailable with default headers values
func NewCreateInboundShipmentPlanServiceUnavailable() *CreateInboundShipmentPlanServiceUnavailable {
	return &CreateInboundShipmentPlanServiceUnavailable{}
}

/*
CreateInboundShipmentPlanServiceUnavailable describes a response with status code 503, with default header values.

Temporary overloading or maintenance of the server.
*/
type CreateInboundShipmentPlanServiceUnavailable struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inbound_v0_models.CreateInboundShipmentPlanResponse
}

// IsSuccess returns true when this create inbound shipment plan service unavailable response has a 2xx status code
func (o *CreateInboundShipmentPlanServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create inbound shipment plan service unavailable response has a 3xx status code
func (o *CreateInboundShipmentPlanServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create inbound shipment plan service unavailable response has a 4xx status code
func (o *CreateInboundShipmentPlanServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this create inbound shipment plan service unavailable response has a 5xx status code
func (o *CreateInboundShipmentPlanServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this create inbound shipment plan service unavailable response a status code equal to that given
func (o *CreateInboundShipmentPlanServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *CreateInboundShipmentPlanServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /fba/inbound/v0/plans][%d] createInboundShipmentPlanServiceUnavailable  %+v", 503, o.Payload)
}

func (o *CreateInboundShipmentPlanServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /fba/inbound/v0/plans][%d] createInboundShipmentPlanServiceUnavailable  %+v", 503, o.Payload)
}

func (o *CreateInboundShipmentPlanServiceUnavailable) GetPayload() *fulfillment_inbound_v0_models.CreateInboundShipmentPlanResponse {
	return o.Payload
}

func (o *CreateInboundShipmentPlanServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(fulfillment_inbound_v0_models.CreateInboundShipmentPlanResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
