// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PostShort post short
//
// swagger:model PostShort
type PostShort struct {

	// date
	// Format: date-time
	Date strfmt.DateTime `json:"date,omitempty"`

	// duration
	Duration string `json:"duration,omitempty"`

	// event type
	EventType string `json:"eventType,omitempty"`

	// hashtag
	Hashtag []string `json:"hashtag"`

	// image
	Image string `json:"image,omitempty"`

	// place
	Place string `json:"place,omitempty"`

	// speakers
	Speakers []string `json:"speakers"`

	// text
	Text string `json:"text,omitempty"`

	// title
	Title string `json:"title,omitempty"`

	// translation link
	TranslationLink string `json:"translationLink,omitempty"`
}

// Validate validates this post short
func (m *PostShort) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostShort) validateDate(formats strfmt.Registry) error {
	if swag.IsZero(m.Date) { // not required
		return nil
	}

	if err := validate.FormatOf("date", "body", "date-time", m.Date.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this post short based on context it is used
func (m *PostShort) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PostShort) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostShort) UnmarshalBinary(b []byte) error {
	var res PostShort
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
