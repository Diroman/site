package server_handler

import (
	"github.com/go-openapi/runtime/middleware"

	"lawSite/internal/logger"
	"lawSite/internal/model"
	"lawSite/internal/tools"
	"lawSite/models"
	"lawSite/restapi/operations/auth"
	"lawSite/restapi/operations/post"
)

func (s *Server) CreatePost(request post.AddNewPostParams, email interface{}) middleware.Responder {
	if request.Body == nil || request.Body.Data == nil {
		errorStruct := tools.CreateError("Empty body or data", nil)
		return post.NewAddNewPostInternalServerError().WithPayload(errorStruct)
	}

	role, err := s.db.GetUserRole(email.(string))
	if err != nil {
		logger.Log.WithError(err).WithField("email", email).Error("Failed to get user role")
		errorStruct := tools.CreateError("Что-то пошло не так", nil)
		return post.NewAddNewPostInternalServerError().WithPayload(errorStruct)
	}

	if role == model.UserRole {
		logger.Log.WithField("email", email).Warn("Forbidden when add post")
		errorStruct := tools.CreateError("Недостаточно прав", nil)
		return post.NewAddNewPostForbidden().WithPayload(errorStruct)
	}

	userID, err := s.db.GetUserIDByEmail(email.(string))
	if err != nil {
		logger.Log.WithError(err).WithField("email", email).Error("Failed to get user by email")
		errorStruct := tools.CreateError("Что-то пошло не так", nil)
		return post.NewAddNewPostInternalServerError().WithPayload(errorStruct)
	}

	id, err := s.db.CreatePost(userID, request.Body.Data)
	if err != nil {
		logger.Log.WithError(err).Error("Failed to add post")
		errorStruct := tools.CreateError("Что-то пошло не так", nil)
		return auth.NewRegisterUserInternalServerError().WithPayload(errorStruct)
	}

	return post.NewAddNewPostOK().WithPayload(&models.ID{Value: id})
}
