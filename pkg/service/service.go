package service

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
	"todo/models"
	"todo/pkg/repository"

	"golang.org/x/crypto/bcrypt"
)

type Auth interface {
	SignUp(models.User) error
	SignIn(models.User) (string, error)
	GetUser(ID string) (models.User, error)
	// CreateToken(models.User) (string, error)
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

func (s *AuthService) SignUp(u models.User) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(u.Password), 1)
	if err != nil {
		return errors.New("error") // TODO здесь ошибку понятную
	} // TODO validate files, shit password etc.

	u.Password = string(passwordHash)

	err = s.repo.CreateUser(u)
	if err != nil {
		return errors.New("DB error")
	}

	return nil
}

func (s *AuthService) SignIn(u models.User) (string, error) {
	user, err := s.repo.GetUser(u.Login) // user from DB
	if err != nil {
		return "", errors.New("User not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(u.Password))
	if err != nil {
		fmt.Println(err)
		return "", errors.New("Password or login is shit")
	}

	token := s.CreateToken()

	err = s.repo.SaveToken(user, token)

	if err != nil {
		return "", errors.New("Error save token")
	}

	return token, nil
}

func (s *AuthService) GetUser(ID string) (models.User, error) {
	return s.repo.GetUser(ID)
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
