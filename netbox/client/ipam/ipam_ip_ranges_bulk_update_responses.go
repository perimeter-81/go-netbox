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

// IpamIPRangesBulkUpdateReader is a Reader for the IpamIPRangesBulkUpdate structure.
type IpamIPRangesBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamIPRangesBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamIPRangesBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIpamIPRangesBulkUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpamIPRangesBulkUpdateOK creates a IpamIPRangesBulkUpdateOK with default headers values
func NewIpamIPRangesBulkUpdateOK() *IpamIPRangesBulkUpdateOK {
	return &IpamIPRangesBulkUpdateOK{}
}

/*
IpamIPRangesBulkUpdateOK describes a response with status code 200, with default header values.

IpamIPRangesBulkUpdateOK ipam Ip ranges bulk update o k
*/
type IpamIPRangesBulkUpdateOK struct {
	Payload *models.IPRange
}

// IsSuccess returns true when this ipam Ip ranges bulk update o k response has a 2xx status code
func (o *IpamIPRangesBulkUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam Ip ranges bulk update o k response has a 3xx status code
func (o *IpamIPRangesBulkUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam Ip ranges bulk update o k response has a 4xx status code
func (o *IpamIPRangesBulkUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam Ip ranges bulk update o k response has a 5xx status code
func (o *IpamIPRangesBulkUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam Ip ranges bulk update o k response a status code equal to that given
func (o *IpamIPRangesBulkUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *IpamIPRangesBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /ipam/ip-ranges/][%d] ipamIpRangesBulkUpdateOK  %+v", 200, o.Payload)
}

func (o *IpamIPRangesBulkUpdateOK) String() string {
	return fmt.Sprintf("[PUT /ipam/ip-ranges/][%d] ipamIpRangesBulkUpdateOK  %+v", 200, o.Payload)
}

func (o *IpamIPRangesBulkUpdateOK) GetPayload() *models.IPRange {
	return o.Payload
}

func (o *IpamIPRangesBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IPRange)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIpamIPRangesBulkUpdateDefault creates a IpamIPRangesBulkUpdateDefault with default headers values
func NewIpamIPRangesBulkUpdateDefault(code int) *IpamIPRangesBulkUpdateDefault {
	return &IpamIPRangesBulkUpdateDefault{
		_statusCode: code,
	}
}

/*
IpamIPRangesBulkUpdateDefault describes a response with status code -1, with default header values.

IpamIPRangesBulkUpdateDefault ipam ip ranges bulk update default
*/
type IpamIPRangesBulkUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the ipam ip ranges bulk update default response
func (o *IpamIPRangesBulkUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this ipam ip ranges bulk update default response has a 2xx status code
func (o *IpamIPRangesBulkUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ipam ip ranges bulk update default response has a 3xx status code
func (o *IpamIPRangesBulkUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ipam ip ranges bulk update default response has a 4xx status code
func (o *IpamIPRangesBulkUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ipam ip ranges bulk update default response has a 5xx status code
func (o *IpamIPRangesBulkUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ipam ip ranges bulk update default response a status code equal to that given
func (o *IpamIPRangesBulkUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *IpamIPRangesBulkUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /ipam/ip-ranges/][%d] ipam_ip-ranges_bulk_update default  %+v", o._statusCode, o.Payload)
}

func (o *IpamIPRangesBulkUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /ipam/ip-ranges/][%d] ipam_ip-ranges_bulk_update default  %+v", o._statusCode, o.Payload)
}

func (o *IpamIPRangesBulkUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamIPRangesBulkUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
