package database

import (
	"context"

	sq "github.com/Masterminds/squirrel"
)

func (c *Client) CheckPostInFavorite(userID, postID int64) (bool, error) {
	ctx := context.Background()

	q := sq.Select("count(*)").
		From(userPostTable).
		Where(sq.And{
			sq.Eq{"user_id": userID},
			sq.Eq{"post_id": postID},
		}).
		Limit(1).
		PlaceholderFormat(sq.Dollar)

	strQ, args, err := q.ToSql()
	if err != nil {
		return false, err
	}

	var count int64
	if err := c.db.QueryRow(ctx, strQ, args...).Scan(&count); err != nil {
		return false, err
	}

	return count > 0, nil
}
