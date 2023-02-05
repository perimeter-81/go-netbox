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

// IpamVlanGroupsCreateReader is a Reader for the IpamVlanGroupsCreate structure.
type IpamVlanGroupsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamVlanGroupsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewIpamVlanGroupsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIpamVlanGroupsCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpamVlanGroupsCreateCreated creates a IpamVlanGroupsCreateCreated with default headers values
func NewIpamVlanGroupsCreateCreated() *IpamVlanGroupsCreateCreated {
	return &IpamVlanGroupsCreateCreated{}
}

/*
IpamVlanGroupsCreateCreated describes a response with status code 201, with default header values.

IpamVlanGroupsCreateCreated ipam vlan groups create created
*/
type IpamVlanGroupsCreateCreated struct {
	Payload *models.VLANGroup
}

// IsSuccess returns true when this ipam vlan groups create created response has a 2xx status code
func (o *IpamVlanGroupsCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam vlan groups create created response has a 3xx status code
func (o *IpamVlanGroupsCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam vlan groups create created response has a 4xx status code
func (o *IpamVlanGroupsCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam vlan groups create created response has a 5xx status code
func (o *IpamVlanGroupsCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam vlan groups create created response a status code equal to that given
func (o *IpamVlanGroupsCreateCreated) IsCode(code int) bool {
	return code == 201
}

func (o *IpamVlanGroupsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /ipam/vlan-groups/][%d] ipamVlanGroupsCreateCreated  %+v", 201, o.Payload)
}

func (o *IpamVlanGroupsCreateCreated) String() string {
	return fmt.Sprintf("[POST /ipam/vlan-groups/][%d] ipamVlanGroupsCreateCreated  %+v", 201, o.Payload)
}

func (o *IpamVlanGroupsCreateCreated) GetPayload() *models.VLANGroup {
	return o.Payload
}

func (o *IpamVlanGroupsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VLANGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIpamVlanGroupsCreateDefault creates a IpamVlanGroupsCreateDefault with default headers values
func NewIpamVlanGroupsCreateDefault(code int) *IpamVlanGroupsCreateDefault {
	return &IpamVlanGroupsCreateDefault{
		_statusCode: code,
	}
}

/*
IpamVlanGroupsCreateDefault describes a response with status code -1, with default header values.

IpamVlanGroupsCreateDefault ipam vlan groups create default
*/
type IpamVlanGroupsCreateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the ipam vlan groups create default response
func (o *IpamVlanGroupsCreateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this ipam vlan groups create default response has a 2xx status code
func (o *IpamVlanGroupsCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ipam vlan groups create default response has a 3xx status code
func (o *IpamVlanGroupsCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ipam vlan groups create default response has a 4xx status code
func (o *IpamVlanGroupsCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ipam vlan groups create default response has a 5xx status code
func (o *IpamVlanGroupsCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ipam vlan groups create default response a status code equal to that given
func (o *IpamVlanGroupsCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *IpamVlanGroupsCreateDefault) Error() string {
	return fmt.Sprintf("[POST /ipam/vlan-groups/][%d] ipam_vlan-groups_create default  %+v", o._statusCode, o.Payload)
}

func (o *IpamVlanGroupsCreateDefault) String() string {
	return fmt.Sprintf("[POST /ipam/vlan-groups/][%d] ipam_vlan-groups_create default  %+v", o._statusCode, o.Payload)
}

func (o *IpamVlanGroupsCreateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamVlanGroupsCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
