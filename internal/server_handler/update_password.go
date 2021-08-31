package server_handler

import (
	"github.com/go-openapi/runtime/middleware"

	"lawSite/internal/logger"
	"lawSite/internal/tools"
	"lawSite/restapi/operations/auth"
)

func (s *Server) UpdatePassword(request auth.UpdatePasswordParams) middleware.Responder {
	token := request.Body.Token

	userID := s.linkGenerator.GetUserIDByToken(token)
	if userID == -1 {
		logger.Log.WithField("token", token).Error("ID by token not found")
		errorStruct := tools.CreateError("Неверная ссылка", nil)
		return auth.NewUpdatePasswordBadRequest().WithPayload(errorStruct)
	}

	password := request.Body.Password

	if err := s.db.UpdatePassword(userID, password); err != nil {
		logger.Log.WithError(err).WithField("token", token).Error("Failed to update password")
		errorStruct := tools.CreateError("Failed to update password", nil)
		return auth.NewUpdatePasswordInternalServerError().WithPayload(errorStruct)
	}

	return auth.NewUpdatePasswordOK()
}
