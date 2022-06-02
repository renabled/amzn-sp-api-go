// Code generated by go-swagger; DO NOT EDIT.

package notifications

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/xamandar/amzn-sp-api-go/api/notifications/notifications_models"
)

// NewCreateSubscriptionParams creates a new CreateSubscriptionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateSubscriptionParams() *CreateSubscriptionParams {
	return &CreateSubscriptionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateSubscriptionParamsWithTimeout creates a new CreateSubscriptionParams object
// with the ability to set a timeout on a request.
func NewCreateSubscriptionParamsWithTimeout(timeout time.Duration) *CreateSubscriptionParams {
	return &CreateSubscriptionParams{
		timeout: timeout,
	}
}

// NewCreateSubscriptionParamsWithContext creates a new CreateSubscriptionParams object
// with the ability to set a context for a request.
func NewCreateSubscriptionParamsWithContext(ctx context.Context) *CreateSubscriptionParams {
	return &CreateSubscriptionParams{
		Context: ctx,
	}
}

// NewCreateSubscriptionParamsWithHTTPClient creates a new CreateSubscriptionParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateSubscriptionParamsWithHTTPClient(client *http.Client) *CreateSubscriptionParams {
	return &CreateSubscriptionParams{
		HTTPClient: client,
	}
}

/* CreateSubscriptionParams contains all the parameters to send to the API endpoint
   for the create subscription operation.

   Typically these are written to a http.Request.
*/
type CreateSubscriptionParams struct {

	// Body.
	Body *notifications_models.CreateSubscriptionRequest

	/* NotificationType.

	    The type of notification.

	For more information about notification types, see [the Notifications API Use Case Guide](doc:notifications-api-v1-use-case-guide).
	*/
	NotificationType string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create subscription params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateSubscriptionParams) WithDefaults() *CreateSubscriptionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create subscription params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateSubscriptionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create subscription params
func (o *CreateSubscriptionParams) WithTimeout(timeout time.Duration) *CreateSubscriptionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create subscription params
func (o *CreateSubscriptionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create subscription params
func (o *CreateSubscriptionParams) WithContext(ctx context.Context) *CreateSubscriptionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create subscription params
func (o *CreateSubscriptionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create subscription params
func (o *CreateSubscriptionParams) WithHTTPClient(client *http.Client) *CreateSubscriptionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create subscription params
func (o *CreateSubscriptionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create subscription params
func (o *CreateSubscriptionParams) WithBody(body *notifications_models.CreateSubscriptionRequest) *CreateSubscriptionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create subscription params
func (o *CreateSubscriptionParams) SetBody(body *notifications_models.CreateSubscriptionRequest) {
	o.Body = body
}

// WithNotificationType adds the notificationType to the create subscription params
func (o *CreateSubscriptionParams) WithNotificationType(notificationType string) *CreateSubscriptionParams {
	o.SetNotificationType(notificationType)
	return o
}

// SetNotificationType adds the notificationType to the create subscription params
func (o *CreateSubscriptionParams) SetNotificationType(notificationType string) {
	o.NotificationType = notificationType
}

// WriteToRequest writes these params to a swagger request
func (o *CreateSubscriptionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param notificationType
	if err := r.SetPathParam("notificationType", o.NotificationType); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
