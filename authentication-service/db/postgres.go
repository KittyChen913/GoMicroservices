package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DbPool *pgxpool.Pool

func InitDb() {
	var err error

	config, err := pgxpool.ParseConfig("postgresql://postgres:password@postgres-db/users")
	if err != nil {
		panic("Unable to connect to the database.")
	}
	config.MaxConns = 10
	config.MinConns = 5

	// 建立連線池
	DbPool, err = pgxpool.NewWithConfig(context.Background(), config)
}
