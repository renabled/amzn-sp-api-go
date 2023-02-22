// Code generated by go-swagger; DO NOT EDIT.

package vendor_shipments_models

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

// Containers containers
//
// swagger:model containers
type Containers struct {

	// Number of cartons per layer on the pallet.
	Block int64 `json:"block,omitempty"`

	// A list of carton identifiers.
	// Required: true
	ContainerIdentifiers []*ContainerIdentification `json:"containerIdentifiers"`

	// An integer that must be submitted for multi-box shipments only, where one item may come in separate packages.
	ContainerSequenceNumber string `json:"containerSequenceNumber,omitempty"`

	// The type of container.
	// Required: true
	// Enum: [carton pallet]
	ContainerType *string `json:"containerType"`

	// dimensions
	Dimensions *Dimensions `json:"dimensions,omitempty"`

	// inner containers details
	InnerContainersDetails *InnerContainersDetails `json:"innerContainersDetails,omitempty"`

	// A list of packed items.
	PackedItems []*PackedItems `json:"packedItems"`

	// Number of layers per pallet.
	Tier int64 `json:"tier,omitempty"`

	// The tracking number used for identifying the shipment.
	TrackingNumber string `json:"trackingNumber,omitempty"`

	// weight
	Weight *Weight `json:"weight,omitempty"`
}

// Validate validates this containers
func (m *Containers) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContainerIdentifiers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContainerType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDimensions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInnerContainersDetails(formats); err != nil {
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

func (m *Containers) validateContainerIdentifiers(formats strfmt.Registry) error {

	if err := validate.Required("containerIdentifiers", "body", m.ContainerIdentifiers); err != nil {
		return err
	}

	for i := 0; i < len(m.ContainerIdentifiers); i++ {
		if swag.IsZero(m.ContainerIdentifiers[i]) { // not required
			continue
		}

		if m.ContainerIdentifiers[i] != nil {
			if err := m.ContainerIdentifiers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("containerIdentifiers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("containerIdentifiers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var containersTypeContainerTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["carton","pallet"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		containersTypeContainerTypePropEnum = append(containersTypeContainerTypePropEnum, v)
	}
}

const (

	// ContainersContainerTypeCarton captures enum value "carton"
	ContainersContainerTypeCarton string = "carton"

	// ContainersContainerTypePallet captures enum value "pallet"
	ContainersContainerTypePallet string = "pallet"
)

// prop value enum
func (m *Containers) validateContainerTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, containersTypeContainerTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Containers) validateContainerType(formats strfmt.Registry) error {

	if err := validate.Required("containerType", "body", m.ContainerType); err != nil {
		return err
	}

	// value enum
	if err := m.validateContainerTypeEnum("containerType", "body", *m.ContainerType); err != nil {
		return err
	}

	return nil
}

func (m *Containers) validateDimensions(formats strfmt.Registry) error {
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

func (m *Containers) validateInnerContainersDetails(formats strfmt.Registry) error {
	if swag.IsZero(m.InnerContainersDetails) { // not required
		return nil
	}

	if m.InnerContainersDetails != nil {
		if err := m.InnerContainersDetails.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("innerContainersDetails")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("innerContainersDetails")
			}
			return err
		}
	}

	return nil
}

func (m *Containers) validatePackedItems(formats strfmt.Registry) error {
	if swag.IsZero(m.PackedItems) { // not required
		return nil
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

func (m *Containers) validateWeight(formats strfmt.Registry) error {
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

// ContextValidate validate this containers based on the context it is used
func (m *Containers) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateContainerIdentifiers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDimensions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInnerContainersDetails(ctx, formats); err != nil {
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

func (m *Containers) contextValidateContainerIdentifiers(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ContainerIdentifiers); i++ {

		if m.ContainerIdentifiers[i] != nil {
			if err := m.ContainerIdentifiers[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("containerIdentifiers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("containerIdentifiers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Containers) contextValidateDimensions(ctx context.Context, formats strfmt.Registry) error {

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

func (m *Containers) contextValidateInnerContainersDetails(ctx context.Context, formats strfmt.Registry) error {

	if m.InnerContainersDetails != nil {
		if err := m.InnerContainersDetails.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("innerContainersDetails")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("innerContainersDetails")
			}
			return err
		}
	}

	return nil
}

func (m *Containers) contextValidatePackedItems(ctx context.Context, formats strfmt.Registry) error {

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

func (m *Containers) contextValidateWeight(ctx context.Context, formats strfmt.Registry) error {

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
func (m *Containers) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Containers) UnmarshalBinary(b []byte) error {
	var res Containers
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}