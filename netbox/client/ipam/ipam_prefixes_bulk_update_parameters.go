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

// NewIpamPrefixesBulkUpdateParams creates a new IpamPrefixesBulkUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIpamPrefixesBulkUpdateParams() *IpamPrefixesBulkUpdateParams {
	return &IpamPrefixesBulkUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIpamPrefixesBulkUpdateParamsWithTimeout creates a new IpamPrefixesBulkUpdateParams object
// with the ability to set a timeout on a request.
func NewIpamPrefixesBulkUpdateParamsWithTimeout(timeout time.Duration) *IpamPrefixesBulkUpdateParams {
	return &IpamPrefixesBulkUpdateParams{
		timeout: timeout,
	}
}

// NewIpamPrefixesBulkUpdateParamsWithContext creates a new IpamPrefixesBulkUpdateParams object
// with the ability to set a context for a request.
func NewIpamPrefixesBulkUpdateParamsWithContext(ctx context.Context) *IpamPrefixesBulkUpdateParams {
	return &IpamPrefixesBulkUpdateParams{
		Context: ctx,
	}
}

// NewIpamPrefixesBulkUpdateParamsWithHTTPClient creates a new IpamPrefixesBulkUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewIpamPrefixesBulkUpdateParamsWithHTTPClient(client *http.Client) *IpamPrefixesBulkUpdateParams {
	return &IpamPrefixesBulkUpdateParams{
		HTTPClient: client,
	}
}

/*
IpamPrefixesBulkUpdateParams contains all the parameters to send to the API endpoint

	for the ipam prefixes bulk update operation.

	Typically these are written to a http.Request.
*/
type IpamPrefixesBulkUpdateParams struct {

	// Data.
	Data *models.WritablePrefix

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ipam prefixes bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamPrefixesBulkUpdateParams) WithDefaults() *IpamPrefixesBulkUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ipam prefixes bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamPrefixesBulkUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ipam prefixes bulk update params
func (o *IpamPrefixesBulkUpdateParams) WithTimeout(timeout time.Duration) *IpamPrefixesBulkUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam prefixes bulk update params
func (o *IpamPrefixesBulkUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam prefixes bulk update params
func (o *IpamPrefixesBulkUpdateParams) WithContext(ctx context.Context) *IpamPrefixesBulkUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam prefixes bulk update params
func (o *IpamPrefixesBulkUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam prefixes bulk update params
func (o *IpamPrefixesBulkUpdateParams) WithHTTPClient(client *http.Client) *IpamPrefixesBulkUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam prefixes bulk update params
func (o *IpamPrefixesBulkUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the ipam prefixes bulk update params
func (o *IpamPrefixesBulkUpdateParams) WithData(data *models.WritablePrefix) *IpamPrefixesBulkUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the ipam prefixes bulk update params
func (o *IpamPrefixesBulkUpdateParams) SetData(data *models.WritablePrefix) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *IpamPrefixesBulkUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
