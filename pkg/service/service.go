package service

import (
	"errors"
	"todo/pkg/repository"
)

type Auth interface {
	SignUp() error
	SignIn() error
}

type Service struct {
	Auth
}

type AuthService struct{}

func (s *AuthService) SignIn() error {
	return errors.New("test")
}

func (s *AuthService) SignUp() error {
	return errors.New("test")
}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Auth: NewAuthService(),
	}
}
