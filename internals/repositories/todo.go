package repositories

import (
	"github.com/malailiyati/todoList/internals/models"
	"gorm.io/gorm"
)

type TodoRepository struct {
	DB *gorm.DB
}

func NewTodoRepository(db *gorm.DB) *TodoRepository {
	return &TodoRepository{DB: db}
}

func (r *TodoRepository) GetAll(search string, limit, offset int) ([]models.Todo, int64, error) {
	var todos []models.Todo
	var total int64

	query := r.DB.Preload("Category").Model(&models.Todo{})

	if search != "" {
		query = query.Where("title ILIKE ? OR description ILIKE ?", "%"+search+"%", "%"+search+"%")
	}

	// ✅ urutkan: belum selesai dulu, lalu priority high > medium > low, terakhir id terbaru
	query = query.
		Order("completed ASC").
		Order(`
            CASE 
                WHEN priority = 'high' THEN 1
                WHEN priority = 'medium' THEN 2
                WHEN priority = 'low' THEN 3
                ELSE 4
            END
        `).
		Order("id DESC")

	query.Count(&total)
	err := query.Limit(limit).Offset(offset).Find(&todos).Error
	return todos, total, err
}

func (r *TodoRepository) Create(todo *models.Todo) error {
	return r.DB.Create(todo).Error
}

func (r *TodoRepository) Update(id uint, data *models.Todo) (*models.Todo, error) {
	result := r.DB.Model(&models.Todo{}).
		Where("id = ?", id).
		Select("title", "description", "priority", "completed", "category_id").
		Updates(data)
	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	if result.Error != nil {
		return nil, result.Error
	}

	var updatedTodo models.Todo
	err := r.DB.Preload("Category").First(&updatedTodo, id).Error
	if err != nil {
		return nil, err
	}

	return &updatedTodo, nil
}

// Delete
func (r *TodoRepository) Delete(id uint) error {
	result := r.DB.Delete(&models.Todo{}, id)
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return result.Error
}

// FindByID — biar bisa dipakai validasi di handler
func (r *TodoRepository) FindByID(id uint) (*models.Todo, error) {
	var todo models.Todo
	err := r.DB.Preload("Category").First(&todo, id).Error
	if err != nil {
		return nil, err
	}
	return &todo, nil
}
