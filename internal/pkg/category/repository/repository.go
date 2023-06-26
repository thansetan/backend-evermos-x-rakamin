package categoryrepository

import (
	"context"
	"final_project/internal/dao"
	"fmt"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	GetCategories(ctx context.Context, params dao.CategoryFilter) (res []*dao.Category, err error)
	GetCategoryByID(ctx context.Context, categoryID string) (res *dao.Category, err error)
	CreateCategory(ctx context.Context, data dao.Category) (categoryID uint, err error)
	UpdateCategoryByID(ctx context.Context, categoryID string, data dao.Category) error
	DeleteCategoryByID(ctx context.Context, categoryID string) error
}

type CategoryRepositoryImpl struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &CategoryRepositoryImpl{
		db: db,
	}
}

func (repo *CategoryRepositoryImpl) GetCategories(ctx context.Context, params dao.CategoryFilter) (res []*dao.Category, err error) {
	if err := repo.db.WithContext(ctx).Where("category_name LIKE ?", fmt.Sprintf("%%%s%%", params.Name)).Limit(params.Limit).Offset(params.Offset).Find(&res).Error; err != nil {
		return res, err
	}
	return res, nil
}

func (repo *CategoryRepositoryImpl) GetCategoryByID(ctx context.Context, categoryID string) (res *dao.Category, err error) {
	if err := repo.db.WithContext(ctx).First(&res, categoryID).Error; err != nil {
		return res, gorm.ErrRecordNotFound
	}
	return res, nil
}

func (repo *CategoryRepositoryImpl) CreateCategory(ctx context.Context, data dao.Category) (categoryID uint, err error) {
	if err := repo.db.WithContext(ctx).Create(&data).Error; err != nil {
		return categoryID, err
	}
	return data.ID, nil
}

func (repo *CategoryRepositoryImpl) UpdateCategoryByID(ctx context.Context, categoryID string, data dao.Category) error {
	var categoryData dao.Category
	if err := repo.db.WithContext(ctx).First(&categoryData, categoryID).Error; err != nil {
		return gorm.ErrRecordNotFound
	}
	if err := repo.db.WithContext(ctx).Model(&categoryData).Updates(&data).Error; err != nil {
		return err
	}
	return nil
}

func (repo *CategoryRepositoryImpl) DeleteCategoryByID(ctx context.Context, categoryID string) error {
	var categoryData dao.Category
	if err := repo.db.WithContext(ctx).First(&categoryData, categoryID).Error; err != nil {
		return gorm.ErrRecordNotFound
	}
	if err := repo.db.WithContext(ctx).Delete(&categoryData).Error; err != nil {
		return err
	}
	return nil
}
