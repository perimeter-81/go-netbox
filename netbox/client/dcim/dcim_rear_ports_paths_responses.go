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

// DcimRearPortsPathsReader is a Reader for the DcimRearPortsPaths structure.
type DcimRearPortsPathsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimRearPortsPathsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimRearPortsPathsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimRearPortsPathsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimRearPortsPathsOK creates a DcimRearPortsPathsOK with default headers values
func NewDcimRearPortsPathsOK() *DcimRearPortsPathsOK {
	return &DcimRearPortsPathsOK{}
}

/*
DcimRearPortsPathsOK describes a response with status code 200, with default header values.

DcimRearPortsPathsOK dcim rear ports paths o k
*/
type DcimRearPortsPathsOK struct {
	Payload *models.RearPort
}

// IsSuccess returns true when this dcim rear ports paths o k response has a 2xx status code
func (o *DcimRearPortsPathsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim rear ports paths o k response has a 3xx status code
func (o *DcimRearPortsPathsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim rear ports paths o k response has a 4xx status code
func (o *DcimRearPortsPathsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim rear ports paths o k response has a 5xx status code
func (o *DcimRearPortsPathsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim rear ports paths o k response a status code equal to that given
func (o *DcimRearPortsPathsOK) IsCode(code int) bool {
	return code == 200
}

func (o *DcimRearPortsPathsOK) Error() string {
	return fmt.Sprintf("[GET /dcim/rear-ports/{id}/paths/][%d] dcimRearPortsPathsOK  %+v", 200, o.Payload)
}

func (o *DcimRearPortsPathsOK) String() string {
	return fmt.Sprintf("[GET /dcim/rear-ports/{id}/paths/][%d] dcimRearPortsPathsOK  %+v", 200, o.Payload)
}

func (o *DcimRearPortsPathsOK) GetPayload() *models.RearPort {
	return o.Payload
}

func (o *DcimRearPortsPathsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RearPort)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimRearPortsPathsDefault creates a DcimRearPortsPathsDefault with default headers values
func NewDcimRearPortsPathsDefault(code int) *DcimRearPortsPathsDefault {
	return &DcimRearPortsPathsDefault{
		_statusCode: code,
	}
}

/*
DcimRearPortsPathsDefault describes a response with status code -1, with default header values.

DcimRearPortsPathsDefault dcim rear ports paths default
*/
type DcimRearPortsPathsDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim rear ports paths default response
func (o *DcimRearPortsPathsDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim rear ports paths default response has a 2xx status code
func (o *DcimRearPortsPathsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim rear ports paths default response has a 3xx status code
func (o *DcimRearPortsPathsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim rear ports paths default response has a 4xx status code
func (o *DcimRearPortsPathsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim rear ports paths default response has a 5xx status code
func (o *DcimRearPortsPathsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim rear ports paths default response a status code equal to that given
func (o *DcimRearPortsPathsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimRearPortsPathsDefault) Error() string {
	return fmt.Sprintf("[GET /dcim/rear-ports/{id}/paths/][%d] dcim_rear-ports_paths default  %+v", o._statusCode, o.Payload)
}

func (o *DcimRearPortsPathsDefault) String() string {
	return fmt.Sprintf("[GET /dcim/rear-ports/{id}/paths/][%d] dcim_rear-ports_paths default  %+v", o._statusCode, o.Payload)
}

func (o *DcimRearPortsPathsDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimRearPortsPathsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
