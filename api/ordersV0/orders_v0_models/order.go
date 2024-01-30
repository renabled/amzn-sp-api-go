// Code generated by go-swagger; DO NOT EDIT.

package orders_v0_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Order Order information.
//
// swagger:model Order
type Order struct {

	// An Amazon-defined order identifier, in 3-7-7 format.
	// Required: true
	AmazonOrderID *string `json:"AmazonOrderId"`

	// Contains information regarding the Shipping Settings Automaton program, such as whether the order's shipping settings were generated automatically, and what those settings are.
	AutomatedShippingSettings *AutomatedShippingSettings `json:"AutomatedShippingSettings,omitempty"`

	// buyer info
	BuyerInfo *BuyerInfo `json:"BuyerInfo,omitempty"`

	// The buyer's invoicing preference. Available only in the TR marketplace.
	// Enum: [INDIVIDUAL BUSINESS]
	BuyerInvoicePreference string `json:"BuyerInvoicePreference,omitempty"`

	// Contains the business invoice tax information.
	BuyerTaxInformation *BuyerTaxInformation `json:"BuyerTaxInformation,omitempty"`

	// Custom ship label for Checkout by Amazon (CBA).
	CbaDisplayableShippingLabel string `json:"CbaDisplayableShippingLabel,omitempty"`

	// The recommended location for the seller to ship the items from. It is calculated at checkout. The seller may or may not choose to ship from this location.
	DefaultShipFromLocationAddress *Address `json:"DefaultShipFromLocationAddress,omitempty"`

	// The start of the time period within which you have committed to fulfill the order. In <a href='https://developer-docs.amazon.com/sp-api/docs/iso-8601'>ISO 8601</a> date time format. Returned only for seller-fulfilled orders.
	EarliestDeliveryDate string `json:"EarliestDeliveryDate,omitempty"`

	// The start of the time period within which you have committed to ship the order. In <a href='https://developer-docs.amazon.com/sp-api/docs/iso-8601'>ISO 8601</a> date time format. Returned only for seller-fulfilled orders.
	//
	// __Note__: EarliestShipDate might not be returned for orders placed before February 1, 2013.
	EarliestShipDate string `json:"EarliestShipDate,omitempty"`

	// The status of the Amazon Easy Ship order. This property is included only for Amazon Easy Ship orders.
	EasyShipShipmentStatus EasyShipShipmentStatus `json:"EasyShipShipmentStatus,omitempty"`

	// The status of the electronic invoice.
	ElectronicInvoiceStatus ElectronicInvoiceStatus `json:"ElectronicInvoiceStatus,omitempty"`

	// Whether the order was fulfilled by Amazon (AFN) or by the seller (MFN).
	// Enum: [MFN AFN]
	FulfillmentChannel string `json:"FulfillmentChannel,omitempty"`

	// Contains the instructions about the fulfillment like where should it be fulfilled from.
	FulfillmentInstruction *FulfillmentInstruction `json:"FulfillmentInstruction,omitempty"`

	// Whether the order contains regulated items which may require additional approval steps before being fulfilled.
	HasRegulatedItems bool `json:"HasRegulatedItems,omitempty"`

	// When true, this order is marked to be delivered to an Access Point. The access location is chosen by the customer. Access Points include Amazon Hub Lockers, Amazon Hub Counters, and pickup points operated by carriers.
	IsAccessPointOrder bool `json:"IsAccessPointOrder,omitempty"`

	// When true, the order is an Amazon Business order. An Amazon Business order is an order where the buyer is a Verified Business Buyer.
	IsBusinessOrder bool `json:"IsBusinessOrder,omitempty"`

	// When true, the estimated ship date is set for the order. Returned only for Sourcing on Demand orders.
	IsEstimatedShipDateSet bool `json:"IsEstimatedShipDateSet,omitempty"`

	// When true, the order is a GlobalExpress order.
	IsGlobalExpressEnabled bool `json:"IsGlobalExpressEnabled,omitempty"`

	// When true, the item within this order was bought and re-sold by Amazon Business EU SARL (ABEU). By buying and instantly re-selling your items, ABEU becomes the seller of record, making your inventory available for sale to customers who would not otherwise purchase from a third-party seller.
	IsIBA bool `json:"IsIBA,omitempty"`

	// When true, this order is marked to be picked up from a store rather than delivered.
	IsISPU bool `json:"IsISPU,omitempty"`

	// When true, the order has a Premium Shipping Service Level Agreement. For more information about Premium Shipping orders, see "Premium Shipping Options" in the Seller Central Help for your marketplace.
	IsPremiumOrder bool `json:"IsPremiumOrder,omitempty"`

	// When true, the order is a seller-fulfilled Amazon Prime order.
	IsPrime bool `json:"IsPrime,omitempty"`

	// When true, this is a replacement order.
	IsReplacementOrder bool `json:"IsReplacementOrder,omitempty"`

	// When true, the item within this order was bought and re-sold by Amazon Business EU SARL (ABEU). By buying and instantly re-selling your items, ABEU becomes the seller of record, making your inventory available for sale to customers who would not otherwise purchase from a third-party seller.
	IsSoldByAB bool `json:"IsSoldByAB,omitempty"`

	// The date when the order was last updated.
	//
	// __Note__: LastUpdateDate is returned with an incorrect date for orders that were last updated before 2009-04-01.
	// Required: true
	LastUpdateDate *string `json:"LastUpdateDate"`

	// The end of the time period within which you have committed to fulfill the order. In <a href='https://developer-docs.amazon.com/sp-api/docs/iso-8601'>ISO 8601</a> date time format. Returned only for seller-fulfilled orders that do not have a PendingAvailability, Pending, or Canceled status.
	LatestDeliveryDate string `json:"LatestDeliveryDate,omitempty"`

	// The end of the time period within which you have committed to ship the order. In <a href='https://developer-docs.amazon.com/sp-api/docs/iso-8601'>ISO 8601</a> date time format. Returned only for seller-fulfilled orders.
	//
	// __Note__: LatestShipDate might not be returned for orders placed before February 1, 2013.
	LatestShipDate string `json:"LatestShipDate,omitempty"`

	// The identifier for the marketplace where the order was placed.
	MarketplaceID string `json:"MarketplaceId,omitempty"`

	// Tax information about the marketplace.
	MarketplaceTaxInfo *MarketplaceTaxInfo `json:"MarketplaceTaxInfo,omitempty"`

	// The number of items shipped.
	NumberOfItemsShipped int64 `json:"NumberOfItemsShipped,omitempty"`

	// The number of items unshipped.
	NumberOfItemsUnshipped int64 `json:"NumberOfItemsUnshipped,omitempty"`

	// The order channel of the first item in the order.
	OrderChannel string `json:"OrderChannel,omitempty"`

	// The current order status.
	// Required: true
	// Enum: [Pending Unshipped PartiallyShipped Shipped Canceled Unfulfillable InvoiceUnconfirmed PendingAvailability]
	OrderStatus *string `json:"OrderStatus"`

	// The total charge for this order.
	OrderTotal *Money `json:"OrderTotal,omitempty"`

	// The type of the order.
	// Enum: [StandardOrder LongLeadTimeOrder Preorder BackOrder SourcingOnDemandOrder]
	OrderType string `json:"OrderType,omitempty"`

	// Information about sub-payment methods for a Cash On Delivery (COD) order.
	//
	// __Note__: For a COD order that is paid for using one sub-payment method, one PaymentExecutionDetailItem object is returned, with PaymentExecutionDetailItem/PaymentMethod = COD. For a COD order that is paid for using multiple sub-payment methods, two or more PaymentExecutionDetailItem objects are returned.
	PaymentExecutionDetail PaymentExecutionDetailItemList `json:"PaymentExecutionDetail,omitempty"`

	// The payment method for the order. This property is limited to Cash On Delivery (COD) and Convenience Store (CVS) payment methods. Unless you need the specific COD payment information provided by the PaymentExecutionDetailItem object, we recommend using the PaymentMethodDetails property to get payment method information.
	// Enum: [COD CVS Other]
	PaymentMethod string `json:"PaymentMethod,omitempty"`

	// A list of payment methods for the order.
	PaymentMethodDetails PaymentMethodDetailItemList `json:"PaymentMethodDetails,omitempty"`

	// Indicates the date by which the seller must respond to the buyer with an estimated ship date. Returned only for Sourcing on Demand orders.
	PromiseResponseDueDate string `json:"PromiseResponseDueDate,omitempty"`

	// The date when the order was created.
	// Required: true
	PurchaseDate *string `json:"PurchaseDate"`

	// The order ID value for the order that is being replaced. Returned only if IsReplacementOrder = true.
	ReplacedOrderID string `json:"ReplacedOrderId,omitempty"`

	// The sales channel of the first item in the order.
	SalesChannel string `json:"SalesChannel,omitempty"`

	// The seller’s friendly name registered in the marketplace.
	SellerDisplayName string `json:"SellerDisplayName,omitempty"`

	// A seller-defined order identifier.
	SellerOrderID string `json:"SellerOrderId,omitempty"`

	// The shipment service level of the order.
	ShipServiceLevel string `json:"ShipServiceLevel,omitempty"`

	// The shipment service level category of the order.
	//
	// Possible values: Expedited, FreeEconomy, NextDay, Priority, SameDay, SecondDay, Scheduled, Standard.
	ShipmentServiceLevelCategory string `json:"ShipmentServiceLevelCategory,omitempty"`

	// shipping address
	ShippingAddress *Address `json:"ShippingAddress,omitempty"`
}

// Validate validates this order
func (m *Order) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmazonOrderID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAutomatedShippingSettings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBuyerInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBuyerInvoicePreference(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBuyerTaxInformation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDefaultShipFromLocationAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEasyShipShipmentStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateElectronicInvoiceStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFulfillmentChannel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFulfillmentInstruction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdateDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMarketplaceTaxInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrderStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrderTotal(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrderType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePaymentExecutionDetail(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePaymentMethod(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePaymentMethodDetails(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePurchaseDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShippingAddress(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Order) validateAmazonOrderID(formats strfmt.Registry) error {

	if err := validate.Required("AmazonOrderId", "body", m.AmazonOrderID); err != nil {
		return err
	}

	return nil
}

func (m *Order) validateAutomatedShippingSettings(formats strfmt.Registry) error {
	if swag.IsZero(m.AutomatedShippingSettings) { // not required
		return nil
	}

	if m.AutomatedShippingSettings != nil {
		if err := m.AutomatedShippingSettings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("AutomatedShippingSettings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("AutomatedShippingSettings")
			}
			return err
		}
	}

	return nil
}

func (m *Order) validateBuyerInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.BuyerInfo) { // not required
		return nil
	}

	if m.BuyerInfo != nil {
		if err := m.BuyerInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("BuyerInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("BuyerInfo")
			}
			return err
		}
	}

	return nil
}

var orderTypeBuyerInvoicePreferencePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["INDIVIDUAL","BUSINESS"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		orderTypeBuyerInvoicePreferencePropEnum = append(orderTypeBuyerInvoicePreferencePropEnum, v)
	}
}

const (

	// OrderBuyerInvoicePreferenceINDIVIDUAL captures enum value "INDIVIDUAL"
	OrderBuyerInvoicePreferenceINDIVIDUAL string = "INDIVIDUAL"

	// OrderBuyerInvoicePreferenceBUSINESS captures enum value "BUSINESS"
	OrderBuyerInvoicePreferenceBUSINESS string = "BUSINESS"
)

// prop value enum
func (m *Order) validateBuyerInvoicePreferenceEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, orderTypeBuyerInvoicePreferencePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Order) validateBuyerInvoicePreference(formats strfmt.Registry) error {
	if swag.IsZero(m.BuyerInvoicePreference) { // not required
		return nil
	}

	// value enum
	if err := m.validateBuyerInvoicePreferenceEnum("BuyerInvoicePreference", "body", m.BuyerInvoicePreference); err != nil {
		return err
	}

	return nil
}

func (m *Order) validateBuyerTaxInformation(formats strfmt.Registry) error {
	if swag.IsZero(m.BuyerTaxInformation) { // not required
		return nil
	}

	if m.BuyerTaxInformation != nil {
		if err := m.BuyerTaxInformation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("BuyerTaxInformation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("BuyerTaxInformation")
			}
			return err
		}
	}

	return nil
}

func (m *Order) validateDefaultShipFromLocationAddress(formats strfmt.Registry) error {
	if swag.IsZero(m.DefaultShipFromLocationAddress) { // not required
		return nil
	}

	if m.DefaultShipFromLocationAddress != nil {
		if err := m.DefaultShipFromLocationAddress.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("DefaultShipFromLocationAddress")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("DefaultShipFromLocationAddress")
			}
			return err
		}
	}

	return nil
}

func (m *Order) validateEasyShipShipmentStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.EasyShipShipmentStatus) { // not required
		return nil
	}

	if err := m.EasyShipShipmentStatus.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("EasyShipShipmentStatus")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("EasyShipShipmentStatus")
		}
		return err
	}

	return nil
}

func (m *Order) validateElectronicInvoiceStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.ElectronicInvoiceStatus) { // not required
		return nil
	}

	if err := m.ElectronicInvoiceStatus.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("ElectronicInvoiceStatus")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("ElectronicInvoiceStatus")
		}
		return err
	}

	return nil
}

var orderTypeFulfillmentChannelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["MFN","AFN"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		orderTypeFulfillmentChannelPropEnum = append(orderTypeFulfillmentChannelPropEnum, v)
	}
}

const (

	// OrderFulfillmentChannelMFN captures enum value "MFN"
	OrderFulfillmentChannelMFN string = "MFN"

	// OrderFulfillmentChannelAFN captures enum value "AFN"
	OrderFulfillmentChannelAFN string = "AFN"
)

// prop value enum
func (m *Order) validateFulfillmentChannelEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, orderTypeFulfillmentChannelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Order) validateFulfillmentChannel(formats strfmt.Registry) error {
	if swag.IsZero(m.FulfillmentChannel) { // not required
		return nil
	}

	// value enum
	if err := m.validateFulfillmentChannelEnum("FulfillmentChannel", "body", m.FulfillmentChannel); err != nil {
		return err
	}

	return nil
}

func (m *Order) validateFulfillmentInstruction(formats strfmt.Registry) error {
	if swag.IsZero(m.FulfillmentInstruction) { // not required
		return nil
	}

	if m.FulfillmentInstruction != nil {
		if err := m.FulfillmentInstruction.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("FulfillmentInstruction")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("FulfillmentInstruction")
			}
			return err
		}
	}

	return nil
}

func (m *Order) validateLastUpdateDate(formats strfmt.Registry) error {

	if err := validate.Required("LastUpdateDate", "body", m.LastUpdateDate); err != nil {
		return err
	}

	return nil
}

func (m *Order) validateMarketplaceTaxInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.MarketplaceTaxInfo) { // not required
		return nil
	}

	if m.MarketplaceTaxInfo != nil {
		if err := m.MarketplaceTaxInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("MarketplaceTaxInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("MarketplaceTaxInfo")
			}
			return err
		}
	}

	return nil
}

var orderTypeOrderStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Pending","Unshipped","PartiallyShipped","Shipped","Canceled","Unfulfillable","InvoiceUnconfirmed","PendingAvailability"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		orderTypeOrderStatusPropEnum = append(orderTypeOrderStatusPropEnum, v)
	}
}

const (

	// OrderOrderStatusPending captures enum value "Pending"
	OrderOrderStatusPending string = "Pending"

	// OrderOrderStatusUnshipped captures enum value "Unshipped"
	OrderOrderStatusUnshipped string = "Unshipped"

	// OrderOrderStatusPartiallyShipped captures enum value "PartiallyShipped"
	OrderOrderStatusPartiallyShipped string = "PartiallyShipped"

	// OrderOrderStatusShipped captures enum value "Shipped"
	OrderOrderStatusShipped string = "Shipped"

	// OrderOrderStatusCanceled captures enum value "Canceled"
	OrderOrderStatusCanceled string = "Canceled"

	// OrderOrderStatusUnfulfillable captures enum value "Unfulfillable"
	OrderOrderStatusUnfulfillable string = "Unfulfillable"

	// OrderOrderStatusInvoiceUnconfirmed captures enum value "InvoiceUnconfirmed"
	OrderOrderStatusInvoiceUnconfirmed string = "InvoiceUnconfirmed"

	// OrderOrderStatusPendingAvailability captures enum value "PendingAvailability"
	OrderOrderStatusPendingAvailability string = "PendingAvailability"
)

// prop value enum
func (m *Order) validateOrderStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, orderTypeOrderStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Order) validateOrderStatus(formats strfmt.Registry) error {

	if err := validate.Required("OrderStatus", "body", m.OrderStatus); err != nil {
		return err
	}

	// value enum
	if err := m.validateOrderStatusEnum("OrderStatus", "body", *m.OrderStatus); err != nil {
		return err
	}

	return nil
}

func (m *Order) validateOrderTotal(formats strfmt.Registry) error {
	if swag.IsZero(m.OrderTotal) { // not required
		return nil
	}

	if m.OrderTotal != nil {
		if err := m.OrderTotal.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("OrderTotal")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("OrderTotal")
			}
			return err
		}
	}

	return nil
}

var orderTypeOrderTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["StandardOrder","LongLeadTimeOrder","Preorder","BackOrder","SourcingOnDemandOrder"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		orderTypeOrderTypePropEnum = append(orderTypeOrderTypePropEnum, v)
	}
}

const (

	// OrderOrderTypeStandardOrder captures enum value "StandardOrder"
	OrderOrderTypeStandardOrder string = "StandardOrder"

	// OrderOrderTypeLongLeadTimeOrder captures enum value "LongLeadTimeOrder"
	OrderOrderTypeLongLeadTimeOrder string = "LongLeadTimeOrder"

	// OrderOrderTypePreorder captures enum value "Preorder"
	OrderOrderTypePreorder string = "Preorder"

	// OrderOrderTypeBackOrder captures enum value "BackOrder"
	OrderOrderTypeBackOrder string = "BackOrder"

	// OrderOrderTypeSourcingOnDemandOrder captures enum value "SourcingOnDemandOrder"
	OrderOrderTypeSourcingOnDemandOrder string = "SourcingOnDemandOrder"
)

// prop value enum
func (m *Order) validateOrderTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, orderTypeOrderTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Order) validateOrderType(formats strfmt.Registry) error {
	if swag.IsZero(m.OrderType) { // not required
		return nil
	}

	// value enum
	if err := m.validateOrderTypeEnum("OrderType", "body", m.OrderType); err != nil {
		return err
	}

	return nil
}

func (m *Order) validatePaymentExecutionDetail(formats strfmt.Registry) error {
	if swag.IsZero(m.PaymentExecutionDetail) { // not required
		return nil
	}

	if err := m.PaymentExecutionDetail.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("PaymentExecutionDetail")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("PaymentExecutionDetail")
		}
		return err
	}

	return nil
}

var orderTypePaymentMethodPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["COD","CVS","Other"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		orderTypePaymentMethodPropEnum = append(orderTypePaymentMethodPropEnum, v)
	}
}

const (

	// OrderPaymentMethodCOD captures enum value "COD"
	OrderPaymentMethodCOD string = "COD"

	// OrderPaymentMethodCVS captures enum value "CVS"
	OrderPaymentMethodCVS string = "CVS"

	// OrderPaymentMethodOther captures enum value "Other"
	OrderPaymentMethodOther string = "Other"
)

// prop value enum
func (m *Order) validatePaymentMethodEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, orderTypePaymentMethodPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Order) validatePaymentMethod(formats strfmt.Registry) error {
	if swag.IsZero(m.PaymentMethod) { // not required
		return nil
	}

	// value enum
	if err := m.validatePaymentMethodEnum("PaymentMethod", "body", m.PaymentMethod); err != nil {
		return err
	}

	return nil
}

func (m *Order) validatePaymentMethodDetails(formats strfmt.Registry) error {
	if swag.IsZero(m.PaymentMethodDetails) { // not required
		return nil
	}

	if err := m.PaymentMethodDetails.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("PaymentMethodDetails")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("PaymentMethodDetails")
		}
		return err
	}

	return nil
}

func (m *Order) validatePurchaseDate(formats strfmt.Registry) error {

	if err := validate.Required("PurchaseDate", "body", m.PurchaseDate); err != nil {
		return err
	}

	return nil
}

func (m *Order) validateShippingAddress(formats strfmt.Registry) error {
	if swag.IsZero(m.ShippingAddress) { // not required
		return nil
	}

	if m.ShippingAddress != nil {
		if err := m.ShippingAddress.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ShippingAddress")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ShippingAddress")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this order based on the context it is used
func (m *Order) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAutomatedShippingSettings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateBuyerInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateBuyerTaxInformation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDefaultShipFromLocationAddress(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEasyShipShipmentStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateElectronicInvoiceStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFulfillmentInstruction(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMarketplaceTaxInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOrderTotal(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePaymentExecutionDetail(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePaymentMethodDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateShippingAddress(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Order) contextValidateAutomatedShippingSettings(ctx context.Context, formats strfmt.Registry) error {

	if m.AutomatedShippingSettings != nil {
		if err := m.AutomatedShippingSettings.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("AutomatedShippingSettings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("AutomatedShippingSettings")
			}
			return err
		}
	}

	return nil
}

func (m *Order) contextValidateBuyerInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.BuyerInfo != nil {
		if err := m.BuyerInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("BuyerInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("BuyerInfo")
			}
			return err
		}
	}

	return nil
}

func (m *Order) contextValidateBuyerTaxInformation(ctx context.Context, formats strfmt.Registry) error {

	if m.BuyerTaxInformation != nil {
		if err := m.BuyerTaxInformation.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("BuyerTaxInformation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("BuyerTaxInformation")
			}
			return err
		}
	}

	return nil
}

func (m *Order) contextValidateDefaultShipFromLocationAddress(ctx context.Context, formats strfmt.Registry) error {

	if m.DefaultShipFromLocationAddress != nil {
		if err := m.DefaultShipFromLocationAddress.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("DefaultShipFromLocationAddress")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("DefaultShipFromLocationAddress")
			}
			return err
		}
	}

	return nil
}

func (m *Order) contextValidateEasyShipShipmentStatus(ctx context.Context, formats strfmt.Registry) error {

	if err := m.EasyShipShipmentStatus.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("EasyShipShipmentStatus")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("EasyShipShipmentStatus")
		}
		return err
	}

	return nil
}

func (m *Order) contextValidateElectronicInvoiceStatus(ctx context.Context, formats strfmt.Registry) error {

	if err := m.ElectronicInvoiceStatus.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("ElectronicInvoiceStatus")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("ElectronicInvoiceStatus")
		}
		return err
	}

	return nil
}

func (m *Order) contextValidateFulfillmentInstruction(ctx context.Context, formats strfmt.Registry) error {

	if m.FulfillmentInstruction != nil {
		if err := m.FulfillmentInstruction.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("FulfillmentInstruction")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("FulfillmentInstruction")
			}
			return err
		}
	}

	return nil
}

func (m *Order) contextValidateMarketplaceTaxInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.MarketplaceTaxInfo != nil {
		if err := m.MarketplaceTaxInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("MarketplaceTaxInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("MarketplaceTaxInfo")
			}
			return err
		}
	}

	return nil
}

func (m *Order) contextValidateOrderTotal(ctx context.Context, formats strfmt.Registry) error {

	if m.OrderTotal != nil {
		if err := m.OrderTotal.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("OrderTotal")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("OrderTotal")
			}
			return err
		}
	}

	return nil
}

func (m *Order) contextValidatePaymentExecutionDetail(ctx context.Context, formats strfmt.Registry) error {

	if err := m.PaymentExecutionDetail.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("PaymentExecutionDetail")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("PaymentExecutionDetail")
		}
		return err
	}

	return nil
}

func (m *Order) contextValidatePaymentMethodDetails(ctx context.Context, formats strfmt.Registry) error {

	if err := m.PaymentMethodDetails.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("PaymentMethodDetails")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("PaymentMethodDetails")
		}
		return err
	}

	return nil
}

func (m *Order) contextValidateShippingAddress(ctx context.Context, formats strfmt.Registry) error {

	if m.ShippingAddress != nil {
		if err := m.ShippingAddress.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ShippingAddress")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ShippingAddress")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Order) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Order) UnmarshalBinary(b []byte) error {
	var res Order
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
