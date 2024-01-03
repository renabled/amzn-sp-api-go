// Code generated by go-swagger; DO NOT EDIT.

package shipping_v2_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// RateItemID Unique ID for the rateItem.
//
// swagger:model RateItemID
type RateItemID string

func NewRateItemID(value RateItemID) *RateItemID {
	return &value
}

// Pointer returns a pointer to a freshly-allocated RateItemID.
func (m RateItemID) Pointer() *RateItemID {
	return &m
}

const (

	// RateItemIDBASERATE captures enum value "BASE_RATE"
	RateItemIDBASERATE RateItemID = "BASE_RATE"

	// RateItemIDTRANSACTIONFEE captures enum value "TRANSACTION_FEE"
	RateItemIDTRANSACTIONFEE RateItemID = "TRANSACTION_FEE"

	// RateItemIDADULTSIGNATURECONFIRMATION captures enum value "ADULT_SIGNATURE_CONFIRMATION"
	RateItemIDADULTSIGNATURECONFIRMATION RateItemID = "ADULT_SIGNATURE_CONFIRMATION"

	// RateItemIDSIGNATURECONFIRMATION captures enum value "SIGNATURE_CONFIRMATION"
	RateItemIDSIGNATURECONFIRMATION RateItemID = "SIGNATURE_CONFIRMATION"

	// RateItemIDNOCONFIRMATION captures enum value "NO_CONFIRMATION"
	RateItemIDNOCONFIRMATION RateItemID = "NO_CONFIRMATION"

	// RateItemIDWAIVESIGNATURE captures enum value "WAIVE_SIGNATURE"
	RateItemIDWAIVESIGNATURE RateItemID = "WAIVE_SIGNATURE"

	// RateItemIDIMPLIEDLIABILITY captures enum value "IMPLIED_LIABILITY"
	RateItemIDIMPLIEDLIABILITY RateItemID = "IMPLIED_LIABILITY"

	// RateItemIDHIDDENPOSTAGE captures enum value "HIDDEN_POSTAGE"
	RateItemIDHIDDENPOSTAGE RateItemID = "HIDDEN_POSTAGE"

	// RateItemIDDECLAREDVALUE captures enum value "DECLARED_VALUE"
	RateItemIDDECLAREDVALUE RateItemID = "DECLARED_VALUE"

	// RateItemIDSUNDAYHOLIDAYDELIVERY captures enum value "SUNDAY_HOLIDAY_DELIVERY"
	RateItemIDSUNDAYHOLIDAYDELIVERY RateItemID = "SUNDAY_HOLIDAY_DELIVERY"

	// RateItemIDDELIVERYCONFIRMATION captures enum value "DELIVERY_CONFIRMATION"
	RateItemIDDELIVERYCONFIRMATION RateItemID = "DELIVERY_CONFIRMATION"

	// RateItemIDIMPORTDUTYCHARGE captures enum value "IMPORT_DUTY_CHARGE"
	RateItemIDIMPORTDUTYCHARGE RateItemID = "IMPORT_DUTY_CHARGE"

	// RateItemIDVAT captures enum value "VAT"
	RateItemIDVAT RateItemID = "VAT"

	// RateItemIDNOSATURDAYDELIVERY captures enum value "NO_SATURDAY_DELIVERY"
	RateItemIDNOSATURDAYDELIVERY RateItemID = "NO_SATURDAY_DELIVERY"

	// RateItemIDINSURANCE captures enum value "INSURANCE"
	RateItemIDINSURANCE RateItemID = "INSURANCE"

	// RateItemIDCOD captures enum value "COD"
	RateItemIDCOD RateItemID = "COD"

	// RateItemIDFUELSURCHARGE captures enum value "FUEL_SURCHARGE"
	RateItemIDFUELSURCHARGE RateItemID = "FUEL_SURCHARGE"

	// RateItemIDINSPECTIONCHARGE captures enum value "INSPECTION_CHARGE"
	RateItemIDINSPECTIONCHARGE RateItemID = "INSPECTION_CHARGE"

	// RateItemIDDELIVERYAREASURCHARGE captures enum value "DELIVERY_AREA_SURCHARGE"
	RateItemIDDELIVERYAREASURCHARGE RateItemID = "DELIVERY_AREA_SURCHARGE"

	// RateItemIDWAYBILLCHARGE captures enum value "WAYBILL_CHARGE"
	RateItemIDWAYBILLCHARGE RateItemID = "WAYBILL_CHARGE"

	// RateItemIDAMAZONSPONSOREDDISCOUNT captures enum value "AMAZON_SPONSORED_DISCOUNT"
	RateItemIDAMAZONSPONSOREDDISCOUNT RateItemID = "AMAZON_SPONSORED_DISCOUNT"

	// RateItemIDINTEGRATORSPONSOREDDISCOUNT captures enum value "INTEGRATOR_SPONSORED_DISCOUNT"
	RateItemIDINTEGRATORSPONSOREDDISCOUNT RateItemID = "INTEGRATOR_SPONSORED_DISCOUNT"

	// RateItemIDOVERSIZESURCHARGE captures enum value "OVERSIZE_SURCHARGE"
	RateItemIDOVERSIZESURCHARGE RateItemID = "OVERSIZE_SURCHARGE"

	// RateItemIDCONGESTIONCHARGE captures enum value "CONGESTION_CHARGE"
	RateItemIDCONGESTIONCHARGE RateItemID = "CONGESTION_CHARGE"

	// RateItemIDRESIDENTIALSURCHARGE captures enum value "RESIDENTIAL_SURCHARGE"
	RateItemIDRESIDENTIALSURCHARGE RateItemID = "RESIDENTIAL_SURCHARGE"

	// RateItemIDADDITIONALSURCHARGE captures enum value "ADDITIONAL_SURCHARGE"
	RateItemIDADDITIONALSURCHARGE RateItemID = "ADDITIONAL_SURCHARGE"

	// RateItemIDSURCHARGE captures enum value "SURCHARGE"
	RateItemIDSURCHARGE RateItemID = "SURCHARGE"

	// RateItemIDREBATE captures enum value "REBATE"
	RateItemIDREBATE RateItemID = "REBATE"
)

// for schema
var rateItemIdEnum []interface{}

func init() {
	var res []RateItemID
	if err := json.Unmarshal([]byte(`["BASE_RATE","TRANSACTION_FEE","ADULT_SIGNATURE_CONFIRMATION","SIGNATURE_CONFIRMATION","NO_CONFIRMATION","WAIVE_SIGNATURE","IMPLIED_LIABILITY","HIDDEN_POSTAGE","DECLARED_VALUE","SUNDAY_HOLIDAY_DELIVERY","DELIVERY_CONFIRMATION","IMPORT_DUTY_CHARGE","VAT","NO_SATURDAY_DELIVERY","INSURANCE","COD","FUEL_SURCHARGE","INSPECTION_CHARGE","DELIVERY_AREA_SURCHARGE","WAYBILL_CHARGE","AMAZON_SPONSORED_DISCOUNT","INTEGRATOR_SPONSORED_DISCOUNT","OVERSIZE_SURCHARGE","CONGESTION_CHARGE","RESIDENTIAL_SURCHARGE","ADDITIONAL_SURCHARGE","SURCHARGE","REBATE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		rateItemIdEnum = append(rateItemIdEnum, v)
	}
}

func (m RateItemID) validateRateItemIDEnum(path, location string, value RateItemID) error {
	if err := validate.EnumCase(path, location, value, rateItemIdEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this rate item ID
func (m RateItemID) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateRateItemIDEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this rate item ID based on context it is used
func (m RateItemID) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}