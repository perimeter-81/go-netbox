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

package dcim

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
	"github.com/go-openapi/swag"

	"github.com/perimeter-81/go-netbox/netbox/models"
)

// NewDcimDevicesPartialUpdateParams creates a new DcimDevicesPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimDevicesPartialUpdateParams() *DcimDevicesPartialUpdateParams {
	return &DcimDevicesPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimDevicesPartialUpdateParamsWithTimeout creates a new DcimDevicesPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewDcimDevicesPartialUpdateParamsWithTimeout(timeout time.Duration) *DcimDevicesPartialUpdateParams {
	return &DcimDevicesPartialUpdateParams{
		timeout: timeout,
	}
}

// NewDcimDevicesPartialUpdateParamsWithContext creates a new DcimDevicesPartialUpdateParams object
// with the ability to set a context for a request.
func NewDcimDevicesPartialUpdateParamsWithContext(ctx context.Context) *DcimDevicesPartialUpdateParams {
	return &DcimDevicesPartialUpdateParams{
		Context: ctx,
	}
}

// NewDcimDevicesPartialUpdateParamsWithHTTPClient creates a new DcimDevicesPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimDevicesPartialUpdateParamsWithHTTPClient(client *http.Client) *DcimDevicesPartialUpdateParams {
	return &DcimDevicesPartialUpdateParams{
		HTTPClient: client,
	}
}

/*
DcimDevicesPartialUpdateParams contains all the parameters to send to the API endpoint

	for the dcim devices partial update operation.

	Typically these are written to a http.Request.
*/
type DcimDevicesPartialUpdateParams struct {

	// Data.
	Data *models.WritableDeviceWithConfigContext

	/* ID.

	   A unique integer value identifying this device.
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim devices partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimDevicesPartialUpdateParams) WithDefaults() *DcimDevicesPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim devices partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimDevicesPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim devices partial update params
func (o *DcimDevicesPartialUpdateParams) WithTimeout(timeout time.Duration) *DcimDevicesPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim devices partial update params
func (o *DcimDevicesPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim devices partial update params
func (o *DcimDevicesPartialUpdateParams) WithContext(ctx context.Context) *DcimDevicesPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim devices partial update params
func (o *DcimDevicesPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim devices partial update params
func (o *DcimDevicesPartialUpdateParams) WithHTTPClient(client *http.Client) *DcimDevicesPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim devices partial update params
func (o *DcimDevicesPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim devices partial update params
func (o *DcimDevicesPartialUpdateParams) WithData(data *models.WritableDeviceWithConfigContext) *DcimDevicesPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim devices partial update params
func (o *DcimDevicesPartialUpdateParams) SetData(data *models.WritableDeviceWithConfigContext) {
	o.Data = data
}

// WithID adds the id to the dcim devices partial update params
func (o *DcimDevicesPartialUpdateParams) WithID(id int64) *DcimDevicesPartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim devices partial update params
func (o *DcimDevicesPartialUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimDevicesPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
