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

func (s *TodoListService) Update(userID int, input models.TodoList) (int, error) {
	return 0, nil
}

func (s *TodoListService) GetAll(userID int) ([]models.TodoList, error) {
	res := []models.TodoList{}
	return res, nil
}

func (s *TodoListService) GetByID(userID, id int) (models.TodoList, error) {
	res := models.TodoList{}
	return res, nil
}

func (s *TodoListService) GetOwnerID(listID int) (int, error) {
	return s.repo.GetOwnerID(listID)
}
