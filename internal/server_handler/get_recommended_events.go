package server_handler

import (
	"errors"

	"github.com/go-openapi/runtime/middleware"

	"lawSite/internal/database"
	"lawSite/internal/logger"
	"lawSite/internal/model"
	"lawSite/internal/tools"
	"lawSite/models"
	"lawSite/restapi/operations/post"
)

func (s *Server) GetRecommendedEvent(request post.GetRecommendedPostParams) middleware.Responder {
	recommendedPostIDs, err := s.db.GetRecommendedEventIDs(request.PostID.Value)
	if err != nil {
		if errors.Is(err, database.ErrNoPost) {
			logger.Log.WithError(err).Error("Recommended post not found")
			return post.NewGetRecommendedPostOK().WithPayload(&models.PostList{Data: []*models.Post{}})
		}
		logger.Log.WithError(err).Error("Failed to get post")
		errorStruct := tools.CreateError("Failed to get post", err)
		return post.NewGetRecommendedPostInternalServerError().WithPayload(errorStruct)
	}

	recommendedPosts, err := s.db.GetPostsByIDs(recommendedPostIDs, model.ALL)
	if err != nil {
		logger.Log.WithError(err).WithField("post_ids", recommendedPostIDs).Error("Failed to get post by IDs")
		errorStruct := tools.CreateError("Failed to get post by IDs", err)
		return post.NewGetRecommendedPostInternalServerError().WithPayload(errorStruct)
	}

	return post.NewGetRecommendedPostOK().WithPayload(&models.PostList{Data: recommendedPosts})
}
