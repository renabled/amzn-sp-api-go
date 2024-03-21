// Code generated by go-swagger; DO NOT EDIT.

package vendor_shipments_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CollectFreightPickupDetails Transport Request pickup date from Vendor Warehouse by Buyer
//
// swagger:model CollectFreightPickupDetails
type CollectFreightPickupDetails struct {

	// Date on which the carrier is being scheduled to pickup items from vendor warehouse by Byer used for WePay/Collect vendors.
	// Format: date-time
	CarrierAssignmentDate strfmt.DateTime `json:"carrierAssignmentDate,omitempty"`

	// Date on which the items can be picked up from vendor warehouse by Buyer used for WePay/Collect vendors.
	// Format: date-time
	RequestedPickUp strfmt.DateTime `json:"requestedPickUp,omitempty"`

	// Date on which the items are scheduled to be picked from vendor warehouse by Buyer used for WePay/Collect vendors.
	// Format: date-time
	ScheduledPickUp strfmt.DateTime `json:"scheduledPickUp,omitempty"`
}

// Validate validates this collect freight pickup details
func (m *CollectFreightPickupDetails) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCarrierAssignmentDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequestedPickUp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScheduledPickUp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CollectFreightPickupDetails) validateCarrierAssignmentDate(formats strfmt.Registry) error {
	if swag.IsZero(m.CarrierAssignmentDate) { // not required
		return nil
	}

	if err := validate.FormatOf("carrierAssignmentDate", "body", "date-time", m.CarrierAssignmentDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CollectFreightPickupDetails) validateRequestedPickUp(formats strfmt.Registry) error {
	if swag.IsZero(m.RequestedPickUp) { // not required
		return nil
	}

	if err := validate.FormatOf("requestedPickUp", "body", "date-time", m.RequestedPickUp.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CollectFreightPickupDetails) validateScheduledPickUp(formats strfmt.Registry) error {
	if swag.IsZero(m.ScheduledPickUp) { // not required
		return nil
	}

	if err := validate.FormatOf("scheduledPickUp", "body", "date-time", m.ScheduledPickUp.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this collect freight pickup details based on context it is used
func (m *CollectFreightPickupDetails) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CollectFreightPickupDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CollectFreightPickupDetails) UnmarshalBinary(b []byte) error {
	var res CollectFreightPickupDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
