package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"

	"lawSite/internal/config"
)

type Client struct {
	db *pgxpool.Pool
}

func NewClient(cfg *config.ConfigStruct) (*Client, error) {
	dbUrl := fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		cfg.Database.User, cfg.Database.Password, cfg.Database.Host, cfg.Database.Port, cfg.Database.Database)
	pool, err := pgxpool.Connect(context.Background(), dbUrl)
	if err != nil {
		return nil, err
	}

	if err := pool.Ping(context.Background()); err != nil {
		return nil, err
	}

	return &Client{db: pool}, nil
}

func (c *Client) Close() {
	c.db.Close()
}
