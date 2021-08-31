// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetUserPostsHandlerFunc turns a function with the right signature into a get user posts handler
type GetUserPostsHandlerFunc func(GetUserPostsParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetUserPostsHandlerFunc) Handle(params GetUserPostsParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetUserPostsHandler interface for that can handle valid get user posts params
type GetUserPostsHandler interface {
	Handle(GetUserPostsParams, interface{}) middleware.Responder
}

// NewGetUserPosts creates a new http.Handler for the get user posts operation
func NewGetUserPosts(ctx *middleware.Context, handler GetUserPostsHandler) *GetUserPosts {
	return &GetUserPosts{Context: ctx, Handler: handler}
}

/* GetUserPosts swagger:route GET /api/user/posts user getUserPosts

Get user posts

*/
type GetUserPosts struct {
	Context *middleware.Context
	Handler GetUserPostsHandler
}

func (o *GetUserPosts) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetUserPostsParams()
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
