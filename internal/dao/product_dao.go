package dao

import (
	"errors"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	StoreID       uint            `json:"id_toko" gorm:"not null"`
	CategoryID    uint            `json:"id_kategori" gorm:"not null"`
	ProductName   string          `json:"nama_produk" gorm:"not null"`
	Slug          string          `json:"slug" gorm:"not null"`
	ResellerPrice uint            `json:"harga_reseller"`
	ConsumerPrice uint            `json:"harga_konsumen" gorm:"not null"`
	Stock         int             `json:"stok" gorm:"not null"`
	Description   string          `json:"deskripsi" gorm:"type:text"`
	ProductPhotos []*ProductPhoto `gorm:"foreignKey:ProductID"`
	ProductLogs   []*ProductLog   `gorm:"foreignKey:ProductID"`
	Store         Store
	Category      Category
}

type ProductLog struct {
	gorm.Model
	StoreID           uint               `json:"id_toko" gorm:"not null"`
	CategoryID        uint               `json:"id_kategori" gorm:"not null"`
	ProductID         uint               `json:"id_produk" gorm:"not null"`
	ProductName       string             `json:"nama_produk" gorm:"not null"`
	Slug              string             `json:"slug" gorm:"not null"`
	ResellerPrice     uint               `json:"harga_reseller" gorm:"not null"`
	ConsumerPrice     uint               `json:"harga_konsumen" gorm:"not null"`
	Description       string             `json:"deskripsi" gorm:"type:text"`
	Stock             uint               `json:"stok" gorm:"not null"`
	TransactionDetail *TransactionDetail `gorm:"foreignKey:ProductLogID"`

	Category Category
	Product  Product
}

type ProductPhoto struct {
	gorm.Model
	ProductID uint   `json:"id_produk"`
	Url       string `json:"url"`
}

type ProductFilter struct {
	ProductName                             string
	Limit, Page                             int
	CategoryID, StoreID, MinPrice, MaxPrice uint
}

type ProductLogRes struct {
	ID, ProductID uint
}

func (p *Product) AfterSave(tx *gorm.DB) error {
	if p.Stock < 0 {
		return errors.New("invalid product stock")
	}
	return nil
}
