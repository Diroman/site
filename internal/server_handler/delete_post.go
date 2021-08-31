package server_handler

import (
	"github.com/go-openapi/runtime/middleware"
	"lawSite/internal/model"

	"lawSite/internal/logger"
	"lawSite/internal/tools"
	"lawSite/restapi/operations/post"
)

func (s *Server) DeletePost(request post.DeletePostParams, email interface{}) middleware.Responder {
	var emailString string
	if email != nil {
		emailString = email.(string)
	}
	userID, err := s.db.GetUserIDByEmail(emailString)
	if err != nil {
		logger.Log.WithError(err).WithField("email", email).Error("Failed to get user ID")
		errorStruct := tools.CreateError("Что-то пошло не так", nil)
		return post.NewDeletePostInternalServerError().WithPayload(errorStruct)
	}

	owner, err := s.db.GetPostOwner(request.Body.Value)
	if err != nil {
		logger.Log.WithError(err).WithField("email", email).WithField("post_id", request.Body.Value).
			Error("Failed to get post owner")
		errorStruct := tools.CreateError("Что-то пошло не так", nil)
		return post.NewDeletePostInternalServerError().WithPayload(errorStruct)
	}

	role, _ := s.db.GetUserRole(emailString)

	if owner != userID && role != model.AdminRole {
		logger.Log.WithError(err).WithField("user_id", userID).WithField("owner_id", owner).
			Error("Try to delete someone else's post")
		errorStruct := tools.CreateError("Недостаточно прав", nil)
		return post.NewDeletePostForbidden().WithPayload(errorStruct)
	}

	err = s.db.DeletePost(request.Body.Value)
	if err != nil {
		logger.Log.WithError(err).WithField("email", email).WithField("post_id", request.Body.Value).
			Error("Failed to delete post")
		errorStruct := tools.CreateError("Что-то пошло не так", nil)
		return post.NewDeletePostInternalServerError().WithPayload(errorStruct)
	}

	return post.NewDeletePostOK()
}
