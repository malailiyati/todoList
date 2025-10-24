package repositories

import (
	"errors"

	"github.com/malailiyati/todoList/internals/models"

	"gorm.io/gorm"
)

type CategoryRepository struct {
	DB *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{DB: db}
}

func (r *CategoryRepository) GetAll() ([]models.Category, error) {
	var categories []models.Category
	err := r.DB.Find(&categories).Error
	return categories, err
}

func (r *CategoryRepository) Create(category *models.Category) error {
	return r.DB.Create(category).Error
}

func (r *CategoryRepository) Update(id uint, data *models.Category) (*models.Category, error) {
	var existing models.Category
	if err := r.DB.First(&existing, id).Error; err != nil {
		return nil, errors.New("category not found")
	}

	if data.Color == "" {
		data.Color = existing.Color
	}

	if err := r.DB.Model(&existing).Updates(data).Error; err != nil {
		return nil, err
	}

	return &existing, nil
}

func (r *CategoryRepository) Delete(id uint) error {
	return r.DB.Delete(&models.Category{}, id).Error
}

func (r *CategoryRepository) FindByID(id uint) (*models.Category, error) {
	var category models.Category
	err := r.DB.First(&category, id).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *CategoryRepository) ExistsByName(name string) (bool, error) {
	var exists bool
	err := r.DB.Model(&models.Category{}).
		Select("count(*) > 0").
		Where("LOWER(name) = LOWER(?)", name).
		Find(&exists).Error
	return exists, err
}
