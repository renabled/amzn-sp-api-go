// Code generated by go-swagger; DO NOT EDIT.

package vendor_orders_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// OrderDetails Details of an order.
//
// swagger:model OrderDetails
type OrderDetails struct {

	// Name/Address and tax details of the bill to party.
	BillToParty *PartyIdentification `json:"billToParty,omitempty"`

	// Name/Address and tax details of the buying party.
	BuyingParty *PartyIdentification `json:"buyingParty,omitempty"`

	// If requested by the recipient, this field will contain a promotional/deal number. The discount code line is optional. It is used to obtain a price discount on items on the order.
	DealCode string `json:"dealCode,omitempty"`

	// This indicates the delivery window. Format is start and end date separated by double hyphen (--). For example, 2007-03-01T13:00:00Z--2007-03-11T15:30:00Z.
	DeliveryWindow DateTimeInterval `json:"deliveryWindow,omitempty"`

	// If the purchase order is an import order, the details for the import order.
	ImportDetails *ImportDetails `json:"importDetails,omitempty"`

	// A list of items in this purchase order.
	// Required: true
	Items []*OrderItem `json:"items"`

	// Payment method used.
	// Enum: [Invoice Consignment CreditCard Prepaid]
	PaymentMethod string `json:"paymentMethod,omitempty"`

	// The date when purchase order was last changed by Amazon after the order was placed. This date will be greater than 'purchaseOrderDate'. This means the PO data was changed on that date and vendors are required to fulfill the  updated PO. The PO changes can be related to Item Quantity, Ship to Location, Ship Window etc. This field will not be present in orders that have not changed after creation. Must be in ISO-8601 date/time format.
	// Format: date-time
	PurchaseOrderChangedDate strfmt.DateTime `json:"purchaseOrderChangedDate,omitempty"`

	// The date the purchase order was placed. Must be in ISO-8601 date/time format.
	// Required: true
	// Format: date-time
	PurchaseOrderDate *strfmt.DateTime `json:"purchaseOrderDate"`

	// The date when current purchase order state was changed. Current purchase order state is available in the field 'purchaseOrderState'. Must be in ISO-8601 date/time format.
	// Required: true
	// Format: date-time
	PurchaseOrderStateChangedDate *strfmt.DateTime `json:"purchaseOrderStateChangedDate"`

	// Type of purchase order.
	// Enum: [RegularOrder ConsignedOrder NewProductIntroduction RushOrder]
	PurchaseOrderType string `json:"purchaseOrderType,omitempty"`

	// Name/Address and tax details of the selling party.
	SellingParty *PartyIdentification `json:"sellingParty,omitempty"`

	// Name/Address and tax details of the ship to party. Find a list of fulfillment center addresses for a region on the [Resources page of Amazon Vendor Central](https://vendorcentral.amazon.com/hz/vendor/members/support/help/node/GPZ88XH8HQM97ZV6).
	ShipToParty *PartyIdentification `json:"shipToParty,omitempty"`

	// This indicates the ship window. Format is start and end date separated by double hyphen (--). For example, 2007-03-01T13:00:00Z--2007-03-11T15:30:00Z.
	ShipWindow DateTimeInterval `json:"shipWindow,omitempty"`
}

// Validate validates this order details
func (m *OrderDetails) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBillToParty(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBuyingParty(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeliveryWindow(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImportDetails(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateItems(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePaymentMethod(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePurchaseOrderChangedDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePurchaseOrderDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePurchaseOrderStateChangedDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePurchaseOrderType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSellingParty(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShipToParty(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShipWindow(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OrderDetails) validateBillToParty(formats strfmt.Registry) error {
	if swag.IsZero(m.BillToParty) { // not required
		return nil
	}

	if m.BillToParty != nil {
		if err := m.BillToParty.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("billToParty")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("billToParty")
			}
			return err
		}
	}

	return nil
}

func (m *OrderDetails) validateBuyingParty(formats strfmt.Registry) error {
	if swag.IsZero(m.BuyingParty) { // not required
		return nil
	}

	if m.BuyingParty != nil {
		if err := m.BuyingParty.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("buyingParty")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("buyingParty")
			}
			return err
		}
	}

	return nil
}

func (m *OrderDetails) validateDeliveryWindow(formats strfmt.Registry) error {
	if swag.IsZero(m.DeliveryWindow) { // not required
		return nil
	}

	if err := m.DeliveryWindow.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("deliveryWindow")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("deliveryWindow")
		}
		return err
	}

	return nil
}

func (m *OrderDetails) validateImportDetails(formats strfmt.Registry) error {
	if swag.IsZero(m.ImportDetails) { // not required
		return nil
	}

	if m.ImportDetails != nil {
		if err := m.ImportDetails.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("importDetails")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("importDetails")
			}
			return err
		}
	}

	return nil
}

func (m *OrderDetails) validateItems(formats strfmt.Registry) error {

	if err := validate.Required("items", "body", m.Items); err != nil {
		return err
	}

	for i := 0; i < len(m.Items); i++ {
		if swag.IsZero(m.Items[i]) { // not required
			continue
		}

		if m.Items[i] != nil {
			if err := m.Items[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("items" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var orderDetailsTypePaymentMethodPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Invoice","Consignment","CreditCard","Prepaid"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		orderDetailsTypePaymentMethodPropEnum = append(orderDetailsTypePaymentMethodPropEnum, v)
	}
}

const (

	// OrderDetailsPaymentMethodInvoice captures enum value "Invoice"
	OrderDetailsPaymentMethodInvoice string = "Invoice"

	// OrderDetailsPaymentMethodConsignment captures enum value "Consignment"
	OrderDetailsPaymentMethodConsignment string = "Consignment"

	// OrderDetailsPaymentMethodCreditCard captures enum value "CreditCard"
	OrderDetailsPaymentMethodCreditCard string = "CreditCard"

	// OrderDetailsPaymentMethodPrepaid captures enum value "Prepaid"
	OrderDetailsPaymentMethodPrepaid string = "Prepaid"
)

// prop value enum
func (m *OrderDetails) validatePaymentMethodEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, orderDetailsTypePaymentMethodPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *OrderDetails) validatePaymentMethod(formats strfmt.Registry) error {
	if swag.IsZero(m.PaymentMethod) { // not required
		return nil
	}

	// value enum
	if err := m.validatePaymentMethodEnum("paymentMethod", "body", m.PaymentMethod); err != nil {
		return err
	}

	return nil
}

func (m *OrderDetails) validatePurchaseOrderChangedDate(formats strfmt.Registry) error {
	if swag.IsZero(m.PurchaseOrderChangedDate) { // not required
		return nil
	}

	if err := validate.FormatOf("purchaseOrderChangedDate", "body", "date-time", m.PurchaseOrderChangedDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *OrderDetails) validatePurchaseOrderDate(formats strfmt.Registry) error {

	if err := validate.Required("purchaseOrderDate", "body", m.PurchaseOrderDate); err != nil {
		return err
	}

	if err := validate.FormatOf("purchaseOrderDate", "body", "date-time", m.PurchaseOrderDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *OrderDetails) validatePurchaseOrderStateChangedDate(formats strfmt.Registry) error {

	if err := validate.Required("purchaseOrderStateChangedDate", "body", m.PurchaseOrderStateChangedDate); err != nil {
		return err
	}

	if err := validate.FormatOf("purchaseOrderStateChangedDate", "body", "date-time", m.PurchaseOrderStateChangedDate.String(), formats); err != nil {
		return err
	}

	return nil
}

var orderDetailsTypePurchaseOrderTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["RegularOrder","ConsignedOrder","NewProductIntroduction","RushOrder"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		orderDetailsTypePurchaseOrderTypePropEnum = append(orderDetailsTypePurchaseOrderTypePropEnum, v)
	}
}

const (

	// OrderDetailsPurchaseOrderTypeRegularOrder captures enum value "RegularOrder"
	OrderDetailsPurchaseOrderTypeRegularOrder string = "RegularOrder"

	// OrderDetailsPurchaseOrderTypeConsignedOrder captures enum value "ConsignedOrder"
	OrderDetailsPurchaseOrderTypeConsignedOrder string = "ConsignedOrder"

	// OrderDetailsPurchaseOrderTypeNewProductIntroduction captures enum value "NewProductIntroduction"
	OrderDetailsPurchaseOrderTypeNewProductIntroduction string = "NewProductIntroduction"

	// OrderDetailsPurchaseOrderTypeRushOrder captures enum value "RushOrder"
	OrderDetailsPurchaseOrderTypeRushOrder string = "RushOrder"
)

// prop value enum
func (m *OrderDetails) validatePurchaseOrderTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, orderDetailsTypePurchaseOrderTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *OrderDetails) validatePurchaseOrderType(formats strfmt.Registry) error {
	if swag.IsZero(m.PurchaseOrderType) { // not required
		return nil
	}

	// value enum
	if err := m.validatePurchaseOrderTypeEnum("purchaseOrderType", "body", m.PurchaseOrderType); err != nil {
		return err
	}

	return nil
}

func (m *OrderDetails) validateSellingParty(formats strfmt.Registry) error {
	if swag.IsZero(m.SellingParty) { // not required
		return nil
	}

	if m.SellingParty != nil {
		if err := m.SellingParty.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sellingParty")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sellingParty")
			}
			return err
		}
	}

	return nil
}

func (m *OrderDetails) validateShipToParty(formats strfmt.Registry) error {
	if swag.IsZero(m.ShipToParty) { // not required
		return nil
	}

	if m.ShipToParty != nil {
		if err := m.ShipToParty.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("shipToParty")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("shipToParty")
			}
			return err
		}
	}

	return nil
}

func (m *OrderDetails) validateShipWindow(formats strfmt.Registry) error {
	if swag.IsZero(m.ShipWindow) { // not required
		return nil
	}

	if err := m.ShipWindow.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("shipWindow")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("shipWindow")
		}
		return err
	}

	return nil
}

// ContextValidate validate this order details based on the context it is used
func (m *OrderDetails) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBillToParty(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateBuyingParty(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDeliveryWindow(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateImportDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateItems(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSellingParty(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateShipToParty(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateShipWindow(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OrderDetails) contextValidateBillToParty(ctx context.Context, formats strfmt.Registry) error {

	if m.BillToParty != nil {
		if err := m.BillToParty.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("billToParty")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("billToParty")
			}
			return err
		}
	}

	return nil
}

func (m *OrderDetails) contextValidateBuyingParty(ctx context.Context, formats strfmt.Registry) error {

	if m.BuyingParty != nil {
		if err := m.BuyingParty.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("buyingParty")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("buyingParty")
			}
			return err
		}
	}

	return nil
}

func (m *OrderDetails) contextValidateDeliveryWindow(ctx context.Context, formats strfmt.Registry) error {

	if err := m.DeliveryWindow.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("deliveryWindow")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("deliveryWindow")
		}
		return err
	}

	return nil
}

func (m *OrderDetails) contextValidateImportDetails(ctx context.Context, formats strfmt.Registry) error {

	if m.ImportDetails != nil {
		if err := m.ImportDetails.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("importDetails")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("importDetails")
			}
			return err
		}
	}

	return nil
}

func (m *OrderDetails) contextValidateItems(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Items); i++ {

		if m.Items[i] != nil {
			if err := m.Items[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("items" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *OrderDetails) contextValidateSellingParty(ctx context.Context, formats strfmt.Registry) error {

	if m.SellingParty != nil {
		if err := m.SellingParty.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sellingParty")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sellingParty")
			}
			return err
		}
	}

	return nil
}

func (m *OrderDetails) contextValidateShipToParty(ctx context.Context, formats strfmt.Registry) error {

	if m.ShipToParty != nil {
		if err := m.ShipToParty.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("shipToParty")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("shipToParty")
			}
			return err
		}
	}

	return nil
}

func (m *OrderDetails) contextValidateShipWindow(ctx context.Context, formats strfmt.Registry) error {

	if err := m.ShipWindow.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("shipWindow")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("shipWindow")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OrderDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OrderDetails) UnmarshalBinary(b []byte) error {
	var res OrderDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
