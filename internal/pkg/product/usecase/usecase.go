package productusecase

import (
	"context"
	"final_project/internal/dao"
	"final_project/internal/helper"
	categorydto "final_project/internal/pkg/category/dto"
	productdto "final_project/internal/pkg/product/dto"
	productrepository "final_project/internal/pkg/product/repository"
	storedto "final_project/internal/pkg/store/dto"
	"final_project/internal/utils"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var currentFilePath = "internal/pkg/product/usecase/usecase.go"

type ProductUseCase interface {
	CreateProduct(ctx context.Context, data productdto.ProductCreateOrUpdate, photos []string) (productID uint, err *helper.ErrorStruct)
	GetProducts(ctx context.Context, params productdto.ProductFilter) (res []*productdto.ProductResponse, err *helper.ErrorStruct)
	GetProductByID(ctx context.Context, productID string) (res *productdto.ProductResponse, err *helper.ErrorStruct)
	UpdateProductByID(ctx context.Context, productID string, data productdto.ProductCreateOrUpdate, photos []string) *helper.ErrorStruct
	DeleteProductByID(ctx context.Context, storeID, productID string) *helper.ErrorStruct
}

type ProductUseCaseImpl struct {
	productrepository productrepository.ProductRepository
}

func NewProductUseCase(productrepository productrepository.ProductRepository) ProductUseCase {
	return &ProductUseCaseImpl{
		productrepository: productrepository,
	}
}

func (uc *ProductUseCaseImpl) CreateProduct(ctx context.Context, data productdto.ProductCreateOrUpdate, photos []string) (productID uint, err *helper.ErrorStruct) {
	if validateErr := helper.Validate.Struct(&data); validateErr != nil {
		return productID, &helper.ErrorStruct{
			Err:  validateErr,
			Code: fiber.StatusBadRequest,
		}
	}
	categoryID, parseErr := utils.StringToUint(data.CategoryID)
	if parseErr != nil {
		return productID, &helper.ErrorStruct{
			Err:  parseErr,
			Code: fiber.StatusBadRequest,
		}
	}
	productSlug := utils.GenerateProductSlug(data.ProductName)
	productStock, parseErr := utils.StringToInt(data.Stock)
	if parseErr != nil {
		return productID, &helper.ErrorStruct{
			Err:  parseErr,
			Code: fiber.StatusBadRequest,
		}
	}
	resellerPrice, parseErr := utils.StringToUint(data.ResellerPrice)
	if parseErr != nil {
		return productID, &helper.ErrorStruct{
			Err:  parseErr,
			Code: fiber.StatusBadRequest,
		}
	}
	consumerPrice, parseErr := utils.StringToUint(data.ConsumerPrice)
	if parseErr != nil {
		return productID, &helper.ErrorStruct{
			Err:  parseErr,
			Code: fiber.StatusBadRequest,
		}
	}
	var productData = &dao.Product{
		StoreID:       data.StoreID,
		ProductName:   data.ProductName,
		CategoryID:    categoryID,
		Slug:          productSlug,
		ResellerPrice: resellerPrice,
		ConsumerPrice: consumerPrice,
		Stock:         productStock,
		Description:   data.Description,
	}
	for _, photo := range photos {
		productData.ProductPhotos = append(productData.ProductPhotos, &dao.ProductPhoto{
			Url: photo,
		})
	}
	productID, productErr := uc.productrepository.CreateProduct(ctx, *productData)
	if productErr != nil {
		helper.Logger(currentFilePath, helper.LoggerLevelError, fmt.Sprintf("Error at CreateProduct: %s", productErr.Error()))
		return productID, &helper.ErrorStruct{
			Err:  productErr,
			Code: fiber.StatusBadRequest,
		}
	}
	return productID, err
}

func (uc *ProductUseCaseImpl) GetProducts(ctx context.Context, params productdto.ProductFilter) (res []*productdto.ProductResponse, err *helper.ErrorStruct) {
	if params.Limit < 1 {
		params.Limit = 10
	}

	if params.Page < 1 {
		params.Page = 0
	} else {
		params.Page = (params.Page - 1) * params.Limit
	}
	productRes, productErr := uc.productrepository.GetProducts(ctx, dao.ProductFilter{
		ProductName: params.ProductName,
		Limit:       params.Limit,
		Page:        params.Page,
		CategoryID:  params.CategoryID,
		StoreID:     params.StoreID,
		MaxPrice:    params.MaxPrice,
		MinPrice:    params.MinPrice,
	})
	if productErr != nil {
		helper.Logger(currentFilePath, helper.LoggerLevelError, fmt.Sprintf("Error at GetProducts: %s", productErr.Error()))
		return res, &helper.ErrorStruct{
			Err:  productErr,
			Code: fiber.StatusBadRequest,
		}
	}
	for i, product := range productRes {
		res = append(res, &productdto.ProductResponse{
			ProductID:     product.ID,
			ProductName:   product.ProductName,
			ResellerPrice: product.ResellerPrice,
			ConsumerPrice: product.ConsumerPrice,
			Stock:         product.Stock,
			Description:   product.Description,
			Store: storedto.StoreResponse{
				ID:        product.Store.ID,
				StoreName: product.Store.StoreName,
				PhotoUrl:  product.Store.PhotoUrl,
			},
			Category: categorydto.CategoryResponse{
				ID:           product.Category.ID,
				CategoryName: product.Category.CategoryName,
			},
		})
		for _, photo := range productRes[i].ProductPhotos {
			res[i].Photos = append(res[i].Photos, &productdto.ProductPhotoResponse{
				PhotoID:   photo.ID,
				ProductID: photo.ProductID,
				Url:       photo.Url,
			})
		}
	}
	return res, err
}

func (uc *ProductUseCaseImpl) GetProductByID(ctx context.Context, productID string) (res *productdto.ProductResponse, err *helper.ErrorStruct) {
	productRes, productErr := uc.productrepository.GetProductByID(ctx, productID)
	if productErr != nil {
		helper.Logger(currentFilePath, helper.LoggerLevelError, fmt.Sprintf("Error at GetProductByID: %s", productErr.Error()))
		return res, &helper.ErrorStruct{
			Err:  productErr,
			Code: fiber.StatusNotFound,
		}
	}

	res = &productdto.ProductResponse{
		ProductID:     productRes.ID,
		ProductName:   productRes.ProductName,
		ResellerPrice: productRes.ResellerPrice,
		ConsumerPrice: productRes.ConsumerPrice,
		Stock:         productRes.Stock,
		Description:   productRes.Description,
		Store: storedto.StoreResponse{
			ID:        productRes.Store.ID,
			StoreName: productRes.Store.StoreName,
			PhotoUrl:  productRes.Store.PhotoUrl,
		},
		Category: categorydto.CategoryResponse{
			ID:           productRes.Category.ID,
			CategoryName: productRes.Category.CategoryName,
		},
	}
	for _, photo := range productRes.ProductPhotos {
		res.Photos = append(res.Photos, &productdto.ProductPhotoResponse{
			PhotoID:   photo.ID,
			ProductID: photo.ProductID,
			Url:       photo.Url,
		})
	}
	return res, err
}

func (uc *ProductUseCaseImpl) UpdateProductByID(ctx context.Context, productID string, data productdto.ProductCreateOrUpdate, photos []string) *helper.ErrorStruct {
	if validateErr := helper.Validate.Struct(&data); validateErr != nil {
		return &helper.ErrorStruct{
			Err:  validateErr,
			Code: fiber.StatusBadRequest,
		}
	}
	categoryID, parseErr := utils.StringToUint(data.CategoryID)
	if parseErr != nil {
		return &helper.ErrorStruct{
			Err:  parseErr,
			Code: fiber.StatusBadRequest,
		}
	}
	productStock, parseErr := utils.StringToInt(data.Stock)
	if parseErr != nil {
		return &helper.ErrorStruct{
			Err:  parseErr,
			Code: fiber.StatusBadRequest,
		}
	}
	resellerPrice, parseErr := utils.StringToUint(data.ResellerPrice)
	if parseErr != nil {
		return &helper.ErrorStruct{
			Err:  parseErr,
			Code: fiber.StatusBadRequest,
		}
	}
	consumerPrice, parseErr := utils.StringToUint(data.ConsumerPrice)
	if parseErr != nil {
		return &helper.ErrorStruct{
			Err:  parseErr,
			Code: fiber.StatusBadRequest,
		}
	}
	var productData = &dao.Product{
		StoreID:       data.StoreID,
		ProductName:   data.ProductName,
		CategoryID:    categoryID,
		ResellerPrice: resellerPrice,
		ConsumerPrice: consumerPrice,
		Stock:         productStock,
		Description:   data.Description,
	}
	for _, photo := range photos {
		productData.ProductPhotos = append(productData.ProductPhotos, &dao.ProductPhoto{
			Url: photo,
		})
	}
	productErr := uc.productrepository.UpdateProductByID(ctx, productID, *productData)
	if productErr != nil {
		helper.Logger(currentFilePath, helper.LoggerLevelError, fmt.Sprintf("Error at UpdateProductByID: %s", productErr.Error()))
		if productErr == gorm.ErrRecordNotFound {
			return &helper.ErrorStruct{
				Err:  productErr,
				Code: fiber.StatusNotFound,
			}
		}
		return &helper.ErrorStruct{
			Err:  productErr,
			Code: fiber.StatusBadRequest,
		}
	}
	return nil
}

func (uc *ProductUseCaseImpl) DeleteProductByID(ctx context.Context, storeID, productID string) *helper.ErrorStruct {
	productErr := uc.productrepository.DeleteProductByID(ctx, storeID, productID)
	if productErr != nil {
		helper.Logger(currentFilePath, helper.LoggerLevelError, fmt.Sprintf("Error at DeleteProductByID: %s", productErr.Error()))
		if productErr == gorm.ErrRecordNotFound {
			return &helper.ErrorStruct{
				Err:  productErr,
				Code: fiber.StatusNotFound,
			}
		}
		return &helper.ErrorStruct{
			Err:  productErr,
			Code: fiber.StatusBadRequest,
		}
	}
	return nil
}
