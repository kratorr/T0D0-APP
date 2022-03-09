package service

import (
	"todo/models"
	"todo/pkg/repository"
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

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Auth: NewAuthService(repos.Auth),
	}
}
