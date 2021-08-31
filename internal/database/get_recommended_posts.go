package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
)

func (c *Client) GetRecommendedEventIDs(postID int64) ([]int64, error) {
	ctx := context.Background()

	q := fmt.Sprintf(`
				SELECT up.post_id, count(*) as us_posts
				FROM users_posts up
				WHERE up.user_id IN (SELECT p.user_id
									FROM users_posts p
									WHERE p.post_id = %d
									GROUP BY user_id)
				AND up.post_id <> %d
				GROUP BY up.post_id
				ORDER BY us_posts DESC
				LIMIT 3`,
				postID, postID)

	rows, err := c.db.Query(ctx, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var count int64
	var IDs []int64
	for rows.Next() {
		var id int64
		err := rows.Scan(&id, &count)
		if err != nil {
			if err == pgx.ErrNoRows {
				return IDs, ErrNoPost
			}
			return nil, err
		}

		IDs = append(IDs, id)
	}

	return IDs, nil
}
