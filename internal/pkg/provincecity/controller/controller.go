package provincecitycontroller

import (
	"final_project/internal/helper"
	provincecityusecase "final_project/internal/pkg/provincecity/usecase"

	"github.com/gofiber/fiber/v2"
)

type ProvinceCityController interface {
	GetProvinces(ctx *fiber.Ctx) error
	GetProvinceByID(ctx *fiber.Ctx) error
	GetCitiesByProvinceID(ctx *fiber.Ctx) error
	GetCityByID(ctx *fiber.Ctx) error
}

type ProvinceCityControllerImpl struct {
	provincecityusecase provincecityusecase.ProvinceCityUseCase
}

func NewProvinceCityController(provincecityusecase provincecityusecase.ProvinceCityUseCase) ProvinceCityController {
	return &ProvinceCityControllerImpl{
		provincecityusecase: provincecityusecase,
	}
}

func (cn *ProvinceCityControllerImpl) GetProvinces(ctx *fiber.Ctx) error {
	res, err := cn.provincecityusecase.GetProvinces()
	if err != nil {
		return helper.ResponseBuilder(*ctx, false, helper.GETDATAFAILED, err.Err.Error(), nil, err.Code)
	}
	return helper.ResponseBuilder(*ctx, true, helper.GETDATASUCCESS, nil, res, fiber.StatusOK)
}

func (cn *ProvinceCityControllerImpl) GetProvinceByID(ctx *fiber.Ctx) error {
	provinceID := ctx.Params("province_id")
	if provinceID == "" {
		return helper.ResponseBuilder(*ctx, false, helper.GETDATAFAILED, fiber.ErrBadRequest.Message, nil, fiber.StatusBadRequest)
	}
	res, err := cn.provincecityusecase.GetProvinceByID(provinceID)
	if err != nil {
		return helper.ResponseBuilder(*ctx, false, helper.GETDATAFAILED, err.Err.Error(), nil, err.Code)
	}
	return helper.ResponseBuilder(*ctx, true, helper.GETDATASUCCESS, nil, res, fiber.StatusOK)
}

func (cn *ProvinceCityControllerImpl) GetCitiesByProvinceID(ctx *fiber.Ctx) error {
	provinceID := ctx.Params("province_id")
	if provinceID == "" {
		return helper.ResponseBuilder(*ctx, false, helper.GETDATAFAILED, fiber.ErrBadRequest.Message, nil, fiber.StatusBadRequest)
	}
	res, err := cn.provincecityusecase.GetCitiesByProvinceID(provinceID)
	if err != nil {
		return helper.ResponseBuilder(*ctx, false, helper.GETDATAFAILED, err.Err.Error(), nil, err.Code)
	}
	return helper.ResponseBuilder(*ctx, true, helper.GETDATASUCCESS, nil, res, fiber.StatusOK)
}

func (cn *ProvinceCityControllerImpl) GetCityByID(ctx *fiber.Ctx) error {
	cityID := ctx.Params("city_id")
	if cityID == "" {
		return helper.ResponseBuilder(*ctx, false, helper.GETDATAFAILED, fiber.ErrBadRequest.Message, nil, fiber.StatusBadRequest)
	}
	res, err := cn.provincecityusecase.GetCityByID(cityID)
	if err != nil {
		return helper.ResponseBuilder(*ctx, false, helper.GETDATAFAILED, err.Err.Error(), nil, err.Code)
	}
	return helper.ResponseBuilder(*ctx, true, helper.GETDATASUCCESS, nil, res, fiber.StatusOK)
}
