// Code generated by go-swagger; DO NOT EDIT.

package post

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetPostListHandlerFunc turns a function with the right signature into a get post list handler
type GetPostListHandlerFunc func(GetPostListParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetPostListHandlerFunc) Handle(params GetPostListParams) middleware.Responder {
	return fn(params)
}

// GetPostListHandler interface for that can handle valid get post list params
type GetPostListHandler interface {
	Handle(GetPostListParams) middleware.Responder
}

// NewGetPostList creates a new http.Handler for the get post list operation
func NewGetPostList(ctx *middleware.Context, handler GetPostListHandler) *GetPostList {
	return &GetPostList{Context: ctx, Handler: handler}
}

/* GetPostList swagger:route GET /api/post/list post getPostList

Popular post for main page

*/
type GetPostList struct {
	Context *middleware.Context
	Handler GetPostListHandler
}

func (o *GetPostList) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetPostListParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
