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

// GetUnmanifestedShipmentsReader is a Reader for the GetUnmanifestedShipments structure.
type GetUnmanifestedShipmentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUnmanifestedShipmentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUnmanifestedShipmentsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetUnmanifestedShipmentsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetUnmanifestedShipmentsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetUnmanifestedShipmentsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetUnmanifestedShipmentsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetUnmanifestedShipmentsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetUnmanifestedShipmentsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetUnmanifestedShipmentsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetUnmanifestedShipmentsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetUnmanifestedShipmentsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetUnmanifestedShipmentsOK creates a GetUnmanifestedShipmentsOK with default headers values
func NewGetUnmanifestedShipmentsOK() *GetUnmanifestedShipmentsOK {
	return &GetUnmanifestedShipmentsOK{}
}

/*
GetUnmanifestedShipmentsOK describes a response with status code 200, with default header values.

Success.
*/
type GetUnmanifestedShipmentsOK struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *shipping_v2_models.GetUnmanifestedShipmentsResponse
}

// IsSuccess returns true when this get unmanifested shipments o k response has a 2xx status code
func (o *GetUnmanifestedShipmentsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get unmanifested shipments o k response has a 3xx status code
func (o *GetUnmanifestedShipmentsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get unmanifested shipments o k response has a 4xx status code
func (o *GetUnmanifestedShipmentsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get unmanifested shipments o k response has a 5xx status code
func (o *GetUnmanifestedShipmentsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get unmanifested shipments o k response a status code equal to that given
func (o *GetUnmanifestedShipmentsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetUnmanifestedShipmentsOK) Error() string {
	return fmt.Sprintf("[PUT /shipping/v2/unmanifestedShipments][%d] getUnmanifestedShipmentsOK  %+v", 200, o.Payload)
}

func (o *GetUnmanifestedShipmentsOK) String() string {
	return fmt.Sprintf("[PUT /shipping/v2/unmanifestedShipments][%d] getUnmanifestedShipmentsOK  %+v", 200, o.Payload)
}

func (o *GetUnmanifestedShipmentsOK) GetPayload() *shipping_v2_models.GetUnmanifestedShipmentsResponse {
	return o.Payload
}

func (o *GetUnmanifestedShipmentsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(shipping_v2_models.GetUnmanifestedShipmentsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUnmanifestedShipmentsBadRequest creates a GetUnmanifestedShipmentsBadRequest with default headers values
func NewGetUnmanifestedShipmentsBadRequest() *GetUnmanifestedShipmentsBadRequest {
	return &GetUnmanifestedShipmentsBadRequest{}
}

/*
GetUnmanifestedShipmentsBadRequest describes a response with status code 400, with default header values.

Request has missing or invalid parameters and cannot be parsed.
*/
type GetUnmanifestedShipmentsBadRequest struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *shipping_v2_models.ErrorList
}

// IsSuccess returns true when this get unmanifested shipments bad request response has a 2xx status code
func (o *GetUnmanifestedShipmentsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get unmanifested shipments bad request response has a 3xx status code
func (o *GetUnmanifestedShipmentsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get unmanifested shipments bad request response has a 4xx status code
func (o *GetUnmanifestedShipmentsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get unmanifested shipments bad request response has a 5xx status code
func (o *GetUnmanifestedShipmentsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get unmanifested shipments bad request response a status code equal to that given
func (o *GetUnmanifestedShipmentsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetUnmanifestedShipmentsBadRequest) Error() string {
	return fmt.Sprintf("[PUT /shipping/v2/unmanifestedShipments][%d] getUnmanifestedShipmentsBadRequest  %+v", 400, o.Payload)
}

func (o *GetUnmanifestedShipmentsBadRequest) String() string {
	return fmt.Sprintf("[PUT /shipping/v2/unmanifestedShipments][%d] getUnmanifestedShipmentsBadRequest  %+v", 400, o.Payload)
}

func (o *GetUnmanifestedShipmentsBadRequest) GetPayload() *shipping_v2_models.ErrorList {
	return o.Payload
}

func (o *GetUnmanifestedShipmentsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetUnmanifestedShipmentsUnauthorized creates a GetUnmanifestedShipmentsUnauthorized with default headers values
func NewGetUnmanifestedShipmentsUnauthorized() *GetUnmanifestedShipmentsUnauthorized {
	return &GetUnmanifestedShipmentsUnauthorized{}
}

/*
GetUnmanifestedShipmentsUnauthorized describes a response with status code 401, with default header values.

The request's Authorization header is not formatted correctly or does not contain a valid token.
*/
type GetUnmanifestedShipmentsUnauthorized struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *shipping_v2_models.ErrorList
}

// IsSuccess returns true when this get unmanifested shipments unauthorized response has a 2xx status code
func (o *GetUnmanifestedShipmentsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get unmanifested shipments unauthorized response has a 3xx status code
func (o *GetUnmanifestedShipmentsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get unmanifested shipments unauthorized response has a 4xx status code
func (o *GetUnmanifestedShipmentsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get unmanifested shipments unauthorized response has a 5xx status code
func (o *GetUnmanifestedShipmentsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get unmanifested shipments unauthorized response a status code equal to that given
func (o *GetUnmanifestedShipmentsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetUnmanifestedShipmentsUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /shipping/v2/unmanifestedShipments][%d] getUnmanifestedShipmentsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetUnmanifestedShipmentsUnauthorized) String() string {
	return fmt.Sprintf("[PUT /shipping/v2/unmanifestedShipments][%d] getUnmanifestedShipmentsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetUnmanifestedShipmentsUnauthorized) GetPayload() *shipping_v2_models.ErrorList {
	return o.Payload
}

func (o *GetUnmanifestedShipmentsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetUnmanifestedShipmentsForbidden creates a GetUnmanifestedShipmentsForbidden with default headers values
func NewGetUnmanifestedShipmentsForbidden() *GetUnmanifestedShipmentsForbidden {
	return &GetUnmanifestedShipmentsForbidden{}
}

/*
GetUnmanifestedShipmentsForbidden describes a response with status code 403, with default header values.

Indicates that access to the resource is forbidden. Possible reasons include Access Denied, Unauthorized, Expired Token, or Invalid Signature.
*/
type GetUnmanifestedShipmentsForbidden struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *shipping_v2_models.ErrorList
}

// IsSuccess returns true when this get unmanifested shipments forbidden response has a 2xx status code
func (o *GetUnmanifestedShipmentsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get unmanifested shipments forbidden response has a 3xx status code
func (o *GetUnmanifestedShipmentsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get unmanifested shipments forbidden response has a 4xx status code
func (o *GetUnmanifestedShipmentsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get unmanifested shipments forbidden response has a 5xx status code
func (o *GetUnmanifestedShipmentsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get unmanifested shipments forbidden response a status code equal to that given
func (o *GetUnmanifestedShipmentsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetUnmanifestedShipmentsForbidden) Error() string {
	return fmt.Sprintf("[PUT /shipping/v2/unmanifestedShipments][%d] getUnmanifestedShipmentsForbidden  %+v", 403, o.Payload)
}

func (o *GetUnmanifestedShipmentsForbidden) String() string {
	return fmt.Sprintf("[PUT /shipping/v2/unmanifestedShipments][%d] getUnmanifestedShipmentsForbidden  %+v", 403, o.Payload)
}

func (o *GetUnmanifestedShipmentsForbidden) GetPayload() *shipping_v2_models.ErrorList {
	return o.Payload
}

func (o *GetUnmanifestedShipmentsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetUnmanifestedShipmentsNotFound creates a GetUnmanifestedShipmentsNotFound with default headers values
func NewGetUnmanifestedShipmentsNotFound() *GetUnmanifestedShipmentsNotFound {
	return &GetUnmanifestedShipmentsNotFound{}
}

/*
GetUnmanifestedShipmentsNotFound describes a response with status code 404, with default header values.

The resource specified does not exist.
*/
type GetUnmanifestedShipmentsNotFound struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *shipping_v2_models.ErrorList
}

// IsSuccess returns true when this get unmanifested shipments not found response has a 2xx status code
func (o *GetUnmanifestedShipmentsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get unmanifested shipments not found response has a 3xx status code
func (o *GetUnmanifestedShipmentsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get unmanifested shipments not found response has a 4xx status code
func (o *GetUnmanifestedShipmentsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get unmanifested shipments not found response has a 5xx status code
func (o *GetUnmanifestedShipmentsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get unmanifested shipments not found response a status code equal to that given
func (o *GetUnmanifestedShipmentsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetUnmanifestedShipmentsNotFound) Error() string {
	return fmt.Sprintf("[PUT /shipping/v2/unmanifestedShipments][%d] getUnmanifestedShipmentsNotFound  %+v", 404, o.Payload)
}

func (o *GetUnmanifestedShipmentsNotFound) String() string {
	return fmt.Sprintf("[PUT /shipping/v2/unmanifestedShipments][%d] getUnmanifestedShipmentsNotFound  %+v", 404, o.Payload)
}

func (o *GetUnmanifestedShipmentsNotFound) GetPayload() *shipping_v2_models.ErrorList {
	return o.Payload
}

func (o *GetUnmanifestedShipmentsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetUnmanifestedShipmentsRequestEntityTooLarge creates a GetUnmanifestedShipmentsRequestEntityTooLarge with default headers values
func NewGetUnmanifestedShipmentsRequestEntityTooLarge() *GetUnmanifestedShipmentsRequestEntityTooLarge {
	return &GetUnmanifestedShipmentsRequestEntityTooLarge{}
}

/*
GetUnmanifestedShipmentsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request size exceeded the maximum accepted size.
*/
type GetUnmanifestedShipmentsRequestEntityTooLarge struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *shipping_v2_models.ErrorList
}

// IsSuccess returns true when this get unmanifested shipments request entity too large response has a 2xx status code
func (o *GetUnmanifestedShipmentsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get unmanifested shipments request entity too large response has a 3xx status code
func (o *GetUnmanifestedShipmentsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get unmanifested shipments request entity too large response has a 4xx status code
func (o *GetUnmanifestedShipmentsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get unmanifested shipments request entity too large response has a 5xx status code
func (o *GetUnmanifestedShipmentsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get unmanifested shipments request entity too large response a status code equal to that given
func (o *GetUnmanifestedShipmentsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetUnmanifestedShipmentsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PUT /shipping/v2/unmanifestedShipments][%d] getUnmanifestedShipmentsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetUnmanifestedShipmentsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[PUT /shipping/v2/unmanifestedShipments][%d] getUnmanifestedShipmentsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetUnmanifestedShipmentsRequestEntityTooLarge) GetPayload() *shipping_v2_models.ErrorList {
	return o.Payload
}

func (o *GetUnmanifestedShipmentsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetUnmanifestedShipmentsUnsupportedMediaType creates a GetUnmanifestedShipmentsUnsupportedMediaType with default headers values
func NewGetUnmanifestedShipmentsUnsupportedMediaType() *GetUnmanifestedShipmentsUnsupportedMediaType {
	return &GetUnmanifestedShipmentsUnsupportedMediaType{}
}

/*
GetUnmanifestedShipmentsUnsupportedMediaType describes a response with status code 415, with default header values.

The request payload is in an unsupported format.
*/
type GetUnmanifestedShipmentsUnsupportedMediaType struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *shipping_v2_models.ErrorList
}

// IsSuccess returns true when this get unmanifested shipments unsupported media type response has a 2xx status code
func (o *GetUnmanifestedShipmentsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get unmanifested shipments unsupported media type response has a 3xx status code
func (o *GetUnmanifestedShipmentsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get unmanifested shipments unsupported media type response has a 4xx status code
func (o *GetUnmanifestedShipmentsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get unmanifested shipments unsupported media type response has a 5xx status code
func (o *GetUnmanifestedShipmentsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get unmanifested shipments unsupported media type response a status code equal to that given
func (o *GetUnmanifestedShipmentsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetUnmanifestedShipmentsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PUT /shipping/v2/unmanifestedShipments][%d] getUnmanifestedShipmentsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetUnmanifestedShipmentsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[PUT /shipping/v2/unmanifestedShipments][%d] getUnmanifestedShipmentsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetUnmanifestedShipmentsUnsupportedMediaType) GetPayload() *shipping_v2_models.ErrorList {
	return o.Payload
}

func (o *GetUnmanifestedShipmentsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetUnmanifestedShipmentsTooManyRequests creates a GetUnmanifestedShipmentsTooManyRequests with default headers values
func NewGetUnmanifestedShipmentsTooManyRequests() *GetUnmanifestedShipmentsTooManyRequests {
	return &GetUnmanifestedShipmentsTooManyRequests{}
}

/*
GetUnmanifestedShipmentsTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type GetUnmanifestedShipmentsTooManyRequests struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *shipping_v2_models.ErrorList
}

// IsSuccess returns true when this get unmanifested shipments too many requests response has a 2xx status code
func (o *GetUnmanifestedShipmentsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get unmanifested shipments too many requests response has a 3xx status code
func (o *GetUnmanifestedShipmentsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get unmanifested shipments too many requests response has a 4xx status code
func (o *GetUnmanifestedShipmentsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get unmanifested shipments too many requests response has a 5xx status code
func (o *GetUnmanifestedShipmentsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get unmanifested shipments too many requests response a status code equal to that given
func (o *GetUnmanifestedShipmentsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetUnmanifestedShipmentsTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /shipping/v2/unmanifestedShipments][%d] getUnmanifestedShipmentsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetUnmanifestedShipmentsTooManyRequests) String() string {
	return fmt.Sprintf("[PUT /shipping/v2/unmanifestedShipments][%d] getUnmanifestedShipmentsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetUnmanifestedShipmentsTooManyRequests) GetPayload() *shipping_v2_models.ErrorList {
	return o.Payload
}

func (o *GetUnmanifestedShipmentsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetUnmanifestedShipmentsInternalServerError creates a GetUnmanifestedShipmentsInternalServerError with default headers values
func NewGetUnmanifestedShipmentsInternalServerError() *GetUnmanifestedShipmentsInternalServerError {
	return &GetUnmanifestedShipmentsInternalServerError{}
}

/*
GetUnmanifestedShipmentsInternalServerError describes a response with status code 500, with default header values.

An unexpected condition occurred that prevented the server from fulfilling the request.
*/
type GetUnmanifestedShipmentsInternalServerError struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *shipping_v2_models.ErrorList
}

// IsSuccess returns true when this get unmanifested shipments internal server error response has a 2xx status code
func (o *GetUnmanifestedShipmentsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get unmanifested shipments internal server error response has a 3xx status code
func (o *GetUnmanifestedShipmentsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get unmanifested shipments internal server error response has a 4xx status code
func (o *GetUnmanifestedShipmentsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get unmanifested shipments internal server error response has a 5xx status code
func (o *GetUnmanifestedShipmentsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get unmanifested shipments internal server error response a status code equal to that given
func (o *GetUnmanifestedShipmentsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetUnmanifestedShipmentsInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /shipping/v2/unmanifestedShipments][%d] getUnmanifestedShipmentsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUnmanifestedShipmentsInternalServerError) String() string {
	return fmt.Sprintf("[PUT /shipping/v2/unmanifestedShipments][%d] getUnmanifestedShipmentsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUnmanifestedShipmentsInternalServerError) GetPayload() *shipping_v2_models.ErrorList {
	return o.Payload
}

func (o *GetUnmanifestedShipmentsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetUnmanifestedShipmentsServiceUnavailable creates a GetUnmanifestedShipmentsServiceUnavailable with default headers values
func NewGetUnmanifestedShipmentsServiceUnavailable() *GetUnmanifestedShipmentsServiceUnavailable {
	return &GetUnmanifestedShipmentsServiceUnavailable{}
}

/*
GetUnmanifestedShipmentsServiceUnavailable describes a response with status code 503, with default header values.

Temporary overloading or maintenance of the server.
*/
type GetUnmanifestedShipmentsServiceUnavailable struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *shipping_v2_models.ErrorList
}

// IsSuccess returns true when this get unmanifested shipments service unavailable response has a 2xx status code
func (o *GetUnmanifestedShipmentsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get unmanifested shipments service unavailable response has a 3xx status code
func (o *GetUnmanifestedShipmentsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get unmanifested shipments service unavailable response has a 4xx status code
func (o *GetUnmanifestedShipmentsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get unmanifested shipments service unavailable response has a 5xx status code
func (o *GetUnmanifestedShipmentsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get unmanifested shipments service unavailable response a status code equal to that given
func (o *GetUnmanifestedShipmentsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetUnmanifestedShipmentsServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /shipping/v2/unmanifestedShipments][%d] getUnmanifestedShipmentsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetUnmanifestedShipmentsServiceUnavailable) String() string {
	return fmt.Sprintf("[PUT /shipping/v2/unmanifestedShipments][%d] getUnmanifestedShipmentsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetUnmanifestedShipmentsServiceUnavailable) GetPayload() *shipping_v2_models.ErrorList {
	return o.Payload
}

func (o *GetUnmanifestedShipmentsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
