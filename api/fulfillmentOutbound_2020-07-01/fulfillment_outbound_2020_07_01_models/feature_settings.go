// Code generated by go-swagger; DO NOT EDIT.

package fulfillment_outbound_2020_07_01_models

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

// FeatureSettings `FeatureSettings` allows users to apply fulfillment features to an order. To block an order from being shipped using Amazon Logistics (AMZL) and an AMZL tracking number, use `featureName` as `BLOCK_AMZL` and `featureFulfillmentPolicy` as `Required`. Blocking AMZL will incur an additional fee surcharge on your MCF orders and increase the risk of some of your orders being unfulfilled or delivered late if there are no alternative carriers available. Using `BLOCK_AMZL` in an order request will take precedence over your Seller Central account setting. To ship in non-Amazon branded packaging (blank boxes), use featureName `BLANK_BOX`.
//
// swagger:model FeatureSettings
type FeatureSettings struct {

	// Specifies the policy to use when fulfilling an order.
	// Enum: [Required NotRequired]
	FeatureFulfillmentPolicy string `json:"featureFulfillmentPolicy,omitempty"`

	// The name of the feature.
	FeatureName string `json:"featureName,omitempty"`
}

// Validate validates this feature settings
func (m *FeatureSettings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFeatureFulfillmentPolicy(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var featureSettingsTypeFeatureFulfillmentPolicyPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Required","NotRequired"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		featureSettingsTypeFeatureFulfillmentPolicyPropEnum = append(featureSettingsTypeFeatureFulfillmentPolicyPropEnum, v)
	}
}

const (

	// FeatureSettingsFeatureFulfillmentPolicyRequired captures enum value "Required"
	FeatureSettingsFeatureFulfillmentPolicyRequired string = "Required"

	// FeatureSettingsFeatureFulfillmentPolicyNotRequired captures enum value "NotRequired"
	FeatureSettingsFeatureFulfillmentPolicyNotRequired string = "NotRequired"
)

// prop value enum
func (m *FeatureSettings) validateFeatureFulfillmentPolicyEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, featureSettingsTypeFeatureFulfillmentPolicyPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *FeatureSettings) validateFeatureFulfillmentPolicy(formats strfmt.Registry) error {
	if swag.IsZero(m.FeatureFulfillmentPolicy) { // not required
		return nil
	}

	// value enum
	if err := m.validateFeatureFulfillmentPolicyEnum("featureFulfillmentPolicy", "body", m.FeatureFulfillmentPolicy); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this feature settings based on context it is used
func (m *FeatureSettings) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *FeatureSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FeatureSettings) UnmarshalBinary(b []byte) error {
	var res FeatureSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
