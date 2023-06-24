package userusecase

import (
	"context"
	"final_project/internal/dao"
	"final_project/internal/helper"
	userdto "final_project/internal/pkg/user/dto"
	userrepository "final_project/internal/pkg/user/repository"
	"final_project/internal/utils"
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var currentFilePath = "internal/pkg/store/usecase/usecase.go"

type UserUseCase interface {
	GetUserByID(ctx context.Context, userID string) (res *userdto.UserResponse, err *helper.ErrorStruct)
	UpdateUserByID(ctx context.Context, userID string, data userdto.UserUpdate) *helper.ErrorStruct
}

type UserUseCaseImpl struct {
	userrepository userrepository.UserRepository
}

func NewUserUseCase(userrepository userrepository.UserRepository) UserUseCase {
	return &UserUseCaseImpl{
		userrepository: userrepository,
	}
}

func (uc *UserUseCaseImpl) GetUserByID(ctx context.Context, userID string) (res *userdto.UserResponse, err *helper.ErrorStruct) {
	userRes, userErr := uc.userrepository.GetUserByID(ctx, userID)
	if userErr != nil {
		helper.Logger(currentFilePath, helper.LoggerLevelError, fmt.Sprintf("Erorr at GetUserByID : %s", userErr.Error()))
		return res, &helper.ErrorStruct{
			Err:  userErr,
			Code: fiber.StatusNotFound,
		}
	}
	res = &userdto.UserResponse{
		ID:          userRes.ID,
		Name:        userRes.Name,
		PhoneNumber: userRes.PhoneNumber,
		DateOfBirth: utils.DateToString(userRes.DateOfBirth),
		Sex:         userRes.Sex,
		About:       userRes.About,
		Occupation:  userRes.Occupation,
		Email:       userRes.Email,
	}
	return res, nil
}

func (uc *UserUseCaseImpl) UpdateUserByID(ctx context.Context, userID string, data userdto.UserUpdate) *helper.ErrorStruct {
	if errValidate := helper.Validate.Struct(data); errValidate != nil {
		log.Println(errValidate)
		return &helper.ErrorStruct{
			Err:  errValidate,
			Code: fiber.StatusBadRequest,
		}
	}
	if data.Password != "" {
		hashPass, hashErr := utils.HashPassword(data.Password)
		if hashErr != nil {
			return &helper.ErrorStruct{
				Err:  hashErr,
				Code: fiber.StatusBadRequest,
			}
		}
		data.Password = hashPass
	}
	var newDoB time.Time
	if data.DateOfBirth != "" {
		parsedDate, parseErr := utils.ParseDate(data.DateOfBirth)
		if parseErr != nil {
			return &helper.ErrorStruct{
				Err:  parseErr,
				Code: fiber.StatusBadRequest,
			}
		}
		newDoB = parsedDate
	}

	userErr := uc.userrepository.UpdateUserByID(ctx, userID, dao.User{
		Name:        data.Name,
		Password:    data.Password,
		PhoneNumber: data.PhoneNumber,
		DateOfBirth: newDoB,
		About:       data.About,
		Occupation:  data.Occupation,
		Email:       data.Email,
		ProvinceID:  data.ProvinceID,
		CityID:      data.CityID,
	})
	if userErr != nil {
		helper.Logger(currentFilePath, helper.LoggerLevelError, fmt.Sprintf("Error at UpdateUserByID: %s", userErr.Error()))
		if userErr == gorm.ErrRecordNotFound {
			return &helper.ErrorStruct{
				Code: fiber.StatusNotFound,
				Err:  userErr,
			}
		}
		return &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  userErr,
		}
	}
	return nil
}
