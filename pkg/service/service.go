package service

import (
	"errors"
	"fmt"
	"todo/models"
	"todo/pkg/repository"

	"golang.org/x/crypto/bcrypt"
)

const salt = "gewstrjgruoieapjfhgeawrjfuiraphjfpurahjfvuraehwf9ahj94-vanrunpier4-"

type Auth interface {
	SignUp(models.User) error
	SignIn(models.User) (string, error)
	GetUser(ID string) (models.User, error)
	CreateToken(models.User) (string, error)
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
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(u.Password), 0)
	if err != nil {
		return errors.New("error") // TODO здесь ошибку понятную
	}

	u.Password = string(passwordHash)

	s.repo.CreateUser(u)

	return nil
}

func (s *AuthService) SignIn(u models.User) (string, error) {
	user, _ := s.repo.GetUser(u.Login) // user from DB

	fmt.Println(user.Password, "DB password")
	fmt.Println(string(u.Password), " REQUEST PASSWORD ")

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(u.Password))
	if err != nil {
		return "", errors.New("shit password")
	}

	token, err := s.CreateToken(u)
	if err != nil {
		return "", errors.New("error") // TODO сделать нормальную ошибку
	}

	return token, nil
}

func (s *AuthService) GetUser(ID string) (models.User, error) {
	return s.repo.GetUser(ID)
}

func (s *AuthService) CreateToken(u models.User) (string, error) {
	return s.repo.CreateToken(u)
}
