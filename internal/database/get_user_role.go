package database

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4"
)

func (c *Client) GetUserRole(email string) (int32, error) {
	ctx := context.Background()

	q := sq.Select("role").
		From(userTable).
		Where(sq.Eq{"email": email}).
		Limit(1).
		PlaceholderFormat(sq.Dollar)

	qStr, _, _ := q.ToSql()

	var role int32
	err := c.db.QueryRow(ctx, qStr, email).Scan(&role)
	if err != nil {
		if err == pgx.ErrNoRows {
			return 0, ErrNoUser
		}
		return 0, err
	}

	return role, nil
}
