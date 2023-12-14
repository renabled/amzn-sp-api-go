// Code generated by go-swagger; DO NOT EDIT.

package listings_items_2021_08_01_models

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

// PatchOperation Individual JSON Patch operation for an HTTP PATCH request.
//
// swagger:model PatchOperation
type PatchOperation struct {

	// Type of JSON Patch operation. Supported JSON Patch operations include add, replace, and delete. Refer to [JavaScript Object Notation (JSON) Patch](https://tools.ietf.org/html/rfc6902) for more information.
	// Required: true
	// Enum: [add replace delete]
	Op *string `json:"op"`

	// JSON Pointer path of the element to patch. Refer to [JavaScript Object Notation (JSON) Patch](https://tools.ietf.org/html/rfc6902) for more information.
	// Required: true
	Path *string `json:"path"`

	// JSON value to add, replace, or delete.
	Value []interface{} `json:"value"`
}

// Validate validates this patch operation
func (m *PatchOperation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePath(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var patchOperationTypeOpPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["add","replace","delete"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		patchOperationTypeOpPropEnum = append(patchOperationTypeOpPropEnum, v)
	}
}

const (

	// PatchOperationOpAdd captures enum value "add"
	PatchOperationOpAdd string = "add"

	// PatchOperationOpReplace captures enum value "replace"
	PatchOperationOpReplace string = "replace"

	// PatchOperationOpDelete captures enum value "delete"
	PatchOperationOpDelete string = "delete"
)

// prop value enum
func (m *PatchOperation) validateOpEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, patchOperationTypeOpPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PatchOperation) validateOp(formats strfmt.Registry) error {

	if err := validate.Required("op", "body", m.Op); err != nil {
		return err
	}

	// value enum
	if err := m.validateOpEnum("op", "body", *m.Op); err != nil {
		return err
	}

	return nil
}

func (m *PatchOperation) validatePath(formats strfmt.Registry) error {

	if err := validate.Required("path", "body", m.Path); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this patch operation based on context it is used
func (m *PatchOperation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PatchOperation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PatchOperation) UnmarshalBinary(b []byte) error {
	var res PatchOperation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
