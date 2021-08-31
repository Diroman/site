package server_handler

import (
	"errors"

	"github.com/go-openapi/runtime/middleware"

	"lawSite/internal/database"
	"lawSite/internal/logger"
	"lawSite/internal/tools"
	"lawSite/restapi/operations/post"
)

func (s *Server) GetNearestEvent(request post.GetNearestPostParams) middleware.Responder {
	findPost, err := s.db.GetNearestEvent()
	if err != nil {
		if errors.Is(err, database.ErrNoPost) {
			logger.Log.WithError(err).Error("Post not found")
			return post.NewGetNearestPostOK().WithPayload(findPost)
		}
		logger.Log.WithError(err).Error("Failed to get post")
		errorStruct := tools.CreateError("Failed to get post", err)
		return post.NewGetNearestPostInternalServerError().WithPayload(errorStruct)
	}

	return post.NewGetNearestPostOK().WithPayload(findPost)
}
