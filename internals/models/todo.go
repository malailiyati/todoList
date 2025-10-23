package models

import "time"

type Todo struct {
	ID          uint       `gorm:"primaryKey" json:"id"`
	Title       string     `gorm:"size:255" json:"title"`
	Description string     `json:"description"`
	Completed   bool       `gorm:"default:false" json:"completed"`
	CategoryID  uint       `json:"category_id"`
	Priority    string     `gorm:"size:10" json:"priority"`
	DueDate     *time.Time `json:"due_date,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	Category    Category   `json:"category" gorm:"foreignKey:CategoryID"`
}
