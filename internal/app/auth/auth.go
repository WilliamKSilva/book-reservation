package auth

import (
	"errors"

	"github.com/WilliamKSilva/book-reservation/internal/app/user"
)

type JwtToken struct {
	Raw    string
	Signed string
}

type IJwtService interface {
	New() (JwtToken, error)
}

type IAuthService interface {
	Login(email string, password string) (LoginResponseDTO, error)
	Register(registerRequestDTO RegisterRequestDTO) (RegisterResponseDTO, error)
}

type AuthService struct {
	JwtService  IJwtService
	UserService user.IUserService
}

func (authService *AuthService) Login(email string, password string) (LoginResponseDTO, error) {
	user, err := authService.UserService.FindByEmail(email)
	if err != nil {
		loginResponse := LoginResponseDTO{
			LoginUser{
				ID:        user.ID,
				Name:      user.Name,
				Email:     user.Email,
				Password:  user.Password,
				CPF:       user.CPF,
				BirthDate: user.BirthDate,
			},
			JwtToken{},
		}

		return loginResponse, err
	}

	if user.Password != password {
		loginResponse := LoginResponseDTO{
			LoginUser{
				ID:        user.ID,
				Name:      user.Name,
				Email:     user.Email,
				Password:  user.Password,
				CPF:       user.CPF,
				BirthDate: user.BirthDate,
			},
			JwtToken{},
		}

		return loginResponse, errors.New("wrong password")
	}

	accessToken, err := authService.JwtService.New()
	if err != nil {
		loginResponse := LoginResponseDTO{
			LoginUser{
				ID:        user.ID,
				Name:      user.Name,
				Email:     user.Email,
				Password:  user.Password,
				CPF:       user.CPF,
				BirthDate: user.BirthDate,
			},
			JwtToken{},
		}

		return loginResponse, err
	}

	loginResponse := LoginResponseDTO{
		LoginUser{
			ID:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			Password:  user.Password,
			CPF:       user.CPF,
			BirthDate: user.BirthDate,
		},
		accessToken,
	}
	return loginResponse, nil
}

func (authService *AuthService) Register(registerRequestDTO RegisterRequestDTO) (RegisterResponseDTO, error) {
	createUserRequestDTO := user.CreateUserRequestDTO{
		Name:      registerRequestDTO.Name,
		Email:     registerRequestDTO.Email,
		Password:  registerRequestDTO.Password,
		CPF:       registerRequestDTO.CPF,
		BirthDate: registerRequestDTO.BirthDate,
	}

	user, err := authService.UserService.Create(createUserRequestDTO)
	if err != nil {
		return RegisterResponseDTO{}, err
	}

	loginResponse, err := authService.Login(user.Email, user.Password)
	if err != nil {
		return RegisterResponseDTO{}, err
	}

	registerResponse := RegisterResponseDTO{
		User: RegisterUser{
			ID:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			Password:  user.Password,
			CPF:       user.CPF,
			BirthDate: user.BirthDate,
		},
		AccessToken: loginResponse.AccessToken,
	}

	return registerResponse, nil
}
