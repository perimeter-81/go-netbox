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

package circuits

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/perimeter-81/go-netbox/netbox/models"
)

// CircuitsProviderNetworksPartialUpdateReader is a Reader for the CircuitsProviderNetworksPartialUpdate structure.
type CircuitsProviderNetworksPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CircuitsProviderNetworksPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCircuitsProviderNetworksPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCircuitsProviderNetworksPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCircuitsProviderNetworksPartialUpdateOK creates a CircuitsProviderNetworksPartialUpdateOK with default headers values
func NewCircuitsProviderNetworksPartialUpdateOK() *CircuitsProviderNetworksPartialUpdateOK {
	return &CircuitsProviderNetworksPartialUpdateOK{}
}

/*
CircuitsProviderNetworksPartialUpdateOK describes a response with status code 200, with default header values.

CircuitsProviderNetworksPartialUpdateOK circuits provider networks partial update o k
*/
type CircuitsProviderNetworksPartialUpdateOK struct {
	Payload *models.ProviderNetwork
}

// IsSuccess returns true when this circuits provider networks partial update o k response has a 2xx status code
func (o *CircuitsProviderNetworksPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this circuits provider networks partial update o k response has a 3xx status code
func (o *CircuitsProviderNetworksPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this circuits provider networks partial update o k response has a 4xx status code
func (o *CircuitsProviderNetworksPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this circuits provider networks partial update o k response has a 5xx status code
func (o *CircuitsProviderNetworksPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this circuits provider networks partial update o k response a status code equal to that given
func (o *CircuitsProviderNetworksPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *CircuitsProviderNetworksPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /circuits/provider-networks/{id}/][%d] circuitsProviderNetworksPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *CircuitsProviderNetworksPartialUpdateOK) String() string {
	return fmt.Sprintf("[PATCH /circuits/provider-networks/{id}/][%d] circuitsProviderNetworksPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *CircuitsProviderNetworksPartialUpdateOK) GetPayload() *models.ProviderNetwork {
	return o.Payload
}

func (o *CircuitsProviderNetworksPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProviderNetwork)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCircuitsProviderNetworksPartialUpdateDefault creates a CircuitsProviderNetworksPartialUpdateDefault with default headers values
func NewCircuitsProviderNetworksPartialUpdateDefault(code int) *CircuitsProviderNetworksPartialUpdateDefault {
	return &CircuitsProviderNetworksPartialUpdateDefault{
		_statusCode: code,
	}
}

/*
CircuitsProviderNetworksPartialUpdateDefault describes a response with status code -1, with default header values.

CircuitsProviderNetworksPartialUpdateDefault circuits provider networks partial update default
*/
type CircuitsProviderNetworksPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the circuits provider networks partial update default response
func (o *CircuitsProviderNetworksPartialUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this circuits provider networks partial update default response has a 2xx status code
func (o *CircuitsProviderNetworksPartialUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this circuits provider networks partial update default response has a 3xx status code
func (o *CircuitsProviderNetworksPartialUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this circuits provider networks partial update default response has a 4xx status code
func (o *CircuitsProviderNetworksPartialUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this circuits provider networks partial update default response has a 5xx status code
func (o *CircuitsProviderNetworksPartialUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this circuits provider networks partial update default response a status code equal to that given
func (o *CircuitsProviderNetworksPartialUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *CircuitsProviderNetworksPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /circuits/provider-networks/{id}/][%d] circuits_provider-networks_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *CircuitsProviderNetworksPartialUpdateDefault) String() string {
	return fmt.Sprintf("[PATCH /circuits/provider-networks/{id}/][%d] circuits_provider-networks_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *CircuitsProviderNetworksPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *CircuitsProviderNetworksPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
