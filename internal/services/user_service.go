package services

import (
	"time"

	"github.com/WilliamKSilva/book-reservation/internal/domain"
	"github.com/WilliamKSilva/book-reservation/internal/services/DTOs"
)

type IUserRepository interface {
	Save(user domain.User) (domain.User, error)
	FindByEmail(email string) (domain.User, error)
}

type IUUIDGenerator interface {
	Generate() string
}

type IUserService interface {
	Create(DTOs.CreateUserRequestDTO) (DTOs.CreateUserResponseDTO, error)
	FindByEmail(email string) (DTOs.FindUserByEmailResponseDTO, error)
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

func (userService *UserService) Create(createUserRequestDTO DTOs.CreateUserRequestDTO) (DTOs.CreateUserResponseDTO, error) {
	birthDateTime, err := time.Parse("2006-01-02", createUserRequestDTO.BirthDate)
	if err != nil {
		return DTOs.CreateUserResponseDTO{}, err
	}

	uuid := userService.UuidGenerator.Generate()

	user := domain.User{
		ID:        uuid,
		Name:      createUserRequestDTO.Name,
		Email:     createUserRequestDTO.Email,
		CPF:       createUserRequestDTO.CPF,
		Password:  createUserRequestDTO.Password,
		BirthDate: birthDateTime,
	}

	createdUser, err := userService.UserRepository.Save(user)
	if err != nil {
		return DTOs.CreateUserResponseDTO{}, err
	}

	return DTOs.CreateUserResponseDTO{
		ID:        createdUser.ID,
		Name:      createdUser.Name,
		Email:     createdUser.Email,
		Password:  createdUser.Password,
		CPF:       createdUser.CPF,
		BirthDate: createdUser.BirthDate.String(),
	}, nil
}

func (userService *UserService) FindByEmail(email string) (DTOs.FindUserByEmailResponseDTO, error) {
	user, err := userService.UserRepository.FindByEmail(email)
	if err != nil {
		return DTOs.FindUserByEmailResponseDTO{}, err
	}

	return DTOs.FindUserByEmailResponseDTO{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		CPF:       user.CPF,
		BirthDate: user.BirthDate.String(),
	}, nil
}