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

// GetItemOffersBatchReader is a Reader for the GetItemOffersBatch structure.
type GetItemOffersBatchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetItemOffersBatchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetItemOffersBatchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetItemOffersBatchBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetItemOffersBatchUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetItemOffersBatchForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetItemOffersBatchNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetItemOffersBatchTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetItemOffersBatchInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetItemOffersBatchServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetItemOffersBatchOK creates a GetItemOffersBatchOK with default headers values
func NewGetItemOffersBatchOK() *GetItemOffersBatchOK {
	return &GetItemOffersBatchOK{}
}

/*
GetItemOffersBatchOK describes a response with status code 200, with default header values.

Indicates that requests were run in batch.  Check the batch response status lines for information on whether a batch request succeeded.
*/
type GetItemOffersBatchOK struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference ID.
	 */
	XAmznRequestID string

	Payload *product_pricing_v0_models.GetItemOffersBatchResponse
}

// IsSuccess returns true when this get item offers batch o k response has a 2xx status code
func (o *GetItemOffersBatchOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get item offers batch o k response has a 3xx status code
func (o *GetItemOffersBatchOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get item offers batch o k response has a 4xx status code
func (o *GetItemOffersBatchOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get item offers batch o k response has a 5xx status code
func (o *GetItemOffersBatchOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get item offers batch o k response a status code equal to that given
func (o *GetItemOffersBatchOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetItemOffersBatchOK) Error() string {
	return fmt.Sprintf("[POST /batches/products/pricing/v0/itemOffers][%d] getItemOffersBatchOK  %+v", 200, o.Payload)
}

func (o *GetItemOffersBatchOK) String() string {
	return fmt.Sprintf("[POST /batches/products/pricing/v0/itemOffers][%d] getItemOffersBatchOK  %+v", 200, o.Payload)
}

func (o *GetItemOffersBatchOK) GetPayload() *product_pricing_v0_models.GetItemOffersBatchResponse {
	return o.Payload
}

func (o *GetItemOffersBatchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(product_pricing_v0_models.GetItemOffersBatchResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetItemOffersBatchBadRequest creates a GetItemOffersBatchBadRequest with default headers values
func NewGetItemOffersBatchBadRequest() *GetItemOffersBatchBadRequest {
	return &GetItemOffersBatchBadRequest{}
}

/*
GetItemOffersBatchBadRequest describes a response with status code 400, with default header values.

Request has missing or invalid parameters and cannot be parsed.
*/
type GetItemOffersBatchBadRequest struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference ID.
	 */
	XAmznRequestID string

	Payload *product_pricing_v0_models.Errors
}

// IsSuccess returns true when this get item offers batch bad request response has a 2xx status code
func (o *GetItemOffersBatchBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get item offers batch bad request response has a 3xx status code
func (o *GetItemOffersBatchBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get item offers batch bad request response has a 4xx status code
func (o *GetItemOffersBatchBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get item offers batch bad request response has a 5xx status code
func (o *GetItemOffersBatchBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get item offers batch bad request response a status code equal to that given
func (o *GetItemOffersBatchBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetItemOffersBatchBadRequest) Error() string {
	return fmt.Sprintf("[POST /batches/products/pricing/v0/itemOffers][%d] getItemOffersBatchBadRequest  %+v", 400, o.Payload)
}

func (o *GetItemOffersBatchBadRequest) String() string {
	return fmt.Sprintf("[POST /batches/products/pricing/v0/itemOffers][%d] getItemOffersBatchBadRequest  %+v", 400, o.Payload)
}

func (o *GetItemOffersBatchBadRequest) GetPayload() *product_pricing_v0_models.Errors {
	return o.Payload
}

func (o *GetItemOffersBatchBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetItemOffersBatchUnauthorized creates a GetItemOffersBatchUnauthorized with default headers values
func NewGetItemOffersBatchUnauthorized() *GetItemOffersBatchUnauthorized {
	return &GetItemOffersBatchUnauthorized{}
}

/*
GetItemOffersBatchUnauthorized describes a response with status code 401, with default header values.

The request's Authorization header is not formatted correctly or does not contain a valid token.
*/
type GetItemOffersBatchUnauthorized struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference ID.
	 */
	XAmznRequestID string

	Payload *product_pricing_v0_models.Errors
}

// IsSuccess returns true when this get item offers batch unauthorized response has a 2xx status code
func (o *GetItemOffersBatchUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get item offers batch unauthorized response has a 3xx status code
func (o *GetItemOffersBatchUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get item offers batch unauthorized response has a 4xx status code
func (o *GetItemOffersBatchUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get item offers batch unauthorized response has a 5xx status code
func (o *GetItemOffersBatchUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get item offers batch unauthorized response a status code equal to that given
func (o *GetItemOffersBatchUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetItemOffersBatchUnauthorized) Error() string {
	return fmt.Sprintf("[POST /batches/products/pricing/v0/itemOffers][%d] getItemOffersBatchUnauthorized  %+v", 401, o.Payload)
}

func (o *GetItemOffersBatchUnauthorized) String() string {
	return fmt.Sprintf("[POST /batches/products/pricing/v0/itemOffers][%d] getItemOffersBatchUnauthorized  %+v", 401, o.Payload)
}

func (o *GetItemOffersBatchUnauthorized) GetPayload() *product_pricing_v0_models.Errors {
	return o.Payload
}

func (o *GetItemOffersBatchUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetItemOffersBatchForbidden creates a GetItemOffersBatchForbidden with default headers values
func NewGetItemOffersBatchForbidden() *GetItemOffersBatchForbidden {
	return &GetItemOffersBatchForbidden{}
}

/*
GetItemOffersBatchForbidden describes a response with status code 403, with default header values.

Indicates access to the resource is forbidden. Possible reasons include Access Denied, Unauthorized, Expired Token, or Invalid Signature.
*/
type GetItemOffersBatchForbidden struct {

	/* Unique request reference ID.
	 */
	XAmznRequestID string

	Payload *product_pricing_v0_models.Errors
}

// IsSuccess returns true when this get item offers batch forbidden response has a 2xx status code
func (o *GetItemOffersBatchForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get item offers batch forbidden response has a 3xx status code
func (o *GetItemOffersBatchForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get item offers batch forbidden response has a 4xx status code
func (o *GetItemOffersBatchForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get item offers batch forbidden response has a 5xx status code
func (o *GetItemOffersBatchForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get item offers batch forbidden response a status code equal to that given
func (o *GetItemOffersBatchForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetItemOffersBatchForbidden) Error() string {
	return fmt.Sprintf("[POST /batches/products/pricing/v0/itemOffers][%d] getItemOffersBatchForbidden  %+v", 403, o.Payload)
}

func (o *GetItemOffersBatchForbidden) String() string {
	return fmt.Sprintf("[POST /batches/products/pricing/v0/itemOffers][%d] getItemOffersBatchForbidden  %+v", 403, o.Payload)
}

func (o *GetItemOffersBatchForbidden) GetPayload() *product_pricing_v0_models.Errors {
	return o.Payload
}

func (o *GetItemOffersBatchForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetItemOffersBatchNotFound creates a GetItemOffersBatchNotFound with default headers values
func NewGetItemOffersBatchNotFound() *GetItemOffersBatchNotFound {
	return &GetItemOffersBatchNotFound{}
}

/*
GetItemOffersBatchNotFound describes a response with status code 404, with default header values.

The specified resource does not exist.
*/
type GetItemOffersBatchNotFound struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference ID.
	 */
	XAmznRequestID string

	Payload *product_pricing_v0_models.Errors
}

// IsSuccess returns true when this get item offers batch not found response has a 2xx status code
func (o *GetItemOffersBatchNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get item offers batch not found response has a 3xx status code
func (o *GetItemOffersBatchNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get item offers batch not found response has a 4xx status code
func (o *GetItemOffersBatchNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get item offers batch not found response has a 5xx status code
func (o *GetItemOffersBatchNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get item offers batch not found response a status code equal to that given
func (o *GetItemOffersBatchNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetItemOffersBatchNotFound) Error() string {
	return fmt.Sprintf("[POST /batches/products/pricing/v0/itemOffers][%d] getItemOffersBatchNotFound  %+v", 404, o.Payload)
}

func (o *GetItemOffersBatchNotFound) String() string {
	return fmt.Sprintf("[POST /batches/products/pricing/v0/itemOffers][%d] getItemOffersBatchNotFound  %+v", 404, o.Payload)
}

func (o *GetItemOffersBatchNotFound) GetPayload() *product_pricing_v0_models.Errors {
	return o.Payload
}

func (o *GetItemOffersBatchNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetItemOffersBatchTooManyRequests creates a GetItemOffersBatchTooManyRequests with default headers values
func NewGetItemOffersBatchTooManyRequests() *GetItemOffersBatchTooManyRequests {
	return &GetItemOffersBatchTooManyRequests{}
}

/*
GetItemOffersBatchTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type GetItemOffersBatchTooManyRequests struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference ID.
	 */
	XAmznRequestID string

	Payload *product_pricing_v0_models.Errors
}

// IsSuccess returns true when this get item offers batch too many requests response has a 2xx status code
func (o *GetItemOffersBatchTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get item offers batch too many requests response has a 3xx status code
func (o *GetItemOffersBatchTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get item offers batch too many requests response has a 4xx status code
func (o *GetItemOffersBatchTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get item offers batch too many requests response has a 5xx status code
func (o *GetItemOffersBatchTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get item offers batch too many requests response a status code equal to that given
func (o *GetItemOffersBatchTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetItemOffersBatchTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /batches/products/pricing/v0/itemOffers][%d] getItemOffersBatchTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetItemOffersBatchTooManyRequests) String() string {
	return fmt.Sprintf("[POST /batches/products/pricing/v0/itemOffers][%d] getItemOffersBatchTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetItemOffersBatchTooManyRequests) GetPayload() *product_pricing_v0_models.Errors {
	return o.Payload
}

func (o *GetItemOffersBatchTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetItemOffersBatchInternalServerError creates a GetItemOffersBatchInternalServerError with default headers values
func NewGetItemOffersBatchInternalServerError() *GetItemOffersBatchInternalServerError {
	return &GetItemOffersBatchInternalServerError{}
}

/*
GetItemOffersBatchInternalServerError describes a response with status code 500, with default header values.

An unexpected condition occurred that prevented the server from fulfilling the request.
*/
type GetItemOffersBatchInternalServerError struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference ID.
	 */
	XAmznRequestID string

	Payload *product_pricing_v0_models.Errors
}

// IsSuccess returns true when this get item offers batch internal server error response has a 2xx status code
func (o *GetItemOffersBatchInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get item offers batch internal server error response has a 3xx status code
func (o *GetItemOffersBatchInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get item offers batch internal server error response has a 4xx status code
func (o *GetItemOffersBatchInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get item offers batch internal server error response has a 5xx status code
func (o *GetItemOffersBatchInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get item offers batch internal server error response a status code equal to that given
func (o *GetItemOffersBatchInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetItemOffersBatchInternalServerError) Error() string {
	return fmt.Sprintf("[POST /batches/products/pricing/v0/itemOffers][%d] getItemOffersBatchInternalServerError  %+v", 500, o.Payload)
}

func (o *GetItemOffersBatchInternalServerError) String() string {
	return fmt.Sprintf("[POST /batches/products/pricing/v0/itemOffers][%d] getItemOffersBatchInternalServerError  %+v", 500, o.Payload)
}

func (o *GetItemOffersBatchInternalServerError) GetPayload() *product_pricing_v0_models.Errors {
	return o.Payload
}

func (o *GetItemOffersBatchInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetItemOffersBatchServiceUnavailable creates a GetItemOffersBatchServiceUnavailable with default headers values
func NewGetItemOffersBatchServiceUnavailable() *GetItemOffersBatchServiceUnavailable {
	return &GetItemOffersBatchServiceUnavailable{}
}

/*
GetItemOffersBatchServiceUnavailable describes a response with status code 503, with default header values.

Temporary overloading or maintenance of the server.
*/
type GetItemOffersBatchServiceUnavailable struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference ID.
	 */
	XAmznRequestID string

	Payload *product_pricing_v0_models.Errors
}

// IsSuccess returns true when this get item offers batch service unavailable response has a 2xx status code
func (o *GetItemOffersBatchServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get item offers batch service unavailable response has a 3xx status code
func (o *GetItemOffersBatchServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get item offers batch service unavailable response has a 4xx status code
func (o *GetItemOffersBatchServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get item offers batch service unavailable response has a 5xx status code
func (o *GetItemOffersBatchServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get item offers batch service unavailable response a status code equal to that given
func (o *GetItemOffersBatchServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetItemOffersBatchServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /batches/products/pricing/v0/itemOffers][%d] getItemOffersBatchServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetItemOffersBatchServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /batches/products/pricing/v0/itemOffers][%d] getItemOffersBatchServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetItemOffersBatchServiceUnavailable) GetPayload() *product_pricing_v0_models.Errors {
	return o.Payload
}

func (o *GetItemOffersBatchServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
