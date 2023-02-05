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

// DcimCablesPartialUpdateReader is a Reader for the DcimCablesPartialUpdate structure.
type DcimCablesPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimCablesPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimCablesPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimCablesPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimCablesPartialUpdateOK creates a DcimCablesPartialUpdateOK with default headers values
func NewDcimCablesPartialUpdateOK() *DcimCablesPartialUpdateOK {
	return &DcimCablesPartialUpdateOK{}
}

/*
DcimCablesPartialUpdateOK describes a response with status code 200, with default header values.

DcimCablesPartialUpdateOK dcim cables partial update o k
*/
type DcimCablesPartialUpdateOK struct {
	Payload *models.Cable
}

// IsSuccess returns true when this dcim cables partial update o k response has a 2xx status code
func (o *DcimCablesPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim cables partial update o k response has a 3xx status code
func (o *DcimCablesPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim cables partial update o k response has a 4xx status code
func (o *DcimCablesPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim cables partial update o k response has a 5xx status code
func (o *DcimCablesPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim cables partial update o k response a status code equal to that given
func (o *DcimCablesPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *DcimCablesPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/cables/{id}/][%d] dcimCablesPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimCablesPartialUpdateOK) String() string {
	return fmt.Sprintf("[PATCH /dcim/cables/{id}/][%d] dcimCablesPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimCablesPartialUpdateOK) GetPayload() *models.Cable {
	return o.Payload
}

func (o *DcimCablesPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Cable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimCablesPartialUpdateDefault creates a DcimCablesPartialUpdateDefault with default headers values
func NewDcimCablesPartialUpdateDefault(code int) *DcimCablesPartialUpdateDefault {
	return &DcimCablesPartialUpdateDefault{
		_statusCode: code,
	}
}

/*
DcimCablesPartialUpdateDefault describes a response with status code -1, with default header values.

DcimCablesPartialUpdateDefault dcim cables partial update default
*/
type DcimCablesPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim cables partial update default response
func (o *DcimCablesPartialUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim cables partial update default response has a 2xx status code
func (o *DcimCablesPartialUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim cables partial update default response has a 3xx status code
func (o *DcimCablesPartialUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim cables partial update default response has a 4xx status code
func (o *DcimCablesPartialUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim cables partial update default response has a 5xx status code
func (o *DcimCablesPartialUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim cables partial update default response a status code equal to that given
func (o *DcimCablesPartialUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimCablesPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /dcim/cables/{id}/][%d] dcim_cables_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimCablesPartialUpdateDefault) String() string {
	return fmt.Sprintf("[PATCH /dcim/cables/{id}/][%d] dcim_cables_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimCablesPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimCablesPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
