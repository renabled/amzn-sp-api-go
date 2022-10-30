// Code generated by go-swagger; DO NOT EDIT.

package listings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/xamandar/amzn-sp-api-go/api/listingsItems_2021-08-01/listings_items_2021_08_01_models"
)

// PatchListingsItemReader is a Reader for the PatchListingsItem structure.
type PatchListingsItemReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchListingsItemReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchListingsItemOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchListingsItemBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchListingsItemForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPatchListingsItemRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPatchListingsItemUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPatchListingsItemTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPatchListingsItemInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPatchListingsItemServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchListingsItemOK creates a PatchListingsItemOK with default headers values
func NewPatchListingsItemOK() *PatchListingsItemOK {
	return &PatchListingsItemOK{}
}

/*
PatchListingsItemOK describes a response with status code 200, with default header values.

Successfully understood the listings item patch request. See the response to determine if the submission was accepted.
*/
type PatchListingsItemOK struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *listings_items_2021_08_01_models.ListingsItemSubmissionResponse
}

// IsSuccess returns true when this patch listings item o k response has a 2xx status code
func (o *PatchListingsItemOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch listings item o k response has a 3xx status code
func (o *PatchListingsItemOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch listings item o k response has a 4xx status code
func (o *PatchListingsItemOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch listings item o k response has a 5xx status code
func (o *PatchListingsItemOK) IsServerError() bool {
	return false
}

// IsCode returns true when this patch listings item o k response a status code equal to that given
func (o *PatchListingsItemOK) IsCode(code int) bool {
	return code == 200
}

func (o *PatchListingsItemOK) Error() string {
	return fmt.Sprintf("[PATCH /listings/2021-08-01/items/{sellerId}/{sku}][%d] patchListingsItemOK  %+v", 200, o.Payload)
}

func (o *PatchListingsItemOK) String() string {
	return fmt.Sprintf("[PATCH /listings/2021-08-01/items/{sellerId}/{sku}][%d] patchListingsItemOK  %+v", 200, o.Payload)
}

func (o *PatchListingsItemOK) GetPayload() *listings_items_2021_08_01_models.ListingsItemSubmissionResponse {
	return o.Payload
}

func (o *PatchListingsItemOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewPatchListingsItemBadRequest creates a PatchListingsItemBadRequest with default headers values
func NewPatchListingsItemBadRequest() *PatchListingsItemBadRequest {
	return &PatchListingsItemBadRequest{}
}

/*
PatchListingsItemBadRequest describes a response with status code 400, with default header values.

Request has missing or invalid parameters and cannot be parsed.
*/
type PatchListingsItemBadRequest struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *listings_items_2021_08_01_models.ErrorList
}

// IsSuccess returns true when this patch listings item bad request response has a 2xx status code
func (o *PatchListingsItemBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch listings item bad request response has a 3xx status code
func (o *PatchListingsItemBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch listings item bad request response has a 4xx status code
func (o *PatchListingsItemBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch listings item bad request response has a 5xx status code
func (o *PatchListingsItemBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this patch listings item bad request response a status code equal to that given
func (o *PatchListingsItemBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PatchListingsItemBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /listings/2021-08-01/items/{sellerId}/{sku}][%d] patchListingsItemBadRequest  %+v", 400, o.Payload)
}

func (o *PatchListingsItemBadRequest) String() string {
	return fmt.Sprintf("[PATCH /listings/2021-08-01/items/{sellerId}/{sku}][%d] patchListingsItemBadRequest  %+v", 400, o.Payload)
}

func (o *PatchListingsItemBadRequest) GetPayload() *listings_items_2021_08_01_models.ErrorList {
	return o.Payload
}

func (o *PatchListingsItemBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewPatchListingsItemForbidden creates a PatchListingsItemForbidden with default headers values
func NewPatchListingsItemForbidden() *PatchListingsItemForbidden {
	return &PatchListingsItemForbidden{}
}

/*
PatchListingsItemForbidden describes a response with status code 403, with default header values.

Indicates that access to the resource is forbidden. Possible reasons include Access Denied, Unauthorized, Expired Token, or Invalid Signature.
*/
type PatchListingsItemForbidden struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *listings_items_2021_08_01_models.ErrorList
}

// IsSuccess returns true when this patch listings item forbidden response has a 2xx status code
func (o *PatchListingsItemForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch listings item forbidden response has a 3xx status code
func (o *PatchListingsItemForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch listings item forbidden response has a 4xx status code
func (o *PatchListingsItemForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch listings item forbidden response has a 5xx status code
func (o *PatchListingsItemForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this patch listings item forbidden response a status code equal to that given
func (o *PatchListingsItemForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PatchListingsItemForbidden) Error() string {
	return fmt.Sprintf("[PATCH /listings/2021-08-01/items/{sellerId}/{sku}][%d] patchListingsItemForbidden  %+v", 403, o.Payload)
}

func (o *PatchListingsItemForbidden) String() string {
	return fmt.Sprintf("[PATCH /listings/2021-08-01/items/{sellerId}/{sku}][%d] patchListingsItemForbidden  %+v", 403, o.Payload)
}

func (o *PatchListingsItemForbidden) GetPayload() *listings_items_2021_08_01_models.ErrorList {
	return o.Payload
}

func (o *PatchListingsItemForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewPatchListingsItemRequestEntityTooLarge creates a PatchListingsItemRequestEntityTooLarge with default headers values
func NewPatchListingsItemRequestEntityTooLarge() *PatchListingsItemRequestEntityTooLarge {
	return &PatchListingsItemRequestEntityTooLarge{}
}

/*
PatchListingsItemRequestEntityTooLarge describes a response with status code 413, with default header values.

The request size exceeded the maximum accepted size.
*/
type PatchListingsItemRequestEntityTooLarge struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *listings_items_2021_08_01_models.ErrorList
}

// IsSuccess returns true when this patch listings item request entity too large response has a 2xx status code
func (o *PatchListingsItemRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch listings item request entity too large response has a 3xx status code
func (o *PatchListingsItemRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch listings item request entity too large response has a 4xx status code
func (o *PatchListingsItemRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch listings item request entity too large response has a 5xx status code
func (o *PatchListingsItemRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this patch listings item request entity too large response a status code equal to that given
func (o *PatchListingsItemRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PatchListingsItemRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PATCH /listings/2021-08-01/items/{sellerId}/{sku}][%d] patchListingsItemRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PatchListingsItemRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[PATCH /listings/2021-08-01/items/{sellerId}/{sku}][%d] patchListingsItemRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PatchListingsItemRequestEntityTooLarge) GetPayload() *listings_items_2021_08_01_models.ErrorList {
	return o.Payload
}

func (o *PatchListingsItemRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewPatchListingsItemUnsupportedMediaType creates a PatchListingsItemUnsupportedMediaType with default headers values
func NewPatchListingsItemUnsupportedMediaType() *PatchListingsItemUnsupportedMediaType {
	return &PatchListingsItemUnsupportedMediaType{}
}

/*
PatchListingsItemUnsupportedMediaType describes a response with status code 415, with default header values.

The request payload is in an unsupported format.
*/
type PatchListingsItemUnsupportedMediaType struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *listings_items_2021_08_01_models.ErrorList
}

// IsSuccess returns true when this patch listings item unsupported media type response has a 2xx status code
func (o *PatchListingsItemUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch listings item unsupported media type response has a 3xx status code
func (o *PatchListingsItemUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch listings item unsupported media type response has a 4xx status code
func (o *PatchListingsItemUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch listings item unsupported media type response has a 5xx status code
func (o *PatchListingsItemUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this patch listings item unsupported media type response a status code equal to that given
func (o *PatchListingsItemUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PatchListingsItemUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PATCH /listings/2021-08-01/items/{sellerId}/{sku}][%d] patchListingsItemUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PatchListingsItemUnsupportedMediaType) String() string {
	return fmt.Sprintf("[PATCH /listings/2021-08-01/items/{sellerId}/{sku}][%d] patchListingsItemUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PatchListingsItemUnsupportedMediaType) GetPayload() *listings_items_2021_08_01_models.ErrorList {
	return o.Payload
}

func (o *PatchListingsItemUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewPatchListingsItemTooManyRequests creates a PatchListingsItemTooManyRequests with default headers values
func NewPatchListingsItemTooManyRequests() *PatchListingsItemTooManyRequests {
	return &PatchListingsItemTooManyRequests{}
}

/*
PatchListingsItemTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type PatchListingsItemTooManyRequests struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *listings_items_2021_08_01_models.ErrorList
}

// IsSuccess returns true when this patch listings item too many requests response has a 2xx status code
func (o *PatchListingsItemTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch listings item too many requests response has a 3xx status code
func (o *PatchListingsItemTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch listings item too many requests response has a 4xx status code
func (o *PatchListingsItemTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch listings item too many requests response has a 5xx status code
func (o *PatchListingsItemTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this patch listings item too many requests response a status code equal to that given
func (o *PatchListingsItemTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PatchListingsItemTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /listings/2021-08-01/items/{sellerId}/{sku}][%d] patchListingsItemTooManyRequests  %+v", 429, o.Payload)
}

func (o *PatchListingsItemTooManyRequests) String() string {
	return fmt.Sprintf("[PATCH /listings/2021-08-01/items/{sellerId}/{sku}][%d] patchListingsItemTooManyRequests  %+v", 429, o.Payload)
}

func (o *PatchListingsItemTooManyRequests) GetPayload() *listings_items_2021_08_01_models.ErrorList {
	return o.Payload
}

func (o *PatchListingsItemTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewPatchListingsItemInternalServerError creates a PatchListingsItemInternalServerError with default headers values
func NewPatchListingsItemInternalServerError() *PatchListingsItemInternalServerError {
	return &PatchListingsItemInternalServerError{}
}

/*
PatchListingsItemInternalServerError describes a response with status code 500, with default header values.

An unexpected condition occurred that prevented the server from fulfilling the request.
*/
type PatchListingsItemInternalServerError struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *listings_items_2021_08_01_models.ErrorList
}

// IsSuccess returns true when this patch listings item internal server error response has a 2xx status code
func (o *PatchListingsItemInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch listings item internal server error response has a 3xx status code
func (o *PatchListingsItemInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch listings item internal server error response has a 4xx status code
func (o *PatchListingsItemInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch listings item internal server error response has a 5xx status code
func (o *PatchListingsItemInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this patch listings item internal server error response a status code equal to that given
func (o *PatchListingsItemInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PatchListingsItemInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /listings/2021-08-01/items/{sellerId}/{sku}][%d] patchListingsItemInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchListingsItemInternalServerError) String() string {
	return fmt.Sprintf("[PATCH /listings/2021-08-01/items/{sellerId}/{sku}][%d] patchListingsItemInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchListingsItemInternalServerError) GetPayload() *listings_items_2021_08_01_models.ErrorList {
	return o.Payload
}

func (o *PatchListingsItemInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewPatchListingsItemServiceUnavailable creates a PatchListingsItemServiceUnavailable with default headers values
func NewPatchListingsItemServiceUnavailable() *PatchListingsItemServiceUnavailable {
	return &PatchListingsItemServiceUnavailable{}
}

/*
PatchListingsItemServiceUnavailable describes a response with status code 503, with default header values.

Temporary overloading or maintenance of the server.
*/
type PatchListingsItemServiceUnavailable struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *listings_items_2021_08_01_models.ErrorList
}

// IsSuccess returns true when this patch listings item service unavailable response has a 2xx status code
func (o *PatchListingsItemServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch listings item service unavailable response has a 3xx status code
func (o *PatchListingsItemServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch listings item service unavailable response has a 4xx status code
func (o *PatchListingsItemServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch listings item service unavailable response has a 5xx status code
func (o *PatchListingsItemServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this patch listings item service unavailable response a status code equal to that given
func (o *PatchListingsItemServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PatchListingsItemServiceUnavailable) Error() string {
	return fmt.Sprintf("[PATCH /listings/2021-08-01/items/{sellerId}/{sku}][%d] patchListingsItemServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PatchListingsItemServiceUnavailable) String() string {
	return fmt.Sprintf("[PATCH /listings/2021-08-01/items/{sellerId}/{sku}][%d] patchListingsItemServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PatchListingsItemServiceUnavailable) GetPayload() *listings_items_2021_08_01_models.ErrorList {
	return o.Payload
}

func (o *PatchListingsItemServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
