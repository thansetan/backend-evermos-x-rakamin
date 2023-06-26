package userdto

import provincecitydto "final_project/internal/pkg/provincecity/dto"

type UserUpdate struct {
	Name        string `json:"nama"`
	Password    string `json:"kata_sandi" validate:"omitempty,min=8"`
	PhoneNumber string `json:"no_telp"`
	DateOfBirth string `json:"tanggal_lahir"`
	Occupation  string `json:"pekerjaan"`
	Email       string `json:"email" validate:"omitempty,email"`
	About       string `json:"about"`
	ProvinceID  string `json:"id_provinsi"`
	CityID      string `json:"id_kota"`
}

type UserResponse struct {
	ID          uint                     `json:"user_id"`
	Name        string                   `json:"nama"`
	PhoneNumber string                   `json:"no_telp"`
	DateOfBirth string                   `json:"tanggal_lahir"`
	Sex         string                   `json:"jenis_kelamin"`
	About       string                   `json:"tentang,omitempty"`
	Occupation  string                   `json:"pekerjaan,omitempty"`
	Email       string                   `json:"email"`
	Province    provincecitydto.Province `json:"provinsi"`
	City        provincecitydto.City     `json:"kota"`
}
