// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ID Id
//
// swagger:model Id
type ID struct {

	// value
	Value int64 `json:"value,omitempty"`
}

// Validate validates this Id
func (m *ID) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this Id based on context it is used
func (m *ID) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ID) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ID) UnmarshalBinary(b []byte) error {
	var res ID
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}