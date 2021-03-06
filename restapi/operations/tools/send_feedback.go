// Code generated by go-swagger; DO NOT EDIT.

package tools

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// SendFeedbackHandlerFunc turns a function with the right signature into a send feedback handler
type SendFeedbackHandlerFunc func(SendFeedbackParams) middleware.Responder

// Handle executing the request and returning a response
func (fn SendFeedbackHandlerFunc) Handle(params SendFeedbackParams) middleware.Responder {
	return fn(params)
}

// SendFeedbackHandler interface for that can handle valid send feedback params
type SendFeedbackHandler interface {
	Handle(SendFeedbackParams) middleware.Responder
}

// NewSendFeedback creates a new http.Handler for the send feedback operation
func NewSendFeedback(ctx *middleware.Context, handler SendFeedbackHandler) *SendFeedback {
	return &SendFeedback{Context: ctx, Handler: handler}
}

/* SendFeedback swagger:route POST /api/tools/feedback tools sendFeedback

Send feedback

*/
type SendFeedback struct {
	Context *middleware.Context
	Handler SendFeedbackHandler
}

func (o *SendFeedback) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewSendFeedbackParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
