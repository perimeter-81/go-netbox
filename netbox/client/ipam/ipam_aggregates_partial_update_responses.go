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

// IpamAggregatesPartialUpdateReader is a Reader for the IpamAggregatesPartialUpdate structure.
type IpamAggregatesPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamAggregatesPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamAggregatesPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIpamAggregatesPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpamAggregatesPartialUpdateOK creates a IpamAggregatesPartialUpdateOK with default headers values
func NewIpamAggregatesPartialUpdateOK() *IpamAggregatesPartialUpdateOK {
	return &IpamAggregatesPartialUpdateOK{}
}

/*
IpamAggregatesPartialUpdateOK describes a response with status code 200, with default header values.

IpamAggregatesPartialUpdateOK ipam aggregates partial update o k
*/
type IpamAggregatesPartialUpdateOK struct {
	Payload *models.Aggregate
}

// IsSuccess returns true when this ipam aggregates partial update o k response has a 2xx status code
func (o *IpamAggregatesPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam aggregates partial update o k response has a 3xx status code
func (o *IpamAggregatesPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam aggregates partial update o k response has a 4xx status code
func (o *IpamAggregatesPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam aggregates partial update o k response has a 5xx status code
func (o *IpamAggregatesPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam aggregates partial update o k response a status code equal to that given
func (o *IpamAggregatesPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *IpamAggregatesPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /ipam/aggregates/{id}/][%d] ipamAggregatesPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *IpamAggregatesPartialUpdateOK) String() string {
	return fmt.Sprintf("[PATCH /ipam/aggregates/{id}/][%d] ipamAggregatesPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *IpamAggregatesPartialUpdateOK) GetPayload() *models.Aggregate {
	return o.Payload
}

func (o *IpamAggregatesPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Aggregate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIpamAggregatesPartialUpdateDefault creates a IpamAggregatesPartialUpdateDefault with default headers values
func NewIpamAggregatesPartialUpdateDefault(code int) *IpamAggregatesPartialUpdateDefault {
	return &IpamAggregatesPartialUpdateDefault{
		_statusCode: code,
	}
}

/*
IpamAggregatesPartialUpdateDefault describes a response with status code -1, with default header values.

IpamAggregatesPartialUpdateDefault ipam aggregates partial update default
*/
type IpamAggregatesPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the ipam aggregates partial update default response
func (o *IpamAggregatesPartialUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this ipam aggregates partial update default response has a 2xx status code
func (o *IpamAggregatesPartialUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ipam aggregates partial update default response has a 3xx status code
func (o *IpamAggregatesPartialUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ipam aggregates partial update default response has a 4xx status code
func (o *IpamAggregatesPartialUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ipam aggregates partial update default response has a 5xx status code
func (o *IpamAggregatesPartialUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ipam aggregates partial update default response a status code equal to that given
func (o *IpamAggregatesPartialUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *IpamAggregatesPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /ipam/aggregates/{id}/][%d] ipam_aggregates_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *IpamAggregatesPartialUpdateDefault) String() string {
	return fmt.Sprintf("[PATCH /ipam/aggregates/{id}/][%d] ipam_aggregates_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *IpamAggregatesPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamAggregatesPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
