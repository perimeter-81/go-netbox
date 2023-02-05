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

// NewIpamServicesBulkUpdateParams creates a new IpamServicesBulkUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIpamServicesBulkUpdateParams() *IpamServicesBulkUpdateParams {
	return &IpamServicesBulkUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIpamServicesBulkUpdateParamsWithTimeout creates a new IpamServicesBulkUpdateParams object
// with the ability to set a timeout on a request.
func NewIpamServicesBulkUpdateParamsWithTimeout(timeout time.Duration) *IpamServicesBulkUpdateParams {
	return &IpamServicesBulkUpdateParams{
		timeout: timeout,
	}
}

// NewIpamServicesBulkUpdateParamsWithContext creates a new IpamServicesBulkUpdateParams object
// with the ability to set a context for a request.
func NewIpamServicesBulkUpdateParamsWithContext(ctx context.Context) *IpamServicesBulkUpdateParams {
	return &IpamServicesBulkUpdateParams{
		Context: ctx,
	}
}

// NewIpamServicesBulkUpdateParamsWithHTTPClient creates a new IpamServicesBulkUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewIpamServicesBulkUpdateParamsWithHTTPClient(client *http.Client) *IpamServicesBulkUpdateParams {
	return &IpamServicesBulkUpdateParams{
		HTTPClient: client,
	}
}

/*
IpamServicesBulkUpdateParams contains all the parameters to send to the API endpoint

	for the ipam services bulk update operation.

	Typically these are written to a http.Request.
*/
type IpamServicesBulkUpdateParams struct {

	// Data.
	Data *models.WritableService

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ipam services bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamServicesBulkUpdateParams) WithDefaults() *IpamServicesBulkUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ipam services bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamServicesBulkUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ipam services bulk update params
func (o *IpamServicesBulkUpdateParams) WithTimeout(timeout time.Duration) *IpamServicesBulkUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam services bulk update params
func (o *IpamServicesBulkUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam services bulk update params
func (o *IpamServicesBulkUpdateParams) WithContext(ctx context.Context) *IpamServicesBulkUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam services bulk update params
func (o *IpamServicesBulkUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam services bulk update params
func (o *IpamServicesBulkUpdateParams) WithHTTPClient(client *http.Client) *IpamServicesBulkUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam services bulk update params
func (o *IpamServicesBulkUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the ipam services bulk update params
func (o *IpamServicesBulkUpdateParams) WithData(data *models.WritableService) *IpamServicesBulkUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the ipam services bulk update params
func (o *IpamServicesBulkUpdateParams) SetData(data *models.WritableService) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *IpamServicesBulkUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
