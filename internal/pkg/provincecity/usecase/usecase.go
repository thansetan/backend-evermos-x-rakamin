package provincecityusecase

import (
	"errors"
	"final_project/internal/helper"
	provincecitydto "final_project/internal/pkg/provincecity/dto"
	provincecityrepository "final_project/internal/pkg/provincecity/repository"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

var currentFilePath = "internal/pkg/provincecity/usecase/usecase.go"

type ProvinceCityUseCase interface {
	GetProvinces() (res []*provincecitydto.Province, err *helper.ErrorStruct)
	GetProvinceByID(provinceID string) (res *provincecitydto.Province, err *helper.ErrorStruct)
	GetCitiesByProvinceID(provinceID string) (res []*provincecitydto.City, err *helper.ErrorStruct)
	GetCityByID(cityID string) (res *provincecitydto.City, err *helper.ErrorStruct)
}

type ProvinceCityUseCaseImpl struct {
	provincecityrepository provincecityrepository.ProvinceCityRepository
}

func NewProvinceCityUseCase(provincecityrepository provincecityrepository.ProvinceCityRepository) ProvinceCityUseCase {
	return &ProvinceCityUseCaseImpl{
		provincecityrepository: provincecityrepository,
	}
}

func (uc *ProvinceCityUseCaseImpl) GetProvinces() (res []*provincecitydto.Province, err *helper.ErrorStruct) {
	provinceRes, provinceErr := uc.provincecityrepository.GetProvinces()
	if provinceErr != nil {
		helper.Logger(currentFilePath, helper.LoggerLevelError, fmt.Sprintf("Error at GetProvinces: %s", provinceErr.Error()))
		return res, &helper.ErrorStruct{
			Err:  provinceErr,
			Code: fiber.StatusBadRequest,
		}
	}
	for _, province := range provinceRes {
		res = append(res, &provincecitydto.Province{
			ID:   province.ID,
			Name: province.Name,
		})
	}
	return res, nil
}

func (uc *ProvinceCityUseCaseImpl) GetProvinceByID(provinceID string) (res *provincecitydto.Province, err *helper.ErrorStruct) {
	provinceRes, provinceErr := uc.provincecityrepository.GetProvinceByID(provinceID)
	if provinceErr != nil {
		helper.Logger(currentFilePath, helper.LoggerLevelError, fmt.Sprintf("Error at GetProvinceByID: %s", provinceErr.Error()))
		return res, &helper.ErrorStruct{
			Err:  errors.New("province with provided ID can't be found"),
			Code: fiber.StatusNotFound,
		}
	}
	res = &provincecitydto.Province{
		ID:   provinceRes.ID,
		Name: provinceRes.Name,
	}
	return res, nil
}

func (uc *ProvinceCityUseCaseImpl) GetCitiesByProvinceID(provinceID string) (res []*provincecitydto.City, err *helper.ErrorStruct) {
	citiesRes, citiesErr := uc.provincecityrepository.GetCitiesByProvinceID(provinceID)
	if citiesErr != nil {
		helper.Logger(currentFilePath, helper.LoggerLevelError, fmt.Sprintf("Error at GetProvinceByID: %s", citiesErr.Error()))
		return res, &helper.ErrorStruct{
			Err:  errors.New("cities within province with provided ID can't be found"),
			Code: fiber.StatusNotFound,
		}
	}
	for _, city := range citiesRes {
		res = append(res, &provincecitydto.City{
			ID:         city.ID,
			ProvinceID: city.ProvinceID,
			Name:       city.Name,
		})
	}
	return res, nil
}

func (uc *ProvinceCityUseCaseImpl) GetCityByID(cityID string) (res *provincecitydto.City, err *helper.ErrorStruct) {
	cityRes, cityErr := uc.provincecityrepository.GetCityByID(cityID)
	if cityErr != nil {
		helper.Logger(currentFilePath, helper.LoggerLevelError, fmt.Sprintf("Error at GetProvinceByID: %s", cityErr.Error()))
		return res, &helper.ErrorStruct{
			Err:  errors.New("city with provided ID can't be found"),
			Code: fiber.StatusNotFound,
		}
	}
	res = &provincecitydto.City{
		ID:         cityRes.ID,
		ProvinceID: cityRes.ProvinceID,
		Name:       cityRes.Name,
	}
	return res, nil
}
