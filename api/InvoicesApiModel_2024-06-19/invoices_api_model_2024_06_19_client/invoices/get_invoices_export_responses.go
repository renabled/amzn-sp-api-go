// Code generated by go-swagger; DO NOT EDIT.

package invoices

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/renabled/amzn-sp-api-go/api/InvoicesApiModel_2024-06-19/invoices_api_model_2024_06_19_models"
)

// GetInvoicesExportReader is a Reader for the GetInvoicesExport structure.
type GetInvoicesExportReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInvoicesExportReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetInvoicesExportOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetInvoicesExportBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetInvoicesExportUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetInvoicesExportForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetInvoicesExportNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetInvoicesExportRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetInvoicesExportUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetInvoicesExportTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetInvoicesExportInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetInvoicesExportServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetInvoicesExportOK creates a GetInvoicesExportOK with default headers values
func NewGetInvoicesExportOK() *GetInvoicesExportOK {
	return &GetInvoicesExportOK{}
}

/*
GetInvoicesExportOK describes a response with status code 200, with default header values.

Success.
*/
type GetInvoicesExportOK struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *invoices_api_model_2024_06_19_models.GetInvoicesExportResponse
}

// IsSuccess returns true when this get invoices export o k response has a 2xx status code
func (o *GetInvoicesExportOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get invoices export o k response has a 3xx status code
func (o *GetInvoicesExportOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get invoices export o k response has a 4xx status code
func (o *GetInvoicesExportOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get invoices export o k response has a 5xx status code
func (o *GetInvoicesExportOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get invoices export o k response a status code equal to that given
func (o *GetInvoicesExportOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetInvoicesExportOK) Error() string {
	return fmt.Sprintf("[GET /tax/invoices/2024-06-19/exports/{exportId}][%d] getInvoicesExportOK  %+v", 200, o.Payload)
}

func (o *GetInvoicesExportOK) String() string {
	return fmt.Sprintf("[GET /tax/invoices/2024-06-19/exports/{exportId}][%d] getInvoicesExportOK  %+v", 200, o.Payload)
}

func (o *GetInvoicesExportOK) GetPayload() *invoices_api_model_2024_06_19_models.GetInvoicesExportResponse {
	return o.Payload
}

func (o *GetInvoicesExportOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(invoices_api_model_2024_06_19_models.GetInvoicesExportResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInvoicesExportBadRequest creates a GetInvoicesExportBadRequest with default headers values
func NewGetInvoicesExportBadRequest() *GetInvoicesExportBadRequest {
	return &GetInvoicesExportBadRequest{}
}

/*
GetInvoicesExportBadRequest describes a response with status code 400, with default header values.

Request has missing or invalid parameters and cannot be parsed.
*/
type GetInvoicesExportBadRequest struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *invoices_api_model_2024_06_19_models.ErrorList
}

// IsSuccess returns true when this get invoices export bad request response has a 2xx status code
func (o *GetInvoicesExportBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get invoices export bad request response has a 3xx status code
func (o *GetInvoicesExportBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get invoices export bad request response has a 4xx status code
func (o *GetInvoicesExportBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get invoices export bad request response has a 5xx status code
func (o *GetInvoicesExportBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get invoices export bad request response a status code equal to that given
func (o *GetInvoicesExportBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetInvoicesExportBadRequest) Error() string {
	return fmt.Sprintf("[GET /tax/invoices/2024-06-19/exports/{exportId}][%d] getInvoicesExportBadRequest  %+v", 400, o.Payload)
}

func (o *GetInvoicesExportBadRequest) String() string {
	return fmt.Sprintf("[GET /tax/invoices/2024-06-19/exports/{exportId}][%d] getInvoicesExportBadRequest  %+v", 400, o.Payload)
}

func (o *GetInvoicesExportBadRequest) GetPayload() *invoices_api_model_2024_06_19_models.ErrorList {
	return o.Payload
}

func (o *GetInvoicesExportBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(invoices_api_model_2024_06_19_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInvoicesExportUnauthorized creates a GetInvoicesExportUnauthorized with default headers values
func NewGetInvoicesExportUnauthorized() *GetInvoicesExportUnauthorized {
	return &GetInvoicesExportUnauthorized{}
}

/*
GetInvoicesExportUnauthorized describes a response with status code 401, with default header values.

A list of error responses returned when a request is unsuccessful.
*/
type GetInvoicesExportUnauthorized struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *invoices_api_model_2024_06_19_models.ErrorList
}

// IsSuccess returns true when this get invoices export unauthorized response has a 2xx status code
func (o *GetInvoicesExportUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get invoices export unauthorized response has a 3xx status code
func (o *GetInvoicesExportUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get invoices export unauthorized response has a 4xx status code
func (o *GetInvoicesExportUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get invoices export unauthorized response has a 5xx status code
func (o *GetInvoicesExportUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get invoices export unauthorized response a status code equal to that given
func (o *GetInvoicesExportUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetInvoicesExportUnauthorized) Error() string {
	return fmt.Sprintf("[GET /tax/invoices/2024-06-19/exports/{exportId}][%d] getInvoicesExportUnauthorized  %+v", 401, o.Payload)
}

func (o *GetInvoicesExportUnauthorized) String() string {
	return fmt.Sprintf("[GET /tax/invoices/2024-06-19/exports/{exportId}][%d] getInvoicesExportUnauthorized  %+v", 401, o.Payload)
}

func (o *GetInvoicesExportUnauthorized) GetPayload() *invoices_api_model_2024_06_19_models.ErrorList {
	return o.Payload
}

func (o *GetInvoicesExportUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(invoices_api_model_2024_06_19_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInvoicesExportForbidden creates a GetInvoicesExportForbidden with default headers values
func NewGetInvoicesExportForbidden() *GetInvoicesExportForbidden {
	return &GetInvoicesExportForbidden{}
}

/*
GetInvoicesExportForbidden describes a response with status code 403, with default header values.

Indicates that access to the resource is forbidden. Possible reasons include Access Denied, Unauthorized, Expired Token, or Invalid Signature.
*/
type GetInvoicesExportForbidden struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *invoices_api_model_2024_06_19_models.ErrorList
}

// IsSuccess returns true when this get invoices export forbidden response has a 2xx status code
func (o *GetInvoicesExportForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get invoices export forbidden response has a 3xx status code
func (o *GetInvoicesExportForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get invoices export forbidden response has a 4xx status code
func (o *GetInvoicesExportForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get invoices export forbidden response has a 5xx status code
func (o *GetInvoicesExportForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get invoices export forbidden response a status code equal to that given
func (o *GetInvoicesExportForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetInvoicesExportForbidden) Error() string {
	return fmt.Sprintf("[GET /tax/invoices/2024-06-19/exports/{exportId}][%d] getInvoicesExportForbidden  %+v", 403, o.Payload)
}

func (o *GetInvoicesExportForbidden) String() string {
	return fmt.Sprintf("[GET /tax/invoices/2024-06-19/exports/{exportId}][%d] getInvoicesExportForbidden  %+v", 403, o.Payload)
}

func (o *GetInvoicesExportForbidden) GetPayload() *invoices_api_model_2024_06_19_models.ErrorList {
	return o.Payload
}

func (o *GetInvoicesExportForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(invoices_api_model_2024_06_19_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInvoicesExportNotFound creates a GetInvoicesExportNotFound with default headers values
func NewGetInvoicesExportNotFound() *GetInvoicesExportNotFound {
	return &GetInvoicesExportNotFound{}
}

/*
GetInvoicesExportNotFound describes a response with status code 404, with default header values.

The resource specified does not exist.
*/
type GetInvoicesExportNotFound struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *invoices_api_model_2024_06_19_models.ErrorList
}

// IsSuccess returns true when this get invoices export not found response has a 2xx status code
func (o *GetInvoicesExportNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get invoices export not found response has a 3xx status code
func (o *GetInvoicesExportNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get invoices export not found response has a 4xx status code
func (o *GetInvoicesExportNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get invoices export not found response has a 5xx status code
func (o *GetInvoicesExportNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get invoices export not found response a status code equal to that given
func (o *GetInvoicesExportNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetInvoicesExportNotFound) Error() string {
	return fmt.Sprintf("[GET /tax/invoices/2024-06-19/exports/{exportId}][%d] getInvoicesExportNotFound  %+v", 404, o.Payload)
}

func (o *GetInvoicesExportNotFound) String() string {
	return fmt.Sprintf("[GET /tax/invoices/2024-06-19/exports/{exportId}][%d] getInvoicesExportNotFound  %+v", 404, o.Payload)
}

func (o *GetInvoicesExportNotFound) GetPayload() *invoices_api_model_2024_06_19_models.ErrorList {
	return o.Payload
}

func (o *GetInvoicesExportNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(invoices_api_model_2024_06_19_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInvoicesExportRequestEntityTooLarge creates a GetInvoicesExportRequestEntityTooLarge with default headers values
func NewGetInvoicesExportRequestEntityTooLarge() *GetInvoicesExportRequestEntityTooLarge {
	return &GetInvoicesExportRequestEntityTooLarge{}
}

/*
GetInvoicesExportRequestEntityTooLarge describes a response with status code 413, with default header values.

The request size exceeded the maximum accepted size.
*/
type GetInvoicesExportRequestEntityTooLarge struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *invoices_api_model_2024_06_19_models.ErrorList
}

// IsSuccess returns true when this get invoices export request entity too large response has a 2xx status code
func (o *GetInvoicesExportRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get invoices export request entity too large response has a 3xx status code
func (o *GetInvoicesExportRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get invoices export request entity too large response has a 4xx status code
func (o *GetInvoicesExportRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get invoices export request entity too large response has a 5xx status code
func (o *GetInvoicesExportRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get invoices export request entity too large response a status code equal to that given
func (o *GetInvoicesExportRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetInvoicesExportRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /tax/invoices/2024-06-19/exports/{exportId}][%d] getInvoicesExportRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetInvoicesExportRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /tax/invoices/2024-06-19/exports/{exportId}][%d] getInvoicesExportRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetInvoicesExportRequestEntityTooLarge) GetPayload() *invoices_api_model_2024_06_19_models.ErrorList {
	return o.Payload
}

func (o *GetInvoicesExportRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(invoices_api_model_2024_06_19_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInvoicesExportUnsupportedMediaType creates a GetInvoicesExportUnsupportedMediaType with default headers values
func NewGetInvoicesExportUnsupportedMediaType() *GetInvoicesExportUnsupportedMediaType {
	return &GetInvoicesExportUnsupportedMediaType{}
}

/*
GetInvoicesExportUnsupportedMediaType describes a response with status code 415, with default header values.

The request payload is in an unsupported format.
*/
type GetInvoicesExportUnsupportedMediaType struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *invoices_api_model_2024_06_19_models.ErrorList
}

// IsSuccess returns true when this get invoices export unsupported media type response has a 2xx status code
func (o *GetInvoicesExportUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get invoices export unsupported media type response has a 3xx status code
func (o *GetInvoicesExportUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get invoices export unsupported media type response has a 4xx status code
func (o *GetInvoicesExportUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get invoices export unsupported media type response has a 5xx status code
func (o *GetInvoicesExportUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get invoices export unsupported media type response a status code equal to that given
func (o *GetInvoicesExportUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetInvoicesExportUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /tax/invoices/2024-06-19/exports/{exportId}][%d] getInvoicesExportUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetInvoicesExportUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /tax/invoices/2024-06-19/exports/{exportId}][%d] getInvoicesExportUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetInvoicesExportUnsupportedMediaType) GetPayload() *invoices_api_model_2024_06_19_models.ErrorList {
	return o.Payload
}

func (o *GetInvoicesExportUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(invoices_api_model_2024_06_19_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInvoicesExportTooManyRequests creates a GetInvoicesExportTooManyRequests with default headers values
func NewGetInvoicesExportTooManyRequests() *GetInvoicesExportTooManyRequests {
	return &GetInvoicesExportTooManyRequests{}
}

/*
GetInvoicesExportTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type GetInvoicesExportTooManyRequests struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *invoices_api_model_2024_06_19_models.ErrorList
}

// IsSuccess returns true when this get invoices export too many requests response has a 2xx status code
func (o *GetInvoicesExportTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get invoices export too many requests response has a 3xx status code
func (o *GetInvoicesExportTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get invoices export too many requests response has a 4xx status code
func (o *GetInvoicesExportTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get invoices export too many requests response has a 5xx status code
func (o *GetInvoicesExportTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get invoices export too many requests response a status code equal to that given
func (o *GetInvoicesExportTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetInvoicesExportTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /tax/invoices/2024-06-19/exports/{exportId}][%d] getInvoicesExportTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetInvoicesExportTooManyRequests) String() string {
	return fmt.Sprintf("[GET /tax/invoices/2024-06-19/exports/{exportId}][%d] getInvoicesExportTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetInvoicesExportTooManyRequests) GetPayload() *invoices_api_model_2024_06_19_models.ErrorList {
	return o.Payload
}

func (o *GetInvoicesExportTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(invoices_api_model_2024_06_19_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInvoicesExportInternalServerError creates a GetInvoicesExportInternalServerError with default headers values
func NewGetInvoicesExportInternalServerError() *GetInvoicesExportInternalServerError {
	return &GetInvoicesExportInternalServerError{}
}

/*
GetInvoicesExportInternalServerError describes a response with status code 500, with default header values.

An unexpected condition occurred that prevented the server from fulfilling the request.
*/
type GetInvoicesExportInternalServerError struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *invoices_api_model_2024_06_19_models.ErrorList
}

// IsSuccess returns true when this get invoices export internal server error response has a 2xx status code
func (o *GetInvoicesExportInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get invoices export internal server error response has a 3xx status code
func (o *GetInvoicesExportInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get invoices export internal server error response has a 4xx status code
func (o *GetInvoicesExportInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get invoices export internal server error response has a 5xx status code
func (o *GetInvoicesExportInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get invoices export internal server error response a status code equal to that given
func (o *GetInvoicesExportInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetInvoicesExportInternalServerError) Error() string {
	return fmt.Sprintf("[GET /tax/invoices/2024-06-19/exports/{exportId}][%d] getInvoicesExportInternalServerError  %+v", 500, o.Payload)
}

func (o *GetInvoicesExportInternalServerError) String() string {
	return fmt.Sprintf("[GET /tax/invoices/2024-06-19/exports/{exportId}][%d] getInvoicesExportInternalServerError  %+v", 500, o.Payload)
}

func (o *GetInvoicesExportInternalServerError) GetPayload() *invoices_api_model_2024_06_19_models.ErrorList {
	return o.Payload
}

func (o *GetInvoicesExportInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(invoices_api_model_2024_06_19_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInvoicesExportServiceUnavailable creates a GetInvoicesExportServiceUnavailable with default headers values
func NewGetInvoicesExportServiceUnavailable() *GetInvoicesExportServiceUnavailable {
	return &GetInvoicesExportServiceUnavailable{}
}

/*
GetInvoicesExportServiceUnavailable describes a response with status code 503, with default header values.

Temporary overloading or maintenance of the server.
*/
type GetInvoicesExportServiceUnavailable struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *invoices_api_model_2024_06_19_models.ErrorList
}

// IsSuccess returns true when this get invoices export service unavailable response has a 2xx status code
func (o *GetInvoicesExportServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get invoices export service unavailable response has a 3xx status code
func (o *GetInvoicesExportServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get invoices export service unavailable response has a 4xx status code
func (o *GetInvoicesExportServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get invoices export service unavailable response has a 5xx status code
func (o *GetInvoicesExportServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get invoices export service unavailable response a status code equal to that given
func (o *GetInvoicesExportServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetInvoicesExportServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /tax/invoices/2024-06-19/exports/{exportId}][%d] getInvoicesExportServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetInvoicesExportServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /tax/invoices/2024-06-19/exports/{exportId}][%d] getInvoicesExportServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetInvoicesExportServiceUnavailable) GetPayload() *invoices_api_model_2024_06_19_models.ErrorList {
	return o.Payload
}

func (o *GetInvoicesExportServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(invoices_api_model_2024_06_19_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
