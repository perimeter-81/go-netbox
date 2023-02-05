// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2020 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package wireless

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

	"github.com/perimeter-81/go-netbox/netbox/models"
)

// NewWirelessWirelessLansBulkPartialUpdateParams creates a new WirelessWirelessLansBulkPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWirelessWirelessLansBulkPartialUpdateParams() *WirelessWirelessLansBulkPartialUpdateParams {
	return &WirelessWirelessLansBulkPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewWirelessWirelessLansBulkPartialUpdateParamsWithTimeout creates a new WirelessWirelessLansBulkPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewWirelessWirelessLansBulkPartialUpdateParamsWithTimeout(timeout time.Duration) *WirelessWirelessLansBulkPartialUpdateParams {
	return &WirelessWirelessLansBulkPartialUpdateParams{
		timeout: timeout,
	}
}

// NewWirelessWirelessLansBulkPartialUpdateParamsWithContext creates a new WirelessWirelessLansBulkPartialUpdateParams object
// with the ability to set a context for a request.
func NewWirelessWirelessLansBulkPartialUpdateParamsWithContext(ctx context.Context) *WirelessWirelessLansBulkPartialUpdateParams {
	return &WirelessWirelessLansBulkPartialUpdateParams{
		Context: ctx,
	}
}

// NewWirelessWirelessLansBulkPartialUpdateParamsWithHTTPClient creates a new WirelessWirelessLansBulkPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewWirelessWirelessLansBulkPartialUpdateParamsWithHTTPClient(client *http.Client) *WirelessWirelessLansBulkPartialUpdateParams {
	return &WirelessWirelessLansBulkPartialUpdateParams{
		HTTPClient: client,
	}
}

/*
WirelessWirelessLansBulkPartialUpdateParams contains all the parameters to send to the API endpoint

	for the wireless wireless lans bulk partial update operation.

	Typically these are written to a http.Request.
*/
type WirelessWirelessLansBulkPartialUpdateParams struct {

	// Data.
	Data *models.WritableWirelessLAN

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the wireless wireless lans bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WirelessWirelessLansBulkPartialUpdateParams) WithDefaults() *WirelessWirelessLansBulkPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the wireless wireless lans bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WirelessWirelessLansBulkPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the wireless wireless lans bulk partial update params
func (o *WirelessWirelessLansBulkPartialUpdateParams) WithTimeout(timeout time.Duration) *WirelessWirelessLansBulkPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the wireless wireless lans bulk partial update params
func (o *WirelessWirelessLansBulkPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the wireless wireless lans bulk partial update params
func (o *WirelessWirelessLansBulkPartialUpdateParams) WithContext(ctx context.Context) *WirelessWirelessLansBulkPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the wireless wireless lans bulk partial update params
func (o *WirelessWirelessLansBulkPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the wireless wireless lans bulk partial update params
func (o *WirelessWirelessLansBulkPartialUpdateParams) WithHTTPClient(client *http.Client) *WirelessWirelessLansBulkPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the wireless wireless lans bulk partial update params
func (o *WirelessWirelessLansBulkPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the wireless wireless lans bulk partial update params
func (o *WirelessWirelessLansBulkPartialUpdateParams) WithData(data *models.WritableWirelessLAN) *WirelessWirelessLansBulkPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the wireless wireless lans bulk partial update params
func (o *WirelessWirelessLansBulkPartialUpdateParams) SetData(data *models.WritableWirelessLAN) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *WirelessWirelessLansBulkPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
