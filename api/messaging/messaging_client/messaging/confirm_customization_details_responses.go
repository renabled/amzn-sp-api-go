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

// ConfirmCustomizationDetailsReader is a Reader for the ConfirmCustomizationDetails structure.
type ConfirmCustomizationDetailsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ConfirmCustomizationDetailsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewConfirmCustomizationDetailsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewConfirmCustomizationDetailsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewConfirmCustomizationDetailsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewConfirmCustomizationDetailsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewConfirmCustomizationDetailsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewConfirmCustomizationDetailsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewConfirmCustomizationDetailsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewConfirmCustomizationDetailsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewConfirmCustomizationDetailsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewConfirmCustomizationDetailsCreated creates a ConfirmCustomizationDetailsCreated with default headers values
func NewConfirmCustomizationDetailsCreated() *ConfirmCustomizationDetailsCreated {
	return &ConfirmCustomizationDetailsCreated{}
}

/*
ConfirmCustomizationDetailsCreated describes a response with status code 201, with default header values.

The message was created for the order.
*/
type ConfirmCustomizationDetailsCreated struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *messaging_models.CreateConfirmCustomizationDetailsResponse
}

// IsSuccess returns true when this confirm customization details created response has a 2xx status code
func (o *ConfirmCustomizationDetailsCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this confirm customization details created response has a 3xx status code
func (o *ConfirmCustomizationDetailsCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this confirm customization details created response has a 4xx status code
func (o *ConfirmCustomizationDetailsCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this confirm customization details created response has a 5xx status code
func (o *ConfirmCustomizationDetailsCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this confirm customization details created response a status code equal to that given
func (o *ConfirmCustomizationDetailsCreated) IsCode(code int) bool {
	return code == 201
}

func (o *ConfirmCustomizationDetailsCreated) Error() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/confirmCustomizationDetails][%d] confirmCustomizationDetailsCreated  %+v", 201, o.Payload)
}

func (o *ConfirmCustomizationDetailsCreated) String() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/confirmCustomizationDetails][%d] confirmCustomizationDetailsCreated  %+v", 201, o.Payload)
}

func (o *ConfirmCustomizationDetailsCreated) GetPayload() *messaging_models.CreateConfirmCustomizationDetailsResponse {
	return o.Payload
}

func (o *ConfirmCustomizationDetailsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(messaging_models.CreateConfirmCustomizationDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConfirmCustomizationDetailsBadRequest creates a ConfirmCustomizationDetailsBadRequest with default headers values
func NewConfirmCustomizationDetailsBadRequest() *ConfirmCustomizationDetailsBadRequest {
	return &ConfirmCustomizationDetailsBadRequest{}
}

/*
ConfirmCustomizationDetailsBadRequest describes a response with status code 400, with default header values.

Request has missing or invalid parameters and cannot be parsed.
*/
type ConfirmCustomizationDetailsBadRequest struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *messaging_models.CreateConfirmCustomizationDetailsResponse
}

// IsSuccess returns true when this confirm customization details bad request response has a 2xx status code
func (o *ConfirmCustomizationDetailsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this confirm customization details bad request response has a 3xx status code
func (o *ConfirmCustomizationDetailsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this confirm customization details bad request response has a 4xx status code
func (o *ConfirmCustomizationDetailsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this confirm customization details bad request response has a 5xx status code
func (o *ConfirmCustomizationDetailsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this confirm customization details bad request response a status code equal to that given
func (o *ConfirmCustomizationDetailsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *ConfirmCustomizationDetailsBadRequest) Error() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/confirmCustomizationDetails][%d] confirmCustomizationDetailsBadRequest  %+v", 400, o.Payload)
}

func (o *ConfirmCustomizationDetailsBadRequest) String() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/confirmCustomizationDetails][%d] confirmCustomizationDetailsBadRequest  %+v", 400, o.Payload)
}

func (o *ConfirmCustomizationDetailsBadRequest) GetPayload() *messaging_models.CreateConfirmCustomizationDetailsResponse {
	return o.Payload
}

func (o *ConfirmCustomizationDetailsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(messaging_models.CreateConfirmCustomizationDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConfirmCustomizationDetailsForbidden creates a ConfirmCustomizationDetailsForbidden with default headers values
func NewConfirmCustomizationDetailsForbidden() *ConfirmCustomizationDetailsForbidden {
	return &ConfirmCustomizationDetailsForbidden{}
}

/*
ConfirmCustomizationDetailsForbidden describes a response with status code 403, with default header values.

Indicates that access to the resource is forbidden. Possible reasons include Access Denied, Unauthorized, Expired Token, or Invalid Signature.
*/
type ConfirmCustomizationDetailsForbidden struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *messaging_models.CreateConfirmCustomizationDetailsResponse
}

// IsSuccess returns true when this confirm customization details forbidden response has a 2xx status code
func (o *ConfirmCustomizationDetailsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this confirm customization details forbidden response has a 3xx status code
func (o *ConfirmCustomizationDetailsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this confirm customization details forbidden response has a 4xx status code
func (o *ConfirmCustomizationDetailsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this confirm customization details forbidden response has a 5xx status code
func (o *ConfirmCustomizationDetailsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this confirm customization details forbidden response a status code equal to that given
func (o *ConfirmCustomizationDetailsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ConfirmCustomizationDetailsForbidden) Error() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/confirmCustomizationDetails][%d] confirmCustomizationDetailsForbidden  %+v", 403, o.Payload)
}

func (o *ConfirmCustomizationDetailsForbidden) String() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/confirmCustomizationDetails][%d] confirmCustomizationDetailsForbidden  %+v", 403, o.Payload)
}

func (o *ConfirmCustomizationDetailsForbidden) GetPayload() *messaging_models.CreateConfirmCustomizationDetailsResponse {
	return o.Payload
}

func (o *ConfirmCustomizationDetailsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(messaging_models.CreateConfirmCustomizationDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConfirmCustomizationDetailsNotFound creates a ConfirmCustomizationDetailsNotFound with default headers values
func NewConfirmCustomizationDetailsNotFound() *ConfirmCustomizationDetailsNotFound {
	return &ConfirmCustomizationDetailsNotFound{}
}

/*
ConfirmCustomizationDetailsNotFound describes a response with status code 404, with default header values.

The resource specified does not exist.
*/
type ConfirmCustomizationDetailsNotFound struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *messaging_models.CreateConfirmCustomizationDetailsResponse
}

// IsSuccess returns true when this confirm customization details not found response has a 2xx status code
func (o *ConfirmCustomizationDetailsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this confirm customization details not found response has a 3xx status code
func (o *ConfirmCustomizationDetailsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this confirm customization details not found response has a 4xx status code
func (o *ConfirmCustomizationDetailsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this confirm customization details not found response has a 5xx status code
func (o *ConfirmCustomizationDetailsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this confirm customization details not found response a status code equal to that given
func (o *ConfirmCustomizationDetailsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *ConfirmCustomizationDetailsNotFound) Error() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/confirmCustomizationDetails][%d] confirmCustomizationDetailsNotFound  %+v", 404, o.Payload)
}

func (o *ConfirmCustomizationDetailsNotFound) String() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/confirmCustomizationDetails][%d] confirmCustomizationDetailsNotFound  %+v", 404, o.Payload)
}

func (o *ConfirmCustomizationDetailsNotFound) GetPayload() *messaging_models.CreateConfirmCustomizationDetailsResponse {
	return o.Payload
}

func (o *ConfirmCustomizationDetailsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(messaging_models.CreateConfirmCustomizationDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConfirmCustomizationDetailsRequestEntityTooLarge creates a ConfirmCustomizationDetailsRequestEntityTooLarge with default headers values
func NewConfirmCustomizationDetailsRequestEntityTooLarge() *ConfirmCustomizationDetailsRequestEntityTooLarge {
	return &ConfirmCustomizationDetailsRequestEntityTooLarge{}
}

/*
ConfirmCustomizationDetailsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request size exceeded the maximum accepted size.
*/
type ConfirmCustomizationDetailsRequestEntityTooLarge struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *messaging_models.CreateConfirmCustomizationDetailsResponse
}

// IsSuccess returns true when this confirm customization details request entity too large response has a 2xx status code
func (o *ConfirmCustomizationDetailsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this confirm customization details request entity too large response has a 3xx status code
func (o *ConfirmCustomizationDetailsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this confirm customization details request entity too large response has a 4xx status code
func (o *ConfirmCustomizationDetailsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this confirm customization details request entity too large response has a 5xx status code
func (o *ConfirmCustomizationDetailsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this confirm customization details request entity too large response a status code equal to that given
func (o *ConfirmCustomizationDetailsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *ConfirmCustomizationDetailsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/confirmCustomizationDetails][%d] confirmCustomizationDetailsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *ConfirmCustomizationDetailsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/confirmCustomizationDetails][%d] confirmCustomizationDetailsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *ConfirmCustomizationDetailsRequestEntityTooLarge) GetPayload() *messaging_models.CreateConfirmCustomizationDetailsResponse {
	return o.Payload
}

func (o *ConfirmCustomizationDetailsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(messaging_models.CreateConfirmCustomizationDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConfirmCustomizationDetailsUnsupportedMediaType creates a ConfirmCustomizationDetailsUnsupportedMediaType with default headers values
func NewConfirmCustomizationDetailsUnsupportedMediaType() *ConfirmCustomizationDetailsUnsupportedMediaType {
	return &ConfirmCustomizationDetailsUnsupportedMediaType{}
}

/*
ConfirmCustomizationDetailsUnsupportedMediaType describes a response with status code 415, with default header values.

The request payload is in an unsupported format.
*/
type ConfirmCustomizationDetailsUnsupportedMediaType struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *messaging_models.CreateConfirmCustomizationDetailsResponse
}

// IsSuccess returns true when this confirm customization details unsupported media type response has a 2xx status code
func (o *ConfirmCustomizationDetailsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this confirm customization details unsupported media type response has a 3xx status code
func (o *ConfirmCustomizationDetailsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this confirm customization details unsupported media type response has a 4xx status code
func (o *ConfirmCustomizationDetailsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this confirm customization details unsupported media type response has a 5xx status code
func (o *ConfirmCustomizationDetailsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this confirm customization details unsupported media type response a status code equal to that given
func (o *ConfirmCustomizationDetailsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *ConfirmCustomizationDetailsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/confirmCustomizationDetails][%d] confirmCustomizationDetailsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *ConfirmCustomizationDetailsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/confirmCustomizationDetails][%d] confirmCustomizationDetailsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *ConfirmCustomizationDetailsUnsupportedMediaType) GetPayload() *messaging_models.CreateConfirmCustomizationDetailsResponse {
	return o.Payload
}

func (o *ConfirmCustomizationDetailsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(messaging_models.CreateConfirmCustomizationDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConfirmCustomizationDetailsTooManyRequests creates a ConfirmCustomizationDetailsTooManyRequests with default headers values
func NewConfirmCustomizationDetailsTooManyRequests() *ConfirmCustomizationDetailsTooManyRequests {
	return &ConfirmCustomizationDetailsTooManyRequests{}
}

/*
ConfirmCustomizationDetailsTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type ConfirmCustomizationDetailsTooManyRequests struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *messaging_models.CreateConfirmCustomizationDetailsResponse
}

// IsSuccess returns true when this confirm customization details too many requests response has a 2xx status code
func (o *ConfirmCustomizationDetailsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this confirm customization details too many requests response has a 3xx status code
func (o *ConfirmCustomizationDetailsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this confirm customization details too many requests response has a 4xx status code
func (o *ConfirmCustomizationDetailsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this confirm customization details too many requests response has a 5xx status code
func (o *ConfirmCustomizationDetailsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this confirm customization details too many requests response a status code equal to that given
func (o *ConfirmCustomizationDetailsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *ConfirmCustomizationDetailsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/confirmCustomizationDetails][%d] confirmCustomizationDetailsTooManyRequests  %+v", 429, o.Payload)
}

func (o *ConfirmCustomizationDetailsTooManyRequests) String() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/confirmCustomizationDetails][%d] confirmCustomizationDetailsTooManyRequests  %+v", 429, o.Payload)
}

func (o *ConfirmCustomizationDetailsTooManyRequests) GetPayload() *messaging_models.CreateConfirmCustomizationDetailsResponse {
	return o.Payload
}

func (o *ConfirmCustomizationDetailsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(messaging_models.CreateConfirmCustomizationDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConfirmCustomizationDetailsInternalServerError creates a ConfirmCustomizationDetailsInternalServerError with default headers values
func NewConfirmCustomizationDetailsInternalServerError() *ConfirmCustomizationDetailsInternalServerError {
	return &ConfirmCustomizationDetailsInternalServerError{}
}

/*
ConfirmCustomizationDetailsInternalServerError describes a response with status code 500, with default header values.

An unexpected condition occurred that prevented the server from fulfilling the request.
*/
type ConfirmCustomizationDetailsInternalServerError struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *messaging_models.CreateConfirmCustomizationDetailsResponse
}

// IsSuccess returns true when this confirm customization details internal server error response has a 2xx status code
func (o *ConfirmCustomizationDetailsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this confirm customization details internal server error response has a 3xx status code
func (o *ConfirmCustomizationDetailsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this confirm customization details internal server error response has a 4xx status code
func (o *ConfirmCustomizationDetailsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this confirm customization details internal server error response has a 5xx status code
func (o *ConfirmCustomizationDetailsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this confirm customization details internal server error response a status code equal to that given
func (o *ConfirmCustomizationDetailsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ConfirmCustomizationDetailsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/confirmCustomizationDetails][%d] confirmCustomizationDetailsInternalServerError  %+v", 500, o.Payload)
}

func (o *ConfirmCustomizationDetailsInternalServerError) String() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/confirmCustomizationDetails][%d] confirmCustomizationDetailsInternalServerError  %+v", 500, o.Payload)
}

func (o *ConfirmCustomizationDetailsInternalServerError) GetPayload() *messaging_models.CreateConfirmCustomizationDetailsResponse {
	return o.Payload
}

func (o *ConfirmCustomizationDetailsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(messaging_models.CreateConfirmCustomizationDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConfirmCustomizationDetailsServiceUnavailable creates a ConfirmCustomizationDetailsServiceUnavailable with default headers values
func NewConfirmCustomizationDetailsServiceUnavailable() *ConfirmCustomizationDetailsServiceUnavailable {
	return &ConfirmCustomizationDetailsServiceUnavailable{}
}

/*
ConfirmCustomizationDetailsServiceUnavailable describes a response with status code 503, with default header values.

Temporary overloading or maintenance of the server.
*/
type ConfirmCustomizationDetailsServiceUnavailable struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *messaging_models.CreateConfirmCustomizationDetailsResponse
}

// IsSuccess returns true when this confirm customization details service unavailable response has a 2xx status code
func (o *ConfirmCustomizationDetailsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this confirm customization details service unavailable response has a 3xx status code
func (o *ConfirmCustomizationDetailsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this confirm customization details service unavailable response has a 4xx status code
func (o *ConfirmCustomizationDetailsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this confirm customization details service unavailable response has a 5xx status code
func (o *ConfirmCustomizationDetailsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this confirm customization details service unavailable response a status code equal to that given
func (o *ConfirmCustomizationDetailsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *ConfirmCustomizationDetailsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/confirmCustomizationDetails][%d] confirmCustomizationDetailsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *ConfirmCustomizationDetailsServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /messaging/v1/orders/{amazonOrderId}/messages/confirmCustomizationDetails][%d] confirmCustomizationDetailsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *ConfirmCustomizationDetailsServiceUnavailable) GetPayload() *messaging_models.CreateConfirmCustomizationDetailsResponse {
	return o.Payload
}

func (o *ConfirmCustomizationDetailsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(messaging_models.CreateConfirmCustomizationDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
