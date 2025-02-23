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

// DcimInventoryItemsDeleteReader is a Reader for the DcimInventoryItemsDelete structure.
type DcimInventoryItemsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimInventoryItemsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimInventoryItemsDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimInventoryItemsDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimInventoryItemsDeleteNoContent creates a DcimInventoryItemsDeleteNoContent with default headers values
func NewDcimInventoryItemsDeleteNoContent() *DcimInventoryItemsDeleteNoContent {
	return &DcimInventoryItemsDeleteNoContent{}
}

/*
DcimInventoryItemsDeleteNoContent describes a response with status code 204, with default header values.

DcimInventoryItemsDeleteNoContent dcim inventory items delete no content
*/
type DcimInventoryItemsDeleteNoContent struct {
}

// IsSuccess returns true when this dcim inventory items delete no content response has a 2xx status code
func (o *DcimInventoryItemsDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim inventory items delete no content response has a 3xx status code
func (o *DcimInventoryItemsDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim inventory items delete no content response has a 4xx status code
func (o *DcimInventoryItemsDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim inventory items delete no content response has a 5xx status code
func (o *DcimInventoryItemsDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim inventory items delete no content response a status code equal to that given
func (o *DcimInventoryItemsDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *DcimInventoryItemsDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/inventory-items/{id}/][%d] dcimInventoryItemsDeleteNoContent ", 204)
}

func (o *DcimInventoryItemsDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /dcim/inventory-items/{id}/][%d] dcimInventoryItemsDeleteNoContent ", 204)
}

func (o *DcimInventoryItemsDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDcimInventoryItemsDeleteDefault creates a DcimInventoryItemsDeleteDefault with default headers values
func NewDcimInventoryItemsDeleteDefault(code int) *DcimInventoryItemsDeleteDefault {
	return &DcimInventoryItemsDeleteDefault{
		_statusCode: code,
	}
}

/*
DcimInventoryItemsDeleteDefault describes a response with status code -1, with default header values.

DcimInventoryItemsDeleteDefault dcim inventory items delete default
*/
type DcimInventoryItemsDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim inventory items delete default response
func (o *DcimInventoryItemsDeleteDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim inventory items delete default response has a 2xx status code
func (o *DcimInventoryItemsDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim inventory items delete default response has a 3xx status code
func (o *DcimInventoryItemsDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim inventory items delete default response has a 4xx status code
func (o *DcimInventoryItemsDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim inventory items delete default response has a 5xx status code
func (o *DcimInventoryItemsDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim inventory items delete default response a status code equal to that given
func (o *DcimInventoryItemsDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimInventoryItemsDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /dcim/inventory-items/{id}/][%d] dcim_inventory-items_delete default  %+v", o._statusCode, o.Payload)
}

func (o *DcimInventoryItemsDeleteDefault) String() string {
	return fmt.Sprintf("[DELETE /dcim/inventory-items/{id}/][%d] dcim_inventory-items_delete default  %+v", o._statusCode, o.Payload)
}

func (o *DcimInventoryItemsDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimInventoryItemsDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
