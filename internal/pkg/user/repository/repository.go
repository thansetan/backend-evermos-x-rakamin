package userrepository

import (
	"context"
	"final_project/internal/dao"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserByID(ctx context.Context, userID string) (res *dao.User, err error)
	UpdateUserByID(ctx context.Context, userID string, data dao.User) error
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{
		db: db,
	}
}

func (repo *UserRepositoryImpl) GetUserByID(ctx context.Context, userID string) (res *dao.User, err error) {
	if err := repo.db.WithContext(ctx).First(&res, userID).Error; err != nil {
		return res, gorm.ErrRecordNotFound
	}
	return res, nil
}

func (repo *UserRepositoryImpl) UpdateUserByID(ctx context.Context, userID string, data dao.User) error {
	var userData dao.User
	if err := repo.db.WithContext(ctx).First(&userData, userID).Error; err != nil {
		return gorm.ErrRecordNotFound
	}
	if err := repo.db.WithContext(ctx).Model(&userData).Updates(data).Where("id = ?", userID).Error; err != nil {
		return err
	}
	return nil
}
