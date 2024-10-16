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

// PayWithAmazonEvent An event related to the seller's Pay with Amazon account.
//
// swagger:model PayWithAmazonEvent
type PayWithAmazonEvent struct {

	// A short description of this payment event.
	AmountDescription string `json:"AmountDescription,omitempty"`

	// The type of business object.
	BusinessObjectType string `json:"BusinessObjectType,omitempty"`

	// The charge associated with the event.
	Charge *ChargeComponent `json:"Charge,omitempty"`

	// A list of fees associated with the event.
	FeeList FeeComponentList `json:"FeeList,omitempty"`

	// The fulfillment channel.
	//
	// Possible values:
	//
	// * `AFN`: Amazon Fulfillment Network (Fulfillment by Amazon)
	//
	// * `MFN`: Merchant Fulfillment Network (self-fulfilled)
	FulfillmentChannel string `json:"FulfillmentChannel,omitempty"`

	// The type of payment.
	//
	// Possible values:
	//
	// * `Sales`
	PaymentAmountType string `json:"PaymentAmountType,omitempty"`

	// The sales channel for the transaction.
	SalesChannel string `json:"SalesChannel,omitempty"`

	// An order identifier that is specified by the seller.
	SellerOrderID string `json:"SellerOrderId,omitempty"`

	// The name of the store where the event occurred.
	StoreName string `json:"StoreName,omitempty"`

	// The date and time when the payment transaction is posted. In [ISO 8601](https://developer-docs.amazon.com/sp-api/docs/iso-8601) date-time format.
	// Format: date-time
	TransactionPostedDate Date `json:"TransactionPostedDate,omitempty"`
}

// Validate validates this pay with amazon event
func (m *PayWithAmazonEvent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCharge(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFeeList(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTransactionPostedDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PayWithAmazonEvent) validateCharge(formats strfmt.Registry) error {
	if swag.IsZero(m.Charge) { // not required
		return nil
	}

	if m.Charge != nil {
		if err := m.Charge.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Charge")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Charge")
			}
			return err
		}
	}

	return nil
}

func (m *PayWithAmazonEvent) validateFeeList(formats strfmt.Registry) error {
	if swag.IsZero(m.FeeList) { // not required
		return nil
	}

	if err := m.FeeList.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("FeeList")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("FeeList")
		}
		return err
	}

	return nil
}

func (m *PayWithAmazonEvent) validateTransactionPostedDate(formats strfmt.Registry) error {
	if swag.IsZero(m.TransactionPostedDate) { // not required
		return nil
	}

	if err := m.TransactionPostedDate.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("TransactionPostedDate")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("TransactionPostedDate")
		}
		return err
	}

	return nil
}

// ContextValidate validate this pay with amazon event based on the context it is used
func (m *PayWithAmazonEvent) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCharge(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFeeList(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTransactionPostedDate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PayWithAmazonEvent) contextValidateCharge(ctx context.Context, formats strfmt.Registry) error {

	if m.Charge != nil {
		if err := m.Charge.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Charge")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Charge")
			}
			return err
		}
	}

	return nil
}

func (m *PayWithAmazonEvent) contextValidateFeeList(ctx context.Context, formats strfmt.Registry) error {

	if err := m.FeeList.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("FeeList")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("FeeList")
		}
		return err
	}

	return nil
}

func (m *PayWithAmazonEvent) contextValidateTransactionPostedDate(ctx context.Context, formats strfmt.Registry) error {

	if err := m.TransactionPostedDate.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("TransactionPostedDate")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("TransactionPostedDate")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PayWithAmazonEvent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PayWithAmazonEvent) UnmarshalBinary(b []byte) error {
	var res PayWithAmazonEvent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
