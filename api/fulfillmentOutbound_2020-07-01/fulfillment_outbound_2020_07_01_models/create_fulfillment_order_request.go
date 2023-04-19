// Code generated by go-swagger; DO NOT EDIT.

package fulfillment_outbound_2020_07_01_models

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

// CreateFulfillmentOrderRequest The request body schema for the createFulfillmentOrder operation.
//
// swagger:model CreateFulfillmentOrderRequest
type CreateFulfillmentOrderRequest struct {

	// cod settings
	CodSettings *CODSettings `json:"codSettings,omitempty"`

	// delivery window
	DeliveryWindow *DeliveryWindow `json:"deliveryWindow,omitempty"`

	// The destination address for the fulfillment order.
	// Required: true
	DestinationAddress *Address `json:"destinationAddress"`

	// Order-specific text that appears in recipient-facing materials such as the outbound shipment packing slip.
	// Required: true
	// Max Length: 1000
	DisplayableOrderComment *string `json:"displayableOrderComment"`

	// The date and time of the fulfillment order. Displays as the order date in recipient-facing materials such as the outbound shipment packing slip.
	// Required: true
	// Format: date-time
	DisplayableOrderDate *Timestamp `json:"displayableOrderDate"`

	// A fulfillment order identifier that the seller creates. This value displays as the order identifier in recipient-facing materials such as the outbound shipment packing slip. The value of DisplayableOrderId should match the order identifier that the seller provides to the recipient. The seller can use the SellerFulfillmentOrderId for this value or they can specify an alternate value if they want the recipient to reference an alternate order identifier.
	//
	// The value must be an alpha-numeric or ISO 8859-1 compliant string from one to 40 characters in length. Cannot contain two spaces in a row. Leading and trailing white space is removed.
	// Required: true
	// Max Length: 40
	DisplayableOrderID *string `json:"displayableOrderId"`

	// A list of features and their fulfillment policies to apply to the order.
	FeatureConstraints []*FeatureSettings `json:"featureConstraints"`

	// fulfillment action
	FulfillmentAction FulfillmentAction `json:"fulfillmentAction,omitempty"`

	// fulfillment policy
	FulfillmentPolicy FulfillmentPolicy `json:"fulfillmentPolicy,omitempty"`

	// A list of items to include in the fulfillment order preview, including quantity.
	// Required: true
	Items CreateFulfillmentOrderItemList `json:"items"`

	// The marketplace the fulfillment order is placed against.
	MarketplaceID string `json:"marketplaceId,omitempty"`

	// notification emails
	NotificationEmails NotificationEmailList `json:"notificationEmails,omitempty"`

	// A fulfillment order identifier that the seller creates to track their fulfillment order. The SellerFulfillmentOrderId must be unique for each fulfillment order that a seller creates. If the seller's system already creates unique order identifiers, then these might be good values for them to use.
	// Required: true
	// Max Length: 40
	SellerFulfillmentOrderID *string `json:"sellerFulfillmentOrderId"`

	// The two-character country code for the country from which the fulfillment order ships. Must be in ISO 3166-1 alpha-2 format.
	ShipFromCountryCode string `json:"shipFromCountryCode,omitempty"`

	// The shipping method for the fulfillment order. When this value is ScheduledDelivery, choose Ship for the fulfillmentAction. Hold is not a valid fulfillmentAction value when the shippingSpeedCategory value is ScheduledDelivery.
	// Required: true
	ShippingSpeedCategory *ShippingSpeedCategory `json:"shippingSpeedCategory"`
}

// Validate validates this create fulfillment order request
func (m *CreateFulfillmentOrderRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCodSettings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeliveryWindow(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDestinationAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDisplayableOrderComment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDisplayableOrderDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDisplayableOrderID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFeatureConstraints(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFulfillmentAction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFulfillmentPolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateItems(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNotificationEmails(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSellerFulfillmentOrderID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShippingSpeedCategory(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateFulfillmentOrderRequest) validateCodSettings(formats strfmt.Registry) error {
	if swag.IsZero(m.CodSettings) { // not required
		return nil
	}

	if m.CodSettings != nil {
		if err := m.CodSettings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("codSettings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("codSettings")
			}
			return err
		}
	}

	return nil
}

func (m *CreateFulfillmentOrderRequest) validateDeliveryWindow(formats strfmt.Registry) error {
	if swag.IsZero(m.DeliveryWindow) { // not required
		return nil
	}

	if m.DeliveryWindow != nil {
		if err := m.DeliveryWindow.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deliveryWindow")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("deliveryWindow")
			}
			return err
		}
	}

	return nil
}

func (m *CreateFulfillmentOrderRequest) validateDestinationAddress(formats strfmt.Registry) error {

	if err := validate.Required("destinationAddress", "body", m.DestinationAddress); err != nil {
		return err
	}

	if m.DestinationAddress != nil {
		if err := m.DestinationAddress.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("destinationAddress")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("destinationAddress")
			}
			return err
		}
	}

	return nil
}

func (m *CreateFulfillmentOrderRequest) validateDisplayableOrderComment(formats strfmt.Registry) error {

	if err := validate.Required("displayableOrderComment", "body", m.DisplayableOrderComment); err != nil {
		return err
	}

	if err := validate.MaxLength("displayableOrderComment", "body", *m.DisplayableOrderComment, 1000); err != nil {
		return err
	}

	return nil
}

func (m *CreateFulfillmentOrderRequest) validateDisplayableOrderDate(formats strfmt.Registry) error {

	if err := validate.Required("displayableOrderDate", "body", m.DisplayableOrderDate); err != nil {
		return err
	}

	if err := validate.Required("displayableOrderDate", "body", m.DisplayableOrderDate); err != nil {
		return err
	}

	if m.DisplayableOrderDate != nil {
		if err := m.DisplayableOrderDate.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("displayableOrderDate")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("displayableOrderDate")
			}
			return err
		}
	}

	return nil
}

func (m *CreateFulfillmentOrderRequest) validateDisplayableOrderID(formats strfmt.Registry) error {

	if err := validate.Required("displayableOrderId", "body", m.DisplayableOrderID); err != nil {
		return err
	}

	if err := validate.MaxLength("displayableOrderId", "body", *m.DisplayableOrderID, 40); err != nil {
		return err
	}

	return nil
}

func (m *CreateFulfillmentOrderRequest) validateFeatureConstraints(formats strfmt.Registry) error {
	if swag.IsZero(m.FeatureConstraints) { // not required
		return nil
	}

	for i := 0; i < len(m.FeatureConstraints); i++ {
		if swag.IsZero(m.FeatureConstraints[i]) { // not required
			continue
		}

		if m.FeatureConstraints[i] != nil {
			if err := m.FeatureConstraints[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("featureConstraints" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("featureConstraints" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CreateFulfillmentOrderRequest) validateFulfillmentAction(formats strfmt.Registry) error {
	if swag.IsZero(m.FulfillmentAction) { // not required
		return nil
	}

	if err := m.FulfillmentAction.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("fulfillmentAction")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("fulfillmentAction")
		}
		return err
	}

	return nil
}

func (m *CreateFulfillmentOrderRequest) validateFulfillmentPolicy(formats strfmt.Registry) error {
	if swag.IsZero(m.FulfillmentPolicy) { // not required
		return nil
	}

	if err := m.FulfillmentPolicy.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("fulfillmentPolicy")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("fulfillmentPolicy")
		}
		return err
	}

	return nil
}

func (m *CreateFulfillmentOrderRequest) validateItems(formats strfmt.Registry) error {

	if err := validate.Required("items", "body", m.Items); err != nil {
		return err
	}

	if err := m.Items.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("items")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("items")
		}
		return err
	}

	return nil
}

func (m *CreateFulfillmentOrderRequest) validateNotificationEmails(formats strfmt.Registry) error {
	if swag.IsZero(m.NotificationEmails) { // not required
		return nil
	}

	if err := m.NotificationEmails.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("notificationEmails")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("notificationEmails")
		}
		return err
	}

	return nil
}

func (m *CreateFulfillmentOrderRequest) validateSellerFulfillmentOrderID(formats strfmt.Registry) error {

	if err := validate.Required("sellerFulfillmentOrderId", "body", m.SellerFulfillmentOrderID); err != nil {
		return err
	}

	if err := validate.MaxLength("sellerFulfillmentOrderId", "body", *m.SellerFulfillmentOrderID, 40); err != nil {
		return err
	}

	return nil
}

func (m *CreateFulfillmentOrderRequest) validateShippingSpeedCategory(formats strfmt.Registry) error {

	if err := validate.Required("shippingSpeedCategory", "body", m.ShippingSpeedCategory); err != nil {
		return err
	}

	if err := validate.Required("shippingSpeedCategory", "body", m.ShippingSpeedCategory); err != nil {
		return err
	}

	if m.ShippingSpeedCategory != nil {
		if err := m.ShippingSpeedCategory.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("shippingSpeedCategory")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("shippingSpeedCategory")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this create fulfillment order request based on the context it is used
func (m *CreateFulfillmentOrderRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCodSettings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDeliveryWindow(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDestinationAddress(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDisplayableOrderDate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFeatureConstraints(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFulfillmentAction(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFulfillmentPolicy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateItems(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNotificationEmails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateShippingSpeedCategory(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateFulfillmentOrderRequest) contextValidateCodSettings(ctx context.Context, formats strfmt.Registry) error {

	if m.CodSettings != nil {
		if err := m.CodSettings.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("codSettings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("codSettings")
			}
			return err
		}
	}

	return nil
}

func (m *CreateFulfillmentOrderRequest) contextValidateDeliveryWindow(ctx context.Context, formats strfmt.Registry) error {

	if m.DeliveryWindow != nil {
		if err := m.DeliveryWindow.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deliveryWindow")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("deliveryWindow")
			}
			return err
		}
	}

	return nil
}

func (m *CreateFulfillmentOrderRequest) contextValidateDestinationAddress(ctx context.Context, formats strfmt.Registry) error {

	if m.DestinationAddress != nil {
		if err := m.DestinationAddress.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("destinationAddress")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("destinationAddress")
			}
			return err
		}
	}

	return nil
}

func (m *CreateFulfillmentOrderRequest) contextValidateDisplayableOrderDate(ctx context.Context, formats strfmt.Registry) error {

	if m.DisplayableOrderDate != nil {
		if err := m.DisplayableOrderDate.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("displayableOrderDate")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("displayableOrderDate")
			}
			return err
		}
	}

	return nil
}

func (m *CreateFulfillmentOrderRequest) contextValidateFeatureConstraints(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.FeatureConstraints); i++ {

		if m.FeatureConstraints[i] != nil {
			if err := m.FeatureConstraints[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("featureConstraints" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("featureConstraints" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CreateFulfillmentOrderRequest) contextValidateFulfillmentAction(ctx context.Context, formats strfmt.Registry) error {

	if err := m.FulfillmentAction.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("fulfillmentAction")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("fulfillmentAction")
		}
		return err
	}

	return nil
}

func (m *CreateFulfillmentOrderRequest) contextValidateFulfillmentPolicy(ctx context.Context, formats strfmt.Registry) error {

	if err := m.FulfillmentPolicy.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("fulfillmentPolicy")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("fulfillmentPolicy")
		}
		return err
	}

	return nil
}

func (m *CreateFulfillmentOrderRequest) contextValidateItems(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Items.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("items")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("items")
		}
		return err
	}

	return nil
}

func (m *CreateFulfillmentOrderRequest) contextValidateNotificationEmails(ctx context.Context, formats strfmt.Registry) error {

	if err := m.NotificationEmails.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("notificationEmails")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("notificationEmails")
		}
		return err
	}

	return nil
}

func (m *CreateFulfillmentOrderRequest) contextValidateShippingSpeedCategory(ctx context.Context, formats strfmt.Registry) error {

	if m.ShippingSpeedCategory != nil {
		if err := m.ShippingSpeedCategory.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("shippingSpeedCategory")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("shippingSpeedCategory")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateFulfillmentOrderRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateFulfillmentOrderRequest) UnmarshalBinary(b []byte) error {
	var res CreateFulfillmentOrderRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
