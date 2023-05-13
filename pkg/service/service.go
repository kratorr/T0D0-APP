package service

import (
	"todo/models"
	"todo/pkg/repository"
)

type Auth interface {
	SignUp(models.CreateUserDTO) error
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
	Create(userID int, input models.TodoElement) (int, error) // input insert
	Delete(elementID int) error
	Update(userID int, input models.TodoElement) error
	GetAll(listID int) ([]models.TodoElement, error)
}

type Service struct {
	Auth
	TodoList
	TodoElement
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Auth:        NewAuthService(repos.Auth),
		TodoList:    NewTodoListService(repos.TodoList),
		TodoElement: NewTodoElementService(repos.TodoElement),
	}
}
