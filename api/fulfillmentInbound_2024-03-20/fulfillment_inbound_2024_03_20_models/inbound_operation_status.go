// Code generated by go-swagger; DO NOT EDIT.

package fulfillment_inbound_2024_03_20_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// InboundOperationStatus GetInboundOperationStatus response.
//
// swagger:model InboundOperationStatus
type InboundOperationStatus struct {

	// The name of the operation in the asynchronous API call.
	// Required: true
	// Max Length: 1024
	// Min Length: 1
	Operation *string `json:"operation"`

	// The operation ID returned by the asynchronous API call.
	// Required: true
	// Max Length: 38
	// Min Length: 36
	// Pattern: ^[a-zA-Z0-9-]*$
	OperationID *string `json:"operationId"`

	// The problems in the processing of the asynchronous operation.
	// Required: true
	OperationProblems []*OperationProblem `json:"operationProblems"`

	// operation status
	// Required: true
	OperationStatus *OperationStatus `json:"operationStatus"`
}

// Validate validates this inbound operation status
func (m *InboundOperationStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOperation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOperationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOperationProblems(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOperationStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InboundOperationStatus) validateOperation(formats strfmt.Registry) error {

	if err := validate.Required("operation", "body", m.Operation); err != nil {
		return err
	}

	if err := validate.MinLength("operation", "body", *m.Operation, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("operation", "body", *m.Operation, 1024); err != nil {
		return err
	}

	return nil
}

func (m *InboundOperationStatus) validateOperationID(formats strfmt.Registry) error {

	if err := validate.Required("operationId", "body", m.OperationID); err != nil {
		return err
	}

	if err := validate.MinLength("operationId", "body", *m.OperationID, 36); err != nil {
		return err
	}

	if err := validate.MaxLength("operationId", "body", *m.OperationID, 38); err != nil {
		return err
	}

	if err := validate.Pattern("operationId", "body", *m.OperationID, `^[a-zA-Z0-9-]*$`); err != nil {
		return err
	}

	return nil
}

func (m *InboundOperationStatus) validateOperationProblems(formats strfmt.Registry) error {

	if err := validate.Required("operationProblems", "body", m.OperationProblems); err != nil {
		return err
	}

	for i := 0; i < len(m.OperationProblems); i++ {
		if swag.IsZero(m.OperationProblems[i]) { // not required
			continue
		}

		if m.OperationProblems[i] != nil {
			if err := m.OperationProblems[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("operationProblems" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("operationProblems" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *InboundOperationStatus) validateOperationStatus(formats strfmt.Registry) error {

	if err := validate.Required("operationStatus", "body", m.OperationStatus); err != nil {
		return err
	}

	if err := validate.Required("operationStatus", "body", m.OperationStatus); err != nil {
		return err
	}

	if m.OperationStatus != nil {
		if err := m.OperationStatus.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("operationStatus")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("operationStatus")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this inbound operation status based on the context it is used
func (m *InboundOperationStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateOperationProblems(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOperationStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InboundOperationStatus) contextValidateOperationProblems(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.OperationProblems); i++ {

		if m.OperationProblems[i] != nil {
			if err := m.OperationProblems[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("operationProblems" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("operationProblems" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *InboundOperationStatus) contextValidateOperationStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.OperationStatus != nil {
		if err := m.OperationStatus.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("operationStatus")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("operationStatus")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InboundOperationStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InboundOperationStatus) UnmarshalBinary(b []byte) error {
	var res InboundOperationStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
