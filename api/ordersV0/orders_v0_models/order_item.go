// Code generated by go-swagger; DO NOT EDIT.

package orders_v0_models

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

// OrderItem A single order item.
//
// swagger:model OrderItem
type OrderItem struct {

	// The item's Amazon Standard Identification Number (ASIN).
	// Required: true
	ASIN *string `json:"ASIN"`

	// Contains the list of programs that are associated with an item.
	AmazonPrograms *AmazonPrograms `json:"AmazonPrograms,omitempty"`

	// A list of associated items that a customer has purchased with a product. For example, a tire installation service purchased with tires.
	AssociatedItems []*AssociatedItem `json:"AssociatedItems"`

	// A single item's buyer information.
	//
	// **Note**: The `BuyerInfo` contains restricted data. Use the Restricted Data Token (RDT) and restricted SPDS roles to access the restricted data in `BuyerInfo`. For example, `BuyerCustomizedInfo` and `GiftMessageText`.
	BuyerInfo *ItemBuyerInfo `json:"BuyerInfo,omitempty"`

	// Information about whether or not a buyer requested cancellation.
	BuyerRequestedCancel *BuyerRequestedCancel `json:"BuyerRequestedCancel,omitempty"`

	// The fee charged for COD service.
	CODFee *Money `json:"CODFee,omitempty"`

	// The discount on the COD fee.
	CODFeeDiscount *Money `json:"CODFeeDiscount,omitempty"`

	// The condition of the item.
	//
	// **Possible values**: `New`, `Used`, `Collectible`, `Refurbished`, `Preorder`, and `Club`.
	ConditionID string `json:"ConditionId,omitempty"`

	// The condition of the item, as described by the seller.
	ConditionNote string `json:"ConditionNote,omitempty"`

	// The subcondition of the item.
	//
	// **Possible values**: `New`, `Mint`, `Very Good`, `Good`, `Acceptable`, `Poor`, `Club`, `OEM`, `Warranty`, `Refurbished Warranty`, `Refurbished`, `Open Box`, `Any`, and `Other`.
	ConditionSubtypeID string `json:"ConditionSubtypeId,omitempty"`

	// The category of deemed reseller. This applies to selling partners that are not based in the EU and is used to help them meet the VAT Deemed Reseller tax laws in the EU and UK.
	// Enum: [IOSS UOSS]
	DeemedResellerCategory string `json:"DeemedResellerCategory,omitempty"`

	// The IOSS number of the marketplace. Sellers shipping to the EU from outside the EU must provide this IOSS number to their carrier when Amazon has collected the VAT on the sale.
	IossNumber string `json:"IossNumber,omitempty"`

	// Indicates whether the item is a gift.
	//
	// **Possible values**: `true` and `false`.
	IsGift string `json:"IsGift,omitempty"`

	// When true, the ASIN is enrolled in Transparency. The Transparency serial number that you must submit is determined by:
	//
	// **1D or 2D Barcode:** This has a **T** logo. Submit either the 29-character alpha-numeric identifier beginning with **AZ** or **ZA**, or the 38-character Serialized Global Trade Item Number (SGTIN).
	// **2D Barcode SN:** Submit the 7- to 20-character serial number barcode, which likely has the prefix **SN**. The serial number is applied to the same side of the packaging as the GTIN (UPC/EAN/ISBN) barcode.
	// **QR code SN:** Submit the URL that the QR code generates.
	IsTransparency bool `json:"IsTransparency,omitempty"`

	// The selling price of the order item. Note that an order item is an item and a quantity. This means that the value of `ItemPrice` is equal to the selling price of the item multiplied by the quantity ordered. `ItemPrice` excludes `ShippingPrice` and GiftWrapPrice.
	ItemPrice *Money `json:"ItemPrice,omitempty"`

	// The tax on the item price.
	ItemTax *Money `json:"ItemTax,omitempty"`

	// Measurement information for the order item.
	Measurement *Measurement `json:"Measurement,omitempty"`

	// An Amazon-defined order item identifier.
	// Required: true
	OrderItemID *string `json:"OrderItemId"`

	// The number and value of Amazon Points granted with the purchase of an item.
	PointsGranted *PointsGrantedDetail `json:"PointsGranted,omitempty"`

	// Indicates that the selling price is a special price that is only available for Amazon Business orders. For more information about the Amazon Business Seller Program, refer to the [Amazon Business website](https://www.amazon.com/b2b/info/amazon-business).
	//
	// **Possible values**: `BusinessPrice`
	PriceDesignation string `json:"PriceDesignation,omitempty"`

	// The item's product information.
	ProductInfo *ProductInfoDetail `json:"ProductInfo,omitempty"`

	// The total of all promotional discounts in the offer.
	PromotionDiscount *Money `json:"PromotionDiscount,omitempty"`

	// The tax on the total of all promotional discounts in the offer.
	PromotionDiscountTax *Money `json:"PromotionDiscountTax,omitempty"`

	// promotion ids
	PromotionIds PromotionIDList `json:"PromotionIds,omitempty"`

	// The number of items in the order.
	// Required: true
	QuantityOrdered *int64 `json:"QuantityOrdered"`

	// The number of items shipped.
	QuantityShipped int64 `json:"QuantityShipped,omitempty"`

	// The end date of the scheduled delivery window in the time zone for the order destination. In [ISO 8601](https://developer-docs.amazon.com/sp-api/docs/iso-8601) date time format.
	ScheduledDeliveryEndDate string `json:"ScheduledDeliveryEndDate,omitempty"`

	// The start date of the scheduled delivery window in the time zone for the order destination. In [ISO 8601](https://developer-docs.amazon.com/sp-api/docs/iso-8601) date time format.
	ScheduledDeliveryStartDate string `json:"ScheduledDeliveryStartDate,omitempty"`

	// The item's seller stock keeping unit (SKU).
	SellerSKU string `json:"SellerSKU,omitempty"`

	// When true, the product type for this item has a serial number.
	//
	//  Only returned for Amazon Easy Ship orders.
	SerialNumberRequired bool `json:"SerialNumberRequired,omitempty"`

	// A list of serial numbers for electronic products that are shipped to customers. Returned for FBA orders only.
	SerialNumbers []string `json:"SerialNumbers"`

	// Shipping constraints applicable to this order.
	ShippingConstraints *ShippingConstraints `json:"ShippingConstraints,omitempty"`

	// The discount on the shipping price.
	ShippingDiscount *Money `json:"ShippingDiscount,omitempty"`

	// The tax on the discount on the shipping price.
	ShippingDiscountTax *Money `json:"ShippingDiscountTax,omitempty"`

	// The item's shipping price.
	ShippingPrice *Money `json:"ShippingPrice,omitempty"`

	// The tax on the shipping price.
	ShippingTax *Money `json:"ShippingTax,omitempty"`

	// The store chain store identifier. Linked to a specific store in a store chain.
	StoreChainStoreID string `json:"StoreChainStoreId,omitempty"`

	// Substitution preferences for the order item. This is an optional field that is only present if a seller supports substitutions, as is the case with some grocery sellers.
	SubstitutionPreferences *SubstitutionPreferences `json:"SubstitutionPreferences,omitempty"`

	// Information about withheld taxes.
	TaxCollection *TaxCollection `json:"TaxCollection,omitempty"`

	// The item's name.
	Title string `json:"Title,omitempty"`
}

// Validate validates this order item
func (m *OrderItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateASIN(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAmazonPrograms(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAssociatedItems(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBuyerInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBuyerRequestedCancel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCODFee(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCODFeeDiscount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeemedResellerCategory(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateItemPrice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateItemTax(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMeasurement(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrderItemID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePointsGranted(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProductInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePromotionDiscount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePromotionDiscountTax(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePromotionIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQuantityOrdered(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShippingConstraints(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShippingDiscount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShippingDiscountTax(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShippingPrice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShippingTax(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubstitutionPreferences(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTaxCollection(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OrderItem) validateASIN(formats strfmt.Registry) error {

	if err := validate.Required("ASIN", "body", m.ASIN); err != nil {
		return err
	}

	return nil
}

func (m *OrderItem) validateAmazonPrograms(formats strfmt.Registry) error {
	if swag.IsZero(m.AmazonPrograms) { // not required
		return nil
	}

	if m.AmazonPrograms != nil {
		if err := m.AmazonPrograms.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("AmazonPrograms")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("AmazonPrograms")
			}
			return err
		}
	}

	return nil
}

func (m *OrderItem) validateAssociatedItems(formats strfmt.Registry) error {
	if swag.IsZero(m.AssociatedItems) { // not required
		return nil
	}

	for i := 0; i < len(m.AssociatedItems); i++ {
		if swag.IsZero(m.AssociatedItems[i]) { // not required
			continue
		}

		if m.AssociatedItems[i] != nil {
			if err := m.AssociatedItems[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("AssociatedItems" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("AssociatedItems" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *OrderItem) validateBuyerInfo(formats strfmt.Registry) error {
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

func (m *OrderItem) validateBuyerRequestedCancel(formats strfmt.Registry) error {
	if swag.IsZero(m.BuyerRequestedCancel) { // not required
		return nil
	}

	if m.BuyerRequestedCancel != nil {
		if err := m.BuyerRequestedCancel.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("BuyerRequestedCancel")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("BuyerRequestedCancel")
			}
			return err
		}
	}

	return nil
}

func (m *OrderItem) validateCODFee(formats strfmt.Registry) error {
	if swag.IsZero(m.CODFee) { // not required
		return nil
	}

	if m.CODFee != nil {
		if err := m.CODFee.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("CODFee")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("CODFee")
			}
			return err
		}
	}

	return nil
}

func (m *OrderItem) validateCODFeeDiscount(formats strfmt.Registry) error {
	if swag.IsZero(m.CODFeeDiscount) { // not required
		return nil
	}

	if m.CODFeeDiscount != nil {
		if err := m.CODFeeDiscount.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("CODFeeDiscount")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("CODFeeDiscount")
			}
			return err
		}
	}

	return nil
}

var orderItemTypeDeemedResellerCategoryPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["IOSS","UOSS"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		orderItemTypeDeemedResellerCategoryPropEnum = append(orderItemTypeDeemedResellerCategoryPropEnum, v)
	}
}

const (

	// OrderItemDeemedResellerCategoryIOSS captures enum value "IOSS"
	OrderItemDeemedResellerCategoryIOSS string = "IOSS"

	// OrderItemDeemedResellerCategoryUOSS captures enum value "UOSS"
	OrderItemDeemedResellerCategoryUOSS string = "UOSS"
)

// prop value enum
func (m *OrderItem) validateDeemedResellerCategoryEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, orderItemTypeDeemedResellerCategoryPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *OrderItem) validateDeemedResellerCategory(formats strfmt.Registry) error {
	if swag.IsZero(m.DeemedResellerCategory) { // not required
		return nil
	}

	// value enum
	if err := m.validateDeemedResellerCategoryEnum("DeemedResellerCategory", "body", m.DeemedResellerCategory); err != nil {
		return err
	}

	return nil
}

func (m *OrderItem) validateItemPrice(formats strfmt.Registry) error {
	if swag.IsZero(m.ItemPrice) { // not required
		return nil
	}

	if m.ItemPrice != nil {
		if err := m.ItemPrice.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ItemPrice")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ItemPrice")
			}
			return err
		}
	}

	return nil
}

func (m *OrderItem) validateItemTax(formats strfmt.Registry) error {
	if swag.IsZero(m.ItemTax) { // not required
		return nil
	}

	if m.ItemTax != nil {
		if err := m.ItemTax.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ItemTax")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ItemTax")
			}
			return err
		}
	}

	return nil
}

func (m *OrderItem) validateMeasurement(formats strfmt.Registry) error {
	if swag.IsZero(m.Measurement) { // not required
		return nil
	}

	if m.Measurement != nil {
		if err := m.Measurement.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Measurement")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Measurement")
			}
			return err
		}
	}

	return nil
}

func (m *OrderItem) validateOrderItemID(formats strfmt.Registry) error {

	if err := validate.Required("OrderItemId", "body", m.OrderItemID); err != nil {
		return err
	}

	return nil
}

func (m *OrderItem) validatePointsGranted(formats strfmt.Registry) error {
	if swag.IsZero(m.PointsGranted) { // not required
		return nil
	}

	if m.PointsGranted != nil {
		if err := m.PointsGranted.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("PointsGranted")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("PointsGranted")
			}
			return err
		}
	}

	return nil
}

func (m *OrderItem) validateProductInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.ProductInfo) { // not required
		return nil
	}

	if m.ProductInfo != nil {
		if err := m.ProductInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ProductInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ProductInfo")
			}
			return err
		}
	}

	return nil
}

func (m *OrderItem) validatePromotionDiscount(formats strfmt.Registry) error {
	if swag.IsZero(m.PromotionDiscount) { // not required
		return nil
	}

	if m.PromotionDiscount != nil {
		if err := m.PromotionDiscount.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("PromotionDiscount")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("PromotionDiscount")
			}
			return err
		}
	}

	return nil
}

func (m *OrderItem) validatePromotionDiscountTax(formats strfmt.Registry) error {
	if swag.IsZero(m.PromotionDiscountTax) { // not required
		return nil
	}

	if m.PromotionDiscountTax != nil {
		if err := m.PromotionDiscountTax.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("PromotionDiscountTax")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("PromotionDiscountTax")
			}
			return err
		}
	}

	return nil
}

func (m *OrderItem) validatePromotionIds(formats strfmt.Registry) error {
	if swag.IsZero(m.PromotionIds) { // not required
		return nil
	}

	if err := m.PromotionIds.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("PromotionIds")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("PromotionIds")
		}
		return err
	}

	return nil
}

func (m *OrderItem) validateQuantityOrdered(formats strfmt.Registry) error {

	if err := validate.Required("QuantityOrdered", "body", m.QuantityOrdered); err != nil {
		return err
	}

	return nil
}

func (m *OrderItem) validateShippingConstraints(formats strfmt.Registry) error {
	if swag.IsZero(m.ShippingConstraints) { // not required
		return nil
	}

	if m.ShippingConstraints != nil {
		if err := m.ShippingConstraints.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ShippingConstraints")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ShippingConstraints")
			}
			return err
		}
	}

	return nil
}

func (m *OrderItem) validateShippingDiscount(formats strfmt.Registry) error {
	if swag.IsZero(m.ShippingDiscount) { // not required
		return nil
	}

	if m.ShippingDiscount != nil {
		if err := m.ShippingDiscount.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ShippingDiscount")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ShippingDiscount")
			}
			return err
		}
	}

	return nil
}

func (m *OrderItem) validateShippingDiscountTax(formats strfmt.Registry) error {
	if swag.IsZero(m.ShippingDiscountTax) { // not required
		return nil
	}

	if m.ShippingDiscountTax != nil {
		if err := m.ShippingDiscountTax.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ShippingDiscountTax")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ShippingDiscountTax")
			}
			return err
		}
	}

	return nil
}

func (m *OrderItem) validateShippingPrice(formats strfmt.Registry) error {
	if swag.IsZero(m.ShippingPrice) { // not required
		return nil
	}

	if m.ShippingPrice != nil {
		if err := m.ShippingPrice.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ShippingPrice")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ShippingPrice")
			}
			return err
		}
	}

	return nil
}

func (m *OrderItem) validateShippingTax(formats strfmt.Registry) error {
	if swag.IsZero(m.ShippingTax) { // not required
		return nil
	}

	if m.ShippingTax != nil {
		if err := m.ShippingTax.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ShippingTax")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ShippingTax")
			}
			return err
		}
	}

	return nil
}

func (m *OrderItem) validateSubstitutionPreferences(formats strfmt.Registry) error {
	if swag.IsZero(m.SubstitutionPreferences) { // not required
		return nil
	}

	if m.SubstitutionPreferences != nil {
		if err := m.SubstitutionPreferences.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("SubstitutionPreferences")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("SubstitutionPreferences")
			}
			return err
		}
	}

	return nil
}

func (m *OrderItem) validateTaxCollection(formats strfmt.Registry) error {
	if swag.IsZero(m.TaxCollection) { // not required
		return nil
	}

	if m.TaxCollection != nil {
		if err := m.TaxCollection.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("TaxCollection")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("TaxCollection")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this order item based on the context it is used
func (m *OrderItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAmazonPrograms(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAssociatedItems(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateBuyerInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateBuyerRequestedCancel(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCODFee(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCODFeeDiscount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateItemPrice(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateItemTax(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMeasurement(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePointsGranted(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProductInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePromotionDiscount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePromotionDiscountTax(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePromotionIds(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateShippingConstraints(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateShippingDiscount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateShippingDiscountTax(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateShippingPrice(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateShippingTax(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSubstitutionPreferences(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTaxCollection(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OrderItem) contextValidateAmazonPrograms(ctx context.Context, formats strfmt.Registry) error {

	if m.AmazonPrograms != nil {
		if err := m.AmazonPrograms.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("AmazonPrograms")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("AmazonPrograms")
			}
			return err
		}
	}

	return nil
}

func (m *OrderItem) contextValidateAssociatedItems(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AssociatedItems); i++ {

		if m.AssociatedItems[i] != nil {
			if err := m.AssociatedItems[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("AssociatedItems" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("AssociatedItems" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *OrderItem) contextValidateBuyerInfo(ctx context.Context, formats strfmt.Registry) error {

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

func (m *OrderItem) contextValidateBuyerRequestedCancel(ctx context.Context, formats strfmt.Registry) error {

	if m.BuyerRequestedCancel != nil {
		if err := m.BuyerRequestedCancel.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("BuyerRequestedCancel")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("BuyerRequestedCancel")
			}
			return err
		}
	}

	return nil
}

func (m *OrderItem) contextValidateCODFee(ctx context.Context, formats strfmt.Registry) error {

	if m.CODFee != nil {
		if err := m.CODFee.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("CODFee")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("CODFee")
			}
			return err
		}
	}

	return nil
}

func (m *OrderItem) contextValidateCODFeeDiscount(ctx context.Context, formats strfmt.Registry) error {

	if m.CODFeeDiscount != nil {
		if err := m.CODFeeDiscount.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("CODFeeDiscount")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("CODFeeDiscount")
			}
			return err
		}
	}

	return nil
}

func (m *OrderItem) contextValidateItemPrice(ctx context.Context, formats strfmt.Registry) error {

	if m.ItemPrice != nil {
		if err := m.ItemPrice.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ItemPrice")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ItemPrice")
			}
			return err
		}
	}

	return nil
}

func (m *OrderItem) contextValidateItemTax(ctx context.Context, formats strfmt.Registry) error {

	if m.ItemTax != nil {
		if err := m.ItemTax.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ItemTax")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ItemTax")
			}
			return err
		}
	}

	return nil
}

func (m *OrderItem) contextValidateMeasurement(ctx context.Context, formats strfmt.Registry) error {

	if m.Measurement != nil {
		if err := m.Measurement.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Measurement")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Measurement")
			}
			return err
		}
	}

	return nil
}

func (m *OrderItem) contextValidatePointsGranted(ctx context.Context, formats strfmt.Registry) error {

	if m.PointsGranted != nil {
		if err := m.PointsGranted.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("PointsGranted")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("PointsGranted")
			}
			return err
		}
	}

	return nil
}

func (m *OrderItem) contextValidateProductInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.ProductInfo != nil {
		if err := m.ProductInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ProductInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ProductInfo")
			}
			return err
		}
	}

	return nil
}

func (m *OrderItem) contextValidatePromotionDiscount(ctx context.Context, formats strfmt.Registry) error {

	if m.PromotionDiscount != nil {
		if err := m.PromotionDiscount.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("PromotionDiscount")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("PromotionDiscount")
			}
			return err
		}
	}

	return nil
}

func (m *OrderItem) contextValidatePromotionDiscountTax(ctx context.Context, formats strfmt.Registry) error {

	if m.PromotionDiscountTax != nil {
		if err := m.PromotionDiscountTax.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("PromotionDiscountTax")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("PromotionDiscountTax")
			}
			return err
		}
	}

	return nil
}

func (m *OrderItem) contextValidatePromotionIds(ctx context.Context, formats strfmt.Registry) error {

	if err := m.PromotionIds.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("PromotionIds")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("PromotionIds")
		}
		return err
	}

	return nil
}

func (m *OrderItem) contextValidateShippingConstraints(ctx context.Context, formats strfmt.Registry) error {

	if m.ShippingConstraints != nil {
		if err := m.ShippingConstraints.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ShippingConstraints")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ShippingConstraints")
			}
			return err
		}
	}

	return nil
}

func (m *OrderItem) contextValidateShippingDiscount(ctx context.Context, formats strfmt.Registry) error {

	if m.ShippingDiscount != nil {
		if err := m.ShippingDiscount.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ShippingDiscount")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ShippingDiscount")
			}
			return err
		}
	}

	return nil
}

func (m *OrderItem) contextValidateShippingDiscountTax(ctx context.Context, formats strfmt.Registry) error {

	if m.ShippingDiscountTax != nil {
		if err := m.ShippingDiscountTax.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ShippingDiscountTax")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ShippingDiscountTax")
			}
			return err
		}
	}

	return nil
}

func (m *OrderItem) contextValidateShippingPrice(ctx context.Context, formats strfmt.Registry) error {

	if m.ShippingPrice != nil {
		if err := m.ShippingPrice.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ShippingPrice")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ShippingPrice")
			}
			return err
		}
	}

	return nil
}

func (m *OrderItem) contextValidateShippingTax(ctx context.Context, formats strfmt.Registry) error {

	if m.ShippingTax != nil {
		if err := m.ShippingTax.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ShippingTax")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ShippingTax")
			}
			return err
		}
	}

	return nil
}

func (m *OrderItem) contextValidateSubstitutionPreferences(ctx context.Context, formats strfmt.Registry) error {

	if m.SubstitutionPreferences != nil {
		if err := m.SubstitutionPreferences.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("SubstitutionPreferences")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("SubstitutionPreferences")
			}
			return err
		}
	}

	return nil
}

func (m *OrderItem) contextValidateTaxCollection(ctx context.Context, formats strfmt.Registry) error {

	if m.TaxCollection != nil {
		if err := m.TaxCollection.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("TaxCollection")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("TaxCollection")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OrderItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OrderItem) UnmarshalBinary(b []byte) error {
	var res OrderItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
