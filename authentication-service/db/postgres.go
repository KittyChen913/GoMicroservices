package db

import (
	"context"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/joho/godotenv/autoload"
)

var DbPool *pgxpool.Pool
var retryCount int = 3

func InitDb() {
	var err error
	connStr := os.Getenv("postgresDB")

	config, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		panic("Unable to parse database config: " + err.Error())
	}
	config.MaxConns = 10
	config.MinConns = 5

	// 建立連線池
	DbPool, err = pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		panic("Unable to connect to the database: " + err.Error())
	}

	// 連線測試
	for i := 1; i <= retryCount; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		err = DbPool.Ping(ctx)
		cancel()
		if err == nil {
			break
		}
		if i == retryCount {
			panic("Database ping failed: " + err.Error())
		}
		time.Sleep(2 * time.Second)
	}
}
