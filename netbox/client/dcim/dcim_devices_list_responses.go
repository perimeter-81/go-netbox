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
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/perimeter-81/go-netbox/netbox/models"
)

// DcimDevicesListReader is a Reader for the DcimDevicesList structure.
type DcimDevicesListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimDevicesListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimDevicesListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimDevicesListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimDevicesListOK creates a DcimDevicesListOK with default headers values
func NewDcimDevicesListOK() *DcimDevicesListOK {
	return &DcimDevicesListOK{}
}

/*
DcimDevicesListOK describes a response with status code 200, with default header values.

DcimDevicesListOK dcim devices list o k
*/
type DcimDevicesListOK struct {
	Payload *DcimDevicesListOKBody
}

// IsSuccess returns true when this dcim devices list o k response has a 2xx status code
func (o *DcimDevicesListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim devices list o k response has a 3xx status code
func (o *DcimDevicesListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim devices list o k response has a 4xx status code
func (o *DcimDevicesListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim devices list o k response has a 5xx status code
func (o *DcimDevicesListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim devices list o k response a status code equal to that given
func (o *DcimDevicesListOK) IsCode(code int) bool {
	return code == 200
}

func (o *DcimDevicesListOK) Error() string {
	return fmt.Sprintf("[GET /dcim/devices/][%d] dcimDevicesListOK  %+v", 200, o.Payload)
}

func (o *DcimDevicesListOK) String() string {
	return fmt.Sprintf("[GET /dcim/devices/][%d] dcimDevicesListOK  %+v", 200, o.Payload)
}

func (o *DcimDevicesListOK) GetPayload() *DcimDevicesListOKBody {
	return o.Payload
}

func (o *DcimDevicesListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(DcimDevicesListOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimDevicesListDefault creates a DcimDevicesListDefault with default headers values
func NewDcimDevicesListDefault(code int) *DcimDevicesListDefault {
	return &DcimDevicesListDefault{
		_statusCode: code,
	}
}

/*
DcimDevicesListDefault describes a response with status code -1, with default header values.

DcimDevicesListDefault dcim devices list default
*/
type DcimDevicesListDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim devices list default response
func (o *DcimDevicesListDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim devices list default response has a 2xx status code
func (o *DcimDevicesListDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim devices list default response has a 3xx status code
func (o *DcimDevicesListDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim devices list default response has a 4xx status code
func (o *DcimDevicesListDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim devices list default response has a 5xx status code
func (o *DcimDevicesListDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim devices list default response a status code equal to that given
func (o *DcimDevicesListDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimDevicesListDefault) Error() string {
	return fmt.Sprintf("[GET /dcim/devices/][%d] dcim_devices_list default  %+v", o._statusCode, o.Payload)
}

func (o *DcimDevicesListDefault) String() string {
	return fmt.Sprintf("[GET /dcim/devices/][%d] dcim_devices_list default  %+v", o._statusCode, o.Payload)
}

func (o *DcimDevicesListDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimDevicesListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
DcimDevicesListOKBody dcim devices list o k body
swagger:model DcimDevicesListOKBody
*/
type DcimDevicesListOKBody struct {

	// count
	// Required: true
	Count *int64 `json:"count"`

	// next
	// Format: uri
	Next *strfmt.URI `json:"next,omitempty"`

	// previous
	// Format: uri
	Previous *strfmt.URI `json:"previous,omitempty"`

	// results
	// Required: true
	Results []*models.DeviceWithConfigContext `json:"results"`
}

// Validate validates this dcim devices list o k body
func (o *DcimDevicesListOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCount(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateNext(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePrevious(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateResults(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DcimDevicesListOKBody) validateCount(formats strfmt.Registry) error {

	if err := validate.Required("dcimDevicesListOK"+"."+"count", "body", o.Count); err != nil {
		return err
	}

	return nil
}

func (o *DcimDevicesListOKBody) validateNext(formats strfmt.Registry) error {
	if swag.IsZero(o.Next) { // not required
		return nil
	}

	if err := validate.FormatOf("dcimDevicesListOK"+"."+"next", "body", "uri", o.Next.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *DcimDevicesListOKBody) validatePrevious(formats strfmt.Registry) error {
	if swag.IsZero(o.Previous) { // not required
		return nil
	}

	if err := validate.FormatOf("dcimDevicesListOK"+"."+"previous", "body", "uri", o.Previous.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *DcimDevicesListOKBody) validateResults(formats strfmt.Registry) error {

	if err := validate.Required("dcimDevicesListOK"+"."+"results", "body", o.Results); err != nil {
		return err
	}

	for i := 0; i < len(o.Results); i++ {
		if swag.IsZero(o.Results[i]) { // not required
			continue
		}

		if o.Results[i] != nil {
			if err := o.Results[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dcimDevicesListOK" + "." + "results" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("dcimDevicesListOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this dcim devices list o k body based on the context it is used
func (o *DcimDevicesListOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateResults(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DcimDevicesListOKBody) contextValidateResults(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Results); i++ {

		if o.Results[i] != nil {
			if err := o.Results[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dcimDevicesListOK" + "." + "results" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("dcimDevicesListOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *DcimDevicesListOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DcimDevicesListOKBody) UnmarshalBinary(b []byte) error {
	var res DcimDevicesListOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
