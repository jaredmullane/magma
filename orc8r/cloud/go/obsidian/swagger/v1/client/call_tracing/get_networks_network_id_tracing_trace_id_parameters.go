// Code generated by go-swagger; DO NOT EDIT.

package call_tracing

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

// NewGetNetworksNetworkIDTracingTraceIDParams creates a new GetNetworksNetworkIDTracingTraceIDParams object
// with the default values initialized.
func NewGetNetworksNetworkIDTracingTraceIDParams() *GetNetworksNetworkIDTracingTraceIDParams {
	var ()
	return &GetNetworksNetworkIDTracingTraceIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetNetworksNetworkIDTracingTraceIDParamsWithTimeout creates a new GetNetworksNetworkIDTracingTraceIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetNetworksNetworkIDTracingTraceIDParamsWithTimeout(timeout time.Duration) *GetNetworksNetworkIDTracingTraceIDParams {
	var ()
	return &GetNetworksNetworkIDTracingTraceIDParams{

		timeout: timeout,
	}
}

// NewGetNetworksNetworkIDTracingTraceIDParamsWithContext creates a new GetNetworksNetworkIDTracingTraceIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetNetworksNetworkIDTracingTraceIDParamsWithContext(ctx context.Context) *GetNetworksNetworkIDTracingTraceIDParams {
	var ()
	return &GetNetworksNetworkIDTracingTraceIDParams{

		Context: ctx,
	}
}

// NewGetNetworksNetworkIDTracingTraceIDParamsWithHTTPClient creates a new GetNetworksNetworkIDTracingTraceIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetNetworksNetworkIDTracingTraceIDParamsWithHTTPClient(client *http.Client) *GetNetworksNetworkIDTracingTraceIDParams {
	var ()
	return &GetNetworksNetworkIDTracingTraceIDParams{
		HTTPClient: client,
	}
}

/*GetNetworksNetworkIDTracingTraceIDParams contains all the parameters to send to the API endpoint
for the get networks network ID tracing trace ID operation typically these are written to a http.Request
*/
type GetNetworksNetworkIDTracingTraceIDParams struct {

	/*NetworkID
	  Network ID

	*/
	NetworkID string
	/*TraceID
	  Unique ID of call trace

	*/
	TraceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get networks network ID tracing trace ID params
func (o *GetNetworksNetworkIDTracingTraceIDParams) WithTimeout(timeout time.Duration) *GetNetworksNetworkIDTracingTraceIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get networks network ID tracing trace ID params
func (o *GetNetworksNetworkIDTracingTraceIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get networks network ID tracing trace ID params
func (o *GetNetworksNetworkIDTracingTraceIDParams) WithContext(ctx context.Context) *GetNetworksNetworkIDTracingTraceIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get networks network ID tracing trace ID params
func (o *GetNetworksNetworkIDTracingTraceIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get networks network ID tracing trace ID params
func (o *GetNetworksNetworkIDTracingTraceIDParams) WithHTTPClient(client *http.Client) *GetNetworksNetworkIDTracingTraceIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get networks network ID tracing trace ID params
func (o *GetNetworksNetworkIDTracingTraceIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the get networks network ID tracing trace ID params
func (o *GetNetworksNetworkIDTracingTraceIDParams) WithNetworkID(networkID string) *GetNetworksNetworkIDTracingTraceIDParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get networks network ID tracing trace ID params
func (o *GetNetworksNetworkIDTracingTraceIDParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithTraceID adds the traceID to the get networks network ID tracing trace ID params
func (o *GetNetworksNetworkIDTracingTraceIDParams) WithTraceID(traceID string) *GetNetworksNetworkIDTracingTraceIDParams {
	o.SetTraceID(traceID)
	return o
}

// SetTraceID adds the traceId to the get networks network ID tracing trace ID params
func (o *GetNetworksNetworkIDTracingTraceIDParams) SetTraceID(traceID string) {
	o.TraceID = traceID
}

// WriteToRequest writes these params to a swagger request
func (o *GetNetworksNetworkIDTracingTraceIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param network_id
	if err := r.SetPathParam("network_id", o.NetworkID); err != nil {
		return err
	}

	// path param trace_id
	if err := r.SetPathParam("trace_id", o.TraceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
