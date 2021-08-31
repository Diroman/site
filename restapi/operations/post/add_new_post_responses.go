// Code generated by go-swagger; DO NOT EDIT.

package post

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"lawSite/models"
)

// AddNewPostOKCode is the HTTP code returned for type AddNewPostOK
const AddNewPostOKCode int = 200

/*AddNewPostOK successful operation

swagger:response addNewPostOK
*/
type AddNewPostOK struct {

	/*
	  In: Body
	*/
	Payload *models.ID `json:"body,omitempty"`
}

// NewAddNewPostOK creates AddNewPostOK with default headers values
func NewAddNewPostOK() *AddNewPostOK {

	return &AddNewPostOK{}
}

// WithPayload adds the payload to the add new post o k response
func (o *AddNewPostOK) WithPayload(payload *models.ID) *AddNewPostOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add new post o k response
func (o *AddNewPostOK) SetPayload(payload *models.ID) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddNewPostOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AddNewPostBadRequestCode is the HTTP code returned for type AddNewPostBadRequest
const AddNewPostBadRequestCode int = 400

/*AddNewPostBadRequest Bad request

swagger:response addNewPostBadRequest
*/
type AddNewPostBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewAddNewPostBadRequest creates AddNewPostBadRequest with default headers values
func NewAddNewPostBadRequest() *AddNewPostBadRequest {

	return &AddNewPostBadRequest{}
}

// WithPayload adds the payload to the add new post bad request response
func (o *AddNewPostBadRequest) WithPayload(payload *models.Error) *AddNewPostBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add new post bad request response
func (o *AddNewPostBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddNewPostBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AddNewPostUnauthorizedCode is the HTTP code returned for type AddNewPostUnauthorized
const AddNewPostUnauthorizedCode int = 401

/*AddNewPostUnauthorized not valid token

swagger:response addNewPostUnauthorized
*/
type AddNewPostUnauthorized struct {
}

// NewAddNewPostUnauthorized creates AddNewPostUnauthorized with default headers values
func NewAddNewPostUnauthorized() *AddNewPostUnauthorized {

	return &AddNewPostUnauthorized{}
}

// WriteResponse to the client
func (o *AddNewPostUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// AddNewPostForbiddenCode is the HTTP code returned for type AddNewPostForbidden
const AddNewPostForbiddenCode int = 403

/*AddNewPostForbidden Forbidden

swagger:response addNewPostForbidden
*/
type AddNewPostForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewAddNewPostForbidden creates AddNewPostForbidden with default headers values
func NewAddNewPostForbidden() *AddNewPostForbidden {

	return &AddNewPostForbidden{}
}

// WithPayload adds the payload to the add new post forbidden response
func (o *AddNewPostForbidden) WithPayload(payload *models.Error) *AddNewPostForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add new post forbidden response
func (o *AddNewPostForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddNewPostForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AddNewPostInternalServerErrorCode is the HTTP code returned for type AddNewPostInternalServerError
const AddNewPostInternalServerErrorCode int = 500

/*AddNewPostInternalServerError Server error

swagger:response addNewPostInternalServerError
*/
type AddNewPostInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewAddNewPostInternalServerError creates AddNewPostInternalServerError with default headers values
func NewAddNewPostInternalServerError() *AddNewPostInternalServerError {

	return &AddNewPostInternalServerError{}
}

// WithPayload adds the payload to the add new post internal server error response
func (o *AddNewPostInternalServerError) WithPayload(payload *models.Error) *AddNewPostInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add new post internal server error response
func (o *AddNewPostInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddNewPostInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
