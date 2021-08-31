package server_handler

import (
	"errors"

	"github.com/go-openapi/runtime/middleware"

	"lawSite/internal/database"
	"lawSite/internal/logger"
	"lawSite/internal/tools"
	"lawSite/restapi/operations/user"
)

func (s *Server) AddPostToFavorite(request user.UserPostsAddParams, email interface{}) middleware.Responder {
	post, err := s.db.GetPostByID(request.Body.Value)
	if err != nil {
		if errors.Is(err, database.ErrNoPost) {
			logger.Log.WithError(err).WithField("email", email).WithField("post_id", request.Body.Value).
				Error("Not found post")
			errorStruct := tools.CreateError("Post not found", nil)
			return user.NewUserPostsAddBadRequest().WithPayload(errorStruct)
		}
		logger.Log.WithError(err).WithField("email", email).WithField("post_id", request.Body.Value).
			Error("Failed to get post")
		errorStruct := tools.CreateError("Failed to get post", err)
		return user.NewUserPostsAddInternalServerError().WithPayload(errorStruct)
	}

	userID, err := s.db.GetUserIDByEmail(email.(string))
	if err != nil {
		if errors.Is(err, database.ErrNoUser) {
			logger.Log.WithError(err).WithField("email", email).WithField("post_id", request.Body.Value).
				Error("Not found user")
			errorStruct := tools.CreateError("User not found", nil)
			return user.NewUserPostsAddBadRequest().WithPayload(errorStruct)
		}
		logger.Log.WithError(err).WithField("email", email).WithField("post_id", request.Body.Value).
			Error("Failed to get user")
		errorStruct := tools.CreateError("Failed to get post", err)
		return user.NewUserPostsAddInternalServerError().WithPayload(errorStruct)
	}

	err = s.db.AddPostToFavorite(userID, post.Data.ID)
	if err != nil {
		logger.Log.WithError(err).WithField("email", email).WithField("post_id", request.Body.Value).
			Error("Failed add to favorite")
		errorStruct := tools.CreateError("Failed add to favorite", nil)
		return user.NewUserPostsAddInternalServerError().WithPayload(errorStruct)
	}

	return user.NewUserPostsAddOK()
}
