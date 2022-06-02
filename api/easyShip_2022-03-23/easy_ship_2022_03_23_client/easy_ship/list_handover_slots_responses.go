// Code generated by go-swagger; DO NOT EDIT.

package easy_ship

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/xamandar/amzn-sp-api-go/api/easyShip_2022-03-23/easy_ship_2022_03_23_models"
)

// ListHandoverSlotsReader is a Reader for the ListHandoverSlots structure.
type ListHandoverSlotsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListHandoverSlotsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListHandoverSlotsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListHandoverSlotsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewListHandoverSlotsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListHandoverSlotsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListHandoverSlotsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewListHandoverSlotsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewListHandoverSlotsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListHandoverSlotsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewListHandoverSlotsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListHandoverSlotsOK creates a ListHandoverSlotsOK with default headers values
func NewListHandoverSlotsOK() *ListHandoverSlotsOK {
	return &ListHandoverSlotsOK{}
}

/* ListHandoverSlotsOK describes a response with status code 200, with default header values.

Success.
*/
type ListHandoverSlotsOK struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *easy_ship_2022_03_23_models.ListHandoverSlotsResponse
}

func (o *ListHandoverSlotsOK) Error() string {
	return fmt.Sprintf("[POST /easyShip/2022-03-23/timeSlot][%d] listHandoverSlotsOK  %+v", 200, o.Payload)
}
func (o *ListHandoverSlotsOK) GetPayload() *easy_ship_2022_03_23_models.ListHandoverSlotsResponse {
	return o.Payload
}

func (o *ListHandoverSlotsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(easy_ship_2022_03_23_models.ListHandoverSlotsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListHandoverSlotsBadRequest creates a ListHandoverSlotsBadRequest with default headers values
func NewListHandoverSlotsBadRequest() *ListHandoverSlotsBadRequest {
	return &ListHandoverSlotsBadRequest{}
}

/* ListHandoverSlotsBadRequest describes a response with status code 400, with default header values.

Request has missing or invalid parameters and cannot be parsed.
*/
type ListHandoverSlotsBadRequest struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *easy_ship_2022_03_23_models.ErrorList
}

func (o *ListHandoverSlotsBadRequest) Error() string {
	return fmt.Sprintf("[POST /easyShip/2022-03-23/timeSlot][%d] listHandoverSlotsBadRequest  %+v", 400, o.Payload)
}
func (o *ListHandoverSlotsBadRequest) GetPayload() *easy_ship_2022_03_23_models.ErrorList {
	return o.Payload
}

func (o *ListHandoverSlotsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(easy_ship_2022_03_23_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListHandoverSlotsUnauthorized creates a ListHandoverSlotsUnauthorized with default headers values
func NewListHandoverSlotsUnauthorized() *ListHandoverSlotsUnauthorized {
	return &ListHandoverSlotsUnauthorized{}
}

/* ListHandoverSlotsUnauthorized describes a response with status code 401, with default header values.

The request's Authorization header is not formatted correctly or does not contain a valid token.
*/
type ListHandoverSlotsUnauthorized struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *easy_ship_2022_03_23_models.ErrorList
}

func (o *ListHandoverSlotsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /easyShip/2022-03-23/timeSlot][%d] listHandoverSlotsUnauthorized  %+v", 401, o.Payload)
}
func (o *ListHandoverSlotsUnauthorized) GetPayload() *easy_ship_2022_03_23_models.ErrorList {
	return o.Payload
}

func (o *ListHandoverSlotsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(easy_ship_2022_03_23_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListHandoverSlotsForbidden creates a ListHandoverSlotsForbidden with default headers values
func NewListHandoverSlotsForbidden() *ListHandoverSlotsForbidden {
	return &ListHandoverSlotsForbidden{}
}

/* ListHandoverSlotsForbidden describes a response with status code 403, with default header values.

Indicates access to the resource is forbidden. Possible reasons include Access Denied, Unauthorized, Expired Token, or Invalid Signature.
*/
type ListHandoverSlotsForbidden struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *easy_ship_2022_03_23_models.ErrorList
}

func (o *ListHandoverSlotsForbidden) Error() string {
	return fmt.Sprintf("[POST /easyShip/2022-03-23/timeSlot][%d] listHandoverSlotsForbidden  %+v", 403, o.Payload)
}
func (o *ListHandoverSlotsForbidden) GetPayload() *easy_ship_2022_03_23_models.ErrorList {
	return o.Payload
}

func (o *ListHandoverSlotsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(easy_ship_2022_03_23_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListHandoverSlotsNotFound creates a ListHandoverSlotsNotFound with default headers values
func NewListHandoverSlotsNotFound() *ListHandoverSlotsNotFound {
	return &ListHandoverSlotsNotFound{}
}

/* ListHandoverSlotsNotFound describes a response with status code 404, with default header values.

The specified resource does not exist.
*/
type ListHandoverSlotsNotFound struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *easy_ship_2022_03_23_models.ErrorList
}

func (o *ListHandoverSlotsNotFound) Error() string {
	return fmt.Sprintf("[POST /easyShip/2022-03-23/timeSlot][%d] listHandoverSlotsNotFound  %+v", 404, o.Payload)
}
func (o *ListHandoverSlotsNotFound) GetPayload() *easy_ship_2022_03_23_models.ErrorList {
	return o.Payload
}

func (o *ListHandoverSlotsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(easy_ship_2022_03_23_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListHandoverSlotsUnsupportedMediaType creates a ListHandoverSlotsUnsupportedMediaType with default headers values
func NewListHandoverSlotsUnsupportedMediaType() *ListHandoverSlotsUnsupportedMediaType {
	return &ListHandoverSlotsUnsupportedMediaType{}
}

/* ListHandoverSlotsUnsupportedMediaType describes a response with status code 415, with default header values.

The request payload is in an unsupported format.
*/
type ListHandoverSlotsUnsupportedMediaType struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *easy_ship_2022_03_23_models.ErrorList
}

func (o *ListHandoverSlotsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /easyShip/2022-03-23/timeSlot][%d] listHandoverSlotsUnsupportedMediaType  %+v", 415, o.Payload)
}
func (o *ListHandoverSlotsUnsupportedMediaType) GetPayload() *easy_ship_2022_03_23_models.ErrorList {
	return o.Payload
}

func (o *ListHandoverSlotsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(easy_ship_2022_03_23_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListHandoverSlotsTooManyRequests creates a ListHandoverSlotsTooManyRequests with default headers values
func NewListHandoverSlotsTooManyRequests() *ListHandoverSlotsTooManyRequests {
	return &ListHandoverSlotsTooManyRequests{}
}

/* ListHandoverSlotsTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type ListHandoverSlotsTooManyRequests struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *easy_ship_2022_03_23_models.ErrorList
}

func (o *ListHandoverSlotsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /easyShip/2022-03-23/timeSlot][%d] listHandoverSlotsTooManyRequests  %+v", 429, o.Payload)
}
func (o *ListHandoverSlotsTooManyRequests) GetPayload() *easy_ship_2022_03_23_models.ErrorList {
	return o.Payload
}

func (o *ListHandoverSlotsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(easy_ship_2022_03_23_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListHandoverSlotsInternalServerError creates a ListHandoverSlotsInternalServerError with default headers values
func NewListHandoverSlotsInternalServerError() *ListHandoverSlotsInternalServerError {
	return &ListHandoverSlotsInternalServerError{}
}

/* ListHandoverSlotsInternalServerError describes a response with status code 500, with default header values.

An unexpected condition occurred that prevented the server from fulfilling the request.
*/
type ListHandoverSlotsInternalServerError struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *easy_ship_2022_03_23_models.ErrorList
}

func (o *ListHandoverSlotsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /easyShip/2022-03-23/timeSlot][%d] listHandoverSlotsInternalServerError  %+v", 500, o.Payload)
}
func (o *ListHandoverSlotsInternalServerError) GetPayload() *easy_ship_2022_03_23_models.ErrorList {
	return o.Payload
}

func (o *ListHandoverSlotsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(easy_ship_2022_03_23_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListHandoverSlotsServiceUnavailable creates a ListHandoverSlotsServiceUnavailable with default headers values
func NewListHandoverSlotsServiceUnavailable() *ListHandoverSlotsServiceUnavailable {
	return &ListHandoverSlotsServiceUnavailable{}
}

/* ListHandoverSlotsServiceUnavailable describes a response with status code 503, with default header values.

Temporary overloading or maintenance of the server.
*/
type ListHandoverSlotsServiceUnavailable struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *easy_ship_2022_03_23_models.ErrorList
}

func (o *ListHandoverSlotsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /easyShip/2022-03-23/timeSlot][%d] listHandoverSlotsServiceUnavailable  %+v", 503, o.Payload)
}
func (o *ListHandoverSlotsServiceUnavailable) GetPayload() *easy_ship_2022_03_23_models.ErrorList {
	return o.Payload
}

func (o *ListHandoverSlotsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(easy_ship_2022_03_23_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
