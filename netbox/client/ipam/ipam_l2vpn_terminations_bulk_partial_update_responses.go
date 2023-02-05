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

// IpamL2vpnTerminationsBulkPartialUpdateReader is a Reader for the IpamL2vpnTerminationsBulkPartialUpdate structure.
type IpamL2vpnTerminationsBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamL2vpnTerminationsBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamL2vpnTerminationsBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIpamL2vpnTerminationsBulkPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpamL2vpnTerminationsBulkPartialUpdateOK creates a IpamL2vpnTerminationsBulkPartialUpdateOK with default headers values
func NewIpamL2vpnTerminationsBulkPartialUpdateOK() *IpamL2vpnTerminationsBulkPartialUpdateOK {
	return &IpamL2vpnTerminationsBulkPartialUpdateOK{}
}

/*
IpamL2vpnTerminationsBulkPartialUpdateOK describes a response with status code 200, with default header values.

IpamL2vpnTerminationsBulkPartialUpdateOK ipam l2vpn terminations bulk partial update o k
*/
type IpamL2vpnTerminationsBulkPartialUpdateOK struct {
	Payload *models.L2VPNTermination
}

// IsSuccess returns true when this ipam l2vpn terminations bulk partial update o k response has a 2xx status code
func (o *IpamL2vpnTerminationsBulkPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam l2vpn terminations bulk partial update o k response has a 3xx status code
func (o *IpamL2vpnTerminationsBulkPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam l2vpn terminations bulk partial update o k response has a 4xx status code
func (o *IpamL2vpnTerminationsBulkPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam l2vpn terminations bulk partial update o k response has a 5xx status code
func (o *IpamL2vpnTerminationsBulkPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam l2vpn terminations bulk partial update o k response a status code equal to that given
func (o *IpamL2vpnTerminationsBulkPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *IpamL2vpnTerminationsBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /ipam/l2vpn-terminations/][%d] ipamL2vpnTerminationsBulkPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *IpamL2vpnTerminationsBulkPartialUpdateOK) String() string {
	return fmt.Sprintf("[PATCH /ipam/l2vpn-terminations/][%d] ipamL2vpnTerminationsBulkPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *IpamL2vpnTerminationsBulkPartialUpdateOK) GetPayload() *models.L2VPNTermination {
	return o.Payload
}

func (o *IpamL2vpnTerminationsBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.L2VPNTermination)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIpamL2vpnTerminationsBulkPartialUpdateDefault creates a IpamL2vpnTerminationsBulkPartialUpdateDefault with default headers values
func NewIpamL2vpnTerminationsBulkPartialUpdateDefault(code int) *IpamL2vpnTerminationsBulkPartialUpdateDefault {
	return &IpamL2vpnTerminationsBulkPartialUpdateDefault{
		_statusCode: code,
	}
}

/*
IpamL2vpnTerminationsBulkPartialUpdateDefault describes a response with status code -1, with default header values.

IpamL2vpnTerminationsBulkPartialUpdateDefault ipam l2vpn terminations bulk partial update default
*/
type IpamL2vpnTerminationsBulkPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the ipam l2vpn terminations bulk partial update default response
func (o *IpamL2vpnTerminationsBulkPartialUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this ipam l2vpn terminations bulk partial update default response has a 2xx status code
func (o *IpamL2vpnTerminationsBulkPartialUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ipam l2vpn terminations bulk partial update default response has a 3xx status code
func (o *IpamL2vpnTerminationsBulkPartialUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ipam l2vpn terminations bulk partial update default response has a 4xx status code
func (o *IpamL2vpnTerminationsBulkPartialUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ipam l2vpn terminations bulk partial update default response has a 5xx status code
func (o *IpamL2vpnTerminationsBulkPartialUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ipam l2vpn terminations bulk partial update default response a status code equal to that given
func (o *IpamL2vpnTerminationsBulkPartialUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *IpamL2vpnTerminationsBulkPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /ipam/l2vpn-terminations/][%d] ipam_l2vpn-terminations_bulk_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *IpamL2vpnTerminationsBulkPartialUpdateDefault) String() string {
	return fmt.Sprintf("[PATCH /ipam/l2vpn-terminations/][%d] ipam_l2vpn-terminations_bulk_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *IpamL2vpnTerminationsBulkPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamL2vpnTerminationsBulkPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
