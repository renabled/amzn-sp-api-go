// Code generated by go-swagger; DO NOT EDIT.

package reports

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/xamandar/amzn-sp-api-go/api/reports_2021-06-30/reports_2021_06_30_models"
)

// GetReportReader is a Reader for the GetReport structure.
type GetReportReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetReportReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetReportOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetReportBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetReportUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetReportForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetReportNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetReportUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetReportTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetReportInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetReportServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetReportOK creates a GetReportOK with default headers values
func NewGetReportOK() *GetReportOK {
	return &GetReportOK{}
}

/* GetReportOK describes a response with status code 200, with default header values.

Success.
*/
type GetReportOK struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *reports_2021_06_30_models.Report
}

func (o *GetReportOK) Error() string {
	return fmt.Sprintf("[GET /reports/2021-06-30/reports/{reportId}][%d] getReportOK  %+v", 200, o.Payload)
}
func (o *GetReportOK) GetPayload() *reports_2021_06_30_models.Report {
	return o.Payload
}

func (o *GetReportOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(reports_2021_06_30_models.Report)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReportBadRequest creates a GetReportBadRequest with default headers values
func NewGetReportBadRequest() *GetReportBadRequest {
	return &GetReportBadRequest{}
}

/* GetReportBadRequest describes a response with status code 400, with default header values.

Request has missing or invalid parameters and cannot be parsed.
*/
type GetReportBadRequest struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *reports_2021_06_30_models.ErrorList
}

func (o *GetReportBadRequest) Error() string {
	return fmt.Sprintf("[GET /reports/2021-06-30/reports/{reportId}][%d] getReportBadRequest  %+v", 400, o.Payload)
}
func (o *GetReportBadRequest) GetPayload() *reports_2021_06_30_models.ErrorList {
	return o.Payload
}

func (o *GetReportBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetReportUnauthorized creates a GetReportUnauthorized with default headers values
func NewGetReportUnauthorized() *GetReportUnauthorized {
	return &GetReportUnauthorized{}
}

/* GetReportUnauthorized describes a response with status code 401, with default header values.

The request's Authorization header is not formatted correctly or does not contain a valid token.
*/
type GetReportUnauthorized struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *reports_2021_06_30_models.ErrorList
}

func (o *GetReportUnauthorized) Error() string {
	return fmt.Sprintf("[GET /reports/2021-06-30/reports/{reportId}][%d] getReportUnauthorized  %+v", 401, o.Payload)
}
func (o *GetReportUnauthorized) GetPayload() *reports_2021_06_30_models.ErrorList {
	return o.Payload
}

func (o *GetReportUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetReportForbidden creates a GetReportForbidden with default headers values
func NewGetReportForbidden() *GetReportForbidden {
	return &GetReportForbidden{}
}

/* GetReportForbidden describes a response with status code 403, with default header values.

Indicates that access to the resource is forbidden. Possible reasons include Access Denied, Unauthorized, Expired Token, or Invalid Signature.
*/
type GetReportForbidden struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *reports_2021_06_30_models.ErrorList
}

func (o *GetReportForbidden) Error() string {
	return fmt.Sprintf("[GET /reports/2021-06-30/reports/{reportId}][%d] getReportForbidden  %+v", 403, o.Payload)
}
func (o *GetReportForbidden) GetPayload() *reports_2021_06_30_models.ErrorList {
	return o.Payload
}

func (o *GetReportForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetReportNotFound creates a GetReportNotFound with default headers values
func NewGetReportNotFound() *GetReportNotFound {
	return &GetReportNotFound{}
}

/* GetReportNotFound describes a response with status code 404, with default header values.

The specified resource does not exist.
*/
type GetReportNotFound struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *reports_2021_06_30_models.ErrorList
}

func (o *GetReportNotFound) Error() string {
	return fmt.Sprintf("[GET /reports/2021-06-30/reports/{reportId}][%d] getReportNotFound  %+v", 404, o.Payload)
}
func (o *GetReportNotFound) GetPayload() *reports_2021_06_30_models.ErrorList {
	return o.Payload
}

func (o *GetReportNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetReportUnsupportedMediaType creates a GetReportUnsupportedMediaType with default headers values
func NewGetReportUnsupportedMediaType() *GetReportUnsupportedMediaType {
	return &GetReportUnsupportedMediaType{}
}

/* GetReportUnsupportedMediaType describes a response with status code 415, with default header values.

The request's Content-Type header is invalid.
*/
type GetReportUnsupportedMediaType struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *reports_2021_06_30_models.ErrorList
}

func (o *GetReportUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /reports/2021-06-30/reports/{reportId}][%d] getReportUnsupportedMediaType  %+v", 415, o.Payload)
}
func (o *GetReportUnsupportedMediaType) GetPayload() *reports_2021_06_30_models.ErrorList {
	return o.Payload
}

func (o *GetReportUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetReportTooManyRequests creates a GetReportTooManyRequests with default headers values
func NewGetReportTooManyRequests() *GetReportTooManyRequests {
	return &GetReportTooManyRequests{}
}

/* GetReportTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type GetReportTooManyRequests struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *reports_2021_06_30_models.ErrorList
}

func (o *GetReportTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /reports/2021-06-30/reports/{reportId}][%d] getReportTooManyRequests  %+v", 429, o.Payload)
}
func (o *GetReportTooManyRequests) GetPayload() *reports_2021_06_30_models.ErrorList {
	return o.Payload
}

func (o *GetReportTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetReportInternalServerError creates a GetReportInternalServerError with default headers values
func NewGetReportInternalServerError() *GetReportInternalServerError {
	return &GetReportInternalServerError{}
}

/* GetReportInternalServerError describes a response with status code 500, with default header values.

An unexpected condition occurred that prevented the server from fulfilling the request.
*/
type GetReportInternalServerError struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *reports_2021_06_30_models.ErrorList
}

func (o *GetReportInternalServerError) Error() string {
	return fmt.Sprintf("[GET /reports/2021-06-30/reports/{reportId}][%d] getReportInternalServerError  %+v", 500, o.Payload)
}
func (o *GetReportInternalServerError) GetPayload() *reports_2021_06_30_models.ErrorList {
	return o.Payload
}

func (o *GetReportInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetReportServiceUnavailable creates a GetReportServiceUnavailable with default headers values
func NewGetReportServiceUnavailable() *GetReportServiceUnavailable {
	return &GetReportServiceUnavailable{}
}

/* GetReportServiceUnavailable describes a response with status code 503, with default header values.

Temporary overloading or maintenance of the server.
*/
type GetReportServiceUnavailable struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *reports_2021_06_30_models.ErrorList
}

func (o *GetReportServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /reports/2021-06-30/reports/{reportId}][%d] getReportServiceUnavailable  %+v", 503, o.Payload)
}
func (o *GetReportServiceUnavailable) GetPayload() *reports_2021_06_30_models.ErrorList {
	return o.Payload
}

func (o *GetReportServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
