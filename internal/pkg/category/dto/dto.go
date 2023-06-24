package categorydto

type CategoryCreateOrUpdate struct {
	CategoryName string `json:"nama_category"`
}

type CategoryResponse struct {
	ID           uint   `json:"id"`
	CategoryName string `json:"nama_category"`
}
