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
)

// DcimConsolePortsDeleteReader is a Reader for the DcimConsolePortsDelete structure.
type DcimConsolePortsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimConsolePortsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimConsolePortsDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimConsolePortsDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimConsolePortsDeleteNoContent creates a DcimConsolePortsDeleteNoContent with default headers values
func NewDcimConsolePortsDeleteNoContent() *DcimConsolePortsDeleteNoContent {
	return &DcimConsolePortsDeleteNoContent{}
}

/*
DcimConsolePortsDeleteNoContent describes a response with status code 204, with default header values.

DcimConsolePortsDeleteNoContent dcim console ports delete no content
*/
type DcimConsolePortsDeleteNoContent struct {
}

// IsSuccess returns true when this dcim console ports delete no content response has a 2xx status code
func (o *DcimConsolePortsDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim console ports delete no content response has a 3xx status code
func (o *DcimConsolePortsDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim console ports delete no content response has a 4xx status code
func (o *DcimConsolePortsDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim console ports delete no content response has a 5xx status code
func (o *DcimConsolePortsDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim console ports delete no content response a status code equal to that given
func (o *DcimConsolePortsDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *DcimConsolePortsDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/console-ports/{id}/][%d] dcimConsolePortsDeleteNoContent ", 204)
}

func (o *DcimConsolePortsDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /dcim/console-ports/{id}/][%d] dcimConsolePortsDeleteNoContent ", 204)
}

func (o *DcimConsolePortsDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDcimConsolePortsDeleteDefault creates a DcimConsolePortsDeleteDefault with default headers values
func NewDcimConsolePortsDeleteDefault(code int) *DcimConsolePortsDeleteDefault {
	return &DcimConsolePortsDeleteDefault{
		_statusCode: code,
	}
}

/*
DcimConsolePortsDeleteDefault describes a response with status code -1, with default header values.

DcimConsolePortsDeleteDefault dcim console ports delete default
*/
type DcimConsolePortsDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim console ports delete default response
func (o *DcimConsolePortsDeleteDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim console ports delete default response has a 2xx status code
func (o *DcimConsolePortsDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim console ports delete default response has a 3xx status code
func (o *DcimConsolePortsDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim console ports delete default response has a 4xx status code
func (o *DcimConsolePortsDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim console ports delete default response has a 5xx status code
func (o *DcimConsolePortsDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim console ports delete default response a status code equal to that given
func (o *DcimConsolePortsDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimConsolePortsDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /dcim/console-ports/{id}/][%d] dcim_console-ports_delete default  %+v", o._statusCode, o.Payload)
}

func (o *DcimConsolePortsDeleteDefault) String() string {
	return fmt.Sprintf("[DELETE /dcim/console-ports/{id}/][%d] dcim_console-ports_delete default  %+v", o._statusCode, o.Payload)
}

func (o *DcimConsolePortsDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimConsolePortsDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
