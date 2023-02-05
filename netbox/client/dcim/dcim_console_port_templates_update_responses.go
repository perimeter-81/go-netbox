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

// DcimConsolePortTemplatesUpdateReader is a Reader for the DcimConsolePortTemplatesUpdate structure.
type DcimConsolePortTemplatesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimConsolePortTemplatesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimConsolePortTemplatesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimConsolePortTemplatesUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimConsolePortTemplatesUpdateOK creates a DcimConsolePortTemplatesUpdateOK with default headers values
func NewDcimConsolePortTemplatesUpdateOK() *DcimConsolePortTemplatesUpdateOK {
	return &DcimConsolePortTemplatesUpdateOK{}
}

/*
DcimConsolePortTemplatesUpdateOK describes a response with status code 200, with default header values.

DcimConsolePortTemplatesUpdateOK dcim console port templates update o k
*/
type DcimConsolePortTemplatesUpdateOK struct {
	Payload *models.ConsolePortTemplate
}

// IsSuccess returns true when this dcim console port templates update o k response has a 2xx status code
func (o *DcimConsolePortTemplatesUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim console port templates update o k response has a 3xx status code
func (o *DcimConsolePortTemplatesUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim console port templates update o k response has a 4xx status code
func (o *DcimConsolePortTemplatesUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim console port templates update o k response has a 5xx status code
func (o *DcimConsolePortTemplatesUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim console port templates update o k response a status code equal to that given
func (o *DcimConsolePortTemplatesUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *DcimConsolePortTemplatesUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/console-port-templates/{id}/][%d] dcimConsolePortTemplatesUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimConsolePortTemplatesUpdateOK) String() string {
	return fmt.Sprintf("[PUT /dcim/console-port-templates/{id}/][%d] dcimConsolePortTemplatesUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimConsolePortTemplatesUpdateOK) GetPayload() *models.ConsolePortTemplate {
	return o.Payload
}

func (o *DcimConsolePortTemplatesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConsolePortTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimConsolePortTemplatesUpdateDefault creates a DcimConsolePortTemplatesUpdateDefault with default headers values
func NewDcimConsolePortTemplatesUpdateDefault(code int) *DcimConsolePortTemplatesUpdateDefault {
	return &DcimConsolePortTemplatesUpdateDefault{
		_statusCode: code,
	}
}

/*
DcimConsolePortTemplatesUpdateDefault describes a response with status code -1, with default header values.

DcimConsolePortTemplatesUpdateDefault dcim console port templates update default
*/
type DcimConsolePortTemplatesUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim console port templates update default response
func (o *DcimConsolePortTemplatesUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim console port templates update default response has a 2xx status code
func (o *DcimConsolePortTemplatesUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim console port templates update default response has a 3xx status code
func (o *DcimConsolePortTemplatesUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim console port templates update default response has a 4xx status code
func (o *DcimConsolePortTemplatesUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim console port templates update default response has a 5xx status code
func (o *DcimConsolePortTemplatesUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim console port templates update default response a status code equal to that given
func (o *DcimConsolePortTemplatesUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimConsolePortTemplatesUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /dcim/console-port-templates/{id}/][%d] dcim_console-port-templates_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimConsolePortTemplatesUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /dcim/console-port-templates/{id}/][%d] dcim_console-port-templates_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimConsolePortTemplatesUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimConsolePortTemplatesUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
