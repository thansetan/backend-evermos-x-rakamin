package categoryusecase

import (
	"context"
	"final_project/internal/dao"
	"final_project/internal/helper"
	categorydto "final_project/internal/pkg/category/dto"
	categoryrepository "final_project/internal/pkg/category/repository"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var currentFilePath = "internal/pkg/category/usecase/usecase.go"

type CategoryUseCase interface {
	GetCategories(ctx context.Context, params categorydto.CategoryFilter) (res []*categorydto.CategoryResponse, err *helper.ErrorStruct)
	GetCategoryByID(ctx context.Context, categoryID string) (res *categorydto.CategoryResponse, err *helper.ErrorStruct)
	CreateCategory(ctx context.Context, data categorydto.CategoryCreateOrUpdate) (res uint, err *helper.ErrorStruct)
	UpdateCategoryByID(ctx context.Context, categoryID string, data categorydto.CategoryCreateOrUpdate) *helper.ErrorStruct
	DeleteCategoryByID(ctx context.Context, categoryID string) *helper.ErrorStruct
}

type CategoryUseCaseImpl struct {
	categoryrepository categoryrepository.CategoryRepository
}

func NewCategoryUseCase(categoryrepository categoryrepository.CategoryRepository) CategoryUseCase {
	return &CategoryUseCaseImpl{
		categoryrepository: categoryrepository,
	}
}

func (uc *CategoryUseCaseImpl) GetCategories(ctx context.Context, params categorydto.CategoryFilter) (res []*categorydto.CategoryResponse, err *helper.ErrorStruct) {
	if params.Limit < 1 {
		params.Limit = 10
	}

	if params.Offset < 1 {
		params.Offset = 0
	} else {
		params.Offset = (params.Offset - 1) * params.Limit
	}
	categoryRes, categoryErr := uc.categoryrepository.GetCategories(ctx, dao.CategoryFilter{
		Name:   params.Name,
		Limit:  params.Limit,
		Offset: params.Offset,
	})
	if categoryErr != nil {
		helper.Logger(currentFilePath, helper.LoggerLevelError, fmt.Sprintf("Error at GetCategories: %s", categoryErr.Error()))
		return res, &helper.ErrorStruct{
			Err:  categoryErr,
			Code: fiber.StatusBadRequest,
		}
	}
	for _, category := range categoryRes {
		res = append(res, &categorydto.CategoryResponse{
			ID:           category.ID,
			CategoryName: category.CategoryName,
		})
	}
	return res, nil
}

func (uc *CategoryUseCaseImpl) GetCategoryByID(ctx context.Context, categoryID string) (res *categorydto.CategoryResponse, err *helper.ErrorStruct) {
	categoryRes, categoryErr := uc.categoryrepository.GetCategoryByID(ctx, categoryID)
	if categoryErr != nil {
		helper.Logger(currentFilePath, helper.LoggerLevelError, fmt.Sprintf("Error at GetCategoryByID: %s", categoryErr.Error()))
		return res, &helper.ErrorStruct{
			Err:  categoryErr,
			Code: fiber.StatusNotFound,
		}
	}
	res = &categorydto.CategoryResponse{
		ID:           categoryRes.ID,
		CategoryName: categoryRes.CategoryName,
	}
	return res, nil
}

func (uc *CategoryUseCaseImpl) CreateCategory(ctx context.Context, data categorydto.CategoryCreateOrUpdate) (res uint, err *helper.ErrorStruct) {
	if validateErr := helper.Validate.Struct(data); validateErr != nil {
		log.Println(validateErr)
		return res, &helper.ErrorStruct{
			Err:  validateErr,
			Code: fiber.StatusBadRequest,
		}
	}
	res, categoryErr := uc.categoryrepository.CreateCategory(ctx, dao.Category{
		CategoryName: data.CategoryName,
	})
	if categoryErr != nil {
		helper.Logger(currentFilePath, helper.LoggerLevelError, fmt.Sprintf("Error at CreateCategory: %s", categoryErr.Error()))
		return res, &helper.ErrorStruct{
			Err:  categoryErr,
			Code: fiber.StatusBadRequest,
		}
	}
	return res, nil
}

func (uc *CategoryUseCaseImpl) UpdateCategoryByID(ctx context.Context, categoryID string, data categorydto.CategoryCreateOrUpdate) *helper.ErrorStruct {
	if validateErr := helper.Validate.Struct(data); validateErr != nil {
		log.Println(validateErr)
		return &helper.ErrorStruct{
			Err:  validateErr,
			Code: fiber.StatusBadRequest,
		}
	}
	categoryErr := uc.categoryrepository.UpdateCategoryByID(ctx, categoryID, dao.Category{
		CategoryName: data.CategoryName,
	})
	if categoryErr != nil {
		helper.Logger(currentFilePath, helper.LoggerLevelError, fmt.Sprintf("Error at UpdateCategoryByID: %s", categoryErr.Error()))
		if categoryErr == gorm.ErrRecordNotFound {
			return &helper.ErrorStruct{
				Err:  categoryErr,
				Code: fiber.StatusNotFound,
			}
		}
		return &helper.ErrorStruct{
			Err:  categoryErr,
			Code: fiber.StatusBadRequest,
		}
	}
	return nil
}

func (uc *CategoryUseCaseImpl) DeleteCategoryByID(ctx context.Context, categoryID string) *helper.ErrorStruct {
	categoryErr := uc.categoryrepository.DeleteCategoryByID(ctx, categoryID)
	if categoryErr != nil {
		helper.Logger(currentFilePath, helper.LoggerLevelError, fmt.Sprintf("Error at DeleteCategoryByID: %s", categoryErr.Error()))
		if categoryErr == gorm.ErrRecordNotFound {
			return &helper.ErrorStruct{
				Err:  categoryErr,
				Code: fiber.StatusNotFound,
			}
		}
		return &helper.ErrorStruct{
			Err:  categoryErr,
			Code: fiber.StatusBadRequest,
		}
	}
	return nil
}
