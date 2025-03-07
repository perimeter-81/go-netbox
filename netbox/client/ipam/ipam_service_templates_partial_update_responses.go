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

package ipam

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/perimeter-81/go-netbox/netbox/models"
)

// IpamServiceTemplatesPartialUpdateReader is a Reader for the IpamServiceTemplatesPartialUpdate structure.
type IpamServiceTemplatesPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamServiceTemplatesPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamServiceTemplatesPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIpamServiceTemplatesPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpamServiceTemplatesPartialUpdateOK creates a IpamServiceTemplatesPartialUpdateOK with default headers values
func NewIpamServiceTemplatesPartialUpdateOK() *IpamServiceTemplatesPartialUpdateOK {
	return &IpamServiceTemplatesPartialUpdateOK{}
}

/*
IpamServiceTemplatesPartialUpdateOK describes a response with status code 200, with default header values.

IpamServiceTemplatesPartialUpdateOK ipam service templates partial update o k
*/
type IpamServiceTemplatesPartialUpdateOK struct {
	Payload *models.ServiceTemplate
}

// IsSuccess returns true when this ipam service templates partial update o k response has a 2xx status code
func (o *IpamServiceTemplatesPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam service templates partial update o k response has a 3xx status code
func (o *IpamServiceTemplatesPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam service templates partial update o k response has a 4xx status code
func (o *IpamServiceTemplatesPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam service templates partial update o k response has a 5xx status code
func (o *IpamServiceTemplatesPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam service templates partial update o k response a status code equal to that given
func (o *IpamServiceTemplatesPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *IpamServiceTemplatesPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /ipam/service-templates/{id}/][%d] ipamServiceTemplatesPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *IpamServiceTemplatesPartialUpdateOK) String() string {
	return fmt.Sprintf("[PATCH /ipam/service-templates/{id}/][%d] ipamServiceTemplatesPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *IpamServiceTemplatesPartialUpdateOK) GetPayload() *models.ServiceTemplate {
	return o.Payload
}

func (o *IpamServiceTemplatesPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIpamServiceTemplatesPartialUpdateDefault creates a IpamServiceTemplatesPartialUpdateDefault with default headers values
func NewIpamServiceTemplatesPartialUpdateDefault(code int) *IpamServiceTemplatesPartialUpdateDefault {
	return &IpamServiceTemplatesPartialUpdateDefault{
		_statusCode: code,
	}
}

/*
IpamServiceTemplatesPartialUpdateDefault describes a response with status code -1, with default header values.

IpamServiceTemplatesPartialUpdateDefault ipam service templates partial update default
*/
type IpamServiceTemplatesPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the ipam service templates partial update default response
func (o *IpamServiceTemplatesPartialUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this ipam service templates partial update default response has a 2xx status code
func (o *IpamServiceTemplatesPartialUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ipam service templates partial update default response has a 3xx status code
func (o *IpamServiceTemplatesPartialUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ipam service templates partial update default response has a 4xx status code
func (o *IpamServiceTemplatesPartialUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ipam service templates partial update default response has a 5xx status code
func (o *IpamServiceTemplatesPartialUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ipam service templates partial update default response a status code equal to that given
func (o *IpamServiceTemplatesPartialUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *IpamServiceTemplatesPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /ipam/service-templates/{id}/][%d] ipam_service-templates_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *IpamServiceTemplatesPartialUpdateDefault) String() string {
	return fmt.Sprintf("[PATCH /ipam/service-templates/{id}/][%d] ipam_service-templates_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *IpamServiceTemplatesPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamServiceTemplatesPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
