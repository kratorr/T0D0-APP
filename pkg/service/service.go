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
	Create(userID int, input models.TodoList) (int, error)
	Delete(userID int, listID int) error
	Update(userID, listID int, input models.TodoList) error
	GetByID(userID int, listID int) (models.TodoList, error)
	GetAll(userID int) ([]models.TodoList, error)
	GetOwnerID(listID int) (int, error)
}

type TodoElement interface {
	Create(userID int, listID int) error // input insert
	Delete(userID int) error
	Update(userID int) error
	GetByID(userID int, ID int) error
	GetAll(userID int) error
}

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
