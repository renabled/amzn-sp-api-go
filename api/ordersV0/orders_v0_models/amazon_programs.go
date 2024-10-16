// Code generated by go-swagger; DO NOT EDIT.

package orders_v0_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AmazonPrograms Contains the list of programs that are associated with an item.
//
// Possible programs are:
//   - **Subscribe and Save**: Offers recurring, scheduled deliveries to Amazon customers and Amazon Business customers for their frequently ordered products.
//
// swagger:model AmazonPrograms
type AmazonPrograms struct {

	// A list of the programs that are associated with the specified order item.
	//
	// **Possible values**: `SUBSCRIBE_AND_SAVE`
	// Required: true
	Programs []string `json:"Programs"`
}

// Validate validates this amazon programs
func (m *AmazonPrograms) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePrograms(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AmazonPrograms) validatePrograms(formats strfmt.Registry) error {

	if err := validate.Required("Programs", "body", m.Programs); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this amazon programs based on context it is used
func (m *AmazonPrograms) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AmazonPrograms) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AmazonPrograms) UnmarshalBinary(b []byte) error {
	var res AmazonPrograms
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
