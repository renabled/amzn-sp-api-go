// Code generated by go-swagger; DO NOT EDIT.

package listings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/renabled/amzn-sp-api-go/api/listingsItems_2021-08-01/listings_items_2021_08_01_models"
)

// PutListingsItemReader is a Reader for the PutListingsItem structure.
type PutListingsItemReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutListingsItemReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutListingsItemOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutListingsItemBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutListingsItemForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPutListingsItemRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPutListingsItemUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPutListingsItemTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutListingsItemInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPutListingsItemServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutListingsItemOK creates a PutListingsItemOK with default headers values
func NewPutListingsItemOK() *PutListingsItemOK {
	return &PutListingsItemOK{}
}

/*
PutListingsItemOK describes a response with status code 200, with default header values.

Successfully understood the request to create or fully-update a listings item. See the response to determine if the submission has been accepted.
*/
type PutListingsItemOK struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *listings_items_2021_08_01_models.ListingsItemSubmissionResponse
}

// IsSuccess returns true when this put listings item o k response has a 2xx status code
func (o *PutListingsItemOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put listings item o k response has a 3xx status code
func (o *PutListingsItemOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put listings item o k response has a 4xx status code
func (o *PutListingsItemOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this put listings item o k response has a 5xx status code
func (o *PutListingsItemOK) IsServerError() bool {
	return false
}

// IsCode returns true when this put listings item o k response a status code equal to that given
func (o *PutListingsItemOK) IsCode(code int) bool {
	return code == 200
}

func (o *PutListingsItemOK) Error() string {
	return fmt.Sprintf("[PUT /listings/2021-08-01/items/{sellerId}/{sku}][%d] putListingsItemOK  %+v", 200, o.Payload)
}

func (o *PutListingsItemOK) String() string {
	return fmt.Sprintf("[PUT /listings/2021-08-01/items/{sellerId}/{sku}][%d] putListingsItemOK  %+v", 200, o.Payload)
}

func (o *PutListingsItemOK) GetPayload() *listings_items_2021_08_01_models.ListingsItemSubmissionResponse {
	return o.Payload
}

func (o *PutListingsItemOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(listings_items_2021_08_01_models.ListingsItemSubmissionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutListingsItemBadRequest creates a PutListingsItemBadRequest with default headers values
func NewPutListingsItemBadRequest() *PutListingsItemBadRequest {
	return &PutListingsItemBadRequest{}
}

/*
PutListingsItemBadRequest describes a response with status code 400, with default header values.

Request has missing or invalid parameters and cannot be parsed.
*/
type PutListingsItemBadRequest struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *listings_items_2021_08_01_models.ErrorList
}

// IsSuccess returns true when this put listings item bad request response has a 2xx status code
func (o *PutListingsItemBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put listings item bad request response has a 3xx status code
func (o *PutListingsItemBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put listings item bad request response has a 4xx status code
func (o *PutListingsItemBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this put listings item bad request response has a 5xx status code
func (o *PutListingsItemBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this put listings item bad request response a status code equal to that given
func (o *PutListingsItemBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PutListingsItemBadRequest) Error() string {
	return fmt.Sprintf("[PUT /listings/2021-08-01/items/{sellerId}/{sku}][%d] putListingsItemBadRequest  %+v", 400, o.Payload)
}

func (o *PutListingsItemBadRequest) String() string {
	return fmt.Sprintf("[PUT /listings/2021-08-01/items/{sellerId}/{sku}][%d] putListingsItemBadRequest  %+v", 400, o.Payload)
}

func (o *PutListingsItemBadRequest) GetPayload() *listings_items_2021_08_01_models.ErrorList {
	return o.Payload
}

func (o *PutListingsItemBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(listings_items_2021_08_01_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutListingsItemForbidden creates a PutListingsItemForbidden with default headers values
func NewPutListingsItemForbidden() *PutListingsItemForbidden {
	return &PutListingsItemForbidden{}
}

/*
PutListingsItemForbidden describes a response with status code 403, with default header values.

Indicates that access to the resource is forbidden. Possible reasons include Access Denied, Unauthorized, Expired Token, or Invalid Signature.
*/
type PutListingsItemForbidden struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *listings_items_2021_08_01_models.ErrorList
}

// IsSuccess returns true when this put listings item forbidden response has a 2xx status code
func (o *PutListingsItemForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put listings item forbidden response has a 3xx status code
func (o *PutListingsItemForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put listings item forbidden response has a 4xx status code
func (o *PutListingsItemForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this put listings item forbidden response has a 5xx status code
func (o *PutListingsItemForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this put listings item forbidden response a status code equal to that given
func (o *PutListingsItemForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PutListingsItemForbidden) Error() string {
	return fmt.Sprintf("[PUT /listings/2021-08-01/items/{sellerId}/{sku}][%d] putListingsItemForbidden  %+v", 403, o.Payload)
}

func (o *PutListingsItemForbidden) String() string {
	return fmt.Sprintf("[PUT /listings/2021-08-01/items/{sellerId}/{sku}][%d] putListingsItemForbidden  %+v", 403, o.Payload)
}

func (o *PutListingsItemForbidden) GetPayload() *listings_items_2021_08_01_models.ErrorList {
	return o.Payload
}

func (o *PutListingsItemForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(listings_items_2021_08_01_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutListingsItemRequestEntityTooLarge creates a PutListingsItemRequestEntityTooLarge with default headers values
func NewPutListingsItemRequestEntityTooLarge() *PutListingsItemRequestEntityTooLarge {
	return &PutListingsItemRequestEntityTooLarge{}
}

/*
PutListingsItemRequestEntityTooLarge describes a response with status code 413, with default header values.

The request size exceeded the maximum accepted size.
*/
type PutListingsItemRequestEntityTooLarge struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *listings_items_2021_08_01_models.ErrorList
}

// IsSuccess returns true when this put listings item request entity too large response has a 2xx status code
func (o *PutListingsItemRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put listings item request entity too large response has a 3xx status code
func (o *PutListingsItemRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put listings item request entity too large response has a 4xx status code
func (o *PutListingsItemRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this put listings item request entity too large response has a 5xx status code
func (o *PutListingsItemRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this put listings item request entity too large response a status code equal to that given
func (o *PutListingsItemRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PutListingsItemRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PUT /listings/2021-08-01/items/{sellerId}/{sku}][%d] putListingsItemRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutListingsItemRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[PUT /listings/2021-08-01/items/{sellerId}/{sku}][%d] putListingsItemRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutListingsItemRequestEntityTooLarge) GetPayload() *listings_items_2021_08_01_models.ErrorList {
	return o.Payload
}

func (o *PutListingsItemRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(listings_items_2021_08_01_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutListingsItemUnsupportedMediaType creates a PutListingsItemUnsupportedMediaType with default headers values
func NewPutListingsItemUnsupportedMediaType() *PutListingsItemUnsupportedMediaType {
	return &PutListingsItemUnsupportedMediaType{}
}

/*
PutListingsItemUnsupportedMediaType describes a response with status code 415, with default header values.

The request payload is in an unsupported format.
*/
type PutListingsItemUnsupportedMediaType struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *listings_items_2021_08_01_models.ErrorList
}

// IsSuccess returns true when this put listings item unsupported media type response has a 2xx status code
func (o *PutListingsItemUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put listings item unsupported media type response has a 3xx status code
func (o *PutListingsItemUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put listings item unsupported media type response has a 4xx status code
func (o *PutListingsItemUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this put listings item unsupported media type response has a 5xx status code
func (o *PutListingsItemUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this put listings item unsupported media type response a status code equal to that given
func (o *PutListingsItemUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PutListingsItemUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PUT /listings/2021-08-01/items/{sellerId}/{sku}][%d] putListingsItemUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutListingsItemUnsupportedMediaType) String() string {
	return fmt.Sprintf("[PUT /listings/2021-08-01/items/{sellerId}/{sku}][%d] putListingsItemUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutListingsItemUnsupportedMediaType) GetPayload() *listings_items_2021_08_01_models.ErrorList {
	return o.Payload
}

func (o *PutListingsItemUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(listings_items_2021_08_01_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutListingsItemTooManyRequests creates a PutListingsItemTooManyRequests with default headers values
func NewPutListingsItemTooManyRequests() *PutListingsItemTooManyRequests {
	return &PutListingsItemTooManyRequests{}
}

/*
PutListingsItemTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type PutListingsItemTooManyRequests struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *listings_items_2021_08_01_models.ErrorList
}

// IsSuccess returns true when this put listings item too many requests response has a 2xx status code
func (o *PutListingsItemTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put listings item too many requests response has a 3xx status code
func (o *PutListingsItemTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put listings item too many requests response has a 4xx status code
func (o *PutListingsItemTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this put listings item too many requests response has a 5xx status code
func (o *PutListingsItemTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this put listings item too many requests response a status code equal to that given
func (o *PutListingsItemTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PutListingsItemTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /listings/2021-08-01/items/{sellerId}/{sku}][%d] putListingsItemTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutListingsItemTooManyRequests) String() string {
	return fmt.Sprintf("[PUT /listings/2021-08-01/items/{sellerId}/{sku}][%d] putListingsItemTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutListingsItemTooManyRequests) GetPayload() *listings_items_2021_08_01_models.ErrorList {
	return o.Payload
}

func (o *PutListingsItemTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(listings_items_2021_08_01_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutListingsItemInternalServerError creates a PutListingsItemInternalServerError with default headers values
func NewPutListingsItemInternalServerError() *PutListingsItemInternalServerError {
	return &PutListingsItemInternalServerError{}
}

/*
PutListingsItemInternalServerError describes a response with status code 500, with default header values.

An unexpected condition occurred that prevented the server from fulfilling the request.
*/
type PutListingsItemInternalServerError struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *listings_items_2021_08_01_models.ErrorList
}

// IsSuccess returns true when this put listings item internal server error response has a 2xx status code
func (o *PutListingsItemInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put listings item internal server error response has a 3xx status code
func (o *PutListingsItemInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put listings item internal server error response has a 4xx status code
func (o *PutListingsItemInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this put listings item internal server error response has a 5xx status code
func (o *PutListingsItemInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this put listings item internal server error response a status code equal to that given
func (o *PutListingsItemInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PutListingsItemInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /listings/2021-08-01/items/{sellerId}/{sku}][%d] putListingsItemInternalServerError  %+v", 500, o.Payload)
}

func (o *PutListingsItemInternalServerError) String() string {
	return fmt.Sprintf("[PUT /listings/2021-08-01/items/{sellerId}/{sku}][%d] putListingsItemInternalServerError  %+v", 500, o.Payload)
}

func (o *PutListingsItemInternalServerError) GetPayload() *listings_items_2021_08_01_models.ErrorList {
	return o.Payload
}

func (o *PutListingsItemInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(listings_items_2021_08_01_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutListingsItemServiceUnavailable creates a PutListingsItemServiceUnavailable with default headers values
func NewPutListingsItemServiceUnavailable() *PutListingsItemServiceUnavailable {
	return &PutListingsItemServiceUnavailable{}
}

/*
PutListingsItemServiceUnavailable describes a response with status code 503, with default header values.

Temporary overloading or maintenance of the server.
*/
type PutListingsItemServiceUnavailable struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *listings_items_2021_08_01_models.ErrorList
}

// IsSuccess returns true when this put listings item service unavailable response has a 2xx status code
func (o *PutListingsItemServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put listings item service unavailable response has a 3xx status code
func (o *PutListingsItemServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put listings item service unavailable response has a 4xx status code
func (o *PutListingsItemServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this put listings item service unavailable response has a 5xx status code
func (o *PutListingsItemServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this put listings item service unavailable response a status code equal to that given
func (o *PutListingsItemServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PutListingsItemServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /listings/2021-08-01/items/{sellerId}/{sku}][%d] putListingsItemServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutListingsItemServiceUnavailable) String() string {
	return fmt.Sprintf("[PUT /listings/2021-08-01/items/{sellerId}/{sku}][%d] putListingsItemServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutListingsItemServiceUnavailable) GetPayload() *listings_items_2021_08_01_models.ErrorList {
	return o.Payload
}

func (o *PutListingsItemServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(listings_items_2021_08_01_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
