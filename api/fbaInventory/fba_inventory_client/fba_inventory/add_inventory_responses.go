// Code generated by go-swagger; DO NOT EDIT.

package fba_inventory

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/renabled/amzn-sp-api-go/api/fbaInventory/fba_inventory_models"
)

// AddInventoryReader is a Reader for the AddInventory structure.
type AddInventoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddInventoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddInventoryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAddInventoryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAddInventoryForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAddInventoryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewAddInventoryTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAddInventoryInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewAddInventoryServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAddInventoryOK creates a AddInventoryOK with default headers values
func NewAddInventoryOK() *AddInventoryOK {
	return &AddInventoryOK{}
}

/*
AddInventoryOK describes a response with status code 200, with default header values.

OK
*/
type AddInventoryOK struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fba_inventory_models.AddInventoryResponse
}

// IsSuccess returns true when this add inventory o k response has a 2xx status code
func (o *AddInventoryOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this add inventory o k response has a 3xx status code
func (o *AddInventoryOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add inventory o k response has a 4xx status code
func (o *AddInventoryOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this add inventory o k response has a 5xx status code
func (o *AddInventoryOK) IsServerError() bool {
	return false
}

// IsCode returns true when this add inventory o k response a status code equal to that given
func (o *AddInventoryOK) IsCode(code int) bool {
	return code == 200
}

func (o *AddInventoryOK) Error() string {
	return fmt.Sprintf("[POST /fba/inventory/v1/items/inventory][%d] addInventoryOK  %+v", 200, o.Payload)
}

func (o *AddInventoryOK) String() string {
	return fmt.Sprintf("[POST /fba/inventory/v1/items/inventory][%d] addInventoryOK  %+v", 200, o.Payload)
}

func (o *AddInventoryOK) GetPayload() *fba_inventory_models.AddInventoryResponse {
	return o.Payload
}

func (o *AddInventoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(fba_inventory_models.AddInventoryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddInventoryBadRequest creates a AddInventoryBadRequest with default headers values
func NewAddInventoryBadRequest() *AddInventoryBadRequest {
	return &AddInventoryBadRequest{}
}

/*
AddInventoryBadRequest describes a response with status code 400, with default header values.

Request has missing or invalid parameters and cannot be parsed.
*/
type AddInventoryBadRequest struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fba_inventory_models.AddInventoryResponse
}

// IsSuccess returns true when this add inventory bad request response has a 2xx status code
func (o *AddInventoryBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add inventory bad request response has a 3xx status code
func (o *AddInventoryBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add inventory bad request response has a 4xx status code
func (o *AddInventoryBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this add inventory bad request response has a 5xx status code
func (o *AddInventoryBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this add inventory bad request response a status code equal to that given
func (o *AddInventoryBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *AddInventoryBadRequest) Error() string {
	return fmt.Sprintf("[POST /fba/inventory/v1/items/inventory][%d] addInventoryBadRequest  %+v", 400, o.Payload)
}

func (o *AddInventoryBadRequest) String() string {
	return fmt.Sprintf("[POST /fba/inventory/v1/items/inventory][%d] addInventoryBadRequest  %+v", 400, o.Payload)
}

func (o *AddInventoryBadRequest) GetPayload() *fba_inventory_models.AddInventoryResponse {
	return o.Payload
}

func (o *AddInventoryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(fba_inventory_models.AddInventoryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddInventoryForbidden creates a AddInventoryForbidden with default headers values
func NewAddInventoryForbidden() *AddInventoryForbidden {
	return &AddInventoryForbidden{}
}

/*
AddInventoryForbidden describes a response with status code 403, with default header values.

Indicates access to the resource is forbidden. Possible reasons include Access Denied, Unauthorized, Expired Token, Invalid Signature or Resource Not Found.
*/
type AddInventoryForbidden struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fba_inventory_models.AddInventoryResponse
}

// IsSuccess returns true when this add inventory forbidden response has a 2xx status code
func (o *AddInventoryForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add inventory forbidden response has a 3xx status code
func (o *AddInventoryForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add inventory forbidden response has a 4xx status code
func (o *AddInventoryForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this add inventory forbidden response has a 5xx status code
func (o *AddInventoryForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this add inventory forbidden response a status code equal to that given
func (o *AddInventoryForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *AddInventoryForbidden) Error() string {
	return fmt.Sprintf("[POST /fba/inventory/v1/items/inventory][%d] addInventoryForbidden  %+v", 403, o.Payload)
}

func (o *AddInventoryForbidden) String() string {
	return fmt.Sprintf("[POST /fba/inventory/v1/items/inventory][%d] addInventoryForbidden  %+v", 403, o.Payload)
}

func (o *AddInventoryForbidden) GetPayload() *fba_inventory_models.AddInventoryResponse {
	return o.Payload
}

func (o *AddInventoryForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(fba_inventory_models.AddInventoryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddInventoryNotFound creates a AddInventoryNotFound with default headers values
func NewAddInventoryNotFound() *AddInventoryNotFound {
	return &AddInventoryNotFound{}
}

/*
AddInventoryNotFound describes a response with status code 404, with default header values.

The specified resource does not exist.
*/
type AddInventoryNotFound struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fba_inventory_models.AddInventoryResponse
}

// IsSuccess returns true when this add inventory not found response has a 2xx status code
func (o *AddInventoryNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add inventory not found response has a 3xx status code
func (o *AddInventoryNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add inventory not found response has a 4xx status code
func (o *AddInventoryNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this add inventory not found response has a 5xx status code
func (o *AddInventoryNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this add inventory not found response a status code equal to that given
func (o *AddInventoryNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *AddInventoryNotFound) Error() string {
	return fmt.Sprintf("[POST /fba/inventory/v1/items/inventory][%d] addInventoryNotFound  %+v", 404, o.Payload)
}

func (o *AddInventoryNotFound) String() string {
	return fmt.Sprintf("[POST /fba/inventory/v1/items/inventory][%d] addInventoryNotFound  %+v", 404, o.Payload)
}

func (o *AddInventoryNotFound) GetPayload() *fba_inventory_models.AddInventoryResponse {
	return o.Payload
}

func (o *AddInventoryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(fba_inventory_models.AddInventoryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddInventoryTooManyRequests creates a AddInventoryTooManyRequests with default headers values
func NewAddInventoryTooManyRequests() *AddInventoryTooManyRequests {
	return &AddInventoryTooManyRequests{}
}

/*
AddInventoryTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type AddInventoryTooManyRequests struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fba_inventory_models.AddInventoryResponse
}

// IsSuccess returns true when this add inventory too many requests response has a 2xx status code
func (o *AddInventoryTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add inventory too many requests response has a 3xx status code
func (o *AddInventoryTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add inventory too many requests response has a 4xx status code
func (o *AddInventoryTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this add inventory too many requests response has a 5xx status code
func (o *AddInventoryTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this add inventory too many requests response a status code equal to that given
func (o *AddInventoryTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *AddInventoryTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /fba/inventory/v1/items/inventory][%d] addInventoryTooManyRequests  %+v", 429, o.Payload)
}

func (o *AddInventoryTooManyRequests) String() string {
	return fmt.Sprintf("[POST /fba/inventory/v1/items/inventory][%d] addInventoryTooManyRequests  %+v", 429, o.Payload)
}

func (o *AddInventoryTooManyRequests) GetPayload() *fba_inventory_models.AddInventoryResponse {
	return o.Payload
}

func (o *AddInventoryTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(fba_inventory_models.AddInventoryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddInventoryInternalServerError creates a AddInventoryInternalServerError with default headers values
func NewAddInventoryInternalServerError() *AddInventoryInternalServerError {
	return &AddInventoryInternalServerError{}
}

/*
AddInventoryInternalServerError describes a response with status code 500, with default header values.

An unexpected condition occurred that prevented the server from fulfilling the request.
*/
type AddInventoryInternalServerError struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fba_inventory_models.AddInventoryResponse
}

// IsSuccess returns true when this add inventory internal server error response has a 2xx status code
func (o *AddInventoryInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add inventory internal server error response has a 3xx status code
func (o *AddInventoryInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add inventory internal server error response has a 4xx status code
func (o *AddInventoryInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this add inventory internal server error response has a 5xx status code
func (o *AddInventoryInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this add inventory internal server error response a status code equal to that given
func (o *AddInventoryInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *AddInventoryInternalServerError) Error() string {
	return fmt.Sprintf("[POST /fba/inventory/v1/items/inventory][%d] addInventoryInternalServerError  %+v", 500, o.Payload)
}

func (o *AddInventoryInternalServerError) String() string {
	return fmt.Sprintf("[POST /fba/inventory/v1/items/inventory][%d] addInventoryInternalServerError  %+v", 500, o.Payload)
}

func (o *AddInventoryInternalServerError) GetPayload() *fba_inventory_models.AddInventoryResponse {
	return o.Payload
}

func (o *AddInventoryInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(fba_inventory_models.AddInventoryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddInventoryServiceUnavailable creates a AddInventoryServiceUnavailable with default headers values
func NewAddInventoryServiceUnavailable() *AddInventoryServiceUnavailable {
	return &AddInventoryServiceUnavailable{}
}

/*
AddInventoryServiceUnavailable describes a response with status code 503, with default header values.

Temporary overloading or maintenance of the server.
*/
type AddInventoryServiceUnavailable struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fba_inventory_models.AddInventoryResponse
}

// IsSuccess returns true when this add inventory service unavailable response has a 2xx status code
func (o *AddInventoryServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add inventory service unavailable response has a 3xx status code
func (o *AddInventoryServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add inventory service unavailable response has a 4xx status code
func (o *AddInventoryServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this add inventory service unavailable response has a 5xx status code
func (o *AddInventoryServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this add inventory service unavailable response a status code equal to that given
func (o *AddInventoryServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *AddInventoryServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /fba/inventory/v1/items/inventory][%d] addInventoryServiceUnavailable  %+v", 503, o.Payload)
}

func (o *AddInventoryServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /fba/inventory/v1/items/inventory][%d] addInventoryServiceUnavailable  %+v", 503, o.Payload)
}

func (o *AddInventoryServiceUnavailable) GetPayload() *fba_inventory_models.AddInventoryResponse {
	return o.Payload
}

func (o *AddInventoryServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(fba_inventory_models.AddInventoryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
