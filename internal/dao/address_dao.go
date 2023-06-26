package dao

import "gorm.io/gorm"

type Address struct {
	gorm.Model
	UserID        uint         `json:"id_user" gorm:"not null"`
	AddressTitle  string       `json:"judul_alamat" gorm:"not null"`
	Recipient     string       `json:"nama_penerima" gorm:"not null"`
	PhoneNumber   string       `json:"no_telp" gorm:"not null"`
	AddressDetail string       `json:"detail_alamat" gorm:" not null"`
	Transaction   *Transaction `gorm:"foreignKey:AddressID"`
}

type AddressFilter struct {
	UserID, AddressTitle string
}
