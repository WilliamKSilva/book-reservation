package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func ResetDatabaseState(ctx context.Context, conn *pgxpool.Pool) error {
	_, err := conn.Exec(ctx, `
		TRUNCATE TABLE users_books_reservations CASCADE;
		TRUNCATE TABLE users CASCADE;
		TRUNCATE TABLE books CASCADE;
	`)

	return err
}

func Connect(ctx context.Context) *pgxpool.Pool {
	config, err := pgxpool.ParseConfig(os.Getenv("DATABASE_URL"))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	return pool
}
