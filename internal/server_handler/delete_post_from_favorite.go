package server_handler

import (
	"errors"

	"github.com/go-openapi/runtime/middleware"

	"lawSite/internal/database"
	"lawSite/internal/logger"
	"lawSite/internal/tools"
	"lawSite/restapi/operations/user"
)

func (s *Server) DeletePostFromFavorite(request user.UserPostsDeleteParams, email interface{}) middleware.Responder {
	_, err := s.db.GetPostByID(request.Body.Value)
	if err != nil {
		if errors.Is(err, database.ErrNoPost) {
			logger.Log.WithError(err).WithField("email", email).WithField("post_id", request.Body.Value).
				Error("Not found post")
			errorStruct := tools.CreateError("Post not found", nil)
			return user.NewUserPostsDeleteBadRequest().WithPayload(errorStruct)
		}
		logger.Log.WithError(err).WithField("email", email).WithField("post_id", request.Body.Value).
			Error("Failed to get post")
		errorStruct := tools.CreateError("Failed to get post", err)
		return user.NewUserPostsDeleteInternalServerError().WithPayload(errorStruct)
	}

	userID, err := s.db.GetUserIDByEmail(email.(string))
	if err != nil {
		if errors.Is(err, database.ErrNoUser) {
			logger.Log.WithError(err).WithField("email", email).WithField("post_id", request.Body.Value).
				Error("Not found user")
			errorStruct := tools.CreateError("User not found", nil)
			return user.NewUserPostsDeleteBadRequest().WithPayload(errorStruct)
		}
		logger.Log.WithError(err).WithField("email", email).WithField("post_id", request.Body.Value).
			Error("Failed to get user")
		errorStruct := tools.CreateError("Failed to get post", err)
		return user.NewUserPostsDeleteInternalServerError().WithPayload(errorStruct)
	}

	err = s.db.DeletePostToFavorite(userID, request.Body.Value)
	if err != nil {
		logger.Log.WithError(err).WithField("email", email).WithField("post_id", request.Body.Value).
			Error("Failed to delete post")
		errorStruct := tools.CreateError("Failed to delete post", err)
		return user.NewUserPostsDeleteInternalServerError().WithPayload(errorStruct)
	}

	return &user.DeleteUserOK{}
}
