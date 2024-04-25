// Code generated by go-swagger; DO NOT EDIT.

package reports

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/renabled/amzn-sp-api-go/api/reports_2021-06-30/reports_2021_06_30_models"
)

// CreateReportScheduleReader is a Reader for the CreateReportSchedule structure.
type CreateReportScheduleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateReportScheduleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateReportScheduleCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateReportScheduleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateReportScheduleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateReportScheduleForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateReportScheduleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewCreateReportScheduleUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewCreateReportScheduleTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateReportScheduleInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewCreateReportScheduleServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateReportScheduleCreated creates a CreateReportScheduleCreated with default headers values
func NewCreateReportScheduleCreated() *CreateReportScheduleCreated {
	return &CreateReportScheduleCreated{}
}

/*
CreateReportScheduleCreated describes a response with status code 201, with default header values.

Success.
*/
type CreateReportScheduleCreated struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *reports_2021_06_30_models.CreateReportScheduleResponse
}

// IsSuccess returns true when this create report schedule created response has a 2xx status code
func (o *CreateReportScheduleCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create report schedule created response has a 3xx status code
func (o *CreateReportScheduleCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create report schedule created response has a 4xx status code
func (o *CreateReportScheduleCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create report schedule created response has a 5xx status code
func (o *CreateReportScheduleCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create report schedule created response a status code equal to that given
func (o *CreateReportScheduleCreated) IsCode(code int) bool {
	return code == 201
}

func (o *CreateReportScheduleCreated) Error() string {
	return fmt.Sprintf("[POST /reports/2021-06-30/schedules][%d] createReportScheduleCreated  %+v", 201, o.Payload)
}

func (o *CreateReportScheduleCreated) String() string {
	return fmt.Sprintf("[POST /reports/2021-06-30/schedules][%d] createReportScheduleCreated  %+v", 201, o.Payload)
}

func (o *CreateReportScheduleCreated) GetPayload() *reports_2021_06_30_models.CreateReportScheduleResponse {
	return o.Payload
}

func (o *CreateReportScheduleCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(reports_2021_06_30_models.CreateReportScheduleResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateReportScheduleBadRequest creates a CreateReportScheduleBadRequest with default headers values
func NewCreateReportScheduleBadRequest() *CreateReportScheduleBadRequest {
	return &CreateReportScheduleBadRequest{}
}

/*
CreateReportScheduleBadRequest describes a response with status code 400, with default header values.

Request has missing or invalid parameters and cannot be parsed.
*/
type CreateReportScheduleBadRequest struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *reports_2021_06_30_models.ErrorList
}

// IsSuccess returns true when this create report schedule bad request response has a 2xx status code
func (o *CreateReportScheduleBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create report schedule bad request response has a 3xx status code
func (o *CreateReportScheduleBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create report schedule bad request response has a 4xx status code
func (o *CreateReportScheduleBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create report schedule bad request response has a 5xx status code
func (o *CreateReportScheduleBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create report schedule bad request response a status code equal to that given
func (o *CreateReportScheduleBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *CreateReportScheduleBadRequest) Error() string {
	return fmt.Sprintf("[POST /reports/2021-06-30/schedules][%d] createReportScheduleBadRequest  %+v", 400, o.Payload)
}

func (o *CreateReportScheduleBadRequest) String() string {
	return fmt.Sprintf("[POST /reports/2021-06-30/schedules][%d] createReportScheduleBadRequest  %+v", 400, o.Payload)
}

func (o *CreateReportScheduleBadRequest) GetPayload() *reports_2021_06_30_models.ErrorList {
	return o.Payload
}

func (o *CreateReportScheduleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(reports_2021_06_30_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateReportScheduleUnauthorized creates a CreateReportScheduleUnauthorized with default headers values
func NewCreateReportScheduleUnauthorized() *CreateReportScheduleUnauthorized {
	return &CreateReportScheduleUnauthorized{}
}

/*
CreateReportScheduleUnauthorized describes a response with status code 401, with default header values.

The request's Authorization header is not formatted correctly or does not contain a valid token.
*/
type CreateReportScheduleUnauthorized struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *reports_2021_06_30_models.ErrorList
}

// IsSuccess returns true when this create report schedule unauthorized response has a 2xx status code
func (o *CreateReportScheduleUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create report schedule unauthorized response has a 3xx status code
func (o *CreateReportScheduleUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create report schedule unauthorized response has a 4xx status code
func (o *CreateReportScheduleUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this create report schedule unauthorized response has a 5xx status code
func (o *CreateReportScheduleUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this create report schedule unauthorized response a status code equal to that given
func (o *CreateReportScheduleUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *CreateReportScheduleUnauthorized) Error() string {
	return fmt.Sprintf("[POST /reports/2021-06-30/schedules][%d] createReportScheduleUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateReportScheduleUnauthorized) String() string {
	return fmt.Sprintf("[POST /reports/2021-06-30/schedules][%d] createReportScheduleUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateReportScheduleUnauthorized) GetPayload() *reports_2021_06_30_models.ErrorList {
	return o.Payload
}

func (o *CreateReportScheduleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(reports_2021_06_30_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateReportScheduleForbidden creates a CreateReportScheduleForbidden with default headers values
func NewCreateReportScheduleForbidden() *CreateReportScheduleForbidden {
	return &CreateReportScheduleForbidden{}
}

/*
CreateReportScheduleForbidden describes a response with status code 403, with default header values.

Indicates that access to the resource is forbidden. Possible reasons include Access Denied, Unauthorized, Expired Token, or Invalid Signature.
*/
type CreateReportScheduleForbidden struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *reports_2021_06_30_models.ErrorList
}

// IsSuccess returns true when this create report schedule forbidden response has a 2xx status code
func (o *CreateReportScheduleForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create report schedule forbidden response has a 3xx status code
func (o *CreateReportScheduleForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create report schedule forbidden response has a 4xx status code
func (o *CreateReportScheduleForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create report schedule forbidden response has a 5xx status code
func (o *CreateReportScheduleForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create report schedule forbidden response a status code equal to that given
func (o *CreateReportScheduleForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *CreateReportScheduleForbidden) Error() string {
	return fmt.Sprintf("[POST /reports/2021-06-30/schedules][%d] createReportScheduleForbidden  %+v", 403, o.Payload)
}

func (o *CreateReportScheduleForbidden) String() string {
	return fmt.Sprintf("[POST /reports/2021-06-30/schedules][%d] createReportScheduleForbidden  %+v", 403, o.Payload)
}

func (o *CreateReportScheduleForbidden) GetPayload() *reports_2021_06_30_models.ErrorList {
	return o.Payload
}

func (o *CreateReportScheduleForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(reports_2021_06_30_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateReportScheduleNotFound creates a CreateReportScheduleNotFound with default headers values
func NewCreateReportScheduleNotFound() *CreateReportScheduleNotFound {
	return &CreateReportScheduleNotFound{}
}

/*
CreateReportScheduleNotFound describes a response with status code 404, with default header values.

The specified resource does not exist.
*/
type CreateReportScheduleNotFound struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *reports_2021_06_30_models.ErrorList
}

// IsSuccess returns true when this create report schedule not found response has a 2xx status code
func (o *CreateReportScheduleNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create report schedule not found response has a 3xx status code
func (o *CreateReportScheduleNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create report schedule not found response has a 4xx status code
func (o *CreateReportScheduleNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this create report schedule not found response has a 5xx status code
func (o *CreateReportScheduleNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this create report schedule not found response a status code equal to that given
func (o *CreateReportScheduleNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *CreateReportScheduleNotFound) Error() string {
	return fmt.Sprintf("[POST /reports/2021-06-30/schedules][%d] createReportScheduleNotFound  %+v", 404, o.Payload)
}

func (o *CreateReportScheduleNotFound) String() string {
	return fmt.Sprintf("[POST /reports/2021-06-30/schedules][%d] createReportScheduleNotFound  %+v", 404, o.Payload)
}

func (o *CreateReportScheduleNotFound) GetPayload() *reports_2021_06_30_models.ErrorList {
	return o.Payload
}

func (o *CreateReportScheduleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(reports_2021_06_30_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateReportScheduleUnsupportedMediaType creates a CreateReportScheduleUnsupportedMediaType with default headers values
func NewCreateReportScheduleUnsupportedMediaType() *CreateReportScheduleUnsupportedMediaType {
	return &CreateReportScheduleUnsupportedMediaType{}
}

/*
CreateReportScheduleUnsupportedMediaType describes a response with status code 415, with default header values.

The request's Content-Type header is invalid.
*/
type CreateReportScheduleUnsupportedMediaType struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *reports_2021_06_30_models.ErrorList
}

// IsSuccess returns true when this create report schedule unsupported media type response has a 2xx status code
func (o *CreateReportScheduleUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create report schedule unsupported media type response has a 3xx status code
func (o *CreateReportScheduleUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create report schedule unsupported media type response has a 4xx status code
func (o *CreateReportScheduleUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this create report schedule unsupported media type response has a 5xx status code
func (o *CreateReportScheduleUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this create report schedule unsupported media type response a status code equal to that given
func (o *CreateReportScheduleUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *CreateReportScheduleUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /reports/2021-06-30/schedules][%d] createReportScheduleUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *CreateReportScheduleUnsupportedMediaType) String() string {
	return fmt.Sprintf("[POST /reports/2021-06-30/schedules][%d] createReportScheduleUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *CreateReportScheduleUnsupportedMediaType) GetPayload() *reports_2021_06_30_models.ErrorList {
	return o.Payload
}

func (o *CreateReportScheduleUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(reports_2021_06_30_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateReportScheduleTooManyRequests creates a CreateReportScheduleTooManyRequests with default headers values
func NewCreateReportScheduleTooManyRequests() *CreateReportScheduleTooManyRequests {
	return &CreateReportScheduleTooManyRequests{}
}

/*
CreateReportScheduleTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type CreateReportScheduleTooManyRequests struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *reports_2021_06_30_models.ErrorList
}

// IsSuccess returns true when this create report schedule too many requests response has a 2xx status code
func (o *CreateReportScheduleTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create report schedule too many requests response has a 3xx status code
func (o *CreateReportScheduleTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create report schedule too many requests response has a 4xx status code
func (o *CreateReportScheduleTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this create report schedule too many requests response has a 5xx status code
func (o *CreateReportScheduleTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this create report schedule too many requests response a status code equal to that given
func (o *CreateReportScheduleTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *CreateReportScheduleTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /reports/2021-06-30/schedules][%d] createReportScheduleTooManyRequests  %+v", 429, o.Payload)
}

func (o *CreateReportScheduleTooManyRequests) String() string {
	return fmt.Sprintf("[POST /reports/2021-06-30/schedules][%d] createReportScheduleTooManyRequests  %+v", 429, o.Payload)
}

func (o *CreateReportScheduleTooManyRequests) GetPayload() *reports_2021_06_30_models.ErrorList {
	return o.Payload
}

func (o *CreateReportScheduleTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(reports_2021_06_30_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateReportScheduleInternalServerError creates a CreateReportScheduleInternalServerError with default headers values
func NewCreateReportScheduleInternalServerError() *CreateReportScheduleInternalServerError {
	return &CreateReportScheduleInternalServerError{}
}

/*
CreateReportScheduleInternalServerError describes a response with status code 500, with default header values.

An unexpected condition occurred that prevented the server from fulfilling the request.
*/
type CreateReportScheduleInternalServerError struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *reports_2021_06_30_models.ErrorList
}

// IsSuccess returns true when this create report schedule internal server error response has a 2xx status code
func (o *CreateReportScheduleInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create report schedule internal server error response has a 3xx status code
func (o *CreateReportScheduleInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create report schedule internal server error response has a 4xx status code
func (o *CreateReportScheduleInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create report schedule internal server error response has a 5xx status code
func (o *CreateReportScheduleInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create report schedule internal server error response a status code equal to that given
func (o *CreateReportScheduleInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *CreateReportScheduleInternalServerError) Error() string {
	return fmt.Sprintf("[POST /reports/2021-06-30/schedules][%d] createReportScheduleInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateReportScheduleInternalServerError) String() string {
	return fmt.Sprintf("[POST /reports/2021-06-30/schedules][%d] createReportScheduleInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateReportScheduleInternalServerError) GetPayload() *reports_2021_06_30_models.ErrorList {
	return o.Payload
}

func (o *CreateReportScheduleInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(reports_2021_06_30_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateReportScheduleServiceUnavailable creates a CreateReportScheduleServiceUnavailable with default headers values
func NewCreateReportScheduleServiceUnavailable() *CreateReportScheduleServiceUnavailable {
	return &CreateReportScheduleServiceUnavailable{}
}

/*
CreateReportScheduleServiceUnavailable describes a response with status code 503, with default header values.

Temporary overloading or maintenance of the server.
*/
type CreateReportScheduleServiceUnavailable struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *reports_2021_06_30_models.ErrorList
}

// IsSuccess returns true when this create report schedule service unavailable response has a 2xx status code
func (o *CreateReportScheduleServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create report schedule service unavailable response has a 3xx status code
func (o *CreateReportScheduleServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create report schedule service unavailable response has a 4xx status code
func (o *CreateReportScheduleServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this create report schedule service unavailable response has a 5xx status code
func (o *CreateReportScheduleServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this create report schedule service unavailable response a status code equal to that given
func (o *CreateReportScheduleServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *CreateReportScheduleServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /reports/2021-06-30/schedules][%d] createReportScheduleServiceUnavailable  %+v", 503, o.Payload)
}

func (o *CreateReportScheduleServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /reports/2021-06-30/schedules][%d] createReportScheduleServiceUnavailable  %+v", 503, o.Payload)
}

func (o *CreateReportScheduleServiceUnavailable) GetPayload() *reports_2021_06_30_models.ErrorList {
	return o.Payload
}

func (o *CreateReportScheduleServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(reports_2021_06_30_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
