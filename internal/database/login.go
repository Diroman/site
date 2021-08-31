package database

import (
	"context"
	"errors"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4"
)

var ErrNoUser = errors.New("Пользователь с таким email отсутствует ")

func (c *Client) GetUserPassword(email string) (string, error) {
	ctx := context.Background()

	q := sq.Select("password").
		From(userTable).
		Where(sq.Eq{"email": email}).
		Limit(1).
		PlaceholderFormat(sq.Dollar)

	qStr, _, _ := q.ToSql()

	var password string
	err := c.db.QueryRow(ctx, qStr, email).Scan(&password)
	if err != nil {
		if err == pgx.ErrNoRows {
			return "", ErrNoUser
		}
		return "", err
	}

	return password, nil
}
