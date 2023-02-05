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

// DcimModulesCreateReader is a Reader for the DcimModulesCreate structure.
type DcimModulesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimModulesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewDcimModulesCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimModulesCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimModulesCreateCreated creates a DcimModulesCreateCreated with default headers values
func NewDcimModulesCreateCreated() *DcimModulesCreateCreated {
	return &DcimModulesCreateCreated{}
}

/*
DcimModulesCreateCreated describes a response with status code 201, with default header values.

DcimModulesCreateCreated dcim modules create created
*/
type DcimModulesCreateCreated struct {
	Payload *models.Module
}

// IsSuccess returns true when this dcim modules create created response has a 2xx status code
func (o *DcimModulesCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim modules create created response has a 3xx status code
func (o *DcimModulesCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim modules create created response has a 4xx status code
func (o *DcimModulesCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim modules create created response has a 5xx status code
func (o *DcimModulesCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim modules create created response a status code equal to that given
func (o *DcimModulesCreateCreated) IsCode(code int) bool {
	return code == 201
}

func (o *DcimModulesCreateCreated) Error() string {
	return fmt.Sprintf("[POST /dcim/modules/][%d] dcimModulesCreateCreated  %+v", 201, o.Payload)
}

func (o *DcimModulesCreateCreated) String() string {
	return fmt.Sprintf("[POST /dcim/modules/][%d] dcimModulesCreateCreated  %+v", 201, o.Payload)
}

func (o *DcimModulesCreateCreated) GetPayload() *models.Module {
	return o.Payload
}

func (o *DcimModulesCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Module)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimModulesCreateDefault creates a DcimModulesCreateDefault with default headers values
func NewDcimModulesCreateDefault(code int) *DcimModulesCreateDefault {
	return &DcimModulesCreateDefault{
		_statusCode: code,
	}
}

/*
DcimModulesCreateDefault describes a response with status code -1, with default header values.

DcimModulesCreateDefault dcim modules create default
*/
type DcimModulesCreateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim modules create default response
func (o *DcimModulesCreateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim modules create default response has a 2xx status code
func (o *DcimModulesCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim modules create default response has a 3xx status code
func (o *DcimModulesCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim modules create default response has a 4xx status code
func (o *DcimModulesCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim modules create default response has a 5xx status code
func (o *DcimModulesCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim modules create default response a status code equal to that given
func (o *DcimModulesCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimModulesCreateDefault) Error() string {
	return fmt.Sprintf("[POST /dcim/modules/][%d] dcim_modules_create default  %+v", o._statusCode, o.Payload)
}

func (o *DcimModulesCreateDefault) String() string {
	return fmt.Sprintf("[POST /dcim/modules/][%d] dcim_modules_create default  %+v", o._statusCode, o.Payload)
}

func (o *DcimModulesCreateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimModulesCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
