package server_handler

import (
	"errors"
	"lawSite/internal/model"

	"github.com/go-openapi/runtime/middleware"

	"lawSite/internal/database"
	"lawSite/internal/logger"
	"lawSite/internal/tools"
	"lawSite/models"
	"lawSite/restapi/operations/user"
)

func (s *Server) GetUserRegisteredPosts(request user.GetUserPostsParams, email interface{}) middleware.Responder {
	userID, err := s.db.GetUserIDByEmail(email.(string))
	if err != nil {
		if errors.Is(err, database.ErrNoUser) {
			logger.Log.WithError(err).WithField("email", email).Error("Not found user")
			errorStruct := tools.CreateError("User not found", nil)
			return user.NewGetUserPostsBadRequest().WithPayload(errorStruct)
		}
		logger.Log.WithError(err).WithField("email", email).Error("Can't get user id by email")
		errorStruct := tools.CreateError("Can't get user id by email", err)
		return user.NewGetUserPostsInternalServerError().WithPayload(errorStruct)
	}

	request.ByTime = applyDefaultByTimeFilter(request.ByTime)

	postIDs, err := s.db.GetUserRegisteredPostIDs(userID)
	if err != nil {
		logger.Log.WithError(err).WithField("user_id", userID).Error("Failed to get user post IDs")
		errorStruct := tools.CreateError("Failed to get user post IDs", err)
		return user.NewGetUserPostsInternalServerError().WithPayload(errorStruct)
	}

	posts, err := s.db.GetPostsByIDs(postIDs, *request.ByTime)
	if err != nil {
		logger.Log.WithError(err).WithField("post_ids", postIDs).Error("Failed to get post by IDs")
		errorStruct := tools.CreateError("Failed to get post by IDs", err)
		return user.NewGetUserPostsInternalServerError().WithPayload(errorStruct)
	}

	return user.NewGetUserPostsOK().WithPayload(&models.PostList{Data: posts})
}

func applyDefaultByTimeFilter(byTime *string) *string {
	var defFilter string
	if byTime != nil {
		return byTime
	}
	defFilter = model.ALL

	return &defFilter
}