// Code generated by go-swagger; DO NOT EDIT.

package post

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// AddNewMaterialsForPostHandlerFunc turns a function with the right signature into a add new materials for post handler
type AddNewMaterialsForPostHandlerFunc func(AddNewMaterialsForPostParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn AddNewMaterialsForPostHandlerFunc) Handle(params AddNewMaterialsForPostParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// AddNewMaterialsForPostHandler interface for that can handle valid add new materials for post params
type AddNewMaterialsForPostHandler interface {
	Handle(AddNewMaterialsForPostParams, interface{}) middleware.Responder
}

// NewAddNewMaterialsForPost creates a new http.Handler for the add new materials for post operation
func NewAddNewMaterialsForPost(ctx *middleware.Context, handler AddNewMaterialsForPostHandler) *AddNewMaterialsForPost {
	return &AddNewMaterialsForPost{Context: ctx, Handler: handler}
}

/* AddNewMaterialsForPost swagger:route POST /api/post/{id}/materials/add post addNewMaterialsForPost

Add new materials

*/
type AddNewMaterialsForPost struct {
	Context *middleware.Context
	Handler AddNewMaterialsForPostHandler
}

func (o *AddNewMaterialsForPost) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewAddNewMaterialsForPostParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc.(interface{}) // this is really a interface{}, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
