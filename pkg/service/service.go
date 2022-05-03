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

type TodoList interface {
	Create(userID int) error
	Delete(userID int) error
	Update(userID int) error
	GetByID(userID int, ID int) error
	GetAll(userID int) error
}

type todoElement interface{}

type Service struct {
	Auth
	TodoList
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Auth:     NewAuthService(repos.Auth),
		TodoList: NewTodoListService(repos.TodoList),
	}
}
