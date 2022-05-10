package repository

import (
	"context"

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
		return todoListId, err
	}

	return todoListId, nil
}

func (r *TodoListPostgres) Delete(userID, listID int) error {
	query := `DELETE FROM todo_lists WHERE user_id = $1 and id = $2 `
	_, err := r.db.Exec(context.Background(), query, userID, listID)
	if err != nil {

		zap.L().Sugar().Error(err.Error())

		return err
	}

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

func (r *TodoListPostgres) GetOwnerID(listID int) (int, error) {
	var ownerID int

	query := `SELECT user_id FROM todo_lists WHERE id = $1 ;`
	err := r.db.QueryRow(context.Background(), query, listID).Scan(&ownerID)
	if err != nil {

		zap.L().Sugar().Error(err.Error())
		return ownerID, err
	}

	return ownerID, nil
}
