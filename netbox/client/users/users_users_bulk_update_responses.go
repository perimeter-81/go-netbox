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
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/perimeter-81/go-netbox/netbox/models"
)

// UsersUsersBulkUpdateReader is a Reader for the UsersUsersBulkUpdate structure.
type UsersUsersBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UsersUsersBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUsersUsersBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUsersUsersBulkUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUsersUsersBulkUpdateOK creates a UsersUsersBulkUpdateOK with default headers values
func NewUsersUsersBulkUpdateOK() *UsersUsersBulkUpdateOK {
	return &UsersUsersBulkUpdateOK{}
}

/*
UsersUsersBulkUpdateOK describes a response with status code 200, with default header values.

UsersUsersBulkUpdateOK users users bulk update o k
*/
type UsersUsersBulkUpdateOK struct {
	Payload *models.User
}

// IsSuccess returns true when this users users bulk update o k response has a 2xx status code
func (o *UsersUsersBulkUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this users users bulk update o k response has a 3xx status code
func (o *UsersUsersBulkUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users users bulk update o k response has a 4xx status code
func (o *UsersUsersBulkUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this users users bulk update o k response has a 5xx status code
func (o *UsersUsersBulkUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this users users bulk update o k response a status code equal to that given
func (o *UsersUsersBulkUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *UsersUsersBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /users/users/][%d] usersUsersBulkUpdateOK  %+v", 200, o.Payload)
}

func (o *UsersUsersBulkUpdateOK) String() string {
	return fmt.Sprintf("[PUT /users/users/][%d] usersUsersBulkUpdateOK  %+v", 200, o.Payload)
}

func (o *UsersUsersBulkUpdateOK) GetPayload() *models.User {
	return o.Payload
}

func (o *UsersUsersBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.User)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersUsersBulkUpdateDefault creates a UsersUsersBulkUpdateDefault with default headers values
func NewUsersUsersBulkUpdateDefault(code int) *UsersUsersBulkUpdateDefault {
	return &UsersUsersBulkUpdateDefault{
		_statusCode: code,
	}
}

/*
UsersUsersBulkUpdateDefault describes a response with status code -1, with default header values.

UsersUsersBulkUpdateDefault users users bulk update default
*/
type UsersUsersBulkUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the users users bulk update default response
func (o *UsersUsersBulkUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this users users bulk update default response has a 2xx status code
func (o *UsersUsersBulkUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this users users bulk update default response has a 3xx status code
func (o *UsersUsersBulkUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this users users bulk update default response has a 4xx status code
func (o *UsersUsersBulkUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this users users bulk update default response has a 5xx status code
func (o *UsersUsersBulkUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this users users bulk update default response a status code equal to that given
func (o *UsersUsersBulkUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *UsersUsersBulkUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /users/users/][%d] users_users_bulk_update default  %+v", o._statusCode, o.Payload)
}

func (o *UsersUsersBulkUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /users/users/][%d] users_users_bulk_update default  %+v", o._statusCode, o.Payload)
}

func (o *UsersUsersBulkUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *UsersUsersBulkUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
