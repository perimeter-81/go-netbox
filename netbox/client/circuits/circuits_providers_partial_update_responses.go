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

// CircuitsProvidersPartialUpdateReader is a Reader for the CircuitsProvidersPartialUpdate structure.
type CircuitsProvidersPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CircuitsProvidersPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCircuitsProvidersPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCircuitsProvidersPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCircuitsProvidersPartialUpdateOK creates a CircuitsProvidersPartialUpdateOK with default headers values
func NewCircuitsProvidersPartialUpdateOK() *CircuitsProvidersPartialUpdateOK {
	return &CircuitsProvidersPartialUpdateOK{}
}

/*
CircuitsProvidersPartialUpdateOK describes a response with status code 200, with default header values.

CircuitsProvidersPartialUpdateOK circuits providers partial update o k
*/
type CircuitsProvidersPartialUpdateOK struct {
	Payload *models.Provider
}

// IsSuccess returns true when this circuits providers partial update o k response has a 2xx status code
func (o *CircuitsProvidersPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this circuits providers partial update o k response has a 3xx status code
func (o *CircuitsProvidersPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this circuits providers partial update o k response has a 4xx status code
func (o *CircuitsProvidersPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this circuits providers partial update o k response has a 5xx status code
func (o *CircuitsProvidersPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this circuits providers partial update o k response a status code equal to that given
func (o *CircuitsProvidersPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *CircuitsProvidersPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /circuits/providers/{id}/][%d] circuitsProvidersPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *CircuitsProvidersPartialUpdateOK) String() string {
	return fmt.Sprintf("[PATCH /circuits/providers/{id}/][%d] circuitsProvidersPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *CircuitsProvidersPartialUpdateOK) GetPayload() *models.Provider {
	return o.Payload
}

func (o *CircuitsProvidersPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Provider)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCircuitsProvidersPartialUpdateDefault creates a CircuitsProvidersPartialUpdateDefault with default headers values
func NewCircuitsProvidersPartialUpdateDefault(code int) *CircuitsProvidersPartialUpdateDefault {
	return &CircuitsProvidersPartialUpdateDefault{
		_statusCode: code,
	}
}

/*
CircuitsProvidersPartialUpdateDefault describes a response with status code -1, with default header values.

CircuitsProvidersPartialUpdateDefault circuits providers partial update default
*/
type CircuitsProvidersPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the circuits providers partial update default response
func (o *CircuitsProvidersPartialUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this circuits providers partial update default response has a 2xx status code
func (o *CircuitsProvidersPartialUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this circuits providers partial update default response has a 3xx status code
func (o *CircuitsProvidersPartialUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this circuits providers partial update default response has a 4xx status code
func (o *CircuitsProvidersPartialUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this circuits providers partial update default response has a 5xx status code
func (o *CircuitsProvidersPartialUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this circuits providers partial update default response a status code equal to that given
func (o *CircuitsProvidersPartialUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *CircuitsProvidersPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /circuits/providers/{id}/][%d] circuits_providers_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *CircuitsProvidersPartialUpdateDefault) String() string {
	return fmt.Sprintf("[PATCH /circuits/providers/{id}/][%d] circuits_providers_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *CircuitsProvidersPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *CircuitsProvidersPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
