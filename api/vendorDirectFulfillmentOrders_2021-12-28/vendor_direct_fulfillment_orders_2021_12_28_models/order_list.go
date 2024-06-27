// Code generated by go-swagger; DO NOT EDIT.

package vendor_direct_fulfillment_orders_2021_12_28_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// OrderList A list of purchase orders.
//
// swagger:model OrderList
type OrderList struct {

	// Represents a purchase order within the OrderList.
	Orders []*Order `json:"orders"`

	// The pagination elements required to retrieve the remaining data.
	Pagination *Pagination `json:"pagination,omitempty"`
}

// Validate validates this order list
func (m *OrderList) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOrders(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePagination(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OrderList) validateOrders(formats strfmt.Registry) error {
	if swag.IsZero(m.Orders) { // not required
		return nil
	}

	for i := 0; i < len(m.Orders); i++ {
		if swag.IsZero(m.Orders[i]) { // not required
			continue
		}

		if m.Orders[i] != nil {
			if err := m.Orders[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("orders" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("orders" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *OrderList) validatePagination(formats strfmt.Registry) error {
	if swag.IsZero(m.Pagination) { // not required
		return nil
	}

	if m.Pagination != nil {
		if err := m.Pagination.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pagination")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pagination")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this order list based on the context it is used
func (m *OrderList) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateOrders(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePagination(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OrderList) contextValidateOrders(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Orders); i++ {

		if m.Orders[i] != nil {
			if err := m.Orders[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("orders" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("orders" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *OrderList) contextValidatePagination(ctx context.Context, formats strfmt.Registry) error {

	if m.Pagination != nil {
		if err := m.Pagination.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pagination")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pagination")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OrderList) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OrderList) UnmarshalBinary(b []byte) error {
	var res OrderList
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
