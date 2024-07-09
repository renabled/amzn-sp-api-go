// Code generated by go-swagger; DO NOT EDIT.

package merchant_fulfillment_v0_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ShipmentRequestDetails Shipment information required for requesting shipping service offers or for creating a shipment.
//
// swagger:model ShipmentRequestDetails
type ShipmentRequestDetails struct {

	// An Amazon-defined order identifier in 3-7-7 format.
	// Required: true
	AmazonOrderID *AmazonOrderID `json:"AmazonOrderId"`

	// item list
	// Required: true
	ItemList ItemList `json:"ItemList"`

	// Label customization options.
	LabelCustomization *LabelCustomization `json:"LabelCustomization,omitempty"`

	// The date by which the package must arrive to keep the promise to the customer, in ISO 8601 datetime format. If MustArriveByDate is specified, only shipping service offers that can be delivered by that date are returned.
	// Format: date-time
	MustArriveByDate Timestamp `json:"MustArriveByDate,omitempty"`

	// The package dimensions.
	// Required: true
	PackageDimensions *PackageDimensions `json:"PackageDimensions"`

	// A seller-defined order identifier.
	SellerOrderID SellerOrderID `json:"SellerOrderId,omitempty"`

	// When used in a request, this is the date and time that the seller wants to ship the package. When used in a response, this is the date and time that the package can be shipped by the indicated method.
	// Format: date-time
	ShipDate Timestamp `json:"ShipDate,omitempty"`

	// The address of the sender.
	// Required: true
	ShipFromAddress *Address `json:"ShipFromAddress"`

	// Extra services offered by the carrier.
	// Required: true
	ShippingServiceOptions *ShippingServiceOptions `json:"ShippingServiceOptions"`

	// The package weight.
	// Required: true
	Weight *Weight `json:"Weight"`
}

// Validate validates this shipment request details
func (m *ShipmentRequestDetails) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmazonOrderID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateItemList(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabelCustomization(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMustArriveByDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePackageDimensions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSellerOrderID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShipDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShipFromAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShippingServiceOptions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWeight(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ShipmentRequestDetails) validateAmazonOrderID(formats strfmt.Registry) error {

	if err := validate.Required("AmazonOrderId", "body", m.AmazonOrderID); err != nil {
		return err
	}

	if err := validate.Required("AmazonOrderId", "body", m.AmazonOrderID); err != nil {
		return err
	}

	if m.AmazonOrderID != nil {
		if err := m.AmazonOrderID.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("AmazonOrderId")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("AmazonOrderId")
			}
			return err
		}
	}

	return nil
}

func (m *ShipmentRequestDetails) validateItemList(formats strfmt.Registry) error {

	if err := validate.Required("ItemList", "body", m.ItemList); err != nil {
		return err
	}

	if err := m.ItemList.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("ItemList")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("ItemList")
		}
		return err
	}

	return nil
}

func (m *ShipmentRequestDetails) validateLabelCustomization(formats strfmt.Registry) error {
	if swag.IsZero(m.LabelCustomization) { // not required
		return nil
	}

	if m.LabelCustomization != nil {
		if err := m.LabelCustomization.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("LabelCustomization")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("LabelCustomization")
			}
			return err
		}
	}

	return nil
}

func (m *ShipmentRequestDetails) validateMustArriveByDate(formats strfmt.Registry) error {
	if swag.IsZero(m.MustArriveByDate) { // not required
		return nil
	}

	if err := m.MustArriveByDate.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("MustArriveByDate")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("MustArriveByDate")
		}
		return err
	}

	return nil
}

func (m *ShipmentRequestDetails) validatePackageDimensions(formats strfmt.Registry) error {

	if err := validate.Required("PackageDimensions", "body", m.PackageDimensions); err != nil {
		return err
	}

	if m.PackageDimensions != nil {
		if err := m.PackageDimensions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("PackageDimensions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("PackageDimensions")
			}
			return err
		}
	}

	return nil
}

func (m *ShipmentRequestDetails) validateSellerOrderID(formats strfmt.Registry) error {
	if swag.IsZero(m.SellerOrderID) { // not required
		return nil
	}

	if err := m.SellerOrderID.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("SellerOrderId")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("SellerOrderId")
		}
		return err
	}

	return nil
}

func (m *ShipmentRequestDetails) validateShipDate(formats strfmt.Registry) error {
	if swag.IsZero(m.ShipDate) { // not required
		return nil
	}

	if err := m.ShipDate.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("ShipDate")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("ShipDate")
		}
		return err
	}

	return nil
}

func (m *ShipmentRequestDetails) validateShipFromAddress(formats strfmt.Registry) error {

	if err := validate.Required("ShipFromAddress", "body", m.ShipFromAddress); err != nil {
		return err
	}

	if m.ShipFromAddress != nil {
		if err := m.ShipFromAddress.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ShipFromAddress")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ShipFromAddress")
			}
			return err
		}
	}

	return nil
}

func (m *ShipmentRequestDetails) validateShippingServiceOptions(formats strfmt.Registry) error {

	if err := validate.Required("ShippingServiceOptions", "body", m.ShippingServiceOptions); err != nil {
		return err
	}

	if m.ShippingServiceOptions != nil {
		if err := m.ShippingServiceOptions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ShippingServiceOptions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ShippingServiceOptions")
			}
			return err
		}
	}

	return nil
}

func (m *ShipmentRequestDetails) validateWeight(formats strfmt.Registry) error {

	if err := validate.Required("Weight", "body", m.Weight); err != nil {
		return err
	}

	if m.Weight != nil {
		if err := m.Weight.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Weight")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Weight")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this shipment request details based on the context it is used
func (m *ShipmentRequestDetails) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAmazonOrderID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateItemList(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLabelCustomization(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMustArriveByDate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePackageDimensions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSellerOrderID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateShipDate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateShipFromAddress(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateShippingServiceOptions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWeight(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ShipmentRequestDetails) contextValidateAmazonOrderID(ctx context.Context, formats strfmt.Registry) error {

	if m.AmazonOrderID != nil {
		if err := m.AmazonOrderID.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("AmazonOrderId")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("AmazonOrderId")
			}
			return err
		}
	}

	return nil
}

func (m *ShipmentRequestDetails) contextValidateItemList(ctx context.Context, formats strfmt.Registry) error {

	if err := m.ItemList.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("ItemList")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("ItemList")
		}
		return err
	}

	return nil
}

func (m *ShipmentRequestDetails) contextValidateLabelCustomization(ctx context.Context, formats strfmt.Registry) error {

	if m.LabelCustomization != nil {
		if err := m.LabelCustomization.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("LabelCustomization")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("LabelCustomization")
			}
			return err
		}
	}

	return nil
}

func (m *ShipmentRequestDetails) contextValidateMustArriveByDate(ctx context.Context, formats strfmt.Registry) error {

	if err := m.MustArriveByDate.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("MustArriveByDate")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("MustArriveByDate")
		}
		return err
	}

	return nil
}

func (m *ShipmentRequestDetails) contextValidatePackageDimensions(ctx context.Context, formats strfmt.Registry) error {

	if m.PackageDimensions != nil {
		if err := m.PackageDimensions.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("PackageDimensions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("PackageDimensions")
			}
			return err
		}
	}

	return nil
}

func (m *ShipmentRequestDetails) contextValidateSellerOrderID(ctx context.Context, formats strfmt.Registry) error {

	if err := m.SellerOrderID.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("SellerOrderId")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("SellerOrderId")
		}
		return err
	}

	return nil
}

func (m *ShipmentRequestDetails) contextValidateShipDate(ctx context.Context, formats strfmt.Registry) error {

	if err := m.ShipDate.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("ShipDate")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("ShipDate")
		}
		return err
	}

	return nil
}

func (m *ShipmentRequestDetails) contextValidateShipFromAddress(ctx context.Context, formats strfmt.Registry) error {

	if m.ShipFromAddress != nil {
		if err := m.ShipFromAddress.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ShipFromAddress")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ShipFromAddress")
			}
			return err
		}
	}

	return nil
}

func (m *ShipmentRequestDetails) contextValidateShippingServiceOptions(ctx context.Context, formats strfmt.Registry) error {

	if m.ShippingServiceOptions != nil {
		if err := m.ShippingServiceOptions.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ShippingServiceOptions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ShippingServiceOptions")
			}
			return err
		}
	}

	return nil
}

func (m *ShipmentRequestDetails) contextValidateWeight(ctx context.Context, formats strfmt.Registry) error {

	if m.Weight != nil {
		if err := m.Weight.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Weight")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Weight")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ShipmentRequestDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ShipmentRequestDetails) UnmarshalBinary(b []byte) error {
	var res ShipmentRequestDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
