// Code generated by go-swagger; DO NOT EDIT.

package post

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"lawSite/models"
)

// CheckPostInFavoriteOKCode is the HTTP code returned for type CheckPostInFavoriteOK
const CheckPostInFavoriteOKCode int = 200

/*CheckPostInFavoriteOK successful operation

swagger:response checkPostInFavoriteOK
*/
type CheckPostInFavoriteOK struct {

	/*
	  In: Body
	*/
	Payload *models.Bool `json:"body,omitempty"`
}

// NewCheckPostInFavoriteOK creates CheckPostInFavoriteOK with default headers values
func NewCheckPostInFavoriteOK() *CheckPostInFavoriteOK {

	return &CheckPostInFavoriteOK{}
}

// WithPayload adds the payload to the check post in favorite o k response
func (o *CheckPostInFavoriteOK) WithPayload(payload *models.Bool) *CheckPostInFavoriteOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the check post in favorite o k response
func (o *CheckPostInFavoriteOK) SetPayload(payload *models.Bool) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CheckPostInFavoriteOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CheckPostInFavoriteBadRequestCode is the HTTP code returned for type CheckPostInFavoriteBadRequest
const CheckPostInFavoriteBadRequestCode int = 400

/*CheckPostInFavoriteBadRequest Bad request

swagger:response checkPostInFavoriteBadRequest
*/
type CheckPostInFavoriteBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCheckPostInFavoriteBadRequest creates CheckPostInFavoriteBadRequest with default headers values
func NewCheckPostInFavoriteBadRequest() *CheckPostInFavoriteBadRequest {

	return &CheckPostInFavoriteBadRequest{}
}

// WithPayload adds the payload to the check post in favorite bad request response
func (o *CheckPostInFavoriteBadRequest) WithPayload(payload *models.Error) *CheckPostInFavoriteBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the check post in favorite bad request response
func (o *CheckPostInFavoriteBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CheckPostInFavoriteBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CheckPostInFavoriteUnauthorizedCode is the HTTP code returned for type CheckPostInFavoriteUnauthorized
const CheckPostInFavoriteUnauthorizedCode int = 401

/*CheckPostInFavoriteUnauthorized not valid token

swagger:response checkPostInFavoriteUnauthorized
*/
type CheckPostInFavoriteUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCheckPostInFavoriteUnauthorized creates CheckPostInFavoriteUnauthorized with default headers values
func NewCheckPostInFavoriteUnauthorized() *CheckPostInFavoriteUnauthorized {

	return &CheckPostInFavoriteUnauthorized{}
}

// WithPayload adds the payload to the check post in favorite unauthorized response
func (o *CheckPostInFavoriteUnauthorized) WithPayload(payload *models.Error) *CheckPostInFavoriteUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the check post in favorite unauthorized response
func (o *CheckPostInFavoriteUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CheckPostInFavoriteUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CheckPostInFavoriteInternalServerErrorCode is the HTTP code returned for type CheckPostInFavoriteInternalServerError
const CheckPostInFavoriteInternalServerErrorCode int = 500

/*CheckPostInFavoriteInternalServerError Server error

swagger:response checkPostInFavoriteInternalServerError
*/
type CheckPostInFavoriteInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCheckPostInFavoriteInternalServerError creates CheckPostInFavoriteInternalServerError with default headers values
func NewCheckPostInFavoriteInternalServerError() *CheckPostInFavoriteInternalServerError {

	return &CheckPostInFavoriteInternalServerError{}
}

// WithPayload adds the payload to the check post in favorite internal server error response
func (o *CheckPostInFavoriteInternalServerError) WithPayload(payload *models.Error) *CheckPostInFavoriteInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the check post in favorite internal server error response
func (o *CheckPostInFavoriteInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CheckPostInFavoriteInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
