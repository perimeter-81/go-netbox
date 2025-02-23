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

// DcimSiteGroupsBulkUpdateReader is a Reader for the DcimSiteGroupsBulkUpdate structure.
type DcimSiteGroupsBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimSiteGroupsBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimSiteGroupsBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimSiteGroupsBulkUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimSiteGroupsBulkUpdateOK creates a DcimSiteGroupsBulkUpdateOK with default headers values
func NewDcimSiteGroupsBulkUpdateOK() *DcimSiteGroupsBulkUpdateOK {
	return &DcimSiteGroupsBulkUpdateOK{}
}

/*
DcimSiteGroupsBulkUpdateOK describes a response with status code 200, with default header values.

DcimSiteGroupsBulkUpdateOK dcim site groups bulk update o k
*/
type DcimSiteGroupsBulkUpdateOK struct {
	Payload *models.SiteGroup
}

// IsSuccess returns true when this dcim site groups bulk update o k response has a 2xx status code
func (o *DcimSiteGroupsBulkUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim site groups bulk update o k response has a 3xx status code
func (o *DcimSiteGroupsBulkUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim site groups bulk update o k response has a 4xx status code
func (o *DcimSiteGroupsBulkUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim site groups bulk update o k response has a 5xx status code
func (o *DcimSiteGroupsBulkUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim site groups bulk update o k response a status code equal to that given
func (o *DcimSiteGroupsBulkUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *DcimSiteGroupsBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/site-groups/][%d] dcimSiteGroupsBulkUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimSiteGroupsBulkUpdateOK) String() string {
	return fmt.Sprintf("[PUT /dcim/site-groups/][%d] dcimSiteGroupsBulkUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimSiteGroupsBulkUpdateOK) GetPayload() *models.SiteGroup {
	return o.Payload
}

func (o *DcimSiteGroupsBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SiteGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimSiteGroupsBulkUpdateDefault creates a DcimSiteGroupsBulkUpdateDefault with default headers values
func NewDcimSiteGroupsBulkUpdateDefault(code int) *DcimSiteGroupsBulkUpdateDefault {
	return &DcimSiteGroupsBulkUpdateDefault{
		_statusCode: code,
	}
}

/*
DcimSiteGroupsBulkUpdateDefault describes a response with status code -1, with default header values.

DcimSiteGroupsBulkUpdateDefault dcim site groups bulk update default
*/
type DcimSiteGroupsBulkUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim site groups bulk update default response
func (o *DcimSiteGroupsBulkUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim site groups bulk update default response has a 2xx status code
func (o *DcimSiteGroupsBulkUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim site groups bulk update default response has a 3xx status code
func (o *DcimSiteGroupsBulkUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim site groups bulk update default response has a 4xx status code
func (o *DcimSiteGroupsBulkUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim site groups bulk update default response has a 5xx status code
func (o *DcimSiteGroupsBulkUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim site groups bulk update default response a status code equal to that given
func (o *DcimSiteGroupsBulkUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimSiteGroupsBulkUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /dcim/site-groups/][%d] dcim_site-groups_bulk_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimSiteGroupsBulkUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /dcim/site-groups/][%d] dcim_site-groups_bulk_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimSiteGroupsBulkUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimSiteGroupsBulkUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
