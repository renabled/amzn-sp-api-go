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

// ArchiveSupplySourceReader is a Reader for the ArchiveSupplySource structure.
type ArchiveSupplySourceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ArchiveSupplySourceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewArchiveSupplySourceNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewArchiveSupplySourceBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewArchiveSupplySourceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewArchiveSupplySourceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewArchiveSupplySourceRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewArchiveSupplySourceUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewArchiveSupplySourceTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewArchiveSupplySourceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewArchiveSupplySourceServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewArchiveSupplySourceNoContent creates a ArchiveSupplySourceNoContent with default headers values
func NewArchiveSupplySourceNoContent() *ArchiveSupplySourceNoContent {
	return &ArchiveSupplySourceNoContent{}
}

/*
ArchiveSupplySourceNoContent describes a response with status code 204, with default header values.

Success.
*/
type ArchiveSupplySourceNoContent struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* The unique request reference ID.
	 */
	XAmznRequestID string

	Payload *supply_sources_2020_07_01_models.ErrorList
}

// IsSuccess returns true when this archive supply source no content response has a 2xx status code
func (o *ArchiveSupplySourceNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this archive supply source no content response has a 3xx status code
func (o *ArchiveSupplySourceNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this archive supply source no content response has a 4xx status code
func (o *ArchiveSupplySourceNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this archive supply source no content response has a 5xx status code
func (o *ArchiveSupplySourceNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this archive supply source no content response a status code equal to that given
func (o *ArchiveSupplySourceNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *ArchiveSupplySourceNoContent) Error() string {
	return fmt.Sprintf("[DELETE /supplySources/2020-07-01/supplySources/{supplySourceId}][%d] archiveSupplySourceNoContent  %+v", 204, o.Payload)
}

func (o *ArchiveSupplySourceNoContent) String() string {
	return fmt.Sprintf("[DELETE /supplySources/2020-07-01/supplySources/{supplySourceId}][%d] archiveSupplySourceNoContent  %+v", 204, o.Payload)
}

func (o *ArchiveSupplySourceNoContent) GetPayload() *supply_sources_2020_07_01_models.ErrorList {
	return o.Payload
}

func (o *ArchiveSupplySourceNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewArchiveSupplySourceBadRequest creates a ArchiveSupplySourceBadRequest with default headers values
func NewArchiveSupplySourceBadRequest() *ArchiveSupplySourceBadRequest {
	return &ArchiveSupplySourceBadRequest{}
}

/*
ArchiveSupplySourceBadRequest describes a response with status code 400, with default header values.

The request has missing or invalid parameters and cannot be parsed.
*/
type ArchiveSupplySourceBadRequest struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* The unique request reference ID.
	 */
	XAmznRequestID string

	Payload *supply_sources_2020_07_01_models.ErrorList
}

// IsSuccess returns true when this archive supply source bad request response has a 2xx status code
func (o *ArchiveSupplySourceBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this archive supply source bad request response has a 3xx status code
func (o *ArchiveSupplySourceBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this archive supply source bad request response has a 4xx status code
func (o *ArchiveSupplySourceBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this archive supply source bad request response has a 5xx status code
func (o *ArchiveSupplySourceBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this archive supply source bad request response a status code equal to that given
func (o *ArchiveSupplySourceBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *ArchiveSupplySourceBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /supplySources/2020-07-01/supplySources/{supplySourceId}][%d] archiveSupplySourceBadRequest  %+v", 400, o.Payload)
}

func (o *ArchiveSupplySourceBadRequest) String() string {
	return fmt.Sprintf("[DELETE /supplySources/2020-07-01/supplySources/{supplySourceId}][%d] archiveSupplySourceBadRequest  %+v", 400, o.Payload)
}

func (o *ArchiveSupplySourceBadRequest) GetPayload() *supply_sources_2020_07_01_models.ErrorList {
	return o.Payload
}

func (o *ArchiveSupplySourceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewArchiveSupplySourceForbidden creates a ArchiveSupplySourceForbidden with default headers values
func NewArchiveSupplySourceForbidden() *ArchiveSupplySourceForbidden {
	return &ArchiveSupplySourceForbidden{}
}

/*
ArchiveSupplySourceForbidden describes a response with status code 403, with default header values.

Indicates that access to the resource is forbidden. Possible reasons include Access Denied, Unauthorized, Expired Token, or Invalid Signature.
*/
type ArchiveSupplySourceForbidden struct {

	/* The unique request reference ID.
	 */
	XAmznRequestID string

	Payload *supply_sources_2020_07_01_models.ErrorList
}

// IsSuccess returns true when this archive supply source forbidden response has a 2xx status code
func (o *ArchiveSupplySourceForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this archive supply source forbidden response has a 3xx status code
func (o *ArchiveSupplySourceForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this archive supply source forbidden response has a 4xx status code
func (o *ArchiveSupplySourceForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this archive supply source forbidden response has a 5xx status code
func (o *ArchiveSupplySourceForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this archive supply source forbidden response a status code equal to that given
func (o *ArchiveSupplySourceForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ArchiveSupplySourceForbidden) Error() string {
	return fmt.Sprintf("[DELETE /supplySources/2020-07-01/supplySources/{supplySourceId}][%d] archiveSupplySourceForbidden  %+v", 403, o.Payload)
}

func (o *ArchiveSupplySourceForbidden) String() string {
	return fmt.Sprintf("[DELETE /supplySources/2020-07-01/supplySources/{supplySourceId}][%d] archiveSupplySourceForbidden  %+v", 403, o.Payload)
}

func (o *ArchiveSupplySourceForbidden) GetPayload() *supply_sources_2020_07_01_models.ErrorList {
	return o.Payload
}

func (o *ArchiveSupplySourceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewArchiveSupplySourceNotFound creates a ArchiveSupplySourceNotFound with default headers values
func NewArchiveSupplySourceNotFound() *ArchiveSupplySourceNotFound {
	return &ArchiveSupplySourceNotFound{}
}

/*
ArchiveSupplySourceNotFound describes a response with status code 404, with default header values.

The resource specified does not exist.
*/
type ArchiveSupplySourceNotFound struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* The unique request reference ID.
	 */
	XAmznRequestID string

	Payload *supply_sources_2020_07_01_models.ErrorList
}

// IsSuccess returns true when this archive supply source not found response has a 2xx status code
func (o *ArchiveSupplySourceNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this archive supply source not found response has a 3xx status code
func (o *ArchiveSupplySourceNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this archive supply source not found response has a 4xx status code
func (o *ArchiveSupplySourceNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this archive supply source not found response has a 5xx status code
func (o *ArchiveSupplySourceNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this archive supply source not found response a status code equal to that given
func (o *ArchiveSupplySourceNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *ArchiveSupplySourceNotFound) Error() string {
	return fmt.Sprintf("[DELETE /supplySources/2020-07-01/supplySources/{supplySourceId}][%d] archiveSupplySourceNotFound  %+v", 404, o.Payload)
}

func (o *ArchiveSupplySourceNotFound) String() string {
	return fmt.Sprintf("[DELETE /supplySources/2020-07-01/supplySources/{supplySourceId}][%d] archiveSupplySourceNotFound  %+v", 404, o.Payload)
}

func (o *ArchiveSupplySourceNotFound) GetPayload() *supply_sources_2020_07_01_models.ErrorList {
	return o.Payload
}

func (o *ArchiveSupplySourceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewArchiveSupplySourceRequestEntityTooLarge creates a ArchiveSupplySourceRequestEntityTooLarge with default headers values
func NewArchiveSupplySourceRequestEntityTooLarge() *ArchiveSupplySourceRequestEntityTooLarge {
	return &ArchiveSupplySourceRequestEntityTooLarge{}
}

/*
ArchiveSupplySourceRequestEntityTooLarge describes a response with status code 413, with default header values.

The request size exceeded the maximum accepted size.
*/
type ArchiveSupplySourceRequestEntityTooLarge struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* The unique request reference ID.
	 */
	XAmznRequestID string

	Payload *supply_sources_2020_07_01_models.ErrorList
}

// IsSuccess returns true when this archive supply source request entity too large response has a 2xx status code
func (o *ArchiveSupplySourceRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this archive supply source request entity too large response has a 3xx status code
func (o *ArchiveSupplySourceRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this archive supply source request entity too large response has a 4xx status code
func (o *ArchiveSupplySourceRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this archive supply source request entity too large response has a 5xx status code
func (o *ArchiveSupplySourceRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this archive supply source request entity too large response a status code equal to that given
func (o *ArchiveSupplySourceRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *ArchiveSupplySourceRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /supplySources/2020-07-01/supplySources/{supplySourceId}][%d] archiveSupplySourceRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *ArchiveSupplySourceRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[DELETE /supplySources/2020-07-01/supplySources/{supplySourceId}][%d] archiveSupplySourceRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *ArchiveSupplySourceRequestEntityTooLarge) GetPayload() *supply_sources_2020_07_01_models.ErrorList {
	return o.Payload
}

func (o *ArchiveSupplySourceRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewArchiveSupplySourceUnsupportedMediaType creates a ArchiveSupplySourceUnsupportedMediaType with default headers values
func NewArchiveSupplySourceUnsupportedMediaType() *ArchiveSupplySourceUnsupportedMediaType {
	return &ArchiveSupplySourceUnsupportedMediaType{}
}

/*
ArchiveSupplySourceUnsupportedMediaType describes a response with status code 415, with default header values.

The request payload is in an unsupported format.
*/
type ArchiveSupplySourceUnsupportedMediaType struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* The unique request reference ID.
	 */
	XAmznRequestID string

	Payload *supply_sources_2020_07_01_models.ErrorList
}

// IsSuccess returns true when this archive supply source unsupported media type response has a 2xx status code
func (o *ArchiveSupplySourceUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this archive supply source unsupported media type response has a 3xx status code
func (o *ArchiveSupplySourceUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this archive supply source unsupported media type response has a 4xx status code
func (o *ArchiveSupplySourceUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this archive supply source unsupported media type response has a 5xx status code
func (o *ArchiveSupplySourceUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this archive supply source unsupported media type response a status code equal to that given
func (o *ArchiveSupplySourceUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *ArchiveSupplySourceUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /supplySources/2020-07-01/supplySources/{supplySourceId}][%d] archiveSupplySourceUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *ArchiveSupplySourceUnsupportedMediaType) String() string {
	return fmt.Sprintf("[DELETE /supplySources/2020-07-01/supplySources/{supplySourceId}][%d] archiveSupplySourceUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *ArchiveSupplySourceUnsupportedMediaType) GetPayload() *supply_sources_2020_07_01_models.ErrorList {
	return o.Payload
}

func (o *ArchiveSupplySourceUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewArchiveSupplySourceTooManyRequests creates a ArchiveSupplySourceTooManyRequests with default headers values
func NewArchiveSupplySourceTooManyRequests() *ArchiveSupplySourceTooManyRequests {
	return &ArchiveSupplySourceTooManyRequests{}
}

/*
ArchiveSupplySourceTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type ArchiveSupplySourceTooManyRequests struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* The unique request reference ID.
	 */
	XAmznRequestID string

	Payload *supply_sources_2020_07_01_models.ErrorList
}

// IsSuccess returns true when this archive supply source too many requests response has a 2xx status code
func (o *ArchiveSupplySourceTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this archive supply source too many requests response has a 3xx status code
func (o *ArchiveSupplySourceTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this archive supply source too many requests response has a 4xx status code
func (o *ArchiveSupplySourceTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this archive supply source too many requests response has a 5xx status code
func (o *ArchiveSupplySourceTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this archive supply source too many requests response a status code equal to that given
func (o *ArchiveSupplySourceTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *ArchiveSupplySourceTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /supplySources/2020-07-01/supplySources/{supplySourceId}][%d] archiveSupplySourceTooManyRequests  %+v", 429, o.Payload)
}

func (o *ArchiveSupplySourceTooManyRequests) String() string {
	return fmt.Sprintf("[DELETE /supplySources/2020-07-01/supplySources/{supplySourceId}][%d] archiveSupplySourceTooManyRequests  %+v", 429, o.Payload)
}

func (o *ArchiveSupplySourceTooManyRequests) GetPayload() *supply_sources_2020_07_01_models.ErrorList {
	return o.Payload
}

func (o *ArchiveSupplySourceTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewArchiveSupplySourceInternalServerError creates a ArchiveSupplySourceInternalServerError with default headers values
func NewArchiveSupplySourceInternalServerError() *ArchiveSupplySourceInternalServerError {
	return &ArchiveSupplySourceInternalServerError{}
}

/*
ArchiveSupplySourceInternalServerError describes a response with status code 500, with default header values.

An unexpected condition occurred that prevented the server from fulfilling the request.
*/
type ArchiveSupplySourceInternalServerError struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* The unique request reference ID.
	 */
	XAmznRequestID string

	Payload *supply_sources_2020_07_01_models.ErrorList
}

// IsSuccess returns true when this archive supply source internal server error response has a 2xx status code
func (o *ArchiveSupplySourceInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this archive supply source internal server error response has a 3xx status code
func (o *ArchiveSupplySourceInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this archive supply source internal server error response has a 4xx status code
func (o *ArchiveSupplySourceInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this archive supply source internal server error response has a 5xx status code
func (o *ArchiveSupplySourceInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this archive supply source internal server error response a status code equal to that given
func (o *ArchiveSupplySourceInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ArchiveSupplySourceInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /supplySources/2020-07-01/supplySources/{supplySourceId}][%d] archiveSupplySourceInternalServerError  %+v", 500, o.Payload)
}

func (o *ArchiveSupplySourceInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /supplySources/2020-07-01/supplySources/{supplySourceId}][%d] archiveSupplySourceInternalServerError  %+v", 500, o.Payload)
}

func (o *ArchiveSupplySourceInternalServerError) GetPayload() *supply_sources_2020_07_01_models.ErrorList {
	return o.Payload
}

func (o *ArchiveSupplySourceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewArchiveSupplySourceServiceUnavailable creates a ArchiveSupplySourceServiceUnavailable with default headers values
func NewArchiveSupplySourceServiceUnavailable() *ArchiveSupplySourceServiceUnavailable {
	return &ArchiveSupplySourceServiceUnavailable{}
}

/*
ArchiveSupplySourceServiceUnavailable describes a response with status code 503, with default header values.

The temporary overloading or maintenance of the server.
*/
type ArchiveSupplySourceServiceUnavailable struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* The unique request reference ID.
	 */
	XAmznRequestID string

	Payload *supply_sources_2020_07_01_models.ErrorList
}

// IsSuccess returns true when this archive supply source service unavailable response has a 2xx status code
func (o *ArchiveSupplySourceServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this archive supply source service unavailable response has a 3xx status code
func (o *ArchiveSupplySourceServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this archive supply source service unavailable response has a 4xx status code
func (o *ArchiveSupplySourceServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this archive supply source service unavailable response has a 5xx status code
func (o *ArchiveSupplySourceServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this archive supply source service unavailable response a status code equal to that given
func (o *ArchiveSupplySourceServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *ArchiveSupplySourceServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /supplySources/2020-07-01/supplySources/{supplySourceId}][%d] archiveSupplySourceServiceUnavailable  %+v", 503, o.Payload)
}

func (o *ArchiveSupplySourceServiceUnavailable) String() string {
	return fmt.Sprintf("[DELETE /supplySources/2020-07-01/supplySources/{supplySourceId}][%d] archiveSupplySourceServiceUnavailable  %+v", 503, o.Payload)
}

func (o *ArchiveSupplySourceServiceUnavailable) GetPayload() *supply_sources_2020_07_01_models.ErrorList {
	return o.Payload
}

func (o *ArchiveSupplySourceServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
