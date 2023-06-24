package dao

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name        string       `json:"nama"`
	Password    string       `json:"kata_sandi"`
	PhoneNumber string       `json:"no_telp" gorm:"unique"`
	DateOfBirth time.Time    `json:"tanggal_lahir"`
	Sex         string       `json:"jenis_kelamin"`
	About       string       `json:"tentang" gorm:"text"`
	Occupation  string       `json:"pekerjaan"`
	Email       string       `json:"email" gorm:"unique"`
	ProvinceID  string       `json:"id_provinsi"`
	CityID      string       `json:"id_kota"`
	IsAdmin     bool         `gorm:"default: false"`
	Addresses   []*Address   `gorm:"foreignKey:UserID"`
	Store       *Store       `gorm:"foreignKey:UserID"`
	Transaction *Transaction `gorm:"foreignKey:UserID"`
}

type UserLogin struct {
	PhoneNumber string `json:"no_telp"`
	Password    string `json:"password"`
}
