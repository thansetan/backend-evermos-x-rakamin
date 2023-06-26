package productcontroller

import (
	"final_project/internal/helper"
	productdto "final_project/internal/pkg/product/dto"
	productusecase "final_project/internal/pkg/product/usecase"
	"final_project/internal/utils"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type ProductController interface {
	CreateProduct(ctx *fiber.Ctx) error
	GetProducts(ctx *fiber.Ctx) error
	GetProductByID(ctx *fiber.Ctx) error
	UpdateProductByID(ctx *fiber.Ctx) error
	DeleteProductByID(ctx *fiber.Ctx) error
}

type ProductControllerImpl struct {
	productusecase productusecase.ProductUseCase
}

func NewProductController(productusecase productusecase.ProductUseCase) ProductController {
	return &ProductControllerImpl{
		productusecase: productusecase,
	}
}

func (cn *ProductControllerImpl) CreateProduct(ctx *fiber.Ctx) error {
	c := ctx.Context()
	userID := fmt.Sprintf("%v", ctx.Locals("userID"))
	if userID == "" {
		return helper.ResponseBuilder(*ctx, false, helper.POSTDATAFAILED, "UNAUTHORIZED", nil, fiber.StatusUnauthorized)
	}
	storeID := fmt.Sprintf("%v", ctx.Locals("storeID"))
	storeIDUint, parseErr := utils.StringToUint(storeID)
	if parseErr != nil {
		return helper.ResponseBuilder(*ctx, false, helper.POSTDATAFAILED, parseErr.Error(), nil, fiber.StatusBadRequest)
	}
	data := new(productdto.ProductCreate)
	data.StoreID = storeIDUint
	if err := ctx.BodyParser(data); err != nil {
		return helper.ResponseBuilder(*ctx, false, helper.POSTDATAFAILED, err.Error(), nil, fiber.StatusBadRequest)
	}
	photos := ctx.Locals("photoUrls").([]string)
	res, err := cn.productusecase.CreateProduct(c, *data, photos)
	if err != nil {
		return helper.ResponseBuilder(*ctx, false, helper.POSTDATAFAILED, err.Err.Error(), nil, err.Code)
	}
	return helper.ResponseBuilder(*ctx, true, helper.POSTDATASUCCESS, nil, res, fiber.StatusCreated)
}

func (cn *ProductControllerImpl) GetProducts(ctx *fiber.Ctx) error {
	c := ctx.Context()
	filter := new(productdto.ProductFilter)
	if err := ctx.QueryParser(filter); err != nil {
		return helper.ResponseBuilder(*ctx, false, helper.GETDATAFAILED, err.Error(), nil, fiber.StatusBadRequest)
	}
	res, err := cn.productusecase.GetProducts(c, *filter)
	if err != nil {
		return helper.ResponseBuilder(*ctx, false, helper.GETDATAFAILED, err.Err.Error(), nil, err.Code)
	}
	return helper.ResponseBuilder(*ctx, true, helper.GETDATASUCCESS, nil, res, fiber.StatusOK)
}

func (cn *ProductControllerImpl) GetProductByID(ctx *fiber.Ctx) error {
	c := ctx.Context()
	productID := ctx.Params("product_id")
	if productID == "" {
		return helper.ResponseBuilder(*ctx, false, helper.GETDATAFAILED, fiber.ErrBadRequest.Message, nil, fiber.StatusBadRequest)
	}
	res, err := cn.productusecase.GetProductByID(c, productID)
	if err != nil {
		return helper.ResponseBuilder(*ctx, false, helper.GETDATAFAILED, err.Err.Error(), nil, err.Code)
	}
	return helper.ResponseBuilder(*ctx, true, helper.GETDATASUCCESS, nil, res, fiber.StatusOK)
}

func (cn *ProductControllerImpl) UpdateProductByID(ctx *fiber.Ctx) error {
	c := ctx.Context()
	userID := fmt.Sprintf("%v", ctx.Locals("userID"))
	if userID == "" {
		return helper.ResponseBuilder(*ctx, false, helper.POSTDATAFAILED, "UNAUTHORIZED", nil, fiber.StatusUnauthorized)
	}
	productID := ctx.Params("product_id")
	if productID == "" {
		return helper.ResponseBuilder(*ctx, false, helper.PUTDATAFAILED, fiber.ErrBadRequest.Message, nil, fiber.StatusBadRequest)
	}
	storeID := fmt.Sprintf("%v", ctx.Locals("storeID"))
	storeIDUint, parseErr := utils.StringToUint(storeID)
	if parseErr != nil {
		return helper.ResponseBuilder(*ctx, false, helper.PUTDATAFAILED, parseErr.Error(), nil, fiber.StatusBadRequest)
	}
	data := new(productdto.ProductUpdate)
	data.StoreID = storeIDUint
	if err := ctx.BodyParser(data); err != nil {
		return helper.ResponseBuilder(*ctx, false, helper.PUTDATAFAILED, err.Error(), nil, fiber.StatusBadRequest)
	}
	photos := ctx.Locals("photoUrls").([]string)
	err := cn.productusecase.UpdateProductByID(c, productID, *data, photos)
	if err != nil {
		return helper.ResponseBuilder(*ctx, false, helper.PUTDATAFAILED, err.Err.Error(), nil, err.Code)
	}
	return helper.ResponseBuilder(*ctx, true, helper.PUTDATASUCCESS, nil, nil, fiber.StatusNoContent)
}

func (cn *ProductControllerImpl) DeleteProductByID(ctx *fiber.Ctx) error {
	c := ctx.Context()
	userID := fmt.Sprintf("%v", ctx.Locals("userID"))
	if userID == "" {
		return helper.ResponseBuilder(*ctx, false, helper.DELETEDATAFAILED, "UNAUTHORIZED", nil, fiber.StatusUnauthorized)
	}
	productID := ctx.Params("product_id")
	if productID == "" {
		return helper.ResponseBuilder(*ctx, false, helper.DELETEDATAFAILED, fiber.ErrBadRequest.Message, nil, fiber.StatusBadRequest)
	}
	storeID := fmt.Sprintf("%v", ctx.Locals("storeID"))
	err := cn.productusecase.DeleteProductByID(c, storeID, productID)
	if err != nil {
		return helper.ResponseBuilder(*ctx, false, helper.DELETEDATAFAILED, err.Err.Error(), nil, err.Code)
	}
	return helper.ResponseBuilder(*ctx, true, helper.DELETEDATASUCCESS, nil, nil, fiber.StatusNoContent)
}
