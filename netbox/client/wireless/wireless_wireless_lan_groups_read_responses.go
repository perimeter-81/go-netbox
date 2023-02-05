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

package wireless

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/perimeter-81/go-netbox/netbox/models"
)

// WirelessWirelessLanGroupsReadReader is a Reader for the WirelessWirelessLanGroupsRead structure.
type WirelessWirelessLanGroupsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WirelessWirelessLanGroupsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWirelessWirelessLanGroupsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWirelessWirelessLanGroupsReadDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWirelessWirelessLanGroupsReadOK creates a WirelessWirelessLanGroupsReadOK with default headers values
func NewWirelessWirelessLanGroupsReadOK() *WirelessWirelessLanGroupsReadOK {
	return &WirelessWirelessLanGroupsReadOK{}
}

/*
WirelessWirelessLanGroupsReadOK describes a response with status code 200, with default header values.

WirelessWirelessLanGroupsReadOK wireless wireless lan groups read o k
*/
type WirelessWirelessLanGroupsReadOK struct {
	Payload *models.WirelessLANGroup
}

// IsSuccess returns true when this wireless wireless lan groups read o k response has a 2xx status code
func (o *WirelessWirelessLanGroupsReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this wireless wireless lan groups read o k response has a 3xx status code
func (o *WirelessWirelessLanGroupsReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this wireless wireless lan groups read o k response has a 4xx status code
func (o *WirelessWirelessLanGroupsReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this wireless wireless lan groups read o k response has a 5xx status code
func (o *WirelessWirelessLanGroupsReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this wireless wireless lan groups read o k response a status code equal to that given
func (o *WirelessWirelessLanGroupsReadOK) IsCode(code int) bool {
	return code == 200
}

func (o *WirelessWirelessLanGroupsReadOK) Error() string {
	return fmt.Sprintf("[GET /wireless/wireless-lan-groups/{id}/][%d] wirelessWirelessLanGroupsReadOK  %+v", 200, o.Payload)
}

func (o *WirelessWirelessLanGroupsReadOK) String() string {
	return fmt.Sprintf("[GET /wireless/wireless-lan-groups/{id}/][%d] wirelessWirelessLanGroupsReadOK  %+v", 200, o.Payload)
}

func (o *WirelessWirelessLanGroupsReadOK) GetPayload() *models.WirelessLANGroup {
	return o.Payload
}

func (o *WirelessWirelessLanGroupsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WirelessLANGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWirelessWirelessLanGroupsReadDefault creates a WirelessWirelessLanGroupsReadDefault with default headers values
func NewWirelessWirelessLanGroupsReadDefault(code int) *WirelessWirelessLanGroupsReadDefault {
	return &WirelessWirelessLanGroupsReadDefault{
		_statusCode: code,
	}
}

/*
WirelessWirelessLanGroupsReadDefault describes a response with status code -1, with default header values.

WirelessWirelessLanGroupsReadDefault wireless wireless lan groups read default
*/
type WirelessWirelessLanGroupsReadDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the wireless wireless lan groups read default response
func (o *WirelessWirelessLanGroupsReadDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this wireless wireless lan groups read default response has a 2xx status code
func (o *WirelessWirelessLanGroupsReadDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this wireless wireless lan groups read default response has a 3xx status code
func (o *WirelessWirelessLanGroupsReadDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this wireless wireless lan groups read default response has a 4xx status code
func (o *WirelessWirelessLanGroupsReadDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this wireless wireless lan groups read default response has a 5xx status code
func (o *WirelessWirelessLanGroupsReadDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this wireless wireless lan groups read default response a status code equal to that given
func (o *WirelessWirelessLanGroupsReadDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *WirelessWirelessLanGroupsReadDefault) Error() string {
	return fmt.Sprintf("[GET /wireless/wireless-lan-groups/{id}/][%d] wireless_wireless-lan-groups_read default  %+v", o._statusCode, o.Payload)
}

func (o *WirelessWirelessLanGroupsReadDefault) String() string {
	return fmt.Sprintf("[GET /wireless/wireless-lan-groups/{id}/][%d] wireless_wireless-lan-groups_read default  %+v", o._statusCode, o.Payload)
}

func (o *WirelessWirelessLanGroupsReadDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *WirelessWirelessLanGroupsReadDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
