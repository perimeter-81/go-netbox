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

// DcimPowerPanelsBulkPartialUpdateReader is a Reader for the DcimPowerPanelsBulkPartialUpdate structure.
type DcimPowerPanelsBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPowerPanelsBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimPowerPanelsBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimPowerPanelsBulkPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimPowerPanelsBulkPartialUpdateOK creates a DcimPowerPanelsBulkPartialUpdateOK with default headers values
func NewDcimPowerPanelsBulkPartialUpdateOK() *DcimPowerPanelsBulkPartialUpdateOK {
	return &DcimPowerPanelsBulkPartialUpdateOK{}
}

/*
DcimPowerPanelsBulkPartialUpdateOK describes a response with status code 200, with default header values.

DcimPowerPanelsBulkPartialUpdateOK dcim power panels bulk partial update o k
*/
type DcimPowerPanelsBulkPartialUpdateOK struct {
	Payload *models.PowerPanel
}

// IsSuccess returns true when this dcim power panels bulk partial update o k response has a 2xx status code
func (o *DcimPowerPanelsBulkPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim power panels bulk partial update o k response has a 3xx status code
func (o *DcimPowerPanelsBulkPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim power panels bulk partial update o k response has a 4xx status code
func (o *DcimPowerPanelsBulkPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim power panels bulk partial update o k response has a 5xx status code
func (o *DcimPowerPanelsBulkPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim power panels bulk partial update o k response a status code equal to that given
func (o *DcimPowerPanelsBulkPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *DcimPowerPanelsBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/power-panels/][%d] dcimPowerPanelsBulkPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimPowerPanelsBulkPartialUpdateOK) String() string {
	return fmt.Sprintf("[PATCH /dcim/power-panels/][%d] dcimPowerPanelsBulkPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimPowerPanelsBulkPartialUpdateOK) GetPayload() *models.PowerPanel {
	return o.Payload
}

func (o *DcimPowerPanelsBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PowerPanel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimPowerPanelsBulkPartialUpdateDefault creates a DcimPowerPanelsBulkPartialUpdateDefault with default headers values
func NewDcimPowerPanelsBulkPartialUpdateDefault(code int) *DcimPowerPanelsBulkPartialUpdateDefault {
	return &DcimPowerPanelsBulkPartialUpdateDefault{
		_statusCode: code,
	}
}

/*
DcimPowerPanelsBulkPartialUpdateDefault describes a response with status code -1, with default header values.

DcimPowerPanelsBulkPartialUpdateDefault dcim power panels bulk partial update default
*/
type DcimPowerPanelsBulkPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim power panels bulk partial update default response
func (o *DcimPowerPanelsBulkPartialUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim power panels bulk partial update default response has a 2xx status code
func (o *DcimPowerPanelsBulkPartialUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim power panels bulk partial update default response has a 3xx status code
func (o *DcimPowerPanelsBulkPartialUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim power panels bulk partial update default response has a 4xx status code
func (o *DcimPowerPanelsBulkPartialUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim power panels bulk partial update default response has a 5xx status code
func (o *DcimPowerPanelsBulkPartialUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim power panels bulk partial update default response a status code equal to that given
func (o *DcimPowerPanelsBulkPartialUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimPowerPanelsBulkPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /dcim/power-panels/][%d] dcim_power-panels_bulk_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimPowerPanelsBulkPartialUpdateDefault) String() string {
	return fmt.Sprintf("[PATCH /dcim/power-panels/][%d] dcim_power-panels_bulk_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimPowerPanelsBulkPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimPowerPanelsBulkPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
