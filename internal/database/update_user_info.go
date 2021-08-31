package database

import (
	"context"

	sq "github.com/Masterminds/squirrel"

	"lawSite/models"
)

func (c *Client) UpdateUserInfo(userID int64, user *models.UserData) error {
	ctx := context.Background()

	q := sq.Update(userTable).
		Where(sq.Eq{"user_id": userID}).
		Set("firstname", user.FirstName).
		Set("lastname", user.LastName).
		Set("phone", user.Phone).
		Set("role", user.Role).
		Set("age", user.Age).
		Set("user_status", user.UserStatus).
		PlaceholderFormat(sq.Dollar)

	qStr, args, _ := q.ToSql()
	_, err := c.db.Exec(ctx, qStr, args...)
	if err != nil {
		return err
	}

	return nil
}
