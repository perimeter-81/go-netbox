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

package virtualization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/perimeter-81/go-netbox/netbox/models"
)

// VirtualizationClusterTypesBulkUpdateReader is a Reader for the VirtualizationClusterTypesBulkUpdate structure.
type VirtualizationClusterTypesBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VirtualizationClusterTypesBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVirtualizationClusterTypesBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewVirtualizationClusterTypesBulkUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewVirtualizationClusterTypesBulkUpdateOK creates a VirtualizationClusterTypesBulkUpdateOK with default headers values
func NewVirtualizationClusterTypesBulkUpdateOK() *VirtualizationClusterTypesBulkUpdateOK {
	return &VirtualizationClusterTypesBulkUpdateOK{}
}

/*
VirtualizationClusterTypesBulkUpdateOK describes a response with status code 200, with default header values.

VirtualizationClusterTypesBulkUpdateOK virtualization cluster types bulk update o k
*/
type VirtualizationClusterTypesBulkUpdateOK struct {
	Payload *models.ClusterType
}

// IsSuccess returns true when this virtualization cluster types bulk update o k response has a 2xx status code
func (o *VirtualizationClusterTypesBulkUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this virtualization cluster types bulk update o k response has a 3xx status code
func (o *VirtualizationClusterTypesBulkUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this virtualization cluster types bulk update o k response has a 4xx status code
func (o *VirtualizationClusterTypesBulkUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this virtualization cluster types bulk update o k response has a 5xx status code
func (o *VirtualizationClusterTypesBulkUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this virtualization cluster types bulk update o k response a status code equal to that given
func (o *VirtualizationClusterTypesBulkUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *VirtualizationClusterTypesBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /virtualization/cluster-types/][%d] virtualizationClusterTypesBulkUpdateOK  %+v", 200, o.Payload)
}

func (o *VirtualizationClusterTypesBulkUpdateOK) String() string {
	return fmt.Sprintf("[PUT /virtualization/cluster-types/][%d] virtualizationClusterTypesBulkUpdateOK  %+v", 200, o.Payload)
}

func (o *VirtualizationClusterTypesBulkUpdateOK) GetPayload() *models.ClusterType {
	return o.Payload
}

func (o *VirtualizationClusterTypesBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterType)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVirtualizationClusterTypesBulkUpdateDefault creates a VirtualizationClusterTypesBulkUpdateDefault with default headers values
func NewVirtualizationClusterTypesBulkUpdateDefault(code int) *VirtualizationClusterTypesBulkUpdateDefault {
	return &VirtualizationClusterTypesBulkUpdateDefault{
		_statusCode: code,
	}
}

/*
VirtualizationClusterTypesBulkUpdateDefault describes a response with status code -1, with default header values.

VirtualizationClusterTypesBulkUpdateDefault virtualization cluster types bulk update default
*/
type VirtualizationClusterTypesBulkUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the virtualization cluster types bulk update default response
func (o *VirtualizationClusterTypesBulkUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this virtualization cluster types bulk update default response has a 2xx status code
func (o *VirtualizationClusterTypesBulkUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this virtualization cluster types bulk update default response has a 3xx status code
func (o *VirtualizationClusterTypesBulkUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this virtualization cluster types bulk update default response has a 4xx status code
func (o *VirtualizationClusterTypesBulkUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this virtualization cluster types bulk update default response has a 5xx status code
func (o *VirtualizationClusterTypesBulkUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this virtualization cluster types bulk update default response a status code equal to that given
func (o *VirtualizationClusterTypesBulkUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *VirtualizationClusterTypesBulkUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /virtualization/cluster-types/][%d] virtualization_cluster-types_bulk_update default  %+v", o._statusCode, o.Payload)
}

func (o *VirtualizationClusterTypesBulkUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /virtualization/cluster-types/][%d] virtualization_cluster-types_bulk_update default  %+v", o._statusCode, o.Payload)
}

func (o *VirtualizationClusterTypesBulkUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *VirtualizationClusterTypesBulkUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
