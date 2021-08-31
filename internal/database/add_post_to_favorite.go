package database

import (
	"context"

	sq "github.com/Masterminds/squirrel"
)

const userPostTable = "public.users_posts"

func (c *Client) AddPostToFavorite(userID, postID int64) error {
	ctx := context.Background()

	q := sq.Insert(userPostTable).
		Columns("user_id", "post_id").
		Values(userID, postID).
		PlaceholderFormat(sq.Dollar)

	qStr, _, _ := q.ToSql()

	_, err := c.db.Exec(ctx, qStr, userID, postID)
	if err != nil {
		return err
	}

	return nil
}
