package addressrepository

import (
	"context"
	"final_project/internal/dao"
	"fmt"

	"gorm.io/gorm"
)

type AddressRepository interface {
	GetAddresses(ctx context.Context, filter dao.AddressFilter) (res []*dao.Address, err error)
	GetAddressByID(ctx context.Context, userID, addressID string) (res *dao.Address, err error)
	CreateAddress(ctx context.Context, data dao.Address) (AddressID uint, err error)
	UpdateAddressByID(ctx context.Context, addressID string, data dao.Address) error
	DeleteAddressByID(ctx context.Context, userID, addressID string) error
}

type AddressRepositoryImpl struct {
	db *gorm.DB
}

func NewAddressRepository(db *gorm.DB) AddressRepository {
	return &AddressRepositoryImpl{
		db: db,
	}
}

func (repo *AddressRepositoryImpl) GetAddresses(ctx context.Context, params dao.AddressFilter) (res []*dao.Address, err error) {
	if err := repo.db.WithContext(ctx).Where("user_id = ? AND address_title LIKE ?", params.UserID, fmt.Sprintf("%%%s%%", params.AddressTitle)).Find(&res).Error; err != nil {
		return res, err
	}
	return res, nil
}

func (repo *AddressRepositoryImpl) GetAddressByID(ctx context.Context, userID, addressID string) (res *dao.Address, err error) {
	fmt.Println(addressID)
	if err := repo.db.WithContext(ctx).Where("user_id = ?", userID).First(&res, addressID).Error; err != nil {
		return res, gorm.ErrRecordNotFound
	}
	return res, nil
}

func (repo *AddressRepositoryImpl) CreateAddress(ctx context.Context, data dao.Address) (AddressID uint, err error) {
	if err := repo.db.WithContext(ctx).Create(&data).Error; err != nil {
		return AddressID, err
	}
	return data.ID, err
}

func (repo *AddressRepositoryImpl) UpdateAddressByID(ctx context.Context, addressID string, data dao.Address) error {
	var addressData dao.Address
	fmt.Println(data)
	if err := repo.db.WithContext(ctx).First(&addressData, addressID).Error; err != nil {
		return gorm.ErrRecordNotFound
	}
	if err := repo.db.WithContext(ctx).Model(addressData).Updates(data).Error; err != nil {
		return err
	}
	return nil
}

func (repo *AddressRepositoryImpl) DeleteAddressByID(ctx context.Context, userID, addressID string) error {
	var addressData dao.Address
	if err := repo.db.WithContext(ctx).Where("user_id = ?", userID).First(&addressData, addressID).Error; err != nil {
		return gorm.ErrRecordNotFound
	}
	if err := repo.db.WithContext(ctx).Delete(&addressData, addressID).Error; err != nil {
		return err
	}
	return nil
}
