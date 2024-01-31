// Code generated by go-swagger; DO NOT EDIT.

package notifications_models

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

// EventFilter A `notificationType` specific filter. This object contains all of the currently available filters and properties that you can use to define a `notificationType` specific filter.
//
// swagger:model EventFilter
type EventFilter struct {
	AggregationFilter

	MarketplaceFilter

	OrderChangeTypeFilter

	// An `eventFilterType` value that is supported by the specific `notificationType`. This is used by the subscription service to determine the type of event filter. Refer to [Notification Type Values](https://developer-docs.amazon.com/sp-api/docs/notification-type-values) to determine if an `eventFilterType` is supported.
	// Required: true
	// Enum: [ANY_OFFER_CHANGED ORDER_CHANGE]
	EventFilterType *string `json:"eventFilterType"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *EventFilter) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 AggregationFilter
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.AggregationFilter = aO0

	// AO1
	var aO1 MarketplaceFilter
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.MarketplaceFilter = aO1

	// AO2
	var aO2 OrderChangeTypeFilter
	if err := swag.ReadJSON(raw, &aO2); err != nil {
		return err
	}
	m.OrderChangeTypeFilter = aO2

	// AO3
	var dataAO3 struct {
		EventFilterType *string `json:"eventFilterType"`
	}
	if err := swag.ReadJSON(raw, &dataAO3); err != nil {
		return err
	}

	m.EventFilterType = dataAO3.EventFilterType

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m EventFilter) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 4)

	aO0, err := swag.WriteJSON(m.AggregationFilter)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.MarketplaceFilter)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	aO2, err := swag.WriteJSON(m.OrderChangeTypeFilter)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO2)
	var dataAO3 struct {
		EventFilterType *string `json:"eventFilterType"`
	}

	dataAO3.EventFilterType = m.EventFilterType

	jsonDataAO3, errAO3 := swag.WriteJSON(dataAO3)
	if errAO3 != nil {
		return nil, errAO3
	}
	_parts = append(_parts, jsonDataAO3)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this event filter
func (m *EventFilter) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with AggregationFilter
	if err := m.AggregationFilter.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with MarketplaceFilter
	if err := m.MarketplaceFilter.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with OrderChangeTypeFilter
	if err := m.OrderChangeTypeFilter.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEventFilterType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var eventFilterTypeEventFilterTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ANY_OFFER_CHANGED","ORDER_CHANGE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		eventFilterTypeEventFilterTypePropEnum = append(eventFilterTypeEventFilterTypePropEnum, v)
	}
}

// property enum
func (m *EventFilter) validateEventFilterTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, eventFilterTypeEventFilterTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *EventFilter) validateEventFilterType(formats strfmt.Registry) error {

	if err := validate.Required("eventFilterType", "body", m.EventFilterType); err != nil {
		return err
	}

	// value enum
	if err := m.validateEventFilterTypeEnum("eventFilterType", "body", *m.EventFilterType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this event filter based on the context it is used
func (m *EventFilter) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with AggregationFilter
	if err := m.AggregationFilter.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with MarketplaceFilter
	if err := m.MarketplaceFilter.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with OrderChangeTypeFilter
	if err := m.OrderChangeTypeFilter.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *EventFilter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EventFilter) UnmarshalBinary(b []byte) error {
	var res EventFilter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
