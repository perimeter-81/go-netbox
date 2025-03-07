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

// DcimRegionsUpdateReader is a Reader for the DcimRegionsUpdate structure.
type DcimRegionsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimRegionsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimRegionsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimRegionsUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimRegionsUpdateOK creates a DcimRegionsUpdateOK with default headers values
func NewDcimRegionsUpdateOK() *DcimRegionsUpdateOK {
	return &DcimRegionsUpdateOK{}
}

/*
DcimRegionsUpdateOK describes a response with status code 200, with default header values.

DcimRegionsUpdateOK dcim regions update o k
*/
type DcimRegionsUpdateOK struct {
	Payload *models.Region
}

// IsSuccess returns true when this dcim regions update o k response has a 2xx status code
func (o *DcimRegionsUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim regions update o k response has a 3xx status code
func (o *DcimRegionsUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim regions update o k response has a 4xx status code
func (o *DcimRegionsUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim regions update o k response has a 5xx status code
func (o *DcimRegionsUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim regions update o k response a status code equal to that given
func (o *DcimRegionsUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *DcimRegionsUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/regions/{id}/][%d] dcimRegionsUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimRegionsUpdateOK) String() string {
	return fmt.Sprintf("[PUT /dcim/regions/{id}/][%d] dcimRegionsUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimRegionsUpdateOK) GetPayload() *models.Region {
	return o.Payload
}

func (o *DcimRegionsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Region)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimRegionsUpdateDefault creates a DcimRegionsUpdateDefault with default headers values
func NewDcimRegionsUpdateDefault(code int) *DcimRegionsUpdateDefault {
	return &DcimRegionsUpdateDefault{
		_statusCode: code,
	}
}

/*
DcimRegionsUpdateDefault describes a response with status code -1, with default header values.

DcimRegionsUpdateDefault dcim regions update default
*/
type DcimRegionsUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim regions update default response
func (o *DcimRegionsUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim regions update default response has a 2xx status code
func (o *DcimRegionsUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim regions update default response has a 3xx status code
func (o *DcimRegionsUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim regions update default response has a 4xx status code
func (o *DcimRegionsUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim regions update default response has a 5xx status code
func (o *DcimRegionsUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim regions update default response a status code equal to that given
func (o *DcimRegionsUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimRegionsUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /dcim/regions/{id}/][%d] dcim_regions_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimRegionsUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /dcim/regions/{id}/][%d] dcim_regions_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimRegionsUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimRegionsUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
