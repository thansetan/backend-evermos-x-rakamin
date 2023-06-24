package authrepository

import (
	"context"
	"final_project/internal/dao"

	"gorm.io/gorm"
)

type AuthRepository interface {
	Login(ctx context.Context, data dao.UserLogin) (res *dao.User, err error)
	Register(ctx context.Context, data dao.User, tx *gorm.DB) (id uint, err error)
}

type AuthRepositoryImpl struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &AuthRepositoryImpl{
		db: db,
	}
}

func (repo *AuthRepositoryImpl) Login(ctx context.Context, data dao.UserLogin) (res *dao.User, err error) {
	if err := repo.db.WithContext(ctx).Where("phone_number = ?", data.PhoneNumber).Preload("Store").First(&res).Error; err != nil {
		return res, gorm.ErrRecordNotFound
	}
	return res, nil
}

func (repo *AuthRepositoryImpl) Register(ctx context.Context, data dao.User, tx *gorm.DB) (id uint, err error) {
	if err := tx.WithContext(ctx).Create(&data).Error; err != nil {
		return id, err
	}
	return data.ID, nil
}
