// Code generated by go-swagger; DO NOT EDIT.

package fulfillment_inbound_2024_03_20_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// PrepType Preparation instructions for shipping an item to Amazon's fulfillment network. For more information about preparing items for shipment to Amazon's fulfillment network, refer to [Seller Central Help for your marketplace](https://developer-docs.amazon.com/sp-api/docs/seller-central-urls).
//
// swagger:model PrepType
type PrepType string

func NewPrepType(value PrepType) *PrepType {
	return &value
}

// Pointer returns a pointer to a freshly-allocated PrepType.
func (m PrepType) Pointer() *PrepType {
	return &m
}

const (

	// PrepTypeITEMBLACKSHRINKWRAP captures enum value "ITEM_BLACK_SHRINKWRAP"
	PrepTypeITEMBLACKSHRINKWRAP PrepType = "ITEM_BLACK_SHRINKWRAP"

	// PrepTypeITEMBLANKSTK captures enum value "ITEM_BLANKSTK"
	PrepTypeITEMBLANKSTK PrepType = "ITEM_BLANKSTK"

	// PrepTypeITEMBOXING captures enum value "ITEM_BOXING"
	PrepTypeITEMBOXING PrepType = "ITEM_BOXING"

	// PrepTypeITEMBUBBLEWRAP captures enum value "ITEM_BUBBLEWRAP"
	PrepTypeITEMBUBBLEWRAP PrepType = "ITEM_BUBBLEWRAP"

	// PrepTypeITEMCAPSEALING captures enum value "ITEM_CAP_SEALING"
	PrepTypeITEMCAPSEALING PrepType = "ITEM_CAP_SEALING"

	// PrepTypeITEMDEBUNDLE captures enum value "ITEM_DEBUNDLE"
	PrepTypeITEMDEBUNDLE PrepType = "ITEM_DEBUNDLE"

	// PrepTypeITEMHANGGARMENT captures enum value "ITEM_HANG_GARMENT"
	PrepTypeITEMHANGGARMENT PrepType = "ITEM_HANG_GARMENT"

	// PrepTypeITEMLABELING captures enum value "ITEM_LABELING"
	PrepTypeITEMLABELING PrepType = "ITEM_LABELING"

	// PrepTypeITEMNOPREP captures enum value "ITEM_NO_PREP"
	PrepTypeITEMNOPREP PrepType = "ITEM_NO_PREP"

	// PrepTypeITEMPOLYBAGGING captures enum value "ITEM_POLYBAGGING"
	PrepTypeITEMPOLYBAGGING PrepType = "ITEM_POLYBAGGING"

	// PrepTypeITEMRMOVHANG captures enum value "ITEM_RMOVHANG"
	PrepTypeITEMRMOVHANG PrepType = "ITEM_RMOVHANG"

	// PrepTypeITEMSETCREAT captures enum value "ITEM_SETCREAT"
	PrepTypeITEMSETCREAT PrepType = "ITEM_SETCREAT"

	// PrepTypeITEMSETSTK captures enum value "ITEM_SETSTK"
	PrepTypeITEMSETSTK PrepType = "ITEM_SETSTK"

	// PrepTypeITEMSIOC captures enum value "ITEM_SIOC"
	PrepTypeITEMSIOC PrepType = "ITEM_SIOC"

	// PrepTypeITEMSUFFOSTK captures enum value "ITEM_SUFFOSTK"
	PrepTypeITEMSUFFOSTK PrepType = "ITEM_SUFFOSTK"

	// PrepTypeITEMTAPING captures enum value "ITEM_TAPING"
	PrepTypeITEMTAPING PrepType = "ITEM_TAPING"
)

// for schema
var prepTypeEnum []interface{}

func init() {
	var res []PrepType
	if err := json.Unmarshal([]byte(`["ITEM_BLACK_SHRINKWRAP","ITEM_BLANKSTK","ITEM_BOXING","ITEM_BUBBLEWRAP","ITEM_CAP_SEALING","ITEM_DEBUNDLE","ITEM_HANG_GARMENT","ITEM_LABELING","ITEM_NO_PREP","ITEM_POLYBAGGING","ITEM_RMOVHANG","ITEM_SETCREAT","ITEM_SETSTK","ITEM_SIOC","ITEM_SUFFOSTK","ITEM_TAPING"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		prepTypeEnum = append(prepTypeEnum, v)
	}
}

func (m PrepType) validatePrepTypeEnum(path, location string, value PrepType) error {
	if err := validate.EnumCase(path, location, value, prepTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this prep type
func (m PrepType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validatePrepTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this prep type based on context it is used
func (m PrepType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
