package database

import (
	"context"

	sq "github.com/Masterminds/squirrel"

	"lawSite/models"
)

func (c *Client) AddMaterialsToPost(postID int64, materials []*models.Material) (int64, error) {
	ctx := context.Background()

	q := sq.Insert(materialsTable).
		Columns("title", "link", "file", "text_field", "post_id").
		PlaceholderFormat(sq.Dollar).
		Suffix("RETURNING material_id")

	for _, material := range materials {
		data := material.Data
		q = q.Values(data.Title, data.Link, data.File, data.Text, postID)
	}

	strQ, args, err := q.ToSql()
	if err != nil {
		return 0, err
	}

	var id int64
	err = c.db.QueryRow(ctx, strQ, args...).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
