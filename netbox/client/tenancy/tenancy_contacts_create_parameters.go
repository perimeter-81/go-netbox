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

package tenancy

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

// NewTenancyContactsCreateParams creates a new TenancyContactsCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTenancyContactsCreateParams() *TenancyContactsCreateParams {
	return &TenancyContactsCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTenancyContactsCreateParamsWithTimeout creates a new TenancyContactsCreateParams object
// with the ability to set a timeout on a request.
func NewTenancyContactsCreateParamsWithTimeout(timeout time.Duration) *TenancyContactsCreateParams {
	return &TenancyContactsCreateParams{
		timeout: timeout,
	}
}

// NewTenancyContactsCreateParamsWithContext creates a new TenancyContactsCreateParams object
// with the ability to set a context for a request.
func NewTenancyContactsCreateParamsWithContext(ctx context.Context) *TenancyContactsCreateParams {
	return &TenancyContactsCreateParams{
		Context: ctx,
	}
}

// NewTenancyContactsCreateParamsWithHTTPClient creates a new TenancyContactsCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewTenancyContactsCreateParamsWithHTTPClient(client *http.Client) *TenancyContactsCreateParams {
	return &TenancyContactsCreateParams{
		HTTPClient: client,
	}
}

/*
TenancyContactsCreateParams contains all the parameters to send to the API endpoint

	for the tenancy contacts create operation.

	Typically these are written to a http.Request.
*/
type TenancyContactsCreateParams struct {

	// Data.
	Data *models.WritableContact

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the tenancy contacts create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TenancyContactsCreateParams) WithDefaults() *TenancyContactsCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the tenancy contacts create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TenancyContactsCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the tenancy contacts create params
func (o *TenancyContactsCreateParams) WithTimeout(timeout time.Duration) *TenancyContactsCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the tenancy contacts create params
func (o *TenancyContactsCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the tenancy contacts create params
func (o *TenancyContactsCreateParams) WithContext(ctx context.Context) *TenancyContactsCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the tenancy contacts create params
func (o *TenancyContactsCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the tenancy contacts create params
func (o *TenancyContactsCreateParams) WithHTTPClient(client *http.Client) *TenancyContactsCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the tenancy contacts create params
func (o *TenancyContactsCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the tenancy contacts create params
func (o *TenancyContactsCreateParams) WithData(data *models.WritableContact) *TenancyContactsCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the tenancy contacts create params
func (o *TenancyContactsCreateParams) SetData(data *models.WritableContact) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *TenancyContactsCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
