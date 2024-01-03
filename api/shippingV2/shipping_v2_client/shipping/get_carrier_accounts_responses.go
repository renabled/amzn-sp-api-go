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

// GetCarrierAccountsReader is a Reader for the GetCarrierAccounts structure.
type GetCarrierAccountsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCarrierAccountsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCarrierAccountsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetCarrierAccountsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetCarrierAccountsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetCarrierAccountsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetCarrierAccountsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetCarrierAccountsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetCarrierAccountsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetCarrierAccountsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetCarrierAccountsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetCarrierAccountsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetCarrierAccountsOK creates a GetCarrierAccountsOK with default headers values
func NewGetCarrierAccountsOK() *GetCarrierAccountsOK {
	return &GetCarrierAccountsOK{}
}

/*
GetCarrierAccountsOK describes a response with status code 200, with default header values.

Success.
*/
type GetCarrierAccountsOK struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *shipping_v2_models.GetCarrierAccountsResponse
}

// IsSuccess returns true when this get carrier accounts o k response has a 2xx status code
func (o *GetCarrierAccountsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get carrier accounts o k response has a 3xx status code
func (o *GetCarrierAccountsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get carrier accounts o k response has a 4xx status code
func (o *GetCarrierAccountsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get carrier accounts o k response has a 5xx status code
func (o *GetCarrierAccountsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get carrier accounts o k response a status code equal to that given
func (o *GetCarrierAccountsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetCarrierAccountsOK) Error() string {
	return fmt.Sprintf("[PUT /shipping/v2/carrierAccounts][%d] getCarrierAccountsOK  %+v", 200, o.Payload)
}

func (o *GetCarrierAccountsOK) String() string {
	return fmt.Sprintf("[PUT /shipping/v2/carrierAccounts][%d] getCarrierAccountsOK  %+v", 200, o.Payload)
}

func (o *GetCarrierAccountsOK) GetPayload() *shipping_v2_models.GetCarrierAccountsResponse {
	return o.Payload
}

func (o *GetCarrierAccountsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(shipping_v2_models.GetCarrierAccountsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCarrierAccountsBadRequest creates a GetCarrierAccountsBadRequest with default headers values
func NewGetCarrierAccountsBadRequest() *GetCarrierAccountsBadRequest {
	return &GetCarrierAccountsBadRequest{}
}

/*
GetCarrierAccountsBadRequest describes a response with status code 400, with default header values.

Request has missing or invalid parameters and cannot be parsed.
*/
type GetCarrierAccountsBadRequest struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *shipping_v2_models.ErrorList
}

// IsSuccess returns true when this get carrier accounts bad request response has a 2xx status code
func (o *GetCarrierAccountsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get carrier accounts bad request response has a 3xx status code
func (o *GetCarrierAccountsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get carrier accounts bad request response has a 4xx status code
func (o *GetCarrierAccountsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get carrier accounts bad request response has a 5xx status code
func (o *GetCarrierAccountsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get carrier accounts bad request response a status code equal to that given
func (o *GetCarrierAccountsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetCarrierAccountsBadRequest) Error() string {
	return fmt.Sprintf("[PUT /shipping/v2/carrierAccounts][%d] getCarrierAccountsBadRequest  %+v", 400, o.Payload)
}

func (o *GetCarrierAccountsBadRequest) String() string {
	return fmt.Sprintf("[PUT /shipping/v2/carrierAccounts][%d] getCarrierAccountsBadRequest  %+v", 400, o.Payload)
}

func (o *GetCarrierAccountsBadRequest) GetPayload() *shipping_v2_models.ErrorList {
	return o.Payload
}

func (o *GetCarrierAccountsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCarrierAccountsUnauthorized creates a GetCarrierAccountsUnauthorized with default headers values
func NewGetCarrierAccountsUnauthorized() *GetCarrierAccountsUnauthorized {
	return &GetCarrierAccountsUnauthorized{}
}

/*
GetCarrierAccountsUnauthorized describes a response with status code 401, with default header values.

The request's Authorization header is not formatted correctly or does not contain a valid token.
*/
type GetCarrierAccountsUnauthorized struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *shipping_v2_models.ErrorList
}

// IsSuccess returns true when this get carrier accounts unauthorized response has a 2xx status code
func (o *GetCarrierAccountsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get carrier accounts unauthorized response has a 3xx status code
func (o *GetCarrierAccountsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get carrier accounts unauthorized response has a 4xx status code
func (o *GetCarrierAccountsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get carrier accounts unauthorized response has a 5xx status code
func (o *GetCarrierAccountsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get carrier accounts unauthorized response a status code equal to that given
func (o *GetCarrierAccountsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetCarrierAccountsUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /shipping/v2/carrierAccounts][%d] getCarrierAccountsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetCarrierAccountsUnauthorized) String() string {
	return fmt.Sprintf("[PUT /shipping/v2/carrierAccounts][%d] getCarrierAccountsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetCarrierAccountsUnauthorized) GetPayload() *shipping_v2_models.ErrorList {
	return o.Payload
}

func (o *GetCarrierAccountsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCarrierAccountsForbidden creates a GetCarrierAccountsForbidden with default headers values
func NewGetCarrierAccountsForbidden() *GetCarrierAccountsForbidden {
	return &GetCarrierAccountsForbidden{}
}

/*
GetCarrierAccountsForbidden describes a response with status code 403, with default header values.

Indicates that access to the resource is forbidden. Possible reasons include Access Denied, Unauthorized, Expired Token, or Invalid Signature.
*/
type GetCarrierAccountsForbidden struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *shipping_v2_models.ErrorList
}

// IsSuccess returns true when this get carrier accounts forbidden response has a 2xx status code
func (o *GetCarrierAccountsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get carrier accounts forbidden response has a 3xx status code
func (o *GetCarrierAccountsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get carrier accounts forbidden response has a 4xx status code
func (o *GetCarrierAccountsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get carrier accounts forbidden response has a 5xx status code
func (o *GetCarrierAccountsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get carrier accounts forbidden response a status code equal to that given
func (o *GetCarrierAccountsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetCarrierAccountsForbidden) Error() string {
	return fmt.Sprintf("[PUT /shipping/v2/carrierAccounts][%d] getCarrierAccountsForbidden  %+v", 403, o.Payload)
}

func (o *GetCarrierAccountsForbidden) String() string {
	return fmt.Sprintf("[PUT /shipping/v2/carrierAccounts][%d] getCarrierAccountsForbidden  %+v", 403, o.Payload)
}

func (o *GetCarrierAccountsForbidden) GetPayload() *shipping_v2_models.ErrorList {
	return o.Payload
}

func (o *GetCarrierAccountsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCarrierAccountsNotFound creates a GetCarrierAccountsNotFound with default headers values
func NewGetCarrierAccountsNotFound() *GetCarrierAccountsNotFound {
	return &GetCarrierAccountsNotFound{}
}

/*
GetCarrierAccountsNotFound describes a response with status code 404, with default header values.

The resource specified does not exist.
*/
type GetCarrierAccountsNotFound struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *shipping_v2_models.ErrorList
}

// IsSuccess returns true when this get carrier accounts not found response has a 2xx status code
func (o *GetCarrierAccountsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get carrier accounts not found response has a 3xx status code
func (o *GetCarrierAccountsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get carrier accounts not found response has a 4xx status code
func (o *GetCarrierAccountsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get carrier accounts not found response has a 5xx status code
func (o *GetCarrierAccountsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get carrier accounts not found response a status code equal to that given
func (o *GetCarrierAccountsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetCarrierAccountsNotFound) Error() string {
	return fmt.Sprintf("[PUT /shipping/v2/carrierAccounts][%d] getCarrierAccountsNotFound  %+v", 404, o.Payload)
}

func (o *GetCarrierAccountsNotFound) String() string {
	return fmt.Sprintf("[PUT /shipping/v2/carrierAccounts][%d] getCarrierAccountsNotFound  %+v", 404, o.Payload)
}

func (o *GetCarrierAccountsNotFound) GetPayload() *shipping_v2_models.ErrorList {
	return o.Payload
}

func (o *GetCarrierAccountsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCarrierAccountsRequestEntityTooLarge creates a GetCarrierAccountsRequestEntityTooLarge with default headers values
func NewGetCarrierAccountsRequestEntityTooLarge() *GetCarrierAccountsRequestEntityTooLarge {
	return &GetCarrierAccountsRequestEntityTooLarge{}
}

/*
GetCarrierAccountsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request size exceeded the maximum accepted size.
*/
type GetCarrierAccountsRequestEntityTooLarge struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *shipping_v2_models.ErrorList
}

// IsSuccess returns true when this get carrier accounts request entity too large response has a 2xx status code
func (o *GetCarrierAccountsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get carrier accounts request entity too large response has a 3xx status code
func (o *GetCarrierAccountsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get carrier accounts request entity too large response has a 4xx status code
func (o *GetCarrierAccountsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get carrier accounts request entity too large response has a 5xx status code
func (o *GetCarrierAccountsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get carrier accounts request entity too large response a status code equal to that given
func (o *GetCarrierAccountsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetCarrierAccountsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PUT /shipping/v2/carrierAccounts][%d] getCarrierAccountsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetCarrierAccountsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[PUT /shipping/v2/carrierAccounts][%d] getCarrierAccountsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetCarrierAccountsRequestEntityTooLarge) GetPayload() *shipping_v2_models.ErrorList {
	return o.Payload
}

func (o *GetCarrierAccountsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCarrierAccountsUnsupportedMediaType creates a GetCarrierAccountsUnsupportedMediaType with default headers values
func NewGetCarrierAccountsUnsupportedMediaType() *GetCarrierAccountsUnsupportedMediaType {
	return &GetCarrierAccountsUnsupportedMediaType{}
}

/*
GetCarrierAccountsUnsupportedMediaType describes a response with status code 415, with default header values.

The request payload is in an unsupported format.
*/
type GetCarrierAccountsUnsupportedMediaType struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *shipping_v2_models.ErrorList
}

// IsSuccess returns true when this get carrier accounts unsupported media type response has a 2xx status code
func (o *GetCarrierAccountsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get carrier accounts unsupported media type response has a 3xx status code
func (o *GetCarrierAccountsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get carrier accounts unsupported media type response has a 4xx status code
func (o *GetCarrierAccountsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get carrier accounts unsupported media type response has a 5xx status code
func (o *GetCarrierAccountsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get carrier accounts unsupported media type response a status code equal to that given
func (o *GetCarrierAccountsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetCarrierAccountsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PUT /shipping/v2/carrierAccounts][%d] getCarrierAccountsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetCarrierAccountsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[PUT /shipping/v2/carrierAccounts][%d] getCarrierAccountsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetCarrierAccountsUnsupportedMediaType) GetPayload() *shipping_v2_models.ErrorList {
	return o.Payload
}

func (o *GetCarrierAccountsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCarrierAccountsTooManyRequests creates a GetCarrierAccountsTooManyRequests with default headers values
func NewGetCarrierAccountsTooManyRequests() *GetCarrierAccountsTooManyRequests {
	return &GetCarrierAccountsTooManyRequests{}
}

/*
GetCarrierAccountsTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type GetCarrierAccountsTooManyRequests struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *shipping_v2_models.ErrorList
}

// IsSuccess returns true when this get carrier accounts too many requests response has a 2xx status code
func (o *GetCarrierAccountsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get carrier accounts too many requests response has a 3xx status code
func (o *GetCarrierAccountsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get carrier accounts too many requests response has a 4xx status code
func (o *GetCarrierAccountsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get carrier accounts too many requests response has a 5xx status code
func (o *GetCarrierAccountsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get carrier accounts too many requests response a status code equal to that given
func (o *GetCarrierAccountsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetCarrierAccountsTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /shipping/v2/carrierAccounts][%d] getCarrierAccountsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetCarrierAccountsTooManyRequests) String() string {
	return fmt.Sprintf("[PUT /shipping/v2/carrierAccounts][%d] getCarrierAccountsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetCarrierAccountsTooManyRequests) GetPayload() *shipping_v2_models.ErrorList {
	return o.Payload
}

func (o *GetCarrierAccountsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCarrierAccountsInternalServerError creates a GetCarrierAccountsInternalServerError with default headers values
func NewGetCarrierAccountsInternalServerError() *GetCarrierAccountsInternalServerError {
	return &GetCarrierAccountsInternalServerError{}
}

/*
GetCarrierAccountsInternalServerError describes a response with status code 500, with default header values.

An unexpected condition occurred that prevented the server from fulfilling the request.
*/
type GetCarrierAccountsInternalServerError struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *shipping_v2_models.ErrorList
}

// IsSuccess returns true when this get carrier accounts internal server error response has a 2xx status code
func (o *GetCarrierAccountsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get carrier accounts internal server error response has a 3xx status code
func (o *GetCarrierAccountsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get carrier accounts internal server error response has a 4xx status code
func (o *GetCarrierAccountsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get carrier accounts internal server error response has a 5xx status code
func (o *GetCarrierAccountsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get carrier accounts internal server error response a status code equal to that given
func (o *GetCarrierAccountsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetCarrierAccountsInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /shipping/v2/carrierAccounts][%d] getCarrierAccountsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCarrierAccountsInternalServerError) String() string {
	return fmt.Sprintf("[PUT /shipping/v2/carrierAccounts][%d] getCarrierAccountsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCarrierAccountsInternalServerError) GetPayload() *shipping_v2_models.ErrorList {
	return o.Payload
}

func (o *GetCarrierAccountsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCarrierAccountsServiceUnavailable creates a GetCarrierAccountsServiceUnavailable with default headers values
func NewGetCarrierAccountsServiceUnavailable() *GetCarrierAccountsServiceUnavailable {
	return &GetCarrierAccountsServiceUnavailable{}
}

/*
GetCarrierAccountsServiceUnavailable describes a response with status code 503, with default header values.

Temporary overloading or maintenance of the server.
*/
type GetCarrierAccountsServiceUnavailable struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *shipping_v2_models.ErrorList
}

// IsSuccess returns true when this get carrier accounts service unavailable response has a 2xx status code
func (o *GetCarrierAccountsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get carrier accounts service unavailable response has a 3xx status code
func (o *GetCarrierAccountsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get carrier accounts service unavailable response has a 4xx status code
func (o *GetCarrierAccountsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get carrier accounts service unavailable response has a 5xx status code
func (o *GetCarrierAccountsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get carrier accounts service unavailable response a status code equal to that given
func (o *GetCarrierAccountsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetCarrierAccountsServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /shipping/v2/carrierAccounts][%d] getCarrierAccountsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetCarrierAccountsServiceUnavailable) String() string {
	return fmt.Sprintf("[PUT /shipping/v2/carrierAccounts][%d] getCarrierAccountsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetCarrierAccountsServiceUnavailable) GetPayload() *shipping_v2_models.ErrorList {
	return o.Payload
}

func (o *GetCarrierAccountsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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