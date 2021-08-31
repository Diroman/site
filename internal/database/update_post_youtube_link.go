package database

import (
	"context"

	sq "github.com/Masterminds/squirrel"
)

func (c *Client) UpdatePostYoutubeLink(postID int64, link string) error {
	ctx := context.Background()

	q := sq.Update(postsTable).
		Where(sq.Eq{"post_id": postID}).
		Set("youtube_link", link).
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
