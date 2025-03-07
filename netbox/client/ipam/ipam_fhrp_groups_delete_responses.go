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
)

// IpamFhrpGroupsDeleteReader is a Reader for the IpamFhrpGroupsDelete structure.
type IpamFhrpGroupsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamFhrpGroupsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewIpamFhrpGroupsDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIpamFhrpGroupsDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpamFhrpGroupsDeleteNoContent creates a IpamFhrpGroupsDeleteNoContent with default headers values
func NewIpamFhrpGroupsDeleteNoContent() *IpamFhrpGroupsDeleteNoContent {
	return &IpamFhrpGroupsDeleteNoContent{}
}

/*
IpamFhrpGroupsDeleteNoContent describes a response with status code 204, with default header values.

IpamFhrpGroupsDeleteNoContent ipam fhrp groups delete no content
*/
type IpamFhrpGroupsDeleteNoContent struct {
}

// IsSuccess returns true when this ipam fhrp groups delete no content response has a 2xx status code
func (o *IpamFhrpGroupsDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam fhrp groups delete no content response has a 3xx status code
func (o *IpamFhrpGroupsDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam fhrp groups delete no content response has a 4xx status code
func (o *IpamFhrpGroupsDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam fhrp groups delete no content response has a 5xx status code
func (o *IpamFhrpGroupsDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam fhrp groups delete no content response a status code equal to that given
func (o *IpamFhrpGroupsDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *IpamFhrpGroupsDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /ipam/fhrp-groups/{id}/][%d] ipamFhrpGroupsDeleteNoContent ", 204)
}

func (o *IpamFhrpGroupsDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /ipam/fhrp-groups/{id}/][%d] ipamFhrpGroupsDeleteNoContent ", 204)
}

func (o *IpamFhrpGroupsDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewIpamFhrpGroupsDeleteDefault creates a IpamFhrpGroupsDeleteDefault with default headers values
func NewIpamFhrpGroupsDeleteDefault(code int) *IpamFhrpGroupsDeleteDefault {
	return &IpamFhrpGroupsDeleteDefault{
		_statusCode: code,
	}
}

/*
IpamFhrpGroupsDeleteDefault describes a response with status code -1, with default header values.

IpamFhrpGroupsDeleteDefault ipam fhrp groups delete default
*/
type IpamFhrpGroupsDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the ipam fhrp groups delete default response
func (o *IpamFhrpGroupsDeleteDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this ipam fhrp groups delete default response has a 2xx status code
func (o *IpamFhrpGroupsDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ipam fhrp groups delete default response has a 3xx status code
func (o *IpamFhrpGroupsDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ipam fhrp groups delete default response has a 4xx status code
func (o *IpamFhrpGroupsDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ipam fhrp groups delete default response has a 5xx status code
func (o *IpamFhrpGroupsDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ipam fhrp groups delete default response a status code equal to that given
func (o *IpamFhrpGroupsDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *IpamFhrpGroupsDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /ipam/fhrp-groups/{id}/][%d] ipam_fhrp-groups_delete default  %+v", o._statusCode, o.Payload)
}

func (o *IpamFhrpGroupsDeleteDefault) String() string {
	return fmt.Sprintf("[DELETE /ipam/fhrp-groups/{id}/][%d] ipam_fhrp-groups_delete default  %+v", o._statusCode, o.Payload)
}

func (o *IpamFhrpGroupsDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamFhrpGroupsDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
