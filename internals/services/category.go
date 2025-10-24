package services

import (
	"errors"
	"regexp"
	"strings"

	"github.com/malailiyati/todoList/internals/models"
	"github.com/malailiyati/todoList/internals/repositories"
)

type CategoryService struct {
	categoryRepo *repositories.CategoryRepository
}

func NewCategoryService(categoryRepo *repositories.CategoryRepository) *CategoryService {
	return &CategoryService{categoryRepo}
}

func (s *CategoryService) Validate(category *models.Category) error {
	if category.Name == "" {
		return errors.New("category name cannot be empty")
	}
	if category.Color == "" {
		return errors.New("category color cannot be empty")
	}

	if !strings.HasPrefix(category.Color, "#") {
		return errors.New("invalid color format: must start with '#'")
	}
	if len(category.Color) != 4 && len(category.Color) != 7 {
		return errors.New("invalid color format: must be #RGB or #RRGGBB")
	}

	exists, err := s.categoryRepo.ExistsByName(category.Name)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("category name already exists")
	}

	return nil
}

func (s *CategoryService) GetAll() ([]models.Category, error) {
	return s.categoryRepo.GetAll()
}

func (s *CategoryService) Create(category *models.Category) error {
	if err := s.Validate(category); err != nil {
		return err
	}
	return s.categoryRepo.Create(category)
}

func (s *CategoryService) Update(id uint, category *models.Category) (*models.Category, error) {
	existing, err := s.categoryRepo.FindByID(id)
	if err != nil || existing == nil {
		return nil, errors.New("category not found")
	}

	if category.Name != "" && category.Name != existing.Name {
		exists, _ := s.categoryRepo.ExistsByName(category.Name)
		if exists {
			return nil, errors.New("category name already exists")
		}
	}

	if category.Color != "" {
		matched, _ := regexp.MatchString(`^(#(?:[0-9a-fA-F]{3}){1,2}|[a-zA-Z]+)$`, category.Color)
		if !matched {
			return nil, errors.New("invalid color format (must be hex like #FF0000 or color name like red)")
		}
	}

	updated, err := s.categoryRepo.Update(id, category)
	if err != nil {
		return nil, err
	}
	return updated, nil
}

func (s *CategoryService) Delete(id uint) error {
	existing, err := s.categoryRepo.FindByID(id)
	if err != nil || existing == nil {
		return errors.New("category not found")
	}
	return s.categoryRepo.Delete(id)
}
