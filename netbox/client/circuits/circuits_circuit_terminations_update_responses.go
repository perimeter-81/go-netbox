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

// CircuitsCircuitTerminationsUpdateReader is a Reader for the CircuitsCircuitTerminationsUpdate structure.
type CircuitsCircuitTerminationsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CircuitsCircuitTerminationsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCircuitsCircuitTerminationsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCircuitsCircuitTerminationsUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCircuitsCircuitTerminationsUpdateOK creates a CircuitsCircuitTerminationsUpdateOK with default headers values
func NewCircuitsCircuitTerminationsUpdateOK() *CircuitsCircuitTerminationsUpdateOK {
	return &CircuitsCircuitTerminationsUpdateOK{}
}

/*
CircuitsCircuitTerminationsUpdateOK describes a response with status code 200, with default header values.

CircuitsCircuitTerminationsUpdateOK circuits circuit terminations update o k
*/
type CircuitsCircuitTerminationsUpdateOK struct {
	Payload *models.CircuitTermination
}

// IsSuccess returns true when this circuits circuit terminations update o k response has a 2xx status code
func (o *CircuitsCircuitTerminationsUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this circuits circuit terminations update o k response has a 3xx status code
func (o *CircuitsCircuitTerminationsUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this circuits circuit terminations update o k response has a 4xx status code
func (o *CircuitsCircuitTerminationsUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this circuits circuit terminations update o k response has a 5xx status code
func (o *CircuitsCircuitTerminationsUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this circuits circuit terminations update o k response a status code equal to that given
func (o *CircuitsCircuitTerminationsUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *CircuitsCircuitTerminationsUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /circuits/circuit-terminations/{id}/][%d] circuitsCircuitTerminationsUpdateOK  %+v", 200, o.Payload)
}

func (o *CircuitsCircuitTerminationsUpdateOK) String() string {
	return fmt.Sprintf("[PUT /circuits/circuit-terminations/{id}/][%d] circuitsCircuitTerminationsUpdateOK  %+v", 200, o.Payload)
}

func (o *CircuitsCircuitTerminationsUpdateOK) GetPayload() *models.CircuitTermination {
	return o.Payload
}

func (o *CircuitsCircuitTerminationsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CircuitTermination)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCircuitsCircuitTerminationsUpdateDefault creates a CircuitsCircuitTerminationsUpdateDefault with default headers values
func NewCircuitsCircuitTerminationsUpdateDefault(code int) *CircuitsCircuitTerminationsUpdateDefault {
	return &CircuitsCircuitTerminationsUpdateDefault{
		_statusCode: code,
	}
}

/*
CircuitsCircuitTerminationsUpdateDefault describes a response with status code -1, with default header values.

CircuitsCircuitTerminationsUpdateDefault circuits circuit terminations update default
*/
type CircuitsCircuitTerminationsUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the circuits circuit terminations update default response
func (o *CircuitsCircuitTerminationsUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this circuits circuit terminations update default response has a 2xx status code
func (o *CircuitsCircuitTerminationsUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this circuits circuit terminations update default response has a 3xx status code
func (o *CircuitsCircuitTerminationsUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this circuits circuit terminations update default response has a 4xx status code
func (o *CircuitsCircuitTerminationsUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this circuits circuit terminations update default response has a 5xx status code
func (o *CircuitsCircuitTerminationsUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this circuits circuit terminations update default response a status code equal to that given
func (o *CircuitsCircuitTerminationsUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *CircuitsCircuitTerminationsUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /circuits/circuit-terminations/{id}/][%d] circuits_circuit-terminations_update default  %+v", o._statusCode, o.Payload)
}

func (o *CircuitsCircuitTerminationsUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /circuits/circuit-terminations/{id}/][%d] circuits_circuit-terminations_update default  %+v", o._statusCode, o.Payload)
}

func (o *CircuitsCircuitTerminationsUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *CircuitsCircuitTerminationsUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
