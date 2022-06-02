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

// GetServiceJobByServiceJobIDReader is a Reader for the GetServiceJobByServiceJobID structure.
type GetServiceJobByServiceJobIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetServiceJobByServiceJobIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetServiceJobByServiceJobIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetServiceJobByServiceJobIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetServiceJobByServiceJobIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetServiceJobByServiceJobIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetServiceJobByServiceJobIDRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetServiceJobByServiceJobIDUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewGetServiceJobByServiceJobIDUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetServiceJobByServiceJobIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetServiceJobByServiceJobIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetServiceJobByServiceJobIDServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetServiceJobByServiceJobIDOK creates a GetServiceJobByServiceJobIDOK with default headers values
func NewGetServiceJobByServiceJobIDOK() *GetServiceJobByServiceJobIDOK {
	return &GetServiceJobByServiceJobIDOK{}
}

/* GetServiceJobByServiceJobIDOK describes a response with status code 200, with default header values.

Success response
*/
type GetServiceJobByServiceJobIDOK struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *services_models.GetServiceJobByServiceJobIDResponse
}

func (o *GetServiceJobByServiceJobIDOK) Error() string {
	return fmt.Sprintf("[GET /service/v1/serviceJobs/{serviceJobId}][%d] getServiceJobByServiceJobIdOK  %+v", 200, o.Payload)
}
func (o *GetServiceJobByServiceJobIDOK) GetPayload() *services_models.GetServiceJobByServiceJobIDResponse {
	return o.Payload
}

func (o *GetServiceJobByServiceJobIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(services_models.GetServiceJobByServiceJobIDResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetServiceJobByServiceJobIDBadRequest creates a GetServiceJobByServiceJobIDBadRequest with default headers values
func NewGetServiceJobByServiceJobIDBadRequest() *GetServiceJobByServiceJobIDBadRequest {
	return &GetServiceJobByServiceJobIDBadRequest{}
}

/* GetServiceJobByServiceJobIDBadRequest describes a response with status code 400, with default header values.

Request has missing or invalid parameters and cannot be parsed.
*/
type GetServiceJobByServiceJobIDBadRequest struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *services_models.GetServiceJobByServiceJobIDResponse
}

func (o *GetServiceJobByServiceJobIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /service/v1/serviceJobs/{serviceJobId}][%d] getServiceJobByServiceJobIdBadRequest  %+v", 400, o.Payload)
}
func (o *GetServiceJobByServiceJobIDBadRequest) GetPayload() *services_models.GetServiceJobByServiceJobIDResponse {
	return o.Payload
}

func (o *GetServiceJobByServiceJobIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(services_models.GetServiceJobByServiceJobIDResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetServiceJobByServiceJobIDForbidden creates a GetServiceJobByServiceJobIDForbidden with default headers values
func NewGetServiceJobByServiceJobIDForbidden() *GetServiceJobByServiceJobIDForbidden {
	return &GetServiceJobByServiceJobIDForbidden{}
}

/* GetServiceJobByServiceJobIDForbidden describes a response with status code 403, with default header values.

Indicates that access to the resource is forbidden. Possible reasons include Access Denied, Unauthorized, Expired Token, or Invalid Signature.
*/
type GetServiceJobByServiceJobIDForbidden struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *services_models.GetServiceJobByServiceJobIDResponse
}

func (o *GetServiceJobByServiceJobIDForbidden) Error() string {
	return fmt.Sprintf("[GET /service/v1/serviceJobs/{serviceJobId}][%d] getServiceJobByServiceJobIdForbidden  %+v", 403, o.Payload)
}
func (o *GetServiceJobByServiceJobIDForbidden) GetPayload() *services_models.GetServiceJobByServiceJobIDResponse {
	return o.Payload
}

func (o *GetServiceJobByServiceJobIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(services_models.GetServiceJobByServiceJobIDResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetServiceJobByServiceJobIDNotFound creates a GetServiceJobByServiceJobIDNotFound with default headers values
func NewGetServiceJobByServiceJobIDNotFound() *GetServiceJobByServiceJobIDNotFound {
	return &GetServiceJobByServiceJobIDNotFound{}
}

/* GetServiceJobByServiceJobIDNotFound describes a response with status code 404, with default header values.

The resource specified does not exist.
*/
type GetServiceJobByServiceJobIDNotFound struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *services_models.GetServiceJobByServiceJobIDResponse
}

func (o *GetServiceJobByServiceJobIDNotFound) Error() string {
	return fmt.Sprintf("[GET /service/v1/serviceJobs/{serviceJobId}][%d] getServiceJobByServiceJobIdNotFound  %+v", 404, o.Payload)
}
func (o *GetServiceJobByServiceJobIDNotFound) GetPayload() *services_models.GetServiceJobByServiceJobIDResponse {
	return o.Payload
}

func (o *GetServiceJobByServiceJobIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(services_models.GetServiceJobByServiceJobIDResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetServiceJobByServiceJobIDRequestEntityTooLarge creates a GetServiceJobByServiceJobIDRequestEntityTooLarge with default headers values
func NewGetServiceJobByServiceJobIDRequestEntityTooLarge() *GetServiceJobByServiceJobIDRequestEntityTooLarge {
	return &GetServiceJobByServiceJobIDRequestEntityTooLarge{}
}

/* GetServiceJobByServiceJobIDRequestEntityTooLarge describes a response with status code 413, with default header values.

The request size exceeded the maximum accepted size.
*/
type GetServiceJobByServiceJobIDRequestEntityTooLarge struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *services_models.GetServiceJobByServiceJobIDResponse
}

func (o *GetServiceJobByServiceJobIDRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /service/v1/serviceJobs/{serviceJobId}][%d] getServiceJobByServiceJobIdRequestEntityTooLarge  %+v", 413, o.Payload)
}
func (o *GetServiceJobByServiceJobIDRequestEntityTooLarge) GetPayload() *services_models.GetServiceJobByServiceJobIDResponse {
	return o.Payload
}

func (o *GetServiceJobByServiceJobIDRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(services_models.GetServiceJobByServiceJobIDResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetServiceJobByServiceJobIDUnsupportedMediaType creates a GetServiceJobByServiceJobIDUnsupportedMediaType with default headers values
func NewGetServiceJobByServiceJobIDUnsupportedMediaType() *GetServiceJobByServiceJobIDUnsupportedMediaType {
	return &GetServiceJobByServiceJobIDUnsupportedMediaType{}
}

/* GetServiceJobByServiceJobIDUnsupportedMediaType describes a response with status code 415, with default header values.

The request payload is in an unsupported format.
*/
type GetServiceJobByServiceJobIDUnsupportedMediaType struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *services_models.GetServiceJobByServiceJobIDResponse
}

func (o *GetServiceJobByServiceJobIDUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /service/v1/serviceJobs/{serviceJobId}][%d] getServiceJobByServiceJobIdUnsupportedMediaType  %+v", 415, o.Payload)
}
func (o *GetServiceJobByServiceJobIDUnsupportedMediaType) GetPayload() *services_models.GetServiceJobByServiceJobIDResponse {
	return o.Payload
}

func (o *GetServiceJobByServiceJobIDUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(services_models.GetServiceJobByServiceJobIDResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetServiceJobByServiceJobIDUnprocessableEntity creates a GetServiceJobByServiceJobIDUnprocessableEntity with default headers values
func NewGetServiceJobByServiceJobIDUnprocessableEntity() *GetServiceJobByServiceJobIDUnprocessableEntity {
	return &GetServiceJobByServiceJobIDUnprocessableEntity{}
}

/* GetServiceJobByServiceJobIDUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable Entity. Unable to process the contained instructions.
*/
type GetServiceJobByServiceJobIDUnprocessableEntity struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *services_models.GetServiceJobByServiceJobIDResponse
}

func (o *GetServiceJobByServiceJobIDUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /service/v1/serviceJobs/{serviceJobId}][%d] getServiceJobByServiceJobIdUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *GetServiceJobByServiceJobIDUnprocessableEntity) GetPayload() *services_models.GetServiceJobByServiceJobIDResponse {
	return o.Payload
}

func (o *GetServiceJobByServiceJobIDUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(services_models.GetServiceJobByServiceJobIDResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetServiceJobByServiceJobIDTooManyRequests creates a GetServiceJobByServiceJobIDTooManyRequests with default headers values
func NewGetServiceJobByServiceJobIDTooManyRequests() *GetServiceJobByServiceJobIDTooManyRequests {
	return &GetServiceJobByServiceJobIDTooManyRequests{}
}

/* GetServiceJobByServiceJobIDTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type GetServiceJobByServiceJobIDTooManyRequests struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *services_models.GetServiceJobByServiceJobIDResponse
}

func (o *GetServiceJobByServiceJobIDTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /service/v1/serviceJobs/{serviceJobId}][%d] getServiceJobByServiceJobIdTooManyRequests  %+v", 429, o.Payload)
}
func (o *GetServiceJobByServiceJobIDTooManyRequests) GetPayload() *services_models.GetServiceJobByServiceJobIDResponse {
	return o.Payload
}

func (o *GetServiceJobByServiceJobIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(services_models.GetServiceJobByServiceJobIDResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetServiceJobByServiceJobIDInternalServerError creates a GetServiceJobByServiceJobIDInternalServerError with default headers values
func NewGetServiceJobByServiceJobIDInternalServerError() *GetServiceJobByServiceJobIDInternalServerError {
	return &GetServiceJobByServiceJobIDInternalServerError{}
}

/* GetServiceJobByServiceJobIDInternalServerError describes a response with status code 500, with default header values.

An unexpected condition occurred that prevented the server from fulfilling the request.
*/
type GetServiceJobByServiceJobIDInternalServerError struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *services_models.GetServiceJobByServiceJobIDResponse
}

func (o *GetServiceJobByServiceJobIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /service/v1/serviceJobs/{serviceJobId}][%d] getServiceJobByServiceJobIdInternalServerError  %+v", 500, o.Payload)
}
func (o *GetServiceJobByServiceJobIDInternalServerError) GetPayload() *services_models.GetServiceJobByServiceJobIDResponse {
	return o.Payload
}

func (o *GetServiceJobByServiceJobIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(services_models.GetServiceJobByServiceJobIDResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetServiceJobByServiceJobIDServiceUnavailable creates a GetServiceJobByServiceJobIDServiceUnavailable with default headers values
func NewGetServiceJobByServiceJobIDServiceUnavailable() *GetServiceJobByServiceJobIDServiceUnavailable {
	return &GetServiceJobByServiceJobIDServiceUnavailable{}
}

/* GetServiceJobByServiceJobIDServiceUnavailable describes a response with status code 503, with default header values.

Temporary overloading or maintenance of the server.
*/
type GetServiceJobByServiceJobIDServiceUnavailable struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *services_models.GetServiceJobByServiceJobIDResponse
}

func (o *GetServiceJobByServiceJobIDServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /service/v1/serviceJobs/{serviceJobId}][%d] getServiceJobByServiceJobIdServiceUnavailable  %+v", 503, o.Payload)
}
func (o *GetServiceJobByServiceJobIDServiceUnavailable) GetPayload() *services_models.GetServiceJobByServiceJobIDResponse {
	return o.Payload
}

func (o *GetServiceJobByServiceJobIDServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(services_models.GetServiceJobByServiceJobIDResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
