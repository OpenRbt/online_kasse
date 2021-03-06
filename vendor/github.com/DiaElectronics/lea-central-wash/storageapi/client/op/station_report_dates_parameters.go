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

// NewStationReportDatesParams creates a new StationReportDatesParams object
// with the default values initialized.
func NewStationReportDatesParams() *StationReportDatesParams {
	var ()
	return &StationReportDatesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewStationReportDatesParamsWithTimeout creates a new StationReportDatesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewStationReportDatesParamsWithTimeout(timeout time.Duration) *StationReportDatesParams {
	var ()
	return &StationReportDatesParams{

		timeout: timeout,
	}
}

// NewStationReportDatesParamsWithContext creates a new StationReportDatesParams object
// with the default values initialized, and the ability to set a context for a request
func NewStationReportDatesParamsWithContext(ctx context.Context) *StationReportDatesParams {
	var ()
	return &StationReportDatesParams{

		Context: ctx,
	}
}

// NewStationReportDatesParamsWithHTTPClient creates a new StationReportDatesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStationReportDatesParamsWithHTTPClient(client *http.Client) *StationReportDatesParams {
	var ()
	return &StationReportDatesParams{
		HTTPClient: client,
	}
}

/*StationReportDatesParams contains all the parameters to send to the API endpoint
for the station report dates operation typically these are written to a http.Request
*/
type StationReportDatesParams struct {

	/*Args*/
	Args StationReportDatesBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the station report dates params
func (o *StationReportDatesParams) WithTimeout(timeout time.Duration) *StationReportDatesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the station report dates params
func (o *StationReportDatesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the station report dates params
func (o *StationReportDatesParams) WithContext(ctx context.Context) *StationReportDatesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the station report dates params
func (o *StationReportDatesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the station report dates params
func (o *StationReportDatesParams) WithHTTPClient(client *http.Client) *StationReportDatesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the station report dates params
func (o *StationReportDatesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithArgs adds the args to the station report dates params
func (o *StationReportDatesParams) WithArgs(args StationReportDatesBody) *StationReportDatesParams {
	o.SetArgs(args)
	return o
}

// SetArgs adds the args to the station report dates params
func (o *StationReportDatesParams) SetArgs(args StationReportDatesBody) {
	o.Args = args
}

// WriteToRequest writes these params to a swagger request
func (o *StationReportDatesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
