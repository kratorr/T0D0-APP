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
	return 0, nil
	// return s.repo.Create(userID, input)
}

func (s *TodoElementService) Delete(userID int, input models.TodoElement) (int, error) {
	return s.repo.Create(userID, input)
}
