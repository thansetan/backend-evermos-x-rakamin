package storeusecase

import (
	"context"
	"final_project/internal/dao"
	"final_project/internal/helper"
	storedto "final_project/internal/pkg/store/dto"
	storerepository "final_project/internal/pkg/store/repository"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var currentFilePath = "internal/pkg/store/usecase/usecase.go"

type StoreUseCase interface {
	GetStoreByID(ctx context.Context, storeID string) (res *storedto.StoreResponse, err *helper.ErrorStruct)
	GetAllStores(ctx context.Context, params storedto.StoreFilter) (res []*storedto.StoreResponse, err *helper.ErrorStruct)
	GetMyStore(ctx context.Context, userID string) (res *storedto.StoreResponse, err *helper.ErrorStruct)
	UpdateStoreByUserID(ctx context.Context, userID string, data storedto.StoreUpdate) *helper.ErrorStruct
}

type StoreUseCaseImpl struct {
	storerepository storerepository.StoreRepository
}

func NewStoreUseCase(storerepository storerepository.StoreRepository) StoreUseCase {
	return &StoreUseCaseImpl{
		storerepository: storerepository,
	}
}

func (uc *StoreUseCaseImpl) GetStoreByID(ctx context.Context, storeID string) (res *storedto.StoreResponse, err *helper.ErrorStruct) {
	storeRes, storeErr := uc.storerepository.GetStoreByID(ctx, storeID)
	if storeErr != nil {
		helper.Logger(currentFilePath, helper.LoggerLevelError, fmt.Sprintf("Error at GetStoreByID: %s", storeErr.Error()))
		return res, &helper.ErrorStruct{
			Err:  storeErr,
			Code: fiber.StatusNotFound,
		}
	}
	res = &storedto.StoreResponse{
		ID:        storeRes.ID,
		StoreName: storeRes.StoreName,
		PhotoUrl:  storeRes.PhotoUrl,
	}
	return res, nil
}

func (uc *StoreUseCaseImpl) GetAllStores(ctx context.Context, params storedto.StoreFilter) (res []*storedto.StoreResponse, err *helper.ErrorStruct) {
	if params.Limit < 1 {
		params.Limit = 10
	}

	if params.Page < 1 {
		params.Page = 0
	} else {
		params.Page = (params.Page - 1) * params.Limit
	}
	storeRes, storeErr := uc.storerepository.GetAllStores(ctx, dao.StoreFilter{
		Limit:  params.Limit,
		Offset: params.Page,
		Name:   params.Name,
	})
	if storeErr != nil {
		helper.Logger(currentFilePath, helper.LoggerLevelError, fmt.Sprintf("Error at GetAllStores: %s", storeErr.Error()))
		return res, &helper.ErrorStruct{
			Err:  storeErr,
			Code: fiber.StatusBadRequest,
		}
	}
	for _, store := range storeRes {
		res = append(res, &storedto.StoreResponse{
			ID:        store.ID,
			StoreName: store.StoreName,
			PhotoUrl:  store.PhotoUrl,
		})
	}
	return res, nil
}

func (uc *StoreUseCaseImpl) GetMyStore(ctx context.Context, userID string) (res *storedto.StoreResponse, err *helper.ErrorStruct) {
	storeRes, storeErr := uc.storerepository.GetStoreByUserID(ctx, userID)
	if storeErr != nil {
		helper.Logger(currentFilePath, helper.LoggerLevelError, fmt.Sprintf("Error at GetMyStore: %s", storeErr.Error()))
		return res, &helper.ErrorStruct{
			Err:  storeErr,
			Code: fiber.StatusNotFound,
		}
	}
	res = &storedto.StoreResponse{
		ID:        storeRes.ID,
		StoreName: storeRes.StoreName,
		PhotoUrl:  storeRes.PhotoUrl,
	}
	return res, nil
}

func (uc *StoreUseCaseImpl) UpdateStoreByUserID(ctx context.Context, userID string, data storedto.StoreUpdate) *helper.ErrorStruct {
	if errValidate := helper.Validate.Struct(data); errValidate != nil {
		log.Println(errValidate)
		return &helper.ErrorStruct{
			Err:  errValidate,
			Code: fiber.StatusBadRequest,
		}
	}
	storeErr := uc.storerepository.UpdateStoreByUserID(ctx, userID, dao.Store{
		StoreName: data.StoreName,
		PhotoUrl:  data.PhotoUrl,
	})
	if storeErr != nil {
		helper.Logger(currentFilePath, helper.LoggerLevelError, fmt.Sprintf("Error at UpdateStoreByUserID: %s", storeErr.Error()))
		if storeErr == gorm.ErrRecordNotFound {
			return &helper.ErrorStruct{
				Code: fiber.StatusNotFound,
				Err:  storeErr,
			}
		}
		return &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  storeErr,
		}
	}
	return nil
}
