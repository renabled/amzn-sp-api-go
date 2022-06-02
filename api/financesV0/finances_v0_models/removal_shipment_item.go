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

// RemovalShipmentItem Item-level information for a removal shipment.
//
// swagger:model RemovalShipmentItem
type RemovalShipmentItem struct {

	// The fee that Amazon charged to the seller for the removal of the item. The amount is a negative number.
	FeeAmount *Currency `json:"FeeAmount,omitempty"`

	// The Amazon fulfillment network SKU for the item.
	FulfillmentNetworkSKU string `json:"FulfillmentNetworkSKU,omitempty"`

	// The quantity of the item.
	Quantity int32 `json:"Quantity,omitempty"`

	// An identifier for an item in a removal shipment.
	RemovalShipmentItemID string `json:"RemovalShipmentItemId,omitempty"`

	// The total amount paid to the seller for the removed item.
	Revenue *Currency `json:"Revenue,omitempty"`

	// Tax collected on the revenue.
	TaxAmount *Currency `json:"TaxAmount,omitempty"`

	// The tax collection model applied to the item.
	//
	// Possible values:
	//
	// * MarketplaceFacilitator - Tax is withheld and remitted to the taxing authority by Amazon on behalf of the seller.
	//
	// * Standard - Tax is paid to the seller and not remitted to the taxing authority by Amazon.
	TaxCollectionModel string `json:"TaxCollectionModel,omitempty"`

	// The tax withheld and remitted to the taxing authority by Amazon on behalf of the seller. If TaxCollectionModel=MarketplaceFacilitator, then TaxWithheld=TaxAmount (except the TaxWithheld amount is a negative number). Otherwise TaxWithheld=0.
	TaxWithheld *Currency `json:"TaxWithheld,omitempty"`
}

// Validate validates this removal shipment item
func (m *RemovalShipmentItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFeeAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRevenue(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTaxAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTaxWithheld(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RemovalShipmentItem) validateFeeAmount(formats strfmt.Registry) error {
	if swag.IsZero(m.FeeAmount) { // not required
		return nil
	}

	if m.FeeAmount != nil {
		if err := m.FeeAmount.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("FeeAmount")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("FeeAmount")
			}
			return err
		}
	}

	return nil
}

func (m *RemovalShipmentItem) validateRevenue(formats strfmt.Registry) error {
	if swag.IsZero(m.Revenue) { // not required
		return nil
	}

	if m.Revenue != nil {
		if err := m.Revenue.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Revenue")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Revenue")
			}
			return err
		}
	}

	return nil
}

func (m *RemovalShipmentItem) validateTaxAmount(formats strfmt.Registry) error {
	if swag.IsZero(m.TaxAmount) { // not required
		return nil
	}

	if m.TaxAmount != nil {
		if err := m.TaxAmount.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("TaxAmount")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("TaxAmount")
			}
			return err
		}
	}

	return nil
}

func (m *RemovalShipmentItem) validateTaxWithheld(formats strfmt.Registry) error {
	if swag.IsZero(m.TaxWithheld) { // not required
		return nil
	}

	if m.TaxWithheld != nil {
		if err := m.TaxWithheld.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("TaxWithheld")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("TaxWithheld")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this removal shipment item based on the context it is used
func (m *RemovalShipmentItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFeeAmount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRevenue(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTaxAmount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTaxWithheld(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RemovalShipmentItem) contextValidateFeeAmount(ctx context.Context, formats strfmt.Registry) error {

	if m.FeeAmount != nil {
		if err := m.FeeAmount.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("FeeAmount")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("FeeAmount")
			}
			return err
		}
	}

	return nil
}

func (m *RemovalShipmentItem) contextValidateRevenue(ctx context.Context, formats strfmt.Registry) error {

	if m.Revenue != nil {
		if err := m.Revenue.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Revenue")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Revenue")
			}
			return err
		}
	}

	return nil
}

func (m *RemovalShipmentItem) contextValidateTaxAmount(ctx context.Context, formats strfmt.Registry) error {

	if m.TaxAmount != nil {
		if err := m.TaxAmount.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("TaxAmount")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("TaxAmount")
			}
			return err
		}
	}

	return nil
}

func (m *RemovalShipmentItem) contextValidateTaxWithheld(ctx context.Context, formats strfmt.Registry) error {

	if m.TaxWithheld != nil {
		if err := m.TaxWithheld.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("TaxWithheld")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("TaxWithheld")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RemovalShipmentItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RemovalShipmentItem) UnmarshalBinary(b []byte) error {
	var res RemovalShipmentItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
