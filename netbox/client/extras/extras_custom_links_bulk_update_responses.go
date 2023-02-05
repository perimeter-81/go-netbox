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

// ExtrasCustomLinksBulkUpdateReader is a Reader for the ExtrasCustomLinksBulkUpdate structure.
type ExtrasCustomLinksBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasCustomLinksBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasCustomLinksBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewExtrasCustomLinksBulkUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewExtrasCustomLinksBulkUpdateOK creates a ExtrasCustomLinksBulkUpdateOK with default headers values
func NewExtrasCustomLinksBulkUpdateOK() *ExtrasCustomLinksBulkUpdateOK {
	return &ExtrasCustomLinksBulkUpdateOK{}
}

/*
ExtrasCustomLinksBulkUpdateOK describes a response with status code 200, with default header values.

ExtrasCustomLinksBulkUpdateOK extras custom links bulk update o k
*/
type ExtrasCustomLinksBulkUpdateOK struct {
	Payload *models.CustomLink
}

// IsSuccess returns true when this extras custom links bulk update o k response has a 2xx status code
func (o *ExtrasCustomLinksBulkUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this extras custom links bulk update o k response has a 3xx status code
func (o *ExtrasCustomLinksBulkUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extras custom links bulk update o k response has a 4xx status code
func (o *ExtrasCustomLinksBulkUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this extras custom links bulk update o k response has a 5xx status code
func (o *ExtrasCustomLinksBulkUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this extras custom links bulk update o k response a status code equal to that given
func (o *ExtrasCustomLinksBulkUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *ExtrasCustomLinksBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /extras/custom-links/][%d] extrasCustomLinksBulkUpdateOK  %+v", 200, o.Payload)
}

func (o *ExtrasCustomLinksBulkUpdateOK) String() string {
	return fmt.Sprintf("[PUT /extras/custom-links/][%d] extrasCustomLinksBulkUpdateOK  %+v", 200, o.Payload)
}

func (o *ExtrasCustomLinksBulkUpdateOK) GetPayload() *models.CustomLink {
	return o.Payload
}

func (o *ExtrasCustomLinksBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CustomLink)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExtrasCustomLinksBulkUpdateDefault creates a ExtrasCustomLinksBulkUpdateDefault with default headers values
func NewExtrasCustomLinksBulkUpdateDefault(code int) *ExtrasCustomLinksBulkUpdateDefault {
	return &ExtrasCustomLinksBulkUpdateDefault{
		_statusCode: code,
	}
}

/*
ExtrasCustomLinksBulkUpdateDefault describes a response with status code -1, with default header values.

ExtrasCustomLinksBulkUpdateDefault extras custom links bulk update default
*/
type ExtrasCustomLinksBulkUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the extras custom links bulk update default response
func (o *ExtrasCustomLinksBulkUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this extras custom links bulk update default response has a 2xx status code
func (o *ExtrasCustomLinksBulkUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this extras custom links bulk update default response has a 3xx status code
func (o *ExtrasCustomLinksBulkUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this extras custom links bulk update default response has a 4xx status code
func (o *ExtrasCustomLinksBulkUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this extras custom links bulk update default response has a 5xx status code
func (o *ExtrasCustomLinksBulkUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this extras custom links bulk update default response a status code equal to that given
func (o *ExtrasCustomLinksBulkUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ExtrasCustomLinksBulkUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /extras/custom-links/][%d] extras_custom-links_bulk_update default  %+v", o._statusCode, o.Payload)
}

func (o *ExtrasCustomLinksBulkUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /extras/custom-links/][%d] extras_custom-links_bulk_update default  %+v", o._statusCode, o.Payload)
}

func (o *ExtrasCustomLinksBulkUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *ExtrasCustomLinksBulkUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
