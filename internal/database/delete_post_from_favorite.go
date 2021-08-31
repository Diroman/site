package database

import (
	"context"

	sq "github.com/Masterminds/squirrel"
)

func (c *Client) DeletePostToFavorite(userID, postID int64) error {
	ctx := context.Background()

	q := sq.Delete(userPostTable).
		Where(
			sq.And{
				sq.Eq{"user_id": userID},
				sq.Eq{"post_id": postID},
			},
		).
		PlaceholderFormat(sq.Dollar)

	qStr, _, _ := q.ToSql()

	_, err := c.db.Exec(ctx, qStr, userID, postID)
	if err != nil {
		return err
	}

	return nil
}
