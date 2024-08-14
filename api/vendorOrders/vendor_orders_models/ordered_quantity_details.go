// Code generated by go-swagger; DO NOT EDIT.

package vendor_orders_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// OrderedQuantityDetails Details of item quantity ordered
//
// swagger:model OrderedQuantityDetails
type OrderedQuantityDetails struct {

	// Item quantity ordered.
	CancelledQuantity *ItemQuantity `json:"cancelledQuantity,omitempty"`

	// Item quantity ordered.
	OrderedQuantity *ItemQuantity `json:"orderedQuantity,omitempty"`

	// The date when the line item quantity was updated by buyer. Must be in ISO-8601 date/time format.
	// Format: date-time
	UpdatedDate strfmt.DateTime `json:"updatedDate,omitempty"`
}

// Validate validates this ordered quantity details
func (m *OrderedQuantityDetails) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCancelledQuantity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrderedQuantity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OrderedQuantityDetails) validateCancelledQuantity(formats strfmt.Registry) error {
	if swag.IsZero(m.CancelledQuantity) { // not required
		return nil
	}

	if m.CancelledQuantity != nil {
		if err := m.CancelledQuantity.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cancelledQuantity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cancelledQuantity")
			}
			return err
		}
	}

	return nil
}

func (m *OrderedQuantityDetails) validateOrderedQuantity(formats strfmt.Registry) error {
	if swag.IsZero(m.OrderedQuantity) { // not required
		return nil
	}

	if m.OrderedQuantity != nil {
		if err := m.OrderedQuantity.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("orderedQuantity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("orderedQuantity")
			}
			return err
		}
	}

	return nil
}

func (m *OrderedQuantityDetails) validateUpdatedDate(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedDate) { // not required
		return nil
	}

	if err := validate.FormatOf("updatedDate", "body", "date-time", m.UpdatedDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this ordered quantity details based on the context it is used
func (m *OrderedQuantityDetails) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCancelledQuantity(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOrderedQuantity(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OrderedQuantityDetails) contextValidateCancelledQuantity(ctx context.Context, formats strfmt.Registry) error {

	if m.CancelledQuantity != nil {
		if err := m.CancelledQuantity.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cancelledQuantity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cancelledQuantity")
			}
			return err
		}
	}

	return nil
}

func (m *OrderedQuantityDetails) contextValidateOrderedQuantity(ctx context.Context, formats strfmt.Registry) error {

	if m.OrderedQuantity != nil {
		if err := m.OrderedQuantity.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("orderedQuantity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("orderedQuantity")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OrderedQuantityDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OrderedQuantityDetails) UnmarshalBinary(b []byte) error {
	var res OrderedQuantityDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
