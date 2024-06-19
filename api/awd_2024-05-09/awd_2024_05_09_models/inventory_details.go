// Code generated by go-swagger; DO NOT EDIT.

package awd_2024_05_09_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// InventoryDetails Additional inventory details. This object is only displayed if the details parameter in the request is set to `SHOW`.
//
// swagger:model InventoryDetails
type InventoryDetails struct {

	// Quantity that is available for downstream channel replenishment.
	AvailableDistributableQuantity int64 `json:"availableDistributableQuantity,omitempty"`

	// Quantity that is reserved for a downstream channel replenishment order that is being prepared for shipment.
	ReservedDistributableQuantity int64 `json:"reservedDistributableQuantity,omitempty"`
}

// Validate validates this inventory details
func (m *InventoryDetails) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this inventory details based on context it is used
func (m *InventoryDetails) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *InventoryDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InventoryDetails) UnmarshalBinary(b []byte) error {
	var res InventoryDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
