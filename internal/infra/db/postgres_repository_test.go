package db

import (
	"reflect"
	"testing"
	"time"

	"github.com/WilliamKSilva/book-reservation/internal/app/user"
)

// it should save a User entity in the database
func mockUser() (user.User, error) {
	birthDateTime, err := time.Parse("2006-01-02", "2024-08-13")
	if err != nil {
		return user.User{}, err
	}
	return user.User{
		ID:        "ff508158-8da7-4840-b891-38c240f9aee1",
		Name:      "johndoe",
		Email:     "johndoe@teste.com",
		CPF:       "12312312323",
		Password:  "teste1234",
		BirthDate: birthDateTime,
	}, err
}

func TestPostgresUserRepository(t *testing.T) {
	ctx, conn := ConnectTestDatabase()
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

		err = ResetDatabaseState(ctx, conn)
		if err != nil {
			t.Errorf("Error trying to reset database state: %s", err)
		}
	})
}
