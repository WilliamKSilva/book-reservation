package internal

import "time"

type User struct {
	ID        string
	Name      string
	Email     string
	CPF       string
	BirthDate time.Time
}

type IUserRepository interface {
	Save(*User) (User, error)
}

type IUserService interface {
	Create(name string, email string, CPF string, birthDate string) (*User, error)
}

type UserService struct {
	userRepository IUserRepository
}

func (userService UserService) Create(name string, email string, CPF string, birthDate string) (*User, error) {
	birthDateTime, err := time.Parse("2006-01-02", birthDate)
	if err != nil {
		return nil, err
	}

	user := User{
		ID:        "1",
		Name:      name,
		Email:     email,
		CPF:       CPF,
		BirthDate: birthDateTime,
	}

	createdUser, err := userService.userRepository.Save(&user)
	if err != nil {
		return nil, err
	}

	return &createdUser, nil
}
