package dao

import "gorm.io/gorm"

type Store struct {
	gorm.Model
	UserID             uint                 `json:"id_user"`
	StoreName          string               `json:"nama_toko"`
	PhotoUrl           string               `json:"url_foto"`
	Products           []*Product           `gorm:"foreignKey:StoreID"`
	ProductLogs        []*ProductLog        `gorm:"foreignKey:StoreID"`
	TransactionDetails []*TransactionDetail `gorm:"foreignKey:StoreID"`
}

type StoreFilter struct {
	Limit, Offset int
}
