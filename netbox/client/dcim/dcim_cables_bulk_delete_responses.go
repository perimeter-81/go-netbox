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

// DcimCablesBulkDeleteReader is a Reader for the DcimCablesBulkDelete structure.
type DcimCablesBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimCablesBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimCablesBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimCablesBulkDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimCablesBulkDeleteNoContent creates a DcimCablesBulkDeleteNoContent with default headers values
func NewDcimCablesBulkDeleteNoContent() *DcimCablesBulkDeleteNoContent {
	return &DcimCablesBulkDeleteNoContent{}
}

/*
DcimCablesBulkDeleteNoContent describes a response with status code 204, with default header values.

DcimCablesBulkDeleteNoContent dcim cables bulk delete no content
*/
type DcimCablesBulkDeleteNoContent struct {
}

// IsSuccess returns true when this dcim cables bulk delete no content response has a 2xx status code
func (o *DcimCablesBulkDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim cables bulk delete no content response has a 3xx status code
func (o *DcimCablesBulkDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim cables bulk delete no content response has a 4xx status code
func (o *DcimCablesBulkDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim cables bulk delete no content response has a 5xx status code
func (o *DcimCablesBulkDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim cables bulk delete no content response a status code equal to that given
func (o *DcimCablesBulkDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *DcimCablesBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/cables/][%d] dcimCablesBulkDeleteNoContent ", 204)
}

func (o *DcimCablesBulkDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /dcim/cables/][%d] dcimCablesBulkDeleteNoContent ", 204)
}

func (o *DcimCablesBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDcimCablesBulkDeleteDefault creates a DcimCablesBulkDeleteDefault with default headers values
func NewDcimCablesBulkDeleteDefault(code int) *DcimCablesBulkDeleteDefault {
	return &DcimCablesBulkDeleteDefault{
		_statusCode: code,
	}
}

/*
DcimCablesBulkDeleteDefault describes a response with status code -1, with default header values.

DcimCablesBulkDeleteDefault dcim cables bulk delete default
*/
type DcimCablesBulkDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim cables bulk delete default response
func (o *DcimCablesBulkDeleteDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim cables bulk delete default response has a 2xx status code
func (o *DcimCablesBulkDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim cables bulk delete default response has a 3xx status code
func (o *DcimCablesBulkDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim cables bulk delete default response has a 4xx status code
func (o *DcimCablesBulkDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim cables bulk delete default response has a 5xx status code
func (o *DcimCablesBulkDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim cables bulk delete default response a status code equal to that given
func (o *DcimCablesBulkDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimCablesBulkDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /dcim/cables/][%d] dcim_cables_bulk_delete default  %+v", o._statusCode, o.Payload)
}

func (o *DcimCablesBulkDeleteDefault) String() string {
	return fmt.Sprintf("[DELETE /dcim/cables/][%d] dcim_cables_bulk_delete default  %+v", o._statusCode, o.Payload)
}

func (o *DcimCablesBulkDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimCablesBulkDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
