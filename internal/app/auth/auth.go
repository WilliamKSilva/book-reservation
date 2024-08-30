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
	Validate(accessToken string) (bool, error)
}

type IAuthService interface {
	Login(accessToken string) (user.User, error)
}

type AuthService struct {
	JwtService     IJwtService
	UserRepository user.IUserRepository
}

type LoginDTO struct {
	User        user.User
	AccessToken JwtToken
}

func (authService *AuthService) Login(email string, password string) (LoginDTO, error) {
	user, err := authService.UserRepository.FindByEmail(email)
	if err != nil {
		return LoginDTO{user, JwtToken{}}, err
	}

	if user.Password != password {
		return LoginDTO{user, JwtToken{}}, errors.New("wrong password")
	}

	accessToken, err := authService.JwtService.New()
	if err != nil {
		return LoginDTO{user, JwtToken{}}, err
	}

	return LoginDTO{user, accessToken}, nil
}
