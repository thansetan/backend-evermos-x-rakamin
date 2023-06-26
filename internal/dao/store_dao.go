package dao

import "gorm.io/gorm"

type Store struct {
	gorm.Model
	UserID             uint                 `json:"id_user" gorm:"not null"`
	StoreName          string               `json:"nama_toko" gorm:"not null"`
	PhotoUrl           string               `json:"url_foto"`
	Products           []*Product           `gorm:"foreignKey:StoreID"`
	ProductLogs        []*ProductLog        `gorm:"foreignKey:StoreID"`
	TransactionDetails []*TransactionDetail `gorm:"foreignKey:StoreID"`
}

type StoreFilter struct {
	Limit, Offset int
	Name          string
}
