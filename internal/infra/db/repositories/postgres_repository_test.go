package repositories

import (
	"context"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/WilliamKSilva/book-reservation/internal/domain"
	"github.com/WilliamKSilva/book-reservation/internal/infra/db"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func ConnectTestDatabase(t *testing.T) (context.Context, *pgxpool.Pool) {
	ctx := context.Background()
	err := godotenv.Load("../../../../.env")
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

func mockUser() (domain.User, error) {
	birthDateTime, err := time.Parse("2006-01-02", "2024-08-13")
	if err != nil {
		return domain.User{}, err
	}
	return domain.User{
		ID:        "ff508158-8da7-4840-b891-38c240f9aee1",
		Name:      "johndoe",
		Email:     "johndoe@teste.com",
		CPF:       "12312312323",
		Password:  "teste1234",
		BirthDate: birthDateTime,
	}, err
}

func TestPostgresUserRepository(t *testing.T) {
	ctx, conn := ConnectTestDatabase(t)
	defer conn.Close()

	t.Run("save user in the database", func(t *testing.T) {
		userRepository := PostgresUserRepository{
			Conn: conn,
			Ctx:  ctx,
		}

		mockedUser, err := mockUser()
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
