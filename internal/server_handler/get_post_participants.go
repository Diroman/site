package server_handler

import (
	"github.com/go-openapi/runtime/middleware"

	"lawSite/internal/logger"
	"lawSite/internal/model"
	"lawSite/internal/tools"
	"lawSite/models"
	"lawSite/restapi/operations/post"
)

func (s *Server) GetPostsParticipants(request post.GetPostParticipantsParams, email interface{}) middleware.Responder {
	role, err := s.db.GetUserRole(email.(string))
	if err != nil {
		logger.Log.WithError(err).WithField("email", email).Error("Failed to get user role")
		errorStruct := tools.CreateError("Что-то пошло не так", nil)
		return post.NewGetPostParticipantsInternalServerError().WithPayload(errorStruct)
	}

	if role == model.UserRole {
		logger.Log.WithField("email", email).Warn("Forbidden when get post participants")
		errorStruct := tools.CreateError("Недостаточно прав", nil)
		return post.NewGetPostParticipantsForbidden().WithPayload(errorStruct)
	}

	users, err := s.db.GetPostParticipants(request.ID)
	if err != nil {
		logger.Log.WithError(err).WithField("post_id", request.ID).Error("Failed to get post participants")
		errorStruct := tools.CreateError("Failed to get post participants", err)
		return post.NewGetPostParticipantsInternalServerError().WithPayload(errorStruct)
	}

	return post.NewGetPostParticipantsOK().WithPayload(&models.UserList{Data: users})
}
