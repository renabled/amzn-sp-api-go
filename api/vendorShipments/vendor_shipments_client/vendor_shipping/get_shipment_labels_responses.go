// Code generated by go-swagger; DO NOT EDIT.

package vendor_shipping

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/xamandar/amzn-sp-api-go/api/vendorShipments/vendor_shipments_models"
)

// GetShipmentLabelsReader is a Reader for the GetShipmentLabels structure.
type GetShipmentLabelsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetShipmentLabelsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetShipmentLabelsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetShipmentLabelsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetShipmentLabelsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetShipmentLabelsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetShipmentLabelsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetShipmentLabelsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetShipmentLabelsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetShipmentLabelsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetShipmentLabelsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetShipmentLabelsOK creates a GetShipmentLabelsOK with default headers values
func NewGetShipmentLabelsOK() *GetShipmentLabelsOK {
	return &GetShipmentLabelsOK{}
}

/*
GetShipmentLabelsOK describes a response with status code 200, with default header values.

Success.
*/
type GetShipmentLabelsOK struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_shipments_models.GetShipmentLabels
}

// IsSuccess returns true when this get shipment labels o k response has a 2xx status code
func (o *GetShipmentLabelsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get shipment labels o k response has a 3xx status code
func (o *GetShipmentLabelsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get shipment labels o k response has a 4xx status code
func (o *GetShipmentLabelsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get shipment labels o k response has a 5xx status code
func (o *GetShipmentLabelsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get shipment labels o k response a status code equal to that given
func (o *GetShipmentLabelsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetShipmentLabelsOK) Error() string {
	return fmt.Sprintf("[GET /vendor/shipping/v1/transportLabels][%d] getShipmentLabelsOK  %+v", 200, o.Payload)
}

func (o *GetShipmentLabelsOK) String() string {
	return fmt.Sprintf("[GET /vendor/shipping/v1/transportLabels][%d] getShipmentLabelsOK  %+v", 200, o.Payload)
}

func (o *GetShipmentLabelsOK) GetPayload() *vendor_shipments_models.GetShipmentLabels {
	return o.Payload
}

func (o *GetShipmentLabelsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(vendor_shipments_models.GetShipmentLabels)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetShipmentLabelsBadRequest creates a GetShipmentLabelsBadRequest with default headers values
func NewGetShipmentLabelsBadRequest() *GetShipmentLabelsBadRequest {
	return &GetShipmentLabelsBadRequest{}
}

/*
GetShipmentLabelsBadRequest describes a response with status code 400, with default header values.

Request has missing or invalid parameters and cannot be parsed.
*/
type GetShipmentLabelsBadRequest struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_shipments_models.GetShipmentLabels
}

// IsSuccess returns true when this get shipment labels bad request response has a 2xx status code
func (o *GetShipmentLabelsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get shipment labels bad request response has a 3xx status code
func (o *GetShipmentLabelsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get shipment labels bad request response has a 4xx status code
func (o *GetShipmentLabelsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get shipment labels bad request response has a 5xx status code
func (o *GetShipmentLabelsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get shipment labels bad request response a status code equal to that given
func (o *GetShipmentLabelsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetShipmentLabelsBadRequest) Error() string {
	return fmt.Sprintf("[GET /vendor/shipping/v1/transportLabels][%d] getShipmentLabelsBadRequest  %+v", 400, o.Payload)
}

func (o *GetShipmentLabelsBadRequest) String() string {
	return fmt.Sprintf("[GET /vendor/shipping/v1/transportLabels][%d] getShipmentLabelsBadRequest  %+v", 400, o.Payload)
}

func (o *GetShipmentLabelsBadRequest) GetPayload() *vendor_shipments_models.GetShipmentLabels {
	return o.Payload
}

func (o *GetShipmentLabelsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(vendor_shipments_models.GetShipmentLabels)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetShipmentLabelsUnauthorized creates a GetShipmentLabelsUnauthorized with default headers values
func NewGetShipmentLabelsUnauthorized() *GetShipmentLabelsUnauthorized {
	return &GetShipmentLabelsUnauthorized{}
}

/*
GetShipmentLabelsUnauthorized describes a response with status code 401, with default header values.

The request's Authorization header is not formatted correctly or does not contain a valid token.
*/
type GetShipmentLabelsUnauthorized struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_shipments_models.GetShipmentLabels
}

// IsSuccess returns true when this get shipment labels unauthorized response has a 2xx status code
func (o *GetShipmentLabelsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get shipment labels unauthorized response has a 3xx status code
func (o *GetShipmentLabelsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get shipment labels unauthorized response has a 4xx status code
func (o *GetShipmentLabelsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get shipment labels unauthorized response has a 5xx status code
func (o *GetShipmentLabelsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get shipment labels unauthorized response a status code equal to that given
func (o *GetShipmentLabelsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetShipmentLabelsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /vendor/shipping/v1/transportLabels][%d] getShipmentLabelsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetShipmentLabelsUnauthorized) String() string {
	return fmt.Sprintf("[GET /vendor/shipping/v1/transportLabels][%d] getShipmentLabelsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetShipmentLabelsUnauthorized) GetPayload() *vendor_shipments_models.GetShipmentLabels {
	return o.Payload
}

func (o *GetShipmentLabelsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(vendor_shipments_models.GetShipmentLabels)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetShipmentLabelsForbidden creates a GetShipmentLabelsForbidden with default headers values
func NewGetShipmentLabelsForbidden() *GetShipmentLabelsForbidden {
	return &GetShipmentLabelsForbidden{}
}

/*
GetShipmentLabelsForbidden describes a response with status code 403, with default header values.

Indicates that access to the resource is forbidden. Possible reasons include Access Denied, Unauthorized, Expired Token, or Invalid Signature.
*/
type GetShipmentLabelsForbidden struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_shipments_models.GetShipmentLabels
}

// IsSuccess returns true when this get shipment labels forbidden response has a 2xx status code
func (o *GetShipmentLabelsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get shipment labels forbidden response has a 3xx status code
func (o *GetShipmentLabelsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get shipment labels forbidden response has a 4xx status code
func (o *GetShipmentLabelsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get shipment labels forbidden response has a 5xx status code
func (o *GetShipmentLabelsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get shipment labels forbidden response a status code equal to that given
func (o *GetShipmentLabelsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetShipmentLabelsForbidden) Error() string {
	return fmt.Sprintf("[GET /vendor/shipping/v1/transportLabels][%d] getShipmentLabelsForbidden  %+v", 403, o.Payload)
}

func (o *GetShipmentLabelsForbidden) String() string {
	return fmt.Sprintf("[GET /vendor/shipping/v1/transportLabels][%d] getShipmentLabelsForbidden  %+v", 403, o.Payload)
}

func (o *GetShipmentLabelsForbidden) GetPayload() *vendor_shipments_models.GetShipmentLabels {
	return o.Payload
}

func (o *GetShipmentLabelsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(vendor_shipments_models.GetShipmentLabels)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetShipmentLabelsNotFound creates a GetShipmentLabelsNotFound with default headers values
func NewGetShipmentLabelsNotFound() *GetShipmentLabelsNotFound {
	return &GetShipmentLabelsNotFound{}
}

/*
GetShipmentLabelsNotFound describes a response with status code 404, with default header values.

The resource specified does not exist.
*/
type GetShipmentLabelsNotFound struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_shipments_models.GetShipmentLabels
}

// IsSuccess returns true when this get shipment labels not found response has a 2xx status code
func (o *GetShipmentLabelsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get shipment labels not found response has a 3xx status code
func (o *GetShipmentLabelsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get shipment labels not found response has a 4xx status code
func (o *GetShipmentLabelsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get shipment labels not found response has a 5xx status code
func (o *GetShipmentLabelsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get shipment labels not found response a status code equal to that given
func (o *GetShipmentLabelsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetShipmentLabelsNotFound) Error() string {
	return fmt.Sprintf("[GET /vendor/shipping/v1/transportLabels][%d] getShipmentLabelsNotFound  %+v", 404, o.Payload)
}

func (o *GetShipmentLabelsNotFound) String() string {
	return fmt.Sprintf("[GET /vendor/shipping/v1/transportLabels][%d] getShipmentLabelsNotFound  %+v", 404, o.Payload)
}

func (o *GetShipmentLabelsNotFound) GetPayload() *vendor_shipments_models.GetShipmentLabels {
	return o.Payload
}

func (o *GetShipmentLabelsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(vendor_shipments_models.GetShipmentLabels)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetShipmentLabelsUnsupportedMediaType creates a GetShipmentLabelsUnsupportedMediaType with default headers values
func NewGetShipmentLabelsUnsupportedMediaType() *GetShipmentLabelsUnsupportedMediaType {
	return &GetShipmentLabelsUnsupportedMediaType{}
}

/*
GetShipmentLabelsUnsupportedMediaType describes a response with status code 415, with default header values.

The request payload is in an unsupported format.
*/
type GetShipmentLabelsUnsupportedMediaType struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_shipments_models.GetShipmentLabels
}

// IsSuccess returns true when this get shipment labels unsupported media type response has a 2xx status code
func (o *GetShipmentLabelsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get shipment labels unsupported media type response has a 3xx status code
func (o *GetShipmentLabelsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get shipment labels unsupported media type response has a 4xx status code
func (o *GetShipmentLabelsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get shipment labels unsupported media type response has a 5xx status code
func (o *GetShipmentLabelsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get shipment labels unsupported media type response a status code equal to that given
func (o *GetShipmentLabelsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetShipmentLabelsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /vendor/shipping/v1/transportLabels][%d] getShipmentLabelsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetShipmentLabelsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /vendor/shipping/v1/transportLabels][%d] getShipmentLabelsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetShipmentLabelsUnsupportedMediaType) GetPayload() *vendor_shipments_models.GetShipmentLabels {
	return o.Payload
}

func (o *GetShipmentLabelsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(vendor_shipments_models.GetShipmentLabels)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetShipmentLabelsTooManyRequests creates a GetShipmentLabelsTooManyRequests with default headers values
func NewGetShipmentLabelsTooManyRequests() *GetShipmentLabelsTooManyRequests {
	return &GetShipmentLabelsTooManyRequests{}
}

/*
GetShipmentLabelsTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type GetShipmentLabelsTooManyRequests struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_shipments_models.GetShipmentLabels
}

// IsSuccess returns true when this get shipment labels too many requests response has a 2xx status code
func (o *GetShipmentLabelsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get shipment labels too many requests response has a 3xx status code
func (o *GetShipmentLabelsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get shipment labels too many requests response has a 4xx status code
func (o *GetShipmentLabelsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get shipment labels too many requests response has a 5xx status code
func (o *GetShipmentLabelsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get shipment labels too many requests response a status code equal to that given
func (o *GetShipmentLabelsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetShipmentLabelsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /vendor/shipping/v1/transportLabels][%d] getShipmentLabelsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetShipmentLabelsTooManyRequests) String() string {
	return fmt.Sprintf("[GET /vendor/shipping/v1/transportLabels][%d] getShipmentLabelsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetShipmentLabelsTooManyRequests) GetPayload() *vendor_shipments_models.GetShipmentLabels {
	return o.Payload
}

func (o *GetShipmentLabelsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(vendor_shipments_models.GetShipmentLabels)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetShipmentLabelsInternalServerError creates a GetShipmentLabelsInternalServerError with default headers values
func NewGetShipmentLabelsInternalServerError() *GetShipmentLabelsInternalServerError {
	return &GetShipmentLabelsInternalServerError{}
}

/*
GetShipmentLabelsInternalServerError describes a response with status code 500, with default header values.

An unexpected condition occurred that prevented the server from fulfilling the request.
*/
type GetShipmentLabelsInternalServerError struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_shipments_models.GetShipmentLabels
}

// IsSuccess returns true when this get shipment labels internal server error response has a 2xx status code
func (o *GetShipmentLabelsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get shipment labels internal server error response has a 3xx status code
func (o *GetShipmentLabelsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get shipment labels internal server error response has a 4xx status code
func (o *GetShipmentLabelsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get shipment labels internal server error response has a 5xx status code
func (o *GetShipmentLabelsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get shipment labels internal server error response a status code equal to that given
func (o *GetShipmentLabelsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetShipmentLabelsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /vendor/shipping/v1/transportLabels][%d] getShipmentLabelsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetShipmentLabelsInternalServerError) String() string {
	return fmt.Sprintf("[GET /vendor/shipping/v1/transportLabels][%d] getShipmentLabelsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetShipmentLabelsInternalServerError) GetPayload() *vendor_shipments_models.GetShipmentLabels {
	return o.Payload
}

func (o *GetShipmentLabelsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(vendor_shipments_models.GetShipmentLabels)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetShipmentLabelsServiceUnavailable creates a GetShipmentLabelsServiceUnavailable with default headers values
func NewGetShipmentLabelsServiceUnavailable() *GetShipmentLabelsServiceUnavailable {
	return &GetShipmentLabelsServiceUnavailable{}
}

/*
GetShipmentLabelsServiceUnavailable describes a response with status code 503, with default header values.

Temporary overloading or maintenance of the server.
*/
type GetShipmentLabelsServiceUnavailable struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_shipments_models.GetShipmentLabels
}

// IsSuccess returns true when this get shipment labels service unavailable response has a 2xx status code
func (o *GetShipmentLabelsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get shipment labels service unavailable response has a 3xx status code
func (o *GetShipmentLabelsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get shipment labels service unavailable response has a 4xx status code
func (o *GetShipmentLabelsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get shipment labels service unavailable response has a 5xx status code
func (o *GetShipmentLabelsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get shipment labels service unavailable response a status code equal to that given
func (o *GetShipmentLabelsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetShipmentLabelsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /vendor/shipping/v1/transportLabels][%d] getShipmentLabelsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetShipmentLabelsServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /vendor/shipping/v1/transportLabels][%d] getShipmentLabelsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetShipmentLabelsServiceUnavailable) GetPayload() *vendor_shipments_models.GetShipmentLabels {
	return o.Payload
}

func (o *GetShipmentLabelsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(vendor_shipments_models.GetShipmentLabels)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}