package database

import (
	"context"

	sq "github.com/Masterminds/squirrel"
)

const feedbackTable = "feedback"

func (c *Client) SendFeedback(email, text string) error {
	ctx := context.Background()

	q := sq.Insert(feedbackTable).
		Columns("email", "text").
		Values(email, text).
		PlaceholderFormat(sq.Dollar)

	qStr, args, _ := q.ToSql()
	if _, err := c.db.Exec(ctx, qStr, args...); err != nil {
		return err
	}

	return nil
}
