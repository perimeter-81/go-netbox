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

// ExtrasConfigContextsCreateReader is a Reader for the ExtrasConfigContextsCreate structure.
type ExtrasConfigContextsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasConfigContextsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewExtrasConfigContextsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewExtrasConfigContextsCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewExtrasConfigContextsCreateCreated creates a ExtrasConfigContextsCreateCreated with default headers values
func NewExtrasConfigContextsCreateCreated() *ExtrasConfigContextsCreateCreated {
	return &ExtrasConfigContextsCreateCreated{}
}

/*
ExtrasConfigContextsCreateCreated describes a response with status code 201, with default header values.

ExtrasConfigContextsCreateCreated extras config contexts create created
*/
type ExtrasConfigContextsCreateCreated struct {
	Payload *models.ConfigContext
}

// IsSuccess returns true when this extras config contexts create created response has a 2xx status code
func (o *ExtrasConfigContextsCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this extras config contexts create created response has a 3xx status code
func (o *ExtrasConfigContextsCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extras config contexts create created response has a 4xx status code
func (o *ExtrasConfigContextsCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this extras config contexts create created response has a 5xx status code
func (o *ExtrasConfigContextsCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this extras config contexts create created response a status code equal to that given
func (o *ExtrasConfigContextsCreateCreated) IsCode(code int) bool {
	return code == 201
}

func (o *ExtrasConfigContextsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /extras/config-contexts/][%d] extrasConfigContextsCreateCreated  %+v", 201, o.Payload)
}

func (o *ExtrasConfigContextsCreateCreated) String() string {
	return fmt.Sprintf("[POST /extras/config-contexts/][%d] extrasConfigContextsCreateCreated  %+v", 201, o.Payload)
}

func (o *ExtrasConfigContextsCreateCreated) GetPayload() *models.ConfigContext {
	return o.Payload
}

func (o *ExtrasConfigContextsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConfigContext)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExtrasConfigContextsCreateDefault creates a ExtrasConfigContextsCreateDefault with default headers values
func NewExtrasConfigContextsCreateDefault(code int) *ExtrasConfigContextsCreateDefault {
	return &ExtrasConfigContextsCreateDefault{
		_statusCode: code,
	}
}

/*
ExtrasConfigContextsCreateDefault describes a response with status code -1, with default header values.

ExtrasConfigContextsCreateDefault extras config contexts create default
*/
type ExtrasConfigContextsCreateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the extras config contexts create default response
func (o *ExtrasConfigContextsCreateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this extras config contexts create default response has a 2xx status code
func (o *ExtrasConfigContextsCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this extras config contexts create default response has a 3xx status code
func (o *ExtrasConfigContextsCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this extras config contexts create default response has a 4xx status code
func (o *ExtrasConfigContextsCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this extras config contexts create default response has a 5xx status code
func (o *ExtrasConfigContextsCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this extras config contexts create default response a status code equal to that given
func (o *ExtrasConfigContextsCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ExtrasConfigContextsCreateDefault) Error() string {
	return fmt.Sprintf("[POST /extras/config-contexts/][%d] extras_config-contexts_create default  %+v", o._statusCode, o.Payload)
}

func (o *ExtrasConfigContextsCreateDefault) String() string {
	return fmt.Sprintf("[POST /extras/config-contexts/][%d] extras_config-contexts_create default  %+v", o._statusCode, o.Payload)
}

func (o *ExtrasConfigContextsCreateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *ExtrasConfigContextsCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
