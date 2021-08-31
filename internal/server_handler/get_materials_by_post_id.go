package server_handler

import (
	"github.com/go-openapi/runtime/middleware"

	"lawSite/internal/logger"
	"lawSite/internal/tools"
	"lawSite/models"
	"lawSite/restapi/operations/post"
)

func (s *Server) GetMaterialsByID(request post.GetMaterialByPostIDParams) middleware.Responder {
	materials, err := s.db.GetMaterialsByPostID(request.ID)
	if err != nil {
		logger.Log.WithError(err).WithField("req", request).Error("Failed to get materials")
		errorStruct := tools.CreateError("Что-то пошло не так", nil)
		return post.NewGetMaterialByPostIDInternalServerError().WithPayload(errorStruct)
	}

	if materials == nil {
		materials = []*models.Material{}
	}
	return post.NewGetMaterialByPostIDOK().WithPayload(&models.MaterialList{Data: materials})
}
