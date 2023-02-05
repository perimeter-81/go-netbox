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

package users

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

// UsersUsersListReader is a Reader for the UsersUsersList structure.
type UsersUsersListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UsersUsersListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUsersUsersListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUsersUsersListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUsersUsersListOK creates a UsersUsersListOK with default headers values
func NewUsersUsersListOK() *UsersUsersListOK {
	return &UsersUsersListOK{}
}

/*
UsersUsersListOK describes a response with status code 200, with default header values.

UsersUsersListOK users users list o k
*/
type UsersUsersListOK struct {
	Payload *UsersUsersListOKBody
}

// IsSuccess returns true when this users users list o k response has a 2xx status code
func (o *UsersUsersListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this users users list o k response has a 3xx status code
func (o *UsersUsersListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users users list o k response has a 4xx status code
func (o *UsersUsersListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this users users list o k response has a 5xx status code
func (o *UsersUsersListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this users users list o k response a status code equal to that given
func (o *UsersUsersListOK) IsCode(code int) bool {
	return code == 200
}

func (o *UsersUsersListOK) Error() string {
	return fmt.Sprintf("[GET /users/users/][%d] usersUsersListOK  %+v", 200, o.Payload)
}

func (o *UsersUsersListOK) String() string {
	return fmt.Sprintf("[GET /users/users/][%d] usersUsersListOK  %+v", 200, o.Payload)
}

func (o *UsersUsersListOK) GetPayload() *UsersUsersListOKBody {
	return o.Payload
}

func (o *UsersUsersListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(UsersUsersListOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersUsersListDefault creates a UsersUsersListDefault with default headers values
func NewUsersUsersListDefault(code int) *UsersUsersListDefault {
	return &UsersUsersListDefault{
		_statusCode: code,
	}
}

/*
UsersUsersListDefault describes a response with status code -1, with default header values.

UsersUsersListDefault users users list default
*/
type UsersUsersListDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the users users list default response
func (o *UsersUsersListDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this users users list default response has a 2xx status code
func (o *UsersUsersListDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this users users list default response has a 3xx status code
func (o *UsersUsersListDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this users users list default response has a 4xx status code
func (o *UsersUsersListDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this users users list default response has a 5xx status code
func (o *UsersUsersListDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this users users list default response a status code equal to that given
func (o *UsersUsersListDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *UsersUsersListDefault) Error() string {
	return fmt.Sprintf("[GET /users/users/][%d] users_users_list default  %+v", o._statusCode, o.Payload)
}

func (o *UsersUsersListDefault) String() string {
	return fmt.Sprintf("[GET /users/users/][%d] users_users_list default  %+v", o._statusCode, o.Payload)
}

func (o *UsersUsersListDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *UsersUsersListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
UsersUsersListOKBody users users list o k body
swagger:model UsersUsersListOKBody
*/
type UsersUsersListOKBody struct {

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
	Results []*models.User `json:"results"`
}

// Validate validates this users users list o k body
func (o *UsersUsersListOKBody) Validate(formats strfmt.Registry) error {
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

func (o *UsersUsersListOKBody) validateCount(formats strfmt.Registry) error {

	if err := validate.Required("usersUsersListOK"+"."+"count", "body", o.Count); err != nil {
		return err
	}

	return nil
}

func (o *UsersUsersListOKBody) validateNext(formats strfmt.Registry) error {
	if swag.IsZero(o.Next) { // not required
		return nil
	}

	if err := validate.FormatOf("usersUsersListOK"+"."+"next", "body", "uri", o.Next.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *UsersUsersListOKBody) validatePrevious(formats strfmt.Registry) error {
	if swag.IsZero(o.Previous) { // not required
		return nil
	}

	if err := validate.FormatOf("usersUsersListOK"+"."+"previous", "body", "uri", o.Previous.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *UsersUsersListOKBody) validateResults(formats strfmt.Registry) error {

	if err := validate.Required("usersUsersListOK"+"."+"results", "body", o.Results); err != nil {
		return err
	}

	for i := 0; i < len(o.Results); i++ {
		if swag.IsZero(o.Results[i]) { // not required
			continue
		}

		if o.Results[i] != nil {
			if err := o.Results[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("usersUsersListOK" + "." + "results" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("usersUsersListOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this users users list o k body based on the context it is used
func (o *UsersUsersListOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateResults(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UsersUsersListOKBody) contextValidateResults(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Results); i++ {

		if o.Results[i] != nil {
			if err := o.Results[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("usersUsersListOK" + "." + "results" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("usersUsersListOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *UsersUsersListOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UsersUsersListOKBody) UnmarshalBinary(b []byte) error {
	var res UsersUsersListOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
