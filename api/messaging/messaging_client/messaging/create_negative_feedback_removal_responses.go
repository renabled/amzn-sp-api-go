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

// CreateNegativeFeedbackRemovalReader is a Reader for the CreateNegativeFeedbackRemoval structure.
type CreateNegativeFeedbackRemovalReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateNegativeFeedbackRemovalReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateNegativeFeedbackRemovalCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateNegativeFeedbackRemovalBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateNegativeFeedbackRemovalForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateNegativeFeedbackRemovalNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewCreateNegativeFeedbackRemovalRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewCreateNegativeFeedbackRemovalUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewCreateNegativeFeedbackRemovalTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateNegativeFeedbackRemovalInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewCreateNegativeFeedbackRemovalServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateNegativeFeedbackRemovalCreated creates a CreateNegativeFeedbackRemovalCreated with default headers values
func NewCreateNegativeFeedbackRemovalCreated() *CreateNegativeFeedbackRemovalCreated {
	return &CreateNegativeFeedbackRemovalCreated{}
}

/*
CreateNegativeFeedbackRemovalCreated describes a response with status code 201, with default header values.

The negativeFeedbackRemoval message was created for the order.
*/
type CreateNegativeFeedbackRemovalCreated struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *messaging_models.CreateNegativeFeedbackRemovalResponse
}

// IsSuccess returns true when this create negative feedback removal created response has a 2xx status code
func (o *CreateNegativeFeedbackRemovalCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create negative feedback removal created response has a 3xx status code
func (o *CreateNegativeFeedbackRemovalCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create negative feedback removal created response has a 4xx status code
func (o *CreateNegativeFeedbackRemovalCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create negative feedback removal created response has a 5xx status code
func (o *CreateNegativeFeedbackRemovalCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create negative feedback removal created response a status code equal to that given
func (o *CreateNegativeFeedbackRemovalCreated) IsCode(code int) bool {
	return code == 201
}

func (o *CreateNegativeFeedbackRemovalCreated) Error() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/negativeFeedbackRemoval][%d] createNegativeFeedbackRemovalCreated  %+v", 201, o.Payload)
}

func (o *CreateNegativeFeedbackRemovalCreated) String() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/negativeFeedbackRemoval][%d] createNegativeFeedbackRemovalCreated  %+v", 201, o.Payload)
}

func (o *CreateNegativeFeedbackRemovalCreated) GetPayload() *messaging_models.CreateNegativeFeedbackRemovalResponse {
	return o.Payload
}

func (o *CreateNegativeFeedbackRemovalCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(messaging_models.CreateNegativeFeedbackRemovalResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateNegativeFeedbackRemovalBadRequest creates a CreateNegativeFeedbackRemovalBadRequest with default headers values
func NewCreateNegativeFeedbackRemovalBadRequest() *CreateNegativeFeedbackRemovalBadRequest {
	return &CreateNegativeFeedbackRemovalBadRequest{}
}

/*
CreateNegativeFeedbackRemovalBadRequest describes a response with status code 400, with default header values.

Request has missing or invalid parameters and cannot be parsed.
*/
type CreateNegativeFeedbackRemovalBadRequest struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *messaging_models.CreateNegativeFeedbackRemovalResponse
}

// IsSuccess returns true when this create negative feedback removal bad request response has a 2xx status code
func (o *CreateNegativeFeedbackRemovalBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create negative feedback removal bad request response has a 3xx status code
func (o *CreateNegativeFeedbackRemovalBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create negative feedback removal bad request response has a 4xx status code
func (o *CreateNegativeFeedbackRemovalBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create negative feedback removal bad request response has a 5xx status code
func (o *CreateNegativeFeedbackRemovalBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create negative feedback removal bad request response a status code equal to that given
func (o *CreateNegativeFeedbackRemovalBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *CreateNegativeFeedbackRemovalBadRequest) Error() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/negativeFeedbackRemoval][%d] createNegativeFeedbackRemovalBadRequest  %+v", 400, o.Payload)
}

func (o *CreateNegativeFeedbackRemovalBadRequest) String() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/negativeFeedbackRemoval][%d] createNegativeFeedbackRemovalBadRequest  %+v", 400, o.Payload)
}

func (o *CreateNegativeFeedbackRemovalBadRequest) GetPayload() *messaging_models.CreateNegativeFeedbackRemovalResponse {
	return o.Payload
}

func (o *CreateNegativeFeedbackRemovalBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(messaging_models.CreateNegativeFeedbackRemovalResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateNegativeFeedbackRemovalForbidden creates a CreateNegativeFeedbackRemovalForbidden with default headers values
func NewCreateNegativeFeedbackRemovalForbidden() *CreateNegativeFeedbackRemovalForbidden {
	return &CreateNegativeFeedbackRemovalForbidden{}
}

/*
CreateNegativeFeedbackRemovalForbidden describes a response with status code 403, with default header values.

Indicates that access to the resource is forbidden. Possible reasons include Access Denied, Unauthorized, Expired Token, or Invalid Signature.
*/
type CreateNegativeFeedbackRemovalForbidden struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *messaging_models.CreateNegativeFeedbackRemovalResponse
}

// IsSuccess returns true when this create negative feedback removal forbidden response has a 2xx status code
func (o *CreateNegativeFeedbackRemovalForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create negative feedback removal forbidden response has a 3xx status code
func (o *CreateNegativeFeedbackRemovalForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create negative feedback removal forbidden response has a 4xx status code
func (o *CreateNegativeFeedbackRemovalForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create negative feedback removal forbidden response has a 5xx status code
func (o *CreateNegativeFeedbackRemovalForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create negative feedback removal forbidden response a status code equal to that given
func (o *CreateNegativeFeedbackRemovalForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *CreateNegativeFeedbackRemovalForbidden) Error() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/negativeFeedbackRemoval][%d] createNegativeFeedbackRemovalForbidden  %+v", 403, o.Payload)
}

func (o *CreateNegativeFeedbackRemovalForbidden) String() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/negativeFeedbackRemoval][%d] createNegativeFeedbackRemovalForbidden  %+v", 403, o.Payload)
}

func (o *CreateNegativeFeedbackRemovalForbidden) GetPayload() *messaging_models.CreateNegativeFeedbackRemovalResponse {
	return o.Payload
}

func (o *CreateNegativeFeedbackRemovalForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(messaging_models.CreateNegativeFeedbackRemovalResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateNegativeFeedbackRemovalNotFound creates a CreateNegativeFeedbackRemovalNotFound with default headers values
func NewCreateNegativeFeedbackRemovalNotFound() *CreateNegativeFeedbackRemovalNotFound {
	return &CreateNegativeFeedbackRemovalNotFound{}
}

/*
CreateNegativeFeedbackRemovalNotFound describes a response with status code 404, with default header values.

The resource specified does not exist.
*/
type CreateNegativeFeedbackRemovalNotFound struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *messaging_models.CreateNegativeFeedbackRemovalResponse
}

// IsSuccess returns true when this create negative feedback removal not found response has a 2xx status code
func (o *CreateNegativeFeedbackRemovalNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create negative feedback removal not found response has a 3xx status code
func (o *CreateNegativeFeedbackRemovalNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create negative feedback removal not found response has a 4xx status code
func (o *CreateNegativeFeedbackRemovalNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this create negative feedback removal not found response has a 5xx status code
func (o *CreateNegativeFeedbackRemovalNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this create negative feedback removal not found response a status code equal to that given
func (o *CreateNegativeFeedbackRemovalNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *CreateNegativeFeedbackRemovalNotFound) Error() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/negativeFeedbackRemoval][%d] createNegativeFeedbackRemovalNotFound  %+v", 404, o.Payload)
}

func (o *CreateNegativeFeedbackRemovalNotFound) String() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/negativeFeedbackRemoval][%d] createNegativeFeedbackRemovalNotFound  %+v", 404, o.Payload)
}

func (o *CreateNegativeFeedbackRemovalNotFound) GetPayload() *messaging_models.CreateNegativeFeedbackRemovalResponse {
	return o.Payload
}

func (o *CreateNegativeFeedbackRemovalNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(messaging_models.CreateNegativeFeedbackRemovalResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateNegativeFeedbackRemovalRequestEntityTooLarge creates a CreateNegativeFeedbackRemovalRequestEntityTooLarge with default headers values
func NewCreateNegativeFeedbackRemovalRequestEntityTooLarge() *CreateNegativeFeedbackRemovalRequestEntityTooLarge {
	return &CreateNegativeFeedbackRemovalRequestEntityTooLarge{}
}

/*
CreateNegativeFeedbackRemovalRequestEntityTooLarge describes a response with status code 413, with default header values.

The request size exceeded the maximum accepted size.
*/
type CreateNegativeFeedbackRemovalRequestEntityTooLarge struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *messaging_models.CreateNegativeFeedbackRemovalResponse
}

// IsSuccess returns true when this create negative feedback removal request entity too large response has a 2xx status code
func (o *CreateNegativeFeedbackRemovalRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create negative feedback removal request entity too large response has a 3xx status code
func (o *CreateNegativeFeedbackRemovalRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create negative feedback removal request entity too large response has a 4xx status code
func (o *CreateNegativeFeedbackRemovalRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this create negative feedback removal request entity too large response has a 5xx status code
func (o *CreateNegativeFeedbackRemovalRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this create negative feedback removal request entity too large response a status code equal to that given
func (o *CreateNegativeFeedbackRemovalRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *CreateNegativeFeedbackRemovalRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/negativeFeedbackRemoval][%d] createNegativeFeedbackRemovalRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *CreateNegativeFeedbackRemovalRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/negativeFeedbackRemoval][%d] createNegativeFeedbackRemovalRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *CreateNegativeFeedbackRemovalRequestEntityTooLarge) GetPayload() *messaging_models.CreateNegativeFeedbackRemovalResponse {
	return o.Payload
}

func (o *CreateNegativeFeedbackRemovalRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(messaging_models.CreateNegativeFeedbackRemovalResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateNegativeFeedbackRemovalUnsupportedMediaType creates a CreateNegativeFeedbackRemovalUnsupportedMediaType with default headers values
func NewCreateNegativeFeedbackRemovalUnsupportedMediaType() *CreateNegativeFeedbackRemovalUnsupportedMediaType {
	return &CreateNegativeFeedbackRemovalUnsupportedMediaType{}
}

/*
CreateNegativeFeedbackRemovalUnsupportedMediaType describes a response with status code 415, with default header values.

The request payload is in an unsupported format.
*/
type CreateNegativeFeedbackRemovalUnsupportedMediaType struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *messaging_models.CreateNegativeFeedbackRemovalResponse
}

// IsSuccess returns true when this create negative feedback removal unsupported media type response has a 2xx status code
func (o *CreateNegativeFeedbackRemovalUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create negative feedback removal unsupported media type response has a 3xx status code
func (o *CreateNegativeFeedbackRemovalUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create negative feedback removal unsupported media type response has a 4xx status code
func (o *CreateNegativeFeedbackRemovalUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this create negative feedback removal unsupported media type response has a 5xx status code
func (o *CreateNegativeFeedbackRemovalUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this create negative feedback removal unsupported media type response a status code equal to that given
func (o *CreateNegativeFeedbackRemovalUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *CreateNegativeFeedbackRemovalUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/negativeFeedbackRemoval][%d] createNegativeFeedbackRemovalUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *CreateNegativeFeedbackRemovalUnsupportedMediaType) String() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/negativeFeedbackRemoval][%d] createNegativeFeedbackRemovalUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *CreateNegativeFeedbackRemovalUnsupportedMediaType) GetPayload() *messaging_models.CreateNegativeFeedbackRemovalResponse {
	return o.Payload
}

func (o *CreateNegativeFeedbackRemovalUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(messaging_models.CreateNegativeFeedbackRemovalResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateNegativeFeedbackRemovalTooManyRequests creates a CreateNegativeFeedbackRemovalTooManyRequests with default headers values
func NewCreateNegativeFeedbackRemovalTooManyRequests() *CreateNegativeFeedbackRemovalTooManyRequests {
	return &CreateNegativeFeedbackRemovalTooManyRequests{}
}

/*
CreateNegativeFeedbackRemovalTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type CreateNegativeFeedbackRemovalTooManyRequests struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *messaging_models.CreateNegativeFeedbackRemovalResponse
}

// IsSuccess returns true when this create negative feedback removal too many requests response has a 2xx status code
func (o *CreateNegativeFeedbackRemovalTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create negative feedback removal too many requests response has a 3xx status code
func (o *CreateNegativeFeedbackRemovalTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create negative feedback removal too many requests response has a 4xx status code
func (o *CreateNegativeFeedbackRemovalTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this create negative feedback removal too many requests response has a 5xx status code
func (o *CreateNegativeFeedbackRemovalTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this create negative feedback removal too many requests response a status code equal to that given
func (o *CreateNegativeFeedbackRemovalTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *CreateNegativeFeedbackRemovalTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/negativeFeedbackRemoval][%d] createNegativeFeedbackRemovalTooManyRequests  %+v", 429, o.Payload)
}

func (o *CreateNegativeFeedbackRemovalTooManyRequests) String() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/negativeFeedbackRemoval][%d] createNegativeFeedbackRemovalTooManyRequests  %+v", 429, o.Payload)
}

func (o *CreateNegativeFeedbackRemovalTooManyRequests) GetPayload() *messaging_models.CreateNegativeFeedbackRemovalResponse {
	return o.Payload
}

func (o *CreateNegativeFeedbackRemovalTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(messaging_models.CreateNegativeFeedbackRemovalResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateNegativeFeedbackRemovalInternalServerError creates a CreateNegativeFeedbackRemovalInternalServerError with default headers values
func NewCreateNegativeFeedbackRemovalInternalServerError() *CreateNegativeFeedbackRemovalInternalServerError {
	return &CreateNegativeFeedbackRemovalInternalServerError{}
}

/*
CreateNegativeFeedbackRemovalInternalServerError describes a response with status code 500, with default header values.

An unexpected condition occurred that prevented the server from fulfilling the request.
*/
type CreateNegativeFeedbackRemovalInternalServerError struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *messaging_models.CreateNegativeFeedbackRemovalResponse
}

// IsSuccess returns true when this create negative feedback removal internal server error response has a 2xx status code
func (o *CreateNegativeFeedbackRemovalInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create negative feedback removal internal server error response has a 3xx status code
func (o *CreateNegativeFeedbackRemovalInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create negative feedback removal internal server error response has a 4xx status code
func (o *CreateNegativeFeedbackRemovalInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create negative feedback removal internal server error response has a 5xx status code
func (o *CreateNegativeFeedbackRemovalInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create negative feedback removal internal server error response a status code equal to that given
func (o *CreateNegativeFeedbackRemovalInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *CreateNegativeFeedbackRemovalInternalServerError) Error() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/negativeFeedbackRemoval][%d] createNegativeFeedbackRemovalInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateNegativeFeedbackRemovalInternalServerError) String() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/negativeFeedbackRemoval][%d] createNegativeFeedbackRemovalInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateNegativeFeedbackRemovalInternalServerError) GetPayload() *messaging_models.CreateNegativeFeedbackRemovalResponse {
	return o.Payload
}

func (o *CreateNegativeFeedbackRemovalInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(messaging_models.CreateNegativeFeedbackRemovalResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateNegativeFeedbackRemovalServiceUnavailable creates a CreateNegativeFeedbackRemovalServiceUnavailable with default headers values
func NewCreateNegativeFeedbackRemovalServiceUnavailable() *CreateNegativeFeedbackRemovalServiceUnavailable {
	return &CreateNegativeFeedbackRemovalServiceUnavailable{}
}

/*
CreateNegativeFeedbackRemovalServiceUnavailable describes a response with status code 503, with default header values.

Temporary overloading or maintenance of the server.
*/
type CreateNegativeFeedbackRemovalServiceUnavailable struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *messaging_models.CreateNegativeFeedbackRemovalResponse
}

// IsSuccess returns true when this create negative feedback removal service unavailable response has a 2xx status code
func (o *CreateNegativeFeedbackRemovalServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create negative feedback removal service unavailable response has a 3xx status code
func (o *CreateNegativeFeedbackRemovalServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create negative feedback removal service unavailable response has a 4xx status code
func (o *CreateNegativeFeedbackRemovalServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this create negative feedback removal service unavailable response has a 5xx status code
func (o *CreateNegativeFeedbackRemovalServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this create negative feedback removal service unavailable response a status code equal to that given
func (o *CreateNegativeFeedbackRemovalServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *CreateNegativeFeedbackRemovalServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/negativeFeedbackRemoval][%d] createNegativeFeedbackRemovalServiceUnavailable  %+v", 503, o.Payload)
}

func (o *CreateNegativeFeedbackRemovalServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/negativeFeedbackRemoval][%d] createNegativeFeedbackRemovalServiceUnavailable  %+v", 503, o.Payload)
}

func (o *CreateNegativeFeedbackRemovalServiceUnavailable) GetPayload() *messaging_models.CreateNegativeFeedbackRemovalResponse {
	return o.Payload
}

func (o *CreateNegativeFeedbackRemovalServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(messaging_models.CreateNegativeFeedbackRemovalResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
