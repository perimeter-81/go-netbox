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

// DcimDeviceBayTemplatesListReader is a Reader for the DcimDeviceBayTemplatesList structure.
type DcimDeviceBayTemplatesListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimDeviceBayTemplatesListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimDeviceBayTemplatesListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimDeviceBayTemplatesListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimDeviceBayTemplatesListOK creates a DcimDeviceBayTemplatesListOK with default headers values
func NewDcimDeviceBayTemplatesListOK() *DcimDeviceBayTemplatesListOK {
	return &DcimDeviceBayTemplatesListOK{}
}

/*
DcimDeviceBayTemplatesListOK describes a response with status code 200, with default header values.

DcimDeviceBayTemplatesListOK dcim device bay templates list o k
*/
type DcimDeviceBayTemplatesListOK struct {
	Payload *DcimDeviceBayTemplatesListOKBody
}

// IsSuccess returns true when this dcim device bay templates list o k response has a 2xx status code
func (o *DcimDeviceBayTemplatesListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim device bay templates list o k response has a 3xx status code
func (o *DcimDeviceBayTemplatesListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim device bay templates list o k response has a 4xx status code
func (o *DcimDeviceBayTemplatesListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim device bay templates list o k response has a 5xx status code
func (o *DcimDeviceBayTemplatesListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim device bay templates list o k response a status code equal to that given
func (o *DcimDeviceBayTemplatesListOK) IsCode(code int) bool {
	return code == 200
}

func (o *DcimDeviceBayTemplatesListOK) Error() string {
	return fmt.Sprintf("[GET /dcim/device-bay-templates/][%d] dcimDeviceBayTemplatesListOK  %+v", 200, o.Payload)
}

func (o *DcimDeviceBayTemplatesListOK) String() string {
	return fmt.Sprintf("[GET /dcim/device-bay-templates/][%d] dcimDeviceBayTemplatesListOK  %+v", 200, o.Payload)
}

func (o *DcimDeviceBayTemplatesListOK) GetPayload() *DcimDeviceBayTemplatesListOKBody {
	return o.Payload
}

func (o *DcimDeviceBayTemplatesListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(DcimDeviceBayTemplatesListOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimDeviceBayTemplatesListDefault creates a DcimDeviceBayTemplatesListDefault with default headers values
func NewDcimDeviceBayTemplatesListDefault(code int) *DcimDeviceBayTemplatesListDefault {
	return &DcimDeviceBayTemplatesListDefault{
		_statusCode: code,
	}
}

/*
DcimDeviceBayTemplatesListDefault describes a response with status code -1, with default header values.

DcimDeviceBayTemplatesListDefault dcim device bay templates list default
*/
type DcimDeviceBayTemplatesListDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim device bay templates list default response
func (o *DcimDeviceBayTemplatesListDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim device bay templates list default response has a 2xx status code
func (o *DcimDeviceBayTemplatesListDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim device bay templates list default response has a 3xx status code
func (o *DcimDeviceBayTemplatesListDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim device bay templates list default response has a 4xx status code
func (o *DcimDeviceBayTemplatesListDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim device bay templates list default response has a 5xx status code
func (o *DcimDeviceBayTemplatesListDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim device bay templates list default response a status code equal to that given
func (o *DcimDeviceBayTemplatesListDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimDeviceBayTemplatesListDefault) Error() string {
	return fmt.Sprintf("[GET /dcim/device-bay-templates/][%d] dcim_device-bay-templates_list default  %+v", o._statusCode, o.Payload)
}

func (o *DcimDeviceBayTemplatesListDefault) String() string {
	return fmt.Sprintf("[GET /dcim/device-bay-templates/][%d] dcim_device-bay-templates_list default  %+v", o._statusCode, o.Payload)
}

func (o *DcimDeviceBayTemplatesListDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimDeviceBayTemplatesListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
DcimDeviceBayTemplatesListOKBody dcim device bay templates list o k body
swagger:model DcimDeviceBayTemplatesListOKBody
*/
type DcimDeviceBayTemplatesListOKBody struct {

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
	Results []*models.DeviceBayTemplate `json:"results"`
}

// Validate validates this dcim device bay templates list o k body
func (o *DcimDeviceBayTemplatesListOKBody) Validate(formats strfmt.Registry) error {
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

func (o *DcimDeviceBayTemplatesListOKBody) validateCount(formats strfmt.Registry) error {

	if err := validate.Required("dcimDeviceBayTemplatesListOK"+"."+"count", "body", o.Count); err != nil {
		return err
	}

	return nil
}

func (o *DcimDeviceBayTemplatesListOKBody) validateNext(formats strfmt.Registry) error {
	if swag.IsZero(o.Next) { // not required
		return nil
	}

	if err := validate.FormatOf("dcimDeviceBayTemplatesListOK"+"."+"next", "body", "uri", o.Next.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *DcimDeviceBayTemplatesListOKBody) validatePrevious(formats strfmt.Registry) error {
	if swag.IsZero(o.Previous) { // not required
		return nil
	}

	if err := validate.FormatOf("dcimDeviceBayTemplatesListOK"+"."+"previous", "body", "uri", o.Previous.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *DcimDeviceBayTemplatesListOKBody) validateResults(formats strfmt.Registry) error {

	if err := validate.Required("dcimDeviceBayTemplatesListOK"+"."+"results", "body", o.Results); err != nil {
		return err
	}

	for i := 0; i < len(o.Results); i++ {
		if swag.IsZero(o.Results[i]) { // not required
			continue
		}

		if o.Results[i] != nil {
			if err := o.Results[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dcimDeviceBayTemplatesListOK" + "." + "results" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("dcimDeviceBayTemplatesListOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this dcim device bay templates list o k body based on the context it is used
func (o *DcimDeviceBayTemplatesListOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateResults(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DcimDeviceBayTemplatesListOKBody) contextValidateResults(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Results); i++ {

		if o.Results[i] != nil {
			if err := o.Results[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dcimDeviceBayTemplatesListOK" + "." + "results" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("dcimDeviceBayTemplatesListOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *DcimDeviceBayTemplatesListOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DcimDeviceBayTemplatesListOKBody) UnmarshalBinary(b []byte) error {
	var res DcimDeviceBayTemplatesListOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
