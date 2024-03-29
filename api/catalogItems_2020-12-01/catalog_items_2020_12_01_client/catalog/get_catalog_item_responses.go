// Code generated by go-swagger; DO NOT EDIT.

package catalog

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/renabled/amzn-sp-api-go/api/catalogItems_2020-12-01/catalog_items_2020_12_01_models"
)

// GetCatalogItemReader is a Reader for the GetCatalogItem structure.
type GetCatalogItemReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCatalogItemReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCatalogItemOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetCatalogItemBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetCatalogItemForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetCatalogItemNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetCatalogItemRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetCatalogItemUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetCatalogItemTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetCatalogItemInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetCatalogItemServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetCatalogItemOK creates a GetCatalogItemOK with default headers values
func NewGetCatalogItemOK() *GetCatalogItemOK {
	return &GetCatalogItemOK{}
}

/*
GetCatalogItemOK describes a response with status code 200, with default header values.

Success.
*/
type GetCatalogItemOK struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *catalog_items_2020_12_01_models.Item
}

// IsSuccess returns true when this get catalog item o k response has a 2xx status code
func (o *GetCatalogItemOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get catalog item o k response has a 3xx status code
func (o *GetCatalogItemOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get catalog item o k response has a 4xx status code
func (o *GetCatalogItemOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get catalog item o k response has a 5xx status code
func (o *GetCatalogItemOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get catalog item o k response a status code equal to that given
func (o *GetCatalogItemOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetCatalogItemOK) Error() string {
	return fmt.Sprintf("[GET /catalog/2020-12-01/items/{asin}][%d] getCatalogItemOK  %+v", 200, o.Payload)
}

func (o *GetCatalogItemOK) String() string {
	return fmt.Sprintf("[GET /catalog/2020-12-01/items/{asin}][%d] getCatalogItemOK  %+v", 200, o.Payload)
}

func (o *GetCatalogItemOK) GetPayload() *catalog_items_2020_12_01_models.Item {
	return o.Payload
}

func (o *GetCatalogItemOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(catalog_items_2020_12_01_models.Item)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCatalogItemBadRequest creates a GetCatalogItemBadRequest with default headers values
func NewGetCatalogItemBadRequest() *GetCatalogItemBadRequest {
	return &GetCatalogItemBadRequest{}
}

/*
GetCatalogItemBadRequest describes a response with status code 400, with default header values.

Request has missing or invalid parameters and cannot be parsed.
*/
type GetCatalogItemBadRequest struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *catalog_items_2020_12_01_models.ErrorList
}

// IsSuccess returns true when this get catalog item bad request response has a 2xx status code
func (o *GetCatalogItemBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get catalog item bad request response has a 3xx status code
func (o *GetCatalogItemBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get catalog item bad request response has a 4xx status code
func (o *GetCatalogItemBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get catalog item bad request response has a 5xx status code
func (o *GetCatalogItemBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get catalog item bad request response a status code equal to that given
func (o *GetCatalogItemBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetCatalogItemBadRequest) Error() string {
	return fmt.Sprintf("[GET /catalog/2020-12-01/items/{asin}][%d] getCatalogItemBadRequest  %+v", 400, o.Payload)
}

func (o *GetCatalogItemBadRequest) String() string {
	return fmt.Sprintf("[GET /catalog/2020-12-01/items/{asin}][%d] getCatalogItemBadRequest  %+v", 400, o.Payload)
}

func (o *GetCatalogItemBadRequest) GetPayload() *catalog_items_2020_12_01_models.ErrorList {
	return o.Payload
}

func (o *GetCatalogItemBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(catalog_items_2020_12_01_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCatalogItemForbidden creates a GetCatalogItemForbidden with default headers values
func NewGetCatalogItemForbidden() *GetCatalogItemForbidden {
	return &GetCatalogItemForbidden{}
}

/*
GetCatalogItemForbidden describes a response with status code 403, with default header values.

Indicates that access to the resource is forbidden. Possible reasons include Access Denied, Unauthorized, Expired Token, or Invalid Signature.
*/
type GetCatalogItemForbidden struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *catalog_items_2020_12_01_models.ErrorList
}

// IsSuccess returns true when this get catalog item forbidden response has a 2xx status code
func (o *GetCatalogItemForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get catalog item forbidden response has a 3xx status code
func (o *GetCatalogItemForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get catalog item forbidden response has a 4xx status code
func (o *GetCatalogItemForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get catalog item forbidden response has a 5xx status code
func (o *GetCatalogItemForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get catalog item forbidden response a status code equal to that given
func (o *GetCatalogItemForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetCatalogItemForbidden) Error() string {
	return fmt.Sprintf("[GET /catalog/2020-12-01/items/{asin}][%d] getCatalogItemForbidden  %+v", 403, o.Payload)
}

func (o *GetCatalogItemForbidden) String() string {
	return fmt.Sprintf("[GET /catalog/2020-12-01/items/{asin}][%d] getCatalogItemForbidden  %+v", 403, o.Payload)
}

func (o *GetCatalogItemForbidden) GetPayload() *catalog_items_2020_12_01_models.ErrorList {
	return o.Payload
}

func (o *GetCatalogItemForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(catalog_items_2020_12_01_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCatalogItemNotFound creates a GetCatalogItemNotFound with default headers values
func NewGetCatalogItemNotFound() *GetCatalogItemNotFound {
	return &GetCatalogItemNotFound{}
}

/*
GetCatalogItemNotFound describes a response with status code 404, with default header values.

The resource specified does not exist.
*/
type GetCatalogItemNotFound struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *catalog_items_2020_12_01_models.ErrorList
}

// IsSuccess returns true when this get catalog item not found response has a 2xx status code
func (o *GetCatalogItemNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get catalog item not found response has a 3xx status code
func (o *GetCatalogItemNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get catalog item not found response has a 4xx status code
func (o *GetCatalogItemNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get catalog item not found response has a 5xx status code
func (o *GetCatalogItemNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get catalog item not found response a status code equal to that given
func (o *GetCatalogItemNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetCatalogItemNotFound) Error() string {
	return fmt.Sprintf("[GET /catalog/2020-12-01/items/{asin}][%d] getCatalogItemNotFound  %+v", 404, o.Payload)
}

func (o *GetCatalogItemNotFound) String() string {
	return fmt.Sprintf("[GET /catalog/2020-12-01/items/{asin}][%d] getCatalogItemNotFound  %+v", 404, o.Payload)
}

func (o *GetCatalogItemNotFound) GetPayload() *catalog_items_2020_12_01_models.ErrorList {
	return o.Payload
}

func (o *GetCatalogItemNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(catalog_items_2020_12_01_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCatalogItemRequestEntityTooLarge creates a GetCatalogItemRequestEntityTooLarge with default headers values
func NewGetCatalogItemRequestEntityTooLarge() *GetCatalogItemRequestEntityTooLarge {
	return &GetCatalogItemRequestEntityTooLarge{}
}

/*
GetCatalogItemRequestEntityTooLarge describes a response with status code 413, with default header values.

The request size exceeded the maximum accepted size.
*/
type GetCatalogItemRequestEntityTooLarge struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *catalog_items_2020_12_01_models.ErrorList
}

// IsSuccess returns true when this get catalog item request entity too large response has a 2xx status code
func (o *GetCatalogItemRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get catalog item request entity too large response has a 3xx status code
func (o *GetCatalogItemRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get catalog item request entity too large response has a 4xx status code
func (o *GetCatalogItemRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get catalog item request entity too large response has a 5xx status code
func (o *GetCatalogItemRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get catalog item request entity too large response a status code equal to that given
func (o *GetCatalogItemRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetCatalogItemRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /catalog/2020-12-01/items/{asin}][%d] getCatalogItemRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetCatalogItemRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /catalog/2020-12-01/items/{asin}][%d] getCatalogItemRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetCatalogItemRequestEntityTooLarge) GetPayload() *catalog_items_2020_12_01_models.ErrorList {
	return o.Payload
}

func (o *GetCatalogItemRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(catalog_items_2020_12_01_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCatalogItemUnsupportedMediaType creates a GetCatalogItemUnsupportedMediaType with default headers values
func NewGetCatalogItemUnsupportedMediaType() *GetCatalogItemUnsupportedMediaType {
	return &GetCatalogItemUnsupportedMediaType{}
}

/*
GetCatalogItemUnsupportedMediaType describes a response with status code 415, with default header values.

The request payload is in an unsupported format.
*/
type GetCatalogItemUnsupportedMediaType struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *catalog_items_2020_12_01_models.ErrorList
}

// IsSuccess returns true when this get catalog item unsupported media type response has a 2xx status code
func (o *GetCatalogItemUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get catalog item unsupported media type response has a 3xx status code
func (o *GetCatalogItemUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get catalog item unsupported media type response has a 4xx status code
func (o *GetCatalogItemUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get catalog item unsupported media type response has a 5xx status code
func (o *GetCatalogItemUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get catalog item unsupported media type response a status code equal to that given
func (o *GetCatalogItemUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetCatalogItemUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /catalog/2020-12-01/items/{asin}][%d] getCatalogItemUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetCatalogItemUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /catalog/2020-12-01/items/{asin}][%d] getCatalogItemUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetCatalogItemUnsupportedMediaType) GetPayload() *catalog_items_2020_12_01_models.ErrorList {
	return o.Payload
}

func (o *GetCatalogItemUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(catalog_items_2020_12_01_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCatalogItemTooManyRequests creates a GetCatalogItemTooManyRequests with default headers values
func NewGetCatalogItemTooManyRequests() *GetCatalogItemTooManyRequests {
	return &GetCatalogItemTooManyRequests{}
}

/*
GetCatalogItemTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type GetCatalogItemTooManyRequests struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *catalog_items_2020_12_01_models.ErrorList
}

// IsSuccess returns true when this get catalog item too many requests response has a 2xx status code
func (o *GetCatalogItemTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get catalog item too many requests response has a 3xx status code
func (o *GetCatalogItemTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get catalog item too many requests response has a 4xx status code
func (o *GetCatalogItemTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get catalog item too many requests response has a 5xx status code
func (o *GetCatalogItemTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get catalog item too many requests response a status code equal to that given
func (o *GetCatalogItemTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetCatalogItemTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /catalog/2020-12-01/items/{asin}][%d] getCatalogItemTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetCatalogItemTooManyRequests) String() string {
	return fmt.Sprintf("[GET /catalog/2020-12-01/items/{asin}][%d] getCatalogItemTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetCatalogItemTooManyRequests) GetPayload() *catalog_items_2020_12_01_models.ErrorList {
	return o.Payload
}

func (o *GetCatalogItemTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(catalog_items_2020_12_01_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCatalogItemInternalServerError creates a GetCatalogItemInternalServerError with default headers values
func NewGetCatalogItemInternalServerError() *GetCatalogItemInternalServerError {
	return &GetCatalogItemInternalServerError{}
}

/*
GetCatalogItemInternalServerError describes a response with status code 500, with default header values.

An unexpected condition occurred that prevented the server from fulfilling the request.
*/
type GetCatalogItemInternalServerError struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *catalog_items_2020_12_01_models.ErrorList
}

// IsSuccess returns true when this get catalog item internal server error response has a 2xx status code
func (o *GetCatalogItemInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get catalog item internal server error response has a 3xx status code
func (o *GetCatalogItemInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get catalog item internal server error response has a 4xx status code
func (o *GetCatalogItemInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get catalog item internal server error response has a 5xx status code
func (o *GetCatalogItemInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get catalog item internal server error response a status code equal to that given
func (o *GetCatalogItemInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetCatalogItemInternalServerError) Error() string {
	return fmt.Sprintf("[GET /catalog/2020-12-01/items/{asin}][%d] getCatalogItemInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCatalogItemInternalServerError) String() string {
	return fmt.Sprintf("[GET /catalog/2020-12-01/items/{asin}][%d] getCatalogItemInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCatalogItemInternalServerError) GetPayload() *catalog_items_2020_12_01_models.ErrorList {
	return o.Payload
}

func (o *GetCatalogItemInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(catalog_items_2020_12_01_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCatalogItemServiceUnavailable creates a GetCatalogItemServiceUnavailable with default headers values
func NewGetCatalogItemServiceUnavailable() *GetCatalogItemServiceUnavailable {
	return &GetCatalogItemServiceUnavailable{}
}

/*
GetCatalogItemServiceUnavailable describes a response with status code 503, with default header values.

Temporary overloading or maintenance of the server.
*/
type GetCatalogItemServiceUnavailable struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *catalog_items_2020_12_01_models.ErrorList
}

// IsSuccess returns true when this get catalog item service unavailable response has a 2xx status code
func (o *GetCatalogItemServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get catalog item service unavailable response has a 3xx status code
func (o *GetCatalogItemServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get catalog item service unavailable response has a 4xx status code
func (o *GetCatalogItemServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get catalog item service unavailable response has a 5xx status code
func (o *GetCatalogItemServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get catalog item service unavailable response a status code equal to that given
func (o *GetCatalogItemServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetCatalogItemServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /catalog/2020-12-01/items/{asin}][%d] getCatalogItemServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetCatalogItemServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /catalog/2020-12-01/items/{asin}][%d] getCatalogItemServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetCatalogItemServiceUnavailable) GetPayload() *catalog_items_2020_12_01_models.ErrorList {
	return o.Payload
}

func (o *GetCatalogItemServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(catalog_items_2020_12_01_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
