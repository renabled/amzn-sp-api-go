// Code generated by go-swagger; DO NOT EDIT.

package fba_inventory

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/renabled/amzn-sp-api-go/api/fbaInventory/fba_inventory_models"
)

// GetInventorySummariesReader is a Reader for the GetInventorySummaries structure.
type GetInventorySummariesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInventorySummariesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetInventorySummariesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetInventorySummariesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetInventorySummariesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetInventorySummariesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetInventorySummariesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetInventorySummariesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetInventorySummariesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetInventorySummariesOK creates a GetInventorySummariesOK with default headers values
func NewGetInventorySummariesOK() *GetInventorySummariesOK {
	return &GetInventorySummariesOK{}
}

/*
GetInventorySummariesOK describes a response with status code 200, with default header values.

OK
*/
type GetInventorySummariesOK struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference ID.
	 */
	XAmznRequestID string

	Payload *fba_inventory_models.GetInventorySummariesResponse
}

// IsSuccess returns true when this get inventory summaries o k response has a 2xx status code
func (o *GetInventorySummariesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get inventory summaries o k response has a 3xx status code
func (o *GetInventorySummariesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get inventory summaries o k response has a 4xx status code
func (o *GetInventorySummariesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get inventory summaries o k response has a 5xx status code
func (o *GetInventorySummariesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get inventory summaries o k response a status code equal to that given
func (o *GetInventorySummariesOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetInventorySummariesOK) Error() string {
	return fmt.Sprintf("[GET /fba/inventory/v1/summaries][%d] getInventorySummariesOK  %+v", 200, o.Payload)
}

func (o *GetInventorySummariesOK) String() string {
	return fmt.Sprintf("[GET /fba/inventory/v1/summaries][%d] getInventorySummariesOK  %+v", 200, o.Payload)
}

func (o *GetInventorySummariesOK) GetPayload() *fba_inventory_models.GetInventorySummariesResponse {
	return o.Payload
}

func (o *GetInventorySummariesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fba_inventory_models.GetInventorySummariesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInventorySummariesBadRequest creates a GetInventorySummariesBadRequest with default headers values
func NewGetInventorySummariesBadRequest() *GetInventorySummariesBadRequest {
	return &GetInventorySummariesBadRequest{}
}

/*
GetInventorySummariesBadRequest describes a response with status code 400, with default header values.

Request has missing or invalid parameters and cannot be parsed.
*/
type GetInventorySummariesBadRequest struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference ID.
	 */
	XAmznRequestID string

	Payload *fba_inventory_models.GetInventorySummariesResponse
}

// IsSuccess returns true when this get inventory summaries bad request response has a 2xx status code
func (o *GetInventorySummariesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get inventory summaries bad request response has a 3xx status code
func (o *GetInventorySummariesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get inventory summaries bad request response has a 4xx status code
func (o *GetInventorySummariesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get inventory summaries bad request response has a 5xx status code
func (o *GetInventorySummariesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get inventory summaries bad request response a status code equal to that given
func (o *GetInventorySummariesBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetInventorySummariesBadRequest) Error() string {
	return fmt.Sprintf("[GET /fba/inventory/v1/summaries][%d] getInventorySummariesBadRequest  %+v", 400, o.Payload)
}

func (o *GetInventorySummariesBadRequest) String() string {
	return fmt.Sprintf("[GET /fba/inventory/v1/summaries][%d] getInventorySummariesBadRequest  %+v", 400, o.Payload)
}

func (o *GetInventorySummariesBadRequest) GetPayload() *fba_inventory_models.GetInventorySummariesResponse {
	return o.Payload
}

func (o *GetInventorySummariesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fba_inventory_models.GetInventorySummariesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInventorySummariesForbidden creates a GetInventorySummariesForbidden with default headers values
func NewGetInventorySummariesForbidden() *GetInventorySummariesForbidden {
	return &GetInventorySummariesForbidden{}
}

/*
GetInventorySummariesForbidden describes a response with status code 403, with default header values.

Indicates access to the resource is forbidden. Possible reasons include Access Denied, Unauthorized, Expired Token, Invalid Signature or Resource Not Found.
*/
type GetInventorySummariesForbidden struct {

	/* Unique request reference ID.
	 */
	XAmznRequestID string

	Payload *fba_inventory_models.GetInventorySummariesResponse
}

// IsSuccess returns true when this get inventory summaries forbidden response has a 2xx status code
func (o *GetInventorySummariesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get inventory summaries forbidden response has a 3xx status code
func (o *GetInventorySummariesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get inventory summaries forbidden response has a 4xx status code
func (o *GetInventorySummariesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get inventory summaries forbidden response has a 5xx status code
func (o *GetInventorySummariesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get inventory summaries forbidden response a status code equal to that given
func (o *GetInventorySummariesForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetInventorySummariesForbidden) Error() string {
	return fmt.Sprintf("[GET /fba/inventory/v1/summaries][%d] getInventorySummariesForbidden  %+v", 403, o.Payload)
}

func (o *GetInventorySummariesForbidden) String() string {
	return fmt.Sprintf("[GET /fba/inventory/v1/summaries][%d] getInventorySummariesForbidden  %+v", 403, o.Payload)
}

func (o *GetInventorySummariesForbidden) GetPayload() *fba_inventory_models.GetInventorySummariesResponse {
	return o.Payload
}

func (o *GetInventorySummariesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(fba_inventory_models.GetInventorySummariesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInventorySummariesNotFound creates a GetInventorySummariesNotFound with default headers values
func NewGetInventorySummariesNotFound() *GetInventorySummariesNotFound {
	return &GetInventorySummariesNotFound{}
}

/*
GetInventorySummariesNotFound describes a response with status code 404, with default header values.

The specified resource does not exist.
*/
type GetInventorySummariesNotFound struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference ID.
	 */
	XAmznRequestID string

	Payload *fba_inventory_models.GetInventorySummariesResponse
}

// IsSuccess returns true when this get inventory summaries not found response has a 2xx status code
func (o *GetInventorySummariesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get inventory summaries not found response has a 3xx status code
func (o *GetInventorySummariesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get inventory summaries not found response has a 4xx status code
func (o *GetInventorySummariesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get inventory summaries not found response has a 5xx status code
func (o *GetInventorySummariesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get inventory summaries not found response a status code equal to that given
func (o *GetInventorySummariesNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetInventorySummariesNotFound) Error() string {
	return fmt.Sprintf("[GET /fba/inventory/v1/summaries][%d] getInventorySummariesNotFound  %+v", 404, o.Payload)
}

func (o *GetInventorySummariesNotFound) String() string {
	return fmt.Sprintf("[GET /fba/inventory/v1/summaries][%d] getInventorySummariesNotFound  %+v", 404, o.Payload)
}

func (o *GetInventorySummariesNotFound) GetPayload() *fba_inventory_models.GetInventorySummariesResponse {
	return o.Payload
}

func (o *GetInventorySummariesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fba_inventory_models.GetInventorySummariesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInventorySummariesTooManyRequests creates a GetInventorySummariesTooManyRequests with default headers values
func NewGetInventorySummariesTooManyRequests() *GetInventorySummariesTooManyRequests {
	return &GetInventorySummariesTooManyRequests{}
}

/*
GetInventorySummariesTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type GetInventorySummariesTooManyRequests struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference ID.
	 */
	XAmznRequestID string

	Payload *fba_inventory_models.GetInventorySummariesResponse
}

// IsSuccess returns true when this get inventory summaries too many requests response has a 2xx status code
func (o *GetInventorySummariesTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get inventory summaries too many requests response has a 3xx status code
func (o *GetInventorySummariesTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get inventory summaries too many requests response has a 4xx status code
func (o *GetInventorySummariesTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get inventory summaries too many requests response has a 5xx status code
func (o *GetInventorySummariesTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get inventory summaries too many requests response a status code equal to that given
func (o *GetInventorySummariesTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetInventorySummariesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /fba/inventory/v1/summaries][%d] getInventorySummariesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetInventorySummariesTooManyRequests) String() string {
	return fmt.Sprintf("[GET /fba/inventory/v1/summaries][%d] getInventorySummariesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetInventorySummariesTooManyRequests) GetPayload() *fba_inventory_models.GetInventorySummariesResponse {
	return o.Payload
}

func (o *GetInventorySummariesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fba_inventory_models.GetInventorySummariesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInventorySummariesInternalServerError creates a GetInventorySummariesInternalServerError with default headers values
func NewGetInventorySummariesInternalServerError() *GetInventorySummariesInternalServerError {
	return &GetInventorySummariesInternalServerError{}
}

/*
GetInventorySummariesInternalServerError describes a response with status code 500, with default header values.

An unexpected condition occurred that prevented the server from fulfilling the request.
*/
type GetInventorySummariesInternalServerError struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference ID.
	 */
	XAmznRequestID string

	Payload *fba_inventory_models.GetInventorySummariesResponse
}

// IsSuccess returns true when this get inventory summaries internal server error response has a 2xx status code
func (o *GetInventorySummariesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get inventory summaries internal server error response has a 3xx status code
func (o *GetInventorySummariesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get inventory summaries internal server error response has a 4xx status code
func (o *GetInventorySummariesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get inventory summaries internal server error response has a 5xx status code
func (o *GetInventorySummariesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get inventory summaries internal server error response a status code equal to that given
func (o *GetInventorySummariesInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetInventorySummariesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /fba/inventory/v1/summaries][%d] getInventorySummariesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetInventorySummariesInternalServerError) String() string {
	return fmt.Sprintf("[GET /fba/inventory/v1/summaries][%d] getInventorySummariesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetInventorySummariesInternalServerError) GetPayload() *fba_inventory_models.GetInventorySummariesResponse {
	return o.Payload
}

func (o *GetInventorySummariesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fba_inventory_models.GetInventorySummariesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInventorySummariesServiceUnavailable creates a GetInventorySummariesServiceUnavailable with default headers values
func NewGetInventorySummariesServiceUnavailable() *GetInventorySummariesServiceUnavailable {
	return &GetInventorySummariesServiceUnavailable{}
}

/*
GetInventorySummariesServiceUnavailable describes a response with status code 503, with default header values.

Temporary overloading or maintenance of the server.
*/
type GetInventorySummariesServiceUnavailable struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference ID.
	 */
	XAmznRequestID string

	Payload *fba_inventory_models.GetInventorySummariesResponse
}

// IsSuccess returns true when this get inventory summaries service unavailable response has a 2xx status code
func (o *GetInventorySummariesServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get inventory summaries service unavailable response has a 3xx status code
func (o *GetInventorySummariesServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get inventory summaries service unavailable response has a 4xx status code
func (o *GetInventorySummariesServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get inventory summaries service unavailable response has a 5xx status code
func (o *GetInventorySummariesServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get inventory summaries service unavailable response a status code equal to that given
func (o *GetInventorySummariesServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetInventorySummariesServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /fba/inventory/v1/summaries][%d] getInventorySummariesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetInventorySummariesServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /fba/inventory/v1/summaries][%d] getInventorySummariesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetInventorySummariesServiceUnavailable) GetPayload() *fba_inventory_models.GetInventorySummariesResponse {
	return o.Payload
}

func (o *GetInventorySummariesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fba_inventory_models.GetInventorySummariesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
