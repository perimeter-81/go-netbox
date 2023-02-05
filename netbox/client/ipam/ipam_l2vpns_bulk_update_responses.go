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

// IpamL2vpnsBulkUpdateReader is a Reader for the IpamL2vpnsBulkUpdate structure.
type IpamL2vpnsBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamL2vpnsBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamL2vpnsBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIpamL2vpnsBulkUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpamL2vpnsBulkUpdateOK creates a IpamL2vpnsBulkUpdateOK with default headers values
func NewIpamL2vpnsBulkUpdateOK() *IpamL2vpnsBulkUpdateOK {
	return &IpamL2vpnsBulkUpdateOK{}
}

/*
IpamL2vpnsBulkUpdateOK describes a response with status code 200, with default header values.

IpamL2vpnsBulkUpdateOK ipam l2vpns bulk update o k
*/
type IpamL2vpnsBulkUpdateOK struct {
	Payload *models.L2VPN
}

// IsSuccess returns true when this ipam l2vpns bulk update o k response has a 2xx status code
func (o *IpamL2vpnsBulkUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam l2vpns bulk update o k response has a 3xx status code
func (o *IpamL2vpnsBulkUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam l2vpns bulk update o k response has a 4xx status code
func (o *IpamL2vpnsBulkUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam l2vpns bulk update o k response has a 5xx status code
func (o *IpamL2vpnsBulkUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam l2vpns bulk update o k response a status code equal to that given
func (o *IpamL2vpnsBulkUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *IpamL2vpnsBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /ipam/l2vpns/][%d] ipamL2vpnsBulkUpdateOK  %+v", 200, o.Payload)
}

func (o *IpamL2vpnsBulkUpdateOK) String() string {
	return fmt.Sprintf("[PUT /ipam/l2vpns/][%d] ipamL2vpnsBulkUpdateOK  %+v", 200, o.Payload)
}

func (o *IpamL2vpnsBulkUpdateOK) GetPayload() *models.L2VPN {
	return o.Payload
}

func (o *IpamL2vpnsBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.L2VPN)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIpamL2vpnsBulkUpdateDefault creates a IpamL2vpnsBulkUpdateDefault with default headers values
func NewIpamL2vpnsBulkUpdateDefault(code int) *IpamL2vpnsBulkUpdateDefault {
	return &IpamL2vpnsBulkUpdateDefault{
		_statusCode: code,
	}
}

/*
IpamL2vpnsBulkUpdateDefault describes a response with status code -1, with default header values.

IpamL2vpnsBulkUpdateDefault ipam l2vpns bulk update default
*/
type IpamL2vpnsBulkUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the ipam l2vpns bulk update default response
func (o *IpamL2vpnsBulkUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this ipam l2vpns bulk update default response has a 2xx status code
func (o *IpamL2vpnsBulkUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ipam l2vpns bulk update default response has a 3xx status code
func (o *IpamL2vpnsBulkUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ipam l2vpns bulk update default response has a 4xx status code
func (o *IpamL2vpnsBulkUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ipam l2vpns bulk update default response has a 5xx status code
func (o *IpamL2vpnsBulkUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ipam l2vpns bulk update default response a status code equal to that given
func (o *IpamL2vpnsBulkUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *IpamL2vpnsBulkUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /ipam/l2vpns/][%d] ipam_l2vpns_bulk_update default  %+v", o._statusCode, o.Payload)
}

func (o *IpamL2vpnsBulkUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /ipam/l2vpns/][%d] ipam_l2vpns_bulk_update default  %+v", o._statusCode, o.Payload)
}

func (o *IpamL2vpnsBulkUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamL2vpnsBulkUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
