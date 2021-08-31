// Code generated by go-swagger; DO NOT EDIT.

package auth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"lawSite/models"
)

// UpdatePasswordOKCode is the HTTP code returned for type UpdatePasswordOK
const UpdatePasswordOKCode int = 200

/*UpdatePasswordOK successful operation

swagger:response updatePasswordOK
*/
type UpdatePasswordOK struct {
}

// NewUpdatePasswordOK creates UpdatePasswordOK with default headers values
func NewUpdatePasswordOK() *UpdatePasswordOK {

	return &UpdatePasswordOK{}
}

// WriteResponse to the client
func (o *UpdatePasswordOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// UpdatePasswordBadRequestCode is the HTTP code returned for type UpdatePasswordBadRequest
const UpdatePasswordBadRequestCode int = 400

/*UpdatePasswordBadRequest User not found

swagger:response updatePasswordBadRequest
*/
type UpdatePasswordBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUpdatePasswordBadRequest creates UpdatePasswordBadRequest with default headers values
func NewUpdatePasswordBadRequest() *UpdatePasswordBadRequest {

	return &UpdatePasswordBadRequest{}
}

// WithPayload adds the payload to the update password bad request response
func (o *UpdatePasswordBadRequest) WithPayload(payload *models.Error) *UpdatePasswordBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update password bad request response
func (o *UpdatePasswordBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdatePasswordBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdatePasswordInternalServerErrorCode is the HTTP code returned for type UpdatePasswordInternalServerError
const UpdatePasswordInternalServerErrorCode int = 500

/*UpdatePasswordInternalServerError Server error

swagger:response updatePasswordInternalServerError
*/
type UpdatePasswordInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUpdatePasswordInternalServerError creates UpdatePasswordInternalServerError with default headers values
func NewUpdatePasswordInternalServerError() *UpdatePasswordInternalServerError {

	return &UpdatePasswordInternalServerError{}
}

// WithPayload adds the payload to the update password internal server error response
func (o *UpdatePasswordInternalServerError) WithPayload(payload *models.Error) *UpdatePasswordInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update password internal server error response
func (o *UpdatePasswordInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdatePasswordInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}