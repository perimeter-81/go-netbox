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

// TenancyContactsBulkUpdateReader is a Reader for the TenancyContactsBulkUpdate structure.
type TenancyContactsBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TenancyContactsBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTenancyContactsBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewTenancyContactsBulkUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewTenancyContactsBulkUpdateOK creates a TenancyContactsBulkUpdateOK with default headers values
func NewTenancyContactsBulkUpdateOK() *TenancyContactsBulkUpdateOK {
	return &TenancyContactsBulkUpdateOK{}
}

/*
TenancyContactsBulkUpdateOK describes a response with status code 200, with default header values.

TenancyContactsBulkUpdateOK tenancy contacts bulk update o k
*/
type TenancyContactsBulkUpdateOK struct {
	Payload *models.Contact
}

// IsSuccess returns true when this tenancy contacts bulk update o k response has a 2xx status code
func (o *TenancyContactsBulkUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this tenancy contacts bulk update o k response has a 3xx status code
func (o *TenancyContactsBulkUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this tenancy contacts bulk update o k response has a 4xx status code
func (o *TenancyContactsBulkUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this tenancy contacts bulk update o k response has a 5xx status code
func (o *TenancyContactsBulkUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this tenancy contacts bulk update o k response a status code equal to that given
func (o *TenancyContactsBulkUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *TenancyContactsBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /tenancy/contacts/][%d] tenancyContactsBulkUpdateOK  %+v", 200, o.Payload)
}

func (o *TenancyContactsBulkUpdateOK) String() string {
	return fmt.Sprintf("[PUT /tenancy/contacts/][%d] tenancyContactsBulkUpdateOK  %+v", 200, o.Payload)
}

func (o *TenancyContactsBulkUpdateOK) GetPayload() *models.Contact {
	return o.Payload
}

func (o *TenancyContactsBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Contact)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTenancyContactsBulkUpdateDefault creates a TenancyContactsBulkUpdateDefault with default headers values
func NewTenancyContactsBulkUpdateDefault(code int) *TenancyContactsBulkUpdateDefault {
	return &TenancyContactsBulkUpdateDefault{
		_statusCode: code,
	}
}

/*
TenancyContactsBulkUpdateDefault describes a response with status code -1, with default header values.

TenancyContactsBulkUpdateDefault tenancy contacts bulk update default
*/
type TenancyContactsBulkUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the tenancy contacts bulk update default response
func (o *TenancyContactsBulkUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this tenancy contacts bulk update default response has a 2xx status code
func (o *TenancyContactsBulkUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this tenancy contacts bulk update default response has a 3xx status code
func (o *TenancyContactsBulkUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this tenancy contacts bulk update default response has a 4xx status code
func (o *TenancyContactsBulkUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this tenancy contacts bulk update default response has a 5xx status code
func (o *TenancyContactsBulkUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this tenancy contacts bulk update default response a status code equal to that given
func (o *TenancyContactsBulkUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *TenancyContactsBulkUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /tenancy/contacts/][%d] tenancy_contacts_bulk_update default  %+v", o._statusCode, o.Payload)
}

func (o *TenancyContactsBulkUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /tenancy/contacts/][%d] tenancy_contacts_bulk_update default  %+v", o._statusCode, o.Payload)
}

func (o *TenancyContactsBulkUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *TenancyContactsBulkUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
