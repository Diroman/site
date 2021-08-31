package server_handler

import (
	"github.com/go-openapi/runtime/middleware"

	"lawSite/internal/logger"
	"lawSite/internal/tools"
	"lawSite/restapi/operations/auth"
)

func (s *Server) CheckResetToken(request auth.CheckResetTokenParams) middleware.Responder {
	token := request.Token.Value

	userID := s.linkGenerator.GetUserIDByToken(token)
	if userID == -1 {
		logger.Log.WithField("token", token).Error("ID by token not found")
		errorStruct := tools.CreateError("Неверная ссылка", nil)
		return auth.NewCheckResetTokenBadRequest().WithPayload(errorStruct)
	}

	return auth.NewCheckResetTokenOK()
}
