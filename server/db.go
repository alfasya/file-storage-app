package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var pool *pgxpool.Pool

func startDb() {
	var err error
	pool, err = pgxpool.New(context.Background(), "postgres://alfasya:2121@localhost:5433/file_storage_db")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Connection has been established")
}
