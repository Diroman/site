// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"lawSite/models"
)

// GetUserByNameOKCode is the HTTP code returned for type GetUserByNameOK
const GetUserByNameOKCode int = 200

/*GetUserByNameOK successful operation

swagger:response getUserByNameOK
*/
type GetUserByNameOK struct {

	/*
	  In: Body
	*/
	Payload *models.User `json:"body,omitempty"`
}

// NewGetUserByNameOK creates GetUserByNameOK with default headers values
func NewGetUserByNameOK() *GetUserByNameOK {

	return &GetUserByNameOK{}
}

// WithPayload adds the payload to the get user by name o k response
func (o *GetUserByNameOK) WithPayload(payload *models.User) *GetUserByNameOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get user by name o k response
func (o *GetUserByNameOK) SetPayload(payload *models.User) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUserByNameOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetUserByNameBadRequestCode is the HTTP code returned for type GetUserByNameBadRequest
const GetUserByNameBadRequestCode int = 400

/*GetUserByNameBadRequest Invalid username supplied

swagger:response getUserByNameBadRequest
*/
type GetUserByNameBadRequest struct {
}

// NewGetUserByNameBadRequest creates GetUserByNameBadRequest with default headers values
func NewGetUserByNameBadRequest() *GetUserByNameBadRequest {

	return &GetUserByNameBadRequest{}
}

// WriteResponse to the client
func (o *GetUserByNameBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// GetUserByNameNotFoundCode is the HTTP code returned for type GetUserByNameNotFound
const GetUserByNameNotFoundCode int = 404

/*GetUserByNameNotFound User not found

swagger:response getUserByNameNotFound
*/
type GetUserByNameNotFound struct {
}

// NewGetUserByNameNotFound creates GetUserByNameNotFound with default headers values
func NewGetUserByNameNotFound() *GetUserByNameNotFound {

	return &GetUserByNameNotFound{}
}

// WriteResponse to the client
func (o *GetUserByNameNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// GetUserByNameInternalServerErrorCode is the HTTP code returned for type GetUserByNameInternalServerError
const GetUserByNameInternalServerErrorCode int = 500

/*GetUserByNameInternalServerError Server error

swagger:response getUserByNameInternalServerError
*/
type GetUserByNameInternalServerError struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewGetUserByNameInternalServerError creates GetUserByNameInternalServerError with default headers values
func NewGetUserByNameInternalServerError() *GetUserByNameInternalServerError {

	return &GetUserByNameInternalServerError{}
}

// WithPayload adds the payload to the get user by name internal server error response
func (o *GetUserByNameInternalServerError) WithPayload(payload string) *GetUserByNameInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get user by name internal server error response
func (o *GetUserByNameInternalServerError) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUserByNameInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}