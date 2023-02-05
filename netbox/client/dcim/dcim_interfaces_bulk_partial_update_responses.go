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
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/perimeter-81/go-netbox/netbox/models"
)

// DcimInterfacesBulkPartialUpdateReader is a Reader for the DcimInterfacesBulkPartialUpdate structure.
type DcimInterfacesBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimInterfacesBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimInterfacesBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimInterfacesBulkPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimInterfacesBulkPartialUpdateOK creates a DcimInterfacesBulkPartialUpdateOK with default headers values
func NewDcimInterfacesBulkPartialUpdateOK() *DcimInterfacesBulkPartialUpdateOK {
	return &DcimInterfacesBulkPartialUpdateOK{}
}

/*
DcimInterfacesBulkPartialUpdateOK describes a response with status code 200, with default header values.

DcimInterfacesBulkPartialUpdateOK dcim interfaces bulk partial update o k
*/
type DcimInterfacesBulkPartialUpdateOK struct {
	Payload *models.Interface
}

// IsSuccess returns true when this dcim interfaces bulk partial update o k response has a 2xx status code
func (o *DcimInterfacesBulkPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim interfaces bulk partial update o k response has a 3xx status code
func (o *DcimInterfacesBulkPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim interfaces bulk partial update o k response has a 4xx status code
func (o *DcimInterfacesBulkPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim interfaces bulk partial update o k response has a 5xx status code
func (o *DcimInterfacesBulkPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim interfaces bulk partial update o k response a status code equal to that given
func (o *DcimInterfacesBulkPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *DcimInterfacesBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/interfaces/][%d] dcimInterfacesBulkPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimInterfacesBulkPartialUpdateOK) String() string {
	return fmt.Sprintf("[PATCH /dcim/interfaces/][%d] dcimInterfacesBulkPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimInterfacesBulkPartialUpdateOK) GetPayload() *models.Interface {
	return o.Payload
}

func (o *DcimInterfacesBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Interface)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimInterfacesBulkPartialUpdateDefault creates a DcimInterfacesBulkPartialUpdateDefault with default headers values
func NewDcimInterfacesBulkPartialUpdateDefault(code int) *DcimInterfacesBulkPartialUpdateDefault {
	return &DcimInterfacesBulkPartialUpdateDefault{
		_statusCode: code,
	}
}

/*
DcimInterfacesBulkPartialUpdateDefault describes a response with status code -1, with default header values.

DcimInterfacesBulkPartialUpdateDefault dcim interfaces bulk partial update default
*/
type DcimInterfacesBulkPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim interfaces bulk partial update default response
func (o *DcimInterfacesBulkPartialUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim interfaces bulk partial update default response has a 2xx status code
func (o *DcimInterfacesBulkPartialUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim interfaces bulk partial update default response has a 3xx status code
func (o *DcimInterfacesBulkPartialUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim interfaces bulk partial update default response has a 4xx status code
func (o *DcimInterfacesBulkPartialUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim interfaces bulk partial update default response has a 5xx status code
func (o *DcimInterfacesBulkPartialUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim interfaces bulk partial update default response a status code equal to that given
func (o *DcimInterfacesBulkPartialUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimInterfacesBulkPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /dcim/interfaces/][%d] dcim_interfaces_bulk_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimInterfacesBulkPartialUpdateDefault) String() string {
	return fmt.Sprintf("[PATCH /dcim/interfaces/][%d] dcim_interfaces_bulk_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimInterfacesBulkPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimInterfacesBulkPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
