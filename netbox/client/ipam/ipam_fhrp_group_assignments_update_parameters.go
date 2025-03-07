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
	"github.com/go-openapi/swag"

	"github.com/perimeter-81/go-netbox/netbox/models"
)

// NewIpamFhrpGroupAssignmentsUpdateParams creates a new IpamFhrpGroupAssignmentsUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIpamFhrpGroupAssignmentsUpdateParams() *IpamFhrpGroupAssignmentsUpdateParams {
	return &IpamFhrpGroupAssignmentsUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIpamFhrpGroupAssignmentsUpdateParamsWithTimeout creates a new IpamFhrpGroupAssignmentsUpdateParams object
// with the ability to set a timeout on a request.
func NewIpamFhrpGroupAssignmentsUpdateParamsWithTimeout(timeout time.Duration) *IpamFhrpGroupAssignmentsUpdateParams {
	return &IpamFhrpGroupAssignmentsUpdateParams{
		timeout: timeout,
	}
}

// NewIpamFhrpGroupAssignmentsUpdateParamsWithContext creates a new IpamFhrpGroupAssignmentsUpdateParams object
// with the ability to set a context for a request.
func NewIpamFhrpGroupAssignmentsUpdateParamsWithContext(ctx context.Context) *IpamFhrpGroupAssignmentsUpdateParams {
	return &IpamFhrpGroupAssignmentsUpdateParams{
		Context: ctx,
	}
}

// NewIpamFhrpGroupAssignmentsUpdateParamsWithHTTPClient creates a new IpamFhrpGroupAssignmentsUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewIpamFhrpGroupAssignmentsUpdateParamsWithHTTPClient(client *http.Client) *IpamFhrpGroupAssignmentsUpdateParams {
	return &IpamFhrpGroupAssignmentsUpdateParams{
		HTTPClient: client,
	}
}

/*
IpamFhrpGroupAssignmentsUpdateParams contains all the parameters to send to the API endpoint

	for the ipam fhrp group assignments update operation.

	Typically these are written to a http.Request.
*/
type IpamFhrpGroupAssignmentsUpdateParams struct {

	// Data.
	Data *models.WritableFHRPGroupAssignment

	/* ID.

	   A unique integer value identifying this FHRP group assignment.
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ipam fhrp group assignments update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamFhrpGroupAssignmentsUpdateParams) WithDefaults() *IpamFhrpGroupAssignmentsUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ipam fhrp group assignments update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamFhrpGroupAssignmentsUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ipam fhrp group assignments update params
func (o *IpamFhrpGroupAssignmentsUpdateParams) WithTimeout(timeout time.Duration) *IpamFhrpGroupAssignmentsUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam fhrp group assignments update params
func (o *IpamFhrpGroupAssignmentsUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam fhrp group assignments update params
func (o *IpamFhrpGroupAssignmentsUpdateParams) WithContext(ctx context.Context) *IpamFhrpGroupAssignmentsUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam fhrp group assignments update params
func (o *IpamFhrpGroupAssignmentsUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam fhrp group assignments update params
func (o *IpamFhrpGroupAssignmentsUpdateParams) WithHTTPClient(client *http.Client) *IpamFhrpGroupAssignmentsUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam fhrp group assignments update params
func (o *IpamFhrpGroupAssignmentsUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the ipam fhrp group assignments update params
func (o *IpamFhrpGroupAssignmentsUpdateParams) WithData(data *models.WritableFHRPGroupAssignment) *IpamFhrpGroupAssignmentsUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the ipam fhrp group assignments update params
func (o *IpamFhrpGroupAssignmentsUpdateParams) SetData(data *models.WritableFHRPGroupAssignment) {
	o.Data = data
}

// WithID adds the id to the ipam fhrp group assignments update params
func (o *IpamFhrpGroupAssignmentsUpdateParams) WithID(id int64) *IpamFhrpGroupAssignmentsUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the ipam fhrp group assignments update params
func (o *IpamFhrpGroupAssignmentsUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *IpamFhrpGroupAssignmentsUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
