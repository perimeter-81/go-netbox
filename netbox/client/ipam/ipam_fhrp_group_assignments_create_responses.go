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

// IpamFhrpGroupAssignmentsCreateReader is a Reader for the IpamFhrpGroupAssignmentsCreate structure.
type IpamFhrpGroupAssignmentsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamFhrpGroupAssignmentsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewIpamFhrpGroupAssignmentsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIpamFhrpGroupAssignmentsCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpamFhrpGroupAssignmentsCreateCreated creates a IpamFhrpGroupAssignmentsCreateCreated with default headers values
func NewIpamFhrpGroupAssignmentsCreateCreated() *IpamFhrpGroupAssignmentsCreateCreated {
	return &IpamFhrpGroupAssignmentsCreateCreated{}
}

/*
IpamFhrpGroupAssignmentsCreateCreated describes a response with status code 201, with default header values.

IpamFhrpGroupAssignmentsCreateCreated ipam fhrp group assignments create created
*/
type IpamFhrpGroupAssignmentsCreateCreated struct {
	Payload *models.FHRPGroupAssignment
}

// IsSuccess returns true when this ipam fhrp group assignments create created response has a 2xx status code
func (o *IpamFhrpGroupAssignmentsCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam fhrp group assignments create created response has a 3xx status code
func (o *IpamFhrpGroupAssignmentsCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam fhrp group assignments create created response has a 4xx status code
func (o *IpamFhrpGroupAssignmentsCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam fhrp group assignments create created response has a 5xx status code
func (o *IpamFhrpGroupAssignmentsCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam fhrp group assignments create created response a status code equal to that given
func (o *IpamFhrpGroupAssignmentsCreateCreated) IsCode(code int) bool {
	return code == 201
}

func (o *IpamFhrpGroupAssignmentsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /ipam/fhrp-group-assignments/][%d] ipamFhrpGroupAssignmentsCreateCreated  %+v", 201, o.Payload)
}

func (o *IpamFhrpGroupAssignmentsCreateCreated) String() string {
	return fmt.Sprintf("[POST /ipam/fhrp-group-assignments/][%d] ipamFhrpGroupAssignmentsCreateCreated  %+v", 201, o.Payload)
}

func (o *IpamFhrpGroupAssignmentsCreateCreated) GetPayload() *models.FHRPGroupAssignment {
	return o.Payload
}

func (o *IpamFhrpGroupAssignmentsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FHRPGroupAssignment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIpamFhrpGroupAssignmentsCreateDefault creates a IpamFhrpGroupAssignmentsCreateDefault with default headers values
func NewIpamFhrpGroupAssignmentsCreateDefault(code int) *IpamFhrpGroupAssignmentsCreateDefault {
	return &IpamFhrpGroupAssignmentsCreateDefault{
		_statusCode: code,
	}
}

/*
IpamFhrpGroupAssignmentsCreateDefault describes a response with status code -1, with default header values.

IpamFhrpGroupAssignmentsCreateDefault ipam fhrp group assignments create default
*/
type IpamFhrpGroupAssignmentsCreateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the ipam fhrp group assignments create default response
func (o *IpamFhrpGroupAssignmentsCreateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this ipam fhrp group assignments create default response has a 2xx status code
func (o *IpamFhrpGroupAssignmentsCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ipam fhrp group assignments create default response has a 3xx status code
func (o *IpamFhrpGroupAssignmentsCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ipam fhrp group assignments create default response has a 4xx status code
func (o *IpamFhrpGroupAssignmentsCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ipam fhrp group assignments create default response has a 5xx status code
func (o *IpamFhrpGroupAssignmentsCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ipam fhrp group assignments create default response a status code equal to that given
func (o *IpamFhrpGroupAssignmentsCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *IpamFhrpGroupAssignmentsCreateDefault) Error() string {
	return fmt.Sprintf("[POST /ipam/fhrp-group-assignments/][%d] ipam_fhrp-group-assignments_create default  %+v", o._statusCode, o.Payload)
}

func (o *IpamFhrpGroupAssignmentsCreateDefault) String() string {
	return fmt.Sprintf("[POST /ipam/fhrp-group-assignments/][%d] ipam_fhrp-group-assignments_create default  %+v", o._statusCode, o.Payload)
}

func (o *IpamFhrpGroupAssignmentsCreateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamFhrpGroupAssignmentsCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
