// Code generated by go-swagger; DO NOT EDIT.

package fba_inbound

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/renabled/amzn-sp-api-go/api/fulfillmentInboundv2024-03-20/fulfillment_inboundv2024_03_20_models"
)

// GetDeliveryChallanDocumentReader is a Reader for the GetDeliveryChallanDocument structure.
type GetDeliveryChallanDocumentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDeliveryChallanDocumentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDeliveryChallanDocumentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetDeliveryChallanDocumentBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetDeliveryChallanDocumentForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetDeliveryChallanDocumentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetDeliveryChallanDocumentRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetDeliveryChallanDocumentUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetDeliveryChallanDocumentTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetDeliveryChallanDocumentInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetDeliveryChallanDocumentServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetDeliveryChallanDocumentOK creates a GetDeliveryChallanDocumentOK with default headers values
func NewGetDeliveryChallanDocumentOK() *GetDeliveryChallanDocumentOK {
	return &GetDeliveryChallanDocumentOK{}
}

/*
GetDeliveryChallanDocumentOK describes a response with status code 200, with default header values.

GetDeliveryChallanDocument 200 response
*/
type GetDeliveryChallanDocumentOK struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inboundv2024_03_20_models.GetDeliveryChallanDocumentResponse
}

// IsSuccess returns true when this get delivery challan document o k response has a 2xx status code
func (o *GetDeliveryChallanDocumentOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get delivery challan document o k response has a 3xx status code
func (o *GetDeliveryChallanDocumentOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get delivery challan document o k response has a 4xx status code
func (o *GetDeliveryChallanDocumentOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get delivery challan document o k response has a 5xx status code
func (o *GetDeliveryChallanDocumentOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get delivery challan document o k response a status code equal to that given
func (o *GetDeliveryChallanDocumentOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetDeliveryChallanDocumentOK) Error() string {
	return fmt.Sprintf("[GET /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/deliveryChallanDocument][%d] getDeliveryChallanDocumentOK  %+v", 200, o.Payload)
}

func (o *GetDeliveryChallanDocumentOK) String() string {
	return fmt.Sprintf("[GET /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/deliveryChallanDocument][%d] getDeliveryChallanDocumentOK  %+v", 200, o.Payload)
}

func (o *GetDeliveryChallanDocumentOK) GetPayload() *fulfillment_inboundv2024_03_20_models.GetDeliveryChallanDocumentResponse {
	return o.Payload
}

func (o *GetDeliveryChallanDocumentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fulfillment_inboundv2024_03_20_models.GetDeliveryChallanDocumentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDeliveryChallanDocumentBadRequest creates a GetDeliveryChallanDocumentBadRequest with default headers values
func NewGetDeliveryChallanDocumentBadRequest() *GetDeliveryChallanDocumentBadRequest {
	return &GetDeliveryChallanDocumentBadRequest{}
}

/*
GetDeliveryChallanDocumentBadRequest describes a response with status code 400, with default header values.

Request has missing or invalid parameters and cannot be parsed.
*/
type GetDeliveryChallanDocumentBadRequest struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inboundv2024_03_20_models.ErrorList
}

// IsSuccess returns true when this get delivery challan document bad request response has a 2xx status code
func (o *GetDeliveryChallanDocumentBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get delivery challan document bad request response has a 3xx status code
func (o *GetDeliveryChallanDocumentBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get delivery challan document bad request response has a 4xx status code
func (o *GetDeliveryChallanDocumentBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get delivery challan document bad request response has a 5xx status code
func (o *GetDeliveryChallanDocumentBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get delivery challan document bad request response a status code equal to that given
func (o *GetDeliveryChallanDocumentBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetDeliveryChallanDocumentBadRequest) Error() string {
	return fmt.Sprintf("[GET /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/deliveryChallanDocument][%d] getDeliveryChallanDocumentBadRequest  %+v", 400, o.Payload)
}

func (o *GetDeliveryChallanDocumentBadRequest) String() string {
	return fmt.Sprintf("[GET /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/deliveryChallanDocument][%d] getDeliveryChallanDocumentBadRequest  %+v", 400, o.Payload)
}

func (o *GetDeliveryChallanDocumentBadRequest) GetPayload() *fulfillment_inboundv2024_03_20_models.ErrorList {
	return o.Payload
}

func (o *GetDeliveryChallanDocumentBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fulfillment_inboundv2024_03_20_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDeliveryChallanDocumentForbidden creates a GetDeliveryChallanDocumentForbidden with default headers values
func NewGetDeliveryChallanDocumentForbidden() *GetDeliveryChallanDocumentForbidden {
	return &GetDeliveryChallanDocumentForbidden{}
}

/*
GetDeliveryChallanDocumentForbidden describes a response with status code 403, with default header values.

Indicates that access to the resource is forbidden. Possible reasons include Access Denied, Unauthorized, Expired Token, or Invalid Signature.
*/
type GetDeliveryChallanDocumentForbidden struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inboundv2024_03_20_models.ErrorList
}

// IsSuccess returns true when this get delivery challan document forbidden response has a 2xx status code
func (o *GetDeliveryChallanDocumentForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get delivery challan document forbidden response has a 3xx status code
func (o *GetDeliveryChallanDocumentForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get delivery challan document forbidden response has a 4xx status code
func (o *GetDeliveryChallanDocumentForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get delivery challan document forbidden response has a 5xx status code
func (o *GetDeliveryChallanDocumentForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get delivery challan document forbidden response a status code equal to that given
func (o *GetDeliveryChallanDocumentForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetDeliveryChallanDocumentForbidden) Error() string {
	return fmt.Sprintf("[GET /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/deliveryChallanDocument][%d] getDeliveryChallanDocumentForbidden  %+v", 403, o.Payload)
}

func (o *GetDeliveryChallanDocumentForbidden) String() string {
	return fmt.Sprintf("[GET /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/deliveryChallanDocument][%d] getDeliveryChallanDocumentForbidden  %+v", 403, o.Payload)
}

func (o *GetDeliveryChallanDocumentForbidden) GetPayload() *fulfillment_inboundv2024_03_20_models.ErrorList {
	return o.Payload
}

func (o *GetDeliveryChallanDocumentForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(fulfillment_inboundv2024_03_20_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDeliveryChallanDocumentNotFound creates a GetDeliveryChallanDocumentNotFound with default headers values
func NewGetDeliveryChallanDocumentNotFound() *GetDeliveryChallanDocumentNotFound {
	return &GetDeliveryChallanDocumentNotFound{}
}

/*
GetDeliveryChallanDocumentNotFound describes a response with status code 404, with default header values.

The resource specified does not exist.
*/
type GetDeliveryChallanDocumentNotFound struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inboundv2024_03_20_models.ErrorList
}

// IsSuccess returns true when this get delivery challan document not found response has a 2xx status code
func (o *GetDeliveryChallanDocumentNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get delivery challan document not found response has a 3xx status code
func (o *GetDeliveryChallanDocumentNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get delivery challan document not found response has a 4xx status code
func (o *GetDeliveryChallanDocumentNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get delivery challan document not found response has a 5xx status code
func (o *GetDeliveryChallanDocumentNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get delivery challan document not found response a status code equal to that given
func (o *GetDeliveryChallanDocumentNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetDeliveryChallanDocumentNotFound) Error() string {
	return fmt.Sprintf("[GET /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/deliveryChallanDocument][%d] getDeliveryChallanDocumentNotFound  %+v", 404, o.Payload)
}

func (o *GetDeliveryChallanDocumentNotFound) String() string {
	return fmt.Sprintf("[GET /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/deliveryChallanDocument][%d] getDeliveryChallanDocumentNotFound  %+v", 404, o.Payload)
}

func (o *GetDeliveryChallanDocumentNotFound) GetPayload() *fulfillment_inboundv2024_03_20_models.ErrorList {
	return o.Payload
}

func (o *GetDeliveryChallanDocumentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fulfillment_inboundv2024_03_20_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDeliveryChallanDocumentRequestEntityTooLarge creates a GetDeliveryChallanDocumentRequestEntityTooLarge with default headers values
func NewGetDeliveryChallanDocumentRequestEntityTooLarge() *GetDeliveryChallanDocumentRequestEntityTooLarge {
	return &GetDeliveryChallanDocumentRequestEntityTooLarge{}
}

/*
GetDeliveryChallanDocumentRequestEntityTooLarge describes a response with status code 413, with default header values.

The request size exceeded the maximum accepted size.
*/
type GetDeliveryChallanDocumentRequestEntityTooLarge struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inboundv2024_03_20_models.ErrorList
}

// IsSuccess returns true when this get delivery challan document request entity too large response has a 2xx status code
func (o *GetDeliveryChallanDocumentRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get delivery challan document request entity too large response has a 3xx status code
func (o *GetDeliveryChallanDocumentRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get delivery challan document request entity too large response has a 4xx status code
func (o *GetDeliveryChallanDocumentRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get delivery challan document request entity too large response has a 5xx status code
func (o *GetDeliveryChallanDocumentRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get delivery challan document request entity too large response a status code equal to that given
func (o *GetDeliveryChallanDocumentRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetDeliveryChallanDocumentRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/deliveryChallanDocument][%d] getDeliveryChallanDocumentRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetDeliveryChallanDocumentRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/deliveryChallanDocument][%d] getDeliveryChallanDocumentRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetDeliveryChallanDocumentRequestEntityTooLarge) GetPayload() *fulfillment_inboundv2024_03_20_models.ErrorList {
	return o.Payload
}

func (o *GetDeliveryChallanDocumentRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(fulfillment_inboundv2024_03_20_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDeliveryChallanDocumentUnsupportedMediaType creates a GetDeliveryChallanDocumentUnsupportedMediaType with default headers values
func NewGetDeliveryChallanDocumentUnsupportedMediaType() *GetDeliveryChallanDocumentUnsupportedMediaType {
	return &GetDeliveryChallanDocumentUnsupportedMediaType{}
}

/*
GetDeliveryChallanDocumentUnsupportedMediaType describes a response with status code 415, with default header values.

The request payload is in an unsupported format.
*/
type GetDeliveryChallanDocumentUnsupportedMediaType struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inboundv2024_03_20_models.ErrorList
}

// IsSuccess returns true when this get delivery challan document unsupported media type response has a 2xx status code
func (o *GetDeliveryChallanDocumentUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get delivery challan document unsupported media type response has a 3xx status code
func (o *GetDeliveryChallanDocumentUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get delivery challan document unsupported media type response has a 4xx status code
func (o *GetDeliveryChallanDocumentUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get delivery challan document unsupported media type response has a 5xx status code
func (o *GetDeliveryChallanDocumentUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get delivery challan document unsupported media type response a status code equal to that given
func (o *GetDeliveryChallanDocumentUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetDeliveryChallanDocumentUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/deliveryChallanDocument][%d] getDeliveryChallanDocumentUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetDeliveryChallanDocumentUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/deliveryChallanDocument][%d] getDeliveryChallanDocumentUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetDeliveryChallanDocumentUnsupportedMediaType) GetPayload() *fulfillment_inboundv2024_03_20_models.ErrorList {
	return o.Payload
}

func (o *GetDeliveryChallanDocumentUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(fulfillment_inboundv2024_03_20_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDeliveryChallanDocumentTooManyRequests creates a GetDeliveryChallanDocumentTooManyRequests with default headers values
func NewGetDeliveryChallanDocumentTooManyRequests() *GetDeliveryChallanDocumentTooManyRequests {
	return &GetDeliveryChallanDocumentTooManyRequests{}
}

/*
GetDeliveryChallanDocumentTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type GetDeliveryChallanDocumentTooManyRequests struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inboundv2024_03_20_models.ErrorList
}

// IsSuccess returns true when this get delivery challan document too many requests response has a 2xx status code
func (o *GetDeliveryChallanDocumentTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get delivery challan document too many requests response has a 3xx status code
func (o *GetDeliveryChallanDocumentTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get delivery challan document too many requests response has a 4xx status code
func (o *GetDeliveryChallanDocumentTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get delivery challan document too many requests response has a 5xx status code
func (o *GetDeliveryChallanDocumentTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get delivery challan document too many requests response a status code equal to that given
func (o *GetDeliveryChallanDocumentTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetDeliveryChallanDocumentTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/deliveryChallanDocument][%d] getDeliveryChallanDocumentTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetDeliveryChallanDocumentTooManyRequests) String() string {
	return fmt.Sprintf("[GET /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/deliveryChallanDocument][%d] getDeliveryChallanDocumentTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetDeliveryChallanDocumentTooManyRequests) GetPayload() *fulfillment_inboundv2024_03_20_models.ErrorList {
	return o.Payload
}

func (o *GetDeliveryChallanDocumentTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(fulfillment_inboundv2024_03_20_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDeliveryChallanDocumentInternalServerError creates a GetDeliveryChallanDocumentInternalServerError with default headers values
func NewGetDeliveryChallanDocumentInternalServerError() *GetDeliveryChallanDocumentInternalServerError {
	return &GetDeliveryChallanDocumentInternalServerError{}
}

/*
GetDeliveryChallanDocumentInternalServerError describes a response with status code 500, with default header values.

An unexpected condition occurred that prevented the server from fulfilling the request.
*/
type GetDeliveryChallanDocumentInternalServerError struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inboundv2024_03_20_models.ErrorList
}

// IsSuccess returns true when this get delivery challan document internal server error response has a 2xx status code
func (o *GetDeliveryChallanDocumentInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get delivery challan document internal server error response has a 3xx status code
func (o *GetDeliveryChallanDocumentInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get delivery challan document internal server error response has a 4xx status code
func (o *GetDeliveryChallanDocumentInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get delivery challan document internal server error response has a 5xx status code
func (o *GetDeliveryChallanDocumentInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get delivery challan document internal server error response a status code equal to that given
func (o *GetDeliveryChallanDocumentInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetDeliveryChallanDocumentInternalServerError) Error() string {
	return fmt.Sprintf("[GET /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/deliveryChallanDocument][%d] getDeliveryChallanDocumentInternalServerError  %+v", 500, o.Payload)
}

func (o *GetDeliveryChallanDocumentInternalServerError) String() string {
	return fmt.Sprintf("[GET /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/deliveryChallanDocument][%d] getDeliveryChallanDocumentInternalServerError  %+v", 500, o.Payload)
}

func (o *GetDeliveryChallanDocumentInternalServerError) GetPayload() *fulfillment_inboundv2024_03_20_models.ErrorList {
	return o.Payload
}

func (o *GetDeliveryChallanDocumentInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(fulfillment_inboundv2024_03_20_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDeliveryChallanDocumentServiceUnavailable creates a GetDeliveryChallanDocumentServiceUnavailable with default headers values
func NewGetDeliveryChallanDocumentServiceUnavailable() *GetDeliveryChallanDocumentServiceUnavailable {
	return &GetDeliveryChallanDocumentServiceUnavailable{}
}

/*
GetDeliveryChallanDocumentServiceUnavailable describes a response with status code 503, with default header values.

Temporary overloading or maintenance of the server.
*/
type GetDeliveryChallanDocumentServiceUnavailable struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inboundv2024_03_20_models.ErrorList
}

// IsSuccess returns true when this get delivery challan document service unavailable response has a 2xx status code
func (o *GetDeliveryChallanDocumentServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get delivery challan document service unavailable response has a 3xx status code
func (o *GetDeliveryChallanDocumentServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get delivery challan document service unavailable response has a 4xx status code
func (o *GetDeliveryChallanDocumentServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get delivery challan document service unavailable response has a 5xx status code
func (o *GetDeliveryChallanDocumentServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get delivery challan document service unavailable response a status code equal to that given
func (o *GetDeliveryChallanDocumentServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetDeliveryChallanDocumentServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/deliveryChallanDocument][%d] getDeliveryChallanDocumentServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetDeliveryChallanDocumentServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/deliveryChallanDocument][%d] getDeliveryChallanDocumentServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetDeliveryChallanDocumentServiceUnavailable) GetPayload() *fulfillment_inboundv2024_03_20_models.ErrorList {
	return o.Payload
}

func (o *GetDeliveryChallanDocumentServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(fulfillment_inboundv2024_03_20_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
