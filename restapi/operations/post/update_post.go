// Code generated by go-swagger; DO NOT EDIT.

package post

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// UpdatePostHandlerFunc turns a function with the right signature into a update post handler
type UpdatePostHandlerFunc func(UpdatePostParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn UpdatePostHandlerFunc) Handle(params UpdatePostParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// UpdatePostHandler interface for that can handle valid update post params
type UpdatePostHandler interface {
	Handle(UpdatePostParams, interface{}) middleware.Responder
}

// NewUpdatePost creates a new http.Handler for the update post operation
func NewUpdatePost(ctx *middleware.Context, handler UpdatePostHandler) *UpdatePost {
	return &UpdatePost{Context: ctx, Handler: handler}
}

/* UpdatePost swagger:route POST /api/post/update post updatePost

Update posts

*/
type UpdatePost struct {
	Context *middleware.Context
	Handler UpdatePostHandler
}

func (o *UpdatePost) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewUpdatePostParams()
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
