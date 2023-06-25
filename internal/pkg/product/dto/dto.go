package productdto

import (
	categorydto "final_project/internal/pkg/category/dto"
	storedto "final_project/internal/pkg/store/dto"
)

type ProductCreateOrUpdate struct {
	StoreID       uint
	ProductName   string `form:"nama_produk"`
	CategoryID    string `form:"category_id"`
	ResellerPrice string `form:"harga_reseller"`
	ConsumerPrice string `form:"harga_konsumen"`
	Stock         string `form:"stok"`
	Description   string `form:"deskripsi"`
}

type ProductResponse struct {
	ProductID     uint                         `json:"id"`
	ProductName   string                       `json:"nama_produk"`
	ResellerPrice uint                         `json:"harga_reseller"`
	ConsumerPrice uint                         `json:"harga_konsumen"`
	Stock         int                          `json:"stok"`
	Description   string                       `json:"deskripsi"`
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
	ProductName string `query:"nama_produk"`
	Limit       int    `query:"limit"`
	Page        int    `query:"page"`
	CategoryID  uint   `query:"category_id"`
	StoreID     uint   `query:"toko_id"`
	MaxPrice    int    `query:"max_harga"`
	MinPrice    uint   `query:"min_harga"`
}
