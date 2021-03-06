// Code generated by go-swagger; DO NOT EDIT.

package op

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

// NewDelStationParams creates a new DelStationParams object
// with the default values initialized.
func NewDelStationParams() *DelStationParams {
	var ()
	return &DelStationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDelStationParamsWithTimeout creates a new DelStationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDelStationParamsWithTimeout(timeout time.Duration) *DelStationParams {
	var ()
	return &DelStationParams{

		timeout: timeout,
	}
}

// NewDelStationParamsWithContext creates a new DelStationParams object
// with the default values initialized, and the ability to set a context for a request
func NewDelStationParamsWithContext(ctx context.Context) *DelStationParams {
	var ()
	return &DelStationParams{

		Context: ctx,
	}
}

// NewDelStationParamsWithHTTPClient creates a new DelStationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDelStationParamsWithHTTPClient(client *http.Client) *DelStationParams {
	var ()
	return &DelStationParams{
		HTTPClient: client,
	}
}

/*DelStationParams contains all the parameters to send to the API endpoint
for the del station operation typically these are written to a http.Request
*/
type DelStationParams struct {

	/*Args*/
	Args DelStationBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the del station params
func (o *DelStationParams) WithTimeout(timeout time.Duration) *DelStationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the del station params
func (o *DelStationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the del station params
func (o *DelStationParams) WithContext(ctx context.Context) *DelStationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the del station params
func (o *DelStationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the del station params
func (o *DelStationParams) WithHTTPClient(client *http.Client) *DelStationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the del station params
func (o *DelStationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithArgs adds the args to the del station params
func (o *DelStationParams) WithArgs(args DelStationBody) *DelStationParams {
	o.SetArgs(args)
	return o
}

// SetArgs adds the args to the del station params
func (o *DelStationParams) SetArgs(args DelStationBody) {
	o.Args = args
}

// WriteToRequest writes these params to a swagger request
func (o *DelStationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.Args); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
