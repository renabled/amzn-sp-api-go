// Code generated by go-swagger; DO NOT EDIT.

package aplus_content

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/xamandar/amzn-sp-api-go/api/aplusContent_2020-11-01/aplus_content_2020_11_01_models"
)

// ListContentDocumentAsinRelationsReader is a Reader for the ListContentDocumentAsinRelations structure.
type ListContentDocumentAsinRelationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListContentDocumentAsinRelationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListContentDocumentAsinRelationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListContentDocumentAsinRelationsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewListContentDocumentAsinRelationsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListContentDocumentAsinRelationsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListContentDocumentAsinRelationsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewListContentDocumentAsinRelationsGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewListContentDocumentAsinRelationsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListContentDocumentAsinRelationsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewListContentDocumentAsinRelationsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListContentDocumentAsinRelationsOK creates a ListContentDocumentAsinRelationsOK with default headers values
func NewListContentDocumentAsinRelationsOK() *ListContentDocumentAsinRelationsOK {
	return &ListContentDocumentAsinRelationsOK{}
}

/* ListContentDocumentAsinRelationsOK describes a response with status code 200, with default header values.

Success.
*/
type ListContentDocumentAsinRelationsOK struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *aplus_content_2020_11_01_models.ListContentDocumentAsinRelationsResponse
}

func (o *ListContentDocumentAsinRelationsOK) Error() string {
	return fmt.Sprintf("[GET /aplus/2020-11-01/contentDocuments/{contentReferenceKey}/asins][%d] listContentDocumentAsinRelationsOK  %+v", 200, o.Payload)
}
func (o *ListContentDocumentAsinRelationsOK) GetPayload() *aplus_content_2020_11_01_models.ListContentDocumentAsinRelationsResponse {
	return o.Payload
}

func (o *ListContentDocumentAsinRelationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(aplus_content_2020_11_01_models.ListContentDocumentAsinRelationsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListContentDocumentAsinRelationsBadRequest creates a ListContentDocumentAsinRelationsBadRequest with default headers values
func NewListContentDocumentAsinRelationsBadRequest() *ListContentDocumentAsinRelationsBadRequest {
	return &ListContentDocumentAsinRelationsBadRequest{}
}

/* ListContentDocumentAsinRelationsBadRequest describes a response with status code 400, with default header values.

Request has missing or invalid parameters and cannot be parsed.
*/
type ListContentDocumentAsinRelationsBadRequest struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *aplus_content_2020_11_01_models.ErrorList
}

func (o *ListContentDocumentAsinRelationsBadRequest) Error() string {
	return fmt.Sprintf("[GET /aplus/2020-11-01/contentDocuments/{contentReferenceKey}/asins][%d] listContentDocumentAsinRelationsBadRequest  %+v", 400, o.Payload)
}
func (o *ListContentDocumentAsinRelationsBadRequest) GetPayload() *aplus_content_2020_11_01_models.ErrorList {
	return o.Payload
}

func (o *ListContentDocumentAsinRelationsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(aplus_content_2020_11_01_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListContentDocumentAsinRelationsUnauthorized creates a ListContentDocumentAsinRelationsUnauthorized with default headers values
func NewListContentDocumentAsinRelationsUnauthorized() *ListContentDocumentAsinRelationsUnauthorized {
	return &ListContentDocumentAsinRelationsUnauthorized{}
}

/* ListContentDocumentAsinRelationsUnauthorized describes a response with status code 401, with default header values.

The request's Authorization header is not formatted correctly or does not contain a valid token.
*/
type ListContentDocumentAsinRelationsUnauthorized struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *aplus_content_2020_11_01_models.ErrorList
}

func (o *ListContentDocumentAsinRelationsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /aplus/2020-11-01/contentDocuments/{contentReferenceKey}/asins][%d] listContentDocumentAsinRelationsUnauthorized  %+v", 401, o.Payload)
}
func (o *ListContentDocumentAsinRelationsUnauthorized) GetPayload() *aplus_content_2020_11_01_models.ErrorList {
	return o.Payload
}

func (o *ListContentDocumentAsinRelationsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(aplus_content_2020_11_01_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListContentDocumentAsinRelationsForbidden creates a ListContentDocumentAsinRelationsForbidden with default headers values
func NewListContentDocumentAsinRelationsForbidden() *ListContentDocumentAsinRelationsForbidden {
	return &ListContentDocumentAsinRelationsForbidden{}
}

/* ListContentDocumentAsinRelationsForbidden describes a response with status code 403, with default header values.

Indicates that access to the resource is forbidden. Possible reasons include Access Denied, Unauthorized, Expired Token, or Invalid Signature.
*/
type ListContentDocumentAsinRelationsForbidden struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *aplus_content_2020_11_01_models.ErrorList
}

func (o *ListContentDocumentAsinRelationsForbidden) Error() string {
	return fmt.Sprintf("[GET /aplus/2020-11-01/contentDocuments/{contentReferenceKey}/asins][%d] listContentDocumentAsinRelationsForbidden  %+v", 403, o.Payload)
}
func (o *ListContentDocumentAsinRelationsForbidden) GetPayload() *aplus_content_2020_11_01_models.ErrorList {
	return o.Payload
}

func (o *ListContentDocumentAsinRelationsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(aplus_content_2020_11_01_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListContentDocumentAsinRelationsNotFound creates a ListContentDocumentAsinRelationsNotFound with default headers values
func NewListContentDocumentAsinRelationsNotFound() *ListContentDocumentAsinRelationsNotFound {
	return &ListContentDocumentAsinRelationsNotFound{}
}

/* ListContentDocumentAsinRelationsNotFound describes a response with status code 404, with default header values.

The specified resource does not exist.
*/
type ListContentDocumentAsinRelationsNotFound struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *aplus_content_2020_11_01_models.ErrorList
}

func (o *ListContentDocumentAsinRelationsNotFound) Error() string {
	return fmt.Sprintf("[GET /aplus/2020-11-01/contentDocuments/{contentReferenceKey}/asins][%d] listContentDocumentAsinRelationsNotFound  %+v", 404, o.Payload)
}
func (o *ListContentDocumentAsinRelationsNotFound) GetPayload() *aplus_content_2020_11_01_models.ErrorList {
	return o.Payload
}

func (o *ListContentDocumentAsinRelationsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(aplus_content_2020_11_01_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListContentDocumentAsinRelationsGone creates a ListContentDocumentAsinRelationsGone with default headers values
func NewListContentDocumentAsinRelationsGone() *ListContentDocumentAsinRelationsGone {
	return &ListContentDocumentAsinRelationsGone{}
}

/* ListContentDocumentAsinRelationsGone describes a response with status code 410, with default header values.

The specified resource no longer exists.
*/
type ListContentDocumentAsinRelationsGone struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *aplus_content_2020_11_01_models.ErrorList
}

func (o *ListContentDocumentAsinRelationsGone) Error() string {
	return fmt.Sprintf("[GET /aplus/2020-11-01/contentDocuments/{contentReferenceKey}/asins][%d] listContentDocumentAsinRelationsGone  %+v", 410, o.Payload)
}
func (o *ListContentDocumentAsinRelationsGone) GetPayload() *aplus_content_2020_11_01_models.ErrorList {
	return o.Payload
}

func (o *ListContentDocumentAsinRelationsGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(aplus_content_2020_11_01_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListContentDocumentAsinRelationsTooManyRequests creates a ListContentDocumentAsinRelationsTooManyRequests with default headers values
func NewListContentDocumentAsinRelationsTooManyRequests() *ListContentDocumentAsinRelationsTooManyRequests {
	return &ListContentDocumentAsinRelationsTooManyRequests{}
}

/* ListContentDocumentAsinRelationsTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type ListContentDocumentAsinRelationsTooManyRequests struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *aplus_content_2020_11_01_models.ErrorList
}

func (o *ListContentDocumentAsinRelationsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /aplus/2020-11-01/contentDocuments/{contentReferenceKey}/asins][%d] listContentDocumentAsinRelationsTooManyRequests  %+v", 429, o.Payload)
}
func (o *ListContentDocumentAsinRelationsTooManyRequests) GetPayload() *aplus_content_2020_11_01_models.ErrorList {
	return o.Payload
}

func (o *ListContentDocumentAsinRelationsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(aplus_content_2020_11_01_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListContentDocumentAsinRelationsInternalServerError creates a ListContentDocumentAsinRelationsInternalServerError with default headers values
func NewListContentDocumentAsinRelationsInternalServerError() *ListContentDocumentAsinRelationsInternalServerError {
	return &ListContentDocumentAsinRelationsInternalServerError{}
}

/* ListContentDocumentAsinRelationsInternalServerError describes a response with status code 500, with default header values.

An unexpected condition occurred that prevented the server from fulfilling the request.
*/
type ListContentDocumentAsinRelationsInternalServerError struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *aplus_content_2020_11_01_models.ErrorList
}

func (o *ListContentDocumentAsinRelationsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /aplus/2020-11-01/contentDocuments/{contentReferenceKey}/asins][%d] listContentDocumentAsinRelationsInternalServerError  %+v", 500, o.Payload)
}
func (o *ListContentDocumentAsinRelationsInternalServerError) GetPayload() *aplus_content_2020_11_01_models.ErrorList {
	return o.Payload
}

func (o *ListContentDocumentAsinRelationsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(aplus_content_2020_11_01_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListContentDocumentAsinRelationsServiceUnavailable creates a ListContentDocumentAsinRelationsServiceUnavailable with default headers values
func NewListContentDocumentAsinRelationsServiceUnavailable() *ListContentDocumentAsinRelationsServiceUnavailable {
	return &ListContentDocumentAsinRelationsServiceUnavailable{}
}

/* ListContentDocumentAsinRelationsServiceUnavailable describes a response with status code 503, with default header values.

Temporary overloading or maintenance of the server.
*/
type ListContentDocumentAsinRelationsServiceUnavailable struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *aplus_content_2020_11_01_models.ErrorList
}

func (o *ListContentDocumentAsinRelationsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /aplus/2020-11-01/contentDocuments/{contentReferenceKey}/asins][%d] listContentDocumentAsinRelationsServiceUnavailable  %+v", 503, o.Payload)
}
func (o *ListContentDocumentAsinRelationsServiceUnavailable) GetPayload() *aplus_content_2020_11_01_models.ErrorList {
	return o.Payload
}

func (o *ListContentDocumentAsinRelationsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(aplus_content_2020_11_01_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
