package server_handler

import (
	"github.com/go-openapi/runtime/middleware"

	"lawSite/internal/logger"
	"lawSite/internal/tools"
	"lawSite/restapi/operations/post"
)

func (s *Server) UpdatePost(request post.UpdatePostParams, email interface{}) middleware.Responder {
	userID, err := s.db.GetUserIDByEmail(email.(string))
	if err != nil {
		logger.Log.WithError(err).WithField("email", email).Error("Failed to get user ID")
		errorStruct := tools.CreateError("Что-то пошло не так", nil)
		return post.NewUpdatePostInternalServerError().WithPayload(errorStruct)
	}

	owner, err := s.db.GetPostOwner(request.Body.Data.ID)
	if err != nil {
		logger.Log.WithError(err).WithField("email", email).WithField("post_id", request.Body.Data.ID).
			Error("Failed to get post owner")
		errorStruct := tools.CreateError("Что-то пошло не так", nil)
		return post.NewUpdatePostInternalServerError().WithPayload(errorStruct)
	}

	if owner != userID {
		logger.Log.WithError(err).WithField("user_id", userID).WithField("owner_id", owner).
			Error("Try to update someone else's post")
		errorStruct := tools.CreateError("Недостаточно прав", nil)
		return post.NewUpdatePostForbidden().WithPayload(errorStruct)
	}

	err = s.db.UpdatePost(request.Body.Data)
	if err != nil {
		logger.Log.WithError(err).WithField("email", email).WithField("post_id", request.Body.Data.ID).
			Error("Failed to update post")
		errorStruct := tools.CreateError("Что-то пошло не так", nil)
		return post.NewUpdatePostInternalServerError().WithPayload(errorStruct)
	}

	return post.NewUpdatePostOK()
}
