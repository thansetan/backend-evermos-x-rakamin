package transactiondto

import (
	addressdto "final_project/internal/pkg/address/dto"
	productdto "final_project/internal/pkg/product/dto"
	storedto "final_project/internal/pkg/store/dto"
)

type TransactionResponse struct {
	ID                 uint                        `json:"id"`
	TotalPrice         uint                        `json:"harga_total"`
	InvoiceNumber      string                      `json:"kode_invoice"`
	PaymentMethod      string                      `json:"metode_pembayaran"`
	Address            addressdto.AddressResponse  `json:"alamat_kirim"`
	TransactionDetails []TransactionDetailResponse `json:"detail_transaksi"`
}

type TransactionDetailResponse struct {
	Product    productdto.ProductResponse `json:"product"`
	Store      storedto.StoreResponse     `json:"toko"`
	Quantity   uint                       `json:"kuantitas"`
	TotalPrice uint                       `json:"harga_total"`
}

type TransactionCreate struct {
	UserID             string
	PaymentMethod      string                    `json:"method_bayar"`
	AddressID          uint                      `json:"alamat_kirim"`
	TransactionDetails []TransactionDetailCreate `json:"detail_trx"`
}

type TransactionDetailCreate struct {
	ProductID uint `json:"product_id"`
	Quantity  uint `json:"kuantitas"`
}
