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
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/perimeter-81/go-netbox/netbox/models"
)

// IpamVlanGroupsUpdateReader is a Reader for the IpamVlanGroupsUpdate structure.
type IpamVlanGroupsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamVlanGroupsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamVlanGroupsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIpamVlanGroupsUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpamVlanGroupsUpdateOK creates a IpamVlanGroupsUpdateOK with default headers values
func NewIpamVlanGroupsUpdateOK() *IpamVlanGroupsUpdateOK {
	return &IpamVlanGroupsUpdateOK{}
}

/*
IpamVlanGroupsUpdateOK describes a response with status code 200, with default header values.

IpamVlanGroupsUpdateOK ipam vlan groups update o k
*/
type IpamVlanGroupsUpdateOK struct {
	Payload *models.VLANGroup
}

// IsSuccess returns true when this ipam vlan groups update o k response has a 2xx status code
func (o *IpamVlanGroupsUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam vlan groups update o k response has a 3xx status code
func (o *IpamVlanGroupsUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam vlan groups update o k response has a 4xx status code
func (o *IpamVlanGroupsUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam vlan groups update o k response has a 5xx status code
func (o *IpamVlanGroupsUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam vlan groups update o k response a status code equal to that given
func (o *IpamVlanGroupsUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *IpamVlanGroupsUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /ipam/vlan-groups/{id}/][%d] ipamVlanGroupsUpdateOK  %+v", 200, o.Payload)
}

func (o *IpamVlanGroupsUpdateOK) String() string {
	return fmt.Sprintf("[PUT /ipam/vlan-groups/{id}/][%d] ipamVlanGroupsUpdateOK  %+v", 200, o.Payload)
}

func (o *IpamVlanGroupsUpdateOK) GetPayload() *models.VLANGroup {
	return o.Payload
}

func (o *IpamVlanGroupsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VLANGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIpamVlanGroupsUpdateDefault creates a IpamVlanGroupsUpdateDefault with default headers values
func NewIpamVlanGroupsUpdateDefault(code int) *IpamVlanGroupsUpdateDefault {
	return &IpamVlanGroupsUpdateDefault{
		_statusCode: code,
	}
}

/*
IpamVlanGroupsUpdateDefault describes a response with status code -1, with default header values.

IpamVlanGroupsUpdateDefault ipam vlan groups update default
*/
type IpamVlanGroupsUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the ipam vlan groups update default response
func (o *IpamVlanGroupsUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this ipam vlan groups update default response has a 2xx status code
func (o *IpamVlanGroupsUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ipam vlan groups update default response has a 3xx status code
func (o *IpamVlanGroupsUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ipam vlan groups update default response has a 4xx status code
func (o *IpamVlanGroupsUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ipam vlan groups update default response has a 5xx status code
func (o *IpamVlanGroupsUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ipam vlan groups update default response a status code equal to that given
func (o *IpamVlanGroupsUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *IpamVlanGroupsUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /ipam/vlan-groups/{id}/][%d] ipam_vlan-groups_update default  %+v", o._statusCode, o.Payload)
}

func (o *IpamVlanGroupsUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /ipam/vlan-groups/{id}/][%d] ipam_vlan-groups_update default  %+v", o._statusCode, o.Payload)
}

func (o *IpamVlanGroupsUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamVlanGroupsUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
