// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Feedback feedback
//
// swagger:model Feedback
type Feedback struct {

	// email
	Email string `json:"email,omitempty"`

	// text
	Text string `json:"text,omitempty"`
}

// Validate validates this feedback
func (m *Feedback) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this feedback based on context it is used
func (m *Feedback) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Feedback) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Feedback) UnmarshalBinary(b []byte) error {
	var res Feedback
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
