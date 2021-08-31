package database

import (
	"context"

	sq "github.com/Masterminds/squirrel"

	"lawSite/models"
)

const (
	materialsTable     = "materials"
	materialsPostTable = "post_materials"
)

func (c *Client) GetMaterialsByPostID(postID int64) ([]*models.Material, error) {
	ctx := context.Background()

	q := sq.Select("material_id", "title", "link", "file", "post_id", "text_field").
		From(materialsTable).
		Where(sq.Eq{"post_id": postID}).
		//OrderBy("link", "file").
		PlaceholderFormat(sq.Dollar)

	strQ, args, err := q.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := c.db.Query(ctx, strQ, args...)
	if err != nil {
		return nil, err
	}

	var materials []*models.Material
	for rows.Next() {
		var material models.MaterialData
		if err := rows.Scan(&material.ID, &material.Title,
			&material.Link, &material.File, &material.PostID, &material.Text); err != nil {
			return nil, err
		}

		material.PostID = postID
		materials = append(materials, &models.Material{Data: &material})
	}

	return materials, nil
}
