package authdto

import (
	provincecitydto "final_project/internal/pkg/provincecity/dto"
	"time"
)

type Register struct {
	Name        string `json:"nama"`
	Password    string `json:"kata_sandi"`
	PhoneNumber string `json:"no_telp"`
	DateOfBirth string `json:"tanggal_lahir"`
	Occupation  string `json:"pekerjaan"`
	Email       string `json:"email"`
	ProvinceID  string `json:"id_provinsi"`
	Sex         string `json:"jenis_kelamin"`
	CityID      string `json:"id_kota"`
}

type Login struct {
	PhoneNumber string `json:"no_telp"`
	Password    string `json:"kata_sandi"`
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
