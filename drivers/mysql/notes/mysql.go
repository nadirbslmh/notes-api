package notes

import (
	"context"
	"notes-api/businesses/notes"

	"gorm.io/gorm"
)

type noteRepository struct {
	conn *gorm.DB
}

func NewMySQLRepository(conn *gorm.DB) notes.Repository {
	return &noteRepository{
		conn: conn,
	}
}

func (nr *noteRepository) GetAll(ctx context.Context) ([]notes.Domain, error) {
	var records []Note

	err := nr.conn.WithContext(ctx).Preload("Category").Find(&records).Error

	if err != nil {
		return nil, err
	}

	noteDomain := []notes.Domain{}

	for _, note := range records {
		noteDomain = append(noteDomain, note.ToDomain())
	}

	return noteDomain, nil
}

func (nr *noteRepository) GetByID(ctx context.Context, id string) (notes.Domain, error) {
	var note Note

	err := nr.conn.WithContext(ctx).Preload("Category").First(&note, "id = ?", id).Error

	if err != nil {
		return notes.Domain{}, err
	}

	return note.ToDomain(), nil
}

func (nr *noteRepository) Create(ctx context.Context, noteDomain *notes.Domain) (notes.Domain, error) {
	record := FromDomain(noteDomain)

	result := nr.conn.Create(&record)

	if err := result.Error; err != nil {
		return notes.Domain{}, err
	}

	err := result.Last(&record).Error

	if err != nil {
		return notes.Domain{}, err
	}

	return record.ToDomain(), nil
}

func (nr *noteRepository) Update(ctx context.Context, id string, noteDomain *notes.Domain) (notes.Domain, error) {
	note, err := nr.GetByID(ctx, id)

	if err != nil {
		return notes.Domain{}, err
	}

	updatedNote := FromDomain(&note)

	updatedNote.Title = noteDomain.Title
	updatedNote.Content = noteDomain.Content
	updatedNote.CategoryID = noteDomain.CategoryID

	err = nr.conn.WithContext(ctx).Save(&updatedNote).Error

	if err != nil {
		return notes.Domain{}, err
	}

	return updatedNote.ToDomain(), nil
}

func (nr *noteRepository) Delete(ctx context.Context, id string) error {
	note, err := nr.GetByID(ctx, id)

	if err != nil {
		return err
	}

	deletedNote := FromDomain(&note)

	err = nr.conn.WithContext(ctx).Delete(&deletedNote).Error

	if err != nil {
		return err
	}

	return nil
}

func (nr *noteRepository) Restore(ctx context.Context, id string) (notes.Domain, error) {
	var trashedNote notes.Domain

	trashed := FromDomain(&trashedNote)

	err := nr.conn.WithContext(ctx).Unscoped().First(&trashed, "id = ?", id).Error

	if err != nil {
		return notes.Domain{}, err
	}

	trashed.DeletedAt = gorm.DeletedAt{}

	err = nr.conn.WithContext(ctx).Unscoped().Save(&trashed).Error

	if err != nil {
		return notes.Domain{}, err
	}

	return trashed.ToDomain(), nil
}

func (nr *noteRepository) ForceDelete(ctx context.Context, id string) error {
	note, err := nr.GetByID(ctx, id)

	if err != nil {
		return err
	}

	deletedNote := FromDomain(&note)

	err = nr.conn.WithContext(ctx).Unscoped().Delete(&deletedNote).Error

	if err != nil {
		return err
	}

	return nil
}
