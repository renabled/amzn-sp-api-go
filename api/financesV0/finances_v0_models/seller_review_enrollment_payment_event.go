// Code generated by go-swagger; DO NOT EDIT.

package finances_v0_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SellerReviewEnrollmentPaymentEvent A fee payment event for the Early Reviewer Program.
//
// swagger:model SellerReviewEnrollmentPaymentEvent
type SellerReviewEnrollmentPaymentEvent struct {

	// charge component
	ChargeComponent *ChargeComponent `json:"ChargeComponent,omitempty"`

	// An enrollment identifier.
	EnrollmentID string `json:"EnrollmentId,omitempty"`

	// fee component
	FeeComponent *FeeComponent `json:"FeeComponent,omitempty"`

	// The Amazon Standard Identification Number (ASIN) of the item that was enrolled in the Early Reviewer Program.
	ParentASIN string `json:"ParentASIN,omitempty"`

	// The date and time when the financial event was posted.
	// Format: date-time
	PostedDate Date `json:"PostedDate,omitempty"`

	// The `FeeComponent` value plus the `ChargeComponent` value.
	TotalAmount *Currency `json:"TotalAmount,omitempty"`
}

// Validate validates this seller review enrollment payment event
func (m *SellerReviewEnrollmentPaymentEvent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChargeComponent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFeeComponent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePostedDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTotalAmount(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SellerReviewEnrollmentPaymentEvent) validateChargeComponent(formats strfmt.Registry) error {
	if swag.IsZero(m.ChargeComponent) { // not required
		return nil
	}

	if m.ChargeComponent != nil {
		if err := m.ChargeComponent.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ChargeComponent")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ChargeComponent")
			}
			return err
		}
	}

	return nil
}

func (m *SellerReviewEnrollmentPaymentEvent) validateFeeComponent(formats strfmt.Registry) error {
	if swag.IsZero(m.FeeComponent) { // not required
		return nil
	}

	if m.FeeComponent != nil {
		if err := m.FeeComponent.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("FeeComponent")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("FeeComponent")
			}
			return err
		}
	}

	return nil
}

func (m *SellerReviewEnrollmentPaymentEvent) validatePostedDate(formats strfmt.Registry) error {
	if swag.IsZero(m.PostedDate) { // not required
		return nil
	}

	if err := m.PostedDate.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("PostedDate")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("PostedDate")
		}
		return err
	}

	return nil
}

func (m *SellerReviewEnrollmentPaymentEvent) validateTotalAmount(formats strfmt.Registry) error {
	if swag.IsZero(m.TotalAmount) { // not required
		return nil
	}

	if m.TotalAmount != nil {
		if err := m.TotalAmount.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("TotalAmount")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("TotalAmount")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this seller review enrollment payment event based on the context it is used
func (m *SellerReviewEnrollmentPaymentEvent) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateChargeComponent(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFeeComponent(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePostedDate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTotalAmount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SellerReviewEnrollmentPaymentEvent) contextValidateChargeComponent(ctx context.Context, formats strfmt.Registry) error {

	if m.ChargeComponent != nil {
		if err := m.ChargeComponent.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ChargeComponent")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ChargeComponent")
			}
			return err
		}
	}

	return nil
}

func (m *SellerReviewEnrollmentPaymentEvent) contextValidateFeeComponent(ctx context.Context, formats strfmt.Registry) error {

	if m.FeeComponent != nil {
		if err := m.FeeComponent.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("FeeComponent")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("FeeComponent")
			}
			return err
		}
	}

	return nil
}

func (m *SellerReviewEnrollmentPaymentEvent) contextValidatePostedDate(ctx context.Context, formats strfmt.Registry) error {

	if err := m.PostedDate.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("PostedDate")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("PostedDate")
		}
		return err
	}

	return nil
}

func (m *SellerReviewEnrollmentPaymentEvent) contextValidateTotalAmount(ctx context.Context, formats strfmt.Registry) error {

	if m.TotalAmount != nil {
		if err := m.TotalAmount.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("TotalAmount")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("TotalAmount")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SellerReviewEnrollmentPaymentEvent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SellerReviewEnrollmentPaymentEvent) UnmarshalBinary(b []byte) error {
	var res SellerReviewEnrollmentPaymentEvent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
