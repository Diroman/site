// Code generated by go-swagger; DO NOT EDIT.

package post

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"lawSite/models"
)

// GetPostListOKCode is the HTTP code returned for type GetPostListOK
const GetPostListOKCode int = 200

/*GetPostListOK successful operation

swagger:response getPostListOK
*/
type GetPostListOK struct {

	/*
	  In: Body
	*/
	Payload *models.PostList `json:"body,omitempty"`
}

// NewGetPostListOK creates GetPostListOK with default headers values
func NewGetPostListOK() *GetPostListOK {

	return &GetPostListOK{}
}

// WithPayload adds the payload to the get post list o k response
func (o *GetPostListOK) WithPayload(payload *models.PostList) *GetPostListOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get post list o k response
func (o *GetPostListOK) SetPayload(payload *models.PostList) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPostListOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetPostListBadRequestCode is the HTTP code returned for type GetPostListBadRequest
const GetPostListBadRequestCode int = 400

/*GetPostListBadRequest Invalid status value

swagger:response getPostListBadRequest
*/
type GetPostListBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetPostListBadRequest creates GetPostListBadRequest with default headers values
func NewGetPostListBadRequest() *GetPostListBadRequest {

	return &GetPostListBadRequest{}
}

// WithPayload adds the payload to the get post list bad request response
func (o *GetPostListBadRequest) WithPayload(payload *models.Error) *GetPostListBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get post list bad request response
func (o *GetPostListBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPostListBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetPostListInternalServerErrorCode is the HTTP code returned for type GetPostListInternalServerError
const GetPostListInternalServerErrorCode int = 500

/*GetPostListInternalServerError Server error

swagger:response getPostListInternalServerError
*/
type GetPostListInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetPostListInternalServerError creates GetPostListInternalServerError with default headers values
func NewGetPostListInternalServerError() *GetPostListInternalServerError {

	return &GetPostListInternalServerError{}
}

// WithPayload adds the payload to the get post list internal server error response
func (o *GetPostListInternalServerError) WithPayload(payload *models.Error) *GetPostListInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get post list internal server error response
func (o *GetPostListInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPostListInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
