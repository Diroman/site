package database

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4"

	"lawSite/internal/model"
	"lawSite/models"
)

func (c *Client) GetNearestEvent() (*models.Post, error) {
	ctx := context.Background()

	q := sq.Select("post_id", "title", "text", "image", "date", "place",
		"duration", "link", "event_type", "contacts", "price").
		From(postsTable).
		OrderBy("date").
		Where(getByTimeFilter(model.NEW)).
		PlaceholderFormat(sq.Dollar).
		Limit(1)

	qStr, _, _ := q.ToSql()

	var p models.PostData
	err := c.db.QueryRow(ctx, qStr).Scan(&p.ID, &p.Title, &p.Body, &p.Image, &p.Date,
		&p.Place, &p.Duration, &p.TranslationLink, &p.EventType, &p.Contacts, &p.Price)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, ErrNoPost
		}
		return nil, err
	}

	p.EventType = translationEvent[p.EventType]

	return &models.Post{Data: &p}, nil
}
