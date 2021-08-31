package server_handler

import (
	"errors"

	"github.com/go-openapi/runtime/middleware"

	"lawSite/internal/database"
	"lawSite/internal/logger"
	"lawSite/internal/tools"
	"lawSite/restapi/operations/post"
)

func (s *Server) GetPostByID(request post.GetPostByIDParams) middleware.Responder {
	findPost, err := s.db.GetPostByID(request.ID)
	if err != nil {
		if errors.Is(err, database.ErrNoPost) {
			logger.Log.WithError(err).WithField("id", request.ID).Error("Post not found")
			errorStruct := tools.CreateError("", err)
			return post.NewGetPostListBadRequest().WithPayload(errorStruct)
		}
		logger.Log.WithError(err).WithField("id", request.ID).Error("Failed to get post")
		errorStruct := tools.CreateError("Failed to get post", err)
		return post.NewGetPostByIDInternalServerError().WithPayload(errorStruct)
	}

	return post.NewGetPostByIDOK().WithPayload(findPost)
}
