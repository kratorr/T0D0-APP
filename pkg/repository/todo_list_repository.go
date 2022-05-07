package repository

import (
	"context"
	"errors"
	"strings"

	"todo/models"

	"github.com/jackc/pgx/v4"
	"go.uber.org/zap"
)

type TodoListPostgres struct {
	db *pgx.Conn
}

func NewTodoListPostgres(db *pgx.Conn) *TodoListPostgres {
	return &TodoListPostgres{db: db}
}

func (r *TodoListPostgres) Create(userID int, input models.TodoList) (int, error) {
	var todoListId int

	query := `INSERT INTO todo_lists (user_id, title, description) VALUES($1, $2, $3) RETURNING id;;`
	err := r.db.QueryRow(context.Background(), query, userID, input.Title, input.Description).Scan(&todoListId)
	if err != nil {

		zap.L().Sugar().Error(err.Error())

		if strings.Contains(err.Error(), "duplicate key") {
			return todoListId, errors.New("user already exists")
		}

		return todoListId, err
	}

	return todoListId, nil
}

func (r *TodoListPostgres) Delete(userID int) error {
	return nil
}

func (r *TodoListPostgres) Update(userID int) error {
	return nil
}

func (r *TodoListPostgres) GetAll(userID int) error {
	return nil
}

func (r *TodoListPostgres) GetByID(userID, id int) error {
	return nil
}
