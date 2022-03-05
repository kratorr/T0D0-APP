package repository

import (
	"todo/models"

	"github.com/jackc/pgx/v4"
)

type Auth interface {
	CreateUser(models.User) error
	GetUser(login string) (models.User, error)
	SaveToken(models.User, string) error
	GetUserByToken(token string) (models.User, error)
}

type Repository struct {
	Auth
}

func NewRepository(db *pgx.Conn) *Repository {
	return &Repository{
		Auth: NewAuthPostgres(db), // DB connection as param
	}
}
