package user

import (
	"time"

	_ "github.com/WilliamKSilva/book-reservation/docs"
)

type User struct {
	ID        string
	Name      string
	Email     string
	CPF       string
	Password  string
	BirthDate time.Time
}

type CreateRequestDTO struct {
	// @example {"name": "John Doe", "email": "john.doe@example.com", "password": "teste1234", "cpf": "23212332112", "birth_date": "2024-08-12"}
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CPF       string `json:"cpf"`
	BirthDate string `json:"birth_date"`
}

type CreateResponseDTO struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CPF       string `json:"cpf"`
	BirthDate string `json:"birth_date"`
}

type IUserRepository interface {
	Save(User) (User, error)
	FindByEmail(email string) (User, error)
}

type IUUIDGenerator interface {
	Generate() string
}

type IUserService interface {
	Create(name string, email string, CPF string, birthDate string) (*User, error)
}

func NewUserService(userRepository IUserRepository, uuidGenerator IUUIDGenerator) *UserService {
	return &UserService{
		UserRepository: userRepository,
		UuidGenerator:  uuidGenerator,
	}
}

type UserService struct {
	UserRepository IUserRepository
	UuidGenerator  IUUIDGenerator
}

func (userService *UserService) Create(name string, email string, CPF string, birthDate string) (*User, error) {
	birthDateTime, err := time.Parse("2006-01-02", birthDate)
	if err != nil {
		return nil, err
	}

	uuid := userService.UuidGenerator.Generate()

	user := User{
		ID:        uuid,
		Name:      name,
		Email:     email,
		CPF:       CPF,
		BirthDate: birthDateTime,
	}

	createdUser, err := userService.UserRepository.Save(user)
	if err != nil {
		return nil, err
	}

	return &createdUser, nil
}
