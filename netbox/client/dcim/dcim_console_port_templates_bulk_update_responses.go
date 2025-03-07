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

// DcimConsolePortTemplatesBulkUpdateReader is a Reader for the DcimConsolePortTemplatesBulkUpdate structure.
type DcimConsolePortTemplatesBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimConsolePortTemplatesBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimConsolePortTemplatesBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimConsolePortTemplatesBulkUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimConsolePortTemplatesBulkUpdateOK creates a DcimConsolePortTemplatesBulkUpdateOK with default headers values
func NewDcimConsolePortTemplatesBulkUpdateOK() *DcimConsolePortTemplatesBulkUpdateOK {
	return &DcimConsolePortTemplatesBulkUpdateOK{}
}

/*
DcimConsolePortTemplatesBulkUpdateOK describes a response with status code 200, with default header values.

DcimConsolePortTemplatesBulkUpdateOK dcim console port templates bulk update o k
*/
type DcimConsolePortTemplatesBulkUpdateOK struct {
	Payload *models.ConsolePortTemplate
}

// IsSuccess returns true when this dcim console port templates bulk update o k response has a 2xx status code
func (o *DcimConsolePortTemplatesBulkUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim console port templates bulk update o k response has a 3xx status code
func (o *DcimConsolePortTemplatesBulkUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim console port templates bulk update o k response has a 4xx status code
func (o *DcimConsolePortTemplatesBulkUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim console port templates bulk update o k response has a 5xx status code
func (o *DcimConsolePortTemplatesBulkUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim console port templates bulk update o k response a status code equal to that given
func (o *DcimConsolePortTemplatesBulkUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *DcimConsolePortTemplatesBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/console-port-templates/][%d] dcimConsolePortTemplatesBulkUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimConsolePortTemplatesBulkUpdateOK) String() string {
	return fmt.Sprintf("[PUT /dcim/console-port-templates/][%d] dcimConsolePortTemplatesBulkUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimConsolePortTemplatesBulkUpdateOK) GetPayload() *models.ConsolePortTemplate {
	return o.Payload
}

func (o *DcimConsolePortTemplatesBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConsolePortTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimConsolePortTemplatesBulkUpdateDefault creates a DcimConsolePortTemplatesBulkUpdateDefault with default headers values
func NewDcimConsolePortTemplatesBulkUpdateDefault(code int) *DcimConsolePortTemplatesBulkUpdateDefault {
	return &DcimConsolePortTemplatesBulkUpdateDefault{
		_statusCode: code,
	}
}

/*
DcimConsolePortTemplatesBulkUpdateDefault describes a response with status code -1, with default header values.

DcimConsolePortTemplatesBulkUpdateDefault dcim console port templates bulk update default
*/
type DcimConsolePortTemplatesBulkUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim console port templates bulk update default response
func (o *DcimConsolePortTemplatesBulkUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim console port templates bulk update default response has a 2xx status code
func (o *DcimConsolePortTemplatesBulkUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim console port templates bulk update default response has a 3xx status code
func (o *DcimConsolePortTemplatesBulkUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim console port templates bulk update default response has a 4xx status code
func (o *DcimConsolePortTemplatesBulkUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim console port templates bulk update default response has a 5xx status code
func (o *DcimConsolePortTemplatesBulkUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim console port templates bulk update default response a status code equal to that given
func (o *DcimConsolePortTemplatesBulkUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimConsolePortTemplatesBulkUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /dcim/console-port-templates/][%d] dcim_console-port-templates_bulk_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimConsolePortTemplatesBulkUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /dcim/console-port-templates/][%d] dcim_console-port-templates_bulk_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimConsolePortTemplatesBulkUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimConsolePortTemplatesBulkUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
