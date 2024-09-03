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

type IUserRepository interface {
	Save(user User) (User, error)
	FindByEmail(email string) (User, error)
}

type IUUIDGenerator interface {
	Generate() string
}

type IUserService interface {
	Create(CreateUserRequestDTO) (CreateUserResponseDTO, error)
	FindByEmail(email string) (FindUserByEmailResponseDTO, error)
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

func (userService *UserService) Create(createUserRequestDTO CreateUserRequestDTO) (CreateUserResponseDTO, error) {
	birthDateTime, err := time.Parse("2006-01-02", createUserRequestDTO.BirthDate)
	if err != nil {
		return CreateUserResponseDTO{}, err
	}

	uuid := userService.UuidGenerator.Generate()

	user := User{
		ID:        uuid,
		Name:      createUserRequestDTO.Name,
		Email:     createUserRequestDTO.Email,
		CPF:       createUserRequestDTO.CPF,
		Password:  createUserRequestDTO.Password,
		BirthDate: birthDateTime,
	}

	createdUser, err := userService.UserRepository.Save(user)
	if err != nil {
		return CreateUserResponseDTO{}, err
	}

	return CreateUserResponseDTO{
		ID:        createdUser.ID,
		Name:      createdUser.Name,
		Email:     createdUser.Email,
		Password:  createdUser.Password,
		CPF:       createdUser.CPF,
		BirthDate: createdUser.BirthDate.String(),
	}, nil
}

func (userService *UserService) FindByEmail(email string) (FindUserByEmailResponseDTO, error) {
	user, err := userService.UserRepository.FindByEmail(email)
	if err != nil {
		return FindUserByEmailResponseDTO{}, err
	}

	return FindUserByEmailResponseDTO{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		CPF:       user.CPF,
		BirthDate: user.BirthDate.String(),
	}, nil
}
