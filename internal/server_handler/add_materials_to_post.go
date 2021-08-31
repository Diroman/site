package server_handler

import (
	"fmt"
	"github.com/go-openapi/runtime/middleware"
	"lawSite/models"

	"lawSite/internal/logger"
	"lawSite/internal/tools"
	"lawSite/restapi/operations/post"
)

const maxTextLen = 100

func (s *Server) AddMaterialsToPost(request post.AddNewMaterialsForPostParams, email interface{}) middleware.Responder {
	for _, material := range request.Body.Data {
		if !checkTextField(material.Data.Text) {
			logger.Log.WithField("request", request).Error("Text field is too long")
			errorStruct := tools.CreateError(fmt.Sprintf("Слишком длинное текстовое поле,  "), nil)
			return post.NewAddNewMaterialsForPostBadRequest().WithPayload(errorStruct)
		}
	}

	id, err := s.db.AddMaterialsToPost(request.ID, request.Body.Data)
	if err != nil {
		logger.Log.WithError(err).WithField("request", request).Error("Failed to add materials")
		errorStruct := tools.CreateError("Failed to add materials", err)
		return post.NewAddNewMaterialsForPostInternalServerError().WithPayload(errorStruct)
	}

	return post.NewAddNewMaterialsForPostOK().WithPayload(&models.ID{Value: id})
}

func checkTextField(text string) bool {
	if len(text) != 0 && len(text) > maxTextLen {
		return false
	}
	return true
}
