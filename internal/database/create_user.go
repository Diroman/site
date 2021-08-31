package database

import (
	"context"

	sq "github.com/Masterminds/squirrel"

	"lawSite/internal/auth"
	"lawSite/models"
)

const userTable = "public.Users"

func (c *Client) ClientIsRegister(email string) (bool, error) {
	ctx := context.Background()

	q := sq.Select("count(*)").
		From(userTable).
		Where(sq.Eq{"email": email}).
		Limit(1).
		PlaceholderFormat(sq.Dollar)

	qStr, _, _ := q.ToSql()

	var count int
	err := c.db.QueryRow(ctx, qStr, email).Scan(&count)
	if err != nil {
		return false, err
	}

	if count > 0 {
		return true, nil
	}
	return false, nil
}

func (c *Client) CreateUser(data *models.UserData) error {
	ctx := context.Background()

	encryptPassword, err := auth.Hash.Generate(data.Password)
	if err != nil {
		return err
	}

	q := sq.Insert(userTable).
		Columns("firstname", "lastname", "email",
			"password", "phone", "age", "user_status").
		PlaceholderFormat(sq.Dollar)

	q = q.Values(data.FirstName, data.LastName, data.Email,
		encryptPassword, data.Phone, data.Age, data.UserStatus)

	qStr, args, _ := q.ToSql()
	_, err = c.db.Exec(ctx, qStr, args...)
	if err != nil {
		return err
	}

	return nil
}
