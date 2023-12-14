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

// AccessPointType The type of access point, like counter (HELIX), lockers, etc.
//
// swagger:model AccessPointType
type AccessPointType string

func NewAccessPointType(value AccessPointType) *AccessPointType {
	return &value
}

// Pointer returns a pointer to a freshly-allocated AccessPointType.
func (m AccessPointType) Pointer() *AccessPointType {
	return &m
}

const (

	// AccessPointTypeHELIX captures enum value "HELIX"
	AccessPointTypeHELIX AccessPointType = "HELIX"

	// AccessPointTypeCAMPUSLOCKER captures enum value "CAMPUS_LOCKER"
	AccessPointTypeCAMPUSLOCKER AccessPointType = "CAMPUS_LOCKER"

	// AccessPointTypeOMNILOCKER captures enum value "OMNI_LOCKER"
	AccessPointTypeOMNILOCKER AccessPointType = "OMNI_LOCKER"

	// AccessPointTypeODINLOCKER captures enum value "ODIN_LOCKER"
	AccessPointTypeODINLOCKER AccessPointType = "ODIN_LOCKER"

	// AccessPointTypeDOBBYLOCKER captures enum value "DOBBY_LOCKER"
	AccessPointTypeDOBBYLOCKER AccessPointType = "DOBBY_LOCKER"

	// AccessPointTypeCORELOCKER captures enum value "CORE_LOCKER"
	AccessPointTypeCORELOCKER AccessPointType = "CORE_LOCKER"

	// AccessPointTypeNr3P captures enum value "3P"
	AccessPointTypeNr3P AccessPointType = "3P"

	// AccessPointTypeCAMPUSROOM captures enum value "CAMPUS_ROOM"
	AccessPointTypeCAMPUSROOM AccessPointType = "CAMPUS_ROOM"
)

// for schema
var accessPointTypeEnum []interface{}

func init() {
	var res []AccessPointType
	if err := json.Unmarshal([]byte(`["HELIX","CAMPUS_LOCKER","OMNI_LOCKER","ODIN_LOCKER","DOBBY_LOCKER","CORE_LOCKER","3P","CAMPUS_ROOM"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		accessPointTypeEnum = append(accessPointTypeEnum, v)
	}
}

func (m AccessPointType) validateAccessPointTypeEnum(path, location string, value AccessPointType) error {
	if err := validate.EnumCase(path, location, value, accessPointTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this access point type
func (m AccessPointType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateAccessPointTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this access point type based on context it is used
func (m AccessPointType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
