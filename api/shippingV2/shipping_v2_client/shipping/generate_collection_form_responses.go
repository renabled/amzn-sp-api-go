// Code generated by go-swagger; DO NOT EDIT.

package shipping

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/renabled/amzn-sp-api-go/api/shippingV2/shipping_v2_models"
)

// GenerateCollectionFormReader is a Reader for the GenerateCollectionForm structure.
type GenerateCollectionFormReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GenerateCollectionFormReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGenerateCollectionFormOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGenerateCollectionFormBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGenerateCollectionFormUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGenerateCollectionFormForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGenerateCollectionFormNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGenerateCollectionFormRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGenerateCollectionFormUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGenerateCollectionFormTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGenerateCollectionFormInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGenerateCollectionFormServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGenerateCollectionFormOK creates a GenerateCollectionFormOK with default headers values
func NewGenerateCollectionFormOK() *GenerateCollectionFormOK {
	return &GenerateCollectionFormOK{}
}

/*
GenerateCollectionFormOK describes a response with status code 200, with default header values.

Success.
*/
type GenerateCollectionFormOK struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *shipping_v2_models.GenerateCollectionFormResponse
}

// IsSuccess returns true when this generate collection form o k response has a 2xx status code
func (o *GenerateCollectionFormOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this generate collection form o k response has a 3xx status code
func (o *GenerateCollectionFormOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this generate collection form o k response has a 4xx status code
func (o *GenerateCollectionFormOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this generate collection form o k response has a 5xx status code
func (o *GenerateCollectionFormOK) IsServerError() bool {
	return false
}

// IsCode returns true when this generate collection form o k response a status code equal to that given
func (o *GenerateCollectionFormOK) IsCode(code int) bool {
	return code == 200
}

func (o *GenerateCollectionFormOK) Error() string {
	return fmt.Sprintf("[POST /shipping/v2/collectionForms][%d] generateCollectionFormOK  %+v", 200, o.Payload)
}

func (o *GenerateCollectionFormOK) String() string {
	return fmt.Sprintf("[POST /shipping/v2/collectionForms][%d] generateCollectionFormOK  %+v", 200, o.Payload)
}

func (o *GenerateCollectionFormOK) GetPayload() *shipping_v2_models.GenerateCollectionFormResponse {
	return o.Payload
}

func (o *GenerateCollectionFormOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(shipping_v2_models.GenerateCollectionFormResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGenerateCollectionFormBadRequest creates a GenerateCollectionFormBadRequest with default headers values
func NewGenerateCollectionFormBadRequest() *GenerateCollectionFormBadRequest {
	return &GenerateCollectionFormBadRequest{}
}

/*
GenerateCollectionFormBadRequest describes a response with status code 400, with default header values.

Request has missing or invalid parameters and cannot be parsed.
*/
type GenerateCollectionFormBadRequest struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *shipping_v2_models.ErrorList
}

// IsSuccess returns true when this generate collection form bad request response has a 2xx status code
func (o *GenerateCollectionFormBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this generate collection form bad request response has a 3xx status code
func (o *GenerateCollectionFormBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this generate collection form bad request response has a 4xx status code
func (o *GenerateCollectionFormBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this generate collection form bad request response has a 5xx status code
func (o *GenerateCollectionFormBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this generate collection form bad request response a status code equal to that given
func (o *GenerateCollectionFormBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GenerateCollectionFormBadRequest) Error() string {
	return fmt.Sprintf("[POST /shipping/v2/collectionForms][%d] generateCollectionFormBadRequest  %+v", 400, o.Payload)
}

func (o *GenerateCollectionFormBadRequest) String() string {
	return fmt.Sprintf("[POST /shipping/v2/collectionForms][%d] generateCollectionFormBadRequest  %+v", 400, o.Payload)
}

func (o *GenerateCollectionFormBadRequest) GetPayload() *shipping_v2_models.ErrorList {
	return o.Payload
}

func (o *GenerateCollectionFormBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(shipping_v2_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGenerateCollectionFormUnauthorized creates a GenerateCollectionFormUnauthorized with default headers values
func NewGenerateCollectionFormUnauthorized() *GenerateCollectionFormUnauthorized {
	return &GenerateCollectionFormUnauthorized{}
}

/*
GenerateCollectionFormUnauthorized describes a response with status code 401, with default header values.

The request's Authorization header is not formatted correctly or does not contain a valid token.
*/
type GenerateCollectionFormUnauthorized struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *shipping_v2_models.ErrorList
}

// IsSuccess returns true when this generate collection form unauthorized response has a 2xx status code
func (o *GenerateCollectionFormUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this generate collection form unauthorized response has a 3xx status code
func (o *GenerateCollectionFormUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this generate collection form unauthorized response has a 4xx status code
func (o *GenerateCollectionFormUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this generate collection form unauthorized response has a 5xx status code
func (o *GenerateCollectionFormUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this generate collection form unauthorized response a status code equal to that given
func (o *GenerateCollectionFormUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GenerateCollectionFormUnauthorized) Error() string {
	return fmt.Sprintf("[POST /shipping/v2/collectionForms][%d] generateCollectionFormUnauthorized  %+v", 401, o.Payload)
}

func (o *GenerateCollectionFormUnauthorized) String() string {
	return fmt.Sprintf("[POST /shipping/v2/collectionForms][%d] generateCollectionFormUnauthorized  %+v", 401, o.Payload)
}

func (o *GenerateCollectionFormUnauthorized) GetPayload() *shipping_v2_models.ErrorList {
	return o.Payload
}

func (o *GenerateCollectionFormUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(shipping_v2_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGenerateCollectionFormForbidden creates a GenerateCollectionFormForbidden with default headers values
func NewGenerateCollectionFormForbidden() *GenerateCollectionFormForbidden {
	return &GenerateCollectionFormForbidden{}
}

/*
GenerateCollectionFormForbidden describes a response with status code 403, with default header values.

Indicates that access to the resource is forbidden. Possible reasons include Access Denied, Unauthorized, Expired Token, or Invalid Signature.
*/
type GenerateCollectionFormForbidden struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *shipping_v2_models.ErrorList
}

// IsSuccess returns true when this generate collection form forbidden response has a 2xx status code
func (o *GenerateCollectionFormForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this generate collection form forbidden response has a 3xx status code
func (o *GenerateCollectionFormForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this generate collection form forbidden response has a 4xx status code
func (o *GenerateCollectionFormForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this generate collection form forbidden response has a 5xx status code
func (o *GenerateCollectionFormForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this generate collection form forbidden response a status code equal to that given
func (o *GenerateCollectionFormForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GenerateCollectionFormForbidden) Error() string {
	return fmt.Sprintf("[POST /shipping/v2/collectionForms][%d] generateCollectionFormForbidden  %+v", 403, o.Payload)
}

func (o *GenerateCollectionFormForbidden) String() string {
	return fmt.Sprintf("[POST /shipping/v2/collectionForms][%d] generateCollectionFormForbidden  %+v", 403, o.Payload)
}

func (o *GenerateCollectionFormForbidden) GetPayload() *shipping_v2_models.ErrorList {
	return o.Payload
}

func (o *GenerateCollectionFormForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(shipping_v2_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGenerateCollectionFormNotFound creates a GenerateCollectionFormNotFound with default headers values
func NewGenerateCollectionFormNotFound() *GenerateCollectionFormNotFound {
	return &GenerateCollectionFormNotFound{}
}

/*
GenerateCollectionFormNotFound describes a response with status code 404, with default header values.

The resource specified does not exist.
*/
type GenerateCollectionFormNotFound struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *shipping_v2_models.ErrorList
}

// IsSuccess returns true when this generate collection form not found response has a 2xx status code
func (o *GenerateCollectionFormNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this generate collection form not found response has a 3xx status code
func (o *GenerateCollectionFormNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this generate collection form not found response has a 4xx status code
func (o *GenerateCollectionFormNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this generate collection form not found response has a 5xx status code
func (o *GenerateCollectionFormNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this generate collection form not found response a status code equal to that given
func (o *GenerateCollectionFormNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GenerateCollectionFormNotFound) Error() string {
	return fmt.Sprintf("[POST /shipping/v2/collectionForms][%d] generateCollectionFormNotFound  %+v", 404, o.Payload)
}

func (o *GenerateCollectionFormNotFound) String() string {
	return fmt.Sprintf("[POST /shipping/v2/collectionForms][%d] generateCollectionFormNotFound  %+v", 404, o.Payload)
}

func (o *GenerateCollectionFormNotFound) GetPayload() *shipping_v2_models.ErrorList {
	return o.Payload
}

func (o *GenerateCollectionFormNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(shipping_v2_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGenerateCollectionFormRequestEntityTooLarge creates a GenerateCollectionFormRequestEntityTooLarge with default headers values
func NewGenerateCollectionFormRequestEntityTooLarge() *GenerateCollectionFormRequestEntityTooLarge {
	return &GenerateCollectionFormRequestEntityTooLarge{}
}

/*
GenerateCollectionFormRequestEntityTooLarge describes a response with status code 413, with default header values.

The request size exceeded the maximum accepted size.
*/
type GenerateCollectionFormRequestEntityTooLarge struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *shipping_v2_models.ErrorList
}

// IsSuccess returns true when this generate collection form request entity too large response has a 2xx status code
func (o *GenerateCollectionFormRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this generate collection form request entity too large response has a 3xx status code
func (o *GenerateCollectionFormRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this generate collection form request entity too large response has a 4xx status code
func (o *GenerateCollectionFormRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this generate collection form request entity too large response has a 5xx status code
func (o *GenerateCollectionFormRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this generate collection form request entity too large response a status code equal to that given
func (o *GenerateCollectionFormRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GenerateCollectionFormRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /shipping/v2/collectionForms][%d] generateCollectionFormRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GenerateCollectionFormRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[POST /shipping/v2/collectionForms][%d] generateCollectionFormRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GenerateCollectionFormRequestEntityTooLarge) GetPayload() *shipping_v2_models.ErrorList {
	return o.Payload
}

func (o *GenerateCollectionFormRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(shipping_v2_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGenerateCollectionFormUnsupportedMediaType creates a GenerateCollectionFormUnsupportedMediaType with default headers values
func NewGenerateCollectionFormUnsupportedMediaType() *GenerateCollectionFormUnsupportedMediaType {
	return &GenerateCollectionFormUnsupportedMediaType{}
}

/*
GenerateCollectionFormUnsupportedMediaType describes a response with status code 415, with default header values.

The request payload is in an unsupported format.
*/
type GenerateCollectionFormUnsupportedMediaType struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *shipping_v2_models.ErrorList
}

// IsSuccess returns true when this generate collection form unsupported media type response has a 2xx status code
func (o *GenerateCollectionFormUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this generate collection form unsupported media type response has a 3xx status code
func (o *GenerateCollectionFormUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this generate collection form unsupported media type response has a 4xx status code
func (o *GenerateCollectionFormUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this generate collection form unsupported media type response has a 5xx status code
func (o *GenerateCollectionFormUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this generate collection form unsupported media type response a status code equal to that given
func (o *GenerateCollectionFormUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GenerateCollectionFormUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /shipping/v2/collectionForms][%d] generateCollectionFormUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GenerateCollectionFormUnsupportedMediaType) String() string {
	return fmt.Sprintf("[POST /shipping/v2/collectionForms][%d] generateCollectionFormUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GenerateCollectionFormUnsupportedMediaType) GetPayload() *shipping_v2_models.ErrorList {
	return o.Payload
}

func (o *GenerateCollectionFormUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(shipping_v2_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGenerateCollectionFormTooManyRequests creates a GenerateCollectionFormTooManyRequests with default headers values
func NewGenerateCollectionFormTooManyRequests() *GenerateCollectionFormTooManyRequests {
	return &GenerateCollectionFormTooManyRequests{}
}

/*
GenerateCollectionFormTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type GenerateCollectionFormTooManyRequests struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *shipping_v2_models.ErrorList
}

// IsSuccess returns true when this generate collection form too many requests response has a 2xx status code
func (o *GenerateCollectionFormTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this generate collection form too many requests response has a 3xx status code
func (o *GenerateCollectionFormTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this generate collection form too many requests response has a 4xx status code
func (o *GenerateCollectionFormTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this generate collection form too many requests response has a 5xx status code
func (o *GenerateCollectionFormTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this generate collection form too many requests response a status code equal to that given
func (o *GenerateCollectionFormTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GenerateCollectionFormTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /shipping/v2/collectionForms][%d] generateCollectionFormTooManyRequests  %+v", 429, o.Payload)
}

func (o *GenerateCollectionFormTooManyRequests) String() string {
	return fmt.Sprintf("[POST /shipping/v2/collectionForms][%d] generateCollectionFormTooManyRequests  %+v", 429, o.Payload)
}

func (o *GenerateCollectionFormTooManyRequests) GetPayload() *shipping_v2_models.ErrorList {
	return o.Payload
}

func (o *GenerateCollectionFormTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(shipping_v2_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGenerateCollectionFormInternalServerError creates a GenerateCollectionFormInternalServerError with default headers values
func NewGenerateCollectionFormInternalServerError() *GenerateCollectionFormInternalServerError {
	return &GenerateCollectionFormInternalServerError{}
}

/*
GenerateCollectionFormInternalServerError describes a response with status code 500, with default header values.

An unexpected condition occurred that prevented the server from fulfilling the request.
*/
type GenerateCollectionFormInternalServerError struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *shipping_v2_models.ErrorList
}

// IsSuccess returns true when this generate collection form internal server error response has a 2xx status code
func (o *GenerateCollectionFormInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this generate collection form internal server error response has a 3xx status code
func (o *GenerateCollectionFormInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this generate collection form internal server error response has a 4xx status code
func (o *GenerateCollectionFormInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this generate collection form internal server error response has a 5xx status code
func (o *GenerateCollectionFormInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this generate collection form internal server error response a status code equal to that given
func (o *GenerateCollectionFormInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GenerateCollectionFormInternalServerError) Error() string {
	return fmt.Sprintf("[POST /shipping/v2/collectionForms][%d] generateCollectionFormInternalServerError  %+v", 500, o.Payload)
}

func (o *GenerateCollectionFormInternalServerError) String() string {
	return fmt.Sprintf("[POST /shipping/v2/collectionForms][%d] generateCollectionFormInternalServerError  %+v", 500, o.Payload)
}

func (o *GenerateCollectionFormInternalServerError) GetPayload() *shipping_v2_models.ErrorList {
	return o.Payload
}

func (o *GenerateCollectionFormInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(shipping_v2_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGenerateCollectionFormServiceUnavailable creates a GenerateCollectionFormServiceUnavailable with default headers values
func NewGenerateCollectionFormServiceUnavailable() *GenerateCollectionFormServiceUnavailable {
	return &GenerateCollectionFormServiceUnavailable{}
}

/*
GenerateCollectionFormServiceUnavailable describes a response with status code 503, with default header values.

Temporary overloading or maintenance of the server.
*/
type GenerateCollectionFormServiceUnavailable struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *shipping_v2_models.ErrorList
}

// IsSuccess returns true when this generate collection form service unavailable response has a 2xx status code
func (o *GenerateCollectionFormServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this generate collection form service unavailable response has a 3xx status code
func (o *GenerateCollectionFormServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this generate collection form service unavailable response has a 4xx status code
func (o *GenerateCollectionFormServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this generate collection form service unavailable response has a 5xx status code
func (o *GenerateCollectionFormServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this generate collection form service unavailable response a status code equal to that given
func (o *GenerateCollectionFormServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GenerateCollectionFormServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /shipping/v2/collectionForms][%d] generateCollectionFormServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GenerateCollectionFormServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /shipping/v2/collectionForms][%d] generateCollectionFormServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GenerateCollectionFormServiceUnavailable) GetPayload() *shipping_v2_models.ErrorList {
	return o.Payload
}

func (o *GenerateCollectionFormServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(shipping_v2_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}