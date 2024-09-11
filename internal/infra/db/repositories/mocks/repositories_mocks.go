package repositories_mocks

import (
	"errors"

	"github.com/WilliamKSilva/book-reservation/internal/domain/user"
)

type MockedUserRepositorySuccess struct{}

func (userRepository *MockedUserRepositorySuccess) Save(_ user.User) (user.User, error) {
	u, err := user.MockUser()
	return u, err
}

func (userRepository *MockedUserRepositorySuccess) FindByEmail(email string) (user.User, error) {
	u, err := user.MockUser()
	return u, err
}

func NewMockedUserRepositorySuccess() *MockedUserRepositorySuccess {
	return &MockedUserRepositorySuccess{}
}

// This is horrible, but we can't override the FindByEmail method of MockedUserRepoistorySuccess
type MockedUserRepositorySuccessFindByEmailNotFound struct{}

func (userRepository *MockedUserRepositorySuccessFindByEmailNotFound) Save(_ user.User) (user.User, error) {
	u, err := user.MockUser()
	return u, err
}

func (userRepository *MockedUserRepositorySuccessFindByEmailNotFound) FindByEmail(email string) (user.User, error) {
	return user.User{}, nil
}

func NewMockedUserRepositorySuccessFindByEmailNotFound() *MockedUserRepositorySuccessFindByEmailNotFound {
	return &MockedUserRepositorySuccessFindByEmailNotFound{}
}

type MockedUserRepositoryFailure struct{}

func (userRepository *MockedUserRepositoryFailure) Save(u user.User) (user.User, error) {
	return user.User{}, errors.New("generic MockedUserRepositoryFailure save error")
}

func (userRepository *MockedUserRepositoryFailure) FindByEmail(email string) (user.User, error) {
	return user.User{}, errors.New("generic MockedUserRepositoryFailure save error")
}

func NewMockedUserRepositoryFailure() *MockedUserRepositoryFailure {
	return &MockedUserRepositoryFailure{}
}
