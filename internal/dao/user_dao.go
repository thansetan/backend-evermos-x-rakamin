package dao

import (
	"errors"
	"fmt"
	"time"

	"net/http"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name        string         `json:"nama"`
	Password    string         `json:"kata_sandi"`
	PhoneNumber string         `json:"no_telp" gorm:"unique"`
	DateOfBirth time.Time      `json:"tanggal_lahir"`
	Sex         string         `json:"jenis_kelamin"`
	About       string         `json:"tentang" gorm:"text"`
	Occupation  string         `json:"pekerjaan"`
	Email       string         `json:"email" gorm:"unique"`
	ProvinceID  string         `json:"id_provinsi"`
	CityID      string         `json:"id_kota"`
	IsAdmin     bool           `gorm:"default: false"`
	Addresses   []*Address     `gorm:"foreignKey:UserID"`
	Store       *Store         `gorm:"foreignKey:UserID"`
	Transaction []*Transaction `gorm:"foreignKey:UserID"`
}

type UserLogin struct {
	PhoneNumber string `json:"no_telp"`
	Password    string `json:"password"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	province, _ := http.Get(fmt.Sprintf("https://www.emsifa.com/api-wilayah-indonesia/api/province/%s.json", u.ProvinceID))
	city, _ := http.Get(fmt.Sprintf("https://www.emsifa.com/api-wilayah-indonesia/api/regency/%s.json", u.CityID))
	if city.StatusCode == 404 || province.StatusCode == 404 {
		return errors.New("invalid city/province ID")
	}
	return nil
}

func (u *User) AfterUpdate(tx *gorm.DB) error {
	city, _ := http.Get(fmt.Sprintf("https://www.emsifa.com/api-wilayah-indonesia/api/regency/%s.json", u.CityID))
	province, _ := http.Get(fmt.Sprintf("https://www.emsifa.com/api-wilayah-indonesia/api/province/%s.json", u.ProvinceID))
	if city.StatusCode == 404 || province.StatusCode == 404 {
		return errors.New("invalid city/province ID")
	}
	return nil
}
