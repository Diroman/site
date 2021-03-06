// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// CreateUserOKCode is the HTTP code returned for type CreateUserOK
const CreateUserOKCode int = 200

/*CreateUserOK successful operation

swagger:response createUserOK
*/
type CreateUserOK struct {
}

// NewCreateUserOK creates CreateUserOK with default headers values
func NewCreateUserOK() *CreateUserOK {

	return &CreateUserOK{}
}

// WriteResponse to the client
func (o *CreateUserOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// CreateUserBadRequestCode is the HTTP code returned for type CreateUserBadRequest
const CreateUserBadRequestCode int = 400

/*CreateUserBadRequest Invalid username/password supplied

swagger:response createUserBadRequest
*/
type CreateUserBadRequest struct {
}

// NewCreateUserBadRequest creates CreateUserBadRequest with default headers values
func NewCreateUserBadRequest() *CreateUserBadRequest {

	return &CreateUserBadRequest{}
}

// WriteResponse to the client
func (o *CreateUserBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// CreateUserConflictCode is the HTTP code returned for type CreateUserConflict
const CreateUserConflictCode int = 409

/*CreateUserConflict Incorrect login

swagger:response createUserConflict
*/
type CreateUserConflict struct {
}

// NewCreateUserConflict creates CreateUserConflict with default headers values
func NewCreateUserConflict() *CreateUserConflict {

	return &CreateUserConflict{}
}

// WriteResponse to the client
func (o *CreateUserConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(409)
}
