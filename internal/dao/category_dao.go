package dao

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	CategoryName string `json:"nama_kategori"`
	Product      *Product
}
