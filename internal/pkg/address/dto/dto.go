package addressdto

type AddressCreate struct {
	UserID        uint
	AddressTitle  string `json:"judul_alamat"`
	Recipient     string `json:"nama_penerima"`
	PhoneNumber   string `json:"no_telp"`
	AddressDetail string `json:"detail_alamat"`
}

type AddressUpdate struct {
	UserID        uint
	Recipient     string `json:"nama_penerima"`
	PhoneNumber   string `json:"no_telp"`
	AddressDetail string `json:"detail_alamat"`
}

type AddressResponse struct {
	ID            uint   `json:"id"`
	AddressTitle  string `json:"judul_alamat"`
	Recipient     string `json:"nama_penerima"`
	PhoneNumber   string `json:"no_telp"`
	AddressDetail string `json:"detail_alamat"`
}

type AddressFilter struct {
	UserID       string
	AddressTitle string `query:"judul_alamat"`
}
