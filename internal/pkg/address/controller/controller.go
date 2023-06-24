package addresscontroller

import (
	"final_project/internal/helper"
	addressdto "final_project/internal/pkg/address/dto"
	addressusecase "final_project/internal/pkg/address/usecase"
	"final_project/internal/utils"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type AddressController interface {
	GetMyAddresses(ctx *fiber.Ctx) error
	GetMyAddressByID(ctx *fiber.Ctx) error
	CreateAddress(ctx *fiber.Ctx) error
	UpdateMyAddressByID(ctx *fiber.Ctx) error
	DeleteMyAddressByID(ctx *fiber.Ctx) error
}

type AddressControllerImpl struct {
	addressusecase addressusecase.AddressUseCase
}

func NewAddressController(addressusecase addressusecase.AddressUseCase) AddressController {
	return &AddressControllerImpl{
		addressusecase: addressusecase,
	}
}

func (cn *AddressControllerImpl) GetMyAddresses(ctx *fiber.Ctx) error {
	c := ctx.Context()
	userID := fmt.Sprintf("%v", ctx.Locals("userID"))
	if userID == "" {
		return helper.ResponseBuilder(*ctx, false, helper.GETDATAFAILED, "UNAUTHORZED", nil, fiber.StatusUnauthorized)
	}
	filter := new(addressdto.AddressFilter)
	if err := ctx.QueryParser(filter); err != nil {
		log.Println(err)
		return helper.ResponseBuilder(*ctx, false, helper.GETDATAFAILED, err.Error(), nil, fiber.StatusBadRequest)
	}
	res, err := cn.addressusecase.GetUserAddresses(c, addressdto.AddressFilter{
		UserID:       userID,
		AddressTitle: filter.AddressTitle,
	})
	if err != nil {
		return helper.ResponseBuilder(*ctx, false, helper.GETDATAFAILED, err.Err.Error(), nil, err.Code)
	}
	return helper.ResponseBuilder(*ctx, true, helper.GETDATASUCCESS, nil, res, fiber.StatusOK)
}

func (cn *AddressControllerImpl) GetMyAddressByID(ctx *fiber.Ctx) error {
	c := ctx.Context()
	userID := fmt.Sprintf("%v", ctx.Locals("userID"))
	if userID == "" {
		return helper.ResponseBuilder(*ctx, false, helper.GETDATAFAILED, "UNAUTHORIZED", nil, fiber.StatusUnauthorized)
	}
	addressID := ctx.Params("address_id")
	if addressID == "" {
		return helper.ResponseBuilder(*ctx, false, helper.GETDATAFAILED, fiber.ErrBadRequest.Message, nil, fiber.StatusBadRequest)
	}
	res, err := cn.addressusecase.GetUserAddressByID(c, userID, addressID)
	if err != nil {
		return helper.ResponseBuilder(*ctx, false, helper.GETDATAFAILED, err.Err.Error(), nil, err.Code)
	}
	return helper.ResponseBuilder(*ctx, true, helper.GETDATASUCCESS, nil, res, fiber.StatusOK)
}

func (cn *AddressControllerImpl) CreateAddress(ctx *fiber.Ctx) error {
	c := ctx.Context()
	userID := fmt.Sprintf("%v", ctx.Locals("userID"))
	if userID == "" {
		return helper.ResponseBuilder(*ctx, false, helper.POSTDATAFAILED, "UNAUTHORIZED", nil, fiber.StatusUnauthorized)
	}
	userIDUint, uintErr := utils.StringToUint(userID)
	if uintErr != nil {
		return helper.ResponseBuilder(*ctx, false, helper.POSTDATAFAILED, uintErr.Error(), nil, fiber.StatusBadRequest)
	}
	data := new(addressdto.AddressCreate)
	if err := ctx.BodyParser(&data); err != nil {
		return helper.ResponseBuilder(*ctx, false, helper.POSTDATAFAILED, err.Error(), nil, fiber.StatusBadRequest)
	}
	data.UserID = userIDUint
	res, err := cn.addressusecase.CreateAddress(c, *data)
	if err != nil {
		return helper.ResponseBuilder(*ctx, false, helper.POSTDATAFAILED, err.Err.Error(), nil, err.Code)
	}
	return helper.ResponseBuilder(*ctx, true, helper.POSTDATASUCCESS, nil, res, fiber.StatusCreated)
}

func (cn *AddressControllerImpl) UpdateMyAddressByID(ctx *fiber.Ctx) error {
	c := ctx.Context()
	userID := fmt.Sprintf("%v", ctx.Locals("userID"))
	if userID == "" {
		return helper.ResponseBuilder(*ctx, false, helper.PUTDATAFAILED, "UNAUTHORIZED", nil, fiber.StatusUnauthorized)
	}
	addressID := ctx.Params("address_id")
	if addressID == "" {
		return helper.ResponseBuilder(*ctx, false, helper.PUTDATAFAILED, fiber.ErrBadRequest.Message, nil, fiber.StatusBadRequest)
	}
	userIDUint, uintErr := utils.StringToUint(userID)
	if uintErr != nil {
		return helper.ResponseBuilder(*ctx, false, helper.POSTDATAFAILED, uintErr.Error(), nil, fiber.StatusBadRequest)
	}
	data := new(addressdto.AddressUpdate)
	if err := ctx.BodyParser(&data); err != nil {
		return helper.ResponseBuilder(*ctx, false, helper.PUTDATAFAILED, err.Error(), nil, fiber.StatusBadRequest)
	}
	data.UserID = userIDUint
	err := cn.addressusecase.UpdateAddressByID(c, addressID, *data)
	if err != nil {
		return helper.ResponseBuilder(*ctx, false, helper.PUTDATAFAILED, err.Err.Error(), nil, err.Code)
	}
	return helper.ResponseBuilder(*ctx, true, helper.PUTDATASUCCESS, nil, nil, fiber.StatusNoContent)
}

func (cn *AddressControllerImpl) DeleteMyAddressByID(ctx *fiber.Ctx) error {
	c := ctx.Context()
	userID := fmt.Sprintf("%v", ctx.Locals("userID"))
	if userID == "" {
		return helper.ResponseBuilder(*ctx, false, helper.DELETEDATAFAILED, "UNAUTHORIZED", nil, fiber.StatusUnauthorized)
	}
	addressID := ctx.Params("address_id")
	if addressID == "" {
		return helper.ResponseBuilder(*ctx, false, helper.GETDATAFAILED, fiber.ErrBadRequest.Message, nil, fiber.StatusBadRequest)
	}
	err := cn.addressusecase.DeleteAddressByID(c, userID, addressID)
	if err != nil {
		return helper.ResponseBuilder(*ctx, false, helper.DELETEDATAFAILED, err.Err.Error(), nil, err.Code)
	}
	return helper.ResponseBuilder(*ctx, true, helper.DELETEDATASUCCESS, nil, nil, fiber.StatusNoContent)
}
