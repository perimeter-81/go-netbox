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
)

// DcimPowerOutletsBulkDeleteReader is a Reader for the DcimPowerOutletsBulkDelete structure.
type DcimPowerOutletsBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPowerOutletsBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimPowerOutletsBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimPowerOutletsBulkDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimPowerOutletsBulkDeleteNoContent creates a DcimPowerOutletsBulkDeleteNoContent with default headers values
func NewDcimPowerOutletsBulkDeleteNoContent() *DcimPowerOutletsBulkDeleteNoContent {
	return &DcimPowerOutletsBulkDeleteNoContent{}
}

/*
DcimPowerOutletsBulkDeleteNoContent describes a response with status code 204, with default header values.

DcimPowerOutletsBulkDeleteNoContent dcim power outlets bulk delete no content
*/
type DcimPowerOutletsBulkDeleteNoContent struct {
}

// IsSuccess returns true when this dcim power outlets bulk delete no content response has a 2xx status code
func (o *DcimPowerOutletsBulkDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim power outlets bulk delete no content response has a 3xx status code
func (o *DcimPowerOutletsBulkDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim power outlets bulk delete no content response has a 4xx status code
func (o *DcimPowerOutletsBulkDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim power outlets bulk delete no content response has a 5xx status code
func (o *DcimPowerOutletsBulkDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim power outlets bulk delete no content response a status code equal to that given
func (o *DcimPowerOutletsBulkDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *DcimPowerOutletsBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/power-outlets/][%d] dcimPowerOutletsBulkDeleteNoContent ", 204)
}

func (o *DcimPowerOutletsBulkDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /dcim/power-outlets/][%d] dcimPowerOutletsBulkDeleteNoContent ", 204)
}

func (o *DcimPowerOutletsBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDcimPowerOutletsBulkDeleteDefault creates a DcimPowerOutletsBulkDeleteDefault with default headers values
func NewDcimPowerOutletsBulkDeleteDefault(code int) *DcimPowerOutletsBulkDeleteDefault {
	return &DcimPowerOutletsBulkDeleteDefault{
		_statusCode: code,
	}
}

/*
DcimPowerOutletsBulkDeleteDefault describes a response with status code -1, with default header values.

DcimPowerOutletsBulkDeleteDefault dcim power outlets bulk delete default
*/
type DcimPowerOutletsBulkDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim power outlets bulk delete default response
func (o *DcimPowerOutletsBulkDeleteDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim power outlets bulk delete default response has a 2xx status code
func (o *DcimPowerOutletsBulkDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim power outlets bulk delete default response has a 3xx status code
func (o *DcimPowerOutletsBulkDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim power outlets bulk delete default response has a 4xx status code
func (o *DcimPowerOutletsBulkDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim power outlets bulk delete default response has a 5xx status code
func (o *DcimPowerOutletsBulkDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim power outlets bulk delete default response a status code equal to that given
func (o *DcimPowerOutletsBulkDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimPowerOutletsBulkDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /dcim/power-outlets/][%d] dcim_power-outlets_bulk_delete default  %+v", o._statusCode, o.Payload)
}

func (o *DcimPowerOutletsBulkDeleteDefault) String() string {
	return fmt.Sprintf("[DELETE /dcim/power-outlets/][%d] dcim_power-outlets_bulk_delete default  %+v", o._statusCode, o.Payload)
}

func (o *DcimPowerOutletsBulkDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimPowerOutletsBulkDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
