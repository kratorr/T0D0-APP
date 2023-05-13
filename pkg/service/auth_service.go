package service

import (
	"errors"
	"fmt"
	"time"

	"todo/models"
	"todo/pkg/repository"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func NewAuthService(repo repository.Auth, secretKey string) *AuthService {
	return &AuthService{
		repo:      repo,
		secretKey: secretKey,
	}
}

type AuthService struct {
	repo      repository.Auth
	secretKey string
}

func (s *AuthService) SignUp(userDto models.CreateUserDTO) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(userDto.Password), 1)
	if err != nil {
		return errors.New("error create user")
	} // TODO validate files, shit password etc. Задачка вроде изян, пока на паузе.

	userDto.Password = string(passwordHash)

	user := models.User{
		Login:    userDto.Login,
		Password: userDto.Password,
	}

	_, err = s.repo.CreateUser(user)
	if err != nil {
		return errors.New(err.Error())
	}

	return nil
}

func (s *AuthService) SignIn(userInput models.SignInUserDTO) (string, error) {
	user, err := s.repo.GetUser(userInput.Login) // user from DB
	if err != nil {
		return "", errors.New("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userInput.Password))

	if err != nil {
		return "", errors.New("authentication failed")
	}

	payload := jwt.MapClaims{
		"sub":      user.ID,
		"nickname": user.Login,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	t, err := token.SignedString([]byte(s.secretKey))
	if err != nil {
		fmt.Println("error jwt signing", err)
		return "", errors.New("error jwt signing")
	}

	return t, nil
}

func (s *AuthService) GetUser(login string) (models.User, error) {
	return s.repo.GetUser(login)
}

func (s *AuthService) GetUserByToken(token string) (models.User, error) {
	return s.repo.GetUserByToken(token)
}

func (s *AuthService) CreateToken() string {
	return ""
}
