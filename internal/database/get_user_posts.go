package database

import (
	"context"

	sq "github.com/Masterminds/squirrel"
)

func (c *Client) GetUserRegisteredPostIDs(id int64) ([]int64, error) {
	ctx := context.Background()

	q := sq.Select("post_id").
		From(userPostTable).
		Where(sq.Eq{"user_id": id}).
		PlaceholderFormat(sq.Dollar)

	qStr, _, _ := q.ToSql()
	rows, err := c.db.Query(ctx, qStr, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var postIDs []int64
	for rows.Next() {
		var postID int64
		if err := rows.Scan(&postID); err != nil {
			return nil, err
		}
		postIDs = append(postIDs, postID)
	}

	return postIDs, nil
}
