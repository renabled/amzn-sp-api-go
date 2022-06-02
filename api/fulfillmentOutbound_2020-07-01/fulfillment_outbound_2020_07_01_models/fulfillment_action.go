// Code generated by go-swagger; DO NOT EDIT.

package fulfillment_outbound_2020_07_01_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// FulfillmentAction Specifies whether the fulfillment order should ship now or have an order hold put on it.
//
// swagger:model FulfillmentAction
type FulfillmentAction string

func NewFulfillmentAction(value FulfillmentAction) *FulfillmentAction {
	return &value
}

// Pointer returns a pointer to a freshly-allocated FulfillmentAction.
func (m FulfillmentAction) Pointer() *FulfillmentAction {
	return &m
}

const (

	// FulfillmentActionShip captures enum value "Ship"
	FulfillmentActionShip FulfillmentAction = "Ship"

	// FulfillmentActionHold captures enum value "Hold"
	FulfillmentActionHold FulfillmentAction = "Hold"
)

// for schema
var fulfillmentActionEnum []interface{}

func init() {
	var res []FulfillmentAction
	if err := json.Unmarshal([]byte(`["Ship","Hold"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		fulfillmentActionEnum = append(fulfillmentActionEnum, v)
	}
}

func (m FulfillmentAction) validateFulfillmentActionEnum(path, location string, value FulfillmentAction) error {
	if err := validate.EnumCase(path, location, value, fulfillmentActionEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this fulfillment action
func (m FulfillmentAction) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateFulfillmentActionEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this fulfillment action based on context it is used
func (m FulfillmentAction) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
