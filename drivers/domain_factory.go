package drivers

import (
	categoryDomain "notes-api/businesses/categories"
	categoryDB "notes-api/drivers/mysql/categories"

	noteDomain "notes-api/businesses/notes"
	noteDB "notes-api/drivers/mysql/notes"

	userDomain "notes-api/businesses/users"
	userDB "notes-api/drivers/mysql/users"

	"gorm.io/gorm"
)

func NewCategoryRepository(conn *gorm.DB) categoryDomain.Repository {
	return categoryDB.NewMySQLRepository(conn)
}

func NewNoteRepository(conn *gorm.DB) noteDomain.Repository {
	return noteDB.NewMySQLRepository(conn)
}

// create factory for user repository
func NewUserRepository(conn *gorm.DB) userDomain.Repository {
	return userDB.NewMySQLRepository(conn)
}
