package categorycontroller

import (
	"final_project/internal/helper"
	categorydto "final_project/internal/pkg/category/dto"
	categoryusecase "final_project/internal/pkg/category/usecase"

	"github.com/gofiber/fiber/v2"
)

type CategoryController interface {
	GetCategories(ctx *fiber.Ctx) error
	GetCategoryByID(ctx *fiber.Ctx) error
	CreateCategory(ctx *fiber.Ctx) error
	UpdateCategoryByID(ctx *fiber.Ctx) error
	DeleteCategoryByID(ctx *fiber.Ctx) error
}

type CategoryControllerImpl struct {
	categoryusecase categoryusecase.CategoryUseCase
}

func NewCategoryController(categoryusecase categoryusecase.CategoryUseCase) CategoryController {
	return &CategoryControllerImpl{
		categoryusecase: categoryusecase,
	}
}

func (cn *CategoryControllerImpl) GetCategories(ctx *fiber.Ctx) error {
	c := ctx.Context()
	filter := new(categorydto.CategoryFilter)
	if err := ctx.QueryParser(filter); err != nil {
		return helper.ResponseBuilder(*ctx, false, helper.GETDATAFAILED, err.Error(), nil, fiber.StatusBadRequest)
	}
	res, err := cn.categoryusecase.GetCategories(c, categorydto.CategoryFilter{
		Name:   filter.Name,
		Limit:  filter.Limit,
		Offset: filter.Offset,
	})
	if err != nil {
		return helper.ResponseBuilder(*ctx, false, helper.GETDATAFAILED, err.Err.Error(), nil, err.Code)
	}
	return helper.ResponseBuilder(*ctx, true, helper.GETDATASUCCESS, nil, res, fiber.StatusOK)
}

func (cn *CategoryControllerImpl) GetCategoryByID(ctx *fiber.Ctx) error {
	c := ctx.Context()
	categoryID := ctx.Params("category_id")
	if categoryID == "" {
		return helper.ResponseBuilder(*ctx, false, helper.GETDATAFAILED, fiber.ErrBadGateway.Message, nil, fiber.StatusBadRequest)
	}
	res, err := cn.categoryusecase.GetCategoryByID(c, categoryID)
	if err != nil {
		return helper.ResponseBuilder(*ctx, false, helper.GETDATAFAILED, err.Err.Error(), nil, err.Code)
	}
	return helper.ResponseBuilder(*ctx, true, helper.GETDATAFAILED, nil, res, fiber.StatusOK)
}

func (cn *CategoryControllerImpl) CreateCategory(ctx *fiber.Ctx) error {
	c := ctx.Context()
	data := new(categorydto.CategoryCreateOrUpdate)
	if err := ctx.BodyParser(&data); err != nil {
		return helper.ResponseBuilder(*ctx, false, helper.POSTDATAFAILED, err.Error(), nil, fiber.StatusBadRequest)
	}
	res, err := cn.categoryusecase.CreateCategory(c, *data)
	if err != nil {
		return helper.ResponseBuilder(*ctx, false, helper.POSTDATAFAILED, err.Err.Error(), nil, err.Code)
	}
	return helper.ResponseBuilder(*ctx, true, helper.POSTDATASUCCESS, nil, res, fiber.StatusCreated)
}

func (cn *CategoryControllerImpl) UpdateCategoryByID(ctx *fiber.Ctx) error {
	c := ctx.Context()
	categoryID := ctx.Params("category_id")
	if categoryID == "" {
		return helper.ResponseBuilder(*ctx, false, helper.PUTDATAFAILED, fiber.ErrBadRequest.Message, nil, fiber.StatusBadRequest)
	}
	data := new(categorydto.CategoryCreateOrUpdate)
	if err := ctx.BodyParser(&data); err != nil {
		return helper.ResponseBuilder(*ctx, false, helper.PUTDATAFAILED, err.Error(), nil, fiber.StatusBadRequest)
	}
	err := cn.categoryusecase.UpdateCategoryByID(c, categoryID, *data)
	if err != nil {
		return helper.ResponseBuilder(*ctx, false, helper.PUTDATAFAILED, err.Err.Error(), nil, err.Code)
	}
	return helper.ResponseBuilder(*ctx, true, helper.PUTDATASUCCESS, nil, nil, fiber.StatusNoContent)
}

func (cn *CategoryControllerImpl) DeleteCategoryByID(ctx *fiber.Ctx) error {
	c := ctx.Context()
	categoryID := ctx.Params("category_id")
	if categoryID == "" {
		return helper.ResponseBuilder(*ctx, false, helper.DELETEDATAFAILED, fiber.ErrBadRequest.Message, nil, fiber.StatusBadRequest)
	}
	err := cn.categoryusecase.DeleteCategoryByID(c, categoryID)
	if err != nil {
		return helper.ResponseBuilder(*ctx, false, helper.DELETEDATAFAILED, err.Err.Error(), nil, err.Code)
	}
	return helper.ResponseBuilder(*ctx, true, helper.DELETEDATASUCCESS, nil, nil, fiber.StatusNoContent)
}
