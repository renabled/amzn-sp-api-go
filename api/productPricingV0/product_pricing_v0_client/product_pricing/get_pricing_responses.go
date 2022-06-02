// Code generated by go-swagger; DO NOT EDIT.

package product_pricing

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/xamandar/amzn-sp-api-go/api/productPricingV0/product_pricing_v0_models"
)

// GetPricingReader is a Reader for the GetPricing structure.
type GetPricingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPricingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPricingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetPricingBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetPricingUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetPricingForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetPricingNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetPricingTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetPricingInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetPricingServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetPricingOK creates a GetPricingOK with default headers values
func NewGetPricingOK() *GetPricingOK {
	return &GetPricingOK{}
}

/* GetPricingOK describes a response with status code 200, with default header values.

Success.
*/
type GetPricingOK struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference ID.
	 */
	XAmznRequestID string

	Payload *product_pricing_v0_models.GetPricingResponse
}

func (o *GetPricingOK) Error() string {
	return fmt.Sprintf("[GET /products/pricing/v0/price][%d] getPricingOK  %+v", 200, o.Payload)
}
func (o *GetPricingOK) GetPayload() *product_pricing_v0_models.GetPricingResponse {
	return o.Payload
}

func (o *GetPricingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(product_pricing_v0_models.GetPricingResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPricingBadRequest creates a GetPricingBadRequest with default headers values
func NewGetPricingBadRequest() *GetPricingBadRequest {
	return &GetPricingBadRequest{}
}

/* GetPricingBadRequest describes a response with status code 400, with default header values.

Request has missing or invalid parameters and cannot be parsed.
*/
type GetPricingBadRequest struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference ID.
	 */
	XAmznRequestID string

	Payload *product_pricing_v0_models.GetPricingResponse
}

func (o *GetPricingBadRequest) Error() string {
	return fmt.Sprintf("[GET /products/pricing/v0/price][%d] getPricingBadRequest  %+v", 400, o.Payload)
}
func (o *GetPricingBadRequest) GetPayload() *product_pricing_v0_models.GetPricingResponse {
	return o.Payload
}

func (o *GetPricingBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(product_pricing_v0_models.GetPricingResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPricingUnauthorized creates a GetPricingUnauthorized with default headers values
func NewGetPricingUnauthorized() *GetPricingUnauthorized {
	return &GetPricingUnauthorized{}
}

/* GetPricingUnauthorized describes a response with status code 401, with default header values.

The request's Authorization header is not formatted correctly or does not contain a valid token.
*/
type GetPricingUnauthorized struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference ID.
	 */
	XAmznRequestID string

	Payload *product_pricing_v0_models.GetPricingResponse
}

func (o *GetPricingUnauthorized) Error() string {
	return fmt.Sprintf("[GET /products/pricing/v0/price][%d] getPricingUnauthorized  %+v", 401, o.Payload)
}
func (o *GetPricingUnauthorized) GetPayload() *product_pricing_v0_models.GetPricingResponse {
	return o.Payload
}

func (o *GetPricingUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(product_pricing_v0_models.GetPricingResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPricingForbidden creates a GetPricingForbidden with default headers values
func NewGetPricingForbidden() *GetPricingForbidden {
	return &GetPricingForbidden{}
}

/* GetPricingForbidden describes a response with status code 403, with default header values.

Indicates access to the resource is forbidden. Possible reasons include Access Denied, Unauthorized, Expired Token, or Invalid Signature.
*/
type GetPricingForbidden struct {

	/* Unique request reference ID.
	 */
	XAmznRequestID string

	Payload *product_pricing_v0_models.GetPricingResponse
}

func (o *GetPricingForbidden) Error() string {
	return fmt.Sprintf("[GET /products/pricing/v0/price][%d] getPricingForbidden  %+v", 403, o.Payload)
}
func (o *GetPricingForbidden) GetPayload() *product_pricing_v0_models.GetPricingResponse {
	return o.Payload
}

func (o *GetPricingForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(product_pricing_v0_models.GetPricingResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPricingNotFound creates a GetPricingNotFound with default headers values
func NewGetPricingNotFound() *GetPricingNotFound {
	return &GetPricingNotFound{}
}

/* GetPricingNotFound describes a response with status code 404, with default header values.

The specified resource does not exist.
*/
type GetPricingNotFound struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference ID.
	 */
	XAmznRequestID string

	Payload *product_pricing_v0_models.GetPricingResponse
}

func (o *GetPricingNotFound) Error() string {
	return fmt.Sprintf("[GET /products/pricing/v0/price][%d] getPricingNotFound  %+v", 404, o.Payload)
}
func (o *GetPricingNotFound) GetPayload() *product_pricing_v0_models.GetPricingResponse {
	return o.Payload
}

func (o *GetPricingNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(product_pricing_v0_models.GetPricingResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPricingTooManyRequests creates a GetPricingTooManyRequests with default headers values
func NewGetPricingTooManyRequests() *GetPricingTooManyRequests {
	return &GetPricingTooManyRequests{}
}

/* GetPricingTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type GetPricingTooManyRequests struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference ID.
	 */
	XAmznRequestID string

	Payload *product_pricing_v0_models.GetPricingResponse
}

func (o *GetPricingTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /products/pricing/v0/price][%d] getPricingTooManyRequests  %+v", 429, o.Payload)
}
func (o *GetPricingTooManyRequests) GetPayload() *product_pricing_v0_models.GetPricingResponse {
	return o.Payload
}

func (o *GetPricingTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(product_pricing_v0_models.GetPricingResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPricingInternalServerError creates a GetPricingInternalServerError with default headers values
func NewGetPricingInternalServerError() *GetPricingInternalServerError {
	return &GetPricingInternalServerError{}
}

/* GetPricingInternalServerError describes a response with status code 500, with default header values.

An unexpected condition occurred that prevented the server from fulfilling the request.
*/
type GetPricingInternalServerError struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference ID.
	 */
	XAmznRequestID string

	Payload *product_pricing_v0_models.GetPricingResponse
}

func (o *GetPricingInternalServerError) Error() string {
	return fmt.Sprintf("[GET /products/pricing/v0/price][%d] getPricingInternalServerError  %+v", 500, o.Payload)
}
func (o *GetPricingInternalServerError) GetPayload() *product_pricing_v0_models.GetPricingResponse {
	return o.Payload
}

func (o *GetPricingInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(product_pricing_v0_models.GetPricingResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPricingServiceUnavailable creates a GetPricingServiceUnavailable with default headers values
func NewGetPricingServiceUnavailable() *GetPricingServiceUnavailable {
	return &GetPricingServiceUnavailable{}
}

/* GetPricingServiceUnavailable describes a response with status code 503, with default header values.

Temporary overloading or maintenance of the server.
*/
type GetPricingServiceUnavailable struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference ID.
	 */
	XAmznRequestID string

	Payload *product_pricing_v0_models.GetPricingResponse
}

func (o *GetPricingServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /products/pricing/v0/price][%d] getPricingServiceUnavailable  %+v", 503, o.Payload)
}
func (o *GetPricingServiceUnavailable) GetPayload() *product_pricing_v0_models.GetPricingResponse {
	return o.Payload
}

func (o *GetPricingServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(product_pricing_v0_models.GetPricingResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
