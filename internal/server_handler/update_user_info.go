package server_handler

import (
	"github.com/go-openapi/runtime/middleware"

	"lawSite/internal/logger"
	"lawSite/internal/tools"
	"lawSite/restapi/operations/user"
)

func (s *Server) UpdateUserInfo(request user.UpdateUserInfoParams, email interface{}) middleware.Responder {
	userID, err := s.db.GetUserIDByEmail(email.(string))
	if err != nil {
		logger.Log.WithError(err).WithField("email", email).Error("Failed to get user ID")
		errorStruct := tools.CreateError("Что-то пошло не так", nil)
		return user.NewUpdateUserInfoInternalServerError().WithPayload(errorStruct)
	}

	err = s.db.UpdateUserInfo(userID, request.Body.Data)
	if err != nil {
		logger.Log.WithError(err).WithField("email", email).Error("Failed to update user info")
		errorStruct := tools.CreateError("Что-то пошло не так", nil)
		return user.NewUpdateUserInfoInternalServerError().WithPayload(errorStruct)
	}

	return nil
}
