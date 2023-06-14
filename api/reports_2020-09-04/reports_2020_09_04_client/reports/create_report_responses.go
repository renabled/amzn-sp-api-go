// Code generated by go-swagger; DO NOT EDIT.

package reports

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/renabled/amzn-sp-api-go/api/reports_2020-09-04/reports_2020_09_04_models"
)

// CreateReportReader is a Reader for the CreateReport structure.
type CreateReportReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateReportReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewCreateReportAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateReportBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateReportUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateReportForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateReportNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewCreateReportUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewCreateReportTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateReportInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewCreateReportServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateReportAccepted creates a CreateReportAccepted with default headers values
func NewCreateReportAccepted() *CreateReportAccepted {
	return &CreateReportAccepted{}
}

/*
CreateReportAccepted describes a response with status code 202, with default header values.

Success.
*/
type CreateReportAccepted struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *reports_2020_09_04_models.CreateReportResponse
}

// IsSuccess returns true when this create report accepted response has a 2xx status code
func (o *CreateReportAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create report accepted response has a 3xx status code
func (o *CreateReportAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create report accepted response has a 4xx status code
func (o *CreateReportAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this create report accepted response has a 5xx status code
func (o *CreateReportAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this create report accepted response a status code equal to that given
func (o *CreateReportAccepted) IsCode(code int) bool {
	return code == 202
}

func (o *CreateReportAccepted) Error() string {
	return fmt.Sprintf("[POST /reports/2020-09-04/reports][%d] createReportAccepted  %+v", 202, o.Payload)
}

func (o *CreateReportAccepted) String() string {
	return fmt.Sprintf("[POST /reports/2020-09-04/reports][%d] createReportAccepted  %+v", 202, o.Payload)
}

func (o *CreateReportAccepted) GetPayload() *reports_2020_09_04_models.CreateReportResponse {
	return o.Payload
}

func (o *CreateReportAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(reports_2020_09_04_models.CreateReportResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateReportBadRequest creates a CreateReportBadRequest with default headers values
func NewCreateReportBadRequest() *CreateReportBadRequest {
	return &CreateReportBadRequest{}
}

/*
CreateReportBadRequest describes a response with status code 400, with default header values.

Request has missing or invalid parameters and cannot be parsed.
*/
type CreateReportBadRequest struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *reports_2020_09_04_models.CreateReportResponse
}

// IsSuccess returns true when this create report bad request response has a 2xx status code
func (o *CreateReportBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create report bad request response has a 3xx status code
func (o *CreateReportBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create report bad request response has a 4xx status code
func (o *CreateReportBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create report bad request response has a 5xx status code
func (o *CreateReportBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create report bad request response a status code equal to that given
func (o *CreateReportBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *CreateReportBadRequest) Error() string {
	return fmt.Sprintf("[POST /reports/2020-09-04/reports][%d] createReportBadRequest  %+v", 400, o.Payload)
}

func (o *CreateReportBadRequest) String() string {
	return fmt.Sprintf("[POST /reports/2020-09-04/reports][%d] createReportBadRequest  %+v", 400, o.Payload)
}

func (o *CreateReportBadRequest) GetPayload() *reports_2020_09_04_models.CreateReportResponse {
	return o.Payload
}

func (o *CreateReportBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(reports_2020_09_04_models.CreateReportResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateReportUnauthorized creates a CreateReportUnauthorized with default headers values
func NewCreateReportUnauthorized() *CreateReportUnauthorized {
	return &CreateReportUnauthorized{}
}

/*
CreateReportUnauthorized describes a response with status code 401, with default header values.

The request's Authorization header is not formatted correctly or does not contain a valid token.
*/
type CreateReportUnauthorized struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *reports_2020_09_04_models.CreateReportResponse
}

// IsSuccess returns true when this create report unauthorized response has a 2xx status code
func (o *CreateReportUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create report unauthorized response has a 3xx status code
func (o *CreateReportUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create report unauthorized response has a 4xx status code
func (o *CreateReportUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this create report unauthorized response has a 5xx status code
func (o *CreateReportUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this create report unauthorized response a status code equal to that given
func (o *CreateReportUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *CreateReportUnauthorized) Error() string {
	return fmt.Sprintf("[POST /reports/2020-09-04/reports][%d] createReportUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateReportUnauthorized) String() string {
	return fmt.Sprintf("[POST /reports/2020-09-04/reports][%d] createReportUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateReportUnauthorized) GetPayload() *reports_2020_09_04_models.CreateReportResponse {
	return o.Payload
}

func (o *CreateReportUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(reports_2020_09_04_models.CreateReportResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateReportForbidden creates a CreateReportForbidden with default headers values
func NewCreateReportForbidden() *CreateReportForbidden {
	return &CreateReportForbidden{}
}

/*
CreateReportForbidden describes a response with status code 403, with default header values.

Indicates that access to the resource is forbidden. Possible reasons include Access Denied, Unauthorized, Expired Token, or Invalid Signature.
*/
type CreateReportForbidden struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *reports_2020_09_04_models.CreateReportResponse
}

// IsSuccess returns true when this create report forbidden response has a 2xx status code
func (o *CreateReportForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create report forbidden response has a 3xx status code
func (o *CreateReportForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create report forbidden response has a 4xx status code
func (o *CreateReportForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create report forbidden response has a 5xx status code
func (o *CreateReportForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create report forbidden response a status code equal to that given
func (o *CreateReportForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *CreateReportForbidden) Error() string {
	return fmt.Sprintf("[POST /reports/2020-09-04/reports][%d] createReportForbidden  %+v", 403, o.Payload)
}

func (o *CreateReportForbidden) String() string {
	return fmt.Sprintf("[POST /reports/2020-09-04/reports][%d] createReportForbidden  %+v", 403, o.Payload)
}

func (o *CreateReportForbidden) GetPayload() *reports_2020_09_04_models.CreateReportResponse {
	return o.Payload
}

func (o *CreateReportForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(reports_2020_09_04_models.CreateReportResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateReportNotFound creates a CreateReportNotFound with default headers values
func NewCreateReportNotFound() *CreateReportNotFound {
	return &CreateReportNotFound{}
}

/*
CreateReportNotFound describes a response with status code 404, with default header values.

The specified resource does not exist.
*/
type CreateReportNotFound struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *reports_2020_09_04_models.CreateReportResponse
}

// IsSuccess returns true when this create report not found response has a 2xx status code
func (o *CreateReportNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create report not found response has a 3xx status code
func (o *CreateReportNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create report not found response has a 4xx status code
func (o *CreateReportNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this create report not found response has a 5xx status code
func (o *CreateReportNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this create report not found response a status code equal to that given
func (o *CreateReportNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *CreateReportNotFound) Error() string {
	return fmt.Sprintf("[POST /reports/2020-09-04/reports][%d] createReportNotFound  %+v", 404, o.Payload)
}

func (o *CreateReportNotFound) String() string {
	return fmt.Sprintf("[POST /reports/2020-09-04/reports][%d] createReportNotFound  %+v", 404, o.Payload)
}

func (o *CreateReportNotFound) GetPayload() *reports_2020_09_04_models.CreateReportResponse {
	return o.Payload
}

func (o *CreateReportNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(reports_2020_09_04_models.CreateReportResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateReportUnsupportedMediaType creates a CreateReportUnsupportedMediaType with default headers values
func NewCreateReportUnsupportedMediaType() *CreateReportUnsupportedMediaType {
	return &CreateReportUnsupportedMediaType{}
}

/*
CreateReportUnsupportedMediaType describes a response with status code 415, with default header values.

The request's Content-Type header is invalid.
*/
type CreateReportUnsupportedMediaType struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *reports_2020_09_04_models.CreateReportResponse
}

// IsSuccess returns true when this create report unsupported media type response has a 2xx status code
func (o *CreateReportUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create report unsupported media type response has a 3xx status code
func (o *CreateReportUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create report unsupported media type response has a 4xx status code
func (o *CreateReportUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this create report unsupported media type response has a 5xx status code
func (o *CreateReportUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this create report unsupported media type response a status code equal to that given
func (o *CreateReportUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *CreateReportUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /reports/2020-09-04/reports][%d] createReportUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *CreateReportUnsupportedMediaType) String() string {
	return fmt.Sprintf("[POST /reports/2020-09-04/reports][%d] createReportUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *CreateReportUnsupportedMediaType) GetPayload() *reports_2020_09_04_models.CreateReportResponse {
	return o.Payload
}

func (o *CreateReportUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(reports_2020_09_04_models.CreateReportResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateReportTooManyRequests creates a CreateReportTooManyRequests with default headers values
func NewCreateReportTooManyRequests() *CreateReportTooManyRequests {
	return &CreateReportTooManyRequests{}
}

/*
CreateReportTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type CreateReportTooManyRequests struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *reports_2020_09_04_models.CreateReportResponse
}

// IsSuccess returns true when this create report too many requests response has a 2xx status code
func (o *CreateReportTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create report too many requests response has a 3xx status code
func (o *CreateReportTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create report too many requests response has a 4xx status code
func (o *CreateReportTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this create report too many requests response has a 5xx status code
func (o *CreateReportTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this create report too many requests response a status code equal to that given
func (o *CreateReportTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *CreateReportTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /reports/2020-09-04/reports][%d] createReportTooManyRequests  %+v", 429, o.Payload)
}

func (o *CreateReportTooManyRequests) String() string {
	return fmt.Sprintf("[POST /reports/2020-09-04/reports][%d] createReportTooManyRequests  %+v", 429, o.Payload)
}

func (o *CreateReportTooManyRequests) GetPayload() *reports_2020_09_04_models.CreateReportResponse {
	return o.Payload
}

func (o *CreateReportTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(reports_2020_09_04_models.CreateReportResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateReportInternalServerError creates a CreateReportInternalServerError with default headers values
func NewCreateReportInternalServerError() *CreateReportInternalServerError {
	return &CreateReportInternalServerError{}
}

/*
CreateReportInternalServerError describes a response with status code 500, with default header values.

An unexpected condition occurred that prevented the server from fulfilling the request.
*/
type CreateReportInternalServerError struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *reports_2020_09_04_models.CreateReportResponse
}

// IsSuccess returns true when this create report internal server error response has a 2xx status code
func (o *CreateReportInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create report internal server error response has a 3xx status code
func (o *CreateReportInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create report internal server error response has a 4xx status code
func (o *CreateReportInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create report internal server error response has a 5xx status code
func (o *CreateReportInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create report internal server error response a status code equal to that given
func (o *CreateReportInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *CreateReportInternalServerError) Error() string {
	return fmt.Sprintf("[POST /reports/2020-09-04/reports][%d] createReportInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateReportInternalServerError) String() string {
	return fmt.Sprintf("[POST /reports/2020-09-04/reports][%d] createReportInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateReportInternalServerError) GetPayload() *reports_2020_09_04_models.CreateReportResponse {
	return o.Payload
}

func (o *CreateReportInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(reports_2020_09_04_models.CreateReportResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateReportServiceUnavailable creates a CreateReportServiceUnavailable with default headers values
func NewCreateReportServiceUnavailable() *CreateReportServiceUnavailable {
	return &CreateReportServiceUnavailable{}
}

/*
CreateReportServiceUnavailable describes a response with status code 503, with default header values.

Temporary overloading or maintenance of the server.
*/
type CreateReportServiceUnavailable struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *reports_2020_09_04_models.CreateReportResponse
}

// IsSuccess returns true when this create report service unavailable response has a 2xx status code
func (o *CreateReportServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create report service unavailable response has a 3xx status code
func (o *CreateReportServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create report service unavailable response has a 4xx status code
func (o *CreateReportServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this create report service unavailable response has a 5xx status code
func (o *CreateReportServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this create report service unavailable response a status code equal to that given
func (o *CreateReportServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *CreateReportServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /reports/2020-09-04/reports][%d] createReportServiceUnavailable  %+v", 503, o.Payload)
}

func (o *CreateReportServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /reports/2020-09-04/reports][%d] createReportServiceUnavailable  %+v", 503, o.Payload)
}

func (o *CreateReportServiceUnavailable) GetPayload() *reports_2020_09_04_models.CreateReportResponse {
	return o.Payload
}

func (o *CreateReportServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(reports_2020_09_04_models.CreateReportResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
