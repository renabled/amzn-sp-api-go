// Code generated by go-swagger; DO NOT EDIT.

package product_pricing

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/renabled/amzn-sp-api-go/api/productPricingV0/product_pricing_v0_models"
)

// GetListingOffersBatchReader is a Reader for the GetListingOffersBatch structure.
type GetListingOffersBatchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetListingOffersBatchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetListingOffersBatchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetListingOffersBatchBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetListingOffersBatchUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetListingOffersBatchForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetListingOffersBatchNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetListingOffersBatchTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetListingOffersBatchInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetListingOffersBatchServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetListingOffersBatchOK creates a GetListingOffersBatchOK with default headers values
func NewGetListingOffersBatchOK() *GetListingOffersBatchOK {
	return &GetListingOffersBatchOK{}
}

/*
GetListingOffersBatchOK describes a response with status code 200, with default header values.

Indicates that requests were run in batch.  Check the batch response status lines for information on whether a batch request succeeded.
*/
type GetListingOffersBatchOK struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference ID.
	 */
	XAmznRequestID string

	Payload *product_pricing_v0_models.GetListingOffersBatchResponse
}

// IsSuccess returns true when this get listing offers batch o k response has a 2xx status code
func (o *GetListingOffersBatchOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get listing offers batch o k response has a 3xx status code
func (o *GetListingOffersBatchOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get listing offers batch o k response has a 4xx status code
func (o *GetListingOffersBatchOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get listing offers batch o k response has a 5xx status code
func (o *GetListingOffersBatchOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get listing offers batch o k response a status code equal to that given
func (o *GetListingOffersBatchOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetListingOffersBatchOK) Error() string {
	return fmt.Sprintf("[POST /batches/products/pricing/v0/listingOffers][%d] getListingOffersBatchOK  %+v", 200, o.Payload)
}

func (o *GetListingOffersBatchOK) String() string {
	return fmt.Sprintf("[POST /batches/products/pricing/v0/listingOffers][%d] getListingOffersBatchOK  %+v", 200, o.Payload)
}

func (o *GetListingOffersBatchOK) GetPayload() *product_pricing_v0_models.GetListingOffersBatchResponse {
	return o.Payload
}

func (o *GetListingOffersBatchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(product_pricing_v0_models.GetListingOffersBatchResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetListingOffersBatchBadRequest creates a GetListingOffersBatchBadRequest with default headers values
func NewGetListingOffersBatchBadRequest() *GetListingOffersBatchBadRequest {
	return &GetListingOffersBatchBadRequest{}
}

/*
GetListingOffersBatchBadRequest describes a response with status code 400, with default header values.

Request has missing or invalid parameters and cannot be parsed.
*/
type GetListingOffersBatchBadRequest struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference ID.
	 */
	XAmznRequestID string

	Payload *product_pricing_v0_models.Errors
}

// IsSuccess returns true when this get listing offers batch bad request response has a 2xx status code
func (o *GetListingOffersBatchBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get listing offers batch bad request response has a 3xx status code
func (o *GetListingOffersBatchBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get listing offers batch bad request response has a 4xx status code
func (o *GetListingOffersBatchBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get listing offers batch bad request response has a 5xx status code
func (o *GetListingOffersBatchBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get listing offers batch bad request response a status code equal to that given
func (o *GetListingOffersBatchBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetListingOffersBatchBadRequest) Error() string {
	return fmt.Sprintf("[POST /batches/products/pricing/v0/listingOffers][%d] getListingOffersBatchBadRequest  %+v", 400, o.Payload)
}

func (o *GetListingOffersBatchBadRequest) String() string {
	return fmt.Sprintf("[POST /batches/products/pricing/v0/listingOffers][%d] getListingOffersBatchBadRequest  %+v", 400, o.Payload)
}

func (o *GetListingOffersBatchBadRequest) GetPayload() *product_pricing_v0_models.Errors {
	return o.Payload
}

func (o *GetListingOffersBatchBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(product_pricing_v0_models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetListingOffersBatchUnauthorized creates a GetListingOffersBatchUnauthorized with default headers values
func NewGetListingOffersBatchUnauthorized() *GetListingOffersBatchUnauthorized {
	return &GetListingOffersBatchUnauthorized{}
}

/*
GetListingOffersBatchUnauthorized describes a response with status code 401, with default header values.

The request's Authorization header is not formatted correctly or does not contain a valid token.
*/
type GetListingOffersBatchUnauthorized struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference ID.
	 */
	XAmznRequestID string

	Payload *product_pricing_v0_models.Errors
}

// IsSuccess returns true when this get listing offers batch unauthorized response has a 2xx status code
func (o *GetListingOffersBatchUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get listing offers batch unauthorized response has a 3xx status code
func (o *GetListingOffersBatchUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get listing offers batch unauthorized response has a 4xx status code
func (o *GetListingOffersBatchUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get listing offers batch unauthorized response has a 5xx status code
func (o *GetListingOffersBatchUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get listing offers batch unauthorized response a status code equal to that given
func (o *GetListingOffersBatchUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetListingOffersBatchUnauthorized) Error() string {
	return fmt.Sprintf("[POST /batches/products/pricing/v0/listingOffers][%d] getListingOffersBatchUnauthorized  %+v", 401, o.Payload)
}

func (o *GetListingOffersBatchUnauthorized) String() string {
	return fmt.Sprintf("[POST /batches/products/pricing/v0/listingOffers][%d] getListingOffersBatchUnauthorized  %+v", 401, o.Payload)
}

func (o *GetListingOffersBatchUnauthorized) GetPayload() *product_pricing_v0_models.Errors {
	return o.Payload
}

func (o *GetListingOffersBatchUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(product_pricing_v0_models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetListingOffersBatchForbidden creates a GetListingOffersBatchForbidden with default headers values
func NewGetListingOffersBatchForbidden() *GetListingOffersBatchForbidden {
	return &GetListingOffersBatchForbidden{}
}

/*
GetListingOffersBatchForbidden describes a response with status code 403, with default header values.

Indicates access to the resource is forbidden. Possible reasons include Access Denied, Unauthorized, Expired Token, or Invalid Signature.
*/
type GetListingOffersBatchForbidden struct {

	/* Unique request reference ID.
	 */
	XAmznRequestID string

	Payload *product_pricing_v0_models.Errors
}

// IsSuccess returns true when this get listing offers batch forbidden response has a 2xx status code
func (o *GetListingOffersBatchForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get listing offers batch forbidden response has a 3xx status code
func (o *GetListingOffersBatchForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get listing offers batch forbidden response has a 4xx status code
func (o *GetListingOffersBatchForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get listing offers batch forbidden response has a 5xx status code
func (o *GetListingOffersBatchForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get listing offers batch forbidden response a status code equal to that given
func (o *GetListingOffersBatchForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetListingOffersBatchForbidden) Error() string {
	return fmt.Sprintf("[POST /batches/products/pricing/v0/listingOffers][%d] getListingOffersBatchForbidden  %+v", 403, o.Payload)
}

func (o *GetListingOffersBatchForbidden) String() string {
	return fmt.Sprintf("[POST /batches/products/pricing/v0/listingOffers][%d] getListingOffersBatchForbidden  %+v", 403, o.Payload)
}

func (o *GetListingOffersBatchForbidden) GetPayload() *product_pricing_v0_models.Errors {
	return o.Payload
}

func (o *GetListingOffersBatchForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(product_pricing_v0_models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetListingOffersBatchNotFound creates a GetListingOffersBatchNotFound with default headers values
func NewGetListingOffersBatchNotFound() *GetListingOffersBatchNotFound {
	return &GetListingOffersBatchNotFound{}
}

/*
GetListingOffersBatchNotFound describes a response with status code 404, with default header values.

The specified resource does not exist.
*/
type GetListingOffersBatchNotFound struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference ID.
	 */
	XAmznRequestID string

	Payload *product_pricing_v0_models.Errors
}

// IsSuccess returns true when this get listing offers batch not found response has a 2xx status code
func (o *GetListingOffersBatchNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get listing offers batch not found response has a 3xx status code
func (o *GetListingOffersBatchNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get listing offers batch not found response has a 4xx status code
func (o *GetListingOffersBatchNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get listing offers batch not found response has a 5xx status code
func (o *GetListingOffersBatchNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get listing offers batch not found response a status code equal to that given
func (o *GetListingOffersBatchNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetListingOffersBatchNotFound) Error() string {
	return fmt.Sprintf("[POST /batches/products/pricing/v0/listingOffers][%d] getListingOffersBatchNotFound  %+v", 404, o.Payload)
}

func (o *GetListingOffersBatchNotFound) String() string {
	return fmt.Sprintf("[POST /batches/products/pricing/v0/listingOffers][%d] getListingOffersBatchNotFound  %+v", 404, o.Payload)
}

func (o *GetListingOffersBatchNotFound) GetPayload() *product_pricing_v0_models.Errors {
	return o.Payload
}

func (o *GetListingOffersBatchNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(product_pricing_v0_models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetListingOffersBatchTooManyRequests creates a GetListingOffersBatchTooManyRequests with default headers values
func NewGetListingOffersBatchTooManyRequests() *GetListingOffersBatchTooManyRequests {
	return &GetListingOffersBatchTooManyRequests{}
}

/*
GetListingOffersBatchTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type GetListingOffersBatchTooManyRequests struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference ID.
	 */
	XAmznRequestID string

	Payload *product_pricing_v0_models.Errors
}

// IsSuccess returns true when this get listing offers batch too many requests response has a 2xx status code
func (o *GetListingOffersBatchTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get listing offers batch too many requests response has a 3xx status code
func (o *GetListingOffersBatchTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get listing offers batch too many requests response has a 4xx status code
func (o *GetListingOffersBatchTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get listing offers batch too many requests response has a 5xx status code
func (o *GetListingOffersBatchTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get listing offers batch too many requests response a status code equal to that given
func (o *GetListingOffersBatchTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetListingOffersBatchTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /batches/products/pricing/v0/listingOffers][%d] getListingOffersBatchTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetListingOffersBatchTooManyRequests) String() string {
	return fmt.Sprintf("[POST /batches/products/pricing/v0/listingOffers][%d] getListingOffersBatchTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetListingOffersBatchTooManyRequests) GetPayload() *product_pricing_v0_models.Errors {
	return o.Payload
}

func (o *GetListingOffersBatchTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(product_pricing_v0_models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetListingOffersBatchInternalServerError creates a GetListingOffersBatchInternalServerError with default headers values
func NewGetListingOffersBatchInternalServerError() *GetListingOffersBatchInternalServerError {
	return &GetListingOffersBatchInternalServerError{}
}

/*
GetListingOffersBatchInternalServerError describes a response with status code 500, with default header values.

An unexpected condition occurred that prevented the server from fulfilling the request.
*/
type GetListingOffersBatchInternalServerError struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference ID.
	 */
	XAmznRequestID string

	Payload *product_pricing_v0_models.Errors
}

// IsSuccess returns true when this get listing offers batch internal server error response has a 2xx status code
func (o *GetListingOffersBatchInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get listing offers batch internal server error response has a 3xx status code
func (o *GetListingOffersBatchInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get listing offers batch internal server error response has a 4xx status code
func (o *GetListingOffersBatchInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get listing offers batch internal server error response has a 5xx status code
func (o *GetListingOffersBatchInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get listing offers batch internal server error response a status code equal to that given
func (o *GetListingOffersBatchInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetListingOffersBatchInternalServerError) Error() string {
	return fmt.Sprintf("[POST /batches/products/pricing/v0/listingOffers][%d] getListingOffersBatchInternalServerError  %+v", 500, o.Payload)
}

func (o *GetListingOffersBatchInternalServerError) String() string {
	return fmt.Sprintf("[POST /batches/products/pricing/v0/listingOffers][%d] getListingOffersBatchInternalServerError  %+v", 500, o.Payload)
}

func (o *GetListingOffersBatchInternalServerError) GetPayload() *product_pricing_v0_models.Errors {
	return o.Payload
}

func (o *GetListingOffersBatchInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(product_pricing_v0_models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetListingOffersBatchServiceUnavailable creates a GetListingOffersBatchServiceUnavailable with default headers values
func NewGetListingOffersBatchServiceUnavailable() *GetListingOffersBatchServiceUnavailable {
	return &GetListingOffersBatchServiceUnavailable{}
}

/*
GetListingOffersBatchServiceUnavailable describes a response with status code 503, with default header values.

Temporary overloading or maintenance of the server.
*/
type GetListingOffersBatchServiceUnavailable struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference ID.
	 */
	XAmznRequestID string

	Payload *product_pricing_v0_models.Errors
}

// IsSuccess returns true when this get listing offers batch service unavailable response has a 2xx status code
func (o *GetListingOffersBatchServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get listing offers batch service unavailable response has a 3xx status code
func (o *GetListingOffersBatchServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get listing offers batch service unavailable response has a 4xx status code
func (o *GetListingOffersBatchServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get listing offers batch service unavailable response has a 5xx status code
func (o *GetListingOffersBatchServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get listing offers batch service unavailable response a status code equal to that given
func (o *GetListingOffersBatchServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetListingOffersBatchServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /batches/products/pricing/v0/listingOffers][%d] getListingOffersBatchServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetListingOffersBatchServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /batches/products/pricing/v0/listingOffers][%d] getListingOffersBatchServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetListingOffersBatchServiceUnavailable) GetPayload() *product_pricing_v0_models.Errors {
	return o.Payload
}

func (o *GetListingOffersBatchServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(product_pricing_v0_models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
