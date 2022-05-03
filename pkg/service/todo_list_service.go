package service

import (
	"todo/pkg/repository"
)

func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{
		repo: repo,
	}
}

type TodoListService struct {
	repo repository.TodoList
}

func (s *TodoListService) Create(userID int) error {
	return nil
}

func (s *TodoListService) Delete(userID int) error {
	return nil
}

func (s *TodoListService) Update(userID int) error {
	return nil
}

func (s *TodoListService) GetAll(userID int) error {
	return nil
}

func (s *TodoListService) GetByID(userID, id int) error {
	return nil
}
