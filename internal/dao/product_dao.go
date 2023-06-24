package dao

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	StoreID       uint            `json:"id_toko"`
	CategoryID    uint            `json:"id_kategori"`
	ProductName   string          `json:"nama_produk"`
	Slug          string          `json:"slug"`
	ResellerPrice string          `json:"harga_reseller"`
	ConsumerPrice string          `json:"harga_konsumen"`
	Stock         int             `json:"stok"`
	Description   string          `json:"deskripsi" gorm:"text"`
	ProductPhotos []*ProductPhoto `gorm:"foreignKey:ProductID"`
	Product       *ProductLog
}

type ProductLog struct {
	gorm.Model
	StoreID       uint            `json:"id_toko"`
	CategoryID    uint            `json:"id_kategori"`
	ProductID     uint            `json:"id_produk"`
	ProductName   string          `json:"nama_produk"`
	Slug          string          `json:"slug"`
	ResellerPrice string          `json:"harga_reseller"`
	ConsumerPrice string          `json:"harga_konsumen"`
	Description   string          `json:"deskripsi" gorm:"text"`
	ProductPhotos []*ProductPhoto `gorm:"foreignKey:ProductID"`
}

type ProductPhoto struct {
	gorm.Model
	ProductID uint   `json:"id_produk"`
	Url       string `json:"url"`
}
