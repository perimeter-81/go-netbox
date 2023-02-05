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

	"github.com/perimeter-81/go-netbox/netbox/models"
)

// ExtrasConfigContextsReadReader is a Reader for the ExtrasConfigContextsRead structure.
type ExtrasConfigContextsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasConfigContextsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasConfigContextsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewExtrasConfigContextsReadDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewExtrasConfigContextsReadOK creates a ExtrasConfigContextsReadOK with default headers values
func NewExtrasConfigContextsReadOK() *ExtrasConfigContextsReadOK {
	return &ExtrasConfigContextsReadOK{}
}

/*
ExtrasConfigContextsReadOK describes a response with status code 200, with default header values.

ExtrasConfigContextsReadOK extras config contexts read o k
*/
type ExtrasConfigContextsReadOK struct {
	Payload *models.ConfigContext
}

// IsSuccess returns true when this extras config contexts read o k response has a 2xx status code
func (o *ExtrasConfigContextsReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this extras config contexts read o k response has a 3xx status code
func (o *ExtrasConfigContextsReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extras config contexts read o k response has a 4xx status code
func (o *ExtrasConfigContextsReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this extras config contexts read o k response has a 5xx status code
func (o *ExtrasConfigContextsReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this extras config contexts read o k response a status code equal to that given
func (o *ExtrasConfigContextsReadOK) IsCode(code int) bool {
	return code == 200
}

func (o *ExtrasConfigContextsReadOK) Error() string {
	return fmt.Sprintf("[GET /extras/config-contexts/{id}/][%d] extrasConfigContextsReadOK  %+v", 200, o.Payload)
}

func (o *ExtrasConfigContextsReadOK) String() string {
	return fmt.Sprintf("[GET /extras/config-contexts/{id}/][%d] extrasConfigContextsReadOK  %+v", 200, o.Payload)
}

func (o *ExtrasConfigContextsReadOK) GetPayload() *models.ConfigContext {
	return o.Payload
}

func (o *ExtrasConfigContextsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConfigContext)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExtrasConfigContextsReadDefault creates a ExtrasConfigContextsReadDefault with default headers values
func NewExtrasConfigContextsReadDefault(code int) *ExtrasConfigContextsReadDefault {
	return &ExtrasConfigContextsReadDefault{
		_statusCode: code,
	}
}

/*
ExtrasConfigContextsReadDefault describes a response with status code -1, with default header values.

ExtrasConfigContextsReadDefault extras config contexts read default
*/
type ExtrasConfigContextsReadDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the extras config contexts read default response
func (o *ExtrasConfigContextsReadDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this extras config contexts read default response has a 2xx status code
func (o *ExtrasConfigContextsReadDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this extras config contexts read default response has a 3xx status code
func (o *ExtrasConfigContextsReadDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this extras config contexts read default response has a 4xx status code
func (o *ExtrasConfigContextsReadDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this extras config contexts read default response has a 5xx status code
func (o *ExtrasConfigContextsReadDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this extras config contexts read default response a status code equal to that given
func (o *ExtrasConfigContextsReadDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ExtrasConfigContextsReadDefault) Error() string {
	return fmt.Sprintf("[GET /extras/config-contexts/{id}/][%d] extras_config-contexts_read default  %+v", o._statusCode, o.Payload)
}

func (o *ExtrasConfigContextsReadDefault) String() string {
	return fmt.Sprintf("[GET /extras/config-contexts/{id}/][%d] extras_config-contexts_read default  %+v", o._statusCode, o.Payload)
}

func (o *ExtrasConfigContextsReadDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *ExtrasConfigContextsReadDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
