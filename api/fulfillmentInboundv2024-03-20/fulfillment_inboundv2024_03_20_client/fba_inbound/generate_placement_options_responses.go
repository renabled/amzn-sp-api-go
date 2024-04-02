// Code generated by go-swagger; DO NOT EDIT.

package fba_inbound

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/renabled/amzn-sp-api-go/api/fulfillmentInboundv2024-03-20/fulfillment_inboundv2024_03_20_models"
)

// GeneratePlacementOptionsReader is a Reader for the GeneratePlacementOptions structure.
type GeneratePlacementOptionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GeneratePlacementOptionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewGeneratePlacementOptionsAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGeneratePlacementOptionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGeneratePlacementOptionsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGeneratePlacementOptionsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGeneratePlacementOptionsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGeneratePlacementOptionsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGeneratePlacementOptionsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGeneratePlacementOptionsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGeneratePlacementOptionsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGeneratePlacementOptionsAccepted creates a GeneratePlacementOptionsAccepted with default headers values
func NewGeneratePlacementOptionsAccepted() *GeneratePlacementOptionsAccepted {
	return &GeneratePlacementOptionsAccepted{}
}

/*
GeneratePlacementOptionsAccepted describes a response with status code 202, with default header values.

GeneratePlacementOptions 202 response
*/
type GeneratePlacementOptionsAccepted struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inboundv2024_03_20_models.GeneratePlacementOptionsResponse
}

// IsSuccess returns true when this generate placement options accepted response has a 2xx status code
func (o *GeneratePlacementOptionsAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this generate placement options accepted response has a 3xx status code
func (o *GeneratePlacementOptionsAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this generate placement options accepted response has a 4xx status code
func (o *GeneratePlacementOptionsAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this generate placement options accepted response has a 5xx status code
func (o *GeneratePlacementOptionsAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this generate placement options accepted response a status code equal to that given
func (o *GeneratePlacementOptionsAccepted) IsCode(code int) bool {
	return code == 202
}

func (o *GeneratePlacementOptionsAccepted) Error() string {
	return fmt.Sprintf("[POST /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/placementOptions][%d] generatePlacementOptionsAccepted  %+v", 202, o.Payload)
}

func (o *GeneratePlacementOptionsAccepted) String() string {
	return fmt.Sprintf("[POST /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/placementOptions][%d] generatePlacementOptionsAccepted  %+v", 202, o.Payload)
}

func (o *GeneratePlacementOptionsAccepted) GetPayload() *fulfillment_inboundv2024_03_20_models.GeneratePlacementOptionsResponse {
	return o.Payload
}

func (o *GeneratePlacementOptionsAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fulfillment_inboundv2024_03_20_models.GeneratePlacementOptionsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGeneratePlacementOptionsBadRequest creates a GeneratePlacementOptionsBadRequest with default headers values
func NewGeneratePlacementOptionsBadRequest() *GeneratePlacementOptionsBadRequest {
	return &GeneratePlacementOptionsBadRequest{}
}

/*
GeneratePlacementOptionsBadRequest describes a response with status code 400, with default header values.

Request has missing or invalid parameters and cannot be parsed.
*/
type GeneratePlacementOptionsBadRequest struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inboundv2024_03_20_models.ErrorList
}

// IsSuccess returns true when this generate placement options bad request response has a 2xx status code
func (o *GeneratePlacementOptionsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this generate placement options bad request response has a 3xx status code
func (o *GeneratePlacementOptionsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this generate placement options bad request response has a 4xx status code
func (o *GeneratePlacementOptionsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this generate placement options bad request response has a 5xx status code
func (o *GeneratePlacementOptionsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this generate placement options bad request response a status code equal to that given
func (o *GeneratePlacementOptionsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GeneratePlacementOptionsBadRequest) Error() string {
	return fmt.Sprintf("[POST /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/placementOptions][%d] generatePlacementOptionsBadRequest  %+v", 400, o.Payload)
}

func (o *GeneratePlacementOptionsBadRequest) String() string {
	return fmt.Sprintf("[POST /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/placementOptions][%d] generatePlacementOptionsBadRequest  %+v", 400, o.Payload)
}

func (o *GeneratePlacementOptionsBadRequest) GetPayload() *fulfillment_inboundv2024_03_20_models.ErrorList {
	return o.Payload
}

func (o *GeneratePlacementOptionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fulfillment_inboundv2024_03_20_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGeneratePlacementOptionsForbidden creates a GeneratePlacementOptionsForbidden with default headers values
func NewGeneratePlacementOptionsForbidden() *GeneratePlacementOptionsForbidden {
	return &GeneratePlacementOptionsForbidden{}
}

/*
GeneratePlacementOptionsForbidden describes a response with status code 403, with default header values.

Indicates that access to the resource is forbidden. Possible reasons include Access Denied, Unauthorized, Expired Token, or Invalid Signature.
*/
type GeneratePlacementOptionsForbidden struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inboundv2024_03_20_models.ErrorList
}

// IsSuccess returns true when this generate placement options forbidden response has a 2xx status code
func (o *GeneratePlacementOptionsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this generate placement options forbidden response has a 3xx status code
func (o *GeneratePlacementOptionsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this generate placement options forbidden response has a 4xx status code
func (o *GeneratePlacementOptionsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this generate placement options forbidden response has a 5xx status code
func (o *GeneratePlacementOptionsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this generate placement options forbidden response a status code equal to that given
func (o *GeneratePlacementOptionsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GeneratePlacementOptionsForbidden) Error() string {
	return fmt.Sprintf("[POST /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/placementOptions][%d] generatePlacementOptionsForbidden  %+v", 403, o.Payload)
}

func (o *GeneratePlacementOptionsForbidden) String() string {
	return fmt.Sprintf("[POST /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/placementOptions][%d] generatePlacementOptionsForbidden  %+v", 403, o.Payload)
}

func (o *GeneratePlacementOptionsForbidden) GetPayload() *fulfillment_inboundv2024_03_20_models.ErrorList {
	return o.Payload
}

func (o *GeneratePlacementOptionsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(fulfillment_inboundv2024_03_20_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGeneratePlacementOptionsNotFound creates a GeneratePlacementOptionsNotFound with default headers values
func NewGeneratePlacementOptionsNotFound() *GeneratePlacementOptionsNotFound {
	return &GeneratePlacementOptionsNotFound{}
}

/*
GeneratePlacementOptionsNotFound describes a response with status code 404, with default header values.

The resource specified does not exist.
*/
type GeneratePlacementOptionsNotFound struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inboundv2024_03_20_models.ErrorList
}

// IsSuccess returns true when this generate placement options not found response has a 2xx status code
func (o *GeneratePlacementOptionsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this generate placement options not found response has a 3xx status code
func (o *GeneratePlacementOptionsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this generate placement options not found response has a 4xx status code
func (o *GeneratePlacementOptionsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this generate placement options not found response has a 5xx status code
func (o *GeneratePlacementOptionsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this generate placement options not found response a status code equal to that given
func (o *GeneratePlacementOptionsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GeneratePlacementOptionsNotFound) Error() string {
	return fmt.Sprintf("[POST /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/placementOptions][%d] generatePlacementOptionsNotFound  %+v", 404, o.Payload)
}

func (o *GeneratePlacementOptionsNotFound) String() string {
	return fmt.Sprintf("[POST /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/placementOptions][%d] generatePlacementOptionsNotFound  %+v", 404, o.Payload)
}

func (o *GeneratePlacementOptionsNotFound) GetPayload() *fulfillment_inboundv2024_03_20_models.ErrorList {
	return o.Payload
}

func (o *GeneratePlacementOptionsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fulfillment_inboundv2024_03_20_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGeneratePlacementOptionsRequestEntityTooLarge creates a GeneratePlacementOptionsRequestEntityTooLarge with default headers values
func NewGeneratePlacementOptionsRequestEntityTooLarge() *GeneratePlacementOptionsRequestEntityTooLarge {
	return &GeneratePlacementOptionsRequestEntityTooLarge{}
}

/*
GeneratePlacementOptionsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request size exceeded the maximum accepted size.
*/
type GeneratePlacementOptionsRequestEntityTooLarge struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inboundv2024_03_20_models.ErrorList
}

// IsSuccess returns true when this generate placement options request entity too large response has a 2xx status code
func (o *GeneratePlacementOptionsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this generate placement options request entity too large response has a 3xx status code
func (o *GeneratePlacementOptionsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this generate placement options request entity too large response has a 4xx status code
func (o *GeneratePlacementOptionsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this generate placement options request entity too large response has a 5xx status code
func (o *GeneratePlacementOptionsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this generate placement options request entity too large response a status code equal to that given
func (o *GeneratePlacementOptionsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GeneratePlacementOptionsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/placementOptions][%d] generatePlacementOptionsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GeneratePlacementOptionsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[POST /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/placementOptions][%d] generatePlacementOptionsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GeneratePlacementOptionsRequestEntityTooLarge) GetPayload() *fulfillment_inboundv2024_03_20_models.ErrorList {
	return o.Payload
}

func (o *GeneratePlacementOptionsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(fulfillment_inboundv2024_03_20_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGeneratePlacementOptionsUnsupportedMediaType creates a GeneratePlacementOptionsUnsupportedMediaType with default headers values
func NewGeneratePlacementOptionsUnsupportedMediaType() *GeneratePlacementOptionsUnsupportedMediaType {
	return &GeneratePlacementOptionsUnsupportedMediaType{}
}

/*
GeneratePlacementOptionsUnsupportedMediaType describes a response with status code 415, with default header values.

The request payload is in an unsupported format.
*/
type GeneratePlacementOptionsUnsupportedMediaType struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inboundv2024_03_20_models.ErrorList
}

// IsSuccess returns true when this generate placement options unsupported media type response has a 2xx status code
func (o *GeneratePlacementOptionsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this generate placement options unsupported media type response has a 3xx status code
func (o *GeneratePlacementOptionsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this generate placement options unsupported media type response has a 4xx status code
func (o *GeneratePlacementOptionsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this generate placement options unsupported media type response has a 5xx status code
func (o *GeneratePlacementOptionsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this generate placement options unsupported media type response a status code equal to that given
func (o *GeneratePlacementOptionsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GeneratePlacementOptionsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/placementOptions][%d] generatePlacementOptionsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GeneratePlacementOptionsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[POST /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/placementOptions][%d] generatePlacementOptionsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GeneratePlacementOptionsUnsupportedMediaType) GetPayload() *fulfillment_inboundv2024_03_20_models.ErrorList {
	return o.Payload
}

func (o *GeneratePlacementOptionsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(fulfillment_inboundv2024_03_20_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGeneratePlacementOptionsTooManyRequests creates a GeneratePlacementOptionsTooManyRequests with default headers values
func NewGeneratePlacementOptionsTooManyRequests() *GeneratePlacementOptionsTooManyRequests {
	return &GeneratePlacementOptionsTooManyRequests{}
}

/*
GeneratePlacementOptionsTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type GeneratePlacementOptionsTooManyRequests struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inboundv2024_03_20_models.ErrorList
}

// IsSuccess returns true when this generate placement options too many requests response has a 2xx status code
func (o *GeneratePlacementOptionsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this generate placement options too many requests response has a 3xx status code
func (o *GeneratePlacementOptionsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this generate placement options too many requests response has a 4xx status code
func (o *GeneratePlacementOptionsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this generate placement options too many requests response has a 5xx status code
func (o *GeneratePlacementOptionsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this generate placement options too many requests response a status code equal to that given
func (o *GeneratePlacementOptionsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GeneratePlacementOptionsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/placementOptions][%d] generatePlacementOptionsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GeneratePlacementOptionsTooManyRequests) String() string {
	return fmt.Sprintf("[POST /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/placementOptions][%d] generatePlacementOptionsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GeneratePlacementOptionsTooManyRequests) GetPayload() *fulfillment_inboundv2024_03_20_models.ErrorList {
	return o.Payload
}

func (o *GeneratePlacementOptionsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(fulfillment_inboundv2024_03_20_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGeneratePlacementOptionsInternalServerError creates a GeneratePlacementOptionsInternalServerError with default headers values
func NewGeneratePlacementOptionsInternalServerError() *GeneratePlacementOptionsInternalServerError {
	return &GeneratePlacementOptionsInternalServerError{}
}

/*
GeneratePlacementOptionsInternalServerError describes a response with status code 500, with default header values.

An unexpected condition occurred that prevented the server from fulfilling the request.
*/
type GeneratePlacementOptionsInternalServerError struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inboundv2024_03_20_models.ErrorList
}

// IsSuccess returns true when this generate placement options internal server error response has a 2xx status code
func (o *GeneratePlacementOptionsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this generate placement options internal server error response has a 3xx status code
func (o *GeneratePlacementOptionsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this generate placement options internal server error response has a 4xx status code
func (o *GeneratePlacementOptionsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this generate placement options internal server error response has a 5xx status code
func (o *GeneratePlacementOptionsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this generate placement options internal server error response a status code equal to that given
func (o *GeneratePlacementOptionsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GeneratePlacementOptionsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/placementOptions][%d] generatePlacementOptionsInternalServerError  %+v", 500, o.Payload)
}

func (o *GeneratePlacementOptionsInternalServerError) String() string {
	return fmt.Sprintf("[POST /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/placementOptions][%d] generatePlacementOptionsInternalServerError  %+v", 500, o.Payload)
}

func (o *GeneratePlacementOptionsInternalServerError) GetPayload() *fulfillment_inboundv2024_03_20_models.ErrorList {
	return o.Payload
}

func (o *GeneratePlacementOptionsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(fulfillment_inboundv2024_03_20_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGeneratePlacementOptionsServiceUnavailable creates a GeneratePlacementOptionsServiceUnavailable with default headers values
func NewGeneratePlacementOptionsServiceUnavailable() *GeneratePlacementOptionsServiceUnavailable {
	return &GeneratePlacementOptionsServiceUnavailable{}
}

/*
GeneratePlacementOptionsServiceUnavailable describes a response with status code 503, with default header values.

Temporary overloading or maintenance of the server.
*/
type GeneratePlacementOptionsServiceUnavailable struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inboundv2024_03_20_models.ErrorList
}

// IsSuccess returns true when this generate placement options service unavailable response has a 2xx status code
func (o *GeneratePlacementOptionsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this generate placement options service unavailable response has a 3xx status code
func (o *GeneratePlacementOptionsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this generate placement options service unavailable response has a 4xx status code
func (o *GeneratePlacementOptionsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this generate placement options service unavailable response has a 5xx status code
func (o *GeneratePlacementOptionsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this generate placement options service unavailable response a status code equal to that given
func (o *GeneratePlacementOptionsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GeneratePlacementOptionsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/placementOptions][%d] generatePlacementOptionsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GeneratePlacementOptionsServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/placementOptions][%d] generatePlacementOptionsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GeneratePlacementOptionsServiceUnavailable) GetPayload() *fulfillment_inboundv2024_03_20_models.ErrorList {
	return o.Payload
}

func (o *GeneratePlacementOptionsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(fulfillment_inboundv2024_03_20_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
