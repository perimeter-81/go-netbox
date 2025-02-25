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

package extras

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ExtrasReportsReadReader is a Reader for the ExtrasReportsRead structure.
type ExtrasReportsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasReportsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasReportsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewExtrasReportsReadDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewExtrasReportsReadOK creates a ExtrasReportsReadOK with default headers values
func NewExtrasReportsReadOK() *ExtrasReportsReadOK {
	return &ExtrasReportsReadOK{}
}

/*
ExtrasReportsReadOK describes a response with status code 200, with default header values.

ExtrasReportsReadOK extras reports read o k
*/
type ExtrasReportsReadOK struct {
}

// IsSuccess returns true when this extras reports read o k response has a 2xx status code
func (o *ExtrasReportsReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this extras reports read o k response has a 3xx status code
func (o *ExtrasReportsReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extras reports read o k response has a 4xx status code
func (o *ExtrasReportsReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this extras reports read o k response has a 5xx status code
func (o *ExtrasReportsReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this extras reports read o k response a status code equal to that given
func (o *ExtrasReportsReadOK) IsCode(code int) bool {
	return code == 200
}

func (o *ExtrasReportsReadOK) Error() string {
	return fmt.Sprintf("[GET /extras/reports/{id}/][%d] extrasReportsReadOK ", 200)
}

func (o *ExtrasReportsReadOK) String() string {
	return fmt.Sprintf("[GET /extras/reports/{id}/][%d] extrasReportsReadOK ", 200)
}

func (o *ExtrasReportsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewExtrasReportsReadDefault creates a ExtrasReportsReadDefault with default headers values
func NewExtrasReportsReadDefault(code int) *ExtrasReportsReadDefault {
	return &ExtrasReportsReadDefault{
		_statusCode: code,
	}
}

/*
ExtrasReportsReadDefault describes a response with status code -1, with default header values.

ExtrasReportsReadDefault extras reports read default
*/
type ExtrasReportsReadDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the extras reports read default response
func (o *ExtrasReportsReadDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this extras reports read default response has a 2xx status code
func (o *ExtrasReportsReadDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this extras reports read default response has a 3xx status code
func (o *ExtrasReportsReadDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this extras reports read default response has a 4xx status code
func (o *ExtrasReportsReadDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this extras reports read default response has a 5xx status code
func (o *ExtrasReportsReadDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this extras reports read default response a status code equal to that given
func (o *ExtrasReportsReadDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ExtrasReportsReadDefault) Error() string {
	return fmt.Sprintf("[GET /extras/reports/{id}/][%d] extras_reports_read default  %+v", o._statusCode, o.Payload)
}

func (o *ExtrasReportsReadDefault) String() string {
	return fmt.Sprintf("[GET /extras/reports/{id}/][%d] extras_reports_read default  %+v", o._statusCode, o.Payload)
}

func (o *ExtrasReportsReadDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *ExtrasReportsReadDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
