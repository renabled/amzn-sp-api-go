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

// CreateUnexpectedProblemReader is a Reader for the CreateUnexpectedProblem structure.
type CreateUnexpectedProblemReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateUnexpectedProblemReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateUnexpectedProblemCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateUnexpectedProblemBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateUnexpectedProblemForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateUnexpectedProblemNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewCreateUnexpectedProblemRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewCreateUnexpectedProblemUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewCreateUnexpectedProblemTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateUnexpectedProblemInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewCreateUnexpectedProblemServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateUnexpectedProblemCreated creates a CreateUnexpectedProblemCreated with default headers values
func NewCreateUnexpectedProblemCreated() *CreateUnexpectedProblemCreated {
	return &CreateUnexpectedProblemCreated{}
}

/*
CreateUnexpectedProblemCreated describes a response with status code 201, with default header values.

The message was created for the order.
*/
type CreateUnexpectedProblemCreated struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *messaging_models.CreateUnexpectedProblemResponse
}

// IsSuccess returns true when this create unexpected problem created response has a 2xx status code
func (o *CreateUnexpectedProblemCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create unexpected problem created response has a 3xx status code
func (o *CreateUnexpectedProblemCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create unexpected problem created response has a 4xx status code
func (o *CreateUnexpectedProblemCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create unexpected problem created response has a 5xx status code
func (o *CreateUnexpectedProblemCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create unexpected problem created response a status code equal to that given
func (o *CreateUnexpectedProblemCreated) IsCode(code int) bool {
	return code == 201
}

func (o *CreateUnexpectedProblemCreated) Error() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/unexpectedProblem][%d] createUnexpectedProblemCreated  %+v", 201, o.Payload)
}

func (o *CreateUnexpectedProblemCreated) String() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/unexpectedProblem][%d] createUnexpectedProblemCreated  %+v", 201, o.Payload)
}

func (o *CreateUnexpectedProblemCreated) GetPayload() *messaging_models.CreateUnexpectedProblemResponse {
	return o.Payload
}

func (o *CreateUnexpectedProblemCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(messaging_models.CreateUnexpectedProblemResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateUnexpectedProblemBadRequest creates a CreateUnexpectedProblemBadRequest with default headers values
func NewCreateUnexpectedProblemBadRequest() *CreateUnexpectedProblemBadRequest {
	return &CreateUnexpectedProblemBadRequest{}
}

/*
CreateUnexpectedProblemBadRequest describes a response with status code 400, with default header values.

Request has missing or invalid parameters and cannot be parsed.
*/
type CreateUnexpectedProblemBadRequest struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *messaging_models.CreateUnexpectedProblemResponse
}

// IsSuccess returns true when this create unexpected problem bad request response has a 2xx status code
func (o *CreateUnexpectedProblemBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create unexpected problem bad request response has a 3xx status code
func (o *CreateUnexpectedProblemBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create unexpected problem bad request response has a 4xx status code
func (o *CreateUnexpectedProblemBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create unexpected problem bad request response has a 5xx status code
func (o *CreateUnexpectedProblemBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create unexpected problem bad request response a status code equal to that given
func (o *CreateUnexpectedProblemBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *CreateUnexpectedProblemBadRequest) Error() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/unexpectedProblem][%d] createUnexpectedProblemBadRequest  %+v", 400, o.Payload)
}

func (o *CreateUnexpectedProblemBadRequest) String() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/unexpectedProblem][%d] createUnexpectedProblemBadRequest  %+v", 400, o.Payload)
}

func (o *CreateUnexpectedProblemBadRequest) GetPayload() *messaging_models.CreateUnexpectedProblemResponse {
	return o.Payload
}

func (o *CreateUnexpectedProblemBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(messaging_models.CreateUnexpectedProblemResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateUnexpectedProblemForbidden creates a CreateUnexpectedProblemForbidden with default headers values
func NewCreateUnexpectedProblemForbidden() *CreateUnexpectedProblemForbidden {
	return &CreateUnexpectedProblemForbidden{}
}

/*
CreateUnexpectedProblemForbidden describes a response with status code 403, with default header values.

Indicates that access to the resource is forbidden. Possible reasons include Access Denied, Unauthorized, Expired Token, or Invalid Signature.
*/
type CreateUnexpectedProblemForbidden struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *messaging_models.CreateUnexpectedProblemResponse
}

// IsSuccess returns true when this create unexpected problem forbidden response has a 2xx status code
func (o *CreateUnexpectedProblemForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create unexpected problem forbidden response has a 3xx status code
func (o *CreateUnexpectedProblemForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create unexpected problem forbidden response has a 4xx status code
func (o *CreateUnexpectedProblemForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create unexpected problem forbidden response has a 5xx status code
func (o *CreateUnexpectedProblemForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create unexpected problem forbidden response a status code equal to that given
func (o *CreateUnexpectedProblemForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *CreateUnexpectedProblemForbidden) Error() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/unexpectedProblem][%d] createUnexpectedProblemForbidden  %+v", 403, o.Payload)
}

func (o *CreateUnexpectedProblemForbidden) String() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/unexpectedProblem][%d] createUnexpectedProblemForbidden  %+v", 403, o.Payload)
}

func (o *CreateUnexpectedProblemForbidden) GetPayload() *messaging_models.CreateUnexpectedProblemResponse {
	return o.Payload
}

func (o *CreateUnexpectedProblemForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(messaging_models.CreateUnexpectedProblemResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateUnexpectedProblemNotFound creates a CreateUnexpectedProblemNotFound with default headers values
func NewCreateUnexpectedProblemNotFound() *CreateUnexpectedProblemNotFound {
	return &CreateUnexpectedProblemNotFound{}
}

/*
CreateUnexpectedProblemNotFound describes a response with status code 404, with default header values.

The resource specified does not exist.
*/
type CreateUnexpectedProblemNotFound struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *messaging_models.CreateUnexpectedProblemResponse
}

// IsSuccess returns true when this create unexpected problem not found response has a 2xx status code
func (o *CreateUnexpectedProblemNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create unexpected problem not found response has a 3xx status code
func (o *CreateUnexpectedProblemNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create unexpected problem not found response has a 4xx status code
func (o *CreateUnexpectedProblemNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this create unexpected problem not found response has a 5xx status code
func (o *CreateUnexpectedProblemNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this create unexpected problem not found response a status code equal to that given
func (o *CreateUnexpectedProblemNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *CreateUnexpectedProblemNotFound) Error() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/unexpectedProblem][%d] createUnexpectedProblemNotFound  %+v", 404, o.Payload)
}

func (o *CreateUnexpectedProblemNotFound) String() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/unexpectedProblem][%d] createUnexpectedProblemNotFound  %+v", 404, o.Payload)
}

func (o *CreateUnexpectedProblemNotFound) GetPayload() *messaging_models.CreateUnexpectedProblemResponse {
	return o.Payload
}

func (o *CreateUnexpectedProblemNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(messaging_models.CreateUnexpectedProblemResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateUnexpectedProblemRequestEntityTooLarge creates a CreateUnexpectedProblemRequestEntityTooLarge with default headers values
func NewCreateUnexpectedProblemRequestEntityTooLarge() *CreateUnexpectedProblemRequestEntityTooLarge {
	return &CreateUnexpectedProblemRequestEntityTooLarge{}
}

/*
CreateUnexpectedProblemRequestEntityTooLarge describes a response with status code 413, with default header values.

The request size exceeded the maximum accepted size.
*/
type CreateUnexpectedProblemRequestEntityTooLarge struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *messaging_models.CreateUnexpectedProblemResponse
}

// IsSuccess returns true when this create unexpected problem request entity too large response has a 2xx status code
func (o *CreateUnexpectedProblemRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create unexpected problem request entity too large response has a 3xx status code
func (o *CreateUnexpectedProblemRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create unexpected problem request entity too large response has a 4xx status code
func (o *CreateUnexpectedProblemRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this create unexpected problem request entity too large response has a 5xx status code
func (o *CreateUnexpectedProblemRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this create unexpected problem request entity too large response a status code equal to that given
func (o *CreateUnexpectedProblemRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *CreateUnexpectedProblemRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/unexpectedProblem][%d] createUnexpectedProblemRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *CreateUnexpectedProblemRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/unexpectedProblem][%d] createUnexpectedProblemRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *CreateUnexpectedProblemRequestEntityTooLarge) GetPayload() *messaging_models.CreateUnexpectedProblemResponse {
	return o.Payload
}

func (o *CreateUnexpectedProblemRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(messaging_models.CreateUnexpectedProblemResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateUnexpectedProblemUnsupportedMediaType creates a CreateUnexpectedProblemUnsupportedMediaType with default headers values
func NewCreateUnexpectedProblemUnsupportedMediaType() *CreateUnexpectedProblemUnsupportedMediaType {
	return &CreateUnexpectedProblemUnsupportedMediaType{}
}

/*
CreateUnexpectedProblemUnsupportedMediaType describes a response with status code 415, with default header values.

The request payload is in an unsupported format.
*/
type CreateUnexpectedProblemUnsupportedMediaType struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *messaging_models.CreateUnexpectedProblemResponse
}

// IsSuccess returns true when this create unexpected problem unsupported media type response has a 2xx status code
func (o *CreateUnexpectedProblemUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create unexpected problem unsupported media type response has a 3xx status code
func (o *CreateUnexpectedProblemUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create unexpected problem unsupported media type response has a 4xx status code
func (o *CreateUnexpectedProblemUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this create unexpected problem unsupported media type response has a 5xx status code
func (o *CreateUnexpectedProblemUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this create unexpected problem unsupported media type response a status code equal to that given
func (o *CreateUnexpectedProblemUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *CreateUnexpectedProblemUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/unexpectedProblem][%d] createUnexpectedProblemUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *CreateUnexpectedProblemUnsupportedMediaType) String() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/unexpectedProblem][%d] createUnexpectedProblemUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *CreateUnexpectedProblemUnsupportedMediaType) GetPayload() *messaging_models.CreateUnexpectedProblemResponse {
	return o.Payload
}

func (o *CreateUnexpectedProblemUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(messaging_models.CreateUnexpectedProblemResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateUnexpectedProblemTooManyRequests creates a CreateUnexpectedProblemTooManyRequests with default headers values
func NewCreateUnexpectedProblemTooManyRequests() *CreateUnexpectedProblemTooManyRequests {
	return &CreateUnexpectedProblemTooManyRequests{}
}

/*
CreateUnexpectedProblemTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type CreateUnexpectedProblemTooManyRequests struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *messaging_models.CreateUnexpectedProblemResponse
}

// IsSuccess returns true when this create unexpected problem too many requests response has a 2xx status code
func (o *CreateUnexpectedProblemTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create unexpected problem too many requests response has a 3xx status code
func (o *CreateUnexpectedProblemTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create unexpected problem too many requests response has a 4xx status code
func (o *CreateUnexpectedProblemTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this create unexpected problem too many requests response has a 5xx status code
func (o *CreateUnexpectedProblemTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this create unexpected problem too many requests response a status code equal to that given
func (o *CreateUnexpectedProblemTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *CreateUnexpectedProblemTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/unexpectedProblem][%d] createUnexpectedProblemTooManyRequests  %+v", 429, o.Payload)
}

func (o *CreateUnexpectedProblemTooManyRequests) String() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/unexpectedProblem][%d] createUnexpectedProblemTooManyRequests  %+v", 429, o.Payload)
}

func (o *CreateUnexpectedProblemTooManyRequests) GetPayload() *messaging_models.CreateUnexpectedProblemResponse {
	return o.Payload
}

func (o *CreateUnexpectedProblemTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(messaging_models.CreateUnexpectedProblemResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateUnexpectedProblemInternalServerError creates a CreateUnexpectedProblemInternalServerError with default headers values
func NewCreateUnexpectedProblemInternalServerError() *CreateUnexpectedProblemInternalServerError {
	return &CreateUnexpectedProblemInternalServerError{}
}

/*
CreateUnexpectedProblemInternalServerError describes a response with status code 500, with default header values.

An unexpected condition occurred that prevented the server from fulfilling the request.
*/
type CreateUnexpectedProblemInternalServerError struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *messaging_models.CreateUnexpectedProblemResponse
}

// IsSuccess returns true when this create unexpected problem internal server error response has a 2xx status code
func (o *CreateUnexpectedProblemInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create unexpected problem internal server error response has a 3xx status code
func (o *CreateUnexpectedProblemInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create unexpected problem internal server error response has a 4xx status code
func (o *CreateUnexpectedProblemInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create unexpected problem internal server error response has a 5xx status code
func (o *CreateUnexpectedProblemInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create unexpected problem internal server error response a status code equal to that given
func (o *CreateUnexpectedProblemInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *CreateUnexpectedProblemInternalServerError) Error() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/unexpectedProblem][%d] createUnexpectedProblemInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateUnexpectedProblemInternalServerError) String() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/unexpectedProblem][%d] createUnexpectedProblemInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateUnexpectedProblemInternalServerError) GetPayload() *messaging_models.CreateUnexpectedProblemResponse {
	return o.Payload
}

func (o *CreateUnexpectedProblemInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(messaging_models.CreateUnexpectedProblemResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateUnexpectedProblemServiceUnavailable creates a CreateUnexpectedProblemServiceUnavailable with default headers values
func NewCreateUnexpectedProblemServiceUnavailable() *CreateUnexpectedProblemServiceUnavailable {
	return &CreateUnexpectedProblemServiceUnavailable{}
}

/*
CreateUnexpectedProblemServiceUnavailable describes a response with status code 503, with default header values.

Temporary overloading or maintenance of the server.
*/
type CreateUnexpectedProblemServiceUnavailable struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *messaging_models.CreateUnexpectedProblemResponse
}

// IsSuccess returns true when this create unexpected problem service unavailable response has a 2xx status code
func (o *CreateUnexpectedProblemServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create unexpected problem service unavailable response has a 3xx status code
func (o *CreateUnexpectedProblemServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create unexpected problem service unavailable response has a 4xx status code
func (o *CreateUnexpectedProblemServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this create unexpected problem service unavailable response has a 5xx status code
func (o *CreateUnexpectedProblemServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this create unexpected problem service unavailable response a status code equal to that given
func (o *CreateUnexpectedProblemServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *CreateUnexpectedProblemServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/unexpectedProblem][%d] createUnexpectedProblemServiceUnavailable  %+v", 503, o.Payload)
}

func (o *CreateUnexpectedProblemServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/unexpectedProblem][%d] createUnexpectedProblemServiceUnavailable  %+v", 503, o.Payload)
}

func (o *CreateUnexpectedProblemServiceUnavailable) GetPayload() *messaging_models.CreateUnexpectedProblemResponse {
	return o.Payload
}

func (o *CreateUnexpectedProblemServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(messaging_models.CreateUnexpectedProblemResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
