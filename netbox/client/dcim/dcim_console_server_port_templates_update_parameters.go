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

// NewDcimConsoleServerPortTemplatesUpdateParams creates a new DcimConsoleServerPortTemplatesUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimConsoleServerPortTemplatesUpdateParams() *DcimConsoleServerPortTemplatesUpdateParams {
	return &DcimConsoleServerPortTemplatesUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimConsoleServerPortTemplatesUpdateParamsWithTimeout creates a new DcimConsoleServerPortTemplatesUpdateParams object
// with the ability to set a timeout on a request.
func NewDcimConsoleServerPortTemplatesUpdateParamsWithTimeout(timeout time.Duration) *DcimConsoleServerPortTemplatesUpdateParams {
	return &DcimConsoleServerPortTemplatesUpdateParams{
		timeout: timeout,
	}
}

// NewDcimConsoleServerPortTemplatesUpdateParamsWithContext creates a new DcimConsoleServerPortTemplatesUpdateParams object
// with the ability to set a context for a request.
func NewDcimConsoleServerPortTemplatesUpdateParamsWithContext(ctx context.Context) *DcimConsoleServerPortTemplatesUpdateParams {
	return &DcimConsoleServerPortTemplatesUpdateParams{
		Context: ctx,
	}
}

// NewDcimConsoleServerPortTemplatesUpdateParamsWithHTTPClient creates a new DcimConsoleServerPortTemplatesUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimConsoleServerPortTemplatesUpdateParamsWithHTTPClient(client *http.Client) *DcimConsoleServerPortTemplatesUpdateParams {
	return &DcimConsoleServerPortTemplatesUpdateParams{
		HTTPClient: client,
	}
}

/*
DcimConsoleServerPortTemplatesUpdateParams contains all the parameters to send to the API endpoint

	for the dcim console server port templates update operation.

	Typically these are written to a http.Request.
*/
type DcimConsoleServerPortTemplatesUpdateParams struct {

	// Data.
	Data *models.WritableConsoleServerPortTemplate

	/* ID.

	   A unique integer value identifying this console server port template.
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim console server port templates update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimConsoleServerPortTemplatesUpdateParams) WithDefaults() *DcimConsoleServerPortTemplatesUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim console server port templates update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimConsoleServerPortTemplatesUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim console server port templates update params
func (o *DcimConsoleServerPortTemplatesUpdateParams) WithTimeout(timeout time.Duration) *DcimConsoleServerPortTemplatesUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim console server port templates update params
func (o *DcimConsoleServerPortTemplatesUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim console server port templates update params
func (o *DcimConsoleServerPortTemplatesUpdateParams) WithContext(ctx context.Context) *DcimConsoleServerPortTemplatesUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim console server port templates update params
func (o *DcimConsoleServerPortTemplatesUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim console server port templates update params
func (o *DcimConsoleServerPortTemplatesUpdateParams) WithHTTPClient(client *http.Client) *DcimConsoleServerPortTemplatesUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim console server port templates update params
func (o *DcimConsoleServerPortTemplatesUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim console server port templates update params
func (o *DcimConsoleServerPortTemplatesUpdateParams) WithData(data *models.WritableConsoleServerPortTemplate) *DcimConsoleServerPortTemplatesUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim console server port templates update params
func (o *DcimConsoleServerPortTemplatesUpdateParams) SetData(data *models.WritableConsoleServerPortTemplate) {
	o.Data = data
}

// WithID adds the id to the dcim console server port templates update params
func (o *DcimConsoleServerPortTemplatesUpdateParams) WithID(id int64) *DcimConsoleServerPortTemplatesUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim console server port templates update params
func (o *DcimConsoleServerPortTemplatesUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimConsoleServerPortTemplatesUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
