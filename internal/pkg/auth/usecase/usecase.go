package authusecase

import (
	"context"
	"errors"
	"final_project/internal/dao"
	"final_project/internal/helper"
	authdto "final_project/internal/pkg/auth/dto"
	authrepository "final_project/internal/pkg/auth/repository"
	provincecitydto "final_project/internal/pkg/provincecity/dto"
	provincecityrepository "final_project/internal/pkg/provincecity/repository"
	storerepository "final_project/internal/pkg/store/repository"
	"final_project/internal/utils"
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

var currentFilePath = "internal/pkg/auth/usecase/usecase.go"

type AuthUseCase interface {
	Register(ctx context.Context, data authdto.Register) (res string, err *helper.ErrorStruct)
	Login(ctx context.Context, data authdto.Login) (res authdto.LoginResponse, err *helper.ErrorStruct)
}

type AuthUseCaseImpl struct {
	authrepository         authrepository.AuthRepository
	storerepository        storerepository.StoreRepository
	provincecityrepository provincecityrepository.ProvinceCityRepository
	db                     *gorm.DB
}

func NewAuthUseCase(authrepository authrepository.AuthRepository,
	storerepository storerepository.StoreRepository,
	provincecityrepository provincecityrepository.ProvinceCityRepository,
	db *gorm.DB) AuthUseCase {
	return &AuthUseCaseImpl{
		authrepository:         authrepository,
		storerepository:        storerepository,
		provincecityrepository: provincecityrepository,
		db:                     db,
	}
}

func (uc *AuthUseCaseImpl) Register(ctx context.Context, data authdto.Register) (res string, err *helper.ErrorStruct) {
	if errValidate := helper.Validate.Struct(data); errValidate != nil {
		log.Println(errValidate)
		return res, &helper.ErrorStruct{
			Err:  errValidate,
			Code: fiber.StatusBadRequest,
		}
	}
	hashedPassword, hashErr := utils.HashPassword(data.Password)
	if hashErr != nil {
		return "", &helper.ErrorStruct{
			Err:  hashErr,
			Code: fiber.StatusBadRequest,
		}
	}
	parsedDate, parseErr := utils.ParseDate(data.DateOfBirth)
	if parseErr != nil {
		return "", &helper.ErrorStruct{
			Err:  parseErr,
			Code: fiber.StatusBadRequest,
		}
	}
	registrationData := dao.User{
		Name:        data.Name,
		Password:    hashedPassword,
		PhoneNumber: data.PhoneNumber,
		DateOfBirth: parsedDate,
		Occupation:  data.Occupation,
		Email:       data.Email,
		ProvinceID:  data.ProvinceID,
		CityID:      data.CityID,
		Sex:         data.Sex,
	}
	tx := uc.db.Begin()
	userID, repoErr := uc.authrepository.Register(ctx, registrationData, tx)
	if helper.CheckDuplicateData(repoErr) {
		helper.Logger(currentFilePath, helper.LoggerLevelError, fmt.Sprintf("Erorr at Register : %s", repoErr.Error()))
		return res, &helper.ErrorStruct{
			Err:  errors.New("email/phone number already in use"),
			Code: fiber.StatusBadRequest,
		}
	}
	if repoErr != nil {
		helper.Logger(currentFilePath, helper.LoggerLevelError, fmt.Sprintf("Erorr at Register : %s", repoErr.Error()))
		return res, &helper.ErrorStruct{
			Err:  repoErr,
			Code: fiber.StatusBadRequest,
		}
	}
	storeData := dao.Store{
		UserID:    userID,
		StoreName: registrationData.Name + "'s Store",
	}
	_, storeErr := uc.storerepository.CreateStore(ctx, storeData, tx)
	if storeErr != nil {
		tx.Rollback()
		helper.Logger(currentFilePath, helper.LoggerLevelError, fmt.Sprintf("Error at store creation: %s", storeErr.Error()))
		return res, &helper.ErrorStruct{
			Err:  storeErr,
			Code: fiber.StatusBadRequest,
		}
	}
	tx.Commit()
	return "Account registered successfully", nil
}

func (uc *AuthUseCaseImpl) Login(ctx context.Context, data authdto.Login) (res authdto.LoginResponse, err *helper.ErrorStruct) {
	if errValidate := helper.Validate.Struct(data); errValidate != nil {
		log.Println(errValidate)
		return res, &helper.ErrorStruct{
			Err:  errValidate,
			Code: fiber.StatusBadRequest,
		}
	}
	user, repoErr := uc.authrepository.Login(ctx, dao.UserLogin{
		PhoneNumber: data.PhoneNumber,
		Password:    data.Password,
	})
	if repoErr != nil {
		helper.Logger(currentFilePath, helper.LoggerLevelError, fmt.Sprintf("Error at Login: %s", repoErr.Error()))
		return res, &helper.ErrorStruct{
			Err:  errors.New("invalid phone number/password"),
			Code: fiber.StatusUnauthorized,
		}
	}
	isValid := utils.CheckPasswordHash(data.Password, user.Password)
	if !isValid {
		return res, &helper.ErrorStruct{
			Err:  errors.New("invalid phone number/password"),
			Code: fiber.StatusUnauthorized,
		}
	}
	claims := jwt.MapClaims{
		"userID":  user.ID,
		"storeID": user.Store.ID,
		"isAdmin": user.IsAdmin,
		"exp":     time.Now().Add(time.Hour * 48).Unix(),
	}
	token, jwtErr := utils.GenerateJWT(claims)
	if jwtErr != nil {
		return res, &helper.ErrorStruct{
			Err:  jwtErr,
			Code: fiber.StatusBadRequest,
		}
	}

	res = authdto.LoginResponse{
		Name:        user.Name,
		PhoneNumber: user.PhoneNumber,
		Sex:         user.Sex,
		DateOfBirth: user.DateOfBirth,
		About:       user.About,
		Occupation:  user.Occupation,
		Email:       user.Email,
		Token:       token,
	}
	city, _ := uc.provincecityrepository.GetCityByID(user.CityID)
	province, _ := uc.provincecityrepository.GetProvinceByID(user.ProvinceID)
	res.City = provincecitydto.City{
		ID:         city.ID,
		ProvinceID: city.ProvinceID,
		Name:       city.Name,
	}
	res.Province = provincecitydto.Province{
		ID:   province.ID,
		Name: province.Name,
	}
	return res, nil
}
