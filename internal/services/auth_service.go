package services

import (
	"errors"

	"github.com/WilliamKSilva/book-reservation/internal/services/DTOs"
)

type IJwtService interface {
	New() (DTOs.JwtToken, error)
}

type IAuthService interface {
	Login(email string, password string) (DTOs.LoginResponseDTO, error)
	Register(registerRequestDTO DTOs.RegisterRequestDTO) (DTOs.RegisterResponseDTO, error)
}

type AuthService struct {
	JwtService  IJwtService
	UserService IUserService
}

func (authService *AuthService) Login(email string, password string) (DTOs.LoginResponseDTO, error) {
	user, err := authService.UserService.FindByEmail(email)
	if err != nil {
		loginResponse := DTOs.LoginResponseDTO{
			User:        DTOs.LoginUser(user),
			AccessToken: DTOs.JwtToken{},
		}

		return loginResponse, err
	}

	if user.Password != password {
		loginResponse := DTOs.LoginResponseDTO{
			User:        DTOs.LoginUser(user),
			AccessToken: DTOs.JwtToken{},
		}

		return loginResponse, errors.New("wrong password")
	}

	accessToken, err := authService.JwtService.New()
	if err != nil {
		loginResponse := DTOs.LoginResponseDTO{
			User:        DTOs.LoginUser(user),
			AccessToken: DTOs.JwtToken{},
		}

		return loginResponse, err
	}

	loginResponse := DTOs.LoginResponseDTO{
		User:        DTOs.LoginUser(user),
		AccessToken: accessToken,
	}
	return loginResponse, nil
}

func (authService *AuthService) Register(registerRequestDTO DTOs.RegisterRequestDTO) (DTOs.RegisterResponseDTO, error) {
	createUserRequestDTO := DTOs.CreateUserRequestDTO(registerRequestDTO)

	user, err := authService.UserService.Create(createUserRequestDTO)
	if err != nil {
		return DTOs.RegisterResponseDTO{}, err
	}

	loginResponse, err := authService.Login(user.Email, user.Password)
	if err != nil {
		return DTOs.RegisterResponseDTO{}, err
	}

	registerResponse := DTOs.RegisterResponseDTO{
		User:        DTOs.RegisterUser(user),
		AccessToken: loginResponse.AccessToken,
	}

	return registerResponse, nil
}
