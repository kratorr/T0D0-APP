package service

import (
	"errors"
	"math/rand"
	"time"

	"todo/models"
	"todo/pkg/repository"

	"golang.org/x/crypto/bcrypt"
)

func NewAuthService(repo repository.Auth) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

type AuthService struct {
	repo repository.Auth
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

func (s *AuthService) SignIn(user models.User) (string, error) {
	userDB, err := s.repo.GetUser(user.Login) // user from DB
	if err != nil {
		return "", errors.New("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(userDB.Password), []byte(user.Password))
	if err != nil {
		return "", errors.New("authentication failed")
	}

	token := s.CreateToken()

	err = s.repo.SaveToken(userDB, token)

	if err != nil {
		return "", errors.New("error save token")
	}

	return token, nil
}

func (s *AuthService) GetUser(login string) (models.User, error) {
	return s.repo.GetUser(login)
}

func (s *AuthService) GetUserByToken(token string) (models.User, error) {
	return s.repo.GetUserByToken(token)
}

func (s *AuthService) CreateToken() string {
	rand.Seed(time.Now().UnixNano())

	letterRunes := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, 40)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	return string(b)
}
