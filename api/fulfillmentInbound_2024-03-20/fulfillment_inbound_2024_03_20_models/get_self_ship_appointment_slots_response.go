// Code generated by go-swagger; DO NOT EDIT.

package fulfillment_inbound_2024_03_20_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GetSelfShipAppointmentSlotsResponse The `getSelfShipAppointmentSlots` response.
//
// swagger:model GetSelfShipAppointmentSlotsResponse
type GetSelfShipAppointmentSlotsResponse struct {

	// pagination
	Pagination *Pagination `json:"pagination,omitempty"`

	// self ship appointment slots availability
	// Required: true
	SelfShipAppointmentSlotsAvailability *SelfShipAppointmentSlotsAvailability `json:"selfShipAppointmentSlotsAvailability"`
}

// Validate validates this get self ship appointment slots response
func (m *GetSelfShipAppointmentSlotsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePagination(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelfShipAppointmentSlotsAvailability(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetSelfShipAppointmentSlotsResponse) validatePagination(formats strfmt.Registry) error {
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

func (m *GetSelfShipAppointmentSlotsResponse) validateSelfShipAppointmentSlotsAvailability(formats strfmt.Registry) error {

	if err := validate.Required("selfShipAppointmentSlotsAvailability", "body", m.SelfShipAppointmentSlotsAvailability); err != nil {
		return err
	}

	if m.SelfShipAppointmentSlotsAvailability != nil {
		if err := m.SelfShipAppointmentSlotsAvailability.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("selfShipAppointmentSlotsAvailability")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("selfShipAppointmentSlotsAvailability")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get self ship appointment slots response based on the context it is used
func (m *GetSelfShipAppointmentSlotsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePagination(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSelfShipAppointmentSlotsAvailability(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetSelfShipAppointmentSlotsResponse) contextValidatePagination(ctx context.Context, formats strfmt.Registry) error {

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

func (m *GetSelfShipAppointmentSlotsResponse) contextValidateSelfShipAppointmentSlotsAvailability(ctx context.Context, formats strfmt.Registry) error {

	if m.SelfShipAppointmentSlotsAvailability != nil {
		if err := m.SelfShipAppointmentSlotsAvailability.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("selfShipAppointmentSlotsAvailability")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("selfShipAppointmentSlotsAvailability")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetSelfShipAppointmentSlotsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetSelfShipAppointmentSlotsResponse) UnmarshalBinary(b []byte) error {
	var res GetSelfShipAppointmentSlotsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
