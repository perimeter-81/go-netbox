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

// IpamL2vpnTerminationsReadReader is a Reader for the IpamL2vpnTerminationsRead structure.
type IpamL2vpnTerminationsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamL2vpnTerminationsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamL2vpnTerminationsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIpamL2vpnTerminationsReadDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpamL2vpnTerminationsReadOK creates a IpamL2vpnTerminationsReadOK with default headers values
func NewIpamL2vpnTerminationsReadOK() *IpamL2vpnTerminationsReadOK {
	return &IpamL2vpnTerminationsReadOK{}
}

/*
IpamL2vpnTerminationsReadOK describes a response with status code 200, with default header values.

IpamL2vpnTerminationsReadOK ipam l2vpn terminations read o k
*/
type IpamL2vpnTerminationsReadOK struct {
	Payload *models.L2VPNTermination
}

// IsSuccess returns true when this ipam l2vpn terminations read o k response has a 2xx status code
func (o *IpamL2vpnTerminationsReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam l2vpn terminations read o k response has a 3xx status code
func (o *IpamL2vpnTerminationsReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam l2vpn terminations read o k response has a 4xx status code
func (o *IpamL2vpnTerminationsReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam l2vpn terminations read o k response has a 5xx status code
func (o *IpamL2vpnTerminationsReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam l2vpn terminations read o k response a status code equal to that given
func (o *IpamL2vpnTerminationsReadOK) IsCode(code int) bool {
	return code == 200
}

func (o *IpamL2vpnTerminationsReadOK) Error() string {
	return fmt.Sprintf("[GET /ipam/l2vpn-terminations/{id}/][%d] ipamL2vpnTerminationsReadOK  %+v", 200, o.Payload)
}

func (o *IpamL2vpnTerminationsReadOK) String() string {
	return fmt.Sprintf("[GET /ipam/l2vpn-terminations/{id}/][%d] ipamL2vpnTerminationsReadOK  %+v", 200, o.Payload)
}

func (o *IpamL2vpnTerminationsReadOK) GetPayload() *models.L2VPNTermination {
	return o.Payload
}

func (o *IpamL2vpnTerminationsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.L2VPNTermination)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIpamL2vpnTerminationsReadDefault creates a IpamL2vpnTerminationsReadDefault with default headers values
func NewIpamL2vpnTerminationsReadDefault(code int) *IpamL2vpnTerminationsReadDefault {
	return &IpamL2vpnTerminationsReadDefault{
		_statusCode: code,
	}
}

/*
IpamL2vpnTerminationsReadDefault describes a response with status code -1, with default header values.

IpamL2vpnTerminationsReadDefault ipam l2vpn terminations read default
*/
type IpamL2vpnTerminationsReadDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the ipam l2vpn terminations read default response
func (o *IpamL2vpnTerminationsReadDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this ipam l2vpn terminations read default response has a 2xx status code
func (o *IpamL2vpnTerminationsReadDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ipam l2vpn terminations read default response has a 3xx status code
func (o *IpamL2vpnTerminationsReadDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ipam l2vpn terminations read default response has a 4xx status code
func (o *IpamL2vpnTerminationsReadDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ipam l2vpn terminations read default response has a 5xx status code
func (o *IpamL2vpnTerminationsReadDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ipam l2vpn terminations read default response a status code equal to that given
func (o *IpamL2vpnTerminationsReadDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *IpamL2vpnTerminationsReadDefault) Error() string {
	return fmt.Sprintf("[GET /ipam/l2vpn-terminations/{id}/][%d] ipam_l2vpn-terminations_read default  %+v", o._statusCode, o.Payload)
}

func (o *IpamL2vpnTerminationsReadDefault) String() string {
	return fmt.Sprintf("[GET /ipam/l2vpn-terminations/{id}/][%d] ipam_l2vpn-terminations_read default  %+v", o._statusCode, o.Payload)
}

func (o *IpamL2vpnTerminationsReadDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamL2vpnTerminationsReadDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
