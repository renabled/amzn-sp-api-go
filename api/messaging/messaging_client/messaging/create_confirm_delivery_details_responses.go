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

// CreateConfirmDeliveryDetailsReader is a Reader for the CreateConfirmDeliveryDetails structure.
type CreateConfirmDeliveryDetailsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateConfirmDeliveryDetailsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateConfirmDeliveryDetailsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateConfirmDeliveryDetailsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateConfirmDeliveryDetailsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateConfirmDeliveryDetailsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewCreateConfirmDeliveryDetailsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewCreateConfirmDeliveryDetailsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewCreateConfirmDeliveryDetailsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateConfirmDeliveryDetailsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewCreateConfirmDeliveryDetailsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateConfirmDeliveryDetailsCreated creates a CreateConfirmDeliveryDetailsCreated with default headers values
func NewCreateConfirmDeliveryDetailsCreated() *CreateConfirmDeliveryDetailsCreated {
	return &CreateConfirmDeliveryDetailsCreated{}
}

/*
CreateConfirmDeliveryDetailsCreated describes a response with status code 201, with default header values.

The message was created for the order.
*/
type CreateConfirmDeliveryDetailsCreated struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *messaging_models.CreateConfirmDeliveryDetailsResponse
}

// IsSuccess returns true when this create confirm delivery details created response has a 2xx status code
func (o *CreateConfirmDeliveryDetailsCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create confirm delivery details created response has a 3xx status code
func (o *CreateConfirmDeliveryDetailsCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create confirm delivery details created response has a 4xx status code
func (o *CreateConfirmDeliveryDetailsCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create confirm delivery details created response has a 5xx status code
func (o *CreateConfirmDeliveryDetailsCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create confirm delivery details created response a status code equal to that given
func (o *CreateConfirmDeliveryDetailsCreated) IsCode(code int) bool {
	return code == 201
}

func (o *CreateConfirmDeliveryDetailsCreated) Error() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/confirmDeliveryDetails][%d] createConfirmDeliveryDetailsCreated  %+v", 201, o.Payload)
}

func (o *CreateConfirmDeliveryDetailsCreated) String() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/confirmDeliveryDetails][%d] createConfirmDeliveryDetailsCreated  %+v", 201, o.Payload)
}

func (o *CreateConfirmDeliveryDetailsCreated) GetPayload() *messaging_models.CreateConfirmDeliveryDetailsResponse {
	return o.Payload
}

func (o *CreateConfirmDeliveryDetailsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(messaging_models.CreateConfirmDeliveryDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateConfirmDeliveryDetailsBadRequest creates a CreateConfirmDeliveryDetailsBadRequest with default headers values
func NewCreateConfirmDeliveryDetailsBadRequest() *CreateConfirmDeliveryDetailsBadRequest {
	return &CreateConfirmDeliveryDetailsBadRequest{}
}

/*
CreateConfirmDeliveryDetailsBadRequest describes a response with status code 400, with default header values.

Request has missing or invalid parameters and cannot be parsed.
*/
type CreateConfirmDeliveryDetailsBadRequest struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *messaging_models.CreateConfirmDeliveryDetailsResponse
}

// IsSuccess returns true when this create confirm delivery details bad request response has a 2xx status code
func (o *CreateConfirmDeliveryDetailsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create confirm delivery details bad request response has a 3xx status code
func (o *CreateConfirmDeliveryDetailsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create confirm delivery details bad request response has a 4xx status code
func (o *CreateConfirmDeliveryDetailsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create confirm delivery details bad request response has a 5xx status code
func (o *CreateConfirmDeliveryDetailsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create confirm delivery details bad request response a status code equal to that given
func (o *CreateConfirmDeliveryDetailsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *CreateConfirmDeliveryDetailsBadRequest) Error() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/confirmDeliveryDetails][%d] createConfirmDeliveryDetailsBadRequest  %+v", 400, o.Payload)
}

func (o *CreateConfirmDeliveryDetailsBadRequest) String() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/confirmDeliveryDetails][%d] createConfirmDeliveryDetailsBadRequest  %+v", 400, o.Payload)
}

func (o *CreateConfirmDeliveryDetailsBadRequest) GetPayload() *messaging_models.CreateConfirmDeliveryDetailsResponse {
	return o.Payload
}

func (o *CreateConfirmDeliveryDetailsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(messaging_models.CreateConfirmDeliveryDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateConfirmDeliveryDetailsForbidden creates a CreateConfirmDeliveryDetailsForbidden with default headers values
func NewCreateConfirmDeliveryDetailsForbidden() *CreateConfirmDeliveryDetailsForbidden {
	return &CreateConfirmDeliveryDetailsForbidden{}
}

/*
CreateConfirmDeliveryDetailsForbidden describes a response with status code 403, with default header values.

Indicates that access to the resource is forbidden. Possible reasons include Access Denied, Unauthorized, Expired Token, or Invalid Signature.
*/
type CreateConfirmDeliveryDetailsForbidden struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *messaging_models.CreateConfirmDeliveryDetailsResponse
}

// IsSuccess returns true when this create confirm delivery details forbidden response has a 2xx status code
func (o *CreateConfirmDeliveryDetailsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create confirm delivery details forbidden response has a 3xx status code
func (o *CreateConfirmDeliveryDetailsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create confirm delivery details forbidden response has a 4xx status code
func (o *CreateConfirmDeliveryDetailsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create confirm delivery details forbidden response has a 5xx status code
func (o *CreateConfirmDeliveryDetailsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create confirm delivery details forbidden response a status code equal to that given
func (o *CreateConfirmDeliveryDetailsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *CreateConfirmDeliveryDetailsForbidden) Error() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/confirmDeliveryDetails][%d] createConfirmDeliveryDetailsForbidden  %+v", 403, o.Payload)
}

func (o *CreateConfirmDeliveryDetailsForbidden) String() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/confirmDeliveryDetails][%d] createConfirmDeliveryDetailsForbidden  %+v", 403, o.Payload)
}

func (o *CreateConfirmDeliveryDetailsForbidden) GetPayload() *messaging_models.CreateConfirmDeliveryDetailsResponse {
	return o.Payload
}

func (o *CreateConfirmDeliveryDetailsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(messaging_models.CreateConfirmDeliveryDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateConfirmDeliveryDetailsNotFound creates a CreateConfirmDeliveryDetailsNotFound with default headers values
func NewCreateConfirmDeliveryDetailsNotFound() *CreateConfirmDeliveryDetailsNotFound {
	return &CreateConfirmDeliveryDetailsNotFound{}
}

/*
CreateConfirmDeliveryDetailsNotFound describes a response with status code 404, with default header values.

The resource specified does not exist.
*/
type CreateConfirmDeliveryDetailsNotFound struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *messaging_models.CreateConfirmDeliveryDetailsResponse
}

// IsSuccess returns true when this create confirm delivery details not found response has a 2xx status code
func (o *CreateConfirmDeliveryDetailsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create confirm delivery details not found response has a 3xx status code
func (o *CreateConfirmDeliveryDetailsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create confirm delivery details not found response has a 4xx status code
func (o *CreateConfirmDeliveryDetailsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this create confirm delivery details not found response has a 5xx status code
func (o *CreateConfirmDeliveryDetailsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this create confirm delivery details not found response a status code equal to that given
func (o *CreateConfirmDeliveryDetailsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *CreateConfirmDeliveryDetailsNotFound) Error() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/confirmDeliveryDetails][%d] createConfirmDeliveryDetailsNotFound  %+v", 404, o.Payload)
}

func (o *CreateConfirmDeliveryDetailsNotFound) String() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/confirmDeliveryDetails][%d] createConfirmDeliveryDetailsNotFound  %+v", 404, o.Payload)
}

func (o *CreateConfirmDeliveryDetailsNotFound) GetPayload() *messaging_models.CreateConfirmDeliveryDetailsResponse {
	return o.Payload
}

func (o *CreateConfirmDeliveryDetailsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(messaging_models.CreateConfirmDeliveryDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateConfirmDeliveryDetailsRequestEntityTooLarge creates a CreateConfirmDeliveryDetailsRequestEntityTooLarge with default headers values
func NewCreateConfirmDeliveryDetailsRequestEntityTooLarge() *CreateConfirmDeliveryDetailsRequestEntityTooLarge {
	return &CreateConfirmDeliveryDetailsRequestEntityTooLarge{}
}

/*
CreateConfirmDeliveryDetailsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request size exceeded the maximum accepted size.
*/
type CreateConfirmDeliveryDetailsRequestEntityTooLarge struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *messaging_models.CreateConfirmDeliveryDetailsResponse
}

// IsSuccess returns true when this create confirm delivery details request entity too large response has a 2xx status code
func (o *CreateConfirmDeliveryDetailsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create confirm delivery details request entity too large response has a 3xx status code
func (o *CreateConfirmDeliveryDetailsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create confirm delivery details request entity too large response has a 4xx status code
func (o *CreateConfirmDeliveryDetailsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this create confirm delivery details request entity too large response has a 5xx status code
func (o *CreateConfirmDeliveryDetailsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this create confirm delivery details request entity too large response a status code equal to that given
func (o *CreateConfirmDeliveryDetailsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *CreateConfirmDeliveryDetailsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/confirmDeliveryDetails][%d] createConfirmDeliveryDetailsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *CreateConfirmDeliveryDetailsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/confirmDeliveryDetails][%d] createConfirmDeliveryDetailsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *CreateConfirmDeliveryDetailsRequestEntityTooLarge) GetPayload() *messaging_models.CreateConfirmDeliveryDetailsResponse {
	return o.Payload
}

func (o *CreateConfirmDeliveryDetailsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(messaging_models.CreateConfirmDeliveryDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateConfirmDeliveryDetailsUnsupportedMediaType creates a CreateConfirmDeliveryDetailsUnsupportedMediaType with default headers values
func NewCreateConfirmDeliveryDetailsUnsupportedMediaType() *CreateConfirmDeliveryDetailsUnsupportedMediaType {
	return &CreateConfirmDeliveryDetailsUnsupportedMediaType{}
}

/*
CreateConfirmDeliveryDetailsUnsupportedMediaType describes a response with status code 415, with default header values.

The request payload is in an unsupported format.
*/
type CreateConfirmDeliveryDetailsUnsupportedMediaType struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *messaging_models.CreateConfirmDeliveryDetailsResponse
}

// IsSuccess returns true when this create confirm delivery details unsupported media type response has a 2xx status code
func (o *CreateConfirmDeliveryDetailsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create confirm delivery details unsupported media type response has a 3xx status code
func (o *CreateConfirmDeliveryDetailsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create confirm delivery details unsupported media type response has a 4xx status code
func (o *CreateConfirmDeliveryDetailsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this create confirm delivery details unsupported media type response has a 5xx status code
func (o *CreateConfirmDeliveryDetailsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this create confirm delivery details unsupported media type response a status code equal to that given
func (o *CreateConfirmDeliveryDetailsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *CreateConfirmDeliveryDetailsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/confirmDeliveryDetails][%d] createConfirmDeliveryDetailsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *CreateConfirmDeliveryDetailsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/confirmDeliveryDetails][%d] createConfirmDeliveryDetailsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *CreateConfirmDeliveryDetailsUnsupportedMediaType) GetPayload() *messaging_models.CreateConfirmDeliveryDetailsResponse {
	return o.Payload
}

func (o *CreateConfirmDeliveryDetailsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(messaging_models.CreateConfirmDeliveryDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateConfirmDeliveryDetailsTooManyRequests creates a CreateConfirmDeliveryDetailsTooManyRequests with default headers values
func NewCreateConfirmDeliveryDetailsTooManyRequests() *CreateConfirmDeliveryDetailsTooManyRequests {
	return &CreateConfirmDeliveryDetailsTooManyRequests{}
}

/*
CreateConfirmDeliveryDetailsTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type CreateConfirmDeliveryDetailsTooManyRequests struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *messaging_models.CreateConfirmDeliveryDetailsResponse
}

// IsSuccess returns true when this create confirm delivery details too many requests response has a 2xx status code
func (o *CreateConfirmDeliveryDetailsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create confirm delivery details too many requests response has a 3xx status code
func (o *CreateConfirmDeliveryDetailsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create confirm delivery details too many requests response has a 4xx status code
func (o *CreateConfirmDeliveryDetailsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this create confirm delivery details too many requests response has a 5xx status code
func (o *CreateConfirmDeliveryDetailsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this create confirm delivery details too many requests response a status code equal to that given
func (o *CreateConfirmDeliveryDetailsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *CreateConfirmDeliveryDetailsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/confirmDeliveryDetails][%d] createConfirmDeliveryDetailsTooManyRequests  %+v", 429, o.Payload)
}

func (o *CreateConfirmDeliveryDetailsTooManyRequests) String() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/confirmDeliveryDetails][%d] createConfirmDeliveryDetailsTooManyRequests  %+v", 429, o.Payload)
}

func (o *CreateConfirmDeliveryDetailsTooManyRequests) GetPayload() *messaging_models.CreateConfirmDeliveryDetailsResponse {
	return o.Payload
}

func (o *CreateConfirmDeliveryDetailsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(messaging_models.CreateConfirmDeliveryDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateConfirmDeliveryDetailsInternalServerError creates a CreateConfirmDeliveryDetailsInternalServerError with default headers values
func NewCreateConfirmDeliveryDetailsInternalServerError() *CreateConfirmDeliveryDetailsInternalServerError {
	return &CreateConfirmDeliveryDetailsInternalServerError{}
}

/*
CreateConfirmDeliveryDetailsInternalServerError describes a response with status code 500, with default header values.

An unexpected condition occurred that prevented the server from fulfilling the request.
*/
type CreateConfirmDeliveryDetailsInternalServerError struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *messaging_models.CreateConfirmDeliveryDetailsResponse
}

// IsSuccess returns true when this create confirm delivery details internal server error response has a 2xx status code
func (o *CreateConfirmDeliveryDetailsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create confirm delivery details internal server error response has a 3xx status code
func (o *CreateConfirmDeliveryDetailsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create confirm delivery details internal server error response has a 4xx status code
func (o *CreateConfirmDeliveryDetailsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create confirm delivery details internal server error response has a 5xx status code
func (o *CreateConfirmDeliveryDetailsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create confirm delivery details internal server error response a status code equal to that given
func (o *CreateConfirmDeliveryDetailsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *CreateConfirmDeliveryDetailsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/confirmDeliveryDetails][%d] createConfirmDeliveryDetailsInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateConfirmDeliveryDetailsInternalServerError) String() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/confirmDeliveryDetails][%d] createConfirmDeliveryDetailsInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateConfirmDeliveryDetailsInternalServerError) GetPayload() *messaging_models.CreateConfirmDeliveryDetailsResponse {
	return o.Payload
}

func (o *CreateConfirmDeliveryDetailsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(messaging_models.CreateConfirmDeliveryDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateConfirmDeliveryDetailsServiceUnavailable creates a CreateConfirmDeliveryDetailsServiceUnavailable with default headers values
func NewCreateConfirmDeliveryDetailsServiceUnavailable() *CreateConfirmDeliveryDetailsServiceUnavailable {
	return &CreateConfirmDeliveryDetailsServiceUnavailable{}
}

/*
CreateConfirmDeliveryDetailsServiceUnavailable describes a response with status code 503, with default header values.

Temporary overloading or maintenance of the server.
*/
type CreateConfirmDeliveryDetailsServiceUnavailable struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *messaging_models.CreateConfirmDeliveryDetailsResponse
}

// IsSuccess returns true when this create confirm delivery details service unavailable response has a 2xx status code
func (o *CreateConfirmDeliveryDetailsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create confirm delivery details service unavailable response has a 3xx status code
func (o *CreateConfirmDeliveryDetailsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create confirm delivery details service unavailable response has a 4xx status code
func (o *CreateConfirmDeliveryDetailsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this create confirm delivery details service unavailable response has a 5xx status code
func (o *CreateConfirmDeliveryDetailsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this create confirm delivery details service unavailable response a status code equal to that given
func (o *CreateConfirmDeliveryDetailsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *CreateConfirmDeliveryDetailsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/confirmDeliveryDetails][%d] createConfirmDeliveryDetailsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *CreateConfirmDeliveryDetailsServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/confirmDeliveryDetails][%d] createConfirmDeliveryDetailsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *CreateConfirmDeliveryDetailsServiceUnavailable) GetPayload() *messaging_models.CreateConfirmDeliveryDetailsResponse {
	return o.Payload
}

func (o *CreateConfirmDeliveryDetailsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(messaging_models.CreateConfirmDeliveryDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
