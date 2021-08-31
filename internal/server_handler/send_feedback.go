package server_handler

import (
	"fmt"
	"github.com/go-openapi/runtime/middleware"

	"lawSite/internal/logger"
	tls "lawSite/internal/tools"
	"lawSite/restapi/operations/auth"
	"lawSite/restapi/operations/tools"
)

const siteEmail = "lawbox.portal@gmail.com"

func (s *Server) SendFeedback(request tools.SendFeedbackParams) middleware.Responder {
	err := s.db.SendFeedback(request.Body.Email, request.Body.Text)
	if err != nil {
		logger.Log.WithError(err).WithField("email", request.Body.Email).Error("Failed to save feedback")
		errorStruct := tls.CreateError("Что-то пошло не так", nil)
		return auth.NewSendResetLinkInternalServerError().WithPayload(errorStruct)
	}

	err = s.mailClient.SendMessage(siteEmail,
		fmt.Sprintf("Email: %s\nСообщение: %s", request.Body.Email, request.Body.Text))
	if err != nil {
		logger.Log.WithError(err).WithField("email", request.Body.Email).Error("Failed to send feedback email")
	}

	return tools.NewSendFeedbackOK()
}
