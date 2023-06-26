package productrepository

import (
	"context"
	"final_project/internal/dao"
	"final_project/internal/utils"
	"fmt"

	"gorm.io/gorm"
)

type ProductRepository interface {
	CreateProduct(ctx context.Context, data dao.Product) (productID uint, err error)
	GetProducts(ctx context.Context, params dao.ProductFilter) (res []*dao.Product, err error)
	GetProductByID(ctx context.Context, productID string) (res *dao.Product, err error)
	UpdateProductByID(ctx context.Context, productID string, data dao.Product) error
	DeleteProductByID(ctx context.Context, storeID, productID string) error

	GetProductDataUsingSliceOfID(ctx context.Context, productIDSlice []uint) (res []*dao.Product, err error)
	CreateProductLog(ctx context.Context, data []dao.ProductLog, tx *gorm.DB) (productLogID []*dao.ProductLogRes, err error)
}

type ProductRepositoryImpl struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &ProductRepositoryImpl{
		db: db,
	}
}

func (repo *ProductRepositoryImpl) CreateProduct(ctx context.Context, data dao.Product) (productID uint, err error) {
	if err := repo.db.WithContext(ctx).Create(&data).Error; err != nil {
		return productID, err
	}
	return data.ID, nil
}

func (repo *ProductRepositoryImpl) GetProducts(ctx context.Context, params dao.ProductFilter) (res []*dao.Product, err error) {
	db := repo.db
	if params.ProductName != "" {
		db = db.WithContext(ctx).Where("product_name LIKE ?", fmt.Sprintf("%%%s%%", params.ProductName))
	}
	if params.MaxPrice > 0 {
		db = db.WithContext(ctx).Where("consumer_price <= ?", params.MaxPrice)
	}
	if params.MinPrice > 0 {
		db = db.WithContext(ctx).Where("consumer_price >= ?", params.MinPrice)
	}
	if params.CategoryID > 0 {
		db = db.WithContext(ctx).Where("category_id = ?", params.CategoryID)
	}
	if params.StoreID > 0 {
		db = db.WithContext(ctx).Where("store_id = ?", params.StoreID)
	}
	if err := db.WithContext(ctx).Preload("Store").Preload("Category").Preload("ProductPhotos").Limit(params.Limit).Offset(params.Page).Find(&res).Error; err != nil {
		return res, err
	}
	return res, nil
}

func (repo *ProductRepositoryImpl) GetProductByID(ctx context.Context, productID string) (res *dao.Product, err error) {
	if err := repo.db.WithContext(ctx).Preload("Store").Preload("Category").Preload("ProductPhotos").First(&res, productID).Error; err != nil {
		return res, gorm.ErrRecordNotFound
	}
	return res, nil
}

func (repo *ProductRepositoryImpl) UpdateProductByID(ctx context.Context, productID string, data dao.Product) error {
	var productData dao.Product
	if err := repo.db.WithContext(ctx).Where("store_id = ?", data.StoreID).Preload("ProductPhotos").First(&productData, productID).Error; err != nil {
		return gorm.ErrRecordNotFound
	}
	if data.ProductPhotos != nil {
		currentPhotos := productData.ProductPhotos
		err := repo.db.WithContext(ctx).Model(&productData).Association("ProductPhotos").Replace(data.ProductPhotos)
		if err == nil {
			for _, photo := range currentPhotos {
				utils.RemoveUnusedPhoto(photo.Url)
			}
		}
		if err != nil {
			return err
		}
	}
	if err := repo.db.WithContext(ctx).Model(&productData).Updates(&data).Error; err != nil {
		return err
	}
	return nil
}

func (repo *ProductRepositoryImpl) DeleteProductByID(ctx context.Context, storeID, productID string) error {
	var productData dao.Product
	if err := repo.db.WithContext(ctx).Where("store_id = ?", storeID).Preload("ProductPhotos").First(&productData, productID).Error; err != nil {
		return gorm.ErrRecordNotFound
	}
	if productData.ProductPhotos != nil {
		currentPhotos := productData.ProductPhotos
		err := repo.db.WithContext(ctx).Model(&productData).Association("ProductPhotos").Delete(productData.ProductPhotos)
		if err == nil {
			for _, photo := range currentPhotos {
				utils.RemoveUnusedPhoto(photo.Url)
			}
		}
		if err != nil {
			return err
		}
	}
	if err := repo.db.WithContext(ctx).Delete(&productData).Error; err != nil {
		return err
	}
	return nil
}

func (repo *ProductRepositoryImpl) GetProductDataUsingSliceOfID(ctx context.Context, productIDSlice []uint) (res []*dao.Product, err error) {
	if err := repo.db.WithContext(ctx).Find(&res, productIDSlice).Error; err != nil {
		return res, err
	}
	if len(res) != len(productIDSlice) { // if the lengths are not the same, it means that some of the products are invalid.
		return res, gorm.ErrRecordNotFound
	}
	return res, nil
}

func (repo *ProductRepositoryImpl) CreateProductLog(ctx context.Context, data []dao.ProductLog, tx *gorm.DB) (res []*dao.ProductLogRes, err error) {
	if err := tx.WithContext(ctx).Create(&data).Error; err != nil {
		return res, err
	}
	for _, productLog := range data {
		res = append(res, &dao.ProductLogRes{
			ID:        productLog.ID,
			ProductID: productLog.ProductID,
		})
	}
	return res, nil
}
