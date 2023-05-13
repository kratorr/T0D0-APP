package service

import (
	"todo/models"
	"todo/pkg/repository"
)

func NewTodoElementService(repo repository.TodoElement) *TodoElementService {
	return &TodoElementService{
		repo: repo,
	}
}

type TodoElementService struct {
	repo repository.TodoElement
}

func (s *TodoElementService) Create(userID int, input models.TodoElement) (int, error) {
	return s.repo.Create(userID, input)
}

func (s *TodoElementService) Delete(elementID int) error {
	return s.repo.Delete(elementID)
}

func (s *TodoElementService) Update(userID int, input models.TodoElement) error {
	return s.repo.Update(userID, input)
}

func (s *TodoElementService) GetAll(listID int) ([]models.TodoElement, error) {
	return s.repo.GetAllByListID(listID)
}
