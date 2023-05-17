package repository

import (
	"context"
	"fmt"

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

func (r *TodoListPostgres) Create(userID int, input models.CreateTodoListDTO) (int, error) {
	var todoListID int

	query := `INSERT INTO todo_lists (user_id, title, description) VALUES($1, $2, $3) RETURNING id;`
	err := r.db.QueryRow(context.Background(), query, userID, input.Title, input.Description).Scan(&todoListID)
	if err != nil {
		zap.L().Sugar().Error(err.Error())

		return todoListID, err
	}

	return todoListID, nil
}

func (r *TodoListPostgres) Delete(listID int) error {
	// query := `DELETE FROM todo_lists WHERE user_id = $1 and id = $2 `
	query := `DELETE FROM todo_lists WHERE id = $1;`

	_, err := r.db.Exec(context.Background(), query, listID)
	if err != nil {
		zap.L().Sugar().Error(err.Error())

		return err
	}

	return nil
}

func (r *TodoListPostgres) Update(userID, listID int, input models.TodoList) error {
	fmt.Println(listID, userID, input)
	query := `UPDATE todo_lists SET title = $1, description  = $2 WHERE id = $3 and user_id = $4;`
	_, err := r.db.Exec(context.Background(), query, input.Title, input.Description, listID, userID)
	if err != nil {
		zap.L().Sugar().Error(err.Error())

		return err
	}
	return nil
}

func (r *TodoListPostgres) GetAll(userID int) ([]models.TodoList, error) {
	result := make([]models.TodoList, 0)

	query := `SELECT id, user_id, title, description FROM todo_lists WHERE user_id = $1;`
	rows, err := r.db.Query(context.Background(), query, userID)
	if err != nil {
		zap.L().Sugar().Error(err.Error())
		return result, err
	}

	for rows.Next() {
		todoList := models.TodoList{}
		err := rows.Scan(&todoList.ID, &todoList.UserID, &todoList.Title, &todoList.Description)
		if err != nil {
			fmt.Println(err)
		}
		result = append(result, todoList)
	}

	return result, nil
}

func (r *TodoListPostgres) GetByID(userID, listID int) (models.TodoList, error) {
	result := models.TodoList{}
	query := `SELECT id, user_id, title, description FROM todo_lists WHERE user_id = $1 and id = $2;`
	err := r.db.QueryRow(context.Background(), query, userID, listID).Scan(&result.ID, &result.UserID, &result.Title, &result.Description)
	if err != nil {
		zap.L().Sugar().Error(err.Error())

		return result, err
	}

	return result, nil
}
