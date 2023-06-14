// Code generated by go-swagger; DO NOT EDIT.

package vendor_direct_fulfillment_shipping_2021_12_28_client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/renabled/amzn-sp-api-go/api/vendorDirectFulfillmentShipping_2021-12-28/vendor_direct_fulfillment_shipping_2021_12_28_client/customer_invoices"
	"github.com/renabled/amzn-sp-api-go/api/vendorDirectFulfillmentShipping_2021-12-28/vendor_direct_fulfillment_shipping_2021_12_28_client/vendor_shipping"
	"github.com/renabled/amzn-sp-api-go/api/vendorDirectFulfillmentShipping_2021-12-28/vendor_direct_fulfillment_shipping_2021_12_28_client/vendor_shipping_labels"
)

// Default vendor direct fulfillment shipping20211228 HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "sellingpartnerapi-na.amazon.com"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"https"}

// NewHTTPClient creates a new vendor direct fulfillment shipping20211228 HTTP client.
func NewHTTPClient(formats strfmt.Registry) *VendorDirectFulfillmentShipping20211228 {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new vendor direct fulfillment shipping20211228 HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *VendorDirectFulfillmentShipping20211228 {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new vendor direct fulfillment shipping20211228 client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *VendorDirectFulfillmentShipping20211228 {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(VendorDirectFulfillmentShipping20211228)
	cli.Transport = transport
	cli.CustomerInvoices = customer_invoices.New(transport, formats)
	cli.VendorShipping = vendor_shipping.New(transport, formats)
	cli.VendorShippingLabels = vendor_shipping_labels.New(transport, formats)
	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// VendorDirectFulfillmentShipping20211228 is a client for vendor direct fulfillment shipping20211228
type VendorDirectFulfillmentShipping20211228 struct {
	CustomerInvoices customer_invoices.ClientService

	VendorShipping vendor_shipping.ClientService

	VendorShippingLabels vendor_shipping_labels.ClientService

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *VendorDirectFulfillmentShipping20211228) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport
	c.CustomerInvoices.SetTransport(transport)
	c.VendorShipping.SetTransport(transport)
	c.VendorShippingLabels.SetTransport(transport)
}
