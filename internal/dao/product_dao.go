package dao

import (
	"errors"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	StoreID       uint            `json:"id_toko"`
	CategoryID    uint            `json:"id_kategori"`
	ProductName   string          `json:"nama_produk"`
	Slug          string          `json:"slug"`
	ResellerPrice uint            `json:"harga_reseller"`
	ConsumerPrice uint            `json:"harga_konsumen"`
	Stock         int             `json:"stok"`
	Description   string          `json:"deskripsi" gorm:"text"`
	ProductPhotos []*ProductPhoto `gorm:"foreignKey:ProductID"`
	ProductLogs   []*ProductLog   `gorm:"foreignKey:ProductID"`
	Store         Store
	Category      Category
}

type ProductLog struct {
	gorm.Model
	StoreID           uint               `json:"id_toko"`
	CategoryID        uint               `json:"id_kategori"`
	ProductID         uint               `json:"id_produk"`
	ProductName       string             `json:"nama_produk"`
	Slug              string             `json:"slug"`
	ResellerPrice     uint               `json:"harga_reseller"`
	ConsumerPrice     uint               `json:"harga_konsumen"`
	Description       string             `json:"deskripsi" gorm:"text"`
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
	ProductName                   string
	Limit, Page, MaxPrice         int
	CategoryID, StoreID, MinPrice uint
}

func (p *Product) AfterSave(tx *gorm.DB) error {
	if p.Stock < 0 {
		return errors.New("invalid product stock")
	}
	return nil
}
