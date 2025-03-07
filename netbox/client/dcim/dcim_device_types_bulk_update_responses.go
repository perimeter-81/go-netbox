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

package dcim

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/perimeter-81/go-netbox/netbox/models"
)

// DcimDeviceTypesBulkUpdateReader is a Reader for the DcimDeviceTypesBulkUpdate structure.
type DcimDeviceTypesBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimDeviceTypesBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimDeviceTypesBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimDeviceTypesBulkUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimDeviceTypesBulkUpdateOK creates a DcimDeviceTypesBulkUpdateOK with default headers values
func NewDcimDeviceTypesBulkUpdateOK() *DcimDeviceTypesBulkUpdateOK {
	return &DcimDeviceTypesBulkUpdateOK{}
}

/*
DcimDeviceTypesBulkUpdateOK describes a response with status code 200, with default header values.

DcimDeviceTypesBulkUpdateOK dcim device types bulk update o k
*/
type DcimDeviceTypesBulkUpdateOK struct {
	Payload *models.DeviceType
}

// IsSuccess returns true when this dcim device types bulk update o k response has a 2xx status code
func (o *DcimDeviceTypesBulkUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim device types bulk update o k response has a 3xx status code
func (o *DcimDeviceTypesBulkUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim device types bulk update o k response has a 4xx status code
func (o *DcimDeviceTypesBulkUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim device types bulk update o k response has a 5xx status code
func (o *DcimDeviceTypesBulkUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim device types bulk update o k response a status code equal to that given
func (o *DcimDeviceTypesBulkUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *DcimDeviceTypesBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/device-types/][%d] dcimDeviceTypesBulkUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimDeviceTypesBulkUpdateOK) String() string {
	return fmt.Sprintf("[PUT /dcim/device-types/][%d] dcimDeviceTypesBulkUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimDeviceTypesBulkUpdateOK) GetPayload() *models.DeviceType {
	return o.Payload
}

func (o *DcimDeviceTypesBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeviceType)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimDeviceTypesBulkUpdateDefault creates a DcimDeviceTypesBulkUpdateDefault with default headers values
func NewDcimDeviceTypesBulkUpdateDefault(code int) *DcimDeviceTypesBulkUpdateDefault {
	return &DcimDeviceTypesBulkUpdateDefault{
		_statusCode: code,
	}
}

/*
DcimDeviceTypesBulkUpdateDefault describes a response with status code -1, with default header values.

DcimDeviceTypesBulkUpdateDefault dcim device types bulk update default
*/
type DcimDeviceTypesBulkUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim device types bulk update default response
func (o *DcimDeviceTypesBulkUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim device types bulk update default response has a 2xx status code
func (o *DcimDeviceTypesBulkUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim device types bulk update default response has a 3xx status code
func (o *DcimDeviceTypesBulkUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim device types bulk update default response has a 4xx status code
func (o *DcimDeviceTypesBulkUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim device types bulk update default response has a 5xx status code
func (o *DcimDeviceTypesBulkUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim device types bulk update default response a status code equal to that given
func (o *DcimDeviceTypesBulkUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimDeviceTypesBulkUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /dcim/device-types/][%d] dcim_device-types_bulk_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimDeviceTypesBulkUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /dcim/device-types/][%d] dcim_device-types_bulk_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimDeviceTypesBulkUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimDeviceTypesBulkUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
