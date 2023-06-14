// Code generated by go-swagger; DO NOT EDIT.

package messaging

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/renabled/amzn-sp-api-go/api/messaging/messaging_models"
)

// CreateLegalDisclosureReader is a Reader for the CreateLegalDisclosure structure.
type CreateLegalDisclosureReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateLegalDisclosureReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateLegalDisclosureCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateLegalDisclosureBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateLegalDisclosureForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateLegalDisclosureNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewCreateLegalDisclosureRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewCreateLegalDisclosureUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewCreateLegalDisclosureTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateLegalDisclosureInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewCreateLegalDisclosureServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateLegalDisclosureCreated creates a CreateLegalDisclosureCreated with default headers values
func NewCreateLegalDisclosureCreated() *CreateLegalDisclosureCreated {
	return &CreateLegalDisclosureCreated{}
}

/*
CreateLegalDisclosureCreated describes a response with status code 201, with default header values.

The legal disclosure message was created for the order.
*/
type CreateLegalDisclosureCreated struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *messaging_models.CreateLegalDisclosureResponse
}

// IsSuccess returns true when this create legal disclosure created response has a 2xx status code
func (o *CreateLegalDisclosureCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create legal disclosure created response has a 3xx status code
func (o *CreateLegalDisclosureCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create legal disclosure created response has a 4xx status code
func (o *CreateLegalDisclosureCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create legal disclosure created response has a 5xx status code
func (o *CreateLegalDisclosureCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create legal disclosure created response a status code equal to that given
func (o *CreateLegalDisclosureCreated) IsCode(code int) bool {
	return code == 201
}

func (o *CreateLegalDisclosureCreated) Error() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/legalDisclosure][%d] createLegalDisclosureCreated  %+v", 201, o.Payload)
}

func (o *CreateLegalDisclosureCreated) String() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/legalDisclosure][%d] createLegalDisclosureCreated  %+v", 201, o.Payload)
}

func (o *CreateLegalDisclosureCreated) GetPayload() *messaging_models.CreateLegalDisclosureResponse {
	return o.Payload
}

func (o *CreateLegalDisclosureCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(messaging_models.CreateLegalDisclosureResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateLegalDisclosureBadRequest creates a CreateLegalDisclosureBadRequest with default headers values
func NewCreateLegalDisclosureBadRequest() *CreateLegalDisclosureBadRequest {
	return &CreateLegalDisclosureBadRequest{}
}

/*
CreateLegalDisclosureBadRequest describes a response with status code 400, with default header values.

Request has missing or invalid parameters and cannot be parsed.
*/
type CreateLegalDisclosureBadRequest struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *messaging_models.CreateLegalDisclosureResponse
}

// IsSuccess returns true when this create legal disclosure bad request response has a 2xx status code
func (o *CreateLegalDisclosureBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create legal disclosure bad request response has a 3xx status code
func (o *CreateLegalDisclosureBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create legal disclosure bad request response has a 4xx status code
func (o *CreateLegalDisclosureBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create legal disclosure bad request response has a 5xx status code
func (o *CreateLegalDisclosureBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create legal disclosure bad request response a status code equal to that given
func (o *CreateLegalDisclosureBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *CreateLegalDisclosureBadRequest) Error() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/legalDisclosure][%d] createLegalDisclosureBadRequest  %+v", 400, o.Payload)
}

func (o *CreateLegalDisclosureBadRequest) String() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/legalDisclosure][%d] createLegalDisclosureBadRequest  %+v", 400, o.Payload)
}

func (o *CreateLegalDisclosureBadRequest) GetPayload() *messaging_models.CreateLegalDisclosureResponse {
	return o.Payload
}

func (o *CreateLegalDisclosureBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(messaging_models.CreateLegalDisclosureResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateLegalDisclosureForbidden creates a CreateLegalDisclosureForbidden with default headers values
func NewCreateLegalDisclosureForbidden() *CreateLegalDisclosureForbidden {
	return &CreateLegalDisclosureForbidden{}
}

/*
CreateLegalDisclosureForbidden describes a response with status code 403, with default header values.

Indicates that access to the resource is forbidden. Possible reasons include Access Denied, Unauthorized, Expired Token, or Invalid Signature.
*/
type CreateLegalDisclosureForbidden struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *messaging_models.CreateLegalDisclosureResponse
}

// IsSuccess returns true when this create legal disclosure forbidden response has a 2xx status code
func (o *CreateLegalDisclosureForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create legal disclosure forbidden response has a 3xx status code
func (o *CreateLegalDisclosureForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create legal disclosure forbidden response has a 4xx status code
func (o *CreateLegalDisclosureForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create legal disclosure forbidden response has a 5xx status code
func (o *CreateLegalDisclosureForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create legal disclosure forbidden response a status code equal to that given
func (o *CreateLegalDisclosureForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *CreateLegalDisclosureForbidden) Error() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/legalDisclosure][%d] createLegalDisclosureForbidden  %+v", 403, o.Payload)
}

func (o *CreateLegalDisclosureForbidden) String() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/legalDisclosure][%d] createLegalDisclosureForbidden  %+v", 403, o.Payload)
}

func (o *CreateLegalDisclosureForbidden) GetPayload() *messaging_models.CreateLegalDisclosureResponse {
	return o.Payload
}

func (o *CreateLegalDisclosureForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(messaging_models.CreateLegalDisclosureResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateLegalDisclosureNotFound creates a CreateLegalDisclosureNotFound with default headers values
func NewCreateLegalDisclosureNotFound() *CreateLegalDisclosureNotFound {
	return &CreateLegalDisclosureNotFound{}
}

/*
CreateLegalDisclosureNotFound describes a response with status code 404, with default header values.

The resource specified does not exist.
*/
type CreateLegalDisclosureNotFound struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *messaging_models.CreateLegalDisclosureResponse
}

// IsSuccess returns true when this create legal disclosure not found response has a 2xx status code
func (o *CreateLegalDisclosureNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create legal disclosure not found response has a 3xx status code
func (o *CreateLegalDisclosureNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create legal disclosure not found response has a 4xx status code
func (o *CreateLegalDisclosureNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this create legal disclosure not found response has a 5xx status code
func (o *CreateLegalDisclosureNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this create legal disclosure not found response a status code equal to that given
func (o *CreateLegalDisclosureNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *CreateLegalDisclosureNotFound) Error() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/legalDisclosure][%d] createLegalDisclosureNotFound  %+v", 404, o.Payload)
}

func (o *CreateLegalDisclosureNotFound) String() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/legalDisclosure][%d] createLegalDisclosureNotFound  %+v", 404, o.Payload)
}

func (o *CreateLegalDisclosureNotFound) GetPayload() *messaging_models.CreateLegalDisclosureResponse {
	return o.Payload
}

func (o *CreateLegalDisclosureNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(messaging_models.CreateLegalDisclosureResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateLegalDisclosureRequestEntityTooLarge creates a CreateLegalDisclosureRequestEntityTooLarge with default headers values
func NewCreateLegalDisclosureRequestEntityTooLarge() *CreateLegalDisclosureRequestEntityTooLarge {
	return &CreateLegalDisclosureRequestEntityTooLarge{}
}

/*
CreateLegalDisclosureRequestEntityTooLarge describes a response with status code 413, with default header values.

The request size exceeded the maximum accepted size.
*/
type CreateLegalDisclosureRequestEntityTooLarge struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *messaging_models.CreateLegalDisclosureResponse
}

// IsSuccess returns true when this create legal disclosure request entity too large response has a 2xx status code
func (o *CreateLegalDisclosureRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create legal disclosure request entity too large response has a 3xx status code
func (o *CreateLegalDisclosureRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create legal disclosure request entity too large response has a 4xx status code
func (o *CreateLegalDisclosureRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this create legal disclosure request entity too large response has a 5xx status code
func (o *CreateLegalDisclosureRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this create legal disclosure request entity too large response a status code equal to that given
func (o *CreateLegalDisclosureRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *CreateLegalDisclosureRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/legalDisclosure][%d] createLegalDisclosureRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *CreateLegalDisclosureRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/legalDisclosure][%d] createLegalDisclosureRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *CreateLegalDisclosureRequestEntityTooLarge) GetPayload() *messaging_models.CreateLegalDisclosureResponse {
	return o.Payload
}

func (o *CreateLegalDisclosureRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(messaging_models.CreateLegalDisclosureResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateLegalDisclosureUnsupportedMediaType creates a CreateLegalDisclosureUnsupportedMediaType with default headers values
func NewCreateLegalDisclosureUnsupportedMediaType() *CreateLegalDisclosureUnsupportedMediaType {
	return &CreateLegalDisclosureUnsupportedMediaType{}
}

/*
CreateLegalDisclosureUnsupportedMediaType describes a response with status code 415, with default header values.

The request payload is in an unsupported format.
*/
type CreateLegalDisclosureUnsupportedMediaType struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *messaging_models.CreateLegalDisclosureResponse
}

// IsSuccess returns true when this create legal disclosure unsupported media type response has a 2xx status code
func (o *CreateLegalDisclosureUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create legal disclosure unsupported media type response has a 3xx status code
func (o *CreateLegalDisclosureUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create legal disclosure unsupported media type response has a 4xx status code
func (o *CreateLegalDisclosureUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this create legal disclosure unsupported media type response has a 5xx status code
func (o *CreateLegalDisclosureUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this create legal disclosure unsupported media type response a status code equal to that given
func (o *CreateLegalDisclosureUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *CreateLegalDisclosureUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/legalDisclosure][%d] createLegalDisclosureUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *CreateLegalDisclosureUnsupportedMediaType) String() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/legalDisclosure][%d] createLegalDisclosureUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *CreateLegalDisclosureUnsupportedMediaType) GetPayload() *messaging_models.CreateLegalDisclosureResponse {
	return o.Payload
}

func (o *CreateLegalDisclosureUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(messaging_models.CreateLegalDisclosureResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateLegalDisclosureTooManyRequests creates a CreateLegalDisclosureTooManyRequests with default headers values
func NewCreateLegalDisclosureTooManyRequests() *CreateLegalDisclosureTooManyRequests {
	return &CreateLegalDisclosureTooManyRequests{}
}

/*
CreateLegalDisclosureTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type CreateLegalDisclosureTooManyRequests struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *messaging_models.CreateLegalDisclosureResponse
}

// IsSuccess returns true when this create legal disclosure too many requests response has a 2xx status code
func (o *CreateLegalDisclosureTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create legal disclosure too many requests response has a 3xx status code
func (o *CreateLegalDisclosureTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create legal disclosure too many requests response has a 4xx status code
func (o *CreateLegalDisclosureTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this create legal disclosure too many requests response has a 5xx status code
func (o *CreateLegalDisclosureTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this create legal disclosure too many requests response a status code equal to that given
func (o *CreateLegalDisclosureTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *CreateLegalDisclosureTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/legalDisclosure][%d] createLegalDisclosureTooManyRequests  %+v", 429, o.Payload)
}

func (o *CreateLegalDisclosureTooManyRequests) String() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/legalDisclosure][%d] createLegalDisclosureTooManyRequests  %+v", 429, o.Payload)
}

func (o *CreateLegalDisclosureTooManyRequests) GetPayload() *messaging_models.CreateLegalDisclosureResponse {
	return o.Payload
}

func (o *CreateLegalDisclosureTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(messaging_models.CreateLegalDisclosureResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateLegalDisclosureInternalServerError creates a CreateLegalDisclosureInternalServerError with default headers values
func NewCreateLegalDisclosureInternalServerError() *CreateLegalDisclosureInternalServerError {
	return &CreateLegalDisclosureInternalServerError{}
}

/*
CreateLegalDisclosureInternalServerError describes a response with status code 500, with default header values.

An unexpected condition occurred that prevented the server from fulfilling the request.
*/
type CreateLegalDisclosureInternalServerError struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *messaging_models.CreateLegalDisclosureResponse
}

// IsSuccess returns true when this create legal disclosure internal server error response has a 2xx status code
func (o *CreateLegalDisclosureInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create legal disclosure internal server error response has a 3xx status code
func (o *CreateLegalDisclosureInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create legal disclosure internal server error response has a 4xx status code
func (o *CreateLegalDisclosureInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create legal disclosure internal server error response has a 5xx status code
func (o *CreateLegalDisclosureInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create legal disclosure internal server error response a status code equal to that given
func (o *CreateLegalDisclosureInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *CreateLegalDisclosureInternalServerError) Error() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/legalDisclosure][%d] createLegalDisclosureInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateLegalDisclosureInternalServerError) String() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/legalDisclosure][%d] createLegalDisclosureInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateLegalDisclosureInternalServerError) GetPayload() *messaging_models.CreateLegalDisclosureResponse {
	return o.Payload
}

func (o *CreateLegalDisclosureInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(messaging_models.CreateLegalDisclosureResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateLegalDisclosureServiceUnavailable creates a CreateLegalDisclosureServiceUnavailable with default headers values
func NewCreateLegalDisclosureServiceUnavailable() *CreateLegalDisclosureServiceUnavailable {
	return &CreateLegalDisclosureServiceUnavailable{}
}

/*
CreateLegalDisclosureServiceUnavailable describes a response with status code 503, with default header values.

Temporary overloading or maintenance of the server.
*/
type CreateLegalDisclosureServiceUnavailable struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *messaging_models.CreateLegalDisclosureResponse
}

// IsSuccess returns true when this create legal disclosure service unavailable response has a 2xx status code
func (o *CreateLegalDisclosureServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create legal disclosure service unavailable response has a 3xx status code
func (o *CreateLegalDisclosureServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create legal disclosure service unavailable response has a 4xx status code
func (o *CreateLegalDisclosureServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this create legal disclosure service unavailable response has a 5xx status code
func (o *CreateLegalDisclosureServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this create legal disclosure service unavailable response a status code equal to that given
func (o *CreateLegalDisclosureServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *CreateLegalDisclosureServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/legalDisclosure][%d] createLegalDisclosureServiceUnavailable  %+v", 503, o.Payload)
}

func (o *CreateLegalDisclosureServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/legalDisclosure][%d] createLegalDisclosureServiceUnavailable  %+v", 503, o.Payload)
}

func (o *CreateLegalDisclosureServiceUnavailable) GetPayload() *messaging_models.CreateLegalDisclosureResponse {
	return o.Payload
}

func (o *CreateLegalDisclosureServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(messaging_models.CreateLegalDisclosureResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
