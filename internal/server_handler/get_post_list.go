package server_handler

import (
	"errors"
	"fmt"
	"time"

	"github.com/go-openapi/runtime/middleware"

	"lawSite/internal/logger"
	"lawSite/internal/model"
	"lawSite/internal/tools"
	"lawSite/restapi/operations/post"
)

func (s *Server) GetPostList(request post.GetPostListParams) middleware.Responder {
	request, err := applyDefaults(request)
	if err != nil {
		logger.Log.WithError(err).WithField("req", request).Error("Invalid request")
		errorStruct := tools.CreateError("Invalid request", err)
		return post.NewGetPostListBadRequest().WithPayload(errorStruct)
	}

	posts, err := s.db.GetPostList(request)
	if err != nil {
		logger.Log.WithError(err).WithField("req", request).Error("Failed to get posts")
		errorStruct := tools.CreateError("Что-то пошло не так", nil)
		return post.NewGetPostListInternalServerError().WithPayload(errorStruct)
	}

	// если производился поиск по заголовку и не было ничего найдено, нужно вернуть 400
	if request.Text != nil && len(posts.Data) == 0 {
		errorStruct := tools.CreateError("Ничего не найдено", nil)
		return post.NewGetPostListBadRequest().WithPayload(errorStruct)
	}

	return post.NewGetPostListOK().WithPayload(posts)
}

func applyDefaults(request post.GetPostListParams) (post.GetPostListParams, error) {
	if request.DateFrom != nil && DateInValidFormat(*request.DateFrom) != nil {
		return request, errors.New("Invalid date from ")
	}
	if request.DateTo != nil && DateInValidFormat(*request.DateTo) != nil {
		return request, errors.New("Invalid date to ")
	}

	if request.SortBy == nil && request.Text == nil {
		def := model.Date
		request.SortBy = &def
	}

	if request.Order == nil {
		def := model.DESC
		request.Order = &def
	}

	if request.Limit == nil {
		def := int64(20)
		request.Limit = &def
	}

	if request.Offset == nil {
		def := int64(0)
		request.Offset = &def
	}

	if request.ByTime == nil {
		def := model.ALL
		request.ByTime = &def
	}

	return request, nil
}

func DateInValidFormat(date string) error {
	if len(date) == 0 {
		return nil
	}
	if len(date) < 10 {
		return errors.New("short date string")
	}
	if _, err := time.Parse("2006-01-02", date[0:10]); err != nil {
		return errors.New(fmt.Sprintf("cant`t parse date_from string: %s", err))
	}
	return nil
}
