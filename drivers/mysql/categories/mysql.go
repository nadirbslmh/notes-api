package categories

import (
	"context"
	"notes-api/businesses/categories"

	"gorm.io/gorm"
)

type categoryRepository struct {
	conn *gorm.DB
}

func NewMySQLRepository(conn *gorm.DB) categories.Repository {
	return &categoryRepository{
		conn: conn,
	}
}

func (cr *categoryRepository) GetAll(ctx context.Context) ([]categories.Domain, error) {
	var records []Category

	err := cr.conn.WithContext(ctx).Find(&records).Error

	if err != nil {
		return nil, err
	}

	categoryDomain := []categories.Domain{}

	for _, category := range records {
		categoryDomain = append(categoryDomain, category.ToDomain())
	}

	return categoryDomain, nil
}

func (cr *categoryRepository) GetByID(ctx context.Context, id string) (categories.Domain, error) {
	var category Category

	err := cr.conn.WithContext(ctx).First(&category, "id = ?", id).Error

	if err != nil {
		return categories.Domain{}, err
	}

	return category.ToDomain(), nil
}

func (cr *categoryRepository) Create(ctx context.Context, categoryDomain *categories.Domain) (categories.Domain, error) {
	record := FromDomain(categoryDomain)

	result := cr.conn.WithContext(ctx).Create(&record)

	if err := result.Error; err != nil {
		return categories.Domain{}, err
	}

	err := result.Last(&record).Error

	if err != nil {
		return categories.Domain{}, err
	}

	return record.ToDomain(), nil
}

func (cr *categoryRepository) Update(ctx context.Context, id string, categoryDomain *categories.Domain) (categories.Domain, error) {
	category, err := cr.GetByID(ctx, id)

	if err != nil {
		return categories.Domain{}, err
	}

	updatedCategory := FromDomain(&category)

	updatedCategory.Name = categoryDomain.Name

	err = cr.conn.WithContext(ctx).Save(&updatedCategory).Error

	if err != nil {
		return categories.Domain{}, err
	}

	return updatedCategory.ToDomain(), nil
}

func (cr *categoryRepository) Delete(ctx context.Context, id string) error {
	category, err := cr.GetByID(ctx, id)

	if err != nil {
		return err
	}

	deletedCategory := FromDomain(&category)

	err = cr.conn.WithContext(ctx).Unscoped().Delete(&deletedCategory).Error

	if err != nil {
		return err
	}

	return nil
}
