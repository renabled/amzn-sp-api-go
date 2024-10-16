// Code generated by go-swagger; DO NOT EDIT.

package sellers_models

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

// Account The response schema for the `getAccount` operation.
//
// swagger:model Account
type Account struct {

	// business
	Business *Business `json:"business,omitempty"`

	// The type of business registered for the seller account.
	// Required: true
	// Enum: [CHARITY CRAFTSMAN NATURAL_PERSON_COMPANY PUBLIC_LISTED PRIVATE_LIMITED SOLE_PROPRIETORSHIP STATE_OWNED INDIVIDUAL]
	BusinessType *string `json:"businessType"`

	// A list of details of the marketplaces where the seller account is active.
	// Required: true
	MarketplaceLevelAttributes []*MarketplaceLevelAttributes `json:"marketplaceLevelAttributes"`

	// primary contact
	PrimaryContact *PrimaryContact `json:"primaryContact,omitempty"`
}

// Validate validates this account
func (m *Account) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBusiness(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBusinessType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMarketplaceLevelAttributes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrimaryContact(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Account) validateBusiness(formats strfmt.Registry) error {
	if swag.IsZero(m.Business) { // not required
		return nil
	}

	if m.Business != nil {
		if err := m.Business.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("business")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("business")
			}
			return err
		}
	}

	return nil
}

var accountTypeBusinessTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CHARITY","CRAFTSMAN","NATURAL_PERSON_COMPANY","PUBLIC_LISTED","PRIVATE_LIMITED","SOLE_PROPRIETORSHIP","STATE_OWNED","INDIVIDUAL"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		accountTypeBusinessTypePropEnum = append(accountTypeBusinessTypePropEnum, v)
	}
}

const (

	// AccountBusinessTypeCHARITY captures enum value "CHARITY"
	AccountBusinessTypeCHARITY string = "CHARITY"

	// AccountBusinessTypeCRAFTSMAN captures enum value "CRAFTSMAN"
	AccountBusinessTypeCRAFTSMAN string = "CRAFTSMAN"

	// AccountBusinessTypeNATURALPERSONCOMPANY captures enum value "NATURAL_PERSON_COMPANY"
	AccountBusinessTypeNATURALPERSONCOMPANY string = "NATURAL_PERSON_COMPANY"

	// AccountBusinessTypePUBLICLISTED captures enum value "PUBLIC_LISTED"
	AccountBusinessTypePUBLICLISTED string = "PUBLIC_LISTED"

	// AccountBusinessTypePRIVATELIMITED captures enum value "PRIVATE_LIMITED"
	AccountBusinessTypePRIVATELIMITED string = "PRIVATE_LIMITED"

	// AccountBusinessTypeSOLEPROPRIETORSHIP captures enum value "SOLE_PROPRIETORSHIP"
	AccountBusinessTypeSOLEPROPRIETORSHIP string = "SOLE_PROPRIETORSHIP"

	// AccountBusinessTypeSTATEOWNED captures enum value "STATE_OWNED"
	AccountBusinessTypeSTATEOWNED string = "STATE_OWNED"

	// AccountBusinessTypeINDIVIDUAL captures enum value "INDIVIDUAL"
	AccountBusinessTypeINDIVIDUAL string = "INDIVIDUAL"
)

// prop value enum
func (m *Account) validateBusinessTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, accountTypeBusinessTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Account) validateBusinessType(formats strfmt.Registry) error {

	if err := validate.Required("businessType", "body", m.BusinessType); err != nil {
		return err
	}

	// value enum
	if err := m.validateBusinessTypeEnum("businessType", "body", *m.BusinessType); err != nil {
		return err
	}

	return nil
}

func (m *Account) validateMarketplaceLevelAttributes(formats strfmt.Registry) error {

	if err := validate.Required("marketplaceLevelAttributes", "body", m.MarketplaceLevelAttributes); err != nil {
		return err
	}

	for i := 0; i < len(m.MarketplaceLevelAttributes); i++ {
		if swag.IsZero(m.MarketplaceLevelAttributes[i]) { // not required
			continue
		}

		if m.MarketplaceLevelAttributes[i] != nil {
			if err := m.MarketplaceLevelAttributes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("marketplaceLevelAttributes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("marketplaceLevelAttributes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Account) validatePrimaryContact(formats strfmt.Registry) error {
	if swag.IsZero(m.PrimaryContact) { // not required
		return nil
	}

	if m.PrimaryContact != nil {
		if err := m.PrimaryContact.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("primaryContact")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("primaryContact")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this account based on the context it is used
func (m *Account) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBusiness(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMarketplaceLevelAttributes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePrimaryContact(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Account) contextValidateBusiness(ctx context.Context, formats strfmt.Registry) error {

	if m.Business != nil {
		if err := m.Business.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("business")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("business")
			}
			return err
		}
	}

	return nil
}

func (m *Account) contextValidateMarketplaceLevelAttributes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.MarketplaceLevelAttributes); i++ {

		if m.MarketplaceLevelAttributes[i] != nil {
			if err := m.MarketplaceLevelAttributes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("marketplaceLevelAttributes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("marketplaceLevelAttributes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Account) contextValidatePrimaryContact(ctx context.Context, formats strfmt.Registry) error {

	if m.PrimaryContact != nil {
		if err := m.PrimaryContact.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("primaryContact")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("primaryContact")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Account) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Account) UnmarshalBinary(b []byte) error {
	var res Account
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
