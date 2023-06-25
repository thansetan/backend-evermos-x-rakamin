package provincecitydto

type Province struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type City struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	ProvinceID string `json:"province_id"`
}
