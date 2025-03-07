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

// IpamIPRangesAvailableIpsCreateReader is a Reader for the IpamIPRangesAvailableIpsCreate structure.
type IpamIPRangesAvailableIpsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamIPRangesAvailableIpsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewIpamIPRangesAvailableIpsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIpamIPRangesAvailableIpsCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpamIPRangesAvailableIpsCreateCreated creates a IpamIPRangesAvailableIpsCreateCreated with default headers values
func NewIpamIPRangesAvailableIpsCreateCreated() *IpamIPRangesAvailableIpsCreateCreated {
	return &IpamIPRangesAvailableIpsCreateCreated{}
}

/*
IpamIPRangesAvailableIpsCreateCreated describes a response with status code 201, with default header values.

IpamIPRangesAvailableIpsCreateCreated ipam Ip ranges available ips create created
*/
type IpamIPRangesAvailableIpsCreateCreated struct {
	Payload []*models.IPAddress
}

// IsSuccess returns true when this ipam Ip ranges available ips create created response has a 2xx status code
func (o *IpamIPRangesAvailableIpsCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam Ip ranges available ips create created response has a 3xx status code
func (o *IpamIPRangesAvailableIpsCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam Ip ranges available ips create created response has a 4xx status code
func (o *IpamIPRangesAvailableIpsCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam Ip ranges available ips create created response has a 5xx status code
func (o *IpamIPRangesAvailableIpsCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam Ip ranges available ips create created response a status code equal to that given
func (o *IpamIPRangesAvailableIpsCreateCreated) IsCode(code int) bool {
	return code == 201
}

func (o *IpamIPRangesAvailableIpsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /ipam/ip-ranges/{id}/available-ips/][%d] ipamIpRangesAvailableIpsCreateCreated  %+v", 201, o.Payload)
}

func (o *IpamIPRangesAvailableIpsCreateCreated) String() string {
	return fmt.Sprintf("[POST /ipam/ip-ranges/{id}/available-ips/][%d] ipamIpRangesAvailableIpsCreateCreated  %+v", 201, o.Payload)
}

func (o *IpamIPRangesAvailableIpsCreateCreated) GetPayload() []*models.IPAddress {
	return o.Payload
}

func (o *IpamIPRangesAvailableIpsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIpamIPRangesAvailableIpsCreateDefault creates a IpamIPRangesAvailableIpsCreateDefault with default headers values
func NewIpamIPRangesAvailableIpsCreateDefault(code int) *IpamIPRangesAvailableIpsCreateDefault {
	return &IpamIPRangesAvailableIpsCreateDefault{
		_statusCode: code,
	}
}

/*
IpamIPRangesAvailableIpsCreateDefault describes a response with status code -1, with default header values.

IpamIPRangesAvailableIpsCreateDefault ipam ip ranges available ips create default
*/
type IpamIPRangesAvailableIpsCreateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the ipam ip ranges available ips create default response
func (o *IpamIPRangesAvailableIpsCreateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this ipam ip ranges available ips create default response has a 2xx status code
func (o *IpamIPRangesAvailableIpsCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ipam ip ranges available ips create default response has a 3xx status code
func (o *IpamIPRangesAvailableIpsCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ipam ip ranges available ips create default response has a 4xx status code
func (o *IpamIPRangesAvailableIpsCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ipam ip ranges available ips create default response has a 5xx status code
func (o *IpamIPRangesAvailableIpsCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ipam ip ranges available ips create default response a status code equal to that given
func (o *IpamIPRangesAvailableIpsCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *IpamIPRangesAvailableIpsCreateDefault) Error() string {
	return fmt.Sprintf("[POST /ipam/ip-ranges/{id}/available-ips/][%d] ipam_ip-ranges_available-ips_create default  %+v", o._statusCode, o.Payload)
}

func (o *IpamIPRangesAvailableIpsCreateDefault) String() string {
	return fmt.Sprintf("[POST /ipam/ip-ranges/{id}/available-ips/][%d] ipam_ip-ranges_available-ips_create default  %+v", o._statusCode, o.Payload)
}

func (o *IpamIPRangesAvailableIpsCreateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamIPRangesAvailableIpsCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
