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

// CreateWarrantyReader is a Reader for the CreateWarranty structure.
type CreateWarrantyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateWarrantyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateWarrantyCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateWarrantyBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateWarrantyForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateWarrantyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewCreateWarrantyRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewCreateWarrantyUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewCreateWarrantyTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateWarrantyInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewCreateWarrantyServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateWarrantyCreated creates a CreateWarrantyCreated with default headers values
func NewCreateWarrantyCreated() *CreateWarrantyCreated {
	return &CreateWarrantyCreated{}
}

/*
CreateWarrantyCreated describes a response with status code 201, with default header values.

The message was created for the order.
*/
type CreateWarrantyCreated struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *messaging_models.CreateWarrantyResponse
}

// IsSuccess returns true when this create warranty created response has a 2xx status code
func (o *CreateWarrantyCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create warranty created response has a 3xx status code
func (o *CreateWarrantyCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create warranty created response has a 4xx status code
func (o *CreateWarrantyCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create warranty created response has a 5xx status code
func (o *CreateWarrantyCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create warranty created response a status code equal to that given
func (o *CreateWarrantyCreated) IsCode(code int) bool {
	return code == 201
}

func (o *CreateWarrantyCreated) Error() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/warranty][%d] createWarrantyCreated  %+v", 201, o.Payload)
}

func (o *CreateWarrantyCreated) String() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/warranty][%d] createWarrantyCreated  %+v", 201, o.Payload)
}

func (o *CreateWarrantyCreated) GetPayload() *messaging_models.CreateWarrantyResponse {
	return o.Payload
}

func (o *CreateWarrantyCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(messaging_models.CreateWarrantyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateWarrantyBadRequest creates a CreateWarrantyBadRequest with default headers values
func NewCreateWarrantyBadRequest() *CreateWarrantyBadRequest {
	return &CreateWarrantyBadRequest{}
}

/*
CreateWarrantyBadRequest describes a response with status code 400, with default header values.

Request has missing or invalid parameters and cannot be parsed.
*/
type CreateWarrantyBadRequest struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *messaging_models.CreateWarrantyResponse
}

// IsSuccess returns true when this create warranty bad request response has a 2xx status code
func (o *CreateWarrantyBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create warranty bad request response has a 3xx status code
func (o *CreateWarrantyBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create warranty bad request response has a 4xx status code
func (o *CreateWarrantyBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create warranty bad request response has a 5xx status code
func (o *CreateWarrantyBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create warranty bad request response a status code equal to that given
func (o *CreateWarrantyBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *CreateWarrantyBadRequest) Error() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/warranty][%d] createWarrantyBadRequest  %+v", 400, o.Payload)
}

func (o *CreateWarrantyBadRequest) String() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/warranty][%d] createWarrantyBadRequest  %+v", 400, o.Payload)
}

func (o *CreateWarrantyBadRequest) GetPayload() *messaging_models.CreateWarrantyResponse {
	return o.Payload
}

func (o *CreateWarrantyBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(messaging_models.CreateWarrantyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateWarrantyForbidden creates a CreateWarrantyForbidden with default headers values
func NewCreateWarrantyForbidden() *CreateWarrantyForbidden {
	return &CreateWarrantyForbidden{}
}

/*
CreateWarrantyForbidden describes a response with status code 403, with default header values.

Indicates that access to the resource is forbidden. Possible reasons include Access Denied, Unauthorized, Expired Token, or Invalid Signature.
*/
type CreateWarrantyForbidden struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *messaging_models.CreateWarrantyResponse
}

// IsSuccess returns true when this create warranty forbidden response has a 2xx status code
func (o *CreateWarrantyForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create warranty forbidden response has a 3xx status code
func (o *CreateWarrantyForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create warranty forbidden response has a 4xx status code
func (o *CreateWarrantyForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create warranty forbidden response has a 5xx status code
func (o *CreateWarrantyForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create warranty forbidden response a status code equal to that given
func (o *CreateWarrantyForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *CreateWarrantyForbidden) Error() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/warranty][%d] createWarrantyForbidden  %+v", 403, o.Payload)
}

func (o *CreateWarrantyForbidden) String() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/warranty][%d] createWarrantyForbidden  %+v", 403, o.Payload)
}

func (o *CreateWarrantyForbidden) GetPayload() *messaging_models.CreateWarrantyResponse {
	return o.Payload
}

func (o *CreateWarrantyForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(messaging_models.CreateWarrantyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateWarrantyNotFound creates a CreateWarrantyNotFound with default headers values
func NewCreateWarrantyNotFound() *CreateWarrantyNotFound {
	return &CreateWarrantyNotFound{}
}

/*
CreateWarrantyNotFound describes a response with status code 404, with default header values.

The resource specified does not exist.
*/
type CreateWarrantyNotFound struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *messaging_models.CreateWarrantyResponse
}

// IsSuccess returns true when this create warranty not found response has a 2xx status code
func (o *CreateWarrantyNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create warranty not found response has a 3xx status code
func (o *CreateWarrantyNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create warranty not found response has a 4xx status code
func (o *CreateWarrantyNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this create warranty not found response has a 5xx status code
func (o *CreateWarrantyNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this create warranty not found response a status code equal to that given
func (o *CreateWarrantyNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *CreateWarrantyNotFound) Error() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/warranty][%d] createWarrantyNotFound  %+v", 404, o.Payload)
}

func (o *CreateWarrantyNotFound) String() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/warranty][%d] createWarrantyNotFound  %+v", 404, o.Payload)
}

func (o *CreateWarrantyNotFound) GetPayload() *messaging_models.CreateWarrantyResponse {
	return o.Payload
}

func (o *CreateWarrantyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(messaging_models.CreateWarrantyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateWarrantyRequestEntityTooLarge creates a CreateWarrantyRequestEntityTooLarge with default headers values
func NewCreateWarrantyRequestEntityTooLarge() *CreateWarrantyRequestEntityTooLarge {
	return &CreateWarrantyRequestEntityTooLarge{}
}

/*
CreateWarrantyRequestEntityTooLarge describes a response with status code 413, with default header values.

The request size exceeded the maximum accepted size.
*/
type CreateWarrantyRequestEntityTooLarge struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *messaging_models.CreateWarrantyResponse
}

// IsSuccess returns true when this create warranty request entity too large response has a 2xx status code
func (o *CreateWarrantyRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create warranty request entity too large response has a 3xx status code
func (o *CreateWarrantyRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create warranty request entity too large response has a 4xx status code
func (o *CreateWarrantyRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this create warranty request entity too large response has a 5xx status code
func (o *CreateWarrantyRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this create warranty request entity too large response a status code equal to that given
func (o *CreateWarrantyRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *CreateWarrantyRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/warranty][%d] createWarrantyRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *CreateWarrantyRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/warranty][%d] createWarrantyRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *CreateWarrantyRequestEntityTooLarge) GetPayload() *messaging_models.CreateWarrantyResponse {
	return o.Payload
}

func (o *CreateWarrantyRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(messaging_models.CreateWarrantyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateWarrantyUnsupportedMediaType creates a CreateWarrantyUnsupportedMediaType with default headers values
func NewCreateWarrantyUnsupportedMediaType() *CreateWarrantyUnsupportedMediaType {
	return &CreateWarrantyUnsupportedMediaType{}
}

/*
CreateWarrantyUnsupportedMediaType describes a response with status code 415, with default header values.

The request payload is in an unsupported format.
*/
type CreateWarrantyUnsupportedMediaType struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *messaging_models.CreateWarrantyResponse
}

// IsSuccess returns true when this create warranty unsupported media type response has a 2xx status code
func (o *CreateWarrantyUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create warranty unsupported media type response has a 3xx status code
func (o *CreateWarrantyUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create warranty unsupported media type response has a 4xx status code
func (o *CreateWarrantyUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this create warranty unsupported media type response has a 5xx status code
func (o *CreateWarrantyUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this create warranty unsupported media type response a status code equal to that given
func (o *CreateWarrantyUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *CreateWarrantyUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/warranty][%d] createWarrantyUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *CreateWarrantyUnsupportedMediaType) String() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/warranty][%d] createWarrantyUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *CreateWarrantyUnsupportedMediaType) GetPayload() *messaging_models.CreateWarrantyResponse {
	return o.Payload
}

func (o *CreateWarrantyUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(messaging_models.CreateWarrantyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateWarrantyTooManyRequests creates a CreateWarrantyTooManyRequests with default headers values
func NewCreateWarrantyTooManyRequests() *CreateWarrantyTooManyRequests {
	return &CreateWarrantyTooManyRequests{}
}

/*
CreateWarrantyTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type CreateWarrantyTooManyRequests struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *messaging_models.CreateWarrantyResponse
}

// IsSuccess returns true when this create warranty too many requests response has a 2xx status code
func (o *CreateWarrantyTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create warranty too many requests response has a 3xx status code
func (o *CreateWarrantyTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create warranty too many requests response has a 4xx status code
func (o *CreateWarrantyTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this create warranty too many requests response has a 5xx status code
func (o *CreateWarrantyTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this create warranty too many requests response a status code equal to that given
func (o *CreateWarrantyTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *CreateWarrantyTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/warranty][%d] createWarrantyTooManyRequests  %+v", 429, o.Payload)
}

func (o *CreateWarrantyTooManyRequests) String() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/warranty][%d] createWarrantyTooManyRequests  %+v", 429, o.Payload)
}

func (o *CreateWarrantyTooManyRequests) GetPayload() *messaging_models.CreateWarrantyResponse {
	return o.Payload
}

func (o *CreateWarrantyTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(messaging_models.CreateWarrantyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateWarrantyInternalServerError creates a CreateWarrantyInternalServerError with default headers values
func NewCreateWarrantyInternalServerError() *CreateWarrantyInternalServerError {
	return &CreateWarrantyInternalServerError{}
}

/*
CreateWarrantyInternalServerError describes a response with status code 500, with default header values.

An unexpected condition occurred that prevented the server from fulfilling the request.
*/
type CreateWarrantyInternalServerError struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *messaging_models.CreateWarrantyResponse
}

// IsSuccess returns true when this create warranty internal server error response has a 2xx status code
func (o *CreateWarrantyInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create warranty internal server error response has a 3xx status code
func (o *CreateWarrantyInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create warranty internal server error response has a 4xx status code
func (o *CreateWarrantyInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create warranty internal server error response has a 5xx status code
func (o *CreateWarrantyInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create warranty internal server error response a status code equal to that given
func (o *CreateWarrantyInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *CreateWarrantyInternalServerError) Error() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/warranty][%d] createWarrantyInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateWarrantyInternalServerError) String() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/warranty][%d] createWarrantyInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateWarrantyInternalServerError) GetPayload() *messaging_models.CreateWarrantyResponse {
	return o.Payload
}

func (o *CreateWarrantyInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(messaging_models.CreateWarrantyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateWarrantyServiceUnavailable creates a CreateWarrantyServiceUnavailable with default headers values
func NewCreateWarrantyServiceUnavailable() *CreateWarrantyServiceUnavailable {
	return &CreateWarrantyServiceUnavailable{}
}

/*
CreateWarrantyServiceUnavailable describes a response with status code 503, with default header values.

Temporary overloading or maintenance of the server.
*/
type CreateWarrantyServiceUnavailable struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *messaging_models.CreateWarrantyResponse
}

// IsSuccess returns true when this create warranty service unavailable response has a 2xx status code
func (o *CreateWarrantyServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create warranty service unavailable response has a 3xx status code
func (o *CreateWarrantyServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create warranty service unavailable response has a 4xx status code
func (o *CreateWarrantyServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this create warranty service unavailable response has a 5xx status code
func (o *CreateWarrantyServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this create warranty service unavailable response a status code equal to that given
func (o *CreateWarrantyServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *CreateWarrantyServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/warranty][%d] createWarrantyServiceUnavailable  %+v", 503, o.Payload)
}

func (o *CreateWarrantyServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/warranty][%d] createWarrantyServiceUnavailable  %+v", 503, o.Payload)
}

func (o *CreateWarrantyServiceUnavailable) GetPayload() *messaging_models.CreateWarrantyResponse {
	return o.Payload
}

func (o *CreateWarrantyServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(messaging_models.CreateWarrantyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
