// Code generated by go-swagger; DO NOT EDIT.

package fba_outbound

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/renabled/amzn-sp-api-go/api/fulfillmentOutbound_2020-07-01/fulfillment_outbound_2020_07_01_models"
)

// GetPackageTrackingDetailsReader is a Reader for the GetPackageTrackingDetails structure.
type GetPackageTrackingDetailsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPackageTrackingDetailsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPackageTrackingDetailsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetPackageTrackingDetailsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetPackageTrackingDetailsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetPackageTrackingDetailsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetPackageTrackingDetailsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetPackageTrackingDetailsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetPackageTrackingDetailsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetPackageTrackingDetailsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetPackageTrackingDetailsOK creates a GetPackageTrackingDetailsOK with default headers values
func NewGetPackageTrackingDetailsOK() *GetPackageTrackingDetailsOK {
	return &GetPackageTrackingDetailsOK{}
}

/*
GetPackageTrackingDetailsOK describes a response with status code 200, with default header values.

Success.
*/
type GetPackageTrackingDetailsOK struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_outbound_2020_07_01_models.GetPackageTrackingDetailsResponse
}

// IsSuccess returns true when this get package tracking details o k response has a 2xx status code
func (o *GetPackageTrackingDetailsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get package tracking details o k response has a 3xx status code
func (o *GetPackageTrackingDetailsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get package tracking details o k response has a 4xx status code
func (o *GetPackageTrackingDetailsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get package tracking details o k response has a 5xx status code
func (o *GetPackageTrackingDetailsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get package tracking details o k response a status code equal to that given
func (o *GetPackageTrackingDetailsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetPackageTrackingDetailsOK) Error() string {
	return fmt.Sprintf("[GET /fba/outbound/2020-07-01/tracking][%d] getPackageTrackingDetailsOK  %+v", 200, o.Payload)
}

func (o *GetPackageTrackingDetailsOK) String() string {
	return fmt.Sprintf("[GET /fba/outbound/2020-07-01/tracking][%d] getPackageTrackingDetailsOK  %+v", 200, o.Payload)
}

func (o *GetPackageTrackingDetailsOK) GetPayload() *fulfillment_outbound_2020_07_01_models.GetPackageTrackingDetailsResponse {
	return o.Payload
}

func (o *GetPackageTrackingDetailsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fulfillment_outbound_2020_07_01_models.GetPackageTrackingDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPackageTrackingDetailsBadRequest creates a GetPackageTrackingDetailsBadRequest with default headers values
func NewGetPackageTrackingDetailsBadRequest() *GetPackageTrackingDetailsBadRequest {
	return &GetPackageTrackingDetailsBadRequest{}
}

/*
GetPackageTrackingDetailsBadRequest describes a response with status code 400, with default header values.

Request has missing or invalid parameters and cannot be parsed.
*/
type GetPackageTrackingDetailsBadRequest struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_outbound_2020_07_01_models.GetPackageTrackingDetailsResponse
}

// IsSuccess returns true when this get package tracking details bad request response has a 2xx status code
func (o *GetPackageTrackingDetailsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get package tracking details bad request response has a 3xx status code
func (o *GetPackageTrackingDetailsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get package tracking details bad request response has a 4xx status code
func (o *GetPackageTrackingDetailsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get package tracking details bad request response has a 5xx status code
func (o *GetPackageTrackingDetailsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get package tracking details bad request response a status code equal to that given
func (o *GetPackageTrackingDetailsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetPackageTrackingDetailsBadRequest) Error() string {
	return fmt.Sprintf("[GET /fba/outbound/2020-07-01/tracking][%d] getPackageTrackingDetailsBadRequest  %+v", 400, o.Payload)
}

func (o *GetPackageTrackingDetailsBadRequest) String() string {
	return fmt.Sprintf("[GET /fba/outbound/2020-07-01/tracking][%d] getPackageTrackingDetailsBadRequest  %+v", 400, o.Payload)
}

func (o *GetPackageTrackingDetailsBadRequest) GetPayload() *fulfillment_outbound_2020_07_01_models.GetPackageTrackingDetailsResponse {
	return o.Payload
}

func (o *GetPackageTrackingDetailsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fulfillment_outbound_2020_07_01_models.GetPackageTrackingDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPackageTrackingDetailsUnauthorized creates a GetPackageTrackingDetailsUnauthorized with default headers values
func NewGetPackageTrackingDetailsUnauthorized() *GetPackageTrackingDetailsUnauthorized {
	return &GetPackageTrackingDetailsUnauthorized{}
}

/*
GetPackageTrackingDetailsUnauthorized describes a response with status code 401, with default header values.

The request's Authorization header is not formatted correctly or does not contain a valid token.
*/
type GetPackageTrackingDetailsUnauthorized struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_outbound_2020_07_01_models.GetPackageTrackingDetailsResponse
}

// IsSuccess returns true when this get package tracking details unauthorized response has a 2xx status code
func (o *GetPackageTrackingDetailsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get package tracking details unauthorized response has a 3xx status code
func (o *GetPackageTrackingDetailsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get package tracking details unauthorized response has a 4xx status code
func (o *GetPackageTrackingDetailsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get package tracking details unauthorized response has a 5xx status code
func (o *GetPackageTrackingDetailsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get package tracking details unauthorized response a status code equal to that given
func (o *GetPackageTrackingDetailsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetPackageTrackingDetailsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /fba/outbound/2020-07-01/tracking][%d] getPackageTrackingDetailsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetPackageTrackingDetailsUnauthorized) String() string {
	return fmt.Sprintf("[GET /fba/outbound/2020-07-01/tracking][%d] getPackageTrackingDetailsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetPackageTrackingDetailsUnauthorized) GetPayload() *fulfillment_outbound_2020_07_01_models.GetPackageTrackingDetailsResponse {
	return o.Payload
}

func (o *GetPackageTrackingDetailsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(fulfillment_outbound_2020_07_01_models.GetPackageTrackingDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPackageTrackingDetailsForbidden creates a GetPackageTrackingDetailsForbidden with default headers values
func NewGetPackageTrackingDetailsForbidden() *GetPackageTrackingDetailsForbidden {
	return &GetPackageTrackingDetailsForbidden{}
}

/*
GetPackageTrackingDetailsForbidden describes a response with status code 403, with default header values.

Indicates that access to the resource is forbidden. Possible reasons include Access Denied, Unauthorized, Expired Token, or Invalid Signature.
*/
type GetPackageTrackingDetailsForbidden struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_outbound_2020_07_01_models.GetPackageTrackingDetailsResponse
}

// IsSuccess returns true when this get package tracking details forbidden response has a 2xx status code
func (o *GetPackageTrackingDetailsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get package tracking details forbidden response has a 3xx status code
func (o *GetPackageTrackingDetailsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get package tracking details forbidden response has a 4xx status code
func (o *GetPackageTrackingDetailsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get package tracking details forbidden response has a 5xx status code
func (o *GetPackageTrackingDetailsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get package tracking details forbidden response a status code equal to that given
func (o *GetPackageTrackingDetailsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetPackageTrackingDetailsForbidden) Error() string {
	return fmt.Sprintf("[GET /fba/outbound/2020-07-01/tracking][%d] getPackageTrackingDetailsForbidden  %+v", 403, o.Payload)
}

func (o *GetPackageTrackingDetailsForbidden) String() string {
	return fmt.Sprintf("[GET /fba/outbound/2020-07-01/tracking][%d] getPackageTrackingDetailsForbidden  %+v", 403, o.Payload)
}

func (o *GetPackageTrackingDetailsForbidden) GetPayload() *fulfillment_outbound_2020_07_01_models.GetPackageTrackingDetailsResponse {
	return o.Payload
}

func (o *GetPackageTrackingDetailsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(fulfillment_outbound_2020_07_01_models.GetPackageTrackingDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPackageTrackingDetailsNotFound creates a GetPackageTrackingDetailsNotFound with default headers values
func NewGetPackageTrackingDetailsNotFound() *GetPackageTrackingDetailsNotFound {
	return &GetPackageTrackingDetailsNotFound{}
}

/*
GetPackageTrackingDetailsNotFound describes a response with status code 404, with default header values.

The specified resource does not exist.
*/
type GetPackageTrackingDetailsNotFound struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_outbound_2020_07_01_models.GetPackageTrackingDetailsResponse
}

// IsSuccess returns true when this get package tracking details not found response has a 2xx status code
func (o *GetPackageTrackingDetailsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get package tracking details not found response has a 3xx status code
func (o *GetPackageTrackingDetailsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get package tracking details not found response has a 4xx status code
func (o *GetPackageTrackingDetailsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get package tracking details not found response has a 5xx status code
func (o *GetPackageTrackingDetailsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get package tracking details not found response a status code equal to that given
func (o *GetPackageTrackingDetailsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetPackageTrackingDetailsNotFound) Error() string {
	return fmt.Sprintf("[GET /fba/outbound/2020-07-01/tracking][%d] getPackageTrackingDetailsNotFound  %+v", 404, o.Payload)
}

func (o *GetPackageTrackingDetailsNotFound) String() string {
	return fmt.Sprintf("[GET /fba/outbound/2020-07-01/tracking][%d] getPackageTrackingDetailsNotFound  %+v", 404, o.Payload)
}

func (o *GetPackageTrackingDetailsNotFound) GetPayload() *fulfillment_outbound_2020_07_01_models.GetPackageTrackingDetailsResponse {
	return o.Payload
}

func (o *GetPackageTrackingDetailsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fulfillment_outbound_2020_07_01_models.GetPackageTrackingDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPackageTrackingDetailsTooManyRequests creates a GetPackageTrackingDetailsTooManyRequests with default headers values
func NewGetPackageTrackingDetailsTooManyRequests() *GetPackageTrackingDetailsTooManyRequests {
	return &GetPackageTrackingDetailsTooManyRequests{}
}

/*
GetPackageTrackingDetailsTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type GetPackageTrackingDetailsTooManyRequests struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_outbound_2020_07_01_models.GetPackageTrackingDetailsResponse
}

// IsSuccess returns true when this get package tracking details too many requests response has a 2xx status code
func (o *GetPackageTrackingDetailsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get package tracking details too many requests response has a 3xx status code
func (o *GetPackageTrackingDetailsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get package tracking details too many requests response has a 4xx status code
func (o *GetPackageTrackingDetailsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get package tracking details too many requests response has a 5xx status code
func (o *GetPackageTrackingDetailsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get package tracking details too many requests response a status code equal to that given
func (o *GetPackageTrackingDetailsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetPackageTrackingDetailsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /fba/outbound/2020-07-01/tracking][%d] getPackageTrackingDetailsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetPackageTrackingDetailsTooManyRequests) String() string {
	return fmt.Sprintf("[GET /fba/outbound/2020-07-01/tracking][%d] getPackageTrackingDetailsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetPackageTrackingDetailsTooManyRequests) GetPayload() *fulfillment_outbound_2020_07_01_models.GetPackageTrackingDetailsResponse {
	return o.Payload
}

func (o *GetPackageTrackingDetailsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(fulfillment_outbound_2020_07_01_models.GetPackageTrackingDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPackageTrackingDetailsInternalServerError creates a GetPackageTrackingDetailsInternalServerError with default headers values
func NewGetPackageTrackingDetailsInternalServerError() *GetPackageTrackingDetailsInternalServerError {
	return &GetPackageTrackingDetailsInternalServerError{}
}

/*
GetPackageTrackingDetailsInternalServerError describes a response with status code 500, with default header values.

An unexpected condition occurred that prevented the server from fulfilling the request.
*/
type GetPackageTrackingDetailsInternalServerError struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_outbound_2020_07_01_models.GetPackageTrackingDetailsResponse
}

// IsSuccess returns true when this get package tracking details internal server error response has a 2xx status code
func (o *GetPackageTrackingDetailsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get package tracking details internal server error response has a 3xx status code
func (o *GetPackageTrackingDetailsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get package tracking details internal server error response has a 4xx status code
func (o *GetPackageTrackingDetailsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get package tracking details internal server error response has a 5xx status code
func (o *GetPackageTrackingDetailsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get package tracking details internal server error response a status code equal to that given
func (o *GetPackageTrackingDetailsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetPackageTrackingDetailsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /fba/outbound/2020-07-01/tracking][%d] getPackageTrackingDetailsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetPackageTrackingDetailsInternalServerError) String() string {
	return fmt.Sprintf("[GET /fba/outbound/2020-07-01/tracking][%d] getPackageTrackingDetailsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetPackageTrackingDetailsInternalServerError) GetPayload() *fulfillment_outbound_2020_07_01_models.GetPackageTrackingDetailsResponse {
	return o.Payload
}

func (o *GetPackageTrackingDetailsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(fulfillment_outbound_2020_07_01_models.GetPackageTrackingDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPackageTrackingDetailsServiceUnavailable creates a GetPackageTrackingDetailsServiceUnavailable with default headers values
func NewGetPackageTrackingDetailsServiceUnavailable() *GetPackageTrackingDetailsServiceUnavailable {
	return &GetPackageTrackingDetailsServiceUnavailable{}
}

/*
GetPackageTrackingDetailsServiceUnavailable describes a response with status code 503, with default header values.

Temporary overloading or maintenance of the server.
*/
type GetPackageTrackingDetailsServiceUnavailable struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_outbound_2020_07_01_models.GetPackageTrackingDetailsResponse
}

// IsSuccess returns true when this get package tracking details service unavailable response has a 2xx status code
func (o *GetPackageTrackingDetailsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get package tracking details service unavailable response has a 3xx status code
func (o *GetPackageTrackingDetailsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get package tracking details service unavailable response has a 4xx status code
func (o *GetPackageTrackingDetailsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get package tracking details service unavailable response has a 5xx status code
func (o *GetPackageTrackingDetailsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get package tracking details service unavailable response a status code equal to that given
func (o *GetPackageTrackingDetailsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetPackageTrackingDetailsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /fba/outbound/2020-07-01/tracking][%d] getPackageTrackingDetailsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetPackageTrackingDetailsServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /fba/outbound/2020-07-01/tracking][%d] getPackageTrackingDetailsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetPackageTrackingDetailsServiceUnavailable) GetPayload() *fulfillment_outbound_2020_07_01_models.GetPackageTrackingDetailsResponse {
	return o.Payload
}

func (o *GetPackageTrackingDetailsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(fulfillment_outbound_2020_07_01_models.GetPackageTrackingDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
