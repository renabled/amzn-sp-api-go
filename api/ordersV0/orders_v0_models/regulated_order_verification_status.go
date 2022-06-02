// Code generated by go-swagger; DO NOT EDIT.

package orders_v0_models

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

// RegulatedOrderVerificationStatus The verification status of the order along with associated approval or rejection metadata.
//
// swagger:model RegulatedOrderVerificationStatus
type RegulatedOrderVerificationStatus struct {

	// The identifier for the order's regulated information reviewer.
	ExternalReviewerID string `json:"ExternalReviewerId,omitempty"`

	// The reason for rejecting the order's regulated information. Not present if the order isn't rejected.
	RejectionReason *RejectionReason `json:"RejectionReason,omitempty"`

	// Whether the regulated information provided in the order requires a review by the merchant.
	// Required: true
	RequiresMerchantAction *bool `json:"RequiresMerchantAction"`

	// The date the order was reviewed. In ISO 8601 date time format.
	ReviewDate string `json:"ReviewDate,omitempty"`

	// The verification status of the order.
	// Required: true
	Status *VerificationStatus `json:"Status"`

	// A list of valid rejection reasons that may be used to reject the order's regulated information.
	// Required: true
	ValidRejectionReasons []*RejectionReason `json:"ValidRejectionReasons"`
}

// Validate validates this regulated order verification status
func (m *RegulatedOrderVerificationStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRejectionReason(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequiresMerchantAction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValidRejectionReasons(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RegulatedOrderVerificationStatus) validateRejectionReason(formats strfmt.Registry) error {
	if swag.IsZero(m.RejectionReason) { // not required
		return nil
	}

	if m.RejectionReason != nil {
		if err := m.RejectionReason.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("RejectionReason")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("RejectionReason")
			}
			return err
		}
	}

	return nil
}

func (m *RegulatedOrderVerificationStatus) validateRequiresMerchantAction(formats strfmt.Registry) error {

	if err := validate.Required("RequiresMerchantAction", "body", m.RequiresMerchantAction); err != nil {
		return err
	}

	return nil
}

func (m *RegulatedOrderVerificationStatus) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("Status", "body", m.Status); err != nil {
		return err
	}

	if err := validate.Required("Status", "body", m.Status); err != nil {
		return err
	}

	if m.Status != nil {
		if err := m.Status.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Status")
			}
			return err
		}
	}

	return nil
}

func (m *RegulatedOrderVerificationStatus) validateValidRejectionReasons(formats strfmt.Registry) error {

	if err := validate.Required("ValidRejectionReasons", "body", m.ValidRejectionReasons); err != nil {
		return err
	}

	for i := 0; i < len(m.ValidRejectionReasons); i++ {
		if swag.IsZero(m.ValidRejectionReasons[i]) { // not required
			continue
		}

		if m.ValidRejectionReasons[i] != nil {
			if err := m.ValidRejectionReasons[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ValidRejectionReasons" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ValidRejectionReasons" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this regulated order verification status based on the context it is used
func (m *RegulatedOrderVerificationStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRejectionReason(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateValidRejectionReasons(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RegulatedOrderVerificationStatus) contextValidateRejectionReason(ctx context.Context, formats strfmt.Registry) error {

	if m.RejectionReason != nil {
		if err := m.RejectionReason.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("RejectionReason")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("RejectionReason")
			}
			return err
		}
	}

	return nil
}

func (m *RegulatedOrderVerificationStatus) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.Status != nil {
		if err := m.Status.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Status")
			}
			return err
		}
	}

	return nil
}

func (m *RegulatedOrderVerificationStatus) contextValidateValidRejectionReasons(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ValidRejectionReasons); i++ {

		if m.ValidRejectionReasons[i] != nil {
			if err := m.ValidRejectionReasons[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ValidRejectionReasons" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ValidRejectionReasons" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *RegulatedOrderVerificationStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RegulatedOrderVerificationStatus) UnmarshalBinary(b []byte) error {
	var res RegulatedOrderVerificationStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
