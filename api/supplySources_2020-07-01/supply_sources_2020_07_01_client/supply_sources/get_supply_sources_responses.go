// Code generated by go-swagger; DO NOT EDIT.

package supply_sources

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/renabled/amzn-sp-api-go/api/supplySources_2020-07-01/supply_sources_2020_07_01_models"
)

// GetSupplySourcesReader is a Reader for the GetSupplySources structure.
type GetSupplySourcesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSupplySourcesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSupplySourcesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetSupplySourcesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetSupplySourcesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetSupplySourcesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetSupplySourcesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetSupplySourcesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetSupplySourcesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetSupplySourcesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetSupplySourcesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetSupplySourcesOK creates a GetSupplySourcesOK with default headers values
func NewGetSupplySourcesOK() *GetSupplySourcesOK {
	return &GetSupplySourcesOK{}
}

/*
GetSupplySourcesOK describes a response with status code 200, with default header values.

Success.
*/
type GetSupplySourcesOK struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* The unique request reference ID.
	 */
	XAmznRequestID string

	Payload *supply_sources_2020_07_01_models.GetSupplySourcesResponse
}

// IsSuccess returns true when this get supply sources o k response has a 2xx status code
func (o *GetSupplySourcesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get supply sources o k response has a 3xx status code
func (o *GetSupplySourcesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get supply sources o k response has a 4xx status code
func (o *GetSupplySourcesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get supply sources o k response has a 5xx status code
func (o *GetSupplySourcesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get supply sources o k response a status code equal to that given
func (o *GetSupplySourcesOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetSupplySourcesOK) Error() string {
	return fmt.Sprintf("[GET /supplySources/2020-07-01/supplySources][%d] getSupplySourcesOK  %+v", 200, o.Payload)
}

func (o *GetSupplySourcesOK) String() string {
	return fmt.Sprintf("[GET /supplySources/2020-07-01/supplySources][%d] getSupplySourcesOK  %+v", 200, o.Payload)
}

func (o *GetSupplySourcesOK) GetPayload() *supply_sources_2020_07_01_models.GetSupplySourcesResponse {
	return o.Payload
}

func (o *GetSupplySourcesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(supply_sources_2020_07_01_models.GetSupplySourcesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSupplySourcesBadRequest creates a GetSupplySourcesBadRequest with default headers values
func NewGetSupplySourcesBadRequest() *GetSupplySourcesBadRequest {
	return &GetSupplySourcesBadRequest{}
}

/*
GetSupplySourcesBadRequest describes a response with status code 400, with default header values.

The request has missing or invalid parameters and cannot be parsed.
*/
type GetSupplySourcesBadRequest struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* The unique request reference ID.
	 */
	XAmznRequestID string

	Payload *supply_sources_2020_07_01_models.ErrorList
}

// IsSuccess returns true when this get supply sources bad request response has a 2xx status code
func (o *GetSupplySourcesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get supply sources bad request response has a 3xx status code
func (o *GetSupplySourcesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get supply sources bad request response has a 4xx status code
func (o *GetSupplySourcesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get supply sources bad request response has a 5xx status code
func (o *GetSupplySourcesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get supply sources bad request response a status code equal to that given
func (o *GetSupplySourcesBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetSupplySourcesBadRequest) Error() string {
	return fmt.Sprintf("[GET /supplySources/2020-07-01/supplySources][%d] getSupplySourcesBadRequest  %+v", 400, o.Payload)
}

func (o *GetSupplySourcesBadRequest) String() string {
	return fmt.Sprintf("[GET /supplySources/2020-07-01/supplySources][%d] getSupplySourcesBadRequest  %+v", 400, o.Payload)
}

func (o *GetSupplySourcesBadRequest) GetPayload() *supply_sources_2020_07_01_models.ErrorList {
	return o.Payload
}

func (o *GetSupplySourcesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(supply_sources_2020_07_01_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSupplySourcesForbidden creates a GetSupplySourcesForbidden with default headers values
func NewGetSupplySourcesForbidden() *GetSupplySourcesForbidden {
	return &GetSupplySourcesForbidden{}
}

/*
GetSupplySourcesForbidden describes a response with status code 403, with default header values.

An error that indicates that access to the resource is forbidden. Possible reasons include Access Denied, Unauthorized, Expired Token, or Invalid Signature.
*/
type GetSupplySourcesForbidden struct {

	/* The unique request reference ID.
	 */
	XAmznRequestID string

	Payload *supply_sources_2020_07_01_models.ErrorList
}

// IsSuccess returns true when this get supply sources forbidden response has a 2xx status code
func (o *GetSupplySourcesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get supply sources forbidden response has a 3xx status code
func (o *GetSupplySourcesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get supply sources forbidden response has a 4xx status code
func (o *GetSupplySourcesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get supply sources forbidden response has a 5xx status code
func (o *GetSupplySourcesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get supply sources forbidden response a status code equal to that given
func (o *GetSupplySourcesForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetSupplySourcesForbidden) Error() string {
	return fmt.Sprintf("[GET /supplySources/2020-07-01/supplySources][%d] getSupplySourcesForbidden  %+v", 403, o.Payload)
}

func (o *GetSupplySourcesForbidden) String() string {
	return fmt.Sprintf("[GET /supplySources/2020-07-01/supplySources][%d] getSupplySourcesForbidden  %+v", 403, o.Payload)
}

func (o *GetSupplySourcesForbidden) GetPayload() *supply_sources_2020_07_01_models.ErrorList {
	return o.Payload
}

func (o *GetSupplySourcesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(supply_sources_2020_07_01_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSupplySourcesNotFound creates a GetSupplySourcesNotFound with default headers values
func NewGetSupplySourcesNotFound() *GetSupplySourcesNotFound {
	return &GetSupplySourcesNotFound{}
}

/*
GetSupplySourcesNotFound describes a response with status code 404, with default header values.

The resource specified does not exist.
*/
type GetSupplySourcesNotFound struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* The unique request reference ID.
	 */
	XAmznRequestID string

	Payload *supply_sources_2020_07_01_models.ErrorList
}

// IsSuccess returns true when this get supply sources not found response has a 2xx status code
func (o *GetSupplySourcesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get supply sources not found response has a 3xx status code
func (o *GetSupplySourcesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get supply sources not found response has a 4xx status code
func (o *GetSupplySourcesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get supply sources not found response has a 5xx status code
func (o *GetSupplySourcesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get supply sources not found response a status code equal to that given
func (o *GetSupplySourcesNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetSupplySourcesNotFound) Error() string {
	return fmt.Sprintf("[GET /supplySources/2020-07-01/supplySources][%d] getSupplySourcesNotFound  %+v", 404, o.Payload)
}

func (o *GetSupplySourcesNotFound) String() string {
	return fmt.Sprintf("[GET /supplySources/2020-07-01/supplySources][%d] getSupplySourcesNotFound  %+v", 404, o.Payload)
}

func (o *GetSupplySourcesNotFound) GetPayload() *supply_sources_2020_07_01_models.ErrorList {
	return o.Payload
}

func (o *GetSupplySourcesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(supply_sources_2020_07_01_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSupplySourcesRequestEntityTooLarge creates a GetSupplySourcesRequestEntityTooLarge with default headers values
func NewGetSupplySourcesRequestEntityTooLarge() *GetSupplySourcesRequestEntityTooLarge {
	return &GetSupplySourcesRequestEntityTooLarge{}
}

/*
GetSupplySourcesRequestEntityTooLarge describes a response with status code 413, with default header values.

The request size exceeded the maximum accepted size.
*/
type GetSupplySourcesRequestEntityTooLarge struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* The unique request reference ID.
	 */
	XAmznRequestID string

	Payload *supply_sources_2020_07_01_models.ErrorList
}

// IsSuccess returns true when this get supply sources request entity too large response has a 2xx status code
func (o *GetSupplySourcesRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get supply sources request entity too large response has a 3xx status code
func (o *GetSupplySourcesRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get supply sources request entity too large response has a 4xx status code
func (o *GetSupplySourcesRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get supply sources request entity too large response has a 5xx status code
func (o *GetSupplySourcesRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get supply sources request entity too large response a status code equal to that given
func (o *GetSupplySourcesRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetSupplySourcesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /supplySources/2020-07-01/supplySources][%d] getSupplySourcesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetSupplySourcesRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /supplySources/2020-07-01/supplySources][%d] getSupplySourcesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetSupplySourcesRequestEntityTooLarge) GetPayload() *supply_sources_2020_07_01_models.ErrorList {
	return o.Payload
}

func (o *GetSupplySourcesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(supply_sources_2020_07_01_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSupplySourcesUnsupportedMediaType creates a GetSupplySourcesUnsupportedMediaType with default headers values
func NewGetSupplySourcesUnsupportedMediaType() *GetSupplySourcesUnsupportedMediaType {
	return &GetSupplySourcesUnsupportedMediaType{}
}

/*
GetSupplySourcesUnsupportedMediaType describes a response with status code 415, with default header values.

The request payload is in an unsupported format.
*/
type GetSupplySourcesUnsupportedMediaType struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* The unique request reference ID.
	 */
	XAmznRequestID string

	Payload *supply_sources_2020_07_01_models.ErrorList
}

// IsSuccess returns true when this get supply sources unsupported media type response has a 2xx status code
func (o *GetSupplySourcesUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get supply sources unsupported media type response has a 3xx status code
func (o *GetSupplySourcesUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get supply sources unsupported media type response has a 4xx status code
func (o *GetSupplySourcesUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get supply sources unsupported media type response has a 5xx status code
func (o *GetSupplySourcesUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get supply sources unsupported media type response a status code equal to that given
func (o *GetSupplySourcesUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetSupplySourcesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /supplySources/2020-07-01/supplySources][%d] getSupplySourcesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetSupplySourcesUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /supplySources/2020-07-01/supplySources][%d] getSupplySourcesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetSupplySourcesUnsupportedMediaType) GetPayload() *supply_sources_2020_07_01_models.ErrorList {
	return o.Payload
}

func (o *GetSupplySourcesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(supply_sources_2020_07_01_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSupplySourcesTooManyRequests creates a GetSupplySourcesTooManyRequests with default headers values
func NewGetSupplySourcesTooManyRequests() *GetSupplySourcesTooManyRequests {
	return &GetSupplySourcesTooManyRequests{}
}

/*
GetSupplySourcesTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type GetSupplySourcesTooManyRequests struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* The unique request reference ID.
	 */
	XAmznRequestID string

	Payload *supply_sources_2020_07_01_models.ErrorList
}

// IsSuccess returns true when this get supply sources too many requests response has a 2xx status code
func (o *GetSupplySourcesTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get supply sources too many requests response has a 3xx status code
func (o *GetSupplySourcesTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get supply sources too many requests response has a 4xx status code
func (o *GetSupplySourcesTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get supply sources too many requests response has a 5xx status code
func (o *GetSupplySourcesTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get supply sources too many requests response a status code equal to that given
func (o *GetSupplySourcesTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetSupplySourcesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /supplySources/2020-07-01/supplySources][%d] getSupplySourcesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetSupplySourcesTooManyRequests) String() string {
	return fmt.Sprintf("[GET /supplySources/2020-07-01/supplySources][%d] getSupplySourcesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetSupplySourcesTooManyRequests) GetPayload() *supply_sources_2020_07_01_models.ErrorList {
	return o.Payload
}

func (o *GetSupplySourcesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(supply_sources_2020_07_01_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSupplySourcesInternalServerError creates a GetSupplySourcesInternalServerError with default headers values
func NewGetSupplySourcesInternalServerError() *GetSupplySourcesInternalServerError {
	return &GetSupplySourcesInternalServerError{}
}

/*
GetSupplySourcesInternalServerError describes a response with status code 500, with default header values.

An unexpected condition occurred that prevented the server from fulfilling the request.
*/
type GetSupplySourcesInternalServerError struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* The unique request reference ID.
	 */
	XAmznRequestID string

	Payload *supply_sources_2020_07_01_models.ErrorList
}

// IsSuccess returns true when this get supply sources internal server error response has a 2xx status code
func (o *GetSupplySourcesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get supply sources internal server error response has a 3xx status code
func (o *GetSupplySourcesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get supply sources internal server error response has a 4xx status code
func (o *GetSupplySourcesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get supply sources internal server error response has a 5xx status code
func (o *GetSupplySourcesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get supply sources internal server error response a status code equal to that given
func (o *GetSupplySourcesInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetSupplySourcesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /supplySources/2020-07-01/supplySources][%d] getSupplySourcesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetSupplySourcesInternalServerError) String() string {
	return fmt.Sprintf("[GET /supplySources/2020-07-01/supplySources][%d] getSupplySourcesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetSupplySourcesInternalServerError) GetPayload() *supply_sources_2020_07_01_models.ErrorList {
	return o.Payload
}

func (o *GetSupplySourcesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(supply_sources_2020_07_01_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSupplySourcesServiceUnavailable creates a GetSupplySourcesServiceUnavailable with default headers values
func NewGetSupplySourcesServiceUnavailable() *GetSupplySourcesServiceUnavailable {
	return &GetSupplySourcesServiceUnavailable{}
}

/*
GetSupplySourcesServiceUnavailable describes a response with status code 503, with default header values.

The temporary overloading or maintenance of the server.
*/
type GetSupplySourcesServiceUnavailable struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* The unique request reference ID.
	 */
	XAmznRequestID string

	Payload *supply_sources_2020_07_01_models.ErrorList
}

// IsSuccess returns true when this get supply sources service unavailable response has a 2xx status code
func (o *GetSupplySourcesServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get supply sources service unavailable response has a 3xx status code
func (o *GetSupplySourcesServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get supply sources service unavailable response has a 4xx status code
func (o *GetSupplySourcesServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get supply sources service unavailable response has a 5xx status code
func (o *GetSupplySourcesServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get supply sources service unavailable response a status code equal to that given
func (o *GetSupplySourcesServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetSupplySourcesServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /supplySources/2020-07-01/supplySources][%d] getSupplySourcesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetSupplySourcesServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /supplySources/2020-07-01/supplySources][%d] getSupplySourcesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetSupplySourcesServiceUnavailable) GetPayload() *supply_sources_2020_07_01_models.ErrorList {
	return o.Payload
}

func (o *GetSupplySourcesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(supply_sources_2020_07_01_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
