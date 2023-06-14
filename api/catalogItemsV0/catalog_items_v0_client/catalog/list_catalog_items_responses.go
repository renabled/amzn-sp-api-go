// Code generated by go-swagger; DO NOT EDIT.

package catalog

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/renabled/amzn-sp-api-go/api/catalogItemsV0/catalog_items_v0_models"
)

// ListCatalogItemsReader is a Reader for the ListCatalogItems structure.
type ListCatalogItemsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListCatalogItemsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListCatalogItemsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListCatalogItemsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewListCatalogItemsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListCatalogItemsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListCatalogItemsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewListCatalogItemsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListCatalogItemsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewListCatalogItemsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListCatalogItemsOK creates a ListCatalogItemsOK with default headers values
func NewListCatalogItemsOK() *ListCatalogItemsOK {
	return &ListCatalogItemsOK{}
}

/*
ListCatalogItemsOK describes a response with status code 200, with default header values.

Success.
*/
type ListCatalogItemsOK struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference ID.
	 */
	XAmznRequestID string

	Payload *catalog_items_v0_models.ListCatalogItemsResponse
}

// IsSuccess returns true when this list catalog items o k response has a 2xx status code
func (o *ListCatalogItemsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list catalog items o k response has a 3xx status code
func (o *ListCatalogItemsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list catalog items o k response has a 4xx status code
func (o *ListCatalogItemsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list catalog items o k response has a 5xx status code
func (o *ListCatalogItemsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list catalog items o k response a status code equal to that given
func (o *ListCatalogItemsOK) IsCode(code int) bool {
	return code == 200
}

func (o *ListCatalogItemsOK) Error() string {
	return fmt.Sprintf("[GET /catalog/v0/items][%d] listCatalogItemsOK  %+v", 200, o.Payload)
}

func (o *ListCatalogItemsOK) String() string {
	return fmt.Sprintf("[GET /catalog/v0/items][%d] listCatalogItemsOK  %+v", 200, o.Payload)
}

func (o *ListCatalogItemsOK) GetPayload() *catalog_items_v0_models.ListCatalogItemsResponse {
	return o.Payload
}

func (o *ListCatalogItemsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(catalog_items_v0_models.ListCatalogItemsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListCatalogItemsBadRequest creates a ListCatalogItemsBadRequest with default headers values
func NewListCatalogItemsBadRequest() *ListCatalogItemsBadRequest {
	return &ListCatalogItemsBadRequest{}
}

/*
ListCatalogItemsBadRequest describes a response with status code 400, with default header values.

Request has missing or invalid parameters and cannot be parsed.
*/
type ListCatalogItemsBadRequest struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference ID.
	 */
	XAmznRequestID string

	Payload *catalog_items_v0_models.ListCatalogItemsResponse
}

// IsSuccess returns true when this list catalog items bad request response has a 2xx status code
func (o *ListCatalogItemsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list catalog items bad request response has a 3xx status code
func (o *ListCatalogItemsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list catalog items bad request response has a 4xx status code
func (o *ListCatalogItemsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this list catalog items bad request response has a 5xx status code
func (o *ListCatalogItemsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this list catalog items bad request response a status code equal to that given
func (o *ListCatalogItemsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *ListCatalogItemsBadRequest) Error() string {
	return fmt.Sprintf("[GET /catalog/v0/items][%d] listCatalogItemsBadRequest  %+v", 400, o.Payload)
}

func (o *ListCatalogItemsBadRequest) String() string {
	return fmt.Sprintf("[GET /catalog/v0/items][%d] listCatalogItemsBadRequest  %+v", 400, o.Payload)
}

func (o *ListCatalogItemsBadRequest) GetPayload() *catalog_items_v0_models.ListCatalogItemsResponse {
	return o.Payload
}

func (o *ListCatalogItemsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(catalog_items_v0_models.ListCatalogItemsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListCatalogItemsUnauthorized creates a ListCatalogItemsUnauthorized with default headers values
func NewListCatalogItemsUnauthorized() *ListCatalogItemsUnauthorized {
	return &ListCatalogItemsUnauthorized{}
}

/*
ListCatalogItemsUnauthorized describes a response with status code 401, with default header values.

The request's Authorization header is not formatted correctly or does not contain a valid token.
*/
type ListCatalogItemsUnauthorized struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference ID.
	 */
	XAmznRequestID string

	Payload *catalog_items_v0_models.ListCatalogItemsResponse
}

// IsSuccess returns true when this list catalog items unauthorized response has a 2xx status code
func (o *ListCatalogItemsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list catalog items unauthorized response has a 3xx status code
func (o *ListCatalogItemsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list catalog items unauthorized response has a 4xx status code
func (o *ListCatalogItemsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this list catalog items unauthorized response has a 5xx status code
func (o *ListCatalogItemsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this list catalog items unauthorized response a status code equal to that given
func (o *ListCatalogItemsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *ListCatalogItemsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /catalog/v0/items][%d] listCatalogItemsUnauthorized  %+v", 401, o.Payload)
}

func (o *ListCatalogItemsUnauthorized) String() string {
	return fmt.Sprintf("[GET /catalog/v0/items][%d] listCatalogItemsUnauthorized  %+v", 401, o.Payload)
}

func (o *ListCatalogItemsUnauthorized) GetPayload() *catalog_items_v0_models.ListCatalogItemsResponse {
	return o.Payload
}

func (o *ListCatalogItemsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(catalog_items_v0_models.ListCatalogItemsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListCatalogItemsForbidden creates a ListCatalogItemsForbidden with default headers values
func NewListCatalogItemsForbidden() *ListCatalogItemsForbidden {
	return &ListCatalogItemsForbidden{}
}

/*
ListCatalogItemsForbidden describes a response with status code 403, with default header values.

Indicates access to the resource is forbidden. Possible reasons include Access Denied, Unauthorized, Expired Token, or Invalid Signature.
*/
type ListCatalogItemsForbidden struct {

	/* Unique request reference ID.
	 */
	XAmznRequestID string

	Payload *catalog_items_v0_models.ListCatalogItemsResponse
}

// IsSuccess returns true when this list catalog items forbidden response has a 2xx status code
func (o *ListCatalogItemsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list catalog items forbidden response has a 3xx status code
func (o *ListCatalogItemsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list catalog items forbidden response has a 4xx status code
func (o *ListCatalogItemsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this list catalog items forbidden response has a 5xx status code
func (o *ListCatalogItemsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this list catalog items forbidden response a status code equal to that given
func (o *ListCatalogItemsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ListCatalogItemsForbidden) Error() string {
	return fmt.Sprintf("[GET /catalog/v0/items][%d] listCatalogItemsForbidden  %+v", 403, o.Payload)
}

func (o *ListCatalogItemsForbidden) String() string {
	return fmt.Sprintf("[GET /catalog/v0/items][%d] listCatalogItemsForbidden  %+v", 403, o.Payload)
}

func (o *ListCatalogItemsForbidden) GetPayload() *catalog_items_v0_models.ListCatalogItemsResponse {
	return o.Payload
}

func (o *ListCatalogItemsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(catalog_items_v0_models.ListCatalogItemsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListCatalogItemsNotFound creates a ListCatalogItemsNotFound with default headers values
func NewListCatalogItemsNotFound() *ListCatalogItemsNotFound {
	return &ListCatalogItemsNotFound{}
}

/*
ListCatalogItemsNotFound describes a response with status code 404, with default header values.

The specified resource does not exist.
*/
type ListCatalogItemsNotFound struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference ID.
	 */
	XAmznRequestID string

	Payload *catalog_items_v0_models.ListCatalogItemsResponse
}

// IsSuccess returns true when this list catalog items not found response has a 2xx status code
func (o *ListCatalogItemsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list catalog items not found response has a 3xx status code
func (o *ListCatalogItemsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list catalog items not found response has a 4xx status code
func (o *ListCatalogItemsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this list catalog items not found response has a 5xx status code
func (o *ListCatalogItemsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this list catalog items not found response a status code equal to that given
func (o *ListCatalogItemsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *ListCatalogItemsNotFound) Error() string {
	return fmt.Sprintf("[GET /catalog/v0/items][%d] listCatalogItemsNotFound  %+v", 404, o.Payload)
}

func (o *ListCatalogItemsNotFound) String() string {
	return fmt.Sprintf("[GET /catalog/v0/items][%d] listCatalogItemsNotFound  %+v", 404, o.Payload)
}

func (o *ListCatalogItemsNotFound) GetPayload() *catalog_items_v0_models.ListCatalogItemsResponse {
	return o.Payload
}

func (o *ListCatalogItemsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(catalog_items_v0_models.ListCatalogItemsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListCatalogItemsTooManyRequests creates a ListCatalogItemsTooManyRequests with default headers values
func NewListCatalogItemsTooManyRequests() *ListCatalogItemsTooManyRequests {
	return &ListCatalogItemsTooManyRequests{}
}

/*
ListCatalogItemsTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type ListCatalogItemsTooManyRequests struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference ID.
	 */
	XAmznRequestID string

	Payload *catalog_items_v0_models.ListCatalogItemsResponse
}

// IsSuccess returns true when this list catalog items too many requests response has a 2xx status code
func (o *ListCatalogItemsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list catalog items too many requests response has a 3xx status code
func (o *ListCatalogItemsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list catalog items too many requests response has a 4xx status code
func (o *ListCatalogItemsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this list catalog items too many requests response has a 5xx status code
func (o *ListCatalogItemsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this list catalog items too many requests response a status code equal to that given
func (o *ListCatalogItemsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *ListCatalogItemsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /catalog/v0/items][%d] listCatalogItemsTooManyRequests  %+v", 429, o.Payload)
}

func (o *ListCatalogItemsTooManyRequests) String() string {
	return fmt.Sprintf("[GET /catalog/v0/items][%d] listCatalogItemsTooManyRequests  %+v", 429, o.Payload)
}

func (o *ListCatalogItemsTooManyRequests) GetPayload() *catalog_items_v0_models.ListCatalogItemsResponse {
	return o.Payload
}

func (o *ListCatalogItemsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(catalog_items_v0_models.ListCatalogItemsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListCatalogItemsInternalServerError creates a ListCatalogItemsInternalServerError with default headers values
func NewListCatalogItemsInternalServerError() *ListCatalogItemsInternalServerError {
	return &ListCatalogItemsInternalServerError{}
}

/*
ListCatalogItemsInternalServerError describes a response with status code 500, with default header values.

An unexpected condition occurred that prevented the server from fulfilling the request.
*/
type ListCatalogItemsInternalServerError struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference ID.
	 */
	XAmznRequestID string

	Payload *catalog_items_v0_models.ListCatalogItemsResponse
}

// IsSuccess returns true when this list catalog items internal server error response has a 2xx status code
func (o *ListCatalogItemsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list catalog items internal server error response has a 3xx status code
func (o *ListCatalogItemsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list catalog items internal server error response has a 4xx status code
func (o *ListCatalogItemsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this list catalog items internal server error response has a 5xx status code
func (o *ListCatalogItemsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this list catalog items internal server error response a status code equal to that given
func (o *ListCatalogItemsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ListCatalogItemsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /catalog/v0/items][%d] listCatalogItemsInternalServerError  %+v", 500, o.Payload)
}

func (o *ListCatalogItemsInternalServerError) String() string {
	return fmt.Sprintf("[GET /catalog/v0/items][%d] listCatalogItemsInternalServerError  %+v", 500, o.Payload)
}

func (o *ListCatalogItemsInternalServerError) GetPayload() *catalog_items_v0_models.ListCatalogItemsResponse {
	return o.Payload
}

func (o *ListCatalogItemsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(catalog_items_v0_models.ListCatalogItemsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListCatalogItemsServiceUnavailable creates a ListCatalogItemsServiceUnavailable with default headers values
func NewListCatalogItemsServiceUnavailable() *ListCatalogItemsServiceUnavailable {
	return &ListCatalogItemsServiceUnavailable{}
}

/*
ListCatalogItemsServiceUnavailable describes a response with status code 503, with default header values.

Temporary overloading or maintenance of the server.
*/
type ListCatalogItemsServiceUnavailable struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference ID.
	 */
	XAmznRequestID string

	Payload *catalog_items_v0_models.ListCatalogItemsResponse
}

// IsSuccess returns true when this list catalog items service unavailable response has a 2xx status code
func (o *ListCatalogItemsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list catalog items service unavailable response has a 3xx status code
func (o *ListCatalogItemsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list catalog items service unavailable response has a 4xx status code
func (o *ListCatalogItemsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this list catalog items service unavailable response has a 5xx status code
func (o *ListCatalogItemsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this list catalog items service unavailable response a status code equal to that given
func (o *ListCatalogItemsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *ListCatalogItemsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /catalog/v0/items][%d] listCatalogItemsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *ListCatalogItemsServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /catalog/v0/items][%d] listCatalogItemsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *ListCatalogItemsServiceUnavailable) GetPayload() *catalog_items_v0_models.ListCatalogItemsResponse {
	return o.Payload
}

func (o *ListCatalogItemsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(catalog_items_v0_models.ListCatalogItemsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
