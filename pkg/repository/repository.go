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
	Create(userID int, input models.TodoList) (int, error)
	Delete(userID, listID int) error
	Update(userID, listID int, input models.TodoList) error
	GetByID(userID int, listID int) (models.TodoList, error)
	GetAll(userID int) ([]models.TodoList, error)
	GetOwnerID(listID int) (int, error)
}

type TodoElement interface {
	Create(userID int, input models.TodoElement) (int, error)
	Delete(userID, listID int) error
	// Update(userID, listID int, input models.TodoList) error
	// GetByID(userID int, listID int) (models.TodoList, error)
	// GetAll(userID int) ([]models.TodoList, error)
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
