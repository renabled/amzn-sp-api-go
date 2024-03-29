// Code generated by go-swagger; DO NOT EDIT.

package listings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/renabled/amzn-sp-api-go/api/listingsRestrictions_2021-08-01/listings_restrictions_2021_08_01_models"
)

// GetListingsRestrictionsReader is a Reader for the GetListingsRestrictions structure.
type GetListingsRestrictionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetListingsRestrictionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetListingsRestrictionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetListingsRestrictionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetListingsRestrictionsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetListingsRestrictionsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetListingsRestrictionsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetListingsRestrictionsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetListingsRestrictionsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetListingsRestrictionsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetListingsRestrictionsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetListingsRestrictionsOK creates a GetListingsRestrictionsOK with default headers values
func NewGetListingsRestrictionsOK() *GetListingsRestrictionsOK {
	return &GetListingsRestrictionsOK{}
}

/*
GetListingsRestrictionsOK describes a response with status code 200, with default header values.

Successfully retrieved the listings restrictions.
*/
type GetListingsRestrictionsOK struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *listings_restrictions_2021_08_01_models.RestrictionList
}

// IsSuccess returns true when this get listings restrictions o k response has a 2xx status code
func (o *GetListingsRestrictionsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get listings restrictions o k response has a 3xx status code
func (o *GetListingsRestrictionsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get listings restrictions o k response has a 4xx status code
func (o *GetListingsRestrictionsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get listings restrictions o k response has a 5xx status code
func (o *GetListingsRestrictionsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get listings restrictions o k response a status code equal to that given
func (o *GetListingsRestrictionsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetListingsRestrictionsOK) Error() string {
	return fmt.Sprintf("[GET /listings/2021-08-01/restrictions][%d] getListingsRestrictionsOK  %+v", 200, o.Payload)
}

func (o *GetListingsRestrictionsOK) String() string {
	return fmt.Sprintf("[GET /listings/2021-08-01/restrictions][%d] getListingsRestrictionsOK  %+v", 200, o.Payload)
}

func (o *GetListingsRestrictionsOK) GetPayload() *listings_restrictions_2021_08_01_models.RestrictionList {
	return o.Payload
}

func (o *GetListingsRestrictionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(listings_restrictions_2021_08_01_models.RestrictionList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetListingsRestrictionsBadRequest creates a GetListingsRestrictionsBadRequest with default headers values
func NewGetListingsRestrictionsBadRequest() *GetListingsRestrictionsBadRequest {
	return &GetListingsRestrictionsBadRequest{}
}

/*
GetListingsRestrictionsBadRequest describes a response with status code 400, with default header values.

Request has missing or invalid parameters and cannot be parsed.
*/
type GetListingsRestrictionsBadRequest struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload listings_restrictions_2021_08_01_models.ErrorList
}

// IsSuccess returns true when this get listings restrictions bad request response has a 2xx status code
func (o *GetListingsRestrictionsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get listings restrictions bad request response has a 3xx status code
func (o *GetListingsRestrictionsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get listings restrictions bad request response has a 4xx status code
func (o *GetListingsRestrictionsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get listings restrictions bad request response has a 5xx status code
func (o *GetListingsRestrictionsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get listings restrictions bad request response a status code equal to that given
func (o *GetListingsRestrictionsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetListingsRestrictionsBadRequest) Error() string {
	return fmt.Sprintf("[GET /listings/2021-08-01/restrictions][%d] getListingsRestrictionsBadRequest  %+v", 400, o.Payload)
}

func (o *GetListingsRestrictionsBadRequest) String() string {
	return fmt.Sprintf("[GET /listings/2021-08-01/restrictions][%d] getListingsRestrictionsBadRequest  %+v", 400, o.Payload)
}

func (o *GetListingsRestrictionsBadRequest) GetPayload() listings_restrictions_2021_08_01_models.ErrorList {
	return o.Payload
}

func (o *GetListingsRestrictionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetListingsRestrictionsForbidden creates a GetListingsRestrictionsForbidden with default headers values
func NewGetListingsRestrictionsForbidden() *GetListingsRestrictionsForbidden {
	return &GetListingsRestrictionsForbidden{}
}

/*
GetListingsRestrictionsForbidden describes a response with status code 403, with default header values.

Indicates that access to the resource is forbidden. Possible reasons include Access Denied, Unauthorized, Expired Token, or Invalid Signature.
*/
type GetListingsRestrictionsForbidden struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload listings_restrictions_2021_08_01_models.ErrorList
}

// IsSuccess returns true when this get listings restrictions forbidden response has a 2xx status code
func (o *GetListingsRestrictionsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get listings restrictions forbidden response has a 3xx status code
func (o *GetListingsRestrictionsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get listings restrictions forbidden response has a 4xx status code
func (o *GetListingsRestrictionsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get listings restrictions forbidden response has a 5xx status code
func (o *GetListingsRestrictionsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get listings restrictions forbidden response a status code equal to that given
func (o *GetListingsRestrictionsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetListingsRestrictionsForbidden) Error() string {
	return fmt.Sprintf("[GET /listings/2021-08-01/restrictions][%d] getListingsRestrictionsForbidden  %+v", 403, o.Payload)
}

func (o *GetListingsRestrictionsForbidden) String() string {
	return fmt.Sprintf("[GET /listings/2021-08-01/restrictions][%d] getListingsRestrictionsForbidden  %+v", 403, o.Payload)
}

func (o *GetListingsRestrictionsForbidden) GetPayload() listings_restrictions_2021_08_01_models.ErrorList {
	return o.Payload
}

func (o *GetListingsRestrictionsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetListingsRestrictionsNotFound creates a GetListingsRestrictionsNotFound with default headers values
func NewGetListingsRestrictionsNotFound() *GetListingsRestrictionsNotFound {
	return &GetListingsRestrictionsNotFound{}
}

/*
GetListingsRestrictionsNotFound describes a response with status code 404, with default header values.

The resource specified does not exist.
*/
type GetListingsRestrictionsNotFound struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload listings_restrictions_2021_08_01_models.ErrorList
}

// IsSuccess returns true when this get listings restrictions not found response has a 2xx status code
func (o *GetListingsRestrictionsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get listings restrictions not found response has a 3xx status code
func (o *GetListingsRestrictionsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get listings restrictions not found response has a 4xx status code
func (o *GetListingsRestrictionsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get listings restrictions not found response has a 5xx status code
func (o *GetListingsRestrictionsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get listings restrictions not found response a status code equal to that given
func (o *GetListingsRestrictionsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetListingsRestrictionsNotFound) Error() string {
	return fmt.Sprintf("[GET /listings/2021-08-01/restrictions][%d] getListingsRestrictionsNotFound  %+v", 404, o.Payload)
}

func (o *GetListingsRestrictionsNotFound) String() string {
	return fmt.Sprintf("[GET /listings/2021-08-01/restrictions][%d] getListingsRestrictionsNotFound  %+v", 404, o.Payload)
}

func (o *GetListingsRestrictionsNotFound) GetPayload() listings_restrictions_2021_08_01_models.ErrorList {
	return o.Payload
}

func (o *GetListingsRestrictionsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetListingsRestrictionsRequestEntityTooLarge creates a GetListingsRestrictionsRequestEntityTooLarge with default headers values
func NewGetListingsRestrictionsRequestEntityTooLarge() *GetListingsRestrictionsRequestEntityTooLarge {
	return &GetListingsRestrictionsRequestEntityTooLarge{}
}

/*
GetListingsRestrictionsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request size exceeded the maximum accepted size.
*/
type GetListingsRestrictionsRequestEntityTooLarge struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload listings_restrictions_2021_08_01_models.ErrorList
}

// IsSuccess returns true when this get listings restrictions request entity too large response has a 2xx status code
func (o *GetListingsRestrictionsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get listings restrictions request entity too large response has a 3xx status code
func (o *GetListingsRestrictionsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get listings restrictions request entity too large response has a 4xx status code
func (o *GetListingsRestrictionsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get listings restrictions request entity too large response has a 5xx status code
func (o *GetListingsRestrictionsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get listings restrictions request entity too large response a status code equal to that given
func (o *GetListingsRestrictionsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetListingsRestrictionsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /listings/2021-08-01/restrictions][%d] getListingsRestrictionsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetListingsRestrictionsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /listings/2021-08-01/restrictions][%d] getListingsRestrictionsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetListingsRestrictionsRequestEntityTooLarge) GetPayload() listings_restrictions_2021_08_01_models.ErrorList {
	return o.Payload
}

func (o *GetListingsRestrictionsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetListingsRestrictionsUnsupportedMediaType creates a GetListingsRestrictionsUnsupportedMediaType with default headers values
func NewGetListingsRestrictionsUnsupportedMediaType() *GetListingsRestrictionsUnsupportedMediaType {
	return &GetListingsRestrictionsUnsupportedMediaType{}
}

/*
GetListingsRestrictionsUnsupportedMediaType describes a response with status code 415, with default header values.

The request payload is in an unsupported format.
*/
type GetListingsRestrictionsUnsupportedMediaType struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload listings_restrictions_2021_08_01_models.ErrorList
}

// IsSuccess returns true when this get listings restrictions unsupported media type response has a 2xx status code
func (o *GetListingsRestrictionsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get listings restrictions unsupported media type response has a 3xx status code
func (o *GetListingsRestrictionsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get listings restrictions unsupported media type response has a 4xx status code
func (o *GetListingsRestrictionsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get listings restrictions unsupported media type response has a 5xx status code
func (o *GetListingsRestrictionsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get listings restrictions unsupported media type response a status code equal to that given
func (o *GetListingsRestrictionsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetListingsRestrictionsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /listings/2021-08-01/restrictions][%d] getListingsRestrictionsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetListingsRestrictionsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /listings/2021-08-01/restrictions][%d] getListingsRestrictionsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetListingsRestrictionsUnsupportedMediaType) GetPayload() listings_restrictions_2021_08_01_models.ErrorList {
	return o.Payload
}

func (o *GetListingsRestrictionsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetListingsRestrictionsTooManyRequests creates a GetListingsRestrictionsTooManyRequests with default headers values
func NewGetListingsRestrictionsTooManyRequests() *GetListingsRestrictionsTooManyRequests {
	return &GetListingsRestrictionsTooManyRequests{}
}

/*
GetListingsRestrictionsTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type GetListingsRestrictionsTooManyRequests struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload listings_restrictions_2021_08_01_models.ErrorList
}

// IsSuccess returns true when this get listings restrictions too many requests response has a 2xx status code
func (o *GetListingsRestrictionsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get listings restrictions too many requests response has a 3xx status code
func (o *GetListingsRestrictionsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get listings restrictions too many requests response has a 4xx status code
func (o *GetListingsRestrictionsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get listings restrictions too many requests response has a 5xx status code
func (o *GetListingsRestrictionsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get listings restrictions too many requests response a status code equal to that given
func (o *GetListingsRestrictionsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetListingsRestrictionsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /listings/2021-08-01/restrictions][%d] getListingsRestrictionsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetListingsRestrictionsTooManyRequests) String() string {
	return fmt.Sprintf("[GET /listings/2021-08-01/restrictions][%d] getListingsRestrictionsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetListingsRestrictionsTooManyRequests) GetPayload() listings_restrictions_2021_08_01_models.ErrorList {
	return o.Payload
}

func (o *GetListingsRestrictionsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetListingsRestrictionsInternalServerError creates a GetListingsRestrictionsInternalServerError with default headers values
func NewGetListingsRestrictionsInternalServerError() *GetListingsRestrictionsInternalServerError {
	return &GetListingsRestrictionsInternalServerError{}
}

/*
GetListingsRestrictionsInternalServerError describes a response with status code 500, with default header values.

An unexpected condition occurred that prevented the server from fulfilling the request.
*/
type GetListingsRestrictionsInternalServerError struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload listings_restrictions_2021_08_01_models.ErrorList
}

// IsSuccess returns true when this get listings restrictions internal server error response has a 2xx status code
func (o *GetListingsRestrictionsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get listings restrictions internal server error response has a 3xx status code
func (o *GetListingsRestrictionsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get listings restrictions internal server error response has a 4xx status code
func (o *GetListingsRestrictionsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get listings restrictions internal server error response has a 5xx status code
func (o *GetListingsRestrictionsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get listings restrictions internal server error response a status code equal to that given
func (o *GetListingsRestrictionsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetListingsRestrictionsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /listings/2021-08-01/restrictions][%d] getListingsRestrictionsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetListingsRestrictionsInternalServerError) String() string {
	return fmt.Sprintf("[GET /listings/2021-08-01/restrictions][%d] getListingsRestrictionsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetListingsRestrictionsInternalServerError) GetPayload() listings_restrictions_2021_08_01_models.ErrorList {
	return o.Payload
}

func (o *GetListingsRestrictionsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetListingsRestrictionsServiceUnavailable creates a GetListingsRestrictionsServiceUnavailable with default headers values
func NewGetListingsRestrictionsServiceUnavailable() *GetListingsRestrictionsServiceUnavailable {
	return &GetListingsRestrictionsServiceUnavailable{}
}

/*
GetListingsRestrictionsServiceUnavailable describes a response with status code 503, with default header values.

Temporary overloading or maintenance of the server.
*/
type GetListingsRestrictionsServiceUnavailable struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload listings_restrictions_2021_08_01_models.ErrorList
}

// IsSuccess returns true when this get listings restrictions service unavailable response has a 2xx status code
func (o *GetListingsRestrictionsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get listings restrictions service unavailable response has a 3xx status code
func (o *GetListingsRestrictionsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get listings restrictions service unavailable response has a 4xx status code
func (o *GetListingsRestrictionsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get listings restrictions service unavailable response has a 5xx status code
func (o *GetListingsRestrictionsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get listings restrictions service unavailable response a status code equal to that given
func (o *GetListingsRestrictionsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetListingsRestrictionsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /listings/2021-08-01/restrictions][%d] getListingsRestrictionsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetListingsRestrictionsServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /listings/2021-08-01/restrictions][%d] getListingsRestrictionsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetListingsRestrictionsServiceUnavailable) GetPayload() listings_restrictions_2021_08_01_models.ErrorList {
	return o.Payload
}

func (o *GetListingsRestrictionsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
