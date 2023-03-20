package notes

import (
	"notes-api/drivers/mysql/categories"
	"time"

	"gorm.io/gorm"
)

type Note struct {
	ID         uint                `json:"id" gorm:"primaryKey"`
	CreatedAt  time.Time           `json:"created_at"`
	UpdatedAt  time.Time           `json:"updated_at"`
	DeletedAt  gorm.DeletedAt      `json:"deleted_at"`
	Title      string              `json:"title"`
	Content    string              `json:"content"`
	Category   categories.Category `json:"category"`
	CategoryID uint                `json:"category_id"`
}
