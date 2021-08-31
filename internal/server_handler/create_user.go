package server_handler

import (
	"github.com/go-openapi/runtime/middleware"

	bearer "lawSite/internal/auth"
	"lawSite/internal/logger"
	"lawSite/internal/tools"
	"lawSite/models"
	"lawSite/restapi/operations/auth"
)

func (s *Server) CreateUser(request auth.RegisterUserParams) middleware.Responder {
	if request.Body == nil || request.Body.Data == nil {
		errorStruct := tools.CreateError("Empty body or data", nil)
		return auth.NewRegisterUserBadRequest().WithPayload(errorStruct)
	}

	isRegister, err := s.db.ClientIsRegister(request.Body.Data.Email)
	if err != nil {
		logger.Log.WithError(err).Error("Failed to check user register")
		errorStruct := tools.CreateError("Failed to check user register", err)
		return auth.NewRegisterUserInternalServerError().WithPayload(errorStruct)
	}

	if isRegister {
		logger.Log.Warn("User already register")
		errorStruct := tools.CreateError("User already register", nil)
		return auth.NewRegisterUserBadRequest().WithPayload(errorStruct)
	}

	err = s.db.CreateUser(request.Body.Data)
	if err != nil {
		logger.Log.WithError(err).Error("Failed to register user")
		errorStruct := tools.CreateError("Failed to register user", err)
		return auth.NewRegisterUserInternalServerError().WithPayload(errorStruct)
	}

	token, err := bearer.CreateNewToken(request.Body.Data.Email)
	if err != nil {
		errorStruct := &models.Error{Data: &models.ErrorData{Error: err.Error()}}
		return auth.NewLoginUserInternalServerError().WithPayload(errorStruct)
	}

	resp := tools.CreateToken(token)
	logger.Log.Warn("Success user register")

	return auth.NewRegisterUserOK().WithPayload(resp)
}
