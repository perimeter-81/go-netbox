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

// ExtrasConfigContextsBulkPartialUpdateReader is a Reader for the ExtrasConfigContextsBulkPartialUpdate structure.
type ExtrasConfigContextsBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasConfigContextsBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasConfigContextsBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewExtrasConfigContextsBulkPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewExtrasConfigContextsBulkPartialUpdateOK creates a ExtrasConfigContextsBulkPartialUpdateOK with default headers values
func NewExtrasConfigContextsBulkPartialUpdateOK() *ExtrasConfigContextsBulkPartialUpdateOK {
	return &ExtrasConfigContextsBulkPartialUpdateOK{}
}

/*
ExtrasConfigContextsBulkPartialUpdateOK describes a response with status code 200, with default header values.

ExtrasConfigContextsBulkPartialUpdateOK extras config contexts bulk partial update o k
*/
type ExtrasConfigContextsBulkPartialUpdateOK struct {
	Payload *models.ConfigContext
}

// IsSuccess returns true when this extras config contexts bulk partial update o k response has a 2xx status code
func (o *ExtrasConfigContextsBulkPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this extras config contexts bulk partial update o k response has a 3xx status code
func (o *ExtrasConfigContextsBulkPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extras config contexts bulk partial update o k response has a 4xx status code
func (o *ExtrasConfigContextsBulkPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this extras config contexts bulk partial update o k response has a 5xx status code
func (o *ExtrasConfigContextsBulkPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this extras config contexts bulk partial update o k response a status code equal to that given
func (o *ExtrasConfigContextsBulkPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *ExtrasConfigContextsBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /extras/config-contexts/][%d] extrasConfigContextsBulkPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *ExtrasConfigContextsBulkPartialUpdateOK) String() string {
	return fmt.Sprintf("[PATCH /extras/config-contexts/][%d] extrasConfigContextsBulkPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *ExtrasConfigContextsBulkPartialUpdateOK) GetPayload() *models.ConfigContext {
	return o.Payload
}

func (o *ExtrasConfigContextsBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConfigContext)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExtrasConfigContextsBulkPartialUpdateDefault creates a ExtrasConfigContextsBulkPartialUpdateDefault with default headers values
func NewExtrasConfigContextsBulkPartialUpdateDefault(code int) *ExtrasConfigContextsBulkPartialUpdateDefault {
	return &ExtrasConfigContextsBulkPartialUpdateDefault{
		_statusCode: code,
	}
}

/*
ExtrasConfigContextsBulkPartialUpdateDefault describes a response with status code -1, with default header values.

ExtrasConfigContextsBulkPartialUpdateDefault extras config contexts bulk partial update default
*/
type ExtrasConfigContextsBulkPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the extras config contexts bulk partial update default response
func (o *ExtrasConfigContextsBulkPartialUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this extras config contexts bulk partial update default response has a 2xx status code
func (o *ExtrasConfigContextsBulkPartialUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this extras config contexts bulk partial update default response has a 3xx status code
func (o *ExtrasConfigContextsBulkPartialUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this extras config contexts bulk partial update default response has a 4xx status code
func (o *ExtrasConfigContextsBulkPartialUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this extras config contexts bulk partial update default response has a 5xx status code
func (o *ExtrasConfigContextsBulkPartialUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this extras config contexts bulk partial update default response a status code equal to that given
func (o *ExtrasConfigContextsBulkPartialUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ExtrasConfigContextsBulkPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /extras/config-contexts/][%d] extras_config-contexts_bulk_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *ExtrasConfigContextsBulkPartialUpdateDefault) String() string {
	return fmt.Sprintf("[PATCH /extras/config-contexts/][%d] extras_config-contexts_bulk_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *ExtrasConfigContextsBulkPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *ExtrasConfigContextsBulkPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
