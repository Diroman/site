package server_handler

import (
	"github.com/go-openapi/runtime/middleware"

	"lawSite/internal/logger"
	"lawSite/internal/tools"
	"lawSite/restapi/operations/post"
)

func (s *Server) UpdatePostYoutubeLink(request post.UpdateLinkInPostParams, email interface{}) middleware.Responder {
	err := s.db.UpdatePostYoutubeLink(request.ID, request.Link.Data)
	if err != nil {
		logger.Log.WithError(err).WithField("req", request).Error("Failed to update youtube link")
		errorStruct := tools.CreateError("", err)
		return post.NewUpdateLinkInPostInternalServerError().WithPayload(errorStruct)
	}

	return post.NewUpdateLinkInPostOK()
}
