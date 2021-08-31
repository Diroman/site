package database

import (
	"context"

	sq "github.com/Masterminds/squirrel"
)

func (c *Client) GetPostOwner(postID int64) (int64, error) {
	ctx := context.Background()

	q := sq.Select("owner").
		From(postsTable).
		Where(sq.Eq{"post_id": postID}).
		PlaceholderFormat(sq.Dollar).
		Limit(1)

	qStr, args, _ := q.ToSql()

	var userID int64
	if err := c.db.QueryRow(ctx, qStr, args...).Scan(&userID); err != nil {
		return 0, err
	}

	return userID, nil
}

func (c *Client) DeletePost(postID int64) error {
	ctx := context.Background()

	q := sq.Delete(postsTable).
		Where(sq.Eq{"post_id": postID}).
		PlaceholderFormat(sq.Dollar)

	qStr, args, _ := q.ToSql()
	if _, err := c.db.Exec(ctx, qStr, args...); err != nil {
		return err
	}

	return nil
}
