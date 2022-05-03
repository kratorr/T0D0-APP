package repository

import "github.com/jackc/pgx/v4"

type TodoListPostgres struct {
	db *pgx.Conn
}

func NewTodoListPostgres(db *pgx.Conn) *TodoListPostgres {
	return &TodoListPostgres{db: db}
}

func (r *TodoListPostgres) Create(userID int) error {
	return nil
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
