package provincecityrepository

import (
	"encoding/json"
	"final_project/internal/dao"
	"fmt"
	"io"
	"net/http"
)

type ProvinceCityRepository interface {
	GetProvinces() (res []*dao.Province, err error)
	GetProvinceByID(provinceID string) (res *dao.Province, err error)
	GetCitiesByProvinceID(provinceiD string) (res []*dao.City, err error)
	GetCityByID(cityID string) (res *dao.City, err error)
}

type ProvinceCityRepositoryImpl struct{}

func NewProvinceCityRepository() ProvinceCityRepository {
	return &ProvinceCityRepositoryImpl{}
}

var baseUrl = "https://www.emsifa.com/api-wilayah-indonesia/api/"

func (repo *ProvinceCityRepositoryImpl) GetProvinces() (res []*dao.Province, err error) {
	provinceJson := "provinces.json"
	resp, err := http.Get(baseUrl + provinceJson)
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err := json.Unmarshal(body, &res); err != nil {
		return res, err
	}
	return res, nil
}

func (repo *ProvinceCityRepositoryImpl) GetProvinceByID(provinceID string) (res *dao.Province, err error) {
	provinceIDJson := fmt.Sprintf("province/%s.json", provinceID)
	resp, err := http.Get(baseUrl + provinceIDJson)
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err := json.Unmarshal(body, &res); err != nil {
		return res, err
	}
	return res, nil
}

func (repo *ProvinceCityRepositoryImpl) GetCitiesByProvinceID(provinceID string) (res []*dao.City, err error) {
	citiesByProvinceIDJson := fmt.Sprintf("regencies/%s.json", provinceID)
	resp, err := http.Get(baseUrl + citiesByProvinceIDJson)
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err := json.Unmarshal(body, &res); err != nil {
		return res, err
	}
	return res, nil
}

func (repo *ProvinceCityRepositoryImpl) GetCityByID(cityID string) (res *dao.City, err error) {
	cityIDJson := fmt.Sprintf("regency/%s.json", cityID)
	resp, err := http.Get(baseUrl + cityIDJson)
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err := json.Unmarshal(body, &res); err != nil {
		return res, err
	}
	return res, nil
}
