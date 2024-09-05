package repositories

import (
	"reflect"
	"testing"
	"time"

	"github.com/WilliamKSilva/book-reservation/internal/domain"
	"github.com/WilliamKSilva/book-reservation/internal/infra/db"
)

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
	ctx, conn := db.ConnectTestDatabase()
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
