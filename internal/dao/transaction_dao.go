package dao

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	UserID             uint                 `json:"id_user" gorm:"not null"`
	AddressID          uint                 `json:"alamat_pengiriman" gorm:"not null"`
	TotalPrice         uint                 `json:"harga_total" gorm:"not null"`
	InvoiceNumber      string               `json:"kode_invoice" gorm:"not null" `
	PaymentMethod      string               `json:"method_bayar" gorm:"not null"`
	TransactionDetails []*TransactionDetail `json:"detail_transaksi" gorm:"foreignKey:TrxID"`

	Address Address
}
type TransactionDetail struct {
	gorm.Model
	TrxID        uint `json:"id_trx" gorm:"not null"`
	ProductLogID uint `json:"id_log_product" gorm:"not null"`
	StoreID      uint `json:"id_toko" gorm:"not null"`
	Quantity     uint `json:"kuantitas" gorm:"not null"`
	TotalPrice   uint `json:"harga_total" gorm:"not null"`

	Store      Store
	ProductLog ProductLog
}
