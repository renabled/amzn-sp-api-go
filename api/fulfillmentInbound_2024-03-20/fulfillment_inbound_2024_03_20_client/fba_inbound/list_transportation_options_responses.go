// Code generated by go-swagger; DO NOT EDIT.

package fba_inbound

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/renabled/amzn-sp-api-go/api/fulfillmentInbound_2024-03-20/fulfillment_inbound_2024_03_20_models"
)

// ListTransportationOptionsReader is a Reader for the ListTransportationOptions structure.
type ListTransportationOptionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListTransportationOptionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListTransportationOptionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListTransportationOptionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListTransportationOptionsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListTransportationOptionsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewListTransportationOptionsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewListTransportationOptionsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewListTransportationOptionsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListTransportationOptionsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewListTransportationOptionsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListTransportationOptionsOK creates a ListTransportationOptionsOK with default headers values
func NewListTransportationOptionsOK() *ListTransportationOptionsOK {
	return &ListTransportationOptionsOK{}
}

/*
ListTransportationOptionsOK describes a response with status code 200, with default header values.

ListTransportationOptions 200 response
*/
type ListTransportationOptionsOK struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inbound_2024_03_20_models.ListTransportationOptionsResponse
}

// IsSuccess returns true when this list transportation options o k response has a 2xx status code
func (o *ListTransportationOptionsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list transportation options o k response has a 3xx status code
func (o *ListTransportationOptionsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list transportation options o k response has a 4xx status code
func (o *ListTransportationOptionsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list transportation options o k response has a 5xx status code
func (o *ListTransportationOptionsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list transportation options o k response a status code equal to that given
func (o *ListTransportationOptionsOK) IsCode(code int) bool {
	return code == 200
}

func (o *ListTransportationOptionsOK) Error() string {
	return fmt.Sprintf("[GET /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/transportationOptions][%d] listTransportationOptionsOK  %+v", 200, o.Payload)
}

func (o *ListTransportationOptionsOK) String() string {
	return fmt.Sprintf("[GET /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/transportationOptions][%d] listTransportationOptionsOK  %+v", 200, o.Payload)
}

func (o *ListTransportationOptionsOK) GetPayload() *fulfillment_inbound_2024_03_20_models.ListTransportationOptionsResponse {
	return o.Payload
}

func (o *ListTransportationOptionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fulfillment_inbound_2024_03_20_models.ListTransportationOptionsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListTransportationOptionsBadRequest creates a ListTransportationOptionsBadRequest with default headers values
func NewListTransportationOptionsBadRequest() *ListTransportationOptionsBadRequest {
	return &ListTransportationOptionsBadRequest{}
}

/*
ListTransportationOptionsBadRequest describes a response with status code 400, with default header values.

Request has missing or invalid parameters and cannot be parsed.
*/
type ListTransportationOptionsBadRequest struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inbound_2024_03_20_models.ErrorList
}

// IsSuccess returns true when this list transportation options bad request response has a 2xx status code
func (o *ListTransportationOptionsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list transportation options bad request response has a 3xx status code
func (o *ListTransportationOptionsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list transportation options bad request response has a 4xx status code
func (o *ListTransportationOptionsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this list transportation options bad request response has a 5xx status code
func (o *ListTransportationOptionsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this list transportation options bad request response a status code equal to that given
func (o *ListTransportationOptionsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *ListTransportationOptionsBadRequest) Error() string {
	return fmt.Sprintf("[GET /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/transportationOptions][%d] listTransportationOptionsBadRequest  %+v", 400, o.Payload)
}

func (o *ListTransportationOptionsBadRequest) String() string {
	return fmt.Sprintf("[GET /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/transportationOptions][%d] listTransportationOptionsBadRequest  %+v", 400, o.Payload)
}

func (o *ListTransportationOptionsBadRequest) GetPayload() *fulfillment_inbound_2024_03_20_models.ErrorList {
	return o.Payload
}

func (o *ListTransportationOptionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fulfillment_inbound_2024_03_20_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListTransportationOptionsForbidden creates a ListTransportationOptionsForbidden with default headers values
func NewListTransportationOptionsForbidden() *ListTransportationOptionsForbidden {
	return &ListTransportationOptionsForbidden{}
}

/*
ListTransportationOptionsForbidden describes a response with status code 403, with default header values.

Indicates that access to the resource is forbidden. Possible reasons include Access Denied, Unauthorized, Expired Token, or Invalid Signature.
*/
type ListTransportationOptionsForbidden struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inbound_2024_03_20_models.ErrorList
}

// IsSuccess returns true when this list transportation options forbidden response has a 2xx status code
func (o *ListTransportationOptionsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list transportation options forbidden response has a 3xx status code
func (o *ListTransportationOptionsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list transportation options forbidden response has a 4xx status code
func (o *ListTransportationOptionsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this list transportation options forbidden response has a 5xx status code
func (o *ListTransportationOptionsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this list transportation options forbidden response a status code equal to that given
func (o *ListTransportationOptionsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ListTransportationOptionsForbidden) Error() string {
	return fmt.Sprintf("[GET /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/transportationOptions][%d] listTransportationOptionsForbidden  %+v", 403, o.Payload)
}

func (o *ListTransportationOptionsForbidden) String() string {
	return fmt.Sprintf("[GET /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/transportationOptions][%d] listTransportationOptionsForbidden  %+v", 403, o.Payload)
}

func (o *ListTransportationOptionsForbidden) GetPayload() *fulfillment_inbound_2024_03_20_models.ErrorList {
	return o.Payload
}

func (o *ListTransportationOptionsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(fulfillment_inbound_2024_03_20_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListTransportationOptionsNotFound creates a ListTransportationOptionsNotFound with default headers values
func NewListTransportationOptionsNotFound() *ListTransportationOptionsNotFound {
	return &ListTransportationOptionsNotFound{}
}

/*
ListTransportationOptionsNotFound describes a response with status code 404, with default header values.

The resource specified does not exist.
*/
type ListTransportationOptionsNotFound struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inbound_2024_03_20_models.ErrorList
}

// IsSuccess returns true when this list transportation options not found response has a 2xx status code
func (o *ListTransportationOptionsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list transportation options not found response has a 3xx status code
func (o *ListTransportationOptionsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list transportation options not found response has a 4xx status code
func (o *ListTransportationOptionsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this list transportation options not found response has a 5xx status code
func (o *ListTransportationOptionsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this list transportation options not found response a status code equal to that given
func (o *ListTransportationOptionsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *ListTransportationOptionsNotFound) Error() string {
	return fmt.Sprintf("[GET /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/transportationOptions][%d] listTransportationOptionsNotFound  %+v", 404, o.Payload)
}

func (o *ListTransportationOptionsNotFound) String() string {
	return fmt.Sprintf("[GET /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/transportationOptions][%d] listTransportationOptionsNotFound  %+v", 404, o.Payload)
}

func (o *ListTransportationOptionsNotFound) GetPayload() *fulfillment_inbound_2024_03_20_models.ErrorList {
	return o.Payload
}

func (o *ListTransportationOptionsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fulfillment_inbound_2024_03_20_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListTransportationOptionsRequestEntityTooLarge creates a ListTransportationOptionsRequestEntityTooLarge with default headers values
func NewListTransportationOptionsRequestEntityTooLarge() *ListTransportationOptionsRequestEntityTooLarge {
	return &ListTransportationOptionsRequestEntityTooLarge{}
}

/*
ListTransportationOptionsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request size exceeded the maximum accepted size.
*/
type ListTransportationOptionsRequestEntityTooLarge struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inbound_2024_03_20_models.ErrorList
}

// IsSuccess returns true when this list transportation options request entity too large response has a 2xx status code
func (o *ListTransportationOptionsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list transportation options request entity too large response has a 3xx status code
func (o *ListTransportationOptionsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list transportation options request entity too large response has a 4xx status code
func (o *ListTransportationOptionsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this list transportation options request entity too large response has a 5xx status code
func (o *ListTransportationOptionsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this list transportation options request entity too large response a status code equal to that given
func (o *ListTransportationOptionsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *ListTransportationOptionsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/transportationOptions][%d] listTransportationOptionsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *ListTransportationOptionsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/transportationOptions][%d] listTransportationOptionsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *ListTransportationOptionsRequestEntityTooLarge) GetPayload() *fulfillment_inbound_2024_03_20_models.ErrorList {
	return o.Payload
}

func (o *ListTransportationOptionsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(fulfillment_inbound_2024_03_20_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListTransportationOptionsUnsupportedMediaType creates a ListTransportationOptionsUnsupportedMediaType with default headers values
func NewListTransportationOptionsUnsupportedMediaType() *ListTransportationOptionsUnsupportedMediaType {
	return &ListTransportationOptionsUnsupportedMediaType{}
}

/*
ListTransportationOptionsUnsupportedMediaType describes a response with status code 415, with default header values.

The request payload is in an unsupported format.
*/
type ListTransportationOptionsUnsupportedMediaType struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inbound_2024_03_20_models.ErrorList
}

// IsSuccess returns true when this list transportation options unsupported media type response has a 2xx status code
func (o *ListTransportationOptionsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list transportation options unsupported media type response has a 3xx status code
func (o *ListTransportationOptionsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list transportation options unsupported media type response has a 4xx status code
func (o *ListTransportationOptionsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this list transportation options unsupported media type response has a 5xx status code
func (o *ListTransportationOptionsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this list transportation options unsupported media type response a status code equal to that given
func (o *ListTransportationOptionsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *ListTransportationOptionsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/transportationOptions][%d] listTransportationOptionsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *ListTransportationOptionsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/transportationOptions][%d] listTransportationOptionsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *ListTransportationOptionsUnsupportedMediaType) GetPayload() *fulfillment_inbound_2024_03_20_models.ErrorList {
	return o.Payload
}

func (o *ListTransportationOptionsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(fulfillment_inbound_2024_03_20_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListTransportationOptionsTooManyRequests creates a ListTransportationOptionsTooManyRequests with default headers values
func NewListTransportationOptionsTooManyRequests() *ListTransportationOptionsTooManyRequests {
	return &ListTransportationOptionsTooManyRequests{}
}

/*
ListTransportationOptionsTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type ListTransportationOptionsTooManyRequests struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inbound_2024_03_20_models.ErrorList
}

// IsSuccess returns true when this list transportation options too many requests response has a 2xx status code
func (o *ListTransportationOptionsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list transportation options too many requests response has a 3xx status code
func (o *ListTransportationOptionsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list transportation options too many requests response has a 4xx status code
func (o *ListTransportationOptionsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this list transportation options too many requests response has a 5xx status code
func (o *ListTransportationOptionsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this list transportation options too many requests response a status code equal to that given
func (o *ListTransportationOptionsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *ListTransportationOptionsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/transportationOptions][%d] listTransportationOptionsTooManyRequests  %+v", 429, o.Payload)
}

func (o *ListTransportationOptionsTooManyRequests) String() string {
	return fmt.Sprintf("[GET /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/transportationOptions][%d] listTransportationOptionsTooManyRequests  %+v", 429, o.Payload)
}

func (o *ListTransportationOptionsTooManyRequests) GetPayload() *fulfillment_inbound_2024_03_20_models.ErrorList {
	return o.Payload
}

func (o *ListTransportationOptionsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(fulfillment_inbound_2024_03_20_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListTransportationOptionsInternalServerError creates a ListTransportationOptionsInternalServerError with default headers values
func NewListTransportationOptionsInternalServerError() *ListTransportationOptionsInternalServerError {
	return &ListTransportationOptionsInternalServerError{}
}

/*
ListTransportationOptionsInternalServerError describes a response with status code 500, with default header values.

An unexpected condition occurred that prevented the server from fulfilling the request.
*/
type ListTransportationOptionsInternalServerError struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inbound_2024_03_20_models.ErrorList
}

// IsSuccess returns true when this list transportation options internal server error response has a 2xx status code
func (o *ListTransportationOptionsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list transportation options internal server error response has a 3xx status code
func (o *ListTransportationOptionsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list transportation options internal server error response has a 4xx status code
func (o *ListTransportationOptionsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this list transportation options internal server error response has a 5xx status code
func (o *ListTransportationOptionsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this list transportation options internal server error response a status code equal to that given
func (o *ListTransportationOptionsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ListTransportationOptionsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/transportationOptions][%d] listTransportationOptionsInternalServerError  %+v", 500, o.Payload)
}

func (o *ListTransportationOptionsInternalServerError) String() string {
	return fmt.Sprintf("[GET /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/transportationOptions][%d] listTransportationOptionsInternalServerError  %+v", 500, o.Payload)
}

func (o *ListTransportationOptionsInternalServerError) GetPayload() *fulfillment_inbound_2024_03_20_models.ErrorList {
	return o.Payload
}

func (o *ListTransportationOptionsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(fulfillment_inbound_2024_03_20_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListTransportationOptionsServiceUnavailable creates a ListTransportationOptionsServiceUnavailable with default headers values
func NewListTransportationOptionsServiceUnavailable() *ListTransportationOptionsServiceUnavailable {
	return &ListTransportationOptionsServiceUnavailable{}
}

/*
ListTransportationOptionsServiceUnavailable describes a response with status code 503, with default header values.

Temporary overloading or maintenance of the server.
*/
type ListTransportationOptionsServiceUnavailable struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inbound_2024_03_20_models.ErrorList
}

// IsSuccess returns true when this list transportation options service unavailable response has a 2xx status code
func (o *ListTransportationOptionsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list transportation options service unavailable response has a 3xx status code
func (o *ListTransportationOptionsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list transportation options service unavailable response has a 4xx status code
func (o *ListTransportationOptionsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this list transportation options service unavailable response has a 5xx status code
func (o *ListTransportationOptionsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this list transportation options service unavailable response a status code equal to that given
func (o *ListTransportationOptionsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *ListTransportationOptionsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/transportationOptions][%d] listTransportationOptionsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *ListTransportationOptionsServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/transportationOptions][%d] listTransportationOptionsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *ListTransportationOptionsServiceUnavailable) GetPayload() *fulfillment_inbound_2024_03_20_models.ErrorList {
	return o.Payload
}

func (o *ListTransportationOptionsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(fulfillment_inbound_2024_03_20_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
