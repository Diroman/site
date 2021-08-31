package database

import (
	"context"

	sq "github.com/Masterminds/squirrel"

	"lawSite/internal/auth"
)

func (c *Client) UpdatePassword(userID int64, password string) error {
	ctx := context.Background()

	q := sq.Update(userTable).
		Where(sq.Eq{"user_id": userID}).
		Set("password", password).
		PlaceholderFormat(sq.Dollar)

	encryptPassword, err := auth.Hash.Generate(password)
	if err != nil {
		return err
	}

	qStr, _, _ := q.ToSql()
	_, err = c.db.Exec(ctx, qStr, encryptPassword, userID)
	if err != nil {
		return err
	}

	return nil
}
