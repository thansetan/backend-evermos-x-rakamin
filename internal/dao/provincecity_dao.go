package dao

type Province struct {
	ID, Name string
}

type City struct {
	ID, Name   string
	ProvinceID string `json:"province_id"`
}
