// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// UpdateUserInfoHandlerFunc turns a function with the right signature into a update user info handler
type UpdateUserInfoHandlerFunc func(UpdateUserInfoParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn UpdateUserInfoHandlerFunc) Handle(params UpdateUserInfoParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// UpdateUserInfoHandler interface for that can handle valid update user info params
type UpdateUserInfoHandler interface {
	Handle(UpdateUserInfoParams, interface{}) middleware.Responder
}

// NewUpdateUserInfo creates a new http.Handler for the update user info operation
func NewUpdateUserInfo(ctx *middleware.Context, handler UpdateUserInfoHandler) *UpdateUserInfo {
	return &UpdateUserInfo{Context: ctx, Handler: handler}
}

/* UpdateUserInfo swagger:route POST /api/user/info/update user updateUserInfo

Update user info

*/
type UpdateUserInfo struct {
	Context *middleware.Context
	Handler UpdateUserInfoHandler
}

func (o *UpdateUserInfo) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewUpdateUserInfoParams()
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