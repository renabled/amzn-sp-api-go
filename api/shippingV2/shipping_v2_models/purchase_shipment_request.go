// Code generated by go-swagger; DO NOT EDIT.

package shipping_v2_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PurchaseShipmentRequest The request schema for the purchaseShipment operation.
//
// swagger:model PurchaseShipmentRequest
type PurchaseShipmentRequest struct {

	// The additional inputs required to purchase a shipping offering, in JSON format. The JSON provided here must adhere to the JSON schema that is returned in the response to the getAdditionalInputs operation.
	//
	// Additional inputs are only required when indicated by the requiresAdditionalInputs property in the response to the getRates operation.
	AdditionalInputs interface{} `json:"additionalInputs,omitempty"`

	// rate Id
	// Required: true
	RateID *RateID `json:"rateId"`

	// request token
	// Required: true
	RequestToken *RequestToken `json:"requestToken"`

	// requested document specification
	// Required: true
	RequestedDocumentSpecification *RequestedDocumentSpecification `json:"requestedDocumentSpecification"`

	// requested value added services
	RequestedValueAddedServices RequestedValueAddedServiceList `json:"requestedValueAddedServices,omitempty"`
}

// Validate validates this purchase shipment request
func (m *PurchaseShipmentRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequestToken(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequestedDocumentSpecification(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequestedValueAddedServices(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PurchaseShipmentRequest) validateRateID(formats strfmt.Registry) error {

	if err := validate.Required("rateId", "body", m.RateID); err != nil {
		return err
	}

	if err := validate.Required("rateId", "body", m.RateID); err != nil {
		return err
	}

	if m.RateID != nil {
		if err := m.RateID.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rateId")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("rateId")
			}
			return err
		}
	}

	return nil
}

func (m *PurchaseShipmentRequest) validateRequestToken(formats strfmt.Registry) error {

	if err := validate.Required("requestToken", "body", m.RequestToken); err != nil {
		return err
	}

	if err := validate.Required("requestToken", "body", m.RequestToken); err != nil {
		return err
	}

	if m.RequestToken != nil {
		if err := m.RequestToken.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("requestToken")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("requestToken")
			}
			return err
		}
	}

	return nil
}

func (m *PurchaseShipmentRequest) validateRequestedDocumentSpecification(formats strfmt.Registry) error {

	if err := validate.Required("requestedDocumentSpecification", "body", m.RequestedDocumentSpecification); err != nil {
		return err
	}

	if m.RequestedDocumentSpecification != nil {
		if err := m.RequestedDocumentSpecification.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("requestedDocumentSpecification")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("requestedDocumentSpecification")
			}
			return err
		}
	}

	return nil
}

func (m *PurchaseShipmentRequest) validateRequestedValueAddedServices(formats strfmt.Registry) error {
	if swag.IsZero(m.RequestedValueAddedServices) { // not required
		return nil
	}

	if err := m.RequestedValueAddedServices.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("requestedValueAddedServices")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("requestedValueAddedServices")
		}
		return err
	}

	return nil
}

// ContextValidate validate this purchase shipment request based on the context it is used
func (m *PurchaseShipmentRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRequestToken(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRequestedDocumentSpecification(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRequestedValueAddedServices(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PurchaseShipmentRequest) contextValidateRateID(ctx context.Context, formats strfmt.Registry) error {

	if m.RateID != nil {
		if err := m.RateID.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rateId")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("rateId")
			}
			return err
		}
	}

	return nil
}

func (m *PurchaseShipmentRequest) contextValidateRequestToken(ctx context.Context, formats strfmt.Registry) error {

	if m.RequestToken != nil {
		if err := m.RequestToken.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("requestToken")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("requestToken")
			}
			return err
		}
	}

	return nil
}

func (m *PurchaseShipmentRequest) contextValidateRequestedDocumentSpecification(ctx context.Context, formats strfmt.Registry) error {

	if m.RequestedDocumentSpecification != nil {
		if err := m.RequestedDocumentSpecification.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("requestedDocumentSpecification")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("requestedDocumentSpecification")
			}
			return err
		}
	}

	return nil
}

func (m *PurchaseShipmentRequest) contextValidateRequestedValueAddedServices(ctx context.Context, formats strfmt.Registry) error {

	if err := m.RequestedValueAddedServices.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("requestedValueAddedServices")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("requestedValueAddedServices")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PurchaseShipmentRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PurchaseShipmentRequest) UnmarshalBinary(b []byte) error {
	var res PurchaseShipmentRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}