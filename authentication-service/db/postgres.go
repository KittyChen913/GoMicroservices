package db

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/joho/godotenv/autoload"
)

var DbPool *pgxpool.Pool

func InitDb() {
	var err error
	connStr := os.Getenv("postgresDB")

	config, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		panic("Unable to connect to the database.")
	}
	config.MaxConns = 10
	config.MinConns = 5

	// 建立連線池
	DbPool, err = pgxpool.NewWithConfig(context.Background(), config)
}
