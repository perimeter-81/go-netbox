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

// WirelessWirelessLinksBulkPartialUpdateReader is a Reader for the WirelessWirelessLinksBulkPartialUpdate structure.
type WirelessWirelessLinksBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WirelessWirelessLinksBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWirelessWirelessLinksBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWirelessWirelessLinksBulkPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWirelessWirelessLinksBulkPartialUpdateOK creates a WirelessWirelessLinksBulkPartialUpdateOK with default headers values
func NewWirelessWirelessLinksBulkPartialUpdateOK() *WirelessWirelessLinksBulkPartialUpdateOK {
	return &WirelessWirelessLinksBulkPartialUpdateOK{}
}

/*
WirelessWirelessLinksBulkPartialUpdateOK describes a response with status code 200, with default header values.

WirelessWirelessLinksBulkPartialUpdateOK wireless wireless links bulk partial update o k
*/
type WirelessWirelessLinksBulkPartialUpdateOK struct {
	Payload *models.WirelessLink
}

// IsSuccess returns true when this wireless wireless links bulk partial update o k response has a 2xx status code
func (o *WirelessWirelessLinksBulkPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this wireless wireless links bulk partial update o k response has a 3xx status code
func (o *WirelessWirelessLinksBulkPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this wireless wireless links bulk partial update o k response has a 4xx status code
func (o *WirelessWirelessLinksBulkPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this wireless wireless links bulk partial update o k response has a 5xx status code
func (o *WirelessWirelessLinksBulkPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this wireless wireless links bulk partial update o k response a status code equal to that given
func (o *WirelessWirelessLinksBulkPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *WirelessWirelessLinksBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /wireless/wireless-links/][%d] wirelessWirelessLinksBulkPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *WirelessWirelessLinksBulkPartialUpdateOK) String() string {
	return fmt.Sprintf("[PATCH /wireless/wireless-links/][%d] wirelessWirelessLinksBulkPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *WirelessWirelessLinksBulkPartialUpdateOK) GetPayload() *models.WirelessLink {
	return o.Payload
}

func (o *WirelessWirelessLinksBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WirelessLink)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWirelessWirelessLinksBulkPartialUpdateDefault creates a WirelessWirelessLinksBulkPartialUpdateDefault with default headers values
func NewWirelessWirelessLinksBulkPartialUpdateDefault(code int) *WirelessWirelessLinksBulkPartialUpdateDefault {
	return &WirelessWirelessLinksBulkPartialUpdateDefault{
		_statusCode: code,
	}
}

/*
WirelessWirelessLinksBulkPartialUpdateDefault describes a response with status code -1, with default header values.

WirelessWirelessLinksBulkPartialUpdateDefault wireless wireless links bulk partial update default
*/
type WirelessWirelessLinksBulkPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the wireless wireless links bulk partial update default response
func (o *WirelessWirelessLinksBulkPartialUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this wireless wireless links bulk partial update default response has a 2xx status code
func (o *WirelessWirelessLinksBulkPartialUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this wireless wireless links bulk partial update default response has a 3xx status code
func (o *WirelessWirelessLinksBulkPartialUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this wireless wireless links bulk partial update default response has a 4xx status code
func (o *WirelessWirelessLinksBulkPartialUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this wireless wireless links bulk partial update default response has a 5xx status code
func (o *WirelessWirelessLinksBulkPartialUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this wireless wireless links bulk partial update default response a status code equal to that given
func (o *WirelessWirelessLinksBulkPartialUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *WirelessWirelessLinksBulkPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /wireless/wireless-links/][%d] wireless_wireless-links_bulk_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *WirelessWirelessLinksBulkPartialUpdateDefault) String() string {
	return fmt.Sprintf("[PATCH /wireless/wireless-links/][%d] wireless_wireless-links_bulk_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *WirelessWirelessLinksBulkPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *WirelessWirelessLinksBulkPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
