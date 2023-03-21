package drivers

import (
	categoryDomain "notes-api/businesses/categories"
	categoryDB "notes-api/drivers/mysql/categories"

	noteDomain "notes-api/businesses/notes"
	noteDB "notes-api/drivers/mysql/notes"

	"gorm.io/gorm"
)

func NewCategoryRepository(conn *gorm.DB) categoryDomain.Repository {
	return categoryDB.NewMySQLRepository(conn)
}

func NewNoteRepository(conn *gorm.DB) noteDomain.Repository {
	return noteDB.NewMySQLRepository(conn)
}
