// Code generated by go-swagger; DO NOT EDIT.

package small_and_light

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/xamandar/amzn-sp-api-go/api/fbaSmallandLight/fba_smalland_light_models"
)

// DeleteSmallAndLightEnrollmentBySellerSKUReader is a Reader for the DeleteSmallAndLightEnrollmentBySellerSKU structure.
type DeleteSmallAndLightEnrollmentBySellerSKUReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteSmallAndLightEnrollmentBySellerSKUReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteSmallAndLightEnrollmentBySellerSKUNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteSmallAndLightEnrollmentBySellerSKUBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteSmallAndLightEnrollmentBySellerSKUForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteSmallAndLightEnrollmentBySellerSKUNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewDeleteSmallAndLightEnrollmentBySellerSKURequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewDeleteSmallAndLightEnrollmentBySellerSKUUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteSmallAndLightEnrollmentBySellerSKUTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteSmallAndLightEnrollmentBySellerSKUInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteSmallAndLightEnrollmentBySellerSKUServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteSmallAndLightEnrollmentBySellerSKUNoContent creates a DeleteSmallAndLightEnrollmentBySellerSKUNoContent with default headers values
func NewDeleteSmallAndLightEnrollmentBySellerSKUNoContent() *DeleteSmallAndLightEnrollmentBySellerSKUNoContent {
	return &DeleteSmallAndLightEnrollmentBySellerSKUNoContent{}
}

/*
DeleteSmallAndLightEnrollmentBySellerSKUNoContent describes a response with status code 204, with default header values.

Success.
*/
type DeleteSmallAndLightEnrollmentBySellerSKUNoContent struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string
}

// IsSuccess returns true when this delete small and light enrollment by seller s k u no content response has a 2xx status code
func (o *DeleteSmallAndLightEnrollmentBySellerSKUNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete small and light enrollment by seller s k u no content response has a 3xx status code
func (o *DeleteSmallAndLightEnrollmentBySellerSKUNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete small and light enrollment by seller s k u no content response has a 4xx status code
func (o *DeleteSmallAndLightEnrollmentBySellerSKUNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete small and light enrollment by seller s k u no content response has a 5xx status code
func (o *DeleteSmallAndLightEnrollmentBySellerSKUNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete small and light enrollment by seller s k u no content response a status code equal to that given
func (o *DeleteSmallAndLightEnrollmentBySellerSKUNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *DeleteSmallAndLightEnrollmentBySellerSKUNoContent) Error() string {
	return fmt.Sprintf("[DELETE /fba/smallAndLight/v1/enrollments/{sellerSKU}][%d] deleteSmallAndLightEnrollmentBySellerSKUNoContent ", 204)
}

func (o *DeleteSmallAndLightEnrollmentBySellerSKUNoContent) String() string {
	return fmt.Sprintf("[DELETE /fba/smallAndLight/v1/enrollments/{sellerSKU}][%d] deleteSmallAndLightEnrollmentBySellerSKUNoContent ", 204)
}

func (o *DeleteSmallAndLightEnrollmentBySellerSKUNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	return nil
}

// NewDeleteSmallAndLightEnrollmentBySellerSKUBadRequest creates a DeleteSmallAndLightEnrollmentBySellerSKUBadRequest with default headers values
func NewDeleteSmallAndLightEnrollmentBySellerSKUBadRequest() *DeleteSmallAndLightEnrollmentBySellerSKUBadRequest {
	return &DeleteSmallAndLightEnrollmentBySellerSKUBadRequest{}
}

/*
DeleteSmallAndLightEnrollmentBySellerSKUBadRequest describes a response with status code 400, with default header values.

Request has missing or invalid parameters and cannot be parsed.
*/
type DeleteSmallAndLightEnrollmentBySellerSKUBadRequest struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fba_smalland_light_models.ErrorList
}

// IsSuccess returns true when this delete small and light enrollment by seller s k u bad request response has a 2xx status code
func (o *DeleteSmallAndLightEnrollmentBySellerSKUBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete small and light enrollment by seller s k u bad request response has a 3xx status code
func (o *DeleteSmallAndLightEnrollmentBySellerSKUBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete small and light enrollment by seller s k u bad request response has a 4xx status code
func (o *DeleteSmallAndLightEnrollmentBySellerSKUBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete small and light enrollment by seller s k u bad request response has a 5xx status code
func (o *DeleteSmallAndLightEnrollmentBySellerSKUBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete small and light enrollment by seller s k u bad request response a status code equal to that given
func (o *DeleteSmallAndLightEnrollmentBySellerSKUBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DeleteSmallAndLightEnrollmentBySellerSKUBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /fba/smallAndLight/v1/enrollments/{sellerSKU}][%d] deleteSmallAndLightEnrollmentBySellerSKUBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteSmallAndLightEnrollmentBySellerSKUBadRequest) String() string {
	return fmt.Sprintf("[DELETE /fba/smallAndLight/v1/enrollments/{sellerSKU}][%d] deleteSmallAndLightEnrollmentBySellerSKUBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteSmallAndLightEnrollmentBySellerSKUBadRequest) GetPayload() *fba_smalland_light_models.ErrorList {
	return o.Payload
}

func (o *DeleteSmallAndLightEnrollmentBySellerSKUBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fba_smalland_light_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteSmallAndLightEnrollmentBySellerSKUForbidden creates a DeleteSmallAndLightEnrollmentBySellerSKUForbidden with default headers values
func NewDeleteSmallAndLightEnrollmentBySellerSKUForbidden() *DeleteSmallAndLightEnrollmentBySellerSKUForbidden {
	return &DeleteSmallAndLightEnrollmentBySellerSKUForbidden{}
}

/*
DeleteSmallAndLightEnrollmentBySellerSKUForbidden describes a response with status code 403, with default header values.

Indicates that access to the resource is forbidden. Possible reasons include Access Denied, Unauthorized, Expired Token, or Invalid Signature.
*/
type DeleteSmallAndLightEnrollmentBySellerSKUForbidden struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fba_smalland_light_models.ErrorList
}

// IsSuccess returns true when this delete small and light enrollment by seller s k u forbidden response has a 2xx status code
func (o *DeleteSmallAndLightEnrollmentBySellerSKUForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete small and light enrollment by seller s k u forbidden response has a 3xx status code
func (o *DeleteSmallAndLightEnrollmentBySellerSKUForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete small and light enrollment by seller s k u forbidden response has a 4xx status code
func (o *DeleteSmallAndLightEnrollmentBySellerSKUForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete small and light enrollment by seller s k u forbidden response has a 5xx status code
func (o *DeleteSmallAndLightEnrollmentBySellerSKUForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete small and light enrollment by seller s k u forbidden response a status code equal to that given
func (o *DeleteSmallAndLightEnrollmentBySellerSKUForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *DeleteSmallAndLightEnrollmentBySellerSKUForbidden) Error() string {
	return fmt.Sprintf("[DELETE /fba/smallAndLight/v1/enrollments/{sellerSKU}][%d] deleteSmallAndLightEnrollmentBySellerSKUForbidden  %+v", 403, o.Payload)
}

func (o *DeleteSmallAndLightEnrollmentBySellerSKUForbidden) String() string {
	return fmt.Sprintf("[DELETE /fba/smallAndLight/v1/enrollments/{sellerSKU}][%d] deleteSmallAndLightEnrollmentBySellerSKUForbidden  %+v", 403, o.Payload)
}

func (o *DeleteSmallAndLightEnrollmentBySellerSKUForbidden) GetPayload() *fba_smalland_light_models.ErrorList {
	return o.Payload
}

func (o *DeleteSmallAndLightEnrollmentBySellerSKUForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(fba_smalland_light_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteSmallAndLightEnrollmentBySellerSKUNotFound creates a DeleteSmallAndLightEnrollmentBySellerSKUNotFound with default headers values
func NewDeleteSmallAndLightEnrollmentBySellerSKUNotFound() *DeleteSmallAndLightEnrollmentBySellerSKUNotFound {
	return &DeleteSmallAndLightEnrollmentBySellerSKUNotFound{}
}

/*
DeleteSmallAndLightEnrollmentBySellerSKUNotFound describes a response with status code 404, with default header values.

The resource specified does not exist.
*/
type DeleteSmallAndLightEnrollmentBySellerSKUNotFound struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fba_smalland_light_models.ErrorList
}

// IsSuccess returns true when this delete small and light enrollment by seller s k u not found response has a 2xx status code
func (o *DeleteSmallAndLightEnrollmentBySellerSKUNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete small and light enrollment by seller s k u not found response has a 3xx status code
func (o *DeleteSmallAndLightEnrollmentBySellerSKUNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete small and light enrollment by seller s k u not found response has a 4xx status code
func (o *DeleteSmallAndLightEnrollmentBySellerSKUNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete small and light enrollment by seller s k u not found response has a 5xx status code
func (o *DeleteSmallAndLightEnrollmentBySellerSKUNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete small and light enrollment by seller s k u not found response a status code equal to that given
func (o *DeleteSmallAndLightEnrollmentBySellerSKUNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *DeleteSmallAndLightEnrollmentBySellerSKUNotFound) Error() string {
	return fmt.Sprintf("[DELETE /fba/smallAndLight/v1/enrollments/{sellerSKU}][%d] deleteSmallAndLightEnrollmentBySellerSKUNotFound  %+v", 404, o.Payload)
}

func (o *DeleteSmallAndLightEnrollmentBySellerSKUNotFound) String() string {
	return fmt.Sprintf("[DELETE /fba/smallAndLight/v1/enrollments/{sellerSKU}][%d] deleteSmallAndLightEnrollmentBySellerSKUNotFound  %+v", 404, o.Payload)
}

func (o *DeleteSmallAndLightEnrollmentBySellerSKUNotFound) GetPayload() *fba_smalland_light_models.ErrorList {
	return o.Payload
}

func (o *DeleteSmallAndLightEnrollmentBySellerSKUNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fba_smalland_light_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteSmallAndLightEnrollmentBySellerSKURequestEntityTooLarge creates a DeleteSmallAndLightEnrollmentBySellerSKURequestEntityTooLarge with default headers values
func NewDeleteSmallAndLightEnrollmentBySellerSKURequestEntityTooLarge() *DeleteSmallAndLightEnrollmentBySellerSKURequestEntityTooLarge {
	return &DeleteSmallAndLightEnrollmentBySellerSKURequestEntityTooLarge{}
}

/*
DeleteSmallAndLightEnrollmentBySellerSKURequestEntityTooLarge describes a response with status code 413, with default header values.

The request size exceeded the maximum accepted size.
*/
type DeleteSmallAndLightEnrollmentBySellerSKURequestEntityTooLarge struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fba_smalland_light_models.ErrorList
}

// IsSuccess returns true when this delete small and light enrollment by seller s k u request entity too large response has a 2xx status code
func (o *DeleteSmallAndLightEnrollmentBySellerSKURequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete small and light enrollment by seller s k u request entity too large response has a 3xx status code
func (o *DeleteSmallAndLightEnrollmentBySellerSKURequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete small and light enrollment by seller s k u request entity too large response has a 4xx status code
func (o *DeleteSmallAndLightEnrollmentBySellerSKURequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete small and light enrollment by seller s k u request entity too large response has a 5xx status code
func (o *DeleteSmallAndLightEnrollmentBySellerSKURequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this delete small and light enrollment by seller s k u request entity too large response a status code equal to that given
func (o *DeleteSmallAndLightEnrollmentBySellerSKURequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *DeleteSmallAndLightEnrollmentBySellerSKURequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /fba/smallAndLight/v1/enrollments/{sellerSKU}][%d] deleteSmallAndLightEnrollmentBySellerSKURequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteSmallAndLightEnrollmentBySellerSKURequestEntityTooLarge) String() string {
	return fmt.Sprintf("[DELETE /fba/smallAndLight/v1/enrollments/{sellerSKU}][%d] deleteSmallAndLightEnrollmentBySellerSKURequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteSmallAndLightEnrollmentBySellerSKURequestEntityTooLarge) GetPayload() *fba_smalland_light_models.ErrorList {
	return o.Payload
}

func (o *DeleteSmallAndLightEnrollmentBySellerSKURequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fba_smalland_light_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteSmallAndLightEnrollmentBySellerSKUUnsupportedMediaType creates a DeleteSmallAndLightEnrollmentBySellerSKUUnsupportedMediaType with default headers values
func NewDeleteSmallAndLightEnrollmentBySellerSKUUnsupportedMediaType() *DeleteSmallAndLightEnrollmentBySellerSKUUnsupportedMediaType {
	return &DeleteSmallAndLightEnrollmentBySellerSKUUnsupportedMediaType{}
}

/*
DeleteSmallAndLightEnrollmentBySellerSKUUnsupportedMediaType describes a response with status code 415, with default header values.

The request payload is in an unsupported format.
*/
type DeleteSmallAndLightEnrollmentBySellerSKUUnsupportedMediaType struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fba_smalland_light_models.ErrorList
}

// IsSuccess returns true when this delete small and light enrollment by seller s k u unsupported media type response has a 2xx status code
func (o *DeleteSmallAndLightEnrollmentBySellerSKUUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete small and light enrollment by seller s k u unsupported media type response has a 3xx status code
func (o *DeleteSmallAndLightEnrollmentBySellerSKUUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete small and light enrollment by seller s k u unsupported media type response has a 4xx status code
func (o *DeleteSmallAndLightEnrollmentBySellerSKUUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete small and light enrollment by seller s k u unsupported media type response has a 5xx status code
func (o *DeleteSmallAndLightEnrollmentBySellerSKUUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this delete small and light enrollment by seller s k u unsupported media type response a status code equal to that given
func (o *DeleteSmallAndLightEnrollmentBySellerSKUUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *DeleteSmallAndLightEnrollmentBySellerSKUUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /fba/smallAndLight/v1/enrollments/{sellerSKU}][%d] deleteSmallAndLightEnrollmentBySellerSKUUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteSmallAndLightEnrollmentBySellerSKUUnsupportedMediaType) String() string {
	return fmt.Sprintf("[DELETE /fba/smallAndLight/v1/enrollments/{sellerSKU}][%d] deleteSmallAndLightEnrollmentBySellerSKUUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteSmallAndLightEnrollmentBySellerSKUUnsupportedMediaType) GetPayload() *fba_smalland_light_models.ErrorList {
	return o.Payload
}

func (o *DeleteSmallAndLightEnrollmentBySellerSKUUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fba_smalland_light_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteSmallAndLightEnrollmentBySellerSKUTooManyRequests creates a DeleteSmallAndLightEnrollmentBySellerSKUTooManyRequests with default headers values
func NewDeleteSmallAndLightEnrollmentBySellerSKUTooManyRequests() *DeleteSmallAndLightEnrollmentBySellerSKUTooManyRequests {
	return &DeleteSmallAndLightEnrollmentBySellerSKUTooManyRequests{}
}

/*
DeleteSmallAndLightEnrollmentBySellerSKUTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type DeleteSmallAndLightEnrollmentBySellerSKUTooManyRequests struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fba_smalland_light_models.ErrorList
}

// IsSuccess returns true when this delete small and light enrollment by seller s k u too many requests response has a 2xx status code
func (o *DeleteSmallAndLightEnrollmentBySellerSKUTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete small and light enrollment by seller s k u too many requests response has a 3xx status code
func (o *DeleteSmallAndLightEnrollmentBySellerSKUTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete small and light enrollment by seller s k u too many requests response has a 4xx status code
func (o *DeleteSmallAndLightEnrollmentBySellerSKUTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete small and light enrollment by seller s k u too many requests response has a 5xx status code
func (o *DeleteSmallAndLightEnrollmentBySellerSKUTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this delete small and light enrollment by seller s k u too many requests response a status code equal to that given
func (o *DeleteSmallAndLightEnrollmentBySellerSKUTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *DeleteSmallAndLightEnrollmentBySellerSKUTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /fba/smallAndLight/v1/enrollments/{sellerSKU}][%d] deleteSmallAndLightEnrollmentBySellerSKUTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteSmallAndLightEnrollmentBySellerSKUTooManyRequests) String() string {
	return fmt.Sprintf("[DELETE /fba/smallAndLight/v1/enrollments/{sellerSKU}][%d] deleteSmallAndLightEnrollmentBySellerSKUTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteSmallAndLightEnrollmentBySellerSKUTooManyRequests) GetPayload() *fba_smalland_light_models.ErrorList {
	return o.Payload
}

func (o *DeleteSmallAndLightEnrollmentBySellerSKUTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fba_smalland_light_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteSmallAndLightEnrollmentBySellerSKUInternalServerError creates a DeleteSmallAndLightEnrollmentBySellerSKUInternalServerError with default headers values
func NewDeleteSmallAndLightEnrollmentBySellerSKUInternalServerError() *DeleteSmallAndLightEnrollmentBySellerSKUInternalServerError {
	return &DeleteSmallAndLightEnrollmentBySellerSKUInternalServerError{}
}

/*
DeleteSmallAndLightEnrollmentBySellerSKUInternalServerError describes a response with status code 500, with default header values.

An unexpected condition occurred that prevented the server from fulfilling the request.
*/
type DeleteSmallAndLightEnrollmentBySellerSKUInternalServerError struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fba_smalland_light_models.ErrorList
}

// IsSuccess returns true when this delete small and light enrollment by seller s k u internal server error response has a 2xx status code
func (o *DeleteSmallAndLightEnrollmentBySellerSKUInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete small and light enrollment by seller s k u internal server error response has a 3xx status code
func (o *DeleteSmallAndLightEnrollmentBySellerSKUInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete small and light enrollment by seller s k u internal server error response has a 4xx status code
func (o *DeleteSmallAndLightEnrollmentBySellerSKUInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete small and light enrollment by seller s k u internal server error response has a 5xx status code
func (o *DeleteSmallAndLightEnrollmentBySellerSKUInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete small and light enrollment by seller s k u internal server error response a status code equal to that given
func (o *DeleteSmallAndLightEnrollmentBySellerSKUInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *DeleteSmallAndLightEnrollmentBySellerSKUInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /fba/smallAndLight/v1/enrollments/{sellerSKU}][%d] deleteSmallAndLightEnrollmentBySellerSKUInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteSmallAndLightEnrollmentBySellerSKUInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /fba/smallAndLight/v1/enrollments/{sellerSKU}][%d] deleteSmallAndLightEnrollmentBySellerSKUInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteSmallAndLightEnrollmentBySellerSKUInternalServerError) GetPayload() *fba_smalland_light_models.ErrorList {
	return o.Payload
}

func (o *DeleteSmallAndLightEnrollmentBySellerSKUInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fba_smalland_light_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteSmallAndLightEnrollmentBySellerSKUServiceUnavailable creates a DeleteSmallAndLightEnrollmentBySellerSKUServiceUnavailable with default headers values
func NewDeleteSmallAndLightEnrollmentBySellerSKUServiceUnavailable() *DeleteSmallAndLightEnrollmentBySellerSKUServiceUnavailable {
	return &DeleteSmallAndLightEnrollmentBySellerSKUServiceUnavailable{}
}

/*
DeleteSmallAndLightEnrollmentBySellerSKUServiceUnavailable describes a response with status code 503, with default header values.

Temporary overloading or maintenance of the server.
*/
type DeleteSmallAndLightEnrollmentBySellerSKUServiceUnavailable struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fba_smalland_light_models.ErrorList
}

// IsSuccess returns true when this delete small and light enrollment by seller s k u service unavailable response has a 2xx status code
func (o *DeleteSmallAndLightEnrollmentBySellerSKUServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete small and light enrollment by seller s k u service unavailable response has a 3xx status code
func (o *DeleteSmallAndLightEnrollmentBySellerSKUServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete small and light enrollment by seller s k u service unavailable response has a 4xx status code
func (o *DeleteSmallAndLightEnrollmentBySellerSKUServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete small and light enrollment by seller s k u service unavailable response has a 5xx status code
func (o *DeleteSmallAndLightEnrollmentBySellerSKUServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this delete small and light enrollment by seller s k u service unavailable response a status code equal to that given
func (o *DeleteSmallAndLightEnrollmentBySellerSKUServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *DeleteSmallAndLightEnrollmentBySellerSKUServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /fba/smallAndLight/v1/enrollments/{sellerSKU}][%d] deleteSmallAndLightEnrollmentBySellerSKUServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteSmallAndLightEnrollmentBySellerSKUServiceUnavailable) String() string {
	return fmt.Sprintf("[DELETE /fba/smallAndLight/v1/enrollments/{sellerSKU}][%d] deleteSmallAndLightEnrollmentBySellerSKUServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteSmallAndLightEnrollmentBySellerSKUServiceUnavailable) GetPayload() *fba_smalland_light_models.ErrorList {
	return o.Payload
}

func (o *DeleteSmallAndLightEnrollmentBySellerSKUServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fba_smalland_light_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
