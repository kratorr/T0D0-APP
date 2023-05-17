package repository

import (
	"todo/models"

	"github.com/jackc/pgx/v4"
)

type Auth interface {
	CreateUser(models.User) (models.User, error)
	GetUser(login string) (models.User, error)
	SaveToken(models.User, string) error
	GetUserByToken(token string) (models.User, error)
}

type TodoList interface {
	Create(userID int, input models.CreateTodoListDTO) (int, error)
	Delete(listID int) error
	Update(userID, listID int, input models.TodoList) error
	GetByID(userID, listID int) (models.TodoList, error)
	GetAll(userID int) ([]models.TodoList, error)
}

type TodoElement interface {
	Create(userID int, input models.TodoElement) (int, error)
	Delete(elementID int) error
	Update(elementID int, input models.TodoElement) error
	GetAllByListID(listID int) ([]models.TodoElement, error)
}

type Repository struct {
	Auth
	TodoList
	TodoElement
}

func NewRepository(db *pgx.Conn) *Repository {
	return &Repository{
		Auth:        NewAuthPostgres(db), // DB connection as param
		TodoList:    NewTodoListPostgres(db),
		TodoElement: NewTodoElementPostgres(db),
	}
}
