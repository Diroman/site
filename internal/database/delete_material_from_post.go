package database

import (
	"context"

	sq "github.com/Masterminds/squirrel"
)

func (c *Client) DeleteMaterialsFromPost(postID, materialID int64) error {
	ctx := context.Background()

	q := sq.Delete(materialsTable).
		Where(sq.And{
			sq.Eq{"post_id": postID},
			sq.Eq{"material_id": materialID},
		}).
		PlaceholderFormat(sq.Dollar)

	strQ, args, err := q.ToSql()
	if err != nil {
		return err
	}

	if _, err := c.db.Exec(ctx, strQ, args...); err != nil {
		return err
	}
	return nil
}
