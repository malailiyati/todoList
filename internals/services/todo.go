package services

import (
	"errors"
	"strings"

	"github.com/malailiyati/todoList/internals/models"
	"github.com/malailiyati/todoList/internals/repositories"
)

type TodoService struct {
	todoRepo     *repositories.TodoRepository
	categoryRepo *repositories.CategoryRepository
}

func NewTodoService(todoRepo *repositories.TodoRepository, categoryRepo *repositories.CategoryRepository) *TodoService {
	return &TodoService{todoRepo, categoryRepo}
}

func newPagination(page, limit, total int) models.Pagination {
	totalPages := (total + limit - 1) / limit
	return models.Pagination{
		CurrentPage: page,
		PerPage:     limit,
		Total:       total,
		TotalPages:  totalPages,
	}
}

func (s *TodoService) validateCommon(todo *models.Todo) error {
	validPriorities := map[string]bool{"low": true, "medium": true, "high": true}
	if todo.Priority != "" && !validPriorities[strings.ToLower(todo.Priority)] {
		return errors.New("invalid priority (must be high, medium, or low)")
	}

	if todo.CategoryID != nil && *todo.CategoryID != 0 {
		exists, err := s.todoRepo.CheckCategoryExists(*todo.CategoryID)
		if err != nil {
			return err
		}
		if !exists {
			return errors.New("invalid category_id")
		}
	}
	return nil
}

func (s *TodoService) validateCreate(todo *models.Todo) error {
	if strings.TrimSpace(todo.Title) == "" {
		return errors.New("title cannot be empty")
	}
	return s.validateCommon(todo)
}

func (s *TodoService) validateUpdate(todo *models.Todo) error {
	// kalau title dikirim tapi kosong → error
	if todo.Title != "" && strings.TrimSpace(todo.Title) == "" {
		return errors.New("title cannot be empty")
	}
	return s.validateCommon(todo)
}

func (s *TodoService) GetAll(search string, page, limit int) (map[string]interface{}, error) {
	offset := (page - 1) * limit
	todos, total, err := s.todoRepo.GetAll(search, limit, offset)
	if err != nil {
		return nil, err
	}

	pagination := newPagination(page, limit, int(total))
	return map[string]interface{}{
		"todos":      todos,
		"pagination": pagination,
	}, nil
}

func (s *TodoService) GetByID(id uint) (*models.Todo, error) {
	todo, err := s.todoRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("todo not found")
	}
	return todo, nil
}

func (s *TodoService) Create(todo *models.Todo) error {
	if err := s.validateCreate(todo); err != nil {
		return err
	}
	return s.todoRepo.Create(todo)
}

func (s *TodoService) Update(id uint, todo *models.Todo) (*models.Todo, error) {
	existing, err := s.todoRepo.FindByID(id)
	if err != nil || existing == nil {
		return nil, errors.New("todo not found")
	}

	if err := s.validateUpdate(todo); err != nil {
		return nil, err
	}

	updated, err := s.todoRepo.Update(id, todo)
	if err != nil {
		return nil, err
	}
	return updated, nil
}

func (s *TodoService) ToggleComplete(id uint) (*models.Todo, error) {
	todo, err := s.todoRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("todo not found")
	}

	todo.Completed = !todo.Completed
	updated, err := s.todoRepo.Update(id, todo)
	if err != nil {
		return nil, err
	}
	return updated, nil
}

func (s *TodoService) Delete(id uint) error {
	existing, err := s.todoRepo.FindByID(id)
	if err != nil || existing == nil {
		return errors.New("todo not found")
	}
	return s.todoRepo.Delete(id)
}
