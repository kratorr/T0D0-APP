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
	Create(userID int) error
	Delete(userID int) error
	Update(userID int) error
	GetByID(userID int, ID int) error
	GetAll(userID int) error
}

type Repository struct {
	Auth
	TodoList
}

func NewRepository(db *pgx.Conn) *Repository {
	return &Repository{
		Auth:     NewAuthPostgres(db), // DB connection as param
		TodoList: NewTodoListPostgres(db),
	}
}
