package db

import (
    "context"
    "log"
    "os"

    "github.com/jackc/pgx/v5/pgxpool"
)

var Pool *pgxpool.Pool

func Init() {
    dbUrl := os.Getenv("DB_URL")
    var err error
    Pool, err = pgxpool.New(context.Background(), dbUrl)
    if err != nil {
        log.Fatalf("Unable to connect to database: %v\n", err)
    }
}

func Close() {
    Pool.Close()
}
