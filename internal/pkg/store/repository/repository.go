package storerepository

import (
	"context"
	"final_project/internal/dao"
	"final_project/internal/utils"

	"gorm.io/gorm"
)

type StoreRepository interface {
	CreateStore(ctx context.Context, data dao.Store, tx *gorm.DB) (storeID uint, err error)
	GetStoreByID(ctx context.Context, storeID string) (res *dao.Store, err error)
	GetAllStores(ctx context.Context, params dao.StoreFilter) (res []*dao.Store, err error)
	GetStoreByUserID(ctx context.Context, userID string) (res *dao.Store, err error)
	UpdateStoreByUserID(ctx context.Context, userID string, data dao.Store) error
}

type StoreRepositoryImpl struct {
	db *gorm.DB
}

func NewStoreRepository(db *gorm.DB) StoreRepository {
	return &StoreRepositoryImpl{
		db: db,
	}
}

func (repo *StoreRepositoryImpl) CreateStore(ctx context.Context, data dao.Store, tx *gorm.DB) (storeID uint, err error) {
	if err := tx.WithContext(ctx).Create(&data).Error; err != nil {
		return storeID, err
	}
	return data.ID, nil
}

func (repo *StoreRepositoryImpl) GetStoreByID(ctx context.Context, storeID string) (res *dao.Store, err error) {
	if err := repo.db.WithContext(ctx).First(&res, storeID).Error; err != nil {
		return res, gorm.ErrRecordNotFound
	}
	return res, nil
}

func (repo *StoreRepositoryImpl) GetAllStores(ctx context.Context, params dao.StoreFilter) (res []*dao.Store, err error) {
	if err := repo.db.WithContext(ctx).Limit(params.Limit).Offset(params.Offset).Find(&res).Error; err != nil {
		return res, err
	}
	return res, nil
}

func (repo *StoreRepositoryImpl) GetStoreByUserID(ctx context.Context, userID string) (res *dao.Store, err error) {
	if err := repo.db.WithContext(ctx).Where("user_id = ?", userID).First(&res).Error; err != nil {
		return res, gorm.ErrRecordNotFound
	}
	return res, nil
}

func (repo *StoreRepositoryImpl) UpdateStoreByUserID(ctx context.Context, userID string, data dao.Store) error {
	var storeData dao.Store
	if err := repo.db.WithContext(ctx).Where("user_id = ?", userID).First(&storeData).Error; err != nil {
		return gorm.ErrRecordNotFound
	}
	if err := repo.db.WithContext(ctx).Model(storeData).Updates(data).Where("user_id = ?", userID).Error; err != nil {
		return err
	}
	if storeData.PhotoUrl != data.PhotoUrl && storeData.PhotoUrl != "" && data.PhotoUrl != "" {
		utils.RemoveUnusedPhoto(storeData.PhotoUrl)
	}
	return nil
}
