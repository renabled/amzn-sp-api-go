// Code generated by go-swagger; DO NOT EDIT.

package vendor_d_f_sandbox

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/xamandar/amzn-sp-api-go/api/vendorDirectFulfillmentSandboxData_2021-10-28/vendor_direct_fulfillment_sandbox_data_2021_10_28_models"
)

// GenerateOrderScenariosReader is a Reader for the GenerateOrderScenarios structure.
type GenerateOrderScenariosReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GenerateOrderScenariosReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewGenerateOrderScenariosAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGenerateOrderScenariosBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGenerateOrderScenariosForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGenerateOrderScenariosNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGenerateOrderScenariosRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGenerateOrderScenariosUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGenerateOrderScenariosTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGenerateOrderScenariosInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGenerateOrderScenariosServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGenerateOrderScenariosAccepted creates a GenerateOrderScenariosAccepted with default headers values
func NewGenerateOrderScenariosAccepted() *GenerateOrderScenariosAccepted {
	return &GenerateOrderScenariosAccepted{}
}

/* GenerateOrderScenariosAccepted describes a response with status code 202, with default header values.

Success.
*/
type GenerateOrderScenariosAccepted struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_sandbox_data_2021_10_28_models.TransactionReference
}

func (o *GenerateOrderScenariosAccepted) Error() string {
	return fmt.Sprintf("[POST /vendor/directFulfillment/sandbox/2021-10-28/orders][%d] generateOrderScenariosAccepted  %+v", 202, o.Payload)
}
func (o *GenerateOrderScenariosAccepted) GetPayload() *vendor_direct_fulfillment_sandbox_data_2021_10_28_models.TransactionReference {
	return o.Payload
}

func (o *GenerateOrderScenariosAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(vendor_direct_fulfillment_sandbox_data_2021_10_28_models.TransactionReference)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGenerateOrderScenariosBadRequest creates a GenerateOrderScenariosBadRequest with default headers values
func NewGenerateOrderScenariosBadRequest() *GenerateOrderScenariosBadRequest {
	return &GenerateOrderScenariosBadRequest{}
}

/* GenerateOrderScenariosBadRequest describes a response with status code 400, with default header values.

Request has missing or invalid parameters and cannot be parsed.
*/
type GenerateOrderScenariosBadRequest struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_sandbox_data_2021_10_28_models.ErrorList
}

func (o *GenerateOrderScenariosBadRequest) Error() string {
	return fmt.Sprintf("[POST /vendor/directFulfillment/sandbox/2021-10-28/orders][%d] generateOrderScenariosBadRequest  %+v", 400, o.Payload)
}
func (o *GenerateOrderScenariosBadRequest) GetPayload() *vendor_direct_fulfillment_sandbox_data_2021_10_28_models.ErrorList {
	return o.Payload
}

func (o *GenerateOrderScenariosBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(vendor_direct_fulfillment_sandbox_data_2021_10_28_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGenerateOrderScenariosForbidden creates a GenerateOrderScenariosForbidden with default headers values
func NewGenerateOrderScenariosForbidden() *GenerateOrderScenariosForbidden {
	return &GenerateOrderScenariosForbidden{}
}

/* GenerateOrderScenariosForbidden describes a response with status code 403, with default header values.

Indicates that access to the resource is forbidden. Possible reasons include Access Denied, Unauthorized, Expired Token, or Invalid Signature.
*/
type GenerateOrderScenariosForbidden struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_sandbox_data_2021_10_28_models.ErrorList
}

func (o *GenerateOrderScenariosForbidden) Error() string {
	return fmt.Sprintf("[POST /vendor/directFulfillment/sandbox/2021-10-28/orders][%d] generateOrderScenariosForbidden  %+v", 403, o.Payload)
}
func (o *GenerateOrderScenariosForbidden) GetPayload() *vendor_direct_fulfillment_sandbox_data_2021_10_28_models.ErrorList {
	return o.Payload
}

func (o *GenerateOrderScenariosForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(vendor_direct_fulfillment_sandbox_data_2021_10_28_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGenerateOrderScenariosNotFound creates a GenerateOrderScenariosNotFound with default headers values
func NewGenerateOrderScenariosNotFound() *GenerateOrderScenariosNotFound {
	return &GenerateOrderScenariosNotFound{}
}

/* GenerateOrderScenariosNotFound describes a response with status code 404, with default header values.

The resource specified does not exist.
*/
type GenerateOrderScenariosNotFound struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_sandbox_data_2021_10_28_models.ErrorList
}

func (o *GenerateOrderScenariosNotFound) Error() string {
	return fmt.Sprintf("[POST /vendor/directFulfillment/sandbox/2021-10-28/orders][%d] generateOrderScenariosNotFound  %+v", 404, o.Payload)
}
func (o *GenerateOrderScenariosNotFound) GetPayload() *vendor_direct_fulfillment_sandbox_data_2021_10_28_models.ErrorList {
	return o.Payload
}

func (o *GenerateOrderScenariosNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(vendor_direct_fulfillment_sandbox_data_2021_10_28_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGenerateOrderScenariosRequestEntityTooLarge creates a GenerateOrderScenariosRequestEntityTooLarge with default headers values
func NewGenerateOrderScenariosRequestEntityTooLarge() *GenerateOrderScenariosRequestEntityTooLarge {
	return &GenerateOrderScenariosRequestEntityTooLarge{}
}

/* GenerateOrderScenariosRequestEntityTooLarge describes a response with status code 413, with default header values.

The request size exceeded the maximum accepted size.
*/
type GenerateOrderScenariosRequestEntityTooLarge struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_sandbox_data_2021_10_28_models.ErrorList
}

func (o *GenerateOrderScenariosRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /vendor/directFulfillment/sandbox/2021-10-28/orders][%d] generateOrderScenariosRequestEntityTooLarge  %+v", 413, o.Payload)
}
func (o *GenerateOrderScenariosRequestEntityTooLarge) GetPayload() *vendor_direct_fulfillment_sandbox_data_2021_10_28_models.ErrorList {
	return o.Payload
}

func (o *GenerateOrderScenariosRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(vendor_direct_fulfillment_sandbox_data_2021_10_28_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGenerateOrderScenariosUnsupportedMediaType creates a GenerateOrderScenariosUnsupportedMediaType with default headers values
func NewGenerateOrderScenariosUnsupportedMediaType() *GenerateOrderScenariosUnsupportedMediaType {
	return &GenerateOrderScenariosUnsupportedMediaType{}
}

/* GenerateOrderScenariosUnsupportedMediaType describes a response with status code 415, with default header values.

The request payload is in an unsupported format.
*/
type GenerateOrderScenariosUnsupportedMediaType struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_sandbox_data_2021_10_28_models.ErrorList
}

func (o *GenerateOrderScenariosUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /vendor/directFulfillment/sandbox/2021-10-28/orders][%d] generateOrderScenariosUnsupportedMediaType  %+v", 415, o.Payload)
}
func (o *GenerateOrderScenariosUnsupportedMediaType) GetPayload() *vendor_direct_fulfillment_sandbox_data_2021_10_28_models.ErrorList {
	return o.Payload
}

func (o *GenerateOrderScenariosUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(vendor_direct_fulfillment_sandbox_data_2021_10_28_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGenerateOrderScenariosTooManyRequests creates a GenerateOrderScenariosTooManyRequests with default headers values
func NewGenerateOrderScenariosTooManyRequests() *GenerateOrderScenariosTooManyRequests {
	return &GenerateOrderScenariosTooManyRequests{}
}

/* GenerateOrderScenariosTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type GenerateOrderScenariosTooManyRequests struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_sandbox_data_2021_10_28_models.ErrorList
}

func (o *GenerateOrderScenariosTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /vendor/directFulfillment/sandbox/2021-10-28/orders][%d] generateOrderScenariosTooManyRequests  %+v", 429, o.Payload)
}
func (o *GenerateOrderScenariosTooManyRequests) GetPayload() *vendor_direct_fulfillment_sandbox_data_2021_10_28_models.ErrorList {
	return o.Payload
}

func (o *GenerateOrderScenariosTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(vendor_direct_fulfillment_sandbox_data_2021_10_28_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGenerateOrderScenariosInternalServerError creates a GenerateOrderScenariosInternalServerError with default headers values
func NewGenerateOrderScenariosInternalServerError() *GenerateOrderScenariosInternalServerError {
	return &GenerateOrderScenariosInternalServerError{}
}

/* GenerateOrderScenariosInternalServerError describes a response with status code 500, with default header values.

An unexpected condition occurred that prevented the server from fulfilling the request.
*/
type GenerateOrderScenariosInternalServerError struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_sandbox_data_2021_10_28_models.ErrorList
}

func (o *GenerateOrderScenariosInternalServerError) Error() string {
	return fmt.Sprintf("[POST /vendor/directFulfillment/sandbox/2021-10-28/orders][%d] generateOrderScenariosInternalServerError  %+v", 500, o.Payload)
}
func (o *GenerateOrderScenariosInternalServerError) GetPayload() *vendor_direct_fulfillment_sandbox_data_2021_10_28_models.ErrorList {
	return o.Payload
}

func (o *GenerateOrderScenariosInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(vendor_direct_fulfillment_sandbox_data_2021_10_28_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGenerateOrderScenariosServiceUnavailable creates a GenerateOrderScenariosServiceUnavailable with default headers values
func NewGenerateOrderScenariosServiceUnavailable() *GenerateOrderScenariosServiceUnavailable {
	return &GenerateOrderScenariosServiceUnavailable{}
}

/* GenerateOrderScenariosServiceUnavailable describes a response with status code 503, with default header values.

Temporary overloading or maintenance of the server.
*/
type GenerateOrderScenariosServiceUnavailable struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *vendor_direct_fulfillment_sandbox_data_2021_10_28_models.ErrorList
}

func (o *GenerateOrderScenariosServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /vendor/directFulfillment/sandbox/2021-10-28/orders][%d] generateOrderScenariosServiceUnavailable  %+v", 503, o.Payload)
}
func (o *GenerateOrderScenariosServiceUnavailable) GetPayload() *vendor_direct_fulfillment_sandbox_data_2021_10_28_models.ErrorList {
	return o.Payload
}

func (o *GenerateOrderScenariosServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(vendor_direct_fulfillment_sandbox_data_2021_10_28_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
