package server_handler

import (
	"github.com/go-openapi/runtime/middleware"

	"lawSite/internal/logger"
	"lawSite/internal/tools"
	"lawSite/models"
	"lawSite/restapi/operations/post"
)

func (s *Server) CheckPostInFavorite(request post.CheckPostInFavoriteParams, email interface{}) middleware.Responder {
	userID, err := s.db.GetUserIDByEmail(email.(string))
	if err != nil {
		logger.Log.WithError(err).WithField("email", email).WithField("postID", request.ID).
			Error("Failed to get user by email")
		errorStruct := tools.CreateError("Что-то пошло не так", nil)
		return post.NewCheckPostInFavoriteInternalServerError().WithPayload(errorStruct)
	}

	isFavorite, err := s.db.CheckPostInFavorite(userID, request.ID)
	if err != nil {
		logger.Log.WithError(err).WithField("email", email).WithField("postID", request.ID).
			Error("Failed to get user`s post")
		errorStruct := tools.CreateError("Что-то пошло не так", nil)
		return post.NewCheckPostInFavoriteInternalServerError().WithPayload(errorStruct)
	}

	return post.NewCheckPostInFavoriteOK().WithPayload(&models.Bool{Data: isFavorite})
}
