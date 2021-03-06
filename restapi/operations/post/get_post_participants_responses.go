// Code generated by go-swagger; DO NOT EDIT.

package post

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"lawSite/models"
)

// GetPostParticipantsOKCode is the HTTP code returned for type GetPostParticipantsOK
const GetPostParticipantsOKCode int = 200

/*GetPostParticipantsOK successful operation

swagger:response getPostParticipantsOK
*/
type GetPostParticipantsOK struct {

	/*
	  In: Body
	*/
	Payload *models.UserList `json:"body,omitempty"`
}

// NewGetPostParticipantsOK creates GetPostParticipantsOK with default headers values
func NewGetPostParticipantsOK() *GetPostParticipantsOK {

	return &GetPostParticipantsOK{}
}

// WithPayload adds the payload to the get post participants o k response
func (o *GetPostParticipantsOK) WithPayload(payload *models.UserList) *GetPostParticipantsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get post participants o k response
func (o *GetPostParticipantsOK) SetPayload(payload *models.UserList) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPostParticipantsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetPostParticipantsBadRequestCode is the HTTP code returned for type GetPostParticipantsBadRequest
const GetPostParticipantsBadRequestCode int = 400

/*GetPostParticipantsBadRequest Bad request

swagger:response getPostParticipantsBadRequest
*/
type GetPostParticipantsBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetPostParticipantsBadRequest creates GetPostParticipantsBadRequest with default headers values
func NewGetPostParticipantsBadRequest() *GetPostParticipantsBadRequest {

	return &GetPostParticipantsBadRequest{}
}

// WithPayload adds the payload to the get post participants bad request response
func (o *GetPostParticipantsBadRequest) WithPayload(payload *models.Error) *GetPostParticipantsBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get post participants bad request response
func (o *GetPostParticipantsBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPostParticipantsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetPostParticipantsForbiddenCode is the HTTP code returned for type GetPostParticipantsForbidden
const GetPostParticipantsForbiddenCode int = 403

/*GetPostParticipantsForbidden Forbidden

swagger:response getPostParticipantsForbidden
*/
type GetPostParticipantsForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetPostParticipantsForbidden creates GetPostParticipantsForbidden with default headers values
func NewGetPostParticipantsForbidden() *GetPostParticipantsForbidden {

	return &GetPostParticipantsForbidden{}
}

// WithPayload adds the payload to the get post participants forbidden response
func (o *GetPostParticipantsForbidden) WithPayload(payload *models.Error) *GetPostParticipantsForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get post participants forbidden response
func (o *GetPostParticipantsForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPostParticipantsForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetPostParticipantsInternalServerErrorCode is the HTTP code returned for type GetPostParticipantsInternalServerError
const GetPostParticipantsInternalServerErrorCode int = 500

/*GetPostParticipantsInternalServerError Server error

swagger:response getPostParticipantsInternalServerError
*/
type GetPostParticipantsInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetPostParticipantsInternalServerError creates GetPostParticipantsInternalServerError with default headers values
func NewGetPostParticipantsInternalServerError() *GetPostParticipantsInternalServerError {

	return &GetPostParticipantsInternalServerError{}
}

// WithPayload adds the payload to the get post participants internal server error response
func (o *GetPostParticipantsInternalServerError) WithPayload(payload *models.Error) *GetPostParticipantsInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get post participants internal server error response
func (o *GetPostParticipantsInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPostParticipantsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
