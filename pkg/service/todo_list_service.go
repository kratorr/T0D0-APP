package service

import (
	"todo/models"
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

func (s *TodoListService) Create(userID int, input models.TodoList) (int, error) {
	return s.repo.Create(userID, input)
}

func (s *TodoListService) Delete(userID int, listID int) error {
	return s.repo.Delete(userID, listID)
}

func (s *TodoListService) Update(userID, listID int, input models.TodoList) error {
	return s.repo.Update(userID, listID, input)
}

func (s *TodoListService) GetAll(userID int) ([]models.TodoList, error) {
	return s.repo.GetAll(userID)
}

func (s *TodoListService) GetByID(userID, listID int) (models.TodoList, error) {
	return s.repo.GetByID(userID, listID)
}

func (s *TodoListService) GetOwnerID(listID int) (int, error) {
	return s.repo.GetOwnerID(listID)
}
