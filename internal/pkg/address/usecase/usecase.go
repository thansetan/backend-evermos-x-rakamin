package addressusecase

import (
	"context"
	"final_project/internal/dao"
	"final_project/internal/helper"
	addressdto "final_project/internal/pkg/address/dto"
	addressrepository "final_project/internal/pkg/address/repository"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type AddressUseCase interface {
	GetUserAddresses(ctx context.Context, filter addressdto.AddressFilter) (res []*addressdto.AddressResponse, err *helper.ErrorStruct)
	GetUserAddressByID(ctx context.Context, userID, addressID string) (res *addressdto.AddressResponse, err *helper.ErrorStruct)
	CreateAddress(ctx context.Context, data addressdto.AddressCreate) (res uint, err *helper.ErrorStruct)
	UpdateAddressByID(ctx context.Context, userID, addressID string, data addressdto.AddressUpdate) *helper.ErrorStruct
	DeleteAddressByID(ctx context.Context, userID, addressID string) *helper.ErrorStruct
}

type AddressUseCaseImpl struct {
	addressrepository addressrepository.AddressRepository
}

func NewAddressUseCase(addressrepository addressrepository.AddressRepository) AddressUseCase {
	return &AddressUseCaseImpl{
		addressrepository: addressrepository,
	}
}

var currentFilePath = "internal/pkg/address/usecase/usecase.go"

func (uc *AddressUseCaseImpl) GetUserAddresses(ctx context.Context, params addressdto.AddressFilter) (res []*addressdto.AddressResponse, err *helper.ErrorStruct) {
	addressRes, addressErr := uc.addressrepository.GetAddresses(ctx, dao.AddressFilter{
		UserID:       params.UserID,
		AddressTitle: params.AddressTitle,
	})
	if addressErr != nil {
		helper.Logger(currentFilePath, helper.LoggerLevelError, fmt.Sprintf("Error at GetUserAddresses : %s", addressErr.Error()))
		return res, &helper.ErrorStruct{
			Err:  addressErr,
			Code: fiber.StatusBadRequest,
		}
	}
	for _, address := range addressRes {
		res = append(res, &addressdto.AddressResponse{
			ID:            address.ID,
			AddressTitle:  address.AddressTitle,
			Recipient:     address.Recipient,
			PhoneNumber:   address.PhoneNumber,
			AddressDetail: address.AddressDetail,
		})
	}
	return res, nil
}

func (uc *AddressUseCaseImpl) GetUserAddressByID(ctx context.Context, userID, addressID string) (res *addressdto.AddressResponse, err *helper.ErrorStruct) {
	addressRes, addressErr := uc.addressrepository.GetAddressByID(ctx, userID, addressID)
	if addressErr != nil {
		helper.Logger(currentFilePath, helper.LoggerLevelError, fmt.Sprintf("Error at GetUserAddressByID : %s", addressErr.Error()))
		return res, &helper.ErrorStruct{
			Err:  addressErr,
			Code: fiber.StatusNotFound,
		}
	}
	res = &addressdto.AddressResponse{
		ID:            addressRes.ID,
		AddressTitle:  addressRes.AddressTitle,
		Recipient:     addressRes.Recipient,
		PhoneNumber:   addressRes.PhoneNumber,
		AddressDetail: addressRes.AddressDetail,
	}
	return res, nil
}

func (uc *AddressUseCaseImpl) CreateAddress(ctx context.Context, data addressdto.AddressCreate) (res uint, err *helper.ErrorStruct) {
	if validateErr := helper.Validate.Struct(&data); validateErr != nil {
		log.Println(validateErr)
		return res, &helper.ErrorStruct{
			Err:  validateErr,
			Code: fiber.StatusBadRequest,
		}
	}
	res, addressErr := uc.addressrepository.CreateAddress(ctx, dao.Address{
		UserID:        data.UserID,
		AddressTitle:  data.AddressTitle,
		Recipient:     data.Recipient,
		PhoneNumber:   data.PhoneNumber,
		AddressDetail: data.AddressDetail,
	})
	if addressErr != nil {
		helper.Logger(currentFilePath, helper.LoggerLevelError, fmt.Sprintf("Error at CreateAddress : %s", addressErr.Error()))
		return res, &helper.ErrorStruct{
			Err:  addressErr,
			Code: fiber.StatusBadRequest,
		}
	}
	return res, nil
}

func (uc *AddressUseCaseImpl) UpdateAddressByID(ctx context.Context, userID, addressID string, data addressdto.AddressUpdate) *helper.ErrorStruct {
	if validateErr := helper.Validate.Struct(&data); validateErr != nil {
		log.Println(validateErr)
		return &helper.ErrorStruct{
			Err:  validateErr,
			Code: fiber.StatusBadRequest,
		}
	}
	addressErr := uc.addressrepository.UpdateAddressByID(ctx, userID, addressID, dao.Address{
		Recipient:     data.Recipient,
		PhoneNumber:   data.PhoneNumber,
		AddressDetail: data.AddressDetail,
	})
	if addressErr != nil {
		helper.Logger(currentFilePath, helper.LoggerLevelError, fmt.Sprintf("Error at UpdateAddressByID : %s", addressErr.Error()))
		if addressErr == gorm.ErrRecordNotFound {
			return &helper.ErrorStruct{
				Err:  addressErr,
				Code: fiber.StatusNotFound,
			}
		}
		return &helper.ErrorStruct{
			Err:  addressErr,
			Code: fiber.StatusBadRequest,
		}
	}
	return nil
}

func (uc *AddressUseCaseImpl) DeleteAddressByID(ctx context.Context, userID, addressID string) *helper.ErrorStruct {
	addressErr := uc.addressrepository.DeleteAddressByID(ctx, userID, addressID)
	if addressErr != nil {
		helper.Logger(currentFilePath, helper.LoggerLevelError, fmt.Sprintf("Error at DeleteAddressByID : %s", addressErr.Error()))
		if addressErr == gorm.ErrRecordNotFound {
			return &helper.ErrorStruct{
				Err:  addressErr,
				Code: fiber.StatusNotFound,
			}
		}
		return &helper.ErrorStruct{
			Err:  addressErr,
			Code: fiber.StatusBadRequest,
		}
	}
	return nil
}
