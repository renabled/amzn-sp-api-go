// Code generated by go-swagger; DO NOT EDIT.

package app_integrations_2024_04_01_models

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

// DeleteNotificationsRequest The request for the `deleteNotifications` operation.
// Example: {"deletionReason":"INCORRECT_CONTENT","templateId":"PRICE_CHANGE"}
//
// swagger:model DeleteNotificationsRequest
type DeleteNotificationsRequest struct {

	// The unique identifier that maps each notification status to a reason code.
	// Required: true
	// Enum: [INCORRECT_CONTENT INCORRECT_RECIPIENT]
	DeletionReason *string `json:"deletionReason"`

	// The unique identifier of the notification template you used to onboard your application.
	// Required: true
	TemplateID *string `json:"templateId"`
}

// Validate validates this delete notifications request
func (m *DeleteNotificationsRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDeletionReason(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTemplateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var deleteNotificationsRequestTypeDeletionReasonPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["INCORRECT_CONTENT","INCORRECT_RECIPIENT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		deleteNotificationsRequestTypeDeletionReasonPropEnum = append(deleteNotificationsRequestTypeDeletionReasonPropEnum, v)
	}
}

const (

	// DeleteNotificationsRequestDeletionReasonINCORRECTCONTENT captures enum value "INCORRECT_CONTENT"
	DeleteNotificationsRequestDeletionReasonINCORRECTCONTENT string = "INCORRECT_CONTENT"

	// DeleteNotificationsRequestDeletionReasonINCORRECTRECIPIENT captures enum value "INCORRECT_RECIPIENT"
	DeleteNotificationsRequestDeletionReasonINCORRECTRECIPIENT string = "INCORRECT_RECIPIENT"
)

// prop value enum
func (m *DeleteNotificationsRequest) validateDeletionReasonEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, deleteNotificationsRequestTypeDeletionReasonPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *DeleteNotificationsRequest) validateDeletionReason(formats strfmt.Registry) error {

	if err := validate.Required("deletionReason", "body", m.DeletionReason); err != nil {
		return err
	}

	// value enum
	if err := m.validateDeletionReasonEnum("deletionReason", "body", *m.DeletionReason); err != nil {
		return err
	}

	return nil
}

func (m *DeleteNotificationsRequest) validateTemplateID(formats strfmt.Registry) error {

	if err := validate.Required("templateId", "body", m.TemplateID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this delete notifications request based on context it is used
func (m *DeleteNotificationsRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DeleteNotificationsRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeleteNotificationsRequest) UnmarshalBinary(b []byte) error {
	var res DeleteNotificationsRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
