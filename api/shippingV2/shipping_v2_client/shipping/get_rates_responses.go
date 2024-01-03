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

// GetRatesReader is a Reader for the GetRates structure.
type GetRatesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRatesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRatesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetRatesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetRatesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetRatesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetRatesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetRatesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetRatesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetRatesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetRatesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetRatesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRatesOK creates a GetRatesOK with default headers values
func NewGetRatesOK() *GetRatesOK {
	return &GetRatesOK{}
}

/*
GetRatesOK describes a response with status code 200, with default header values.

Success.
*/
type GetRatesOK struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *shipping_v2_models.GetRatesResponse
}

// IsSuccess returns true when this get rates o k response has a 2xx status code
func (o *GetRatesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get rates o k response has a 3xx status code
func (o *GetRatesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get rates o k response has a 4xx status code
func (o *GetRatesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get rates o k response has a 5xx status code
func (o *GetRatesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get rates o k response a status code equal to that given
func (o *GetRatesOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetRatesOK) Error() string {
	return fmt.Sprintf("[POST /shipping/v2/shipments/rates][%d] getRatesOK  %+v", 200, o.Payload)
}

func (o *GetRatesOK) String() string {
	return fmt.Sprintf("[POST /shipping/v2/shipments/rates][%d] getRatesOK  %+v", 200, o.Payload)
}

func (o *GetRatesOK) GetPayload() *shipping_v2_models.GetRatesResponse {
	return o.Payload
}

func (o *GetRatesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(shipping_v2_models.GetRatesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRatesBadRequest creates a GetRatesBadRequest with default headers values
func NewGetRatesBadRequest() *GetRatesBadRequest {
	return &GetRatesBadRequest{}
}

/*
GetRatesBadRequest describes a response with status code 400, with default header values.

Request has missing or invalid parameters and cannot be parsed.
*/
type GetRatesBadRequest struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *shipping_v2_models.ErrorList
}

// IsSuccess returns true when this get rates bad request response has a 2xx status code
func (o *GetRatesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get rates bad request response has a 3xx status code
func (o *GetRatesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get rates bad request response has a 4xx status code
func (o *GetRatesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get rates bad request response has a 5xx status code
func (o *GetRatesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get rates bad request response a status code equal to that given
func (o *GetRatesBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetRatesBadRequest) Error() string {
	return fmt.Sprintf("[POST /shipping/v2/shipments/rates][%d] getRatesBadRequest  %+v", 400, o.Payload)
}

func (o *GetRatesBadRequest) String() string {
	return fmt.Sprintf("[POST /shipping/v2/shipments/rates][%d] getRatesBadRequest  %+v", 400, o.Payload)
}

func (o *GetRatesBadRequest) GetPayload() *shipping_v2_models.ErrorList {
	return o.Payload
}

func (o *GetRatesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetRatesUnauthorized creates a GetRatesUnauthorized with default headers values
func NewGetRatesUnauthorized() *GetRatesUnauthorized {
	return &GetRatesUnauthorized{}
}

/*
GetRatesUnauthorized describes a response with status code 401, with default header values.

The request's Authorization header is not formatted correctly or does not contain a valid token.
*/
type GetRatesUnauthorized struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *shipping_v2_models.ErrorList
}

// IsSuccess returns true when this get rates unauthorized response has a 2xx status code
func (o *GetRatesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get rates unauthorized response has a 3xx status code
func (o *GetRatesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get rates unauthorized response has a 4xx status code
func (o *GetRatesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get rates unauthorized response has a 5xx status code
func (o *GetRatesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get rates unauthorized response a status code equal to that given
func (o *GetRatesUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetRatesUnauthorized) Error() string {
	return fmt.Sprintf("[POST /shipping/v2/shipments/rates][%d] getRatesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetRatesUnauthorized) String() string {
	return fmt.Sprintf("[POST /shipping/v2/shipments/rates][%d] getRatesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetRatesUnauthorized) GetPayload() *shipping_v2_models.ErrorList {
	return o.Payload
}

func (o *GetRatesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetRatesForbidden creates a GetRatesForbidden with default headers values
func NewGetRatesForbidden() *GetRatesForbidden {
	return &GetRatesForbidden{}
}

/*
GetRatesForbidden describes a response with status code 403, with default header values.

Indicates that access to the resource is forbidden. Possible reasons include Access Denied, Unauthorized, Expired Token, or Invalid Signature.
*/
type GetRatesForbidden struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *shipping_v2_models.ErrorList
}

// IsSuccess returns true when this get rates forbidden response has a 2xx status code
func (o *GetRatesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get rates forbidden response has a 3xx status code
func (o *GetRatesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get rates forbidden response has a 4xx status code
func (o *GetRatesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get rates forbidden response has a 5xx status code
func (o *GetRatesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get rates forbidden response a status code equal to that given
func (o *GetRatesForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetRatesForbidden) Error() string {
	return fmt.Sprintf("[POST /shipping/v2/shipments/rates][%d] getRatesForbidden  %+v", 403, o.Payload)
}

func (o *GetRatesForbidden) String() string {
	return fmt.Sprintf("[POST /shipping/v2/shipments/rates][%d] getRatesForbidden  %+v", 403, o.Payload)
}

func (o *GetRatesForbidden) GetPayload() *shipping_v2_models.ErrorList {
	return o.Payload
}

func (o *GetRatesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetRatesNotFound creates a GetRatesNotFound with default headers values
func NewGetRatesNotFound() *GetRatesNotFound {
	return &GetRatesNotFound{}
}

/*
GetRatesNotFound describes a response with status code 404, with default header values.

The resource specified does not exist.
*/
type GetRatesNotFound struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *shipping_v2_models.ErrorList
}

// IsSuccess returns true when this get rates not found response has a 2xx status code
func (o *GetRatesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get rates not found response has a 3xx status code
func (o *GetRatesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get rates not found response has a 4xx status code
func (o *GetRatesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get rates not found response has a 5xx status code
func (o *GetRatesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get rates not found response a status code equal to that given
func (o *GetRatesNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetRatesNotFound) Error() string {
	return fmt.Sprintf("[POST /shipping/v2/shipments/rates][%d] getRatesNotFound  %+v", 404, o.Payload)
}

func (o *GetRatesNotFound) String() string {
	return fmt.Sprintf("[POST /shipping/v2/shipments/rates][%d] getRatesNotFound  %+v", 404, o.Payload)
}

func (o *GetRatesNotFound) GetPayload() *shipping_v2_models.ErrorList {
	return o.Payload
}

func (o *GetRatesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetRatesRequestEntityTooLarge creates a GetRatesRequestEntityTooLarge with default headers values
func NewGetRatesRequestEntityTooLarge() *GetRatesRequestEntityTooLarge {
	return &GetRatesRequestEntityTooLarge{}
}

/*
GetRatesRequestEntityTooLarge describes a response with status code 413, with default header values.

The request size exceeded the maximum accepted size.
*/
type GetRatesRequestEntityTooLarge struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *shipping_v2_models.ErrorList
}

// IsSuccess returns true when this get rates request entity too large response has a 2xx status code
func (o *GetRatesRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get rates request entity too large response has a 3xx status code
func (o *GetRatesRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get rates request entity too large response has a 4xx status code
func (o *GetRatesRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get rates request entity too large response has a 5xx status code
func (o *GetRatesRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get rates request entity too large response a status code equal to that given
func (o *GetRatesRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetRatesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /shipping/v2/shipments/rates][%d] getRatesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetRatesRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[POST /shipping/v2/shipments/rates][%d] getRatesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetRatesRequestEntityTooLarge) GetPayload() *shipping_v2_models.ErrorList {
	return o.Payload
}

func (o *GetRatesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetRatesUnsupportedMediaType creates a GetRatesUnsupportedMediaType with default headers values
func NewGetRatesUnsupportedMediaType() *GetRatesUnsupportedMediaType {
	return &GetRatesUnsupportedMediaType{}
}

/*
GetRatesUnsupportedMediaType describes a response with status code 415, with default header values.

The request payload is in an unsupported format.
*/
type GetRatesUnsupportedMediaType struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *shipping_v2_models.ErrorList
}

// IsSuccess returns true when this get rates unsupported media type response has a 2xx status code
func (o *GetRatesUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get rates unsupported media type response has a 3xx status code
func (o *GetRatesUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get rates unsupported media type response has a 4xx status code
func (o *GetRatesUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get rates unsupported media type response has a 5xx status code
func (o *GetRatesUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get rates unsupported media type response a status code equal to that given
func (o *GetRatesUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetRatesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /shipping/v2/shipments/rates][%d] getRatesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetRatesUnsupportedMediaType) String() string {
	return fmt.Sprintf("[POST /shipping/v2/shipments/rates][%d] getRatesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetRatesUnsupportedMediaType) GetPayload() *shipping_v2_models.ErrorList {
	return o.Payload
}

func (o *GetRatesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetRatesTooManyRequests creates a GetRatesTooManyRequests with default headers values
func NewGetRatesTooManyRequests() *GetRatesTooManyRequests {
	return &GetRatesTooManyRequests{}
}

/*
GetRatesTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type GetRatesTooManyRequests struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *shipping_v2_models.ErrorList
}

// IsSuccess returns true when this get rates too many requests response has a 2xx status code
func (o *GetRatesTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get rates too many requests response has a 3xx status code
func (o *GetRatesTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get rates too many requests response has a 4xx status code
func (o *GetRatesTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get rates too many requests response has a 5xx status code
func (o *GetRatesTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get rates too many requests response a status code equal to that given
func (o *GetRatesTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetRatesTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /shipping/v2/shipments/rates][%d] getRatesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetRatesTooManyRequests) String() string {
	return fmt.Sprintf("[POST /shipping/v2/shipments/rates][%d] getRatesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetRatesTooManyRequests) GetPayload() *shipping_v2_models.ErrorList {
	return o.Payload
}

func (o *GetRatesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetRatesInternalServerError creates a GetRatesInternalServerError with default headers values
func NewGetRatesInternalServerError() *GetRatesInternalServerError {
	return &GetRatesInternalServerError{}
}

/*
GetRatesInternalServerError describes a response with status code 500, with default header values.

An unexpected condition occurred that prevented the server from fulfilling the request.
*/
type GetRatesInternalServerError struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *shipping_v2_models.ErrorList
}

// IsSuccess returns true when this get rates internal server error response has a 2xx status code
func (o *GetRatesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get rates internal server error response has a 3xx status code
func (o *GetRatesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get rates internal server error response has a 4xx status code
func (o *GetRatesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get rates internal server error response has a 5xx status code
func (o *GetRatesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get rates internal server error response a status code equal to that given
func (o *GetRatesInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetRatesInternalServerError) Error() string {
	return fmt.Sprintf("[POST /shipping/v2/shipments/rates][%d] getRatesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetRatesInternalServerError) String() string {
	return fmt.Sprintf("[POST /shipping/v2/shipments/rates][%d] getRatesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetRatesInternalServerError) GetPayload() *shipping_v2_models.ErrorList {
	return o.Payload
}

func (o *GetRatesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetRatesServiceUnavailable creates a GetRatesServiceUnavailable with default headers values
func NewGetRatesServiceUnavailable() *GetRatesServiceUnavailable {
	return &GetRatesServiceUnavailable{}
}

/*
GetRatesServiceUnavailable describes a response with status code 503, with default header values.

Temporary overloading or maintenance of the server.
*/
type GetRatesServiceUnavailable struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *shipping_v2_models.ErrorList
}

// IsSuccess returns true when this get rates service unavailable response has a 2xx status code
func (o *GetRatesServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get rates service unavailable response has a 3xx status code
func (o *GetRatesServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get rates service unavailable response has a 4xx status code
func (o *GetRatesServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get rates service unavailable response has a 5xx status code
func (o *GetRatesServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get rates service unavailable response a status code equal to that given
func (o *GetRatesServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetRatesServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /shipping/v2/shipments/rates][%d] getRatesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetRatesServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /shipping/v2/shipments/rates][%d] getRatesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetRatesServiceUnavailable) GetPayload() *shipping_v2_models.ErrorList {
	return o.Payload
}

func (o *GetRatesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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