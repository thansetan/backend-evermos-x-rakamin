package categorydto

type CategoryCreateOrUpdate struct {
	CategoryName string `json:"nama_category" validate:"required"`
}

type CategoryResponse struct {
	ID           uint   `json:"id"`
	CategoryName string `json:"nama_category"`
}

type CategoryFilter struct {
	Name   string `query:"nama"`
	Limit  int    `query:"limit"`
	Offset int    `query:"page"`
}
