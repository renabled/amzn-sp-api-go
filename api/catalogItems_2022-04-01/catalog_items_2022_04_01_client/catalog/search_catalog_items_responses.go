// Code generated by go-swagger; DO NOT EDIT.

package catalog

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/xamandar/amzn-sp-api-go/api/catalogItems_2022-04-01/catalog_items_2022_04_01_models"
)

// SearchCatalogItemsReader is a Reader for the SearchCatalogItems structure.
type SearchCatalogItemsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SearchCatalogItemsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSearchCatalogItemsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSearchCatalogItemsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSearchCatalogItemsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSearchCatalogItemsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewSearchCatalogItemsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewSearchCatalogItemsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewSearchCatalogItemsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSearchCatalogItemsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewSearchCatalogItemsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSearchCatalogItemsOK creates a SearchCatalogItemsOK with default headers values
func NewSearchCatalogItemsOK() *SearchCatalogItemsOK {
	return &SearchCatalogItemsOK{}
}

/*
SearchCatalogItemsOK describes a response with status code 200, with default header values.

Success.
*/
type SearchCatalogItemsOK struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *catalog_items_2022_04_01_models.ItemSearchResults
}

// IsSuccess returns true when this search catalog items o k response has a 2xx status code
func (o *SearchCatalogItemsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this search catalog items o k response has a 3xx status code
func (o *SearchCatalogItemsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search catalog items o k response has a 4xx status code
func (o *SearchCatalogItemsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this search catalog items o k response has a 5xx status code
func (o *SearchCatalogItemsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this search catalog items o k response a status code equal to that given
func (o *SearchCatalogItemsOK) IsCode(code int) bool {
	return code == 200
}

func (o *SearchCatalogItemsOK) Error() string {
	return fmt.Sprintf("[GET /catalog/2022-04-01/items][%d] searchCatalogItemsOK  %+v", 200, o.Payload)
}

func (o *SearchCatalogItemsOK) String() string {
	return fmt.Sprintf("[GET /catalog/2022-04-01/items][%d] searchCatalogItemsOK  %+v", 200, o.Payload)
}

func (o *SearchCatalogItemsOK) GetPayload() *catalog_items_2022_04_01_models.ItemSearchResults {
	return o.Payload
}

func (o *SearchCatalogItemsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(catalog_items_2022_04_01_models.ItemSearchResults)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchCatalogItemsBadRequest creates a SearchCatalogItemsBadRequest with default headers values
func NewSearchCatalogItemsBadRequest() *SearchCatalogItemsBadRequest {
	return &SearchCatalogItemsBadRequest{}
}

/*
SearchCatalogItemsBadRequest describes a response with status code 400, with default header values.

Request has missing or invalid parameters and cannot be parsed.
*/
type SearchCatalogItemsBadRequest struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *catalog_items_2022_04_01_models.ErrorList
}

// IsSuccess returns true when this search catalog items bad request response has a 2xx status code
func (o *SearchCatalogItemsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search catalog items bad request response has a 3xx status code
func (o *SearchCatalogItemsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search catalog items bad request response has a 4xx status code
func (o *SearchCatalogItemsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this search catalog items bad request response has a 5xx status code
func (o *SearchCatalogItemsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this search catalog items bad request response a status code equal to that given
func (o *SearchCatalogItemsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *SearchCatalogItemsBadRequest) Error() string {
	return fmt.Sprintf("[GET /catalog/2022-04-01/items][%d] searchCatalogItemsBadRequest  %+v", 400, o.Payload)
}

func (o *SearchCatalogItemsBadRequest) String() string {
	return fmt.Sprintf("[GET /catalog/2022-04-01/items][%d] searchCatalogItemsBadRequest  %+v", 400, o.Payload)
}

func (o *SearchCatalogItemsBadRequest) GetPayload() *catalog_items_2022_04_01_models.ErrorList {
	return o.Payload
}

func (o *SearchCatalogItemsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(catalog_items_2022_04_01_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchCatalogItemsForbidden creates a SearchCatalogItemsForbidden with default headers values
func NewSearchCatalogItemsForbidden() *SearchCatalogItemsForbidden {
	return &SearchCatalogItemsForbidden{}
}

/*
SearchCatalogItemsForbidden describes a response with status code 403, with default header values.

Indicates that access to the resource is forbidden. Possible reasons include Access Denied, Unauthorized, Expired Token, or Invalid Signature.
*/
type SearchCatalogItemsForbidden struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *catalog_items_2022_04_01_models.ErrorList
}

// IsSuccess returns true when this search catalog items forbidden response has a 2xx status code
func (o *SearchCatalogItemsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search catalog items forbidden response has a 3xx status code
func (o *SearchCatalogItemsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search catalog items forbidden response has a 4xx status code
func (o *SearchCatalogItemsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this search catalog items forbidden response has a 5xx status code
func (o *SearchCatalogItemsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this search catalog items forbidden response a status code equal to that given
func (o *SearchCatalogItemsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *SearchCatalogItemsForbidden) Error() string {
	return fmt.Sprintf("[GET /catalog/2022-04-01/items][%d] searchCatalogItemsForbidden  %+v", 403, o.Payload)
}

func (o *SearchCatalogItemsForbidden) String() string {
	return fmt.Sprintf("[GET /catalog/2022-04-01/items][%d] searchCatalogItemsForbidden  %+v", 403, o.Payload)
}

func (o *SearchCatalogItemsForbidden) GetPayload() *catalog_items_2022_04_01_models.ErrorList {
	return o.Payload
}

func (o *SearchCatalogItemsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(catalog_items_2022_04_01_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchCatalogItemsNotFound creates a SearchCatalogItemsNotFound with default headers values
func NewSearchCatalogItemsNotFound() *SearchCatalogItemsNotFound {
	return &SearchCatalogItemsNotFound{}
}

/*
SearchCatalogItemsNotFound describes a response with status code 404, with default header values.

The resource specified does not exist.
*/
type SearchCatalogItemsNotFound struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *catalog_items_2022_04_01_models.ErrorList
}

// IsSuccess returns true when this search catalog items not found response has a 2xx status code
func (o *SearchCatalogItemsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search catalog items not found response has a 3xx status code
func (o *SearchCatalogItemsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search catalog items not found response has a 4xx status code
func (o *SearchCatalogItemsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this search catalog items not found response has a 5xx status code
func (o *SearchCatalogItemsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this search catalog items not found response a status code equal to that given
func (o *SearchCatalogItemsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *SearchCatalogItemsNotFound) Error() string {
	return fmt.Sprintf("[GET /catalog/2022-04-01/items][%d] searchCatalogItemsNotFound  %+v", 404, o.Payload)
}

func (o *SearchCatalogItemsNotFound) String() string {
	return fmt.Sprintf("[GET /catalog/2022-04-01/items][%d] searchCatalogItemsNotFound  %+v", 404, o.Payload)
}

func (o *SearchCatalogItemsNotFound) GetPayload() *catalog_items_2022_04_01_models.ErrorList {
	return o.Payload
}

func (o *SearchCatalogItemsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(catalog_items_2022_04_01_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchCatalogItemsRequestEntityTooLarge creates a SearchCatalogItemsRequestEntityTooLarge with default headers values
func NewSearchCatalogItemsRequestEntityTooLarge() *SearchCatalogItemsRequestEntityTooLarge {
	return &SearchCatalogItemsRequestEntityTooLarge{}
}

/*
SearchCatalogItemsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request size exceeded the maximum accepted size.
*/
type SearchCatalogItemsRequestEntityTooLarge struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *catalog_items_2022_04_01_models.ErrorList
}

// IsSuccess returns true when this search catalog items request entity too large response has a 2xx status code
func (o *SearchCatalogItemsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search catalog items request entity too large response has a 3xx status code
func (o *SearchCatalogItemsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search catalog items request entity too large response has a 4xx status code
func (o *SearchCatalogItemsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this search catalog items request entity too large response has a 5xx status code
func (o *SearchCatalogItemsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this search catalog items request entity too large response a status code equal to that given
func (o *SearchCatalogItemsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *SearchCatalogItemsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /catalog/2022-04-01/items][%d] searchCatalogItemsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *SearchCatalogItemsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /catalog/2022-04-01/items][%d] searchCatalogItemsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *SearchCatalogItemsRequestEntityTooLarge) GetPayload() *catalog_items_2022_04_01_models.ErrorList {
	return o.Payload
}

func (o *SearchCatalogItemsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(catalog_items_2022_04_01_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchCatalogItemsUnsupportedMediaType creates a SearchCatalogItemsUnsupportedMediaType with default headers values
func NewSearchCatalogItemsUnsupportedMediaType() *SearchCatalogItemsUnsupportedMediaType {
	return &SearchCatalogItemsUnsupportedMediaType{}
}

/*
SearchCatalogItemsUnsupportedMediaType describes a response with status code 415, with default header values.

The request payload is in an unsupported format.
*/
type SearchCatalogItemsUnsupportedMediaType struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *catalog_items_2022_04_01_models.ErrorList
}

// IsSuccess returns true when this search catalog items unsupported media type response has a 2xx status code
func (o *SearchCatalogItemsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search catalog items unsupported media type response has a 3xx status code
func (o *SearchCatalogItemsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search catalog items unsupported media type response has a 4xx status code
func (o *SearchCatalogItemsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this search catalog items unsupported media type response has a 5xx status code
func (o *SearchCatalogItemsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this search catalog items unsupported media type response a status code equal to that given
func (o *SearchCatalogItemsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *SearchCatalogItemsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /catalog/2022-04-01/items][%d] searchCatalogItemsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *SearchCatalogItemsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /catalog/2022-04-01/items][%d] searchCatalogItemsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *SearchCatalogItemsUnsupportedMediaType) GetPayload() *catalog_items_2022_04_01_models.ErrorList {
	return o.Payload
}

func (o *SearchCatalogItemsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(catalog_items_2022_04_01_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchCatalogItemsTooManyRequests creates a SearchCatalogItemsTooManyRequests with default headers values
func NewSearchCatalogItemsTooManyRequests() *SearchCatalogItemsTooManyRequests {
	return &SearchCatalogItemsTooManyRequests{}
}

/*
SearchCatalogItemsTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type SearchCatalogItemsTooManyRequests struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *catalog_items_2022_04_01_models.ErrorList
}

// IsSuccess returns true when this search catalog items too many requests response has a 2xx status code
func (o *SearchCatalogItemsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search catalog items too many requests response has a 3xx status code
func (o *SearchCatalogItemsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search catalog items too many requests response has a 4xx status code
func (o *SearchCatalogItemsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this search catalog items too many requests response has a 5xx status code
func (o *SearchCatalogItemsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this search catalog items too many requests response a status code equal to that given
func (o *SearchCatalogItemsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *SearchCatalogItemsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /catalog/2022-04-01/items][%d] searchCatalogItemsTooManyRequests  %+v", 429, o.Payload)
}

func (o *SearchCatalogItemsTooManyRequests) String() string {
	return fmt.Sprintf("[GET /catalog/2022-04-01/items][%d] searchCatalogItemsTooManyRequests  %+v", 429, o.Payload)
}

func (o *SearchCatalogItemsTooManyRequests) GetPayload() *catalog_items_2022_04_01_models.ErrorList {
	return o.Payload
}

func (o *SearchCatalogItemsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(catalog_items_2022_04_01_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchCatalogItemsInternalServerError creates a SearchCatalogItemsInternalServerError with default headers values
func NewSearchCatalogItemsInternalServerError() *SearchCatalogItemsInternalServerError {
	return &SearchCatalogItemsInternalServerError{}
}

/*
SearchCatalogItemsInternalServerError describes a response with status code 500, with default header values.

An unexpected condition occurred that prevented the server from fulfilling the request.
*/
type SearchCatalogItemsInternalServerError struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *catalog_items_2022_04_01_models.ErrorList
}

// IsSuccess returns true when this search catalog items internal server error response has a 2xx status code
func (o *SearchCatalogItemsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search catalog items internal server error response has a 3xx status code
func (o *SearchCatalogItemsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search catalog items internal server error response has a 4xx status code
func (o *SearchCatalogItemsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this search catalog items internal server error response has a 5xx status code
func (o *SearchCatalogItemsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this search catalog items internal server error response a status code equal to that given
func (o *SearchCatalogItemsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *SearchCatalogItemsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /catalog/2022-04-01/items][%d] searchCatalogItemsInternalServerError  %+v", 500, o.Payload)
}

func (o *SearchCatalogItemsInternalServerError) String() string {
	return fmt.Sprintf("[GET /catalog/2022-04-01/items][%d] searchCatalogItemsInternalServerError  %+v", 500, o.Payload)
}

func (o *SearchCatalogItemsInternalServerError) GetPayload() *catalog_items_2022_04_01_models.ErrorList {
	return o.Payload
}

func (o *SearchCatalogItemsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(catalog_items_2022_04_01_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchCatalogItemsServiceUnavailable creates a SearchCatalogItemsServiceUnavailable with default headers values
func NewSearchCatalogItemsServiceUnavailable() *SearchCatalogItemsServiceUnavailable {
	return &SearchCatalogItemsServiceUnavailable{}
}

/*
SearchCatalogItemsServiceUnavailable describes a response with status code 503, with default header values.

Temporary overloading or maintenance of the server.
*/
type SearchCatalogItemsServiceUnavailable struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *catalog_items_2022_04_01_models.ErrorList
}

// IsSuccess returns true when this search catalog items service unavailable response has a 2xx status code
func (o *SearchCatalogItemsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search catalog items service unavailable response has a 3xx status code
func (o *SearchCatalogItemsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search catalog items service unavailable response has a 4xx status code
func (o *SearchCatalogItemsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this search catalog items service unavailable response has a 5xx status code
func (o *SearchCatalogItemsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this search catalog items service unavailable response a status code equal to that given
func (o *SearchCatalogItemsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *SearchCatalogItemsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /catalog/2022-04-01/items][%d] searchCatalogItemsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *SearchCatalogItemsServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /catalog/2022-04-01/items][%d] searchCatalogItemsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *SearchCatalogItemsServiceUnavailable) GetPayload() *catalog_items_2022_04_01_models.ErrorList {
	return o.Payload
}

func (o *SearchCatalogItemsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(catalog_items_2022_04_01_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
