package repository

import (
	"context"
	"fmt"

	"todo/models"

	"github.com/jackc/pgx/v4"
	"go.uber.org/zap"
)

type TodoElementPostgres struct {
	db *pgx.Conn
}

func NewTodoElementPostgres(db *pgx.Conn) *TodoElementPostgres {
	return &TodoElementPostgres{db: db}
}

func (r *TodoElementPostgres) Create(userID int, input models.TodoElement) (int, error) {
	var todoElementId int
	query := `INSERT INTO todo_element (todo_list_id, title) VALUES($1, $2) RETURNING id;;`
	err := r.db.QueryRow(context.Background(), query, input.TodoListID, input.Title).Scan(&todoElementId)
	if err != nil {

		zap.L().Sugar().Error(err.Error())
		return todoElementId, err
	}

	return todoElementId, nil
}

func (r *TodoElementPostgres) Delete(elementID int) error {
	query := `DELETE FROM todo_element WHERE id = $1;;`
	_, err := r.db.Exec(context.Background(), query, elementID)
	if err != nil {
		zap.L().Sugar().Error(err.Error())
		return err
	}

	return nil
}

func (r *TodoElementPostgres) Update(elementID int, input models.TodoElement) error {
	query := `UPDATE todo_element SET title = $1 WHERE id = $3`
	_, err := r.db.Exec(context.Background(), query, input.Title)
	if err != nil {
		zap.L().Sugar().Error(err.Error())
		return err
	}
	return nil
}

func (r *TodoElementPostgres) GetAllByListID(listID int) ([]models.TodoElement, error) {
	result := make([]models.TodoElement, 0)
	userID := 1
	query := `SELECT id, todo_list_id, title, state_id FROM todo_elemnt WHERE id = $1;`
	rows, err := r.db.Query(context.Background(), query, userID)
	if err != nil {
		zap.L().Sugar().Error(err.Error())
		return result, err
	}

	for rows.Next() {
		todoElement := models.TodoElement{}
		err := rows.Scan(&todoElement.ID, &todoElement.Title)
		if err != nil {
			fmt.Println(err)
		}
		result = append(result, todoElement)
	}

	return result, nil
}
