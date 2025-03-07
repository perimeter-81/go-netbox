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

// IpamL2vpnTerminationsListReader is a Reader for the IpamL2vpnTerminationsList structure.
type IpamL2vpnTerminationsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamL2vpnTerminationsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamL2vpnTerminationsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIpamL2vpnTerminationsListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpamL2vpnTerminationsListOK creates a IpamL2vpnTerminationsListOK with default headers values
func NewIpamL2vpnTerminationsListOK() *IpamL2vpnTerminationsListOK {
	return &IpamL2vpnTerminationsListOK{}
}

/*
IpamL2vpnTerminationsListOK describes a response with status code 200, with default header values.

IpamL2vpnTerminationsListOK ipam l2vpn terminations list o k
*/
type IpamL2vpnTerminationsListOK struct {
	Payload *IpamL2vpnTerminationsListOKBody
}

// IsSuccess returns true when this ipam l2vpn terminations list o k response has a 2xx status code
func (o *IpamL2vpnTerminationsListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam l2vpn terminations list o k response has a 3xx status code
func (o *IpamL2vpnTerminationsListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam l2vpn terminations list o k response has a 4xx status code
func (o *IpamL2vpnTerminationsListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam l2vpn terminations list o k response has a 5xx status code
func (o *IpamL2vpnTerminationsListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam l2vpn terminations list o k response a status code equal to that given
func (o *IpamL2vpnTerminationsListOK) IsCode(code int) bool {
	return code == 200
}

func (o *IpamL2vpnTerminationsListOK) Error() string {
	return fmt.Sprintf("[GET /ipam/l2vpn-terminations/][%d] ipamL2vpnTerminationsListOK  %+v", 200, o.Payload)
}

func (o *IpamL2vpnTerminationsListOK) String() string {
	return fmt.Sprintf("[GET /ipam/l2vpn-terminations/][%d] ipamL2vpnTerminationsListOK  %+v", 200, o.Payload)
}

func (o *IpamL2vpnTerminationsListOK) GetPayload() *IpamL2vpnTerminationsListOKBody {
	return o.Payload
}

func (o *IpamL2vpnTerminationsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(IpamL2vpnTerminationsListOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIpamL2vpnTerminationsListDefault creates a IpamL2vpnTerminationsListDefault with default headers values
func NewIpamL2vpnTerminationsListDefault(code int) *IpamL2vpnTerminationsListDefault {
	return &IpamL2vpnTerminationsListDefault{
		_statusCode: code,
	}
}

/*
IpamL2vpnTerminationsListDefault describes a response with status code -1, with default header values.

IpamL2vpnTerminationsListDefault ipam l2vpn terminations list default
*/
type IpamL2vpnTerminationsListDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the ipam l2vpn terminations list default response
func (o *IpamL2vpnTerminationsListDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this ipam l2vpn terminations list default response has a 2xx status code
func (o *IpamL2vpnTerminationsListDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ipam l2vpn terminations list default response has a 3xx status code
func (o *IpamL2vpnTerminationsListDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ipam l2vpn terminations list default response has a 4xx status code
func (o *IpamL2vpnTerminationsListDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ipam l2vpn terminations list default response has a 5xx status code
func (o *IpamL2vpnTerminationsListDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ipam l2vpn terminations list default response a status code equal to that given
func (o *IpamL2vpnTerminationsListDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *IpamL2vpnTerminationsListDefault) Error() string {
	return fmt.Sprintf("[GET /ipam/l2vpn-terminations/][%d] ipam_l2vpn-terminations_list default  %+v", o._statusCode, o.Payload)
}

func (o *IpamL2vpnTerminationsListDefault) String() string {
	return fmt.Sprintf("[GET /ipam/l2vpn-terminations/][%d] ipam_l2vpn-terminations_list default  %+v", o._statusCode, o.Payload)
}

func (o *IpamL2vpnTerminationsListDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamL2vpnTerminationsListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
IpamL2vpnTerminationsListOKBody ipam l2vpn terminations list o k body
swagger:model IpamL2vpnTerminationsListOKBody
*/
type IpamL2vpnTerminationsListOKBody struct {

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
	Results []*models.L2VPNTermination `json:"results"`
}

// Validate validates this ipam l2vpn terminations list o k body
func (o *IpamL2vpnTerminationsListOKBody) Validate(formats strfmt.Registry) error {
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

func (o *IpamL2vpnTerminationsListOKBody) validateCount(formats strfmt.Registry) error {

	if err := validate.Required("ipamL2vpnTerminationsListOK"+"."+"count", "body", o.Count); err != nil {
		return err
	}

	return nil
}

func (o *IpamL2vpnTerminationsListOKBody) validateNext(formats strfmt.Registry) error {
	if swag.IsZero(o.Next) { // not required
		return nil
	}

	if err := validate.FormatOf("ipamL2vpnTerminationsListOK"+"."+"next", "body", "uri", o.Next.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *IpamL2vpnTerminationsListOKBody) validatePrevious(formats strfmt.Registry) error {
	if swag.IsZero(o.Previous) { // not required
		return nil
	}

	if err := validate.FormatOf("ipamL2vpnTerminationsListOK"+"."+"previous", "body", "uri", o.Previous.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *IpamL2vpnTerminationsListOKBody) validateResults(formats strfmt.Registry) error {

	if err := validate.Required("ipamL2vpnTerminationsListOK"+"."+"results", "body", o.Results); err != nil {
		return err
	}

	for i := 0; i < len(o.Results); i++ {
		if swag.IsZero(o.Results[i]) { // not required
			continue
		}

		if o.Results[i] != nil {
			if err := o.Results[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ipamL2vpnTerminationsListOK" + "." + "results" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ipamL2vpnTerminationsListOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this ipam l2vpn terminations list o k body based on the context it is used
func (o *IpamL2vpnTerminationsListOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateResults(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *IpamL2vpnTerminationsListOKBody) contextValidateResults(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Results); i++ {

		if o.Results[i] != nil {
			if err := o.Results[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ipamL2vpnTerminationsListOK" + "." + "results" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ipamL2vpnTerminationsListOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *IpamL2vpnTerminationsListOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *IpamL2vpnTerminationsListOKBody) UnmarshalBinary(b []byte) error {
	var res IpamL2vpnTerminationsListOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
