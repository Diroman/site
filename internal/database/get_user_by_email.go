package database

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4"

	"lawSite/models"
)

func (c *Client) GetUserByEmail(email string) (*models.User, error) {
	ctx := context.Background()

	q := sq.Select("firstname", "lastname", "email", "phone", "role", "age", "user_status").
		From(userTable).
		Where(sq.Eq{"email": email}).
		Limit(1).
		PlaceholderFormat(sq.Dollar)

	qStr, _, _ := q.ToSql()

	var user models.UserData
	err := c.db.QueryRow(ctx, qStr, email).Scan(&user.FirstName, &user.LastName,
		&user.Email, &user.Phone, &user.Role, &user.Age, &user.UserStatus)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, ErrNoUser
		}
		return nil, err
	}

	return &models.User{Data: &user}, nil
}

func (c *Client) GetUserIDByEmail(email string) (int64, error) {
	ctx := context.Background()

	q := sq.Select("user_id").
		From(userTable).
		Where(sq.Eq{"email": email}).
		Limit(1).
		PlaceholderFormat(sq.Dollar)

	qStr, _, _ := q.ToSql()

	var ID int64
	err := c.db.QueryRow(ctx, qStr, email).Scan(&ID)
	if err != nil {
		if err == pgx.ErrNoRows {
			return 0, ErrNoUser
		}
		return 0, err
	}

	return ID, nil
}
