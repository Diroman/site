package server_handler

import (
	"github.com/go-openapi/runtime/middleware"

	"lawSite/internal/logger"
	"lawSite/internal/tools"
	"lawSite/models"
	"lawSite/restapi/operations/user"
)

func (s *Server) GetPostsCreatedByUser(request user.GetUserCreatedPostsParams, email interface{}) middleware.Responder {
	userID, err := s.db.GetUserIDByEmail(email.(string))
	if err != nil {
		logger.Log.WithError(err).WithField("email", email).Error("Failed to get user by email")
		errorStruct := tools.CreateError("Что-то пошло не так", nil)
		return user.NewGetUserCreatedPostsInternalServerError().WithPayload(errorStruct)
	}

	request.ByTime = applyDefaultByTimeFilter(request.ByTime)

	createdPosts, err := s.db.GetPostsCreatedByUserID(userID, *request.ByTime)
	if err != nil {
		logger.Log.WithError(err).WithField("user_id", userID).Error("Failed to get post created by user")
		errorStruct := tools.CreateError("Failed to get post created by user", err)
		return user.NewGetUserCreatedPostsInternalServerError().WithPayload(errorStruct)
	}

	return user.NewGetUserCreatedPostsOK().WithPayload(&models.PostList{Data: createdPosts})
}
