// Code generated by go-swagger; DO NOT EDIT.

package lte_networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteLTENetworkIDSubscriberConfigRuleNamesRuleIDParams creates a new DeleteLTENetworkIDSubscriberConfigRuleNamesRuleIDParams object
// with the default values initialized.
func NewDeleteLTENetworkIDSubscriberConfigRuleNamesRuleIDParams() *DeleteLTENetworkIDSubscriberConfigRuleNamesRuleIDParams {
	var ()
	return &DeleteLTENetworkIDSubscriberConfigRuleNamesRuleIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteLTENetworkIDSubscriberConfigRuleNamesRuleIDParamsWithTimeout creates a new DeleteLTENetworkIDSubscriberConfigRuleNamesRuleIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteLTENetworkIDSubscriberConfigRuleNamesRuleIDParamsWithTimeout(timeout time.Duration) *DeleteLTENetworkIDSubscriberConfigRuleNamesRuleIDParams {
	var ()
	return &DeleteLTENetworkIDSubscriberConfigRuleNamesRuleIDParams{

		timeout: timeout,
	}
}

// NewDeleteLTENetworkIDSubscriberConfigRuleNamesRuleIDParamsWithContext creates a new DeleteLTENetworkIDSubscriberConfigRuleNamesRuleIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteLTENetworkIDSubscriberConfigRuleNamesRuleIDParamsWithContext(ctx context.Context) *DeleteLTENetworkIDSubscriberConfigRuleNamesRuleIDParams {
	var ()
	return &DeleteLTENetworkIDSubscriberConfigRuleNamesRuleIDParams{

		Context: ctx,
	}
}

// NewDeleteLTENetworkIDSubscriberConfigRuleNamesRuleIDParamsWithHTTPClient creates a new DeleteLTENetworkIDSubscriberConfigRuleNamesRuleIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteLTENetworkIDSubscriberConfigRuleNamesRuleIDParamsWithHTTPClient(client *http.Client) *DeleteLTENetworkIDSubscriberConfigRuleNamesRuleIDParams {
	var ()
	return &DeleteLTENetworkIDSubscriberConfigRuleNamesRuleIDParams{
		HTTPClient: client,
	}
}

/*DeleteLTENetworkIDSubscriberConfigRuleNamesRuleIDParams contains all the parameters to send to the API endpoint
for the delete LTE network ID subscriber config rule names rule ID operation typically these are written to a http.Request
*/
type DeleteLTENetworkIDSubscriberConfigRuleNamesRuleIDParams struct {

	/*NetworkID
	  Network ID

	*/
	NetworkID string
	/*RuleID
	  Rule Id

	*/
	RuleID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete LTE network ID subscriber config rule names rule ID params
func (o *DeleteLTENetworkIDSubscriberConfigRuleNamesRuleIDParams) WithTimeout(timeout time.Duration) *DeleteLTENetworkIDSubscriberConfigRuleNamesRuleIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete LTE network ID subscriber config rule names rule ID params
func (o *DeleteLTENetworkIDSubscriberConfigRuleNamesRuleIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete LTE network ID subscriber config rule names rule ID params
func (o *DeleteLTENetworkIDSubscriberConfigRuleNamesRuleIDParams) WithContext(ctx context.Context) *DeleteLTENetworkIDSubscriberConfigRuleNamesRuleIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete LTE network ID subscriber config rule names rule ID params
func (o *DeleteLTENetworkIDSubscriberConfigRuleNamesRuleIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete LTE network ID subscriber config rule names rule ID params
func (o *DeleteLTENetworkIDSubscriberConfigRuleNamesRuleIDParams) WithHTTPClient(client *http.Client) *DeleteLTENetworkIDSubscriberConfigRuleNamesRuleIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete LTE network ID subscriber config rule names rule ID params
func (o *DeleteLTENetworkIDSubscriberConfigRuleNamesRuleIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the delete LTE network ID subscriber config rule names rule ID params
func (o *DeleteLTENetworkIDSubscriberConfigRuleNamesRuleIDParams) WithNetworkID(networkID string) *DeleteLTENetworkIDSubscriberConfigRuleNamesRuleIDParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the delete LTE network ID subscriber config rule names rule ID params
func (o *DeleteLTENetworkIDSubscriberConfigRuleNamesRuleIDParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithRuleID adds the ruleID to the delete LTE network ID subscriber config rule names rule ID params
func (o *DeleteLTENetworkIDSubscriberConfigRuleNamesRuleIDParams) WithRuleID(ruleID string) *DeleteLTENetworkIDSubscriberConfigRuleNamesRuleIDParams {
	o.SetRuleID(ruleID)
	return o
}

// SetRuleID adds the ruleId to the delete LTE network ID subscriber config rule names rule ID params
func (o *DeleteLTENetworkIDSubscriberConfigRuleNamesRuleIDParams) SetRuleID(ruleID string) {
	o.RuleID = ruleID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteLTENetworkIDSubscriberConfigRuleNamesRuleIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param network_id
	if err := r.SetPathParam("network_id", o.NetworkID); err != nil {
		return err
	}

	// path param rule_id
	if err := r.SetPathParam("rule_id", o.RuleID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
