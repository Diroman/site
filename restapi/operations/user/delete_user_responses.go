// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// DeleteUserOKCode is the HTTP code returned for type DeleteUserOK
const DeleteUserOKCode int = 200

/*DeleteUserOK Success delete

swagger:response deleteUserOK
*/
type DeleteUserOK struct {
}

// NewDeleteUserOK creates DeleteUserOK with default headers values
func NewDeleteUserOK() *DeleteUserOK {

	return &DeleteUserOK{}
}

// WriteResponse to the client
func (o *DeleteUserOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// DeleteUserBadRequestCode is the HTTP code returned for type DeleteUserBadRequest
const DeleteUserBadRequestCode int = 400

/*DeleteUserBadRequest Invalid username supplied

swagger:response deleteUserBadRequest
*/
type DeleteUserBadRequest struct {
}

// NewDeleteUserBadRequest creates DeleteUserBadRequest with default headers values
func NewDeleteUserBadRequest() *DeleteUserBadRequest {

	return &DeleteUserBadRequest{}
}

// WriteResponse to the client
func (o *DeleteUserBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// DeleteUserNotFoundCode is the HTTP code returned for type DeleteUserNotFound
const DeleteUserNotFoundCode int = 404

/*DeleteUserNotFound User not found

swagger:response deleteUserNotFound
*/
type DeleteUserNotFound struct {
}

// NewDeleteUserNotFound creates DeleteUserNotFound with default headers values
func NewDeleteUserNotFound() *DeleteUserNotFound {

	return &DeleteUserNotFound{}
}

// WriteResponse to the client
func (o *DeleteUserNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// DeleteUserInternalServerErrorCode is the HTTP code returned for type DeleteUserInternalServerError
const DeleteUserInternalServerErrorCode int = 500

/*DeleteUserInternalServerError Server error

swagger:response deleteUserInternalServerError
*/
type DeleteUserInternalServerError struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewDeleteUserInternalServerError creates DeleteUserInternalServerError with default headers values
func NewDeleteUserInternalServerError() *DeleteUserInternalServerError {

	return &DeleteUserInternalServerError{}
}

// WithPayload adds the payload to the delete user internal server error response
func (o *DeleteUserInternalServerError) WithPayload(payload string) *DeleteUserInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete user internal server error response
func (o *DeleteUserInternalServerError) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteUserInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}