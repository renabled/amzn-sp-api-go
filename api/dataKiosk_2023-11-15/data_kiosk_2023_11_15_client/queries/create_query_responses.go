// Code generated by go-swagger; DO NOT EDIT.

package queries

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/renabled/amzn-sp-api-go/api/dataKiosk_2023-11-15/data_kiosk_2023_11_15_models"
)

// CreateQueryReader is a Reader for the CreateQuery structure.
type CreateQueryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateQueryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewCreateQueryAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateQueryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateQueryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateQueryForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateQueryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewCreateQueryUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewCreateQueryTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateQueryInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewCreateQueryServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateQueryAccepted creates a CreateQueryAccepted with default headers values
func NewCreateQueryAccepted() *CreateQueryAccepted {
	return &CreateQueryAccepted{}
}

/*
CreateQueryAccepted describes a response with status code 202, with default header values.

Success.
*/
type CreateQueryAccepted struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *data_kiosk_2023_11_15_models.CreateQueryResponse
}

// IsSuccess returns true when this create query accepted response has a 2xx status code
func (o *CreateQueryAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create query accepted response has a 3xx status code
func (o *CreateQueryAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create query accepted response has a 4xx status code
func (o *CreateQueryAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this create query accepted response has a 5xx status code
func (o *CreateQueryAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this create query accepted response a status code equal to that given
func (o *CreateQueryAccepted) IsCode(code int) bool {
	return code == 202
}

func (o *CreateQueryAccepted) Error() string {
	return fmt.Sprintf("[POST /dataKiosk/2023-11-15/queries][%d] createQueryAccepted  %+v", 202, o.Payload)
}

func (o *CreateQueryAccepted) String() string {
	return fmt.Sprintf("[POST /dataKiosk/2023-11-15/queries][%d] createQueryAccepted  %+v", 202, o.Payload)
}

func (o *CreateQueryAccepted) GetPayload() *data_kiosk_2023_11_15_models.CreateQueryResponse {
	return o.Payload
}

func (o *CreateQueryAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(data_kiosk_2023_11_15_models.CreateQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateQueryBadRequest creates a CreateQueryBadRequest with default headers values
func NewCreateQueryBadRequest() *CreateQueryBadRequest {
	return &CreateQueryBadRequest{}
}

/*
CreateQueryBadRequest describes a response with status code 400, with default header values.

Request has missing or invalid parameters and cannot be parsed.
*/
type CreateQueryBadRequest struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *data_kiosk_2023_11_15_models.ErrorList
}

// IsSuccess returns true when this create query bad request response has a 2xx status code
func (o *CreateQueryBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create query bad request response has a 3xx status code
func (o *CreateQueryBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create query bad request response has a 4xx status code
func (o *CreateQueryBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create query bad request response has a 5xx status code
func (o *CreateQueryBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create query bad request response a status code equal to that given
func (o *CreateQueryBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *CreateQueryBadRequest) Error() string {
	return fmt.Sprintf("[POST /dataKiosk/2023-11-15/queries][%d] createQueryBadRequest  %+v", 400, o.Payload)
}

func (o *CreateQueryBadRequest) String() string {
	return fmt.Sprintf("[POST /dataKiosk/2023-11-15/queries][%d] createQueryBadRequest  %+v", 400, o.Payload)
}

func (o *CreateQueryBadRequest) GetPayload() *data_kiosk_2023_11_15_models.ErrorList {
	return o.Payload
}

func (o *CreateQueryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(data_kiosk_2023_11_15_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateQueryUnauthorized creates a CreateQueryUnauthorized with default headers values
func NewCreateQueryUnauthorized() *CreateQueryUnauthorized {
	return &CreateQueryUnauthorized{}
}

/*
CreateQueryUnauthorized describes a response with status code 401, with default header values.

The request's Authorization header is not formatted correctly or does not contain a valid token.
*/
type CreateQueryUnauthorized struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *data_kiosk_2023_11_15_models.ErrorList
}

// IsSuccess returns true when this create query unauthorized response has a 2xx status code
func (o *CreateQueryUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create query unauthorized response has a 3xx status code
func (o *CreateQueryUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create query unauthorized response has a 4xx status code
func (o *CreateQueryUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this create query unauthorized response has a 5xx status code
func (o *CreateQueryUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this create query unauthorized response a status code equal to that given
func (o *CreateQueryUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *CreateQueryUnauthorized) Error() string {
	return fmt.Sprintf("[POST /dataKiosk/2023-11-15/queries][%d] createQueryUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateQueryUnauthorized) String() string {
	return fmt.Sprintf("[POST /dataKiosk/2023-11-15/queries][%d] createQueryUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateQueryUnauthorized) GetPayload() *data_kiosk_2023_11_15_models.ErrorList {
	return o.Payload
}

func (o *CreateQueryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(data_kiosk_2023_11_15_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateQueryForbidden creates a CreateQueryForbidden with default headers values
func NewCreateQueryForbidden() *CreateQueryForbidden {
	return &CreateQueryForbidden{}
}

/*
CreateQueryForbidden describes a response with status code 403, with default header values.

Indicates that access to the resource is forbidden. Possible reasons include Access Denied, Unauthorized, Expired Token, or Invalid Signature.
*/
type CreateQueryForbidden struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *data_kiosk_2023_11_15_models.ErrorList
}

// IsSuccess returns true when this create query forbidden response has a 2xx status code
func (o *CreateQueryForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create query forbidden response has a 3xx status code
func (o *CreateQueryForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create query forbidden response has a 4xx status code
func (o *CreateQueryForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create query forbidden response has a 5xx status code
func (o *CreateQueryForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create query forbidden response a status code equal to that given
func (o *CreateQueryForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *CreateQueryForbidden) Error() string {
	return fmt.Sprintf("[POST /dataKiosk/2023-11-15/queries][%d] createQueryForbidden  %+v", 403, o.Payload)
}

func (o *CreateQueryForbidden) String() string {
	return fmt.Sprintf("[POST /dataKiosk/2023-11-15/queries][%d] createQueryForbidden  %+v", 403, o.Payload)
}

func (o *CreateQueryForbidden) GetPayload() *data_kiosk_2023_11_15_models.ErrorList {
	return o.Payload
}

func (o *CreateQueryForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(data_kiosk_2023_11_15_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateQueryNotFound creates a CreateQueryNotFound with default headers values
func NewCreateQueryNotFound() *CreateQueryNotFound {
	return &CreateQueryNotFound{}
}

/*
CreateQueryNotFound describes a response with status code 404, with default header values.

The specified resource does not exist.
*/
type CreateQueryNotFound struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *data_kiosk_2023_11_15_models.ErrorList
}

// IsSuccess returns true when this create query not found response has a 2xx status code
func (o *CreateQueryNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create query not found response has a 3xx status code
func (o *CreateQueryNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create query not found response has a 4xx status code
func (o *CreateQueryNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this create query not found response has a 5xx status code
func (o *CreateQueryNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this create query not found response a status code equal to that given
func (o *CreateQueryNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *CreateQueryNotFound) Error() string {
	return fmt.Sprintf("[POST /dataKiosk/2023-11-15/queries][%d] createQueryNotFound  %+v", 404, o.Payload)
}

func (o *CreateQueryNotFound) String() string {
	return fmt.Sprintf("[POST /dataKiosk/2023-11-15/queries][%d] createQueryNotFound  %+v", 404, o.Payload)
}

func (o *CreateQueryNotFound) GetPayload() *data_kiosk_2023_11_15_models.ErrorList {
	return o.Payload
}

func (o *CreateQueryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(data_kiosk_2023_11_15_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateQueryUnsupportedMediaType creates a CreateQueryUnsupportedMediaType with default headers values
func NewCreateQueryUnsupportedMediaType() *CreateQueryUnsupportedMediaType {
	return &CreateQueryUnsupportedMediaType{}
}

/*
CreateQueryUnsupportedMediaType describes a response with status code 415, with default header values.

The request's Content-Type header is invalid.
*/
type CreateQueryUnsupportedMediaType struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *data_kiosk_2023_11_15_models.ErrorList
}

// IsSuccess returns true when this create query unsupported media type response has a 2xx status code
func (o *CreateQueryUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create query unsupported media type response has a 3xx status code
func (o *CreateQueryUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create query unsupported media type response has a 4xx status code
func (o *CreateQueryUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this create query unsupported media type response has a 5xx status code
func (o *CreateQueryUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this create query unsupported media type response a status code equal to that given
func (o *CreateQueryUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *CreateQueryUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /dataKiosk/2023-11-15/queries][%d] createQueryUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *CreateQueryUnsupportedMediaType) String() string {
	return fmt.Sprintf("[POST /dataKiosk/2023-11-15/queries][%d] createQueryUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *CreateQueryUnsupportedMediaType) GetPayload() *data_kiosk_2023_11_15_models.ErrorList {
	return o.Payload
}

func (o *CreateQueryUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(data_kiosk_2023_11_15_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateQueryTooManyRequests creates a CreateQueryTooManyRequests with default headers values
func NewCreateQueryTooManyRequests() *CreateQueryTooManyRequests {
	return &CreateQueryTooManyRequests{}
}

/*
CreateQueryTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type CreateQueryTooManyRequests struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *data_kiosk_2023_11_15_models.ErrorList
}

// IsSuccess returns true when this create query too many requests response has a 2xx status code
func (o *CreateQueryTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create query too many requests response has a 3xx status code
func (o *CreateQueryTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create query too many requests response has a 4xx status code
func (o *CreateQueryTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this create query too many requests response has a 5xx status code
func (o *CreateQueryTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this create query too many requests response a status code equal to that given
func (o *CreateQueryTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *CreateQueryTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /dataKiosk/2023-11-15/queries][%d] createQueryTooManyRequests  %+v", 429, o.Payload)
}

func (o *CreateQueryTooManyRequests) String() string {
	return fmt.Sprintf("[POST /dataKiosk/2023-11-15/queries][%d] createQueryTooManyRequests  %+v", 429, o.Payload)
}

func (o *CreateQueryTooManyRequests) GetPayload() *data_kiosk_2023_11_15_models.ErrorList {
	return o.Payload
}

func (o *CreateQueryTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(data_kiosk_2023_11_15_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateQueryInternalServerError creates a CreateQueryInternalServerError with default headers values
func NewCreateQueryInternalServerError() *CreateQueryInternalServerError {
	return &CreateQueryInternalServerError{}
}

/*
CreateQueryInternalServerError describes a response with status code 500, with default header values.

An unexpected condition occurred that prevented the server from fulfilling the request.
*/
type CreateQueryInternalServerError struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *data_kiosk_2023_11_15_models.ErrorList
}

// IsSuccess returns true when this create query internal server error response has a 2xx status code
func (o *CreateQueryInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create query internal server error response has a 3xx status code
func (o *CreateQueryInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create query internal server error response has a 4xx status code
func (o *CreateQueryInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create query internal server error response has a 5xx status code
func (o *CreateQueryInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create query internal server error response a status code equal to that given
func (o *CreateQueryInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *CreateQueryInternalServerError) Error() string {
	return fmt.Sprintf("[POST /dataKiosk/2023-11-15/queries][%d] createQueryInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateQueryInternalServerError) String() string {
	return fmt.Sprintf("[POST /dataKiosk/2023-11-15/queries][%d] createQueryInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateQueryInternalServerError) GetPayload() *data_kiosk_2023_11_15_models.ErrorList {
	return o.Payload
}

func (o *CreateQueryInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(data_kiosk_2023_11_15_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateQueryServiceUnavailable creates a CreateQueryServiceUnavailable with default headers values
func NewCreateQueryServiceUnavailable() *CreateQueryServiceUnavailable {
	return &CreateQueryServiceUnavailable{}
}

/*
CreateQueryServiceUnavailable describes a response with status code 503, with default header values.

Temporary overloading or maintenance of the server.
*/
type CreateQueryServiceUnavailable struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *data_kiosk_2023_11_15_models.ErrorList
}

// IsSuccess returns true when this create query service unavailable response has a 2xx status code
func (o *CreateQueryServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create query service unavailable response has a 3xx status code
func (o *CreateQueryServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create query service unavailable response has a 4xx status code
func (o *CreateQueryServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this create query service unavailable response has a 5xx status code
func (o *CreateQueryServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this create query service unavailable response a status code equal to that given
func (o *CreateQueryServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *CreateQueryServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /dataKiosk/2023-11-15/queries][%d] createQueryServiceUnavailable  %+v", 503, o.Payload)
}

func (o *CreateQueryServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /dataKiosk/2023-11-15/queries][%d] createQueryServiceUnavailable  %+v", 503, o.Payload)
}

func (o *CreateQueryServiceUnavailable) GetPayload() *data_kiosk_2023_11_15_models.ErrorList {
	return o.Payload
}

func (o *CreateQueryServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(data_kiosk_2023_11_15_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
