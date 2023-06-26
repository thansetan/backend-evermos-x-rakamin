package transactionusecase

import (
	"context"
	"final_project/internal/dao"
	"final_project/internal/helper"
	productdto "final_project/internal/pkg/product/dto"
	productrepository "final_project/internal/pkg/product/repository"
	transactiondto "final_project/internal/pkg/transaction/dto"
	transactionrepository "final_project/internal/pkg/transaction/repository"
	"final_project/internal/utils"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var currentFilePath = "internal/pkg/store/usecase/usecase.go"

type TransactionUseCase interface {
	GetUserTransactions(ctx context.Context, userID string) (res []*transactiondto.TransactionResponse, err *helper.ErrorStruct)
	GetUserTransactionByID(ctx context.Context, userID, transactionID string) (res *transactiondto.TransactionResponse, err *helper.ErrorStruct)
	CreateTransaction(ctx context.Context, data transactiondto.TransactionCreate) (res uint, err *helper.ErrorStruct)
}

type TransactionUseCaseImpl struct {
	transactionrepository transactionrepository.TransactionRepository
	productrepository     productrepository.ProductRepository
	db                    *gorm.DB
}

func NewTransactionUseCase(transactionrepository transactionrepository.TransactionRepository, productrepository productrepository.ProductRepository, db *gorm.DB) TransactionUseCase {
	return &TransactionUseCaseImpl{
		transactionrepository: transactionrepository,
		productrepository:     productrepository,
		db:                    db,
	}
}

func (uc *TransactionUseCaseImpl) GetUserTransactions(ctx context.Context, userID string) (res []*transactiondto.TransactionResponse, err *helper.ErrorStruct) {
	trxRes, trxErr := uc.transactionrepository.GetTransactionsByUserID(ctx, userID)
	if trxErr != nil {
		helper.Logger(currentFilePath, helper.LoggerLevelError, fmt.Sprintf("Error at GetStoreByID: %s", trxErr.Error()))
		return res, &helper.ErrorStruct{
			Err:  trxErr,
			Code: fiber.StatusBadRequest,
		}
	}
	for _, trx := range trxRes {
		trxRes := &transactiondto.TransactionResponse{
			ID:            trx.ID,
			TotalPrice:    trx.TotalPrice,
			InvoiceNumber: trx.InvoiceNumber,
			PaymentMethod: trx.PaymentMethod,
		}
		// Address
		trxRes.Address.ID = trx.Address.ID
		trxRes.Address.AddressDetail = trx.Address.AddressDetail
		trxRes.Address.PhoneNumber = trx.Address.PhoneNumber
		trxRes.Address.AddressTitle = trx.Address.AddressTitle
		trxRes.Address.Recipient = trx.Address.Recipient

		// Transaction details
		trxDetails := []transactiondto.TransactionDetailResponse{}
		for _, trxDetail := range trx.TransactionDetails {
			trxDetailRes := &transactiondto.TransactionDetailResponse{
				Quantity:   trxDetail.Quantity,
				TotalPrice: trxDetail.TotalPrice,
			}
			trxDetailRes.Store.ID = trxDetail.Store.ID
			trxDetailRes.Store.PhotoUrl = trxDetail.Store.PhotoUrl
			trxDetailRes.Store.StoreName = trxDetail.Store.StoreName

			// Product
			trxDetailRes.Product.ProductID = trxDetail.ProductLog.ID
			trxDetailRes.Product.ProductName = trxDetail.ProductLog.ProductName
			trxDetailRes.Product.ResellerPrice = trxDetail.ProductLog.ResellerPrice
			trxDetailRes.Product.ConsumerPrice = trxDetail.ProductLog.ConsumerPrice
			trxDetailRes.Product.Description = trxDetail.ProductLog.Description
			trxDetailRes.Product.Store = trxDetailRes.Store
			trxDetailRes.Product.Category.ID = trxDetail.ProductLog.Category.ID
			trxDetailRes.Product.Category.CategoryName = trxDetail.ProductLog.Category.CategoryName

			// Photos
			for _, photo := range trxDetail.ProductLog.Product.ProductPhotos {
				photoRes := &productdto.ProductPhotoResponse{
					PhotoID:   photo.ID,
					ProductID: photo.ProductID,
					Url:       photo.Url,
				}
				trxDetailRes.Product.Photos = append(trxDetailRes.Product.Photos, photoRes)
			}
			trxDetails = append(trxDetails, *trxDetailRes)
		}
		trxRes.TransactionDetails = trxDetails
		res = append(res, trxRes)
	}
	return res, nil
}

func (uc *TransactionUseCaseImpl) GetUserTransactionByID(ctx context.Context, userID, transactionID string) (res *transactiondto.TransactionResponse, err *helper.ErrorStruct) {
	trxRes, trxErr := uc.transactionrepository.GetTransactionByUserIDAndTransactionID(ctx, userID, transactionID)
	if trxErr != nil {
		helper.Logger(currentFilePath, helper.LoggerLevelError, fmt.Sprintf("Error at GetStoreByID: %s", trxErr.Error()))
		return res, &helper.ErrorStruct{
			Err:  trxErr,
			Code: fiber.StatusNotFound,
		}
	}
	res = &transactiondto.TransactionResponse{
		ID:            trxRes.ID,
		TotalPrice:    trxRes.TotalPrice,
		InvoiceNumber: trxRes.InvoiceNumber,
		PaymentMethod: trxRes.PaymentMethod,
	}

	// Address
	res.Address.ID = trxRes.Address.ID
	res.Address.AddressTitle = trxRes.Address.AddressTitle
	res.Address.AddressDetail = trxRes.Address.AddressDetail
	res.Address.PhoneNumber = trxRes.Address.PhoneNumber
	res.Address.Recipient = trxRes.Address.Recipient

	// Transaction details
	trxDetails := []transactiondto.TransactionDetailResponse{}
	for _, trxDetail := range trxRes.TransactionDetails {
		trxDetailRes := &transactiondto.TransactionDetailResponse{
			Quantity:   trxDetail.Quantity,
			TotalPrice: trxDetail.TotalPrice,
		}
		trxDetailRes.Store.ID = trxDetail.Store.ID
		trxDetailRes.Store.PhotoUrl = trxDetail.Store.PhotoUrl
		trxDetailRes.Store.StoreName = trxDetail.Store.StoreName

		// Product
		trxDetailRes.Product.ProductID = trxDetail.ProductLog.ProductID
		trxDetailRes.Product.ProductName = trxDetail.ProductLog.ProductName
		trxDetailRes.Product.ResellerPrice = trxDetail.ProductLog.ResellerPrice
		trxDetailRes.Product.ConsumerPrice = trxDetail.ProductLog.ConsumerPrice
		trxDetailRes.Product.Description = trxDetail.ProductLog.Description
		trxDetailRes.Product.Store = trxDetailRes.Store
		trxDetailRes.Product.Category.ID = trxDetail.ProductLog.Category.ID
		trxDetailRes.Product.Category.CategoryName = trxDetail.ProductLog.Category.CategoryName

		// Photos
		for _, photo := range trxDetail.ProductLog.Product.ProductPhotos {
			photoRes := &productdto.ProductPhotoResponse{
				PhotoID:   photo.ID,
				ProductID: photo.ProductID,
				Url:       photo.Url,
			}
			trxDetailRes.Product.Photos = append(trxDetailRes.Product.Photos, photoRes)
		}
		trxDetails = append(trxDetails, *trxDetailRes)
	}
	res.TransactionDetails = trxDetails
	return res, nil
}

func (uc *TransactionUseCaseImpl) CreateTransaction(ctx context.Context, data transactiondto.TransactionCreate) (res uint, err *helper.ErrorStruct) {
	if validateErr := helper.Validate.Struct(data); validateErr != nil {
		log.Println(validateErr)
		return res, &helper.ErrorStruct{
			Err:  validateErr,
			Code: fiber.StatusBadRequest,
		}
	}
	tx := uc.db.Begin()
	productIDSlice := make([]uint, len(data.TransactionDetails))
	productQtySlice := make([]uint, len(data.TransactionDetails))
	for i, product := range data.TransactionDetails {
		productIDSlice[i] = product.ProductID
		productQtySlice[i] = product.Quantity
	}

	// First we query product data with those IDs
	productRecords, productErr := uc.productrepository.GetProductDataUsingSliceOfID(ctx, productIDSlice)
	if productErr != nil {
		helper.Logger(currentFilePath, helper.LoggerLevelError, fmt.Sprintf("Error at CreateTransaction: %s", productErr.Error()))
		return res, &helper.ErrorStruct{
			Err:  productErr,
			Code: fiber.StatusBadRequest,
		}
	}

	// and then we map those data to productLog object
	var productLogData []dao.ProductLog
	for _, product := range productRecords {
		productLogData = append(productLogData, dao.ProductLog{
			StoreID:       product.StoreID,
			CategoryID:    product.CategoryID,
			ProductID:     product.ID,
			ProductName:   product.ProductName,
			Slug:          product.Slug,
			ResellerPrice: product.ResellerPrice,
			ConsumerPrice: product.ConsumerPrice,
			Description:   product.Description,
		})
	}

	// and then we insert the productLog to the database
	productLogRes, productLogErr := uc.productrepository.CreateProductLog(ctx, productLogData)
	if productLogErr != nil {
		helper.Logger(currentFilePath, helper.LoggerLevelError, fmt.Sprintf("Error at CreateTransaction: %s", productErr.Error()))
		return res, &helper.ErrorStruct{
			Err:  productErr,
			Code: fiber.StatusBadRequest,
		}
	}

	// then, we create the transaction
	var trxDetailsData []*dao.TransactionDetail
	for i, productLog := range productLogRes {
		trxDetail := &dao.TransactionDetail{
			ProductLogID: productLog,
			StoreID:      productLogData[i].StoreID,
			Quantity:     productQtySlice[i],
			TotalPrice:   productQtySlice[i] * productLogData[i].ConsumerPrice,
		}
		trxDetailsData = append(trxDetailsData, trxDetail)
	}

	var totalTrxPrice uint
	for _, trxDetails := range trxDetailsData {
		totalTrxPrice += trxDetails.TotalPrice
	}
	userIDUint, parseErr := utils.StringToUint(data.UserID)
	if parseErr != nil {
		return res, &helper.ErrorStruct{
			Err:  parseErr,
			Code: fiber.StatusBadRequest,
		}
	}
	transactionRes, transactionErr := uc.transactionrepository.CreateTransaction(ctx, dao.Transaction{
		UserID:             userIDUint,
		AddressID:          data.AddressID,
		PaymentMethod:      data.PaymentMethod,
		InvoiceNumber:      utils.GenerateInvoiceNumber(),
		TotalPrice:         totalTrxPrice,
		TransactionDetails: trxDetailsData,
	})
	if transactionErr != nil {
		helper.Logger(currentFilePath, helper.LoggerLevelError, fmt.Sprintf("Error at CreateTransaction: %s", transactionErr.Error()))
		return res, &helper.ErrorStruct{
			Err:  transactionErr,
			Code: fiber.StatusBadRequest,
		}
	}

	// finally we update the product data after the transaction

	for i, product := range data.TransactionDetails {
		productErr := uc.productrepository.UpdateProductByID(ctx, fmt.Sprintf("%d", product.ProductID), dao.Product{
			Stock:   productRecords[i].Stock - int(product.Quantity),
			StoreID: productRecords[i].StoreID,
		})
		if productErr != nil {
			helper.Logger(currentFilePath, helper.LoggerLevelError, fmt.Sprintf("Error at CreateTransaction: %s", transactionErr.Error()))
			return res, &helper.ErrorStruct{
				Err:  transactionErr,
				Code: fiber.StatusBadRequest,
			}
		}
	}
	tx.Commit()
	return transactionRes, nil
}
