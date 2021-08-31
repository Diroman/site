package server_handler

import (
	"errors"

	"github.com/go-openapi/runtime/middleware"

	bearer "lawSite/internal/auth"
	"lawSite/internal/database"
	"lawSite/internal/logger"
	"lawSite/internal/tools"
	"lawSite/models"
	"lawSite/restapi/operations/auth"
)

func (s *Server) LoginUser(request auth.LoginUserParams) middleware.Responder {
	password, err := s.db.GetUserPassword(request.LoginInfo.Data.Login)
	if err != nil {
		if errors.Is(err, database.ErrNoUser) {
			errorStruct := tools.CreateError("", err)
			return auth.NewLoginUserBadRequest().WithPayload(errorStruct)
		}
		logger.Log.WithError(err).WithField("email", request.LoginInfo.Data.Login).Error("Failed to get user")
		errorStruct := tools.CreateError("Failed to get user", err)
		return auth.NewLoginUserInternalServerError().WithPayload(errorStruct)
	}

	err = bearer.Hash.ComparePassword(password, request.LoginInfo.Data.Password)
	if err != nil {
		errorStruct := tools.CreateError("Неверный пароль", nil)
		return auth.NewLoginUserBadRequest().WithPayload(errorStruct)
	}

	token, err := bearer.CreateNewToken(request.LoginInfo.Data.Login)
	if err != nil {
		errorStruct := &models.Error{Data: &models.ErrorData{Error: err.Error()}}
		return auth.NewLoginUserInternalServerError().WithPayload(errorStruct)
	}

	resp := tools.CreateToken(token)
	return auth.NewLoginUserOK().WithPayload(resp)
}
