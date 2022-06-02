// Code generated by go-swagger; DO NOT EDIT.

package service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/xamandar/amzn-sp-api-go/api/services/services_models"
)

// CompleteServiceJobByServiceJobIDReader is a Reader for the CompleteServiceJobByServiceJobID structure.
type CompleteServiceJobByServiceJobIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CompleteServiceJobByServiceJobIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCompleteServiceJobByServiceJobIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCompleteServiceJobByServiceJobIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCompleteServiceJobByServiceJobIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCompleteServiceJobByServiceJobIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewCompleteServiceJobByServiceJobIDRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewCompleteServiceJobByServiceJobIDUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewCompleteServiceJobByServiceJobIDUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewCompleteServiceJobByServiceJobIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCompleteServiceJobByServiceJobIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewCompleteServiceJobByServiceJobIDServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCompleteServiceJobByServiceJobIDOK creates a CompleteServiceJobByServiceJobIDOK with default headers values
func NewCompleteServiceJobByServiceJobIDOK() *CompleteServiceJobByServiceJobIDOK {
	return &CompleteServiceJobByServiceJobIDOK{}
}

/* CompleteServiceJobByServiceJobIDOK describes a response with status code 200, with default header values.

Success response
*/
type CompleteServiceJobByServiceJobIDOK struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *services_models.CompleteServiceJobByServiceJobIDResponse
}

func (o *CompleteServiceJobByServiceJobIDOK) Error() string {
	return fmt.Sprintf("[PUT /service/v1/serviceJobs/{serviceJobId}/completions][%d] completeServiceJobByServiceJobIdOK  %+v", 200, o.Payload)
}
func (o *CompleteServiceJobByServiceJobIDOK) GetPayload() *services_models.CompleteServiceJobByServiceJobIDResponse {
	return o.Payload
}

func (o *CompleteServiceJobByServiceJobIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(services_models.CompleteServiceJobByServiceJobIDResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCompleteServiceJobByServiceJobIDBadRequest creates a CompleteServiceJobByServiceJobIDBadRequest with default headers values
func NewCompleteServiceJobByServiceJobIDBadRequest() *CompleteServiceJobByServiceJobIDBadRequest {
	return &CompleteServiceJobByServiceJobIDBadRequest{}
}

/* CompleteServiceJobByServiceJobIDBadRequest describes a response with status code 400, with default header values.

Request has missing or invalid parameters and cannot be parsed.
*/
type CompleteServiceJobByServiceJobIDBadRequest struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *services_models.CompleteServiceJobByServiceJobIDResponse
}

func (o *CompleteServiceJobByServiceJobIDBadRequest) Error() string {
	return fmt.Sprintf("[PUT /service/v1/serviceJobs/{serviceJobId}/completions][%d] completeServiceJobByServiceJobIdBadRequest  %+v", 400, o.Payload)
}
func (o *CompleteServiceJobByServiceJobIDBadRequest) GetPayload() *services_models.CompleteServiceJobByServiceJobIDResponse {
	return o.Payload
}

func (o *CompleteServiceJobByServiceJobIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(services_models.CompleteServiceJobByServiceJobIDResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCompleteServiceJobByServiceJobIDForbidden creates a CompleteServiceJobByServiceJobIDForbidden with default headers values
func NewCompleteServiceJobByServiceJobIDForbidden() *CompleteServiceJobByServiceJobIDForbidden {
	return &CompleteServiceJobByServiceJobIDForbidden{}
}

/* CompleteServiceJobByServiceJobIDForbidden describes a response with status code 403, with default header values.

Indicates that access to the resource is forbidden. Possible reasons include Access Denied, Unauthorized, Expired Token, or Invalid Signature.
*/
type CompleteServiceJobByServiceJobIDForbidden struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *services_models.CompleteServiceJobByServiceJobIDResponse
}

func (o *CompleteServiceJobByServiceJobIDForbidden) Error() string {
	return fmt.Sprintf("[PUT /service/v1/serviceJobs/{serviceJobId}/completions][%d] completeServiceJobByServiceJobIdForbidden  %+v", 403, o.Payload)
}
func (o *CompleteServiceJobByServiceJobIDForbidden) GetPayload() *services_models.CompleteServiceJobByServiceJobIDResponse {
	return o.Payload
}

func (o *CompleteServiceJobByServiceJobIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(services_models.CompleteServiceJobByServiceJobIDResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCompleteServiceJobByServiceJobIDNotFound creates a CompleteServiceJobByServiceJobIDNotFound with default headers values
func NewCompleteServiceJobByServiceJobIDNotFound() *CompleteServiceJobByServiceJobIDNotFound {
	return &CompleteServiceJobByServiceJobIDNotFound{}
}

/* CompleteServiceJobByServiceJobIDNotFound describes a response with status code 404, with default header values.

The resource specified does not exist.
*/
type CompleteServiceJobByServiceJobIDNotFound struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *services_models.CompleteServiceJobByServiceJobIDResponse
}

func (o *CompleteServiceJobByServiceJobIDNotFound) Error() string {
	return fmt.Sprintf("[PUT /service/v1/serviceJobs/{serviceJobId}/completions][%d] completeServiceJobByServiceJobIdNotFound  %+v", 404, o.Payload)
}
func (o *CompleteServiceJobByServiceJobIDNotFound) GetPayload() *services_models.CompleteServiceJobByServiceJobIDResponse {
	return o.Payload
}

func (o *CompleteServiceJobByServiceJobIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(services_models.CompleteServiceJobByServiceJobIDResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCompleteServiceJobByServiceJobIDRequestEntityTooLarge creates a CompleteServiceJobByServiceJobIDRequestEntityTooLarge with default headers values
func NewCompleteServiceJobByServiceJobIDRequestEntityTooLarge() *CompleteServiceJobByServiceJobIDRequestEntityTooLarge {
	return &CompleteServiceJobByServiceJobIDRequestEntityTooLarge{}
}

/* CompleteServiceJobByServiceJobIDRequestEntityTooLarge describes a response with status code 413, with default header values.

The request size exceeded the maximum accepted size.
*/
type CompleteServiceJobByServiceJobIDRequestEntityTooLarge struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *services_models.CompleteServiceJobByServiceJobIDResponse
}

func (o *CompleteServiceJobByServiceJobIDRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PUT /service/v1/serviceJobs/{serviceJobId}/completions][%d] completeServiceJobByServiceJobIdRequestEntityTooLarge  %+v", 413, o.Payload)
}
func (o *CompleteServiceJobByServiceJobIDRequestEntityTooLarge) GetPayload() *services_models.CompleteServiceJobByServiceJobIDResponse {
	return o.Payload
}

func (o *CompleteServiceJobByServiceJobIDRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(services_models.CompleteServiceJobByServiceJobIDResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCompleteServiceJobByServiceJobIDUnsupportedMediaType creates a CompleteServiceJobByServiceJobIDUnsupportedMediaType with default headers values
func NewCompleteServiceJobByServiceJobIDUnsupportedMediaType() *CompleteServiceJobByServiceJobIDUnsupportedMediaType {
	return &CompleteServiceJobByServiceJobIDUnsupportedMediaType{}
}

/* CompleteServiceJobByServiceJobIDUnsupportedMediaType describes a response with status code 415, with default header values.

The request payload is in an unsupported format.
*/
type CompleteServiceJobByServiceJobIDUnsupportedMediaType struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *services_models.CompleteServiceJobByServiceJobIDResponse
}

func (o *CompleteServiceJobByServiceJobIDUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PUT /service/v1/serviceJobs/{serviceJobId}/completions][%d] completeServiceJobByServiceJobIdUnsupportedMediaType  %+v", 415, o.Payload)
}
func (o *CompleteServiceJobByServiceJobIDUnsupportedMediaType) GetPayload() *services_models.CompleteServiceJobByServiceJobIDResponse {
	return o.Payload
}

func (o *CompleteServiceJobByServiceJobIDUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(services_models.CompleteServiceJobByServiceJobIDResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCompleteServiceJobByServiceJobIDUnprocessableEntity creates a CompleteServiceJobByServiceJobIDUnprocessableEntity with default headers values
func NewCompleteServiceJobByServiceJobIDUnprocessableEntity() *CompleteServiceJobByServiceJobIDUnprocessableEntity {
	return &CompleteServiceJobByServiceJobIDUnprocessableEntity{}
}

/* CompleteServiceJobByServiceJobIDUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable Entity. Unable to process the contained instructions
*/
type CompleteServiceJobByServiceJobIDUnprocessableEntity struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *services_models.CompleteServiceJobByServiceJobIDResponse
}

func (o *CompleteServiceJobByServiceJobIDUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /service/v1/serviceJobs/{serviceJobId}/completions][%d] completeServiceJobByServiceJobIdUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *CompleteServiceJobByServiceJobIDUnprocessableEntity) GetPayload() *services_models.CompleteServiceJobByServiceJobIDResponse {
	return o.Payload
}

func (o *CompleteServiceJobByServiceJobIDUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(services_models.CompleteServiceJobByServiceJobIDResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCompleteServiceJobByServiceJobIDTooManyRequests creates a CompleteServiceJobByServiceJobIDTooManyRequests with default headers values
func NewCompleteServiceJobByServiceJobIDTooManyRequests() *CompleteServiceJobByServiceJobIDTooManyRequests {
	return &CompleteServiceJobByServiceJobIDTooManyRequests{}
}

/* CompleteServiceJobByServiceJobIDTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type CompleteServiceJobByServiceJobIDTooManyRequests struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *services_models.CompleteServiceJobByServiceJobIDResponse
}

func (o *CompleteServiceJobByServiceJobIDTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /service/v1/serviceJobs/{serviceJobId}/completions][%d] completeServiceJobByServiceJobIdTooManyRequests  %+v", 429, o.Payload)
}
func (o *CompleteServiceJobByServiceJobIDTooManyRequests) GetPayload() *services_models.CompleteServiceJobByServiceJobIDResponse {
	return o.Payload
}

func (o *CompleteServiceJobByServiceJobIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(services_models.CompleteServiceJobByServiceJobIDResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCompleteServiceJobByServiceJobIDInternalServerError creates a CompleteServiceJobByServiceJobIDInternalServerError with default headers values
func NewCompleteServiceJobByServiceJobIDInternalServerError() *CompleteServiceJobByServiceJobIDInternalServerError {
	return &CompleteServiceJobByServiceJobIDInternalServerError{}
}

/* CompleteServiceJobByServiceJobIDInternalServerError describes a response with status code 500, with default header values.

An unexpected condition occurred that prevented the server from fulfilling the request.
*/
type CompleteServiceJobByServiceJobIDInternalServerError struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *services_models.CompleteServiceJobByServiceJobIDResponse
}

func (o *CompleteServiceJobByServiceJobIDInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /service/v1/serviceJobs/{serviceJobId}/completions][%d] completeServiceJobByServiceJobIdInternalServerError  %+v", 500, o.Payload)
}
func (o *CompleteServiceJobByServiceJobIDInternalServerError) GetPayload() *services_models.CompleteServiceJobByServiceJobIDResponse {
	return o.Payload
}

func (o *CompleteServiceJobByServiceJobIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(services_models.CompleteServiceJobByServiceJobIDResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCompleteServiceJobByServiceJobIDServiceUnavailable creates a CompleteServiceJobByServiceJobIDServiceUnavailable with default headers values
func NewCompleteServiceJobByServiceJobIDServiceUnavailable() *CompleteServiceJobByServiceJobIDServiceUnavailable {
	return &CompleteServiceJobByServiceJobIDServiceUnavailable{}
}

/* CompleteServiceJobByServiceJobIDServiceUnavailable describes a response with status code 503, with default header values.

Temporary overloading or maintenance of the server.
*/
type CompleteServiceJobByServiceJobIDServiceUnavailable struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *services_models.CompleteServiceJobByServiceJobIDResponse
}

func (o *CompleteServiceJobByServiceJobIDServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /service/v1/serviceJobs/{serviceJobId}/completions][%d] completeServiceJobByServiceJobIdServiceUnavailable  %+v", 503, o.Payload)
}
func (o *CompleteServiceJobByServiceJobIDServiceUnavailable) GetPayload() *services_models.CompleteServiceJobByServiceJobIDResponse {
	return o.Payload
}

func (o *CompleteServiceJobByServiceJobIDServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(services_models.CompleteServiceJobByServiceJobIDResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
