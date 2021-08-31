package server_handler

import (
	"github.com/go-openapi/runtime/middleware"

	"lawSite/internal/database"
	"lawSite/internal/logger"
	"lawSite/internal/tools"
	"lawSite/restapi/operations/user"
)

func (s *Server) GetUserInfo(request user.GetUserInfoParams, email interface{}) middleware.Responder {
	resp, err := s.db.GetUserByEmail(email.(string))
	if err != nil {
		if err != database.ErrNoUser {
			logger.Log.WithError(err).WithField("email", email).Error("User not found")
			errorStruct := tools.CreateError("User not found", nil)
			return user.NewGetUserInfoBadRequest().WithPayload(errorStruct)
		}
		logger.Log.WithError(err).WithField("email", email).Error("Failed to get user by ID")
		errorStruct := tools.CreateError("Что-то пошло не так", err)
		return user.NewGetUserInfoBadRequest().WithPayload(errorStruct)
	}

	return user.NewGetUserInfoOK().WithPayload(resp)
}
