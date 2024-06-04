// Code generated by go-swagger; DO NOT EDIT.

package fba_inbound

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/renabled/amzn-sp-api-go/api/fulfillmentInbound_2024-03-20/fulfillment_inbound_2024_03_20_models"
)

// ListShipmentContentUpdatePreviewsReader is a Reader for the ListShipmentContentUpdatePreviews structure.
type ListShipmentContentUpdatePreviewsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListShipmentContentUpdatePreviewsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListShipmentContentUpdatePreviewsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListShipmentContentUpdatePreviewsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListShipmentContentUpdatePreviewsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListShipmentContentUpdatePreviewsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewListShipmentContentUpdatePreviewsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewListShipmentContentUpdatePreviewsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewListShipmentContentUpdatePreviewsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListShipmentContentUpdatePreviewsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewListShipmentContentUpdatePreviewsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListShipmentContentUpdatePreviewsOK creates a ListShipmentContentUpdatePreviewsOK with default headers values
func NewListShipmentContentUpdatePreviewsOK() *ListShipmentContentUpdatePreviewsOK {
	return &ListShipmentContentUpdatePreviewsOK{}
}

/*
ListShipmentContentUpdatePreviewsOK describes a response with status code 200, with default header values.

ListShipmentContentUpdatePreviews 200 response
*/
type ListShipmentContentUpdatePreviewsOK struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inbound_2024_03_20_models.ListShipmentContentUpdatePreviewsResponse
}

// IsSuccess returns true when this list shipment content update previews o k response has a 2xx status code
func (o *ListShipmentContentUpdatePreviewsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list shipment content update previews o k response has a 3xx status code
func (o *ListShipmentContentUpdatePreviewsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list shipment content update previews o k response has a 4xx status code
func (o *ListShipmentContentUpdatePreviewsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list shipment content update previews o k response has a 5xx status code
func (o *ListShipmentContentUpdatePreviewsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list shipment content update previews o k response a status code equal to that given
func (o *ListShipmentContentUpdatePreviewsOK) IsCode(code int) bool {
	return code == 200
}

func (o *ListShipmentContentUpdatePreviewsOK) Error() string {
	return fmt.Sprintf("[GET /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/contentUpdatePreviews][%d] listShipmentContentUpdatePreviewsOK  %+v", 200, o.Payload)
}

func (o *ListShipmentContentUpdatePreviewsOK) String() string {
	return fmt.Sprintf("[GET /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/contentUpdatePreviews][%d] listShipmentContentUpdatePreviewsOK  %+v", 200, o.Payload)
}

func (o *ListShipmentContentUpdatePreviewsOK) GetPayload() *fulfillment_inbound_2024_03_20_models.ListShipmentContentUpdatePreviewsResponse {
	return o.Payload
}

func (o *ListShipmentContentUpdatePreviewsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fulfillment_inbound_2024_03_20_models.ListShipmentContentUpdatePreviewsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListShipmentContentUpdatePreviewsBadRequest creates a ListShipmentContentUpdatePreviewsBadRequest with default headers values
func NewListShipmentContentUpdatePreviewsBadRequest() *ListShipmentContentUpdatePreviewsBadRequest {
	return &ListShipmentContentUpdatePreviewsBadRequest{}
}

/*
ListShipmentContentUpdatePreviewsBadRequest describes a response with status code 400, with default header values.

Request has missing or invalid parameters and cannot be parsed.
*/
type ListShipmentContentUpdatePreviewsBadRequest struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inbound_2024_03_20_models.ErrorList
}

// IsSuccess returns true when this list shipment content update previews bad request response has a 2xx status code
func (o *ListShipmentContentUpdatePreviewsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list shipment content update previews bad request response has a 3xx status code
func (o *ListShipmentContentUpdatePreviewsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list shipment content update previews bad request response has a 4xx status code
func (o *ListShipmentContentUpdatePreviewsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this list shipment content update previews bad request response has a 5xx status code
func (o *ListShipmentContentUpdatePreviewsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this list shipment content update previews bad request response a status code equal to that given
func (o *ListShipmentContentUpdatePreviewsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *ListShipmentContentUpdatePreviewsBadRequest) Error() string {
	return fmt.Sprintf("[GET /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/contentUpdatePreviews][%d] listShipmentContentUpdatePreviewsBadRequest  %+v", 400, o.Payload)
}

func (o *ListShipmentContentUpdatePreviewsBadRequest) String() string {
	return fmt.Sprintf("[GET /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/contentUpdatePreviews][%d] listShipmentContentUpdatePreviewsBadRequest  %+v", 400, o.Payload)
}

func (o *ListShipmentContentUpdatePreviewsBadRequest) GetPayload() *fulfillment_inbound_2024_03_20_models.ErrorList {
	return o.Payload
}

func (o *ListShipmentContentUpdatePreviewsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fulfillment_inbound_2024_03_20_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListShipmentContentUpdatePreviewsForbidden creates a ListShipmentContentUpdatePreviewsForbidden with default headers values
func NewListShipmentContentUpdatePreviewsForbidden() *ListShipmentContentUpdatePreviewsForbidden {
	return &ListShipmentContentUpdatePreviewsForbidden{}
}

/*
ListShipmentContentUpdatePreviewsForbidden describes a response with status code 403, with default header values.

Indicates that access to the resource is forbidden. Possible reasons include Access Denied, Unauthorized, Expired Token, or Invalid Signature.
*/
type ListShipmentContentUpdatePreviewsForbidden struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inbound_2024_03_20_models.ErrorList
}

// IsSuccess returns true when this list shipment content update previews forbidden response has a 2xx status code
func (o *ListShipmentContentUpdatePreviewsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list shipment content update previews forbidden response has a 3xx status code
func (o *ListShipmentContentUpdatePreviewsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list shipment content update previews forbidden response has a 4xx status code
func (o *ListShipmentContentUpdatePreviewsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this list shipment content update previews forbidden response has a 5xx status code
func (o *ListShipmentContentUpdatePreviewsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this list shipment content update previews forbidden response a status code equal to that given
func (o *ListShipmentContentUpdatePreviewsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ListShipmentContentUpdatePreviewsForbidden) Error() string {
	return fmt.Sprintf("[GET /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/contentUpdatePreviews][%d] listShipmentContentUpdatePreviewsForbidden  %+v", 403, o.Payload)
}

func (o *ListShipmentContentUpdatePreviewsForbidden) String() string {
	return fmt.Sprintf("[GET /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/contentUpdatePreviews][%d] listShipmentContentUpdatePreviewsForbidden  %+v", 403, o.Payload)
}

func (o *ListShipmentContentUpdatePreviewsForbidden) GetPayload() *fulfillment_inbound_2024_03_20_models.ErrorList {
	return o.Payload
}

func (o *ListShipmentContentUpdatePreviewsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(fulfillment_inbound_2024_03_20_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListShipmentContentUpdatePreviewsNotFound creates a ListShipmentContentUpdatePreviewsNotFound with default headers values
func NewListShipmentContentUpdatePreviewsNotFound() *ListShipmentContentUpdatePreviewsNotFound {
	return &ListShipmentContentUpdatePreviewsNotFound{}
}

/*
ListShipmentContentUpdatePreviewsNotFound describes a response with status code 404, with default header values.

The resource specified does not exist.
*/
type ListShipmentContentUpdatePreviewsNotFound struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inbound_2024_03_20_models.ErrorList
}

// IsSuccess returns true when this list shipment content update previews not found response has a 2xx status code
func (o *ListShipmentContentUpdatePreviewsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list shipment content update previews not found response has a 3xx status code
func (o *ListShipmentContentUpdatePreviewsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list shipment content update previews not found response has a 4xx status code
func (o *ListShipmentContentUpdatePreviewsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this list shipment content update previews not found response has a 5xx status code
func (o *ListShipmentContentUpdatePreviewsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this list shipment content update previews not found response a status code equal to that given
func (o *ListShipmentContentUpdatePreviewsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *ListShipmentContentUpdatePreviewsNotFound) Error() string {
	return fmt.Sprintf("[GET /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/contentUpdatePreviews][%d] listShipmentContentUpdatePreviewsNotFound  %+v", 404, o.Payload)
}

func (o *ListShipmentContentUpdatePreviewsNotFound) String() string {
	return fmt.Sprintf("[GET /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/contentUpdatePreviews][%d] listShipmentContentUpdatePreviewsNotFound  %+v", 404, o.Payload)
}

func (o *ListShipmentContentUpdatePreviewsNotFound) GetPayload() *fulfillment_inbound_2024_03_20_models.ErrorList {
	return o.Payload
}

func (o *ListShipmentContentUpdatePreviewsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fulfillment_inbound_2024_03_20_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListShipmentContentUpdatePreviewsRequestEntityTooLarge creates a ListShipmentContentUpdatePreviewsRequestEntityTooLarge with default headers values
func NewListShipmentContentUpdatePreviewsRequestEntityTooLarge() *ListShipmentContentUpdatePreviewsRequestEntityTooLarge {
	return &ListShipmentContentUpdatePreviewsRequestEntityTooLarge{}
}

/*
ListShipmentContentUpdatePreviewsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request size exceeded the maximum accepted size.
*/
type ListShipmentContentUpdatePreviewsRequestEntityTooLarge struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inbound_2024_03_20_models.ErrorList
}

// IsSuccess returns true when this list shipment content update previews request entity too large response has a 2xx status code
func (o *ListShipmentContentUpdatePreviewsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list shipment content update previews request entity too large response has a 3xx status code
func (o *ListShipmentContentUpdatePreviewsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list shipment content update previews request entity too large response has a 4xx status code
func (o *ListShipmentContentUpdatePreviewsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this list shipment content update previews request entity too large response has a 5xx status code
func (o *ListShipmentContentUpdatePreviewsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this list shipment content update previews request entity too large response a status code equal to that given
func (o *ListShipmentContentUpdatePreviewsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *ListShipmentContentUpdatePreviewsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/contentUpdatePreviews][%d] listShipmentContentUpdatePreviewsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *ListShipmentContentUpdatePreviewsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/contentUpdatePreviews][%d] listShipmentContentUpdatePreviewsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *ListShipmentContentUpdatePreviewsRequestEntityTooLarge) GetPayload() *fulfillment_inbound_2024_03_20_models.ErrorList {
	return o.Payload
}

func (o *ListShipmentContentUpdatePreviewsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(fulfillment_inbound_2024_03_20_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListShipmentContentUpdatePreviewsUnsupportedMediaType creates a ListShipmentContentUpdatePreviewsUnsupportedMediaType with default headers values
func NewListShipmentContentUpdatePreviewsUnsupportedMediaType() *ListShipmentContentUpdatePreviewsUnsupportedMediaType {
	return &ListShipmentContentUpdatePreviewsUnsupportedMediaType{}
}

/*
ListShipmentContentUpdatePreviewsUnsupportedMediaType describes a response with status code 415, with default header values.

The request payload is in an unsupported format.
*/
type ListShipmentContentUpdatePreviewsUnsupportedMediaType struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inbound_2024_03_20_models.ErrorList
}

// IsSuccess returns true when this list shipment content update previews unsupported media type response has a 2xx status code
func (o *ListShipmentContentUpdatePreviewsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list shipment content update previews unsupported media type response has a 3xx status code
func (o *ListShipmentContentUpdatePreviewsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list shipment content update previews unsupported media type response has a 4xx status code
func (o *ListShipmentContentUpdatePreviewsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this list shipment content update previews unsupported media type response has a 5xx status code
func (o *ListShipmentContentUpdatePreviewsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this list shipment content update previews unsupported media type response a status code equal to that given
func (o *ListShipmentContentUpdatePreviewsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *ListShipmentContentUpdatePreviewsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/contentUpdatePreviews][%d] listShipmentContentUpdatePreviewsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *ListShipmentContentUpdatePreviewsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/contentUpdatePreviews][%d] listShipmentContentUpdatePreviewsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *ListShipmentContentUpdatePreviewsUnsupportedMediaType) GetPayload() *fulfillment_inbound_2024_03_20_models.ErrorList {
	return o.Payload
}

func (o *ListShipmentContentUpdatePreviewsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(fulfillment_inbound_2024_03_20_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListShipmentContentUpdatePreviewsTooManyRequests creates a ListShipmentContentUpdatePreviewsTooManyRequests with default headers values
func NewListShipmentContentUpdatePreviewsTooManyRequests() *ListShipmentContentUpdatePreviewsTooManyRequests {
	return &ListShipmentContentUpdatePreviewsTooManyRequests{}
}

/*
ListShipmentContentUpdatePreviewsTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type ListShipmentContentUpdatePreviewsTooManyRequests struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inbound_2024_03_20_models.ErrorList
}

// IsSuccess returns true when this list shipment content update previews too many requests response has a 2xx status code
func (o *ListShipmentContentUpdatePreviewsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list shipment content update previews too many requests response has a 3xx status code
func (o *ListShipmentContentUpdatePreviewsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list shipment content update previews too many requests response has a 4xx status code
func (o *ListShipmentContentUpdatePreviewsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this list shipment content update previews too many requests response has a 5xx status code
func (o *ListShipmentContentUpdatePreviewsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this list shipment content update previews too many requests response a status code equal to that given
func (o *ListShipmentContentUpdatePreviewsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *ListShipmentContentUpdatePreviewsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/contentUpdatePreviews][%d] listShipmentContentUpdatePreviewsTooManyRequests  %+v", 429, o.Payload)
}

func (o *ListShipmentContentUpdatePreviewsTooManyRequests) String() string {
	return fmt.Sprintf("[GET /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/contentUpdatePreviews][%d] listShipmentContentUpdatePreviewsTooManyRequests  %+v", 429, o.Payload)
}

func (o *ListShipmentContentUpdatePreviewsTooManyRequests) GetPayload() *fulfillment_inbound_2024_03_20_models.ErrorList {
	return o.Payload
}

func (o *ListShipmentContentUpdatePreviewsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(fulfillment_inbound_2024_03_20_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListShipmentContentUpdatePreviewsInternalServerError creates a ListShipmentContentUpdatePreviewsInternalServerError with default headers values
func NewListShipmentContentUpdatePreviewsInternalServerError() *ListShipmentContentUpdatePreviewsInternalServerError {
	return &ListShipmentContentUpdatePreviewsInternalServerError{}
}

/*
ListShipmentContentUpdatePreviewsInternalServerError describes a response with status code 500, with default header values.

An unexpected condition occurred that prevented the server from fulfilling the request.
*/
type ListShipmentContentUpdatePreviewsInternalServerError struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inbound_2024_03_20_models.ErrorList
}

// IsSuccess returns true when this list shipment content update previews internal server error response has a 2xx status code
func (o *ListShipmentContentUpdatePreviewsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list shipment content update previews internal server error response has a 3xx status code
func (o *ListShipmentContentUpdatePreviewsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list shipment content update previews internal server error response has a 4xx status code
func (o *ListShipmentContentUpdatePreviewsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this list shipment content update previews internal server error response has a 5xx status code
func (o *ListShipmentContentUpdatePreviewsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this list shipment content update previews internal server error response a status code equal to that given
func (o *ListShipmentContentUpdatePreviewsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ListShipmentContentUpdatePreviewsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/contentUpdatePreviews][%d] listShipmentContentUpdatePreviewsInternalServerError  %+v", 500, o.Payload)
}

func (o *ListShipmentContentUpdatePreviewsInternalServerError) String() string {
	return fmt.Sprintf("[GET /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/contentUpdatePreviews][%d] listShipmentContentUpdatePreviewsInternalServerError  %+v", 500, o.Payload)
}

func (o *ListShipmentContentUpdatePreviewsInternalServerError) GetPayload() *fulfillment_inbound_2024_03_20_models.ErrorList {
	return o.Payload
}

func (o *ListShipmentContentUpdatePreviewsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(fulfillment_inbound_2024_03_20_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListShipmentContentUpdatePreviewsServiceUnavailable creates a ListShipmentContentUpdatePreviewsServiceUnavailable with default headers values
func NewListShipmentContentUpdatePreviewsServiceUnavailable() *ListShipmentContentUpdatePreviewsServiceUnavailable {
	return &ListShipmentContentUpdatePreviewsServiceUnavailable{}
}

/*
ListShipmentContentUpdatePreviewsServiceUnavailable describes a response with status code 503, with default header values.

Temporary overloading or maintenance of the server.
*/
type ListShipmentContentUpdatePreviewsServiceUnavailable struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inbound_2024_03_20_models.ErrorList
}

// IsSuccess returns true when this list shipment content update previews service unavailable response has a 2xx status code
func (o *ListShipmentContentUpdatePreviewsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list shipment content update previews service unavailable response has a 3xx status code
func (o *ListShipmentContentUpdatePreviewsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list shipment content update previews service unavailable response has a 4xx status code
func (o *ListShipmentContentUpdatePreviewsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this list shipment content update previews service unavailable response has a 5xx status code
func (o *ListShipmentContentUpdatePreviewsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this list shipment content update previews service unavailable response a status code equal to that given
func (o *ListShipmentContentUpdatePreviewsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *ListShipmentContentUpdatePreviewsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/contentUpdatePreviews][%d] listShipmentContentUpdatePreviewsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *ListShipmentContentUpdatePreviewsServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/contentUpdatePreviews][%d] listShipmentContentUpdatePreviewsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *ListShipmentContentUpdatePreviewsServiceUnavailable) GetPayload() *fulfillment_inbound_2024_03_20_models.ErrorList {
	return o.Payload
}

func (o *ListShipmentContentUpdatePreviewsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(fulfillment_inbound_2024_03_20_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
