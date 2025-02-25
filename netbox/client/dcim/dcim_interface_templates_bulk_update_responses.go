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

// DcimInterfaceTemplatesBulkUpdateReader is a Reader for the DcimInterfaceTemplatesBulkUpdate structure.
type DcimInterfaceTemplatesBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimInterfaceTemplatesBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimInterfaceTemplatesBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimInterfaceTemplatesBulkUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimInterfaceTemplatesBulkUpdateOK creates a DcimInterfaceTemplatesBulkUpdateOK with default headers values
func NewDcimInterfaceTemplatesBulkUpdateOK() *DcimInterfaceTemplatesBulkUpdateOK {
	return &DcimInterfaceTemplatesBulkUpdateOK{}
}

/*
DcimInterfaceTemplatesBulkUpdateOK describes a response with status code 200, with default header values.

DcimInterfaceTemplatesBulkUpdateOK dcim interface templates bulk update o k
*/
type DcimInterfaceTemplatesBulkUpdateOK struct {
	Payload *models.InterfaceTemplate
}

// IsSuccess returns true when this dcim interface templates bulk update o k response has a 2xx status code
func (o *DcimInterfaceTemplatesBulkUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim interface templates bulk update o k response has a 3xx status code
func (o *DcimInterfaceTemplatesBulkUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim interface templates bulk update o k response has a 4xx status code
func (o *DcimInterfaceTemplatesBulkUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim interface templates bulk update o k response has a 5xx status code
func (o *DcimInterfaceTemplatesBulkUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim interface templates bulk update o k response a status code equal to that given
func (o *DcimInterfaceTemplatesBulkUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *DcimInterfaceTemplatesBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/interface-templates/][%d] dcimInterfaceTemplatesBulkUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimInterfaceTemplatesBulkUpdateOK) String() string {
	return fmt.Sprintf("[PUT /dcim/interface-templates/][%d] dcimInterfaceTemplatesBulkUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimInterfaceTemplatesBulkUpdateOK) GetPayload() *models.InterfaceTemplate {
	return o.Payload
}

func (o *DcimInterfaceTemplatesBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InterfaceTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimInterfaceTemplatesBulkUpdateDefault creates a DcimInterfaceTemplatesBulkUpdateDefault with default headers values
func NewDcimInterfaceTemplatesBulkUpdateDefault(code int) *DcimInterfaceTemplatesBulkUpdateDefault {
	return &DcimInterfaceTemplatesBulkUpdateDefault{
		_statusCode: code,
	}
}

/*
DcimInterfaceTemplatesBulkUpdateDefault describes a response with status code -1, with default header values.

DcimInterfaceTemplatesBulkUpdateDefault dcim interface templates bulk update default
*/
type DcimInterfaceTemplatesBulkUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim interface templates bulk update default response
func (o *DcimInterfaceTemplatesBulkUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim interface templates bulk update default response has a 2xx status code
func (o *DcimInterfaceTemplatesBulkUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim interface templates bulk update default response has a 3xx status code
func (o *DcimInterfaceTemplatesBulkUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim interface templates bulk update default response has a 4xx status code
func (o *DcimInterfaceTemplatesBulkUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim interface templates bulk update default response has a 5xx status code
func (o *DcimInterfaceTemplatesBulkUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim interface templates bulk update default response a status code equal to that given
func (o *DcimInterfaceTemplatesBulkUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimInterfaceTemplatesBulkUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /dcim/interface-templates/][%d] dcim_interface-templates_bulk_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimInterfaceTemplatesBulkUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /dcim/interface-templates/][%d] dcim_interface-templates_bulk_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimInterfaceTemplatesBulkUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimInterfaceTemplatesBulkUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
