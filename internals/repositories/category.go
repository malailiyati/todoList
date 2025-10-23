package repositories

import (
	"errors"
	"strings"

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
	// Validasi nama unik
	var existing models.Category
	err := r.DB.Where("LOWER(name) = ?", strings.ToLower(category.Name)).First(&existing).Error
	if err == nil {
		return errors.New("category name already exists")
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	return r.DB.Create(category).Error
}

func (r *CategoryRepository) Update(id uint, data *models.Category) (*models.Category, error) {
	var existing models.Category
	if err := r.DB.First(&existing, id).Error; err != nil {
		return nil, errors.New("category not found")
	}

	// Kalau color kosong, pertahankan yang lama
	if data.Color == "" {
		data.Color = existing.Color
	}

	if err := r.DB.Model(&existing).Updates(data).Error; err != nil {
		return nil, err
	}

	return &existing, nil
}

func (r *CategoryRepository) Delete(id uint) error {
	var category models.Category
	if err := r.DB.First(&category, id).Error; err != nil {
		return errors.New("category not found")
	}
	return r.DB.Delete(&category).Error
}

func (r *CategoryRepository) GetByID(id uint) (*models.Category, error) {
	var category models.Category
	err := r.DB.First(&category, id).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *TodoRepository) CheckCategoryExists(categoryID uint) (bool, error) {
	var exists bool
	err := r.DB.Model(&models.Category{}).
		Select("count(*) > 0").
		Where("id = ?", categoryID).
		Find(&exists).Error
	return exists, err
}

func (r *CategoryRepository) CheckNameExists(name string) (bool, error) {
	var exists bool
	err := r.DB.Model(&models.Category{}).
		Select("count(*) > 0").
		Where("LOWER(name) = LOWER(?)", name).
		Find(&exists).Error
	return exists, err
}

func (r *CategoryRepository) FindByID(id uint) (*models.Category, error) {
	var category models.Category
	err := r.DB.First(&category, id).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}
