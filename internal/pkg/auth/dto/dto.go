package authdto

import (
	provincecitydto "final_project/internal/pkg/provincecity/dto"
	"time"
)

type Register struct {
	Name        string `json:"nama" validate:"required"`
	Password    string `json:"kata_sandi" validate:"required,min=8"`
	PhoneNumber string `json:"no_telp" validate:"required"`
	DateOfBirth string `json:"tanggal_lahir" validate:"required"`
	Occupation  string `json:"pekerjaan"`
	Email       string `json:"email" validate:"email,required"`
	Sex         string `json:"jenis_kelamin" validate:"required"`
	ProvinceID  string `json:"id_provinsi" validate:"required"`
	CityID      string `json:"id_kota" validate:"required"`
}

type Login struct {
	PhoneNumber string `json:"no_telp" validate:"required"`
	Password    string `json:"kata_sandi" validate:"required"`
}

type LoginResponse struct {
	Name        string                   `json:"nama"`
	PhoneNumber string                   `json:"no_telp"`
	DateOfBirth time.Time                `json:"tanggal_lahir"`
	Sex         string                   `json:"jenis_kelamin"`
	About       string                   `json:"tentang"`
	Occupation  string                   `json:"pekerjaan"`
	Email       string                   `json:"email"`
	Province    provincecitydto.Province `json:"provinsi"`
	City        provincecitydto.City     `json:"kota"`
	Token       string                   `json:"token"`
}
