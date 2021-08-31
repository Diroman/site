package database

import (
	"context"

	sq "github.com/Masterminds/squirrel"

	"lawSite/models"
)

func (c *Client) GetPostParticipants(postID int64) ([]*models.User, error) {
	ctx := context.Background()

	q := sq.Select("firstname", "lastname", "email", "phone", "age", "user_status").
		From("users_posts up").
		Join("users u USING (user_id)").
		Where(sq.Eq{"up.post_id": postID}).
		PlaceholderFormat(sq.Dollar)

	qStr, _, _ := q.ToSql()
	rows, err := c.db.Query(ctx, qStr, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*models.User
	for rows.Next() {
		user := &models.UserData{}
		err := rows.Scan(&user.FirstName, &user.LastName, &user.Email, &user.Phone, &user.Age, &user.UserStatus)
		if err != nil {
			return nil, err
		}

		users = append(users, &models.User{Data: user})
	}

	return users, nil
}
