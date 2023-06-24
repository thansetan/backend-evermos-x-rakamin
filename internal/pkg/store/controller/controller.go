package storecontroller

import (
	"final_project/internal/helper"
	storedto "final_project/internal/pkg/store/dto"
	storeusecase "final_project/internal/pkg/store/usecase"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type StoreController interface {
	GetStoreByID(ctx *fiber.Ctx) error
	GetAllStores(ctx *fiber.Ctx) error
	GetMyStore(ctx *fiber.Ctx) error
	UpdateStoreByID(ctx *fiber.Ctx) error
}

type StoreControllerImpl struct {
	storeusecase storeusecase.StoreUseCase
}

func NewStoreController(storeusecase storeusecase.StoreUseCase) StoreController {
	return &StoreControllerImpl{
		storeusecase: storeusecase,
	}
}

func (cn *StoreControllerImpl) GetStoreByID(ctx *fiber.Ctx) error {
	c := ctx.Context()
	storeID := ctx.Params("store_id")
	if storeID == "" {
		return helper.ResponseBuilder(*ctx, false, helper.GETDATAFAILED, fiber.ErrBadRequest.Message, nil, fiber.StatusBadRequest)
	}

	res, err := cn.storeusecase.GetStoreByID(c, storeID)
	if err != nil {
		return helper.ResponseBuilder(*ctx, false, helper.GETDATAFAILED, err.Err.Error(), nil, err.Code)
	}
	return helper.ResponseBuilder(*ctx, true, helper.GETDATASUCCESS, nil, res, fiber.StatusOK)
}

func (cn *StoreControllerImpl) GetAllStores(ctx *fiber.Ctx) error {
	c := ctx.Context()
	filter := new(storedto.StoreFilter)
	if err := ctx.QueryParser(filter); err != nil {
		log.Println(err)
		return helper.ResponseBuilder(*ctx, false, helper.GETDATAFAILED, err.Error(), nil, fiber.StatusBadRequest)
	}
	res, err := cn.storeusecase.GetAllStores(c, storedto.StoreFilter{
		Limit: filter.Limit,
		Page:  filter.Page,
	})
	if err != nil {
		return helper.ResponseBuilder(*ctx, false, helper.GETDATAFAILED, err.Err.Error(), nil, err.Code)
	}
	return helper.ResponseBuilder(*ctx, true, helper.GETDATASUCCESS, nil, res, fiber.StatusOK)
}

func (cn *StoreControllerImpl) GetMyStore(ctx *fiber.Ctx) error {
	c := ctx.Context()
	userID := fmt.Sprintf("%v", ctx.Locals("userID"))
	if userID == "" {
		return helper.ResponseBuilder(*ctx, false, helper.GETDATAFAILED, "UNAUTHORIZED", nil, fiber.StatusUnauthorized)
	}
	res, err := cn.storeusecase.GetMyStore(c, userID)
	if err != nil {
		return helper.ResponseBuilder(*ctx, false, helper.GETDATAFAILED, err.Err.Error(), nil, err.Code)
	}
	return helper.ResponseBuilder(*ctx, true, helper.GETDATASUCCESS, nil, res, fiber.StatusOK)
}

func (cn *StoreControllerImpl) UpdateStoreByID(ctx *fiber.Ctx) error {
	c := ctx.Context()
	userID := fmt.Sprintf("%v", ctx.Locals("userID"))
	if userID == "" {
		return helper.ResponseBuilder(*ctx, false, helper.PUTDATAFAILED, "UNAUTHORIZED", nil, fiber.StatusUnauthorized)
	}
	data := new(storedto.StoreUpdate)
	data.StoreName = ctx.FormValue("nama_toko")
	newPhotoUrl := ctx.Locals("photoUrl").(string)
	if newPhotoUrl != "" {
		data.PhotoUrl = newPhotoUrl
	}
	err := cn.storeusecase.UpdateStoreByUserID(c, userID, *data)
	if err != nil {
		return helper.ResponseBuilder(*ctx, false, helper.PUTDATAFAILED, err.Err.Error(), nil, err.Code)
	}

	return helper.ResponseBuilder(*ctx, true, helper.PUTDATASUCCESS, nil, nil, fiber.StatusNoContent)
}
