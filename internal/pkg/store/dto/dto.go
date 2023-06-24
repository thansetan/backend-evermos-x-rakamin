package storedto

type StoreFilter struct {
	Limit int `query:"limit"`
	Page  int `query:"page"`
}

type StoreUpdate struct {
	StoreName string `form:"nama_toko"`
	PhotoUrl  string `form:"photo"`
}

type StoreResponse struct {
	ID        uint   `json:"id"`
	StoreName string `json:"nama_toko"`
	PhotoUrl  string `json:"url_foto"`
}
