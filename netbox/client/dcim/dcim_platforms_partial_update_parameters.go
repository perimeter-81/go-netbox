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

// NewDcimPlatformsPartialUpdateParams creates a new DcimPlatformsPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimPlatformsPartialUpdateParams() *DcimPlatformsPartialUpdateParams {
	return &DcimPlatformsPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimPlatformsPartialUpdateParamsWithTimeout creates a new DcimPlatformsPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewDcimPlatformsPartialUpdateParamsWithTimeout(timeout time.Duration) *DcimPlatformsPartialUpdateParams {
	return &DcimPlatformsPartialUpdateParams{
		timeout: timeout,
	}
}

// NewDcimPlatformsPartialUpdateParamsWithContext creates a new DcimPlatformsPartialUpdateParams object
// with the ability to set a context for a request.
func NewDcimPlatformsPartialUpdateParamsWithContext(ctx context.Context) *DcimPlatformsPartialUpdateParams {
	return &DcimPlatformsPartialUpdateParams{
		Context: ctx,
	}
}

// NewDcimPlatformsPartialUpdateParamsWithHTTPClient creates a new DcimPlatformsPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimPlatformsPartialUpdateParamsWithHTTPClient(client *http.Client) *DcimPlatformsPartialUpdateParams {
	return &DcimPlatformsPartialUpdateParams{
		HTTPClient: client,
	}
}

/*
DcimPlatformsPartialUpdateParams contains all the parameters to send to the API endpoint

	for the dcim platforms partial update operation.

	Typically these are written to a http.Request.
*/
type DcimPlatformsPartialUpdateParams struct {

	// Data.
	Data *models.WritablePlatform

	/* ID.

	   A unique integer value identifying this platform.
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim platforms partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimPlatformsPartialUpdateParams) WithDefaults() *DcimPlatformsPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim platforms partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimPlatformsPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim platforms partial update params
func (o *DcimPlatformsPartialUpdateParams) WithTimeout(timeout time.Duration) *DcimPlatformsPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim platforms partial update params
func (o *DcimPlatformsPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim platforms partial update params
func (o *DcimPlatformsPartialUpdateParams) WithContext(ctx context.Context) *DcimPlatformsPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim platforms partial update params
func (o *DcimPlatformsPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim platforms partial update params
func (o *DcimPlatformsPartialUpdateParams) WithHTTPClient(client *http.Client) *DcimPlatformsPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim platforms partial update params
func (o *DcimPlatformsPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim platforms partial update params
func (o *DcimPlatformsPartialUpdateParams) WithData(data *models.WritablePlatform) *DcimPlatformsPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim platforms partial update params
func (o *DcimPlatformsPartialUpdateParams) SetData(data *models.WritablePlatform) {
	o.Data = data
}

// WithID adds the id to the dcim platforms partial update params
func (o *DcimPlatformsPartialUpdateParams) WithID(id int64) *DcimPlatformsPartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim platforms partial update params
func (o *DcimPlatformsPartialUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimPlatformsPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
