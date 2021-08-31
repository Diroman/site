package server_handler

import (
	"github.com/go-openapi/runtime/middleware"

	"lawSite/internal/logger"
	"lawSite/internal/tools"
	"lawSite/restapi/operations/post"
)

func (s *Server) DeleteMaterialsFromPost(request post.DeletePostMaterialParams, email interface{}) middleware.Responder {
	if err := s.db.DeleteMaterialsFromPost(request.ID, request.Body.Value); err != nil {
		logger.Log.WithError(err).WithField("request", request).Error("Failed to delete materials")
		errorStruct := tools.CreateError("Failed to delete materials", err)
		return post.NewDeletePostInternalServerError().WithPayload(errorStruct)
	}

	return post.NewDeletePostOK()
}
