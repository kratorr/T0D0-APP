package repository

import (
	"todo/models"

	"github.com/jackc/pgx/v4"
)

type TodoElementPostgres struct {
	db *pgx.Conn
}

func NewTodoElementPostgres(db *pgx.Conn) *TodoElementPostgres {
	return &TodoElementPostgres{db: db}
}

func (r *TodoElementPostgres) Create(userID int, input models.TodoElement) (int, error) {
	return 0, nil
}

func (r *TodoElementPostgres) Delete(userID, listID int) error {
	return nil
}

func (r *TodoElementPostgres) GetAll(userID, listID int) error {
	return nil
}
