package repositories_tests

import (
	"context"
	"os"
	"reflect"
	"testing"

	"github.com/WilliamKSilva/book-reservation/internal/domain/user"
	"github.com/WilliamKSilva/book-reservation/internal/infra/db"
	"github.com/WilliamKSilva/book-reservation/internal/infra/db/repositories"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func ConnectTestDatabase(t *testing.T) (context.Context, *pgxpool.Pool) {
	ctx := context.Background()
	err := godotenv.Load("../../../../../.env")
	if err != nil {
		t.Errorf("Error loading .env: %s", err.Error())
	}

	dbURL := os.Getenv("DATABASE_TEST_URL")
	poolConfig, err := pgxpool.ParseConfig(dbURL)
	if err != nil {
		t.Errorf("Unable to parse database URL: %s", err.Error())
	}

	pool, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		t.Errorf("Unable to create connection pool: %v\n", err)
	}

	if err := pool.Ping(ctx); err != nil {
		t.Errorf("Error trying to connect to database: %v\n", err)
	}

	return ctx, pool
}

func TestPostgresUserRepository(t *testing.T) {
	ctx, conn := ConnectTestDatabase(t)
	defer conn.Close()

	t.Run("save user in the database", func(t *testing.T) {
		userRepository := repositories.PostgresUserRepository{
			Conn: conn,
			Ctx:  ctx,
		}

		mockedUser, err := user.MockUser()
		if err != nil {
			t.Error("Error mocking User", err)
			return
		}

		u, err := userRepository.Save(mockedUser)
		if err != nil {
			t.Errorf("Error trying to save User: %s", err)
			return
		}

		if !reflect.DeepEqual(mockedUser, u) {
			t.Errorf("Saved User does not match")
			return
		}

		err = db.ResetDatabaseState(ctx, conn)
		if err != nil {
			t.Errorf("Error trying to reset database state: %s", err)
		}
	})
}
