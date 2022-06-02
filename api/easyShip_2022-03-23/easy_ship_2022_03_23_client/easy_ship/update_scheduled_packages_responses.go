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

// UpdateScheduledPackagesReader is a Reader for the UpdateScheduledPackages structure.
type UpdateScheduledPackagesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateScheduledPackagesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateScheduledPackagesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateScheduledPackagesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateScheduledPackagesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateScheduledPackagesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateScheduledPackagesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewUpdateScheduledPackagesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewUpdateScheduledPackagesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateScheduledPackagesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewUpdateScheduledPackagesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateScheduledPackagesOK creates a UpdateScheduledPackagesOK with default headers values
func NewUpdateScheduledPackagesOK() *UpdateScheduledPackagesOK {
	return &UpdateScheduledPackagesOK{}
}

/* UpdateScheduledPackagesOK describes a response with status code 200, with default header values.

Success
*/
type UpdateScheduledPackagesOK struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *easy_ship_2022_03_23_models.Packages
}

func (o *UpdateScheduledPackagesOK) Error() string {
	return fmt.Sprintf("[PATCH /easyShip/2022-03-23/package][%d] updateScheduledPackagesOK  %+v", 200, o.Payload)
}
func (o *UpdateScheduledPackagesOK) GetPayload() *easy_ship_2022_03_23_models.Packages {
	return o.Payload
}

func (o *UpdateScheduledPackagesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(easy_ship_2022_03_23_models.Packages)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateScheduledPackagesBadRequest creates a UpdateScheduledPackagesBadRequest with default headers values
func NewUpdateScheduledPackagesBadRequest() *UpdateScheduledPackagesBadRequest {
	return &UpdateScheduledPackagesBadRequest{}
}

/* UpdateScheduledPackagesBadRequest describes a response with status code 400, with default header values.

Request has missing or invalid parameters and cannot be parsed.
*/
type UpdateScheduledPackagesBadRequest struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *easy_ship_2022_03_23_models.ErrorList
}

func (o *UpdateScheduledPackagesBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /easyShip/2022-03-23/package][%d] updateScheduledPackagesBadRequest  %+v", 400, o.Payload)
}
func (o *UpdateScheduledPackagesBadRequest) GetPayload() *easy_ship_2022_03_23_models.ErrorList {
	return o.Payload
}

func (o *UpdateScheduledPackagesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateScheduledPackagesUnauthorized creates a UpdateScheduledPackagesUnauthorized with default headers values
func NewUpdateScheduledPackagesUnauthorized() *UpdateScheduledPackagesUnauthorized {
	return &UpdateScheduledPackagesUnauthorized{}
}

/* UpdateScheduledPackagesUnauthorized describes a response with status code 401, with default header values.

The request's Authorization header is not formatted correctly or does not contain a valid token.
*/
type UpdateScheduledPackagesUnauthorized struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *easy_ship_2022_03_23_models.ErrorList
}

func (o *UpdateScheduledPackagesUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /easyShip/2022-03-23/package][%d] updateScheduledPackagesUnauthorized  %+v", 401, o.Payload)
}
func (o *UpdateScheduledPackagesUnauthorized) GetPayload() *easy_ship_2022_03_23_models.ErrorList {
	return o.Payload
}

func (o *UpdateScheduledPackagesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateScheduledPackagesForbidden creates a UpdateScheduledPackagesForbidden with default headers values
func NewUpdateScheduledPackagesForbidden() *UpdateScheduledPackagesForbidden {
	return &UpdateScheduledPackagesForbidden{}
}

/* UpdateScheduledPackagesForbidden describes a response with status code 403, with default header values.

Indicates access to the resource is forbidden. Possible reasons include Access Denied, Unauthorized, Expired Token, or Invalid Signature.
*/
type UpdateScheduledPackagesForbidden struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *easy_ship_2022_03_23_models.ErrorList
}

func (o *UpdateScheduledPackagesForbidden) Error() string {
	return fmt.Sprintf("[PATCH /easyShip/2022-03-23/package][%d] updateScheduledPackagesForbidden  %+v", 403, o.Payload)
}
func (o *UpdateScheduledPackagesForbidden) GetPayload() *easy_ship_2022_03_23_models.ErrorList {
	return o.Payload
}

func (o *UpdateScheduledPackagesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateScheduledPackagesNotFound creates a UpdateScheduledPackagesNotFound with default headers values
func NewUpdateScheduledPackagesNotFound() *UpdateScheduledPackagesNotFound {
	return &UpdateScheduledPackagesNotFound{}
}

/* UpdateScheduledPackagesNotFound describes a response with status code 404, with default header values.

The specified resource does not exist.
*/
type UpdateScheduledPackagesNotFound struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *easy_ship_2022_03_23_models.ErrorList
}

func (o *UpdateScheduledPackagesNotFound) Error() string {
	return fmt.Sprintf("[PATCH /easyShip/2022-03-23/package][%d] updateScheduledPackagesNotFound  %+v", 404, o.Payload)
}
func (o *UpdateScheduledPackagesNotFound) GetPayload() *easy_ship_2022_03_23_models.ErrorList {
	return o.Payload
}

func (o *UpdateScheduledPackagesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateScheduledPackagesUnsupportedMediaType creates a UpdateScheduledPackagesUnsupportedMediaType with default headers values
func NewUpdateScheduledPackagesUnsupportedMediaType() *UpdateScheduledPackagesUnsupportedMediaType {
	return &UpdateScheduledPackagesUnsupportedMediaType{}
}

/* UpdateScheduledPackagesUnsupportedMediaType describes a response with status code 415, with default header values.

The request payload is in an unsupported format.
*/
type UpdateScheduledPackagesUnsupportedMediaType struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *easy_ship_2022_03_23_models.ErrorList
}

func (o *UpdateScheduledPackagesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PATCH /easyShip/2022-03-23/package][%d] updateScheduledPackagesUnsupportedMediaType  %+v", 415, o.Payload)
}
func (o *UpdateScheduledPackagesUnsupportedMediaType) GetPayload() *easy_ship_2022_03_23_models.ErrorList {
	return o.Payload
}

func (o *UpdateScheduledPackagesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateScheduledPackagesTooManyRequests creates a UpdateScheduledPackagesTooManyRequests with default headers values
func NewUpdateScheduledPackagesTooManyRequests() *UpdateScheduledPackagesTooManyRequests {
	return &UpdateScheduledPackagesTooManyRequests{}
}

/* UpdateScheduledPackagesTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type UpdateScheduledPackagesTooManyRequests struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *easy_ship_2022_03_23_models.ErrorList
}

func (o *UpdateScheduledPackagesTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /easyShip/2022-03-23/package][%d] updateScheduledPackagesTooManyRequests  %+v", 429, o.Payload)
}
func (o *UpdateScheduledPackagesTooManyRequests) GetPayload() *easy_ship_2022_03_23_models.ErrorList {
	return o.Payload
}

func (o *UpdateScheduledPackagesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateScheduledPackagesInternalServerError creates a UpdateScheduledPackagesInternalServerError with default headers values
func NewUpdateScheduledPackagesInternalServerError() *UpdateScheduledPackagesInternalServerError {
	return &UpdateScheduledPackagesInternalServerError{}
}

/* UpdateScheduledPackagesInternalServerError describes a response with status code 500, with default header values.

An unexpected condition occurred that prevented the server from fulfilling the request.
*/
type UpdateScheduledPackagesInternalServerError struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *easy_ship_2022_03_23_models.ErrorList
}

func (o *UpdateScheduledPackagesInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /easyShip/2022-03-23/package][%d] updateScheduledPackagesInternalServerError  %+v", 500, o.Payload)
}
func (o *UpdateScheduledPackagesInternalServerError) GetPayload() *easy_ship_2022_03_23_models.ErrorList {
	return o.Payload
}

func (o *UpdateScheduledPackagesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateScheduledPackagesServiceUnavailable creates a UpdateScheduledPackagesServiceUnavailable with default headers values
func NewUpdateScheduledPackagesServiceUnavailable() *UpdateScheduledPackagesServiceUnavailable {
	return &UpdateScheduledPackagesServiceUnavailable{}
}

/* UpdateScheduledPackagesServiceUnavailable describes a response with status code 503, with default header values.

Temporary overloading or maintenance of the server.
*/
type UpdateScheduledPackagesServiceUnavailable struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *easy_ship_2022_03_23_models.ErrorList
}

func (o *UpdateScheduledPackagesServiceUnavailable) Error() string {
	return fmt.Sprintf("[PATCH /easyShip/2022-03-23/package][%d] updateScheduledPackagesServiceUnavailable  %+v", 503, o.Payload)
}
func (o *UpdateScheduledPackagesServiceUnavailable) GetPayload() *easy_ship_2022_03_23_models.ErrorList {
	return o.Payload
}

func (o *UpdateScheduledPackagesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
