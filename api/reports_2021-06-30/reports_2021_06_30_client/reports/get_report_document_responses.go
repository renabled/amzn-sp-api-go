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

// GetReportDocumentReader is a Reader for the GetReportDocument structure.
type GetReportDocumentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetReportDocumentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetReportDocumentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetReportDocumentBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetReportDocumentUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetReportDocumentForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetReportDocumentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetReportDocumentUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetReportDocumentTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetReportDocumentInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetReportDocumentServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetReportDocumentOK creates a GetReportDocumentOK with default headers values
func NewGetReportDocumentOK() *GetReportDocumentOK {
	return &GetReportDocumentOK{}
}

/*
GetReportDocumentOK describes a response with status code 200, with default header values.

Success.
*/
type GetReportDocumentOK struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *reports_2021_06_30_models.ReportDocument
}

// IsSuccess returns true when this get report document o k response has a 2xx status code
func (o *GetReportDocumentOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get report document o k response has a 3xx status code
func (o *GetReportDocumentOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get report document o k response has a 4xx status code
func (o *GetReportDocumentOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get report document o k response has a 5xx status code
func (o *GetReportDocumentOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get report document o k response a status code equal to that given
func (o *GetReportDocumentOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetReportDocumentOK) Error() string {
	return fmt.Sprintf("[GET /reports/2021-06-30/documents/{reportDocumentId}][%d] getReportDocumentOK  %+v", 200, o.Payload)
}

func (o *GetReportDocumentOK) String() string {
	return fmt.Sprintf("[GET /reports/2021-06-30/documents/{reportDocumentId}][%d] getReportDocumentOK  %+v", 200, o.Payload)
}

func (o *GetReportDocumentOK) GetPayload() *reports_2021_06_30_models.ReportDocument {
	return o.Payload
}

func (o *GetReportDocumentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(reports_2021_06_30_models.ReportDocument)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReportDocumentBadRequest creates a GetReportDocumentBadRequest with default headers values
func NewGetReportDocumentBadRequest() *GetReportDocumentBadRequest {
	return &GetReportDocumentBadRequest{}
}

/*
GetReportDocumentBadRequest describes a response with status code 400, with default header values.

Request has missing or invalid parameters and cannot be parsed.
*/
type GetReportDocumentBadRequest struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *reports_2021_06_30_models.ErrorList
}

// IsSuccess returns true when this get report document bad request response has a 2xx status code
func (o *GetReportDocumentBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get report document bad request response has a 3xx status code
func (o *GetReportDocumentBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get report document bad request response has a 4xx status code
func (o *GetReportDocumentBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get report document bad request response has a 5xx status code
func (o *GetReportDocumentBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get report document bad request response a status code equal to that given
func (o *GetReportDocumentBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetReportDocumentBadRequest) Error() string {
	return fmt.Sprintf("[GET /reports/2021-06-30/documents/{reportDocumentId}][%d] getReportDocumentBadRequest  %+v", 400, o.Payload)
}

func (o *GetReportDocumentBadRequest) String() string {
	return fmt.Sprintf("[GET /reports/2021-06-30/documents/{reportDocumentId}][%d] getReportDocumentBadRequest  %+v", 400, o.Payload)
}

func (o *GetReportDocumentBadRequest) GetPayload() *reports_2021_06_30_models.ErrorList {
	return o.Payload
}

func (o *GetReportDocumentBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetReportDocumentUnauthorized creates a GetReportDocumentUnauthorized with default headers values
func NewGetReportDocumentUnauthorized() *GetReportDocumentUnauthorized {
	return &GetReportDocumentUnauthorized{}
}

/*
GetReportDocumentUnauthorized describes a response with status code 401, with default header values.

The request's Authorization header is not formatted correctly or does not contain a valid token.
*/
type GetReportDocumentUnauthorized struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *reports_2021_06_30_models.ErrorList
}

// IsSuccess returns true when this get report document unauthorized response has a 2xx status code
func (o *GetReportDocumentUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get report document unauthorized response has a 3xx status code
func (o *GetReportDocumentUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get report document unauthorized response has a 4xx status code
func (o *GetReportDocumentUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get report document unauthorized response has a 5xx status code
func (o *GetReportDocumentUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get report document unauthorized response a status code equal to that given
func (o *GetReportDocumentUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetReportDocumentUnauthorized) Error() string {
	return fmt.Sprintf("[GET /reports/2021-06-30/documents/{reportDocumentId}][%d] getReportDocumentUnauthorized  %+v", 401, o.Payload)
}

func (o *GetReportDocumentUnauthorized) String() string {
	return fmt.Sprintf("[GET /reports/2021-06-30/documents/{reportDocumentId}][%d] getReportDocumentUnauthorized  %+v", 401, o.Payload)
}

func (o *GetReportDocumentUnauthorized) GetPayload() *reports_2021_06_30_models.ErrorList {
	return o.Payload
}

func (o *GetReportDocumentUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetReportDocumentForbidden creates a GetReportDocumentForbidden with default headers values
func NewGetReportDocumentForbidden() *GetReportDocumentForbidden {
	return &GetReportDocumentForbidden{}
}

/*
GetReportDocumentForbidden describes a response with status code 403, with default header values.

Indicates access to the resource is forbidden. Possible reasons include Access Denied, Unauthorized, Expired Token, or Invalid Signature.
*/
type GetReportDocumentForbidden struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *reports_2021_06_30_models.ErrorList
}

// IsSuccess returns true when this get report document forbidden response has a 2xx status code
func (o *GetReportDocumentForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get report document forbidden response has a 3xx status code
func (o *GetReportDocumentForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get report document forbidden response has a 4xx status code
func (o *GetReportDocumentForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get report document forbidden response has a 5xx status code
func (o *GetReportDocumentForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get report document forbidden response a status code equal to that given
func (o *GetReportDocumentForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetReportDocumentForbidden) Error() string {
	return fmt.Sprintf("[GET /reports/2021-06-30/documents/{reportDocumentId}][%d] getReportDocumentForbidden  %+v", 403, o.Payload)
}

func (o *GetReportDocumentForbidden) String() string {
	return fmt.Sprintf("[GET /reports/2021-06-30/documents/{reportDocumentId}][%d] getReportDocumentForbidden  %+v", 403, o.Payload)
}

func (o *GetReportDocumentForbidden) GetPayload() *reports_2021_06_30_models.ErrorList {
	return o.Payload
}

func (o *GetReportDocumentForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetReportDocumentNotFound creates a GetReportDocumentNotFound with default headers values
func NewGetReportDocumentNotFound() *GetReportDocumentNotFound {
	return &GetReportDocumentNotFound{}
}

/*
GetReportDocumentNotFound describes a response with status code 404, with default header values.

The specified resource does not exist.
*/
type GetReportDocumentNotFound struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *reports_2021_06_30_models.ErrorList
}

// IsSuccess returns true when this get report document not found response has a 2xx status code
func (o *GetReportDocumentNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get report document not found response has a 3xx status code
func (o *GetReportDocumentNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get report document not found response has a 4xx status code
func (o *GetReportDocumentNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get report document not found response has a 5xx status code
func (o *GetReportDocumentNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get report document not found response a status code equal to that given
func (o *GetReportDocumentNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetReportDocumentNotFound) Error() string {
	return fmt.Sprintf("[GET /reports/2021-06-30/documents/{reportDocumentId}][%d] getReportDocumentNotFound  %+v", 404, o.Payload)
}

func (o *GetReportDocumentNotFound) String() string {
	return fmt.Sprintf("[GET /reports/2021-06-30/documents/{reportDocumentId}][%d] getReportDocumentNotFound  %+v", 404, o.Payload)
}

func (o *GetReportDocumentNotFound) GetPayload() *reports_2021_06_30_models.ErrorList {
	return o.Payload
}

func (o *GetReportDocumentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetReportDocumentUnsupportedMediaType creates a GetReportDocumentUnsupportedMediaType with default headers values
func NewGetReportDocumentUnsupportedMediaType() *GetReportDocumentUnsupportedMediaType {
	return &GetReportDocumentUnsupportedMediaType{}
}

/*
GetReportDocumentUnsupportedMediaType describes a response with status code 415, with default header values.

The request's Content-Type header is invalid.
*/
type GetReportDocumentUnsupportedMediaType struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *reports_2021_06_30_models.ErrorList
}

// IsSuccess returns true when this get report document unsupported media type response has a 2xx status code
func (o *GetReportDocumentUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get report document unsupported media type response has a 3xx status code
func (o *GetReportDocumentUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get report document unsupported media type response has a 4xx status code
func (o *GetReportDocumentUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get report document unsupported media type response has a 5xx status code
func (o *GetReportDocumentUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get report document unsupported media type response a status code equal to that given
func (o *GetReportDocumentUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetReportDocumentUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /reports/2021-06-30/documents/{reportDocumentId}][%d] getReportDocumentUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetReportDocumentUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /reports/2021-06-30/documents/{reportDocumentId}][%d] getReportDocumentUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetReportDocumentUnsupportedMediaType) GetPayload() *reports_2021_06_30_models.ErrorList {
	return o.Payload
}

func (o *GetReportDocumentUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetReportDocumentTooManyRequests creates a GetReportDocumentTooManyRequests with default headers values
func NewGetReportDocumentTooManyRequests() *GetReportDocumentTooManyRequests {
	return &GetReportDocumentTooManyRequests{}
}

/*
GetReportDocumentTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type GetReportDocumentTooManyRequests struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *reports_2021_06_30_models.ErrorList
}

// IsSuccess returns true when this get report document too many requests response has a 2xx status code
func (o *GetReportDocumentTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get report document too many requests response has a 3xx status code
func (o *GetReportDocumentTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get report document too many requests response has a 4xx status code
func (o *GetReportDocumentTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get report document too many requests response has a 5xx status code
func (o *GetReportDocumentTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get report document too many requests response a status code equal to that given
func (o *GetReportDocumentTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetReportDocumentTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /reports/2021-06-30/documents/{reportDocumentId}][%d] getReportDocumentTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetReportDocumentTooManyRequests) String() string {
	return fmt.Sprintf("[GET /reports/2021-06-30/documents/{reportDocumentId}][%d] getReportDocumentTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetReportDocumentTooManyRequests) GetPayload() *reports_2021_06_30_models.ErrorList {
	return o.Payload
}

func (o *GetReportDocumentTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetReportDocumentInternalServerError creates a GetReportDocumentInternalServerError with default headers values
func NewGetReportDocumentInternalServerError() *GetReportDocumentInternalServerError {
	return &GetReportDocumentInternalServerError{}
}

/*
GetReportDocumentInternalServerError describes a response with status code 500, with default header values.

An unexpected condition occurred that prevented the server from fulfilling the request.
*/
type GetReportDocumentInternalServerError struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *reports_2021_06_30_models.ErrorList
}

// IsSuccess returns true when this get report document internal server error response has a 2xx status code
func (o *GetReportDocumentInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get report document internal server error response has a 3xx status code
func (o *GetReportDocumentInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get report document internal server error response has a 4xx status code
func (o *GetReportDocumentInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get report document internal server error response has a 5xx status code
func (o *GetReportDocumentInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get report document internal server error response a status code equal to that given
func (o *GetReportDocumentInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetReportDocumentInternalServerError) Error() string {
	return fmt.Sprintf("[GET /reports/2021-06-30/documents/{reportDocumentId}][%d] getReportDocumentInternalServerError  %+v", 500, o.Payload)
}

func (o *GetReportDocumentInternalServerError) String() string {
	return fmt.Sprintf("[GET /reports/2021-06-30/documents/{reportDocumentId}][%d] getReportDocumentInternalServerError  %+v", 500, o.Payload)
}

func (o *GetReportDocumentInternalServerError) GetPayload() *reports_2021_06_30_models.ErrorList {
	return o.Payload
}

func (o *GetReportDocumentInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetReportDocumentServiceUnavailable creates a GetReportDocumentServiceUnavailable with default headers values
func NewGetReportDocumentServiceUnavailable() *GetReportDocumentServiceUnavailable {
	return &GetReportDocumentServiceUnavailable{}
}

/*
GetReportDocumentServiceUnavailable describes a response with status code 503, with default header values.

Temporary overloading or maintenance of the server.
*/
type GetReportDocumentServiceUnavailable struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *reports_2021_06_30_models.ErrorList
}

// IsSuccess returns true when this get report document service unavailable response has a 2xx status code
func (o *GetReportDocumentServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get report document service unavailable response has a 3xx status code
func (o *GetReportDocumentServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get report document service unavailable response has a 4xx status code
func (o *GetReportDocumentServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get report document service unavailable response has a 5xx status code
func (o *GetReportDocumentServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get report document service unavailable response a status code equal to that given
func (o *GetReportDocumentServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetReportDocumentServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /reports/2021-06-30/documents/{reportDocumentId}][%d] getReportDocumentServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetReportDocumentServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /reports/2021-06-30/documents/{reportDocumentId}][%d] getReportDocumentServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetReportDocumentServiceUnavailable) GetPayload() *reports_2021_06_30_models.ErrorList {
	return o.Payload
}

func (o *GetReportDocumentServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
