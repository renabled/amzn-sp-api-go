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

// GetItemOffersReader is a Reader for the GetItemOffers structure.
type GetItemOffersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetItemOffersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetItemOffersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetItemOffersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetItemOffersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetItemOffersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetItemOffersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetItemOffersTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetItemOffersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetItemOffersServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetItemOffersOK creates a GetItemOffersOK with default headers values
func NewGetItemOffersOK() *GetItemOffersOK {
	return &GetItemOffersOK{}
}

/* GetItemOffersOK describes a response with status code 200, with default header values.

Success.
*/
type GetItemOffersOK struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference ID.
	 */
	XAmznRequestID string

	Payload *product_pricing_v0_models.GetOffersResponse
}

func (o *GetItemOffersOK) Error() string {
	return fmt.Sprintf("[GET /products/pricing/v0/items/{Asin}/offers][%d] getItemOffersOK  %+v", 200, o.Payload)
}
func (o *GetItemOffersOK) GetPayload() *product_pricing_v0_models.GetOffersResponse {
	return o.Payload
}

func (o *GetItemOffersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(product_pricing_v0_models.GetOffersResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetItemOffersBadRequest creates a GetItemOffersBadRequest with default headers values
func NewGetItemOffersBadRequest() *GetItemOffersBadRequest {
	return &GetItemOffersBadRequest{}
}

/* GetItemOffersBadRequest describes a response with status code 400, with default header values.

Request has missing or invalid parameters and cannot be parsed.
*/
type GetItemOffersBadRequest struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference ID.
	 */
	XAmznRequestID string

	Payload *product_pricing_v0_models.GetOffersResponse
}

func (o *GetItemOffersBadRequest) Error() string {
	return fmt.Sprintf("[GET /products/pricing/v0/items/{Asin}/offers][%d] getItemOffersBadRequest  %+v", 400, o.Payload)
}
func (o *GetItemOffersBadRequest) GetPayload() *product_pricing_v0_models.GetOffersResponse {
	return o.Payload
}

func (o *GetItemOffersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(product_pricing_v0_models.GetOffersResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetItemOffersUnauthorized creates a GetItemOffersUnauthorized with default headers values
func NewGetItemOffersUnauthorized() *GetItemOffersUnauthorized {
	return &GetItemOffersUnauthorized{}
}

/* GetItemOffersUnauthorized describes a response with status code 401, with default header values.

The request's Authorization header is not formatted correctly or does not contain a valid token.
*/
type GetItemOffersUnauthorized struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference ID.
	 */
	XAmznRequestID string

	Payload *product_pricing_v0_models.GetOffersResponse
}

func (o *GetItemOffersUnauthorized) Error() string {
	return fmt.Sprintf("[GET /products/pricing/v0/items/{Asin}/offers][%d] getItemOffersUnauthorized  %+v", 401, o.Payload)
}
func (o *GetItemOffersUnauthorized) GetPayload() *product_pricing_v0_models.GetOffersResponse {
	return o.Payload
}

func (o *GetItemOffersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(product_pricing_v0_models.GetOffersResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetItemOffersForbidden creates a GetItemOffersForbidden with default headers values
func NewGetItemOffersForbidden() *GetItemOffersForbidden {
	return &GetItemOffersForbidden{}
}

/* GetItemOffersForbidden describes a response with status code 403, with default header values.

Indicates access to the resource is forbidden. Possible reasons include Access Denied, Unauthorized, Expired Token, or Invalid Signature.
*/
type GetItemOffersForbidden struct {

	/* Unique request reference ID.
	 */
	XAmznRequestID string

	Payload *product_pricing_v0_models.GetOffersResponse
}

func (o *GetItemOffersForbidden) Error() string {
	return fmt.Sprintf("[GET /products/pricing/v0/items/{Asin}/offers][%d] getItemOffersForbidden  %+v", 403, o.Payload)
}
func (o *GetItemOffersForbidden) GetPayload() *product_pricing_v0_models.GetOffersResponse {
	return o.Payload
}

func (o *GetItemOffersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(product_pricing_v0_models.GetOffersResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetItemOffersNotFound creates a GetItemOffersNotFound with default headers values
func NewGetItemOffersNotFound() *GetItemOffersNotFound {
	return &GetItemOffersNotFound{}
}

/* GetItemOffersNotFound describes a response with status code 404, with default header values.

The specified resource does not exist.
*/
type GetItemOffersNotFound struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference ID.
	 */
	XAmznRequestID string

	Payload *product_pricing_v0_models.GetOffersResponse
}

func (o *GetItemOffersNotFound) Error() string {
	return fmt.Sprintf("[GET /products/pricing/v0/items/{Asin}/offers][%d] getItemOffersNotFound  %+v", 404, o.Payload)
}
func (o *GetItemOffersNotFound) GetPayload() *product_pricing_v0_models.GetOffersResponse {
	return o.Payload
}

func (o *GetItemOffersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(product_pricing_v0_models.GetOffersResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetItemOffersTooManyRequests creates a GetItemOffersTooManyRequests with default headers values
func NewGetItemOffersTooManyRequests() *GetItemOffersTooManyRequests {
	return &GetItemOffersTooManyRequests{}
}

/* GetItemOffersTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type GetItemOffersTooManyRequests struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference ID.
	 */
	XAmznRequestID string

	Payload *product_pricing_v0_models.GetOffersResponse
}

func (o *GetItemOffersTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /products/pricing/v0/items/{Asin}/offers][%d] getItemOffersTooManyRequests  %+v", 429, o.Payload)
}
func (o *GetItemOffersTooManyRequests) GetPayload() *product_pricing_v0_models.GetOffersResponse {
	return o.Payload
}

func (o *GetItemOffersTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(product_pricing_v0_models.GetOffersResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetItemOffersInternalServerError creates a GetItemOffersInternalServerError with default headers values
func NewGetItemOffersInternalServerError() *GetItemOffersInternalServerError {
	return &GetItemOffersInternalServerError{}
}

/* GetItemOffersInternalServerError describes a response with status code 500, with default header values.

An unexpected condition occurred that prevented the server from fulfilling the request.
*/
type GetItemOffersInternalServerError struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference ID.
	 */
	XAmznRequestID string

	Payload *product_pricing_v0_models.GetOffersResponse
}

func (o *GetItemOffersInternalServerError) Error() string {
	return fmt.Sprintf("[GET /products/pricing/v0/items/{Asin}/offers][%d] getItemOffersInternalServerError  %+v", 500, o.Payload)
}
func (o *GetItemOffersInternalServerError) GetPayload() *product_pricing_v0_models.GetOffersResponse {
	return o.Payload
}

func (o *GetItemOffersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(product_pricing_v0_models.GetOffersResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetItemOffersServiceUnavailable creates a GetItemOffersServiceUnavailable with default headers values
func NewGetItemOffersServiceUnavailable() *GetItemOffersServiceUnavailable {
	return &GetItemOffersServiceUnavailable{}
}

/* GetItemOffersServiceUnavailable describes a response with status code 503, with default header values.

Temporary overloading or maintenance of the server.
*/
type GetItemOffersServiceUnavailable struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference ID.
	 */
	XAmznRequestID string

	Payload *product_pricing_v0_models.GetOffersResponse
}

func (o *GetItemOffersServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /products/pricing/v0/items/{Asin}/offers][%d] getItemOffersServiceUnavailable  %+v", 503, o.Payload)
}
func (o *GetItemOffersServiceUnavailable) GetPayload() *product_pricing_v0_models.GetOffersResponse {
	return o.Payload
}

func (o *GetItemOffersServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(product_pricing_v0_models.GetOffersResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
