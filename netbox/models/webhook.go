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

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Webhook webhook
//
// swagger:model Webhook
type Webhook struct {

	// Additional headers
	//
	// User-supplied HTTP headers to be sent with the request in addition to the HTTP content type. Headers should be defined in the format <code>Name: Value</code>. Jinja2 template processing is supported with the same context as the request body (below).
	AdditionalHeaders string `json:"additional_headers,omitempty"`

	// Body template
	//
	// Jinja2 template for a custom request body. If blank, a JSON object representing the change will be included. Available context data includes: <code>event</code>, <code>model</code>, <code>timestamp</code>, <code>username</code>, <code>request_id</code>, and <code>data</code>.
	BodyTemplate string `json:"body_template,omitempty"`

	// CA File Path
	//
	// The specific CA certificate file to use for SSL verification. Leave blank to use the system defaults.
	// Max Length: 4096
	CaFilePath *string `json:"ca_file_path,omitempty"`

	// Conditions
	//
	// A set of conditions which determine whether the webhook will be generated.
	Conditions interface{} `json:"conditions,omitempty"`

	// content types
	// Required: true
	// Unique: true
	ContentTypes []string `json:"content_types"`

	// Created
	// Read Only: true
	// Format: date-time
	Created *strfmt.DateTime `json:"created,omitempty"`

	// Display
	// Read Only: true
	Display string `json:"display,omitempty"`

	// Enabled
	Enabled bool `json:"enabled,omitempty"`

	// HTTP content type
	//
	// The complete list of official content types is available <a href="https://www.iana.org/assignments/media-types/media-types.xhtml">here</a>.
	// Max Length: 100
	// Min Length: 1
	HTTPContentType string `json:"http_content_type,omitempty"`

	// HTTP method
	// Enum: [GET POST PUT PATCH DELETE]
	HTTPMethod string `json:"http_method,omitempty"`

	// ID
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Last updated
	// Read Only: true
	// Format: date-time
	LastUpdated *strfmt.DateTime `json:"last_updated,omitempty"`

	// Name
	// Required: true
	// Max Length: 150
	// Min Length: 1
	Name *string `json:"name"`

	// URL
	//
	// This URL will be called using the HTTP method defined when the webhook is called. Jinja2 template processing is supported with the same context as the request body.
	// Required: true
	// Max Length: 500
	// Min Length: 1
	PayloadURL *string `json:"payload_url"`

	// Secret
	//
	// When provided, the request will include a 'X-Hook-Signature' header containing a HMAC hex digest of the payload body using the secret as the key. The secret is not transmitted in the request.
	// Max Length: 255
	Secret string `json:"secret,omitempty"`

	// SSL verification
	//
	// Enable SSL certificate verification. Disable with caution!
	SslVerification bool `json:"ssl_verification,omitempty"`

	// Type create
	//
	// Call this webhook when a matching object is created.
	TypeCreate *bool `json:"type_create,omitempty"`

	// Type delete
	//
	// Call this webhook when a matching object is deleted.
	TypeDelete *bool `json:"type_delete,omitempty"`

	// Type update
	//
	// Call this webhook when a matching object is updated.
	TypeUpdate *bool `json:"type_update,omitempty"`

	// Url
	// Read Only: true
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`
}

// Validate validates this webhook
func (m *Webhook) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCaFilePath(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContentTypes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHTTPContentType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHTTPMethod(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePayloadURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecret(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Webhook) validateCaFilePath(formats strfmt.Registry) error {
	if swag.IsZero(m.CaFilePath) { // not required
		return nil
	}

	if err := validate.MaxLength("ca_file_path", "body", *m.CaFilePath, 4096); err != nil {
		return err
	}

	return nil
}

func (m *Webhook) validateContentTypes(formats strfmt.Registry) error {

	if err := validate.Required("content_types", "body", m.ContentTypes); err != nil {
		return err
	}

	if err := validate.UniqueItems("content_types", "body", m.ContentTypes); err != nil {
		return err
	}

	return nil
}

func (m *Webhook) validateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date-time", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Webhook) validateHTTPContentType(formats strfmt.Registry) error {
	if swag.IsZero(m.HTTPContentType) { // not required
		return nil
	}

	if err := validate.MinLength("http_content_type", "body", m.HTTPContentType, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("http_content_type", "body", m.HTTPContentType, 100); err != nil {
		return err
	}

	return nil
}

var webhookTypeHTTPMethodPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["GET","POST","PUT","PATCH","DELETE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		webhookTypeHTTPMethodPropEnum = append(webhookTypeHTTPMethodPropEnum, v)
	}
}

const (

	// WebhookHTTPMethodGET captures enum value "GET"
	WebhookHTTPMethodGET string = "GET"

	// WebhookHTTPMethodPOST captures enum value "POST"
	WebhookHTTPMethodPOST string = "POST"

	// WebhookHTTPMethodPUT captures enum value "PUT"
	WebhookHTTPMethodPUT string = "PUT"

	// WebhookHTTPMethodPATCH captures enum value "PATCH"
	WebhookHTTPMethodPATCH string = "PATCH"

	// WebhookHTTPMethodDELETE captures enum value "DELETE"
	WebhookHTTPMethodDELETE string = "DELETE"
)

// prop value enum
func (m *Webhook) validateHTTPMethodEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, webhookTypeHTTPMethodPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Webhook) validateHTTPMethod(formats strfmt.Registry) error {
	if swag.IsZero(m.HTTPMethod) { // not required
		return nil
	}

	// value enum
	if err := m.validateHTTPMethodEnum("http_method", "body", m.HTTPMethod); err != nil {
		return err
	}

	return nil
}

func (m *Webhook) validateLastUpdated(formats strfmt.Registry) error {
	if swag.IsZero(m.LastUpdated) { // not required
		return nil
	}

	if err := validate.FormatOf("last_updated", "body", "date-time", m.LastUpdated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Webhook) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", *m.Name, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", *m.Name, 150); err != nil {
		return err
	}

	return nil
}

func (m *Webhook) validatePayloadURL(formats strfmt.Registry) error {

	if err := validate.Required("payload_url", "body", m.PayloadURL); err != nil {
		return err
	}

	if err := validate.MinLength("payload_url", "body", *m.PayloadURL, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("payload_url", "body", *m.PayloadURL, 500); err != nil {
		return err
	}

	return nil
}

func (m *Webhook) validateSecret(formats strfmt.Registry) error {
	if swag.IsZero(m.Secret) { // not required
		return nil
	}

	if err := validate.MaxLength("secret", "body", m.Secret, 255); err != nil {
		return err
	}

	return nil
}

func (m *Webhook) validateURL(formats strfmt.Registry) error {
	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this webhook based on the context it is used
func (m *Webhook) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCreated(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDisplay(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLastUpdated(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateURL(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Webhook) contextValidateCreated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "created", "body", m.Created); err != nil {
		return err
	}

	return nil
}

func (m *Webhook) contextValidateDisplay(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "display", "body", string(m.Display)); err != nil {
		return err
	}

	return nil
}

func (m *Webhook) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", int64(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *Webhook) contextValidateLastUpdated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "last_updated", "body", m.LastUpdated); err != nil {
		return err
	}

	return nil
}

func (m *Webhook) contextValidateURL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "url", "body", strfmt.URI(m.URL)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Webhook) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Webhook) UnmarshalBinary(b []byte) error {
	var res Webhook
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
