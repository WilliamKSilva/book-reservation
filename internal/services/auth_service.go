package services

import (
	"github.com/WilliamKSilva/book-reservation/internal/services/DTOs"
	services_errors "github.com/WilliamKSilva/book-reservation/internal/services/errors"
)

type IAuthService interface {
	Login(email string, password string) (DTOs.LoginResponseDTO, error)
	Register(registerRequestDTO DTOs.RegisterRequestDTO) (DTOs.RegisterResponseDTO, error)
}

type AuthService struct {
	JwtService  JwtServiceAdapter
	UserService IUserService
}

func (authService *AuthService) Login(email string, password string) (DTOs.LoginResponseDTO, error) {
	user, err := authService.UserService.FindByEmail(email)
	if err != nil {
		return DTOs.LoginResponseDTO{}, err
	}

	// TODO: add password validation based on encryption that will be added
	if user.Password != password {
		res := DTOs.LoginResponseDTO{
			User:        DTOs.LoginUser{},
			AccessToken: DTOs.JwtToken{},
		}

		return res, &services_errors.WrongPasswordError{}
	}

	accessToken, err := authService.JwtService.New()
	if err != nil {
		res := DTOs.LoginResponseDTO{
			User:        DTOs.LoginUser{},
			AccessToken: DTOs.JwtToken{},
		}

		return res, err
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
