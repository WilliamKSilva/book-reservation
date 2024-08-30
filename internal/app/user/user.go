package user

import "time"

type User struct {
	ID        string
	Name      string
	Email     string
	CPF       string
	Password  string
	BirthDate time.Time
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
