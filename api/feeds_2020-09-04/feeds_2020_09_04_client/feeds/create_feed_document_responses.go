// Code generated by go-swagger; DO NOT EDIT.

package feeds

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/xamandar/amzn-sp-api-go/api/feeds_2020-09-04/feeds_2020_09_04_models"
)

// CreateFeedDocumentReader is a Reader for the CreateFeedDocument structure.
type CreateFeedDocumentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateFeedDocumentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateFeedDocumentCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateFeedDocumentBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateFeedDocumentForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateFeedDocumentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewCreateFeedDocumentRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewCreateFeedDocumentUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewCreateFeedDocumentTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateFeedDocumentInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewCreateFeedDocumentServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateFeedDocumentCreated creates a CreateFeedDocumentCreated with default headers values
func NewCreateFeedDocumentCreated() *CreateFeedDocumentCreated {
	return &CreateFeedDocumentCreated{}
}

/* CreateFeedDocumentCreated describes a response with status code 201, with default header values.

Successfully created a feed document that is ready to receive contents.
*/
type CreateFeedDocumentCreated struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference ID.
	 */
	XAmznRequestID string

	Payload *feeds_2020_09_04_models.CreateFeedDocumentResponse
}

func (o *CreateFeedDocumentCreated) Error() string {
	return fmt.Sprintf("[POST /feeds/2020-09-04/documents][%d] createFeedDocumentCreated  %+v", 201, o.Payload)
}
func (o *CreateFeedDocumentCreated) GetPayload() *feeds_2020_09_04_models.CreateFeedDocumentResponse {
	return o.Payload
}

func (o *CreateFeedDocumentCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(feeds_2020_09_04_models.CreateFeedDocumentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateFeedDocumentBadRequest creates a CreateFeedDocumentBadRequest with default headers values
func NewCreateFeedDocumentBadRequest() *CreateFeedDocumentBadRequest {
	return &CreateFeedDocumentBadRequest{}
}

/* CreateFeedDocumentBadRequest describes a response with status code 400, with default header values.

Request has missing or invalid parameters and cannot be parsed.
*/
type CreateFeedDocumentBadRequest struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference ID.
	 */
	XAmznRequestID string

	Payload *feeds_2020_09_04_models.CreateFeedDocumentResponse
}

func (o *CreateFeedDocumentBadRequest) Error() string {
	return fmt.Sprintf("[POST /feeds/2020-09-04/documents][%d] createFeedDocumentBadRequest  %+v", 400, o.Payload)
}
func (o *CreateFeedDocumentBadRequest) GetPayload() *feeds_2020_09_04_models.CreateFeedDocumentResponse {
	return o.Payload
}

func (o *CreateFeedDocumentBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(feeds_2020_09_04_models.CreateFeedDocumentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateFeedDocumentForbidden creates a CreateFeedDocumentForbidden with default headers values
func NewCreateFeedDocumentForbidden() *CreateFeedDocumentForbidden {
	return &CreateFeedDocumentForbidden{}
}

/* CreateFeedDocumentForbidden describes a response with status code 403, with default header values.

Indicates that access to the resource is forbidden. Possible reasons include Access Denied, Unauthorized, Expired Token, or Invalid Signature.
*/
type CreateFeedDocumentForbidden struct {

	/* Unique request reference ID.
	 */
	XAmznRequestID string

	Payload *feeds_2020_09_04_models.CreateFeedDocumentResponse
}

func (o *CreateFeedDocumentForbidden) Error() string {
	return fmt.Sprintf("[POST /feeds/2020-09-04/documents][%d] createFeedDocumentForbidden  %+v", 403, o.Payload)
}
func (o *CreateFeedDocumentForbidden) GetPayload() *feeds_2020_09_04_models.CreateFeedDocumentResponse {
	return o.Payload
}

func (o *CreateFeedDocumentForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(feeds_2020_09_04_models.CreateFeedDocumentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateFeedDocumentNotFound creates a CreateFeedDocumentNotFound with default headers values
func NewCreateFeedDocumentNotFound() *CreateFeedDocumentNotFound {
	return &CreateFeedDocumentNotFound{}
}

/* CreateFeedDocumentNotFound describes a response with status code 404, with default header values.

The resource specified does not exist.
*/
type CreateFeedDocumentNotFound struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference ID.
	 */
	XAmznRequestID string

	Payload *feeds_2020_09_04_models.CreateFeedDocumentResponse
}

func (o *CreateFeedDocumentNotFound) Error() string {
	return fmt.Sprintf("[POST /feeds/2020-09-04/documents][%d] createFeedDocumentNotFound  %+v", 404, o.Payload)
}
func (o *CreateFeedDocumentNotFound) GetPayload() *feeds_2020_09_04_models.CreateFeedDocumentResponse {
	return o.Payload
}

func (o *CreateFeedDocumentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(feeds_2020_09_04_models.CreateFeedDocumentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateFeedDocumentRequestEntityTooLarge creates a CreateFeedDocumentRequestEntityTooLarge with default headers values
func NewCreateFeedDocumentRequestEntityTooLarge() *CreateFeedDocumentRequestEntityTooLarge {
	return &CreateFeedDocumentRequestEntityTooLarge{}
}

/* CreateFeedDocumentRequestEntityTooLarge describes a response with status code 413, with default header values.

The request size exceeded the maximum accepted size.
*/
type CreateFeedDocumentRequestEntityTooLarge struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference ID.
	 */
	XAmznRequestID string

	Payload *feeds_2020_09_04_models.CreateFeedDocumentResponse
}

func (o *CreateFeedDocumentRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /feeds/2020-09-04/documents][%d] createFeedDocumentRequestEntityTooLarge  %+v", 413, o.Payload)
}
func (o *CreateFeedDocumentRequestEntityTooLarge) GetPayload() *feeds_2020_09_04_models.CreateFeedDocumentResponse {
	return o.Payload
}

func (o *CreateFeedDocumentRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(feeds_2020_09_04_models.CreateFeedDocumentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateFeedDocumentUnsupportedMediaType creates a CreateFeedDocumentUnsupportedMediaType with default headers values
func NewCreateFeedDocumentUnsupportedMediaType() *CreateFeedDocumentUnsupportedMediaType {
	return &CreateFeedDocumentUnsupportedMediaType{}
}

/* CreateFeedDocumentUnsupportedMediaType describes a response with status code 415, with default header values.

The request payload is in an unsupported format.
*/
type CreateFeedDocumentUnsupportedMediaType struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference ID.
	 */
	XAmznRequestID string

	Payload *feeds_2020_09_04_models.CreateFeedDocumentResponse
}

func (o *CreateFeedDocumentUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /feeds/2020-09-04/documents][%d] createFeedDocumentUnsupportedMediaType  %+v", 415, o.Payload)
}
func (o *CreateFeedDocumentUnsupportedMediaType) GetPayload() *feeds_2020_09_04_models.CreateFeedDocumentResponse {
	return o.Payload
}

func (o *CreateFeedDocumentUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(feeds_2020_09_04_models.CreateFeedDocumentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateFeedDocumentTooManyRequests creates a CreateFeedDocumentTooManyRequests with default headers values
func NewCreateFeedDocumentTooManyRequests() *CreateFeedDocumentTooManyRequests {
	return &CreateFeedDocumentTooManyRequests{}
}

/* CreateFeedDocumentTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type CreateFeedDocumentTooManyRequests struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference ID.
	 */
	XAmznRequestID string

	Payload *feeds_2020_09_04_models.CreateFeedDocumentResponse
}

func (o *CreateFeedDocumentTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /feeds/2020-09-04/documents][%d] createFeedDocumentTooManyRequests  %+v", 429, o.Payload)
}
func (o *CreateFeedDocumentTooManyRequests) GetPayload() *feeds_2020_09_04_models.CreateFeedDocumentResponse {
	return o.Payload
}

func (o *CreateFeedDocumentTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(feeds_2020_09_04_models.CreateFeedDocumentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateFeedDocumentInternalServerError creates a CreateFeedDocumentInternalServerError with default headers values
func NewCreateFeedDocumentInternalServerError() *CreateFeedDocumentInternalServerError {
	return &CreateFeedDocumentInternalServerError{}
}

/* CreateFeedDocumentInternalServerError describes a response with status code 500, with default header values.

An unexpected condition occurred that prevented the server from fulfilling the request.
*/
type CreateFeedDocumentInternalServerError struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference ID.
	 */
	XAmznRequestID string

	Payload *feeds_2020_09_04_models.CreateFeedDocumentResponse
}

func (o *CreateFeedDocumentInternalServerError) Error() string {
	return fmt.Sprintf("[POST /feeds/2020-09-04/documents][%d] createFeedDocumentInternalServerError  %+v", 500, o.Payload)
}
func (o *CreateFeedDocumentInternalServerError) GetPayload() *feeds_2020_09_04_models.CreateFeedDocumentResponse {
	return o.Payload
}

func (o *CreateFeedDocumentInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(feeds_2020_09_04_models.CreateFeedDocumentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateFeedDocumentServiceUnavailable creates a CreateFeedDocumentServiceUnavailable with default headers values
func NewCreateFeedDocumentServiceUnavailable() *CreateFeedDocumentServiceUnavailable {
	return &CreateFeedDocumentServiceUnavailable{}
}

/* CreateFeedDocumentServiceUnavailable describes a response with status code 503, with default header values.

Temporary overloading or maintenance of the server.
*/
type CreateFeedDocumentServiceUnavailable struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference ID.
	 */
	XAmznRequestID string

	Payload *feeds_2020_09_04_models.CreateFeedDocumentResponse
}

func (o *CreateFeedDocumentServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /feeds/2020-09-04/documents][%d] createFeedDocumentServiceUnavailable  %+v", 503, o.Payload)
}
func (o *CreateFeedDocumentServiceUnavailable) GetPayload() *feeds_2020_09_04_models.CreateFeedDocumentResponse {
	return o.Payload
}

func (o *CreateFeedDocumentServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(feeds_2020_09_04_models.CreateFeedDocumentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
