// Code generated by go-swagger; DO NOT EDIT.

package shipping_v2_models

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

// ClientReferenceDetail Client Reference Details
//
// swagger:model ClientReferenceDetail
type ClientReferenceDetail struct {

	// The Client Reference Id.
	// Required: true
	ClientReferenceID *string `json:"clientReferenceId"`

	// Client Reference type.
	// Required: true
	// Enum: [IntegratorShipperId IntegratorMerchantId]
	ClientReferenceType *string `json:"clientReferenceType"`
}

// Validate validates this client reference detail
func (m *ClientReferenceDetail) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClientReferenceID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClientReferenceType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClientReferenceDetail) validateClientReferenceID(formats strfmt.Registry) error {

	if err := validate.Required("clientReferenceId", "body", m.ClientReferenceID); err != nil {
		return err
	}

	return nil
}

var clientReferenceDetailTypeClientReferenceTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["IntegratorShipperId","IntegratorMerchantId"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		clientReferenceDetailTypeClientReferenceTypePropEnum = append(clientReferenceDetailTypeClientReferenceTypePropEnum, v)
	}
}

const (

	// ClientReferenceDetailClientReferenceTypeIntegratorShipperID captures enum value "IntegratorShipperId"
	ClientReferenceDetailClientReferenceTypeIntegratorShipperID string = "IntegratorShipperId"

	// ClientReferenceDetailClientReferenceTypeIntegratorMerchantID captures enum value "IntegratorMerchantId"
	ClientReferenceDetailClientReferenceTypeIntegratorMerchantID string = "IntegratorMerchantId"
)

// prop value enum
func (m *ClientReferenceDetail) validateClientReferenceTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, clientReferenceDetailTypeClientReferenceTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ClientReferenceDetail) validateClientReferenceType(formats strfmt.Registry) error {

	if err := validate.Required("clientReferenceType", "body", m.ClientReferenceType); err != nil {
		return err
	}

	// value enum
	if err := m.validateClientReferenceTypeEnum("clientReferenceType", "body", *m.ClientReferenceType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this client reference detail based on context it is used
func (m *ClientReferenceDetail) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ClientReferenceDetail) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClientReferenceDetail) UnmarshalBinary(b []byte) error {
	var res ClientReferenceDetail
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
