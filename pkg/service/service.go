package service

import (
	"errors"
	"math/rand"
	"time"
	"todo/models"
	"todo/pkg/repository"

	"golang.org/x/crypto/bcrypt"
)

type Auth interface {
	SignUp(models.User) error
	SignIn(models.User) (string, error)
	GetUser(login string) (models.User, error)
	GetUserByToken(token string) (models.User, error)
}

type Service struct {
	Auth
}

type AuthService struct {
	repo repository.Auth
}

func NewAuthService(repo repository.Auth) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Auth: NewAuthService(repos.Auth),
	}
}

func (s *AuthService) SignUp(user models.User) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 1)
	if err != nil {
		return errors.New("Error create user")
	} // TODO validate files, shit password etc.

	user.Password = string(passwordHash)

	err = s.repo.CreateUser(user)
	if err != nil {
		return errors.New("DB error")
	}

	return nil
}

func (s *AuthService) SignIn(user models.User) (string, error) {
	userDB, err := s.repo.GetUser(user.Login) // user from DB
	if err != nil {
		return "", errors.New("User not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(userDB.Password), []byte(user.Password))
	if err != nil {
		return "", errors.New("Password or login is shit")
	}

	token := s.CreateToken()

	err = s.repo.SaveToken(userDB, token)

	if err != nil {
		return "", errors.New("Error save token")
	}

	return token, nil
}

func (s *AuthService) GetUser(Login string) (models.User, error) {
	return s.repo.GetUser(Login)
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

func (s *AuthService) GetUserByToken(token string) (models.User, error) {
	return s.repo.GetUserByToken(token)
}
