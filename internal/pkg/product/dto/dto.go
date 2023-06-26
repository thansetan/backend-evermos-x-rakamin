package productdto

import (
	categorydto "final_project/internal/pkg/category/dto"
	storedto "final_project/internal/pkg/store/dto"
)

type ProductCreate struct {
	StoreID       uint
	ProductName   string  `form:"nama_produk" validate:"required"`
	CategoryID    string  `form:"category_id" validate:"required"`
	ResellerPrice float64 `form:"harga_reseller" validate:"omitempty,gte=0"`
	ConsumerPrice float64 `form:"harga_konsumen" validate:"required,gte=0"`
	Stock         string  `form:"stok" validate:"required,gte=0"`
	Description   string  `form:"deskripsi"`
}

type ProductUpdate struct {
	StoreID       uint
	ProductName   string  `form:"nama_produk"`
	CategoryID    string  `form:"category_id"`
	ResellerPrice float64 `form:"harga_reseller" validate:"omitempty,gte=0"`
	ConsumerPrice float64 `form:"harga_konsumen" validate:"omitempty,gte=0"`
	Stock         string  `form:"stok" validate:"gte=0"`
	Description   string  `form:"deskripsi"`
}

type ProductResponse struct {
	ProductID     uint                         `json:"id"`
	ProductName   string                       `json:"nama_produk"`
	ResellerPrice uint                         `json:"harga_reseller,omitempty"`
	ConsumerPrice uint                         `json:"harga_konsumen"`
	Stock         int                          `json:"stok"`
	Description   string                       `json:"deskripsi,omitempty"`
	Store         storedto.StoreResponse       `json:"toko"`
	Category      categorydto.CategoryResponse `json:"category"`
	Photos        []*ProductPhotoResponse      `json:"photos"`
}

type ProductPhotoResponse struct {
	PhotoID   uint   `json:"id"`
	ProductID uint   `json:"id_produk"`
	Url       string `json:"url"`
}

type ProductFilter struct {
	ProductName string  `query:"nama_produk"`
	Limit       int     `query:"limit"`
	Page        int     `query:"page"`
	CategoryID  uint    `query:"category_id"`
	StoreID     uint    `query:"toko_id"`
	MinPrice    float64 `query:"min_harga" validate:"omitempty,gte=0"`
	MaxPrice    float64 `query:"max_harga" validate:"omitempty,gtefield=MinPrice"`
}
