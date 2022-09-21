// Code generated by go-swagger; DO NOT EDIT.

package vendor_direct_fulfillment_shipping_2021_12_28_models

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

// Container container
//
// swagger:model Container
type Container struct {

	// Carrier required for EU VOC vendors only.
	Carrier string `json:"carrier,omitempty"`

	// The container identifier.
	// Required: true
	ContainerIdentifier *string `json:"containerIdentifier"`

	// An integer that must be submitted for multi-box shipments only, where one item may come in separate packages.
	ContainerSequenceNumber int64 `json:"containerSequenceNumber,omitempty"`

	// The type of container.
	// Required: true
	// Enum: [Carton Pallet]
	ContainerType *string `json:"containerType"`

	// dimensions
	Dimensions *Dimensions `json:"dimensions,omitempty"`

	// The date of the manifest.
	ManifestDate string `json:"manifestDate,omitempty"`

	// The manifest identifier.
	ManifestID string `json:"manifestId,omitempty"`

	// A list of packed items.
	// Required: true
	PackedItems []*PackedItem `json:"packedItems"`

	// SCAC code required for NA VOC vendors only.
	ScacCode string `json:"scacCode,omitempty"`

	// The shipment method. This property is required when calling the submitShipmentConfirmations operation, and optional otherwise.
	ShipMethod string `json:"shipMethod,omitempty"`

	// The tracking number.
	TrackingNumber string `json:"trackingNumber,omitempty"`

	// weight
	Weight *Weight `json:"weight,omitempty"`
}

// Validate validates this container
func (m *Container) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContainerIdentifier(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContainerType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDimensions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePackedItems(formats); err != nil {
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

func (m *Container) validateContainerIdentifier(formats strfmt.Registry) error {

	if err := validate.Required("containerIdentifier", "body", m.ContainerIdentifier); err != nil {
		return err
	}

	return nil
}

var containerTypeContainerTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Carton","Pallet"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		containerTypeContainerTypePropEnum = append(containerTypeContainerTypePropEnum, v)
	}
}

const (

	// ContainerContainerTypeCarton captures enum value "Carton"
	ContainerContainerTypeCarton string = "Carton"

	// ContainerContainerTypePallet captures enum value "Pallet"
	ContainerContainerTypePallet string = "Pallet"
)

// prop value enum
func (m *Container) validateContainerTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, containerTypeContainerTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Container) validateContainerType(formats strfmt.Registry) error {

	if err := validate.Required("containerType", "body", m.ContainerType); err != nil {
		return err
	}

	// value enum
	if err := m.validateContainerTypeEnum("containerType", "body", *m.ContainerType); err != nil {
		return err
	}

	return nil
}

func (m *Container) validateDimensions(formats strfmt.Registry) error {
	if swag.IsZero(m.Dimensions) { // not required
		return nil
	}

	if m.Dimensions != nil {
		if err := m.Dimensions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dimensions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dimensions")
			}
			return err
		}
	}

	return nil
}

func (m *Container) validatePackedItems(formats strfmt.Registry) error {

	if err := validate.Required("packedItems", "body", m.PackedItems); err != nil {
		return err
	}

	for i := 0; i < len(m.PackedItems); i++ {
		if swag.IsZero(m.PackedItems[i]) { // not required
			continue
		}

		if m.PackedItems[i] != nil {
			if err := m.PackedItems[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("packedItems" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("packedItems" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Container) validateWeight(formats strfmt.Registry) error {
	if swag.IsZero(m.Weight) { // not required
		return nil
	}

	if m.Weight != nil {
		if err := m.Weight.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("weight")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("weight")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this container based on the context it is used
func (m *Container) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDimensions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePackedItems(ctx, formats); err != nil {
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

func (m *Container) contextValidateDimensions(ctx context.Context, formats strfmt.Registry) error {

	if m.Dimensions != nil {
		if err := m.Dimensions.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dimensions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dimensions")
			}
			return err
		}
	}

	return nil
}

func (m *Container) contextValidatePackedItems(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.PackedItems); i++ {

		if m.PackedItems[i] != nil {
			if err := m.PackedItems[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("packedItems" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("packedItems" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Container) contextValidateWeight(ctx context.Context, formats strfmt.Registry) error {

	if m.Weight != nil {
		if err := m.Weight.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("weight")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("weight")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Container) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Container) UnmarshalBinary(b []byte) error {
	var res Container
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
