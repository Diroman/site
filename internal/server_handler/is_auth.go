package server_handler

import (
	"github.com/go-openapi/runtime/middleware"

	"lawSite/restapi/operations/auth"
)

func (s *Server) IsAuth(request auth.IsAuthParams, a interface{}) middleware.Responder {
	return auth.NewIsAuthOK()
}
