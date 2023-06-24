package dao

import "gorm.io/gorm"

type Address struct {
	gorm.Model
	UserID        uint   `json:"id_user"`
	AddressTitle  string `json:"judul_alamat"`
	Recipient     string `json:"nama_penerima"`
	PhoneNumber   string `json:"no_telp"`
	AddressDetail string `json:"detail_alamat"`
}

type AddressFilter struct {
	UserID, AddressTitle string
}
