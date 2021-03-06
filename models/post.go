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

// Post post
//
// swagger:model Post
type Post struct {

	// data
	Data *PostData `json:"data,omitempty"`
}

// Validate validates this post
func (m *Post) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Post) validateData(formats strfmt.Registry) error {
	if swag.IsZero(m.Data) { // not required
		return nil
	}

	if m.Data != nil {
		if err := m.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this post based on the context it is used
func (m *Post) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Post) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if m.Data != nil {
		if err := m.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Post) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Post) UnmarshalBinary(b []byte) error {
	var res Post
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PostData post data
//
// swagger:model PostData
type PostData struct {

	// body
	Body string `json:"body,omitempty"`

	// companies
	Companies *List `json:"companies,omitempty"`

	// contact phone
	ContactPhone string `json:"contactPhone,omitempty"`

	// contacts
	Contacts string `json:"contacts,omitempty"`

	// date
	// Format: date
	Date strfmt.Date `json:"date,omitempty"`

	// duration
	Duration string `json:"duration,omitempty"`

	// event type
	EventType string `json:"eventType,omitempty"`

	// hashtag
	Hashtag *List `json:"hashtag,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// image
	Image string `json:"image,omitempty"`

	// member count
	MemberCount int64 `json:"memberCount,omitempty"`

	// owner
	Owner int64 `json:"owner,omitempty"`

	// place
	Place string `json:"place,omitempty"`

	// price
	Price string `json:"price,omitempty"`

	// social media
	SocialMedia *List `json:"socialMedia,omitempty"`

	// speakers
	Speakers *List `json:"speakers,omitempty"`

	// time end
	TimeEnd string `json:"timeEnd,omitempty"`

	// time start
	TimeStart string `json:"timeStart,omitempty"`

	// title
	Title string `json:"title,omitempty"`

	// translation link
	TranslationLink string `json:"translationLink,omitempty"`

	// youtube link
	YoutubeLink string `json:"youtubeLink,omitempty"`
}

// Validate validates this post data
func (m *PostData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCompanies(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHashtag(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSocialMedia(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSpeakers(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostData) validateCompanies(formats strfmt.Registry) error {
	if swag.IsZero(m.Companies) { // not required
		return nil
	}

	if m.Companies != nil {
		if err := m.Companies.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data" + "." + "companies")
			}
			return err
		}
	}

	return nil
}

func (m *PostData) validateDate(formats strfmt.Registry) error {
	if swag.IsZero(m.Date) { // not required
		return nil
	}

	if err := validate.FormatOf("data"+"."+"date", "body", "date", m.Date.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PostData) validateHashtag(formats strfmt.Registry) error {
	if swag.IsZero(m.Hashtag) { // not required
		return nil
	}

	if m.Hashtag != nil {
		if err := m.Hashtag.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data" + "." + "hashtag")
			}
			return err
		}
	}

	return nil
}

func (m *PostData) validateSocialMedia(formats strfmt.Registry) error {
	if swag.IsZero(m.SocialMedia) { // not required
		return nil
	}

	if m.SocialMedia != nil {
		if err := m.SocialMedia.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data" + "." + "socialMedia")
			}
			return err
		}
	}

	return nil
}

func (m *PostData) validateSpeakers(formats strfmt.Registry) error {
	if swag.IsZero(m.Speakers) { // not required
		return nil
	}

	if m.Speakers != nil {
		if err := m.Speakers.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data" + "." + "speakers")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this post data based on the context it is used
func (m *PostData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCompanies(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHashtag(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSocialMedia(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSpeakers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostData) contextValidateCompanies(ctx context.Context, formats strfmt.Registry) error {

	if m.Companies != nil {
		if err := m.Companies.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data" + "." + "companies")
			}
			return err
		}
	}

	return nil
}

func (m *PostData) contextValidateHashtag(ctx context.Context, formats strfmt.Registry) error {

	if m.Hashtag != nil {
		if err := m.Hashtag.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data" + "." + "hashtag")
			}
			return err
		}
	}

	return nil
}

func (m *PostData) contextValidateSocialMedia(ctx context.Context, formats strfmt.Registry) error {

	if m.SocialMedia != nil {
		if err := m.SocialMedia.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data" + "." + "socialMedia")
			}
			return err
		}
	}

	return nil
}

func (m *PostData) contextValidateSpeakers(ctx context.Context, formats strfmt.Registry) error {

	if m.Speakers != nil {
		if err := m.Speakers.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data" + "." + "speakers")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PostData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostData) UnmarshalBinary(b []byte) error {
	var res PostData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
