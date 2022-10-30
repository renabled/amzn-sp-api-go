// Code generated by go-swagger; DO NOT EDIT.

package fees

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/xamandar/amzn-sp-api-go/api/productFeesV0/product_fees_v0_models"
)

// GetMyFeesEstimateForASINReader is a Reader for the GetMyFeesEstimateForASIN structure.
type GetMyFeesEstimateForASINReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMyFeesEstimateForASINReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetMyFeesEstimateForASINOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetMyFeesEstimateForASINBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetMyFeesEstimateForASINUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetMyFeesEstimateForASINForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetMyFeesEstimateForASINNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetMyFeesEstimateForASINTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetMyFeesEstimateForASINInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetMyFeesEstimateForASINServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetMyFeesEstimateForASINOK creates a GetMyFeesEstimateForASINOK with default headers values
func NewGetMyFeesEstimateForASINOK() *GetMyFeesEstimateForASINOK {
	return &GetMyFeesEstimateForASINOK{}
}

/*
GetMyFeesEstimateForASINOK describes a response with status code 200, with default header values.

Success.
*/
type GetMyFeesEstimateForASINOK struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *product_fees_v0_models.GetMyFeesEstimateResponse
}

// IsSuccess returns true when this get my fees estimate for a s i n o k response has a 2xx status code
func (o *GetMyFeesEstimateForASINOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get my fees estimate for a s i n o k response has a 3xx status code
func (o *GetMyFeesEstimateForASINOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get my fees estimate for a s i n o k response has a 4xx status code
func (o *GetMyFeesEstimateForASINOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get my fees estimate for a s i n o k response has a 5xx status code
func (o *GetMyFeesEstimateForASINOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get my fees estimate for a s i n o k response a status code equal to that given
func (o *GetMyFeesEstimateForASINOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetMyFeesEstimateForASINOK) Error() string {
	return fmt.Sprintf("[POST /products/fees/v0/items/{Asin}/feesEstimate][%d] getMyFeesEstimateForASINOK  %+v", 200, o.Payload)
}

func (o *GetMyFeesEstimateForASINOK) String() string {
	return fmt.Sprintf("[POST /products/fees/v0/items/{Asin}/feesEstimate][%d] getMyFeesEstimateForASINOK  %+v", 200, o.Payload)
}

func (o *GetMyFeesEstimateForASINOK) GetPayload() *product_fees_v0_models.GetMyFeesEstimateResponse {
	return o.Payload
}

func (o *GetMyFeesEstimateForASINOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(product_fees_v0_models.GetMyFeesEstimateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMyFeesEstimateForASINBadRequest creates a GetMyFeesEstimateForASINBadRequest with default headers values
func NewGetMyFeesEstimateForASINBadRequest() *GetMyFeesEstimateForASINBadRequest {
	return &GetMyFeesEstimateForASINBadRequest{}
}

/*
GetMyFeesEstimateForASINBadRequest describes a response with status code 400, with default header values.

Request has missing or invalid parameters and cannot be parsed.
*/
type GetMyFeesEstimateForASINBadRequest struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *product_fees_v0_models.GetMyFeesEstimateResponse
}

// IsSuccess returns true when this get my fees estimate for a s i n bad request response has a 2xx status code
func (o *GetMyFeesEstimateForASINBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get my fees estimate for a s i n bad request response has a 3xx status code
func (o *GetMyFeesEstimateForASINBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get my fees estimate for a s i n bad request response has a 4xx status code
func (o *GetMyFeesEstimateForASINBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get my fees estimate for a s i n bad request response has a 5xx status code
func (o *GetMyFeesEstimateForASINBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get my fees estimate for a s i n bad request response a status code equal to that given
func (o *GetMyFeesEstimateForASINBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetMyFeesEstimateForASINBadRequest) Error() string {
	return fmt.Sprintf("[POST /products/fees/v0/items/{Asin}/feesEstimate][%d] getMyFeesEstimateForASINBadRequest  %+v", 400, o.Payload)
}

func (o *GetMyFeesEstimateForASINBadRequest) String() string {
	return fmt.Sprintf("[POST /products/fees/v0/items/{Asin}/feesEstimate][%d] getMyFeesEstimateForASINBadRequest  %+v", 400, o.Payload)
}

func (o *GetMyFeesEstimateForASINBadRequest) GetPayload() *product_fees_v0_models.GetMyFeesEstimateResponse {
	return o.Payload
}

func (o *GetMyFeesEstimateForASINBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(product_fees_v0_models.GetMyFeesEstimateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMyFeesEstimateForASINUnauthorized creates a GetMyFeesEstimateForASINUnauthorized with default headers values
func NewGetMyFeesEstimateForASINUnauthorized() *GetMyFeesEstimateForASINUnauthorized {
	return &GetMyFeesEstimateForASINUnauthorized{}
}

/*
GetMyFeesEstimateForASINUnauthorized describes a response with status code 401, with default header values.

The request's Authorization header is not formatted correctly or does not contain a valid token.
*/
type GetMyFeesEstimateForASINUnauthorized struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *product_fees_v0_models.GetMyFeesEstimateResponse
}

// IsSuccess returns true when this get my fees estimate for a s i n unauthorized response has a 2xx status code
func (o *GetMyFeesEstimateForASINUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get my fees estimate for a s i n unauthorized response has a 3xx status code
func (o *GetMyFeesEstimateForASINUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get my fees estimate for a s i n unauthorized response has a 4xx status code
func (o *GetMyFeesEstimateForASINUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get my fees estimate for a s i n unauthorized response has a 5xx status code
func (o *GetMyFeesEstimateForASINUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get my fees estimate for a s i n unauthorized response a status code equal to that given
func (o *GetMyFeesEstimateForASINUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetMyFeesEstimateForASINUnauthorized) Error() string {
	return fmt.Sprintf("[POST /products/fees/v0/items/{Asin}/feesEstimate][%d] getMyFeesEstimateForASINUnauthorized  %+v", 401, o.Payload)
}

func (o *GetMyFeesEstimateForASINUnauthorized) String() string {
	return fmt.Sprintf("[POST /products/fees/v0/items/{Asin}/feesEstimate][%d] getMyFeesEstimateForASINUnauthorized  %+v", 401, o.Payload)
}

func (o *GetMyFeesEstimateForASINUnauthorized) GetPayload() *product_fees_v0_models.GetMyFeesEstimateResponse {
	return o.Payload
}

func (o *GetMyFeesEstimateForASINUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(product_fees_v0_models.GetMyFeesEstimateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMyFeesEstimateForASINForbidden creates a GetMyFeesEstimateForASINForbidden with default headers values
func NewGetMyFeesEstimateForASINForbidden() *GetMyFeesEstimateForASINForbidden {
	return &GetMyFeesEstimateForASINForbidden{}
}

/*
GetMyFeesEstimateForASINForbidden describes a response with status code 403, with default header values.

Indicates access to the resource is forbidden. Possible reasons include Access Denied, Unauthorized, Expired Token, or Invalid Signature.
*/
type GetMyFeesEstimateForASINForbidden struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *product_fees_v0_models.GetMyFeesEstimateResponse
}

// IsSuccess returns true when this get my fees estimate for a s i n forbidden response has a 2xx status code
func (o *GetMyFeesEstimateForASINForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get my fees estimate for a s i n forbidden response has a 3xx status code
func (o *GetMyFeesEstimateForASINForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get my fees estimate for a s i n forbidden response has a 4xx status code
func (o *GetMyFeesEstimateForASINForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get my fees estimate for a s i n forbidden response has a 5xx status code
func (o *GetMyFeesEstimateForASINForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get my fees estimate for a s i n forbidden response a status code equal to that given
func (o *GetMyFeesEstimateForASINForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetMyFeesEstimateForASINForbidden) Error() string {
	return fmt.Sprintf("[POST /products/fees/v0/items/{Asin}/feesEstimate][%d] getMyFeesEstimateForASINForbidden  %+v", 403, o.Payload)
}

func (o *GetMyFeesEstimateForASINForbidden) String() string {
	return fmt.Sprintf("[POST /products/fees/v0/items/{Asin}/feesEstimate][%d] getMyFeesEstimateForASINForbidden  %+v", 403, o.Payload)
}

func (o *GetMyFeesEstimateForASINForbidden) GetPayload() *product_fees_v0_models.GetMyFeesEstimateResponse {
	return o.Payload
}

func (o *GetMyFeesEstimateForASINForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(product_fees_v0_models.GetMyFeesEstimateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMyFeesEstimateForASINNotFound creates a GetMyFeesEstimateForASINNotFound with default headers values
func NewGetMyFeesEstimateForASINNotFound() *GetMyFeesEstimateForASINNotFound {
	return &GetMyFeesEstimateForASINNotFound{}
}

/*
GetMyFeesEstimateForASINNotFound describes a response with status code 404, with default header values.

The specified resource does not exist.
*/
type GetMyFeesEstimateForASINNotFound struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *product_fees_v0_models.GetMyFeesEstimateResponse
}

// IsSuccess returns true when this get my fees estimate for a s i n not found response has a 2xx status code
func (o *GetMyFeesEstimateForASINNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get my fees estimate for a s i n not found response has a 3xx status code
func (o *GetMyFeesEstimateForASINNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get my fees estimate for a s i n not found response has a 4xx status code
func (o *GetMyFeesEstimateForASINNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get my fees estimate for a s i n not found response has a 5xx status code
func (o *GetMyFeesEstimateForASINNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get my fees estimate for a s i n not found response a status code equal to that given
func (o *GetMyFeesEstimateForASINNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetMyFeesEstimateForASINNotFound) Error() string {
	return fmt.Sprintf("[POST /products/fees/v0/items/{Asin}/feesEstimate][%d] getMyFeesEstimateForASINNotFound  %+v", 404, o.Payload)
}

func (o *GetMyFeesEstimateForASINNotFound) String() string {
	return fmt.Sprintf("[POST /products/fees/v0/items/{Asin}/feesEstimate][%d] getMyFeesEstimateForASINNotFound  %+v", 404, o.Payload)
}

func (o *GetMyFeesEstimateForASINNotFound) GetPayload() *product_fees_v0_models.GetMyFeesEstimateResponse {
	return o.Payload
}

func (o *GetMyFeesEstimateForASINNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(product_fees_v0_models.GetMyFeesEstimateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMyFeesEstimateForASINTooManyRequests creates a GetMyFeesEstimateForASINTooManyRequests with default headers values
func NewGetMyFeesEstimateForASINTooManyRequests() *GetMyFeesEstimateForASINTooManyRequests {
	return &GetMyFeesEstimateForASINTooManyRequests{}
}

/*
GetMyFeesEstimateForASINTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type GetMyFeesEstimateForASINTooManyRequests struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *product_fees_v0_models.GetMyFeesEstimateResponse
}

// IsSuccess returns true when this get my fees estimate for a s i n too many requests response has a 2xx status code
func (o *GetMyFeesEstimateForASINTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get my fees estimate for a s i n too many requests response has a 3xx status code
func (o *GetMyFeesEstimateForASINTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get my fees estimate for a s i n too many requests response has a 4xx status code
func (o *GetMyFeesEstimateForASINTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get my fees estimate for a s i n too many requests response has a 5xx status code
func (o *GetMyFeesEstimateForASINTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get my fees estimate for a s i n too many requests response a status code equal to that given
func (o *GetMyFeesEstimateForASINTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetMyFeesEstimateForASINTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /products/fees/v0/items/{Asin}/feesEstimate][%d] getMyFeesEstimateForASINTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetMyFeesEstimateForASINTooManyRequests) String() string {
	return fmt.Sprintf("[POST /products/fees/v0/items/{Asin}/feesEstimate][%d] getMyFeesEstimateForASINTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetMyFeesEstimateForASINTooManyRequests) GetPayload() *product_fees_v0_models.GetMyFeesEstimateResponse {
	return o.Payload
}

func (o *GetMyFeesEstimateForASINTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(product_fees_v0_models.GetMyFeesEstimateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMyFeesEstimateForASINInternalServerError creates a GetMyFeesEstimateForASINInternalServerError with default headers values
func NewGetMyFeesEstimateForASINInternalServerError() *GetMyFeesEstimateForASINInternalServerError {
	return &GetMyFeesEstimateForASINInternalServerError{}
}

/*
GetMyFeesEstimateForASINInternalServerError describes a response with status code 500, with default header values.

An unexpected condition occurred that prevented the server from fulfilling the request.
*/
type GetMyFeesEstimateForASINInternalServerError struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *product_fees_v0_models.GetMyFeesEstimateResponse
}

// IsSuccess returns true when this get my fees estimate for a s i n internal server error response has a 2xx status code
func (o *GetMyFeesEstimateForASINInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get my fees estimate for a s i n internal server error response has a 3xx status code
func (o *GetMyFeesEstimateForASINInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get my fees estimate for a s i n internal server error response has a 4xx status code
func (o *GetMyFeesEstimateForASINInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get my fees estimate for a s i n internal server error response has a 5xx status code
func (o *GetMyFeesEstimateForASINInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get my fees estimate for a s i n internal server error response a status code equal to that given
func (o *GetMyFeesEstimateForASINInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetMyFeesEstimateForASINInternalServerError) Error() string {
	return fmt.Sprintf("[POST /products/fees/v0/items/{Asin}/feesEstimate][%d] getMyFeesEstimateForASINInternalServerError  %+v", 500, o.Payload)
}

func (o *GetMyFeesEstimateForASINInternalServerError) String() string {
	return fmt.Sprintf("[POST /products/fees/v0/items/{Asin}/feesEstimate][%d] getMyFeesEstimateForASINInternalServerError  %+v", 500, o.Payload)
}

func (o *GetMyFeesEstimateForASINInternalServerError) GetPayload() *product_fees_v0_models.GetMyFeesEstimateResponse {
	return o.Payload
}

func (o *GetMyFeesEstimateForASINInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(product_fees_v0_models.GetMyFeesEstimateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMyFeesEstimateForASINServiceUnavailable creates a GetMyFeesEstimateForASINServiceUnavailable with default headers values
func NewGetMyFeesEstimateForASINServiceUnavailable() *GetMyFeesEstimateForASINServiceUnavailable {
	return &GetMyFeesEstimateForASINServiceUnavailable{}
}

/*
GetMyFeesEstimateForASINServiceUnavailable describes a response with status code 503, with default header values.

Temporary overloading or maintenance of the server.
*/
type GetMyFeesEstimateForASINServiceUnavailable struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *product_fees_v0_models.GetMyFeesEstimateResponse
}

// IsSuccess returns true when this get my fees estimate for a s i n service unavailable response has a 2xx status code
func (o *GetMyFeesEstimateForASINServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get my fees estimate for a s i n service unavailable response has a 3xx status code
func (o *GetMyFeesEstimateForASINServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get my fees estimate for a s i n service unavailable response has a 4xx status code
func (o *GetMyFeesEstimateForASINServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get my fees estimate for a s i n service unavailable response has a 5xx status code
func (o *GetMyFeesEstimateForASINServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get my fees estimate for a s i n service unavailable response a status code equal to that given
func (o *GetMyFeesEstimateForASINServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetMyFeesEstimateForASINServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /products/fees/v0/items/{Asin}/feesEstimate][%d] getMyFeesEstimateForASINServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetMyFeesEstimateForASINServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /products/fees/v0/items/{Asin}/feesEstimate][%d] getMyFeesEstimateForASINServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetMyFeesEstimateForASINServiceUnavailable) GetPayload() *product_fees_v0_models.GetMyFeesEstimateResponse {
	return o.Payload
}

func (o *GetMyFeesEstimateForASINServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(product_fees_v0_models.GetMyFeesEstimateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
