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
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/perimeter-81/go-netbox/netbox/models"
)

// TenancyContactRolesBulkPartialUpdateReader is a Reader for the TenancyContactRolesBulkPartialUpdate structure.
type TenancyContactRolesBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TenancyContactRolesBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTenancyContactRolesBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewTenancyContactRolesBulkPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewTenancyContactRolesBulkPartialUpdateOK creates a TenancyContactRolesBulkPartialUpdateOK with default headers values
func NewTenancyContactRolesBulkPartialUpdateOK() *TenancyContactRolesBulkPartialUpdateOK {
	return &TenancyContactRolesBulkPartialUpdateOK{}
}

/*
TenancyContactRolesBulkPartialUpdateOK describes a response with status code 200, with default header values.

TenancyContactRolesBulkPartialUpdateOK tenancy contact roles bulk partial update o k
*/
type TenancyContactRolesBulkPartialUpdateOK struct {
	Payload *models.ContactRole
}

// IsSuccess returns true when this tenancy contact roles bulk partial update o k response has a 2xx status code
func (o *TenancyContactRolesBulkPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this tenancy contact roles bulk partial update o k response has a 3xx status code
func (o *TenancyContactRolesBulkPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this tenancy contact roles bulk partial update o k response has a 4xx status code
func (o *TenancyContactRolesBulkPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this tenancy contact roles bulk partial update o k response has a 5xx status code
func (o *TenancyContactRolesBulkPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this tenancy contact roles bulk partial update o k response a status code equal to that given
func (o *TenancyContactRolesBulkPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *TenancyContactRolesBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /tenancy/contact-roles/][%d] tenancyContactRolesBulkPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *TenancyContactRolesBulkPartialUpdateOK) String() string {
	return fmt.Sprintf("[PATCH /tenancy/contact-roles/][%d] tenancyContactRolesBulkPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *TenancyContactRolesBulkPartialUpdateOK) GetPayload() *models.ContactRole {
	return o.Payload
}

func (o *TenancyContactRolesBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ContactRole)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTenancyContactRolesBulkPartialUpdateDefault creates a TenancyContactRolesBulkPartialUpdateDefault with default headers values
func NewTenancyContactRolesBulkPartialUpdateDefault(code int) *TenancyContactRolesBulkPartialUpdateDefault {
	return &TenancyContactRolesBulkPartialUpdateDefault{
		_statusCode: code,
	}
}

/*
TenancyContactRolesBulkPartialUpdateDefault describes a response with status code -1, with default header values.

TenancyContactRolesBulkPartialUpdateDefault tenancy contact roles bulk partial update default
*/
type TenancyContactRolesBulkPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the tenancy contact roles bulk partial update default response
func (o *TenancyContactRolesBulkPartialUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this tenancy contact roles bulk partial update default response has a 2xx status code
func (o *TenancyContactRolesBulkPartialUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this tenancy contact roles bulk partial update default response has a 3xx status code
func (o *TenancyContactRolesBulkPartialUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this tenancy contact roles bulk partial update default response has a 4xx status code
func (o *TenancyContactRolesBulkPartialUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this tenancy contact roles bulk partial update default response has a 5xx status code
func (o *TenancyContactRolesBulkPartialUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this tenancy contact roles bulk partial update default response a status code equal to that given
func (o *TenancyContactRolesBulkPartialUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *TenancyContactRolesBulkPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /tenancy/contact-roles/][%d] tenancy_contact-roles_bulk_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *TenancyContactRolesBulkPartialUpdateDefault) String() string {
	return fmt.Sprintf("[PATCH /tenancy/contact-roles/][%d] tenancy_contact-roles_bulk_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *TenancyContactRolesBulkPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *TenancyContactRolesBulkPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
