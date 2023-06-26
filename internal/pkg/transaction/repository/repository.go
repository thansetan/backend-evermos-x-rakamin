package transactionrepository

import (
	"context"
	"final_project/internal/dao"
	"fmt"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	GetTransactionsByUserID(ctx context.Context, userID string) (res []*dao.Transaction, err error)
	GetTransactionByUserIDAndTransactionID(ctx context.Context, userID, transactionID string) (res *dao.Transaction, err error)
	CreateTransaction(ctx context.Context, data dao.Transaction) (transactionID uint, err error)
}

type TransactionRepositoryImpl struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &TransactionRepositoryImpl{
		db: db,
	}
}

func (repo *TransactionRepositoryImpl) GetTransactionsByUserID(ctx context.Context, userID string) (res []*dao.Transaction, err error) {
	if err := repo.db.WithContext(ctx).Preload("Address").Preload("TransactionDetails").Preload("TransactionDetails.Store").
		Preload("TransactionDetails.ProductLog").Preload("TransactionDetails.ProductLog.Category").
		Preload("TransactionDetails.ProductLog.Product.ProductPhotos").Where("user_id = ?", userID).
		Find(&res).Error; err != nil {
		return res, err
	}
	return res, nil
}

func (repo *TransactionRepositoryImpl) GetTransactionByUserIDAndTransactionID(ctx context.Context, userID, transactionID string) (res *dao.Transaction, err error) {
	if err := repo.db.WithContext(ctx).Where("user_id = ?", userID).Preload("Address").Preload("TransactionDetails").
		Preload("TransactionDetails.Store").Preload("TransactionDetails.ProductLog").
		Preload("TransactionDetails.ProductLog.Category").Preload("TransactionDetails.ProductLog.Product.ProductPhotos").
		Find(&res, transactionID).Error; err != nil {
		return res, gorm.ErrRecordNotFound
	}
	fmt.Println(res.Address)
	return res, nil
}
func (repo *TransactionRepositoryImpl) CreateTransaction(ctx context.Context, data dao.Transaction) (transactionID uint, err error) {
	if err := repo.db.WithContext(ctx).Create(&data).Error; err != nil {
		return transactionID, err
	}
	return data.ID, nil
}
