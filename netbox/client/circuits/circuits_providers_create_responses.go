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

package circuits

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/perimeter-81/go-netbox/netbox/models"
)

// CircuitsProvidersCreateReader is a Reader for the CircuitsProvidersCreate structure.
type CircuitsProvidersCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CircuitsProvidersCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCircuitsProvidersCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCircuitsProvidersCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCircuitsProvidersCreateCreated creates a CircuitsProvidersCreateCreated with default headers values
func NewCircuitsProvidersCreateCreated() *CircuitsProvidersCreateCreated {
	return &CircuitsProvidersCreateCreated{}
}

/*
CircuitsProvidersCreateCreated describes a response with status code 201, with default header values.

CircuitsProvidersCreateCreated circuits providers create created
*/
type CircuitsProvidersCreateCreated struct {
	Payload *models.Provider
}

// IsSuccess returns true when this circuits providers create created response has a 2xx status code
func (o *CircuitsProvidersCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this circuits providers create created response has a 3xx status code
func (o *CircuitsProvidersCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this circuits providers create created response has a 4xx status code
func (o *CircuitsProvidersCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this circuits providers create created response has a 5xx status code
func (o *CircuitsProvidersCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this circuits providers create created response a status code equal to that given
func (o *CircuitsProvidersCreateCreated) IsCode(code int) bool {
	return code == 201
}

func (o *CircuitsProvidersCreateCreated) Error() string {
	return fmt.Sprintf("[POST /circuits/providers/][%d] circuitsProvidersCreateCreated  %+v", 201, o.Payload)
}

func (o *CircuitsProvidersCreateCreated) String() string {
	return fmt.Sprintf("[POST /circuits/providers/][%d] circuitsProvidersCreateCreated  %+v", 201, o.Payload)
}

func (o *CircuitsProvidersCreateCreated) GetPayload() *models.Provider {
	return o.Payload
}

func (o *CircuitsProvidersCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Provider)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCircuitsProvidersCreateDefault creates a CircuitsProvidersCreateDefault with default headers values
func NewCircuitsProvidersCreateDefault(code int) *CircuitsProvidersCreateDefault {
	return &CircuitsProvidersCreateDefault{
		_statusCode: code,
	}
}

/*
CircuitsProvidersCreateDefault describes a response with status code -1, with default header values.

CircuitsProvidersCreateDefault circuits providers create default
*/
type CircuitsProvidersCreateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the circuits providers create default response
func (o *CircuitsProvidersCreateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this circuits providers create default response has a 2xx status code
func (o *CircuitsProvidersCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this circuits providers create default response has a 3xx status code
func (o *CircuitsProvidersCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this circuits providers create default response has a 4xx status code
func (o *CircuitsProvidersCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this circuits providers create default response has a 5xx status code
func (o *CircuitsProvidersCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this circuits providers create default response a status code equal to that given
func (o *CircuitsProvidersCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *CircuitsProvidersCreateDefault) Error() string {
	return fmt.Sprintf("[POST /circuits/providers/][%d] circuits_providers_create default  %+v", o._statusCode, o.Payload)
}

func (o *CircuitsProvidersCreateDefault) String() string {
	return fmt.Sprintf("[POST /circuits/providers/][%d] circuits_providers_create default  %+v", o._statusCode, o.Payload)
}

func (o *CircuitsProvidersCreateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *CircuitsProvidersCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
