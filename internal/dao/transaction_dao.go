package dao

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	UserID            uint               `json:"id_user"`
	AddressID         uint               `json:"alamat_pengiriman"`
	TotalPrice        uint               `json:"harga_total"`
	InvoiceNumber     string             `json:"kode_invocie"`
	PaymentMethod     string             `json:"method_bayar"`
	TransactionDetail *TransactionDetail `json:"detail_transaksi" gorm:"foreignKey:TrxID"`
}
type TransactionDetail struct {
	gorm.Model
	TrxID        uint `json:"id_trx"`
	LogProductID uint `json:"id_log_product"`
	StoreID      uint `json:"id_toko"`
	Quantity     int  `json:"kuantitas"`
	TotalPrice   int  `json:"harga_total"`
}
