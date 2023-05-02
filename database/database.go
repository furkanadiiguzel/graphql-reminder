package database

import (
	"context"
	"database/sql"
	"time"
)

var connectionString = "postgres://postgres:1234@localhost:5432/dbname?sslmode=require"

func Connect() *sql.DB {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	err = db.PingContext(ctx)
	if err != nil {
		panic(err)
	}

	return db
}
