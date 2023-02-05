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

package ipam

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

// NewIpamServicesBulkPartialUpdateParams creates a new IpamServicesBulkPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIpamServicesBulkPartialUpdateParams() *IpamServicesBulkPartialUpdateParams {
	return &IpamServicesBulkPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIpamServicesBulkPartialUpdateParamsWithTimeout creates a new IpamServicesBulkPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewIpamServicesBulkPartialUpdateParamsWithTimeout(timeout time.Duration) *IpamServicesBulkPartialUpdateParams {
	return &IpamServicesBulkPartialUpdateParams{
		timeout: timeout,
	}
}

// NewIpamServicesBulkPartialUpdateParamsWithContext creates a new IpamServicesBulkPartialUpdateParams object
// with the ability to set a context for a request.
func NewIpamServicesBulkPartialUpdateParamsWithContext(ctx context.Context) *IpamServicesBulkPartialUpdateParams {
	return &IpamServicesBulkPartialUpdateParams{
		Context: ctx,
	}
}

// NewIpamServicesBulkPartialUpdateParamsWithHTTPClient creates a new IpamServicesBulkPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewIpamServicesBulkPartialUpdateParamsWithHTTPClient(client *http.Client) *IpamServicesBulkPartialUpdateParams {
	return &IpamServicesBulkPartialUpdateParams{
		HTTPClient: client,
	}
}

/*
IpamServicesBulkPartialUpdateParams contains all the parameters to send to the API endpoint

	for the ipam services bulk partial update operation.

	Typically these are written to a http.Request.
*/
type IpamServicesBulkPartialUpdateParams struct {

	// Data.
	Data *models.WritableService

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ipam services bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamServicesBulkPartialUpdateParams) WithDefaults() *IpamServicesBulkPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ipam services bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamServicesBulkPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ipam services bulk partial update params
func (o *IpamServicesBulkPartialUpdateParams) WithTimeout(timeout time.Duration) *IpamServicesBulkPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam services bulk partial update params
func (o *IpamServicesBulkPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam services bulk partial update params
func (o *IpamServicesBulkPartialUpdateParams) WithContext(ctx context.Context) *IpamServicesBulkPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam services bulk partial update params
func (o *IpamServicesBulkPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam services bulk partial update params
func (o *IpamServicesBulkPartialUpdateParams) WithHTTPClient(client *http.Client) *IpamServicesBulkPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam services bulk partial update params
func (o *IpamServicesBulkPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the ipam services bulk partial update params
func (o *IpamServicesBulkPartialUpdateParams) WithData(data *models.WritableService) *IpamServicesBulkPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the ipam services bulk partial update params
func (o *IpamServicesBulkPartialUpdateParams) SetData(data *models.WritableService) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *IpamServicesBulkPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
