package server_handler

import (
	"errors"

	"github.com/go-openapi/runtime/middleware"

	"lawSite/internal/database"
	"lawSite/internal/logger"
	"lawSite/internal/tools"
	"lawSite/restapi/operations/auth"
)

func (s *Server) SendResetLink(request auth.SendResetLinkParams) middleware.Responder {
	email := request.Email.Value

	userID, err := s.db.GetUserIDByEmail(email)
	if err != nil {
		if errors.Is(err, database.ErrNoUser) {
			logger.Log.WithError(err).WithField("email", email).Error("Not found user")
			errorStruct := tools.CreateError("Пользователь не найден", nil)
			return auth.NewSendResetLinkBadRequest().WithPayload(errorStruct)
		}
		logger.Log.WithError(err).WithField("email", email).Error("Can't get user id by email")
		errorStruct := tools.CreateError("Что-то пошло не так", nil)
		return auth.NewSendResetLinkInternalServerError().WithPayload(errorStruct)
	}

	token := s.linkGenerator.GenerateLink()

	err = s.mailClient.SendPasswordUpdate(email, token)
	if err != nil {
		logger.Log.WithError(err).WithField("email", email).Error("Can't send email")
		errorStruct := tools.CreateError("Что-то пошло не так", nil)
		return auth.NewSendResetLinkInternalServerError().WithPayload(errorStruct)
	}

	s.linkGenerator.AddLink(userID, token)

	return auth.NewSendResetLinkOK()
}
