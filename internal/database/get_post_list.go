package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgtype"
	"strings"
	"time"

	"lawSite/internal/model"
	"lawSite/models"
	"lawSite/restapi/operations/post"
)

const postsTable = "public.posts"

const popularQuery = `(SELECT post_id, count(user_id) as popular
					   FROM users_posts
					   GROUP BY post_id) up ON posts.post_id = up.post_id`

var translationEvent = map[string]string{
	model.Lecture:     "Лекции и материалы",
	model.MasterClass: "Мастер-классы",
	model.Conference:  "Конференции",
	model.Training:    "Тренинги",
	model.Webinars:    "Вебинары и эфиры",
}

func (c *Client) GetPostList(request post.GetPostListParams) (*models.PostList, error) {
	ctx := context.Background()

	var filters []string

	filters = append(filters, getByTimeFilter(*request.ByTime))
	if request.EventType != nil {
		filters = append(filters, fmt.Sprintf("event_type = '%s'", *request.EventType))
	}

	if request.DateFrom != nil && request.DateTo != nil {
		filters = append(filters, fmt.Sprintf("date::DATE >= TO_DATE('%s','YYYY-MM-DD')", *request.DateFrom))
		filters = append(filters, fmt.Sprintf("date::DATE <= TO_DATE('%s','YYYY-MM-DD')", *request.DateTo))
	} else {
		if request.DateFrom != nil {
			filters = append(filters, fmt.Sprintf("date::DATE = TO_DATE('%s','YYYY-MM-DD')", *request.DateFrom))
		}
	}

	orderBy := getOrderBy(*request.Order, request.SortBy, request.Text)

	q := fmt.Sprintf(
		`SELECT posts.post_id, title, text, image, date, place, 
						duration, link, event_type, contacts, price, tags, 
						speakers, coalesce(popular, 0)
				FROM %s LEFT JOIN %s
				WHERE %s
				ORDER BY %s, start_time
				LIMIT %d
				OFFSET %d`,
		postsTable, popularQuery,
		strings.Join(filters, " AND "),
		orderBy,
		*request.Limit, *request.Offset)

	rows, err := c.db.Query(ctx, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []*models.Post
	for rows.Next() {
		var p models.PostData
		var tags, speakers pgtype.TextArray
		err := rows.Scan(&p.ID, &p.Title, &p.Body, &p.Image, &p.Date, &p.Place, &p.Duration, &p.TranslationLink,
			&p.EventType, &p.Contacts, &p.Price, &tags, &speakers, &p.MemberCount)
		if err != nil {
			return nil, err
		}

		p.Hashtag = &models.List{Data: make([]string, len(tags.Elements))}
		p.Speakers = &models.List{Data: make([]string, len(speakers.Elements))}
		err = tags.AssignTo(&p.Hashtag.Data)
		if err != nil {
			return nil, err
		}
		err = speakers.AssignTo(&p.Speakers.Data)
		if err != nil {
			return nil, err
		}
		p.EventType = translationEvent[p.EventType]

		posts = append(posts, &models.Post{Data: &p})
	}

	return &models.PostList{Data: posts}, nil
}

func getOrderBy(orderBy string, sortBy, text *string) string {
	if sortBy != nil {
		var sort string
		switch *sortBy {
		case model.Date:
			sort = "date"
		case model.Popularity:
			sort = "popular"
		case model.CreateDate:
			sort = "timestamp"
		}

		return fmt.Sprintf("%s %s", sort, orderBy)
	}

	if text != nil {
		return fmt.Sprintf("ts_rank(to_tsvector(title), plainto_tsquery('%s')) DESC", *text)
	}

	return ""
}

func getByTimeFilter(byTime string) string {
	now := time.Now().Format("2006-01-02 15:04:05")

	switch byTime {
	case model.ALL:
		return "TRUE"
	case model.NEW:
		return fmt.Sprintf("end_time > '%s'", now)
	case model.OLD:
		return fmt.Sprintf("end_time < '%s'", now)
	}
	return "TRUE"
}