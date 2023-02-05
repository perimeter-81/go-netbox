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

// DcimInterfaceTemplatesListReader is a Reader for the DcimInterfaceTemplatesList structure.
type DcimInterfaceTemplatesListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimInterfaceTemplatesListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimInterfaceTemplatesListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimInterfaceTemplatesListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimInterfaceTemplatesListOK creates a DcimInterfaceTemplatesListOK with default headers values
func NewDcimInterfaceTemplatesListOK() *DcimInterfaceTemplatesListOK {
	return &DcimInterfaceTemplatesListOK{}
}

/*
DcimInterfaceTemplatesListOK describes a response with status code 200, with default header values.

DcimInterfaceTemplatesListOK dcim interface templates list o k
*/
type DcimInterfaceTemplatesListOK struct {
	Payload *DcimInterfaceTemplatesListOKBody
}

// IsSuccess returns true when this dcim interface templates list o k response has a 2xx status code
func (o *DcimInterfaceTemplatesListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim interface templates list o k response has a 3xx status code
func (o *DcimInterfaceTemplatesListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim interface templates list o k response has a 4xx status code
func (o *DcimInterfaceTemplatesListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim interface templates list o k response has a 5xx status code
func (o *DcimInterfaceTemplatesListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim interface templates list o k response a status code equal to that given
func (o *DcimInterfaceTemplatesListOK) IsCode(code int) bool {
	return code == 200
}

func (o *DcimInterfaceTemplatesListOK) Error() string {
	return fmt.Sprintf("[GET /dcim/interface-templates/][%d] dcimInterfaceTemplatesListOK  %+v", 200, o.Payload)
}

func (o *DcimInterfaceTemplatesListOK) String() string {
	return fmt.Sprintf("[GET /dcim/interface-templates/][%d] dcimInterfaceTemplatesListOK  %+v", 200, o.Payload)
}

func (o *DcimInterfaceTemplatesListOK) GetPayload() *DcimInterfaceTemplatesListOKBody {
	return o.Payload
}

func (o *DcimInterfaceTemplatesListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(DcimInterfaceTemplatesListOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimInterfaceTemplatesListDefault creates a DcimInterfaceTemplatesListDefault with default headers values
func NewDcimInterfaceTemplatesListDefault(code int) *DcimInterfaceTemplatesListDefault {
	return &DcimInterfaceTemplatesListDefault{
		_statusCode: code,
	}
}

/*
DcimInterfaceTemplatesListDefault describes a response with status code -1, with default header values.

DcimInterfaceTemplatesListDefault dcim interface templates list default
*/
type DcimInterfaceTemplatesListDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim interface templates list default response
func (o *DcimInterfaceTemplatesListDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim interface templates list default response has a 2xx status code
func (o *DcimInterfaceTemplatesListDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim interface templates list default response has a 3xx status code
func (o *DcimInterfaceTemplatesListDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim interface templates list default response has a 4xx status code
func (o *DcimInterfaceTemplatesListDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim interface templates list default response has a 5xx status code
func (o *DcimInterfaceTemplatesListDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim interface templates list default response a status code equal to that given
func (o *DcimInterfaceTemplatesListDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimInterfaceTemplatesListDefault) Error() string {
	return fmt.Sprintf("[GET /dcim/interface-templates/][%d] dcim_interface-templates_list default  %+v", o._statusCode, o.Payload)
}

func (o *DcimInterfaceTemplatesListDefault) String() string {
	return fmt.Sprintf("[GET /dcim/interface-templates/][%d] dcim_interface-templates_list default  %+v", o._statusCode, o.Payload)
}

func (o *DcimInterfaceTemplatesListDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimInterfaceTemplatesListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
DcimInterfaceTemplatesListOKBody dcim interface templates list o k body
swagger:model DcimInterfaceTemplatesListOKBody
*/
type DcimInterfaceTemplatesListOKBody struct {

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
	Results []*models.InterfaceTemplate `json:"results"`
}

// Validate validates this dcim interface templates list o k body
func (o *DcimInterfaceTemplatesListOKBody) Validate(formats strfmt.Registry) error {
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

func (o *DcimInterfaceTemplatesListOKBody) validateCount(formats strfmt.Registry) error {

	if err := validate.Required("dcimInterfaceTemplatesListOK"+"."+"count", "body", o.Count); err != nil {
		return err
	}

	return nil
}

func (o *DcimInterfaceTemplatesListOKBody) validateNext(formats strfmt.Registry) error {
	if swag.IsZero(o.Next) { // not required
		return nil
	}

	if err := validate.FormatOf("dcimInterfaceTemplatesListOK"+"."+"next", "body", "uri", o.Next.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *DcimInterfaceTemplatesListOKBody) validatePrevious(formats strfmt.Registry) error {
	if swag.IsZero(o.Previous) { // not required
		return nil
	}

	if err := validate.FormatOf("dcimInterfaceTemplatesListOK"+"."+"previous", "body", "uri", o.Previous.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *DcimInterfaceTemplatesListOKBody) validateResults(formats strfmt.Registry) error {

	if err := validate.Required("dcimInterfaceTemplatesListOK"+"."+"results", "body", o.Results); err != nil {
		return err
	}

	for i := 0; i < len(o.Results); i++ {
		if swag.IsZero(o.Results[i]) { // not required
			continue
		}

		if o.Results[i] != nil {
			if err := o.Results[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dcimInterfaceTemplatesListOK" + "." + "results" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("dcimInterfaceTemplatesListOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this dcim interface templates list o k body based on the context it is used
func (o *DcimInterfaceTemplatesListOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateResults(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DcimInterfaceTemplatesListOKBody) contextValidateResults(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Results); i++ {

		if o.Results[i] != nil {
			if err := o.Results[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dcimInterfaceTemplatesListOK" + "." + "results" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("dcimInterfaceTemplatesListOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *DcimInterfaceTemplatesListOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DcimInterfaceTemplatesListOKBody) UnmarshalBinary(b []byte) error {
	var res DcimInterfaceTemplatesListOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
