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

// CancelReportReader is a Reader for the CancelReport structure.
type CancelReportReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CancelReportReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCancelReportOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCancelReportBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCancelReportUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCancelReportForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCancelReportNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewCancelReportUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewCancelReportTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCancelReportInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewCancelReportServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCancelReportOK creates a CancelReportOK with default headers values
func NewCancelReportOK() *CancelReportOK {
	return &CancelReportOK{}
}

/*
CancelReportOK describes a response with status code 200, with default header values.

Success.
*/
type CancelReportOK struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string
}

// IsSuccess returns true when this cancel report o k response has a 2xx status code
func (o *CancelReportOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cancel report o k response has a 3xx status code
func (o *CancelReportOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cancel report o k response has a 4xx status code
func (o *CancelReportOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this cancel report o k response has a 5xx status code
func (o *CancelReportOK) IsServerError() bool {
	return false
}

// IsCode returns true when this cancel report o k response a status code equal to that given
func (o *CancelReportOK) IsCode(code int) bool {
	return code == 200
}

func (o *CancelReportOK) Error() string {
	return fmt.Sprintf("[DELETE /reports/2021-06-30/reports/{reportId}][%d] cancelReportOK ", 200)
}

func (o *CancelReportOK) String() string {
	return fmt.Sprintf("[DELETE /reports/2021-06-30/reports/{reportId}][%d] cancelReportOK ", 200)
}

func (o *CancelReportOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	return nil
}

// NewCancelReportBadRequest creates a CancelReportBadRequest with default headers values
func NewCancelReportBadRequest() *CancelReportBadRequest {
	return &CancelReportBadRequest{}
}

/*
CancelReportBadRequest describes a response with status code 400, with default header values.

Request has missing or invalid parameters and cannot be parsed.
*/
type CancelReportBadRequest struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *reports_2021_06_30_models.ErrorList
}

// IsSuccess returns true when this cancel report bad request response has a 2xx status code
func (o *CancelReportBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cancel report bad request response has a 3xx status code
func (o *CancelReportBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cancel report bad request response has a 4xx status code
func (o *CancelReportBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this cancel report bad request response has a 5xx status code
func (o *CancelReportBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this cancel report bad request response a status code equal to that given
func (o *CancelReportBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *CancelReportBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /reports/2021-06-30/reports/{reportId}][%d] cancelReportBadRequest  %+v", 400, o.Payload)
}

func (o *CancelReportBadRequest) String() string {
	return fmt.Sprintf("[DELETE /reports/2021-06-30/reports/{reportId}][%d] cancelReportBadRequest  %+v", 400, o.Payload)
}

func (o *CancelReportBadRequest) GetPayload() *reports_2021_06_30_models.ErrorList {
	return o.Payload
}

func (o *CancelReportBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCancelReportUnauthorized creates a CancelReportUnauthorized with default headers values
func NewCancelReportUnauthorized() *CancelReportUnauthorized {
	return &CancelReportUnauthorized{}
}

/*
CancelReportUnauthorized describes a response with status code 401, with default header values.

The request's Authorization header is not formatted correctly or does not contain a valid token.
*/
type CancelReportUnauthorized struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *reports_2021_06_30_models.ErrorList
}

// IsSuccess returns true when this cancel report unauthorized response has a 2xx status code
func (o *CancelReportUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cancel report unauthorized response has a 3xx status code
func (o *CancelReportUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cancel report unauthorized response has a 4xx status code
func (o *CancelReportUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this cancel report unauthorized response has a 5xx status code
func (o *CancelReportUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this cancel report unauthorized response a status code equal to that given
func (o *CancelReportUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *CancelReportUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /reports/2021-06-30/reports/{reportId}][%d] cancelReportUnauthorized  %+v", 401, o.Payload)
}

func (o *CancelReportUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /reports/2021-06-30/reports/{reportId}][%d] cancelReportUnauthorized  %+v", 401, o.Payload)
}

func (o *CancelReportUnauthorized) GetPayload() *reports_2021_06_30_models.ErrorList {
	return o.Payload
}

func (o *CancelReportUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCancelReportForbidden creates a CancelReportForbidden with default headers values
func NewCancelReportForbidden() *CancelReportForbidden {
	return &CancelReportForbidden{}
}

/*
CancelReportForbidden describes a response with status code 403, with default header values.

Indicates that access to the resource is forbidden. Possible reasons include Access Denied, Unauthorized, Expired Token, or Invalid Signature.
*/
type CancelReportForbidden struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *reports_2021_06_30_models.ErrorList
}

// IsSuccess returns true when this cancel report forbidden response has a 2xx status code
func (o *CancelReportForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cancel report forbidden response has a 3xx status code
func (o *CancelReportForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cancel report forbidden response has a 4xx status code
func (o *CancelReportForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this cancel report forbidden response has a 5xx status code
func (o *CancelReportForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this cancel report forbidden response a status code equal to that given
func (o *CancelReportForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *CancelReportForbidden) Error() string {
	return fmt.Sprintf("[DELETE /reports/2021-06-30/reports/{reportId}][%d] cancelReportForbidden  %+v", 403, o.Payload)
}

func (o *CancelReportForbidden) String() string {
	return fmt.Sprintf("[DELETE /reports/2021-06-30/reports/{reportId}][%d] cancelReportForbidden  %+v", 403, o.Payload)
}

func (o *CancelReportForbidden) GetPayload() *reports_2021_06_30_models.ErrorList {
	return o.Payload
}

func (o *CancelReportForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCancelReportNotFound creates a CancelReportNotFound with default headers values
func NewCancelReportNotFound() *CancelReportNotFound {
	return &CancelReportNotFound{}
}

/*
CancelReportNotFound describes a response with status code 404, with default header values.

The specified resource does not exist.
*/
type CancelReportNotFound struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *reports_2021_06_30_models.ErrorList
}

// IsSuccess returns true when this cancel report not found response has a 2xx status code
func (o *CancelReportNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cancel report not found response has a 3xx status code
func (o *CancelReportNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cancel report not found response has a 4xx status code
func (o *CancelReportNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this cancel report not found response has a 5xx status code
func (o *CancelReportNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this cancel report not found response a status code equal to that given
func (o *CancelReportNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *CancelReportNotFound) Error() string {
	return fmt.Sprintf("[DELETE /reports/2021-06-30/reports/{reportId}][%d] cancelReportNotFound  %+v", 404, o.Payload)
}

func (o *CancelReportNotFound) String() string {
	return fmt.Sprintf("[DELETE /reports/2021-06-30/reports/{reportId}][%d] cancelReportNotFound  %+v", 404, o.Payload)
}

func (o *CancelReportNotFound) GetPayload() *reports_2021_06_30_models.ErrorList {
	return o.Payload
}

func (o *CancelReportNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCancelReportUnsupportedMediaType creates a CancelReportUnsupportedMediaType with default headers values
func NewCancelReportUnsupportedMediaType() *CancelReportUnsupportedMediaType {
	return &CancelReportUnsupportedMediaType{}
}

/*
CancelReportUnsupportedMediaType describes a response with status code 415, with default header values.

The request's Content-Type header is invalid.
*/
type CancelReportUnsupportedMediaType struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *reports_2021_06_30_models.ErrorList
}

// IsSuccess returns true when this cancel report unsupported media type response has a 2xx status code
func (o *CancelReportUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cancel report unsupported media type response has a 3xx status code
func (o *CancelReportUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cancel report unsupported media type response has a 4xx status code
func (o *CancelReportUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this cancel report unsupported media type response has a 5xx status code
func (o *CancelReportUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this cancel report unsupported media type response a status code equal to that given
func (o *CancelReportUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *CancelReportUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /reports/2021-06-30/reports/{reportId}][%d] cancelReportUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *CancelReportUnsupportedMediaType) String() string {
	return fmt.Sprintf("[DELETE /reports/2021-06-30/reports/{reportId}][%d] cancelReportUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *CancelReportUnsupportedMediaType) GetPayload() *reports_2021_06_30_models.ErrorList {
	return o.Payload
}

func (o *CancelReportUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCancelReportTooManyRequests creates a CancelReportTooManyRequests with default headers values
func NewCancelReportTooManyRequests() *CancelReportTooManyRequests {
	return &CancelReportTooManyRequests{}
}

/*
CancelReportTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type CancelReportTooManyRequests struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *reports_2021_06_30_models.ErrorList
}

// IsSuccess returns true when this cancel report too many requests response has a 2xx status code
func (o *CancelReportTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cancel report too many requests response has a 3xx status code
func (o *CancelReportTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cancel report too many requests response has a 4xx status code
func (o *CancelReportTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this cancel report too many requests response has a 5xx status code
func (o *CancelReportTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this cancel report too many requests response a status code equal to that given
func (o *CancelReportTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *CancelReportTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /reports/2021-06-30/reports/{reportId}][%d] cancelReportTooManyRequests  %+v", 429, o.Payload)
}

func (o *CancelReportTooManyRequests) String() string {
	return fmt.Sprintf("[DELETE /reports/2021-06-30/reports/{reportId}][%d] cancelReportTooManyRequests  %+v", 429, o.Payload)
}

func (o *CancelReportTooManyRequests) GetPayload() *reports_2021_06_30_models.ErrorList {
	return o.Payload
}

func (o *CancelReportTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCancelReportInternalServerError creates a CancelReportInternalServerError with default headers values
func NewCancelReportInternalServerError() *CancelReportInternalServerError {
	return &CancelReportInternalServerError{}
}

/*
CancelReportInternalServerError describes a response with status code 500, with default header values.

An unexpected condition occurred that prevented the server from fulfilling the request.
*/
type CancelReportInternalServerError struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *reports_2021_06_30_models.ErrorList
}

// IsSuccess returns true when this cancel report internal server error response has a 2xx status code
func (o *CancelReportInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cancel report internal server error response has a 3xx status code
func (o *CancelReportInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cancel report internal server error response has a 4xx status code
func (o *CancelReportInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this cancel report internal server error response has a 5xx status code
func (o *CancelReportInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this cancel report internal server error response a status code equal to that given
func (o *CancelReportInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *CancelReportInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /reports/2021-06-30/reports/{reportId}][%d] cancelReportInternalServerError  %+v", 500, o.Payload)
}

func (o *CancelReportInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /reports/2021-06-30/reports/{reportId}][%d] cancelReportInternalServerError  %+v", 500, o.Payload)
}

func (o *CancelReportInternalServerError) GetPayload() *reports_2021_06_30_models.ErrorList {
	return o.Payload
}

func (o *CancelReportInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCancelReportServiceUnavailable creates a CancelReportServiceUnavailable with default headers values
func NewCancelReportServiceUnavailable() *CancelReportServiceUnavailable {
	return &CancelReportServiceUnavailable{}
}

/*
CancelReportServiceUnavailable describes a response with status code 503, with default header values.

Temporary overloading or maintenance of the server.
*/
type CancelReportServiceUnavailable struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *reports_2021_06_30_models.ErrorList
}

// IsSuccess returns true when this cancel report service unavailable response has a 2xx status code
func (o *CancelReportServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cancel report service unavailable response has a 3xx status code
func (o *CancelReportServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cancel report service unavailable response has a 4xx status code
func (o *CancelReportServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this cancel report service unavailable response has a 5xx status code
func (o *CancelReportServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this cancel report service unavailable response a status code equal to that given
func (o *CancelReportServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *CancelReportServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /reports/2021-06-30/reports/{reportId}][%d] cancelReportServiceUnavailable  %+v", 503, o.Payload)
}

func (o *CancelReportServiceUnavailable) String() string {
	return fmt.Sprintf("[DELETE /reports/2021-06-30/reports/{reportId}][%d] cancelReportServiceUnavailable  %+v", 503, o.Payload)
}

func (o *CancelReportServiceUnavailable) GetPayload() *reports_2021_06_30_models.ErrorList {
	return o.Payload
}

func (o *CancelReportServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
