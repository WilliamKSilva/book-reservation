package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func ConnectTestDatabase() (context.Context, *pgxpool.Pool) {
	ctx := context.Background()
	err := godotenv.Load("../../../../.env")
	if err != nil {
		log.Printf("Error loading .env: %s", err.Error())
		os.Exit(1)
	}

	dbURL := os.Getenv("DATABASE_TEST_URL")
	poolConfig, err := pgxpool.ParseConfig(dbURL)
	if err != nil {
		log.Printf("Unable to parse database URL: %s", err.Error())
		os.Exit(1)
	}

	pool, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		log.Printf("Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}

	if err := pool.Ping(ctx); err != nil {
		fmt.Printf("Error trying to connect to database: %v\n", err)
		os.Exit(1)
	}

	return ctx, pool
}

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
