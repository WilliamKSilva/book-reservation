package services

import (
	"time"

	"github.com/WilliamKSilva/book-reservation/internal/domain/user"
	"github.com/WilliamKSilva/book-reservation/internal/services/DTOs"
	services_errors "github.com/WilliamKSilva/book-reservation/internal/services/errors"
)

type UserServiceInterface interface {
	Create(DTOs.CreateUserRequestDTO) (DTOs.CreateUserResponseDTO, error)
	FindByEmail(email string) (DTOs.FindUserByEmailResponseDTO, error)
}

func NewUserService(userRepository UserRepositoryInterface, uuidGenerator UuidServiceInterface) *UserService {
	return &UserService{
		UserRepository: userRepository,
		UuidService:    uuidGenerator,
	}
}

type UserService struct {
	UserRepository   UserRepositoryInterface
	UuidService      UuidServiceInterface
	EncrypterService EncrypterServiceInterface
}

func (userService *UserService) Create(createUserRequestDTO DTOs.CreateUserRequestDTO) (DTOs.CreateUserResponseDTO, error) {
	vErr := ValidateStructData(createUserRequestDTO)
	if vErr != nil {
		return DTOs.CreateUserResponseDTO{}, vErr
	}

	birthDateTime, err := time.Parse("2006-01-02", createUserRequestDTO.BirthDate)
	if err != nil {
		return DTOs.CreateUserResponseDTO{}, err
	}

	uuid, err := userService.UuidService.Generate()
	if err != nil {
		return DTOs.CreateUserResponseDTO{}, err
	}

	hashed, err := userService.EncrypterService.Hash(createUserRequestDTO.Password)
	if err != nil {
		return DTOs.CreateUserResponseDTO{}, err
	}

	user := user.User{
		ID:        uuid,
		Name:      createUserRequestDTO.Name,
		Email:     createUserRequestDTO.Email,
		CPF:       createUserRequestDTO.CPF,
		Password:  hashed,
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

	// TODO: add Password encryption

	if user.ID == "" {
		return DTOs.FindUserByEmailResponseDTO{}, &services_errors.UserNotFoundError{
			Message: "user not found",
		}
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
